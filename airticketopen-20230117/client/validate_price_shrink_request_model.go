// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *ValidatePriceShrinkRequest
	GetAccountNo() *int64
	SetAdults(v int32) *ValidatePriceShrinkRequest
	GetAdults() *int32
	SetChildren(v int32) *ValidatePriceShrinkRequest
	GetChildren() *int32
	SetChildrenAgesShrink(v string) *ValidatePriceShrinkRequest
	GetChildrenAgesShrink() *string
	SetItemOfferKey(v string) *ValidatePriceShrinkRequest
	GetItemOfferKey() *string
	SetRoomCount(v int32) *ValidatePriceShrinkRequest
	GetRoomCount() *int32
	SetTracerId(v string) *ValidatePriceShrinkRequest
	GetTracerId() *string
}

type ValidatePriceShrinkRequest struct {
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
	// 0
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
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValidatePriceShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *ValidatePriceShrinkRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *ValidatePriceShrinkRequest) GetChildren() *int32 {
	return s.Children
}

func (s *ValidatePriceShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *ValidatePriceShrinkRequest) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *ValidatePriceShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *ValidatePriceShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceShrinkRequest) SetAccountNo(v int64) *ValidatePriceShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetAdults(v int32) *ValidatePriceShrinkRequest {
	s.Adults = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetChildren(v int32) *ValidatePriceShrinkRequest {
	s.Children = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetChildrenAgesShrink(v string) *ValidatePriceShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetItemOfferKey(v string) *ValidatePriceShrinkRequest {
	s.ItemOfferKey = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetRoomCount(v int32) *ValidatePriceShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *ValidatePriceShrinkRequest) SetTracerId(v string) *ValidatePriceShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
