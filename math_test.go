package main
import "testing"

func TesteSoma(t *testing.T) {
	total := Soma(15,15)

	if total != 30 {
		t.Errorf("esperado 30, mas obteve %d", total)
	}
}