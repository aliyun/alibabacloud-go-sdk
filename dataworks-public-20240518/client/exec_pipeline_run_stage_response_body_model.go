// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecPipelineRunStageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecPipelineRunStageResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecPipelineRunStageResponseBody
  GetSuccess() *bool 
}

type ExecPipelineRunStageResponseBody struct {
  // The request ID. Used for locating logs and troubleshooting issues.
  // 
  // example:
  // 
  // AFBB799F-8578-51C5-A766-E922EDB8XXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the call is successful. Valid values:
  // 
  // - true: The call is successful.
  // 
  // - false: The call failed.
  // 
  // 	Notice: This only indicates whether the stage is triggered, not the execution result of the publish stage.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecPipelineRunStageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecPipelineRunStageResponseBody) GoString() string {
  return s.String()
}

func (s *ExecPipelineRunStageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecPipelineRunStageResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecPipelineRunStageResponseBody) SetRequestId(v string) *ExecPipelineRunStageResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecPipelineRunStageResponseBody) SetSuccess(v bool) *ExecPipelineRunStageResponseBody {
  s.Success = &v
  return s
}

func (s *ExecPipelineRunStageResponseBody) Validate() error {
  return dara.Validate(s)
}

