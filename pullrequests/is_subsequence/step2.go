//lint:file-ignore U1000 Ignore all unused code
package issubsequence

/*
よりシンプルにした。
*/
func isSubsequenceStep2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	current := 0
	for i := 0; i < len(t); i++ {
		if s[current] == t[i] {
			current++
		}
		if current == len(s) {
			return true
		}
	}
	return false
}
