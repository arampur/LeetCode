package main

func rearrangeCharacters(s string, target string) int {
	tMap := make(map[rune]int)
	sMap := make(map[rune]int)

	for _, val := range target {
		tMap[val]++
	}

	for _, val := range s {
		sMap[val]++
	}

	numOfCopies := 0
	curCount := 0
	for i := 0; i < len(s); i++ {
		if val, ok := sMap[rune(s[i])]; ok {
			curCount++
			val -= 1
			tMap[rune(s[i])] = val
		} else {
			if curCount > 0 {
				curCount = 0
			}
		}

		if curCount == len(target) {
			numOfCopies++
		}
	}
	return numOfCopies
}
