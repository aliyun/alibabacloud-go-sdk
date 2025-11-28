// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceGroupRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDeviceGroupId(v string) *EnableDeviceGroupRequest
  GetDeviceGroupId() *string 
  SetRegionId(v string) *EnableDeviceGroupRequest
  GetRegionId() *string 
}

type EnableDeviceGroupRequest struct {
  // This parameter is required.
  DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableDeviceGroupRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceGroupRequest) GoString() string {
  return s.String()
}

func (s *EnableDeviceGroupRequest) GetDeviceGroupId() *string  {
  return s.DeviceGroupId
}

func (s *EnableDeviceGroupRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableDeviceGroupRequest) SetDeviceGroupId(v string) *EnableDeviceGroupRequest {
  s.DeviceGroupId = &v
  return s
}

func (s *EnableDeviceGroupRequest) SetRegionId(v string) *EnableDeviceGroupRequest {
  s.RegionId = &v
  return s
}

func (s *EnableDeviceGroupRequest) Validate() error {
  return dara.Validate(s)
}

