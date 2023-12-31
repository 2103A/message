// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: user/user.proto

package user

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

const (
	User_UserRegister_FullMethodName         = "/user.User/UserRegister"
	User_UserLogin_FullMethodName            = "/user.User/UserLogin"
	User_UserLaunch_FullMethodName           = "/user.User/UserLaunch"
	User_UserUpdatePassword_FullMethodName   = "/user.User/UserUpdatePassword"
	User_CreateUserStarGood_FullMethodName   = "/user.User/CreateUserStarGood"
	User_DeleteUserStarGood_FullMethodName   = "/user.User/DeleteUserStarGood"
	User_DeleteUserStarsGood_FullMethodName  = "/user.User/DeleteUserStarsGood"
	User_GetUserStarGoodInfo_FullMethodName  = "/user.User/GetUserStarGoodInfo"
	User_GetUserStarGoodInfos_FullMethodName = "/user.User/GetUserStarGoodInfos"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 注册
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 登录
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 退出登录
	UserLaunch(ctx context.Context, in *UserLaunchRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	// 修改密码
	UserUpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*UserUpdateLoginResponse, error)
	CreateUserStarGood(ctx context.Context, in *CreateUserStarGoodRequest, opts ...grpc.CallOption) (*CreateUserStarGoodResponse, error)
	DeleteUserStarGood(ctx context.Context, in *DeleteUserStarGoodRequest, opts ...grpc.CallOption) (*DeleteUserStarGoodResponse, error)
	DeleteUserStarsGood(ctx context.Context, in *DeleteUserStarGoodsRequest, opts ...grpc.CallOption) (*DeleteUserStarGoodsResponse, error)
	GetUserStarGoodInfo(ctx context.Context, in *GetUserStarGoodInfoRequest, opts ...grpc.CallOption) (*GetUserStarGoodInfoRequest, error)
	GetUserStarGoodInfos(ctx context.Context, in *GetUserStarGoodInfosRequest, opts ...grpc.CallOption) (*GetUserStarGoodInfosRequest, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, User_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, User_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLaunch(ctx context.Context, in *UserLaunchRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, User_UserLaunch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserUpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*UserUpdateLoginResponse, error) {
	out := new(UserUpdateLoginResponse)
	err := c.cc.Invoke(ctx, User_UserUpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUserStarGood(ctx context.Context, in *CreateUserStarGoodRequest, opts ...grpc.CallOption) (*CreateUserStarGoodResponse, error) {
	out := new(CreateUserStarGoodResponse)
	err := c.cc.Invoke(ctx, User_CreateUserStarGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserStarGood(ctx context.Context, in *DeleteUserStarGoodRequest, opts ...grpc.CallOption) (*DeleteUserStarGoodResponse, error) {
	out := new(DeleteUserStarGoodResponse)
	err := c.cc.Invoke(ctx, User_DeleteUserStarGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserStarsGood(ctx context.Context, in *DeleteUserStarGoodsRequest, opts ...grpc.CallOption) (*DeleteUserStarGoodsResponse, error) {
	out := new(DeleteUserStarGoodsResponse)
	err := c.cc.Invoke(ctx, User_DeleteUserStarsGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserStarGoodInfo(ctx context.Context, in *GetUserStarGoodInfoRequest, opts ...grpc.CallOption) (*GetUserStarGoodInfoRequest, error) {
	out := new(GetUserStarGoodInfoRequest)
	err := c.cc.Invoke(ctx, User_GetUserStarGoodInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserStarGoodInfos(ctx context.Context, in *GetUserStarGoodInfosRequest, opts ...grpc.CallOption) (*GetUserStarGoodInfosRequest, error) {
	out := new(GetUserStarGoodInfosRequest)
	err := c.cc.Invoke(ctx, User_GetUserStarGoodInfos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 注册
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	// 登录
	UserLogin(context.Context, *UserLoginRequest) (*UserRegisterResponse, error)
	// 退出登录
	UserLaunch(context.Context, *UserLaunchRequest) (*UserRegisterResponse, error)
	// 修改密码
	UserUpdatePassword(context.Context, *UserUpdatePasswordRequest) (*UserUpdateLoginResponse, error)
	CreateUserStarGood(context.Context, *CreateUserStarGoodRequest) (*CreateUserStarGoodResponse, error)
	DeleteUserStarGood(context.Context, *DeleteUserStarGoodRequest) (*DeleteUserStarGoodResponse, error)
	DeleteUserStarsGood(context.Context, *DeleteUserStarGoodsRequest) (*DeleteUserStarGoodsResponse, error)
	GetUserStarGoodInfo(context.Context, *GetUserStarGoodInfoRequest) (*GetUserStarGoodInfoRequest, error)
	GetUserStarGoodInfos(context.Context, *GetUserStarGoodInfosRequest) (*GetUserStarGoodInfosRequest, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServer) UserLogin(context.Context, *UserLoginRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServer) UserLaunch(context.Context, *UserLaunchRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLaunch not implemented")
}
func (UnimplementedUserServer) UserUpdatePassword(context.Context, *UserUpdatePasswordRequest) (*UserUpdateLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdatePassword not implemented")
}
func (UnimplementedUserServer) CreateUserStarGood(context.Context, *CreateUserStarGoodRequest) (*CreateUserStarGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserStarGood not implemented")
}
func (UnimplementedUserServer) DeleteUserStarGood(context.Context, *DeleteUserStarGoodRequest) (*DeleteUserStarGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserStarGood not implemented")
}
func (UnimplementedUserServer) DeleteUserStarsGood(context.Context, *DeleteUserStarGoodsRequest) (*DeleteUserStarGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserStarsGood not implemented")
}
func (UnimplementedUserServer) GetUserStarGoodInfo(context.Context, *GetUserStarGoodInfoRequest) (*GetUserStarGoodInfoRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStarGoodInfo not implemented")
}
func (UnimplementedUserServer) GetUserStarGoodInfos(context.Context, *GetUserStarGoodInfosRequest) (*GetUserStarGoodInfosRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStarGoodInfos not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLaunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserLaunch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLaunch(ctx, req.(*UserLaunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserUpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserUpdatePassword(ctx, req.(*UserUpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUserStarGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserStarGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUserStarGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUserStarGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUserStarGood(ctx, req.(*CreateUserStarGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserStarGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserStarGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserStarGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUserStarGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserStarGood(ctx, req.(*DeleteUserStarGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserStarsGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserStarGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserStarsGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUserStarsGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserStarsGood(ctx, req.(*DeleteUserStarGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserStarGoodInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStarGoodInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserStarGoodInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserStarGoodInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserStarGoodInfo(ctx, req.(*GetUserStarGoodInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserStarGoodInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStarGoodInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserStarGoodInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserStarGoodInfos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserStarGoodInfos(ctx, req.(*GetUserStarGoodInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _User_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "UserLaunch",
			Handler:    _User_UserLaunch_Handler,
		},
		{
			MethodName: "UserUpdatePassword",
			Handler:    _User_UserUpdatePassword_Handler,
		},
		{
			MethodName: "CreateUserStarGood",
			Handler:    _User_CreateUserStarGood_Handler,
		},
		{
			MethodName: "DeleteUserStarGood",
			Handler:    _User_DeleteUserStarGood_Handler,
		},
		{
			MethodName: "DeleteUserStarsGood",
			Handler:    _User_DeleteUserStarsGood_Handler,
		},
		{
			MethodName: "GetUserStarGoodInfo",
			Handler:    _User_GetUserStarGoodInfo_Handler,
		},
		{
			MethodName: "GetUserStarGoodInfos",
			Handler:    _User_GetUserStarGoodInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
