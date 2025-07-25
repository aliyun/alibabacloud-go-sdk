// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolCidrsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *ListIpamPoolCidrsRequest
	GetCidr() *string
	SetIpamPoolId(v string) *ListIpamPoolCidrsRequest
	GetIpamPoolId() *string
	SetMaxResults(v int32) *ListIpamPoolCidrsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamPoolCidrsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpamPoolCidrsRequest
	GetRegionId() *string
}

type ListIpamPoolCidrsRequest struct {
	// The provisioned CIDR block that you want to query.
	//
	// >  Only IPv4 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
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
	// The ID of the region where the IPAM instance is hosted.
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

func (s ListIpamPoolCidrsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolCidrsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsRequest) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamPoolCidrsRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamPoolCidrsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamPoolCidrsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolCidrsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamPoolCidrsRequest) SetCidr(v string) *ListIpamPoolCidrsRequest {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetIpamPoolId(v string) *ListIpamPoolCidrsRequest {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetMaxResults(v int32) *ListIpamPoolCidrsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetNextToken(v string) *ListIpamPoolCidrsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) SetRegionId(v string) *ListIpamPoolCidrsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamPoolCidrsRequest) Validate() error {
	return dara.Validate(s)
}
