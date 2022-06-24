package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "", "\v", "", "\f", "").Replace(input)

	if input == "" {
		return "", fmt.Errorf("error in operands: %w", errorEmptyInput)
	}

	values := strings.Split(input, "+")

	if len(values) == 1 {
		lastIndex := strings.LastIndex(input, "-")

		switch {
		case lastIndex <= 0:
			return "", fmt.Errorf("error in operands: %w", errorNotTwoOperands)
		case len(input)-1 == lastIndex:
			return "", fmt.Errorf("error in operands: %w", errorNotTwoOperands)
		default:
			return calculate(input[:lastIndex], input[lastIndex:])
		}
	}

	if len(values) != 2 {
		return "", fmt.Errorf("error in operands: %w", errorNotTwoOperands)
	}

	return calculate(values[0], values[1])
}

func calculate(value1, value2 string) (string, error) {
	number1, err1 := strconv.Atoi(value1)
	if err1 != nil {
		return "", fmt.Errorf("error in calculation: %w", err1)
	}

	number2, err2 := strconv.Atoi(value2)
	if err2 != nil {
		return "", fmt.Errorf("error in calculation: %w", err2)
	}

	return strconv.Itoa(number1 + number2), nil
}
