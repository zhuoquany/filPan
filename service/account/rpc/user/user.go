// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"ipfsdisk/service/account/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChangePasswdReq = pb.ChangePasswdReq
	LoginReq        = pb.LoginReq
	LoginResponse   = pb.LoginResponse
	MsgResponse     = pb.MsgResponse
	ResisterReq     = pb.ResisterReq

	User interface {
		// UserLogin 用户登录
		UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResponse, error)
		// UserResister 用户注册
		UserResister(ctx context.Context, in *ResisterReq, opts ...grpc.CallOption) (*LoginResponse, error)
		// UserChangePasswd 修改密码
		UserChangePasswd(ctx context.Context, in *ChangePasswdReq, opts ...grpc.CallOption) (*MsgResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// UserLogin 用户登录
func (m *defaultUser) UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// UserResister 用户注册
func (m *defaultUser) UserResister(ctx context.Context, in *ResisterReq, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserResister(ctx, in, opts...)
}

// UserChangePasswd 修改密码
func (m *defaultUser) UserChangePasswd(ctx context.Context, in *ChangePasswdReq, opts ...grpc.CallOption) (*MsgResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserChangePasswd(ctx, in, opts...)
}
