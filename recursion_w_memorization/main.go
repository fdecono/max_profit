// Top Down Dp (Memorization)

package main

import "fmt"

func main() {
	prices := []int{0, 2, 1, 2, 3, 7, 1, 10, 0, 3}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

func findProfit(currentDay int, canBuy int, prices []int, memo [][]int) int {
	amountOfDays := len(prices)

	if currentDay >= amountOfDays {
		return 0
	}

	if memo[currentDay][canBuy] != -1 {
		fmt.Println("memo hit", currentDay, canBuy)
		return memo[currentDay][canBuy]
	}

	if canBuy == 1 {
		buy := -prices[currentDay] + findProfit(currentDay+1, 0, prices, memo)
		skip := findProfit(currentDay+1, 1, prices, memo)
		memo[currentDay][canBuy] = max(buy, skip)
	} else if canBuy == 0 {
		sell := prices[currentDay] + findProfit(currentDay+2, 1, prices, memo)
		skip := findProfit(currentDay+1, 0, prices, memo)
		memo[currentDay][canBuy] = max(sell, skip)
	}

	return memo[currentDay][canBuy]
}

func maxProfit(prices []int) int {
	amountOfDays := len(prices)
	memo := make([][]int, amountOfDays)
	for i := range memo {
		memo[i] = []int{-1, -1}
	}
	return findProfit(0, 1, prices, memo)
}
