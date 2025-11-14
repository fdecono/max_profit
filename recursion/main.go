// Stock Buy & Sell with Cooldown, to maximize profit:
// 	- After selling, you have to wait at least 1 day before buying again.
//	- You operate with only 1 stock.

package main

import (
	"fmt"
)

func main() {
	prices := []int{0, 2, 1, 2, 3, 7, 1, 10, 0, 3}
	profit := findProfit(0, true, prices)
	fmt.Println(profit)
}

func findProfit(currentDay int, canBuy bool, prices []int) int {
	amountOfDays := len(prices)

	if currentDay >= amountOfDays {
		return 0
	}

	if canBuy {
		buy := -prices[currentDay] + findProfit(currentDay+1, false, prices)
		skip := findProfit(currentDay+1, true, prices)
		return max(buy, skip)
	} else {
		sell := prices[currentDay] + findProfit(currentDay+2, true, prices)
		skip := findProfit(currentDay+1, false, prices)
		return max(sell, skip)
	}
}
