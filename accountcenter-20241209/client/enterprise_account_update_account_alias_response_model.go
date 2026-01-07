// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateAccountAliasResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountAliasResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateAccountAliasResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateAccountAliasResponseBody) *EnterpriseAccountUpdateAccountAliasResponse
  GetBody() *EnterpriseAccountUpdateAccountAliasResponseBody 
}

type EnterpriseAccountUpdateAccountAliasResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateAccountAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateAccountAliasResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateAccountAliasResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) GetBody() *EnterpriseAccountUpdateAccountAliasResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateAccountAliasResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateAccountAliasResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) SetBody(v *EnterpriseAccountUpdateAccountAliasResponseBody) *EnterpriseAccountUpdateAccountAliasResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateAccountAliasResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

