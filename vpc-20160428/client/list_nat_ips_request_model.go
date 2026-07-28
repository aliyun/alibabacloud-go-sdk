// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListNatIpsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ListNatIpsRequest
	GetDryRun() *bool
	SetIpOrigin(v string) *ListNatIpsRequest
	GetIpOrigin() *string
	SetIpv4Prefix(v string) *ListNatIpsRequest
	GetIpv4Prefix() *string
	SetMaxResults(v string) *ListNatIpsRequest
	GetMaxResults() *string
	SetNatGatewayId(v string) *ListNatIpsRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *ListNatIpsRequest
	GetNatIpCidr() *string
	SetNatIpIds(v []*string) *ListNatIpsRequest
	GetNatIpIds() []*string
	SetNatIpName(v []*string) *ListNatIpsRequest
	GetNatIpName() []*string
	SetNatIpStatus(v string) *ListNatIpsRequest
	GetNatIpStatus() *string
	SetNextToken(v string) *ListNatIpsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListNatIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListNatIpsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListNatIpsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListNatIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListNatIpsRequest
	GetResourceOwnerId() *int64
}

type ListNatIpsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The origin of the NAT IP address to query. Valid values:
	//
	// - prefix: a NAT IP address that belongs to an IP prefix.
	//
	// - cidr: a standalone NAT IP address that does not belong to any IP prefix.
	//
	// - Empty: queries all NAT IP addresses.
	//
	// example:
	//
	// cidr
	IpOrigin *string `json:"IpOrigin,omitempty" xml:"IpOrigin,omitempty"`
	// The CIDR block of the IP prefix to query.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The number of entries per page for a paged query. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The instance ID of the NAT gateway to which the NAT IP addresses belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The CIDR block to which the NAT IP addresses belong.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The instance ID of the NAT IP address to query. Valid values of **N**: **1*	- to **20**.
	//
	// example:
	//
	// vpcnatip-gw8a863sut1zijxh0****
	NatIpIds []*string `json:"NatIpIds,omitempty" xml:"NatIpIds,omitempty" type:"Repeated"`
	// The name of the NAT IP address to query. Valid values of **N**: **1*	- to **20**.
	//
	// example:
	//
	// test
	NatIpName []*string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty" type:"Repeated"`
	// The status of the NAT IP addresses to query. Valid values:
	//
	// - **Available**: available.
	//
	// - **Deleting**: being deleted.
	//
	// - **Creating**: being created.
	//
	// example:
	//
	// Available
	NatIpStatus *string `json:"NatIpStatus,omitempty" xml:"NatIpStatus,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first request or if no subsequent query exists.
	//
	// - If a next query exists, set the value to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT IP addresses belong.
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

func (s ListNatIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpsRequest) GoString() string {
	return s.String()
}

func (s *ListNatIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListNatIpsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListNatIpsRequest) GetIpOrigin() *string {
	return s.IpOrigin
}

func (s *ListNatIpsRequest) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *ListNatIpsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListNatIpsRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListNatIpsRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *ListNatIpsRequest) GetNatIpIds() []*string {
	return s.NatIpIds
}

func (s *ListNatIpsRequest) GetNatIpName() []*string {
	return s.NatIpName
}

func (s *ListNatIpsRequest) GetNatIpStatus() *string {
	return s.NatIpStatus
}

func (s *ListNatIpsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNatIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListNatIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListNatIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNatIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListNatIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListNatIpsRequest) SetClientToken(v string) *ListNatIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListNatIpsRequest) SetDryRun(v bool) *ListNatIpsRequest {
	s.DryRun = &v
	return s
}

func (s *ListNatIpsRequest) SetIpOrigin(v string) *ListNatIpsRequest {
	s.IpOrigin = &v
	return s
}

func (s *ListNatIpsRequest) SetIpv4Prefix(v string) *ListNatIpsRequest {
	s.Ipv4Prefix = &v
	return s
}

func (s *ListNatIpsRequest) SetMaxResults(v string) *ListNatIpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNatIpsRequest) SetNatGatewayId(v string) *ListNatIpsRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ListNatIpsRequest) SetNatIpCidr(v string) *ListNatIpsRequest {
	s.NatIpCidr = &v
	return s
}

func (s *ListNatIpsRequest) SetNatIpIds(v []*string) *ListNatIpsRequest {
	s.NatIpIds = v
	return s
}

func (s *ListNatIpsRequest) SetNatIpName(v []*string) *ListNatIpsRequest {
	s.NatIpName = v
	return s
}

func (s *ListNatIpsRequest) SetNatIpStatus(v string) *ListNatIpsRequest {
	s.NatIpStatus = &v
	return s
}

func (s *ListNatIpsRequest) SetNextToken(v string) *ListNatIpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNatIpsRequest) SetOwnerAccount(v string) *ListNatIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListNatIpsRequest) SetOwnerId(v int64) *ListNatIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListNatIpsRequest) SetRegionId(v string) *ListNatIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNatIpsRequest) SetResourceOwnerAccount(v string) *ListNatIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListNatIpsRequest) SetResourceOwnerId(v int64) *ListNatIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListNatIpsRequest) Validate() error {
	return dara.Validate(s)
}
