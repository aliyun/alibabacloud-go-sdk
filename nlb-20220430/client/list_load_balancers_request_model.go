// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *ListLoadBalancersRequest
	GetAddressIpVersion() *string
	SetAddressType(v string) *ListLoadBalancersRequest
	GetAddressType() *string
	SetDNSName(v string) *ListLoadBalancersRequest
	GetDNSName() *string
	SetIpv6AddressType(v string) *ListLoadBalancersRequest
	GetIpv6AddressType() *string
	SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest
	GetLoadBalancerBusinessStatus() *string
	SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest
	GetLoadBalancerIds() []*string
	SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest
	GetLoadBalancerNames() []*string
	SetLoadBalancerStatus(v string) *ListLoadBalancersRequest
	GetLoadBalancerStatus() *string
	SetLoadBalancerType(v string) *ListLoadBalancersRequest
	GetLoadBalancerType() *string
	SetMaxResults(v int32) *ListLoadBalancersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLoadBalancersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListLoadBalancersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListLoadBalancersRequest
	GetResourceGroupId() *string
	SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest
	GetTag() []*ListLoadBalancersRequestTag
	SetVpcIds(v []*string) *ListLoadBalancersRequest
	GetVpcIds() []*string
	SetZoneId(v string) *ListLoadBalancersRequest
	GetZoneId() *string
}

type ListLoadBalancersRequest struct {
	// The protocol version. Valid values:
	//
	// - **ipv4**: IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The IPv4 address type of the NLB instance. Valid values:
	//
	// - **Internet**: The load balancing instance has a public IP address. The DNS domain name is resolved to the public IP address. Therefore, the instance is accessible over the Internet.
	//
	// - **Intranet**: The load balancing instance has only a private IP address. The DNS domain name is resolved to the private IP address. Therefore, the instance is accessible only within the internal network environment of the VPC where the load balancing instance is deployed.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The DNS domain name.
	//
	// example:
	//
	// nlb-wb7r6dlwetvt5j****.cn-hangzhou.nlb.aliyuncs.com
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The IPv6 address type of the NLB instance. Valid values:
	//
	// - **Internet**: The load balancing instance has a public IP address. The DNS domain name is resolved to the public IP address. Therefore, the instance is accessible over the Internet.
	//
	// - **Intranet**: The load balancing instance has only a private IP address. The DNS domain name is resolved to the private IP address. Therefore, the instance is accessible only within the internal network environment of the VPC where the load balancing instance is deployed.
	//
	// example:
	//
	// Internet
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The business status of the NLB instance. Valid values:
	//
	// - **Abnormal**: The instance is in an abnormal state.
	//
	// - **Normal**: The instance is in a normal state.
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The IDs of the NLB instances. You can query up to 20 instances at a time.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The names of the NLB instances. You can query up to 20 instance names at a time.
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// The status of the NLB instance. Valid values:
	//
	// - **Inactive**: The instance is stopped. Listeners of the instance no longer forward traffic.
	//
	// - **Active**: The instance is running.
	//
	// - **Provisioning**: The instance is being created.
	//
	// - **Configuring**: Upgrade/Downgrade in progress.
	//
	// - **Deleting**: The instance is being deleted.
	//
	// - **Deleted**: The instance is deleted.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The type of load balancing. Only **network*	- is returned, which indicates Network Load Balancer (NLB).
	//
	// example:
	//
	// network
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - If this is the first request or no next request exists, you do not need to specify this parameter.
	//
	// - If a next request exists, set the value to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the NLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The labels of the NLB instance.
	Tag []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID of the NLB instance. You can query up to 10 VPC IDs at a time.
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// The zone name. You can call the [DescribeZones](https://help.aliyun.com/document_detail/443890.html) operation to query zones.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListLoadBalancersRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *ListLoadBalancersRequest) GetDNSName() *string {
	return s.DNSName
}

func (s *ListLoadBalancersRequest) GetIpv6AddressType() *string {
	return s.Ipv6AddressType
}

func (s *ListLoadBalancersRequest) GetLoadBalancerBusinessStatus() *string {
	return s.LoadBalancerBusinessStatus
}

func (s *ListLoadBalancersRequest) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *ListLoadBalancersRequest) GetLoadBalancerNames() []*string {
	return s.LoadBalancerNames
}

func (s *ListLoadBalancersRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *ListLoadBalancersRequest) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *ListLoadBalancersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLoadBalancersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLoadBalancersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLoadBalancersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLoadBalancersRequest) GetTag() []*ListLoadBalancersRequestTag {
	return s.Tag
}

func (s *ListLoadBalancersRequest) GetVpcIds() []*string {
	return s.VpcIds
}

func (s *ListLoadBalancersRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListLoadBalancersRequest) SetAddressIpVersion(v string) *ListLoadBalancersRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersRequest) SetAddressType(v string) *ListLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetDNSName(v string) *ListLoadBalancersRequest {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersRequest) SetIpv6AddressType(v string) *ListLoadBalancersRequest {
	s.Ipv6AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerNames = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerType(v string) *ListLoadBalancersRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetRegionId(v string) *ListLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneId(v string) *ListLoadBalancersRequest {
	s.ZoneId = &v
	return s
}

func (s *ListLoadBalancersRequest) Validate() error {
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

type ListLoadBalancersRequestTag struct {
	// The tag key of the instance. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListLoadBalancersRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListLoadBalancersRequestTag) SetKey(v string) *ListLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersRequestTag) SetValue(v string) *ListLoadBalancersRequestTag {
	s.Value = &v
	return s
}

func (s *ListLoadBalancersRequestTag) Validate() error {
	return dara.Validate(s)
}
