// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNatGatewayEcsMetricResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableNatGatewayEcsMetricResponseBody
  GetRequestId() *string 
}

type EnableNatGatewayEcsMetricResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableNatGatewayEcsMetricResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableNatGatewayEcsMetricResponseBody) GoString() string {
  return s.String()
}

func (s *EnableNatGatewayEcsMetricResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableNatGatewayEcsMetricResponseBody) SetRequestId(v string) *EnableNatGatewayEcsMetricResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableNatGatewayEcsMetricResponseBody) Validate() error {
  return dara.Validate(s)
}

