package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genImports(withCache, timeImport bool, table Table, postgreSql bool) (string, error) {
	if withCache {
		text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
			"time": timeImport,
			"data": table,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}
	var text string
	var err error
	if postgreSql {
		text, err = pathx.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCachePg)
		if err != nil {
			return "", err
		}
	} else {
		text, err = pathx.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
		if err != nil {
			return "", err
		}
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time": timeImport,
		"data": table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
