// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportResponse
  GetStatusCode() *int32 
  SetBody(v *ExportResponseBody) *ExportResponse
  GetBody() *ExportResponseBody 
}

type ExportResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportResponse) GoString() string {
  return s.String()
}

func (s *ExportResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportResponse) GetBody() *ExportResponseBody  {
  return s.Body
}

func (s *ExportResponse) SetHeaders(v map[string]*string) *ExportResponse {
  s.Headers = v
  return s
}

func (s *ExportResponse) SetStatusCode(v int32) *ExportResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportResponse) SetBody(v *ExportResponseBody) *ExportResponse {
  s.Body = v
  return s
}

func (s *ExportResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

