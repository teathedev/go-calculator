package main

import "testing"

func TestAddition(t *testing.T) {
	f := addition(5, 5)

	if f != 10 {
		t.Error("Addition is wrong")
	}
}

func TestSubtraction(t *testing.T) {
	f := subtraction(5, 10)

	if f != -5 {
		t.Error("Subtraction is wrong")
	}
}

func TestMultiplication(t *testing.T) {
	f := multiplication(5, 5)

	if f != 25 {
		t.Error("Multiplication is wrong")
	}
}

func TestDivision(t *testing.T) {
	f := division(25, 5)

	if f != 5 {
		t.Error("Division is wrong")
	}
}
