package gcloudproxy

import (
	"fmt"
	"os"

	"github.com/xUnholy/go-proxy/pkg/execute"
)

func EnableProxyConfiguration(port string) []execute.Command {
	//cmds := make([]execute.Command{}, 3, 4)
	addr := execute.Command{Cmd: "gcloud", Args: []string{"config", "set", "proxy/address", "127.0.0.1"}}
	port_ := execute.Command{Cmd: "gcloud", Args: []string{"config", "set", "proxy/port", port}}
	http := execute.Command{Cmd: "gcloud", Args: []string{"config", "set", "proxy/type", "http"}}
	cmds := append([]execute.Command{}, addr, port_, http)
	return cmds
}

func GCloudSetCACert(cmds []execute.Command, path string) []execute.Command {
	ca_cert := execute.Command{Cmd: "gcloud", Args: []string{"config", "set", "core/custom_ca_certs_file", fmt.Sprintf("%v/.proxyca", os.Getenv("HOME"))}}
	cmds = append(cmds, ca_cert)
	return cmds
}

func DisableProxyConfiguration() []execute.Command {
	addr := execute.Command{Cmd: "gcloud", Args: []string{"config", "unset", "proxy/address"}}
	port_ := execute.Command{Cmd: "gcloud", Args: []string{"config", "unset", "proxy/port"}}
	http := execute.Command{Cmd: "gcloud", Args: []string{"config", "unset", "proxy/type"}}
	ca_cert := execute.Command{Cmd: "gcloud", Args: []string{"config", "unset", "core/custom_ca_certs_file"}}
	cmds := append([]execute.Command{}, addr, port_, http, ca_cert)
	return cmds
}
