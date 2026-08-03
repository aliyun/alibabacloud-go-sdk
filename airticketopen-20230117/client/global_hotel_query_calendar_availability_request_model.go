// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryCalendarAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelQueryCalendarAvailabilityRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest
	GetAdultCount() *int32
	SetCheckInDateEnd(v string) *GlobalHotelQueryCalendarAvailabilityRequest
	GetCheckInDateEnd() *string
	SetCheckInDateStart(v string) *GlobalHotelQueryCalendarAvailabilityRequest
	GetCheckInDateStart() *string
	SetChildCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest
	GetChildCount() *int32
	SetChildrenAges(v []*int32) *GlobalHotelQueryCalendarAvailabilityRequest
	GetChildrenAges() []*int32
	SetRoomCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest
	GetRoomCount() *int32
	SetStandardHotelIds(v []*string) *GlobalHotelQueryCalendarAvailabilityRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityRequest
	GetTracerId() *string
}

type GlobalHotelQueryCalendarAvailabilityRequest struct {
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
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryCalendarAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetCheckInDateEnd() *string {
	return s.CheckInDateEnd
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetCheckInDateStart() *string {
	return s.CheckInDateStart
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetAccountNo(v int64) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetAdultCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.AdultCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetCheckInDateEnd(v string) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.CheckInDateEnd = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetCheckInDateStart(v string) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.CheckInDateStart = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetChildCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.ChildCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetChildrenAges(v []*int32) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.ChildrenAges = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetRoomCount(v int32) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetStandardHotelIds(v []*string) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.StandardHotelIds = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) SetTracerId(v string) *GlobalHotelQueryCalendarAvailabilityRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
