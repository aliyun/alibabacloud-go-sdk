// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemVersionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightWorkitemVersionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightWorkitemVersionResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightWorkitemVersionResponseBody) *ExportInsightWorkitemVersionResponse
  GetBody() *ExportInsightWorkitemVersionResponseBody 
}

type ExportInsightWorkitemVersionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightWorkitemVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightWorkitemVersionResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemVersionResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemVersionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightWorkitemVersionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightWorkitemVersionResponse) GetBody() *ExportInsightWorkitemVersionResponseBody  {
  return s.Body
}

func (s *ExportInsightWorkitemVersionResponse) SetHeaders(v map[string]*string) *ExportInsightWorkitemVersionResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightWorkitemVersionResponse) SetStatusCode(v int32) *ExportInsightWorkitemVersionResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightWorkitemVersionResponse) SetBody(v *ExportInsightWorkitemVersionResponseBody) *ExportInsightWorkitemVersionResponse {
  s.Body = v
  return s
}

func (s *ExportInsightWorkitemVersionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

