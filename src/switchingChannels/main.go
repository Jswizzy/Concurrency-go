package switchingChannels

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"frodo@underhill.me"},
		From:    "gandalf@whitecouncil.org",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message intercepted by black rider",
		OriginalMessage: "Keep it secret, keep it safe",
	}

	msgCh <- msg
	errCh <- failedMessage


	select {
	case receiveMsg := <-msgCh:
		fmt.Println(receiveMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("No messages to print.")
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage string
}
