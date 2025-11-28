// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeviceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDeviceId(v string) *EnableDeviceRequest
  GetDeviceId() *string 
  SetRegionId(v string) *EnableDeviceRequest
  GetRegionId() *string 
}

type EnableDeviceRequest struct {
  // This parameter is required.
  DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableDeviceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDeviceRequest) GoString() string {
  return s.String()
}

func (s *EnableDeviceRequest) GetDeviceId() *string  {
  return s.DeviceId
}

func (s *EnableDeviceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableDeviceRequest) SetDeviceId(v string) *EnableDeviceRequest {
  s.DeviceId = &v
  return s
}

func (s *EnableDeviceRequest) SetRegionId(v string) *EnableDeviceRequest {
  s.RegionId = &v
  return s
}

func (s *EnableDeviceRequest) Validate() error {
  return dara.Validate(s)
}

