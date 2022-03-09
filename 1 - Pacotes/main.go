package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo arquivo main")
	auxiliar.Escrevendo()
	erroNil := checkmail.ValidateFormat("guilherme.castro.fernandes@gmail.com")
	fmt.Println("Exemplo de email válido (Deve retornar <nil>): ", erroNil)
	fmt.Println("####################################")
	erroNotNil := checkmail.ValidateFormat("123")
	fmt.Println("Exemplo de email inválido (Deve retornar 'invalid format')", erroNotNil)

}
