// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCalendarAvailabilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *QueryCalendarAvailabilityShrinkRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *QueryCalendarAvailabilityShrinkRequest
	GetAdultCount() *int32
	SetCheckInDateEnd(v string) *QueryCalendarAvailabilityShrinkRequest
	GetCheckInDateEnd() *string
	SetCheckInDateStart(v string) *QueryCalendarAvailabilityShrinkRequest
	GetCheckInDateStart() *string
	SetChildCount(v int32) *QueryCalendarAvailabilityShrinkRequest
	GetChildCount() *int32
	SetChildrenAgesShrink(v string) *QueryCalendarAvailabilityShrinkRequest
	GetChildrenAgesShrink() *string
	SetRoomCount(v int32) *QueryCalendarAvailabilityShrinkRequest
	GetRoomCount() *int32
	SetStandardHotelIdsShrink(v string) *QueryCalendarAvailabilityShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *QueryCalendarAvailabilityShrinkRequest
	GetTracerId() *string
}

type QueryCalendarAvailabilityShrinkRequest struct {
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
	AdultCount *int32 `json:"AdultCount,omitempty" xml:"AdultCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-03
	CheckInDateEnd *string `json:"CheckInDateEnd,omitempty" xml:"CheckInDateEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-01
	CheckInDateStart *string `json:"CheckInDateStart,omitempty" xml:"CheckInDateStart,omitempty"`
	// example:
	//
	// 0
	ChildCount *int32 `json:"ChildCount,omitempty" xml:"ChildCount,omitempty"`
	// example:
	//
	// []
	ChildrenAgesShrink *string `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty"`
	// This parameter is required.
	//
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
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryCalendarAvailabilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetCheckInDateEnd() *string {
	return s.CheckInDateEnd
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetCheckInDateStart() *string {
	return s.CheckInDateStart
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *QueryCalendarAvailabilityShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetAccountNo(v int64) *QueryCalendarAvailabilityShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetAdultCount(v int32) *QueryCalendarAvailabilityShrinkRequest {
	s.AdultCount = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetCheckInDateEnd(v string) *QueryCalendarAvailabilityShrinkRequest {
	s.CheckInDateEnd = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetCheckInDateStart(v string) *QueryCalendarAvailabilityShrinkRequest {
	s.CheckInDateStart = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetChildCount(v int32) *QueryCalendarAvailabilityShrinkRequest {
	s.ChildCount = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetChildrenAgesShrink(v string) *QueryCalendarAvailabilityShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetRoomCount(v int32) *QueryCalendarAvailabilityShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetStandardHotelIdsShrink(v string) *QueryCalendarAvailabilityShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) SetTracerId(v string) *QueryCalendarAvailabilityShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *QueryCalendarAvailabilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
