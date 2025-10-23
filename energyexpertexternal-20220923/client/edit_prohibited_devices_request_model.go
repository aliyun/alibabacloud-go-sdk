// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditProhibitedDevicesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFactoryId(v string) *EditProhibitedDevicesRequest
  GetFactoryId() *string 
  SetHvacDeviceConfigVOList(v []*EditProhibitedDevicesRequestHvacDeviceConfigVOList) *EditProhibitedDevicesRequest
  GetHvacDeviceConfigVOList() []*EditProhibitedDevicesRequestHvacDeviceConfigVOList 
  SetSystemId(v string) *EditProhibitedDevicesRequest
  GetSystemId() *string 
}

type EditProhibitedDevicesRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // ***
  FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
  // This parameter is required.
  HvacDeviceConfigVOList []*EditProhibitedDevicesRequestHvacDeviceConfigVOList `json:"hvacDeviceConfigVOList,omitempty" xml:"hvacDeviceConfigVOList,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // system1
  SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
}

func (s EditProhibitedDevicesRequest) String() string {
  return dara.Prettify(s)
}

func (s EditProhibitedDevicesRequest) GoString() string {
  return s.String()
}

func (s *EditProhibitedDevicesRequest) GetFactoryId() *string  {
  return s.FactoryId
}

func (s *EditProhibitedDevicesRequest) GetHvacDeviceConfigVOList() []*EditProhibitedDevicesRequestHvacDeviceConfigVOList  {
  return s.HvacDeviceConfigVOList
}

func (s *EditProhibitedDevicesRequest) GetSystemId() *string  {
  return s.SystemId
}

func (s *EditProhibitedDevicesRequest) SetFactoryId(v string) *EditProhibitedDevicesRequest {
  s.FactoryId = &v
  return s
}

func (s *EditProhibitedDevicesRequest) SetHvacDeviceConfigVOList(v []*EditProhibitedDevicesRequestHvacDeviceConfigVOList) *EditProhibitedDevicesRequest {
  s.HvacDeviceConfigVOList = v
  return s
}

func (s *EditProhibitedDevicesRequest) SetSystemId(v string) *EditProhibitedDevicesRequest {
  s.SystemId = &v
  return s
}

func (s *EditProhibitedDevicesRequest) Validate() error {
  if s.HvacDeviceConfigVOList != nil {
    for _, item := range s.HvacDeviceConfigVOList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EditProhibitedDevicesRequestHvacDeviceConfigVOList struct {
  // example:
  // 
  // build_01
  BuildingId *string `json:"buildingId,omitempty" xml:"buildingId,omitempty"`
  // example:
  // 
  // id1
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
  // example:
  // 
  // name1
  DeviceName *string `json:"deviceName,omitempty" xml:"deviceName,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
  // example:
  // 
  // fence_01
  FenceId *string `json:"fenceId,omitempty" xml:"fenceId,omitempty"`
  // example:
  // 
  // floor_01
  FloorId *string `json:"floorId,omitempty" xml:"floorId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  IsForbidden *int32 `json:"isForbidden,omitempty" xml:"isForbidden,omitempty"`
  // example:
  // 
  // 1
  IsUnfavorableArea *int32 `json:"isUnfavorableArea,omitempty" xml:"isUnfavorableArea,omitempty"`
}

func (s EditProhibitedDevicesRequestHvacDeviceConfigVOList) String() string {
  return dara.Prettify(s)
}

func (s EditProhibitedDevicesRequestHvacDeviceConfigVOList) GoString() string {
  return s.String()
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetBuildingId() *string  {
  return s.BuildingId
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetDeviceId() *string  {
  return s.DeviceId
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetDeviceName() *string  {
  return s.DeviceName
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetDeviceType() *string  {
  return s.DeviceType
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetFenceId() *string  {
  return s.FenceId
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetFloorId() *string  {
  return s.FloorId
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetIsForbidden() *int32  {
  return s.IsForbidden
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) GetIsUnfavorableArea() *int32  {
  return s.IsUnfavorableArea
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetBuildingId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.BuildingId = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.DeviceId = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceName(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.DeviceName = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetDeviceType(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.DeviceType = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetFenceId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.FenceId = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetFloorId(v string) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.FloorId = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetIsForbidden(v int32) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.IsForbidden = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) SetIsUnfavorableArea(v int32) *EditProhibitedDevicesRequestHvacDeviceConfigVOList {
  s.IsUnfavorableArea = &v
  return s
}

func (s *EditProhibitedDevicesRequestHvacDeviceConfigVOList) Validate() error {
  return dara.Validate(s)
}

