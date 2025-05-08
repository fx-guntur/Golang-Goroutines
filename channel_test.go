package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
	biasakan kalau bikin channel langsung sekalian closenya
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Mega"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	defer close(channel)
}
