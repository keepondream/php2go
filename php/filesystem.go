package php

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// Chmod - Changes file mode
func Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

// Chown - Chown changes the numeric uid and gid of the named file.
func Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}

// Mkdir - Makes directory
func Mkdir(name string, mode os.FileMode) error {
	return os.Mkdir(name, mode)
}

// IsDir - Tells whether the filename is a directory
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

// Copy - Copies file
func Copy(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// Fclose - Closes an open file pointer
func Fclose(file *os.File) error {
	return file.Close()
}

// Dirname - Returns a parent directory's path
func Dirname(dirPth string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirPth)
}

// Delete - Deletes a file
func Delete(name string) error {
	return Unlink(name)
}

// Unlink - Deletes a file
func Unlink(name string) error {
	return os.Remove(name)
}

// Filemtime - Gets file modification time
func Filemtime(file string) time.Time {
	fi, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}
	return fi.ModTime()
}

// FileExists - Checks whether a file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsReadable - Tells whether a file exists and is readable
func IsReadable(name string) bool {
	_, err := syscall.Open(name, syscall.O_RDONLY, 0)
	return err == nil
}

// IsWritable - Tells whether the filename is writable
func IsWritable(name string) bool {
	_, err := syscall.Open(name, syscall.O_WRONLY, 0)
	return err == nil
}

// IsWriteable - Alias of IsWritable()
func IsWriteable(name string) bool {
	return IsWritable(name)
}

// Realpath - Returns canonicalized absolute pathname
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// Rename - Renames a file or directory
func Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}
