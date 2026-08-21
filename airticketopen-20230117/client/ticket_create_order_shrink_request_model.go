// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *TicketCreateOrderShrinkRequest
	GetAccountNo() *int64
	SetContactShrink(v string) *TicketCreateOrderShrinkRequest
	GetContactShrink() *string
	SetDistributorOrderId(v string) *TicketCreateOrderShrinkRequest
	GetDistributorOrderId() *string
	SetOrderProductShrink(v string) *TicketCreateOrderShrinkRequest
	GetOrderProductShrink() *string
	SetQuantity(v int32) *TicketCreateOrderShrinkRequest
	GetQuantity() *int32
	SetTotalDistributionPriceShrink(v string) *TicketCreateOrderShrinkRequest
	GetTotalDistributionPriceShrink() *string
	SetTravelersShrink(v string) *TicketCreateOrderShrinkRequest
	GetTravelersShrink() *string
}

type TicketCreateOrderShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	ContactShrink *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	DistributorOrderId *string `json:"DistributorOrderId,omitempty" xml:"DistributorOrderId,omitempty"`
	// This parameter is required.
	OrderProductShrink *string `json:"OrderProduct,omitempty" xml:"OrderProduct,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// This parameter is required.
	TotalDistributionPriceShrink *string `json:"TotalDistributionPrice,omitempty" xml:"TotalDistributionPrice,omitempty"`
	TravelersShrink              *string `json:"Travelers,omitempty" xml:"Travelers,omitempty"`
}

func (s TicketCreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketCreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *TicketCreateOrderShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *TicketCreateOrderShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *TicketCreateOrderShrinkRequest) GetDistributorOrderId() *string {
	return s.DistributorOrderId
}

func (s *TicketCreateOrderShrinkRequest) GetOrderProductShrink() *string {
	return s.OrderProductShrink
}

func (s *TicketCreateOrderShrinkRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *TicketCreateOrderShrinkRequest) GetTotalDistributionPriceShrink() *string {
	return s.TotalDistributionPriceShrink
}

func (s *TicketCreateOrderShrinkRequest) GetTravelersShrink() *string {
	return s.TravelersShrink
}

func (s *TicketCreateOrderShrinkRequest) SetAccountNo(v int64) *TicketCreateOrderShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetContactShrink(v string) *TicketCreateOrderShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetDistributorOrderId(v string) *TicketCreateOrderShrinkRequest {
	s.DistributorOrderId = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetOrderProductShrink(v string) *TicketCreateOrderShrinkRequest {
	s.OrderProductShrink = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetQuantity(v int32) *TicketCreateOrderShrinkRequest {
	s.Quantity = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetTotalDistributionPriceShrink(v string) *TicketCreateOrderShrinkRequest {
	s.TotalDistributionPriceShrink = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) SetTravelersShrink(v string) *TicketCreateOrderShrinkRequest {
	s.TravelersShrink = &v
	return s
}

func (s *TicketCreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
