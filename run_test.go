package run

import (
	"os/exec"
	"testing"
)

func TestRun(t *testing.T) {
	input := `package main
import "fmt"
func main() {
fmt.Println("hello") }`
	output := `package main

import "fmt"

func main() {
	fmt.Println("hello")
}
`
	stdout, stderr, err := Run([]byte(input), "gofmt")
	if err == exec.ErrNotFound {
		t.Errorf("can't perform test: gofmt is not in $PATH.\n")
	} else if err != nil {
		t.Errorf("unexpected error in running gofmt: %v\n", err)
	} else if len(stderr) != 0 {
		t.Errorf("unexpected stderr in running gofmt: %s\n",
			string(stderr))
	} else if string(stdout) != output {
		t.Errorf("expecting:\n%s, got:\n%s\n", output, string(stdout))
	}

	inputBuggy := `package main
import "fmt"
func main( {
fmt.Println("hello") }`
	stdout, stderr, err = Run([]byte(inputBuggy), "gofmt")
	if len(stderr) == 0 {
		t.Errorf("expecting error, got nothing\n",)
	}
}
