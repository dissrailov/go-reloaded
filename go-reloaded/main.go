package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regexMap = map[string]*regexp.Regexp{
	"hex":        regexp.MustCompile(`(\w+)([\s*[:punct:]]*)(\(\s*(?i)hex\s*\)) *`),
	"bin":        regexp.MustCompile(`(\w+)([\s*[:punct:]]*)(\(\s*(?i)bin\s*\)) *`),
	"capnum":     regexp.MustCompile(`([\s\S]+?)\s*\(\s*(?i)cap\s*,\s*(\w+)\s*\) *`),
	"cap":        regexp.MustCompile(`(\w+?)([\s*[:punct:]]*)(\(\s*(?i)cap\s*\)) *`),
	"lownum":     regexp.MustCompile(`([\s\S]+?)\s*\(\s*(?i)low\s*,\s*(\w+)\s*\) *`),
	"low":        regexp.MustCompile(`(\w+?)([\s*[:punct:]]*)(\(\s*(?i)low\s*\)) *`),
	"upnum":      regexp.MustCompile(`([\s\S]+?)\s*\(\s*(?i)up\s*,\s*(\w+)\s*\) *`),
	"up":         regexp.MustCompile(`(\w+?)([\s*[:punct:]]*)(\(\s*(?i)up\s*\)) *`),
	"replace":    regexp.MustCompile(` *▓ *`),
	"quot":       regexp.MustCompile(` *'([^']*)' *`),
	"doublequot": regexp.MustCompile(` *"([^"]*)" *`),
	"newline":    regexp.MustCompile(`\n *`),
	"all":        regexp.MustCompile(`\(\s*(?i)low\s*,\s*(\d+)\s*\)|\(\s*(?i)up\s*,\s*(\d+)\s*\)|(\(\s*(?i)up\s*\))\s*|(\(\s*hex\s*\)) *|(\(\s*bin\s*\))\s*|(\(\s*(?i)low\s*\))\s*|(\(\s*(?i)cap\s*\))\s*|\(\s*(?i)cap\s*,\s*(\d+)\s*\)`),
	"space":      regexp.MustCompile(`\n +`),
	"spacef":     regexp.MustCompile(` +`),
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]
	if !strings.HasSuffix(inputFileName, ".txt") {
		fmt.Println("Ошибка: входной файл должен иметь расширение .txt")
		os.Exit(1)
	}

	if !strings.HasSuffix(outputFileName, ".txt") {
		fmt.Println("Ошибка: выходной файл должен иметь расширение .txt")
		os.Exit(1)
	}

	inputFile, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("Ошибка при чтении входного файла: %v", err)
	}

	line := string(inputFile)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Ошибка при создании выходного файла: %v", err)
	}
	defer outputFile.Close()
	line = Process(line)
	fmt.Fprint(outputFile, line)
}

