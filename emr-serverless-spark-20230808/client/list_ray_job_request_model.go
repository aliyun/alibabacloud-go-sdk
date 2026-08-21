// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListRayJobRequest
	GetName() *string
	SetPageNum(v int32) *ListRayJobRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRayJobRequest
	GetPageSize() *int32
	SetSubmissionId(v string) *ListRayJobRequest
	GetSubmissionId() *string
	SetSubmitTime(v *ListRayJobRequestSubmitTime) *ListRayJobRequest
	GetSubmitTime() *ListRayJobRequestSubmitTime
	SetTaskBizId(v string) *ListRayJobRequest
	GetTaskBizId() *string
}

type ListRayJobRequest struct {
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
	SubmitTime *ListRayJobRequestSubmitTime `json:"submitTime,omitempty" xml:"submitTime,omitempty" type:"Struct"`
	// The data development node ID.
	//
	// example:
	//
	// TSK-db8b870d901e443ba0aebba40c923e02
	TaskBizId *string `json:"taskBizId,omitempty" xml:"taskBizId,omitempty"`
}

func (s ListRayJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRayJobRequest) GoString() string {
	return s.String()
}

func (s *ListRayJobRequest) GetName() *string {
	return s.Name
}

func (s *ListRayJobRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRayJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRayJobRequest) GetSubmissionId() *string {
	return s.SubmissionId
}

func (s *ListRayJobRequest) GetSubmitTime() *ListRayJobRequestSubmitTime {
	return s.SubmitTime
}

func (s *ListRayJobRequest) GetTaskBizId() *string {
	return s.TaskBizId
}

func (s *ListRayJobRequest) SetName(v string) *ListRayJobRequest {
	s.Name = &v
	return s
}

func (s *ListRayJobRequest) SetPageNum(v int32) *ListRayJobRequest {
	s.PageNum = &v
	return s
}

func (s *ListRayJobRequest) SetPageSize(v int32) *ListRayJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListRayJobRequest) SetSubmissionId(v string) *ListRayJobRequest {
	s.SubmissionId = &v
	return s
}

func (s *ListRayJobRequest) SetSubmitTime(v *ListRayJobRequestSubmitTime) *ListRayJobRequest {
	s.SubmitTime = v
	return s
}

func (s *ListRayJobRequest) SetTaskBizId(v string) *ListRayJobRequest {
	s.TaskBizId = &v
	return s
}

func (s *ListRayJobRequest) Validate() error {
	if s.SubmitTime != nil {
		if err := s.SubmitTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRayJobRequestSubmitTime struct {
	// The end of the submit time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1780018822000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start of the submit time range. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1780017822000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListRayJobRequestSubmitTime) String() string {
	return dara.Prettify(s)
}

func (s ListRayJobRequestSubmitTime) GoString() string {
	return s.String()
}

func (s *ListRayJobRequestSubmitTime) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListRayJobRequestSubmitTime) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRayJobRequestSubmitTime) SetEndTime(v int64) *ListRayJobRequestSubmitTime {
	s.EndTime = &v
	return s
}

func (s *ListRayJobRequestSubmitTime) SetStartTime(v int64) *ListRayJobRequestSubmitTime {
	s.StartTime = &v
	return s
}

func (s *ListRayJobRequestSubmitTime) Validate() error {
	return dara.Validate(s)
}
