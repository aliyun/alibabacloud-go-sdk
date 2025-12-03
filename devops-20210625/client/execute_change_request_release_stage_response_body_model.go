// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteChangeRequestReleaseStageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetObject(v int64) *ExecuteChangeRequestReleaseStageResponseBody
  GetObject() *int64 
  SetPipelineId(v int64) *ExecuteChangeRequestReleaseStageResponseBody
  GetPipelineId() *int64 
  SetPipelineRunId(v int64) *ExecuteChangeRequestReleaseStageResponseBody
  GetPipelineRunId() *int64 
}

type ExecuteChangeRequestReleaseStageResponseBody struct {
  // example:
  // 
  // 1
  Object *int64 `json:"object,omitempty" xml:"object,omitempty"`
  // example:
  // 
  // 3259***
  PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
  // example:
  // 
  // 1
  PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
}

func (s ExecuteChangeRequestReleaseStageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteChangeRequestReleaseStageResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) GetObject() *int64  {
  return s.Object
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) GetPipelineId() *int64  {
  return s.PipelineId
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) GetPipelineRunId() *int64  {
  return s.PipelineRunId
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) SetObject(v int64) *ExecuteChangeRequestReleaseStageResponseBody {
  s.Object = &v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) SetPipelineId(v int64) *ExecuteChangeRequestReleaseStageResponseBody {
  s.PipelineId = &v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) SetPipelineRunId(v int64) *ExecuteChangeRequestReleaseStageResponseBody {
  s.PipelineRunId = &v
  return s
}

func (s *ExecuteChangeRequestReleaseStageResponseBody) Validate() error {
  return dara.Validate(s)
}

