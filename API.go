package goex

type API interface {
	// 挂单买
	LimitBuy(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error)
	// 挂单卖
	LimitSell(amount, price string, currency CurrencyPair, opt ...LimitOrderOptionalParameter) (*Order, error)
	// 市场价买
	MarketBuy(amount, price string, currency CurrencyPair) (*Order, error)
	// 市场价卖
	MarketSell(amount, price string, currency CurrencyPair) (*Order, error)
	// 取消挂单
	CancelOrder(orderId string, currency CurrencyPair) (bool, error)
	// 获取订单信息
	GetOneOrder(orderId string, currency CurrencyPair) (*Order, error)
	// 获取未完成订单
	GetUnfinishOrders(currency CurrencyPair) ([]Order, error)
	// 获取历史订单
	GetOrderHistorys(currency CurrencyPair, opt ...OptionalParameter) ([]Order, error)
	// 获取账户信息
	GetAccount() (*Account, error)
	// 获取标的价格
	GetTicker(currency CurrencyPair) (*Ticker, error)
	// 获取交易深度
	GetDepth(size int, currency CurrencyPair) (*Depth, error)
	// 获取k线
	GetKlineRecords(currency CurrencyPair, period KlinePeriod, size int, optional ...OptionalParameter) ([]Kline, error)
	// 非个人，整个交易所的交易记录
	GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error)
	// 获取当前实体的对应的交易所
	GetExchangeName() string
}
