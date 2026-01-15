// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteNatIpRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteNatIpRequest
	GetDryRun() *bool
	SetIpv4Prefix(v string) *DeleteNatIpRequest
	GetIpv4Prefix() *string
	SetNatGatewayId(v string) *DeleteNatIpRequest
	GetNatGatewayId() *string
	SetNatIpId(v string) *DeleteNatIpRequest
	GetNatIpId() *string
	SetOwnerAccount(v string) *DeleteNatIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNatIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNatIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNatIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNatIpRequest
	GetResourceOwnerId() *int64
}

type DeleteNatIpRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP prefix address to be deleted.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The ID of the NAT gateway instance to which the IP prefix to be deleted belongs. Required when deleting an IP prefix.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the NAT IP address that you want to delete.
	//
	// example:
	//
	// vpcnatip-gw8y7q3cpk3fggs87****
	NatIpId      *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway to which the NAT IP address that you want to delete belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteNatIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNatIpRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteNatIpRequest) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *DeleteNatIpRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatIpRequest) GetNatIpId() *string {
	return s.NatIpId
}

func (s *DeleteNatIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNatIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNatIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNatIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNatIpRequest) SetClientToken(v string) *DeleteNatIpRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNatIpRequest) SetDryRun(v bool) *DeleteNatIpRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteNatIpRequest) SetIpv4Prefix(v string) *DeleteNatIpRequest {
	s.Ipv4Prefix = &v
	return s
}

func (s *DeleteNatIpRequest) SetNatGatewayId(v string) *DeleteNatIpRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatIpRequest) SetNatIpId(v string) *DeleteNatIpRequest {
	s.NatIpId = &v
	return s
}

func (s *DeleteNatIpRequest) SetOwnerAccount(v string) *DeleteNatIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNatIpRequest) SetOwnerId(v int64) *DeleteNatIpRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNatIpRequest) SetRegionId(v string) *DeleteNatIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatIpRequest) SetResourceOwnerAccount(v string) *DeleteNatIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNatIpRequest) SetResourceOwnerId(v int64) *DeleteNatIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNatIpRequest) Validate() error {
	return dara.Validate(s)
}
