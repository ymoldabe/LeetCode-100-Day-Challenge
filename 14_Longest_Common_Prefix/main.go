package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		curr := strs[i]
		j := 0
		for j < len(pref) && j < len(curr) && pref[j] == curr[j] {
			j++
		}
		pref = pref[:j]
		if pref == "" {
			return ""
		}
	}
	return pref
}
