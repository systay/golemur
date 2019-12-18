package golemur

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strconv"
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

func rewrite(n *ast.BinaryExpr) ast.Node {
	xLit, xIsLit := n.X.(*ast.BasicLit)
	yLit, yIsLit := n.Y.(*ast.BasicLit)
	// left to the reader:
	//  - check types. we're pretending it's always integers here
	//  - implement other operators. here we're only doing `+`
	if xIsLit && yIsLit {
		// how do we replace the current node!?
		xval, _ := strconv.Atoi(xLit.Value)
		yval, _ := strconv.Atoi(yLit.Value)
		lit := ast.BasicLit{
			ValuePos: xLit.ValuePos,
			Kind:     token.INT,
			Value:    strconv.FormatInt(int64(xval+yval), 10),
		}

		return &lit
	}

	return n
}