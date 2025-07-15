// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandDataVolumeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetOrderId(v string) *ExpandDataVolumeResponseBody
  GetOrderId() *string 
  SetRequestId(v string) *ExpandDataVolumeResponseBody
  GetRequestId() *string 
}

type ExpandDataVolumeResponseBody struct {
  // example:
  // 
  // 22326560487****
  OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // example:
  // 
  // 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExpandDataVolumeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExpandDataVolumeResponseBody) GoString() string {
  return s.String()
}

func (s *ExpandDataVolumeResponseBody) GetOrderId() *string  {
  return s.OrderId
}

func (s *ExpandDataVolumeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExpandDataVolumeResponseBody) SetOrderId(v string) *ExpandDataVolumeResponseBody {
  s.OrderId = &v
  return s
}

func (s *ExpandDataVolumeResponseBody) SetRequestId(v string) *ExpandDataVolumeResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExpandDataVolumeResponseBody) Validate() error {
  return dara.Validate(s)
}

