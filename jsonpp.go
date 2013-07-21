// package jsonpp provides a simple interface to pretty-printing JSON
package jsonpp

import (
	"bytes"
	"encoding/json"
)

// Pretty-prints the given byte slice with indention of two spaces
func Pretty(src []byte) ([]byte, error) {
	return PrettyIndent(src, "  ")
}

// Pretty-prints the given byte slice using the provided indention string
func PrettyIndent(src []byte, indent string) ([]byte, error) {
	dst := new(bytes.Buffer)
	err := json.Indent(dst, src, "", indent)
	return dst.Bytes(), err
}
