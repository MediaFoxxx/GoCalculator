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
var exit bool

const operators = "+-/*"

func errExit(i int) {
	fmt.Println(errStrings[i])
	exit = true
}

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
		exit = false
		operator = ""
		numbers = numbers[:0]

		fmt.Println("Введите значение:")
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
					errExit(2)
					break
				} else if len(numbers) < 2 {
					numbers = append(numbers, romanNumber)
					isRoman = true
				} else { // ошибка если больше двух чисел в строке
					errExit(3)
					break
				}

			} else if toNumber, err := strconv.Atoi(v); err == nil {
				if isRoman { // ошибка если уже есть римские цифры
					errExit(1)
					break
				} else if toNumber < 1 || toNumber > 10 { // ошибка если число выходит за допустимые пределы
					errExit(4)
					break
				} else if len(numbers) < 2 {
					numbers = append(numbers, toNumber)
					isArabic = true
				} else {
					errExit(3) // ошибка если больше двух чисел в строке
					break
				}

			} else if len(v) == 1 && strings.Contains(operators, v) {

				if i != 1 { // ошибка если оператор не между числами
					errExit(2)
					break
				}

				if operator == "" {
					operator = v
				} else { // ошибка если больше одного оператора
					errExit(3)
					break
				}
			}
		}

		if !exit { // пропуск вычисления если были ошибки
			if len(numbers) != 2 || operator == "" {
				fmt.Println(errStrings[2])
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
						fmt.Println(errStrings[0])
					} else {
						fmt.Println(getRomanNum(result))
					}
				}
			}
		}

	}
}