// tam gde vse rabotaet
func Process(line string) string {
	line = ArticleA(line)
	found := false
	line = ArticleAn(line)
	line = regexMap["quot"].ReplaceAllStringFunc(line, func(match string) string {
		matches := regexMap["quot"].FindStringSubmatch(match)
		return " '" + strings.TrimSpace(matches[1]) + "' "
	})
	line = regexMap["doublequot"].ReplaceAllStringFunc(line, func(match string) string {
		matches := regexMap["doublequot"].FindStringSubmatch(match)
		return " \"" + strings.TrimSpace(matches[1]) + "\" "
	})
	for {
		match := regexMap["all"].FindString(line)
		if len(match) == 0 {
			break
		}

		switch {
		case regexp.MustCompile(`(?i)\(up\)`).MatchString(match):
			zashel := false
			line = regexMap["up"].ReplaceAllStringFunc(line, func(match string) string {
				if zashel {
					return match
				}
				found = true
				word := regexMap["up"].FindStringSubmatch(match)[1]
				group2 := regexMap["up"].FindStringSubmatch(match)[2]
				up := regexMap["up"].FindStringSubmatch(match)[3]
				changed := strings.ToUpper(word)
				if strings.Contains(match, word+" "+up) {
					zashel = true
					return changed + group2
				} else {
					zashel = true
					return changed + group2 + " "
				}
			})
			if !found {
				return ""
			}
		case regexp.MustCompile(`(?i)\(low\)`).MatchString(match):
			zashel := false
			line = regexMap["low"].ReplaceAllStringFunc(line, func(match string) string {
				if zashel {
					return match
				}
				found = true
				word := regexMap["low"].FindStringSubmatch(match)[1]
				group2 := regexMap["low"].FindStringSubmatch(match)[2]
				low := regexMap["low"].FindStringSubmatch(match)[3]
				changed := strings.ToLower(word)
				if strings.Contains(match, word+" "+low) {
					zashel = true
					return changed + group2
				} else {
					zashel = true
					return changed + group2 + " "
				}
			})
			if !found {
				return ""
			}
		case regexp.MustCompile(`(?i)\(cap\)`).MatchString(match):
			zashel := false
			line = regexMap["cap"].ReplaceAllStringFunc(line, func(match string) string {
				if zashel {
					return match
				}
				found = true
				word := regexMap["cap"].FindStringSubmatch(match)[1]
				group2 := regexMap["cap"].FindStringSubmatch(match)[2]
				cap := regexMap["cap"].FindStringSubmatch(match)[3]
				changed := strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
				if strings.Contains(match, word+" "+cap) {
					zashel = true
					return changed + group2
				} else {
					zashel = true
					return changed + group2 + " "
				}
			})
			if !found {
				return ""
			}
		case regexp.MustCompile(`\(\s*(?i)up\s*,\s*(\w+)\s*\)`).MatchString(match):
			line = regexMap["newline"].ReplaceAllStringFunc(line, func(match string) string {
				return "▓ "
			})
			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "▓ "
			})
			line = regexMap["upnum"].ReplaceAllStringFunc(line, func(match string) string {
				found = true
				lines := strings.Split(match, "\n")
				submatches := regexMap["upnum"].FindStringSubmatch(lines[0])
				if len(submatches) < 1 {
					return match
				}
				words := strings.Fields(submatches[1])
				num, err := strconv.Atoi(submatches[2])
				words = Nurma(words)
				if err != nil {
					log.Fatal(err)
				}
				valid := false

				for i := len(words) - 1; i >= 0 && num > 0; i-- {
					for _, ch := range words[i] {
						if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
							valid = true
							break
						}
					}
					if valid {
						words[i] = strings.ToUpper(words[i])
						num--
					}
				}

				if num > len(words) {
					log.Fatalf("Error: change num")
				}

				return strings.Join(words, " ") + " "
			})
			if !found {
				return ""
			}

			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "\n"
			})
		case regexp.MustCompile(`\(\s*(?i)low\s*,\s*(\w+)\s*\)`).MatchString(match):
			line = regexMap["newline"].ReplaceAllStringFunc(line, func(match string) string {
				return "▓ "
			})
			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "▓ "
			})
			line = regexMap["lownum"].ReplaceAllStringFunc(line, func(match string) string {
				lines := strings.Split(match, "\n")
				submatches := regexMap["lownum"].FindStringSubmatch(lines[0])
				found = true
				if len(submatches) < 1 {
					return match
				}
				words := strings.Fields(submatches[1])
				num, err := strconv.Atoi(submatches[2])
				words = Nurma(words)
				if err != nil {
					log.Fatal(err)
				}
				valid := false

				for i := len(words) - 1; i >= 0 && num > 0; i-- {
					for _, ch := range words[i] {
						if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
							valid = true
							break
						}
					}
					if valid {
						words[i] = strings.ToLower(words[i])
						num--
					}
				}

				if num > len(words) {
					log.Fatalf("Error: change num")
				}

				return strings.Join(words, " ") + " "
			})
			if !found {
				return ""
			}
			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "\n"
			})
		case regexp.MustCompile(`\(\s*(?i)cap\s*,\s*(\w+)\s*\)`).MatchString(match):
			line = regexMap["newline"].ReplaceAllStringFunc(line, func(match string) string {
				return "▓ "
			})
			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "▓ "
			})
			line = regexMap["capnum"].ReplaceAllStringFunc(line, func(match string) string {
				found = true
				lines := strings.Split(match, "\n")
				submatches := regexMap["capnum"].FindStringSubmatch(lines[0])
				if len(submatches) < 1 {
					return match
				}
				words := strings.Fields(submatches[1])
				num, err := strconv.Atoi(submatches[2])
				words = Nurma(words)
				if err != nil {
					log.Fatal(err)
				}
				valid := false

				for i := len(words) - 1; i >= 0 && num > 0; i-- {
					for _, ch := range words[i] {
						if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
							valid = true
							break
						}
					}
					if valid {
						words[i] = strings.Title(strings.ToLower(words[i]))
						num--
					}
				}

				if num > len(words) {
					log.Fatalf("Error: change num")
				}
				return strings.Join(words, " ") + " "
			})
			if !found {
				return ""
			}
			line = regexMap["replace"].ReplaceAllStringFunc(line, func(string) string {
				return "\n"
			})
		case regexp.MustCompile(`(?i)\(hex\)`).MatchString(match):
			zashel := false
			line = regexMap["hex"].ReplaceAllStringFunc(line, func(match string) string {
				if zashel {
					return match
				}
				found = true
				hexStr := regexMap["hex"].FindStringSubmatch(match)[1]
				group2 := regexMap["hex"].FindStringSubmatch(match)[2]
				decStr := fmt.Sprintf("%d", ParserInt(hexStr, 16))
				hex := regexMap["hex"].FindStringSubmatch(match)[3]
				if strings.Contains(match, hexStr+" "+hex) {
					if decStr == "0" {
						return strings.Replace(match, hexStr+group2+hex, hexStr, 1)
					}
					zashel = true
					line = strings.Replace(match, hexStr+group2+hex, decStr, 1)
					return line
				} else {
					zashel = true
					return strings.Replace(match, hexStr+group2+hex, decStr+group2, 1)
				}
			})
			if !found {
				return ""
			}
			// бин в десимал
		case regexp.MustCompile(`(?i)\(bin\)`).MatchString(match):
			zashel := false
			line = regexMap["bin"].ReplaceAllStringFunc(line, func(match string) string {
				if zashel {
					return match
				}
				found = true
				binStr := regexMap["bin"].FindStringSubmatch(match)[1]
				group2 := regexMap["bin"].FindStringSubmatch(match)[2]
				decStr := fmt.Sprintf("%d", ParserInt(binStr, 2))
				bin := regexMap["bin"].FindStringSubmatch(match)[3]
				if strings.Contains(match, binStr+" "+bin) {
					if decStr == "0" {
						return strings.Replace(match, binStr+group2+bin, binStr, 1)
					}
					zashel = true
					line = strings.Replace(match, binStr+group2+bin, decStr, 1)
					return line
				} else {
					zashel = true
					return strings.Replace(match, binStr+group2+bin, decStr+group2, 1)
				}
			})
			if !found {
				return ""
			}

		default:
			rall := regexp.MustCompile(`\(\s*(?i)low\s*\)|\(\s*(?i)cap\s*\)|\(\s*(?i)up\s*\)|\(\s*(?i)hex\s*\)|\(\s*(?i)bin\s*\)|\(\s*(?i)cap\s*,\s*\d+\s*\)|\(\s*(?i)up\s*,\s*\d+\s*\)|\(\s*(?i)low\s*,\s*\d+\s*\)`)
			line = rall.ReplaceAllString(line, "")
		}
	}
	line = Punctuation(line)
	line = FixQuotes(line)
	line = FixDoubleQuotes(line)
	line = ArticleA(line)
	line = regexMap["space"].ReplaceAllStringFunc(line, func(s string) string {
		return "\n"
	})
	line = regexMap["spacef"].ReplaceAllStringFunc(line, func(s string) string {
		return " "
	})
	return line
}

