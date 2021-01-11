package fs

import (
	"archive/zip"
	"bazil.org/fuse/fs"
)

type FS struct {
	archive *zip.Reader
}

//var _ fs.FS = (*FS)(nil)

func (f *FS) Root() (fs.Node, error) {
	n := &Dir{
		archive: f.archive,
	}
	return n, nil
}
