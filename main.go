package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var errStrings [5]string
var numbers []int
var operator string
var result int

var isRoman bool
var isArabic bool

const operators = "+-/*"

func getRomanNum(number int) string {

	var result = ""
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, recoding := range conversions {
		for number >= recoding.value {
			result = result + recoding.digit
			number -= recoding.value
		}
	}

	return result
}

func main() {
	romanMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	errStrings[0] = "Выдача паники, так как в римской системе нет отрицательных чисел."
	errStrings[1] = "Выдача паники, так как используются одновременно разные системы счисления."
	errStrings[2] = "Выдача паники, так как строка не является математической операцией."
	errStrings[3] = "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	errStrings[4] = "Выдача паники, так как входное число не подходит под рабочий диапазон (1, 10)."

	reader := bufio.NewReader(os.Stdin)

	for {
		// Обнуление переменных
		isRoman = false
		isArabic = false
		operator = ""
		numbers = numbers[:0]

		fmt.Println("Введите значение (или 'exit' для завершения программы):")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		// Выход из программы
		if text == "exit" {
			break
		}

		splitText := strings.Split(text, " ")

		// Проход по каждому "слову" в переданной строке
		for i, v := range splitText {

			romanNumber := romanMap[v] // римская цифра, если такая имеется
			if romanNumber != 0 {
				if isArabic { // ошибка если уже есть арабские цифры
					panic(errStrings[2])

				} else if len(numbers) < 2 {
					numbers = append(numbers, romanNumber)
					isRoman = true

				} else { // ошибка если больше двух чисел в строке
					panic(errStrings[3])
				}

			} else if toNumber, err := strconv.Atoi(v); err == nil {
				if isRoman { // ошибка если уже есть римские цифры
					panic(errStrings[1])

				} else if toNumber < 1 || toNumber > 10 { // ошибка если число выходит за допустимые пределы
					panic(errStrings[4])

				} else if len(numbers) < 2 {
					numbers = append(numbers, toNumber)
					isArabic = true

				} else {
					panic(errStrings[3])
				}

			} else if len(v) == 1 && strings.Contains(operators, v) {

				if i != 1 { // ошибка если оператор не между числами
					panic(errStrings[2])
				}

				if operator == "" {
					operator = v
				} else { // ошибка если больше одного оператора
					panic(errStrings[3])
				}
			}
		}

		if len(numbers) != 2 || operator == "" {
			panic(errStrings[2])

		} else {

			switch operator {
			case "+":
				result = numbers[0] + numbers[1]
			case "-":
				result = numbers[0] - numbers[1]
			case "*":
				result = numbers[0] * numbers[1]
			case "/":
				result = numbers[0] / numbers[1]
			}

			if !isRoman {
				fmt.Println(result)
			} else {
				if result <= 0 {
					panic(errStrings[0])
				} else {
					fmt.Println(getRomanNum(result))
				}
			}
		}

	}
}
