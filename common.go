package golemur

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

func smallExpression() ast.Node {
	var lit1 ast.Expr = &ast.BasicLit{
		ValuePos: token.Pos(0),
		Kind:     token.INT,
		Value:    "1",
	}
	var lit2 ast.Expr = &ast.BasicLit{
		ValuePos: token.Pos(0),
		Kind:     token.INT,
		Value:    "1",
	}
	return &ast.BinaryExpr{
		X:     lit1,
		OpPos: token.Pos(0),
		Op:    token.ADD,
		Y:     lit2,
	}
}

func printAst(n ast.Node) string {
	var buf bytes.Buffer
	fset := token.NewFileSet()
	err := printer.Fprint(&buf, fset, n)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
