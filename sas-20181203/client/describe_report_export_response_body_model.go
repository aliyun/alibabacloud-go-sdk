// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportExportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDownloadUrl(v string) *DescribeReportExportResponseBody
	GetDownloadUrl() *string
	SetExportDate(v string) *DescribeReportExportResponseBody
	GetExportDate() *string
	SetExportId(v int64) *DescribeReportExportResponseBody
	GetExportId() *int64
	SetExportStatus(v string) *DescribeReportExportResponseBody
	GetExportStatus() *string
	SetReportId(v int64) *DescribeReportExportResponseBody
	GetReportId() *int64
	SetRequestId(v string) *DescribeReportExportResponseBody
	GetRequestId() *string
	SetUrlExpiredTime(v int64) *DescribeReportExportResponseBody
	GetUrlExpiredTime() *int64
}

type DescribeReportExportResponseBody struct {
	// The download URL of the report.
	//
	// example:
	//
	// https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the report was exported.
	//
	// example:
	//
	// 2022-12-15
	ExportDate *string `json:"ExportDate,omitempty" xml:"ExportDate,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 2
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The status of the export task. Valid values:
	//
	// 	- **-1**: The export task fails.
	//
	// 	- **0**: The export task is being initialized.
	//
	// 	- **1**: The export task is being executed.
	//
	// 	- **2**: The export task is successful.
	//
	// example:
	//
	// 2
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// The ID of the report.
	//
	// example:
	//
	// 377665
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp when the download URL expires. Unit: milliseconds.
	//
	// example:
	//
	// 1660113647000
	UrlExpiredTime *int64 `json:"UrlExpiredTime,omitempty" xml:"UrlExpiredTime,omitempty"`
}

func (s DescribeReportExportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportExportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReportExportResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeReportExportResponseBody) GetExportDate() *string {
	return s.ExportDate
}

func (s *DescribeReportExportResponseBody) GetExportId() *int64 {
	return s.ExportId
}

func (s *DescribeReportExportResponseBody) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *DescribeReportExportResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeReportExportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReportExportResponseBody) GetUrlExpiredTime() *int64 {
	return s.UrlExpiredTime
}

func (s *DescribeReportExportResponseBody) SetDownloadUrl(v string) *DescribeReportExportResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetExportDate(v string) *DescribeReportExportResponseBody {
	s.ExportDate = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetExportId(v int64) *DescribeReportExportResponseBody {
	s.ExportId = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetExportStatus(v string) *DescribeReportExportResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetReportId(v int64) *DescribeReportExportResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetRequestId(v string) *DescribeReportExportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReportExportResponseBody) SetUrlExpiredTime(v int64) *DescribeReportExportResponseBody {
	s.UrlExpiredTime = &v
	return s
}

func (s *DescribeReportExportResponseBody) Validate() error {
	return dara.Validate(s)
}
