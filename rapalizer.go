package rapalizer

import (
	"fmt"
	"strings"
	"encoding/json"
)

const NUM_WORDS_TO_FIND_RHYMES int = 20
const MIN_CHARS_IN_ELIGIBLE_WORD int = 2

type Rapalizer struct {
	Lyrics 	[]string 		`json:"lyrics"`
	Score 	int 			`json:"score"`
	Artist 	string 			`json:"artist"`
	Title 	string 			`json:"title"`
	Pairs 	[]LyricPairs 	`json:"pairs"`
}

type LyricPairs struct {
	First 	string
	Second 	string
}
func (rap *Rapalizer) ToJson() string {
	j, err := json.Marshal(rap)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(j)
}
func (rap *Rapalizer) SetArtist(artist string) {
	rap.Artist = SanitizeString(artist)
}

func (rap *Rapalizer) SetSongTitle(title string) {
	rap.Title = SanitizeString(title)
}

func SanitizeString(input string) string {
	cleansedString := strings.Replace(input, ")", "", -1)
	cleansedString = strings.Replace(cleansedString, "(", "", -1)
	cleansedString = strings.Replace(cleansedString, "/", "", -1)
	cleansedString = strings.Replace(cleansedString, "*", "", -1)
	cleansedString = strings.Replace(cleansedString, "//", "", -1)
	cleansedString = strings.Replace(cleansedString, "(", "", -1)
	cleansedString = strings.Replace(cleansedString, "\n", " ", -1)
	cleansedString = strings.Replace(cleansedString, "\r", " ", -1)
	cleansedString = strings.Replace(cleansedString, "\\n", " ", -1)
	cleansedString = strings.Replace(cleansedString, "\\r", " ", -1)
	cleansedString = strings.Replace(cleansedString, "\\", "", -1)
	return cleansedString
}

func (rap *Rapalizer) LoadStringIntoWordArray(lyrics string) {
	sanitizedString := SanitizeString(lyrics)
	importedLyrics := strings.Split(sanitizedString, " ")
	if len(importedLyrics) > 0 {
		for _, word := range importedLyrics {
			rap.Lyrics = append(rap.Lyrics, word)
		}
	}
}

func (rap *Rapalizer) CalculateScore() {
	wordsToSkip := []string{"is", "but", "if", "in", "as", "a", "i", 
	"the", "of", "or", "an", "and", "you", "me", " ", ""}
	if len(rap.Lyrics) == 0 {
		rap.Score = -9999
	} else {
		rap.CalculateRapScore(rap.Lyrics, wordsToSkip)
		fmt.Println("Final Score:", rap.Score)
	}
}

func (rap *Rapalizer) CalculateRapScore(words []string, skipWords []string) {
	for i, word := range words {
		if !IsWordInSlice(word, skipWords) {
			if (len(words) - i) < NUM_WORDS_TO_FIND_RHYMES {
				rap.CompareSuffixes(NormalizeWord(word), words[i:], skipWords)
			} else {
				rap.CompareSuffixes(NormalizeWord(word), words[i:i+NUM_WORDS_TO_FIND_RHYMES], skipWords)
			}
		}
	}
}

func (rap *Rapalizer) CompareSuffixes(target string, words []string, skipWords []string) {
	for i := 1; i < len(words); i++ {
		thisword := NormalizeWord(words[i])
		if !IsWordInSlice(thisword, skipWords) && len(thisword) > 1 {
			rap.CompareSuffixPair(target, thisword)
		}
	}
}

func (rap *Rapalizer) CompareSuffixPair(firstWord string, compareWord string) {
	firstSuffix := firstWord[len(firstWord)-MIN_CHARS_IN_ELIGIBLE_WORD:]
	compareSuffix := compareWord[len(compareWord)-MIN_CHARS_IN_ELIGIBLE_WORD:]

	// TODO:  Put special situation logic here. EX: noise/boys does not currently rhyme.
	if firstSuffix == compareSuffix {
		if firstWord == compareWord {
			rap.Score--
			fmt.Println("Rem One:", firstWord, compareWord)
		} else {
			fmt.Println("Add One:", firstWord, compareWord)
			rap.Pairs = append(rap.Pairs, LyricPairs{firstWord, compareWord})
			rap.Score++
		}
	}
}

func NormalizeWord(word string) string {
	return StripPunctuation(strings.ToLower(word))
}

func StripPunctuation(word string) string {
	if strings.ContainsAny(word, ".|!|,|?|:|;|)") {
		if strings.ContainsAny(word, "(") {
			return word[1:len(word)-1]
		} else {
			return word[:len(word)-1]
		}
	}
	return word
}

func IsWordInSlice(entry string, wordslice []string) bool {
	for _, word := range wordslice {

		if NormalizeWord(entry) == NormalizeWord(word) {
			return true
		}
	}
	return false
}