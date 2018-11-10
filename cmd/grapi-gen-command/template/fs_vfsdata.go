// Code generated by vfsgen; DO NOT EDIT.

// +build !vfsgen

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FS statically implements the virtual filesystem provided to vfsgen.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 9, 1, 5, 41, 0, time.UTC),
		},
		"/cmd": &vfsgen۰DirInfo{
			name:    "cmd",
			modTime: time.Date(2018, 11, 9, 1, 5, 41, 0, time.UTC),
		},
		"/cmd/{{ .name }}": &vfsgen۰DirInfo{
			name:    "{{ .name }}",
			modTime: time.Date(2018, 11, 9, 1, 5, 41, 0, time.UTC),
		},
		"/cmd/{{ .name }}/run.go.tmpl": &vfsgen۰FileInfo{
			name:    "run.go.tmpl",
			modTime: time.Date(2018, 11, 9, 1, 5, 41, 0, time.UTC),
			content: []byte("\x70\x61\x63\x6b\x61\x67\x65\x20\x6d\x61\x69\x6e\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x28\x0a\x09\x22\x66\x6d\x74\x22\x0a\x09\x22\x6f\x73\x22\x0a\x29\x0a\x0a\x66\x75\x6e\x63\x20\x6d\x61\x69\x6e\x28\x29\x20\x7b\x0a\x09\x6f\x73\x2e\x45\x78\x69\x74\x28\x72\x75\x6e\x28\x29\x29\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x20\x72\x75\x6e\x28\x29\x20\x69\x6e\x74\x20\x7b\x0a\x09\x66\x6d\x74\x2e\x50\x72\x69\x6e\x74\x6c\x6e\x28\x22\x49\x74\x20\x77\x6f\x72\x6b\x73\x21\x22\x29\x0a\x09\x72\x65\x74\x75\x72\x6e\x20\x30\x0a\x7d\x0a"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/cmd"].(os.FileInfo),
	}
	fs["/cmd"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/cmd/{{ .name }}"].(os.FileInfo),
	}
	fs["/cmd/{{ .name }}"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/cmd/{{ .name }}/run.go.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// We already imported "compress/gzip" and "io/ioutil", but ended up not using them. Avoid unused import error.
var _ = gzip.Reader{}
var _ = ioutil.Discard

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
