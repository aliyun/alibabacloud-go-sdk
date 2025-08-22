// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *DescribeDcdnReportListRequest
	GetReportId() *int64
}

type DescribeDcdnReportListRequest struct {
	// The ID of the operations report that you want to query. If you do not specify an ID, all operations reports are queried.
	//
	// example:
	//
	// 2
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeDcdnReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeDcdnReportListRequest) SetReportId(v int64) *DescribeDcdnReportListRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeDcdnReportListRequest) Validate() error {
	return dara.Validate(s)
}
