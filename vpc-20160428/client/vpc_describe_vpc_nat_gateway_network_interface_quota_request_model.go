// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetClientToken() *string
	SetNatGatewayId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetOwnerId() *int64
	SetRegionId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetResourceOwnerId() *int64
	SetResourceUid(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest
	GetResourceUid() *int64
}

type VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the VPC NAT gateway.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceUid *int64 `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) GetResourceUid() *int64 {
	return s.ResourceUid
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetClientToken(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetNatGatewayId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.NatGatewayId = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetOwnerAccount(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetOwnerId(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetRegionId(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetResourceOwnerAccount(v string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetResourceOwnerId(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) SetResourceUid(v int64) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest {
	s.ResourceUid = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
