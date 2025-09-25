// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvictTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTaskId(v string) *EvictTaskRequest
  GetTaskId() *string 
}

type EvictTaskRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 17071319
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s EvictTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s EvictTaskRequest) GoString() string {
  return s.String()
}

func (s *EvictTaskRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *EvictTaskRequest) SetTaskId(v string) *EvictTaskRequest {
  s.TaskId = &v
  return s
}

func (s *EvictTaskRequest) Validate() error {
  return dara.Validate(s)
}

