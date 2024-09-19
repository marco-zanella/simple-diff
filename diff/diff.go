package diff

import (
	"bufio"
	"os"
	"strings"
)

// DiffResult holds the results of the diff operation
type DiffResult struct {
	Added   []string
	Removed []string
}

// ComputeDiff calculates the added and removed lines between two slices of lines.
func ComputeDiff(lines1, lines2 []string) DiffResult {
	lineMap1 := make(map[string]bool)
	lineMap2 := make(map[string]bool)

	for _, line := range lines1 {
		lineMap1[line] = true
	}

	for _, line := range lines2 {
		lineMap2[line] = true
	}

	var added []string
	var removed []string

	// Check for lines in lines2 that are not in lines1 (added)
	for line := range lineMap2 {
		if !lineMap1[line] {
			added = append(added, line)
		}
	}

	// Check for lines in lines1 that are not in lines2 (removed)
	for line := range lineMap1 {
		if !lineMap2[line] {
			removed = append(removed, line)
		}
	}

	return DiffResult{
		Added:   added,
		Removed: removed,
	}
}

// ReadFileLines reads all lines from a file and returns them as a slice of strings.
func ReadFileLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

func DiffFilesStandard(path1, path2 string) (DiffResult, error) {
	lines1 := []string{}
	var err error

	if path1 != "" {
		lines1, err = ReadFileLines(path1)
		if err != nil {
			return DiffResult{}, err
		}
	}

	lines2, err := ReadFileLines(path2)
	if err != nil {
		return DiffResult{}, err
	}

	return ComputeDiff(lines1, lines2), nil
}

func DiffFilesFast(path1, path2 string) (DiffResult, error) {
	lineMap := make(map[string]bool)
	result := DiffResult{
		Added:   []string{},
		Removed: []string{},
	}

	if path1 != "" {
		file, err := os.Open(path1)
		if err != nil {
			return result, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			lineMap[line] = false
		}

		if err := scanner.Err(); err != nil {
			return result, err
		}
	}

	file, err := os.Open(path2)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		_, ok := lineMap[line]
		if ok {
			lineMap[line] = true
		} else {
			result.Added = append(result.Added, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	for line, seeninSecondFile := range lineMap {
		if !seeninSecondFile {
			result.Removed = append(result.Removed, line)
		}
	}

	return result, nil
}

// DiffFiles compares two files and returns the diff result.
func DiffFiles(path1, path2 string) (DiffResult, error) {
	return DiffFilesFast(path1, path2)
}
