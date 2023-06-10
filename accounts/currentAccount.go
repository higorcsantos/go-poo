package accounts

import "github.com/alura/clients"

type CurrentAccount struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(value float64) string { //aponta para a conta que está sendo chamada
	canWithdraw := value <= c.balance && value > 0
	if canWithdraw {
		c.balance -= value
		return "Saque realizado com sucesso"
	}
	return "Saldo Insuficiente"
}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Depósito realizado com sucesso", c.balance
	}
	return "Valor menor do que 0", c.balance
}

func (c *CurrentAccount) Transferir(valorDaTransferencia float64, contaDestino CurrentAccount) bool {
	if valorDaTransferencia < c.balance {
		c.balance -= valorDaTransferencia
		contaDestino.Deposit(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
