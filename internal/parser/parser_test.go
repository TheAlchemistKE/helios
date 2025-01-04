package parser

import (
	"testing"

	"github.com/TheAlchemistKE/helios/internal/ast"
	"github.com/TheAlchemistKE/helios/internal/lexer"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "let x = 5;",
			expected: "let x = 5;",
		},
		{
			input:    "return x;",
			expected: "return x;",
		},
		{
			input:    "x + y;",
			expected: "(x + y)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		// Normalize the string output (e.g., trimming spaces)
		actual := program.String()

		// Trim spaces or newline characters from the start and end of the strings before comparing
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParseLetStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "let x = 5;",
			expected: "let x = 5;",
		},
		{
			input:    "let y = true;",
			expected: "let y = true;",
		},
		{
			input:    "let z = x + y;",
			expected: "let z = (x + y);",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt := p.parseLetStatement()

		// Assuming LetStatement has a String() method
		actual := stmt.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParseReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "return 5;",
			expected: "return 5;",
		},
		{
			input:    "return x + y;",
			expected: "return (x + y);",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt := p.parseReturnStatement()

		// Assuming ReturnStatement has a String() method
		actual := stmt.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParseExpressionStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "x + y;",
			expected: "(x + y)",
		},
		{
			input:    "true;",
			expected: "true",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		stmt := p.parseExpressionStatement()

		// Assuming ExpressionStatement has a String() method
		actual := stmt.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParseIntegerLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{
			input:    "5",
			expected: 5,
		},
		{
			input:    "123",
			expected: 123,
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr := p.parseIntegerLiteral()

		// Assuming IntegerLiteral has a Value field
		actual := expr.(*ast.IntegerLiteral).Value
		if actual != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, actual)
		}
	}
}

func TestParseStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    `"hello"`,
			expected: "hello",
		},
		{
			input:    `"world"`,
			expected: "world",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr := p.parseStringLiteral()

		// Assuming StringLiteral has a Value field
		actual := expr.(*ast.StringLiteral).Value
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParseInfixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "x + y",
			expected: "(x + y)",
		},
		{
			input:    "x - y",
			expected: "(x - y)",
		},
		{
			input:    "x * y",
			expected: "(x * y)",
		},
		{
			input:    "x / y",
			expected: "(x / y)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr := p.parseExpression(LOWEST)

		// Assuming InfixExpression has a String() method for easier comparison
		actual := expr.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}

func TestParsePrefixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "-x",
			expected: "(-x)",
		},
		{
			input:    "!x",
			expected: "(!x)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		expr := p.parsePrefixExpression()

		// Assuming PrefixExpression has a String() method
		actual := expr.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
		}
	}
}
