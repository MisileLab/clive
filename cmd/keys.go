package cmd

import "github.com/go-rod/rod/lib/input"

func shift(k input.Key) input.Key {
	k, _ = k.Shift()
	return k
}

var keymap = map[rune]input.Key{
	// numbers row
	'`':  input.Backquote,
	'~':  shift(input.Backquote),
	'1':  input.Digit1,
	'!':  shift(input.Digit1),
	'2':  input.Digit2,
	'@':  shift(input.Digit2),
	'3':  input.Digit3,
	'#':  shift(input.Digit3),
	'4':  input.Digit4,
	'$':  shift(input.Digit4),
	'5':  input.Digit5,
	'%':  shift(input.Digit5),
	'6':  input.Digit6,
	'^':  shift(input.Digit6),
	'7':  input.Digit7,
	'&':  shift(input.Digit7),
	'8':  input.Digit8,
	'*':  shift(input.Digit8),
	'9':  input.Digit9,
	'(':  shift(input.Digit9),
	'0':  input.Digit0,
	')':  shift(input.Digit0),
	'-':  input.Minus,
	'_':  shift(input.Minus),
	'=':  input.Equal,
	'+':  shift(input.Equal),
	'\\': input.Backslash,
	'|':  shift(input.Backslash),

	// first row
	'q': input.KeyQ,
	'Q': shift(input.KeyQ),
	'w': input.KeyW,
	'W': shift(input.KeyW),
	'e': input.KeyE,
	'E': shift(input.KeyE),
	'r': input.KeyR,
	'R': shift(input.KeyR),
	't': input.KeyT,
	'T': shift(input.KeyT),
	'y': input.KeyY,
	'Y': shift(input.KeyY),
	'u': input.KeyU,
	'U': shift(input.KeyU),
	'i': input.KeyI,
	'I': shift(input.KeyI),
	'o': input.KeyO,
	'O': shift(input.KeyO),
	'p': input.KeyP,
	'P': shift(input.KeyP),
	'[': input.BracketLeft,
	'{': shift(input.BracketLeft),
	']': input.BracketRight,
	'}': shift(input.BracketRight),

	// second row
	'a':  input.KeyA,
	'A':  shift(input.KeyA),
	's':  input.KeyS,
	'S':  shift(input.KeyS),
	'd':  input.KeyD,
	'D':  shift(input.KeyD),
	'f':  input.KeyF,
	'F':  shift(input.KeyF),
	'g':  input.KeyG,
	'G':  shift(input.KeyG),
	'h':  input.KeyH,
	'H':  shift(input.KeyH),
	'j':  input.KeyJ,
	'J':  shift(input.KeyJ),
	'k':  input.KeyK,
	'K':  shift(input.KeyK),
	'l':  input.KeyL,
	'L':  shift(input.KeyL),
	';':  input.Semicolon,
	':':  shift(input.Semicolon),
	'\'': input.Quote,
	'"':  shift(input.Quote),

	// third row
	'z': input.KeyZ,
	'Z': shift(input.KeyZ),
	'x': input.KeyX,
	'X': shift(input.KeyX),
	'c': input.KeyC,
	'C': shift(input.KeyC),
	'v': input.KeyV,
	'V': shift(input.KeyV),
	'b': input.KeyB,
	'B': shift(input.KeyB),
	'n': input.KeyN,
	'N': shift(input.KeyN),
	'm': input.KeyM,
	'M': shift(input.KeyM),
	',': input.Comma,
	'<': shift(input.Comma),
	'.': input.Period,
	'>': shift(input.Period),
	'/': input.Slash,
	'?': shift(input.Slash),

	// other
	'\n': input.Enter,
	'\t': input.Tab,
}

var specialkeymap = map[string]input.Key{
	"esc":       input.Escape,
	"backspace": input.Backspace,
	"tab":       input.Tab,
	"enter":     input.Enter,
	"left":      input.ArrowLeft,
	"up":        input.ArrowUp,
	"right":     input.ArrowRight,
	"down":      input.ArrowDown,
}
