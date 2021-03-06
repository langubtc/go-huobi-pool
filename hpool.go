package hpool

import (
	"errors"
	"fmt"
	"strconv"
)

// 分页查询子账号的收益明细
func (p *SubAccount) ListRecord(page, size int) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"page":       strconv.Itoa(page),
		"size":       strconv.Itoa(size),
	}
	request("GET", p.user.secretKey,
		"/open/api/account/v1/list-record", params)
}

// 查询子账号的实时算力
func (p *SubAccount) GetHashRate() (*HashRates, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/get-hash-rate", params)
	if err != nil {
		return nil, err
	}
	r := HashRatesResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 查询子账号的矿工统计
func (p *SubAccount) GetWorkerStats(coin string) (*WorkerStats, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"coin_name":  coin,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/get-worker-stats", params)
	if err != nil {
		return nil, err
	}
	r := WorkerStatsResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 今日预估收益
func (p *SubAccount) GetTodayProfit() (*TodayProfit, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/get-today-profit", params)
	if err != nil {
		return nil, err
	}
	r := TodayProfitResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 今日预估收益V2
func (p *SubAccount) GetTodayProfitV2() ([]TodayProfit, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v2/get-today-profit", params)
	if err != nil {
		return nil, err
	}
	r := TodayProfitResultV2{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s (body: %s)", err, res))
	}
	return r.Data, nil
}

// 切换挖矿币种
func (p *SubAccount) ChangeCoin(coin string) (bool, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"currency":   coin,
	}
	res, err := request("POST", p.user.secretKey,
		"/open/api/user/v1/change-sub-user-currency", params)
	if err != nil {
		return false, err
	}
	r := ChangeCoinResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 查询子账号下的矿工
func (p *SubAccount) GetWorkers() (*WorkerList, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"status":     "1",
		"page":       "1",
		"size":       "10",
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v2/get-worker", params)
	if err != nil {
		return nil, err
	}
	r := WorkersResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 查询用户历史转让的收益
func (p *SubAccount) GetTransferProfit(date, currency string) (*TransferProfit, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"date":       date,
		"currency":   currency,
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/transfer-profit", params)
	if err != nil {
		return nil, err
	}
	r := TransferProfitResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r.Data, nil
}

// 查询单位理论收益和平均网络手续费
func (p *SubAccount) GetUnitCurrencyProfits(date, currency string) (UnitCurrencyProfits, error) {
	params := map[string]string{
		"access_key": p.user.accessKey,
		"sub_code":   p.SubCode,
		"date":       date,
	}
	if currency != "" {
		params["coin_name"] = currency
	}
	res, err := request("GET", p.user.secretKey,
		"/open/api/user/v1/get-user-unit-currency-profit", params)
	if err != nil {
		return nil, err
	}
	r := UnitCurrencyProfitResult{}
	err = unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
