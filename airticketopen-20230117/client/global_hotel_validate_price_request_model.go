// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelValidatePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelValidatePriceRequest
	GetAccountNo() *int64
	SetAdults(v int32) *GlobalHotelValidatePriceRequest
	GetAdults() *int32
	SetChildren(v int32) *GlobalHotelValidatePriceRequest
	GetChildren() *int32
	SetChildrenAges(v []*int32) *GlobalHotelValidatePriceRequest
	GetChildrenAges() []*int32
	SetItemOfferKey(v string) *GlobalHotelValidatePriceRequest
	GetItemOfferKey() *string
	SetRoomCount(v int32) *GlobalHotelValidatePriceRequest
	GetRoomCount() *int32
	SetTracerId(v string) *GlobalHotelValidatePriceRequest
	GetTracerId() *string
}

type GlobalHotelValidatePriceRequest struct {
	// The distributor account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The number of adults per room.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Adults *int32 `json:"Adults,omitempty" xml:"Adults,omitempty"`
	// The number of children per room.
	//
	// example:
	//
	// 0
	Children *int32 `json:"Children,omitempty" xml:"Children,omitempty"`
	// The list of children ages.
	//
	// example:
	//
	// [8]
	ChildrenAges []*int32 `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty" type:"Repeated"`
	// The offer key.
	//
	// This parameter is required.
	//
	// example:
	//
	// itemOfferKey_abc123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// The number of rooms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelValidatePriceRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *GlobalHotelValidatePriceRequest) GetChildren() *int32 {
	return s.Children
}

func (s *GlobalHotelValidatePriceRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *GlobalHotelValidatePriceRequest) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *GlobalHotelValidatePriceRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelValidatePriceRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceRequest) SetAccountNo(v int64) *GlobalHotelValidatePriceRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetAdults(v int32) *GlobalHotelValidatePriceRequest {
	s.Adults = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetChildren(v int32) *GlobalHotelValidatePriceRequest {
	s.Children = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetChildrenAges(v []*int32) *GlobalHotelValidatePriceRequest {
	s.ChildrenAges = v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetItemOfferKey(v string) *GlobalHotelValidatePriceRequest {
	s.ItemOfferKey = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetRoomCount(v int32) *GlobalHotelValidatePriceRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) SetTracerId(v string) *GlobalHotelValidatePriceRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceRequest) Validate() error {
	return dara.Validate(s)
}
