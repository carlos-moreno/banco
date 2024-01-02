package main

import (
	"fmt"

	"github.com/carlos-moreno/banco/contas"

	"github.com/carlos-moreno/banco/clientes"
)

func main() {

	clienteCarlos := clientes.Titular{
		Nome:      "Carlos Moreno",
		CPF:       "999.999.123-90",
		Profissao: "Desenvolvedor"}
	contaDoCarlos := contas.ContaCorrente{Titular: clienteCarlos,
		NumeroAgencia: 1432,
		NumeroConta:   223123,
		Saldo:         10500.50}
	fmt.Println(contaDoCarlos)

}
