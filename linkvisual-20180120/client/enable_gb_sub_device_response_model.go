// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableGbSubDeviceResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableGbSubDeviceResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableGbSubDeviceResponse
  GetStatusCode() *int32 
  SetBody(v *EnableGbSubDeviceResponseBody) *EnableGbSubDeviceResponse
  GetBody() *EnableGbSubDeviceResponseBody 
}

type EnableGbSubDeviceResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableGbSubDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableGbSubDeviceResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableGbSubDeviceResponse) GoString() string {
  return s.String()
}

func (s *EnableGbSubDeviceResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableGbSubDeviceResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableGbSubDeviceResponse) GetBody() *EnableGbSubDeviceResponseBody  {
  return s.Body
}

func (s *EnableGbSubDeviceResponse) SetHeaders(v map[string]*string) *EnableGbSubDeviceResponse {
  s.Headers = v
  return s
}

func (s *EnableGbSubDeviceResponse) SetStatusCode(v int32) *EnableGbSubDeviceResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableGbSubDeviceResponse) SetBody(v *EnableGbSubDeviceResponseBody) *EnableGbSubDeviceResponse {
  s.Body = v
  return s
}

func (s *EnableGbSubDeviceResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

