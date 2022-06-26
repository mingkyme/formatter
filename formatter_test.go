package formatter_test

import (
	"testing"

	"github.com/mingkyme/formatter"
)

func TestFormating(t *testing.T) {
	str := formatter.Formating("Hi {@name}", map[string]string{
		"name": "MINGKYME",
	})
	if str != "Hi MINGKYME" {
		t.Error(str, `is not equals "Hi MINGKYME`)
	}
}
