/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	isUseWeb     bool
	isInverted   bool
	outputMode   string
	outputWidth  int
	asciiPattern string
)

var rootCmd = &cobra.Command{
	Use:   "imgcli-cobra",
	Short: "A rough copy of imgcli written with cobra",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(func() {
		switch outputMode {
		case "ascii":
		case "color":
		case "box":
		default:
			fmt.Println("Please provide a valid print mode (color, ascii, or box)")
			os.Exit(1)
		}

		if rootCmd.Flag("ascii").Changed {
			outputMode = "ascii"
		}
	})

	rootCmd.PersistentFlags().BoolVarP(&isUseWeb, "web", "W", false, "Whether the source image is in the filesystem or fetched from the web")
	rootCmd.PersistentFlags().BoolVarP(&isInverted, "invert", "i", false, "Whether the the print will be inverted or not")
	rootCmd.PersistentFlags().StringVarP(&outputMode, "mode", "m", "ascii", "he mode the image will be printed in")
	rootCmd.PersistentFlags().IntVarP(&outputWidth, "width", "w", 100, "The number of characters in each row of the output")
	rootCmd.PersistentFlags().StringVarP(&asciiPattern, "ascii", "p", " .-+*#%@", "The pattern of ascii characters from least to greatest visibility. Patterns of over 8 characters are not recommended")

}
