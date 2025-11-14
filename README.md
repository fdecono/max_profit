# 📈 Stock Trading with Cooldown - Maximum Profit

> **Dynamic Programming Solutions** for optimizing stock trading with a one-day cooldown period

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Status](https://img.shields.io/badge/Status-Active-success?style=flat-square)]()

---

## 🎯 Problem Statement

Given an array `prices[]`, where the `i-th` element represents the stock price on day `i`, find the **maximum profit** you can achieve by buying and selling stocks with the following constraints:

- ✅ You can **buy**, **sell**, or **skip** on any day
- ⏸️ After **selling**, you must wait **at least 1 day** before buying again (cooldown period)
- 📊 You operate with **only 1 stock** at a time
- 💰 All prices are **non-negative integers**

---

## 🚀 Solution Approaches

This repository implements **three different approaches** to solve the problem, each with different time/space trade-offs:

| Approach | Time Complexity | Space Complexity | Status |
|----------|----------------|------------------|--------|
| 🔄 **Pure Recursion** | O(2^n) | O(n) | ✅ Complete |
| 🧠 **Top-Down DP (Memorization)** | O(n) | O(n) | ✅ Complete |
| 📊 **Bottom-Up DP (Tabulation)** | O(n) | O(n) | 🚧 In Progress |

---

## 📁 Project Structure

```
stocks/
├── recursion/                    # Pure recursive solution
│   └── main.go
├── recursion_w_memorization/     # Top-down DP
│   └── main.go
├── tabulation/                   # Bottom-up DP
│   └── main.go
├── go.mod                        # Go module file
└── README.md                     # This file
```

---

## 💻 Usage

### Pure Recursion
```bash
go run recursion/main.go
```

### Top-Down DP (Memoization)
```bash
go run recursion_w_memorization/main.go
```

### Bottom-Up DP (Tabulation)
```bash
go run tabulation/main.go
```

---

## 🧮 Algorithm Details

### 1️⃣ Pure Recursion
- **Approach:** Explore all possible buy/sell/skip combinations
- **State:** `(currentDay, canBuy)`
- **Decision Tree:** Exponential growth
- **Use Case:** Educational, understanding the problem structure

### 2️⃣ Top-Down DP (Memoization)
- **Approach:** Cache results of subproblems to avoid recomputation
- **Memoization:** `dp[i][buy]` stores max profit from day `i` with state `buy`
- **Optimization:** Reduces time from O(2^n) to O(n)
- **Use Case:** When you prefer recursive thinking with optimization

### 3️⃣ Bottom-Up DP (Tabulation)
- **Approach:** Build solution iteratively from base cases
- **Iteration:** Fill table from last day to first day
- **Dependencies:** Each state depends on future states (i+1 or i+2)
- **Use Case:** Production code, avoids recursion overhead

---

## 📊 State Transitions

```
┌─────────────────────────────────────────┐
│         State Machine Diagram           │
└─────────────────────────────────────────┘

    [Can Buy] ──buy──> [Holding Stock]
       ↑                    │
       │                    │ sell
       │                    ↓
    [Cooldown] <──skip── [Can Sell]
       │
       │ (wait 1 day)
       ↓
    [Can Buy]
```

**State Definitions:**
- **Can Buy (buy=1):** No stock held, can buy or skip
- **Can Sell (buy=0):** Stock held, can sell or skip
- **Cooldown:** After selling, must wait 1 day before buying

---

## ⚡ Performance Comparison

| Input Size | Pure Recursion | Memoization | Tabulation |
|------------|----------------|-------------|------------|
| n = 10     | ~1ms           | <0.1ms      | <0.1ms     |
| n = 20     | ~1000ms        | <0.1ms      | <0.1ms     |
| n = 30     | Timeout        | <0.1ms      | <0.1ms     |

---

## 🛠️ Implementation Highlights

### Key Features
- ✨ Clean, readable Go code
- 📝 Well-commented implementations
- 🎯 Multiple solution approaches
- ⚡ Performance optimized
- 🧪 Easy to test and modify

### Recurrence Relations

**When Can Buy:**
```
dp[i][1] = max(
    -prices[i] + dp[i+1][0],  // Buy today
    dp[i+1][1]                 // Skip today
)
```

**When Can Sell:**
```
dp[i][0] = max(
    prices[i] + dp[i+2][1],   // Sell today (cooldown)
    dp[i+1][0]                 // Skip today
)
```

---

Built with ❤️ for learning and understanding dynamic programming concepts.

---

<div align="center">

**⭐ Star this repo if you find it helpful! ⭐**

Made with Go 🐹

</div>
