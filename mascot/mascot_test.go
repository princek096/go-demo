package mascot_test

import (
	"myapp/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.Best() != "Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
