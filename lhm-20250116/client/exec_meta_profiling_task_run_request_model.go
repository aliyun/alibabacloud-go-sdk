// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaProfilingTaskRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchNo(v string) *ExecMetaProfilingTaskRunRequest
  GetBatchNo() *string 
  SetComponentId(v int64) *ExecMetaProfilingTaskRunRequest
  GetComponentId() *int64 
}

type ExecMetaProfilingTaskRunRequest struct {
  // example:
  // 
  // 1/0
  BatchNo *string `json:"batchNo,omitempty" xml:"batchNo,omitempty"`
  // example:
  // 
  // 123
  ComponentId *int64 `json:"componentId,omitempty" xml:"componentId,omitempty"`
}

func (s ExecMetaProfilingTaskRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaProfilingTaskRunRequest) GoString() string {
  return s.String()
}

func (s *ExecMetaProfilingTaskRunRequest) GetBatchNo() *string  {
  return s.BatchNo
}

func (s *ExecMetaProfilingTaskRunRequest) GetComponentId() *int64  {
  return s.ComponentId
}

func (s *ExecMetaProfilingTaskRunRequest) SetBatchNo(v string) *ExecMetaProfilingTaskRunRequest {
  s.BatchNo = &v
  return s
}

func (s *ExecMetaProfilingTaskRunRequest) SetComponentId(v int64) *ExecMetaProfilingTaskRunRequest {
  s.ComponentId = &v
  return s
}

func (s *ExecMetaProfilingTaskRunRequest) Validate() error {
  return dara.Validate(s)
}

