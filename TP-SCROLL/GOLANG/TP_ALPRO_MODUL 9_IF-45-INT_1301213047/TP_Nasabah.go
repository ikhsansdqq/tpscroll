package main

import . "fmt"

type customer struct {
	assistantCode, name, bank string
	account                   int
}

type tabCustomer [2022]customer

func main() {
	var (
		X      string
		T      tabCustomer
		N, num int
	)

	N = 0
	num = 0
	for num < 10 {
		addCustomer(&T, N)
		num += 1
		N += 1
	}

	Scanln(&X)
	N = num + 1
	print(T, N, X)
}

func addCustomer(T *tabCustomer, N int) {
	var a customer
	if N < 2022 {
		Scanln(&a.assistantCode, &a.name, &a.bank, &a.account)
		T[N].assistantCode = a.assistantCode
		T[N].name = a.name
		T[N].bank = a.bank
		T[N].account = a.account
	} else {
		Println("Data is Very Full")
	}
}

func print(T tabCustomer, N int, X string) {
	var num int
	for num < N && num < 2022 {
		if T[num].bank == X {
			Println(T[num])
		}
		num += 1
	}
}
