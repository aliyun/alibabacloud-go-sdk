// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharId(v string) *DescribeChartDataRequest
	GetCharId() *string
	SetChartId(v string) *DescribeChartDataRequest
	GetChartId() *string
	SetLang(v string) *DescribeChartDataRequest
	GetLang() *string
	SetReportId(v int64) *DescribeChartDataRequest
	GetReportId() *int64
	SetResourceDirectoryAccountId(v int64) *DescribeChartDataRequest
	GetResourceDirectoryAccountId() *int64
	SetTimeEnd(v int64) *DescribeChartDataRequest
	GetTimeEnd() *int64
	SetTimeStart(v int64) *DescribeChartDataRequest
	GetTimeStart() *int64
}

type DescribeChartDataRequest struct {
	// The chart ID.
	//
	// > Call the [DescribeChartList](~~DescribeChartList~~) operation to obtain this parameter. This parameter is required if the report version is 1.0.0.
	//
	// example:
	//
	// CID_ASSET_RISK_TREND
	CharId *string `json:"CharId,omitempty" xml:"CharId,omitempty"`
	// The chart ID.
	//
	// > Call the [DescribeChartList](~~DescribeChartList~~) operation to obtain this parameter. This parameter is required if the report version is 2.0.0.
	//
	// example:
	//
	// CID_VUL_SUMMARY
	ChartId *string `json:"ChartId,omitempty" xml:"ChartId,omitempty"`
	// The language type for the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The report ID.
	//
	// > Call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to obtain this parameter.
	//
	// example:
	//
	// 721734
	ReportId                   *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The end timestamp for statistics. Unit: milliseconds.
	//
	// example:
	//
	// 1683862286000
	TimeEnd *int64 `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
	// The start timestamp for statistics. Unit: milliseconds.
	//
	// example:
	//
	// 1683603086000
	TimeStart *int64 `json:"TimeStart,omitempty" xml:"TimeStart,omitempty"`
}

func (s DescribeChartDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChartDataRequest) GetCharId() *string {
	return s.CharId
}

func (s *DescribeChartDataRequest) GetChartId() *string {
	return s.ChartId
}

func (s *DescribeChartDataRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeChartDataRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeChartDataRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeChartDataRequest) GetTimeEnd() *int64 {
	return s.TimeEnd
}

func (s *DescribeChartDataRequest) GetTimeStart() *int64 {
	return s.TimeStart
}

func (s *DescribeChartDataRequest) SetCharId(v string) *DescribeChartDataRequest {
	s.CharId = &v
	return s
}

func (s *DescribeChartDataRequest) SetChartId(v string) *DescribeChartDataRequest {
	s.ChartId = &v
	return s
}

func (s *DescribeChartDataRequest) SetLang(v string) *DescribeChartDataRequest {
	s.Lang = &v
	return s
}

func (s *DescribeChartDataRequest) SetReportId(v int64) *DescribeChartDataRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeChartDataRequest) SetResourceDirectoryAccountId(v int64) *DescribeChartDataRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeChartDataRequest) SetTimeEnd(v int64) *DescribeChartDataRequest {
	s.TimeEnd = &v
	return s
}

func (s *DescribeChartDataRequest) SetTimeStart(v int64) *DescribeChartDataRequest {
	s.TimeStart = &v
	return s
}

func (s *DescribeChartDataRequest) Validate() error {
	return dara.Validate(s)
}
