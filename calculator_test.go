package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := calculator.Add(1, 2)
	if sum != 3 {
		t.Error("La suma no es correcta")
		t.Fail()
	} else {
		t.Log("TestAdd finalizado corretamente")
		t.Log(sum)
	}
}
