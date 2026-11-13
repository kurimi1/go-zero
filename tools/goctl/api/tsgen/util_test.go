package tsgen

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
			expected: "\tname: string // name in path\n",
		},
		{
			name: "multiline block comment",
			comment: "/* first\r\n}\nfunction injected() {}\rvar escaped = true\u2028" +
				"const executed = true\u2029*/",
			expected: "\tname: string // /* first } function injected() {} " +
				"var escaped = true const executed = true */\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var builder strings.Builder
			err := writeProperty(&builder, spec.Member{
				Name:    "Name",
				Type:    spec.PrimitiveType{RawName: "string"},
				Tag:     "`json:\"name\"`",
				Comment: test.comment,
			}, 1)

			assert.Nil(t, err)
			assert.Equal(t, test.expected, builder.String())
		})
	}
}
