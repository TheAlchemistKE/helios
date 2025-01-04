package ast

import (
	"testing"

	"github.com/TheAlchemistKE/helios/internal/token"
)

func TestLetStatementString(t *testing.T) {
	// Creating a mock token for LET
	letToken := token.Token{Type: token.LET, Literal: "let"}

	// Creating the mock identifier and expression
	name := &Identifier{Token: token.Token{Literal: "x"}, Value: "x"}
	value := &IntegerLiteral{Token: token.Token{Literal: "5"}, Value: 5}

	// Create the LetStatement
	letStmt := &LetStatement{Token: letToken, Name: name, Value: value}

	// Check if the string representation is as expected
	expected := "let x = 5;"
	if letStmt.String() != expected {
		t.Errorf("expected %s, got %s", expected, letStmt.String())
	}
}

func TestReturnStatementString(t *testing.T) {
	// Creating mock tokens
	returnToken := token.Token{Type: token.RETURN, Literal: "return"}
	returnValue := &IntegerLiteral{Token: token.Token{Literal: "10"}, Value: 10}

	// Create ReturnStatement
	returnStmt := &ReturnStatement{Token: returnToken, ReturnValue: returnValue}

	// Check string representation
	expected := "return 10;"
	if returnStmt.String() != expected {
		t.Errorf("expected %s, got %s", expected, returnStmt.String())
	}
}

func TestInfixExpressionString(t *testing.T) {
	// Create tokens and mock expressions
	left := &IntegerLiteral{Token: token.Token{Literal: "5"}, Value: 5}
	operator := "+"
	right := &IntegerLiteral{Token: token.Token{Literal: "10"}, Value: 10}

	// Create the InfixExpression
	infixExpr := &InfixExpression{
		Token:    token.Token{Literal: operator},
		Left:     left,
		Operator: operator,
		Right:    right,
	}

	// Expected string representation
	expected := "(5 + 10)"
	if infixExpr.String() != expected {
		t.Errorf("expected %s, got %s", expected, infixExpr.String())
	}
}

func TestBooleanString(t *testing.T) {
	// Create a mock boolean token
	booleanToken := token.Token{Literal: "true"}
	boolean := &Boolean{Token: booleanToken, Value: true}

	// Check string representation
	expected := "true"
	if boolean.String() != expected {
		t.Errorf("expected %s, got %s", expected, boolean.String())
	}
}

func TestIfExpressionString(t *testing.T) {
	// Creating mock tokens
	ifToken := token.Token{Literal: "if"}
	condition := &Boolean{Token: token.Token{Literal: "true"}, Value: true}
	consequence := &BlockStatement{Token: token.Token{Literal: "{"}, Statements: []Statement{}}

	// Creating the IfExpression
	ifExpr := &IfExpression{
		Token:       ifToken,
		Condition:   condition,
		Consequence: consequence,
	}

	// Check string representation
	expected := "if true {  }"
	if ifExpr.String() != expected {
		t.Errorf("expected %s, got %s", expected, ifExpr.String())
	}
}

func TestFunctionLiteralString(t *testing.T) {
	// Mock tokens and params
	fnToken := token.Token{Literal: "fn"}
	params := []*Identifier{
		&Identifier{Token: token.Token{Literal: "x"}, Value: "x"},
		&Identifier{Token: token.Token{Literal: "y"}, Value: "y"},
	}
	body := &BlockStatement{
		Token:      token.Token{Literal: "{"},
		Statements: []Statement{},
	}

	// Create the FunctionLiteral
	fnLit := &FunctionLiteral{
		Token:      fnToken,
		Parameters: params,
		Body:       body,
	}

	// Check the expected string output
	expected := "fn(x, y) {}"
	if fnLit.String() != expected {
		t.Errorf("expected %s, got %s", expected, fnLit.String())
	}
}

func TestArrayLiteralString(t *testing.T) {
	// Mock tokens
	arrayToken := token.Token{Literal: "["}
	elements := []Expression{
		&IntegerLiteral{Token: token.Token{Literal: "1"}, Value: 1},
		&IntegerLiteral{Token: token.Token{Literal: "2"}, Value: 2},
	}

	// Create the ArrayLiteral
	arrayLiteral := &ArrayLiteral{
		Token:    arrayToken,
		Elements: elements,
	}

	// Check the string output
	expected := "[1, 2]"
	if arrayLiteral.String() != expected {
		t.Errorf("expected %s, got %s", expected, arrayLiteral.String())
	}
}

