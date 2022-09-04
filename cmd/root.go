/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"v0-simulator/v0"

	"github.com/spf13/cobra"
)

var apiMethod string //"order", "label", "tracking"
var reqMethod string //"get", "delete", "create", "update"
var numbers string   //customer number, or order number

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "v0",
	Short: "v0 simulator.",
	Long:  `v0 simulator, for testing.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//viper.GetString()

		if apiMethod == "order" {
			if reqMethod == "create" {
				v0.CreateOrderProcess(numbers)
				return
			}
			if reqMethod == "update" {
				v0.UpdateOrderProcess(numbers)
				return
			}
			if reqMethod == "get" {
				v0.GetOrderProcess(numbers)
				return
			}
			if reqMethod == "delete" {
				v0.DeleteOrderProcess(numbers)
				return
			}
		}

		if apiMethod == "label" {
			v0.GetLabelProcess(numbers)
			return
		}

		if apiMethod == "tracking" {
			v0.GetTrackingProcess(numbers)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func init() {

	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xpower.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./conf/config.ini)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.Flags().StringVarP(&url, "url", "u", "", "url (Mandatory)")
	//rootCmd.Flags().StringVarP(&dir, "dir", "d", "", "Destination file's directory")

	rootCmd.Flags().StringVarP(&apiMethod, "apiMethod", "a", "", "order, label, tracking")
	rootCmd.Flags().StringVarP(&reqMethod, "reqMethod", "r", "", "create, get, update, delete")
	rootCmd.Flags().StringVarP(&numbers, "numbers", "n", "",
		"customer number, or order no. When update, should be customerNo,orderNo. When Multi numbers, splitted by , ")

}

func initConfig() {
	//if cfgFile != "" {
	//	viper.SetConfigFile(cfgFile)
	//} else {
	//	viper.SetConfigFile("./conf/config.ini")
	//}

	//viper.AutomaticEnv()

	//err := viper.ReadInConfig()
	//if err == nil {
	//	fmt.Println("Using config file:", viper.ConfigFileUsed())
	//}

	//goroutines = viper.GetInt("default.goroutines")
	//rangeSize = viper.GetInt64("default.range_size")
	//chanLen = viper.GetInt64("default.channel_length")
}
