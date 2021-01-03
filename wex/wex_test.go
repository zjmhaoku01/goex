package wex

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/zjmhaoku01/goex"
)

var w = New(http.DefaultClient, "", "")

func TestGetTicker(t *testing.T) {
	params := url.Values{"key": []string{"1"}, "kee": []string{"hha"}}
	fmt.Println(params.Encode())
	ticker, err := w.GetTicker(goex.BTC_USDT)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("ticker:%v\n", ticker)
}

func TestGetAccount(t *testing.T) {
	_, err := w.GetAccount()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
