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
	"strconv"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = CreateSumCmd()

// CreateSumCmd wrapper fot testing
func CreateSumCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sum",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			SumRun(cmd, args)
			return nil
		},
	}

	cmd.Flags().BoolP("creator", "c", false, "Get get get get")

	return cmd
}

// SumRun is a RunE func
func SumRun(cmd *cobra.Command, args []string) {
	creator, _ := cmd.Flags().GetBool("creator")

	if creator {
		cmd.Println("Created by: mcruzdev")
		return
	}

	length := len(args)

	if length > 0 {

		sum := SumArr(args)

		cmd.Println("Sum is:", sum)
	}
}

func SumArr(arr []string) int64 {

	var sum int64 = 0

	for _, number := range arr {
		parseInt, _ := strconv.ParseInt(number, 10, 64)
		sum += parseInt
	}

	return sum
}

func init() {
	rootCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
