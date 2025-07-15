// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableVpcIpv4GatewayResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableVpcIpv4GatewayResponseBody
  GetRequestId() *string 
}

type EnableVpcIpv4GatewayResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 54B48E3D-DF70-471B-AA93-08E683A1B45
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVpcIpv4GatewayResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableVpcIpv4GatewayResponseBody) GoString() string {
  return s.String()
}

func (s *EnableVpcIpv4GatewayResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableVpcIpv4GatewayResponseBody) SetRequestId(v string) *EnableVpcIpv4GatewayResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableVpcIpv4GatewayResponseBody) Validate() error {
  return dara.Validate(s)
}

