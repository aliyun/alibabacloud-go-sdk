// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceIpv6AddressResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableInstanceIpv6AddressResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableInstanceIpv6AddressResponse
  GetStatusCode() *int32 
  SetBody(v *EnableInstanceIpv6AddressResponseBody) *EnableInstanceIpv6AddressResponse
  GetBody() *EnableInstanceIpv6AddressResponseBody 
}

type EnableInstanceIpv6AddressResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableInstanceIpv6AddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInstanceIpv6AddressResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponse) GoString() string {
  return s.String()
}

func (s *EnableInstanceIpv6AddressResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableInstanceIpv6AddressResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableInstanceIpv6AddressResponse) GetBody() *EnableInstanceIpv6AddressResponseBody  {
  return s.Body
}

func (s *EnableInstanceIpv6AddressResponse) SetHeaders(v map[string]*string) *EnableInstanceIpv6AddressResponse {
  s.Headers = v
  return s
}

func (s *EnableInstanceIpv6AddressResponse) SetStatusCode(v int32) *EnableInstanceIpv6AddressResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableInstanceIpv6AddressResponse) SetBody(v *EnableInstanceIpv6AddressResponseBody) *EnableInstanceIpv6AddressResponse {
  s.Body = v
  return s
}

func (s *EnableInstanceIpv6AddressResponse) Validate() error {
  return dara.Validate(s)
}

