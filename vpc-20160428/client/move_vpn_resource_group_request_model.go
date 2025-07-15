// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveVpnResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *MoveVpnResourceGroupRequest
	GetInstanceId() *string
	SetNewResourceGroupId(v string) *MoveVpnResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *MoveVpnResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MoveVpnResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MoveVpnResourceGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *MoveVpnResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveVpnResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *MoveVpnResourceGroupRequest
	GetResourceType() *string
}

type MoveVpnResourceGroupRequest struct {
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-8vb3lzn7biepthri8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmzs372yg****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource.
	//
	// 	- **VpnGateway**: VPN gateway
	//
	//     After you move a VPN gateway to a new resource group, the following associated resources are also moved to the new resource group: IPsec servers, SSL servers, SSL client certificates, and IPsec-VPN connections.
	//
	// 	- **CustomerGateway**: customer gateway
	//
	// 	- **VpnAttachment**: IPsec-VPN connection
	//
	//     An IPsec-VPN connection associated with a transit router or not associate with a resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// VpnGateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveVpnResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveVpnResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveVpnResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveVpnResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveVpnResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MoveVpnResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MoveVpnResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveVpnResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MoveVpnResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MoveVpnResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *MoveVpnResourceGroupRequest) SetInstanceId(v string) *MoveVpnResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetNewResourceGroupId(v string) *MoveVpnResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetOwnerAccount(v string) *MoveVpnResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetOwnerId(v int64) *MoveVpnResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetRegionId(v string) *MoveVpnResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetResourceOwnerAccount(v string) *MoveVpnResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetResourceOwnerId(v int64) *MoveVpnResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) SetResourceType(v string) *MoveVpnResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveVpnResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
