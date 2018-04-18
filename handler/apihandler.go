package handler

import (
	"net/http"
	"github.com/zm-dev/yzm-client/pkg/httputils"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/zm-dev/yzm-client/distinguish"
	"io/ioutil"
)

func CreateHTTPAPIHandler() (http.Handler) {
	r := mux.NewRouter()
	r.Handle("/batch_upload", httputils.APPHandler(batchUpload))
	r.Handle("/upload", httputils.APPHandler(upload))
	return r
}

func batchUpload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	batchImageFile, batchImageFileHeader, err := r.FormFile("batch_image")
	defer batchImageFile.Close()

	if err != nil {
		return httputils.InternalServerError("图片压缩包上传失败！").WithError(err)
	}

	mappings, err := distinguish.BatchProcess(category, batchImageFile, batchImageFileHeader.Size, distinguish.BatchDistinguish)

	if err != nil {
		return httputils.InternalServerError("图片压缩包处理失败！").WithError(err)
	}

	b, _ := ioutil.ReadAll(mappings)
	w.Write(b)
	return nil
}

func upload(w http.ResponseWriter, r *http.Request) httputils.HTTPError {

	category, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		return httputils.BadRequest("category 错误！").WithError(err)
	}

	imageFile, _, err := r.FormFile("image")
	defer imageFile.Close()

	yzmStr, err := distinguish.Process(category, imageFile)
	if err != nil {
		return httputils.InternalServerError("图片识别出错!").WithError(err)
	}
	w.Write([]byte(yzmStr))
	return nil

}
