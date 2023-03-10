package main

import (
	"fmt"

	"github.com/xunzhuo/kube-prow-bot/cmd/kube-prow-bot/config"
	"github.com/xunzhuo/kube-prow-bot/pkg/core"
	"k8s.io/klog"
)

func main() {
	klog.Info("Starting Kube Prow Bot...")

	if err := config.InitConfig(); err != nil {
		klog.Error(err)
		return
	}
	if config.Get().TYPE == "created" {
		core.Greeting()
		core.Help()
		if err := core.RunCommands(); err != nil {
			core.Response(fmt.Sprintf("kube prow bot 🤖️ occurred errors:\n\n```\n%s\n```", err.Error()))
			core.Help()
		}
	}
	if config.Get().TYPE == "comment" {
		if err := core.RunCommands(); err != nil {
			core.Response(fmt.Sprintf("kube prow bot 🤖️ occurred errors:\n\n```\n%s\n```", err.Error()))
			core.Help()
		}
	}
}
