// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableFederatedCredentialProviderResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableFederatedCredentialProviderResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableFederatedCredentialProviderResponse
  GetStatusCode() *int32 
  SetBody(v *EnableFederatedCredentialProviderResponseBody) *EnableFederatedCredentialProviderResponse
  GetBody() *EnableFederatedCredentialProviderResponseBody 
}

type EnableFederatedCredentialProviderResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFederatedCredentialProviderResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableFederatedCredentialProviderResponse) GoString() string {
  return s.String()
}

func (s *EnableFederatedCredentialProviderResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableFederatedCredentialProviderResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableFederatedCredentialProviderResponse) GetBody() *EnableFederatedCredentialProviderResponseBody  {
  return s.Body
}

func (s *EnableFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *EnableFederatedCredentialProviderResponse {
  s.Headers = v
  return s
}

func (s *EnableFederatedCredentialProviderResponse) SetStatusCode(v int32) *EnableFederatedCredentialProviderResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableFederatedCredentialProviderResponse) SetBody(v *EnableFederatedCredentialProviderResponseBody) *EnableFederatedCredentialProviderResponse {
  s.Body = v
  return s
}

func (s *EnableFederatedCredentialProviderResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

