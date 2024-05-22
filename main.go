package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var x1, x2 float64

	fmt.Println("Реши квадратное уравнение")

	fmt.Println("Введи a")
	fmt.Scan(&a)

	fmt.Println("Введи b")
	fmt.Scan(&b)

	fmt.Println("Введи c")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)
	switch {
	case D < 0:
		fmt.Println("Корней нет")
	case D > 0:
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Ваше уравнение имеет 2 корня: ", x1, x2)
	case D == 0:
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Ваше уравнение имеет 1 корень: ", x)

	}

}
