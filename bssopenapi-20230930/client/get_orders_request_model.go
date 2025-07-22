// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *GetOrdersRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *GetOrdersRequest
	GetCreateTimeStart() *string
	SetMemberUid(v int64) *GetOrdersRequest
	GetMemberUid() *int64
	SetOrderType(v string) *GetOrdersRequest
	GetOrderType() *string
	SetOwnerId(v int64) *GetOrdersRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *GetOrdersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetOrdersRequest
	GetPageSize() *int32
	SetPaymentStatus(v string) *GetOrdersRequest
	GetPaymentStatus() *string
	SetProductCode(v string) *GetOrdersRequest
	GetProductCode() *string
	SetProductType(v string) *GetOrdersRequest
	GetProductType() *string
	SetSubscriptionType(v string) *GetOrdersRequest
	GetSubscriptionType() *string
}

type GetOrdersRequest struct {
	// example:
	//
	// 2016-05-23T12:00:00Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2016-05-23T13:00:00Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1779628988149763
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// New
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// Paid
	PaymentStatus *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
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
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s GetOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrdersRequest) GoString() string {
	return s.String()
}

func (s *GetOrdersRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *GetOrdersRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *GetOrdersRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *GetOrdersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetOrdersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetOrdersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOrdersRequest) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetOrdersRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetOrdersRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetOrdersRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *GetOrdersRequest) SetCreateTimeEnd(v string) *GetOrdersRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *GetOrdersRequest) SetCreateTimeStart(v string) *GetOrdersRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *GetOrdersRequest) SetMemberUid(v int64) *GetOrdersRequest {
	s.MemberUid = &v
	return s
}

func (s *GetOrdersRequest) SetOrderType(v string) *GetOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *GetOrdersRequest) SetOwnerId(v int64) *GetOrdersRequest {
	s.OwnerId = &v
	return s
}

func (s *GetOrdersRequest) SetPageNum(v int32) *GetOrdersRequest {
	s.PageNum = &v
	return s
}

func (s *GetOrdersRequest) SetPageSize(v int32) *GetOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *GetOrdersRequest) SetPaymentStatus(v string) *GetOrdersRequest {
	s.PaymentStatus = &v
	return s
}

func (s *GetOrdersRequest) SetProductCode(v string) *GetOrdersRequest {
	s.ProductCode = &v
	return s
}

func (s *GetOrdersRequest) SetProductType(v string) *GetOrdersRequest {
	s.ProductType = &v
	return s
}

func (s *GetOrdersRequest) SetSubscriptionType(v string) *GetOrdersRequest {
	s.SubscriptionType = &v
	return s
}

func (s *GetOrdersRequest) Validate() error {
	return dara.Validate(s)
}
