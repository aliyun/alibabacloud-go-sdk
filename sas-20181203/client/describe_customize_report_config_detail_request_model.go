// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportConfigDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustomizeReportConfigDetailRequest
	GetLang() *string
	SetReportId(v int64) *DescribeCustomizeReportConfigDetailRequest
	GetReportId() *int64
	SetSourceIp(v string) *DescribeCustomizeReportConfigDetailRequest
	GetSourceIp() *string
}

type DescribeCustomizeReportConfigDetailRequest struct {
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
	// The ID of the report.\\
	//
	// You can call the [DescribeCustomizeReportList](https://help.aliyun.com/document_detail/271655.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 619031
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 202.104.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCustomizeReportConfigDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportConfigDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportConfigDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustomizeReportConfigDetailRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCustomizeReportConfigDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCustomizeReportConfigDetailRequest) SetLang(v string) *DescribeCustomizeReportConfigDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailRequest) SetReportId(v int64) *DescribeCustomizeReportConfigDetailRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailRequest) SetSourceIp(v string) *DescribeCustomizeReportConfigDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailRequest) Validate() error {
	return dara.Validate(s)
}
