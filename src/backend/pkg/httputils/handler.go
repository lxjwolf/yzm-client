package httputils

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"fmt"
)

type APPHandler func(w http.ResponseWriter, r *http.Request) HTTPError

func (fn APPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err HTTPError) {

	for key, vals := range err.Headers() {
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}

	if err.StatusCode() >= http.StatusInternalServerError {
		log.Warnf("出了一个大错：%+v, InsideError:%+v", err, err.InsideError())
	}

	// http.Error(w, err.Error(), err.StatusCode())
	w.WriteHeader(err.StatusCode())
	fmt.Fprintln(w, err.Error())
}
