// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *GetDataCheckReportStatusRequest
	GetBatchId() *int64
}

type GetDataCheckReportStatusRequest struct {
	// The batch ID returned by the ExecDataCheckSaveTask operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s GetDataCheckReportStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportStatusRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDataCheckReportStatusRequest) SetBatchId(v int64) *GetDataCheckReportStatusRequest {
	s.BatchId = &v
	return s
}

func (s *GetDataCheckReportStatusRequest) Validate() error {
	return dara.Validate(s)
}
