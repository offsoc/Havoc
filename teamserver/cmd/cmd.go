package cmd

import (
	"fmt"
	"os"

	"Havoc/cmd/server"
	"Havoc/pkg/colors"

	"github.com/spf13/cobra"
)

var (
	VersionNumber = "0.4.1"
	VersionName   = "The Fool"
	DatabasePath  = "data/teamserver.db"

	HavocCli = &cobra.Command{
		Use:          "havoc",
		Short:        "Havoc C2 Framework",
		SilenceUsage: true,
		RunE:         teamserverFunc,
	}

	flags server.TeamserverFlags
)

// init all flags
func init() {
	HavocCli.CompletionOptions.DisableDefaultCmd = true

	// server flags
	CobraServer.Flags().SortFlags = false
	CobraServer.Flags().StringVarP(&flags.Server.Profile, "profile", "", "", "set havoc teamserver profile")
	CobraServer.Flags().BoolVarP(&flags.Server.Debug, "debug", "", false, "enable debug mode")
	CobraServer.Flags().BoolVarP(&flags.Server.DebugDev, "debug-dev", "", false, "enable debug mode for developers (compiles the agent with the debug mode/macro enabled)")
	CobraServer.Flags().BoolVarP(&flags.Server.Default, "default", "d", false, "uses default profile (overwrites --profile)")
	CobraServer.Flags().BoolVarP(&flags.Server.Verbose, "verbose", "v", false, "verbose messages")

	// add commands to the teamserver cli
	HavocCli.Flags().SortFlags = false
	HavocCli.AddCommand(CobraServer)
	HavocCli.AddCommand(CobraClient)
}

func teamserverFunc(cmd *cobra.Command, args []string) error {
	startMenu()

	if len(os.Args) <= 2 {
		err := cmd.Help()
		if err != nil {
			return err
		}
		os.Exit(0)
	}

	return nil
}

func startMenu() {
	fmt.Println(colors.Red("              _______           _______  _______ \n    │\\     /│(  ___  )│\\     /│(  ___  )(  ____ \\\n    │ )   ( ││ (   ) ││ )   ( ││ (   ) ││ (    \\/\n    │ (___) ││ (___) ││ │   │ ││ │   │ ││ │      \n    │  ___  ││  ___  │( (   ) )│ │   │ ││ │      \n    │ (   ) ││ (   ) │ \\ \\_/ / │ │   │ ││ │      \n    │ )   ( ││ )   ( │  \\   /  │ (___) ││ (____/\\\n    │/     \\││/     \\│   \\_/   (_______)(_______/"))
	fmt.Println()
	fmt.Println("  	", colors.Red("pwn"), "and", colors.Blue("elevate"), "until it's done")
	fmt.Println()
}
