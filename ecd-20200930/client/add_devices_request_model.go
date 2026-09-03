// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientType(v int32) *AddDevicesRequest
	GetClientType() *int32
	SetDeviceIds(v []*string) *AddDevicesRequest
	GetDeviceIds() []*string
	SetRegionId(v string) *AddDevicesRequest
	GetRegionId() *string
}

type AddDevicesRequest struct {
	// The device type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The list of device IDs. You can specify 1 to 200 device IDs.
	//
	// This parameter is required.
	DeviceIds []*string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddDevicesRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *AddDevicesRequest) GetDeviceIds() []*string {
	return s.DeviceIds
}

func (s *AddDevicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDevicesRequest) SetClientType(v int32) *AddDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *AddDevicesRequest) SetDeviceIds(v []*string) *AddDevicesRequest {
	s.DeviceIds = v
	return s
}

func (s *AddDevicesRequest) SetRegionId(v string) *AddDevicesRequest {
	s.RegionId = &v
	return s
}

func (s *AddDevicesRequest) Validate() error {
	return dara.Validate(s)
}
