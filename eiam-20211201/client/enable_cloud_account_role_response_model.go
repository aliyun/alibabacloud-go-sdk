// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudAccountRoleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCloudAccountRoleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCloudAccountRoleResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCloudAccountRoleResponseBody) *EnableCloudAccountRoleResponse
  GetBody() *EnableCloudAccountRoleResponseBody 
}

type EnableCloudAccountRoleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCloudAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCloudAccountRoleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudAccountRoleResponse) GoString() string {
  return s.String()
}

func (s *EnableCloudAccountRoleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCloudAccountRoleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCloudAccountRoleResponse) GetBody() *EnableCloudAccountRoleResponseBody  {
  return s.Body
}

func (s *EnableCloudAccountRoleResponse) SetHeaders(v map[string]*string) *EnableCloudAccountRoleResponse {
  s.Headers = v
  return s
}

func (s *EnableCloudAccountRoleResponse) SetStatusCode(v int32) *EnableCloudAccountRoleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCloudAccountRoleResponse) SetBody(v *EnableCloudAccountRoleResponseBody) *EnableCloudAccountRoleResponse {
  s.Body = v
  return s
}

func (s *EnableCloudAccountRoleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

