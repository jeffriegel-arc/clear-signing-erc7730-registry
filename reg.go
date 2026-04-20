package registry

import (
	"embed"
	"io/fs"
)

//go:embed registry specs tools ercs
var regFS embed.FS

func WalkDir(fn fs.WalkDirFunc) error {
	return fs.WalkDir(regFS, "registry", fn)
}

func Open(fn string) (fs.File, error) {
	return regFS.Open(fn)
}
