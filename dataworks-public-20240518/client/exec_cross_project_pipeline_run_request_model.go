// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecCrossProjectPipelineRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPipelineRunId(v string) *ExecCrossProjectPipelineRunRequest
  GetPipelineRunId() *string 
  SetProjectId(v int64) *ExecCrossProjectPipelineRunRequest
  GetProjectId() *int64 
}

type ExecCrossProjectPipelineRunRequest struct {
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
}

func (s ExecCrossProjectPipelineRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecCrossProjectPipelineRunRequest) GoString() string {
  return s.String()
}

func (s *ExecCrossProjectPipelineRunRequest) GetPipelineRunId() *string  {
  return s.PipelineRunId
}

func (s *ExecCrossProjectPipelineRunRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecCrossProjectPipelineRunRequest) SetPipelineRunId(v string) *ExecCrossProjectPipelineRunRequest {
  s.PipelineRunId = &v
  return s
}

func (s *ExecCrossProjectPipelineRunRequest) SetProjectId(v int64) *ExecCrossProjectPipelineRunRequest {
  s.ProjectId = &v
  return s
}

func (s *ExecCrossProjectPipelineRunRequest) Validate() error {
  return dara.Validate(s)
}

