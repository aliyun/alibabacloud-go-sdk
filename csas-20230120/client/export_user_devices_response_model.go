// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportUserDevicesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportUserDevicesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportUserDevicesResponse
  GetStatusCode() *int32 
  SetBody(v *ExportUserDevicesResponseBody) *ExportUserDevicesResponse
  GetBody() *ExportUserDevicesResponseBody 
}

type ExportUserDevicesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportUserDevicesResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportUserDevicesResponse) GoString() string {
  return s.String()
}

func (s *ExportUserDevicesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportUserDevicesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportUserDevicesResponse) GetBody() *ExportUserDevicesResponseBody  {
  return s.Body
}

func (s *ExportUserDevicesResponse) SetHeaders(v map[string]*string) *ExportUserDevicesResponse {
  s.Headers = v
  return s
}

func (s *ExportUserDevicesResponse) SetStatusCode(v int32) *ExportUserDevicesResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportUserDevicesResponse) SetBody(v *ExportUserDevicesResponseBody) *ExportUserDevicesResponse {
  s.Body = v
  return s
}

func (s *ExportUserDevicesResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

