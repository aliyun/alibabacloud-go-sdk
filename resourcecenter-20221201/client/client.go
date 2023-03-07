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

type DisableMultiAccountResourceCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *DisableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *DisableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *DisableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetBody(v *DisableMultiAccountResourceCenterResponseBody) *DisableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type DisableResourceCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponseBody) SetRequestId(v string) *DisableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableResourceCenterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponse) SetHeaders(v map[string]*string) *DisableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceCenterResponse) SetStatusCode(v int32) *DisableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceCenterResponse) SetBody(v *DisableResourceCenterResponseBody) *DisableResourceCenterResponse {
	s.Body = v
	return s
}

type EnableMultiAccountResourceCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetStatus(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *EnableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *EnableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetBody(v *EnableMultiAccountResourceCenterResponseBody) *EnableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type EnableResourceCenterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponseBody) SetRequestId(v string) *EnableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableResourceCenterResponseBody) SetStatus(v string) *EnableResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableResourceCenterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponse) SetHeaders(v map[string]*string) *EnableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableResourceCenterResponse) SetStatusCode(v int32) *EnableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableResourceCenterResponse) SetBody(v *EnableResourceCenterResponseBody) *EnableResourceCenterResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponseBody struct {
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMultiAccountResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetBody(v *GetMultiAccountResourceCenterServiceStatusResponseBody) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceConfigurationRequest struct {
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMultiAccountResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationRequest) SetAccountId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceRegionId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceType(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBody struct {
	AccountId       *string                                                 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Configuration   map[string]interface{}                                  `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	CreateTime      *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IpAddresses     []*string                                               `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	RegionId        *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string                                                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName    *string                                                 `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType    *string                                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags            []*GetMultiAccountResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneId          *string                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetAccountId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetMultiAccountResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetCreateTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRegionId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRequestId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceName(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceType(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetTags(v []*GetMultiAccountResourceConfigurationResponseBodyTags) *GetMultiAccountResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetZoneId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetKey(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetValue(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetMultiAccountResourceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMultiAccountResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMultiAccountResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetStatusCode(v int32) *GetMultiAccountResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetBody(v *GetMultiAccountResourceConfigurationResponseBody) *GetMultiAccountResourceConfigurationResponse {
	s.Body = v
	return s
}

type GetResourceCenterServiceStatusResponseBody struct {
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetBody(v *GetResourceCenterServiceStatusResponseBody) *GetResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetResourceConfigurationRequest struct {
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationRequest) SetResourceId(v string) *GetResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceRegionId(v string) *GetResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceType(v string) *GetResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetResourceConfigurationResponseBody struct {
	AccountId       *string                                     `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Configuration   map[string]interface{}                      `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	CreateTime      *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IpAddresses     []*string                                   `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	RegionId        *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string                                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName    *string                                     `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType    *string                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags            []*GetResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneId          *string                                     `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBody) SetAccountId(v string) *GetResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetCreateTime(v string) *GetResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRegionId(v string) *GetResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRequestId(v string) *GetResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceName(v string) *GetResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceType(v string) *GetResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetTags(v []*GetResourceConfigurationResponseBodyTags) *GetResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetZoneId(v string) *GetResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetResourceConfigurationResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyTags) SetKey(v string) *GetResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyTags) SetValue(v string) *GetResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetResourceConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationResponse) SetStatusCode(v int32) *GetResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationResponse) SetBody(v *GetResourceConfigurationResponseBody) *GetResourceConfigurationResponse {
	s.Body = v
	return s
}

