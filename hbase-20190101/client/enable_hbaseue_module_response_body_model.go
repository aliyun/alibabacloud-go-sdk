// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueModuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *EnableHBaseueModuleResponseBody
  GetClusterId() *string 
  SetOrderId(v string) *EnableHBaseueModuleResponseBody
  GetOrderId() *string 
  SetRequestId(v string) *EnableHBaseueModuleResponseBody
  GetRequestId() *string 
}

type EnableHBaseueModuleResponseBody struct {
  // example:
  // 
  // ld-bp150tns0sjxs****-m1-ps
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // example:
  // 
  // 21474915573****
  OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // example:
  // 
  // 407075EA-47F5-5A2D-888F-C1F90B8F3FCA
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableHBaseueModuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueModuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHBaseueModuleResponseBody) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnableHBaseueModuleResponseBody) GetOrderId() *string  {
  return s.OrderId
}

func (s *EnableHBaseueModuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHBaseueModuleResponseBody) SetClusterId(v string) *EnableHBaseueModuleResponseBody {
  s.ClusterId = &v
  return s
}

func (s *EnableHBaseueModuleResponseBody) SetOrderId(v string) *EnableHBaseueModuleResponseBody {
  s.OrderId = &v
  return s
}

func (s *EnableHBaseueModuleResponseBody) SetRequestId(v string) *EnableHBaseueModuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHBaseueModuleResponseBody) Validate() error {
  return dara.Validate(s)
}

