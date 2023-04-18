package mycrypt

import "unicode"

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
		if unicode.IsUpper(melding[i]) {
			kryptertMelding[i] = 'd'
		} else if melding[i] == 'd' {
			kryptertMelding[i] = alphabet[sokIAlfabetet('K', alphabet)]
			kryptertMelding = append(kryptertMelding, alphabet[sokIAlfabetet('S', alphabet)])
			kryptertMelding = append(kryptertMelding, alphabet[sokIAlfabetet('N', alphabet)])
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i, c := range alfabet {
		if c == symbol {
			return i
		}
	}
	return -1
}
