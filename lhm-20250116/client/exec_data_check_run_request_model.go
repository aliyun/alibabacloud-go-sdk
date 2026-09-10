// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckRunRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckRunRequest
  GetBatchId() *int64 
}

type ExecDataCheckRunRequest struct {
  // The check job ID.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s ExecDataCheckRunRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckRunRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckRunRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckRunRequest) SetBatchId(v int64) *ExecDataCheckRunRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckRunRequest) Validate() error {
  return dara.Validate(s)
}

