// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftPersonnelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPaging(v *ListShiftPersonnelsResponseBodyPaging) *ListShiftPersonnelsResponseBody
	GetPaging() *ListShiftPersonnelsResponseBodyPaging
	SetRequestId(v string) *ListShiftPersonnelsResponseBody
	GetRequestId() *string
}

type ListShiftPersonnelsResponseBody struct {
	// The pagination information.
	Paging *ListShiftPersonnelsResponseBodyPaging `json:"Paging,omitempty" xml:"Paging,omitempty" type:"Struct"`
	// The request ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListShiftPersonnelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListShiftPersonnelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListShiftPersonnelsResponseBody) GetPaging() *ListShiftPersonnelsResponseBodyPaging {
	return s.Paging
}

func (s *ListShiftPersonnelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListShiftPersonnelsResponseBody) SetPaging(v *ListShiftPersonnelsResponseBodyPaging) *ListShiftPersonnelsResponseBody {
	s.Paging = v
	return s
}

func (s *ListShiftPersonnelsResponseBody) SetRequestId(v string) *ListShiftPersonnelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListShiftPersonnelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListShiftPersonnelsResponseBodyPaging struct {
	// The page number. Valid values: 1 to 100. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The on-duty engineers in the shift schedule.
	ShiftPersons []*ListShiftPersonnelsResponseBodyPagingShiftPersons `json:"ShiftPersons,omitempty" xml:"ShiftPersons,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListShiftPersonnelsResponseBodyPaging) String() string {
	return dara.Prettify(s)
}

func (s ListShiftPersonnelsResponseBodyPaging) GoString() string {
	return s.String()
}

func (s *ListShiftPersonnelsResponseBodyPaging) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListShiftPersonnelsResponseBodyPaging) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListShiftPersonnelsResponseBodyPaging) GetShiftPersons() []*ListShiftPersonnelsResponseBodyPagingShiftPersons {
	return s.ShiftPersons
}

func (s *ListShiftPersonnelsResponseBodyPaging) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListShiftPersonnelsResponseBodyPaging) SetPageNumber(v int32) *ListShiftPersonnelsResponseBodyPaging {
	s.PageNumber = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPaging) SetPageSize(v int32) *ListShiftPersonnelsResponseBodyPaging {
	s.PageSize = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPaging) SetShiftPersons(v []*ListShiftPersonnelsResponseBodyPagingShiftPersons) *ListShiftPersonnelsResponseBodyPaging {
	s.ShiftPersons = v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPaging) SetTotalCount(v int32) *ListShiftPersonnelsResponseBodyPaging {
	s.TotalCount = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPaging) Validate() error {
	return dara.Validate(s)
}

type ListShiftPersonnelsResponseBodyPagingShiftPersons struct {
	// The time when the on-duty engineer starts the shift.
	//
	// example:
	//
	// 1593950832000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The time when the on-duty engineer ends the shift.
	//
	// example:
	//
	// 1593950832000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the on-duty engineer.
	//
	// example:
	//
	// Zhang San
	ShiftPersonName *string `json:"ShiftPersonName,omitempty" xml:"ShiftPersonName,omitempty"`
	// The UID of the on-duty engineer.
	//
	// example:
	//
	// 3726346****
	ShiftPersonUID *string `json:"ShiftPersonUID,omitempty" xml:"ShiftPersonUID,omitempty"`
}

func (s ListShiftPersonnelsResponseBodyPagingShiftPersons) String() string {
	return dara.Prettify(s)
}

func (s ListShiftPersonnelsResponseBodyPagingShiftPersons) GoString() string {
	return s.String()
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) GetShiftPersonName() *string {
	return s.ShiftPersonName
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) GetShiftPersonUID() *string {
	return s.ShiftPersonUID
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) SetBeginTime(v int64) *ListShiftPersonnelsResponseBodyPagingShiftPersons {
	s.BeginTime = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) SetEndTime(v int64) *ListShiftPersonnelsResponseBodyPagingShiftPersons {
	s.EndTime = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) SetShiftPersonName(v string) *ListShiftPersonnelsResponseBodyPagingShiftPersons {
	s.ShiftPersonName = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) SetShiftPersonUID(v string) *ListShiftPersonnelsResponseBodyPagingShiftPersons {
	s.ShiftPersonUID = &v
	return s
}

func (s *ListShiftPersonnelsResponseBodyPagingShiftPersons) Validate() error {
	return dara.Validate(s)
}
