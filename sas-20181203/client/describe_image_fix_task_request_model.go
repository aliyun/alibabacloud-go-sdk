// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFixTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageFixTaskRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeImageFixTaskRequest
	GetEndTime() *int64
	SetPageSize(v int32) *DescribeImageFixTaskRequest
	GetPageSize() *int32
	SetStartTime(v int64) *DescribeImageFixTaskRequest
	GetStartTime() *int64
	SetStatus(v string) *DescribeImageFixTaskRequest
	GetStatus() *string
}

type DescribeImageFixTaskRequest struct {
	// The number of the page to return. Default value: **1**
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The timestamp when the task ends. Unit: milliseconds.
	//
	// example:
	//
	// 1635575219000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Default value: **20**
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The timestamp when the task starts. Unit: milliseconds.
	//
	// example:
	//
	// 1634725571000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **1**: The task is running.
	//
	// 	- **2**: The task is successful.
	//
	// 	- **3**: The task failed.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageFixTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageFixTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageFixTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeImageFixTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageFixTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeImageFixTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageFixTaskRequest) SetCurrentPage(v int32) *DescribeImageFixTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageFixTaskRequest) SetEndTime(v int64) *DescribeImageFixTaskRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeImageFixTaskRequest) SetPageSize(v int32) *DescribeImageFixTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageFixTaskRequest) SetStartTime(v int64) *DescribeImageFixTaskRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeImageFixTaskRequest) SetStatus(v string) *DescribeImageFixTaskRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageFixTaskRequest) Validate() error {
	return dara.Validate(s)
}
