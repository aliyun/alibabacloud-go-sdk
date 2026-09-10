// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunFailedRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckRunFailedRequest
  GetBatchId() *int64 
  SetType(v int32) *ExecDataCheckRunFailedRequest
  GetType() *int32 
}

type ExecDataCheckRunFailedRequest struct {
  // The batch ID returned by the ExecDataCheckSaveTask operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
  // The rerun type. Valid values:
  // 
  // - 0: Reruns only execution-failed subtasks.
  // 
  // - 1: Reruns execution-failed and validation-failed subtasks.
  // 
  // - 2: Reruns execution-failed and stopped subtasks.
  // 
  // Default value: 1.
  // 
  // example:
  // 
  // 0
  Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ExecDataCheckRunFailedRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunFailedRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunFailedRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckRunFailedRequest) GetType() *int32  {
  return s.Type
}

func (s *ExecDataCheckRunFailedRequest) SetBatchId(v int64) *ExecDataCheckRunFailedRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckRunFailedRequest) SetType(v int32) *ExecDataCheckRunFailedRequest {
  s.Type = &v
  return s
}

func (s *ExecDataCheckRunFailedRequest) Validate() error {
  return dara.Validate(s)
}

