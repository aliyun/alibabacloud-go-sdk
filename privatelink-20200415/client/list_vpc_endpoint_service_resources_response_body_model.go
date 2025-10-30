// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServiceResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServiceResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpcEndpointServiceResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*ListVpcEndpointServiceResourcesResponseBodyResources) *ListVpcEndpointServiceResourcesResponseBody
	GetResources() []*ListVpcEndpointServiceResourcesResponseBodyResources
}

type ListVpcEndpointServiceResourcesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next requests are performed.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service resources.
	Resources []*ListVpcEndpointServiceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListVpcEndpointServiceResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServiceResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServiceResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpcEndpointServiceResourcesResponseBody) GetResources() []*ListVpcEndpointServiceResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetMaxResults(v int32) *ListVpcEndpointServiceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetNextToken(v string) *ListVpcEndpointServiceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetRequestId(v string) *ListVpcEndpointServiceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) SetResources(v []*ListVpcEndpointServiceResourcesResponseBodyResources) *ListVpcEndpointServiceResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpcEndpointServiceResourcesResponseBodyResources struct {
	// Indicates whether automatic resource allocation is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoAllocatedEnabled *bool `json:"AutoAllocatedEnabled,omitempty" xml:"AutoAllocatedEnabled,omitempty"`
	// The IP address of the service resource.
	//
	// example:
	//
	// 192.168.10.23
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region where the service resource is deployed.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of endpoints that are associated with the service resource that is smoothly migrated.
	//
	// example:
	//
	// 10
	RelatedDeprecatedEndpointCount *int64 `json:"RelatedDeprecatedEndpointCount,omitempty" xml:"RelatedDeprecatedEndpointCount,omitempty"`
	// The number of endpoints that are associated with the service resource.
	//
	// example:
	//
	// 10
	RelatedEndpointCount *int64 `json:"RelatedEndpointCount,omitempty" xml:"RelatedEndpointCount,omitempty"`
	// The service resource ID.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Indicates whether IPv6 is enabled for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ResourceSupportIPv6 *bool `json:"ResourceSupportIPv6,omitempty" xml:"ResourceSupportIPv6,omitempty"`
	// The type of the service resource.
	//
	// Only **slb*	- is returned. This value indicates a Classic Load Balancer (CLB) instance.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the vSwitch to which the service resource belongs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the service resource belongs.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the service resource belongs.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetAutoAllocatedEnabled() *bool {
	return s.AutoAllocatedEnabled
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetIp() *string {
	return s.Ip
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetRelatedDeprecatedEndpointCount() *int64 {
	return s.RelatedDeprecatedEndpointCount
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetRelatedEndpointCount() *int64 {
	return s.RelatedEndpointCount
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetResourceSupportIPv6() *bool {
	return s.ResourceSupportIPv6
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetAutoAllocatedEnabled(v bool) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.AutoAllocatedEnabled = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetIp(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.Ip = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRegionId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRelatedDeprecatedEndpointCount(v int64) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RelatedDeprecatedEndpointCount = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetRelatedEndpointCount(v int64) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.RelatedEndpointCount = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceSupportIPv6(v bool) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceSupportIPv6 = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetResourceType(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetVSwitchId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetVpcId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) SetZoneId(v string) *ListVpcEndpointServiceResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
