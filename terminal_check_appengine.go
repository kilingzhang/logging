// +build appengine

package logging

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
