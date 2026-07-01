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
	// The allocated CIDR block that you want to query.
	//
	// > Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// A list of instance IDs of the IPAM pool CIDR block allocations.
	IpamPoolAllocationIds []*string `json:"IpamPoolAllocationIds,omitempty" xml:"IpamPoolAllocationIds,omitempty" type:"Repeated"`
	// The name of the IPAM pool CIDR block allocation.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test name
	IpamPoolAllocationName *string `json:"IpamPoolAllocationName,omitempty" xml:"IpamPoolAllocationName,omitempty"`
	// The instance ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no next page exists.
	//
	// - If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region of the IPAM pool that contains the CIDR block allocation.
	//
	// > If the IPAM pool has a specific region, this parameter specifies that region. If the IPAM pool does not have a specific region, this parameter specifies the managed region of IPAM.
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
