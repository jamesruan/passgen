/*
 * Just generate base64 encoded sha1 word-by-word sum of passphrase.
 * Copyright (C) 2016 James Ruan <ruanbeihong@gmail.com>

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = printUsage
	flag.Parse()

	phrase := flag.Args()

	switch pc := len(phrase); {
	case pc < 1:
		printUsage()
	default:
		printPass(genPass(phrase))
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s passphrase\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func printPass(pass string) {
	fmt.Println(pass)
}

func genPass(phrase []string) string {
	pc := len(phrase)
	acc := sha1.New()
	for i := 0; i < pc; i += 1 {
		io.WriteString(acc, phrase[i])
	}
	sum := acc.Sum(nil)
	return base64.StdEncoding.EncodeToString(sum)
}
