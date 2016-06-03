package zmq4_test

import (
	zmq "github.com/pebbe/zmq4"

	"testing"
	"time"
)

func TestUdp(t *testing.T) {

	if _, minor, _ := zmq.Version(); minor < 2 {
		t.Skip("Sockets RADIO and DISH not avalable in ZeroMQ versions prior to 4.2")
	}

	ctx, err := zmq.NewContext()
	if err != nil {
		t.Fatal("NewContext:", err)
	}
	defer ctx.Term()

	radio, err := ctx.NewSocket(zmq.RADIO)
	if err != nil {
		t.Fatal("NewSocket RADIO:", err)
	}
	defer radio.Close()
	dish, err := ctx.NewSocket(zmq.DISH)
	if err != nil {
		t.Fatal("NewSocket DISH:", err)
	}
	defer dish.Close()

	//  Connecting dish should fail
	err = dish.Connect("udp://127.0.0.1:5556")
	if err == nil {
		t.Fatal("Expected fail on dish.Connect")
	}

	err = dish.Bind("udp://*:5556")
	if err != nil {
		t.Fatal("dish.Bind:", err)
	}

	//  Bind radio should fail
	err = radio.Bind("udp://*:5556")
	if err == nil {
		t.Fatal("Expected fail on radio.Bind")
	}

	err = radio.Connect("udp://127.0.0.1:5556")
	if err != nil {
		t.Fatal("radio.Connect:", err)
	}

	time.Sleep(300 * time.Millisecond)

	err = dish.Join("TV")
	if err != nil {
		t.Fatal("dish.Join:", err)
	}

	_, err = radio.Send("Friends", 0, zmq.OptSetGroup("TV"))
	if err != nil {
		t.Fatal("radio.SendMessage:", err)
	}

	msg, opt, err := dish.RecvWithOpts(0, zmq.OptGetGroup(true))
	if err != nil {
		t.Fatal("dish.RecvWithOpt:", err)
	}
	if len(opt) != 1 {
		t.Fatal("dish.RecvWithOpt: wrong number off options")
	}
	if opt[0].(string) != "TV" {
		t.Fatal("dish.RecvWithOpt: wrong group: %v", opt[0])
	}
	if msg != "Friends" {
		t.Fatal("dish.RecvWithOpt: wrong message: %q", msg)
	}
}
