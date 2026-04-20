// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListIpamDiscoveredIpAddressesResponseBody
	GetCount() *int32
	SetIpamDiscoveredIpAddresses(v []*ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) *ListIpamDiscoveredIpAddressesResponseBody
	GetIpamDiscoveredIpAddresses() []*ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses
	SetMaxResults(v int32) *ListIpamDiscoveredIpAddressesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamDiscoveredIpAddressesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamDiscoveredIpAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIpamDiscoveredIpAddressesResponseBody
	GetTotalCount() *int32
}

type ListIpamDiscoveredIpAddressesResponseBody struct {
	// example:
	//
	// 10
	Count                     *int32                                                                `json:"Count,omitempty" xml:"Count,omitempty"`
	IpamDiscoveredIpAddresses []*ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses `json:"IpamDiscoveredIpAddresses,omitempty" xml:"IpamDiscoveredIpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 9F8315CB-560E-5F1E-B069-6E44B440CAF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamDiscoveredIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetIpamDiscoveredIpAddresses() []*ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	return s.IpamDiscoveredIpAddresses
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetCount(v int32) *ListIpamDiscoveredIpAddressesResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetIpamDiscoveredIpAddresses(v []*ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) *ListIpamDiscoveredIpAddressesResponseBody {
	s.IpamDiscoveredIpAddresses = v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetMaxResults(v int32) *ListIpamDiscoveredIpAddressesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetNextToken(v string) *ListIpamDiscoveredIpAddressesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetRequestId(v string) *ListIpamDiscoveredIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) SetTotalCount(v int32) *ListIpamDiscoveredIpAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBody) Validate() error {
	if s.IpamDiscoveredIpAddresses != nil {
		for _, item := range s.IpamDiscoveredIpAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses struct {
	// example:
	//
	// 192.168.10.40/32
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// eni-bp1001jpjhzmgc5m****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// cn-shenzhen
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// ENI
	ResourceServiceType *string `json:"ResourceServiceType,omitempty" xml:"ResourceServiceType,omitempty"`
	// example:
	//
	// vsw-bp1bmwg5u07e1l3q0is4w
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1fjfnrg3av6zb9e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetResourceServiceType() *string {
	return s.ResourceServiceType
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetIpAddress(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.IpAddress = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetIpVersion(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.IpVersion = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetResourceId(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.ResourceId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetResourceRegionId(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetResourceServiceType(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.ResourceServiceType = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetVSwitchId(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.VSwitchId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) SetVpcId(v string) *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses {
	s.VpcId = &v
	return s
}

func (s *ListIpamDiscoveredIpAddressesResponseBodyIpamDiscoveredIpAddresses) Validate() error {
	return dara.Validate(s)
}
