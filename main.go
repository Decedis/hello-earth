package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, eg. en, ur...")
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language code
type language string

// phrasebook holds greetings for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello World!",      // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese”
}

func greet(l language) string {
	// return a simple greeting message
	greeting, ok := phrasebook[l]
	// greeting = stored message
	// ok = if greeting exists then true, else then false
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
