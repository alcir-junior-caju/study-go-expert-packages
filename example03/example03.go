package example03

import "fmt"

func ShowDefer() {
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
