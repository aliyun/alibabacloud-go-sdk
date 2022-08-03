// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListResourceRelationshipsRequest struct {
	MaxResults         *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken          *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SourceRegionId     *string   `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	SourceResourceId   []*string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty" type:"Repeated"`
	SourceResourceType *string   `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	TargetResourceType []*string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsRequest) SetMaxResults(v int32) *ListResourceRelationshipsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetNextToken(v string) *ListResourceRelationshipsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetSourceRegionId(v string) *ListResourceRelationshipsRequest {
	s.SourceRegionId = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetSourceResourceId(v []*string) *ListResourceRelationshipsRequest {
	s.SourceResourceId = v
	return s
}

func (s *ListResourceRelationshipsRequest) SetSourceResourceType(v string) *ListResourceRelationshipsRequest {
	s.SourceResourceType = &v
	return s
}

func (s *ListResourceRelationshipsRequest) SetTargetResourceType(v []*string) *ListResourceRelationshipsRequest {
	s.TargetResourceType = v
	return s
}

type ListResourceRelationshipsResponseBody struct {
	MaxResults            *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRelationships []*ListResourceRelationshipsResponseBodyResourceRelationships `json:"ResourceRelationships,omitempty" xml:"ResourceRelationships,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBody) SetMaxResults(v int32) *ListResourceRelationshipsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetNextToken(v string) *ListResourceRelationshipsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetRequestId(v string) *ListResourceRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetResourceRelationships(v []*ListResourceRelationshipsResponseBodyResourceRelationships) *ListResourceRelationshipsResponseBody {
	s.ResourceRelationships = v
	return s
}

type ListResourceRelationshipsResponseBodyResourceRelationships struct {
	AccountId          *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	SourceRegionId     *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	SourceResourceId   *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	SourceResourceType *string `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	TargetRegionId     *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	TargetResourceId   *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
	TargetResourceType *string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty"`
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetAccountId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.AccountId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetSourceRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.SourceRegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetSourceResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.SourceResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetSourceResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.SourceResourceType = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetTargetRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.TargetRegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetTargetResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.TargetResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetTargetResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.TargetResourceType = &v
	return s
}

type ListResourceRelationshipsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceRelationshipsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceRelationshipsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponse) SetHeaders(v map[string]*string) *ListResourceRelationshipsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceRelationshipsResponse) SetStatusCode(v int32) *ListResourceRelationshipsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceRelationshipsResponse) SetBody(v *ListResourceRelationshipsResponseBody) *ListResourceRelationshipsResponse {
	s.Body = v
	return s
}

type SearchResourcesRequest struct {
	Filter          []*SearchResourcesRequestFilter      `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults      *int32                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SortCriterion   *SearchResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequest) SetFilter(v []*SearchResourcesRequestFilter) *SearchResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchResourcesRequest) SetMaxResults(v int32) *SearchResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesRequest) SetNextToken(v string) *SearchResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesRequest) SetResourceGroupId(v string) *SearchResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesRequest) SetSortCriterion(v *SearchResourcesRequestSortCriterion) *SearchResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchResourcesRequestFilter struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestFilter) SetKey(v string) *SearchResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetMatchType(v string) *SearchResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetValue(v []*string) *SearchResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchResourcesRequestSortCriterion struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestSortCriterion) SetKey(v string) *SearchResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestSortCriterion) SetOrder(v string) *SearchResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchResourcesResponseBody struct {
	Filters    []*SearchResourcesResponseBodyFilters   `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*SearchResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBody) GoString() string {
	return s.String()
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

type SearchResourcesResponseBodyFilters struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Values    []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyFilters) GoString() string {
	return s.String()
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

type SearchResourcesResponseBodyResources struct {
	AccountId       *string                                     `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateTime      *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IpAddresses     []*string                                   `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string                                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName    *string                                     `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType    *string                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags            []*SearchResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneId          *string                                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResources) SetAccountId(v string) *SearchResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetCreateTime(v string) *SearchResourcesResponseBodyResources {
	s.CreateTime = &v
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

type SearchResourcesResponseBodyResourcesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesTags) SetKey(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesTags) SetValue(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponse) SetHeaders(v map[string]*string) *SearchResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchResourcesResponse) SetStatusCode(v int32) *SearchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResourcesResponse) SetBody(v *SearchResourcesResponseBody) *SearchResourcesResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rmc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListResourceRelationshipsWithOptions(request *ListResourceRelationshipsRequest, runtime *util.RuntimeOptions) (_result *ListResourceRelationshipsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRegionId)) {
		query["SourceRegionId"] = request.SourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourceId)) {
		query["SourceResourceId"] = request.SourceResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourceType)) {
		query["SourceResourceType"] = request.SourceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetResourceType)) {
		query["TargetResourceType"] = request.TargetResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceRelationships"),
		Version:     tea.String("2021-11-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceRelationships(request *ListResourceRelationshipsRequest) (_result *ListResourceRelationshipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceRelationshipsResponse{}
	_body, _err := client.ListResourceRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchResourcesWithOptions(request *SearchResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SortCriterion))) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchResources"),
		Version:     tea.String("2021-11-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchResources(request *SearchResourcesRequest) (_result *SearchResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchResourcesResponse{}
	_body, _err := client.SearchResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
