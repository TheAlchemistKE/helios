package token

import "testing"

func TestLookupIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"fn", FUNCTION},
		{"let", LET},
		{"if", IF},
		{"else", ELSE},
		{"return", RETURN},
		{"myVariable", IDENT},
		{"x", IDENT},
		{"notAKeyword", IDENT},
	}

	for _, tt := range tests {
		actual := LookupIdent(tt.input)
		if actual != tt.expected {
			t.Errorf("LookupIdent(%q) = %q, want %q",
				tt.input, actual, tt.expected)
		}
	}
}
