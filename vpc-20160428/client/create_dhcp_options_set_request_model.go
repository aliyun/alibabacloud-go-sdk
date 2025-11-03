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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the DHCP options set.
	//
	// The description must be 1 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	DhcpOptionsSetDescription *string `json:"DhcpOptionsSetDescription,omitempty" xml:"DhcpOptionsSetDescription,omitempty"`
	// The name of the DHCP options set.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
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
	// >  If no IP address is specified, the Elastic Compute Service (ECS) instance uses the IP addresses 100.100.2.136 and 100.100.2.138, which are provided by Alibaba Cloud by default.
	//
	// example:
	//
	// 192.XX.XX.123
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// >  When you enter a value, you must also specify the unit.
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
	// >  When you enter a value, you must also specify the unit.
	//
	// example:
	//
	// 3650d
	LeaseTime    *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region to which the DHCP options set belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// The tag of the resource.
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
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be at most 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
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
