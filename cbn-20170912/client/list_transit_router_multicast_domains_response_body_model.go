// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterMulticastDomainsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterMulticastDomainsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterMulticastDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterMulticastDomainsResponseBody
	GetTotalCount() *int32
	SetTransitRouterMulticastDomains(v []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) *ListTransitRouterMulticastDomainsResponseBody
	GetTransitRouterMulticastDomains() []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains
}

type ListTransitRouterMulticastDomainsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8A0F93D1-FD6C-56FC-B6D2-668FC92D12D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of multicast domains.
	TransitRouterMulticastDomains []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains `json:"TransitRouterMulticastDomains,omitempty" xml:"TransitRouterMulticastDomains,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterMulticastDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterMulticastDomainsResponseBody) GetTransitRouterMulticastDomains() []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	return s.TransitRouterMulticastDomains
}

func (s *ListTransitRouterMulticastDomainsResponseBody) SetMaxResults(v int32) *ListTransitRouterMulticastDomainsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBody) SetNextToken(v string) *ListTransitRouterMulticastDomainsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBody) SetRequestId(v string) *ListTransitRouterMulticastDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBody) SetTotalCount(v int32) *ListTransitRouterMulticastDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBody) SetTransitRouterMulticastDomains(v []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) *ListTransitRouterMulticastDomainsResponseBody {
	s.TransitRouterMulticastDomains = v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains struct {
	// The CEN instance ID.
	//
	// example:
	//
	// cen-a7syd349kne38g****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Multicast domain feature.
	Options *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the multicast domain.
	//
	// The valid value is **Active**, which indicates that the multicast domain is available.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The transit router ID.
	//
	// example:
	//
	// tr-bp1c23ijrl6d6c226h***
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The description of the multicast domain.
	//
	// example:
	//
	// desctest
	TransitRouterMulticastDomainDescription *string `json:"TransitRouterMulticastDomainDescription,omitempty" xml:"TransitRouterMulticastDomainDescription,omitempty"`
	// The ID of the multicast domain.
	//
	// example:
	//
	// tr-mcast-domain-3r3bvbypxqheej****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The name of the multicast domain.
	//
	// example:
	//
	// nametest
	TransitRouterMulticastDomainName *string `json:"TransitRouterMulticastDomainName,omitempty" xml:"TransitRouterMulticastDomainName,omitempty"`
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetOptions() *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions {
	return s.Options
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetTags() []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags {
	return s.Tags
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetTransitRouterMulticastDomainDescription() *string {
	return s.TransitRouterMulticastDomainDescription
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) GetTransitRouterMulticastDomainName() *string {
	return s.TransitRouterMulticastDomainName
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetCenId(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetOptions(v *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.Options = v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetRegionId(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetStatus(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.Status = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetTags(v []*ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.Tags = v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetTransitRouterId(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetTransitRouterMulticastDomainDescription(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.TransitRouterMulticastDomainDescription = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) SetTransitRouterMulticastDomainName(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains {
	s.TransitRouterMulticastDomainName = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomains) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions struct {
	// Indicates whether the IGMP feature is enabled for the multicast domain.
	//
	// example:
	//
	// enable
	Igmpv2Support *string `json:"Igmpv2Support,omitempty" xml:"Igmpv2Support,omitempty"`
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) GetIgmpv2Support() *string {
	return s.Igmpv2Support
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) SetIgmpv2Support(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions {
	s.Igmpv2Support = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsOptions) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags struct {
	// The tag key.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) SetKey(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) SetValue(v string) *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterMulticastDomainsResponseBodyTransitRouterMulticastDomainsTags) Validate() error {
	return dara.Validate(s)
}
