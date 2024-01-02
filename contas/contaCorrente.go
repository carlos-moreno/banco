package contas

import "github.com/carlos-moreno/banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor para deposito inválido", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
