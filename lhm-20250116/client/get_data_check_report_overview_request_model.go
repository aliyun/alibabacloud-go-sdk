// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *GetDataCheckReportOverviewRequest
	GetBatchId() *int64
}

type GetDataCheckReportOverviewRequest struct {
	// The ID of the validation job (batch).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s GetDataCheckReportOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportOverviewRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDataCheckReportOverviewRequest) SetBatchId(v int64) *GetDataCheckReportOverviewRequest {
	s.BatchId = &v
	return s
}

func (s *GetDataCheckReportOverviewRequest) Validate() error {
	return dara.Validate(s)
}
