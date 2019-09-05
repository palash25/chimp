package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "INDENT" // indentifiers
	INT   = "INT"    // literals

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	SUB    = "-"
	MUL    = "*"
	DIV    = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywordMap = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(identifier string) TokenType {
	if tokenType, ok := keywordMap[identifier]; ok {
		return tokenType
	}
	return IDENT
}
