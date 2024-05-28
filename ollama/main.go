package main

import (
	"context"

	"github.com/uppercaveman/ollama-server/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
}
