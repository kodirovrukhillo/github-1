package main

import (
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	var cards = []types.Card{

		{
			Number:  "5058 xxxx xxxx 0001",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Number:  "5058 xxxx xxxx 1010",
			Balance: 20_000_00,
			Active:  false,
		},
		{
			Number:  "5058 xxxx xxxx 8888",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Number:  "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
	}

	result := payment.PaymentSources(cards)

	fmt.Println(result)

}
