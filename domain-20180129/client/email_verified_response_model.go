// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmailVerifiedResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EmailVerifiedResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EmailVerifiedResponse
  GetStatusCode() *int32 
  SetBody(v *EmailVerifiedResponseBody) *EmailVerifiedResponse
  GetBody() *EmailVerifiedResponseBody 
}

type EmailVerifiedResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EmailVerifiedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EmailVerifiedResponse) String() string {
  return dara.Prettify(s)
}

func (s EmailVerifiedResponse) GoString() string {
  return s.String()
}

func (s *EmailVerifiedResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EmailVerifiedResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EmailVerifiedResponse) GetBody() *EmailVerifiedResponseBody  {
  return s.Body
}

func (s *EmailVerifiedResponse) SetHeaders(v map[string]*string) *EmailVerifiedResponse {
  s.Headers = v
  return s
}

func (s *EmailVerifiedResponse) SetStatusCode(v int32) *EmailVerifiedResponse {
  s.StatusCode = &v
  return s
}

func (s *EmailVerifiedResponse) SetBody(v *EmailVerifiedResponseBody) *EmailVerifiedResponse {
  s.Body = v
  return s
}

func (s *EmailVerifiedResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

