// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckReRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckReRunRequest
  GetBatchId() *int64 
}

type ExecDataCheckReRunRequest struct {
  // The batch ID returned by the ExecDataCheckSaveTask operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s ExecDataCheckReRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckReRunRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckReRunRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckReRunRequest) SetBatchId(v int64) *ExecDataCheckReRunRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckReRunRequest) Validate() error {
  return dara.Validate(s)
}

