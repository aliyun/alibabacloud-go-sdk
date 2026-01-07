// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountBizRoleGrantResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountBizRoleGrantResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateAccountBizRoleGrantResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) *EnterpriseAccountUpdateAccountBizRoleGrantResponse
  GetBody() *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody 
}

type EnterpriseAccountUpdateAccountBizRoleGrantResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountBizRoleGrantResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) GetBody() *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) SetBody(v *EnterpriseAccountUpdateAccountBizRoleGrantResponseBody) *EnterpriseAccountUpdateAccountBizRoleGrantResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateAccountBizRoleGrantResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

