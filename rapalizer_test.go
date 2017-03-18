package rapalizer
import (
	"testing"
	"fmt"
)

const singleRhyme string = "You missed two classes and no homework.  But your teacher preaches class like you're some kind of jerk!"

const fullSong string = "Kick it! " +
	"You wake up late for school man you don't wanna go! " +
	"You ask you mom, Please? but she still says, No! " +
	"You missed two classes and no homework " +
	"But your teacher preaches class like you're some kind of jerk " +
	"You gotta fight for your right to party " +
	"You pops caught you smoking and he said, No way! " +
	"That hypocrite smokes two packs a day " +
	"Man, living at home is such a drag " +
	"Now your mom threw away your best porno mag (Busted!) " +
	"Dont step out of this house if thats the clothes youre gonna wear " +
	"Ill kick you out of my home if you dont cut that hair " +
	"Your mom busted in and said, Whats that noise? " +
	"Aw, mom youre just jealous its the Beastie Boys!"

func TestNullString(t *testing.T) {
	var rap Rapalizer
	rap.CalculateScore()
	if rap.Score != -9999 {
		t.Error("Empty String did not return negative score: ",  rap.Score)
	}
}

func TestSingleWordString(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word.")
	rap.CalculateScore()
	if rap.Score != 0 {
		fmt.Println("Score: ", rap.Score)
		t.Error("Single word cannot rhyme and should return a 0 score.")
	}
}

func TestReplicateWord(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word word.")
	rap.CalculateScore()
	if rap.Score != -1 {
		fmt.Println("Score: ", rap.Score)
		t.Error("double word.  should return -1 score.")
	}
}

func TestSingleRhyme(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word Turd.")
	rap.CalculateScore()
	if rap.Score != 1 {
		fmt.Println("Score: ", rap.Score)
		t.Error("single rhyme.  should return 1 score.")
	}
}

func TestSingleNoRhyme(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word Nope.")
	rap.CalculateScore()
	if rap.Score != 0 {
		fmt.Println("Score: ", rap.Score)
		t.Error("no rhyme.  should return 0 score.")
	}
}

func TestSingleRhymeWithSkipWords(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word and a Turd.")
	rap.CalculateScore()
	if rap.Score != 1 {
		fmt.Println("Score: ", rap.Score)
		t.Error("single rhyme with skip words.  should return 1 score.")
	}
}

func TestTripleRhymeWithSkipWords(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Word and a Turd and a Bird.")
	rap.CalculateScore()
	if rap.Score != 3 {
		fmt.Println("Score: ", rap.Score)
		t.Error("single rhyme with skip words.  should return 3 score.")
	}
}

func TestRealSingleRhymeLyricUnderTwentyWords(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray(singleRhyme)
	rap.CalculateScore()
	if rap.Score != 2 {
		fmt.Println("Score: ", rap.Score)
		t.Error("single rhyme with skip words.  should return 1 score.")
	}
}

func TestRealSingleRhymeLyricOverTwentyWords(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray(fullSong)
	rap.CalculateScore()
	if rap.Score != 6 {
		fmt.Println("Score: ", rap.Score)
		t.Error("full rhyme with skip words.  should return SOME score.")
	}
}

func TestisWordInSlice(t *testing.T) {
	warray := []string{ "I", "Will", "Take", "A", "Giant", "Dump"}
	if IsWordInSlice("Will", warray) != true {
		t.Error("Word is in slice but returning false.")
	}
	if IsWordInSlice("Farts", warray) != false {
		t.Error("Word is not in slice but returning true.")
	}
}

func TestStripPunctuation(t *testing.T) {
	if "Farting" != StripPunctuation("Farting!") {
		t.Error("Ending punctuation on a word is not stripped.")
	}
	if "Farting" != StripPunctuation("(Farting!") {
		t.Error("Starting punctuation on a word is not stripped.")
	}
	if "Farting" != StripPunctuation("(Farting?") {
		t.Error("Double sided punctuation on a word is not stripped.")
	}
	if "Farting" != StripPunctuation("Farting") {
		t.Error("Clean word is not properly passing through logic.")
	}
}

func TestNormalizingWords(t *testing.T) {
	if "farting" != NormalizeWord("(FaRtiNg?") {
		t.Error("Word Normalization failed.")
	}
}

func TestLoadingStringIntoWordArray(t *testing.T) {
	var rap Rapalizer
	rap.LoadStringIntoWordArray("")
	if len(rap.Lyrics) != 1 {
		t.Error("Lyric-Word array should be empty. Size: ", len(rap.Lyrics))
	}
	var rap1 Rapalizer
	rap1.LoadStringIntoWordArray("word")
	if len(rap1.Lyrics) != 1 || rap1.Lyrics[0] != "word" {
		t.Error("Should only have one word in array: ", rap1.Lyrics[0])
	}
	var rap2 Rapalizer
	rap2.LoadStringIntoWordArray("This is a complete Sentence.")
	ideal := []string{"This", "is", "a", "complete", "Sentence."}
	if len(rap2.Lyrics) != 5  {
		t.Error("Sentence loaded incorrectly: ", rap2.Lyrics)
	}
	for i := range rap2.Lyrics {
		if rap2.Lyrics[i] != ideal[i] {
			t.Error("Slice is not equal to ideal slice.", rap2.Lyrics)
		}
	}
	var rap3 Rapalizer
	rap3.LoadStringIntoWordArray(singleRhyme)
	if len(rap3.Lyrics) != 19 {
		t.Error("Beastie Boys Lyric did not properly load: ", len(rap3.Lyrics), rap3.Lyrics)
	}
}

func TestSanitizeString(t *testing.T) {
	stringWithNewlines := `This\nhas\nnewlines`
	if SanitizeString(stringWithNewlines) != "This has newlines" {
		fmt.Println("string: %s", SanitizeString(stringWithNewlines))
		t.Error("String with newlines incorrectly sanitized.")
	}
	stringWithSpecialChars := `*()//\\Redman`
	if SanitizeString(stringWithSpecialChars) != "Redman" {
		t.Error("String with special chars incorrectly sanitized.")
	}
}

func TestSetArtist(t *testing.T) {
	var rap Rapalizer
	rap.SetArtist("Method Man")
	if rap.Artist != "Method Man" {
		t.Error("SetArtist incorrectly assigned string to object variable")
	}
}

func TestSetTitle(t *testing.T) {
	var rap Rapalizer
	rap.SetSongTitle("Method Man")
	if rap.Title != "Method Man" {
		t.Error("SetSongTitle incorrectly assigned string to object variable")
	}
}

func TestRapalizerToJson(t *testing.T) {
	idealString := `{"lyrics":["Fight","for","your","right."],"score":0,"artist":"Beastie Boys","title":"Fight","pairs":null}`
	var rap Rapalizer
	rap.LoadStringIntoWordArray("Fight for your right.")
	rap.SetSongTitle("Fight")
	rap.SetArtist("Beastie Boys")
	objJson := rap.ToJson()
	if objJson != idealString {
		t.Error("ToJson not properly converting.")
	}
}