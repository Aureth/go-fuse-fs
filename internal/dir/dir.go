package dir

import (
	"archive/zip"
	"bazil.org/fuse"
)

type Dir struct {
	archive *zip.Reader
	// nil for the root directory, which has no entry in the zip
	file *zip.File
}

//var _ fs.Node = (*Dir)(nil)

func ZipAttr(f *zip.File, a *fuse.Attr) {
	a.Size = f.UncompressedSize64
	a.Mode = f.Mode()
	a.Mtime = f.ModTime()
	a.Ctime = f.ModTime()
	a.Crtime = f.ModTime()
}
