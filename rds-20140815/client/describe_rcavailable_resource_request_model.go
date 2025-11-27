// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCores(v int32) *DescribeRCAvailableResourceRequest
	GetCores() *int32
	SetDataDiskCategory(v string) *DescribeRCAvailableResourceRequest
	GetDataDiskCategory() *string
	SetDedicatedHostId(v string) *DescribeRCAvailableResourceRequest
	GetDedicatedHostId() *string
	SetDestinationResource(v string) *DescribeRCAvailableResourceRequest
	GetDestinationResource() *string
	SetInstanceChargeType(v string) *DescribeRCAvailableResourceRequest
	GetInstanceChargeType() *string
	SetInstanceType(v string) *DescribeRCAvailableResourceRequest
	GetInstanceType() *string
	SetIoOptimized(v string) *DescribeRCAvailableResourceRequest
	GetIoOptimized() *string
	SetMemory(v float32) *DescribeRCAvailableResourceRequest
	GetMemory() *float32
	SetNetworkCategory(v string) *DescribeRCAvailableResourceRequest
	GetNetworkCategory() *string
	SetRegionId(v string) *DescribeRCAvailableResourceRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeRCAvailableResourceRequest
	GetResourceType() *string
	SetScope(v string) *DescribeRCAvailableResourceRequest
	GetScope() *string
	SetSpotDuration(v int32) *DescribeRCAvailableResourceRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *DescribeRCAvailableResourceRequest
	GetSpotStrategy() *string
	SetSystemDiskCategory(v string) *DescribeRCAvailableResourceRequest
	GetSystemDiskCategory() *string
	SetZoneId(v string) *DescribeRCAvailableResourceRequest
	GetZoneId() *string
}

type DescribeRCAvailableResourceRequest struct {
	Cores            *int32  `json:"Cores,omitempty" xml:"Cores,omitempty"`
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DedicatedHostId  *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// This parameter is required.
	DestinationResource *string  `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	InstanceChargeType  *string  `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType        *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IoOptimized         *string  `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Memory              *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkCategory     *string  `json:"NetworkCategory,omitempty" xml:"NetworkCategory,omitempty"`
	// This parameter is required.
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope              *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SpotDuration       *int32  `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotStrategy       *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeRCAvailableResourceRequest) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeRCAvailableResourceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeRCAvailableResourceRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeRCAvailableResourceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRCAvailableResourceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCAvailableResourceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeRCAvailableResourceRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeRCAvailableResourceRequest) GetNetworkCategory() *string {
	return s.NetworkCategory
}

func (s *DescribeRCAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCAvailableResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCAvailableResourceRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeRCAvailableResourceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeRCAvailableResourceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRCAvailableResourceRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeRCAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCAvailableResourceRequest) SetCores(v int32) *DescribeRCAvailableResourceRequest {
	s.Cores = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetDataDiskCategory(v string) *DescribeRCAvailableResourceRequest {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetDedicatedHostId(v string) *DescribeRCAvailableResourceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetDestinationResource(v string) *DescribeRCAvailableResourceRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetInstanceChargeType(v string) *DescribeRCAvailableResourceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetInstanceType(v string) *DescribeRCAvailableResourceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetIoOptimized(v string) *DescribeRCAvailableResourceRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetMemory(v float32) *DescribeRCAvailableResourceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetNetworkCategory(v string) *DescribeRCAvailableResourceRequest {
	s.NetworkCategory = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetRegionId(v string) *DescribeRCAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetResourceType(v string) *DescribeRCAvailableResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetScope(v string) *DescribeRCAvailableResourceRequest {
	s.Scope = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetSpotDuration(v int32) *DescribeRCAvailableResourceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetSpotStrategy(v string) *DescribeRCAvailableResourceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetSystemDiskCategory(v string) *DescribeRCAvailableResourceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) SetZoneId(v string) *DescribeRCAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
