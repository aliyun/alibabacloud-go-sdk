// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCreateOrderShrinkRequest
	GetAccountNo() *int64
	SetContactShrink(v string) *GlobalHotelCreateOrderShrinkRequest
	GetContactShrink() *string
	SetExternalOrderNo(v string) *GlobalHotelCreateOrderShrinkRequest
	GetExternalOrderNo() *string
	SetGuestsShrink(v string) *GlobalHotelCreateOrderShrinkRequest
	GetGuestsShrink() *string
	SetItemOfferId(v string) *GlobalHotelCreateOrderShrinkRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *GlobalHotelCreateOrderShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelCreateOrderShrinkRequest
	GetTracerId() *string
}

type GlobalHotelCreateOrderShrinkRequest struct {
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

func (s GlobalHotelCreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetGuestsShrink() *string {
	return s.GuestsShrink
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelCreateOrderShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetAccountNo(v int64) *GlobalHotelCreateOrderShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetContactShrink(v string) *GlobalHotelCreateOrderShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetExternalOrderNo(v string) *GlobalHotelCreateOrderShrinkRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetGuestsShrink(v string) *GlobalHotelCreateOrderShrinkRequest {
	s.GuestsShrink = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetItemOfferId(v string) *GlobalHotelCreateOrderShrinkRequest {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetRoomCount(v int32) *GlobalHotelCreateOrderShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) SetTracerId(v string) *GlobalHotelCreateOrderShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
