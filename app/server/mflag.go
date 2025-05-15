package server

import (
	"github.com/wangsendi/250514-s-mod/app"
	"github.com/wangsendi/250514-s-mod/pkg/sgin"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"github.com/spf13/cobra"
)

func Cmd() *mflag.Ts {
	mc := mflag.New(app.Flag).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}, "run", "", "app", "mlog", "server")
	return mc
}

func run(cmd *cobra.Command, args []string) {
	_ = map[string]any{"cmd": cmd, "args": args}
	mlog.Info(mlog.H{"msg": "app.Flag", "data": app.Flag})

	defer mlog.Close()

	// 启动gin服务
	ginServer := sgin.New().
		WithAddr(app.Flag.Server.ListenAddr).
		WithMode("debug").
		SetupRouter(SetupRouter)
	if err := ginServer.Start(); err != nil {
		mlog.Fatal(mlog.H{"msg": "gin server start failed", "error": err})
	}
}
