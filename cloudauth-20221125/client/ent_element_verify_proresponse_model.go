// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyPROResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntElementVerifyPROResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntElementVerifyPROResponse
  GetStatusCode() *int32 
  SetBody(v *EntElementVerifyPROResponseBody) *EntElementVerifyPROResponse
  GetBody() *EntElementVerifyPROResponseBody 
}

type EntElementVerifyPROResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntElementVerifyPROResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntElementVerifyPROResponse) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyPROResponse) GoString() string {
  return s.String()
}

func (s *EntElementVerifyPROResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntElementVerifyPROResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntElementVerifyPROResponse) GetBody() *EntElementVerifyPROResponseBody  {
  return s.Body
}

func (s *EntElementVerifyPROResponse) SetHeaders(v map[string]*string) *EntElementVerifyPROResponse {
  s.Headers = v
  return s
}

func (s *EntElementVerifyPROResponse) SetStatusCode(v int32) *EntElementVerifyPROResponse {
  s.StatusCode = &v
  return s
}

func (s *EntElementVerifyPROResponse) SetBody(v *EntElementVerifyPROResponseBody) *EntElementVerifyPROResponse {
  s.Body = v
  return s
}

func (s *EntElementVerifyPROResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

