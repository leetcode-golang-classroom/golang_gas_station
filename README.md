# golang_gas_station

There are `n` gas stations along a circular route, where the amount of gas at the `ith` station is `gas[i]`.

You have a car with an unlimited gas tank and it costs `cost[i]` of gas to travel from the `ith` station to its next `(i + 1)th` station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays `gas` and `cost`, return *the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return* `-1`. If there exists a solution, it is **guaranteed** to be **unique**

## Examples

**Example 1:**

```
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

```

**Example 2:**

```
Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.

```

**Constraints:**

- `n == gas.length == cost.length`
- `1 <= n <= 105`
- `0 <= gas[i], cost[i] <= 104`

## 解析

給定二個整數陣列 gas , cost

gas 中的每個元素值 gas[pos] 代表在該 pos 所會增加的油量

cost 中的每個元素值 cost[pos] 代表在該 pos 開到下一個位置 pos +1 所會消耗的油量

假設這個 pos 是一個環狀的位置

也就是 假設 n = len(gas) 

則 i= 0, … , n-1 個站點

每個 i ≤ n-2 都是往下一站  而 n-1 是往 0 站

也就是每個站 都是往 (pos+1)%n 前進

要求寫一個演算法來計算要從哪一個位置的站出發可以走完全程， 如果不存在則回傳 -1

假設從每個 index = 0, … , n-1 個試著走

因為一開始 tank 是 0

每次 都是 total += gas[index] - cost[index]

當發現 total ≤ 0 代表不可能走完

這樣走的話

每個出發點有 n 個 每次最多走 n 次

所以這樣時間複雜度是 O($n^2$) 

要透過 Greedy algorithm 必須對能夠走完的條件做觀察

如果要走完 必須要 sum(gas) - sum(cost) ≥ 0 

所以如果 sum(gas) - sum(cost) < 0 代表一定無法走完回傳 -1

第2件事要找出可以出發的點 

先初始化出發點是 start =0, 然後延申出發點 i=0 開始往後累加 total 

每次累加 total += gas[i] - cost[i]

當發現 total < 0

代表該點不是出發點 所以把出發點往後 start = i + 1 , 並且重新累計 total = 0

因為整體假設是可行的代表在某個點其所增益的值 可以補償其他點的損失

![](https://i.imgur.com/Onk71ac.png)

## 程式碼
```go
package sol

func canCompleteCircuit(gas []int, cost []int) int {
	nGas := len(gas)
	remain := 0
	total := 0
	start := 0
	for pos := 0; pos < nGas; pos++ {
		total += gas[pos] - cost[pos]
		if total < 0 {
			remain += total
			total = 0
			start = pos + 1
		}
	}
	if remain+total >= 0 {
		return start
	}
	return -1
}

```
## 困難點

1. 要透過 Greedy Algorithm 必須要理解 從某個 start 開始之後與其 remain 之間的關係

## Solve Point

- [x]  初始化 total = 0, remain = 0, start = 0
- [x]  每次計算 total += gas[pos] - cost[pos] 如果 total  < 0, start = pos +1,  更新 remain += total, total = 0
- [x]  判斷 remain + total 是否 ≥ 0 如果是 則回傳 start， 否則回傳 -1