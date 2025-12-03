// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceRefResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightSpaceRefResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightSpaceRefResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightSpaceRefResponseBody) *ExportInsightSpaceRefResponse
  GetBody() *ExportInsightSpaceRefResponseBody 
}

type ExportInsightSpaceRefResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightSpaceRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightSpaceRefResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceRefResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceRefResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightSpaceRefResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightSpaceRefResponse) GetBody() *ExportInsightSpaceRefResponseBody  {
  return s.Body
}

func (s *ExportInsightSpaceRefResponse) SetHeaders(v map[string]*string) *ExportInsightSpaceRefResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightSpaceRefResponse) SetStatusCode(v int32) *ExportInsightSpaceRefResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightSpaceRefResponse) SetBody(v *ExportInsightSpaceRefResponseBody) *ExportInsightSpaceRefResponse {
  s.Body = v
  return s
}

func (s *ExportInsightSpaceRefResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

