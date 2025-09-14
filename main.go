package main

import (
	_ "blog/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"blog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
