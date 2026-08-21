// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryPriceStockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryPriceStockResponseBodyData) *TicketQueryPriceStockResponseBody
	GetData() *TicketQueryPriceStockResponseBodyData
	SetErrorCode(v string) *TicketQueryPriceStockResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryPriceStockResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryPriceStockResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryPriceStockResponseBody
	GetSuccess() *bool
}

type TicketQueryPriceStockResponseBody struct {
	Data *TicketQueryPriceStockResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ScenicIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ScenicId不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketQueryPriceStockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBody) GetData() *TicketQueryPriceStockResponseBodyData {
	return s.Data
}

func (s *TicketQueryPriceStockResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryPriceStockResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryPriceStockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryPriceStockResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryPriceStockResponseBody) SetData(v *TicketQueryPriceStockResponseBodyData) *TicketQueryPriceStockResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryPriceStockResponseBody) SetErrorCode(v string) *TicketQueryPriceStockResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBody) SetErrorMsg(v string) *TicketQueryPriceStockResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryPriceStockResponseBody) SetRequestId(v string) *TicketQueryPriceStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryPriceStockResponseBody) SetSuccess(v bool) *TicketQueryPriceStockResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryPriceStockResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryPriceStockResponseBodyData struct {
	CalendarPriceStocks []*TicketQueryPriceStockResponseBodyDataCalendarPriceStocks `json:"CalendarPriceStocks,omitempty" xml:"CalendarPriceStocks,omitempty" type:"Repeated"`
	NormalPriceStock    *TicketQueryPriceStockResponseBodyDataNormalPriceStock      `json:"NormalPriceStock,omitempty" xml:"NormalPriceStock,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	StockType *int32 `json:"StockType,omitempty" xml:"StockType,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyData) GetCalendarPriceStocks() []*TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	return s.CalendarPriceStocks
}

func (s *TicketQueryPriceStockResponseBodyData) GetNormalPriceStock() *TicketQueryPriceStockResponseBodyDataNormalPriceStock {
	return s.NormalPriceStock
}

func (s *TicketQueryPriceStockResponseBodyData) GetProductId() *string {
	return s.ProductId
}

func (s *TicketQueryPriceStockResponseBodyData) GetStockType() *int32 {
	return s.StockType
}

func (s *TicketQueryPriceStockResponseBodyData) SetCalendarPriceStocks(v []*TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) *TicketQueryPriceStockResponseBodyData {
	s.CalendarPriceStocks = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyData) SetNormalPriceStock(v *TicketQueryPriceStockResponseBodyDataNormalPriceStock) *TicketQueryPriceStockResponseBodyData {
	s.NormalPriceStock = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyData) SetProductId(v string) *TicketQueryPriceStockResponseBodyData {
	s.ProductId = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyData) SetStockType(v int32) *TicketQueryPriceStockResponseBodyData {
	s.StockType = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyData) Validate() error {
	if s.CalendarPriceStocks != nil {
		for _, item := range s.CalendarPriceStocks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NormalPriceStock != nil {
		if err := s.NormalPriceStock.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryPriceStockResponseBodyDataCalendarPriceStocks struct {
	// example:
	//
	// 2026-10-01
	Date              *string                                                                    `json:"Date,omitempty" xml:"Date,omitempty"`
	DistributionPrice *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice `json:"DistributionPrice,omitempty" xml:"DistributionPrice,omitempty" type:"Struct"`
	MarketPrice       *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice       `json:"MarketPrice,omitempty" xml:"MarketPrice,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Stock          *int64                                                                  `json:"Stock,omitempty" xml:"Stock,omitempty"`
	SuggestedPrice *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice `json:"SuggestedPrice,omitempty" xml:"SuggestedPrice,omitempty" type:"Struct"`
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GetDate() *string {
	return s.Date
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GetDistributionPrice() *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice {
	return s.DistributionPrice
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GetMarketPrice() *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice {
	return s.MarketPrice
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GetStock() *int64 {
	return s.Stock
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) GetSuggestedPrice() *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice {
	return s.SuggestedPrice
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) SetDate(v string) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	s.Date = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) SetDistributionPrice(v *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	s.DistributionPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) SetMarketPrice(v *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	s.MarketPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) SetStock(v int64) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	s.Stock = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) SetSuggestedPrice(v *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks {
	s.SuggestedPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocks) Validate() error {
	if s.DistributionPrice != nil {
		if err := s.DistributionPrice.Validate(); err != nil {
			return err
		}
	}
	if s.MarketPrice != nil {
		if err := s.MarketPrice.Validate(); err != nil {
			return err
		}
	}
	if s.SuggestedPrice != nil {
		if err := s.SuggestedPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksDistributionPrice) Validate() error {
	return dara.Validate(s)
}

type TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksMarketPrice) Validate() error {
	return dara.Validate(s)
}

type TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataCalendarPriceStocksSuggestedPrice) Validate() error {
	return dara.Validate(s)
}

type TicketQueryPriceStockResponseBodyDataNormalPriceStock struct {
	DistributionPrice *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice `json:"DistributionPrice,omitempty" xml:"DistributionPrice,omitempty" type:"Struct"`
	MarketPrice       *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice       `json:"MarketPrice,omitempty" xml:"MarketPrice,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Stock          *int64                                                               `json:"Stock,omitempty" xml:"Stock,omitempty"`
	SuggestedPrice *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice `json:"SuggestedPrice,omitempty" xml:"SuggestedPrice,omitempty" type:"Struct"`
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStock) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStock) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) GetDistributionPrice() *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice {
	return s.DistributionPrice
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) GetMarketPrice() *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice {
	return s.MarketPrice
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) GetStock() *int64 {
	return s.Stock
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) GetSuggestedPrice() *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice {
	return s.SuggestedPrice
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) SetDistributionPrice(v *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) *TicketQueryPriceStockResponseBodyDataNormalPriceStock {
	s.DistributionPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) SetMarketPrice(v *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) *TicketQueryPriceStockResponseBodyDataNormalPriceStock {
	s.MarketPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) SetStock(v int64) *TicketQueryPriceStockResponseBodyDataNormalPriceStock {
	s.Stock = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) SetSuggestedPrice(v *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) *TicketQueryPriceStockResponseBodyDataNormalPriceStock {
	s.SuggestedPrice = v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStock) Validate() error {
	if s.DistributionPrice != nil {
		if err := s.DistributionPrice.Validate(); err != nil {
			return err
		}
	}
	if s.MarketPrice != nil {
		if err := s.MarketPrice.Validate(); err != nil {
			return err
		}
	}
	if s.SuggestedPrice != nil {
		if err := s.SuggestedPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockDistributionPrice) Validate() error {
	return dara.Validate(s)
}

type TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockMarketPrice) Validate() error {
	return dara.Validate(s)
}

type TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice struct {
	// example:
	//
	// 10000
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) GoString() string {
	return s.String()
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) GetAmount() *int64 {
	return s.Amount
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) SetAmount(v int64) *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice {
	s.Amount = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) SetCurrencyCode(v string) *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice {
	s.CurrencyCode = &v
	return s
}

func (s *TicketQueryPriceStockResponseBodyDataNormalPriceStockSuggestedPrice) Validate() error {
	return dara.Validate(s)
}
