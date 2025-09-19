// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomizeReportResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportCustomizeReportResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportCustomizeReportResponse
  GetStatusCode() *int32 
  SetBody(v *ExportCustomizeReportResponseBody) *ExportCustomizeReportResponse
  GetBody() *ExportCustomizeReportResponseBody 
}

type ExportCustomizeReportResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportCustomizeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportCustomizeReportResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomizeReportResponse) GoString() string {
  return s.String()
}

func (s *ExportCustomizeReportResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportCustomizeReportResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportCustomizeReportResponse) GetBody() *ExportCustomizeReportResponseBody  {
  return s.Body
}

func (s *ExportCustomizeReportResponse) SetHeaders(v map[string]*string) *ExportCustomizeReportResponse {
  s.Headers = v
  return s
}

func (s *ExportCustomizeReportResponse) SetStatusCode(v int32) *ExportCustomizeReportResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportCustomizeReportResponse) SetBody(v *ExportCustomizeReportResponseBody) *ExportCustomizeReportResponse {
  s.Body = v
  return s
}

func (s *ExportCustomizeReportResponse) Validate() error {
  return dara.Validate(s)
}

