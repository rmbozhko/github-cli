/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Username struct {
	Login string `json:"login"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-cli",
	Short: "Welcome to the github-cli!",
	Long: `This is a CLI application written in Go to interact with GitHub.

All you need to do is change the name of the project "go-cli-template"
to the name of your project and you are good to go!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	username string
	priority string
	showAll  bool
)

var activityCmd = &cobra.Command{
	Use:   "activity <username>",
	Short: "Set user for a GitHub activity context",
	Long:  "Set the GitHub username for activities",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("GitHub username is required")
		} else if len(args) > 1 {
			return errors.New("only one username can be set at a time")
		}
		return nil
	},
	Run: setUsername,
}

const UserURL = "https://api.github.com/users/%s"

func setUsername(cmd *cobra.Command, args []string) {
	u := Username{Login: args[0]}

	f, err := os.OpenFile(viper.GetString("file"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	_, err = f.Write(b)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("GitHub username set to '%s'\n", u.Login)

	// req, err := http.NewRequest("GET", fmt.Sprintf(UserURL, username), nil)
	// if err != nil {
	// 	return nil, err
	// }
	// req.Header.Set(
	// 	"Accept", "application/vnd.github.v3.text-match+json")
	// resp, err := http.DefaultClient.Do(req)
}

// func getUsername() string {
// }

func setupConfig() {
	viper.SetConfigName("github-cli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.SetDefault("file", filepath.Join(os.Getenv("HOME"), ".github-cli.yaml"))

	viper.ReadInConfig()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cli-template.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	activityCmd.SetUsageTemplate("Usage: github activity <username>")
	rootCmd.AddCommand(activityCmd)

	setupConfig()
}
