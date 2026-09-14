// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrossProjectPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineRunId(v string) *GetCrossProjectPipelineRunRequest
	GetPipelineRunId() *string
	SetProjectId(v int64) *GetCrossProjectPipelineRunRequest
	GetProjectId() *int64
}

type GetCrossProjectPipelineRunRequest struct {
	// The cross-workspace deployment flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fcfd4160-e2ff-4603-9719-09128fe733df
	PipelineRunId *string `json:"PipelineRunId,omitempty" xml:"PipelineRunId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetCrossProjectPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCrossProjectPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *GetCrossProjectPipelineRunRequest) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *GetCrossProjectPipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCrossProjectPipelineRunRequest) SetPipelineRunId(v string) *GetCrossProjectPipelineRunRequest {
	s.PipelineRunId = &v
	return s
}

func (s *GetCrossProjectPipelineRunRequest) SetProjectId(v int64) *GetCrossProjectPipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCrossProjectPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
