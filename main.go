package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errStrings [4]string
var numbers []int
var operator string

const operators = "+-/*"

func main() {
	//hasOperator := false

	errStrings[0] = "Выдача паники, так как в римской системе нет отрицательных чисел."
	errStrings[1] = "Выдача паники, так как используются одновременно разные системы счисления."
	errStrings[2] = "Выдача паники, так как строка не является математической операцией."
	errStrings[3] = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите значение:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	splitText := strings.Split(text, " ")

	for _, v := range splitText {
		//fmt.Println(v)

		if toNumber, err := strconv.Atoi(v); err == nil {

			if len(numbers) < 2 {
				fmt.Println(toNumber)
				numbers = append(numbers, toNumber)
			} else {
				fmt.Println("So many numbers!")
				break
			}

		} else if len(v) == 1 && strings.Contains(operators, v) {

			if operator == "" {
				fmt.Println(v)
				operator = v
			} else {
				fmt.Println("So many operators!")
				break
			}
		}
	}

	fmt.Println(numbers, operator)
}
