// commmand jsonpp pretty-prints stdin and prints to stdout
package main

import (
	"fmt"
	"github.com/traviscline/jsonpp"
	"io/ioutil"
	"os"
)

// Return true if os.Stdin appears to be interactive
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}

func main() {
	if isInteractive() {
		fmt.Fprintln(os.Stderr, os.Args[0], "expects stdin")
		os.Exit(1)
	}
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading stdin:", err)
		os.Exit(1)
	}
	out, err := jsonpp.Pretty(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error formatting:", err)
		os.Exit(1)
	}
	_, err = os.Stdout.Write(out)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing:", err)
		os.Exit(1)
	}
}
