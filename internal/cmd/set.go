package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/xUnholy/go-proxy/pkg/execute"
	"github.com/xUnholy/go-proxy/pkg/prompt"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	git "github.com/xUnholy/go-proxy/internal/git"
	npm "github.com/xUnholy/go-proxy/internal/npm"
)

var (
	cntlmFile = "/usr/local/etc/cntlm.conf"
	port      int
	setAll    bool
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
				Description: "This command will set the NPM proxy values. Both https-proxy and proxy will be set",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:        "port, p",
						Value:       3128,
						Usage:       "set custom CNTLM `PORT`",
						Destination: &port,
					},
				},
				Action: func(_ *cli.Context) {
					p := makeProxyURL(port)
					cmds := npm.EnableProxyConfiguration(p)
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Set npm config successfully")
				},
			},
			{
				Name:        "git",
				Usage:       "set git proxy config",
				Description: "This command will set the GIT global proxy values. Both http.proxy and https.proxy will be set",
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:        "port, p",
						Value:       3128,
						Usage:       "set custom CNTLM `PORT`",
						Destination: &port,
					},
				},
				Action: func(_ *cli.Context) {
					p := makeProxyURL(port)
					cmds := git.EnableProxyConfiguration(p)
					_, err := execute.RunCommands(cmds)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("Set git config successfully")
				},
			},
			{
				Name:        "ca-cert",
				Usage:       "set custom CA cert for proxy",
				Description: "This command imports a CA cert for use with other tools (eg. gcloud)",
				Action: func(c *cli.Context) {
					if !c.Args().Present() {
						fmt.Println("You must supply a filepath")
						return
					}

					caCertPath := c.Args().First()

					if _, err := os.Stat(caCertPath); err == nil {
						err := copyFile(caCertPath, fmt.Sprintf("%v/.proxyca", os.Getenv("HOME")))
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("Stored custom CA cert at ~/.proxyca")
					} else {
						fmt.Println("File not found:", caCertPath)
					}
				},
			},
			{
				Name:        "username",
				Usage:       "proxy set username",
				Description: "This command will update the Username value in your CNTLM.conf file",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Username: ")
					output, err := prompt.GetInput()
					if err != nil {
						log.Fatal(err)
					}
					update := fmt.Sprintln("Username\t", output)
					cntlm.UpdateFile(cntlmFile, update)
					fmt.Println("Set CNTLM username successfully")
				},
			},
			{
				Name:        "password",
				Usage:       "proxy set password",
				Description: "This command will update the Password value in your CNTLM.conf file",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Password: ")
					e := execute.Command{Cmd: "cntlm", Args: []string{"-H"}}
					output, err := execute.RunCommand(e)
					if err != nil {
						log.Fatal(err)
					}
					cntlm.UpdateFile(cntlmFile, output)
					fmt.Println("Set CNTLM password successfully")
				},
			},
			{
				Name:        "domain",
				Usage:       "proxy set domain",
				Description: "This command will update the domain value in your CNTLM.conf file",
				Action: func(_ *cli.Context) {
					fmt.Printf("Enter Proxy Domain: ")
					output, err := prompt.GetInput()
					if err != nil {
						log.Fatal(err)
					}
					update := fmt.Sprintln("Domain\t", output)
					cntlm.UpdateFile(cntlmFile, update)
					fmt.Println("Set CNTLM domain successfully")
				},
			},
		},
	}
}

func makeProxyURL(port int) string {
	return fmt.Sprintf("http://localhost:%d", port)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
