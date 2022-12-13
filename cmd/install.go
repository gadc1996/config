/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install software configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configDir := "config/zsh/"
		configFilename := "zshrc"
		homeDir := os.Getenv("HOME") + "/"

		oldname := homeDir + configDir + configFilename
		newname := homeDir + "." + configFilename

		fmt.Println(homeDir + configDir)
		if _, err := os.Stat(newname); err != nil {
			os.Remove(newname)
		} 
		  
		err := os.Symlink(oldname, newname)

		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println("Config files installation succeded")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
