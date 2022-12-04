package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func subtraction(a int, b int) int {
	return a - b
}

func division(a int, b int) int {
	return a / b
}

func multiplication(a int, b int) int {
	return a * b
}

func operation(a int, b int, sign string) int {
	operator := map[string]func(a int, b int) int{
		"+": add,
		"-": subtraction,
		"/": division,
		"*": multiplication}
	return operator[sign](a, b)
}

func romanToArabic(a string) int {
	var digit int
	romanToArabicMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	for key := range romanToArabicMap {
		if a == key {
			digit = romanToArabicMap[key]
		}
	}
	return digit
}

func romanOperations(a string, b string, sign string) string {
	aInt := romanToArabic(a)
	bInt := romanToArabic(b)
	resultInt := operation(aInt, bInt, sign)
	if resultInt < 1 {
		panic("roman numeral system does not contain negative numbers and 0")
	}

	var resultString string

	arabicToRomanKeys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	arabicToRomanValues := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(arabicToRomanKeys); i++ {
		for intermediateResult := resultInt / arabicToRomanKeys[i]; intermediateResult > 0; intermediateResult -= 1 {
			resultString += arabicToRomanValues[i]
			resultInt -= arabicToRomanKeys[i]
		}
	}
	return resultString
}

func validOperation(input string) {
	validOperations := []string{"+", "-", "*", "/"}
	var inputSlice []string
	var sign string
	for _, operator := range validOperations {
		if strings.Contains(input, operator) {
			inputSlice = strings.Split(input, operator)
			sign = operator
			break
		}
	}
	if sign == "" || len(inputSlice) != 2 {
		panic("unsupported operation")
	}

	romanSigns := "IVIIIX"
	var aRoman, bRoman string
	a, err := strconv.Atoi(inputSlice[0])
	if err != nil {
		if strings.Contains(romanSigns, inputSlice[0]) {
			aRoman = inputSlice[0]
		} else {
			panic("unknown symbols")
		}
	}
	b, err := strconv.Atoi(inputSlice[1])
	if err != nil {
		if strings.Contains(romanSigns, inputSlice[1]) {
			bRoman = inputSlice[1]
		} else {
			panic("unknown symbols")
		}
	}
	if aRoman != "" && bRoman != "" {
		fmt.Println(romanOperations(aRoman, bRoman, sign))
	} else if a > 0 && a <= 10 && b > 0 && b <= 10 {
		fmt.Println(operation(a, b, sign))
	} else {
		panic("your number or numbers are not between 0 and 10 or you use arabic and roman digit simultaneously")
	}
}

func main() {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.Trim(input, "\n")
	input = strings.Replace(input, " ", "", -1)
	validOperation(input)
}
