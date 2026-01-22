// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrdersResponseBody
	GetCode() *string
	SetData(v *GetOrdersResponseBodyData) *GetOrdersResponseBody
	GetData() *GetOrdersResponseBodyData
	SetMessage(v string) *GetOrdersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrdersResponseBody
	GetSuccess() *bool
}

type GetOrdersResponseBody struct {
	// example:
	//
	// Success
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOrdersResponseBody) GetData() *GetOrdersResponseBodyData {
	return s.Data
}

func (s *GetOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrdersResponseBody) SetCode(v string) *GetOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrdersResponseBody) SetData(v *GetOrdersResponseBodyData) *GetOrdersResponseBody {
	s.Data = v
	return s
}

func (s *GetOrdersResponseBody) SetMessage(v string) *GetOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrdersResponseBody) SetRequestId(v string) *GetOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrdersResponseBody) SetSuccess(v bool) *GetOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrdersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrdersResponseBodyData struct {
	// example:
	//
	// test
	HostName  *string                             `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OrderList *GetOrdersResponseBodyDataOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrdersResponseBodyData) GetHostName() *string {
	return s.HostName
}

func (s *GetOrdersResponseBodyData) GetOrderList() *GetOrdersResponseBodyDataOrderList {
	return s.OrderList
}

func (s *GetOrdersResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetOrdersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOrdersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetOrdersResponseBodyData) SetHostName(v string) *GetOrdersResponseBodyData {
	s.HostName = &v
	return s
}

func (s *GetOrdersResponseBodyData) SetOrderList(v *GetOrdersResponseBodyDataOrderList) *GetOrdersResponseBodyData {
	s.OrderList = v
	return s
}

func (s *GetOrdersResponseBodyData) SetPageNum(v int32) *GetOrdersResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOrdersResponseBodyData) SetPageSize(v int32) *GetOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOrdersResponseBodyData) SetTotalCount(v int32) *GetOrdersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetOrdersResponseBodyData) Validate() error {
	if s.OrderList != nil {
		if err := s.OrderList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrdersResponseBodyDataOrderList struct {
	Order []*GetOrdersResponseBodyDataOrderListOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s GetOrdersResponseBodyDataOrderList) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *GetOrdersResponseBodyDataOrderList) GetOrder() []*GetOrdersResponseBodyDataOrderListOrder {
	return s.Order
}

func (s *GetOrdersResponseBodyDataOrderList) SetOrder(v []*GetOrdersResponseBodyDataOrderListOrder) *GetOrdersResponseBodyDataOrderList {
	s.Order = v
	return s
}

func (s *GetOrdersResponseBodyDataOrderList) Validate() error {
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

type GetOrdersResponseBodyDataOrderListOrder struct {
	// example:
	//
	// 66
	AfterTaxAmount *string `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2017-06-08T09:41:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 237258627070169
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// New
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// CNY
	PaymentCurrency *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	// example:
	//
	// Paid
	PaymentStatus *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	// example:
	//
	// 2017-06-08T09:41:30Z
	PaymentTime *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	// example:
	//
	// 0
	PretaxAmount *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// example:
	//
	// 0
	PretaxAmountLocal *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	// example:
	//
	// 0
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 234535345345342
	RelatedOrderId *string `json:"RelatedOrderId,omitempty" xml:"RelatedOrderId,omitempty"`
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// example:
	//
	// 1
	Tax *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
}

func (s GetOrdersResponseBodyDataOrderListOrder) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersResponseBodyDataOrderListOrder) GoString() string {
	return s.String()
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetAfterTaxAmount() *string {
	return s.AfterTaxAmount
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetCurrency() *string {
	return s.Currency
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetOrderType() *string {
	return s.OrderType
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPretaxAmountLocal() *string {
	return s.PretaxAmountLocal
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetPretaxGrossAmount() *string {
	return s.PretaxGrossAmount
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetProductType() *string {
	return s.ProductType
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetRelatedOrderId() *string {
	return s.RelatedOrderId
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *GetOrdersResponseBodyDataOrderListOrder) GetTax() *string {
	return s.Tax
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetAfterTaxAmount(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.AfterTaxAmount = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetCommodityCode(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.CommodityCode = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetCreateTime(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.CreateTime = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetCurrency(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.Currency = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetOrderId(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.OrderId = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetOrderType(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.OrderType = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPaymentCurrency(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PaymentCurrency = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPaymentStatus(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PaymentStatus = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPaymentTime(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PaymentTime = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPretaxAmount(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmount = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPretaxAmountLocal(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmountLocal = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetPretaxGrossAmount(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.PretaxGrossAmount = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetProductCode(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.ProductCode = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetProductType(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.ProductType = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetRelatedOrderId(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.RelatedOrderId = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetSubscriptionType(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.SubscriptionType = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) SetTax(v string) *GetOrdersResponseBodyDataOrderListOrder {
	s.Tax = &v
	return s
}

func (s *GetOrdersResponseBodyDataOrderListOrder) Validate() error {
	return dara.Validate(s)
}
