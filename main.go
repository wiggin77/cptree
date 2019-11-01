package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	// APPNAME is the original application name.
	APPNAME = "cptree"
	// APPVERSION is the application version number.
	APPVERSION = "0.0.1"
)

// Opts contains options from the command line.
type Opts struct {
	update  bool
	perms   bool
	help    bool
	version bool
	src     string
	dst     string
}

func main() {
	opts := Opts{}
	flag.BoolVar(&opts.update, "u", true, "update; copy files newer in src than dst")
	flag.BoolVar(&opts.perms, "p", true, "copy permissions")
	flag.BoolVar(&opts.help, "h", false, "display help")
	flag.BoolVar(&opts.version, "version", false, "display version info")
	flag.StringVar(&opts.src, "src", "", "src directory to copy")
	flag.StringVar(&opts.dst, "dst", "", "destination directory")
	flag.Parse()

	if opts.help {
		flag.PrintDefaults()
		return
	}
	if opts.version {
		printAppInfo(os.Stdout)
	}

	code := 0
	err := cptree(opts)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		code = -1
	}
	os.Exit(code)
}

func printAppInfo(w io.Writer) {
	fmt.Fprintf(w, "%s v%s", APPNAME, APPVERSION)
}
