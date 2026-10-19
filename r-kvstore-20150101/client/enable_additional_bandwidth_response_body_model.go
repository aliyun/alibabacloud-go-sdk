// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdditionalBandwidthResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetOrderId(v string) *EnableAdditionalBandwidthResponseBody
  GetOrderId() *string 
  SetRequestId(v string) *EnableAdditionalBandwidthResponseBody
  GetRequestId() *string 
}

type EnableAdditionalBandwidthResponseBody struct {
  // The ID of the order.
  // 
  // example:
  // 
  // 2084452111111
  OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // D622714-AEDD-4609-9167-F5DDD3D1****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAdditionalBandwidthResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAdditionalBandwidthResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAdditionalBandwidthResponseBody) GetOrderId() *string  {
  return s.OrderId
}

func (s *EnableAdditionalBandwidthResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAdditionalBandwidthResponseBody) SetOrderId(v string) *EnableAdditionalBandwidthResponseBody {
  s.OrderId = &v
  return s
}

func (s *EnableAdditionalBandwidthResponseBody) SetRequestId(v string) *EnableAdditionalBandwidthResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAdditionalBandwidthResponseBody) Validate() error {
  return dara.Validate(s)
}

