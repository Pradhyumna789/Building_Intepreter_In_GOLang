package token 

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

// Identifiers and literals
  IDENT = "IDENT" 
  INT = "INT"

  // Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  LT = "<"
  GT = ">"

  // Delimeters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"
)

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}


























