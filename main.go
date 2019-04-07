package main

import (
	"github.com/micro-plat/hydra/context"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
)

func main() {
	app := hydra.NewApp(
		hydra.WithPlatName("myplat"),
		hydra.WithSystemName("file"),
		hydra.WithServerTypes("rpc"),
		hydra.WithDebug())
	app.Conf.API.SetMainConf(`{"address":":8090"}`)

	app.Conf.API.SetSubConf("static", `{
		"exts":[".txt"]
	}`)
	app.RPC("/hello", (component.ServiceFunc)(helloWorld))
	app.Start()
}

func helloWorld(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------数据库报警数据收集-------------------")

	ctx.Log.Info("1. 校验参数")

	ctx.Log.Infof("2. 收集消息(%d)", ctx.Request.GetInt("csc_id"))

	ctx.Log.Info("数据库报警消息-收集成功")
	return "hello rpc"
}
