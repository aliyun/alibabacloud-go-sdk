// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountChangeSecurityEmailResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityEmailResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnterpriseAccountChangeSecurityEmailResponse
  GetStatusCode() *int32 
  SetBody(v *EnterpriseAccountChangeSecurityEmailResponseBody) *EnterpriseAccountChangeSecurityEmailResponse
  GetBody() *EnterpriseAccountChangeSecurityEmailResponseBody 
}

type EnterpriseAccountChangeSecurityEmailResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnterpriseAccountChangeSecurityEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnterpriseAccountChangeSecurityEmailResponse) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountChangeSecurityEmailResponse) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) GetBody() *EnterpriseAccountChangeSecurityEmailResponseBody  {
  return s.Body
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetHeaders(v map[string]*string) *EnterpriseAccountChangeSecurityEmailResponse {
  s.Headers = v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetStatusCode(v int32) *EnterpriseAccountChangeSecurityEmailResponse {
  s.StatusCode = &v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) SetBody(v *EnterpriseAccountChangeSecurityEmailResponseBody) *EnterpriseAccountChangeSecurityEmailResponse {
  s.Body = v
  return s
}

func (s *EnterpriseAccountChangeSecurityEmailResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

