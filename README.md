# Exec Time

Print execution time of a function.

## Usage

```golang
package main

import (
	"time"

	"github.com/proximisfr/exec-time"
)

func main() {
	defer exectime.ExecTime(time.Now(), false)
	test()
	testFullName()
	time.Sleep(3000 * time.Millisecond)
}

func test() {
	defer exectime.ExecTime(time.Now(), false)

	time.Sleep(1000 * time.Millisecond)
}

func testFullName() {
	defer exectime.ExecTime(time.Now(), true)

	time.Sleep(1000 * time.Millisecond)
}

```

```shell
$#: go run main.go
2018/08/08 16:57:42 test took 1.000100813s
2018/08/08 16:57:43 main.testFullName took 1.00012207s
2018/08/08 16:57:46 main took 5.000898751s
```
