// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAuthnResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableIdentityProviderAuthnResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableIdentityProviderAuthnResponse
  GetStatusCode() *int32 
  SetBody(v *EnableIdentityProviderAuthnResponseBody) *EnableIdentityProviderAuthnResponse
  GetBody() *EnableIdentityProviderAuthnResponseBody 
}

type EnableIdentityProviderAuthnResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableIdentityProviderAuthnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableIdentityProviderAuthnResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAuthnResponse) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAuthnResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableIdentityProviderAuthnResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableIdentityProviderAuthnResponse) GetBody() *EnableIdentityProviderAuthnResponseBody  {
  return s.Body
}

func (s *EnableIdentityProviderAuthnResponse) SetHeaders(v map[string]*string) *EnableIdentityProviderAuthnResponse {
  s.Headers = v
  return s
}

func (s *EnableIdentityProviderAuthnResponse) SetStatusCode(v int32) *EnableIdentityProviderAuthnResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableIdentityProviderAuthnResponse) SetBody(v *EnableIdentityProviderAuthnResponseBody) *EnableIdentityProviderAuthnResponse {
  s.Body = v
  return s
}

func (s *EnableIdentityProviderAuthnResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

