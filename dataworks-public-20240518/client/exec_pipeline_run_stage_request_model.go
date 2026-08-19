// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecPipelineRunStageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExecPipelineRunStageRequest
  GetCode() *string 
  SetId(v string) *ExecPipelineRunStageRequest
  GetId() *string 
  SetProjectId(v int64) *ExecPipelineRunStageRequest
  GetProjectId() *int64 
}

type ExecPipelineRunStageRequest struct {
  // The code of the publish flow stage. For the specific value, see the response of the GetPipelineRun operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // DEV_CHECK
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The unique identifier of the publish flow.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // a7ef0634-20ec-4a7c-a214-54020f91XXXX
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
  // The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace settings page to obtain the workspace ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10000
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecPipelineRunStageRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecPipelineRunStageRequest) GoString() string {
  return s.String()
}

func (s *ExecPipelineRunStageRequest) GetCode() *string  {
  return s.Code
}

func (s *ExecPipelineRunStageRequest) GetId() *string  {
  return s.Id
}

func (s *ExecPipelineRunStageRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecPipelineRunStageRequest) SetCode(v string) *ExecPipelineRunStageRequest {
  s.Code = &v
  return s
}

func (s *ExecPipelineRunStageRequest) SetId(v string) *ExecPipelineRunStageRequest {
  s.Id = &v
  return s
}

func (s *ExecPipelineRunStageRequest) SetProjectId(v int64) *ExecPipelineRunStageRequest {
  s.ProjectId = &v
  return s
}

func (s *ExecPipelineRunStageRequest) Validate() error {
  return dara.Validate(s)
}

