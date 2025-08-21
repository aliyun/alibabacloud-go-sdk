// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessSendVerificationCodeResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EcologyOpennessSendVerificationCodeResponse
  GetStatusCode() *int32 
  SetBody(v *EcologyOpennessSendVerificationCodeResponseBody) *EcologyOpennessSendVerificationCodeResponse
  GetBody() *EcologyOpennessSendVerificationCodeResponseBody 
}

type EcologyOpennessSendVerificationCodeResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EcologyOpennessSendVerificationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeResponse) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponse) GoString() string {
  return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EcologyOpennessSendVerificationCodeResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EcologyOpennessSendVerificationCodeResponse) GetBody() *EcologyOpennessSendVerificationCodeResponseBody  {
  return s.Body
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetHeaders(v map[string]*string) *EcologyOpennessSendVerificationCodeResponse {
  s.Headers = v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetStatusCode(v int32) *EcologyOpennessSendVerificationCodeResponse {
  s.StatusCode = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponse) SetBody(v *EcologyOpennessSendVerificationCodeResponseBody) *EcologyOpennessSendVerificationCodeResponse {
  s.Body = v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponse) Validate() error {
  return dara.Validate(s)
}

