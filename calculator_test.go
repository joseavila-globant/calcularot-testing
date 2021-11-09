package calculator_test

import (
	"calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := calculator.Add(1, 2)
	if sum != 3 {
		t.Error("La suma no es correcta")
		t.Fail()
	} else {
		t.Log("TestAdd finalizado corretamente")
	}
}

func TestAddTestify(t *testing.T) {

	sum := calculator.AddTestify(1, 2)

	assert.Equal(t, 3, sum)

}
