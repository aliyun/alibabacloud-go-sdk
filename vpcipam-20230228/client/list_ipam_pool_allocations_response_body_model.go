// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolAllocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamPoolAllocationsResponseBody
	GetCount() *int64
	SetIpamPoolAllocations(v []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) *ListIpamPoolAllocationsResponseBody
	GetIpamPoolAllocations() []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations
	SetMaxResults(v int64) *ListIpamPoolAllocationsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamPoolAllocationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamPoolAllocationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamPoolAllocationsResponseBody
	GetTotalCount() *int64
}

type ListIpamPoolAllocationsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IDs of the instances to which CIDR blocks are allocated from the IPAM pool.
	IpamPoolAllocations []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations `json:"IpamPoolAllocations,omitempty" xml:"IpamPoolAllocations,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
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

func (s ListIpamPoolAllocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolAllocationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamPoolAllocationsResponseBody) GetIpamPoolAllocations() []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	return s.IpamPoolAllocations
}

func (s *ListIpamPoolAllocationsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamPoolAllocationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolAllocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamPoolAllocationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamPoolAllocationsResponseBody) SetCount(v int64) *ListIpamPoolAllocationsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetIpamPoolAllocations(v []*ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) *ListIpamPoolAllocationsResponseBody {
	s.IpamPoolAllocations = v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetMaxResults(v int64) *ListIpamPoolAllocationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetNextToken(v string) *ListIpamPoolAllocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetRequestId(v string) *ListIpamPoolAllocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) SetTotalCount(v int64) *ListIpamPoolAllocationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamPoolAllocationsResponseBodyIpamPoolAllocations struct {
	// The allocated CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-05-19T08:59:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the allocation.
	//
	// example:
	//
	// test description
	IpamPoolAllocationDescription *string `json:"IpamPoolAllocationDescription,omitempty" xml:"IpamPoolAllocationDescription,omitempty"`
	// The ID of the instance to which CIDR blocks are allocated from the IPAM pool.
	//
	// example:
	//
	// ipam-pool-alloc-112za33e4****
	IpamPoolAllocationId *string `json:"IpamPoolAllocationId,omitempty" xml:"IpamPoolAllocationId,omitempty"`
	// The name of the allocation.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource to which the CIDR block is allocated.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 132193271328****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource to which the CIDR block is allocated. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **IpamPool**
	//
	// 	- **Custom**
	//
	// example:
	//
	// Custom
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 192.168.0.0/16
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetIpamPoolAllocationDescription() *string {
	return s.IpamPoolAllocationDescription
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetIpamPoolAllocationId() *string {
	return s.IpamPoolAllocationId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetIpamPoolAllocationName() *string {
	return s.IpamPoolAllocationName
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) GetStatus() *string {
	return s.Status
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetCidr(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetCreationTime(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.CreationTime = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationDescription(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationDescription = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolAllocationName(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetIpamPoolId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetRegionId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceOwnerId(v int64) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceRegionId(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetResourceType(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.ResourceType = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetSourceCidr(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.SourceCidr = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) SetStatus(v string) *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations {
	s.Status = &v
	return s
}

func (s *ListIpamPoolAllocationsResponseBodyIpamPoolAllocations) Validate() error {
	return dara.Validate(s)
}
