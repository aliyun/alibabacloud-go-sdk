// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpCidrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListNatIpCidrsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListNatIpCidrsRequest
	GetDryRun() *bool
	SetMaxResults(v string) *ListNatIpCidrsRequest
	GetMaxResults() *string
	SetNatGatewayId(v string) *ListNatIpCidrsRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *ListNatIpCidrsRequest
	GetNatIpCidr() *string
	SetNatIpCidrName(v []*string) *ListNatIpCidrsRequest
	GetNatIpCidrName() []*string
	SetNatIpCidrStatus(v string) *ListNatIpCidrsRequest
	GetNatIpCidrStatus() *string
	SetNatIpCidrs(v []*string) *ListNatIpCidrsRequest
	GetNatIpCidrs() []*string
	SetNextToken(v string) *ListNatIpCidrsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListNatIpCidrsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListNatIpCidrsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListNatIpCidrsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListNatIpCidrsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListNatIpCidrsRequest
	GetResourceOwnerId() *int64
}

type ListNatIpCidrsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// 	- **true**: checks the API request. The CIDR blocks of the NAT gateway are not queried if the API request passes the precheck. The system checks whether your AccessKey pair is valid, whether the Resource Access Management (RAM) user is authorized, and whether the required parameters are set. If the request fails to pass the precheck, the corresponding error message is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// 	- **false**: sends the API request. If the request passes the precheck, 2xx HTTP status code is returned and the CIDR blocks of the NAT gateway are queried. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the VPC NAT gateway that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The CIDR block of the NAT gateway that you want to query.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The name of the CIDR block that you want to query. Valid values of **N**: **1*	- to **20**.
	//
	// example:
	//
	// test
	NatIpCidrName []*string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty" type:"Repeated"`
	// The status of the CIDR block that you want to query. Set the value to **Available**.
	//
	// example:
	//
	// Available
	NatIpCidrStatus *string `json:"NatIpCidrStatus,omitempty" xml:"NatIpCidrStatus,omitempty"`
	// The CIDR block of the NAT gateway that you want to query. Valid values of **N**: **1*	- to **20**.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidrs []*string `json:"NatIpCidrs,omitempty" xml:"NatIpCidrs,omitempty" type:"Repeated"`
	// The token that is used for the next query. Set the value as needed.
	//
	// 	- If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Virtual Private Cloud (VPC) NAT gateway that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListNatIpCidrsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpCidrsRequest) GoString() string {
	return s.String()
}

func (s *ListNatIpCidrsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListNatIpCidrsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListNatIpCidrsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListNatIpCidrsRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListNatIpCidrsRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *ListNatIpCidrsRequest) GetNatIpCidrName() []*string {
	return s.NatIpCidrName
}

func (s *ListNatIpCidrsRequest) GetNatIpCidrStatus() *string {
	return s.NatIpCidrStatus
}

func (s *ListNatIpCidrsRequest) GetNatIpCidrs() []*string {
	return s.NatIpCidrs
}

func (s *ListNatIpCidrsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNatIpCidrsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListNatIpCidrsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListNatIpCidrsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNatIpCidrsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListNatIpCidrsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListNatIpCidrsRequest) SetClientToken(v string) *ListNatIpCidrsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetDryRun(v bool) *ListNatIpCidrsRequest {
	s.DryRun = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetMaxResults(v string) *ListNatIpCidrsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetNatGatewayId(v string) *ListNatIpCidrsRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetNatIpCidr(v string) *ListNatIpCidrsRequest {
	s.NatIpCidr = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetNatIpCidrName(v []*string) *ListNatIpCidrsRequest {
	s.NatIpCidrName = v
	return s
}

func (s *ListNatIpCidrsRequest) SetNatIpCidrStatus(v string) *ListNatIpCidrsRequest {
	s.NatIpCidrStatus = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetNatIpCidrs(v []*string) *ListNatIpCidrsRequest {
	s.NatIpCidrs = v
	return s
}

func (s *ListNatIpCidrsRequest) SetNextToken(v string) *ListNatIpCidrsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetOwnerAccount(v string) *ListNatIpCidrsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetOwnerId(v int64) *ListNatIpCidrsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetRegionId(v string) *ListNatIpCidrsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetResourceOwnerAccount(v string) *ListNatIpCidrsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListNatIpCidrsRequest) SetResourceOwnerId(v int64) *ListNatIpCidrsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListNatIpCidrsRequest) Validate() error {
	return dara.Validate(s)
}
