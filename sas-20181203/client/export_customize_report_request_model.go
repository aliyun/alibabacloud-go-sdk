// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomizeReportRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExportType(v string) *ExportCustomizeReportRequest
  GetExportType() *string 
  SetReportId(v int64) *ExportCustomizeReportRequest
  GetReportId() *int64 
}

type ExportCustomizeReportRequest struct {
  // The type of the security report that you want to export. Valid values:
  // 
  // 	- **HTML**
  // 
  // 	- **PDF**
  // 
  // >  The default value is HTML. PDF is supported only for security reports in version 2.0.0.
  // 
  // example:
  // 
  // HTML
  ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
  // The ID of the security report.
  // 
  // >  You can call the [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) operation to query the ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s ExportCustomizeReportRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomizeReportRequest) GoString() string {
  return s.String()
}

func (s *ExportCustomizeReportRequest) GetExportType() *string  {
  return s.ExportType
}

func (s *ExportCustomizeReportRequest) GetReportId() *int64  {
  return s.ReportId
}

func (s *ExportCustomizeReportRequest) SetExportType(v string) *ExportCustomizeReportRequest {
  s.ExportType = &v
  return s
}

func (s *ExportCustomizeReportRequest) SetReportId(v int64) *ExportCustomizeReportRequest {
  s.ReportId = &v
  return s
}

func (s *ExportCustomizeReportRequest) Validate() error {
  return dara.Validate(s)
}

