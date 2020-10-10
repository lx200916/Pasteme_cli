/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"pastemecli/api"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	//fmt.Println(fileInfo.Size())
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [Content] ([Password] [Custom Key])",
	Short: "Create A Paste.",
	Long: `Create a Paste "xxx" with Password "1234" and Assign Key "abcd":
pastemecli create xxx 1234 abcd
pastemecli create xxx -p 1234 -key abcd
-o flag to enable Burn after reading feature.
-r flag to enable Raw output(Link Address Only).
-c to disable auto copy
Content Can also be provided by Pipe`,
	Args: cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("create called",args)

		pass, _ := cmd.Flags().GetString("pass")
		key, _ := cmd.Flags().GetString("key")
		lang, _ := cmd.Flags().GetString("lang")
		once, _ := cmd.Flags().GetBool("once")
		raw, _ := cmd.Flags().GetBool("raw")
		copyF, _ := cmd.Flags().GetBool("copy")
		if len(args) == 0 {
			if isInputFromPipe() {
				//fmt.Println(111111111111)
				scanner := bufio.NewScanner(os.Stdin)
				var buffer bytes.Buffer
				for scanner.Scan() {
					buffer.WriteString(scanner.Text())
				}
				//fmt.Println(buffer.String())
				contents := buffer.String()
				if len(contents) == 0 {
					fmt.Println("👋Please at least Input the Content of Paste")
					return
				}
				api.CreateBase(contents, pass, lang, key, once, raw, copyF)

			} else {
				fmt.Println("👋Please at least Input the Content of Paste")

			}
			return
		}
		if key != "" {
			flag, _ := regexp.MatchString("[a-z]", key)

			if (len(key) < 2 && len(key) > 9) || !flag {
				fmt.Println("👋Key should only contains digits and lowercase letters, at least one alpha is required")
				return
			}
		}
		if len(args) == 1 {
			api.CreateBase(args[0], pass, lang, key, once, raw, copyF)
		}
		if len(args) == 2 {
			api.CreateBase(args[0], args[1], lang, key, once, raw, copyF)

		}
		if len(args) == 3 {
			flag, _ := regexp.MatchString("[a-z]", args[2])
			if args[2] != "" && ((len(args[2]) < 2 && len(args[2]) > 9) || !flag) {
				fmt.Println("👋Key should only contains digits and lowercase letters, at least one alpha is required")
				return
			}
			api.CreateBase(args[0], args[1], lang, args[2], once, raw, copyF)

		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createCmd.PersistentFlags().StringP("pass", "p", "", "🔑Password for Paste.")
	createCmd.PersistentFlags().StringP("lang", "l", "", "Define the Coding Language of Paste.")
	createCmd.PersistentFlags().StringP("key", "k", "", "Custom Temporary for Paste.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().BoolP("once", "o", false, "Burn after Reading Once.")
	createCmd.Flags().BoolP("raw", "r", false, "Raw URL Output.")
	createCmd.Flags().BoolP("copy", "c", false, "Disable Auto Copy.")

}
