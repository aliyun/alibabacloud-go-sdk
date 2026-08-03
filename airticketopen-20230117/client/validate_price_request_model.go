// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *ValidatePriceRequest
	GetAccountNo() *int64
	SetAdults(v int32) *ValidatePriceRequest
	GetAdults() *int32
	SetChildren(v int32) *ValidatePriceRequest
	GetChildren() *int32
	SetChildrenAges(v []*int32) *ValidatePriceRequest
	GetChildrenAges() []*int32
	SetItemOfferKey(v string) *ValidatePriceRequest
	GetItemOfferKey() *string
	SetRoomCount(v int32) *ValidatePriceRequest
	GetRoomCount() *int32
	SetTracerId(v string) *ValidatePriceRequest
	GetTracerId() *string
}

type ValidatePriceRequest struct {
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
	ChildrenAges []*int32 `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty" type:"Repeated"`
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

func (s ValidatePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceRequest) GoString() string {
	return s.String()
}

func (s *ValidatePriceRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *ValidatePriceRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *ValidatePriceRequest) GetChildren() *int32 {
	return s.Children
}

func (s *ValidatePriceRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *ValidatePriceRequest) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *ValidatePriceRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *ValidatePriceRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceRequest) SetAccountNo(v int64) *ValidatePriceRequest {
	s.AccountNo = &v
	return s
}

func (s *ValidatePriceRequest) SetAdults(v int32) *ValidatePriceRequest {
	s.Adults = &v
	return s
}

func (s *ValidatePriceRequest) SetChildren(v int32) *ValidatePriceRequest {
	s.Children = &v
	return s
}

func (s *ValidatePriceRequest) SetChildrenAges(v []*int32) *ValidatePriceRequest {
	s.ChildrenAges = v
	return s
}

func (s *ValidatePriceRequest) SetItemOfferKey(v string) *ValidatePriceRequest {
	s.ItemOfferKey = &v
	return s
}

func (s *ValidatePriceRequest) SetRoomCount(v int32) *ValidatePriceRequest {
	s.RoomCount = &v
	return s
}

func (s *ValidatePriceRequest) SetTracerId(v string) *ValidatePriceRequest {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceRequest) Validate() error {
	return dara.Validate(s)
}
