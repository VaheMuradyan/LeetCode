package main

func lengthOfLongestSubstring(s string) int {
	sls := []string{}
	maxValue := 0

	for i := 0; i < len(s); i++ {
		val := string(s[i])

		if contains(sls, val) {
			for len(sls) > 0 && sls[0] != val {
				sls = sls[1:]
			}
			sls = sls[1:]
		}

		sls = append(sls, val)

		if len(sls) > maxValue {
			maxValue = len(sls)
		}
	}

	return maxValue
}

func contains(arr []string, val string) bool {
	for _, i := range arr {
		if i == val {
			return true
		}
	}
	return false
}
