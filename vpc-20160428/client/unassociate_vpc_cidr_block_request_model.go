// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateVpcCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPv6CidrBlock(v string) *UnassociateVpcCidrBlockRequest
	GetIPv6CidrBlock() *string
	SetOwnerAccount(v string) *UnassociateVpcCidrBlockRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateVpcCidrBlockRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateVpcCidrBlockRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateVpcCidrBlockRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateVpcCidrBlockRequest
	GetResourceOwnerId() *int64
	SetSecondaryCidrBlock(v string) *UnassociateVpcCidrBlockRequest
	GetSecondaryCidrBlock() *string
	SetVpcId(v string) *UnassociateVpcCidrBlockRequest
	GetVpcId() *string
}

type UnassociateVpcCidrBlockRequest struct {
	// The secondary IPv6 CIDR block to be deleted.
	//
	// >  You must set one of the **Ipv6CidrBlock*	- and **SecondaryCidrBlock*	- parameters.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	IPv6CidrBlock *string `json:"IPv6CidrBlock,omitempty" xml:"IPv6CidrBlock,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC to which the secondary CIDR block to be deleted belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ch-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The secondary IPv4 CIDR block to be deleted.
	//
	// >  You must set one of the **SecondaryCidrBlock*	- and **Ipv6CidrBlock*	- parameters.
	//
	// example:
	//
	// 192.168.0.0/16
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// The ID of the VPC from which you want to delete a secondary CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-o6wrloqsdqc9io3mg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UnassociateVpcCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateVpcCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *UnassociateVpcCidrBlockRequest) GetIPv6CidrBlock() *string {
	return s.IPv6CidrBlock
}

func (s *UnassociateVpcCidrBlockRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateVpcCidrBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateVpcCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateVpcCidrBlockRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateVpcCidrBlockRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateVpcCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *UnassociateVpcCidrBlockRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UnassociateVpcCidrBlockRequest) SetIPv6CidrBlock(v string) *UnassociateVpcCidrBlockRequest {
	s.IPv6CidrBlock = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetOwnerAccount(v string) *UnassociateVpcCidrBlockRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetOwnerId(v int64) *UnassociateVpcCidrBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetRegionId(v string) *UnassociateVpcCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetResourceOwnerAccount(v string) *UnassociateVpcCidrBlockRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetResourceOwnerId(v int64) *UnassociateVpcCidrBlockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetSecondaryCidrBlock(v string) *UnassociateVpcCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) SetVpcId(v string) *UnassociateVpcCidrBlockRequest {
	s.VpcId = &v
	return s
}

func (s *UnassociateVpcCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
