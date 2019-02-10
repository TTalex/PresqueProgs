package tiredcalculator

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	s := "3 + 4 * ( 5 - 6 )"
	expected := -1.0
	res := Calculate(s)
	if res != expected {
		t.Errorf("TestCalculate fail %f", res)
	}

	s = "5*4+ (3 - 1)*9"
	expected = 38.0
	res = Calculate(s)
	if res != expected {
		t.Errorf("TestCalculate fail %f", res)
	}
	s = "1.5 * 2"
	expected = 3.0
	res = Calculate(s)
	if res != expected {
		t.Errorf("TestCalculate fail %f", res)
	}
}

func TestAddSpaces(t *testing.T) {
	expected := "3 - 4"
	res := AddSpaces("3-4")
	if res != expected {
		t.Errorf("AddSpaces fail %s", res)
	}
	expected = "3 * 4"
	res = AddSpaces("3*4")
	if res != expected {
		t.Errorf("AddSpaces fail %s", res)
	}
	expected = "( 3 * 4 ) + 2"
	res = AddSpaces("(3* 4)+2")
	if res != expected {
		t.Errorf("AddSpaces fail %s", res)
	}
	expected = "1.5 * 2"
	res = AddSpaces("1.5 * 2")
	if res != expected {
		t.Errorf("AddSpaces fail %s", res)
	}
	// expected = "-1 + 1"
	// res = AddSpaces("-1 + 1")
	// if res != expected {
	// 	t.Errorf("AddSpaces fail %s", res)
	// }
}

func TestTiredCalculator(t *testing.T) {
	expected := 15430.0
	res := TiredCalculator(15432.32, 4)
	if res != expected {
		t.Errorf("TiredCalculator fail %f", res)
	}
	expected = 3.14
	res = TiredCalculator(3.1411641, 4)
	if res != expected {
		t.Errorf("TiredCalculator fail %f", res)
	}
}

func TestTiredCalculate(t *testing.T) {
	s := "3 + 4 * ( 5 - 6 )"
	expected := 3.0
	res := TiredCalculate(s, 1)
	if res != expected {
		t.Errorf("TestTiredCalculate fail %f", res)
	}
	s = "3 + 45 * ( 50 - 6 )"
	expected = 1980.0
	res = TiredCalculate(s, 3)
	if res != expected {
		t.Errorf("TestTiredCalculate fail %f", res)
	}
}
