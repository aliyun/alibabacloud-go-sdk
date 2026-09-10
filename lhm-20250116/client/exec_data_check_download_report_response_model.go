// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckDownloadReportResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckDownloadReportResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckDownloadReportResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckDownloadReportResponseBody) *ExecDataCheckDownloadReportResponse
  GetBody() *ExecDataCheckDownloadReportResponseBody 
}

type ExecDataCheckDownloadReportResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckDownloadReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckDownloadReportResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckDownloadReportResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckDownloadReportResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckDownloadReportResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckDownloadReportResponse) GetBody() *ExecDataCheckDownloadReportResponseBody  {
  return s.Body
}

func (s *ExecDataCheckDownloadReportResponse) SetHeaders(v map[string]*string) *ExecDataCheckDownloadReportResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckDownloadReportResponse) SetStatusCode(v int32) *ExecDataCheckDownloadReportResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckDownloadReportResponse) SetBody(v *ExecDataCheckDownloadReportResponseBody) *ExecDataCheckDownloadReportResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckDownloadReportResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

