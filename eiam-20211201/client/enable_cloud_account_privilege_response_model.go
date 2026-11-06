// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudAccountPrivilegeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCloudAccountPrivilegeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCloudAccountPrivilegeResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCloudAccountPrivilegeResponseBody) *EnableCloudAccountPrivilegeResponse
  GetBody() *EnableCloudAccountPrivilegeResponseBody 
}

type EnableCloudAccountPrivilegeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCloudAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCloudAccountPrivilegeResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudAccountPrivilegeResponse) GoString() string {
  return s.String()
}

func (s *EnableCloudAccountPrivilegeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCloudAccountPrivilegeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCloudAccountPrivilegeResponse) GetBody() *EnableCloudAccountPrivilegeResponseBody  {
  return s.Body
}

func (s *EnableCloudAccountPrivilegeResponse) SetHeaders(v map[string]*string) *EnableCloudAccountPrivilegeResponse {
  s.Headers = v
  return s
}

func (s *EnableCloudAccountPrivilegeResponse) SetStatusCode(v int32) *EnableCloudAccountPrivilegeResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCloudAccountPrivilegeResponse) SetBody(v *EnableCloudAccountPrivilegeResponseBody) *EnableCloudAccountPrivilegeResponse {
  s.Body = v
  return s
}

func (s *EnableCloudAccountPrivilegeResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

