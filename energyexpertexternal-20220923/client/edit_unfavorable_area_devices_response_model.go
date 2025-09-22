// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditUnfavorableAreaDevicesResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditUnfavorableAreaDevicesResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditUnfavorableAreaDevicesResponse
  GetStatusCode() *int32 
  SetBody(v *EditUnfavorableAreaDevicesResponseBody) *EditUnfavorableAreaDevicesResponse
  GetBody() *EditUnfavorableAreaDevicesResponseBody 
}

type EditUnfavorableAreaDevicesResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditUnfavorableAreaDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditUnfavorableAreaDevicesResponse) String() string {
  return dara.Prettify(s)
}

func (s EditUnfavorableAreaDevicesResponse) GoString() string {
  return s.String()
}

func (s *EditUnfavorableAreaDevicesResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditUnfavorableAreaDevicesResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditUnfavorableAreaDevicesResponse) GetBody() *EditUnfavorableAreaDevicesResponseBody  {
  return s.Body
}

func (s *EditUnfavorableAreaDevicesResponse) SetHeaders(v map[string]*string) *EditUnfavorableAreaDevicesResponse {
  s.Headers = v
  return s
}

func (s *EditUnfavorableAreaDevicesResponse) SetStatusCode(v int32) *EditUnfavorableAreaDevicesResponse {
  s.StatusCode = &v
  return s
}

func (s *EditUnfavorableAreaDevicesResponse) SetBody(v *EditUnfavorableAreaDevicesResponseBody) *EditUnfavorableAreaDevicesResponse {
  s.Body = v
  return s
}

func (s *EditUnfavorableAreaDevicesResponse) Validate() error {
  return dara.Validate(s)
}

