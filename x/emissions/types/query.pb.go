// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e578782beb6ef82, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e578782beb6ef82, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryListPoolAddressesRequest struct {
}

func (m *QueryListPoolAddressesRequest) Reset()         { *m = QueryListPoolAddressesRequest{} }
func (m *QueryListPoolAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListPoolAddressesRequest) ProtoMessage()    {}
func (*QueryListPoolAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e578782beb6ef82, []int{2}
}
func (m *QueryListPoolAddressesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListPoolAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListPoolAddressesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListPoolAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListPoolAddressesRequest.Merge(m, src)
}
func (m *QueryListPoolAddressesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListPoolAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListPoolAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListPoolAddressesRequest proto.InternalMessageInfo

type QueryListPoolAddressesResponse struct {
	UndistributedObserverBalancesAddress string `protobuf:"bytes,1,opt,name=undistributed_observer_balances_address,json=undistributedObserverBalancesAddress,proto3" json:"undistributed_observer_balances_address,omitempty"`
	UndistributedTssBalancesAddress      string `protobuf:"bytes,2,opt,name=undistributed_tss_balances_address,json=undistributedTssBalancesAddress,proto3" json:"undistributed_tss_balances_address,omitempty"`
	EmissionModuleAddress                string `protobuf:"bytes,3,opt,name=emission_module_address,json=emissionModuleAddress,proto3" json:"emission_module_address,omitempty"`
}

func (m *QueryListPoolAddressesResponse) Reset()         { *m = QueryListPoolAddressesResponse{} }
func (m *QueryListPoolAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListPoolAddressesResponse) ProtoMessage()    {}
func (*QueryListPoolAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e578782beb6ef82, []int{3}
}
func (m *QueryListPoolAddressesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListPoolAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListPoolAddressesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListPoolAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListPoolAddressesResponse.Merge(m, src)
}
func (m *QueryListPoolAddressesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListPoolAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListPoolAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListPoolAddressesResponse proto.InternalMessageInfo

func (m *QueryListPoolAddressesResponse) GetUndistributedObserverBalancesAddress() string {
	if m != nil {
		return m.UndistributedObserverBalancesAddress
	}
	return ""
}

func (m *QueryListPoolAddressesResponse) GetUndistributedTssBalancesAddress() string {
	if m != nil {
		return m.UndistributedTssBalancesAddress
	}
	return ""
}

func (m *QueryListPoolAddressesResponse) GetEmissionModuleAddress() string {
	if m != nil {
		return m.EmissionModuleAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "zetachain.zetacore.emissions.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "zetachain.zetacore.emissions.QueryParamsResponse")
	proto.RegisterType((*QueryListPoolAddressesRequest)(nil), "zetachain.zetacore.emissions.QueryListPoolAddressesRequest")
	proto.RegisterType((*QueryListPoolAddressesResponse)(nil), "zetachain.zetacore.emissions.QueryListPoolAddressesResponse")
}

func init() { proto.RegisterFile("emissions/query.proto", fileDescriptor_6e578782beb6ef82) }

var fileDescriptor_6e578782beb6ef82 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3d, 0x6f, 0x13, 0x41,
	0x10, 0x86, 0x7d, 0x06, 0x2c, 0xb1, 0x54, 0x2c, 0xe1, 0x43, 0x56, 0x38, 0xa3, 0x53, 0x50, 0x10,
	0x92, 0x6f, 0xe3, 0x20, 0x68, 0xa0, 0xc1, 0x25, 0x1f, 0x22, 0x58, 0x50, 0x40, 0x63, 0xed, 0xf9,
	0x46, 0x97, 0x95, 0xee, 0x76, 0x2e, 0x3b, 0x7b, 0x11, 0xa1, 0xe4, 0x17, 0x20, 0xd1, 0xf2, 0x77,
	0x90, 0x52, 0x46, 0xa2, 0xa1, 0x42, 0xc8, 0xa6, 0xe6, 0x17, 0x50, 0x20, 0xef, 0xad, 0x0d, 0x8e,
	0x89, 0xf9, 0xe8, 0x56, 0x73, 0xef, 0xfb, 0xbc, 0xb3, 0x33, 0x7b, 0xec, 0x22, 0x14, 0x8a, 0x48,
	0xa1, 0x26, 0xb1, 0x57, 0x81, 0x39, 0x88, 0x4b, 0x83, 0x16, 0xf9, 0xfa, 0x6b, 0xb0, 0x72, 0xb4,
	0x2b, 0x95, 0x8e, 0xdd, 0x09, 0x0d, 0xc4, 0x73, 0x65, 0xfb, 0xe6, 0x08, 0xa9, 0x40, 0x12, 0x89,
	0x24, 0xa8, 0x6d, 0x62, 0xbf, 0x97, 0x80, 0x95, 0x3d, 0x51, 0xca, 0x4c, 0x69, 0x69, 0x15, 0xea,
	0x9a, 0xd4, 0xbe, 0xf4, 0x33, 0xa0, 0x94, 0x46, 0x16, 0xe4, 0xeb, 0x6b, 0x19, 0x66, 0xe8, 0x8e,
	0x62, 0x7a, 0xf2, 0xd5, 0xf5, 0x0c, 0x31, 0xcb, 0x41, 0xc8, 0x52, 0x09, 0xa9, 0x35, 0x5a, 0x87,
	0xf2, 0x9e, 0x68, 0x8d, 0xf1, 0xa7, 0xd3, 0xb4, 0x1d, 0x07, 0x1a, 0xc0, 0x5e, 0x05, 0x64, 0xa3,
	0x17, 0xec, 0xc2, 0x42, 0x95, 0x4a, 0xd4, 0x04, 0xbc, 0xcf, 0x5a, 0x75, 0xe0, 0x95, 0xe0, 0x5a,
	0x70, 0xe3, 0xdc, 0xf6, 0x46, 0xbc, 0xea, 0x4e, 0x71, 0xed, 0xee, 0x9f, 0x3e, 0xfc, 0xdc, 0x69,
	0x0c, 0xbc, 0x33, 0xea, 0xb0, 0xab, 0x0e, 0xfd, 0x48, 0x91, 0xdd, 0x41, 0xcc, 0xef, 0xa7, 0xa9,
	0x01, 0x22, 0x98, 0x67, 0x7f, 0x0f, 0x58, 0x78, 0x92, 0xc2, 0xf7, 0xf1, 0x9c, 0x6d, 0x56, 0x3a,
	0x55, 0x64, 0x8d, 0x4a, 0x2a, 0x0b, 0xe9, 0x10, 0x13, 0x02, 0xb3, 0x0f, 0x66, 0x98, 0xc8, 0x5c,
	0xea, 0x11, 0xd0, 0x50, 0xd6, 0x26, 0xd7, 0xe8, 0xd9, 0xc1, 0xc6, 0x82, 0xfc, 0x89, 0x57, 0xf7,
	0xbd, 0xd8, 0x07, 0xf0, 0x87, 0x2c, 0x5a, 0xc4, 0x5a, 0xa2, 0x65, 0x62, 0xd3, 0x11, 0x3b, 0x0b,
	0xca, 0x67, 0x44, 0xc7, 0x61, 0x77, 0xd8, 0xe5, 0xd9, 0x24, 0x86, 0x05, 0xa6, 0x55, 0x0e, 0x73,
	0xc2, 0x29, 0x47, 0x98, 0x3f, 0x93, 0xc7, 0xee, 0xab, 0xf7, 0x6d, 0x7f, 0x6b, 0xb2, 0x33, 0xee,
	0xfa, 0xfc, 0x7d, 0xc0, 0x5a, 0xf5, 0x08, 0xf9, 0xd6, 0xea, 0x41, 0x2f, 0x6f, 0xb0, 0xdd, 0xfb,
	0x07, 0x47, 0x3d, 0xd5, 0xa8, 0xfb, 0xe6, 0xe3, 0xd7, 0x77, 0xcd, 0x4d, 0x7e, 0x5d, 0x4c, 0x0d,
	0x5d, 0xe7, 0x15, 0x33, 0xaf, 0x38, 0xfe, 0xe6, 0xf8, 0x87, 0x80, 0x9d, 0x5f, 0x5a, 0x11, 0xbf,
	0xfb, 0x17, 0xb9, 0x27, 0xad, 0xbe, 0x7d, 0xef, 0xff, 0xcc, 0xbe, 0xff, 0xdb, 0xae, 0x7f, 0xc1,
	0xbb, 0x7f, 0xe8, 0x3f, 0x57, 0x64, 0x67, 0xbb, 0x00, 0xea, 0x3f, 0x38, 0x1c, 0x87, 0xc1, 0xd1,
	0x38, 0x0c, 0xbe, 0x8c, 0xc3, 0xe0, 0xed, 0x24, 0x6c, 0x1c, 0x4d, 0xc2, 0xc6, 0xa7, 0x49, 0xd8,
	0x78, 0xb9, 0x95, 0x29, 0xbb, 0x5b, 0x25, 0xf1, 0x08, 0x8b, 0xdf, 0x22, 0x5f, 0xfd, 0x02, 0xb5,
	0x07, 0x25, 0x50, 0xd2, 0x72, 0x3f, 0xd5, 0xad, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xf9,
	0x61, 0x4c, 0x03, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ListBalances items.
	ListPoolAddresses(ctx context.Context, in *QueryListPoolAddressesRequest, opts ...grpc.CallOption) (*QueryListPoolAddressesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.emissions.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListPoolAddresses(ctx context.Context, in *QueryListPoolAddressesRequest, opts ...grpc.CallOption) (*QueryListPoolAddressesResponse, error) {
	out := new(QueryListPoolAddressesResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.emissions.Query/ListPoolAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ListBalances items.
	ListPoolAddresses(context.Context, *QueryListPoolAddressesRequest) (*QueryListPoolAddressesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ListPoolAddresses(ctx context.Context, req *QueryListPoolAddressesRequest) (*QueryListPoolAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPoolAddresses not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.emissions.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListPoolAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListPoolAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListPoolAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.emissions.Query/ListPoolAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListPoolAddresses(ctx, req.(*QueryListPoolAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.emissions.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ListPoolAddresses",
			Handler:    _Query_ListPoolAddresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emissions/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryListPoolAddressesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListPoolAddressesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListPoolAddressesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryListPoolAddressesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListPoolAddressesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListPoolAddressesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EmissionModuleAddress) > 0 {
		i -= len(m.EmissionModuleAddress)
		copy(dAtA[i:], m.EmissionModuleAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EmissionModuleAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UndistributedTssBalancesAddress) > 0 {
		i -= len(m.UndistributedTssBalancesAddress)
		copy(dAtA[i:], m.UndistributedTssBalancesAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.UndistributedTssBalancesAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UndistributedObserverBalancesAddress) > 0 {
		i -= len(m.UndistributedObserverBalancesAddress)
		copy(dAtA[i:], m.UndistributedObserverBalancesAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.UndistributedObserverBalancesAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryListPoolAddressesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryListPoolAddressesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UndistributedObserverBalancesAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.UndistributedTssBalancesAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.EmissionModuleAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListPoolAddressesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListPoolAddressesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListPoolAddressesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListPoolAddressesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListPoolAddressesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListPoolAddressesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UndistributedObserverBalancesAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UndistributedObserverBalancesAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UndistributedTssBalancesAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UndistributedTssBalancesAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmissionModuleAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmissionModuleAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
