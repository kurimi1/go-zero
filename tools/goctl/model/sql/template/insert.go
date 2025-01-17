package template

// Insert defines a template for insert code in model
var Insert = `
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{if .containsIndexCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.ExecNoCacheCtx(ctx, query, {{.expressionValues}})
	{{end}}{{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}
`

// pg
var InsertPg = `
func (m *default{{.upperStartCamelObject}}Model) Insert(data *{{.upperStartCamelObject}}) (int64,error) {
	{{if .withCache}}{{if .containsIndexCache}}{{.keys}}
    ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.Exec(query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.ExecNoCache(query, {{.expressionValues}})
	{{end}}{{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}}) RETURNING id", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	var id int64
    err:=m.conn.QueryRow(&id, query, {{.expressionValues}}){{end}}
	return id,err
}
`

// InsertMethod defines an interface method template for insert code in model
var InsertMethod = `Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error)`

// pg
var InsertMethodPg = `Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (int64,error)`
