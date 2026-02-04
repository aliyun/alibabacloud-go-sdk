// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpireLoginTokenResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExpireLoginTokenResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExpireLoginTokenResponse
  GetStatusCode() *int32 
  SetBody(v *ExpireLoginTokenResponseBody) *ExpireLoginTokenResponse
  GetBody() *ExpireLoginTokenResponseBody 
}

type ExpireLoginTokenResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExpireLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExpireLoginTokenResponse) String() string {
  return dara.Prettify(s)
}

func (s ExpireLoginTokenResponse) GoString() string {
  return s.String()
}

func (s *ExpireLoginTokenResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExpireLoginTokenResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExpireLoginTokenResponse) GetBody() *ExpireLoginTokenResponseBody  {
  return s.Body
}

func (s *ExpireLoginTokenResponse) SetHeaders(v map[string]*string) *ExpireLoginTokenResponse {
  s.Headers = v
  return s
}

func (s *ExpireLoginTokenResponse) SetStatusCode(v int32) *ExpireLoginTokenResponse {
  s.StatusCode = &v
  return s
}

func (s *ExpireLoginTokenResponse) SetBody(v *ExpireLoginTokenResponseBody) *ExpireLoginTokenResponse {
  s.Body = v
  return s
}

func (s *ExpireLoginTokenResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

