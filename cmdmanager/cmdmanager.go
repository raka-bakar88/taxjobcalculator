package cmdmanager

import "fmt"

type CmdManager struct {
}

func New() CmdManager {
	return CmdManager{}
}

func (cmd CmdManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}
