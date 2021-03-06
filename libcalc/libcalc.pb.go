// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: libcalc/libcalc.proto

package libcalc

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Libcalc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addend1 int32 `protobuf:"varint,1,opt,name=addend1,proto3" json:"addend1,omitempty"`
	Addend2 int32 `protobuf:"varint,2,opt,name=addend2,proto3" json:"addend2,omitempty"`
}

func (x *Libcalc) Reset() {
	*x = Libcalc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcalc_libcalc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Libcalc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Libcalc) ProtoMessage() {}

func (x *Libcalc) ProtoReflect() protoreflect.Message {
	mi := &file_libcalc_libcalc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Libcalc.ProtoReflect.Descriptor instead.
func (*Libcalc) Descriptor() ([]byte, []int) {
	return file_libcalc_libcalc_proto_rawDescGZIP(), []int{0}
}

func (x *Libcalc) GetAddend1() int32 {
	if x != nil {
		return x.Addend1
	}
	return 0
}

func (x *Libcalc) GetAddend2() int32 {
	if x != nil {
		return x.Addend2
	}
	return 0
}

type LibcalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *Libcalc `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *LibcalcRequest) Reset() {
	*x = LibcalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcalc_libcalc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibcalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibcalcRequest) ProtoMessage() {}

func (x *LibcalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libcalc_libcalc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibcalcRequest.ProtoReflect.Descriptor instead.
func (*LibcalcRequest) Descriptor() ([]byte, []int) {
	return file_libcalc_libcalc_proto_rawDescGZIP(), []int{1}
}

func (x *LibcalcRequest) GetRequest() *Libcalc {
	if x != nil {
		return x.Request
	}
	return nil
}

type LibcalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LibcalcResponse) Reset() {
	*x = LibcalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcalc_libcalc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibcalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibcalcResponse) ProtoMessage() {}

func (x *LibcalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libcalc_libcalc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibcalcResponse.ProtoReflect.Descriptor instead.
func (*LibcalcResponse) Descriptor() ([]byte, []int) {
	return file_libcalc_libcalc_proto_rawDescGZIP(), []int{2}
}

func (x *LibcalcResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type LibcalcFactorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nint int32 `protobuf:"varint,1,opt,name=nint,proto3" json:"nint,omitempty"`
}

func (x *LibcalcFactorsRequest) Reset() {
	*x = LibcalcFactorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcalc_libcalc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibcalcFactorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibcalcFactorsRequest) ProtoMessage() {}

func (x *LibcalcFactorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_libcalc_libcalc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibcalcFactorsRequest.ProtoReflect.Descriptor instead.
func (*LibcalcFactorsRequest) Descriptor() ([]byte, []int) {
	return file_libcalc_libcalc_proto_rawDescGZIP(), []int{3}
}

func (x *LibcalcFactorsRequest) GetNint() int32 {
	if x != nil {
		return x.Nint
	}
	return 0
}

type LibcalcFactorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LibcalcFactorsResponse) Reset() {
	*x = LibcalcFactorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libcalc_libcalc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LibcalcFactorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LibcalcFactorsResponse) ProtoMessage() {}

func (x *LibcalcFactorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_libcalc_libcalc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LibcalcFactorsResponse.ProtoReflect.Descriptor instead.
func (*LibcalcFactorsResponse) Descriptor() ([]byte, []int) {
	return file_libcalc_libcalc_proto_rawDescGZIP(), []int{4}
}

func (x *LibcalcFactorsResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_libcalc_libcalc_proto protoreflect.FileDescriptor

var file_libcalc_libcalc_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63,
	0x22, 0x3d, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x65, 0x6e, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x65, 0x6e, 0x64, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x32, 0x22,
	0x3c, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4c, 0x69, 0x62,
	0x63, 0x61, 0x6c, 0x63, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a,
	0x0f, 0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2b, 0x0a, 0x15, 0x4c, 0x69, 0x62, 0x63,
	0x61, 0x6c, 0x63, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x6e, 0x69, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x16, 0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9f, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x63,
	0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x4c, 0x69,
	0x62, 0x63, 0x61, 0x6c, 0x63, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2e,
	0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x46, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4c,
	0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4c,
	0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6c,
	0x69, 0x62, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libcalc_libcalc_proto_rawDescOnce sync.Once
	file_libcalc_libcalc_proto_rawDescData = file_libcalc_libcalc_proto_rawDesc
)

func file_libcalc_libcalc_proto_rawDescGZIP() []byte {
	file_libcalc_libcalc_proto_rawDescOnce.Do(func() {
		file_libcalc_libcalc_proto_rawDescData = protoimpl.X.CompressGZIP(file_libcalc_libcalc_proto_rawDescData)
	})
	return file_libcalc_libcalc_proto_rawDescData
}

var file_libcalc_libcalc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_libcalc_libcalc_proto_goTypes = []interface{}{
	(*Libcalc)(nil),                // 0: libcalc.Libcalc
	(*LibcalcRequest)(nil),         // 1: libcalc.LibcalcRequest
	(*LibcalcResponse)(nil),        // 2: libcalc.LibcalcResponse
	(*LibcalcFactorsRequest)(nil),  // 3: libcalc.LibcalcFactorsRequest
	(*LibcalcFactorsResponse)(nil), // 4: libcalc.LibcalcFactorsResponse
}
var file_libcalc_libcalc_proto_depIdxs = []int32{
	0, // 0: libcalc.LibcalcRequest.request:type_name -> libcalc.Libcalc
	1, // 1: libcalc.LibcalcService.Libcalc:input_type -> libcalc.LibcalcRequest
	3, // 2: libcalc.LibcalcService.Factor:input_type -> libcalc.LibcalcFactorsRequest
	2, // 3: libcalc.LibcalcService.Libcalc:output_type -> libcalc.LibcalcResponse
	4, // 4: libcalc.LibcalcService.Factor:output_type -> libcalc.LibcalcFactorsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_libcalc_libcalc_proto_init() }
func file_libcalc_libcalc_proto_init() {
	if File_libcalc_libcalc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libcalc_libcalc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Libcalc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libcalc_libcalc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibcalcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libcalc_libcalc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibcalcResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libcalc_libcalc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibcalcFactorsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_libcalc_libcalc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LibcalcFactorsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_libcalc_libcalc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_libcalc_libcalc_proto_goTypes,
		DependencyIndexes: file_libcalc_libcalc_proto_depIdxs,
		MessageInfos:      file_libcalc_libcalc_proto_msgTypes,
	}.Build()
	File_libcalc_libcalc_proto = out.File
	file_libcalc_libcalc_proto_rawDesc = nil
	file_libcalc_libcalc_proto_goTypes = nil
	file_libcalc_libcalc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LibcalcServiceClient is the client API for LibcalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LibcalcServiceClient interface {
	// define unary API
	Libcalc(ctx context.Context, in *LibcalcRequest, opts ...grpc.CallOption) (*LibcalcResponse, error)
	Factor(ctx context.Context, in *LibcalcFactorsRequest, opts ...grpc.CallOption) (LibcalcService_FactorClient, error)
}

type libcalcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibcalcServiceClient(cc grpc.ClientConnInterface) LibcalcServiceClient {
	return &libcalcServiceClient{cc}
}

func (c *libcalcServiceClient) Libcalc(ctx context.Context, in *LibcalcRequest, opts ...grpc.CallOption) (*LibcalcResponse, error) {
	out := new(LibcalcResponse)
	err := c.cc.Invoke(ctx, "/libcalc.LibcalcService/Libcalc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libcalcServiceClient) Factor(ctx context.Context, in *LibcalcFactorsRequest, opts ...grpc.CallOption) (LibcalcService_FactorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LibcalcService_serviceDesc.Streams[0], "/libcalc.LibcalcService/Factor", opts...)
	if err != nil {
		return nil, err
	}
	x := &libcalcServiceFactorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LibcalcService_FactorClient interface {
	Recv() (*LibcalcFactorsResponse, error)
	grpc.ClientStream
}

type libcalcServiceFactorClient struct {
	grpc.ClientStream
}

func (x *libcalcServiceFactorClient) Recv() (*LibcalcFactorsResponse, error) {
	m := new(LibcalcFactorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LibcalcServiceServer is the server API for LibcalcService service.
type LibcalcServiceServer interface {
	// define unary API
	Libcalc(context.Context, *LibcalcRequest) (*LibcalcResponse, error)
	Factor(*LibcalcFactorsRequest, LibcalcService_FactorServer) error
}

// UnimplementedLibcalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLibcalcServiceServer struct {
}

func (*UnimplementedLibcalcServiceServer) Libcalc(context.Context, *LibcalcRequest) (*LibcalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Libcalc not implemented")
}
func (*UnimplementedLibcalcServiceServer) Factor(*LibcalcFactorsRequest, LibcalcService_FactorServer) error {
	return status.Errorf(codes.Unimplemented, "method Factor not implemented")
}

func RegisterLibcalcServiceServer(s *grpc.Server, srv LibcalcServiceServer) {
	s.RegisterService(&_LibcalcService_serviceDesc, srv)
}

func _LibcalcService_Libcalc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LibcalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibcalcServiceServer).Libcalc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/libcalc.LibcalcService/Libcalc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibcalcServiceServer).Libcalc(ctx, req.(*LibcalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibcalcService_Factor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LibcalcFactorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LibcalcServiceServer).Factor(m, &libcalcServiceFactorServer{stream})
}

type LibcalcService_FactorServer interface {
	Send(*LibcalcFactorsResponse) error
	grpc.ServerStream
}

type libcalcServiceFactorServer struct {
	grpc.ServerStream
}

func (x *libcalcServiceFactorServer) Send(m *LibcalcFactorsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _LibcalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "libcalc.LibcalcService",
	HandlerType: (*LibcalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Libcalc",
			Handler:    _LibcalcService_Libcalc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Factor",
			Handler:       _LibcalcService_Factor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "libcalc/libcalc.proto",
}
