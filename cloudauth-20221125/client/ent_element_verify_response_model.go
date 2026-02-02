// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntElementVerifyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntElementVerifyResponse
  GetStatusCode() *int32 
  SetBody(v *EntElementVerifyResponseBody) *EntElementVerifyResponse
  GetBody() *EntElementVerifyResponseBody 
}

type EntElementVerifyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntElementVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntElementVerifyResponse) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyResponse) GoString() string {
  return s.String()
}

func (s *EntElementVerifyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntElementVerifyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntElementVerifyResponse) GetBody() *EntElementVerifyResponseBody  {
  return s.Body
}

func (s *EntElementVerifyResponse) SetHeaders(v map[string]*string) *EntElementVerifyResponse {
  s.Headers = v
  return s
}

func (s *EntElementVerifyResponse) SetStatusCode(v int32) *EntElementVerifyResponse {
  s.StatusCode = &v
  return s
}

func (s *EntElementVerifyResponse) SetBody(v *EntElementVerifyResponseBody) *EntElementVerifyResponse {
  s.Body = v
  return s
}

func (s *EntElementVerifyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

