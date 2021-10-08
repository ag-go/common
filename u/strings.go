package u

// NormalizeNewlinesInPlace changes CRLF (Windows) and
// CR (Mac) to LF (Unix)
// Optimized for speed, modifies data in place
func NormalizeNewlinesInPlace(d []byte) []byte {
	wi := 0
	n := len(d)
	for i := 0; i < n; i++ {
		c := d[i]
		// 13 is CR
		if c != 13 {
			d[wi] = c
			wi++
			continue
		}
		// replace CR (mac / win) with LF (unix)
		d[wi] = 10
		wi++
		if i < n-1 && d[i+1] == 10 {
			// this was CRLF, so skip the LF
			i++
		}

	}
	return d[:wi]
}

// NormalizeNewlines is like NormalizeNewlinesInPlace but
// slower because it makes a copy of data
func NormalizeNewlines(d []byte) []byte {
	d = append([]byte{}, d...)
	return NormalizeNewlinesInPlace(d)
}
