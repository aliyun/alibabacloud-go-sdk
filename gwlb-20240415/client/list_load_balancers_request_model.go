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
	SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest
	GetLoadBalancerBusinessStatus() *string
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
	SetResourceGroupId(v string) *ListLoadBalancersRequest
	GetResourceGroupId() *string
	SetSkip(v int32) *ListLoadBalancersRequest
	GetSkip() *int32
	SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest
	GetTag() []*ListLoadBalancersRequestTag
	SetTrafficMode(v string) *ListLoadBalancersRequest
	GetTrafficMode() *string
	SetVpcIds(v []*string) *ListLoadBalancersRequest
	GetVpcIds() []*string
	SetZoneIds(v []*string) *ListLoadBalancersRequest
	GetZoneIds() []*string
}

type ListLoadBalancersRequest struct {
	// The IP version of the NLB instance. Valid values:
	//
	// 	- **Ipv4**
	//
	// Enumeration values:
	//
	// 	- IPv4: IPv4
	//
	// 	- DualStack: dual-stack
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The business status of the GWLB instance. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The GWLB instance IDs. You can query at most 20 GWLB instances in each call.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The GWLB instance names. You can specify at most 20 GWLB instance names in each call.
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// The GWLB instance status. Valid values:
	//
	// 	- **Active**: The GWLB instance is running.
	//
	// 	- **Inactive**: The GWLB instance is disabled. Listeners of GWLB instances in the Inactive state do not forward traffic.
	//
	// 	- **Provisioning**: The GWLB instance is being created.
	//
	// 	- **Configuring**: The GWLB instance is being modified.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// WyJyb290IiwibiIsIm4iLDEsMCwxNjg1MDY1NTgyNzYwLCI2NDcwMGY2ZTc2Zjc0MWFiZGEyZjQyNzc4ZDk2MmJkOTk3ZGZmM2Nm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2htf5qsyrn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tags. You can specify up to 20 tags in each call.
	Tag []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies the traffic processing mode. Valid values:
	//
	// 	- **LoadBalance**: load balancing mode. GWLB continues to forward traffic to backend servers.
	//
	// 	- **ByPass**: bypass mode. GWLB directly returns traffic to the GWLB endpoint without forwarding it to the backend servers.
	//
	// example:
	//
	// LoadBalance
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The virtual private cloud (VPC) IDs. You can query at most 20 IDs in each call.
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// The zone IDs. You can query at most 20 zone IDs in each call.
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
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

func (s *ListLoadBalancersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLoadBalancersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLoadBalancersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLoadBalancersRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListLoadBalancersRequest) GetTag() []*ListLoadBalancersRequestTag {
	return s.Tag
}

func (s *ListLoadBalancersRequest) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *ListLoadBalancersRequest) GetVpcIds() []*string {
	return s.VpcIds
}

func (s *ListLoadBalancersRequest) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *ListLoadBalancersRequest) SetAddressIpVersion(v string) *ListLoadBalancersRequest {
	s.AddressIpVersion = &v
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

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetSkip(v int32) *ListLoadBalancersRequest {
	s.Skip = &v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetTrafficMode(v string) *ListLoadBalancersRequest {
	s.TrafficMode = &v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneIds(v []*string) *ListLoadBalancersRequest {
	s.ZoneIds = v
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
	// The tag key. You cannot specify an empty string as a tag key.
	//
	// The tag key can be up to 128 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. It can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
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
