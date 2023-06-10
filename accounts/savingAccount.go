package accounts

import "github.com/alura/clients"

type SavingAccount struct {
	Holder        clients.Holder
	AgencyNumber  int
	AccountNumber int
	Operation     int
	balance       float64
}

func (c *SavingAccount) Withdraw(value float64) string { //aponta para a conta que está sendo chamada
	canWithdraw := value <= c.balance && value > 0
	if canWithdraw {
		c.balance -= value
		return "Saque realizado com sucesso"
	}
	return "Saldo Insuficiente"
}

func (c *SavingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Depósito realizado com sucesso", c.balance
	}
	return "Valor menor do que 0", c.balance
}

func (c *SavingAccount) Transferir(valorDaTransferencia float64, contaDestino SavingAccount) bool {
	if valorDaTransferencia < c.balance {
		c.balance -= valorDaTransferencia
		contaDestino.Deposit(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
