// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *ChangeResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ChangeResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ChangeResourceGroupRequest
	GetRegionId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *ChangeResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChangeResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek3ctkufaw****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the hosted region. Call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to get the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource whose resource group you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-uq5dcfc2eqhpf4****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - Ipam: an IPAM instance
	//
	// - IpamScope: an IPAM scope
	//
	// - IpamPool: an IPAM pool
	//
	// - IpamResourceDiscovery: an IPAM resource discovery
	//
	// This parameter is required.
	//
	// example:
	//
	// Ipam
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ChangeResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChangeResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetOwnerAccount(v string) *ChangeResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetOwnerId(v int64) *ChangeResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceOwnerAccount(v string) *ChangeResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceOwnerId(v int64) *ChangeResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
