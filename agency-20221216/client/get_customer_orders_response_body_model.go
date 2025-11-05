// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerOrdersResponseBody
	GetCode() *string
	SetData(v []*GetCustomerOrdersResponseBodyData) *GetCustomerOrdersResponseBody
	GetData() []*GetCustomerOrdersResponseBodyData
	SetMessage(v string) *GetCustomerOrdersResponseBody
	GetMessage() *string
	SetMsg(v string) *GetCustomerOrdersResponseBody
	GetMsg() *string
	SetPageNo(v int32) *GetCustomerOrdersResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetCustomerOrdersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetCustomerOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerOrdersResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetCustomerOrdersResponseBody
	GetTotal() *int32
}

type GetCustomerOrdersResponseBody struct {
	// example:
	//
	// 200
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetCustomerOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Msg     *string                              `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 23309219-4A34-589D-A3E0-9B2A3BFFD24F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCustomerOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerOrdersResponseBody) GetData() []*GetCustomerOrdersResponseBodyData {
	return s.Data
}

func (s *GetCustomerOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerOrdersResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetCustomerOrdersResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCustomerOrdersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCustomerOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerOrdersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetCustomerOrdersResponseBody) SetCode(v string) *GetCustomerOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetData(v []*GetCustomerOrdersResponseBodyData) *GetCustomerOrdersResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetMessage(v string) *GetCustomerOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetMsg(v string) *GetCustomerOrdersResponseBody {
	s.Msg = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetPageNo(v int32) *GetCustomerOrdersResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetPageSize(v int32) *GetCustomerOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetRequestId(v string) *GetCustomerOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetSuccess(v bool) *GetCustomerOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) SetTotal(v int32) *GetCustomerOrdersResponseBody {
	s.Total = &v
	return s
}

func (s *GetCustomerOrdersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomerOrdersResponseBodyData struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// myBd
	CustomerManager *string `json:"CustomerManager,omitempty" xml:"CustomerManager,omitempty"`
	// example:
	//
	// 123456
	CustomerNo *int64 `json:"CustomerNo,omitempty" xml:"CustomerNo,omitempty"`
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 0
	OrderSource *string `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 3.92
	OriginalCost *float64 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	// example:
	//
	// 3:32
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// 2024-08-13 13:02:02
	PaymentTime *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	// example:
	//
	// 3.92
	PretaxCost *float64 `json:"PretaxCost,omitempty" xml:"PretaxCost,omitempty"`
	// example:
	//
	// oss
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// example:
	//
	// snapshot
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 2024-08-13 13:02:02
	TimeToOrder *string `json:"TimeToOrder,omitempty" xml:"TimeToOrder,omitempty"`
}

func (s GetCustomerOrdersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersResponseBodyData) GetCustomerAccount() *string {
	return s.CustomerAccount
}

func (s *GetCustomerOrdersResponseBodyData) GetCustomerManager() *string {
	return s.CustomerManager
}

func (s *GetCustomerOrdersResponseBodyData) GetCustomerNo() *int64 {
	return s.CustomerNo
}

func (s *GetCustomerOrdersResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetCustomerOrdersResponseBodyData) GetOrderSource() *string {
	return s.OrderSource
}

func (s *GetCustomerOrdersResponseBodyData) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetCustomerOrdersResponseBodyData) GetOrderType() *string {
	return s.OrderType
}

func (s *GetCustomerOrdersResponseBodyData) GetOriginalCost() *float64 {
	return s.OriginalCost
}

func (s *GetCustomerOrdersResponseBodyData) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *GetCustomerOrdersResponseBodyData) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *GetCustomerOrdersResponseBodyData) GetPretaxCost() *float64 {
	return s.PretaxCost
}

func (s *GetCustomerOrdersResponseBodyData) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *GetCustomerOrdersResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetCustomerOrdersResponseBodyData) GetTimeToOrder() *string {
	return s.TimeToOrder
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerAccount(v string) *GetCustomerOrdersResponseBodyData {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerManager(v string) *GetCustomerOrdersResponseBodyData {
	s.CustomerManager = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetCustomerNo(v int64) *GetCustomerOrdersResponseBodyData {
	s.CustomerNo = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderId(v int64) *GetCustomerOrdersResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderSource(v string) *GetCustomerOrdersResponseBodyData {
	s.OrderSource = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderStatus(v int32) *GetCustomerOrdersResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOrderType(v string) *GetCustomerOrdersResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetOriginalCost(v float64) *GetCustomerOrdersResponseBodyData {
	s.OriginalCost = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPaymentMethod(v string) *GetCustomerOrdersResponseBodyData {
	s.PaymentMethod = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPaymentTime(v string) *GetCustomerOrdersResponseBodyData {
	s.PaymentTime = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetPretaxCost(v float64) *GetCustomerOrdersResponseBodyData {
	s.PretaxCost = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetProductDetail(v string) *GetCustomerOrdersResponseBodyData {
	s.ProductDetail = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetProductType(v string) *GetCustomerOrdersResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) SetTimeToOrder(v string) *GetCustomerOrdersResponseBodyData {
	s.TimeToOrder = &v
	return s
}

func (s *GetCustomerOrdersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
