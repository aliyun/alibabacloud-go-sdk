// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeBatchTasksShrinkRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeBatchTasksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBatchTasksShrinkRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeBatchTasksShrinkRequest
	GetStartTime() *string
	SetStatusShrink(v string) *DescribeBatchTasksShrinkRequest
	GetStatusShrink() *string
	SetTaskType(v string) *DescribeBatchTasksShrinkRequest
	GetTaskType() *string
}

type DescribeBatchTasksShrinkRequest struct {
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
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusShrink *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// polarclaw_install_skills
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeBatchTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchTasksShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBatchTasksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBatchTasksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBatchTasksShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBatchTasksShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *DescribeBatchTasksShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeBatchTasksShrinkRequest) SetEndTime(v string) *DescribeBatchTasksShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) SetPageNumber(v int32) *DescribeBatchTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) SetPageSize(v int32) *DescribeBatchTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) SetStartTime(v string) *DescribeBatchTasksShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) SetStatusShrink(v string) *DescribeBatchTasksShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) SetTaskType(v string) *DescribeBatchTasksShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeBatchTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