type ListMultiAccountResourceGroupsRequest struct {
	AccountId        *string   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsRequest) SetAccountId(v string) *ListMultiAccountResourceGroupsRequest {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetMaxResults(v int32) *ListMultiAccountResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetNextToken(v string) *ListMultiAccountResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListMultiAccountResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

type ListMultiAccountResourceGroupsResponseBody struct {
	NextToken      *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroups []*ListMultiAccountResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetNextToken(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetRequestId(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetResourceGroups(v []*ListMultiAccountResourceGroupsResponseBodyResourceGroups) *ListMultiAccountResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

type ListMultiAccountResourceGroupsResponseBodyResourceGroups struct {
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetAccountId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetCreateDate(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.CreateDate = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetDisplayName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.DisplayName = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Id = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Name = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetStatus(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Status = &v
	return s
}

type ListMultiAccountResourceGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMultiAccountResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMultiAccountResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetStatusCode(v int32) *ListMultiAccountResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetBody(v *ListMultiAccountResourceGroupsResponseBody) *ListMultiAccountResourceGroupsResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagKeysRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Scope      *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListMultiAccountTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysRequest) SetMatchType(v string) *ListMultiAccountTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetMaxResults(v int32) *ListMultiAccountTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetNextToken(v string) *ListMultiAccountTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetScope(v string) *ListMultiAccountTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetTagKey(v string) *ListMultiAccountTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListMultiAccountTagKeysResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponseBody) SetNextToken(v string) *ListMultiAccountTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetRequestId(v string) *ListMultiAccountTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetTagKeys(v []*string) *ListMultiAccountTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListMultiAccountTagKeysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMultiAccountTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMultiAccountTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetStatusCode(v int32) *ListMultiAccountTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetBody(v *ListMultiAccountTagKeysResponseBody) *ListMultiAccountTagKeysResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagValuesRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Scope      *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListMultiAccountTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesRequest) SetMatchType(v string) *ListMultiAccountTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetMaxResults(v int32) *ListMultiAccountTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetNextToken(v string) *ListMultiAccountTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetScope(v string) *ListMultiAccountTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagKey(v string) *ListMultiAccountTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagValue(v string) *ListMultiAccountTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListMultiAccountTagValuesResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponseBody) SetNextToken(v string) *ListMultiAccountTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetRequestId(v string) *ListMultiAccountTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetTagValues(v []*string) *ListMultiAccountTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListMultiAccountTagValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMultiAccountTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMultiAccountTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetStatusCode(v int32) *ListMultiAccountTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetBody(v *ListMultiAccountTagValuesResponseBody) *ListMultiAccountTagValuesResponse {
	s.Body = v
	return s
}

type ListResourceTypesRequest struct {
	AcceptLanguage *string   `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	Query          []*string `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	ResourceType   *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetAcceptLanguage(v string) *ListResourceTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesRequest) SetQuery(v []*string) *ListResourceTypesRequest {
	s.Query = v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

type ListResourceTypesResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

type ListResourceTypesResponseBodyResourceTypes struct {
	CodeMapping      *ListResourceTypesResponseBodyResourceTypesCodeMapping `json:"CodeMapping,omitempty" xml:"CodeMapping,omitempty" type:"Struct"`
	FilterKeys       []*string                                              `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	ProductName      *string                                                `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ResourceType     *string                                                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceTypeName *string                                                `json:"ResourceTypeName,omitempty" xml:"ResourceTypeName,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCodeMapping(v *ListResourceTypesResponseBodyResourceTypesCodeMapping) *ListResourceTypesResponseBodyResourceTypes {
	s.CodeMapping = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterKeys(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterKeys = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceTypeName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceTypeName = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypesCodeMapping struct {
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetResourceGroup(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.ResourceGroup = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetTag(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.Tag = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetMatchType(v string) *ListTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetTagKey(v string) *ListTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListTagKeysResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
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

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue   *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMatchType(v string) *ListTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetTagValue(v string) *ListTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListTagValuesResponseBody struct {
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
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

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type SearchMultiAccountResourcesRequest struct {
	Filter        []*SearchMultiAccountResourcesRequestFilter      `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults    *int32                                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Scope         *string                                          `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SortCriterion *SearchMultiAccountResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchMultiAccountResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequest) SetFilter(v []*SearchMultiAccountResourcesRequestFilter) *SearchMultiAccountResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetMaxResults(v int32) *SearchMultiAccountResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetNextToken(v string) *SearchMultiAccountResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetScope(v string) *SearchMultiAccountResourcesRequest {
	s.Scope = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetSortCriterion(v *SearchMultiAccountResourcesRequestSortCriterion) *SearchMultiAccountResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchMultiAccountResourcesRequestFilter struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestFilter) SetKey(v string) *SearchMultiAccountResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetMatchType(v string) *SearchMultiAccountResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetValue(v []*string) *SearchMultiAccountResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchMultiAccountResourcesRequestSortCriterion struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchMultiAccountResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetKey(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetOrder(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchMultiAccountResourcesResponseBody struct {
	Filters    []*SearchMultiAccountResourcesResponseBodyFilters   `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	MaxResults *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*SearchMultiAccountResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	Scope      *string                                             `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBody) GoString() string {
	return s.String()
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

type SearchMultiAccountResourcesResponseBodyFilters struct {
	Key       *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Values    []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyFilters) GoString() string {
	return s.String()
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

type SearchMultiAccountResourcesResponseBodyResources struct {
	AccountId       *string                                                 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CreateTime      *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IpAddresses     []*string                                               `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	RegionId        *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string                                                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName    *string                                                 `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType    *string                                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags            []*SearchMultiAccountResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneId          *string                                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetAccountId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetCreateTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.CreateTime = &v
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

type SearchMultiAccountResourcesResponseBodyResourcesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetKey(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetValue(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchMultiAccountResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchMultiAccountResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchMultiAccountResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponse) SetHeaders(v map[string]*string) *SearchMultiAccountResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetStatusCode(v int32) *SearchMultiAccountResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetBody(v *SearchMultiAccountResourcesResponseBody) *SearchMultiAccountResourcesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcecenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DisableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableMultiAccountResourceCenter() (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.DisableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableResourceCenter() (_result *DisableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.DisableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableMultiAccountResourceCenter() (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.EnableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableResourceCenter() (_result *EnableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.EnableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultiAccountResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultiAccountResourceCenterServiceStatus() (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.GetMultiAccountResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultiAccountResourceConfigurationWithOptions(request *GetMultiAccountResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultiAccountResourceConfiguration(request *GetMultiAccountResourceConfigurationRequest) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.GetMultiAccountResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceCenterServiceStatus() (_result *GetResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.GetResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceConfigurationWithOptions(request *GetResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceConfiguration(request *GetResourceConfigurationRequest) (_result *GetResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.GetResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMultiAccountResourceGroupsWithOptions(request *ListMultiAccountResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountResourceGroups"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMultiAccountResourceGroups(request *ListMultiAccountResourceGroupsRequest) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.ListMultiAccountResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMultiAccountTagKeysWithOptions(request *ListMultiAccountTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMultiAccountTagKeys(request *ListMultiAccountTagKeysRequest) (_result *ListMultiAccountTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.ListMultiAccountTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMultiAccountTagValuesWithOptions(request *ListMultiAccountTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMultiAccountTagValues(request *ListMultiAccountTagValuesRequest) (_result *ListMultiAccountTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.ListMultiAccountTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SearchMultiAccountResourcesWithOptions(request *SearchMultiAccountResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchMultiAccountResources"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchMultiAccountResources(request *SearchMultiAccountResourcesRequest) (_result *SearchMultiAccountResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.SearchMultiAccountResourcesWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchResources"),
		Version:     tea.String("2022-12-01"),
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
