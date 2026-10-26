// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationStatus(v string) *DescribeDedicatedHostsRequest
	GetAllocationStatus() *string
	SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostGroupId() *string
	SetDedicatedHostId(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostId() *string
	SetHostStatus(v string) *DescribeDedicatedHostsRequest
	GetHostStatus() *string
	SetHostType(v string) *DescribeDedicatedHostsRequest
	GetHostType() *string
	SetOrderId(v int64) *DescribeDedicatedHostsRequest
	GetOrderId() *int64
	SetOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDedicatedHostsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeDedicatedHostsRequest
	GetZoneId() *string
}

type DescribeDedicatedHostsRequest struct {
	// Specifies whether the host allows instance allocation. Valid values:
	//
	// 	- **0**: Not allowed.
	//
	// 	- **1**: Allowed.
	//
	// example:
	//
	// 1
	AllocationStatus *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	// The dedicated cluster ID. You can call the DescribeDedicatedHostGroups operation to obtain the dedicated cluster ID.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host ID of the dedicated cluster.
	//
	// example:
	//
	// ch-t4nn100ddxxxxxxxx
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The host status. Valid values:
	//
	//  	- **0**: Creating.
	//
	//  	- **1**: Running.
	//
	//  	- **2**: Down.
	//
	//  	- **3**: Down and offline (host replacement in progress).
	//
	//  	- **4**: Offline.
	//
	//  	- **5**: Deleted.
	//
	//  	- **6**: Restarting.
	//
	// example:
	//
	// 1
	HostStatus *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	// The storage type of the host. Valid values:
	//
	// 	- **dhg_cloud_ssd**: ESSD cloud disk.
	//
	// 	- **dhg_local_ssd**: local SSD.
	//
	// example:
	//
	// dhg_cloud_ssd
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 102565235
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequest) GetAllocationStatus() *string {
	return s.AllocationStatus
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeDedicatedHostsRequest) GetHostStatus() *string {
	return s.HostStatus
}

func (s *DescribeDedicatedHostsRequest) GetHostType() *string {
	return s.HostType
}

func (s *DescribeDedicatedHostsRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeDedicatedHostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostsRequest) SetAllocationStatus(v string) *DescribeDedicatedHostsRequest {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostStatus(v string) *DescribeDedicatedHostsRequest {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetHostType(v string) *DescribeDedicatedHostsRequest {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOrderId(v int64) *DescribeDedicatedHostsRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetRegionId(v string) *DescribeDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetZoneId(v string) *DescribeDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) Validate() error {
	return dara.Validate(s)
}
