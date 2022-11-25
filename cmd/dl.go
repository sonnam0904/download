package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	downloader "github.com/sonnn/download"
)

func main() {
	url := flag.String("u", "", "* Url file tải")
	concurrency := flag.Int("n", 3, "Concurrency level")
	filename := flag.String("f", "", "Output file")
	bufferSize := flag.Int("buffer-size", 32*1024, "The buffer size to copy from http response body")
	resume := flag.Bool("resume", false, "Tiếp tục tải")

	flag.Parse()
	if *url == "" {
		log.Fatal("Please specify the url using -u parameter")
	}

	config := &downloader.Config{
		Url:            *url,
		Concurrency:    *concurrency,
		OutFilename:    *filename,
		CopyBufferSize: *bufferSize,
		Resume:         *resume,
	}
	d, err := downloader.NewFromConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	termCh := make(chan os.Signal)
	signal.Notify(termCh, os.Interrupt)
	go func() {
		<-termCh
		println("\nĐang thoát ...")
		d.Pause()
	}()

	d.Download()
	if d.Paused {
		println("\nTải xuống đang tạm dừng. Để tiếp tục hãy sử dụng tham số -resume=true.")
	} else {
		println("Tải xuống hoàn tất.")
	}
}
