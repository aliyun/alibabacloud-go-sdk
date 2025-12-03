// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightFieldResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightFieldResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightFieldResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightFieldResponseBody) *ExportInsightFieldResponse
  GetBody() *ExportInsightFieldResponseBody 
}

type ExportInsightFieldResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightFieldResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightFieldResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightFieldResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightFieldResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightFieldResponse) GetBody() *ExportInsightFieldResponseBody  {
  return s.Body
}

func (s *ExportInsightFieldResponse) SetHeaders(v map[string]*string) *ExportInsightFieldResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightFieldResponse) SetStatusCode(v int32) *ExportInsightFieldResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightFieldResponse) SetBody(v *ExportInsightFieldResponseBody) *ExportInsightFieldResponse {
  s.Body = v
  return s
}

func (s *ExportInsightFieldResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

