// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateNetworkAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateNetworkAclRequest
	GetDryRun() *bool
	SetNetworkAclId(v string) *AssociateNetworkAclRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *AssociateNetworkAclRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateNetworkAclRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateNetworkAclRequest
	GetRegionId() *string
	SetResource(v []*AssociateNetworkAclRequestResource) *AssociateNetworkAclRequest
	GetResource() []*AssociateNetworkAclRequestResource
	SetResourceOwnerAccount(v string) *AssociateNetworkAclRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateNetworkAclRequest
	GetResourceOwnerId() *int64
}

type AssociateNetworkAclRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0sp****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the associated resources.
	Resource             []*AssociateNetworkAclRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssociateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *AssociateNetworkAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateNetworkAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *AssociateNetworkAclRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateNetworkAclRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateNetworkAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateNetworkAclRequest) GetResource() []*AssociateNetworkAclRequestResource {
	return s.Resource
}

func (s *AssociateNetworkAclRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateNetworkAclRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateNetworkAclRequest) SetClientToken(v string) *AssociateNetworkAclRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetDryRun(v bool) *AssociateNetworkAclRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetNetworkAclId(v string) *AssociateNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetOwnerAccount(v string) *AssociateNetworkAclRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetOwnerId(v int64) *AssociateNetworkAclRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetRegionId(v string) *AssociateNetworkAclRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetResource(v []*AssociateNetworkAclRequestResource) *AssociateNetworkAclRequest {
	s.Resource = v
	return s
}

func (s *AssociateNetworkAclRequest) SetResourceOwnerAccount(v string) *AssociateNetworkAclRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateNetworkAclRequest) SetResourceOwnerId(v int64) *AssociateNetworkAclRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}

type AssociateNetworkAclRequestResource struct {
	// The ID of the associated resource.
	//
	// example:
	//
	// vsw-bp1de348lntdw****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource with which you want to associate the network ACL. Set the value to **VSwitch**.
	//
	// Valid values of **N**: **0*	- to **29**. You can associate a network ACL with up to 30 vSwitches.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AssociateNetworkAclRequestResource) String() string {
	return dara.Prettify(s)
}

func (s AssociateNetworkAclRequestResource) GoString() string {
	return s.String()
}

func (s *AssociateNetworkAclRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *AssociateNetworkAclRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *AssociateNetworkAclRequestResource) SetResourceId(v string) *AssociateNetworkAclRequestResource {
	s.ResourceId = &v
	return s
}

func (s *AssociateNetworkAclRequestResource) SetResourceType(v string) *AssociateNetworkAclRequestResource {
	s.ResourceType = &v
	return s
}

func (s *AssociateNetworkAclRequestResource) Validate() error {
	return dara.Validate(s)
}
