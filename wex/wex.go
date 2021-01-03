package wex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	. "github.com/zjmhaoku01/goex"
)

const (
	baseUrl = "https://api.wex.app"
)

type Wex struct {
	httpClient *http.Client
	accessKey  string
	secretKey  string
}

func New(client *http.Client, accessKey, secretKey string) *Wex {
	return &Wex{client, accessKey, secretKey}
}

func (w *Wex) LimitBuy(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	panic("implement me")
}

func (w *Wex) LimitSell(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error) {
	panic("implement me")
}

func (w *Wex) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("implement me")
}

func (w *Wex) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("implement me")
}

func (w *Wex) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
	panic("implement me")
}

func (w *Wex) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
	panic("implement me")
}

func (w *Wex) GetUnfinishOrders(currency CurrencyPair) ([]Order, error) {
	panic("implement me")
}

func (w *Wex) GetOrderHistorys(currency CurrencyPair, opt ...OptionalParameter) ([]Order, error) {
	panic("implement me")
}

func (w *Wex) GetAccount() (*Account, error) {
	data, err := w.doRequest("GET", "/api/v1/private/asset", &url.Values{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("data:%v", data)
	return nil, nil
}

func (w *Wex) GetTicker(currency CurrencyPair) (*Ticker, error) {
	data, err := w.doRequest("GET", "/api/v1/public/ticker", &url.Values{"market": []string{currency.ToSymbol("_")}})
	if err != nil {
		return nil, err
	}
	return &Ticker{
		Date: ToUint64(data["time"]) / 1000,
		Last: ToFloat64(data["last"]),
		Buy:  ToFloat64(data["buy"]),
		Sell: ToFloat64(data["sell"]),
		High: ToFloat64(data["high"]),
		Low:  ToFloat64(data["low"]),
		Vol:  ToFloat64(data["vol"])}, nil
}

func (w *Wex) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	panic("implement me")
}

func (w *Wex) GetKlineRecords(currency CurrencyPair, period KlinePeriod, size int, optional ...OptionalParameter) ([]Kline, error) {
	panic("implement me")
}

func (w *Wex) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
	panic("implement me")
}

func (w *Wex) GetExchangeName() string {
	return WEX
}

func (w *Wex) doRequest(method, uri string, params *url.Values) (map[string]interface{}, error) {
	resp, err := w.doRequestInner(method, uri, params)

	if err != nil {
		return nil, err
	}

	retMap := make(map[string]interface{})
	err = json.Unmarshal(resp, &retMap)
	if err != nil {
		return nil, err
	}

	if _, ok := retMap["data"]; !ok {
		return nil, fmt.Errorf("doRequest: failed, retMap:%v", retMap)
	}

	dataMap := retMap["data"].(map[string]interface{})

	return dataMap, nil
}

func (w *Wex) doRequestInner(method, uri string, params *url.Values) (buf []byte, err error) {
	reqUrl := baseUrl + uri

	headerMap := map[string]string{
		"Content-Type": "application/json; charset=utf-8"}

	// 添加token
	if strings.HasPrefix(uri, "/api/v1/private") {
		params.Set("access_key", w.accessKey)
		params.Set("nonce", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
		signData := params.Encode()
		mac := hmac.New(sha256.New, []byte(w.secretKey))
		mac.Write([]byte(uri + "|" + signData))
		sign := hex.EncodeToString(mac.Sum(nil))
		params.Set("sign", sign)
	}

	if ("GET" == method || "DELETE" == method) && len(params.Encode()) > 0 {
		reqUrl += "?" + params.Encode()
	}

	var paramStr = ""
	if "POST" == method {
		//to json
		paramStr = params.Encode()
		var paraMap map[string]string = make(map[string]string, 2)
		for _, v := range strings.Split(paramStr, "&") {
			vv := strings.Split(v, "=")
			paraMap[vv[0]] = vv[1]
		}
		jsonData, _ := json.Marshal(paraMap)
		paramStr = string(jsonData)
	}

	return NewHttpRequest(w.httpClient, method, reqUrl, paramStr, headerMap)
}
