// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CreateAndPayShrinkRequest
	GetAccountNo() *int64
	SetContactShrink(v string) *CreateAndPayShrinkRequest
	GetContactShrink() *string
	SetExternalOrderNo(v string) *CreateAndPayShrinkRequest
	GetExternalOrderNo() *string
	SetGuestsShrink(v string) *CreateAndPayShrinkRequest
	GetGuestsShrink() *string
	SetItemOfferId(v string) *CreateAndPayShrinkRequest
	GetItemOfferId() *string
	SetRoomCount(v int32) *CreateAndPayShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *CreateAndPayShrinkRequest
	GetTracerId() *string
}

type CreateAndPayShrinkRequest struct {
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CreateAndPayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPayShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CreateAndPayShrinkRequest) GetContactShrink() *string {
	return s.ContactShrink
}

func (s *CreateAndPayShrinkRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *CreateAndPayShrinkRequest) GetGuestsShrink() *string {
	return s.GuestsShrink
}

func (s *CreateAndPayShrinkRequest) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *CreateAndPayShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *CreateAndPayShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CreateAndPayShrinkRequest) SetAccountNo(v int64) *CreateAndPayShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetContactShrink(v string) *CreateAndPayShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetExternalOrderNo(v string) *CreateAndPayShrinkRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetGuestsShrink(v string) *CreateAndPayShrinkRequest {
	s.GuestsShrink = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetItemOfferId(v string) *CreateAndPayShrinkRequest {
	s.ItemOfferId = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetRoomCount(v int32) *CreateAndPayShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *CreateAndPayShrinkRequest) SetTracerId(v string) *CreateAndPayShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *CreateAndPayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
