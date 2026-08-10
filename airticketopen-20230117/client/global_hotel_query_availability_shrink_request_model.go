// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryAvailabilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelQueryAvailabilityShrinkRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest
	GetAdultCount() *int32
	SetCheckInDate(v string) *GlobalHotelQueryAvailabilityShrinkRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *GlobalHotelQueryAvailabilityShrinkRequest
	GetCheckOutDate() *string
	SetChildCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest
	GetChildCount() *int32
	SetChildrenAgesShrink(v string) *GlobalHotelQueryAvailabilityShrinkRequest
	GetChildrenAgesShrink() *string
	SetRoomCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest
	GetRoomCount() *int32
	SetStandardHotelIdsShrink(v string) *GlobalHotelQueryAvailabilityShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *GlobalHotelQueryAvailabilityShrinkRequest
	GetTracerId() *string
}

type GlobalHotelQueryAvailabilityShrinkRequest struct {
	// The account ID of the distributor.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The number of adults.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AdultCount *int32 `json:"AdultCount,omitempty" xml:"AdultCount,omitempty"`
	// The check-in date in the format of yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-01
	CheckInDate *string `json:"CheckInDate,omitempty" xml:"CheckInDate,omitempty"`
	// The check-out date in the format of yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-07-03
	CheckOutDate *string `json:"CheckOutDate,omitempty" xml:"CheckOutDate,omitempty"`
	// The number of children.
	//
	// example:
	//
	// 1
	ChildCount *int32 `json:"ChildCount,omitempty" xml:"ChildCount,omitempty"`
	// The list of children\\"s ages.
	//
	// example:
	//
	// [8]
	ChildrenAgesShrink *string `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty"`
	// The number of rooms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// The list of standard hotel IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["H001"]
	StandardHotelIdsShrink *string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryAvailabilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryAvailabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetAccountNo(v int64) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetAdultCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.AdultCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetCheckInDate(v string) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.CheckInDate = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetCheckOutDate(v string) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.CheckOutDate = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetChildCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.ChildCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetChildrenAgesShrink(v string) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetRoomCount(v int32) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetStandardHotelIdsShrink(v string) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) SetTracerId(v string) *GlobalHotelQueryAvailabilityShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
