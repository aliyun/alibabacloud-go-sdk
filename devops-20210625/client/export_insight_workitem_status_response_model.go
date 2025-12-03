// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightWorkitemStatusResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightWorkitemStatusResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightWorkitemStatusResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightWorkitemStatusResponseBody) *ExportInsightWorkitemStatusResponse
  GetBody() *ExportInsightWorkitemStatusResponseBody 
}

type ExportInsightWorkitemStatusResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightWorkitemStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightWorkitemStatusResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightWorkitemStatusResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightWorkitemStatusResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightWorkitemStatusResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightWorkitemStatusResponse) GetBody() *ExportInsightWorkitemStatusResponseBody  {
  return s.Body
}

func (s *ExportInsightWorkitemStatusResponse) SetHeaders(v map[string]*string) *ExportInsightWorkitemStatusResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightWorkitemStatusResponse) SetStatusCode(v int32) *ExportInsightWorkitemStatusResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightWorkitemStatusResponse) SetBody(v *ExportInsightWorkitemStatusResponseBody) *ExportInsightWorkitemStatusResponse {
  s.Body = v
  return s
}

func (s *ExportInsightWorkitemStatusResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

