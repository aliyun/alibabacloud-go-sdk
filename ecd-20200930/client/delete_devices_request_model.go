// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v int32) *DeleteDevicesRequest
	GetClientType() *int32
	SetDeviceIds(v []*string) *DeleteDevicesRequest
	GetDeviceIds() []*string
	SetForce(v int32) *DeleteDevicesRequest
	GetForce() *int32
	SetRegionId(v string) *DeleteDevicesRequest
	GetRegionId() *string
}

type DeleteDevicesRequest struct {
	// The type of the client.
	//
	// Valid values:
	//
	// 	- 1: hardware client.
	//
	// 	- 2: software client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The IDs of the devices. You can specify up to 200 IDs.
	//
	// This parameter is required.
	DeviceIds []*string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Repeated"`
	// Specifies whether to forcefully delete the device if the device is bound to a user.
	//
	// Valid values:
	//
	// 	- 0: do not forcefully delete the device.
	//
	// 	- 1: forcefully delete the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Force *int32 `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevicesRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *DeleteDevicesRequest) GetDeviceIds() []*string {
	return s.DeviceIds
}

func (s *DeleteDevicesRequest) GetForce() *int32 {
	return s.Force
}

func (s *DeleteDevicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDevicesRequest) SetClientType(v int32) *DeleteDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *DeleteDevicesRequest) SetDeviceIds(v []*string) *DeleteDevicesRequest {
	s.DeviceIds = v
	return s
}

func (s *DeleteDevicesRequest) SetForce(v int32) *DeleteDevicesRequest {
	s.Force = &v
	return s
}

func (s *DeleteDevicesRequest) SetRegionId(v string) *DeleteDevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDevicesRequest) Validate() error {
	return dara.Validate(s)
}
