// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCalendarAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *QueryCalendarAvailabilityRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *QueryCalendarAvailabilityRequest
	GetAdultCount() *int32
	SetCheckInDateEnd(v string) *QueryCalendarAvailabilityRequest
	GetCheckInDateEnd() *string
	SetCheckInDateStart(v string) *QueryCalendarAvailabilityRequest
	GetCheckInDateStart() *string
	SetChildCount(v int32) *QueryCalendarAvailabilityRequest
	GetChildCount() *int32
	SetChildrenAges(v []*int32) *QueryCalendarAvailabilityRequest
	GetChildrenAges() []*int32
	SetRoomCount(v int32) *QueryCalendarAvailabilityRequest
	GetRoomCount() *int32
	SetStandardHotelIds(v []*string) *QueryCalendarAvailabilityRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *QueryCalendarAvailabilityRequest
	GetTracerId() *string
}

type QueryCalendarAvailabilityRequest struct {
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
	ChildrenAges []*int32 `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty" type:"Repeated"`
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
	StandardHotelIds []*string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty" type:"Repeated"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryCalendarAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *QueryCalendarAvailabilityRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *QueryCalendarAvailabilityRequest) GetCheckInDateEnd() *string {
	return s.CheckInDateEnd
}

func (s *QueryCalendarAvailabilityRequest) GetCheckInDateStart() *string {
	return s.CheckInDateStart
}

func (s *QueryCalendarAvailabilityRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *QueryCalendarAvailabilityRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *QueryCalendarAvailabilityRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *QueryCalendarAvailabilityRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *QueryCalendarAvailabilityRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryCalendarAvailabilityRequest) SetAccountNo(v int64) *QueryCalendarAvailabilityRequest {
	s.AccountNo = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetAdultCount(v int32) *QueryCalendarAvailabilityRequest {
	s.AdultCount = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetCheckInDateEnd(v string) *QueryCalendarAvailabilityRequest {
	s.CheckInDateEnd = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetCheckInDateStart(v string) *QueryCalendarAvailabilityRequest {
	s.CheckInDateStart = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetChildCount(v int32) *QueryCalendarAvailabilityRequest {
	s.ChildCount = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetChildrenAges(v []*int32) *QueryCalendarAvailabilityRequest {
	s.ChildrenAges = v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetRoomCount(v int32) *QueryCalendarAvailabilityRequest {
	s.RoomCount = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetStandardHotelIds(v []*string) *QueryCalendarAvailabilityRequest {
	s.StandardHotelIds = v
	return s
}

func (s *QueryCalendarAvailabilityRequest) SetTracerId(v string) *QueryCalendarAvailabilityRequest {
	s.TracerId = &v
	return s
}

func (s *QueryCalendarAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
