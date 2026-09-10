// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckDownloadReportRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBatchId(v int64) *ExecDataCheckDownloadReportRequest
  GetBatchId() *int64 
}

type ExecDataCheckDownloadReportRequest struct {
  // The batch ID. Prerequisite: the report status must be 2 (Generated).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20001
  BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s ExecDataCheckDownloadReportRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckDownloadReportRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckDownloadReportRequest) GetBatchId() *int64  {
  return s.BatchId
}

func (s *ExecDataCheckDownloadReportRequest) SetBatchId(v int64) *ExecDataCheckDownloadReportRequest {
  s.BatchId = &v
  return s
}

func (s *ExecDataCheckDownloadReportRequest) Validate() error {
  return dara.Validate(s)
}

