// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportType(v string) *DescribeBenchmarkTaskReportRequest
	GetReportType() *string
}

type DescribeBenchmarkTaskReportRequest struct {
	// The report type of the stress testing task. Valid values: RAW and Report.
	//
	// example:
	//
	// report
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s DescribeBenchmarkTaskReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportRequest) GetReportType() *string {
	return s.ReportType
}

func (s *DescribeBenchmarkTaskReportRequest) SetReportType(v string) *DescribeBenchmarkTaskReportRequest {
	s.ReportType = &v
	return s
}

func (s *DescribeBenchmarkTaskReportRequest) Validate() error {
	return dara.Validate(s)
}
