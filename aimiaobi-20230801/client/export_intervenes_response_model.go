// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportIntervenesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportIntervenesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportIntervenesResponse
  GetStatusCode() *int32 
  SetBody(v *ExportIntervenesResponseBody) *ExportIntervenesResponse
  GetBody() *ExportIntervenesResponseBody 
}

type ExportIntervenesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportIntervenesResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportIntervenesResponse) GoString() string {
  return s.String()
}

func (s *ExportIntervenesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportIntervenesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportIntervenesResponse) GetBody() *ExportIntervenesResponseBody  {
  return s.Body
}

func (s *ExportIntervenesResponse) SetHeaders(v map[string]*string) *ExportIntervenesResponse {
  s.Headers = v
  return s
}

func (s *ExportIntervenesResponse) SetStatusCode(v int32) *ExportIntervenesResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportIntervenesResponse) SetBody(v *ExportIntervenesResponseBody) *ExportIntervenesResponse {
  s.Body = v
  return s
}

func (s *ExportIntervenesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

