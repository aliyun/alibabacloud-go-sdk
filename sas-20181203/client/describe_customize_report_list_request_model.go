// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomizeReportListRequest
	GetLang() *string
	SetPinned(v bool) *DescribeCustomizeReportListRequest
	GetPinned() *bool
	SetReportStatus(v int32) *DescribeCustomizeReportListRequest
	GetReportStatus() *int32
	SetReportType(v int32) *DescribeCustomizeReportListRequest
	GetReportType() *int32
	SetReportVersion(v string) *DescribeCustomizeReportListRequest
	GetReportVersion() *string
	SetTitle(v string) *DescribeCustomizeReportListRequest
	GetTitle() *string
}

type DescribeCustomizeReportListRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to pin the report. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	Pinned *bool `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
	// The state of the report. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	ReportStatus *int32 `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	// The type of the report. Valid values:
	//
	// 	- **0**: daily report
	//
	// 	- **1**: weekly report
	//
	// 	- **2**: monthly report
	//
	// 	- **3**: report whose statistics are collected in a custom time range
	//
	// example:
	//
	// 0
	ReportType *int32 `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The report version. Valid values:
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// example:
	//
	// 2.0.0
	ReportVersion *string `json:"ReportVersion,omitempty" xml:"ReportVersion,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeCustomizeReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomizeReportListRequest) GetPinned() *bool {
	return s.Pinned
}

func (s *DescribeCustomizeReportListRequest) GetReportStatus() *int32 {
	return s.ReportStatus
}

func (s *DescribeCustomizeReportListRequest) GetReportType() *int32 {
	return s.ReportType
}

func (s *DescribeCustomizeReportListRequest) GetReportVersion() *string {
	return s.ReportVersion
}

func (s *DescribeCustomizeReportListRequest) GetTitle() *string {
	return s.Title
}

func (s *DescribeCustomizeReportListRequest) SetLang(v string) *DescribeCustomizeReportListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) SetPinned(v bool) *DescribeCustomizeReportListRequest {
	s.Pinned = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) SetReportStatus(v int32) *DescribeCustomizeReportListRequest {
	s.ReportStatus = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) SetReportType(v int32) *DescribeCustomizeReportListRequest {
	s.ReportType = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) SetReportVersion(v string) *DescribeCustomizeReportListRequest {
	s.ReportVersion = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) SetTitle(v string) *DescribeCustomizeReportListRequest {
	s.Title = &v
	return s
}

func (s *DescribeCustomizeReportListRequest) Validate() error {
	return dara.Validate(s)
}
