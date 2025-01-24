package main

import (
	"fmt"
)

// Loops
// Laços de repetição
// Repetir tarefas

func main() {

	//sum := 0

	// i++ -> soma 1
	// i-- -> subtraindo 1
	// for i := 1; i < 10; i++ {
	// 	fmt.Println("Sum:", sum)
	// 	fmt.Println("Indice:", i)
	// 	sum += i
	// 	// 	// é a mesma coisa que: sum = sum + i
	// 	// 	// sum -= i -> sum = sum - i

	// 	// 	//  ao final do loop
	// 	// 	// e adicionado 1 ao indice i
	// 	// 	// i++
	// 	// 	// i = i + 1
	// }
	// fmt.Println(sum)

	// // loop infinito
	// for {
	// 	fmt.Println("Golang do zero")
	// 	time.Sleep(2 * time.Second)
	// }

	// // for range
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println("Fruta", fruta)
	}
}
