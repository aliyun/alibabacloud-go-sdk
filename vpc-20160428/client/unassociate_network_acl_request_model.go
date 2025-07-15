// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassociateNetworkAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *UnassociateNetworkAclRequest
	GetDryRun() *bool
	SetNetworkAclId(v string) *UnassociateNetworkAclRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *UnassociateNetworkAclRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateNetworkAclRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateNetworkAclRequest
	GetRegionId() *string
	SetResource(v []*UnassociateNetworkAclRequestResource) *UnassociateNetworkAclRequest
	GetResource() []*UnassociateNetworkAclRequestResource
	SetResourceOwnerAccount(v string) *UnassociateNetworkAclRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateNetworkAclRequest
	GetResourceOwnerId() *int64
}

type UnassociateNetworkAclRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the network ACL that you want to disassociate from a resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0sp****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the associated resource.
	Resource             []*UnassociateNetworkAclRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                                 `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateNetworkAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UnassociateNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *UnassociateNetworkAclRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateNetworkAclRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateNetworkAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateNetworkAclRequest) GetResource() []*UnassociateNetworkAclRequestResource {
	return s.Resource
}

func (s *UnassociateNetworkAclRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateNetworkAclRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateNetworkAclRequest) SetClientToken(v string) *UnassociateNetworkAclRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetDryRun(v bool) *UnassociateNetworkAclRequest {
	s.DryRun = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetNetworkAclId(v string) *UnassociateNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetOwnerAccount(v string) *UnassociateNetworkAclRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetOwnerId(v int64) *UnassociateNetworkAclRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetRegionId(v string) *UnassociateNetworkAclRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetResource(v []*UnassociateNetworkAclRequestResource) *UnassociateNetworkAclRequest {
	s.Resource = v
	return s
}

func (s *UnassociateNetworkAclRequest) SetResourceOwnerAccount(v string) *UnassociateNetworkAclRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetResourceOwnerId(v int64) *UnassociateNetworkAclRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}

type UnassociateNetworkAclRequestResource struct {
	// The ID of the resource from which you want to disassociate the network ACL.
	//
	// example:
	//
	// vsw-bp1de348lntdw****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource from which you want to disassociate the network ACL. Set the value to **VSwitch**.
	//
	// Valid values of **N**: 0 to 29. You can disassociate a network ACL from at most 30 resources at a time.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UnassociateNetworkAclRequestResource) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclRequestResource) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *UnassociateNetworkAclRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnassociateNetworkAclRequestResource) SetResourceId(v string) *UnassociateNetworkAclRequestResource {
	s.ResourceId = &v
	return s
}

func (s *UnassociateNetworkAclRequestResource) SetResourceType(v string) *UnassociateNetworkAclRequestResource {
	s.ResourceType = &v
	return s
}

func (s *UnassociateNetworkAclRequestResource) Validate() error {
	return dara.Validate(s)
}
