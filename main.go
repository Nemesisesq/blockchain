package main

func main() {
	bc := NewBlockchain()

	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more BTC to Ivan")

	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
