package main

import (
	"os"

	"github.com/jacksoom/gotabit/app"
	"github.com/jacksoom/gotabit/cmd/cosmoscmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	cmdOptions := GetWasmCmdOptions()

	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.CoinType,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
		cmdOptions...,
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
