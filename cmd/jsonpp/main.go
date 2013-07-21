// commmand jsonpp pretty-prints stdin and prints to stdout
package main

import (
	"fmt"
	"github.com/traviscline/jsonpp"
	"io/ioutil"
	"os"
)

func main() {
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
