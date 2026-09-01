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
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the DHCP options set.
	//
	// The description can be empty or 2 to 256 characters in length. It must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	DhcpOptionsSetDescription *string `json:"DhcpOptionsSetDescription,omitempty" xml:"DhcpOptionsSetDescription,omitempty"`
	// The ID of the DHCP options set to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// The name of the DHCP options set.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// name
	DhcpOptionsSetName *string `json:"DhcpOptionsSetName,omitempty" xml:"DhcpOptionsSetName,omitempty"`
	// The hostname suffix, such as example.com.
	//
	// After you attach the DHCP options set to associate VPC, the hostname suffix is automatically subject to synchronization to ECS instances in the VPC.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of DNS servers. You can specify up to four DNS server IP addresses, separated by commas (,).
	//
	// >If you do not specify any DNS server IP addresses, ECS instances use the DNS server IP addresses provided by Alibaba Cloud (100.100.2.136 and 100.100.2.138) by default.
	//
	// example:
	//
	// 192.XX.XX.123
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the DHCP options set configuration is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The lease time of the IPv6 DHCP options set.
	//
	// - When the lease time is set in hours: Unit: h. Valid values: **24h to 1176h*	- and **87600h to 175200h**. Default value: **24h**.
	//
	// - When the lease time is set in days: Unit: d. Valid values: **1d to 49d*	- and **3650d to 7300d**. Default value: **1d**.
	//
	// > You must include the unit when specifying the value.
	//
	// example:
	//
	// 3650d
	Ipv6LeaseTime *string `json:"Ipv6LeaseTime,omitempty" xml:"Ipv6LeaseTime,omitempty"`
	// The lease time of the IPv4 DHCP options set.
	//
	// - When the lease time is set in hours: Unit: h. Valid values: **24h to 1176h*	- and **87600h to 175200h**. Default value: **87600h**.
	//
	// - When the lease time is set in days: Unit: d. Valid values: **1d to 49d*	- and **3650d to 7300d**. Default value: **3650d**.
	//
	// > You must include the unit when specifying the value.
	//
	// example:
	//
	// 3650d
	LeaseTime    *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the DHCP options set to modify. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
