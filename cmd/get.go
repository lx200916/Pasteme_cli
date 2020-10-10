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
	"github.com/spf13/cobra"
	"os"
	"pastemecli/api"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [Key] ([password])",
	Short: "Get A Paste",
	Long: `Access a paste with key "1234" password "1111":
pastemecli get 1234 1111
pastemecli get 1234 -p 1111
-c Disable Auto Copy
`,
	Args: cobra.MaximumNArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("get called")
		pass, _ := cmd.Flags().GetString("pass")
		copyF, _ := cmd.Flags().GetBool("copy")
		if len(args) == 0 {
			if isInputFromPipe() {
				scanner := bufio.NewScanner(os.Stdin)
				var buffer bytes.Buffer
				for scanner.Scan() {
					buffer.WriteString(scanner.Text())
				}
				//fmt.Println(buffer.String())
				contents := buffer.String()
				if len(contents) == 0 {
					fmt.Println("👋Please at least Input the Key of Paste")
					return
				}
				api.Get(pass, contents, copyF)

			} else {
				fmt.Println("👋Please at least Input the Key of Paste")

			}
			return
		}

		if len(args) == 1 {
			api.Get(pass, args[0], copyF)
			return
		}
		if len(args) == 2 {
			api.Get(args[1], args[0], copyF)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("copy", "c", false, "🚫Disable Auto Copy.")
	getCmd.PersistentFlags().StringP("pass", "p", "", "🔑Password for Paste.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
