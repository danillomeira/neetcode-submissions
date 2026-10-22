func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)

	for _, letra := range s {
		counts[letra]++
	}
	for _, letra := range t {
		counts[letra]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}
