// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDeviceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDeviceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDeviceResponseBody) *EnableDeviceResponse
  GetBody() *EnableDeviceResponseBody 
}

type EnableDeviceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDeviceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceResponse) GoString() string {
  return s.String()
}

func (s *EnableDeviceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDeviceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDeviceResponse) GetBody() *EnableDeviceResponseBody  {
  return s.Body
}

func (s *EnableDeviceResponse) SetHeaders(v map[string]*string) *EnableDeviceResponse {
  s.Headers = v
  return s
}

func (s *EnableDeviceResponse) SetStatusCode(v int32) *EnableDeviceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDeviceResponse) SetBody(v *EnableDeviceResponseBody) *EnableDeviceResponse {
  s.Body = v
  return s
}

func (s *EnableDeviceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

