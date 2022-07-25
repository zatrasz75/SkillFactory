package calc

import "fmt"

//Сложение
func Add(num1, num2 float32) float32 {
	return num1 + num2
}

//Вычитание
func Subtract(num1, num2 float32) float32 {
	return num1 - num2
}

//Умножение
func Multipy(num1, num2 float32) float32 {
	return num1 * num2
}

//Деление
func Divide(num1, num2 float32) float32 {
	return num1 / num2
}

//Первое число
func ReadIntager1() float32 {
	var number1 float32
	fmt.Print("Введите первое число :")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось считать число: %s"), err)
	}
	return number1
}

//Стока операции
func ReadStr() string {
	var str string
	fmt.Print("Введите операцию +, -, *, /  :")
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось считать : %s"), err)
	}
	return str
}

//Второе число
func ReadIntager2() float32 {
	var number2 float32
	fmt.Print("Введите вторе число :")
	_, err := fmt.Scanln(&number2)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось считать число: %s"), err)
	}
	return number2
}
