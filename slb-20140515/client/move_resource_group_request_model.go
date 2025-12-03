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
	SetResourceGroupId(v string) *MoveResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *MoveResourceGroupRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *MoveResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MoveResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *MoveResourceGroupRequest
	GetResourceType() *string
	SetAccessKeyId(v string) *MoveResourceGroupRequest
	GetAccessKeyId() *string
}

type MoveResourceGroupRequest struct {
	// The ID of the resource group to which you want to move the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2rpsek5h****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2dmxj56z****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource for which you want to modify the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-uf6ghek7ds2btzt65****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **loadbalancer**: a CLB instance
	//
	// 	- **certificate**: a certificate
	//
	// 	- **acl**: an access control list (ACL)
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The AccessKey ID provided to you by Alibaba Cloud for accessing the service.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
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

func (s *MoveResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *MoveResourceGroupRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
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

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
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

func (s *MoveResourceGroupRequest) SetAccessKeyId(v string) *MoveResourceGroupRequest {
	s.AccessKeyId = &v
	return s
}

func (s *MoveResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
