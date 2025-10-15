// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationFederatedCredentialResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationFederatedCredentialResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationFederatedCredentialResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationFederatedCredentialResponseBody) *EnableApplicationFederatedCredentialResponse
  GetBody() *EnableApplicationFederatedCredentialResponseBody 
}

type EnableApplicationFederatedCredentialResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationFederatedCredentialResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationFederatedCredentialResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationFederatedCredentialResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationFederatedCredentialResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationFederatedCredentialResponse) GetBody() *EnableApplicationFederatedCredentialResponseBody  {
  return s.Body
}

func (s *EnableApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *EnableApplicationFederatedCredentialResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationFederatedCredentialResponse) SetStatusCode(v int32) *EnableApplicationFederatedCredentialResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationFederatedCredentialResponse) SetBody(v *EnableApplicationFederatedCredentialResponseBody) *EnableApplicationFederatedCredentialResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationFederatedCredentialResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

