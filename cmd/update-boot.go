package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vanilla-os/abroot/core"
)

func updateBootUsage(*cobra.Command) error {
	fmt.Print(`Description:
	Update the boot partition for maintenance purposes (only for advanced users).

Usage:
	_update-boot

Options:
	--help/-h		show this message
	--assume-yes/-y		assume yes to all questions
`)

	return nil
}

func NewUpdateBootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "_update-boot",
		Short: "Update the boot partition",
		RunE:  status,
	}
	cmd.SetUsageFunc(updateBootUsage)

	return cmd
}

func status(cmd *cobra.Command, args []string) error {
	if !core.RootCheck(true) {
		return nil
	}

	assumeYes, _ := cmd.Flags().GetBool("assume-yes")
	if !assumeYes {
		if !core.AskConfirmation(`Are you sure you want to proceed?
The boot partition should be updated only if a transaction succeeded. This 
command should be used by advanced users for maintenance purposes.`) {
			return nil
		}
	}

	if err := core.UpdateRootBoot(false); err != nil {
		return err
	}

	return nil
}
