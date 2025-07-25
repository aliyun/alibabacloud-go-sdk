// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolCidrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListIpamPoolCidrsResponseBody
	GetCount() *int64
	SetIpamPoolCidrs(v []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs) *ListIpamPoolCidrsResponseBody
	GetIpamPoolCidrs() []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs
	SetMaxResults(v int64) *ListIpamPoolCidrsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamPoolCidrsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpamPoolCidrsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIpamPoolCidrsResponseBody
	GetTotalCount() *int64
}

type ListIpamPoolCidrsResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The IDs of IPAM pools.
	IpamPoolCidrs []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs `json:"IpamPoolCidrs,omitempty" xml:"IpamPoolCidrs,omitempty" type:"Repeated"`
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
	// 9E7CCB95-62E0-534D-9B9A-71F27E8B71B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpamPoolCidrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListIpamPoolCidrsResponseBody) GetIpamPoolCidrs() []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	return s.IpamPoolCidrs
}

func (s *ListIpamPoolCidrsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamPoolCidrsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamPoolCidrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpamPoolCidrsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIpamPoolCidrsResponseBody) SetCount(v int64) *ListIpamPoolCidrsResponseBody {
	s.Count = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetIpamPoolCidrs(v []*ListIpamPoolCidrsResponseBodyIpamPoolCidrs) *ListIpamPoolCidrsResponseBody {
	s.IpamPoolCidrs = v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetMaxResults(v int64) *ListIpamPoolCidrsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetNextToken(v string) *ListIpamPoolCidrsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetRequestId(v string) *ListIpamPoolCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) SetTotalCount(v int64) *ListIpamPoolCidrsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpamPoolCidrsResponseBodyIpamPoolCidrs struct {
	// The provisioned CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The status of the CIDR block provisioned to the IPAM pool. Valid values:
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

func (s ListIpamPoolCidrsResponseBodyIpamPoolCidrs) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolCidrsResponseBodyIpamPoolCidrs) GoString() string {
	return s.String()
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) GetCidr() *string {
	return s.Cidr
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) GetStatus() *string {
	return s.Status
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetCidr(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.Cidr = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetIpamPoolId(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.IpamPoolId = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) SetStatus(v string) *ListIpamPoolCidrsResponseBodyIpamPoolCidrs {
	s.Status = &v
	return s
}

func (s *ListIpamPoolCidrsResponseBodyIpamPoolCidrs) Validate() error {
	return dara.Validate(s)
}
