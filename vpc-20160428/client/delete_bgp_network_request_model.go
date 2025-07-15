// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteBgpNetworkRequest
	GetClientToken() *string
	SetDstCidrBlock(v string) *DeleteBgpNetworkRequest
	GetDstCidrBlock() *string
	SetOwnerAccount(v string) *DeleteBgpNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBgpNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBgpNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBgpNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBgpNetworkRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *DeleteBgpNetworkRequest
	GetRouterId() *string
}

type DeleteBgpNetworkRequest struct {
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
	// The CIDR block of the virtual private cloud (VPC) or vSwitch that you want to connect to a data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.110.192.12/32
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the BGP group.
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
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
}

func (s DeleteBgpNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteBgpNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBgpNetworkRequest) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *DeleteBgpNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBgpNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBgpNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBgpNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBgpNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBgpNetworkRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DeleteBgpNetworkRequest) SetClientToken(v string) *DeleteBgpNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetDstCidrBlock(v string) *DeleteBgpNetworkRequest {
	s.DstCidrBlock = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetOwnerAccount(v string) *DeleteBgpNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetOwnerId(v int64) *DeleteBgpNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetRegionId(v string) *DeleteBgpNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetResourceOwnerAccount(v string) *DeleteBgpNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetResourceOwnerId(v int64) *DeleteBgpNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBgpNetworkRequest) SetRouterId(v string) *DeleteBgpNetworkRequest {
	s.RouterId = &v
	return s
}

func (s *DeleteBgpNetworkRequest) Validate() error {
	return dara.Validate(s)
}
