# How To Use
```bash
go get github.com/mingkyme/formatter
```

```go
import "github.com/mingkyme/formatter"

str := formatter.Formating("Hi {@name}", map[string]string{
  "name": "MINGKYME",
}) // Hi MINGKYME
```