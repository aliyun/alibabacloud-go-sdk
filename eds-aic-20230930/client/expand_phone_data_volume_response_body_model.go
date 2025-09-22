// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandPhoneDataVolumeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetOrderId(v string) *ExpandPhoneDataVolumeResponseBody
  GetOrderId() *string 
  SetRequestId(v string) *ExpandPhoneDataVolumeResponseBody
  GetRequestId() *string 
}

type ExpandPhoneDataVolumeResponseBody struct {
  // example:
  // 
  // 223684716098****
  OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // example:
  // 
  // DB070C80-45AC-52CA-8101-937C25DA****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExpandPhoneDataVolumeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExpandPhoneDataVolumeResponseBody) GoString() string {
  return s.String()
}

func (s *ExpandPhoneDataVolumeResponseBody) GetOrderId() *string  {
  return s.OrderId
}

func (s *ExpandPhoneDataVolumeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExpandPhoneDataVolumeResponseBody) SetOrderId(v string) *ExpandPhoneDataVolumeResponseBody {
  s.OrderId = &v
  return s
}

func (s *ExpandPhoneDataVolumeResponseBody) SetRequestId(v string) *ExpandPhoneDataVolumeResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExpandPhoneDataVolumeResponseBody) Validate() error {
  return dara.Validate(s)
}

