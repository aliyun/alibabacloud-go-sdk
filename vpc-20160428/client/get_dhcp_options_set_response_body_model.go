// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDhcpOptionsSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociateVpcs(v []*GetDhcpOptionsSetResponseBodyAssociateVpcs) *GetDhcpOptionsSetResponseBody
	GetAssociateVpcs() []*GetDhcpOptionsSetResponseBodyAssociateVpcs
	SetCreationTime(v string) *GetDhcpOptionsSetResponseBody
	GetCreationTime() *string
	SetDhcpOptions(v *GetDhcpOptionsSetResponseBodyDhcpOptions) *GetDhcpOptionsSetResponseBody
	GetDhcpOptions() *GetDhcpOptionsSetResponseBodyDhcpOptions
	SetDhcpOptionsSetDescription(v string) *GetDhcpOptionsSetResponseBody
	GetDhcpOptionsSetDescription() *string
	SetDhcpOptionsSetId(v string) *GetDhcpOptionsSetResponseBody
	GetDhcpOptionsSetId() *string
	SetDhcpOptionsSetName(v string) *GetDhcpOptionsSetResponseBody
	GetDhcpOptionsSetName() *string
	SetOwnerId(v int64) *GetDhcpOptionsSetResponseBody
	GetOwnerId() *int64
	SetRequestId(v string) *GetDhcpOptionsSetResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetDhcpOptionsSetResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetDhcpOptionsSetResponseBody
	GetStatus() *string
	SetTags(v []*GetDhcpOptionsSetResponseBodyTags) *GetDhcpOptionsSetResponseBody
	GetTags() []*GetDhcpOptionsSetResponseBodyTags
}

