// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableIdentityProviderAdvancedAbilityResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableIdentityProviderAdvancedAbilityResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableIdentityProviderAdvancedAbilityResponse
  GetStatusCode() *int32 
  SetBody(v *EnableIdentityProviderAdvancedAbilityResponseBody) *EnableIdentityProviderAdvancedAbilityResponse
  GetBody() *EnableIdentityProviderAdvancedAbilityResponseBody 
}

type EnableIdentityProviderAdvancedAbilityResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableIdentityProviderAdvancedAbilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableIdentityProviderAdvancedAbilityResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableIdentityProviderAdvancedAbilityResponse) GoString() string {
  return s.String()
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) GetBody() *EnableIdentityProviderAdvancedAbilityResponseBody  {
  return s.Body
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) SetHeaders(v map[string]*string) *EnableIdentityProviderAdvancedAbilityResponse {
  s.Headers = v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) SetStatusCode(v int32) *EnableIdentityProviderAdvancedAbilityResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) SetBody(v *EnableIdentityProviderAdvancedAbilityResponseBody) *EnableIdentityProviderAdvancedAbilityResponse {
  s.Body = v
  return s
}

func (s *EnableIdentityProviderAdvancedAbilityResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

