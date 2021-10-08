package main

import (
	"fmt"
	"today-i-learned/bankaccount"
)

func main() {
	account := bankaccount.NewAccount() // ❶ 계좌 생성
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
