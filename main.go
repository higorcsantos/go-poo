package main

func PayBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	//execute the functions you want to test
}
