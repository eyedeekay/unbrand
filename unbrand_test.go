package ubrand

import (
	"testing"
)

func TestMain(t *testing.T) {
	filesfirefox := FindFilesFirefox()
	if len(filesfirefox) < 1 {
		t.Error("Search failed, do you have an unmodified Firefox checked out?")
	}
	for _, filefirefox := range filesfirefox {
		t.Log(filefirefox)
	}
}
