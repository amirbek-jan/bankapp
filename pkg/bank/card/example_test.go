package card

import (
	"bank/pkg/bank/types"	
	"fmt"
)


func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active: false,
		},
	}

	fmt.Println(Total(cards))
	// Output: 100000
}