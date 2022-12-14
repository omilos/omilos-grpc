// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: omilos.proto

package omilos_grpc

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

// OmilosClient is the client API for Omilos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmilosClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	PostCast(ctx context.Context, in *PostCastRequest, opts ...grpc.CallOption) (*Cast, error)
	LikeCast(ctx context.Context, in *LikeCastRequest, opts ...grpc.CallOption) (*LikeCastResponse, error)
	GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*User, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*Notifications, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GetCast(ctx context.Context, in *GetCastRequest, opts ...grpc.CallOption) (*Cast, error)
	GetCasts(ctx context.Context, in *GetCastsRequest, opts ...grpc.CallOption) (*Casts, error)
}

type omilosClient struct {
	cc grpc.ClientConnInterface
}

func NewOmilosClient(cc grpc.ClientConnInterface) OmilosClient {
	return &omilosClient{cc}
}

func (c *omilosClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) PostCast(ctx context.Context, in *PostCastRequest, opts ...grpc.CallOption) (*Cast, error) {
	out := new(Cast)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/PostCast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) LikeCast(ctx context.Context, in *LikeCastRequest, opts ...grpc.CallOption) (*LikeCastResponse, error) {
	out := new(LikeCastResponse)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/LikeCast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) GetMe(ctx context.Context, in *GetMeRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/GetMe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*Notifications, error) {
	out := new(Notifications)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) GetCast(ctx context.Context, in *GetCastRequest, opts ...grpc.CallOption) (*Cast, error) {
	out := new(Cast)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/GetCast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omilosClient) GetCasts(ctx context.Context, in *GetCastsRequest, opts ...grpc.CallOption) (*Casts, error) {
	out := new(Casts)
	err := c.cc.Invoke(ctx, "/omilos_grpc.Omilos/GetCasts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmilosServer is the server API for Omilos service.
// All implementations must embed UnimplementedOmilosServer
// for forward compatibility
type OmilosServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	PostCast(context.Context, *PostCastRequest) (*Cast, error)
	LikeCast(context.Context, *LikeCastRequest) (*LikeCastResponse, error)
	GetMe(context.Context, *GetMeRequest) (*User, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*Notifications, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GetCast(context.Context, *GetCastRequest) (*Cast, error)
	GetCasts(context.Context, *GetCastsRequest) (*Casts, error)
	mustEmbedUnimplementedOmilosServer()
}

// UnimplementedOmilosServer must be embedded to have forward compatible implementations.
type UnimplementedOmilosServer struct {
}

func (UnimplementedOmilosServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedOmilosServer) PostCast(context.Context, *PostCastRequest) (*Cast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCast not implemented")
}
func (UnimplementedOmilosServer) LikeCast(context.Context, *LikeCastRequest) (*LikeCastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeCast not implemented")
}
func (UnimplementedOmilosServer) GetMe(context.Context, *GetMeRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMe not implemented")
}
func (UnimplementedOmilosServer) GetNotifications(context.Context, *GetNotificationsRequest) (*Notifications, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedOmilosServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedOmilosServer) GetCast(context.Context, *GetCastRequest) (*Cast, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCast not implemented")
}
func (UnimplementedOmilosServer) GetCasts(context.Context, *GetCastsRequest) (*Casts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCasts not implemented")
}
func (UnimplementedOmilosServer) mustEmbedUnimplementedOmilosServer() {}

// UnsafeOmilosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmilosServer will
// result in compilation errors.
type UnsafeOmilosServer interface {
	mustEmbedUnimplementedOmilosServer()
}

func RegisterOmilosServer(s grpc.ServiceRegistrar, srv OmilosServer) {
	s.RegisterService(&Omilos_ServiceDesc, srv)
}

func _Omilos_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_PostCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).PostCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/PostCast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).PostCast(ctx, req.(*PostCastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_LikeCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeCastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).LikeCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/LikeCast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).LikeCast(ctx, req.(*LikeCastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/GetMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).GetMe(ctx, req.(*GetMeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_GetCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).GetCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/GetCast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).GetCast(ctx, req.(*GetCastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omilos_GetCasts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCastsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmilosServer).GetCasts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omilos_grpc.Omilos/GetCasts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmilosServer).GetCasts(ctx, req.(*GetCastsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Omilos_ServiceDesc is the grpc.ServiceDesc for Omilos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Omilos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omilos_grpc.Omilos",
	HandlerType: (*OmilosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Omilos_Login_Handler,
		},
		{
			MethodName: "PostCast",
			Handler:    _Omilos_PostCast_Handler,
		},
		{
			MethodName: "LikeCast",
			Handler:    _Omilos_LikeCast_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _Omilos_GetMe_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _Omilos_GetNotifications_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Omilos_GetUser_Handler,
		},
		{
			MethodName: "GetCast",
			Handler:    _Omilos_GetCast_Handler,
		},
		{
			MethodName: "GetCasts",
			Handler:    _Omilos_GetCasts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omilos.proto",
}
