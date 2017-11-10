# example
```go
package main

import (
    "sync"

    "github.com/parkjy76/throttle"
)

func main() {
    t := throttle.NewThrottle(100)
    defer t.Close()

    wg := new(sync.WaitGroup)

    // omit testChan's definition code
    for _ := range testChan {
        t.Increase()
        wg.Add(1)
        go func() {
            something()
            wg.Done()
            t.Decrease()
        }()
    }

    wg.Wait()
}
```
