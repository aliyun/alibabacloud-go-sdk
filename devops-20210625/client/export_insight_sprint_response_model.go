// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSprintResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightSprintResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightSprintResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightSprintResponseBody) *ExportInsightSprintResponse
  GetBody() *ExportInsightSprintResponseBody 
}

type ExportInsightSprintResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightSprintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightSprintResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSprintResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightSprintResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightSprintResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightSprintResponse) GetBody() *ExportInsightSprintResponseBody  {
  return s.Body
}

func (s *ExportInsightSprintResponse) SetHeaders(v map[string]*string) *ExportInsightSprintResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightSprintResponse) SetStatusCode(v int32) *ExportInsightSprintResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightSprintResponse) SetBody(v *ExportInsightSprintResponseBody) *ExportInsightSprintResponse {
  s.Body = v
  return s
}

func (s *ExportInsightSprintResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

