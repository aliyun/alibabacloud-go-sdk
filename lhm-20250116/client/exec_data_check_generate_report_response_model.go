// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckGenerateReportResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecDataCheckGenerateReportResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecDataCheckGenerateReportResponse
  GetStatusCode() *int32 
  SetBody(v *ExecDataCheckGenerateReportResponseBody) *ExecDataCheckGenerateReportResponse
  GetBody() *ExecDataCheckGenerateReportResponseBody 
}

type ExecDataCheckGenerateReportResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecDataCheckGenerateReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDataCheckGenerateReportResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckGenerateReportResponse) GoString() string {
  return s.String()
}

func (s *ExecDataCheckGenerateReportResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecDataCheckGenerateReportResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecDataCheckGenerateReportResponse) GetBody() *ExecDataCheckGenerateReportResponseBody  {
  return s.Body
}

func (s *ExecDataCheckGenerateReportResponse) SetHeaders(v map[string]*string) *ExecDataCheckGenerateReportResponse {
  s.Headers = v
  return s
}

func (s *ExecDataCheckGenerateReportResponse) SetStatusCode(v int32) *ExecDataCheckGenerateReportResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecDataCheckGenerateReportResponse) SetBody(v *ExecDataCheckGenerateReportResponseBody) *ExecDataCheckGenerateReportResponse {
  s.Body = v
  return s
}

func (s *ExecDataCheckGenerateReportResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

