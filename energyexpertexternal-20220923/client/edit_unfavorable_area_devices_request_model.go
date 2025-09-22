// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditUnfavorableAreaDevicesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFactoryId(v string) *EditUnfavorableAreaDevicesRequest
  GetFactoryId() *string 
  SetHvacDeviceConfigVOList(v []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) *EditUnfavorableAreaDevicesRequest
  GetHvacDeviceConfigVOList() []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList 
  SetSystemId(v string) *EditUnfavorableAreaDevicesRequest
  GetSystemId() *string 
}

type EditUnfavorableAreaDevicesRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // ***
  FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
  // This parameter is required.
  HvacDeviceConfigVOList []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList `json:"hvacDeviceConfigVOList,omitempty" xml:"hvacDeviceConfigVOList,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // system1
  SystemId *string `json:"systemId,omitempty" xml:"systemId,omitempty"`
}

func (s EditUnfavorableAreaDevicesRequest) String() string {
  return dara.Prettify(s)
}

func (s EditUnfavorableAreaDevicesRequest) GoString() string {
  return s.String()
}

func (s *EditUnfavorableAreaDevicesRequest) GetFactoryId() *string  {
  return s.FactoryId
}

func (s *EditUnfavorableAreaDevicesRequest) GetHvacDeviceConfigVOList() []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList  {
  return s.HvacDeviceConfigVOList
}

func (s *EditUnfavorableAreaDevicesRequest) GetSystemId() *string  {
  return s.SystemId
}

func (s *EditUnfavorableAreaDevicesRequest) SetFactoryId(v string) *EditUnfavorableAreaDevicesRequest {
  s.FactoryId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequest) SetHvacDeviceConfigVOList(v []*EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) *EditUnfavorableAreaDevicesRequest {
  s.HvacDeviceConfigVOList = v
  return s
}

func (s *EditUnfavorableAreaDevicesRequest) SetSystemId(v string) *EditUnfavorableAreaDevicesRequest {
  s.SystemId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequest) Validate() error {
  return dara.Validate(s)
}

type EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList struct {
  // example:
  // 
  // buildingId1
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
  // fenceId1
  FenceId *string `json:"fenceId,omitempty" xml:"fenceId,omitempty"`
  // example:
  // 
  // floorId2
  FloorId *string `json:"floorId,omitempty" xml:"floorId,omitempty"`
  // example:
  // 
  // 1
  IsForbidden *int32 `json:"isForbidden,omitempty" xml:"isForbidden,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  IsUnfavorableArea *int32 `json:"isUnfavorableArea,omitempty" xml:"isUnfavorableArea,omitempty"`
}

func (s EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) String() string {
  return dara.Prettify(s)
}

func (s EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GoString() string {
  return s.String()
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetBuildingId() *string  {
  return s.BuildingId
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetDeviceId() *string  {
  return s.DeviceId
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetDeviceName() *string  {
  return s.DeviceName
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetDeviceType() *string  {
  return s.DeviceType
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetFenceId() *string  {
  return s.FenceId
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetFloorId() *string  {
  return s.FloorId
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetIsForbidden() *int32  {
  return s.IsForbidden
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) GetIsUnfavorableArea() *int32  {
  return s.IsUnfavorableArea
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetBuildingId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.BuildingId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.DeviceId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceName(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.DeviceName = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetDeviceType(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.DeviceType = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetFenceId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.FenceId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetFloorId(v string) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.FloorId = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetIsForbidden(v int32) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.IsForbidden = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) SetIsUnfavorableArea(v int32) *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList {
  s.IsUnfavorableArea = &v
  return s
}

func (s *EditUnfavorableAreaDevicesRequestHvacDeviceConfigVOList) Validate() error {
  return dara.Validate(s)
}

