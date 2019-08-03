package fileline

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func FileLineFunc(offset int) string {
	if offset < 0 {
		offset = 0
	}
	depth := 1 + offset
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return "FileLineFunc"
	}

	baseFile := filepath.Base(file)
	name := runtime.FuncForPC(pc).Name()
	idx := strings.LastIndex(name, ".")
	if idx != -1 {
		name = string(name[idx+1:])
	}

	return fmt.Sprintf("%s:%v:%s", baseFile, line, name)
}