func TestNullLiteralString(t *testing.T) {
	// Mock token
	nullToken := token.Token{Literal: "null"}

	// Create NullLiteral
	nullLiteral := &NullLiteral{Token: nullToken}

	// Check the string output
	expected := "null"
	if nullLiteral.String() != expected {
		t.Errorf("expected %s, got %s", expected, nullLiteral.String())
	}
}

func TestForExpressionString(t *testing.T) {
	// Mock tokens and expressions
	forToken := token.Token{Literal: "for"}
	identifier := &Identifier{Token: token.Token{Literal: "i"}, Value: "i"}
	iterator := &IntegerLiteral{Token: token.Token{Literal: "10"}, Value: 10}
	body := &BlockStatement{
		Token:      token.Token{Literal: "{"},
		Statements: []Statement{},
	}

	// Create the ForExpression
	forExpr := &ForExpression{
		Token:      forToken,
		Identifier: identifier,
		Iterator:   iterator,
		Body:       body,
	}

	// Check the string output
	expected := "for i in 10 {}"
	if forExpr.String() != expected {
		t.Errorf("expected %s, got %s", expected, forExpr.String())
	}
}

func TestProgramWithMultipleStatements(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Literal: "let"},
				Name:  &Identifier{Value: "x"},
				Value: &IntegerLiteral{Token: token.Token{Literal: "5"}, Value: 5},
			},
			&ReturnStatement{
				Token:       token.Token{Literal: "return"},
				ReturnValue: &Identifier{Token: token.Token{Literal: "x"}, Value: "x"},
			},
		},
	}
	expected := "let x = 5; return x;"
	if program.String() != expected {
		t.Errorf("expected %v, got %v", expected, program.String())
	}
}

func TestLetStatementWithComplexExpression(t *testing.T) {
	letStmt := &LetStatement{
		Token: token.Token{Literal: "let"},
		Name:  &Identifier{Value: "result"},
		Value: &InfixExpression{
			Token:    token.Token{Literal: "+"},
			Left:     &IntegerLiteral{Token: token.Token{Literal: "5"}, Value: 5},
			Operator: "+",
			Right:    &IntegerLiteral{Token: token.Token{Literal: "10"}, Value: 10},
		},
	}
	expected := "let result = (5 + 10);"
	if letStmt.String() != expected {
		t.Errorf("expected %v, got %v", expected, letStmt.String())
	}
}

func TestReturnStatementWithInfixExpression(t *testing.T) {
	returnStmt := &ReturnStatement{
		Token: token.Token{Literal: "return"},
		ReturnValue: &InfixExpression{
			Token:    token.Token{Literal: "+"},
			Left:     &IntegerLiteral{Token: token.Token{Literal: "1"}, Value: 1},
			Operator: "+",
			Right:    &IntegerLiteral{Token: token.Token{Literal: "2"}, Value: 2},
		},
	}
	expected := "return (1 + 2);"
	if returnStmt.String() != expected {
		t.Errorf("expected %v, got %v", expected, returnStmt.String())
	}
}

func TestPrefixExpression(t *testing.T) {
	prefixExpr := &PrefixExpression{
		Token:    token.Token{Literal: "!"},
		Operator: "!",
		Right:    &Boolean{Token: token.Token{Literal: "true"}, Value: true},
	}
	expected := "(!true)"
	if prefixExpr.String() != expected {
		t.Errorf("expected %v, got %v", expected, prefixExpr.String())
	}
}

func TestIfExpressionWithConsequenceAndAlternative(t *testing.T) {
	ifExpr := &IfExpression{
		Token:     token.Token{Literal: "if"},
		Condition: &Boolean{Token: token.Token{Literal: "true"}, Value: true},
		Consequence: &BlockStatement{
			Token: token.Token{Literal: "{"},
			Statements: []Statement{
				&ReturnStatement{
					Token: token.Token{Literal: "return"},
					ReturnValue: &IntegerLiteral{
						Token: token.Token{Literal: "1"},
						Value: 1,
					},
				},
			},
		},
		Alternative: &BlockStatement{
			Token: token.Token{Literal: "{"},
			Statements: []Statement{
				&ReturnStatement{
					Token: token.Token{Literal: "return"},
					ReturnValue: &IntegerLiteral{
						Token: token.Token{Literal: "2"},
						Value: 2,
					},
				},
			},
		},
	}
	expected := "if true { return 1; } else { return 2; }"
	if ifExpr.String() != expected {
		t.Errorf("expected %v, got %v", expected, ifExpr.String())
	}
}

