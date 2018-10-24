package exectime

import (
	"log"
	"regexp"
	"runtime"
	"time"
)

// ExecTime return time elapsed for a function.
func ExecTime(start time.Time, fullName bool) {
	var fname string
	elapsed := time.Since(start)
	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	if fullName {
		fname = funcObj.Name()
	} else {
		// Regex to extract just the function name (and not the module path).
		runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
		fname = runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	}

	log.Printf("%s took %s", fname, elapsed)
}
