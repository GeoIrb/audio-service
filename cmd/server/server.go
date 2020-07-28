package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/GeoIrb/sound-ethernet-streaming/pkg/converter"
	"github.com/GeoIrb/sound-ethernet-streaming/pkg/server"
	udp "github.com/GeoIrb/sound-ethernet-streaming/pkg/udp/server"
	"github.com/GeoIrb/sound-ethernet-streaming/pkg/wav"
)

const (
	sizeData = 5
	dstAddr  = "255.255.255.255:8080"
	file     = "/home/geo/go/src/github.com/GeoIrb/sound-ethernet-streaming/audio/test.wav"
)

func main() {
	var (
		err  error
		data []byte
	)
	udpSrv := udp.NewServerUDP(dstAddr)
	if err = udpSrv.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer udpSrv.Shutdown()

	c7v := converter.NewConverter()
	s4v := server.NewServer(
		udpSrv,
		c7v,
		sizeData,
	)

	if data, err = ioutil.ReadFile(file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	audio := wav.NewWAV()
	if err = audio.Parse(data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go s4v.Streaming(ctx, audio)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	sig := <-c
	fmt.Printf("received signal, exiting signal %v\n", sig)
	cancel()
}
