// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *DescribeReportExportRequest
	GetExportId() *int64
	SetLang(v string) *DescribeReportExportRequest
	GetLang() *string
}

type DescribeReportExportRequest struct {
	// The ID of the export task.
	//
	// >  You can call the [ExportCustomizeReport](https://help.aliyun.com/document_detail/2842677.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
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
}

func (s DescribeReportExportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportExportRequest) GoString() string {
	return s.String()
}

func (s *DescribeReportExportRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeReportExportRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeReportExportRequest) SetExportId(v int64) *DescribeReportExportRequest {
	s.ExportId = &v
	return s
}

func (s *DescribeReportExportRequest) SetLang(v string) *DescribeReportExportRequest {
	s.Lang = &v
	return s
}

func (s *DescribeReportExportRequest) Validate() error {
	return dara.Validate(s)
}
