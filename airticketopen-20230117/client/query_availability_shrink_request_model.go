// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailabilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *QueryAvailabilityShrinkRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *QueryAvailabilityShrinkRequest
	GetAdultCount() *int32
	SetCheckInDate(v string) *QueryAvailabilityShrinkRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *QueryAvailabilityShrinkRequest
	GetCheckOutDate() *string
	SetChildCount(v int32) *QueryAvailabilityShrinkRequest
	GetChildCount() *int32
	SetChildrenAgesShrink(v string) *QueryAvailabilityShrinkRequest
	GetChildrenAgesShrink() *string
	SetRoomCount(v int32) *QueryAvailabilityShrinkRequest
	GetRoomCount() *int32
	SetStandardHotelIdsShrink(v string) *QueryAvailabilityShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *QueryAvailabilityShrinkRequest
	GetTracerId() *string
}

type QueryAvailabilityShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// 2
	AdultCount *int32 `json:"AdultCount,omitempty" xml:"AdultCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-01
	CheckInDate *string `json:"CheckInDate,omitempty" xml:"CheckInDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-03
	CheckOutDate *string `json:"CheckOutDate,omitempty" xml:"CheckOutDate,omitempty"`
	// example:
	//
	// 1
	ChildCount *int32 `json:"ChildCount,omitempty" xml:"ChildCount,omitempty"`
	// example:
	//
	// [8]
	ChildrenAgesShrink *string `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty"`
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["H001"]
	StandardHotelIdsShrink *string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty"`
	// example:
	//
	// string
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryAvailabilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailabilityShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *QueryAvailabilityShrinkRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *QueryAvailabilityShrinkRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *QueryAvailabilityShrinkRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *QueryAvailabilityShrinkRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *QueryAvailabilityShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *QueryAvailabilityShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *QueryAvailabilityShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *QueryAvailabilityShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryAvailabilityShrinkRequest) SetAccountNo(v int64) *QueryAvailabilityShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetAdultCount(v int32) *QueryAvailabilityShrinkRequest {
	s.AdultCount = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetCheckInDate(v string) *QueryAvailabilityShrinkRequest {
	s.CheckInDate = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetCheckOutDate(v string) *QueryAvailabilityShrinkRequest {
	s.CheckOutDate = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetChildCount(v int32) *QueryAvailabilityShrinkRequest {
	s.ChildCount = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetChildrenAgesShrink(v string) *QueryAvailabilityShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetRoomCount(v int32) *QueryAvailabilityShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetStandardHotelIdsShrink(v string) *QueryAvailabilityShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) SetTracerId(v string) *QueryAvailabilityShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *QueryAvailabilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
