// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpPoolListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *DedicatedIpPoolListResponseBody
	GetCurrentPage() *string
	SetHasMore(v bool) *DedicatedIpPoolListResponseBody
	GetHasMore() *bool
	SetIpPools(v []*DedicatedIpPoolListResponseBodyIpPools) *DedicatedIpPoolListResponseBody
	GetIpPools() []*DedicatedIpPoolListResponseBodyIpPools
	SetPageSize(v string) *DedicatedIpPoolListResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DedicatedIpPoolListResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *DedicatedIpPoolListResponseBody
	GetTotalCounts() *int32
}

type DedicatedIpPoolListResponseBody struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Whether there is a next page
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// List of IP pools
	IpPools []*DedicatedIpPoolListResponseBodyIpPools `json:"IpPools,omitempty" xml:"IpPools,omitempty" type:"Repeated"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data under the current request conditions
	//
	// example:
	//
	// 5
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s DedicatedIpPoolListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListResponseBody) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DedicatedIpPoolListResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *DedicatedIpPoolListResponseBody) GetIpPools() []*DedicatedIpPoolListResponseBodyIpPools {
	return s.IpPools
}

func (s *DedicatedIpPoolListResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DedicatedIpPoolListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpPoolListResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *DedicatedIpPoolListResponseBody) SetCurrentPage(v string) *DedicatedIpPoolListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DedicatedIpPoolListResponseBody) SetHasMore(v bool) *DedicatedIpPoolListResponseBody {
	s.HasMore = &v
	return s
}

func (s *DedicatedIpPoolListResponseBody) SetIpPools(v []*DedicatedIpPoolListResponseBodyIpPools) *DedicatedIpPoolListResponseBody {
	s.IpPools = v
	return s
}

func (s *DedicatedIpPoolListResponseBody) SetPageSize(v string) *DedicatedIpPoolListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DedicatedIpPoolListResponseBody) SetRequestId(v string) *DedicatedIpPoolListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpPoolListResponseBody) SetTotalCounts(v int32) *DedicatedIpPoolListResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *DedicatedIpPoolListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DedicatedIpPoolListResponseBodyIpPools struct {
	// Creation time
	//
	// example:
	//
	// 2025-05-23T07:41:43Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// IP pool ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Number of source IP addresses
	//
	// example:
	//
	// 1
	IpCount *int32 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// List of IPs
	Ips []*DedicatedIpPoolListResponseBodyIpPoolsIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// IP pool name
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DedicatedIpPoolListResponseBodyIpPools) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListResponseBodyIpPools) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListResponseBodyIpPools) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DedicatedIpPoolListResponseBodyIpPools) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolListResponseBodyIpPools) GetIpCount() *int32 {
	return s.IpCount
}

func (s *DedicatedIpPoolListResponseBodyIpPools) GetIps() []*DedicatedIpPoolListResponseBodyIpPoolsIps {
	return s.Ips
}

func (s *DedicatedIpPoolListResponseBodyIpPools) GetName() *string {
	return s.Name
}

func (s *DedicatedIpPoolListResponseBodyIpPools) SetCreateTime(v string) *DedicatedIpPoolListResponseBodyIpPools {
	s.CreateTime = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPools) SetId(v string) *DedicatedIpPoolListResponseBodyIpPools {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPools) SetIpCount(v int32) *DedicatedIpPoolListResponseBodyIpPools {
	s.IpCount = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPools) SetIps(v []*DedicatedIpPoolListResponseBodyIpPoolsIps) *DedicatedIpPoolListResponseBodyIpPools {
	s.Ips = v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPools) SetName(v string) *DedicatedIpPoolListResponseBodyIpPools {
	s.Name = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPools) Validate() error {
	return dara.Validate(s)
}

type DedicatedIpPoolListResponseBodyIpPoolsIps struct {
	// Instance purchase ID
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// IP address
	//
	// example:
	//
	// xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DedicatedIpPoolListResponseBodyIpPoolsIps) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpPoolListResponseBodyIpPoolsIps) GoString() string {
	return s.String()
}

func (s *DedicatedIpPoolListResponseBodyIpPoolsIps) GetId() *string {
	return s.Id
}

func (s *DedicatedIpPoolListResponseBodyIpPoolsIps) GetIp() *string {
	return s.Ip
}

func (s *DedicatedIpPoolListResponseBodyIpPoolsIps) SetId(v string) *DedicatedIpPoolListResponseBodyIpPoolsIps {
	s.Id = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPoolsIps) SetIp(v string) *DedicatedIpPoolListResponseBodyIpPoolsIps {
	s.Ip = &v
	return s
}

func (s *DedicatedIpPoolListResponseBodyIpPoolsIps) Validate() error {
	return dara.Validate(s)
}
