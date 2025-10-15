// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderUdPullResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableIdentityProviderUdPullResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableIdentityProviderUdPullResponse
  GetStatusCode() *int32 
  SetBody(v *EnableIdentityProviderUdPullResponseBody) *EnableIdentityProviderUdPullResponse
  GetBody() *EnableIdentityProviderUdPullResponseBody 
}

type EnableIdentityProviderUdPullResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableIdentityProviderUdPullResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableIdentityProviderUdPullResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderUdPullResponse) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderUdPullResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableIdentityProviderUdPullResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableIdentityProviderUdPullResponse) GetBody() *EnableIdentityProviderUdPullResponseBody  {
  return s.Body
}

func (s *EnableIdentityProviderUdPullResponse) SetHeaders(v map[string]*string) *EnableIdentityProviderUdPullResponse {
  s.Headers = v
  return s
}

func (s *EnableIdentityProviderUdPullResponse) SetStatusCode(v int32) *EnableIdentityProviderUdPullResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableIdentityProviderUdPullResponse) SetBody(v *EnableIdentityProviderUdPullResponseBody) *EnableIdentityProviderUdPullResponse {
  s.Body = v
  return s
}

func (s *EnableIdentityProviderUdPullResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

