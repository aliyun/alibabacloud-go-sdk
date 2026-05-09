// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeBatchTasksRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeBatchTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBatchTasksRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeBatchTasksRequest
	GetStartTime() *string
	SetStatus(v []*string) *DescribeBatchTasksRequest
	GetStatus() []*string
	SetTaskType(v string) *DescribeBatchTasksRequest
	GetTaskType() *string
}

type DescribeBatchTasksRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-04-06T20:52:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2026-04-06T20:51:40Z
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// polarclaw_install_skills
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeBatchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBatchTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBatchTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBatchTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBatchTasksRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeBatchTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeBatchTasksRequest) SetEndTime(v string) *DescribeBatchTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBatchTasksRequest) SetPageNumber(v int32) *DescribeBatchTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchTasksRequest) SetPageSize(v int32) *DescribeBatchTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchTasksRequest) SetStartTime(v string) *DescribeBatchTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBatchTasksRequest) SetStatus(v []*string) *DescribeBatchTasksRequest {
	s.Status = v
	return s
}

func (s *DescribeBatchTasksRequest) SetTaskType(v string) *DescribeBatchTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeBatchTasksRequest) Validate() error {
	return dara.Validate(s)
}
