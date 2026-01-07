// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleUpdateBizRoleResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseRoleUpdateBizRoleResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseRoleUpdateBizRoleResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseRoleUpdateBizRoleResponseBody) *EnterpriseRoleUpdateBizRoleResponse
  GetBody() *EnterpriseRoleUpdateBizRoleResponseBody 
}

type EnterpriseRoleUpdateBizRoleResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseRoleUpdateBizRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseRoleUpdateBizRoleResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseRoleUpdateBizRoleResponse) GetBody() *EnterpriseRoleUpdateBizRoleResponseBody  {
  return s.Body
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetHeaders(v map[string]*string) *EnterpriseRoleUpdateBizRoleResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetStatusCode(v int32) *EnterpriseRoleUpdateBizRoleResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponse) SetBody(v *EnterpriseRoleUpdateBizRoleResponseBody) *EnterpriseRoleUpdateBizRoleResponse {
  s.Body = v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

