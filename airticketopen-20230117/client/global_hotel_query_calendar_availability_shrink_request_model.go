// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryCalendarAvailabilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetAdultCount() *int32
	SetCheckInDateEnd(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetCheckInDateEnd() *string
	SetCheckInDateStart(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetCheckInDateStart() *string
	SetChildCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetChildCount() *int32
	SetChildrenAgesShrink(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetChildrenAgesShrink() *string
	SetRoomCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetRoomCount() *int32
	SetStandardHotelIdsShrink(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetStandardHotelIdsShrink() *string
	SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest
	GetTracerId() *string
}

type GlobalHotelQueryCalendarAvailabilityShrinkRequest struct {
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
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryCalendarAvailabilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetCheckInDateEnd() *string {
	return s.CheckInDateEnd
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetCheckInDateStart() *string {
	return s.CheckInDateStart
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetChildrenAgesShrink() *string {
	return s.ChildrenAgesShrink
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetStandardHotelIdsShrink() *string {
	return s.StandardHotelIdsShrink
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetAccountNo(v int64) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetAdultCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.AdultCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetCheckInDateEnd(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.CheckInDateEnd = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetCheckInDateStart(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.CheckInDateStart = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetChildCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.ChildCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetChildrenAgesShrink(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.ChildrenAgesShrink = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetRoomCount(v int32) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetStandardHotelIdsShrink(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.StandardHotelIdsShrink = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityShrinkRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
