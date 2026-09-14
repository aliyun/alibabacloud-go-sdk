// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishCrossProjectPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineRunId(v string) *AbolishCrossProjectPipelineRunRequest
	GetPipelineRunId() *string
	SetProjectId(v int64) *AbolishCrossProjectPipelineRunRequest
	GetProjectId() *int64
	SetReason(v string) *AbolishCrossProjectPipelineRunRequest
	GetReason() *string
}

type AbolishCrossProjectPipelineRunRequest struct {
	// The ID of the cross-workspace publish flow.
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
	// The reason for stopping the cross-workspace publish flow.
	//
	// example:
	//
	// The target publish plan has changed. Stop the flow that has not been executed
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s AbolishCrossProjectPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s AbolishCrossProjectPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *AbolishCrossProjectPipelineRunRequest) GetPipelineRunId() *string {
	return s.PipelineRunId
}

func (s *AbolishCrossProjectPipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AbolishCrossProjectPipelineRunRequest) GetReason() *string {
	return s.Reason
}

func (s *AbolishCrossProjectPipelineRunRequest) SetPipelineRunId(v string) *AbolishCrossProjectPipelineRunRequest {
	s.PipelineRunId = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunRequest) SetProjectId(v int64) *AbolishCrossProjectPipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunRequest) SetReason(v string) *AbolishCrossProjectPipelineRunRequest {
	s.Reason = &v
	return s
}

func (s *AbolishCrossProjectPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
