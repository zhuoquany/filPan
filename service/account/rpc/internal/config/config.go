// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
}
