// +build js nacl plan9

package logging

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
