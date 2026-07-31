// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterIds(v string) *DescribeDedicatedHostClustersRequest
	GetDedicatedHostClusterIds() *string
	SetDedicatedHostClusterName(v string) *DescribeDedicatedHostClustersRequest
	GetDedicatedHostClusterName() *string
	SetLockReason(v string) *DescribeDedicatedHostClustersRequest
	GetLockReason() *string
	SetOwnerAccount(v string) *DescribeDedicatedHostClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedHostClustersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDedicatedHostClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedHostClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDedicatedHostClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedHostClustersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostClustersRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeDedicatedHostClustersRequest
	GetStatus() *string
	SetTag(v []*DescribeDedicatedHostClustersRequestTag) *DescribeDedicatedHostClustersRequest
	GetTag() []*DescribeDedicatedHostClustersRequestTag
	SetZoneId(v string) *DescribeDedicatedHostClustersRequest
	GetZoneId() *string
}

type DescribeDedicatedHostClustersRequest struct {
	// The IDs of dedicated host clusters. The value is a JSON array of dedicated host cluster IDs in the format of `["dc-xxxxxxxxx", "dc-yyyyyyyyy", … ,"dc-zzzzzzzzz"]`. A maximum of 100 IDs are supported. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["dc-bp12wlf6am0vz9v2****", "dc-bp12wlf6am0vz9v3****"]
	DedicatedHostClusterIds *string `json:"DedicatedHostClusterIds,omitempty" xml:"DedicatedHostClusterIds,omitempty"`
	// The name of the dedicated host cluster.
	//
	// example:
	//
	// myDDHCluster
	DedicatedHostClusterName *string `json:"DedicatedHostClusterName,omitempty" xml:"DedicatedHostClusterName,omitempty"`
	// >This parameter is not yet available.
	//
	// example:
	//
	// null
	LockReason   *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the dedicated host cluster list.
	//
	// Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for the paged query. Settings for paging:
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the dedicated host cluster. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host cluster belongs. When you use this parameter to filter resources, the resource count cannot exceed 1,000.
	//
	// >Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// >This parameter is not yet available.
	//
	// example:
	//
	// null
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeDedicatedHostClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the dedicated host cluster. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the available zones.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersRequest) GetDedicatedHostClusterIds() *string {
	return s.DedicatedHostClusterIds
}

func (s *DescribeDedicatedHostClustersRequest) GetDedicatedHostClusterName() *string {
	return s.DedicatedHostClusterName
}

func (s *DescribeDedicatedHostClustersRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDedicatedHostClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedHostClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedHostClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedHostClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedHostClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostClustersRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedHostClustersRequest) GetTag() []*DescribeDedicatedHostClustersRequestTag {
	return s.Tag
}

func (s *DescribeDedicatedHostClustersRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostClustersRequest) SetDedicatedHostClusterIds(v string) *DescribeDedicatedHostClustersRequest {
	s.DedicatedHostClusterIds = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetDedicatedHostClusterName(v string) *DescribeDedicatedHostClustersRequest {
	s.DedicatedHostClusterName = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetLockReason(v string) *DescribeDedicatedHostClustersRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetOwnerAccount(v string) *DescribeDedicatedHostClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetOwnerId(v int64) *DescribeDedicatedHostClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetPageNumber(v int32) *DescribeDedicatedHostClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetPageSize(v int32) *DescribeDedicatedHostClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetRegionId(v string) *DescribeDedicatedHostClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetResourceGroupId(v string) *DescribeDedicatedHostClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetStatus(v string) *DescribeDedicatedHostClustersRequest {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetTag(v []*DescribeDedicatedHostClustersRequestTag) *DescribeDedicatedHostClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) SetZoneId(v string) *DescribeDedicatedHostClustersRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostClustersRequestTag struct {
	// The tag key of the dedicated host cluster. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// If you use a single tag to filter resources, the resource count with the tag cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources that have all specified tags attached cannot exceed 1,000. If the resource count exceeds 1,000, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the dedicated host cluster. Valid values of N: 1 to 20. The tag value cannot be an empty string. The tag value can be up to 64 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDedicatedHostClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDedicatedHostClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDedicatedHostClustersRequestTag) SetKey(v string) *DescribeDedicatedHostClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequestTag) SetValue(v string) *DescribeDedicatedHostClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDedicatedHostClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
