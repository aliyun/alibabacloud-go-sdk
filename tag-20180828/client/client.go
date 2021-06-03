// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagProduct           *string `json:"TagProduct,omitempty" xml:"TagProduct,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerAccount(v string) *ListTagKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetTagProduct(v string) *ListTagKeysRequest {
	s.TagProduct = &v
	return s
}

type ListTagKeysResponseBody struct {
	NextToken *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Keys      *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

type ListTagKeysResponseBodyKeys struct {
	Key []*ListTagKeysResponseBodyKeysKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v []*ListTagKeysResponseBodyKeysKey) *ListTagKeysResponseBodyKeys {
	s.Key = v
	return s
}

type ListTagKeysResponseBodyKeysKey struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ListTagKeysResponseBodyKeysKey) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeysKey) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeysKey) SetKey(v string) *ListTagKeysResponseBodyKeysKey {
	s.Key = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) SetCategory(v string) *ListTagKeysResponseBodyKeysKey {
	s.Category = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tags                 *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListTagResourcesRequest) SetCategory(v string) *ListTagResourcesRequest {
	s.Category = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceARN(v []*string) *ListTagResourcesRequest {
	s.ResourceARN = v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceARN *string                                         `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Tags        []*ListTagResourcesResponseBodyTagResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceARN(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceARN = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTags(v []*ListTagResourcesResponseBodyTagResourcesTags) *ListTagResourcesResponseBodyTagResources {
	s.Tags = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTags struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetKey(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetValue(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Value = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetCategory(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Category = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagProduct           *string `json:"TagProduct,omitempty" xml:"TagProduct,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerAccount(v string) *ListTagValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagProduct(v string) *ListTagValuesRequest {
	s.TagProduct = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values    *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponseBodyValues struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyValues) SetValue(v []*string) *ListTagValuesResponseBodyValues {
	s.Value = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags                 *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *TagResourcesRequest) SetResourceARN(v []*string) *TagResourcesRequest {
	s.ResourceARN = v
	return s
}

type TagResourcesResponseBody struct {
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedResources *TagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetFailedResources(v *TagResourcesResponseBodyFailedResources) *TagResourcesResponseBody {
	s.FailedResources = v
	return s
}

type TagResourcesResponseBodyFailedResources struct {
	FailedResource []*TagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s TagResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResources) SetFailedResource(v []*TagResourcesResponseBodyFailedResourcesFailedResource) *TagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

type TagResourcesResponseBodyFailedResourcesFailedResource struct {
	ResourceARN *string                                                      `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Result      *TagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *TagResourcesResponseBodyFailedResourcesFailedResourceResult) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

type TagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceARN(v []*string) *UntagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedResources *UntagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetFailedResources(v *UntagResourcesResponseBodyFailedResources) *UntagResourcesResponseBody {
	s.FailedResources = v
	return s
}

type UntagResourcesResponseBodyFailedResources struct {
	FailedResource []*UntagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s UntagResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResources) SetFailedResource(v []*UntagResourcesResponseBodyFailedResourcesFailedResource) *UntagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

type UntagResourcesResponseBodyFailedResourcesFailedResource struct {
	ResourceARN *string                                                        `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Result      *UntagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

type UntagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("tag.aliyuncs.com"),
		"cn-beijing":                  tea.String("tag.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("tag.aliyuncs.com"),
		"cn-shanghai":                 tea.String("tag.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("tag.aliyuncs.com"),
		"cn-hongkong":                 tea.String("tag.aliyuncs.com"),
		"ap-southeast-1":              tea.String("tag.aliyuncs.com"),
		"us-west-1":                   tea.String("tag.aliyuncs.com"),
		"us-east-1":                   tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("tag.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("tag.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("tag.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("tag.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("tag.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("tag.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("tag.aliyuncs.com"),
		"cn-edge-1":                   tea.String("tag.aliyuncs.com"),
		"cn-fujian":                   tea.String("tag.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("tag.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("tag.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("tag.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("tag.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("tag.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("tag.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("tag.aliyuncs.com"),
		"cn-wuhan":                    tea.String("tag.aliyuncs.com"),
		"cn-yushanfang":               tea.String("tag.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("tag.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("tag.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("tag.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("tag.cn-shenzhen-cloudstone.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("tag.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tag"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-08-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
