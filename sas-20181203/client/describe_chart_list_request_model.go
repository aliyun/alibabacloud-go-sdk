// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeChartListRequest
	GetLang() *string
	SetProjectCode(v string) *DescribeChartListRequest
	GetProjectCode() *string
	SetReportId(v string) *DescribeChartListRequest
	GetReportId() *string
}

type DescribeChartListRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The code of the report. Valid value:
	//
	// 	- **customize_report**
	//
	// This parameter is required.
	//
	// example:
	//
	// customize_report
	ProjectCode *string `json:"ProjectCode,omitempty" xml:"ProjectCode,omitempty"`
	// The ID of the report.
	//
	// >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
	//
	// example:
	//
	// 720549
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeChartListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartListRequest) GoString() string {
	return s.String()
}

func (s *DescribeChartListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeChartListRequest) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *DescribeChartListRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeChartListRequest) SetLang(v string) *DescribeChartListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeChartListRequest) SetProjectCode(v string) *DescribeChartListRequest {
	s.ProjectCode = &v
	return s
}

func (s *DescribeChartListRequest) SetReportId(v string) *DescribeChartListRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeChartListRequest) Validate() error {
	return dara.Validate(s)
}
