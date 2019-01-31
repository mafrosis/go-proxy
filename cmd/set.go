package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	password string
)

func SetCommand() cli.Command {
	return cli.Command{
		Name:        "set",
		Aliases:     []string{""},
		Usage:       "proxy set",
		Description: "Set CNTLM Proxy Config",
		Subcommands: []cli.Command{
			{
				Name:        "npm",
				Usage:       "set npm proxy config",
				Description: "additional description?",
				Action: func(c *cli.Context) {
					fmt.Println("new task template: ", c.Args().First())
				},
			},
			{
				Name:  "gradle",
				Usage: "set gradle proxy config",
				Action: func(c *cli.Context) {
					fmt.Println("new task template: ", c.Args().First())
				},
			},
			{
				Name:  "git",
				Usage: "set git proxy config",
				Action: func(c *cli.Context) {
					fmt.Println("new task template: ", c.Args().First())
				},
			},
			{
				Name:  "bash",
				Usage: "set bash profile proxy config",
				Action: func(c *cli.Context) {
					fmt.Println("new task template: ", c.Args().First())
				},
			},
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "password, p",
				Usage:       "set CNTLM `PASSWORD` config",
				Value:       "",
				Destination: &password,
			},
		},
		Action: func(c *cli.Context) {
			fmt.Println("All Command Executed: ", c.Args().First())
		},
	}
}
