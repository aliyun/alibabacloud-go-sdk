// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceIpv6AddressResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *EnableInstanceIpv6AddressResponseBody
  GetData() *bool 
  SetRequestId(v string) *EnableInstanceIpv6AddressResponseBody
  GetRequestId() *string 
}

type EnableInstanceIpv6AddressResponseBody struct {
  // The result of the request.
  // 
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 7E9F7F16-05B5-42DA-94D6-E36402******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstanceIpv6AddressResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceIpv6AddressResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInstanceIpv6AddressResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EnableInstanceIpv6AddressResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInstanceIpv6AddressResponseBody) SetData(v bool) *EnableInstanceIpv6AddressResponseBody {
  s.Data = &v
  return s
}

func (s *EnableInstanceIpv6AddressResponseBody) SetRequestId(v string) *EnableInstanceIpv6AddressResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInstanceIpv6AddressResponseBody) Validate() error {
  return dara.Validate(s)
}

