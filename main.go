package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

const appVersion = "0.0.2"
const appName = "gosha256"

// commandLineOptions just separates the definition of command line options ==> creating a shorter main
func commandLineOptions(portNumber *uint, privKeyFile *string, publicKeyFile *string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug, d",
			Usage: "OPTIONAL: enable debug",
		},
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// main start routine
func main() {
	var portNumber uint = 65536 // set to non acceptable value, if unchanged then a CLI option was missing => error&exit
	var privKeyFile string
	var publicKeyFile string

	app := cli.NewApp()
	app.Flags = commandLineOptions(&portNumber, &privKeyFile, &publicKeyFile)
	app.Name = appName
	app.Version = appVersion
	app.Usage = "TODO usage"

	app.Action = func(c *cli.Context) error {
		var filename string
		if c.NArg() == 1 {
			filename = c.Args().Get(0)
		} else {
			return errors.New("Must be called with 1 arg")
		}
		dat, err := os.ReadFile(filename)
		check(err)
		h := sha256.New()
		h.Write(dat)
		sha256_hash := hex.EncodeToString(h.Sum(nil))

		fmt.Println("SHA256 is:", sha256_hash)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}

// eof
