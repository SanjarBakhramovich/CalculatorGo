package calculator

import (
	"calculator/roman"
	"fmt"
	"strconv"
	"strings"
)

const (
	Sum            = "+"
	Multiplication = "*"
	Division       = "/"
	Subtraction    = "-"
)

type Calculator struct {
	Input         string
	MathOperation string
	FirstValue    float32
	SecondValue   float32
}

func (c *Calculator) ParseInput() {
	parts := strings.Fields(c.Input)
	if len(parts) != 3 {
		fmt.Println("неверный формат ввода")
		return
	}

	firstValueStr := parts[0]
	secondValueStr := parts[2]

	firstValue, err := roman.RomanToArabic(firstValueStr)
	// check if
	// esli firstValue = Roman
	// Prebrazovat v arabic
	//
	if err != nil {
		floatValue, err := strconv.ParseFloat(firstValueStr, 32)
		if err != nil {
			fmt.Println("Ошибка при преобразовании первого значения", err)
			return
		}
		firstValue = float32(floatValue)
	}

	secondValue, err := roman.RomanToArabic(secondValueStr)
	if err != nil {
		floatValue, err := strconv.ParseFloat(secondValueStr, 32)
		if err != nil {
			fmt.Println("Ошибка при преобразовании второго значения", err)
			return
		}
		secondValue = float32(floatValue)
	}

	c.MathOperation = parts[1]
	c.FirstValue = float32(firstValue)
	c.SecondValue = float32(secondValue)

}
func NewCalculator(input string) *Calculator {
	calculator := &Calculator{Input: input}
	calculator.ParseInput()
	return calculator
}

func (c *Calculator) Calculate() float32 {
	switch c.MathOperation {
	case Sum:
		return c.FirstValue + c.SecondValue
	case Subtraction:
		return c.FirstValue - c.SecondValue
	case Multiplication:
		return c.FirstValue * c.SecondValue
	case Division:
		if c.SecondValue == 0 {
			fmt.Println("Ошибка деление на ноль")
			return 0
		}
		return c.FirstValue / c.SecondValue
	default:
		fmt.Println("Неверная математическая операция")
		return -1
	}
}
