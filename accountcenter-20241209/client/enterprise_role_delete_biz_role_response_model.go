// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleDeleteBizRoleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleDeleteBizRoleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleDeleteBizRoleResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleDeleteBizRoleResponseBody) *EnterpriseRoleDeleteBizRoleResponse
  GetBody() *EnterpriseRoleDeleteBizRoleResponseBody 
}

type EnterpriseRoleDeleteBizRoleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleDeleteBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleDeleteBizRoleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleDeleteBizRoleResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleDeleteBizRoleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleDeleteBizRoleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleDeleteBizRoleResponse) GetBody() *EnterpriseRoleDeleteBizRoleResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleDeleteBizRoleResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleDeleteBizRoleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponse) SetBody(v *EnterpriseRoleDeleteBizRoleResponseBody) *EnterpriseRoleDeleteBizRoleResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleDeleteBizRoleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

