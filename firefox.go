package ubrand

var FirefoxVersionMajor = "78"
var FirefoxVersionMinor = "10"
var FirefoxSubVersion = "0"
var UnbrandedVersion = "1"
var FirefoxMirror = "https://ftp.mozilla.org/pub/mozilla.org/firefox/releases/"

// FirefoxVersion
func FirefoxVersion() string {
	return FirefoxVersionMajor + "." + FirefoxVersionMinor + "." + FirefoxSubVersion
}

// FindFirefoxMirrorURL
func FindFirefoxMirrorURL() string {
	return FirefoxMirror + FirefoxVersion()
}

// FindFirefoxSourceURL
func FindFirefoxSourceURL() string {
	return FindFirefoxMirrorURL() + "esr/source/"
}

// FindFirefoxSourceBallURL
func FindFirefoxSourceBallURL() string {
	return FindFirefoxSourceURL() + SourceBall()
}

// SourceBall
func SourceBall() string {
	return "firefox-" + FirefoxVersion() + "esr.source.tar.xz"
}

// FindFilesFennec
func FindFilesFennec() (filelist []string) {
	return FindFiles("fennec")
}

// FindFilesFirefox
func FindFilesFirefox() (filelist []string) {
	return FindFiles("firefox")
}

// FindFilesRunMozilla
func FindFilesRunMozilla() (filelist []string) {
	return FindFiles("run-mozilla")
}
