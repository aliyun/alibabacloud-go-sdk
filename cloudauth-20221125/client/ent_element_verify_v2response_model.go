// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyV2Response interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntElementVerifyV2Response
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntElementVerifyV2Response
  GetStatusCode() *int32 
  SetBody(v *EntElementVerifyV2ResponseBody) *EntElementVerifyV2Response
  GetBody() *EntElementVerifyV2ResponseBody 
}

type EntElementVerifyV2Response struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntElementVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntElementVerifyV2Response) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyV2Response) GoString() string {
  return s.String()
}

func (s *EntElementVerifyV2Response) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntElementVerifyV2Response) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntElementVerifyV2Response) GetBody() *EntElementVerifyV2ResponseBody  {
  return s.Body
}

func (s *EntElementVerifyV2Response) SetHeaders(v map[string]*string) *EntElementVerifyV2Response {
  s.Headers = v
  return s
}

func (s *EntElementVerifyV2Response) SetStatusCode(v int32) *EntElementVerifyV2Response {
  s.StatusCode = &v
  return s
}

func (s *EntElementVerifyV2Response) SetBody(v *EntElementVerifyV2ResponseBody) *EntElementVerifyV2Response {
  s.Body = v
  return s
}

func (s *EntElementVerifyV2Response) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

