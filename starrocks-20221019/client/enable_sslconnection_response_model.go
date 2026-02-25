// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSSLConnectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSSLConnectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSSLConnectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSSLConnectionResponseBody) *EnableSSLConnectionResponse
  GetBody() *EnableSSLConnectionResponseBody 
}

type EnableSSLConnectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSSLConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSSLConnectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSSLConnectionResponse) GoString() string {
  return s.String()
}

func (s *EnableSSLConnectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSSLConnectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSSLConnectionResponse) GetBody() *EnableSSLConnectionResponseBody  {
  return s.Body
}

func (s *EnableSSLConnectionResponse) SetHeaders(v map[string]*string) *EnableSSLConnectionResponse {
  s.Headers = v
  return s
}

func (s *EnableSSLConnectionResponse) SetStatusCode(v int32) *EnableSSLConnectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSSLConnectionResponse) SetBody(v *EnableSSLConnectionResponseBody) *EnableSSLConnectionResponse {
  s.Body = v
  return s
}

func (s *EnableSSLConnectionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

