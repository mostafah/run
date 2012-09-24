// Package run provides an easy way to run external commands. It's not a very
// big package (less than 100 lines of code), but comes handy when one is using
// an external command inside a Go program. Because it gets stdin and returns
// stdout and stderr in []byte.
//
//         stdout, stderr, err := Run("hello", "tr", "eo", "EO")
//         // string(stdout) is now "hEllO"
package run

import (
	"io/ioutil"
	"os/exec"
)

// Run executes commands with given args and stdin and returns its stdout and
// stderr.
func Run(stdin []byte, command string, args ...string) (stdout, stderr []byte, err error) {
	// Setup command and its stdin, stdout, and stderr.
	cmd := exec.Command(command, args...)
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return
	}
	// Start command execution. After calling Start stdin, stdout, and
	// stderr will be writable and readable.
	if err = cmd.Start(); err != nil {
		return
	}
	// Write stdin data.
	if _, err = stdinPipe.Write(stdin); err != nil {
		return
	}
	if err = stdinPipe.Close(); err != nil {
		return
	}
	// Read stderr.
	if stderr, err = ioutil.ReadAll(stderrPipe); err != nil {
		return
	}
	// Read stdout.
	if stdout, err = ioutil.ReadAll(stdoutPipe); err != nil {
		return
	}
	// Wait for the command to finish.
	if err = cmd.Wait(); err != nil {
		return
	}
	// Return result.
	return
}
