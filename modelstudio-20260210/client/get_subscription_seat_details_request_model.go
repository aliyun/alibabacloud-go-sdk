// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionSeatDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *GetSubscriptionSeatDetailsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetSubscriptionSeatDetailsRequest
	GetPageSize() *int32
	SetQueryAssigned(v bool) *GetSubscriptionSeatDetailsRequest
	GetQueryAssigned() *bool
	SetSeatId(v string) *GetSubscriptionSeatDetailsRequest
	GetSeatId() *string
	SetSeatType(v string) *GetSubscriptionSeatDetailsRequest
	GetSeatType() *string
	SetStatusList(v []*string) *GetSubscriptionSeatDetailsRequest
	GetStatusList() []*string
}

type GetSubscriptionSeatDetailsRequest struct {
	// The page number. Default value: 1. Valid values: positive integers.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The seat assignment status filter. Set to true for assigned seats, false for unassigned seats, or null for all seats.
	//
	// example:
	//
	// true
	QueryAssigned *bool `json:"QueryAssigned,omitempty" xml:"QueryAssigned,omitempty"`
	// The seat ID.
	//
	// example:
	//
	// seat-123456
	SeatId *string `json:"SeatId,omitempty" xml:"SeatId,omitempty"`
	// The seat type (specType). Valid values:
	//
	// - standard: standard seat.
	//
	// - pro: pro seat.
	//
	// - max: premium seat.
	//
	// example:
	//
	// standard
	SeatType *string `json:"SeatType,omitempty" xml:"SeatType,omitempty"`
	// The list of seat statuses used for filtering.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s GetSubscriptionSeatDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubscriptionSeatDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubscriptionSeatDetailsRequest) GetQueryAssigned() *bool {
	return s.QueryAssigned
}

func (s *GetSubscriptionSeatDetailsRequest) GetSeatId() *string {
	return s.SeatId
}

func (s *GetSubscriptionSeatDetailsRequest) GetSeatType() *string {
	return s.SeatType
}

func (s *GetSubscriptionSeatDetailsRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *GetSubscriptionSeatDetailsRequest) SetPageNo(v int32) *GetSubscriptionSeatDetailsRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) SetPageSize(v int32) *GetSubscriptionSeatDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) SetQueryAssigned(v bool) *GetSubscriptionSeatDetailsRequest {
	s.QueryAssigned = &v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) SetSeatId(v string) *GetSubscriptionSeatDetailsRequest {
	s.SeatId = &v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) SetSeatType(v string) *GetSubscriptionSeatDetailsRequest {
	s.SeatType = &v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) SetStatusList(v []*string) *GetSubscriptionSeatDetailsRequest {
	s.StatusList = v
	return s
}

func (s *GetSubscriptionSeatDetailsRequest) Validate() error {
	return dara.Validate(s)
}
