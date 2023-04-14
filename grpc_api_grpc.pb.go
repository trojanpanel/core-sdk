// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: grpc_api.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiNodeServiceClient is the client API for ApiNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiNodeServiceClient interface {
	AddNode(ctx context.Context, in *NodeAddDto, opts ...grpc.CallOption) (*Response, error)
	RemoveNode(ctx context.Context, in *NodeRemoveDto, opts ...grpc.CallOption) (*Response, error)
}

type apiNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiNodeServiceClient(cc grpc.ClientConnInterface) ApiNodeServiceClient {
	return &apiNodeServiceClient{cc}
}

func (c *apiNodeServiceClient) AddNode(ctx context.Context, in *NodeAddDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiNodeService/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiNodeServiceClient) RemoveNode(ctx context.Context, in *NodeRemoveDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiNodeService/RemoveNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiNodeServiceServer is the server API for ApiNodeService service.
// All implementations must embed UnimplementedApiNodeServiceServer
// for forward compatibility
type ApiNodeServiceServer interface {
	AddNode(context.Context, *NodeAddDto) (*Response, error)
	RemoveNode(context.Context, *NodeRemoveDto) (*Response, error)
	mustEmbedUnimplementedApiNodeServiceServer()
}

// UnimplementedApiNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiNodeServiceServer struct {
}

func (UnimplementedApiNodeServiceServer) AddNode(context.Context, *NodeAddDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedApiNodeServiceServer) RemoveNode(context.Context, *NodeRemoveDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedApiNodeServiceServer) mustEmbedUnimplementedApiNodeServiceServer() {}

// UnsafeApiNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiNodeServiceServer will
// result in compilation errors.
type UnsafeApiNodeServiceServer interface {
	mustEmbedUnimplementedApiNodeServiceServer()
}

func RegisterApiNodeServiceServer(s grpc.ServiceRegistrar, srv ApiNodeServiceServer) {
	s.RegisterService(&ApiNodeService_ServiceDesc, srv)
}

func _ApiNodeService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAddDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiNodeServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiNodeService/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiNodeServiceServer).AddNode(ctx, req.(*NodeAddDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiNodeService_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeRemoveDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiNodeServiceServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiNodeService/RemoveNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiNodeServiceServer).RemoveNode(ctx, req.(*NodeRemoveDto))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiNodeService_ServiceDesc is the grpc.ServiceDesc for ApiNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ApiNodeService",
	HandlerType: (*ApiNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNode",
			Handler:    _ApiNodeService_AddNode_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _ApiNodeService_RemoveNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api.proto",
}

// ApiNodeServerServiceClient is the client API for ApiNodeServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiNodeServerServiceClient interface {
	GetNodeServerInfo(ctx context.Context, in *NodeServerInfoDto, opts ...grpc.CallOption) (*Response, error)
}

type apiNodeServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiNodeServerServiceClient(cc grpc.ClientConnInterface) ApiNodeServerServiceClient {
	return &apiNodeServerServiceClient{cc}
}

func (c *apiNodeServerServiceClient) GetNodeServerInfo(ctx context.Context, in *NodeServerInfoDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiNodeServerService/GetNodeServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiNodeServerServiceServer is the server API for ApiNodeServerService service.
// All implementations must embed UnimplementedApiNodeServerServiceServer
// for forward compatibility
type ApiNodeServerServiceServer interface {
	GetNodeServerInfo(context.Context, *NodeServerInfoDto) (*Response, error)
	mustEmbedUnimplementedApiNodeServerServiceServer()
}

// UnimplementedApiNodeServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiNodeServerServiceServer struct {
}

func (UnimplementedApiNodeServerServiceServer) GetNodeServerInfo(context.Context, *NodeServerInfoDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeServerInfo not implemented")
}
func (UnimplementedApiNodeServerServiceServer) mustEmbedUnimplementedApiNodeServerServiceServer() {}

// UnsafeApiNodeServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiNodeServerServiceServer will
// result in compilation errors.
type UnsafeApiNodeServerServiceServer interface {
	mustEmbedUnimplementedApiNodeServerServiceServer()
}

func RegisterApiNodeServerServiceServer(s grpc.ServiceRegistrar, srv ApiNodeServerServiceServer) {
	s.RegisterService(&ApiNodeServerService_ServiceDesc, srv)
}

func _ApiNodeServerService_GetNodeServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeServerInfoDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiNodeServerServiceServer).GetNodeServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiNodeServerService/GetNodeServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiNodeServerServiceServer).GetNodeServerInfo(ctx, req.(*NodeServerInfoDto))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiNodeServerService_ServiceDesc is the grpc.ServiceDesc for ApiNodeServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiNodeServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ApiNodeServerService",
	HandlerType: (*ApiNodeServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeServerInfo",
			Handler:    _ApiNodeServerService_GetNodeServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api.proto",
}

