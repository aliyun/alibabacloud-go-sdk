// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntVerifyResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EntVerifyResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EntVerifyResponse
  GetStatusCode() *int32 
  SetBody(v *EntVerifyResponseBody) *EntVerifyResponse
  GetBody() *EntVerifyResponseBody 
}

type EntVerifyResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EntVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntVerifyResponse) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyResponse) GoString() string {
  return s.String()
}

func (s *EntVerifyResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EntVerifyResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EntVerifyResponse) GetBody() *EntVerifyResponseBody  {
  return s.Body
}

func (s *EntVerifyResponse) SetHeaders(v map[string]*string) *EntVerifyResponse {
  s.Headers = v
  return s
}

func (s *EntVerifyResponse) SetStatusCode(v int32) *EntVerifyResponse {
  s.StatusCode = &v
  return s
}

func (s *EntVerifyResponse) SetBody(v *EntVerifyResponseBody) *EntVerifyResponse {
  s.Body = v
  return s
}

func (s *EntVerifyResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

