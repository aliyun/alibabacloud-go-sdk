// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolAllocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *ListIpamPoolAllocationsRequest
	GetCidr() *string
	SetIpamPoolAllocationIds(v []*string) *ListIpamPoolAllocationsRequest
	GetIpamPoolAllocationIds() []*string
	SetIpamPoolAllocationName(v string) *ListIpamPoolAllocationsRequest
	GetIpamPoolAllocationName() *string
	SetIpamPoolId(v string) *ListIpamPoolAllocationsRequest
	GetIpamPoolId() *string
	SetMaxResults(v int32) *ListIpamPoolAllocationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamPoolAllocationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpamPoolAllocationsRequest
	GetRegionId() *string
}

type ListIpamPoolAllocationsRequest struct {
	// Specifies whether to query allocations by specifying allocated CIDR blocks.
	//
	// **
	//
	// **Usage notes*	- Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The IDs of the instances to which CIDR blocks are allocated from the IPAM pool.
	IpamPoolAllocationIds []*string `json:"IpamPoolAllocationIds,omitempty" xml:"IpamPoolAllocationIds,omitempty" type:"Repeated"`
	// The name of  allocations.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	// The ID of the region where you want to perform the operation.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListIpamPoolAllocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolAllocationsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsRequest) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamPoolAllocationsRequest) GetIpamPoolAllocationIds() []*string {
	return s.IpamPoolAllocationIds
}

func (s *ListIpamPoolAllocationsRequest) GetIpamPoolAllocationName() *string {
	return s.IpamPoolAllocationName
}

func (s *ListIpamPoolAllocationsRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamPoolAllocationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamPoolAllocationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolAllocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamPoolAllocationsRequest) SetCidr(v string) *ListIpamPoolAllocationsRequest {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolAllocationIds(v []*string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolAllocationIds = v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolAllocationName(v string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolAllocationName = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetIpamPoolId(v string) *ListIpamPoolAllocationsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetMaxResults(v int32) *ListIpamPoolAllocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetNextToken(v string) *ListIpamPoolAllocationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) SetRegionId(v string) *ListIpamPoolAllocationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolAllocationsRequest) Validate() error {
	return dara.Validate(s)
}
