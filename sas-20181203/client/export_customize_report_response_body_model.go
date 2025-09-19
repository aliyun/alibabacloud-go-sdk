// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomizeReportResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDownloadUrl(v string) *ExportCustomizeReportResponseBody
  GetDownloadUrl() *string 
  SetExportDate(v string) *ExportCustomizeReportResponseBody
  GetExportDate() *string 
  SetExportId(v int64) *ExportCustomizeReportResponseBody
  GetExportId() *int64 
  SetExportStatus(v string) *ExportCustomizeReportResponseBody
  GetExportStatus() *string 
  SetFileName(v string) *ExportCustomizeReportResponseBody
  GetFileName() *string 
  SetReportId(v int64) *ExportCustomizeReportResponseBody
  GetReportId() *int64 
  SetRequestId(v string) *ExportCustomizeReportResponseBody
  GetRequestId() *string 
  SetUrlExpiredTime(v int64) *ExportCustomizeReportResponseBody
  GetUrlExpiredTime() *int64 
}

type ExportCustomizeReportResponseBody struct {
  // The download URL of the security report.
  // 
  // example:
  // 
  // https://xxxxxxxx.oss-cn-hangzhou-1.aliyuncs.com/xxxxx/xxxxxxxxxxxxxx?Expires=1671448125&OSSAccessKeyId=xxx
  DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
  // The time when the security report was exported.
  // 
  // example:
  // 
  // 2023-01-10
  ExportDate *string `json:"ExportDate,omitempty" xml:"ExportDate,omitempty"`
  // The ID of the export task.
  // 
  // example:
  // 
  // 22
  ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
  // The status of the export task. Valid values:
  // 
  // 	- **fail**: The export task fails.
  // 
  // 	- **exporting**: The export task is being executed.
  // 
  // 	- **success**: The export task is successful.
  // 
  // example:
  // 
  // exporting
  ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
  // The name of the report file that is exported.
  // 
  // example:
  // 
  // test_xxxx.html
  FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
  // The ID of the security report.
  // 
  // example:
  // 
  // 123
  ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // FFDFCEB3-A5EE-590A-8E70-283EBC5D****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The timestamp when the download URL expires. Unit: seconds.
  // 
  // example:
  // 
  // 1673335497000
  UrlExpiredTime *int64 `json:"UrlExpiredTime,omitempty" xml:"UrlExpiredTime,omitempty"`
}

func (s ExportCustomizeReportResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomizeReportResponseBody) GoString() string {
  return s.String()
}

func (s *ExportCustomizeReportResponseBody) GetDownloadUrl() *string  {
  return s.DownloadUrl
}

func (s *ExportCustomizeReportResponseBody) GetExportDate() *string  {
  return s.ExportDate
}

func (s *ExportCustomizeReportResponseBody) GetExportId() *int64  {
  return s.ExportId
}

func (s *ExportCustomizeReportResponseBody) GetExportStatus() *string  {
  return s.ExportStatus
}

func (s *ExportCustomizeReportResponseBody) GetFileName() *string  {
  return s.FileName
}

func (s *ExportCustomizeReportResponseBody) GetReportId() *int64  {
  return s.ReportId
}

func (s *ExportCustomizeReportResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportCustomizeReportResponseBody) GetUrlExpiredTime() *int64  {
  return s.UrlExpiredTime
}

func (s *ExportCustomizeReportResponseBody) SetDownloadUrl(v string) *ExportCustomizeReportResponseBody {
  s.DownloadUrl = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetExportDate(v string) *ExportCustomizeReportResponseBody {
  s.ExportDate = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetExportId(v int64) *ExportCustomizeReportResponseBody {
  s.ExportId = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetExportStatus(v string) *ExportCustomizeReportResponseBody {
  s.ExportStatus = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetFileName(v string) *ExportCustomizeReportResponseBody {
  s.FileName = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetReportId(v int64) *ExportCustomizeReportResponseBody {
  s.ReportId = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetRequestId(v string) *ExportCustomizeReportResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) SetUrlExpiredTime(v int64) *ExportCustomizeReportResponseBody {
  s.UrlExpiredTime = &v
  return s
}

func (s *ExportCustomizeReportResponseBody) Validate() error {
  return dara.Validate(s)
}

