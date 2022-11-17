/*
Copyright © 2022 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package ota

import (
	"path/filepath"

	"github.com/apex/log"
	"github.com/blacktop/ipsw/pkg/ota"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	OtaCmd.AddCommand(extractCmd)

	extractCmd.Flags().StringP("output", "o", "", "Output folder")
	viper.BindPFlag("ota.extract.output", extractCmd.Flags().Lookup("output"))
}

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:     "extract <OTA> <PATTERN>",
	Aliases: []string{"e"},
	Short:   "Extract OTA payload files",
	Args:    cobra.MinimumNArgs(2),
	// SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		otaPath := filepath.Clean(args[0])

		log.Infof("Extracting files that match %#v", args[1])
		return ota.Extract(otaPath, args[1], viper.GetString("ota.extract.output"))
	},
}
