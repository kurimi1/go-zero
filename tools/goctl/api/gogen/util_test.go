package gogen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func TestWritePropertyComment(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		expected string
	}{
		{
			name:     "line comment",
			comment:  "// name in path",
			expected: "\tName string `json:\"name\"` // name in path\n",
		},
		{
			name:    "multiline block comment",
			comment: "/* first\r\n}\nfunc init() {}\rvar injected = true\u2028const escaped = true\u2029*/",
			expected: "\tName string `json:\"name\"` ///* first } func init() {} " +
				"var injected = true const escaped = true */\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var builder strings.Builder
			err := writeProperty(&builder, "name", "`json:\"name\"`", test.comment,
				spec.PrimitiveType{RawName: "string"}, 1)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, builder.String())
		})
	}
}
