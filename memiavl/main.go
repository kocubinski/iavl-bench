package main

import (
	"fmt"
	"os"

	"github.com/cosmos/iavl-bench/memiavl/cmd"
	"github.com/spf13/cobra"
)

func rootCommand() (*cobra.Command, error) {
	root := &cobra.Command{
		Use:   "memiavl-bench",
		Short: "benchmark crypto-org-chain/cronos/memiavl",
	}
	return root, nil
}

func main() {
	root, err := rootCommand()
	if err != nil {
		os.Exit(1)
	}
	root.AddCommand(cmd.RunCommand(), cmd.BuildCommand())

	if err := root.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}
