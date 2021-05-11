package ubrand

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

import (
	"github.com/rwtodd/Go.Sed/sed"
)

import (
	"github.com/jorgemarey/filelib"
)

var mercurial = "hg"

var hgpath, hgerr = FindMercurial()

// SearchPath looks in the $PATH to find an executable and outputs the found path or nothing.
func SearchPath(filename string) (foundpath string) {
	pathsplit := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range pathsplit {
		foundpath = filepath.Join(path, filename)
		if FileExists(foundpath) {
			return
		}
	}
	foundpath = ""
	return
}

// Checks if a File exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// SedHelper
func SedHelper(expression, file string) (err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	presum := md5.Sum(content)
	engine, err := sed.New(strings.NewReader(expression))
	if err != nil {
		return
	}
	postcontent, err := engine.RunString(string(content))
	if err != nil {
		return
	}
	postsum := md5.Sum([]byte(postcontent))
	if postsum == presum {
		err = fmt.Errorf("File %s was not changed. Sums were: \tbefore sum: %s\n\t after:%s\n", file, presum, postsum)
		log.Printf("%s", err.Error())
		return
	}
	return
}

// FindMercurial
func FindMercurial() (mercurialpath string, err error) {
	mercurialpath = SearchPath(mercurial)
	err = fmt.Errorf("Mercurial is required to clone Firefox. You can install:\n\tapt-get install mercurial - on Ubuntu and Debian distributions\n\t pacman install mercurial - on Arch distributions")
	return
}

// SearchSelector
func SearchSelector(search string) func(os.FileInfo, string) bool {
	return func(fi os.FileInfo, fn string) bool {
		return !fi.IsDir() && strings.Contains(fn, search)
	}
}

// Rename
func Rename(element, before, after string) (err error) {
	final := strings.Replace(element, before, after, -1)
	return os.Rename(element, final)
}

// RenameBatch
func RenameBatch(batch []string, before, after string) (err error) {
	for _, element := range batch {
		if err := Rename(element, before, after); err != nil {
			return err
		}
	}
	return nil
}

// FindFiles
func FindFiles(search string) (filelist []string) {
	wd, _ := os.Getwd()
	files := filelib.Search(SearchSelector(search), wd)
	for file := range files {
		filelist = append(filelist, file)
	}
	return filelist
}
