// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTicketsShrinkRequest
	GetEndTime() *string
	SetHotelId(v string) *ListTicketsShrinkRequest
	GetHotelId() *string
	SetIsDesc(v bool) *ListTicketsShrinkRequest
	GetIsDesc() *bool
	SetIsNeedCallback(v bool) *ListTicketsShrinkRequest
	GetIsNeedCallback() *bool
	SetIsNeedCharges(v bool) *ListTicketsShrinkRequest
	GetIsNeedCharges() *bool
	SetPageShrink(v string) *ListTicketsShrinkRequest
	GetPageShrink() *string
	SetRoomNo(v string) *ListTicketsShrinkRequest
	GetRoomNo() *string
	SetSortField(v string) *ListTicketsShrinkRequest
	GetSortField() *string
	SetStartTime(v string) *ListTicketsShrinkRequest
	GetStartTime() *string
	SetStatus(v string) *ListTicketsShrinkRequest
	GetStatus() *string
	SetType(v string) *ListTicketsShrinkRequest
	GetType() *string
}

type ListTicketsShrinkRequest struct {
	// example:
	//
	// 2022-09-14 14:23:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// true
	IsDesc *bool `json:"IsDesc,omitempty" xml:"IsDesc,omitempty"`
	// example:
	//
	// false
	IsNeedCallback *bool `json:"IsNeedCallback,omitempty" xml:"IsNeedCallback,omitempty"`
	// example:
	//
	// false
	IsNeedCharges *bool   `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	PageShrink    *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// gmtCalled
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// example:
	//
	// 2022-04-08 09:39:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTicketsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTicketsShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListTicketsShrinkRequest) GetIsDesc() *bool {
	return s.IsDesc
}

func (s *ListTicketsShrinkRequest) GetIsNeedCallback() *bool {
	return s.IsNeedCallback
}

func (s *ListTicketsShrinkRequest) GetIsNeedCharges() *bool {
	return s.IsNeedCharges
}

func (s *ListTicketsShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListTicketsShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListTicketsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTicketsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTicketsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTicketsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListTicketsShrinkRequest) SetEndTime(v string) *ListTicketsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetHotelId(v string) *ListTicketsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsDesc(v bool) *ListTicketsShrinkRequest {
	s.IsDesc = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsNeedCallback(v bool) *ListTicketsShrinkRequest {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetIsNeedCharges(v bool) *ListTicketsShrinkRequest {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetPageShrink(v string) *ListTicketsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetRoomNo(v string) *ListTicketsShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetSortField(v string) *ListTicketsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStartTime(v string) *ListTicketsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetStatus(v string) *ListTicketsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTicketsShrinkRequest) SetType(v string) *ListTicketsShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListTicketsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
