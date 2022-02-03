package lexer

import (
	"fmt"
	"strings"

	"github.com/zhangzheheng12345/FlowScript/tools"
)

// TODO: Support string and list syntax and more
/* Core lexer code */
func Lex(str []string) []Token {
	result := make([]Token, 0)
	for index := 0; index < len(str); index++ {
		if str[index] == " " || str[index] == "\t" {
			// Do nothing
		} else if str[index] == "#" {
			for ; index < len(str) && str[index] != "\n"; index++ {
				// Do nothing
			}
			index--
		} else if tools.IsSingleAlpha(str[index]) || str[index] == "_" {
			var word string
			word, index = tools.PickSymbol(str, index)
			if word == "add" {
				result = append(result, Token{ADD, ""})
			} else if word == "sub" {
				result = append(result, Token{SUB, ""})
			} else if word == "multi" {
				result = append(result, Token{MULTI, ""})
			} else if word == "div" {
				result = append(result, Token{DIV, ""})
			} else if word == "mod" {
				result = append(result, Token{MOD, ""})
			} else if word == "equal" {
				result = append(result, Token{EQUAL, ""})
			} else if word == "bigr" {
				result = append(result, Token{BIGR, ""})
			} else if word == "smlr" {
				result = append(result, Token{SMLR, ""})
			} else if word == "not" {
				result = append(result, Token{NOT, ""})
			} else if word == "and" {
				result = append(result, Token{AND, ""})
			} else if word == "or" {
				result = append(result, Token{OR, ""})
			} else if word == "var" {
				result = append(result, Token{VAR, ""})
			} else if word == "if" {
				result = append(result, Token{IF, ""})
			} else if word == "def" {
				result = append(result, Token{DEF, ""})
			} else if word == "begin" {
				result = append(result, Token{BEGIN, ""})
			} else if word == "end" {
				result = append(result, Token{END, ""})
			} else if word == "echo" {
				result = append(result, Token{ECHO, ""})
			} else if word == "input" {
				result = append(result, Token{INPUT, ""})
			} else if word == "list" {
				result = append(result, Token{LIST, ""})
			} else if word == "len" {
				result = append(result, Token{LEN, ""})
			} else if word == "index" {
				result = append(result, Token{INDEX, ""})
			} else if word == "app" {
				result = append(result, Token{APP, ""})
			} else if word == "slice" {
				result = append(result, Token{SLICE, ""})
			} else {
				result = append(result, Token{SYMBOL, word})
			}
		} else if tools.IsSingleNum(str[index]) {
			var num string
			num, index = tools.PickNum(str, index)
			result = append(result, Token{NUM, num})
		} else if str[index] == ">" {
			result = append(result, Token{SEND, ""})
		} else if str[index] == "-" {
			if index+1 < len(str) && tools.IsSingleNum(str[index+1]) {
				var num string
				num, index = tools.PickNum(str, index)
				result = append(result, Token{NUM, num})
			} else {
				result = append(result, Token{TMP, ""})
			}
		} else if str[index] == ";" {
			result = append(result, Token{STOP, ""})
		} else if str[index] == "\n" {
			if result[len(result)-1].Type() != SEND && result[len(result)-1].Type() != BEGIN {
				/*
				   begin \n ... => not end
				   > \n ... => not end
				   or will autoly end
				*/
				result = append(result, Token{STOP, ""})
			}
		} else if str[index] == "(" {
			index++
			begin := index
			count := 1
			for ; index < len(str) && count != 0; index++ {
				if str[index] == "(" {
					count++
				} else if str[index] == ")" {
					count--
				}
			}
			index--
			result = append(result, Token{XEXP, strings.Join(str[begin:index], "")})
		} else if str[index] == ")" {
			fmt.Println("Error: too much back parenthesis. Letter: ", index)
		} else if str[index] == "\n" {
			// ; => end, obviously
			result = append(result, Token{STOP, ""})
		} else if str[index] == "\"" {
			index++
			// TODO: Avoid connect string very often ( res is connected very often
			res := ""
			for index < len(str) && str[index] != "\"" {
				if str[index] == "\\" {
					/* Escape character */
					index++
					escChar := tools.PickEscapeChar(str, index)
					res += escChar
				} else {
					res += str[index]
				}
				index++
			}
			if index >= len(str) && str[index-1] == "\"" {
				fmt.Println("Error: lose \" while giving a string value.")
			}
			result = append(result, Token{STR, res})
		} else if str[index] == "'" {
			/* Char */
			index++
			if index+1 < len(str) {
				if len(str[index]) > 1 {
					fmt.Println("Error: a char value must be ascii but not other wide codeset. Letter: ", index)
				} else {
					res := ""
					if str[index] == "\\" {
						/* Escape character */
						index++
						res = tools.PickEscapeChar(str, index)
					} else {
						res = str[index]
					}
					index++
					if str[index] == "'" {
						result = append(result, Token{CHAR, res})
					} else {
						fmt.Println("Error: lose ' while giving a char value")
					}
				}
			} else {
				fmt.Println("Error: lose ' while giving a char value")
			}
		} else {
			fmt.Println("Warn: unexpected token of: ", index+1, str[index])
		}
	}
	return result
}
