// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorJobId(v string) *GetSemanticJobDetailRequest
	GetExecutorJobId() *string
	SetProjectId(v int64) *GetSemanticJobDetailRequest
	GetProjectId() *int64
}

type GetSemanticJobDetailRequest struct {
	// The executor job ID. Use the Data.ExecutorJobId from the RunSemanticJob response or the ExecutorJobId from a ListSemanticJobRuns record.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The ID of the DataWorks workspace to which the job belongs. Use the ProjectId from the CreateSemanticJob response or a ListSemanticJobs list item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetSemanticJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetSemanticJobDetailRequest) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *GetSemanticJobDetailRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSemanticJobDetailRequest) SetExecutorJobId(v string) *GetSemanticJobDetailRequest {
	s.ExecutorJobId = &v
	return s
}

func (s *GetSemanticJobDetailRequest) SetProjectId(v int64) *GetSemanticJobDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSemanticJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
