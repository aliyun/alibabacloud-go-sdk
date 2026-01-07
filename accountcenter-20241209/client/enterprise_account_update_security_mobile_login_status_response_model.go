// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountUpdateSecurityMobileLoginStatusResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
  GetBody() *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody 
}

type EnterpriseAccountUpdateSecurityMobileLoginStatusResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) GetBody() *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetHeaders(v map[string]*string) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetStatusCode(v int32) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) SetBody(v *EnterpriseAccountUpdateSecurityMobileLoginStatusResponseBody) *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

