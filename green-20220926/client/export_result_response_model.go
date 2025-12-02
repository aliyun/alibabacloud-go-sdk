// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportResultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportResultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportResultResponse
  GetStatusCode() *int32 
  SetBody(v *ExportResultResponseBody) *ExportResultResponse
  GetBody() *ExportResultResponseBody 
}

type ExportResultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportResultResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportResultResponse) GoString() string {
  return s.String()
}

func (s *ExportResultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportResultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportResultResponse) GetBody() *ExportResultResponseBody  {
  return s.Body
}

func (s *ExportResultResponse) SetHeaders(v map[string]*string) *ExportResultResponse {
  s.Headers = v
  return s
}

func (s *ExportResultResponse) SetStatusCode(v int32) *ExportResultResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportResultResponse) SetBody(v *ExportResultResponseBody) *ExportResultResponse {
  s.Body = v
  return s
}

func (s *ExportResultResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

