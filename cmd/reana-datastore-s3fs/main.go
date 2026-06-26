/*
This file is part of REANA.
Copyright (C) 2026 CERN.

REANA is free software; you can redistribute it and/or modify it
under the terms of the MIT License; see LICENSE file for more details.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/reanahub/reana-datastore-s3fs/internal/version"
)

func main() {
	showVersion := flag.Bool("version", false, "show version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version.Version)
		return
	}

	fmt.Fprintln(os.Stdout, "reana-datastore-s3fs sidecar scaffold")
}
