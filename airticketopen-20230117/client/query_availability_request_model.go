// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *QueryAvailabilityRequest
	GetAccountNo() *int64
	SetAdultCount(v int32) *QueryAvailabilityRequest
	GetAdultCount() *int32
	SetCheckInDate(v string) *QueryAvailabilityRequest
	GetCheckInDate() *string
	SetCheckOutDate(v string) *QueryAvailabilityRequest
	GetCheckOutDate() *string
	SetChildCount(v int32) *QueryAvailabilityRequest
	GetChildCount() *int32
	SetChildrenAges(v []*int32) *QueryAvailabilityRequest
	GetChildrenAges() []*int32
	SetRoomCount(v int32) *QueryAvailabilityRequest
	GetRoomCount() *int32
	SetStandardHotelIds(v []*string) *QueryAvailabilityRequest
	GetStandardHotelIds() []*string
	SetTracerId(v string) *QueryAvailabilityRequest
	GetTracerId() *string
}

type QueryAvailabilityRequest struct {
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
	ChildrenAges []*int32 `json:"ChildrenAges,omitempty" xml:"ChildrenAges,omitempty" type:"Repeated"`
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
	// string
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailabilityRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *QueryAvailabilityRequest) GetAdultCount() *int32 {
	return s.AdultCount
}

func (s *QueryAvailabilityRequest) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *QueryAvailabilityRequest) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *QueryAvailabilityRequest) GetChildCount() *int32 {
	return s.ChildCount
}

func (s *QueryAvailabilityRequest) GetChildrenAges() []*int32 {
	return s.ChildrenAges
}

func (s *QueryAvailabilityRequest) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *QueryAvailabilityRequest) GetStandardHotelIds() []*string {
	return s.StandardHotelIds
}

func (s *QueryAvailabilityRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryAvailabilityRequest) SetAccountNo(v int64) *QueryAvailabilityRequest {
	s.AccountNo = &v
	return s
}

func (s *QueryAvailabilityRequest) SetAdultCount(v int32) *QueryAvailabilityRequest {
	s.AdultCount = &v
	return s
}

func (s *QueryAvailabilityRequest) SetCheckInDate(v string) *QueryAvailabilityRequest {
	s.CheckInDate = &v
	return s
}

func (s *QueryAvailabilityRequest) SetCheckOutDate(v string) *QueryAvailabilityRequest {
	s.CheckOutDate = &v
	return s
}

func (s *QueryAvailabilityRequest) SetChildCount(v int32) *QueryAvailabilityRequest {
	s.ChildCount = &v
	return s
}

func (s *QueryAvailabilityRequest) SetChildrenAges(v []*int32) *QueryAvailabilityRequest {
	s.ChildrenAges = v
	return s
}

func (s *QueryAvailabilityRequest) SetRoomCount(v int32) *QueryAvailabilityRequest {
	s.RoomCount = &v
	return s
}

func (s *QueryAvailabilityRequest) SetStandardHotelIds(v []*string) *QueryAvailabilityRequest {
	s.StandardHotelIds = v
	return s
}

func (s *QueryAvailabilityRequest) SetTracerId(v string) *QueryAvailabilityRequest {
	s.TracerId = &v
	return s
}

func (s *QueryAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
