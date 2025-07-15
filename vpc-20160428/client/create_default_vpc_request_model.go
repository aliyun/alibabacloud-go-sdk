// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDefaultVpcRequest
	GetClientToken() *string
	SetEnableIpv6(v bool) *CreateDefaultVpcRequest
	GetEnableIpv6() *bool
	SetIpv6CidrBlock(v string) *CreateDefaultVpcRequest
	GetIpv6CidrBlock() *string
	SetOwnerAccount(v string) *CreateDefaultVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDefaultVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDefaultVpcRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDefaultVpcRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDefaultVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDefaultVpcRequest
	GetResourceOwnerId() *int64
}

type CreateDefaultVpcRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The IPv6 CIDR block of the default VPC.
	//
	// > When **EnableIpv6*	- is set to **true**, this parameter is required.
	//
	// example:
	//
	// 2408:XXXX:346:b600::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the default VPC belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmystnjq4****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDefaultVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateDefaultVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDefaultVpcRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateDefaultVpcRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *CreateDefaultVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDefaultVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDefaultVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefaultVpcRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDefaultVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDefaultVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDefaultVpcRequest) SetClientToken(v string) *CreateDefaultVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetEnableIpv6(v bool) *CreateDefaultVpcRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetIpv6CidrBlock(v string) *CreateDefaultVpcRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetOwnerAccount(v string) *CreateDefaultVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetOwnerId(v int64) *CreateDefaultVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetRegionId(v string) *CreateDefaultVpcRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetResourceGroupId(v string) *CreateDefaultVpcRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetResourceOwnerAccount(v string) *CreateDefaultVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDefaultVpcRequest) SetResourceOwnerId(v int64) *CreateDefaultVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDefaultVpcRequest) Validate() error {
	return dara.Validate(s)
}
