package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func pruneNetworks(cli *client.Client) {
	fmt.Println("started to delete unused networks")
	networksPruneReport, err := cli.NetworksPrune(context.Background(), filters.NewArgs())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range networksPruneReport.NetworksDeleted {
		fmt.Println("deleted network ", v)
	}
	fmt.Println("deleted ", len(networksPruneReport.NetworksDeleted), " networks")

}

func pruneImages(cli *client.Client) {
	fmt.Println("deleting unused images")
	imagePruneReport, err := cli.ImagesPrune(context.Background(), filters.NewArgs())

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range imagePruneReport.ImagesDeleted {
		fmt.Println("deleted image ", v)
	}
	fmt.Println("deleted ", len(imagePruneReport.ImagesDeleted), " images")
	fmt.Println("freed ", imagePruneReport.SpaceReclaimed/(1024*1024), " MiB")
}

func main() {
	args := os.Args[1:]
	networks := false
	images := true
	if len(args) != 0 {
		if args[0] == "all" {
			networks = true
			images = true
		} else if args[0] == "networks" {
			networks = true
			images = false
		} else {
			fmt.Println("unknown command, usage : dockerclean all|networks (default images only)")
			return
		}
	}
	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	if networks {
		pruneNetworks(cli)
	}
	if images {
		pruneImages(cli)
	}
}
