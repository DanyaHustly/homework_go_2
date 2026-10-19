package text

import (
	"strings"
	"unicode"
)

// Блок text закрепляет работу со строками:
// string, byte, rune, immutable string, Unicode,
// пакет strings и простые операции над текстом.

// ByteLen возвращает длину строки в байтах.
//
// TODO: верните размер s в байтах. Для Unicode-строк количество байт может быть больше количества символов.
func ByteLen(s string) int {
	return len(s)
}

// RuneLen возвращает количество Unicode-символов в строке.
//
// TODO: верните количество Unicode-символов в s. Кириллица и emoji должны считаться как отдельные символы.
func RuneLen(s string) int {
	return len([]rune(s))
}

// FirstRune возвращает первый Unicode-символ строки.
//
// TODO: верните первый Unicode-символ s как строку. Для пустой строки верните "".
func FirstRune(s string) string {
	if len(s) == 0 {
		return ""
	}
	runs := []rune(s)
	return string(runs[0])
}

// LastRune возвращает последний Unicode-символ строки.
//
// TODO: верните последний Unicode-символ s как строку. Для пустой строки верните "".
func LastRune(s string) string {
	if len(s) == 0 {
		return ""
	}
	runs := []rune(s)
	return string(runs[len(runs)-1])

}

// Trim убирает пробелы по краям строки.
//
// TODO: верните s без пробельных символов по краям. Внутреннее содержимое строки не изменяйте.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// ToLower переводит строку в нижний регистр.
//
// TODO: верните s в нижнем регистре с корректной обработкой латиницы и кириллицы.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper переводит строку в верхний регистр.
//
// TODO: верните s в верхнем регистре с корректной обработкой латиницы и кириллицы.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// NormalizeEmail нормализует email.
//
// TODO: удалите пробельные символы по краям email и приведите всю строку к нижнему регистру.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// ContainsWord проверяет, содержит ли text подстроку word.
//
// TODO: верните true, если word встречается внутри text. Поиск чувствителен к регистру; пустая строка word считается найденной.
func ContainsWord(text, word string) bool {
	return strings.Contains(text, word)
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
//
// TODO: верните новую строку, заменив первый Unicode-символ s на r. Для пустой строки верните "".
func ReplaceFirstRune(s string, r rune) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	runes[0] = r
	return string(runes)
}

// ReverseRunes разворачивает строку по Unicode-символам.
//
// TODO: верните s в обратном порядке по Unicode-символам. Кириллица и emoji не должны повреждаться.
func ReverseRunes(s string) string {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return string(runes)
}

// Initials возвращает инициалы имени и фамилии.
//
// TODO: очистите имя и фамилию по краям, возьмите первые Unicode-символы непустых частей, переведите их в верхний регистр и соедините без разделителя.
func Initials(firstName, lastName string) string {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	result := ""

	if firstName != "" {
		first := []rune(firstName)[0]
		result += string(unicode.ToUpper(first))
	}

	if lastName != "" {
		last := []rune(lastName)[0]
		result += string(unicode.ToUpper(last))
	}

	return result

}

// RepeatWord повторяет слово count раз.
//
// TODO: верните word, повторённое count раз без разделителя. При count <= 0 верните "".
func RepeatWord(word string, count int) string {
	if count <= 0 {
		return ""
	}

	return strings.Repeat(word, count)
}

// JoinWithComma объединяет строки через запятую.
//
// TODO: соедините элементы values через запятую без дополнительных пробелов. Для пустого или nil-среза верните "".
func JoinWithComma(values []string) string {
	return strings.Join(values, ",")
}

// IsPalindrome проверяет, является ли строка палиндромом.
//
// TODO: верните true, если s читается одинаково в обоих направлениях после удаления пробелов и приведения к одному регистру. Поддержите Unicode.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	s = strings.Join(strings.Fields(s), "")

	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}

		left++
		right--
	}

	return true
}
