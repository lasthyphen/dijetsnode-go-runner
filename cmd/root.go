// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	"fmt"
	"os"

	"github.com/lasthyphen/dijetsnode-go-runner/cmd/control"
	"github.com/lasthyphen/dijetsnode-go-runner/cmd/ping"
	"github.com/lasthyphen/dijetsnode-go-runner/cmd/server"
	"github.com/spf13/cobra"
)

var Version = ""

var rootCmd = &cobra.Command{
	Use:        "dijetsnode-go-runner",
	Short:      "avalanche-network-runner commands",
	SuggestFor: []string{"network-runner"},
	Version:    Version,
}

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCmd.AddCommand(
		server.NewCommand(),
		ping.NewCommand(),
		control.NewCommand(),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "avalanche-network-runner failed %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
