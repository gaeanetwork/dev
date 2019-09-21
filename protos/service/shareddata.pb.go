// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/gaeanetwork/gaea-core/protos/service/shareddata.proto

package service

import (
	context "context"
	fmt "fmt"
	tee "github.com/gaeanetwork/gaea-core/protos/tee"
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

type UploadRequest struct {
	Data                 string          `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Hash                 string          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Description          string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Owner                string          `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Signatures           *SignatureGroup `protobuf:"bytes,5,opt,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c1df5d6268a97b, []int{0}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *UploadRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *UploadRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UploadRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UploadRequest) GetSignatures() *SignatureGroup {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type UploadResponse struct {
	Data                 *tee.SharedData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c1df5d6268a97b, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetData() *tee.SharedData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SignatureGroup struct {
	Signature            []string `protobuf:"bytes,1,rep,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureGroup) Reset()         { *m = SignatureGroup{} }
func (m *SignatureGroup) String() string { return proto.CompactTextString(m) }
func (*SignatureGroup) ProtoMessage()    {}
func (*SignatureGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_94c1df5d6268a97b, []int{2}
}

func (m *SignatureGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureGroup.Unmarshal(m, b)
}
func (m *SignatureGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureGroup.Marshal(b, m, deterministic)
}
func (m *SignatureGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureGroup.Merge(m, src)
}
func (m *SignatureGroup) XXX_Size() int {
	return xxx_messageInfo_SignatureGroup.Size(m)
}
func (m *SignatureGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureGroup proto.InternalMessageInfo

func (m *SignatureGroup) GetSignature() []string {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadRequest)(nil), "service.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "service.UploadResponse")
	proto.RegisterType((*SignatureGroup)(nil), "service.SignatureGroup")
}

func init() {
	proto.RegisterFile("github.com/gaeanetwork/gaea-core/protos/service/shareddata.proto", fileDescriptor_94c1df5d6268a97b)
}

var fileDescriptor_94c1df5d6268a97b = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xfd, 0x23, 0x99, 0x60, 0x85, 0x45, 0xec, 0x52, 0x3c, 0x84, 0x78, 0xc9, 0xc5,
	0x0d, 0x46, 0x44, 0x04, 0x0f, 0x22, 0x42, 0xef, 0x29, 0x5e, 0xbc, 0x6d, 0x93, 0x21, 0x09, 0x6a,
	0x36, 0xee, 0x4e, 0xec, 0x57, 0xf2, 0x63, 0x4a, 0x36, 0xfd, 0x93, 0xd2, 0x8b, 0xbd, 0xbd, 0x79,
	0xbf, 0xe5, 0xed, 0x9b, 0x5d, 0x78, 0xce, 0x4b, 0x2a, 0x9a, 0xa5, 0x48, 0xd5, 0x57, 0x94, 0x4b,
	0x94, 0x15, 0xd2, 0x4a, 0xe9, 0x0f, 0xab, 0x6f, 0x52, 0xa5, 0x31, 0xaa, 0xb5, 0x22, 0x65, 0x22,
	0x83, 0xfa, 0xa7, 0x4c, 0x31, 0x32, 0x85, 0xd4, 0x98, 0x65, 0x92, 0xa4, 0xb0, 0x84, 0x9d, 0xae,
	0xc9, 0xec, 0xe9, 0xbf, 0x51, 0x84, 0x87, 0x31, 0xc1, 0xaf, 0x03, 0x67, 0x6f, 0xf5, 0xa7, 0x92,
	0x59, 0x82, 0xdf, 0x0d, 0x1a, 0x62, 0x0c, 0x86, 0x2d, 0xe7, 0x8e, 0xef, 0x84, 0x6e, 0x62, 0x75,
	0xeb, 0x15, 0xd2, 0x14, 0xfc, 0xa4, 0xf3, 0x5a, 0xcd, 0x7c, 0xf0, 0x32, 0x34, 0xa9, 0x2e, 0x6b,
	0x2a, 0x55, 0xc5, 0x07, 0x16, 0xf5, 0x2d, 0x76, 0x01, 0x23, 0xb5, 0xaa, 0x50, 0xf3, 0xa1, 0x65,
	0xdd, 0xc0, 0x1e, 0x00, 0x4c, 0x99, 0x57, 0x92, 0x1a, 0x8d, 0x86, 0x8f, 0x7c, 0x27, 0xf4, 0xe2,
	0xa9, 0x58, 0x6f, 0x23, 0x16, 0x1b, 0x34, 0xd7, 0xaa, 0xa9, 0x93, 0xde, 0xd1, 0xe0, 0x1e, 0x26,
	0x9b, 0xa6, 0xa6, 0x56, 0x95, 0x41, 0x76, 0xdd, 0xab, 0xea, 0xc5, 0xe7, 0x82, 0x10, 0xc5, 0xc2,
	0x6e, 0xf8, 0x2a, 0x49, 0x76, 0xdd, 0x03, 0x01, 0x93, 0xfd, 0x50, 0x76, 0x05, 0xee, 0x36, 0x96,
	0x3b, 0xfe, 0x20, 0x74, 0x93, 0x9d, 0x11, 0xcf, 0x01, 0x76, 0x19, 0xec, 0x11, 0xc6, 0xdd, 0xa5,
	0xec, 0x72, 0xdb, 0x71, 0xef, 0xbd, 0x66, 0xd3, 0x03, 0xbf, 0x6b, 0xf7, 0x72, 0xfb, 0x1e, 0x1d,
	0xf9, 0xcb, 0xcb, 0xb1, 0x9d, 0xef, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x8e, 0x4c, 0xb5,
	0x1f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SharedDataClient is the client API for SharedData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SharedDataClient interface {
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type sharedDataClient struct {
	cc *grpc.ClientConn
}

func NewSharedDataClient(cc *grpc.ClientConn) SharedDataClient {
	return &sharedDataClient{cc}
}

func (c *sharedDataClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/service.SharedData/Upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedDataServer is the server API for SharedData service.
type SharedDataServer interface {
	Upload(context.Context, *UploadRequest) (*UploadResponse, error)
}

// UnimplementedSharedDataServer can be embedded to have forward compatible implementations.
type UnimplementedSharedDataServer struct {
}

func (*UnimplementedSharedDataServer) Upload(ctx context.Context, req *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterSharedDataServer(s *grpc.Server, srv SharedDataServer) {
	s.RegisterService(&_SharedData_serviceDesc, srv)
}

func _SharedData_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedDataServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SharedData/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedDataServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SharedData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.SharedData",
	HandlerType: (*SharedDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _SharedData_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/gaeanetwork/gaea-core/protos/service/shareddata.proto",
}
