// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *ListShiftSchedulesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListShiftSchedulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListShiftSchedulesRequest
	GetPageSize() *int32
	SetShiftScheduleName(v string) *ListShiftSchedulesRequest
	GetShiftScheduleName() *string
}

type ListShiftSchedulesRequest struct {
	// The Alibaba Cloud account ID. You can log on to the DataWorks console and move the pointer over the profile picture in the upper-right corner to view the ID.
	//
	// example:
	//
	// 1933790683****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Minimum value:1. Maximum value: 100. Default value: 1.
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
	// The keyword used to perform a fuzzy search on shift schedules.
	//
	// example:
	//
	// Duty table name keyword
	ShiftScheduleName *string `json:"ShiftScheduleName,omitempty" xml:"ShiftScheduleName,omitempty"`
}

func (s ListShiftSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShiftSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListShiftSchedulesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListShiftSchedulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListShiftSchedulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListShiftSchedulesRequest) GetShiftScheduleName() *string {
	return s.ShiftScheduleName
}

func (s *ListShiftSchedulesRequest) SetOwner(v string) *ListShiftSchedulesRequest {
	s.Owner = &v
	return s
}

func (s *ListShiftSchedulesRequest) SetPageNumber(v int32) *ListShiftSchedulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListShiftSchedulesRequest) SetPageSize(v int32) *ListShiftSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListShiftSchedulesRequest) SetShiftScheduleName(v string) *ListShiftSchedulesRequest {
	s.ShiftScheduleName = &v
	return s
}

func (s *ListShiftSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
