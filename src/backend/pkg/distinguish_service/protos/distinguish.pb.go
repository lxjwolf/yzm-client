// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distinguish.proto

/*
Package distinguish is a generated protocol buffer package.

It is generated from these files:
	distinguish.proto

It has these top-level messages:
	Image
	Yzm
*/
package distinguish

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Image_Category int32

const (
	Image_DATA1 Image_Category = 0
	Image_DATA2 Image_Category = 1
	Image_DATA3 Image_Category = 2
	Image_DATA4 Image_Category = 3
	Image_DATA5 Image_Category = 4
)

var Image_Category_name = map[int32]string{
	0: "DATA1",
	1: "DATA2",
	2: "DATA3",
	3: "DATA4",
	4: "DATA5",
}
var Image_Category_value = map[string]int32{
	"DATA1": 0,
	"DATA2": 1,
	"DATA3": 2,
	"DATA4": 3,
	"DATA5": 4,
}

func (x Image_Category) String() string {
	return proto.EnumName(Image_Category_name, int32(x))
}
func (Image_Category) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Image struct {
	Category Image_Category `protobuf:"varint,1,opt,name=category,enum=distinguish.Image_Category" json:"category,omitempty"`
	Data     []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Image) GetCategory() Image_Category {
	if m != nil {
		return m.Category
	}
	return Image_DATA1
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Yzm struct {
	Yzm string `protobuf:"bytes,1,opt,name=yzm" json:"yzm,omitempty"`
}

func (m *Yzm) Reset()                    { *m = Yzm{} }
func (m *Yzm) String() string            { return proto.CompactTextString(m) }
func (*Yzm) ProtoMessage()               {}
func (*Yzm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Yzm) GetYzm() string {
	if m != nil {
		return m.Yzm
	}
	return ""
}

func init() {
	proto.RegisterType((*Image)(nil), "distinguish.Image")
	proto.RegisterType((*Yzm)(nil), "distinguish.Yzm")
	proto.RegisterEnum("distinguish.Image_Category", Image_Category_name, Image_Category_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Distinguish service

type DistinguishClient interface {
	Distinguish(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Yzm, error)
}

type distinguishClient struct {
	cc *grpc.ClientConn
}

func NewDistinguishClient(cc *grpc.ClientConn) DistinguishClient {
	return &distinguishClient{cc}
}

func (c *distinguishClient) Distinguish(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Yzm, error) {
	out := new(Yzm)
	err := grpc.Invoke(ctx, "/distinguish.Distinguish/Distinguish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Distinguish service

type DistinguishServer interface {
	Distinguish(context.Context, *Image) (*Yzm, error)
}

func RegisterDistinguishServer(s *grpc.Server, srv DistinguishServer) {
	s.RegisterService(&_Distinguish_serviceDesc, srv)
}

func _Distinguish_Distinguish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistinguishServer).Distinguish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/distinguish.Distinguish/Distinguish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistinguishServer).Distinguish(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

var _Distinguish_serviceDesc = grpc.ServiceDesc{
	ServiceName: "distinguish.Distinguish",
	HandlerType: (*DistinguishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Distinguish",
			Handler:    _Distinguish_Distinguish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "distinguish.proto",
}

func init() { proto.RegisterFile("distinguish.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xc9, 0x2c, 0x2e,
	0xc9, 0xcc, 0x4b, 0x2f, 0xcd, 0x2c, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x9a, 0xce, 0xc8, 0xc5, 0xea, 0x99, 0x9b, 0x98, 0x9e, 0x2a, 0x64, 0xce, 0xc5, 0x91,
	0x9c, 0x58, 0x92, 0x9a, 0x9e, 0x5f, 0x54, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xad,
	0x87, 0xac, 0x19, 0xac, 0x4a, 0xcf, 0x19, 0xaa, 0x24, 0x08, 0xae, 0x58, 0x48, 0x88, 0x8b, 0x25,
	0x25, 0xb1, 0x24, 0x51, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x56, 0x72, 0xe4, 0xe2,
	0x80, 0xa9, 0x14, 0xe2, 0xe4, 0x62, 0x75, 0x71, 0x0c, 0x71, 0x34, 0x14, 0x60, 0x80, 0x31, 0x8d,
	0x04, 0x18, 0x61, 0x4c, 0x63, 0x01, 0x26, 0x18, 0xd3, 0x44, 0x80, 0x19, 0xc6, 0x34, 0x15, 0x60,
	0x51, 0x12, 0xe7, 0x62, 0x8e, 0xac, 0xca, 0x15, 0x12, 0xe0, 0x62, 0xae, 0xac, 0xca, 0x05, 0xbb,
	0x88, 0x33, 0x08, 0xc4, 0x34, 0x72, 0xe1, 0xe2, 0x76, 0x41, 0xb8, 0x4b, 0xc8, 0x14, 0x95, 0x2b,
	0x84, 0xe9, 0x68, 0x29, 0x01, 0x14, 0xb1, 0xc8, 0xaa, 0x5c, 0x25, 0x86, 0x24, 0x36, 0x70, 0x60,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x7f, 0xe0, 0xdf, 0x21, 0x01, 0x00, 0x00,
}
