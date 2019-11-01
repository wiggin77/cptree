package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func cptree(opts Opts) error {
	if err := checkOpts(opts); err != nil {
		return nil
	}

	filepath.Walk(opts.src, func(path string, info os.FileInfo, err error) error {
		return walkFunc(opts, path, info, err)
	})

	return nil
}

func walkFunc(opts Opts, path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - %v\n", path, err)
		return nil
	}

	rel, err := filepath.Rel(opts.src, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - %v\n", path, err)
		return err
	}
	dstPath := filepath.Join(opts.dst, rel)

	if info.IsDir() {
		err = os.MkdirAll(dstPath, info.Mode())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: cannot create dst dir - %v\n", dstPath, err)
		}
		return nil
	}

	fmt.Print(path, ": ")

	if !info.Mode().IsRegular() {
		fmt.Fprintf(os.Stderr, "%s: not a regular file\n", path)
		return nil
	}

	var msg string
	dstFile, err := os.Stat(dstPath)
	if err != nil {
		msg = "missing from dst, copying... "
	} else if opts.update && (info.ModTime().After(dstFile.ModTime()) || info.Size() != dstFile.Size()) {
		msg = "update needed... "
	}

	if msg != "" {
		fmt.Print(msg)
		err = copyFile(path, dstPath)
		if err == nil {
			fmt.Println("done")
		} else {
			fmt.Println(err)
			fmt.Fprintf(os.Stderr, "%s -> %s: %v\n", path, dstPath, err)
		}
	} else {
		fmt.Println("up to date")
	}
	return nil
}

func copyFile(src string, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func checkOpts(opts Opts) error {
	var err error
	opts.src, err = filepath.Abs(opts.src)
	if err != nil {
		return err
	}
	opts.dst, err = filepath.Abs(opts.dst)
	if err != nil {
		return err
	}

	if _, err := os.Stat(opts.src); os.IsNotExist(err) {
		return fmt.Errorf("src '%s' does not exist", opts.src)
	}
	if _, err := os.Stat(opts.dst); os.IsNotExist(err) {
		return fmt.Errorf("dst '%s' does not exist", opts.dst)
	}
	return nil
}
