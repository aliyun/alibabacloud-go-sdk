// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateAndPayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelCreateAndPayShrinkRequest
	GetAccountNo() *int64
	SetContactShrink(v string) *GlobalHotelCreateAndPayShrinkRequest
	GetContactShrink() *string
	SetExternalOrderNo(v string) *GlobalHotelCreateAndPayShrinkRequest
	GetExternalOrderNo() *string
	SetGuestsShrink(v string) *GlobalHotelCreateAndPayShrinkRequest
	GetGuestsShrink() *string
	SetItemOfferId(v string) *GlobalHotelCreateAndPayShrinkRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *GlobalHotelCreateAndPayShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelCreateAndPayShrinkRequest
	GetTracerId() *string
}

type GlobalHotelCreateAndPayShrinkRequest struct {
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

func (s GlobalHotelCreateAndPayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetGuestsShrink() *string {
	return s.GuestsShrink
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelCreateAndPayShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetAccountNo(v int64) *GlobalHotelCreateAndPayShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetContactShrink(v string) *GlobalHotelCreateAndPayShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetExternalOrderNo(v string) *GlobalHotelCreateAndPayShrinkRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetGuestsShrink(v string) *GlobalHotelCreateAndPayShrinkRequest {
	s.GuestsShrink = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetItemOfferId(v string) *GlobalHotelCreateAndPayShrinkRequest {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetRoomCount(v int32) *GlobalHotelCreateAndPayShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) SetTracerId(v string) *GlobalHotelCreateAndPayShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelCreateAndPayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
