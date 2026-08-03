// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CreateOrderShrinkRequest
	GetAccountNo() *int64
	SetContactShrink(v string) *CreateOrderShrinkRequest
	GetContactShrink() *string
	SetExternalOrderNo(v string) *CreateOrderShrinkRequest
	GetExternalOrderNo() *string
	SetGuestsShrink(v string) *CreateOrderShrinkRequest
	GetGuestsShrink() *string
	SetItemOfferId(v string) *CreateOrderShrinkRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *CreateOrderShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *CreateOrderShrinkRequest
	GetTracerId() *string
}

type CreateOrderShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	ContactShrink *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// This parameter is required.
	GuestsShrink *string `json:"Guests,omitempty" xml:"Guests,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// itemOffer_123
	ItemOfferId *string `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CreateOrderShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *CreateOrderShrinkRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *CreateOrderShrinkRequest) GetGuestsShrink() *string {
	return s.GuestsShrink
}

func (s *CreateOrderShrinkRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *CreateOrderShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *CreateOrderShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateOrderShrinkRequest) SetAccountNo(v int64) *CreateOrderShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetContactShrink(v string) *CreateOrderShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetExternalOrderNo(v string) *CreateOrderShrinkRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetGuestsShrink(v string) *CreateOrderShrinkRequest {
	s.GuestsShrink = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetItemOfferId(v string) *CreateOrderShrinkRequest {
	s.ItemOfferId = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetRoomCount(v int32) *CreateOrderShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *CreateOrderShrinkRequest) SetTracerId(v string) *CreateOrderShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *CreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
