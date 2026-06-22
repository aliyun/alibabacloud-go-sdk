// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCustomizeReportChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChartIds(v string) *OperationCustomizeReportChartRequest
	GetChartIds() *string
	SetReportId(v int64) *OperationCustomizeReportChartRequest
	GetReportId() *int64
}

type OperationCustomizeReportChartRequest struct {
	// The IDs of report chart configurations. Separate multiple IDs with commas (,).
	//
	// >Call [DescribeChartList](~~DescribeChartList~~) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// CID_VUL_SUMMARY,CID_VUL_TREND,CID_VUL_OPERATION_TREND,CID_BASELINE_CHECK_SUMMARY,CID_BASELINE_CHECK_TREND,CID_BASELINE_CHECK_OPERATION_TREND
	ChartIds *string `json:"ChartIds,omitempty" xml:"ChartIds,omitempty"`
	// The report ID.
	//
	// >Call [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s OperationCustomizeReportChartRequest) String() string {
	return dara.Prettify(s)
}

func (s OperationCustomizeReportChartRequest) GoString() string {
	return s.String()
}

func (s *OperationCustomizeReportChartRequest) GetChartIds() *string {
	return s.ChartIds
}

func (s *OperationCustomizeReportChartRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *OperationCustomizeReportChartRequest) SetChartIds(v string) *OperationCustomizeReportChartRequest {
	s.ChartIds = &v
	return s
}

func (s *OperationCustomizeReportChartRequest) SetReportId(v int64) *OperationCustomizeReportChartRequest {
	s.ReportId = &v
	return s
}

func (s *OperationCustomizeReportChartRequest) Validate() error {
	return dara.Validate(s)
}
