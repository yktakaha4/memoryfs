package memoryfs

import (
	path2 "path"
	"strings"
)

func cleanse(path string) string {
	path = strings.ReplaceAll(path, "/", separator)
	path = path2.Clean(path)
	path = strings.TrimPrefix(path, "."+separator)
	path = strings.TrimPrefix(path, separator)
	if path == "." {
		return ""
	}
	return path
}
