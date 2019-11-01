# cptree

[![GoDoc](https://godoc.org/github.com/wiggin77/cptree?status.svg)](https://godoc.org/github.com/wiggin77/cptree)
[![Build Status](https://travis-ci.org/wiggin77/cptree.svg?branch=master)](https://travis-ci.org/wiggin77/cptree)

cptree copies a tree of files and directories from one location to another.

<!-- markdownlint-disable MD026 -->
### Why create yet another file copy?

I need something that can run on a Linux-based NAS box where the rsync provided has been modified to not work correctly when copying from RAID volume to external drive attached via USB.

```help
./cptree -h
  -dst string
        destination directory
  -h    display help
  -src string
        src directory to copy
  -u    update; copy files newer in src than dst (default true)
  -version
        display version info
```
