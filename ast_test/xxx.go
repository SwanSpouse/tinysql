package main

import (
	"fmt"
	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	driver "github.com/pingcap/tidb/types/parser_driver"
)

type visitor struct{}

func (v *visitor) Enter(in ast.Node) (out ast.Node, skipChildren bool) {
	fmt.Printf("%s %T\n", in.Text(), in)

	switch cur := in.(type) {
	case *ast.SelectStmt:
		fmt.Printf("--->cur:%+v\n", cur)
	case *ast.TableOptimizerHint:
		fmt.Printf("--->tableOptimizierHint:%+v\n", cur)
	case *ast.TableRefsClause:
		fmt.Printf("--->TableRefsClause:%+v\n", cur)
	case *ast.Join:
		fmt.Printf("--->Join:%+v\n", cur)
	case *ast.TableSource:
		fmt.Printf("--->TableSource:%+v\n", cur)
	case *ast.TableName:
		fmt.Printf("--->TableName:%+v\n", cur)
	case *ast.BinaryOperationExpr:
		fmt.Printf("--->BinaryOperationExpr:%+v\n", cur)
	case *ast.ColumnNameExpr:
		fmt.Printf("--->ColumnNameExpr:%+v\n", cur)
	case *ast.ColumnName:
		fmt.Printf("--->ColumnNameExpr:%+v\n", cur)
	case *driver.ValueExpr:
		fmt.Printf("--->ValueExpr:%+v\n", cur)
	case *ast.FieldList:
		fmt.Printf("--->FieldList:%+v\n", cur)
	case *ast.SelectField:
		fmt.Printf("--->SelectField:%+v\n", cur)
	}

	return in, false
}

func (v *visitor) Leave(in ast.Node) (out ast.Node, ok bool) {
	return in, true
}

func main() {
	sql := "SELECT /*+ TIDB_SMJ(employees) */ emp_no, first_name, last_name " +
		"FROM employees USE INDEX (last_name) " +
		"where last_name='Aamodt' and gender='F' and birth_date > '1960-01-01'"

	sqlParser := parser.New()
	stmtNodes, _, err := sqlParser.Parse(sql, "", "")
	if err != nil {
		fmt.Printf("parse error:\n%v\n%s", err, sql)
		return
	}

	for _, stmtNode := range stmtNodes {
		v := visitor{}
		stmtNode.Accept(&v)
	}
}
