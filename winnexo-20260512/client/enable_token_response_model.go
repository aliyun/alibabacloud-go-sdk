// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTokenResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableTokenResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableTokenResponse
  GetStatusCode() *int32 
  SetBody(v *EnableTokenResponseBody) *EnableTokenResponse
  GetBody() *EnableTokenResponseBody 
}

type EnableTokenResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableTokenResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableTokenResponse) GoString() string {
  return s.String()
}

func (s *EnableTokenResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableTokenResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableTokenResponse) GetBody() *EnableTokenResponseBody  {
  return s.Body
}

func (s *EnableTokenResponse) SetHeaders(v map[string]*string) *EnableTokenResponse {
  s.Headers = v
  return s
}

func (s *EnableTokenResponse) SetStatusCode(v int32) *EnableTokenResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableTokenResponse) SetBody(v *EnableTokenResponseBody) *EnableTokenResponse {
  s.Body = v
  return s
}

func (s *EnableTokenResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

