// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityMobileResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityMobileResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountChangeSecurityMobileResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountChangeSecurityMobileResponseBody) *EnterpriseAccountChangeSecurityMobileResponse
  GetBody() *EnterpriseAccountChangeSecurityMobileResponseBody 
}

type EnterpriseAccountChangeSecurityMobileResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountChangeSecurityMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeSecurityMobileResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityMobileResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) GetBody() *EnterpriseAccountChangeSecurityMobileResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityMobileResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetStatusCode(v int32) *EnterpriseAccountChangeSecurityMobileResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) SetBody(v *EnterpriseAccountChangeSecurityMobileResponseBody) *EnterpriseAccountChangeSecurityMobileResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountChangeSecurityMobileResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

