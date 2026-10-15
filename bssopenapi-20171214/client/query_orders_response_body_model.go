// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryOrdersResponseBody
	GetCode() *string
	SetData(v *QueryOrdersResponseBodyData) *QueryOrdersResponseBody
	GetData() *QueryOrdersResponseBodyData
	SetMessage(v string) *QueryOrdersResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryOrdersResponseBody
	GetSuccess() *bool
}

type QueryOrdersResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *QueryOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 61293E7A-3406-4447-8620-EC88B0AA66AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryOrdersResponseBody) GetData() *QueryOrdersResponseBodyData {
	return s.Data
}

func (s *QueryOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrdersResponseBody) SetCode(v string) *QueryOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrdersResponseBody) SetData(v *QueryOrdersResponseBodyData) *QueryOrdersResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrdersResponseBody) SetMessage(v string) *QueryOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrdersResponseBody) SetRequestId(v string) *QueryOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrdersResponseBody) SetSuccess(v bool) *QueryOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrdersResponseBodyData struct {
	// The hostname.
	//
	// example:
	//
	// test
	HostName  *string                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OrderList *QueryOrdersResponseBodyDataOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Set this parameter to a maximum of 20. A larger value may cause a timeout.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyData) GetHostName() *string {
	return s.HostName
}

func (s *QueryOrdersResponseBodyData) GetOrderList() *QueryOrdersResponseBodyDataOrderList {
	return s.OrderList
}

func (s *QueryOrdersResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryOrdersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOrdersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryOrdersResponseBodyData) SetHostName(v string) *QueryOrdersResponseBodyData {
	s.HostName = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetOrderList(v *QueryOrdersResponseBodyDataOrderList) *QueryOrdersResponseBodyData {
	s.OrderList = v
	return s
}

func (s *QueryOrdersResponseBodyData) SetPageNum(v int32) *QueryOrdersResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetPageSize(v int32) *QueryOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetTotalCount(v int32) *QueryOrdersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryOrdersResponseBodyData) Validate() error {
	if s.OrderList != nil {
		if err := s.OrderList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrdersResponseBodyDataOrderList struct {
	Order []*QueryOrdersResponseBodyDataOrderListOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s QueryOrdersResponseBodyDataOrderList) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyDataOrderList) GetOrder() []*QueryOrdersResponseBodyDataOrderListOrder {
	return s.Order
}

func (s *QueryOrdersResponseBodyDataOrderList) SetOrder(v []*QueryOrdersResponseBodyDataOrderListOrder) *QueryOrdersResponseBodyDataOrderList {
	s.Order = v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderList) Validate() error {
	if s.Order != nil {
		for _, item := range s.Order {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryOrdersResponseBodyDataOrderListOrder struct {
	AfterTaxAmount    *string `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	CommodityCode     *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Currency          *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PaymentCurrency   *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	PaymentStatus     *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PaymentTime       *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	PretaxAmount      *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxAmountLocal *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelatedOrderId    *string `json:"RelatedOrderId,omitempty" xml:"RelatedOrderId,omitempty"`
	SubscriptionType  *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tax               *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
}

func (s QueryOrdersResponseBodyDataOrderListOrder) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersResponseBodyDataOrderListOrder) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetAfterTaxAmount() *string {
	return s.AfterTaxAmount
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPretaxAmountLocal() *string {
	return s.PretaxAmountLocal
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetPretaxGrossAmount() *string {
	return s.PretaxGrossAmount
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetProductType() *string {
	return s.ProductType
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetRelatedOrderId() *string {
	return s.RelatedOrderId
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) GetTax() *string {
	return s.Tax
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetAfterTaxAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCommodityCode(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.CommodityCode = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCreateTime(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.CreateTime = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCurrency(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.Currency = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetOrderId(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.OrderId = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetOrderType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.OrderType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentCurrency(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentStatus(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentStatus = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentTime(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentTime = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxAmountLocal(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxGrossAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetProductCode(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.ProductCode = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetProductType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.ProductType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetRelatedOrderId(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.RelatedOrderId = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetSubscriptionType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.SubscriptionType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetTax(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.Tax = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) Validate() error {
	return dara.Validate(s)
}
