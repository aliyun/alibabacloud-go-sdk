// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsByGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListJobsByGroupRequest
	GetInstanceId() *string
	SetJobFailureReason(v string) *ListJobsByGroupRequest
	GetJobFailureReason() *string
	SetJobGroupId(v string) *ListJobsByGroupRequest
	GetJobGroupId() *string
	SetJobStatus(v string) *ListJobsByGroupRequest
	GetJobStatus() *string
	SetPageNumber(v int32) *ListJobsByGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsByGroupRequest
	GetPageSize() *int32
}

type ListJobsByGroupRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job failure reason
	//
	// example:
	//
	// NoAnswer
	JobFailureReason *string `json:"JobFailureReason,omitempty" xml:"JobFailureReason,omitempty"`
	// Job group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// de48407d-309e-451a-81ec-6fb11f8fdbf3
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job status
	//
	// example:
	//
	// Succeeded
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsByGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsByGroupRequest) GoString() string {
	return s.String()
}

func (s *ListJobsByGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListJobsByGroupRequest) GetJobFailureReason() *string {
	return s.JobFailureReason
}

func (s *ListJobsByGroupRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ListJobsByGroupRequest) GetJobStatus() *string {
	return s.JobStatus
}

func (s *ListJobsByGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsByGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsByGroupRequest) SetInstanceId(v string) *ListJobsByGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListJobsByGroupRequest) SetJobFailureReason(v string) *ListJobsByGroupRequest {
	s.JobFailureReason = &v
	return s
}

func (s *ListJobsByGroupRequest) SetJobGroupId(v string) *ListJobsByGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *ListJobsByGroupRequest) SetJobStatus(v string) *ListJobsByGroupRequest {
	s.JobStatus = &v
	return s
}

func (s *ListJobsByGroupRequest) SetPageNumber(v int32) *ListJobsByGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsByGroupRequest) SetPageSize(v int32) *ListJobsByGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsByGroupRequest) Validate() error {
	return dara.Validate(s)
}
