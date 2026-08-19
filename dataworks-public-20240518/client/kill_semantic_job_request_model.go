// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSemanticJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorJobId(v string) *KillSemanticJobRequest
	GetExecutorJobId() *string
	SetProjectId(v int64) *KillSemanticJobRequest
	GetProjectId() *int64
	SetRetryTimes(v int32) *KillSemanticJobRequest
	GetRetryTimes() *int32
}

type KillSemanticJobRequest struct {
	// The executor job ID of the run to stop. Use the Data.ExecutorJobId value from the RunSemanticJob response or the ExecutorJobId value from a ListSemanticJobRuns record.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The ID of the DataWorks workspace to which the job belongs. Use the ProjectId value from the CreateSemanticJob response or a ListSemanticJobs list item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The number of retries when sending the stop request to the executor. This parameter is typically optional. If specified, use a non-negative integer. After the call, confirm the final status by calling GetSemanticJobDetail.
	//
	// example:
	//
	// 1
	RetryTimes *int32 `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
}

func (s KillSemanticJobRequest) String() string {
	return dara.Prettify(s)
}

func (s KillSemanticJobRequest) GoString() string {
	return s.String()
}

func (s *KillSemanticJobRequest) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *KillSemanticJobRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *KillSemanticJobRequest) GetRetryTimes() *int32 {
	return s.RetryTimes
}

func (s *KillSemanticJobRequest) SetExecutorJobId(v string) *KillSemanticJobRequest {
	s.ExecutorJobId = &v
	return s
}

func (s *KillSemanticJobRequest) SetProjectId(v int64) *KillSemanticJobRequest {
	s.ProjectId = &v
	return s
}

func (s *KillSemanticJobRequest) SetRetryTimes(v int32) *KillSemanticJobRequest {
	s.RetryTimes = &v
	return s
}

func (s *KillSemanticJobRequest) Validate() error {
	return dara.Validate(s)
}
