package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	codePointPlus  = '+'
	codePointMinus = '-'
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")

	cashNum     string                // текущее число, составляемое итерационно (цифра к цифре)
	numberCount []string = []string{} // количество чисел в строке (если нужно будет ограничивать количество слагаемых)
	result      int                   // итоговая сумма
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var resString string
	var cashOperand rune //   +/-

	input = replaceWhiteSpacies(input)
	if input == "" {
		return "", fmt.Errorf("error occured - input string is empty: %w", errorEmptyInput)
	}

	// для выявления ошибки "not digit" "на берегу"
	inputWithoutOperands := strings.ReplaceAll(input, "+", "")
	inputWithoutOperands = strings.ReplaceAll(inputWithoutOperands, "-", "")
	_, err = strconv.ParseInt(inputWithoutOperands, 10, 32)
	if err != nil {
		return "", fmt.Errorf("error occured - not digit with err: %w", err)
	}

	// посимвольный перебор (основной цикл)
	for _, v := range input {
		if unicode.IsDigit(v) {
			cashNum += string(v) // собрать числа из цифр
		} else {
			afterDigitReading(cashOperand)
			cashOperand = v
		}
	}
	afterDigitReading(cashOperand)

	// Errors - operand count
	if len(numberCount) == 1 {
		return "", fmt.Errorf("error occured: an expression contains only one operand: %w", errorNotTwoOperands)
	}
	if len(numberCount) > 2 {
		return "", fmt.Errorf("error occured: an expression contains greater than two operands (%d): %w", len(numberCount), errorNotTwoOperands)
	}

	resString = strconv.Itoa(result)
	return resString, nil
}

func afterDigitReading(cashOperand rune) {
	if cashNum != "" {
		numberCount = append(numberCount, cashNum)
	}
	getResultSum(cashOperand)
	cashNum = ""
}

func getResultSum(cashOperand rune) {
	cash := getOperand(string(cashOperand))
	cashNumToInt, _ := strconv.Atoi(string(cashNum))
	if cash == string(codePointPlus) {
		result += cashNumToInt
	}
	if cash == string(codePointMinus) {
		result -= cashNumToInt
	}
}

func replaceWhiteSpacies(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\v", "")
	str = strings.ReplaceAll(str, "\f", "")
	str = strings.ReplaceAll(str, "\r", "")
	return str
}

// если в начале не стоит знак (+/-), то будем считать, что cashOperand = '+'
func getOperand(cash string) string {
	if cash == string(codePointMinus) {
		return cash
	}
	if cash == string(codePointPlus) {
		return cash
	}
	return string(codePointPlus)
}
