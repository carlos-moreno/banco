package main

import (
	"fmt"

	"github.com/carlos-moreno/banco/contas"

	"github.com/carlos-moreno/banco/clientes"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteCarlos := clientes.Titular{
		Nome:      "Carlos Moreno",
		CPF:       "999.999.999-90",
		Profissao: "Desenvolvedor"}
	contaDoCarlos := contas.ContaCorrente{Titular: clienteCarlos,
		NumeroAgencia: 1432,
		NumeroConta:   223123}

	contaDoCarlos.Depositar(100)

	fmt.Println(contaDoCarlos)
	fmt.Println(contaDoCarlos.ObterSaldo())

	clienteLuiz := clientes.Titular{
		Nome:      "Luíz Gabriel",
		CPF:       "999.999.887-92",
		Profissao: "SRE"}
	contaDoLuiz := contas.ContaPoupanca{Titular: clienteLuiz}
	fmt.Println(contaDoLuiz)

	PagarBoleto(&contaDoCarlos, 40)
	fmt.Println(contaDoCarlos.ObterSaldo())
}
