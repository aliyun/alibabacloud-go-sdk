// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecIdentityProviderMetadataUrlResolutionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecIdentityProviderMetadataUrlResolutionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecIdentityProviderMetadataUrlResolutionResponse
  GetStatusCode() *int32 
  SetBody(v *ExecIdentityProviderMetadataUrlResolutionResponseBody) *ExecIdentityProviderMetadataUrlResolutionResponse
  GetBody() *ExecIdentityProviderMetadataUrlResolutionResponseBody 
}

type ExecIdentityProviderMetadataUrlResolutionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecIdentityProviderMetadataUrlResolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecIdentityProviderMetadataUrlResolutionResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecIdentityProviderMetadataUrlResolutionResponse) GoString() string {
  return s.String()
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) GetBody() *ExecIdentityProviderMetadataUrlResolutionResponseBody  {
  return s.Body
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) SetHeaders(v map[string]*string) *ExecIdentityProviderMetadataUrlResolutionResponse {
  s.Headers = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) SetStatusCode(v int32) *ExecIdentityProviderMetadataUrlResolutionResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) SetBody(v *ExecIdentityProviderMetadataUrlResolutionResponseBody) *ExecIdentityProviderMetadataUrlResolutionResponse {
  s.Body = v
  return s
}

func (s *ExecIdentityProviderMetadataUrlResolutionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

