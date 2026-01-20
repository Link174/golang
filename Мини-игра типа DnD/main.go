package main

/* Этот код симулирует битву по правилам игры DnD. Удачи!*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	strong := 3

	def := 14

	fmt.Println()

	fmt.Println("На тебя напал Безумный старец! Его модификатор защиты: 14.")

	time.Sleep(2 * time.Second)

	fmt.Println()

	fmt.Println("Бросаю кубик на атаку...")

	time.Sleep(2 * time.Second)

	fmt.Println()

	a := rand.Intn(21)

	if a == 20 {
		fmt.Println("Критическая удача!")
	} else if a == 1 {
		fmt.Println("Критический провал!")
	} else {
		fmt.Println("Выпало:", a)
	}

	fmt.Println()

	time.Sleep(2 * time.Second)

	fmt.Println("Плюс твой модификатор силы...")

	time.Sleep(2 * time.Second)

	fmt.Println()

	result := a + strong

	fmt.Println("Результат:", result)

	fmt.Println()

	time.Sleep(2 * time.Second)

	if result >= def {
		fmt.Println("Ты попал!")
	} else {
		fmt.Println("Промах!")
	}

	fmt.Println()

}
