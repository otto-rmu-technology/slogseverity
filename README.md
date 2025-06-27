# slogseverity

`slogseverity` provides a simple implementation of the `ReplaceAttr`
function for `JSONHandler` from [slog](https://pkg.go.dev/log/slog@master).

This implementation renames the log level attribute into severity to be displayed correctly in the gcp.

## Usage

### Override default attributes

```go
package main

import (
	"github.com/otto-rmu-technology/slogseverity"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: slogseverity.ReplaceAttr,
		Level:       slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}

```
