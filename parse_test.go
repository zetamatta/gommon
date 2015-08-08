package gommon

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	ParseString("(a (b c d) e d)").Print(os.Stdout)
	println()

	ParseString("(a (b () c d) e d)").Print(os.Stdout)
	println()
}
