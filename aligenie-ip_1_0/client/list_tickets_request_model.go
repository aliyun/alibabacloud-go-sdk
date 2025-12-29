// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTicketsRequest
	GetEndTime() *string
	SetHotelId(v string) *ListTicketsRequest
	GetHotelId() *string
	SetIsDesc(v bool) *ListTicketsRequest
	GetIsDesc() *bool
	SetIsNeedCallback(v bool) *ListTicketsRequest
	GetIsNeedCallback() *bool
	SetIsNeedCharges(v bool) *ListTicketsRequest
	GetIsNeedCharges() *bool
	SetPage(v *ListTicketsRequestPage) *ListTicketsRequest
	GetPage() *ListTicketsRequestPage
	SetRoomNo(v string) *ListTicketsRequest
	GetRoomNo() *string
	SetSortField(v string) *ListTicketsRequest
	GetSortField() *string
	SetStartTime(v string) *ListTicketsRequest
	GetStartTime() *string
	SetStatus(v string) *ListTicketsRequest
	GetStatus() *string
	SetType(v string) *ListTicketsRequest
	GetType() *string
}

type ListTicketsRequest struct {
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
	IsNeedCharges *bool                   `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	Page          *ListTicketsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
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

func (s ListTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListTicketsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTicketsRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListTicketsRequest) GetIsDesc() *bool {
	return s.IsDesc
}

func (s *ListTicketsRequest) GetIsNeedCallback() *bool {
	return s.IsNeedCallback
}

func (s *ListTicketsRequest) GetIsNeedCharges() *bool {
	return s.IsNeedCharges
}

func (s *ListTicketsRequest) GetPage() *ListTicketsRequestPage {
	return s.Page
}

func (s *ListTicketsRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListTicketsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTicketsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTicketsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTicketsRequest) GetType() *string {
	return s.Type
}

func (s *ListTicketsRequest) SetEndTime(v string) *ListTicketsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTicketsRequest) SetHotelId(v string) *ListTicketsRequest {
	s.HotelId = &v
	return s
}

func (s *ListTicketsRequest) SetIsDesc(v bool) *ListTicketsRequest {
	s.IsDesc = &v
	return s
}

func (s *ListTicketsRequest) SetIsNeedCallback(v bool) *ListTicketsRequest {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsRequest) SetIsNeedCharges(v bool) *ListTicketsRequest {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsRequest) SetPage(v *ListTicketsRequestPage) *ListTicketsRequest {
	s.Page = v
	return s
}

func (s *ListTicketsRequest) SetRoomNo(v string) *ListTicketsRequest {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsRequest) SetSortField(v string) *ListTicketsRequest {
	s.SortField = &v
	return s
}

func (s *ListTicketsRequest) SetStartTime(v string) *ListTicketsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTicketsRequest) SetStatus(v string) *ListTicketsRequest {
	s.Status = &v
	return s
}

func (s *ListTicketsRequest) SetType(v string) *ListTicketsRequest {
	s.Type = &v
	return s
}

func (s *ListTicketsRequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketsRequestPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTicketsRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsRequestPage) GoString() string {
	return s.String()
}

func (s *ListTicketsRequestPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketsRequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsRequestPage) SetPageNumber(v int32) *ListTicketsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsRequestPage) SetPageSize(v int32) *ListTicketsRequestPage {
	s.PageSize = &v
	return s
}

func (s *ListTicketsRequestPage) Validate() error {
	return dara.Validate(s)
}
