package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/autobrr/go-rtorrent"
	"github.com/thoas/go-funk"
	"os"
)

func main() {
	client := getRtorrentClient()
	torrents, err := client.GetTorrents(context.Background(), rtorrent.ViewMain)
	if err != nil {
		fmt.Printf("Error getting torrents: %v\n", err)
		return
	}
	size := funk.Reduce(torrents, func(acc int64, cur rtorrent.Torrent) int64 {
		return acc + int64(cur.Size)
	}, 0)

	fmt.Printf("%d\n", size)
}

func getRtorrentClient() *rtorrent.Client {
	flag.Usage = usage
	username := flag.String("username", "", "HTTP Basic Authentication username")
	pass := flag.String("password", "", "HTTP Basic Authentication password")

	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		flag.Usage()
		os.Exit(1)
		return nil
	}
	cfg := rtorrent.Config{Addr: url}
	if username != nil {
		cfg.BasicUser = *username
	}
	if pass != nil {
		cfg.BasicPass = *pass
	}

	client := rtorrent.NewClient(cfg)
	return client
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] url\n", os.Args[0])
	flag.PrintDefaults()
}
