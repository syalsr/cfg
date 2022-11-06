// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/config.proto

package proto

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

// ConfigWrapperClient is the client API for ConfigWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigWrapperClient interface {
	GetConfig(ctx context.Context, in *RequestService, opts ...grpc.CallOption) (*ResponseService, error)
	CreateConfig(ctx context.Context, in *RequestConfig, opts ...grpc.CallOption) (*Status, error)
	UpdateConfig(ctx context.Context, in *RequestConfig, opts ...grpc.CallOption) (*Status, error)
	DeleteUnusedConfig(ctx context.Context, in *RequestService, opts ...grpc.CallOption) (*Status, error)
}

type configWrapperClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigWrapperClient(cc grpc.ClientConnInterface) ConfigWrapperClient {
	return &configWrapperClient{cc}
}

func (c *configWrapperClient) GetConfig(ctx context.Context, in *RequestService, opts ...grpc.CallOption) (*ResponseService, error) {
	out := new(ResponseService)
	err := c.cc.Invoke(ctx, "/proto.ConfigWrapper/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configWrapperClient) CreateConfig(ctx context.Context, in *RequestConfig, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.ConfigWrapper/CreateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configWrapperClient) UpdateConfig(ctx context.Context, in *RequestConfig, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.ConfigWrapper/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configWrapperClient) DeleteUnusedConfig(ctx context.Context, in *RequestService, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.ConfigWrapper/DeleteUnusedConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigWrapperServer is the server API for ConfigWrapper service.
// All implementations must embed UnimplementedConfigWrapperServer
// for forward compatibility
type ConfigWrapperServer interface {
	GetConfig(context.Context, *RequestService) (*ResponseService, error)
	CreateConfig(context.Context, *RequestConfig) (*Status, error)
	UpdateConfig(context.Context, *RequestConfig) (*Status, error)
	DeleteUnusedConfig(context.Context, *RequestService) (*Status, error)
	mustEmbedUnimplementedConfigWrapperServer()
}

// UnimplementedConfigWrapperServer must be embedded to have forward compatible implementations.
type UnimplementedConfigWrapperServer struct {
}

func (UnimplementedConfigWrapperServer) GetConfig(context.Context, *RequestService) (*ResponseService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigWrapperServer) CreateConfig(context.Context, *RequestConfig) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfig not implemented")
}
func (UnimplementedConfigWrapperServer) UpdateConfig(context.Context, *RequestConfig) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigWrapperServer) DeleteUnusedConfig(context.Context, *RequestService) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUnusedConfig not implemented")
}
func (UnimplementedConfigWrapperServer) mustEmbedUnimplementedConfigWrapperServer() {}

// UnsafeConfigWrapperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigWrapperServer will
// result in compilation errors.
type UnsafeConfigWrapperServer interface {
	mustEmbedUnimplementedConfigWrapperServer()
}

func RegisterConfigWrapperServer(s grpc.ServiceRegistrar, srv ConfigWrapperServer) {
	s.RegisterService(&ConfigWrapper_ServiceDesc, srv)
}

func _ConfigWrapper_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigWrapperServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConfigWrapper/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigWrapperServer).GetConfig(ctx, req.(*RequestService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigWrapper_CreateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigWrapperServer).CreateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConfigWrapper/CreateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigWrapperServer).CreateConfig(ctx, req.(*RequestConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigWrapper_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigWrapperServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConfigWrapper/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigWrapperServer).UpdateConfig(ctx, req.(*RequestConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigWrapper_DeleteUnusedConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigWrapperServer).DeleteUnusedConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConfigWrapper/DeleteUnusedConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigWrapperServer).DeleteUnusedConfig(ctx, req.(*RequestService))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigWrapper_ServiceDesc is the grpc.ServiceDesc for ConfigWrapper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigWrapper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConfigWrapper",
	HandlerType: (*ConfigWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ConfigWrapper_GetConfig_Handler,
		},
		{
			MethodName: "CreateConfig",
			Handler:    _ConfigWrapper_CreateConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _ConfigWrapper_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteUnusedConfig",
			Handler:    _ConfigWrapper_DeleteUnusedConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/config.proto",
}
