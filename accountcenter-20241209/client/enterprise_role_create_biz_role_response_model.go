// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleCreateBizRoleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleCreateBizRoleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleCreateBizRoleResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleCreateBizRoleResponseBody) *EnterpriseRoleCreateBizRoleResponse
  GetBody() *EnterpriseRoleCreateBizRoleResponseBody 
}

type EnterpriseRoleCreateBizRoleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleCreateBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleCreateBizRoleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleCreateBizRoleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleCreateBizRoleResponse) GetBody() *EnterpriseRoleCreateBizRoleResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleCreateBizRoleResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleCreateBizRoleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponse) SetBody(v *EnterpriseRoleCreateBizRoleResponseBody) *EnterpriseRoleCreateBizRoleResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleCreateBizRoleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

