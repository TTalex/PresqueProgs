package tiredcalculator

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/TTalex/stack"
)

func precedence(op string) int {
	m := make(map[string]int)
	m["*"] = 3
	m["/"] = 3
	m["+"] = 2
	m["-"] = 2
	return m[op]
}

func AddSpaces(s string) string {
	last_c := 'q'
	res := ""
	for _, c := range s {
		if last_c != 'q' {
			if c == '(' || c == ')' || c == '*' || c == '/' || c == '+' || c == '-' {
				if last_c != ' ' {
					res += string(' ')
				}
			} else if last_c == '(' || last_c == ')' || last_c == '*' || last_c == '/' || last_c == '+' || last_c == '-' {
				if c != ' ' {
					res += string(' ')
				}
			}
		}
		res += string(c)
		last_c = c
	}
	return res
}

func ShutingYard(s string) []interface{} {
	output := stack.Stack{}
	operators := stack.Stack{}
	s = AddSpaces(s)
	slice := strings.Split(s, " ")
	for k := 0; k < len(slice); k++ {
		v := slice[k]
		if _, err := strconv.ParseFloat(v, 64); err == nil {
			output.Push(v)
		} else {
			switch v {
			case "(":
				operators.Push(v)
			case ")":
				for topop := operators.Get(); topop != nil && topop != "("; topop = operators.Get() {
					output.Push(operators.Pop())
				}
				operators.Pop()
			default:
				for topop := operators.Get(); topop != nil && precedence(topop.(string)) >= precedence(v); topop = operators.Get() {
					output.Push(operators.Pop())
				}
				operators.Push(v)
			}
		}
	}
	for k := 0; operators.Size() > 0; k++ {
		output.Push(operators.Pop())
	}
	return output.GetList()
}

func Rpn(postfix []interface{}, tired_factor int) float64 {
	res_stack := stack.Stack{}
	for i := 0; i < len(postfix); i++ {
		e := postfix[i].(string)
		if _, err := strconv.ParseFloat(e, 64); err != nil {
			// This is an operator
			op2 := res_stack.Pop()
			op1 := res_stack.Pop()
			operand_1, _ := strconv.ParseFloat(op1.(string), 64)
			operand_2, _ := strconv.ParseFloat(op2.(string), 64)
			operand_1 = TiredCalculator(operand_1, tired_factor)
			operand_2 = TiredCalculator(operand_2, tired_factor)
			result := 0.0
			switch e {
			case "*":
				result = operand_1 * operand_2
			case "/":
				result = operand_1 / operand_2
			case "-":
				result = operand_1 - operand_2
			case "+":
				result = operand_1 + operand_2
			}
			res_stack.Push(strconv.FormatFloat(result, 'f', -1, 64))
		} else {
			// This is a number
			res_stack.Push(e)
		}
	}
	r, _ := strconv.ParseFloat(res_stack.Pop().(string), 64)
	return TiredCalculator(r, tired_factor)
}
func Calculate(s string) float64 {
	sy := ShutingYard(s)
	res := Rpn(sy, 0)
	return res
}
func TiredCalculate(s string, factor int) float64 {
	sy := ShutingYard(s)
	res := Rpn(sy, factor)
	return res
}
func TiredCalculator(f float64, factor int) float64 {
	if factor == 0 {
		return f
	}
	s := strconv.FormatFloat(f, 'f', -1, 64)
	split := strings.Split(s, ".")
	ints := split[0]
	if len(ints) > factor {
		f = f / 10
		f = math.Round(f)
		f = f * 10
	} else if len(split) > 1 {
		decimals := split[1]
		if len(decimals) > factor {
			f = f * 100
			f = math.Round(f)
			f = f / 100
		}
	}
	return f
}
func Calculator() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(TiredCalculate(text[:len(text)-2], 1))
	}
}
