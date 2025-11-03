package main

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/gateway"
)

func main() {
	var c gateway.GatewayConf
	conf.MustLoad("etc/gateway.yaml", &c)

	gw := gateway.MustNewServer(c)
	defer gw.Stop()

	gw.Start()
}
