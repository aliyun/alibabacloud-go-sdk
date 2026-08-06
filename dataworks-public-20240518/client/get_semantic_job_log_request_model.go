// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorJobId(v string) *GetSemanticJobLogRequest
	GetExecutorJobId() *string
	SetProjectId(v int64) *GetSemanticJobLogRequest
	GetProjectId() *int64
}

type GetSemanticJobLogRequest struct {
	// The executor job ID. Use the Data.ExecutorJobId from the RunSemanticJob response or the ExecutorJobId from a ListSemanticJobRuns record.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The ID of the DataWorks workspace to which the task belongs. Use the ProjectId from the CreateSemanticJob response or a ListSemanticJobs list item.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetSemanticJobLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetSemanticJobLogRequest) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *GetSemanticJobLogRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSemanticJobLogRequest) SetExecutorJobId(v string) *GetSemanticJobLogRequest {
	s.ExecutorJobId = &v
	return s
}

func (s *GetSemanticJobLogRequest) SetProjectId(v int64) *GetSemanticJobLogRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSemanticJobLogRequest) Validate() error {
	return dara.Validate(s)
}
