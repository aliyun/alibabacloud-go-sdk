// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *MoveResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *MoveResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MoveResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MoveResourceGroupRequest
	GetRegionId() *string
	SetResourceId(v string) *MoveResourceGroupRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *MoveResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *MoveResourceGroupRequest
	GetResourceType() *string
}

type MoveResourceGroupRequest struct {
	// The ID of the resource group to which you want to move the cloud resource instance.
	//
	// > A resource group is a mechanism for grouping and managing resources within an Alibaba Cloud account. Resource groups help you address the complexity of resource grouping and authorization management within a single cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfm3peow3k****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cloud resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the cloud resource for which you want to modify the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-hp31psbg8ec3023s6****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the cloud resource for which you want to modify the resource group. Valid values:
	//
	// - **Vpc**: virtual private cloud (VPC)
	//
	// - **Eip**: elastic IP address (EIP)
	//
	// - **BandwidthPackage**: Internet Shared Bandwidth
	//
	// - **PrefixList**: prefix list
	//
	// - **PublicIpAddressPool**: IPAM pool
	//
	// - **FlowLog**: flow log
	//
	// - **HaVip**: high-availability virtual IP address
	//
	// - **TrafficMirrorFilter**: traffic mirror filter
	//
	// - **TrafficMirrorSession**: traffic mirror session
	//
	// - **IPv4Gateway**: IPv4 gateway
	//
	// - **IPv6Gateway**: IPv6 gateway
	//
	// - **IPv6Address**: IPv6 address
	//
	// - **DhcpOptionsSet**: DHCP options set
	//
	// - **GatewayEndpoint**: gateway endpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MoveResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MoveResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MoveResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MoveResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerAccount(v string) *MoveResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetOwnerId(v int64) *MoveResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerAccount(v string) *MoveResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceOwnerId(v int64) *MoveResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
