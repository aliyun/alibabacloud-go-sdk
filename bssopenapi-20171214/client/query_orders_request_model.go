// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *QueryOrdersRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *QueryOrdersRequest
	GetCreateTimeStart() *string
	SetOrderType(v string) *QueryOrdersRequest
	GetOrderType() *string
	SetOwnerId(v int64) *QueryOrdersRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryOrdersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryOrdersRequest
	GetPageSize() *int32
	SetPaymentStatus(v string) *QueryOrdersRequest
	GetPaymentStatus() *string
	SetProductCode(v string) *QueryOrdersRequest
	GetProductCode() *string
	SetProductType(v string) *QueryOrdersRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QueryOrdersRequest
	GetSubscriptionType() *string
}

type QueryOrdersRequest struct {
	// The end time of the period during which the orders were created. By default, orders within the last hour are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-05-23T12:00:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start time of the period during which the orders were created. By default, orders within the last hour are queried. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2016-05-23T13:00:00Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The type of the order. Valid values:
	//
	// 	- New: purchases an instance.
	//
	// 	- Renew: renews an instance.
	//
	// 	- Upgrade: upgrades the configurations of an instance.
	//
	// 	- Refund: applies for a refund.
	//
	// example:
	//
	// New
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of payment. Valid values for a non-refund order:
	//
	// 	- Unpaid: The order is not paid.
	//
	// 	- Paid: The order is paid.
	//
	// 	- Cancelled: The order is canceled.
	//
	// > : You can set this parameter to NULL for a refund order.
	//
	// example:
	//
	// Paid
	PaymentStatus *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrdersRequest) GoString() string {
	return s.String()
}

func (s *QueryOrdersRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *QueryOrdersRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *QueryOrdersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryOrdersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryOrdersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOrdersRequest) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *QueryOrdersRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryOrdersRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryOrdersRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryOrdersRequest) SetCreateTimeEnd(v string) *QueryOrdersRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryOrdersRequest) SetCreateTimeStart(v string) *QueryOrdersRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryOrdersRequest) SetOrderType(v string) *QueryOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *QueryOrdersRequest) SetOwnerId(v int64) *QueryOrdersRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryOrdersRequest) SetPageNum(v int32) *QueryOrdersRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOrdersRequest) SetPageSize(v int32) *QueryOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrdersRequest) SetPaymentStatus(v string) *QueryOrdersRequest {
	s.PaymentStatus = &v
	return s
}

func (s *QueryOrdersRequest) SetProductCode(v string) *QueryOrdersRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryOrdersRequest) SetProductType(v string) *QueryOrdersRequest {
	s.ProductType = &v
	return s
}

func (s *QueryOrdersRequest) SetSubscriptionType(v string) *QueryOrdersRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryOrdersRequest) Validate() error {
	return dara.Validate(s)
}
