package formatter

import (
	"fmt"
	"strings"
)

func Formating(str string, data map[string]string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	for _, key := range keys {
		str = strings.ReplaceAll(str, fmt.Sprintf(`{@%s}`, key), data[key])
	}
	return str
}
