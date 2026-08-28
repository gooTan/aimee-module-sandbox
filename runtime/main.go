package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JBailes/aimee/server-go/bus"
	handler "github.com/JBailes/aimee/server-go/modules/sandbox"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s DAEMON_MODULE_BUS_SOCKET\n", os.Args[0])
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config := bus.ModuleProcessConfig{
		SocketPath: os.Args[1], ModuleName: "sandbox",
		PrincipalClass: 1, PrincipalRef: 26,
		Stages: []bus.ModuleStage{
		{EventKind: 10753, StageID: 1},
		{EventKind: 10754, StageID: 2},
		{EventKind: 10755, StageID: 3},
		{EventKind: 10756, StageID: 4},
		},
		Handler: handler.Handle,
	}
	if err := bus.RunModuleProcess(ctx, config); err != nil {
		fmt.Fprintf(os.Stderr, "aimee-module-sandbox: %v\n", err)
		os.Exit(1)
	}
}
