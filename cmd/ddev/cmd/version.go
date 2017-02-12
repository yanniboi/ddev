package cmd

import (
	"fmt"

	"github.com/drud/ddev/pkg/version"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print DRUD version",
	Long:  `Display the version of this DRUD binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		table := uitable.New()
		table.MaxColWidth = 200

		table.AddRow("cli:", version.CliVersion)
		table.AddRow("web:", version.WebImg+":"+version.WebTag)
		table.AddRow("db:", version.DBImg+":"+version.DBTag)
		table.AddRow("commit", version.VERSION)

		fmt.Println(table)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
