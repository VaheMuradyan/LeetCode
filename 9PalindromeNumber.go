package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	counter := 0
	anqam := 1

	values := []int{}

	for y > 0 {
		y /= 10
		counter++
	}
	counter2 := counter
	for counter > 1 {
		anqam *= 10
		counter--
	}

	for counter2 > 0 {
		values = append(values, x/anqam)
		x %= anqam
		anqam /= 10
		counter2--
	}

	i := 0
	j := len(values) - 1

	for i < j {
		if values[i] != values[j] {
			return false
		}
		i++
		j--
	}

	return true
}
