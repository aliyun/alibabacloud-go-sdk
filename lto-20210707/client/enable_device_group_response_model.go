// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceGroupResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDeviceGroupResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDeviceGroupResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDeviceGroupResponseBody) *EnableDeviceGroupResponse
  GetBody() *EnableDeviceGroupResponseBody 
}

type EnableDeviceGroupResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDeviceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDeviceGroupResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceGroupResponse) GoString() string {
  return s.String()
}

func (s *EnableDeviceGroupResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDeviceGroupResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDeviceGroupResponse) GetBody() *EnableDeviceGroupResponseBody  {
  return s.Body
}

func (s *EnableDeviceGroupResponse) SetHeaders(v map[string]*string) *EnableDeviceGroupResponse {
  s.Headers = v
  return s
}

func (s *EnableDeviceGroupResponse) SetStatusCode(v int32) *EnableDeviceGroupResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDeviceGroupResponse) SetBody(v *EnableDeviceGroupResponseBody) *EnableDeviceGroupResponse {
  s.Body = v
  return s
}

func (s *EnableDeviceGroupResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

