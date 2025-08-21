// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessAuthenticateResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EcologyOpennessAuthenticateResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EcologyOpennessAuthenticateResponse
  GetStatusCode() *int32 
  SetBody(v *EcologyOpennessAuthenticateResponseBody) *EcologyOpennessAuthenticateResponse
  GetBody() *EcologyOpennessAuthenticateResponseBody 
}

type EcologyOpennessAuthenticateResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EcologyOpennessAuthenticateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EcologyOpennessAuthenticateResponse) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessAuthenticateResponse) GoString() string {
  return s.String()
}

func (s *EcologyOpennessAuthenticateResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EcologyOpennessAuthenticateResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EcologyOpennessAuthenticateResponse) GetBody() *EcologyOpennessAuthenticateResponseBody  {
  return s.Body
}

func (s *EcologyOpennessAuthenticateResponse) SetHeaders(v map[string]*string) *EcologyOpennessAuthenticateResponse {
  s.Headers = v
  return s
}

func (s *EcologyOpennessAuthenticateResponse) SetStatusCode(v int32) *EcologyOpennessAuthenticateResponse {
  s.StatusCode = &v
  return s
}

func (s *EcologyOpennessAuthenticateResponse) SetBody(v *EcologyOpennessAuthenticateResponseBody) *EcologyOpennessAuthenticateResponse {
  s.Body = v
  return s
}

func (s *EcologyOpennessAuthenticateResponse) Validate() error {
  return dara.Validate(s)
}

