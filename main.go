// Package main imports the SIFI Daemon swagger documentation and writes it to a
// file. This requires access to SofIron's private git instance, and is not
// intended to be used outside of SoftIron.
package main

import (
	"os"

	"git.softiron.com/sw/hc/sifi.git/cmd/sifid/docs"
)

func main() {
	fout, err := os.OpenFile("swagger.json", os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}

	defer fout.Close()

	doc := docs.SwaggerInfo.ReadDoc()
	if _, err := fout.WriteString(doc); err != nil {
		panic(err)
	}
}
