package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	nodeHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}
VERSION:
   {{.Version}}
   {{end}}
`

	// portFlag defines a flag for setting the port on which the node will listen for connections
	portFlag = cli.IntFlag{
		Name:  "port",
		Usage: "The `[p2p port]` number on which the application will start",
		Value: 0,
	}

	// initialConnectionFlag defines a flag for the initial connection of the node
	// to find more peers
	initialConnectionFlag = cli.StringFlag{
		Name: "initial-connection",
		Usage: "The initial connection this node will try to dial to find more" +
			"peers. It will be something like /ip4/10.19.0.5/tcp/33819/p2p/16Uiu2HAkwNw1TQA4u6xasJmWaSyWqxRZAGSx4kYdo22cJhSuVaHZ",
		Value: "",
	}

	// seedFlag defines a flag for the seed string used in identity generation
	seedFlag = cli.StringFlag{
		Name: "seed",
		Usage: "The seed used in identity generation. If this field is left blank, the " +
			"identity will be randomly generated",
		Value: "",
	}
)

func main() {
	app := cli.NewApp()
	cli.AppHelpTemplate = nodeHelpTemplate
	app.Name = "Elrond Node CLI App"
	app.Version = "0.0.0"
	app.Usage = "This is the entry point for starting a new Elrond p2p connection application"
	app.Flags = []cli.Flag{
		portFlag,
		initialConnectionFlag,
		seedFlag,
	}

	app.Authors = []cli.Author{
		{
			Name:  "The Elrond Team",
			Email: "contact@elrond.com",
		},
	}

	app.Action = startNode

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		os.Exit(1)
	}
}

func startNode(c *cli.Context) error {

}
