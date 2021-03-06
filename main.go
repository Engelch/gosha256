package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	ce "github.com/engelch/go_libs/v2"
	"github.com/urfave/cli/v2"
)

const appVersion = "0.7.0"
const appName = "gosha256"

const _debug = "debug" // long (normal) name of CLI option
const _raw = "raw"     // long (normal) name of CLI option for mode, write signature not base64 encoded

// commandLineOptions just separates the definition of command line options ==> creating a shorter main
func commandLineOptions(portNumber *uint, privKeyFile *string, publicKeyFile *string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    _debug,
			Aliases: []string{"d"},
			Value:   false,
			Usage:   "OPTIONAL: enable debug",
		},
		&cli.BoolFlag{
			Name:    _raw,
			Aliases: []string{"r"},
			Value:   false,
			Usage:   "Write digest in pure form, not hex-encoded which is the default",
		},
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcSha256(filename string, raw bool) {
	dat, err := os.ReadFile(filename)
	check(err)
	h := sha256.New()
	h.Write(dat)
	sha256_hash := (h.Sum(nil))
	ce.CondDebugln("SHA256 is:", hex.EncodeToString(sha256_hash))
	if raw {
		_, _ = os.Stdout.Write(sha256_hash)
	} else { // pretty print stdin without filename
		if filename != "/dev/stdin" {
			fmt.Printf("%x %s\n", sha256_hash, filename)
		} else {
			fmt.Printf("%x\n", sha256_hash)
		}
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
	app.Usage = appName + " [--raw | -r] [--debug | -d] [<<file>>]...\n\n\tIf no file is given, then read from stdin.\n\t" +
		"The SHA-256 is output as hex encoded by default.\n\tUsing, raw, it is output as a 32 byte sequence."

	app.Action = func(c *cli.Context) error {
		switch c.NArg() {
		case 0:
			calcSha256("/dev/stdin", c.Bool(_raw))
		default:
			if c.Bool(_raw) {
				return errors.New("Multiple files digest calculation with raw output not supported.")
			}
			for index, _ := range make([]int, c.NArg()) {
				calcSha256(c.Args().Get(index), c.Bool(_raw))
			}
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}

// eof
