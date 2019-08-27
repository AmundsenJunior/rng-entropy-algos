package main

import (
	"fmt"
	"time"
)

// Run an infinite loop of calls to TimeNowDelta(), printing to stdout until
// program receives a kill signal (Ctrl+C).
func main() {
	for {
		fmt.Printf(TimeNowDelta())
	}
}

// Use the variability in the execution time of reading time.Now() twice as an
// entropy source to produce a difference value that is returned as a Unicode
// character.
//
// Algorithm is based on the Memory Access noise source defined in
// http://www.chronox.de/jent/doc/CPU-Jitter-NPTRNG.html (sections 3.1.1 and
// 3.1.2).
//
// TODO: Provide an alternative function that produces noise by including
//   memory access calls.
func TimeNowDelta() string {
	before := time.Now().UnixNano()
	after := time.Now().UnixNano()

	return fmt.Sprintf("%c", after - before)
}
