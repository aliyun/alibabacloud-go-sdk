// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkTimeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightWorkTimeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightWorkTimeResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightWorkTimeResponseBody) *ExportInsightWorkTimeResponse
  GetBody() *ExportInsightWorkTimeResponseBody 
}

type ExportInsightWorkTimeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightWorkTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightWorkTimeResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkTimeResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkTimeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightWorkTimeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightWorkTimeResponse) GetBody() *ExportInsightWorkTimeResponseBody  {
  return s.Body
}

func (s *ExportInsightWorkTimeResponse) SetHeaders(v map[string]*string) *ExportInsightWorkTimeResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightWorkTimeResponse) SetStatusCode(v int32) *ExportInsightWorkTimeResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightWorkTimeResponse) SetBody(v *ExportInsightWorkTimeResponseBody) *ExportInsightWorkTimeResponse {
  s.Body = v
  return s
}

func (s *ExportInsightWorkTimeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

