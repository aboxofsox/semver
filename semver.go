package semver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Semver struct {
	Major      int    // 1.x.x
	Minor      int    // x.1.x
	Patch      int    // x.x.1
	Prerelease string // x.x.x-alpha
	Meta       string // x.x.x-x+001
}

var re = regexp.MustCompile(`\d+\.\d+\.\d+(-[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?`)

func compareInts(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// Compare takes two version strings, normalizes and parses them into Semver structures,
// and then compares them according to the rules of semantic versioning.
//
// The function first compares the prerelease tags of the two versions. If both versions
// have prerelease tags, it returns -1 if the tag of the first version is lexicographically
// less than the tag of the second version, 1 if it's greater, and 0 if they're equal.
// If only one version has a prerelease tag, that version is considered smaller.
//
// If the prerelease tags are equal or nonexistent, the function compares the major, minor,
// and patch versions in that order. For each component, it returns -1 if the component of
// the first version is less than the component of the second version, 1 if it's greater,
// and 0 if they're equal.
//
// If all components are equal, the function returns 0, indicating that the two versions
// are equal.
//
// If there is an error parsing either version string, the function returns 0 and the error.
//
// Example:
//
//	result, err := Compare("1.0.0-alpha", "1.0.0-beta")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(result) // prints -1
func Compare(v1, v2 string) (int, error) {
	ver1, err := ParseVersion(normalize(v1))
	if err != nil {
		return 0, err
	}
	ver2, err := ParseVersion(normalize(v2))
	if err != nil {
		return 0, err
	}

	// compare prerelease tag
	if ver1.Prerelease != "" && ver2.Prerelease != "" {
		if ver1.Prerelease < ver2.Prerelease {
			return -1, nil
		} else if ver1.Prerelease > ver2.Prerelease {
			return 1, nil
		}
	} else if ver1.Prerelease != "" {
		return -1, nil
	} else if ver2.Prerelease != "" {
		return 1, nil
	}

	// compare version 1 major and version 2 major
	if result := compareInts(ver1.Major, ver2.Major); result != 0 {
		return result, nil
	}

	// compare version 1 minor and version 2 minor
	if result := compareInts(ver1.Minor, ver2.Minor); result != 0 {
		return result, nil
	}

	// compare version 1 pach and version 2 patch
	if result := compareInts(ver1.Patch, ver2.Patch); result != 0 {
		return result, nil
	}

	return 0, nil
}

// ParseVersion takes a version string, normalizes it, and parses it into a Semver structure.
//
// The function first checks if the version string contains a "+" or a "-" character, which
// indicate the presence of metadata or a prerelease tag, respectively. If a "+" is found,
// the function splits the string at the "+" and assigns the second part to the Meta field
// of the Semver structure. If a "-" is found, the function splits the string at the "-"
// and assigns the second part to the Prerelease field of the Semver structure.
//
// After processing the metadata and prerelease tag, the function splits the remaining
// version string at the "." characters to get the major, minor, and patch versions. These
// are converted to integers and assigned to the Major, Minor, and Patch fields of the
// Semver structure, respectively.
//
// If there is an error parsing the version string, the function returns an empty Semver
// structure and the error.
//
// Example:
//
//	ver, err := ParseVersion("1.0.0-alpha+001")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(ver) // prints {Major:1 Minor:0 Patch:0 Prerelease:alpha Meta:001}
func ParseVersion(v string) (Semver, error) {
	var (
		pre  string
		meta string
	)

	if strings.Contains(v, "+") {
		split := strings.Split(v, "+")
		v = split[0]
		if len(split) > 1 {
			meta = split[1]
		}
	}

	if strings.Contains(v, "-") {
		split := strings.Split(v, "-")
		v = split[0]
		if len(split) > 1 {
			pre = split[1]
		}
	}

	major, minor, patch, err := splitVer(v)
	if err != nil {
		return Semver{}, err
	}

	return Semver{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		Prerelease: pre,
		Meta:       meta,
	}, nil
}

func splitVer(v string) (int, int, int, error) {
	if strings.Contains(v, "+") {
		v = strings.Split(v, "+")[0]
	}

	vers := make([]int, 3)

	split := strings.Split(v, ".")
	if len(split) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid semver format")
	}

	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, 0, 0, err
		}
		vers[i] = n
	}

	return vers[0], vers[1], vers[2], nil
}

func normalize(v string) string {
	match := re.FindString(v)
	return match
}
