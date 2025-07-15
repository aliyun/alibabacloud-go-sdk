// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeNetworkAclAttributesRequest
	GetClientToken() *string
	SetNetworkAclId(v string) *DescribeNetworkAclAttributesRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *DescribeNetworkAclAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkAclAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNetworkAclAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkAclAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkAclAttributesRequest
	GetResourceOwnerId() *int64
}

type DescribeNetworkAclAttributesRequest struct {
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
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
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
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeNetworkAclAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNetworkAclAttributesRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkAclAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkAclAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkAclAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkAclAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkAclAttributesRequest) SetClientToken(v string) *DescribeNetworkAclAttributesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetNetworkAclId(v string) *DescribeNetworkAclAttributesRequest {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetOwnerAccount(v string) *DescribeNetworkAclAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetOwnerId(v int64) *DescribeNetworkAclAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetRegionId(v string) *DescribeNetworkAclAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetResourceOwnerAccount(v string) *DescribeNetworkAclAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) SetResourceOwnerId(v int64) *DescribeNetworkAclAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkAclAttributesRequest) Validate() error {
	return dara.Validate(s)
}
