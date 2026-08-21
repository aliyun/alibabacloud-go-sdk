// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListRayJobShrinkRequest
	GetName() *string
	SetPageNum(v int32) *ListRayJobShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRayJobShrinkRequest
	GetPageSize() *int32
	SetSubmissionId(v string) *ListRayJobShrinkRequest
	GetSubmissionId() *string
	SetSubmitTimeShrink(v string) *ListRayJobShrinkRequest
	GetSubmitTimeShrink() *string
	SetTaskBizId(v string) *ListRayJobShrinkRequest
	GetTaskBizId() *string
}

type ListRayJobShrinkRequest struct {
	// The name of the Ray Job (exact match).
	//
	// example:
	//
	// myrayjob
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The submission ID of the Ray Job.
	//
	// example:
	//
	// rj-xxxxxxxxxxx
	SubmissionId *string `json:"submissionId,omitempty" xml:"submissionId,omitempty"`
	// The submit time range.
	SubmitTimeShrink *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The data development node ID.
	//
	// example:
	//
	// TSK-db8b870d901e443ba0aebba40c923e02
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
}

func (s ListRayJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRayJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRayJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListRayJobShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRayJobShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRayJobShrinkRequest) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *ListRayJobShrinkRequest) GetSubmitTimeShrink() *string {
	return s.SubmitTimeShrink
}

func (s *ListRayJobShrinkRequest) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *ListRayJobShrinkRequest) SetName(v string) *ListRayJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListRayJobShrinkRequest) SetPageNum(v int32) *ListRayJobShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListRayJobShrinkRequest) SetPageSize(v int32) *ListRayJobShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListRayJobShrinkRequest) SetSubmissionId(v string) *ListRayJobShrinkRequest {
	s.SubmissionId = &v
	return s
}

func (s *ListRayJobShrinkRequest) SetSubmitTimeShrink(v string) *ListRayJobShrinkRequest {
	s.SubmitTimeShrink = &v
	return s
}

func (s *ListRayJobShrinkRequest) SetTaskBizId(v string) *ListRayJobShrinkRequest {
	s.TaskBizId = &v
	return s
}

func (s *ListRayJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
