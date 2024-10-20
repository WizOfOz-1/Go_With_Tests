package maps

import (
	"testing"
)

func TestMain(t *testing.T) {
	dictionary := Dictionary{"test": "This is a test"}
	t.Run("Search for an existing word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is a test"
		assertStrings(t, got, want)
	})
	t.Run("Search for a non-existing word", func(t *testing.T) {
		_, err := dictionary.Search("ugabuga")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "This is a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "This is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func assertError(t testing.TB, got error, expected error) {
	t.Helper()
	if got != expected {
		t.Fatal("Expected to get an error here, but the errors are different", got, expected)
	}

}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word: ", err)
	}
	assertStrings(t, got, definition)
}