type GetDhcpOptionsSetResponseBody struct {
	// The information about the virtual private cloud (VPC) that is associated with the DHCP options set.
	AssociateVpcs []*GetDhcpOptionsSetResponseBodyAssociateVpcs `json:"AssociateVpcs,omitempty" xml:"AssociateVpcs,omitempty" type:"Repeated"`
	CreationTime  *string                                       `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The configuration information about the DHCP options set.
	DhcpOptions *GetDhcpOptionsSetResponseBodyDhcpOptions `json:"DhcpOptions,omitempty" xml:"DhcpOptions,omitempty" type:"Struct"`
	// The description of the DHCP options set.
	//
	// example:
	//
	// test
	DhcpOptionsSetDescription *string `json:"DhcpOptionsSetDescription,omitempty" xml:"DhcpOptionsSetDescription,omitempty"`
	// The ID of the DHCP options set.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// The name of the DHCP options set.
	//
	// example:
	//
	// test
	DhcpOptionsSetName *string `json:"DhcpOptionsSetName,omitempty" xml:"DhcpOptionsSetName,omitempty"`
	// The ID of the Alibaba Cloud account to which the DHCP options set belongs.
	//
	// example:
	//
	// 283117732402483989
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the DHCP options set. Valid values:
	//
	// 	- **Available**: available
	//
	// 	- **InUse**: in use
	//
	// 	- **Deleted**: deleted
	//
	// 	- **Pending**: being configured
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*GetDhcpOptionsSetResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetDhcpOptionsSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetResponseBody) GetAssociateVpcs() []*GetDhcpOptionsSetResponseBodyAssociateVpcs {
	return s.AssociateVpcs
}

func (s *GetDhcpOptionsSetResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetDhcpOptionsSetResponseBody) GetDhcpOptions() *GetDhcpOptionsSetResponseBodyDhcpOptions {
	return s.DhcpOptions
}

func (s *GetDhcpOptionsSetResponseBody) GetDhcpOptionsSetDescription() *string {
	return s.DhcpOptionsSetDescription
}

func (s *GetDhcpOptionsSetResponseBody) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *GetDhcpOptionsSetResponseBody) GetDhcpOptionsSetName() *string {
	return s.DhcpOptionsSetName
}

func (s *GetDhcpOptionsSetResponseBody) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetDhcpOptionsSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDhcpOptionsSetResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetDhcpOptionsSetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDhcpOptionsSetResponseBody) GetTags() []*GetDhcpOptionsSetResponseBodyTags {
	return s.Tags
}

func (s *GetDhcpOptionsSetResponseBody) SetAssociateVpcs(v []*GetDhcpOptionsSetResponseBodyAssociateVpcs) *GetDhcpOptionsSetResponseBody {
	s.AssociateVpcs = v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetCreationTime(v string) *GetDhcpOptionsSetResponseBody {
	s.CreationTime = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetDhcpOptions(v *GetDhcpOptionsSetResponseBodyDhcpOptions) *GetDhcpOptionsSetResponseBody {
	s.DhcpOptions = v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetDhcpOptionsSetDescription(v string) *GetDhcpOptionsSetResponseBody {
	s.DhcpOptionsSetDescription = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetDhcpOptionsSetId(v string) *GetDhcpOptionsSetResponseBody {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetDhcpOptionsSetName(v string) *GetDhcpOptionsSetResponseBody {
	s.DhcpOptionsSetName = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetOwnerId(v int64) *GetDhcpOptionsSetResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetRequestId(v string) *GetDhcpOptionsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetResourceGroupId(v string) *GetDhcpOptionsSetResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetStatus(v string) *GetDhcpOptionsSetResponseBody {
	s.Status = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) SetTags(v []*GetDhcpOptionsSetResponseBodyTags) *GetDhcpOptionsSetResponseBody {
	s.Tags = v
	return s
}

func (s *GetDhcpOptionsSetResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDhcpOptionsSetResponseBodyAssociateVpcs struct {
	// The status of the VPC that is associated with the DHCP options set. Valid values:
	//
	// 	- **InUse**: in use
	//
	// 	- **Pending**: being configured
	//
	// example:
	//
	// InUse
	AssociateStatus *string `json:"AssociateStatus,omitempty" xml:"AssociateStatus,omitempty"`
	// The ID of the VPC that is associated with the DHCP options set.
	//
	// example:
	//
	// vpc-eb3b54r6otues4tjj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetDhcpOptionsSetResponseBodyAssociateVpcs) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetResponseBodyAssociateVpcs) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetResponseBodyAssociateVpcs) GetAssociateStatus() *string {
	return s.AssociateStatus
}

func (s *GetDhcpOptionsSetResponseBodyAssociateVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetDhcpOptionsSetResponseBodyAssociateVpcs) SetAssociateStatus(v string) *GetDhcpOptionsSetResponseBodyAssociateVpcs {
	s.AssociateStatus = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyAssociateVpcs) SetVpcId(v string) *GetDhcpOptionsSetResponseBodyAssociateVpcs {
	s.VpcId = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyAssociateVpcs) Validate() error {
	return dara.Validate(s)
}

type GetDhcpOptionsSetResponseBodyDhcpOptions struct {
	// The suffix of the hostname.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP address of the DNS server.
	//
	// example:
	//
	// 192.XX.XX.123
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// The lease time of the IPv6 addresses for the DHCP options set.
	//
	// 	- If you use hours as the unit, Valid values are **24h to 1176h*	- and **87600h to 175200h**. Default value: **87600h**.
	//
	// 	- If you use days as the unit, Valid values are **1d to 49d*	- and **3650d to 7300d**. Default value: **3650d**.
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
	// example:
	//
	// 3650d
	LeaseTime *string `json:"LeaseTime,omitempty" xml:"LeaseTime,omitempty"`
}

func (s GetDhcpOptionsSetResponseBodyDhcpOptions) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetResponseBodyDhcpOptions) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) GetDomainNameServers() *string {
	return s.DomainNameServers
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) GetIpv6LeaseTime() *string {
	return s.Ipv6LeaseTime
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) GetLeaseTime() *string {
	return s.LeaseTime
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) SetDomainName(v string) *GetDhcpOptionsSetResponseBodyDhcpOptions {
	s.DomainName = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) SetDomainNameServers(v string) *GetDhcpOptionsSetResponseBodyDhcpOptions {
	s.DomainNameServers = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) SetIpv6LeaseTime(v string) *GetDhcpOptionsSetResponseBodyDhcpOptions {
	s.Ipv6LeaseTime = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) SetLeaseTime(v string) *GetDhcpOptionsSetResponseBodyDhcpOptions {
	s.LeaseTime = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyDhcpOptions) Validate() error {
	return dara.Validate(s)
}

type GetDhcpOptionsSetResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDhcpOptionsSetResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetDhcpOptionsSetResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetDhcpOptionsSetResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetDhcpOptionsSetResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetDhcpOptionsSetResponseBodyTags) SetKey(v string) *GetDhcpOptionsSetResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyTags) SetValue(v string) *GetDhcpOptionsSetResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetDhcpOptionsSetResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
