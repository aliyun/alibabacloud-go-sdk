// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportNacosConfigResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportNacosConfigResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportNacosConfigResponse
  GetStatusCode() *int32 
  SetBody(v *ExportNacosConfigResponseBody) *ExportNacosConfigResponse
  GetBody() *ExportNacosConfigResponseBody 
}

type ExportNacosConfigResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportNacosConfigResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportNacosConfigResponse) GoString() string {
  return s.String()
}

func (s *ExportNacosConfigResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportNacosConfigResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportNacosConfigResponse) GetBody() *ExportNacosConfigResponseBody  {
  return s.Body
}

func (s *ExportNacosConfigResponse) SetHeaders(v map[string]*string) *ExportNacosConfigResponse {
  s.Headers = v
  return s
}

func (s *ExportNacosConfigResponse) SetStatusCode(v int32) *ExportNacosConfigResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportNacosConfigResponse) SetBody(v *ExportNacosConfigResponseBody) *ExportNacosConfigResponse {
  s.Body = v
  return s
}

func (s *ExportNacosConfigResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

