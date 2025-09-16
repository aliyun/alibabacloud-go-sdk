// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTaskId(v int32) *ExecuteTaskRequest
  GetTaskId() *int32 
}

type ExecuteTaskRequest struct {
  // example:
  // 
  // 2615
  TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTaskRequest) GetTaskId() *int32  {
  return s.TaskId
}

func (s *ExecuteTaskRequest) SetTaskId(v int32) *ExecuteTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExecuteTaskRequest) Validate() error {
  return dara.Validate(s)
}

