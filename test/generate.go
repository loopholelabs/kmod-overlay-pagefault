/*
    Copyright (C) 2024 Loophole Labs

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	pageSize := os.Getpagesize()
	totalSize := pageSize * 1024 * 1024 * 4

	fmt.Printf("using pageSize %d with totalSize %d\n", pageSize, totalSize)

	{
		out, err := os.OpenFile("base.bin", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		if err := out.Truncate(int64(totalSize)); err != nil {
			panic(err)
		}
	}

	{
		out, err := os.OpenFile("overlay.bin", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		in, err := os.Open("/dev/random")
		if err != nil {
			panic(err)
		}
		defer in.Close()

		if _, err := io.CopyN(out, in, int64(totalSize)); err != nil {
			panic(err)
		}
	}
}