// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightTagRefResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightTagRefResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightTagRefResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightTagRefResponseBody) *ExportInsightTagRefResponse
  GetBody() *ExportInsightTagRefResponseBody 
}

type ExportInsightTagRefResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightTagRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightTagRefResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightTagRefResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightTagRefResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightTagRefResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightTagRefResponse) GetBody() *ExportInsightTagRefResponseBody  {
  return s.Body
}

func (s *ExportInsightTagRefResponse) SetHeaders(v map[string]*string) *ExportInsightTagRefResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightTagRefResponse) SetStatusCode(v int32) *ExportInsightTagRefResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightTagRefResponse) SetBody(v *ExportInsightTagRefResponseBody) *ExportInsightTagRefResponse {
  s.Body = v
  return s
}

func (s *ExportInsightTagRefResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

