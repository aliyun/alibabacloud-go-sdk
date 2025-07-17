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
  // The request ID.
  // 
  // example:
  // 
  // AFBB799F-8578-51C5-A766-E922EDB8XXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  //     **
  // 
  //     **Note:*	- The value of this parameter indicates only whether the stage is triggered but does not indicate whether the execution of the stage is successful.
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

