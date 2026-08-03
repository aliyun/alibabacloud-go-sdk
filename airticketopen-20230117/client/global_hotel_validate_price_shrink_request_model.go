// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelValidatePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelValidatePriceShrinkRequest
	GetAccountNo() *int64
	SetAdults(v int32) *GlobalHotelValidatePriceShrinkRequest
	GetAdults() *int32
	SetChildren(v int32) *GlobalHotelValidatePriceShrinkRequest
	GetChildren() *int32
	SetChildrenAgesShrink(v string) *GlobalHotelValidatePriceShrinkRequest
	GetChildrenAgesShrink() *string
	SetItemOfferKey(v string) *GlobalHotelValidatePriceShrinkRequest
	GetItemOfferKey() *string
	SetRoomCount(v int32) *GlobalHotelValidatePriceShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelValidatePriceShrinkRequest
	GetTracerId() *string
}

type GlobalHotelValidatePriceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Adults *int32 `json:"Adults,omitempty" xml:"Adults,omitempty"`
	// example:
	//
	// 0
	Children *int32 `json:"Children,omitempty" xml:"Children,omitempty"`
	// example:
	//
	// [8]
	ChildrenAgesShrink *string `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// itemOfferKey_abc123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32  `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	TracerId  *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetChildren() *int32 {
	return s.Children
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelValidatePriceShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetAccountNo(v int64) *GlobalHotelValidatePriceShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetAdults(v int32) *GlobalHotelValidatePriceShrinkRequest {
	s.Adults = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetChildren(v int32) *GlobalHotelValidatePriceShrinkRequest {
	s.Children = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetChildrenAgesShrink(v string) *GlobalHotelValidatePriceShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetItemOfferKey(v string) *GlobalHotelValidatePriceShrinkRequest {
	s.ItemOfferKey = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetRoomCount(v int32) *GlobalHotelValidatePriceShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) SetTracerId(v string) *GlobalHotelValidatePriceShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
