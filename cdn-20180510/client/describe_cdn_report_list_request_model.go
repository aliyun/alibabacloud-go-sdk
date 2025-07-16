// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *DescribeCdnReportListRequest
	GetReportId() *int64
}

type DescribeCdnReportListRequest struct {
	// The ID of the operations report that you want to query. If you do not specify an ID, all operations reports are queried.
	//
	// example:
	//
	// 1
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeCdnReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportListRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCdnReportListRequest) SetReportId(v int64) *DescribeCdnReportListRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnReportListRequest) Validate() error {
	return dara.Validate(s)
}
