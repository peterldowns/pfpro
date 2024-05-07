package root

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/peterldowns/localias/cmd/localias/shared"
	"github.com/peterldowns/localias/pkg/config"
	"github.com/peterldowns/localias/pkg/daemon"
)

var startCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "start",
	Short: "start the proxy server as a daemon process",
	RunE:  startImpl,
}

func startImpl(_ *cobra.Command, _ []string) error {
	// Warn if the daemon was already running
	existing, err := daemon.Status()
	if err != nil {
		return err
	}
	if existing != nil {
		fmt.Printf("replacing existing daemon on pid %d\n", existing.Pid)
	}
	// Apply the config to /etc/hosts
	hctl := shared.Controller()
	cfg := shared.Config()
	if err := config.Apply(hctl, cfg); err != nil {
		return err
	}
	// Start the daemon with the new config.
	return daemon.Start(cfg)
}

func init() {
	Command.AddCommand(startCmd)
}
