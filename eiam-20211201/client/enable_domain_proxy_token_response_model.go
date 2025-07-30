// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDomainProxyTokenResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDomainProxyTokenResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDomainProxyTokenResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDomainProxyTokenResponseBody) *EnableDomainProxyTokenResponse
  GetBody() *EnableDomainProxyTokenResponseBody 
}

type EnableDomainProxyTokenResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDomainProxyTokenResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDomainProxyTokenResponse) GoString() string {
  return s.String()
}

func (s *EnableDomainProxyTokenResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDomainProxyTokenResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDomainProxyTokenResponse) GetBody() *EnableDomainProxyTokenResponseBody  {
  return s.Body
}

func (s *EnableDomainProxyTokenResponse) SetHeaders(v map[string]*string) *EnableDomainProxyTokenResponse {
  s.Headers = v
  return s
}

func (s *EnableDomainProxyTokenResponse) SetStatusCode(v int32) *EnableDomainProxyTokenResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDomainProxyTokenResponse) SetBody(v *EnableDomainProxyTokenResponseBody) *EnableDomainProxyTokenResponse {
  s.Body = v
  return s
}

func (s *EnableDomainProxyTokenResponse) Validate() error {
  return dara.Validate(s)
}

