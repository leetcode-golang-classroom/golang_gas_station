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
