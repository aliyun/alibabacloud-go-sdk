// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBgpNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddBgpNetworkRequest
	GetClientToken() *string
	SetDstCidrBlock(v string) *AddBgpNetworkRequest
	GetDstCidrBlock() *string
	SetOwnerAccount(v string) *AddBgpNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddBgpNetworkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddBgpNetworkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddBgpNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBgpNetworkRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *AddBgpNetworkRequest
	GetRouterId() *string
	SetVpcId(v string) *AddBgpNetworkRequest
	GetVpcId() *string
}

type AddBgpNetworkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
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
	// 10.10.XX.XX/32
	DstCidrBlock *string `json:"DstCidrBlock,omitempty" xml:"DstCidrBlock,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the virtual border router (VBR) group.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the router that is associated with the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// vrt-2zeo3xzyf38r4u******
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1qpo0kug3a2*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddBgpNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBgpNetworkRequest) GoString() string {
	return s.String()
}

func (s *AddBgpNetworkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddBgpNetworkRequest) GetDstCidrBlock() *string {
	return s.DstCidrBlock
}

func (s *AddBgpNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddBgpNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBgpNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBgpNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBgpNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBgpNetworkRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *AddBgpNetworkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AddBgpNetworkRequest) SetClientToken(v string) *AddBgpNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *AddBgpNetworkRequest) SetDstCidrBlock(v string) *AddBgpNetworkRequest {
	s.DstCidrBlock = &v
	return s
}

func (s *AddBgpNetworkRequest) SetOwnerAccount(v string) *AddBgpNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddBgpNetworkRequest) SetOwnerId(v int64) *AddBgpNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBgpNetworkRequest) SetRegionId(v string) *AddBgpNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *AddBgpNetworkRequest) SetResourceOwnerAccount(v string) *AddBgpNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBgpNetworkRequest) SetResourceOwnerId(v int64) *AddBgpNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBgpNetworkRequest) SetRouterId(v string) *AddBgpNetworkRequest {
	s.RouterId = &v
	return s
}

func (s *AddBgpNetworkRequest) SetVpcId(v string) *AddBgpNetworkRequest {
	s.VpcId = &v
	return s
}

func (s *AddBgpNetworkRequest) Validate() error {
	return dara.Validate(s)
}
