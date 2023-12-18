package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "gcurl",
		Version: "0.0.1",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

}