// ApiAccountServiceClient is the client API for ApiAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiAccountServiceClient interface {
	RemoveAccount(ctx context.Context, in *AccountRemoveDto, opts ...grpc.CallOption) (*Response, error)
}

type apiAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiAccountServiceClient(cc grpc.ClientConnInterface) ApiAccountServiceClient {
	return &apiAccountServiceClient{cc}
}

func (c *apiAccountServiceClient) RemoveAccount(ctx context.Context, in *AccountRemoveDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiAccountService/RemoveAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiAccountServiceServer is the server API for ApiAccountService service.
// All implementations must embed UnimplementedApiAccountServiceServer
// for forward compatibility
type ApiAccountServiceServer interface {
	RemoveAccount(context.Context, *AccountRemoveDto) (*Response, error)
	mustEmbedUnimplementedApiAccountServiceServer()
}

// UnimplementedApiAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiAccountServiceServer struct {
}

func (UnimplementedApiAccountServiceServer) RemoveAccount(context.Context, *AccountRemoveDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAccount not implemented")
}
func (UnimplementedApiAccountServiceServer) mustEmbedUnimplementedApiAccountServiceServer() {}

// UnsafeApiAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiAccountServiceServer will
// result in compilation errors.
type UnsafeApiAccountServiceServer interface {
	mustEmbedUnimplementedApiAccountServiceServer()
}

func RegisterApiAccountServiceServer(s grpc.ServiceRegistrar, srv ApiAccountServiceServer) {
	s.RegisterService(&ApiAccountService_ServiceDesc, srv)
}

func _ApiAccountService_RemoveAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRemoveDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiAccountServiceServer).RemoveAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiAccountService/RemoveAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiAccountServiceServer).RemoveAccount(ctx, req.(*AccountRemoveDto))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiAccountService_ServiceDesc is the grpc.ServiceDesc for ApiAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ApiAccountService",
	HandlerType: (*ApiAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveAccount",
			Handler:    _ApiAccountService_RemoveAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api.proto",
}

// ApiStateServiceClient is the client API for ApiStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiStateServiceClient interface {
	GetNodeState(ctx context.Context, in *NodeStateDto, opts ...grpc.CallOption) (*Response, error)
	GetNodeServerState(ctx context.Context, in *NodeServerStateDto, opts ...grpc.CallOption) (*Response, error)
}

type apiStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiStateServiceClient(cc grpc.ClientConnInterface) ApiStateServiceClient {
	return &apiStateServiceClient{cc}
}

func (c *apiStateServiceClient) GetNodeState(ctx context.Context, in *NodeStateDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiStateService/GetNodeState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiStateServiceClient) GetNodeServerState(ctx context.Context, in *NodeServerStateDto, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ApiStateService/GetNodeServerState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiStateServiceServer is the server API for ApiStateService service.
// All implementations must embed UnimplementedApiStateServiceServer
// for forward compatibility
type ApiStateServiceServer interface {
	GetNodeState(context.Context, *NodeStateDto) (*Response, error)
	GetNodeServerState(context.Context, *NodeServerStateDto) (*Response, error)
	mustEmbedUnimplementedApiStateServiceServer()
}

// UnimplementedApiStateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiStateServiceServer struct {
}

func (UnimplementedApiStateServiceServer) GetNodeState(context.Context, *NodeStateDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeState not implemented")
}
func (UnimplementedApiStateServiceServer) GetNodeServerState(context.Context, *NodeServerStateDto) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeServerState not implemented")
}
func (UnimplementedApiStateServiceServer) mustEmbedUnimplementedApiStateServiceServer() {}

// UnsafeApiStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiStateServiceServer will
// result in compilation errors.
type UnsafeApiStateServiceServer interface {
	mustEmbedUnimplementedApiStateServiceServer()
}

func RegisterApiStateServiceServer(s grpc.ServiceRegistrar, srv ApiStateServiceServer) {
	s.RegisterService(&ApiStateService_ServiceDesc, srv)
}

func _ApiStateService_GetNodeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStateServiceServer).GetNodeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiStateService/GetNodeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStateServiceServer).GetNodeState(ctx, req.(*NodeStateDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiStateService_GetNodeServerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeServerStateDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiStateServiceServer).GetNodeServerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApiStateService/GetNodeServerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiStateServiceServer).GetNodeServerState(ctx, req.(*NodeServerStateDto))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiStateService_ServiceDesc is the grpc.ServiceDesc for ApiStateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiStateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ApiStateService",
	HandlerType: (*ApiStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeState",
			Handler:    _ApiStateService_GetNodeState_Handler,
		},
		{
			MethodName: "GetNodeServerState",
			Handler:    _ApiStateService_GetNodeServerState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_api.proto",
}
