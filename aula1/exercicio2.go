package main

import (
	"errors"
	"fmt"
)

func getInfo(input string) (int, error) {
	infor := len(input)
	if infor == 0 {
		return 0, errors.New("String vazia")
	}
	return infor, nil
}

func main() {
	valor := "Olá, mundo! Rodrigo"
	tamanho, err := getInfo(valor)
	if err != nil {

		fmt.Println("Erro:", err)
	}
	fmt.Printf("Tamanho da string: %d\n", tamanho)
}
