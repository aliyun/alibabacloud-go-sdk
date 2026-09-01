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
	SetResourceDirectoryAccountId(v int64) *DescribeCustomizeReportConfigDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetSourceIp(v string) *DescribeCustomizeReportConfigDetailRequest
	GetSourceIp() *string
}

type DescribeCustomizeReportConfigDetailRequest struct {
	// The language type. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The report ID.\\
	//
	// The ReportId returned by calling the [DescribeCustomizeReportList](https://help.aliyun.com/document_detail/271655.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 619031
	ReportId                   *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The IP address of the access source.
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

func (s *DescribeCustomizeReportConfigDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *DescribeCustomizeReportConfigDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeCustomizeReportConfigDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailRequest) SetSourceIp(v string) *DescribeCustomizeReportConfigDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailRequest) Validate() error {
	return dara.Validate(s)
}
