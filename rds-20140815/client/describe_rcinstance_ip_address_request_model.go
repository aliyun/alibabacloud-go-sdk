// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRCInstanceIpAddressRequest
	GetCurrentPage() *int32
	SetDdosRegionId(v string) *DescribeRCInstanceIpAddressRequest
	GetDdosRegionId() *string
	SetDdosStatus(v string) *DescribeRCInstanceIpAddressRequest
	GetDdosStatus() *string
	SetInstanceId(v string) *DescribeRCInstanceIpAddressRequest
	GetInstanceId() *string
	SetInstanceIp(v string) *DescribeRCInstanceIpAddressRequest
	GetInstanceIp() *string
	SetInstanceName(v string) *DescribeRCInstanceIpAddressRequest
	GetInstanceName() *string
	SetInstanceType(v string) *DescribeRCInstanceIpAddressRequest
	GetInstanceType() *string
	SetPageSize(v int32) *DescribeRCInstanceIpAddressRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCInstanceIpAddressRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeRCInstanceIpAddressRequest
	GetResourceType() *string
}

type DescribeRCInstanceIpAddressRequest struct {
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-beijing
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The DDoS mitigation status of the asset. Valid values:
	//
	// 	- **defense**: queries assets for which traffic scrubbing is performed.
	//
	// 	- **blackhole**: queries assets for which blackhole filtering is triggered.
	//
	// example:
	//
	// defense
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The ID of the RDS Custom instance to which the asset to query is added.
	//
	// example:
	//
	// rc-y6dn4pyuub1r89******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the asset to query.
	//
	// example:
	//
	// 39.105.XXX.XXX
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// The name of the RDS Custom instance to which the asset to query is added.
	//
	// example:
	//
	// rc-y6dn4pyuub1r89******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the asset that is assigned a public IP address. Set the value to **ecs**.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of instances on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the RDS Custom instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Set the value to **ecs**.
	//
	// example:
	//
	// ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRCInstanceIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceIpAddressRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRCInstanceIpAddressRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeRCInstanceIpAddressRequest) GetDdosStatus() *string {
	return s.DdosStatus
}

func (s *DescribeRCInstanceIpAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceIpAddressRequest) GetInstanceIp() *string {
	return s.InstanceIp
}

func (s *DescribeRCInstanceIpAddressRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstanceIpAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCInstanceIpAddressRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCInstanceIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceIpAddressRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRCInstanceIpAddressRequest) SetCurrentPage(v int32) *DescribeRCInstanceIpAddressRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetDdosRegionId(v string) *DescribeRCInstanceIpAddressRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetDdosStatus(v string) *DescribeRCInstanceIpAddressRequest {
	s.DdosStatus = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetInstanceId(v string) *DescribeRCInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetInstanceIp(v string) *DescribeRCInstanceIpAddressRequest {
	s.InstanceIp = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetInstanceName(v string) *DescribeRCInstanceIpAddressRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetInstanceType(v string) *DescribeRCInstanceIpAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetPageSize(v int32) *DescribeRCInstanceIpAddressRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetRegionId(v string) *DescribeRCInstanceIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) SetResourceType(v string) *DescribeRCInstanceIpAddressRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRCInstanceIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
