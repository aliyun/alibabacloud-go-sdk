// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPaging(v *ListShiftSchedulesResponseBodyPaging) *ListShiftSchedulesResponseBody
	GetPaging() *ListShiftSchedulesResponseBodyPaging
	SetRequestId(v string) *ListShiftSchedulesResponseBody
	GetRequestId() *string
}

type ListShiftSchedulesResponseBody struct {
	// The pagination information.
	Paging *ListShiftSchedulesResponseBodyPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListShiftSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShiftSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListShiftSchedulesResponseBody) GetPaging() *ListShiftSchedulesResponseBodyPaging {
	return s.Paging
}

func (s *ListShiftSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListShiftSchedulesResponseBody) SetPaging(v *ListShiftSchedulesResponseBodyPaging) *ListShiftSchedulesResponseBody {
	s.Paging = v
	return s
}

func (s *ListShiftSchedulesResponseBody) SetRequestId(v string) *ListShiftSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListShiftSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListShiftSchedulesResponseBodyPaging struct {
	// The page number. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The shift schedules.
	ShiftSchedules []*ListShiftSchedulesResponseBodyPagingShiftSchedules `json:"ShiftSchedules,omitempty" xml:"ShiftSchedules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListShiftSchedulesResponseBodyPaging) String() string {
	return dara.Prettify(s)
}

func (s ListShiftSchedulesResponseBodyPaging) GoString() string {
	return s.String()
}

func (s *ListShiftSchedulesResponseBodyPaging) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListShiftSchedulesResponseBodyPaging) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListShiftSchedulesResponseBodyPaging) GetShiftSchedules() []*ListShiftSchedulesResponseBodyPagingShiftSchedules {
	return s.ShiftSchedules
}

func (s *ListShiftSchedulesResponseBodyPaging) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListShiftSchedulesResponseBodyPaging) SetPageNumber(v int32) *ListShiftSchedulesResponseBodyPaging {
	s.PageNumber = &v
	return s
}

func (s *ListShiftSchedulesResponseBodyPaging) SetPageSize(v int32) *ListShiftSchedulesResponseBodyPaging {
	s.PageSize = &v
	return s
}

func (s *ListShiftSchedulesResponseBodyPaging) SetShiftSchedules(v []*ListShiftSchedulesResponseBodyPagingShiftSchedules) *ListShiftSchedulesResponseBodyPaging {
	s.ShiftSchedules = v
	return s
}

func (s *ListShiftSchedulesResponseBodyPaging) SetTotalCount(v int32) *ListShiftSchedulesResponseBodyPaging {
	s.TotalCount = &v
	return s
}

func (s *ListShiftSchedulesResponseBodyPaging) Validate() error {
	return dara.Validate(s)
}

type ListShiftSchedulesResponseBodyPagingShiftSchedules struct {
	// The unique identifier of the shift schedule. You can use the identifier to query the on-duty engineers in the shift schedule.
	//
	// example:
	//
	// 2ab6456ada634b2f938ee******9b45b
	ShiftScheduleIdentifier *string `json:"ShiftScheduleIdentifier,omitempty" xml:"ShiftScheduleIdentifier,omitempty"`
	// The name of the shift schedule.
	//
	// example:
	//
	// Duty table name
	ShiftScheduleName *string `json:"ShiftScheduleName,omitempty" xml:"ShiftScheduleName,omitempty"`
}

func (s ListShiftSchedulesResponseBodyPagingShiftSchedules) String() string {
	return dara.Prettify(s)
}

func (s ListShiftSchedulesResponseBodyPagingShiftSchedules) GoString() string {
	return s.String()
}

func (s *ListShiftSchedulesResponseBodyPagingShiftSchedules) GetShiftScheduleIdentifier() *string {
	return s.ShiftScheduleIdentifier
}

func (s *ListShiftSchedulesResponseBodyPagingShiftSchedules) GetShiftScheduleName() *string {
	return s.ShiftScheduleName
}

func (s *ListShiftSchedulesResponseBodyPagingShiftSchedules) SetShiftScheduleIdentifier(v string) *ListShiftSchedulesResponseBodyPagingShiftSchedules {
	s.ShiftScheduleIdentifier = &v
	return s
}

func (s *ListShiftSchedulesResponseBodyPagingShiftSchedules) SetShiftScheduleName(v string) *ListShiftSchedulesResponseBodyPagingShiftSchedules {
	s.ShiftScheduleName = &v
	return s
}

func (s *ListShiftSchedulesResponseBodyPagingShiftSchedules) Validate() error {
	return dara.Validate(s)
}
