// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightCustomValueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightCustomValueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightCustomValueResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightCustomValueResponseBody) *ExportInsightCustomValueResponse
  GetBody() *ExportInsightCustomValueResponseBody 
}

type ExportInsightCustomValueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightCustomValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightCustomValueResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightCustomValueResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightCustomValueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightCustomValueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightCustomValueResponse) GetBody() *ExportInsightCustomValueResponseBody  {
  return s.Body
}

func (s *ExportInsightCustomValueResponse) SetHeaders(v map[string]*string) *ExportInsightCustomValueResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightCustomValueResponse) SetStatusCode(v int32) *ExportInsightCustomValueResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightCustomValueResponse) SetBody(v *ExportInsightCustomValueResponseBody) *ExportInsightCustomValueResponse {
  s.Body = v
  return s
}

func (s *ExportInsightCustomValueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

