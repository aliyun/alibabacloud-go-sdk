// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckStopRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckStopRequest
  GetBatchId() *int64 
}

type ExecDataCheckStopRequest struct {
  // The check job ID.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s ExecDataCheckStopRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckStopRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckStopRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckStopRequest) SetBatchId(v int64) *ExecDataCheckStopRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckStopRequest) Validate() error {
  return dara.Validate(s)
}

