// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkuPriceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySkuPriceListResponseBody
	GetCode() *string
	SetData(v *QuerySkuPriceListResponseBodyData) *QuerySkuPriceListResponseBody
	GetData() *QuerySkuPriceListResponseBodyData
	SetMessage(v string) *QuerySkuPriceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySkuPriceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySkuPriceListResponseBody
	GetSuccess() *bool
}

type QuerySkuPriceListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *QuerySkuPriceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5B803CF-94D8-43AF-ADB3-D819AAD30E27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySkuPriceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySkuPriceListResponseBody) GetData() *QuerySkuPriceListResponseBodyData {
	return s.Data
}

func (s *QuerySkuPriceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySkuPriceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySkuPriceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySkuPriceListResponseBody) SetCode(v string) *QuerySkuPriceListResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySkuPriceListResponseBody) SetData(v *QuerySkuPriceListResponseBodyData) *QuerySkuPriceListResponseBody {
	s.Data = v
	return s
}

func (s *QuerySkuPriceListResponseBody) SetMessage(v string) *QuerySkuPriceListResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySkuPriceListResponseBody) SetRequestId(v string) *QuerySkuPriceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySkuPriceListResponseBody) SetSuccess(v bool) *QuerySkuPriceListResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySkuPriceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySkuPriceListResponseBodyData struct {
	// The SKUs of the pricing object.
	SkuPricePage *QuerySkuPriceListResponseBodyDataSkuPricePage `json:"SkuPricePage,omitempty" xml:"SkuPricePage,omitempty" type:"Struct"`
}

func (s QuerySkuPriceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBodyData) GetSkuPricePage() *QuerySkuPriceListResponseBodyDataSkuPricePage {
	return s.SkuPricePage
}

func (s *QuerySkuPriceListResponseBodyData) SetSkuPricePage(v *QuerySkuPriceListResponseBodyDataSkuPricePage) *QuerySkuPriceListResponseBodyData {
	s.SkuPricePage = v
	return s
}

func (s *QuerySkuPriceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QuerySkuPriceListResponseBodyDataSkuPricePage struct {
	// The token that is used to query the next page.
	//
	// example:
	//
	// 080112060a0422020800180022490a470342000000315333303332363436363336333433393636333136333338333733373333333133373336363336323634363336363337333836333636333636313336363433363332
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The SKUs.
	SkuPriceList []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList `json:"SkuPriceList,omitempty" xml:"SkuPriceList,omitempty" type:"Repeated"`
	// The total number of SKUs.
	//
	// example:
	//
	// 18732
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePage) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePage) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) GetSkuPriceList() []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList {
	return s.SkuPriceList
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) SetNextPageToken(v string) *QuerySkuPriceListResponseBodyDataSkuPricePage {
	s.NextPageToken = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) SetSkuPriceList(v []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) *QuerySkuPriceListResponseBodyDataSkuPricePage {
	s.SkuPriceList = v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) SetTotalCount(v int32) *QuerySkuPriceListResponseBodyDataSkuPricePage {
	s.TotalCount = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePage) Validate() error {
	return dara.Validate(s)
}

type QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList struct {
	// The prices of the SKUs.
	CskuPriceList []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList `json:"CskuPriceList,omitempty" xml:"CskuPriceList,omitempty" type:"Repeated"`
	// The code of the SKU.
	//
	// example:
	//
	// 017c15a31507bc6de22aa93777461adc
	SkuCode *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	// The values of the pricing factors.
	SkuFactorMap map[string]*string `json:"SkuFactorMap,omitempty" xml:"SkuFactorMap,omitempty"`
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) GetCskuPriceList() []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	return s.CskuPriceList
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) GetSkuCode() *string {
	return s.SkuCode
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) GetSkuFactorMap() map[string]*string {
	return s.SkuFactorMap
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) SetCskuPriceList(v []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList {
	s.CskuPriceList = v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) SetSkuCode(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList {
	s.SkuCode = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) SetSkuFactorMap(v map[string]*string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList {
	s.SkuFactorMap = v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceList) Validate() error {
	return dara.Validate(s)
}

type QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList struct {
	// The unique code of the SKU price.
	//
	// example:
	//
	// ac74dd7b52ae6389ddef099283fb8275
	CskuCode *string `json:"CskuCode,omitempty" xml:"CskuCode,omitempty"`
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The price.
	//
	// example:
	//
	// 100
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// The pricing mode.
	//
	// example:
	//
	// STEP_ARRIVE
	PriceMode *string `json:"PriceMode,omitempty" xml:"PriceMode,omitempty"`
	// The pricing type.
	//
	// example:
	//
	// hourPrice
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// The unit of the price.
	//
	// example:
	//
	// USD (per unit)
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// If the PriceMode parameter is set to STEP_ACCUMULATION or STEP_ARRIVE, the value of this field exists and specifies the range. If the PriceMode parameter is set to NORMAL_PRICE, the value of this field is null.
	RangeList []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList `json:"RangeList,omitempty" xml:"RangeList,omitempty" type:"Repeated"`
	// The usage unit.
	//
	// example:
	//
	// Count
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetCskuCode() *string {
	return s.CskuCode
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetCurrency() *string {
	return s.Currency
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetPrice() *string {
	return s.Price
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetPriceMode() *string {
	return s.PriceMode
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetPriceType() *string {
	return s.PriceType
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetRangeList() []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList {
	return s.RangeList
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetCskuCode(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.CskuCode = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetCurrency(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.Currency = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetPrice(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.Price = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetPriceMode(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.PriceMode = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetPriceType(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.PriceType = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetPriceUnit(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.PriceUnit = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetRangeList(v []*QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.RangeList = v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) SetUsageUnit(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList {
	s.UsageUnit = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceList) Validate() error {
	return dara.Validate(s)
}

type QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList struct {
	// The code of the pricing factor.
	//
	// example:
	//
	// 212fbd27866307fc79ecf06934a88b2c
	FactorCode *string `json:"FactorCode,omitempty" xml:"FactorCode,omitempty"`
	// The maximum value.
	//
	// example:
	//
	// 10
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum value.
	//
	// example:
	//
	// 1
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
	// The closure type of the interval.
	//
	// example:
	//
	// LORC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) String() string {
	return dara.Prettify(s)
}

func (s QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) GoString() string {
	return s.String()
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) GetFactorCode() *string {
	return s.FactorCode
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) GetMax() *string {
	return s.Max
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) GetMin() *string {
	return s.Min
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) GetType() *string {
	return s.Type
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) SetFactorCode(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList {
	s.FactorCode = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) SetMax(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList {
	s.Max = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) SetMin(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList {
	s.Min = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) SetType(v string) *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList {
	s.Type = &v
	return s
}

func (s *QuerySkuPriceListResponseBodyDataSkuPricePageSkuPriceListCskuPriceListRangeList) Validate() error {
	return dara.Validate(s)
}
