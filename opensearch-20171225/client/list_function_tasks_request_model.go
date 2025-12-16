// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListFunctionTasksRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *ListFunctionTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFunctionTasksRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListFunctionTasksRequest
	GetStartTime() *int64
	SetStatus(v string) *ListFunctionTasksRequest
	GetStatus() *string
}

type ListFunctionTasksRequest struct {
	// The end time is less than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1582646399
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 1.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The start time is greater than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1582214400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// 	- running
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListFunctionTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionTasksRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListFunctionTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionTasksRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListFunctionTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFunctionTasksRequest) SetEndTime(v int64) *ListFunctionTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageNumber(v int32) *ListFunctionTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageSize(v int32) *ListFunctionTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStartTime(v int64) *ListFunctionTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStatus(v string) *ListFunctionTasksRequest {
	s.Status = &v
	return s
}

func (s *ListFunctionTasksRequest) Validate() error {
	return dara.Validate(s)
}
