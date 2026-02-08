// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SubmitBatchJobsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *SubmitBatchJobsRequest
	GetJobGroupId() *string
}

type SubmitBatchJobsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b3ea2a1-32b3-4041-842b-9bde5de9dda0
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s SubmitBatchJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitBatchJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitBatchJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *SubmitBatchJobsRequest) SetInstanceId(v string) *SubmitBatchJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *SubmitBatchJobsRequest) SetJobGroupId(v string) *SubmitBatchJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *SubmitBatchJobsRequest) Validate() error {
	return dara.Validate(s)
}
