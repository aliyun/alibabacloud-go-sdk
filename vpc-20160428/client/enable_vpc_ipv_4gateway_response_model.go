// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcIpv4GatewayResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableVpcIpv4GatewayResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableVpcIpv4GatewayResponse
  GetStatusCode() *int32 
  SetBody(v *EnableVpcIpv4GatewayResponseBody) *EnableVpcIpv4GatewayResponse
  GetBody() *EnableVpcIpv4GatewayResponseBody 
}

type EnableVpcIpv4GatewayResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableVpcIpv4GatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableVpcIpv4GatewayResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcIpv4GatewayResponse) GoString() string {
  return s.String()
}

func (s *EnableVpcIpv4GatewayResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableVpcIpv4GatewayResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableVpcIpv4GatewayResponse) GetBody() *EnableVpcIpv4GatewayResponseBody  {
  return s.Body
}

func (s *EnableVpcIpv4GatewayResponse) SetHeaders(v map[string]*string) *EnableVpcIpv4GatewayResponse {
  s.Headers = v
  return s
}

func (s *EnableVpcIpv4GatewayResponse) SetStatusCode(v int32) *EnableVpcIpv4GatewayResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableVpcIpv4GatewayResponse) SetBody(v *EnableVpcIpv4GatewayResponseBody) *EnableVpcIpv4GatewayResponse {
  s.Body = v
  return s
}

func (s *EnableVpcIpv4GatewayResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

