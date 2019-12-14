package golemur

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"testing"
)

func TestStandardLib(t *testing.T) {
	var s StandardLibConstantFolder
	input := smallExpression()
	ast.Walk(s, input)
	fmt.Println(printAst(input))
}

var _ ast.Visitor = (*StandardLibConstantFolder)(nil)

type StandardLibConstantFolder struct {}

func (s StandardLibConstantFolder) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.BinaryExpr:
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

			node = &lit
			fmt.Println(printAst(node))
			return nil // don't descend more
		}

	}
	return s
}
