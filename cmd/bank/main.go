package main

import (
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	var cards = []types.Card{

		{
			PAN:     "5058 xxxx xxxx 0000",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 20_000_00,
			Active:  false,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
	}

	result := payment.PaymentSources(cards)
	fmt.Println(result[0].PAN)

}
