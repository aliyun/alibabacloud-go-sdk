// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListResourceRelationshipsRequest struct {
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the resource whose associated resources you want to query.
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The IDs of the resource whose associated resources you want to query.
	//
	// You can specify a maximum of 10 resource IDs.
	SourceResourceId []*string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty" type:"Repeated"`
	// The type of the resource whose associated resources you want to query.
	SourceResourceType *string `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	// The types of the associated resources that you want to query.
	//
	// You can specify a maximum of 10 resource types.
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
	// The maximum number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the associated resources.
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
	// The ID of the Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The region ID of the specified resource.
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The ID of the specified resource.
	SourceResourceId *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	// The type of the specified resource.
	SourceResourceType *string `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	// The region ID of the associated resource.
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The ID of the associated resource.
	TargetResourceId *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
	// The type of the associated resource.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceRelationshipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The filter conditions.
	Filter []*SearchResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The method that is used to sort the entries.
	SortCriterion *SearchResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
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
	// The key of the filter condition. Valid values:
	//
	// *   ResourceType: resource type
	// *   RegionId: region ID
	// *   ResourceId: resource ID
	// *   ResourceGroupId: resource group ID
	// *   ResourceName: resource name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method. Set the value to Equals. This value indicates that resources that match the filter conditions are queried.
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
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
	// The attribute based on which the entries are sorted.
	//
	// The value `CreateTime` indicates the creation time of resources.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// *   ASC: The entries are sorted in ascending order. This value is the default value.
	// *   DESC: The entries are sorted in descending order.
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
	// The filter conditions.
	Filters []*SearchResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resources.
	Resources []*SearchResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
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
	// The key of the filter condition.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method.
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	// The ID of the Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

/**
 * This section provides the types of resources that can be queried. Two-way queries are supported. For example, you can query the disks (ACS::ECS::Disk) that are associated with a specific Elastic Compute Service (ECS) instance (ACS::ECS::Instance) or query the ECS instance that is associated with a specific disk.
 * - For ECS instances, the following types of resources can be queried:    - ACS::ECS::Disk
 *   - ACS::EIP::EipAddress
 *   - ACS::VPC::VPC
 *   - ACS::ECS::KeyPair
 *   - ACS::ECS::SecurityGroup
 *   - ACS::ECS::NetworkInterface
 *   - ACS::ECS::Image
 * - For virtual private clouds (VPCs), which are indicated by ACS::VPC::VPC, the following types of resources can be queried:    - ACS::ECS::Instance
 *   - ACS::RDS::DBInstance
 *   - ACS::SLB::LoadBalancer
 *   - ACS::ALB::LoadBalancer
 *   - ACS::Elasticsearch::Instance
 *   - ACS::Redis::DBInstance
 *   - ACS::PolarDB::DBCluster
 *   - ACS::MongoDB::DBInstance
 *   - ACS::DRDS::PolarDBXInstance
 *   - ACS::EDAS::Cluster
 *   - ACS::ECI::ContainerGroup
 *   - ACS::ADB::DBCluster
 *   - ACS::DRDS::DBInstance
 *   - ACS::HBase::Cluster
 *   - ACS::EMR::Cluster
 * This topic provides an example on how to call the API operation to query the resources that are associated with the ECS instance `i-uf6imlgyr1nudhud****` in the China (Shanghai) region.
 * ## Prerequisites
 * Resource Meta Center (RMC) is enabled. For more information, see [Query resources that belong to different resource groups](~~310198~~).
 * ## QPS limits
 * You can call this API operation up to 20 times per second per account. Requests that exceed this limit will fail, and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request ListResourceRelationshipsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceRelationshipsResponse
 */
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

/**
 * This section provides the types of resources that can be queried. Two-way queries are supported. For example, you can query the disks (ACS::ECS::Disk) that are associated with a specific Elastic Compute Service (ECS) instance (ACS::ECS::Instance) or query the ECS instance that is associated with a specific disk.
 * - For ECS instances, the following types of resources can be queried:    - ACS::ECS::Disk
 *   - ACS::EIP::EipAddress
 *   - ACS::VPC::VPC
 *   - ACS::ECS::KeyPair
 *   - ACS::ECS::SecurityGroup
 *   - ACS::ECS::NetworkInterface
 *   - ACS::ECS::Image
 * - For virtual private clouds (VPCs), which are indicated by ACS::VPC::VPC, the following types of resources can be queried:    - ACS::ECS::Instance
 *   - ACS::RDS::DBInstance
 *   - ACS::SLB::LoadBalancer
 *   - ACS::ALB::LoadBalancer
 *   - ACS::Elasticsearch::Instance
 *   - ACS::Redis::DBInstance
 *   - ACS::PolarDB::DBCluster
 *   - ACS::MongoDB::DBInstance
 *   - ACS::DRDS::PolarDBXInstance
 *   - ACS::EDAS::Cluster
 *   - ACS::ECI::ContainerGroup
 *   - ACS::ADB::DBCluster
 *   - ACS::DRDS::DBInstance
 *   - ACS::HBase::Cluster
 *   - ACS::EMR::Cluster
 * This topic provides an example on how to call the API operation to query the resources that are associated with the ECS instance `i-uf6imlgyr1nudhud****` in the China (Shanghai) region.
 * ## Prerequisites
 * Resource Meta Center (RMC) is enabled. For more information, see [Query resources that belong to different resource groups](~~310198~~).
 * ## QPS limits
 * You can call this API operation up to 20 times per second per account. Requests that exceed this limit will fail, and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request ListResourceRelationshipsRequest
 * @return ListResourceRelationshipsResponse
 */
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

/**
 * For more information about resource types that support RMC, see [Resource types that support RMC](https://www.alibabacloud.com/help/en/resource-management/latest/resource-types-that-support-rmc).
 * This topic provides an example on how to call the API operation to query the resources that can be accessed within the current account in the China (Hangzhou) region.
 * ## Prerequisites
 * Resource Meta Center (RMC) is enabled. For more information, see [Query resources that belong to different resource groups](~~310198~~).
 * ## QPS limits
 * You can call this API operation up to 20 times per second per account. Requests that exceed this limit will fail, and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request SearchResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SearchResourcesResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
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

/**
 * For more information about resource types that support RMC, see [Resource types that support RMC](https://www.alibabacloud.com/help/en/resource-management/latest/resource-types-that-support-rmc).
 * This topic provides an example on how to call the API operation to query the resources that can be accessed within the current account in the China (Hangzhou) region.
 * ## Prerequisites
 * Resource Meta Center (RMC) is enabled. For more information, see [Query resources that belong to different resource groups](~~310198~~).
 * ## QPS limits
 * You can call this API operation up to 20 times per second per account. Requests that exceed this limit will fail, and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
 *
 * @param request SearchResourcesRequest
 * @return SearchResourcesResponse
 */
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
