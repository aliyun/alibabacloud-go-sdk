// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditProhibitedDevicesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditProhibitedDevicesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditProhibitedDevicesResponse
  GetStatusCode() *int32 
  SetBody(v *EditProhibitedDevicesResponseBody) *EditProhibitedDevicesResponse
  GetBody() *EditProhibitedDevicesResponseBody 
}

type EditProhibitedDevicesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditProhibitedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditProhibitedDevicesResponse) String() string {
  return dara.Prettify(s)
}

func (s EditProhibitedDevicesResponse) GoString() string {
  return s.String()
}

func (s *EditProhibitedDevicesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditProhibitedDevicesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditProhibitedDevicesResponse) GetBody() *EditProhibitedDevicesResponseBody  {
  return s.Body
}

func (s *EditProhibitedDevicesResponse) SetHeaders(v map[string]*string) *EditProhibitedDevicesResponse {
  s.Headers = v
  return s
}

func (s *EditProhibitedDevicesResponse) SetStatusCode(v int32) *EditProhibitedDevicesResponse {
  s.StatusCode = &v
  return s
}

func (s *EditProhibitedDevicesResponse) SetBody(v *EditProhibitedDevicesResponseBody) *EditProhibitedDevicesResponse {
  s.Body = v
  return s
}

func (s *EditProhibitedDevicesResponse) Validate() error {
  return dara.Validate(s)
}

