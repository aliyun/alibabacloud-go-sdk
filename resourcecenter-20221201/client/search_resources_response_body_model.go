// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*SearchResourcesResponseBodyFilters) *SearchResourcesResponseBody
	GetFilters() []*SearchResourcesResponseBodyFilters
	SetMaxResults(v int32) *SearchResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*SearchResourcesResponseBodyResources) *SearchResourcesResponseBody
	GetResources() []*SearchResourcesResponseBodyResources
}

type SearchResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBody) GetFilters() []*SearchResourcesResponseBodyFilters {
	return s.Filters
}

func (s *SearchResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchResourcesResponseBody) GetResources() []*SearchResourcesResponseBodyResources {
	return s.Resources
}

func (s *SearchResourcesResponseBody) SetFilters(v []*SearchResourcesResponseBodyFilters) *SearchResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchResourcesResponseBody) SetMaxResults(v int32) *SearchResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesResponseBody) SetNextToken(v string) *SearchResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesResponseBody) SetRequestId(v string) *SearchResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResourcesResponseBody) SetResources(v []*SearchResourcesResponseBodyResources) *SearchResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *SearchResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchResourcesResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyFilters) GetKey() *string {
	return s.Key
}

func (s *SearchResourcesResponseBodyFilters) GetMatchType() *string {
	return s.MatchType
}

func (s *SearchResourcesResponseBodyFilters) GetValues() []*string {
	return s.Values
}

func (s *SearchResourcesResponseBodyFilters) SetKey(v string) *SearchResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetMatchType(v string) *SearchResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetValues(v []*string) *SearchResourcesResponseBodyFilters {
	s.Values = v
	return s
}

func (s *SearchResourcesResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}

type SearchResourcesResponseBodyResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vtb-bp11lbh452fr8940s****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResources) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchResourcesResponseBodyResources) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *SearchResourcesResponseBodyResources) GetIpAddressAttributes() []*SearchResourcesResponseBodyResourcesIpAddressAttributes {
	return s.IpAddressAttributes
}

func (s *SearchResourcesResponseBodyResources) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *SearchResourcesResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchResourcesResponseBodyResources) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *SearchResourcesResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *SearchResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *SearchResourcesResponseBodyResources) GetTags() []*SearchResourcesResponseBodyResourcesTags {
	return s.Tags
}

func (s *SearchResourcesResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchResourcesResponseBodyResources) SetAccountId(v string) *SearchResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetCreateTime(v string) *SearchResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetExpireTime(v string) *SearchResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchResourcesResponseBodyResourcesIpAddressAttributes) *SearchResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetRegionId(v string) *SearchResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceName(v string) *SearchResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceType(v string) *SearchResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetTags(v []*SearchResourcesResponseBodyResourcesTags) *SearchResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetZoneId(v string) *SearchResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type SearchResourcesResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) GetVersion() *string {
	return s.Version
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) Validate() error {
	return dara.Validate(s)
}

type SearchResourcesResponseBodyResourcesTags struct {
	// The key of tag N.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesTags) GetKey() *string {
	return s.Key
}

func (s *SearchResourcesResponseBodyResourcesTags) GetValue() *string {
	return s.Value
}

func (s *SearchResourcesResponseBodyResourcesTags) SetKey(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesTags) SetValue(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesTags) Validate() error {
	return dara.Validate(s)
}
