// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSSLResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSSLResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSSLResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSSLResponseBody) *EnableSSLResponse
  GetBody() *EnableSSLResponseBody 
}

type EnableSSLResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSSLResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSSLResponse) GoString() string {
  return s.String()
}

func (s *EnableSSLResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSSLResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSSLResponse) GetBody() *EnableSSLResponseBody  {
  return s.Body
}

func (s *EnableSSLResponse) SetHeaders(v map[string]*string) *EnableSSLResponse {
  s.Headers = v
  return s
}

func (s *EnableSSLResponse) SetStatusCode(v int32) *EnableSSLResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSSLResponse) SetBody(v *EnableSSLResponseBody) *EnableSSLResponse {
  s.Body = v
  return s
}

func (s *EnableSSLResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

