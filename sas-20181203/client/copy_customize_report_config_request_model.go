// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCustomizeReportConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CopyCustomizeReportConfigRequest
	GetLang() *string
	SetReportId(v int64) *CopyCustomizeReportConfigRequest
	GetReportId() *int64
	SetSourceIp(v string) *CopyCustomizeReportConfigRequest
	GetSourceIp() *string
}

type CopyCustomizeReportConfigRequest struct {
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
	// The ID of the report.
	//
	// >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 492742
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 59.46.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CopyCustomizeReportConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyCustomizeReportConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyCustomizeReportConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *CopyCustomizeReportConfigRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *CopyCustomizeReportConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CopyCustomizeReportConfigRequest) SetLang(v string) *CopyCustomizeReportConfigRequest {
	s.Lang = &v
	return s
}

func (s *CopyCustomizeReportConfigRequest) SetReportId(v int64) *CopyCustomizeReportConfigRequest {
	s.ReportId = &v
	return s
}

func (s *CopyCustomizeReportConfigRequest) SetSourceIp(v string) *CopyCustomizeReportConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CopyCustomizeReportConfigRequest) Validate() error {
	return dara.Validate(s)
}
