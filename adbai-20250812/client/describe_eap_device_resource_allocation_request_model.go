// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEapDeviceResourceAllocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEapDeviceResourceAllocationRequest
	GetDBClusterId() *string
	SetDeviceCount(v int32) *DescribeEapDeviceResourceAllocationRequest
	GetDeviceCount() *int32
	SetRegionId(v string) *DescribeEapDeviceResourceAllocationRequest
	GetRegionId() *string
}

type DescribeEapDeviceResourceAllocationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 3
	DeviceCount *int32 `json:"DeviceCount,omitempty" xml:"DeviceCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEapDeviceResourceAllocationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEapDeviceResourceAllocationRequest) GoString() string {
	return s.String()
}

func (s *DescribeEapDeviceResourceAllocationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEapDeviceResourceAllocationRequest) GetDeviceCount() *int32 {
	return s.DeviceCount
}

func (s *DescribeEapDeviceResourceAllocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEapDeviceResourceAllocationRequest) SetDBClusterId(v string) *DescribeEapDeviceResourceAllocationRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationRequest) SetDeviceCount(v int32) *DescribeEapDeviceResourceAllocationRequest {
	s.DeviceCount = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationRequest) SetRegionId(v string) *DescribeEapDeviceResourceAllocationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEapDeviceResourceAllocationRequest) Validate() error {
	return dara.Validate(s)
}