func TestForExpression(t *testing.T) {
	forExpr := &ForExpression{
		Token:      token.Token{Literal: "for"},
		Identifier: &Identifier{Value: "i"},
		Iterator:   &IntegerLiteral{Token: token.Token{Literal: "10"}, Value: 10},
		Body: &BlockStatement{
			Token: token.Token{Literal: "{"},
			Statements: []Statement{
				&ExpressionStatement{
					Token: token.Token{Literal: "print"},
					Expression: &CallExpression{
						Token:     token.Token{Literal: "("},
						Function:  &Identifier{Value: "print"},
						Arguments: []Expression{&Identifier{Value: "i"}},
					},
				},
			},
		},
	}
	expected := "for i in 10 { print(i); }"
	if forExpr.String() != expected {
		t.Errorf("expected %v, got %v", expected, forExpr.String())
	}
}

func TestArrayLiteral(t *testing.T) {
	arrayLit := &ArrayLiteral{
		Token: token.Token{Literal: "["},
		Elements: []Expression{
			&IntegerLiteral{Token: token.Token{Literal: "1"}, Value: 1},
			&IntegerLiteral{Token: token.Token{Literal: "2"}, Value: 2},
			&IntegerLiteral{Token: token.Token{Literal: "3"}, Value: 3},
		},
	}
	expected := "[1, 2, 3]"
	if arrayLit.String() != expected {
		t.Errorf("expected %v, got %v", expected, arrayLit.String())
	}
}

func TestHashLiteral(t *testing.T) {
	hashLit := &HashLiteral{
		Token: token.Token{Literal: "{"},
		Pairs: map[Expression]Expression{
			//TODO: Make the line below work.
			//&StringLiteral{Token: token.Token{Literal: "name"}, Value: "Alice"}: &StringLiteral{Token: token.Token{Literal: "value"}, Value: "123"},
			&StringLiteral{Token: token.Token{Literal: "age"}, Value: "age"}: &IntegerLiteral{Token: token.Token{Literal: "30"}, Value: 30},
		},
	}
	expected := "{age:30}"
	if hashLit.String() != expected {
		t.Errorf("expected %v, got %v", expected, hashLit.String())
	}
}

func TestNullLiteral(t *testing.T) {
	nullLit := &NullLiteral{Token: token.Token{Literal: "null"}}
	expected := "null"
	if nullLit.String() != expected {
		t.Errorf("expected %v, got %v", expected, nullLit.String())
	}
}

func TestTernaryExpression(t *testing.T) {
	ternaryExpr := &TernaryExpression{
		Token:       token.Token{Literal: "?"},
		Condition:   &Boolean{Token: token.Token{Literal: "true"}, Value: true},
		TrueBranch:  &IntegerLiteral{Token: token.Token{Literal: "1"}, Value: 1},
		FalseBranch: &IntegerLiteral{Token: token.Token{Literal: "0"}, Value: 0},
	}
	expected := "(true ? 1 : 0)"
	if ternaryExpr.String() != expected {
		t.Errorf("expected %v, got %v", expected, ternaryExpr.String())
	}
}

func TestWhileExpression(t *testing.T) {
	whileExpr := &WhileExpression{
		Token: token.Token{Literal: "while"},
		Condition: &Boolean{
			Token: token.Token{Literal: "true"},
			Value: true,
		},
		Body: &BlockStatement{
			Token: token.Token{Literal: "{"},
			Statements: []Statement{
				&ReturnStatement{
					Token: token.Token{Literal: "return"},
					ReturnValue: &IntegerLiteral{
						Token: token.Token{Literal: "1"},
						Value: 1,
					},
				},
			},
		},
	}
	expected := "while true { return 1; }"
	if whileExpr.String() != expected {
		t.Errorf("expected %v, got %v", expected, whileExpr.String())
	}
}
