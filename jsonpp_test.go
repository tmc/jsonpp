package jsonpp

import (
	"testing"
)

var (
	src      = []byte(`{"this": "is", "a": {"test": ["of", "jsonpp"]}}`)
	expected = `{
  "this": "is",
  "a": {
    "test": [
      "of",
      "jsonpp"
    ]
  }
}`
)

func TestBasics(t *testing.T) {
	actual, err := Pretty(src)
	if err != nil {
		t.Error(err)
	}
	if string(actual) != expected {
		t.Errorf("'%s' was expected but got '%s'", expected, string(actual))
	}
}
