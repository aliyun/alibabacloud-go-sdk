// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkitemActivityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportWorkitemActivityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportWorkitemActivityResponse
  GetStatusCode() *int32 
  SetBody(v *ExportWorkitemActivityResponseBody) *ExportWorkitemActivityResponse
  GetBody() *ExportWorkitemActivityResponseBody 
}

type ExportWorkitemActivityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportWorkitemActivityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportWorkitemActivityResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkitemActivityResponse) GoString() string {
  return s.String()
}

func (s *ExportWorkitemActivityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportWorkitemActivityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportWorkitemActivityResponse) GetBody() *ExportWorkitemActivityResponseBody  {
  return s.Body
}

func (s *ExportWorkitemActivityResponse) SetHeaders(v map[string]*string) *ExportWorkitemActivityResponse {
  s.Headers = v
  return s
}

func (s *ExportWorkitemActivityResponse) SetStatusCode(v int32) *ExportWorkitemActivityResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportWorkitemActivityResponse) SetBody(v *ExportWorkitemActivityResponseBody) *ExportWorkitemActivityResponse {
  s.Body = v
  return s
}

func (s *ExportWorkitemActivityResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

