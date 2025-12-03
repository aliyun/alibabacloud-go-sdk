// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightExpectedWorkTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightExpectedWorkTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightExpectedWorkTimeResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightExpectedWorkTimeResponseBody) *ExportInsightExpectedWorkTimeResponse
  GetBody() *ExportInsightExpectedWorkTimeResponseBody 
}

type ExportInsightExpectedWorkTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightExpectedWorkTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightExpectedWorkTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightExpectedWorkTimeResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightExpectedWorkTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightExpectedWorkTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightExpectedWorkTimeResponse) GetBody() *ExportInsightExpectedWorkTimeResponseBody  {
  return s.Body
}

func (s *ExportInsightExpectedWorkTimeResponse) SetHeaders(v map[string]*string) *ExportInsightExpectedWorkTimeResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponse) SetStatusCode(v int32) *ExportInsightExpectedWorkTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponse) SetBody(v *ExportInsightExpectedWorkTimeResponseBody) *ExportInsightExpectedWorkTimeResponse {
  s.Body = v
  return s
}

func (s *ExportInsightExpectedWorkTimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

