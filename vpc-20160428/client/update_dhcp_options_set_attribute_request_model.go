// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDhcpOptionsSetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetClientToken() *string
	SetDhcpOptionsSetDescription(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetDhcpOptionsSetDescription() *string
	SetDhcpOptionsSetId(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetDhcpOptionsSetId() *string
	SetDhcpOptionsSetName(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetDhcpOptionsSetName() *string
	SetDomainName(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetDomainName() *string
	SetDomainNameServers(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetDomainNameServers() *string
	SetDryRun(v bool) *UpdateDhcpOptionsSetAttributeRequest
	GetDryRun() *bool
	SetIpv6LeaseTime(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetIpv6LeaseTime() *string
	SetLeaseTime(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetLeaseTime() *string
	SetOwnerAccount(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateDhcpOptionsSetAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateDhcpOptionsSetAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateDhcpOptionsSetAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdateDhcpOptionsSetAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Enter a description for the DHCP options set.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`. You can also leave the description empty.
	//
	// example:
	//
	// description
	DhcpOptionsSetDescription *string `json:"DhcpOptionsSetDescription,omitempty" xml:"DhcpOptionsSetDescription,omitempty"`
	// The ID of the DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// The name of the DHCP options set.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// name
	DhcpOptionsSetName *string `json:"DhcpOptionsSetName,omitempty" xml:"DhcpOptionsSetName,omitempty"`
	// The root domain. For example, you can set the value to example.com.
	//
	// After a DHCP options set is associated with a virtual private cloud (VPC), the root domain in the DHCP options set is automatically synchronized with the ECS instances in the VPC.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP address of the DNS server. You can enter at most four DNS server IP addresses. Separate IP addresses with commas (,).
	//
	// >  If you do not specify a DNS server IP address, Elastic Compute Service (ECS) instances use the IP addresses of the Alibaba Cloud DNS servers, which are 100.100.2.136 and 100.100.2.138.
	//
	// example:
	//
	// 192.XX.XX.123
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The lease time of the IPv6 addresses for the DHCP options set.
	//
	// 	- If you use hours as the unit, valid values are **24h to 1176h*	- and **87600h to 175200h**. Default value: **87600h**.
	//
	// 	- If you use days as the unit, valid values are **1d to 49d*	- and **3650d to 7300d**. Default value: **3650d**.
	//
	// >  If you specify a value, you must also specify the unit.
	//
	// example:
	//
	// 3650d
	Ipv6LeaseTime *string `json:"Ipv6LeaseTime,omitempty" xml:"Ipv6LeaseTime,omitempty"`
	// The lease time of the IPv4 addresses for the DHCP options set.
	//
	// 	- If you use hours as the unit, valid values are **24h to 1176h*	- and **87600h to 175200h**. Default value: **87600h**.
	//
	// 	- If you use days as the unit, valid values are **1d to 49d*	- and **3650d to 7300d**. Default value: **3650d**.
	//
	// >  If you specify a value, you must also specify the unit.
	//
	// example:
	//
	// 3650d
	LeaseTime    *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the DHCP options set is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s UpdateDhcpOptionsSetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDhcpOptionsSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDhcpOptionsSetDescription() *string {
	return s.DhcpOptionsSetDescription
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDhcpOptionsSetName() *string {
	return s.DhcpOptionsSetName
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDomainNameServers() *string {
	return s.DomainNameServers
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetIpv6LeaseTime() *string {
	return s.Ipv6LeaseTime
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetLeaseTime() *string {
	return s.LeaseTime
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateDhcpOptionsSetAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetClientToken(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDhcpOptionsSetDescription(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.DhcpOptionsSetDescription = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDhcpOptionsSetId(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDhcpOptionsSetName(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.DhcpOptionsSetName = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDomainName(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDomainNameServers(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.DomainNameServers = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetDryRun(v bool) *UpdateDhcpOptionsSetAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetIpv6LeaseTime(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.Ipv6LeaseTime = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetLeaseTime(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.LeaseTime = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetOwnerAccount(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetOwnerId(v int64) *UpdateDhcpOptionsSetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetRegionId(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetResourceOwnerAccount(v string) *UpdateDhcpOptionsSetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) SetResourceOwnerId(v int64) *UpdateDhcpOptionsSetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateDhcpOptionsSetAttributeRequest) Validate() error {
	return dara.Validate(s)
}
