// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelQueryAvailabilityRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *GlobalHotelQueryAvailabilityRequest
	GetAdultCount() *int32
	SetCheckInDate(v string) *GlobalHotelQueryAvailabilityRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *GlobalHotelQueryAvailabilityRequest
	GetCheckOutDate() *string
	SetChildCount(v int32) *GlobalHotelQueryAvailabilityRequest
	GetChildCount() *int32
	SetChildrenAges(v []*int32) *GlobalHotelQueryAvailabilityRequest
	GetChildrenAges() []*int32
	SetRoomCount(v int32) *GlobalHotelQueryAvailabilityRequest
	GetRoomCount() *int32
	SetStandardHotelIds(v []*string) *GlobalHotelQueryAvailabilityRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *GlobalHotelQueryAvailabilityRequest
	GetTracerId() *string
}

type GlobalHotelQueryAvailabilityRequest struct {
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
	ChildrenAges []*int32 `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty" type:"Repeated"`
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
	StandardHotelIds []*string `json:"StandardHotelIds,omitempty" xml:"StandardHotelIds,omitempty" type:"Repeated"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryAvailabilityRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelQueryAvailabilityRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *GlobalHotelQueryAvailabilityRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *GlobalHotelQueryAvailabilityRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *GlobalHotelQueryAvailabilityRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *GlobalHotelQueryAvailabilityRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *GlobalHotelQueryAvailabilityRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelQueryAvailabilityRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *GlobalHotelQueryAvailabilityRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryAvailabilityRequest) SetAccountNo(v int64) *GlobalHotelQueryAvailabilityRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetAdultCount(v int32) *GlobalHotelQueryAvailabilityRequest {
	s.AdultCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetCheckInDate(v string) *GlobalHotelQueryAvailabilityRequest {
	s.CheckInDate = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetCheckOutDate(v string) *GlobalHotelQueryAvailabilityRequest {
	s.CheckOutDate = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetChildCount(v int32) *GlobalHotelQueryAvailabilityRequest {
	s.ChildCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetChildrenAges(v []*int32) *GlobalHotelQueryAvailabilityRequest {
	s.ChildrenAges = v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetRoomCount(v int32) *GlobalHotelQueryAvailabilityRequest {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetStandardHotelIds(v []*string) *GlobalHotelQueryAvailabilityRequest {
	s.StandardHotelIds = v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) SetTracerId(v string) *GlobalHotelQueryAvailabilityRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
