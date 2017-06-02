// Package getpass reads passwords without echoing output.
//
// The current implementation is compatible with modern Windows and *nix systems
// that use cmd.exe, PowerShell, Bash, or Bourne shell environments.
package getpass

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// Get reads a line from stdin with local echoing disabled and returns it.
//
// If prompt is not an empty string, Get displays it before reading input. If
// echoing cannot be disabled for the current terminal, Get quits with an error
// instead of attempting to read input.
func Get(prompt string) ([]byte, error) {
	fd := int(os.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		state, err := terminal.MakeRaw(fd)
		if err != nil {
			f := "getpass: cannot disable terminal echoing — %v"
			return nil, fmt.Errorf(f, err)
		}
		defer terminal.Restore(fd, state)
		defer fmt.Println()
	}

	if prompt == "" {
		prompt = "enter password: "
	}
	fmt.Fprint(os.Stderr, prompt)
	return terminal.ReadPassword(fd)
}
