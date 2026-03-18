// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCredentialProviderResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableCredentialProviderResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableCredentialProviderResponse
  GetStatusCode() *int32 
  SetBody(v *EnableCredentialProviderResponseBody) *EnableCredentialProviderResponse
  GetBody() *EnableCredentialProviderResponseBody 
}

type EnableCredentialProviderResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCredentialProviderResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableCredentialProviderResponse) GoString() string {
  return s.String()
}

func (s *EnableCredentialProviderResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableCredentialProviderResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableCredentialProviderResponse) GetBody() *EnableCredentialProviderResponseBody  {
  return s.Body
}

func (s *EnableCredentialProviderResponse) SetHeaders(v map[string]*string) *EnableCredentialProviderResponse {
  s.Headers = v
  return s
}

func (s *EnableCredentialProviderResponse) SetStatusCode(v int32) *EnableCredentialProviderResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableCredentialProviderResponse) SetBody(v *EnableCredentialProviderResponseBody) *EnableCredentialProviderResponse {
  s.Body = v
  return s
}

func (s *EnableCredentialProviderResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

