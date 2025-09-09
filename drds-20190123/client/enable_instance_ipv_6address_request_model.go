// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceIpv6AddressRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDrdsInstanceId(v string) *EnableInstanceIpv6AddressRequest
  GetDrdsInstanceId() *string 
  SetRegionId(v string) *EnableInstanceIpv6AddressRequest
  GetRegionId() *string 
}

type EnableInstanceIpv6AddressRequest struct {
  // The ID of the PolarDB-X 1.0 instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // drds************
  DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
  // The ID of the region in which the instance resides.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstanceIpv6AddressRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceIpv6AddressRequest) GoString() string {
  return s.String()
}

func (s *EnableInstanceIpv6AddressRequest) GetDrdsInstanceId() *string  {
  return s.DrdsInstanceId
}

func (s *EnableInstanceIpv6AddressRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableInstanceIpv6AddressRequest) SetDrdsInstanceId(v string) *EnableInstanceIpv6AddressRequest {
  s.DrdsInstanceId = &v
  return s
}

func (s *EnableInstanceIpv6AddressRequest) SetRegionId(v string) *EnableInstanceIpv6AddressRequest {
  s.RegionId = &v
  return s
}

func (s *EnableInstanceIpv6AddressRequest) Validate() error {
  return dara.Validate(s)
}

