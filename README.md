# Simple diff
Simplified `diff` tool for computing line-by-line differences between two files, ignoring order.

## Command-line Usage
You can run `simple-diff` by using `go run` and supplying the paths to the file to analyse:
```bash
go run cmd/simple-diff/main.go path-to-file-1 path-to-file-2
< this line is in file 1, but not in file 2
> this line is in file 2, but not in file 1
```

Alternatively you can compile this tool by running:
```bash
go build -o simple-diff cms/simple-diff/main.go
```

Providing a single file will consider the first file empty:
```bash
go run cmd/simple-diff/main.go path-to-file
< this line is in file
```
which will consider every line in the file as added.

## Inclusion as a Dependency
To use this tool as part of a Go project you can import this package:
```go
package main

import (
    "fmt"

	"github.com/marco-zanella/simple-diff/diff"
)

func main() {
    result, err := diff.DiffFiles("path-to-file-1", "path-to-file-2")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, line := range result.Removed {
		fmt.Printf("This line was in file 1, but not in file 2: %s\n", line)
	}

	for _, line := range result.Added {
		fmt.Printf("This line was in file2, but not in file 1: %s\n", line)
	}
}
```
Differences will be availabed in `result.Removed` and `result.Added`, representing lines removed from the first file and added to the second one, respectively.

As with the command-line tool, providing `""` as the first file will be considered as an empty file, and every line in the second one will be considered as added.