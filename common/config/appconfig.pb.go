// Code generated by protoc-gen-go. DO NOT EDIT.
// source: appconfig.proto

// 定义包名

package config

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 定义 Res 消息结构
type Params struct {
	Keys                 string   `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6183fdf07ef5608d, []int{0}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Params.Unmarshal(m, b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Params.Marshal(b, m, deterministic)
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return xxx_messageInfo_Params.Size(m)
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeys() string {
	if m != nil {
		return m.Keys
	}
	return ""
}

type Result struct {
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_6183fdf07ef5608d, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "appconfig.Params")
	proto.RegisterType((*Result)(nil), "appconfig.Result")
}

func init() { proto.RegisterFile("appconfig.proto", fileDescriptor_6183fdf07ef5608d) }

var fileDescriptor_6183fdf07ef5608d = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x28, 0x48,
	0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xc9, 0x70, 0xb1, 0x05, 0x24, 0x16, 0x25, 0xe6, 0x16, 0x0b, 0x09, 0x71, 0xb1, 0x64, 0xa7, 0x56,
	0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x06, 0x5c, 0x6c, 0x41, 0xa9,
	0xc5, 0xa5, 0x39, 0x25, 0x20, 0x59, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x30, 0x1b, 0x24, 0xe6, 0x92, 0x58, 0x92, 0x08, 0xd3, 0x01, 0x62, 0x1b, 0x39, 0x73, 0x71,
	0x3a, 0xc2, 0x0c, 0x17, 0x32, 0xe3, 0xe2, 0x71, 0x4f, 0x2d, 0x71, 0x2c, 0x28, 0x70, 0x86, 0xf0,
	0x05, 0xf5, 0x10, 0x2e, 0x81, 0xd8, 0x2a, 0x85, 0x2c, 0x04, 0xb1, 0x4a, 0x89, 0x21, 0x89, 0x0d,
	0xec, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0xef, 0x2a, 0x33, 0xb9, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppconfigClient is the client API for Appconfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppconfigClient interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	GetAppConfig(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Result, error)
}

type appconfigClient struct {
	cc *grpc.ClientConn
}

func NewAppconfigClient(cc *grpc.ClientConn) AppconfigClient {
	return &appconfigClient{cc}
}

func (c *appconfigClient) GetAppConfig(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/appconfig.Appconfig/GetAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppconfigServer is the server API for Appconfig service.
type AppconfigServer interface {
	// 定义接口 (结构体可以复用)
	// 方法 (请求消息结构体) returns (返回消息结构体) {}
	GetAppConfig(context.Context, *Params) (*Result, error)
}

// UnimplementedAppconfigServer can be embedded to have forward compatible implementations.
type UnimplementedAppconfigServer struct {
}

func (*UnimplementedAppconfigServer) GetAppConfig(ctx context.Context, req *Params) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}

func RegisterAppconfigServer(s *grpc.Server, srv AppconfigServer) {
	s.RegisterService(&_Appconfig_serviceDesc, srv)
}

func _Appconfig_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppconfigServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appconfig.Appconfig/GetAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppconfigServer).GetAppConfig(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

var _Appconfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appconfig.Appconfig",
	HandlerType: (*AppconfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppConfig",
			Handler:    _Appconfig_GetAppConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appconfig.proto",
}
