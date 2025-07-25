// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListIpamDiscoveredResourceResponseBody
	GetCount() *int32
	SetIpamDiscoveredResources(v []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) *ListIpamDiscoveredResourceResponseBody
	GetIpamDiscoveredResources() []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources
	SetMaxResults(v int32) *ListIpamDiscoveredResourceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamDiscoveredResourceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamDiscoveredResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamDiscoveredResourceResponseBody
	GetTotalCount() *int64
}

type ListIpamDiscoveredResourceResponseBody struct {
	// The maximum number of entries on each page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of resources.
	IpamDiscoveredResources []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources `json:"IpamDiscoveredResources,omitempty" xml:"IpamDiscoveredResources,omitempty" type:"Repeated"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, there is no next page.
	//
	// 	- If a value of **NextToken*	- is returned, it indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3748DEFF-68BE-5EED-9937-7C1D0C21BAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListIpamDiscoveredResourceResponseBody) GetIpamDiscoveredResources() []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	return s.IpamDiscoveredResources
}

func (s *ListIpamDiscoveredResourceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamDiscoveredResourceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamDiscoveredResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamDiscoveredResourceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamDiscoveredResourceResponseBody) SetCount(v int32) *ListIpamDiscoveredResourceResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetIpamDiscoveredResources(v []*ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) *ListIpamDiscoveredResourceResponseBody {
	s.IpamDiscoveredResources = v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetMaxResults(v int32) *ListIpamDiscoveredResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetNextToken(v string) *ListIpamDiscoveredResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetRequestId(v string) *ListIpamDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) SetTotalCount(v int64) *ListIpamDiscoveredResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 132193271328****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The CIDR block of the resource.
	//
	// example:
	//
	// 192.168.1.0/32
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the resource was discovered.
	//
	// >  If the resource has not been modified since it was created, the discovery time remains unchanged.
	//
	// example:
	//
	// 2024-01-01 00:00:00
	DiscoveryTime *string                                                                     `json:"DiscoveryTime,omitempty" xml:"DiscoveryTime,omitempty"`
	IpCountDetail *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail `json:"IpCountDetail,omitempty" xml:"IpCountDetail,omitempty" type:"Struct"`
	// The IP usage in decimal form.
	//
	// example:
	//
	// 0
	IpUsage *string `json:"IpUsage,omitempty" xml:"IpUsage,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// vpc-uf611fp465c7dyb4z****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the region to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VSwitch**
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The ID of the VPC to which the resource belongs.
	//
	// example:
	//
	// vpc-uf611fp465c7dyb4z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetDiscoveryTime() *string {
	return s.DiscoveryTime
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetIpCountDetail() *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	return s.IpCountDetail
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetIpUsage() *string {
	return s.IpUsage
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetAliUid(v int64) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.AliUid = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetCidr(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.Cidr = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetDiscoveryTime(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.DiscoveryTime = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpCountDetail(v *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpCountDetail = v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpUsage(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpUsage = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceOwnerId(v int64) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceRegionId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetResourceType(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.ResourceType = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetSourceCidr(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) SetVpcId(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources {
	s.VpcId = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResources) Validate() error {
	return dara.Validate(s)
}

type ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail struct {
	FreeIpCount  *string `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	TotalIpCount *string `json:"TotalIpCount,omitempty" xml:"TotalIpCount,omitempty"`
	UsedIpCount  *string `json:"UsedIpCount,omitempty" xml:"UsedIpCount,omitempty"`
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) GetFreeIpCount() *string {
	return s.FreeIpCount
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) GetTotalIpCount() *string {
	return s.TotalIpCount
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) GetUsedIpCount() *string {
	return s.UsedIpCount
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetFreeIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.FreeIpCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetTotalIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.TotalIpCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) SetUsedIpCount(v string) *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail {
	s.UsedIpCount = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponseBodyIpamDiscoveredResourcesIpCountDetail) Validate() error {
	return dara.Validate(s)
}
