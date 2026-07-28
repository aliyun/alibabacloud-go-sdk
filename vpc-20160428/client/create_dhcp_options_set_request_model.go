// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDhcpOptionsSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDhcpOptionsSetRequest
	GetClientToken() *string
	SetDhcpOptionsSetDescription(v string) *CreateDhcpOptionsSetRequest
	GetDhcpOptionsSetDescription() *string
	SetDhcpOptionsSetName(v string) *CreateDhcpOptionsSetRequest
	GetDhcpOptionsSetName() *string
	SetDomainName(v string) *CreateDhcpOptionsSetRequest
	GetDomainName() *string
	SetDomainNameServers(v string) *CreateDhcpOptionsSetRequest
	GetDomainNameServers() *string
	SetDryRun(v bool) *CreateDhcpOptionsSetRequest
	GetDryRun() *bool
	SetIpv6LeaseTime(v string) *CreateDhcpOptionsSetRequest
	GetIpv6LeaseTime() *string
	SetLeaseTime(v string) *CreateDhcpOptionsSetRequest
	GetLeaseTime() *string
	SetOwnerAccount(v string) *CreateDhcpOptionsSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDhcpOptionsSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDhcpOptionsSetRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDhcpOptionsSetRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDhcpOptionsSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDhcpOptionsSetRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateDhcpOptionsSetRequestTag) *CreateDhcpOptionsSetRequest
	GetTag() []*CreateDhcpOptionsSetRequestTag
}

type CreateDhcpOptionsSetRequest struct {
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
	// The description can be empty or 1 to 256 characters in length. It must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	DhcpOptionsSetDescription *string `json:"DhcpOptionsSetDescription,omitempty" xml:"DhcpOptionsSetDescription,omitempty"`
	// The name of the DHCP options set.
	//
	// The name must be 1 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// name
	DhcpOptionsSetName *string `json:"DhcpOptionsSetName,omitempty" xml:"DhcpOptionsSetName,omitempty"`
	// The hostname suffix. Example: example.com.
	//
	// After the DHCP options set is used to associate VPC, the hostname suffix is automatically synchronized to the ECS instances in the VPC.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of DNS servers. You can specify up to four DNS server IP addresses. Separate multiple IP addresses with commas (,).
	//
	// >If you do not specify DNS server IP addresses, ECS instances use the DNS server IP addresses provided by Alibaba Cloud (100.100.2.136 and 100.100.2.138) by default.
	//
	// example:
	//
	// 192.XX.XX.123
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the DHCP options set is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The lease time of the IPv6 DHCP options set.
	//
	// - If the lease time is in hours, the unit is h. Valid values: **24h to 1176h*	- and **87600h to 175200h**. Default value: **24h**.
	//
	// - If the lease time is in days, the unit is d. Valid values: **1d to 49d*	- and **3650d to 7300d**. Default value: **1d**.
	//
	// > You must include the unit when specifying the value.
	//
	// example:
	//
	// 3650d
	Ipv6LeaseTime *string `json:"Ipv6LeaseTime,omitempty" xml:"Ipv6LeaseTime,omitempty"`
	// The lease time of the IPv4 DHCP options set.
	//
	// - If the lease time is in hours, the unit is h. Valid values: **24h to 1176h*	- and **87600h to 175200h**. Default value: **87600h**.
	//
	// - If the lease time is in days, the unit is d. Valid values: **1d to 49d*	- and **3650d to 7300d**. Default value: **3650d**.
	//
	// > You must include the unit when specifying the value.
	//
	// example:
	//
	// 3650d
	LeaseTime    *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region in which the DHCP options set resides.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the DHCP options set belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateDhcpOptionsSetRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDhcpOptionsSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDhcpOptionsSetRequest) GoString() string {
	return s.String()
}

func (s *CreateDhcpOptionsSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDhcpOptionsSetRequest) GetDhcpOptionsSetDescription() *string {
	return s.DhcpOptionsSetDescription
}

func (s *CreateDhcpOptionsSetRequest) GetDhcpOptionsSetName() *string {
	return s.DhcpOptionsSetName
}

func (s *CreateDhcpOptionsSetRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDhcpOptionsSetRequest) GetDomainNameServers() *string {
	return s.DomainNameServers
}

func (s *CreateDhcpOptionsSetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDhcpOptionsSetRequest) GetIpv6LeaseTime() *string {
	return s.Ipv6LeaseTime
}

func (s *CreateDhcpOptionsSetRequest) GetLeaseTime() *string {
	return s.LeaseTime
}

func (s *CreateDhcpOptionsSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDhcpOptionsSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDhcpOptionsSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDhcpOptionsSetRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDhcpOptionsSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDhcpOptionsSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDhcpOptionsSetRequest) GetTag() []*CreateDhcpOptionsSetRequestTag {
	return s.Tag
}

func (s *CreateDhcpOptionsSetRequest) SetClientToken(v string) *CreateDhcpOptionsSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetDhcpOptionsSetDescription(v string) *CreateDhcpOptionsSetRequest {
	s.DhcpOptionsSetDescription = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetDhcpOptionsSetName(v string) *CreateDhcpOptionsSetRequest {
	s.DhcpOptionsSetName = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetDomainName(v string) *CreateDhcpOptionsSetRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetDomainNameServers(v string) *CreateDhcpOptionsSetRequest {
	s.DomainNameServers = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetDryRun(v bool) *CreateDhcpOptionsSetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetIpv6LeaseTime(v string) *CreateDhcpOptionsSetRequest {
	s.Ipv6LeaseTime = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetLeaseTime(v string) *CreateDhcpOptionsSetRequest {
	s.LeaseTime = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetOwnerAccount(v string) *CreateDhcpOptionsSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetOwnerId(v int64) *CreateDhcpOptionsSetRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetRegionId(v string) *CreateDhcpOptionsSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetResourceGroupId(v string) *CreateDhcpOptionsSetRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetResourceOwnerAccount(v string) *CreateDhcpOptionsSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetResourceOwnerId(v int64) *CreateDhcpOptionsSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDhcpOptionsSetRequest) SetTag(v []*CreateDhcpOptionsSetRequestTag) *CreateDhcpOptionsSetRequest {
	s.Tag = v
	return s
}

func (s *CreateDhcpOptionsSetRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDhcpOptionsSetRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDhcpOptionsSetRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDhcpOptionsSetRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDhcpOptionsSetRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDhcpOptionsSetRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDhcpOptionsSetRequestTag) SetKey(v string) *CreateDhcpOptionsSetRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDhcpOptionsSetRequestTag) SetValue(v string) *CreateDhcpOptionsSetRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDhcpOptionsSetRequestTag) Validate() error {
	return dara.Validate(s)
}
