// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDhcpOptionsSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDhcpOptionsSets(v []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) *ListDhcpOptionsSetsResponseBody
	GetDhcpOptionsSets() []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSets
	SetNextToken(v string) *ListDhcpOptionsSetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDhcpOptionsSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListDhcpOptionsSetsResponseBody
	GetTotalCount() *string
}

type ListDhcpOptionsSetsResponseBody struct {
	// The list of the DHCP options sets.
	DhcpOptionsSets []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSets `json:"DhcpOptionsSets,omitempty" xml:"DhcpOptionsSets,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd********
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries.
	//
	// example:
	//
	// 10
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDhcpOptionsSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsResponseBody) GetDhcpOptionsSets() []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	return s.DhcpOptionsSets
}

func (s *ListDhcpOptionsSetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDhcpOptionsSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDhcpOptionsSetsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListDhcpOptionsSetsResponseBody) SetDhcpOptionsSets(v []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) *ListDhcpOptionsSetsResponseBody {
	s.DhcpOptionsSets = v
	return s
}

func (s *ListDhcpOptionsSetsResponseBody) SetNextToken(v string) *ListDhcpOptionsSetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBody) SetRequestId(v string) *ListDhcpOptionsSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBody) SetTotalCount(v string) *ListDhcpOptionsSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBody) Validate() error {
	if s.DhcpOptionsSets != nil {
		for _, item := range s.DhcpOptionsSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDhcpOptionsSetsResponseBodyDhcpOptionsSets struct {
	// The number of VPCs with which the DHCP options set is associated.
	//
	// example:
	//
	// 2
	AssociateVpcCount *int32 `json:"AssociateVpcCount,omitempty" xml:"AssociateVpcCount,omitempty"`
	// The creation time of the DHCP options sets.
	//
	// example:
	//
	// 2025-08-21 ***
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The details of DHCP options.
	DhcpOptions *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions `json:"DhcpOptions,omitempty" xml:"DhcpOptions,omitempty" type:"Struct"`
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
	// 253460731706911258
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the DHCP options set belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the DHCP options set. Valid values:
	//
	// 	- **Available**
	//
	// 	- **InUse**
	//
	// 	- **Pending**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetAssociateVpcCount() *int32 {
	return s.AssociateVpcCount
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetDhcpOptions() *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions {
	return s.DhcpOptions
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetDhcpOptionsSetDescription() *string {
	return s.DhcpOptionsSetDescription
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetDhcpOptionsSetName() *string {
	return s.DhcpOptionsSetName
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetStatus() *string {
	return s.Status
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) GetTags() []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags {
	return s.Tags
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetAssociateVpcCount(v int32) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.AssociateVpcCount = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetCreationTime(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.CreationTime = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetDhcpOptions(v *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.DhcpOptions = v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetDhcpOptionsSetDescription(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.DhcpOptionsSetDescription = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetDhcpOptionsSetId(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetDhcpOptionsSetName(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.DhcpOptionsSetName = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetOwnerId(v int64) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.OwnerId = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetResourceGroupId(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetStatus(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.Status = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) SetTags(v []*ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets {
	s.Tags = v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSets) Validate() error {
	if s.DhcpOptions != nil {
		if err := s.DhcpOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions struct {
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
	// 192.168.XX.XX
	DomainNameServers *string `json:"DomainNameServers,omitempty" xml:"DomainNameServers,omitempty"`
	// The lease time of the IPv6 DHCP options set.
	//
	// 	- If you use hours as the unit, Unit: h. Valid values are **24h to 1176h*	- and **87600h to 175200h**. Default value: **24h**.
	//
	// 	- If you use days as the unit, Unit: d. Valid values are **1d to 49d*	- and **3650d to 7300d**. Default value: **1d**.
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

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) GetDomainNameServers() *string {
	return s.DomainNameServers
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) GetIpv6LeaseTime() *string {
	return s.Ipv6LeaseTime
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) GetLeaseTime() *string {
	return s.LeaseTime
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) SetDomainName(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions {
	s.DomainName = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) SetDomainNameServers(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions {
	s.DomainNameServers = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) SetIpv6LeaseTime(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions {
	s.Ipv6LeaseTime = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) SetLeaseTime(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions {
	s.LeaseTime = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsDhcpOptions) Validate() error {
	return dara.Validate(s)
}

type ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) GetKey() *string {
	return s.Key
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) GetValue() *string {
	return s.Value
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) SetKey(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags {
	s.Key = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) SetValue(v string) *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags {
	s.Value = &v
	return s
}

func (s *ListDhcpOptionsSetsResponseBodyDhcpOptionsSetsTags) Validate() error {
	return dara.Validate(s)
}
