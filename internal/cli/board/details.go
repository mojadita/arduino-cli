// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package board

import (
	"context"
	"fmt"
	"os"

	"github.com/arduino/arduino-cli/commands/board"
	"github.com/arduino/arduino-cli/internal/cli/arguments"
	"github.com/arduino/arduino-cli/internal/cli/feedback"
	"github.com/arduino/arduino-cli/internal/cli/instance"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"github.com/arduino/arduino-cli/table"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	showFullDetails bool
	listProgrammers bool
	fqbn            arguments.Fqbn
)

func initDetailsCommand() *cobra.Command {
	var detailsCommand = &cobra.Command{
		Use:     fmt.Sprintf("details -b <%s>", tr("FQBN")),
		Short:   tr("Print details about a board."),
		Long:    tr("Show information about a board, in particular if the board has options to be specified in the FQBN."),
		Example: "  " + os.Args[0] + " board details -b arduino:avr:nano",
		Args:    cobra.NoArgs,
		Run:     runDetailsCommand,
	}

	fqbn.AddToCommand(detailsCommand)
	detailsCommand.Flags().BoolVarP(&showFullDetails, "full", "f", false, tr("Show full board details"))
	detailsCommand.Flags().BoolVarP(&listProgrammers, "list-programmers", "", false, tr("Show list of available programmers"))
	detailsCommand.MarkFlagRequired("fqbn")

	return detailsCommand
}

func runDetailsCommand(cmd *cobra.Command, args []string) {
	inst := instance.CreateAndInit()

	logrus.Info("Executing `arduino-cli board details`")

	res, err := board.Details(context.Background(), &rpc.BoardDetailsRequest{
		Instance: inst,
		Fqbn:     fqbn.String(),
	})

	if err != nil {
		feedback.Fatal(tr("Error getting board details: %v", err), feedback.ErrGeneric)
	}

	feedback.PrintResult(detailsResult{details: res})
}

// output from this command requires special formatting, let's create a dedicated
// feedback.Result implementation
type detailsResult struct {
	details *rpc.BoardDetailsResponse
}

func (dr detailsResult) Data() interface{} {
	return dr.details
}

func (dr detailsResult) String() string {
	details := dr.details

	if listProgrammers {
		t := table.New()
		t.AddRow(tr("Id"), tr("Programmer name"))
		for _, programmer := range details.Programmers {
			t.AddRow(programmer.GetId(), programmer.GetName())
		}
		return t.Render()
	}

	// Table is 4 columns wide:
	// |               |                             | |                       |
	// Board name:     Arduino Nano
	//
	// Required tools: arduino:avr-gcc                 5.4.0-atmel3.6.1-arduino2
	//                 arduino:avrdude                 6.3.0-arduino14
	//                 arduino:arduinoOTA              1.2.1
	//
	// Option:         Processor                       cpu
	//                 ATmega328P                    ✔ cpu=atmega328
	//                 ATmega328P (Old Bootloader)     cpu=atmega328old
	//                 ATmega168                       cpu=atmega168
	t := table.New()
	tab := table.New()
	addIfNotEmpty := func(label, content string) {
		if content != "" {
			t.AddRow(label, content)
		}
	}

	t.SetColumnWidthMode(1, table.Average)
	t.AddRow(tr("Board name:"), details.Name)
	t.AddRow(tr("FQBN:"), details.Fqbn)
	addIfNotEmpty(tr("Board version:"), details.Version)
	if details.GetDebuggingSupported() {
		t.AddRow(tr("Debugging supported:"), table.NewCell("✔", color.New(color.FgGreen)))
	}

	if details.Official {
		t.AddRow() // get some space from above
		t.AddRow(tr("Official Arduino board:"),
			table.NewCell("✔", color.New(color.FgGreen)))
	}

	for _, idp := range details.GetIdentificationProperties() {
		t.AddRow() // get some space from above
		header := tr("Identification properties:")
		for k, v := range idp.GetProperties() {
			t.AddRow(header, k+"="+v)
			header = ""
		}
	}

	t.AddRow() // get some space from above
	addIfNotEmpty(tr("Package name:"), details.Package.Name)
	addIfNotEmpty(tr("Package maintainer:"), details.Package.Maintainer)
	addIfNotEmpty(tr("Package URL:"), details.Package.Url)
	addIfNotEmpty(tr("Package website:"), details.Package.WebsiteUrl)
	addIfNotEmpty(tr("Package online help:"), details.Package.Help.Online)

	t.AddRow() // get some space from above
	addIfNotEmpty(tr("Platform name:"), details.Platform.Name)
	addIfNotEmpty(tr("Platform category:"), details.Platform.Category)
	addIfNotEmpty(tr("Platform architecture:"), details.Platform.Architecture)
	addIfNotEmpty(tr("Platform URL:"), details.Platform.Url)
	addIfNotEmpty(tr("Platform file name:"), details.Platform.ArchiveFilename)
	if details.Platform.Size != 0 {
		addIfNotEmpty(tr("Platform size (bytes):"), fmt.Sprint(details.Platform.Size))
	}
	addIfNotEmpty(tr("Platform checksum:"), details.Platform.Checksum)

	t.AddRow() // get some space from above

	tab.SetColumnWidthMode(1, table.Average)
	for _, tool := range details.ToolsDependencies {
		tab.AddRow(tr("Required tool:"), tool.Packager+":"+tool.Name, tool.Version)
		if showFullDetails {
			for _, sys := range tool.Systems {
				tab.AddRow("", tr("OS:"), sys.Host)
				tab.AddRow("", tr("File:"), sys.ArchiveFilename)
				tab.AddRow("", tr("Size (bytes):"), fmt.Sprint(sys.Size))
				tab.AddRow("", tr("Checksum:"), sys.Checksum)
				tab.AddRow("", tr("URL:"), sys.Url)
				tab.AddRow() // get some space from above
			}
		}
	}

	tab.AddRow() // get some space from above
	for _, option := range details.ConfigOptions {
		tab.AddRow(tr("Option:"), option.OptionLabel, "", option.Option)
		for _, value := range option.Values {
			green := color.New(color.FgGreen)
			if value.Selected {
				tab.AddRow("",
					table.NewCell(value.ValueLabel, green),
					table.NewCell("✔", green),
					table.NewCell(option.Option+"="+value.Value, green))
			} else {
				tab.AddRow("",
					value.ValueLabel,
					"",
					option.Option+"="+value.Value)
			}
		}
	}

	tab.AddRow(tr("Programmers:"), tr("ID"), tr("Name"))
	for _, programmer := range details.Programmers {
		tab.AddRow("", programmer.GetId(), programmer.GetName())
	}

	return t.Render() + tab.Render()
}
