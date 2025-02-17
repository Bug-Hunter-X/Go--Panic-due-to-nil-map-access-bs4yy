# Go: Handling nil map access to prevent panics

This repository demonstrates a common error in Go: accessing a map without checking for `nil`.  Accessing a key in a `nil` map results in a runtime panic.

The `bug.go` file shows the problematic code.  The `bugSolution.go` file provides the corrected code, demonstrating safe map access.

This issue is easily avoidable by using a check for `nil` before map access or by using the comma ok idiom.