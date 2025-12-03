// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportInsightSpaceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportInsightSpaceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportInsightSpaceResponse
  GetStatusCode() *int32 
  SetBody(v *ExportInsightSpaceResponseBody) *ExportInsightSpaceResponse
  GetBody() *ExportInsightSpaceResponseBody 
}

type ExportInsightSpaceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportInsightSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportInsightSpaceResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportInsightSpaceResponse) GoString() string {
  return s.String()
}

func (s *ExportInsightSpaceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportInsightSpaceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportInsightSpaceResponse) GetBody() *ExportInsightSpaceResponseBody  {
  return s.Body
}

func (s *ExportInsightSpaceResponse) SetHeaders(v map[string]*string) *ExportInsightSpaceResponse {
  s.Headers = v
  return s
}

func (s *ExportInsightSpaceResponse) SetStatusCode(v int32) *ExportInsightSpaceResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportInsightSpaceResponse) SetBody(v *ExportInsightSpaceResponseBody) *ExportInsightSpaceResponse {
  s.Body = v
  return s
}

func (s *ExportInsightSpaceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