func FixQuotes(text string) string {
	text = regexp.MustCompile(`'\s*([^']+?)\s*'`).ReplaceAllString(text, "'$1'")

	return text
}

func FixDoubleQuotes(text string) string {
	text = regexp.MustCompile(`"\s*([^"]+?)\s*"`).ReplaceAllString(text, "\"$1\"")

	return text
}

// newline
func Nurma(str []string) []string {
	var res []string
	var move bool = false
	for i := 1; i < len(str); i++ {
		if str[i] == "▓" {
			str[i-1] = str[i-1] + str[i]
			str[i] = ""
			move = true
		}
	}
	if move {
		str = Nurma(str)
	}
	for i := 0; i < len(str); i++ {
		if len(str[i]) > 0 {
			res = append(res, str[i])
		}
	}
	return res
}

// rabota hex and bin
func ParserInt(str string, base int) int {
	val, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(val)
}

// article
func ArticleA(text string) string {
	patt := regexp.MustCompile(`((?i)\s+a|^a)(\s+)((?i)[aeiouh])`)
	for patt.MatchString(text) {
		text = patt.ReplaceAllStringFunc(text, func(s string) string {
			match := patt.FindStringSubmatch(s)
			if len(match) >= 2 {
				return match[1] + "n" + match[2] + match[3]
			}
			return s
		})
	}
	return text
}

func ArticleAn(text string) string {
	patt := `((?i)\s+an|^an)(\s+)((?i)[^aeiouh])`
	state := regexp.MustCompile(patt)
	for state.MatchString(text) {
		text = state.ReplaceAllStringFunc(text, func(s string) string {
			match := state.FindStringSubmatch(s)
			if len(match) > 2 {
				return match[1][:len(match[1])-1] + match[2] + match[3]
			}
			return s
		})
	}
	return text
}

// punct
func Punctuation(text string) string {
	punc := regexp.MustCompile(`\s*([.,:;!?]+)[ 	]*`)
	text = punc.ReplaceAllString(text, "$1")
	puncSpaces := regexp.MustCompile(`([.,:;!?]+)([^.,:;!?]+)`)
	text = puncSpaces.ReplaceAllString(text, "$1 $2")
	return text
}
