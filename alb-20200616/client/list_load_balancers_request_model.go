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
	SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersRequest
	GetLoadBalancerBussinessStatus() *string
	SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest
	GetLoadBalancerIds() []*string
	SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest
	GetLoadBalancerNames() []*string
	SetLoadBalancerStatus(v string) *ListLoadBalancersRequest
	GetLoadBalancerStatus() *string
	SetMaxResults(v int32) *ListLoadBalancersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLoadBalancersRequest
	GetNextToken() *string
	SetPayType(v string) *ListLoadBalancersRequest
	GetPayType() *string
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
	// The IP version. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the VPC where the ALB instance is deployed.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The domain name.
	//
	// example:
	//
	// alb-95qnr2itwu9orb****.cn-hangzhou.alb.aliyuncs.com
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The type of IPv6 address that is used by the ALB instance. Valid values:
	//
	// 	- **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. Therefore, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	//
	// example:
	//
	// Intranet
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The service status of the ALB instance. Valid values:
	//
	// 	- **Abnormal**
	//
	// 	- **Normal**
	//
	// example:
	//
	// Normal
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The instance IDs. You can specify at most 20 ALB instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The instance names. You can specify at most 10 instance names.
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// The status of the ALB instance. Valid values:
	//
	// 	- **Inactive**: The ALB instance is disabled. The listeners do not forward traffic.
	//
	// 	- **Active**: The ALB instance is running.
	//
	// 	- **Provisioning**: The ALB instance is being created.
	//
	// 	- **Configuring**: The ALB instance is being modified.
	//
	// 	- **CreateFailed**: The system failed to create the ALB instance. In this case, you are not charged for the ALB instance. You can only delete the ALB instance. By default, the system deletes the ALB instances that are in the CreateFailed state within the last day.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The billing method of the ALB instance. Set the value to
	//
	// **PostPay**, which specifies the pay-as-you-go billing method. This is the default value.
	//
	// example:
	//
	// PostPay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the ALB instance.
	Tag []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the ALB instance belongs. You can specify at most 10 VPC IDs.
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// The ID of the zone where the ALB instance is deployed.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/189196.html) operation to query zones.
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

func (s *ListLoadBalancersRequest) GetLoadBalancerBussinessStatus() *string {
	return s.LoadBalancerBussinessStatus
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

func (s *ListLoadBalancersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLoadBalancersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLoadBalancersRequest) GetPayType() *string {
	return s.PayType
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

func (s *ListLoadBalancersRequest) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBussinessStatus = &v
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

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetPayType(v string) *ListLoadBalancersRequest {
	s.PayType = &v
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
	return dara.Validate(s)
}

type ListLoadBalancersRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// alueTest
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
