package version

import (
	"fmt"
	"runtime"

	"github.com/yuk7/wsldl/lib/cmdline"
)

// GetCommand returns the version command structure
func GetCommand() cmdline.Command {
	return cmdline.Command{
		Names: []string{"version", "-v", "--version"},
		Run: func(distroName string, args []string) {
			Execute()
		},
	}
}

// Execute is default version entrypoint. prints version information
func Execute() {
	fmt.Printf("%s, version %s  (%s)\n", project, version, runtime.GOARCH)
	fmt.Printf("%s\n", url)
}
