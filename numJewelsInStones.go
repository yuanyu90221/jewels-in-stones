package jewels_stone

import "strings"

func numsJewelsInStone(J string, S string) int {
	count := 0
	for _, r := range S {
		temp_s := string(r)
		if strings.Contains(J, temp_s) {
			count++
		}
	}
	return count
}
