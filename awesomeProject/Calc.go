package main

import (
	"fmt"
	"strconv"
)

func main() {
	var dig1 string
	var dig2 string
	var oper string
	var dig3 string
	fmt.Println("Введите выражение: ")
	fmt.Scanln(&dig1, &oper, &dig2, &dig3)
	if dig3 != "" {
		fmt.Println("Нельзя использовать больше 2ух операндов")
		return
	}
	if dig1 == "" || oper == "" || dig2 == "" {
		fmt.Println("Нужно больше 1 операнда, повторите ввод")
	} else {
		fmt.Println(changeRoma(dig1, oper, dig2))
	}
}

func changeRoma(a string, b string, c string) string {
	var a2, c2, t int

	romans := []string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	for i, value := range romans {
		if a == value {
			a2 += i
		}
		if c == value {
			c2 += i
		}
	}

	if a2 > 10 || c2 > 10 {
		t := "Aргументы должны быть кратны или меньше 10, повторите ввод"
		return t
	}
	if a2 > 0 && c2 > 0 {
		switch b {
		case "+":
			t = a2 + c2
		case "-":
			t = a2 - c2
		case "*":
			t = a2 * c2
		case "/":
			t = a2 / c2
		default:
			strconv.Itoa(t)
			t := "Не корректный операнд, повторите ввод"
			return t
		}
	} else {
		a2, _ = strconv.Atoi(a)
		c2, _ = strconv.Atoi(c)
		if a2 == 0 || c2 == 0 {
			t := "Нельзя использовать Римские и Арабские числа одновременно, а также использовать 0 или дробные числа, повторите ввод"
			return t
		}
		if a2 > 10 || c2 > 10 {
			t := "Aргументы должны быть кратны или меньше 10, повторите ввод"
			return t
		}
		switch b {
		case "+":
			t = a2 + c2
		case "-":
			t = a2 - c2
		case "*":
			t = a2 * c2
		case "/":
			t = a2 / c2
		default:
			strconv.Itoa(t)
			t := "Не корректный операнд, повторите ввод"
			return t
		}
		return strconv.Itoa(t)
	}
	if t < 0 {
		strconv.Itoa(t)
		t := "Римские цифры не могу быть отрицательными повторите ввод"
		return t
	}

	return romans[t]
}
