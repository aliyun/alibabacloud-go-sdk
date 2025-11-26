// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportApplicationConfigsResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportApplicationConfigsResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportApplicationConfigsResponse
  GetStatusCode() *int32 
  SetBody(v *ExportApplicationConfigsResponseBody) *ExportApplicationConfigsResponse
  GetBody() *ExportApplicationConfigsResponseBody 
}

type ExportApplicationConfigsResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportApplicationConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportApplicationConfigsResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportApplicationConfigsResponse) GoString() string {
  return s.String()
}

func (s *ExportApplicationConfigsResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportApplicationConfigsResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportApplicationConfigsResponse) GetBody() *ExportApplicationConfigsResponseBody  {
  return s.Body
}

func (s *ExportApplicationConfigsResponse) SetHeaders(v map[string]*string) *ExportApplicationConfigsResponse {
  s.Headers = v
  return s
}

func (s *ExportApplicationConfigsResponse) SetStatusCode(v int32) *ExportApplicationConfigsResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportApplicationConfigsResponse) SetBody(v *ExportApplicationConfigsResponseBody) *ExportApplicationConfigsResponse {
  s.Body = v
  return s
}

func (s *ExportApplicationConfigsResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

