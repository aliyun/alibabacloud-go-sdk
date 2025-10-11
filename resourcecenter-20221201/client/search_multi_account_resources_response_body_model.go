// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMultiAccountResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*SearchMultiAccountResourcesResponseBodyFilters) *SearchMultiAccountResourcesResponseBody
	GetFilters() []*SearchMultiAccountResourcesResponseBodyFilters
	SetMaxResults(v int32) *SearchMultiAccountResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchMultiAccountResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchMultiAccountResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*SearchMultiAccountResourcesResponseBodyResources) *SearchMultiAccountResourcesResponseBody
	GetResources() []*SearchMultiAccountResourcesResponseBodyResources
	SetScope(v string) *SearchMultiAccountResourcesResponseBody
	GetScope() *string
}

type SearchMultiAccountResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchMultiAccountResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchMultiAccountResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The search scope.
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBody) GetFilters() []*SearchMultiAccountResourcesResponseBodyFilters {
	return s.Filters
}

func (s *SearchMultiAccountResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchMultiAccountResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchMultiAccountResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMultiAccountResourcesResponseBody) GetResources() []*SearchMultiAccountResourcesResponseBodyResources {
	return s.Resources
}

func (s *SearchMultiAccountResourcesResponseBody) GetScope() *string {
	return s.Scope
}

func (s *SearchMultiAccountResourcesResponseBody) SetFilters(v []*SearchMultiAccountResourcesResponseBodyFilters) *SearchMultiAccountResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetMaxResults(v int32) *SearchMultiAccountResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetNextToken(v string) *SearchMultiAccountResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetRequestId(v string) *SearchMultiAccountResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetResources(v []*SearchMultiAccountResourcesResponseBodyResources) *SearchMultiAccountResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetScope(v string) *SearchMultiAccountResourcesResponseBody {
	s.Scope = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchMultiAccountResourcesResponseBodyFilters struct {
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

func (s SearchMultiAccountResourcesResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) GetKey() *string {
	return s.Key
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) GetMatchType() *string {
	return s.MatchType
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) GetValues() []*string {
	return s.Values
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetKey(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetMatchType(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetValues(v []*string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Values = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}

type SearchMultiAccountResourcesResponseBodyResources struct {
	// The ID of the management account or member of the resource directory.
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
	// 2023-06-14T14:35:45Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
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
	Tags []*SearchMultiAccountResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetIpAddressAttributes() []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	return s.IpAddressAttributes
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetTags() []*SearchMultiAccountResourcesResponseBodyResourcesTags {
	return s.Tags
}

func (s *SearchMultiAccountResourcesResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetAccountId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetCreateTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetExpireTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetRegionId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceName(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceType(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetTags(v []*SearchMultiAccountResourcesResponseBodyResourcesTags) *SearchMultiAccountResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetZoneId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes struct {
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

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GetVersion() *string {
	return s.Version
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) Validate() error {
	return dara.Validate(s)
}

type SearchMultiAccountResourcesResponseBodyResourcesTags struct {
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

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) GetKey() *string {
	return s.Key
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) GetValue() *string {
	return s.Value
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetKey(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetValue(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) Validate() error {
	return dara.Validate(s)
}
