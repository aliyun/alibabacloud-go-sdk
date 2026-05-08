// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListAICoachTaskPageRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListAICoachTaskPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAICoachTaskPageRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListAICoachTaskPageRequest
	GetStartTime() *string
	SetStatus(v string) *ListAICoachTaskPageRequest
	GetStatus() *string
	SetStudentId(v string) *ListAICoachTaskPageRequest
	GetStudentId() *string
	SetTaskId(v string) *ListAICoachTaskPageRequest
	GetTaskId() *string
}

type ListAICoachTaskPageRequest struct {
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 111
	StudentId *string `json:"studentId,omitempty" xml:"studentId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAICoachTaskPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskPageRequest) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskPageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAICoachTaskPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAICoachTaskPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAICoachTaskPageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAICoachTaskPageRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAICoachTaskPageRequest) GetStudentId() *string {
	return s.StudentId
}

func (s *ListAICoachTaskPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAICoachTaskPageRequest) SetEndTime(v string) *ListAICoachTaskPageRequest {
	s.EndTime = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetPageNumber(v int32) *ListAICoachTaskPageRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetPageSize(v int32) *ListAICoachTaskPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStartTime(v string) *ListAICoachTaskPageRequest {
	s.StartTime = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStatus(v string) *ListAICoachTaskPageRequest {
	s.Status = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetStudentId(v string) *ListAICoachTaskPageRequest {
	s.StudentId = &v
	return s
}

func (s *ListAICoachTaskPageRequest) SetTaskId(v string) *ListAICoachTaskPageRequest {
	s.TaskId = &v
	return s
}

func (s *ListAICoachTaskPageRequest) Validate() error {
	return dara.Validate(s)
}
