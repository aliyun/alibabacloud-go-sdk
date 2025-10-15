// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationTokenResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableApplicationTokenResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableApplicationTokenResponse
  GetStatusCode() *int32 
  SetBody(v *EnableApplicationTokenResponseBody) *EnableApplicationTokenResponse
  GetBody() *EnableApplicationTokenResponseBody 
}

type EnableApplicationTokenResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableApplicationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationTokenResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationTokenResponse) GoString() string {
  return s.String()
}

func (s *EnableApplicationTokenResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableApplicationTokenResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableApplicationTokenResponse) GetBody() *EnableApplicationTokenResponseBody  {
  return s.Body
}

func (s *EnableApplicationTokenResponse) SetHeaders(v map[string]*string) *EnableApplicationTokenResponse {
  s.Headers = v
  return s
}

func (s *EnableApplicationTokenResponse) SetStatusCode(v int32) *EnableApplicationTokenResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableApplicationTokenResponse) SetBody(v *EnableApplicationTokenResponseBody) *EnableApplicationTokenResponse {
  s.Body = v
  return s
}

func (s *EnableApplicationTokenResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

