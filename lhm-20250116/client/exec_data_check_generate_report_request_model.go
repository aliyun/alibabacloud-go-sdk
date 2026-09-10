// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckGenerateReportRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckGenerateReportRequest
  GetBatchId() *int64 
}

type ExecDataCheckGenerateReportRequest struct {
  // The batch ID returned by the ExecDataCheckSaveTask operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s ExecDataCheckGenerateReportRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckGenerateReportRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckGenerateReportRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckGenerateReportRequest) SetBatchId(v int64) *ExecDataCheckGenerateReportRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckGenerateReportRequest) Validate() error {
  return dara.Validate(s)
}

