// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ModifyFullLogTtlRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ttl             *int    `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) SetSourceIp(v string) *ModifyFullLogTtlRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetLang(v string) *ModifyFullLogTtlRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyFullLogTtlResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyFullLogTtlResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponse) SetRequestId(v string) *ModifyFullLogTtlResponse {
	s.RequestId = &v
	return s
}

type ReleaseValueAddedRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ReleaseValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedRequest) SetSourceIp(v string) *ReleaseValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseValueAddedRequest) SetInstanceId(v string) *ReleaseValueAddedRequest {
	s.InstanceId = &v
	return s
}

type ReleaseValueAddedResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleaseValueAddedResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedResponse) SetRequestId(v string) *ReleaseValueAddedResponse {
	s.RequestId = &v
	return s
}

type ListValueAddedRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ListValueAddedRequest) SetSourceIp(v string) *ListValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ListValueAddedRequest) SetResourceGroupId(v string) *ListValueAddedRequest {
	s.ResourceGroupId = &v
	return s
}

type ListValueAddedResponse struct {
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ValueAddedList []*ListValueAddedResponseValueAddedList `json:"ValueAddedList,omitempty" xml:"ValueAddedList,omitempty" require:"true" type:"Repeated"`
}

func (s ListValueAddedResponse) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponse) SetRequestId(v string) *ListValueAddedResponse {
	s.RequestId = &v
	return s
}

func (s *ListValueAddedResponse) SetValueAddedList(v []*ListValueAddedResponseValueAddedList) *ListValueAddedResponse {
	s.ValueAddedList = v
	return s
}

type ListValueAddedResponseValueAddedList struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Status     *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	LogSize    *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty" require:"true"`
}

func (s ListValueAddedResponseValueAddedList) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponseValueAddedList) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseValueAddedList) SetInstanceId(v string) *ListValueAddedResponseValueAddedList {
	s.InstanceId = &v
	return s
}

func (s *ListValueAddedResponseValueAddedList) SetStatus(v int) *ListValueAddedResponseValueAddedList {
	s.Status = &v
	return s
}

func (s *ListValueAddedResponseValueAddedList) SetExpireTime(v int64) *ListValueAddedResponseValueAddedList {
	s.ExpireTime = &v
	return s
}

func (s *ListValueAddedResponseValueAddedList) SetGmtCreate(v int64) *ListValueAddedResponseValueAddedList {
	s.GmtCreate = &v
	return s
}

func (s *ListValueAddedResponseValueAddedList) SetLogSize(v int64) *ListValueAddedResponseValueAddedList {
	s.LogSize = &v
	return s
}

type ListLayer7CustomPortsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListLayer7CustomPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsRequest) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsRequest) SetSourceIp(v string) *ListLayer7CustomPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetLang(v string) *ListLayer7CustomPortsRequest {
	s.Lang = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetResourceGroupId(v string) *ListLayer7CustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListLayer7CustomPortsResponse struct {
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Layer7CustomPorts []*ListLayer7CustomPortsResponseLayer7CustomPorts `json:"Layer7CustomPorts,omitempty" xml:"Layer7CustomPorts,omitempty" require:"true" type:"Repeated"`
}

func (s ListLayer7CustomPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponse) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponse) SetRequestId(v string) *ListLayer7CustomPortsResponse {
	s.RequestId = &v
	return s
}

func (s *ListLayer7CustomPortsResponse) SetLayer7CustomPorts(v []*ListLayer7CustomPortsResponseLayer7CustomPorts) *ListLayer7CustomPortsResponse {
	s.Layer7CustomPorts = v
	return s
}

type ListLayer7CustomPortsResponseLayer7CustomPorts struct {
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty" require:"true"`
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" require:"true" type:"Repeated"`
}

func (s ListLayer7CustomPortsResponseLayer7CustomPorts) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponseLayer7CustomPorts) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseLayer7CustomPorts) SetProxyType(v string) *ListLayer7CustomPortsResponseLayer7CustomPorts {
	s.ProxyType = &v
	return s
}

func (s *ListLayer7CustomPortsResponseLayer7CustomPorts) SetProxyPorts(v []*string) *ListLayer7CustomPortsResponseLayer7CustomPorts {
	s.ProxyPorts = v
	return s
}

type DescribeSimpleDomainsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeSimpleDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsRequest) SetSourceIp(v string) *DescribeSimpleDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetLang(v string) *DescribeSimpleDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetResourceGroupId(v string) *DescribeSimpleDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetInstanceIds(v []*string) *DescribeSimpleDomainsRequest {
	s.InstanceIds = v
	return s
}

type DescribeSimpleDomainsResponse struct {
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSimpleDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponse) SetRequestId(v string) *DescribeSimpleDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSimpleDomainsResponse) SetDomainList(v []*string) *DescribeSimpleDomainsResponse {
	s.DomainList = v
	return s
}

type DescribeDefenseCountStatisticsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDefenseCountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsRequest) SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDefenseCountStatisticsResponse struct {
	RequestId              *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDefenseCountStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponse) SetRequestId(v string) *DescribeDefenseCountStatisticsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseDefenseCountStatistics) *DescribeDefenseCountStatisticsResponse {
	s.DefenseCountStatistics = v
	return s
}

type DescribeDefenseCountStatisticsResponseDefenseCountStatistics struct {
	DefenseCountTotalUsageOfCurrentMonth *int `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty" require:"true"`
	FlowPackCountRemain                  *int `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty" require:"true"`
	MaxUsableDefenseCountCurrentMonth    *int `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty" require:"true"`
}

func (s DescribeDefenseCountStatisticsResponseDefenseCountStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseDefenseCountStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseDefenseCountStatistics) SetDefenseCountTotalUsageOfCurrentMonth(v int) *DescribeDefenseCountStatisticsResponseDefenseCountStatistics {
	s.DefenseCountTotalUsageOfCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseDefenseCountStatistics) SetFlowPackCountRemain(v int) *DescribeDefenseCountStatisticsResponseDefenseCountStatistics {
	s.FlowPackCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseDefenseCountStatistics) SetMaxUsableDefenseCountCurrentMonth(v int) *DescribeDefenseCountStatisticsResponseDefenseCountStatistics {
	s.MaxUsableDefenseCountCurrentMonth = &v
	return s
}

type UntagResourcesRequest struct {
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag             []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type ListTagResourcesRequest struct {
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	NextToken       *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponse struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TagResources *ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Struct"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v *ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	TagResource []*ListTagResourcesResponseTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetTagResource(v []*ListTagResourcesResponseTagResourcesTagResource) *ListTagResourcesResponseTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagKeysRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

type ListTagKeysResponse struct {
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CurrentPage *int                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount  *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TagKeys     []*ListTagKeysResponseTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetRequestId(v string) *ListTagKeysResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponse) SetCurrentPage(v int) *ListTagKeysResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponse) SetPageSize(v int) *ListTagKeysResponse {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponse) SetTotalCount(v int) *ListTagKeysResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponse) SetTagKeys(v []*ListTagKeysResponseTagKeys) *ListTagKeysResponse {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseTagKeys struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagCount *int    `json:"TagCount,omitempty" xml:"TagCount,omitempty" require:"true"`
}

func (s ListTagKeysResponseTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseTagKeys) SetTagKey(v string) *ListTagKeysResponseTagKeys {
	s.TagKey = &v
	return s
}

func (s *ListTagKeysResponseTagKeys) SetTagCount(v int) *ListTagKeysResponseTagKeys {
	s.TagCount = &v
	return s
}

type DescribeDomainQpsWithCacheRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeDomainQpsWithCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheRequest) SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetDomain(v string) *DescribeDomainQpsWithCacheRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainQpsWithCacheResponse struct {
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Interval      *int      `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Totals        []*string `json:"Totals,omitempty" xml:"Totals,omitempty" require:"true" type:"Repeated"`
	Blocks        []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" require:"true" type:"Repeated"`
	CacheHits     []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" require:"true" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" require:"true" type:"Repeated"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" require:"true" type:"Repeated"`
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" require:"true" type:"Repeated"`
	CcJsQps       []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" require:"true" type:"Repeated"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" require:"true" type:"Repeated"`
	CcBlockQps    []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainQpsWithCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponse) SetRequestId(v string) *DescribeDomainQpsWithCacheResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetInterval(v int) *DescribeDomainQpsWithCacheResponse {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetStartTime(v int64) *DescribeDomainQpsWithCacheResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetTotals(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.CacheHits = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponse {
	s.CcBlockQps = v
	return s
}

type DescribeLogStoreExistStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLogStoreExistStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) SetSourceIp(v string) *DescribeLogStoreExistStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetLang(v string) *DescribeLogStoreExistStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeLogStoreExistStatusResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ExistStatus *bool   `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty" require:"true"`
}

func (s DescribeLogStoreExistStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponse) SetRequestId(v string) *DescribeLogStoreExistStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetExistStatus(v bool) *DescribeLogStoreExistStatusResponse {
	s.ExistStatus = &v
	return s
}

type DescribeBatchSlsDispatchStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetSourceIp(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetLang(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetResourceGroupId(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageNo(v int) *DescribeBatchSlsDispatchStatusRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageSize(v int) *DescribeBatchSlsDispatchStatusRequest {
	s.PageSize = &v
	return s
}

type DescribeBatchSlsDispatchStatusResponse struct {
	RequestId           *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount          *int                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	SlsConfigStatusList []*DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList `json:"SlsConfigStatusList,omitempty" xml:"SlsConfigStatusList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBatchSlsDispatchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetTotalCount(v int) *DescribeBatchSlsDispatchStatusResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetSlsConfigStatusList(v []*DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList) *DescribeBatchSlsDispatchStatusResponse {
	s.SlsConfigStatusList = v
	return s
}

type DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList struct {
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty" require:"true"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList) SetEnable(v bool) *DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList {
	s.Enable = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList) SetDomain(v string) *DescribeBatchSlsDispatchStatusResponseSlsConfigStatusList {
	s.Domain = &v
	return s
}

type DescribeSlsEmptyCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsEmptyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountRequest) SetSourceIp(v string) *DescribeSlsEmptyCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetLang(v string) *DescribeSlsEmptyCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetResourceGroupId(v string) *DescribeSlsEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsEmptyCountResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AvailableCount *int    `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty" require:"true"`
}

func (s DescribeSlsEmptyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponse) SetRequestId(v string) *DescribeSlsEmptyCountResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsEmptyCountResponse) SetAvailableCount(v int) *DescribeSlsEmptyCountResponse {
	s.AvailableCount = &v
	return s
}

type EmptySlsLogstoreRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreRequest) SetSourceIp(v string) *EmptySlsLogstoreRequest {
	s.SourceIp = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetLang(v string) *EmptySlsLogstoreRequest {
	s.Lang = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
	s.ResourceGroupId = &v
	return s
}

type EmptySlsLogstoreResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EmptySlsLogstoreResponse) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreResponse) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreResponse) SetRequestId(v string) *EmptySlsLogstoreResponse {
	s.RequestId = &v
	return s
}

type CloseDomainSlsConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s CloseDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigRequest) SetSourceIp(v string) *CloseDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetLang(v string) *CloseDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetResourceGroupId(v string) *CloseDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetDomain(v string) *CloseDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

type CloseDomainSlsConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CloseDomainSlsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigResponse) SetRequestId(v string) *CloseDomainSlsConfigResponse {
	s.RequestId = &v
	return s
}

type OpenDomainSlsConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s OpenDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigRequest) SetSourceIp(v string) *OpenDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetLang(v string) *OpenDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetResourceGroupId(v string) *OpenDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetDomain(v string) *OpenDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

type OpenDomainSlsConfigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s OpenDomainSlsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigResponse) SetRequestId(v string) *OpenDomainSlsConfigResponse {
	s.RequestId = &v
	return s
}

type DescribeSlsLogstoreInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetLang(v string) *DescribeSlsLogstoreInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsLogstoreInfoResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Quota     *int64  `json:"Quota,omitempty" xml:"Quota,omitempty" require:"true"`
	LogStore  *string `json:"LogStore,omitempty" xml:"LogStore,omitempty" require:"true"`
	Used      *int64  `json:"Used,omitempty" xml:"Used,omitempty" require:"true"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty" require:"true"`
	Ttl       *int    `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
}

func (s DescribeSlsLogstoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponse) SetRequestId(v string) *DescribeSlsLogstoreInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetQuota(v int64) *DescribeSlsLogstoreInfoResponse {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetLogStore(v string) *DescribeSlsLogstoreInfoResponse {
	s.LogStore = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetUsed(v int64) *DescribeSlsLogstoreInfoResponse {
	s.Used = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetProject(v string) *DescribeSlsLogstoreInfoResponse {
	s.Project = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetTtl(v int) *DescribeSlsLogstoreInfoResponse {
	s.Ttl = &v
	return s
}

type DescribeSlsAuthStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) SetSourceIp(v string) *DescribeSlsAuthStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetLang(v string) *DescribeSlsAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsAuthStatusResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SlsAuthStatus *bool   `json:"SlsAuthStatus,omitempty" xml:"SlsAuthStatus,omitempty" require:"true"`
}

func (s DescribeSlsAuthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponse) SetRequestId(v string) *DescribeSlsAuthStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetSlsAuthStatus(v bool) *DescribeSlsAuthStatusResponse {
	s.SlsAuthStatus = &v
	return s
}

type DescribeSlsOpenStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) SetSourceIp(v string) *DescribeSlsOpenStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetLang(v string) *DescribeSlsOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsOpenStatusResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SlsOpenStatus *bool   `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty" require:"true"`
}

func (s DescribeSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponse) SetRequestId(v string) *DescribeSlsOpenStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetSlsOpenStatus(v bool) *DescribeSlsOpenStatusResponse {
	s.SlsOpenStatus = &v
	return s
}

type DescribeDomainSlsStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DescribeDomainSlsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusRequest) SetSourceIp(v string) *DescribeDomainSlsStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetLang(v string) *DescribeDomainSlsStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetResourceGroupId(v string) *DescribeDomainSlsStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetDomain(v string) *DescribeDomainSlsStatusRequest {
	s.Domain = &v
	return s
}

type DescribeDomainSlsStatusResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SlsStatus   *bool   `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty" require:"true"`
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty" require:"true"`
	SlsProject  *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty" require:"true"`
}

func (s DescribeDomainSlsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponse) SetRequestId(v string) *DescribeDomainSlsStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetSlsStatus(v bool) *DescribeDomainSlsStatusResponse {
	s.SlsStatus = &v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetSlsLogstore(v string) *DescribeDomainSlsStatusResponse {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetSlsProject(v string) *DescribeDomainSlsStatusResponse {
	s.SlsProject = &v
	return s
}

type ConfigDomainAccessModeRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	AccessMode *int    `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" require:"true"`
}

func (s ConfigDomainAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *ConfigDomainAccessModeRequest) SetSourceIp(v string) *ConfigDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigDomainAccessModeRequest) SetDomain(v string) *ConfigDomainAccessModeRequest {
	s.Domain = &v
	return s
}

func (s *ConfigDomainAccessModeRequest) SetAccessMode(v int) *ConfigDomainAccessModeRequest {
	s.AccessMode = &v
	return s
}

type ConfigDomainAccessModeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigDomainAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigDomainAccessModeResponse) GoString() string {
	return s.String()
}

func (s *ConfigDomainAccessModeResponse) SetRequestId(v string) *ConfigDomainAccessModeResponse {
	s.RequestId = &v
	return s
}

type DescribeDomainAccessModeRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeRequest) SetSourceIp(v string) *DescribeDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAccessModeRequest) SetDomainList(v []*string) *DescribeDomainAccessModeRequest {
	s.DomainList = v
	return s
}

type DescribeDomainAccessModeResponse struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DomainModeList []*DescribeDomainAccessModeResponseDomainModeList `json:"DomainModeList,omitempty" xml:"DomainModeList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponse) SetRequestId(v string) *DescribeDomainAccessModeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAccessModeResponse) SetDomainModeList(v []*DescribeDomainAccessModeResponseDomainModeList) *DescribeDomainAccessModeResponse {
	s.DomainModeList = v
	return s
}

type DescribeDomainAccessModeResponseDomainModeList struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	AccessMode *int    `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" require:"true"`
}

func (s DescribeDomainAccessModeResponseDomainModeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponseDomainModeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseDomainModeList) SetDomain(v string) *DescribeDomainAccessModeResponseDomainModeList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAccessModeResponseDomainModeList) SetAccessMode(v int) *DescribeDomainAccessModeResponseDomainModeList {
	s.AccessMode = &v
	return s
}

type DeleteAsyncTaskRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId          *int    `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s DeleteAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskRequest) SetResourceGroupId(v string) *DeleteAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetTaskId(v int) *DeleteAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteAsyncTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponse) SetRequestId(v string) *DeleteAsyncTaskResponse {
	s.RequestId = &v
	return s
}

type CreateAsyncTaskRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskType        *int    `json:"TaskType,omitempty" xml:"TaskType,omitempty" require:"true"`
	TaskParams      *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty" require:"true"`
}

func (s CreateAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

type CreateAsyncTaskResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponse) SetRequestId(v string) *CreateAsyncTaskResponse {
	s.RequestId = &v
	return s
}

type ListAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s ListAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskRequest) SetSourceIp(v string) *ListAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *ListAsyncTaskRequest) SetLang(v string) *ListAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListAsyncTaskRequest) SetResourceGroupId(v string) *ListAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageNo(v int) *ListAsyncTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageSize(v int) *ListAsyncTaskRequest {
	s.PageSize = &v
	return s
}

type ListAsyncTaskResponse struct {
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total      *int                               `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	AsyncTasks []*ListAsyncTaskResponseAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" require:"true" type:"Repeated"`
}

func (s ListAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponse) SetRequestId(v string) *ListAsyncTaskResponse {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTaskResponse) SetTotal(v int) *ListAsyncTaskResponse {
	s.Total = &v
	return s
}

func (s *ListAsyncTaskResponse) SetAsyncTasks(v []*ListAsyncTaskResponseAsyncTasks) *ListAsyncTaskResponse {
	s.AsyncTasks = v
	return s
}

type ListAsyncTaskResponseAsyncTasks struct {
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	TaskStatus *int    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" require:"true"`
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty" require:"true"`
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty" require:"true"`
	TaskType   *int    `json:"TaskType,omitempty" xml:"TaskType,omitempty" require:"true"`
}

func (s ListAsyncTaskResponseAsyncTasks) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponseAsyncTasks) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseAsyncTasks) SetTaskId(v int64) *ListAsyncTaskResponseAsyncTasks {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetEndTime(v int64) *ListAsyncTaskResponseAsyncTasks {
	s.EndTime = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetStartTime(v int64) *ListAsyncTaskResponseAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetTaskStatus(v int) *ListAsyncTaskResponseAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetTaskResult(v string) *ListAsyncTaskResponseAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetTaskParams(v string) *ListAsyncTaskResponseAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *ListAsyncTaskResponseAsyncTasks) SetTaskType(v int) *ListAsyncTaskResponseAsyncTasks {
	s.TaskType = &v
	return s
}

type EnableLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s EnableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleRequest) SetSourceIp(v string) *EnableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetResourceGroupId(v string) *EnableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetDomain(v string) *EnableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

type EnableLayer7CCRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleResponse) SetRequestId(v string) *EnableLayer7CCRuleResponse {
	s.RequestId = &v
	return s
}

type EnableLayer7CCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s EnableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRequest) SetSourceIp(v string) *EnableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableLayer7CCRequest) SetResourceGroupId(v string) *EnableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRequest) SetDomain(v string) *EnableLayer7CCRequest {
	s.Domain = &v
	return s
}

type EnableLayer7CCResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableLayer7CCResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCResponse) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCResponse) SetRequestId(v string) *EnableLayer7CCResponse {
	s.RequestId = &v
	return s
}

type DisableLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DisableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleRequest) SetSourceIp(v string) *DisableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetResourceGroupId(v string) *DisableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetDomain(v string) *DisableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

type DisableLayer7CCRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleResponse) SetRequestId(v string) *DisableLayer7CCRuleResponse {
	s.RequestId = &v
	return s
}

type DisableLayer7CCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DisableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRequest) SetSourceIp(v string) *DisableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRequest) SetResourceGroupId(v string) *DisableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRequest) SetDomain(v string) *DisableLayer7CCRequest {
	s.Domain = &v
	return s
}

type DisableLayer7CCResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableLayer7CCResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCResponse) SetRequestId(v string) *DisableLayer7CCResponse {
	s.RequestId = &v
	return s
}

type DescribleLayer7InstanceRelationsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DomainList      []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsRequest) SetSourceIp(v string) *DescribleLayer7InstanceRelationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetResourceGroupId(v string) *DescribleLayer7InstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetDomainList(v []*string) *DescribleLayer7InstanceRelationsRequest {
	s.DomainList = v
	return s
}

type DescribleLayer7InstanceRelationsResponse struct {
	RequestId               *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Layer7InstanceRelations []*DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations `json:"Layer7InstanceRelations,omitempty" xml:"Layer7InstanceRelations,omitempty" require:"true" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponse) SetRequestId(v string) *DescribleLayer7InstanceRelationsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) SetLayer7InstanceRelations(v []*DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations) *DescribleLayer7InstanceRelationsResponse {
	s.Layer7InstanceRelations = v
	return s
}

type DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations struct {
	Domain          *string                                                                           `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	InstanceDetails []*DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" require:"true" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations) SetDomain(v string) *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations {
	s.Domain = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations) SetInstanceDetails(v []*DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelations {
	s.InstanceDetails = v
	return s
}

type DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails struct {
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	FunctionVersion *string   `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty" require:"true"`
	EipList         []*string `json:"EipList,omitempty" xml:"EipList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) SetInstanceId(v string) *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) SetFunctionVersion(v string) *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails {
	s.FunctionVersion = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails) SetEipList(v []*string) *DescribleLayer7InstanceRelationsResponseLayer7InstanceRelationsInstanceDetails {
	s.EipList = v
	return s
}

type DescribleCertListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribleCertListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListRequest) GoString() string {
	return s.String()
}

func (s *DescribleCertListRequest) SetSourceIp(v string) *DescribleCertListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleCertListRequest) SetResourceGroupId(v string) *DescribleCertListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleCertListRequest) SetDomain(v string) *DescribleCertListRequest {
	s.Domain = &v
	return s
}

type DescribleCertListResponse struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CertList  []*DescribleCertListResponseCertList `json:"CertList,omitempty" xml:"CertList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribleCertListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponse) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponse) SetRequestId(v string) *DescribleCertListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribleCertListResponse) SetCertList(v []*DescribleCertListResponseCertList) *DescribleCertListResponse {
	s.CertList = v
	return s
}

type DescribleCertListResponseCertList struct {
	Id            *int    `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Common        *string `json:"Common,omitempty" xml:"Common,omitempty" require:"true"`
	Issuer        *string `json:"Issuer,omitempty" xml:"Issuer,omitempty" require:"true"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty" require:"true"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
	DomainRelated *bool   `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty" require:"true"`
}

func (s DescribleCertListResponseCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponseCertList) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseCertList) SetId(v int) *DescribleCertListResponseCertList {
	s.Id = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetName(v string) *DescribleCertListResponseCertList {
	s.Name = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetCommon(v string) *DescribleCertListResponseCertList {
	s.Common = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetIssuer(v string) *DescribleCertListResponseCertList {
	s.Issuer = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetStartDate(v string) *DescribleCertListResponseCertList {
	s.StartDate = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetEndDate(v string) *DescribleCertListResponseCertList {
	s.EndDate = &v
	return s
}

func (s *DescribleCertListResponseCertList) SetDomainRelated(v bool) *DescribleCertListResponseCertList {
	s.DomainRelated = &v
	return s
}

type DescribeLayer7CCRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Offset          *int    `json:"Offset,omitempty" xml:"Offset,omitempty" require:"true"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeLayer7CCRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesRequest) SetSourceIp(v string) *DescribeLayer7CCRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetResourceGroupId(v string) *DescribeLayer7CCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetDomain(v string) *DescribeLayer7CCRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetOffset(v int) *DescribeLayer7CCRulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetPageSize(v string) *DescribeLayer7CCRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeLayer7CCRulesResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total         *int64                                        `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Layer7CCRules []*DescribeLayer7CCRulesResponseLayer7CCRules `json:"Layer7CCRules,omitempty" xml:"Layer7CCRules,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLayer7CCRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponse) SetRequestId(v string) *DescribeLayer7CCRulesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetTotal(v int64) *DescribeLayer7CCRulesResponse {
	s.Total = &v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetLayer7CCRules(v []*DescribeLayer7CCRulesResponseLayer7CCRules) *DescribeLayer7CCRulesResponse {
	s.Layer7CCRules = v
	return s
}

type DescribeLayer7CCRulesResponseLayer7CCRules struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Act      *string `json:"Act,omitempty" xml:"Act,omitempty" require:"true"`
	Count    *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	Interval *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	Ttl      *int    `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty" require:"true"`
}

func (s DescribeLayer7CCRulesResponseLayer7CCRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseLayer7CCRules) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetName(v string) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Name = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetAct(v string) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Act = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetCount(v int) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Count = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetInterval(v int) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Interval = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetMode(v string) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Mode = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetTtl(v int) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseLayer7CCRules) SetUri(v string) *DescribeLayer7CCRulesResponseLayer7CCRules {
	s.Uri = &v
	return s
}

type DescribeDomainsRequest struct {
	SourceIp           *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId    *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain             *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	QueryDomainPattern *string   `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	Offset             *int      `json:"Offset,omitempty" xml:"Offset,omitempty" require:"true"`
	PageSize           *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetSourceIp(v string) *DescribeDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetQueryDomainPattern(v string) *DescribeDomainsRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeDomainsRequest) SetOffset(v int) *DescribeDomainsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v string) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
	return s
}

type DescribeDomainsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                            `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Domains   []*DescribeDomainsResponseDomains `json:"Domains,omitempty" xml:"Domains,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) SetRequestId(v string) *DescribeDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponse) SetTotal(v int64) *DescribeDomainsResponse {
	s.Total = &v
	return s
}

func (s *DescribeDomainsResponse) SetDomains(v []*DescribeDomainsResponseDomains) *DescribeDomainsResponse {
	s.Domains = v
	return s
}

type DescribeDomainsResponseDomains struct {
	Domain        *string                                        `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	CcEnabled     *bool                                          `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty" require:"true"`
	CcRuleEnabled *bool                                          `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty" require:"true"`
	CcTemplate    *string                                        `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty" require:"true"`
	SslProtocols  *string                                        `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty" require:"true"`
	SslCiphers    *string                                        `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty" require:"true"`
	Http2Enable   *bool                                          `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty" require:"true"`
	Cname         *string                                        `json:"Cname,omitempty" xml:"Cname,omitempty" require:"true"`
	CertName      *string                                        `json:"CertName,omitempty" xml:"CertName,omitempty" require:"true"`
	ProxyTypeList []*DescribeDomainsResponseDomainsProxyTypeList `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty" require:"true" type:"Repeated"`
	RealServers   []*DescribeDomainsResponseDomainsRealServers   `json:"RealServers,omitempty" xml:"RealServers,omitempty" require:"true" type:"Repeated"`
	WhiteList     []*string                                      `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" require:"true" type:"Repeated"`
	BlackList     []*string                                      `json:"BlackList,omitempty" xml:"BlackList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainsResponseDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseDomains) SetDomain(v string) *DescribeDomainsResponseDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetCcEnabled(v bool) *DescribeDomainsResponseDomains {
	s.CcEnabled = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetCcRuleEnabled(v bool) *DescribeDomainsResponseDomains {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetCcTemplate(v string) *DescribeDomainsResponseDomains {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetSslProtocols(v string) *DescribeDomainsResponseDomains {
	s.SslProtocols = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetSslCiphers(v string) *DescribeDomainsResponseDomains {
	s.SslCiphers = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetHttp2Enable(v bool) *DescribeDomainsResponseDomains {
	s.Http2Enable = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetCname(v string) *DescribeDomainsResponseDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetCertName(v string) *DescribeDomainsResponseDomains {
	s.CertName = &v
	return s
}

func (s *DescribeDomainsResponseDomains) SetProxyTypeList(v []*DescribeDomainsResponseDomainsProxyTypeList) *DescribeDomainsResponseDomains {
	s.ProxyTypeList = v
	return s
}

func (s *DescribeDomainsResponseDomains) SetRealServers(v []*DescribeDomainsResponseDomainsRealServers) *DescribeDomainsResponseDomains {
	s.RealServers = v
	return s
}

func (s *DescribeDomainsResponseDomains) SetWhiteList(v []*string) *DescribeDomainsResponseDomains {
	s.WhiteList = v
	return s
}

func (s *DescribeDomainsResponseDomains) SetBlackList(v []*string) *DescribeDomainsResponseDomains {
	s.BlackList = v
	return s
}

type DescribeDomainsResponseDomainsProxyTypeList struct {
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty" require:"true"`
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainsResponseDomainsProxyTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseDomainsProxyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseDomainsProxyTypeList) SetProxyType(v string) *DescribeDomainsResponseDomainsProxyTypeList {
	s.ProxyType = &v
	return s
}

func (s *DescribeDomainsResponseDomainsProxyTypeList) SetProxyPorts(v []*string) *DescribeDomainsResponseDomainsProxyTypeList {
	s.ProxyPorts = v
	return s
}

type DescribeDomainsResponseDomainsRealServers struct {
	RsType     *int    `json:"RsType,omitempty" xml:"RsType,omitempty" require:"true"`
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty" require:"true"`
}

func (s DescribeDomainsResponseDomainsRealServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseDomainsRealServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseDomainsRealServers) SetRsType(v int) *DescribeDomainsResponseDomainsRealServers {
	s.RsType = &v
	return s
}

func (s *DescribeDomainsResponseDomainsRealServers) SetRealServer(v string) *DescribeDomainsResponseDomainsRealServers {
	s.RealServer = &v
	return s
}

type DescribeDomainQpsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeDomainQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsRequest) SetSourceIp(v string) *DescribeDomainQpsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetResourceGroupId(v string) *DescribeDomainQpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetDomain(v string) *DescribeDomainQpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetStartTime(v int64) *DescribeDomainQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetEndTime(v int64) *DescribeDomainQpsRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainQpsResponse struct {
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Interval      *int      `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Totals        []*string `json:"Totals,omitempty" xml:"Totals,omitempty" require:"true" type:"Repeated"`
	Blocks        []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" require:"true" type:"Repeated"`
	CacheHits     []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" require:"true" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" require:"true" type:"Repeated"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" require:"true" type:"Repeated"`
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" require:"true" type:"Repeated"`
	CcJsQps       []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" require:"true" type:"Repeated"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" require:"true" type:"Repeated"`
	CcBlockQps    []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsResponse) SetRequestId(v string) *DescribeDomainQpsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsResponse) SetInterval(v int) *DescribeDomainQpsResponse {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsResponse) SetStartTime(v int64) *DescribeDomainQpsResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsResponse) SetTotals(v []*string) *DescribeDomainQpsResponse {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsResponse) SetBlocks(v []*string) *DescribeDomainQpsResponse {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsResponse) SetCacheHits(v []*string) *DescribeDomainQpsResponse {
	s.CacheHits = v
	return s
}

func (s *DescribeDomainQpsResponse) SetPreciseBlocks(v []*string) *DescribeDomainQpsResponse {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsResponse) SetRegionBlocks(v []*string) *DescribeDomainQpsResponse {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsResponse) SetIpBlockQps(v []*string) *DescribeDomainQpsResponse {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsResponse) SetCcJsQps(v []*string) *DescribeDomainQpsResponse {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsResponse) SetPreciseJsQps(v []*string) *DescribeDomainQpsResponse {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsResponse) SetCcBlockQps(v []*string) *DescribeDomainQpsResponse {
	s.CcBlockQps = v
	return s
}

type DescribeDomainAttackEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Offset          *int    `json:"Offset,omitempty" xml:"Offset,omitempty" require:"true"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeDomainAttackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsRequest) SetSourceIp(v string) *DescribeDomainAttackEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetStartTime(v int64) *DescribeDomainAttackEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetEndTime(v int64) *DescribeDomainAttackEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetDomain(v string) *DescribeDomainAttackEventsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetOffset(v int) *DescribeDomainAttackEventsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageSize(v string) *DescribeDomainAttackEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainAttackEventsResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                                      `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Events    []*DescribeDomainAttackEventsResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDomainAttackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponse) SetRequestId(v string) *DescribeDomainAttackEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetTotal(v int64) *DescribeDomainAttackEventsResponse {
	s.Total = &v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetEvents(v []*DescribeDomainAttackEventsResponseEvents) *DescribeDomainAttackEventsResponse {
	s.Events = v
	return s
}

type DescribeDomainAttackEventsResponseEvents struct {
	StartTime  *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Duration   *int   `json:"Duration,omitempty" xml:"Duration,omitempty" require:"true"`
	Finished   *bool  `json:"Finished,omitempty" xml:"Finished,omitempty" require:"true"`
	MaxQps     *int   `json:"MaxQps,omitempty" xml:"MaxQps,omitempty" require:"true"`
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty" require:"true"`
}

func (s DescribeDomainAttackEventsResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseEvents) SetStartTime(v int64) *DescribeDomainAttackEventsResponseEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseEvents) SetEndTime(v int64) *DescribeDomainAttackEventsResponseEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseEvents) SetDuration(v int) *DescribeDomainAttackEventsResponseEvents {
	s.Duration = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseEvents) SetFinished(v bool) *DescribeDomainAttackEventsResponseEvents {
	s.Finished = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseEvents) SetMaxQps(v int) *DescribeDomainAttackEventsResponseEvents {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseEvents) SetBlockCount(v int64) *DescribeDomainAttackEventsResponseEvents {
	s.BlockCount = &v
	return s
}

type DescribeBackSourceCidrRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Line            *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s DescribeBackSourceCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrRequest) SetSourceIp(v string) *DescribeBackSourceCidrRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetResourceGroupId(v string) *DescribeBackSourceCidrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

type DescribeBackSourceCidrResponse struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CidrList  []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeBackSourceCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponse) SetRequestId(v string) *DescribeBackSourceCidrResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetCidrList(v []*string) *DescribeBackSourceCidrResponse {
	s.CidrList = v
	return s
}

type DeleteLayer7RuleRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
}

func (s DeleteLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleRequest) SetResourceGroupId(v string) *DeleteLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7RuleRequest) SetDomain(v string) *DeleteLayer7RuleRequest {
	s.Domain = &v
	return s
}

type DeleteLayer7RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleResponse) SetRequestId(v string) *DeleteLayer7RuleResponse {
	s.RequestId = &v
	return s
}

type DeleteLayer7CCRuleRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s DeleteLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleRequest) SetResourceGroupId(v string) *DeleteLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetDomain(v string) *DeleteLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetName(v string) *DeleteLayer7CCRuleRequest {
	s.Name = &v
	return s
}

type DeleteLayer7CCRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleResponse) SetRequestId(v string) *DeleteLayer7CCRuleResponse {
	s.RequestId = &v
	return s
}

type CreateLayer7RuleRequest struct {
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	RsType          *int      `json:"RsType,omitempty" xml:"RsType,omitempty" require:"true"`
	Rules           *string   `json:"Rules,omitempty" xml:"Rules,omitempty" require:"true"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleRequest) SetResourceGroupId(v string) *CreateLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetDomain(v string) *CreateLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRsType(v int) *CreateLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRules(v string) *CreateLayer7RuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetInstanceIds(v []*string) *CreateLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

type CreateLayer7RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleResponse) SetRequestId(v string) *CreateLayer7RuleResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer7RuleRequest struct {
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	ProxyTypeList   *string   `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty"`
	ProxyTypes      []*string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	RsType          *int      `json:"RsType,omitempty" xml:"RsType,omitempty" require:"true"`
	RealServers     []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" require:"true" type:"Repeated"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s ConfigLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleRequest) SetResourceGroupId(v string) *ConfigLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetDomain(v string) *ConfigLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypeList(v string) *ConfigLayer7RuleRequest {
	s.ProxyTypeList = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypes(v []*string) *ConfigLayer7RuleRequest {
	s.ProxyTypes = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRsType(v int) *ConfigLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRealServers(v []*string) *ConfigLayer7RuleRequest {
	s.RealServers = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetInstanceIds(v []*string) *ConfigLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

type ConfigLayer7RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleResponse) SetRequestId(v string) *ConfigLayer7RuleResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer7CertRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	CertId          *int    `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cert            *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ConfigLayer7CertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertRequest) SetResourceGroupId(v string) *ConfigLayer7CertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetDomain(v string) *ConfigLayer7CertRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertId(v int) *ConfigLayer7CertRequest {
	s.CertId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertName(v string) *ConfigLayer7CertRequest {
	s.CertName = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCert(v string) *ConfigLayer7CertRequest {
	s.Cert = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetKey(v string) *ConfigLayer7CertRequest {
	s.Key = &v
	return s
}

type ConfigLayer7CertResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer7CertResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertResponse) SetRequestId(v string) *ConfigLayer7CertResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer7CCTemplateRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty" require:"true"`
}

func (s ConfigLayer7CCTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateRequest) SetResourceGroupId(v string) *ConfigLayer7CCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetDomain(v string) *ConfigLayer7CCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetTemplate(v string) *ConfigLayer7CCTemplateRequest {
	s.Template = &v
	return s
}

type ConfigLayer7CCTemplateResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer7CCTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateResponse) SetRequestId(v string) *ConfigLayer7CCTemplateResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer7CCRuleRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty" require:"true"`
	Count           *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	Interval        *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	Ttl             *int    `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty" require:"true"`
}

func (s ConfigLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleRequest) SetResourceGroupId(v string) *ConfigLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetDomain(v string) *ConfigLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetName(v string) *ConfigLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetAct(v string) *ConfigLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetCount(v int) *ConfigLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetInterval(v int) *ConfigLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetMode(v string) *ConfigLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetTtl(v int) *ConfigLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetUri(v string) *ConfigLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

type ConfigLayer7CCRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleResponse) SetRequestId(v string) *ConfigLayer7CCRuleResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer7BlackWhiteListRequest struct {
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	BlackList       []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	WhiteList       []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigLayer7BlackWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListRequest) SetResourceGroupId(v string) *ConfigLayer7BlackWhiteListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetDomain(v string) *ConfigLayer7BlackWhiteListRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetBlackList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.BlackList = v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetWhiteList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.WhiteList = v
	return s
}

type ConfigLayer7BlackWhiteListResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer7BlackWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListResponse) SetRequestId(v string) *ConfigLayer7BlackWhiteListResponse {
	s.RequestId = &v
	return s
}

type AddLayer7CCRuleRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty" require:"true"`
	Count           *int    `json:"Count,omitempty" xml:"Count,omitempty" require:"true"`
	Interval        *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	Ttl             *int    `json:"Ttl,omitempty" xml:"Ttl,omitempty" require:"true"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty" require:"true"`
}

func (s AddLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleRequest) SetResourceGroupId(v string) *AddLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetDomain(v string) *AddLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetName(v string) *AddLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetAct(v string) *AddLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetCount(v int) *AddLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetInterval(v int) *AddLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetMode(v string) *AddLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetTtl(v int) *AddLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetUri(v string) *AddLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

type AddLayer7CCRuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleResponse) SetRequestId(v string) *AddLayer7CCRuleResponse {
	s.RequestId = &v
	return s
}

type ReleaseInstanceRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetSourceIp(v string) *ReleaseInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetRequestId(v string) *ReleaseInstanceResponse {
	s.RequestId = &v
	return s
}

type ModifyInstanceRemarkRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyInstanceRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkRequest) SetSourceIp(v string) *ModifyInstanceRemarkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

type ModifyInstanceRemarkResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyInstanceRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponse) SetRequestId(v string) *ModifyInstanceRemarkResponse {
	s.RequestId = &v
	return s
}

type ModifyElasticBandWidthRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ElasticBandwidth *int    `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty" require:"true"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) SetSourceIp(v string) *ModifyElasticBandWidthRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

type ModifyElasticBandWidthResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyElasticBandWidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponse) SetRequestId(v string) *ModifyElasticBandWidthResponse {
	s.RequestId = &v
	return s
}

type DescribeOpEntitiesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EntityType      *int    `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject    *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageNo(v int) *DescribeOpEntitiesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total      *int64                                  `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	OpEntities []*DescribeOpEntitiesResponseOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetRequestId(v string) *DescribeOpEntitiesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetTotal(v int64) *DescribeOpEntitiesResponse {
	s.Total = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetOpEntities(v []*DescribeOpEntitiesResponseOpEntities) *DescribeOpEntitiesResponse {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseOpEntities struct {
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	EntityType   *int    `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty" require:"true"`
	OpAction     *int    `json:"OpAction,omitempty" xml:"OpAction,omitempty" require:"true"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty" require:"true"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponseOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityType(v int) *DescribeOpEntitiesResponseOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAction(v int) *DescribeOpEntitiesResponseOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeLayer4RulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int    `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Offset          *int    `json:"Offset,omitempty" xml:"Offset,omitempty" require:"true"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesRequest) SetSourceIp(v string) *DescribeLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetInstanceId(v string) *DescribeLayer4RulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetForwardProtocol(v string) *DescribeLayer4RulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetFrontendPort(v int) *DescribeLayer4RulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetOffset(v int) *DescribeLayer4RulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetPageSize(v string) *DescribeLayer4RulesRequest {
	s.PageSize = &v
	return s
}

type DescribeLayer4RulesResponse struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                                  `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Listeners []*DescribeLayer4RulesResponseListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponse) SetRequestId(v string) *DescribeLayer4RulesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RulesResponse) SetTotal(v int64) *DescribeLayer4RulesResponse {
	s.Total = &v
	return s
}

func (s *DescribeLayer4RulesResponse) SetListeners(v []*DescribeLayer4RulesResponseListeners) *DescribeLayer4RulesResponse {
	s.Listeners = v
	return s
}

type DescribeLayer4RulesResponseListeners struct {
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Protocol     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	FrontendPort *int      `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	BackendPort  *int      `json:"BackendPort,omitempty" xml:"BackendPort,omitempty" require:"true"`
	IsAutoCreate *bool     `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty" require:"true"`
	RealServers  []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLayer4RulesResponseListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseListeners) SetInstanceId(v string) *DescribeLayer4RulesResponseListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesResponseListeners) SetProtocol(v string) *DescribeLayer4RulesResponseListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulesResponseListeners) SetFrontendPort(v int) *DescribeLayer4RulesResponseListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseListeners) SetBackendPort(v int) *DescribeLayer4RulesResponseListeners {
	s.BackendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseListeners) SetIsAutoCreate(v bool) *DescribeLayer4RulesResponseListeners {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeLayer4RulesResponseListeners) SetRealServers(v []*string) *DescribeLayer4RulesResponseListeners {
	s.RealServers = v
	return s
}

type DescribeLayer4RuleAttributesRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesRequest) SetSourceIp(v string) *DescribeLayer4RuleAttributesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RuleAttributesRequest) SetListeners(v string) *DescribeLayer4RuleAttributesRequest {
	s.Listeners = &v
	return s
}

type DescribeLayer4RuleAttributesResponse struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Listeners []*DescribeLayer4RuleAttributesResponseListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLayer4RuleAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponse) SetRequestId(v string) *DescribeLayer4RuleAttributesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) SetListeners(v []*DescribeLayer4RuleAttributesResponseListeners) *DescribeLayer4RuleAttributesResponse {
	s.Listeners = v
	return s
}

type DescribeLayer4RuleAttributesResponseListeners struct {
	InstanceId   *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Protocol     *string                                              `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	FrontendPort *int                                                 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	Config       *DescribeLayer4RuleAttributesResponseListenersConfig `json:"Config,omitempty" xml:"Config,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLayer4RuleAttributesResponseListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListeners) SetInstanceId(v string) *DescribeLayer4RuleAttributesResponseListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListeners) SetProtocol(v string) *DescribeLayer4RuleAttributesResponseListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListeners) SetFrontendPort(v int) *DescribeLayer4RuleAttributesResponseListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListeners) SetConfig(v *DescribeLayer4RuleAttributesResponseListenersConfig) *DescribeLayer4RuleAttributesResponseListeners {
	s.Config = v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfig struct {
	PersistenceTimeout *int                                                           `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty" require:"true"`
	Synproxy           *string                                                        `json:"Synproxy,omitempty" xml:"Synproxy,omitempty" require:"true"`
	NodataConn         *string                                                        `json:"NodataConn,omitempty" xml:"NodataConn,omitempty" require:"true"`
	Sla                *DescribeLayer4RuleAttributesResponseListenersConfigSla        `json:"Sla,omitempty" xml:"Sla,omitempty" require:"true" type:"Struct"`
	Slimit             *DescribeLayer4RuleAttributesResponseListenersConfigSlimit     `json:"Slimit,omitempty" xml:"Slimit,omitempty" require:"true" type:"Struct"`
	PayloadLen         *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" require:"true" type:"Struct"`
	Cc                 *DescribeLayer4RuleAttributesResponseListenersConfigCc         `json:"Cc,omitempty" xml:"Cc,omitempty" require:"true" type:"Struct"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfig) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetPersistenceTimeout(v int) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetSynproxy(v string) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.Synproxy = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetNodataConn(v string) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.NodataConn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetSla(v *DescribeLayer4RuleAttributesResponseListenersConfigSla) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.Sla = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetSlimit(v *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.Slimit = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetPayloadLen(v *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.PayloadLen = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfig) SetCc(v *DescribeLayer4RuleAttributesResponseListenersConfigCc) *DescribeLayer4RuleAttributesResponseListenersConfig {
	s.Cc = v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfigSla struct {
	Cps           *int `json:"Cps,omitempty" xml:"Cps,omitempty" require:"true"`
	Maxconn       *int `json:"Maxconn,omitempty" xml:"Maxconn,omitempty" require:"true"`
	CpsEnable     *int `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty" require:"true"`
	MaxconnEnable *int `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigSla) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSla) SetCps(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSla) SetMaxconn(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSla {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSla) SetCpsEnable(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSla) SetMaxconnEnable(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSla {
	s.MaxconnEnable = &v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfigSlimit struct {
	Cps           *int   `json:"Cps,omitempty" xml:"Cps,omitempty" require:"true"`
	Maxconn       *int   `json:"Maxconn,omitempty" xml:"Maxconn,omitempty" require:"true"`
	CpsEnable     *int   `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty" require:"true"`
	CpsMode       *int   `json:"CpsMode,omitempty" xml:"CpsMode,omitempty" require:"true"`
	MaxconnEnable *int   `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty" require:"true"`
	Bps           *int64 `json:"Bps,omitempty" xml:"Bps,omitempty" require:"true"`
	Pps           *int64 `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigSlimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetCps(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetMaxconn(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetCpsEnable(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetCpsMode(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.CpsMode = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetMaxconnEnable(v int) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetBps(v int64) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigSlimit) SetPps(v int64) *DescribeLayer4RuleAttributesResponseListenersConfigSlimit {
	s.Pps = &v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen struct {
	Min *int `json:"Min,omitempty" xml:"Min,omitempty" require:"true"`
	Max *int `json:"Max,omitempty" xml:"Max,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen) SetMin(v int) *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen {
	s.Min = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen) SetMax(v int) *DescribeLayer4RuleAttributesResponseListenersConfigPayloadLen {
	s.Max = &v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfigCc struct {
	Sblack []*DescribeLayer4RuleAttributesResponseListenersConfigCcSblack `json:"Sblack,omitempty" xml:"Sblack,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigCc) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigCc) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigCc) SetSblack(v []*DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) *DescribeLayer4RuleAttributesResponseListenersConfigCc {
	s.Sblack = v
	return s
}

type DescribeLayer4RuleAttributesResponseListenersConfigCcSblack struct {
	During  *int `json:"During,omitempty" xml:"During,omitempty" require:"true"`
	Expires *int `json:"Expires,omitempty" xml:"Expires,omitempty" require:"true"`
	Cnt     *int `json:"Cnt,omitempty" xml:"Cnt,omitempty" require:"true"`
	Type    *int `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) SetDuring(v int) *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) SetExpires(v int) *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) SetCnt(v int) *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack {
	s.Cnt = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack) SetType(v int) *DescribeLayer4RuleAttributesResponseListenersConfigCcSblack {
	s.Type = &v
	return s
}

type DescribeIpTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Interval        *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	Port            *int    `json:"Port,omitempty" xml:"Port,omitempty"`
	QueryProtocol   *string `json:"QueryProtocol,omitempty" xml:"QueryProtocol,omitempty"`
}

func (s DescribeIpTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficRequest) SetSourceIp(v string) *DescribeIpTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetResourceGroupId(v string) *DescribeIpTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetStartTime(v int64) *DescribeIpTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetInterval(v int) *DescribeIpTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEndTime(v int64) *DescribeIpTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEip(v string) *DescribeIpTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetPort(v int) *DescribeIpTrafficRequest {
	s.Port = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetQueryProtocol(v string) *DescribeIpTrafficRequest {
	s.QueryProtocol = &v
	return s
}

type DescribeIpTrafficResponse struct {
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MaxInBps        *int64                                      `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty" require:"true"`
	AvgInBps        *int64                                      `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty" require:"true"`
	MaxOutBps       *int64                                      `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty" require:"true"`
	AvgOutBps       *int64                                      `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty" require:"true"`
	IpTrafficPoints []*DescribeIpTrafficResponseIpTrafficPoints `json:"IpTrafficPoints,omitempty" xml:"IpTrafficPoints,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeIpTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponse) SetRequestId(v string) *DescribeIpTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetMaxInBps(v int64) *DescribeIpTrafficResponse {
	s.MaxInBps = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetAvgInBps(v int64) *DescribeIpTrafficResponse {
	s.AvgInBps = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetMaxOutBps(v int64) *DescribeIpTrafficResponse {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetAvgOutBps(v int64) *DescribeIpTrafficResponse {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetIpTrafficPoints(v []*DescribeIpTrafficResponseIpTrafficPoints) *DescribeIpTrafficResponse {
	s.IpTrafficPoints = v
	return s
}

type DescribeIpTrafficResponseIpTrafficPoints struct {
	Time       *int64 `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	MaxInbps   *int64 `json:"MaxInbps,omitempty" xml:"MaxInbps,omitempty" require:"true"`
	MaxOutbps  *int64 `json:"MaxOutbps,omitempty" xml:"MaxOutbps,omitempty" require:"true"`
	Cps        *int   `json:"Cps,omitempty" xml:"Cps,omitempty" require:"true"`
	ActConns   *int   `json:"ActConns,omitempty" xml:"ActConns,omitempty" require:"true"`
	InactConns *int   `json:"InactConns,omitempty" xml:"InactConns,omitempty" require:"true"`
}

func (s DescribeIpTrafficResponseIpTrafficPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponseIpTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetTime(v int64) *DescribeIpTrafficResponseIpTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetMaxInbps(v int64) *DescribeIpTrafficResponseIpTrafficPoints {
	s.MaxInbps = &v
	return s
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetMaxOutbps(v int64) *DescribeIpTrafficResponseIpTrafficPoints {
	s.MaxOutbps = &v
	return s
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetCps(v int) *DescribeIpTrafficResponseIpTrafficPoints {
	s.Cps = &v
	return s
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetActConns(v int) *DescribeIpTrafficResponseIpTrafficPoints {
	s.ActConns = &v
	return s
}

func (s *DescribeIpTrafficResponseIpTrafficPoints) SetInactConns(v int) *DescribeIpTrafficResponseIpTrafficPoints {
	s.InactConns = &v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceStatisticsResponse struct {
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceStatistics []*DescribeInstanceStatisticsResponseInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponse) SetRequestId(v string) *DescribeInstanceStatisticsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetInstanceStatistics(v []*DescribeInstanceStatisticsResponseInstanceStatistics) *DescribeInstanceStatisticsResponse {
	s.InstanceStatistics = v
	return s
}

type DescribeInstanceStatisticsResponseInstanceStatistics struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	PortUsage         *int    `json:"PortUsage,omitempty" xml:"PortUsage,omitempty" require:"true"`
	DomainUsage       *int    `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty" require:"true"`
	SiteUsage         *int    `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty" require:"true"`
	DefenseCountUsage *int    `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty" require:"true"`
}

func (s DescribeInstanceStatisticsResponseInstanceStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseInstanceStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseInstanceStatistics) SetInstanceId(v string) *DescribeInstanceStatisticsResponseInstanceStatistics {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseInstanceStatistics) SetPortUsage(v int) *DescribeInstanceStatisticsResponseInstanceStatistics {
	s.PortUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseInstanceStatistics) SetDomainUsage(v int) *DescribeInstanceStatisticsResponseInstanceStatistics {
	s.DomainUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseInstanceStatistics) SetSiteUsage(v int) *DescribeInstanceStatisticsResponseInstanceStatistics {
	s.SiteUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseInstanceStatistics) SetDefenseCountUsage(v int) *DescribeInstanceStatisticsResponseInstanceStatistics {
	s.DefenseCountUsage = &v
	return s
}

type DescribeInstanceSpecsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceSpecs []*DescribeInstanceSpecsResponseInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetRequestId(v string) *DescribeInstanceSpecsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseInstanceSpecs) *DescribeInstanceSpecsResponse {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecsResponseInstanceSpecs struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	BaseBandwidth    *int    `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty" require:"true"`
	ElasticBandwidth *int    `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty" require:"true"`
	PortLimit        *int    `json:"PortLimit,omitempty" xml:"PortLimit,omitempty" require:"true"`
	SiteLimit        *int    `json:"SiteLimit,omitempty" xml:"SiteLimit,omitempty" require:"true"`
	DomainLimit      *int    `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty" require:"true"`
	BandwidthMbps    *int    `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty" require:"true"`
	DefenseCount     *int    `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty" require:"true"`
	FunctionVersion  *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty" require:"true"`
	QpsLimit         *int    `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponseInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetBaseBandwidth(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetElasticBandwidth(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetPortLimit(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.PortLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetSiteLimit(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.SiteLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetDomainLimit(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.DomainLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetBandwidthMbps(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.BandwidthMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetDefenseCount(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.DefenseCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetFunctionVersion(v string) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetQpsLimit(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.QpsLimit = &v
	return s
}

type DescribeInstancesRequest struct {
	SourceIp        *string                        `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     *string                        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNo          *string                        `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize        *string                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	Ip              *string                        `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark          *string                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Edition         *int                           `json:"Edition,omitempty" xml:"Edition,omitempty"`
	Enabled         *int                           `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExpireStartTime *int64                         `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	ExpireEndTime   *int64                         `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	Status          []*int                         `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNo(v string) *DescribeInstancesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesRequest) SetRemark(v string) *DescribeInstancesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesRequest) SetEdition(v int) *DescribeInstancesRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnabled(v int) *DescribeInstancesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireStartTime(v int64) *DescribeInstancesRequest {
	s.ExpireStartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireEndTime(v int64) *DescribeInstancesRequest {
	s.ExpireEndTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v []*int) *DescribeInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                                `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Instances []*DescribeInstancesResponseInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetRequestId(v string) *DescribeInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponse) SetTotal(v int64) *DescribeInstancesResponse {
	s.Total = &v
	return s
}

func (s *DescribeInstancesResponse) SetInstances(v []*DescribeInstancesResponseInstances) *DescribeInstancesResponse {
	s.Instances = v
	return s
}

type DescribeInstancesResponseInstances struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	Status     *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	DebtStatus *int    `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty" require:"true"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	Edition    *int    `json:"Edition,omitempty" xml:"Edition,omitempty" require:"true"`
	Enabled    *int    `json:"Enabled,omitempty" xml:"Enabled,omitempty" require:"true"`
}

func (s DescribeInstancesResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseInstances) SetInstanceId(v string) *DescribeInstancesResponseInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetRemark(v string) *DescribeInstancesResponseInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetStatus(v int) *DescribeInstancesResponseInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetDebtStatus(v int) *DescribeInstancesResponseInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetExpireTime(v int64) *DescribeInstancesResponseInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetGmtCreate(v int64) *DescribeInstancesResponseInstances {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetEdition(v int) *DescribeInstancesResponseInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseInstances) SetEnabled(v int) *DescribeInstancesResponseInstances {
	s.Enabled = &v
	return s
}

type DescribeInstanceDetailsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" require:"true"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) SetSourceIp(v string) *DescribeInstanceDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceDetailsResponse struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceDetails []*DescribeInstanceDetailsResponseInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponse) SetRequestId(v string) *DescribeInstanceDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetInstanceDetails(v []*DescribeInstanceDetailsResponseInstanceDetails) *DescribeInstanceDetailsResponse {
	s.InstanceDetails = v
	return s
}

type DescribeInstanceDetailsResponseInstanceDetails struct {
	InstanceId  *string                                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Line        *string                                                      `json:"Line,omitempty" xml:"Line,omitempty" require:"true"`
	EipInfoList []*DescribeInstanceDetailsResponseInstanceDetailsEipInfoList `json:"EipInfoList,omitempty" xml:"EipInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceDetailsResponseInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseInstanceDetails {
	s.Line = &v
	return s
}

func (s *DescribeInstanceDetailsResponseInstanceDetails) SetEipInfoList(v []*DescribeInstanceDetailsResponseInstanceDetailsEipInfoList) *DescribeInstanceDetailsResponseInstanceDetails {
	s.EipInfoList = v
	return s
}

type DescribeInstanceDetailsResponseInstanceDetailsEipInfoList struct {
	Eip    *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeInstanceDetailsResponseInstanceDetailsEipInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseInstanceDetailsEipInfoList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseInstanceDetailsEipInfoList) SetEip(v string) *DescribeInstanceDetailsResponseInstanceDetailsEipInfoList {
	s.Eip = &v
	return s
}

func (s *DescribeInstanceDetailsResponseInstanceDetailsEipInfoList) SetStatus(v string) *DescribeInstanceDetailsResponseInstanceDetailsEipInfoList {
	s.Status = &v
	return s
}

type DescribeHealthCheckStatusListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s DescribeHealthCheckStatusListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListRequest) SetSourceIp(v string) *DescribeHealthCheckStatusListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckStatusListRequest) SetListeners(v string) *DescribeHealthCheckStatusListRequest {
	s.Listeners = &v
	return s
}

type DescribeHealthCheckStatusListResponse struct {
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HealthCheckStatusList []*DescribeHealthCheckStatusListResponseHealthCheckStatusList `json:"HealthCheckStatusList,omitempty" xml:"HealthCheckStatusList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHealthCheckStatusListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponse) SetRequestId(v string) *DescribeHealthCheckStatusListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) SetHealthCheckStatusList(v []*DescribeHealthCheckStatusListResponseHealthCheckStatusList) *DescribeHealthCheckStatusListResponse {
	s.HealthCheckStatusList = v
	return s
}

type DescribeHealthCheckStatusListResponseHealthCheckStatusList struct {
	InstanceId           *string                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Protocol             *string                                                                           `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	FrontendPort         *int                                                                              `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	Status               *string                                                                           `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	RealServerStatusList []*DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHealthCheckStatusListResponseHealthCheckStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseHealthCheckStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusList) SetInstanceId(v string) *DescribeHealthCheckStatusListResponseHealthCheckStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusList) SetProtocol(v string) *DescribeHealthCheckStatusListResponseHealthCheckStatusList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusList) SetFrontendPort(v int) *DescribeHealthCheckStatusListResponseHealthCheckStatusList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseHealthCheckStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusList) SetRealServerStatusList(v []*DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList) *DescribeHealthCheckStatusListResponseHealthCheckStatusList {
	s.RealServerStatusList = v
	return s
}

type DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty" require:"true"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList {
	s.Address = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseHealthCheckStatusListRealServerStatusList {
	s.Status = &v
	return s
}

type DescribeHealthCheckListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s DescribeHealthCheckListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListRequest) SetSourceIp(v string) *DescribeHealthCheckListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckListRequest) SetListeners(v string) *DescribeHealthCheckListRequest {
	s.Listeners = &v
	return s
}

type DescribeHealthCheckListResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Listeners []*DescribeHealthCheckListResponseListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeHealthCheckListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponse) SetRequestId(v string) *DescribeHealthCheckListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckListResponse) SetListeners(v []*DescribeHealthCheckListResponseListeners) *DescribeHealthCheckListResponse {
	s.Listeners = v
	return s
}

type DescribeHealthCheckListResponseListeners struct {
	InstanceId   *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Protocol     *string                                              `json:"Protocol,omitempty" xml:"Protocol,omitempty" require:"true"`
	FrontendPort *int                                                 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	HealthCheck  *DescribeHealthCheckListResponseListenersHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" require:"true" type:"Struct"`
}

func (s DescribeHealthCheckListResponseListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseListeners) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseListeners) SetInstanceId(v string) *DescribeHealthCheckListResponseListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseListeners) SetProtocol(v string) *DescribeHealthCheckListResponseListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckListResponseListeners) SetFrontendPort(v int) *DescribeHealthCheckListResponseListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckListResponseListeners) SetHealthCheck(v *DescribeHealthCheckListResponseListenersHealthCheck) *DescribeHealthCheckListResponseListeners {
	s.HealthCheck = v
	return s
}

type DescribeHealthCheckListResponseListenersHealthCheck struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty" require:"true"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty" require:"true"`
	Down     *int    `json:"Down,omitempty" xml:"Down,omitempty" require:"true"`
	Interval *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Port     *int    `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	Timeout  *int    `json:"Timeout,omitempty" xml:"Timeout,omitempty" require:"true"`
	Up       *int    `json:"Up,omitempty" xml:"Up,omitempty" require:"true"`
}

func (s DescribeHealthCheckListResponseListenersHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseListenersHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetType(v string) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Uri = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetDown(v int) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetInterval(v int) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetPort(v int) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetTimeout(v int) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseListenersHealthCheck) SetUp(v int) *DescribeHealthCheckListResponseListenersHealthCheck {
	s.Up = &v
	return s
}

type DescribeElasticBandwidthSpecRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s DescribeElasticBandwidthSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecRequest) SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

type DescribeElasticBandwidthSpecResponse struct {
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeElasticBandwidthSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponse) SetRequestId(v string) *DescribeElasticBandwidthSpecResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponse {
	s.ElasticBandwidthSpec = v
	return s
}

type DescribeDDoSTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Interval        *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
}

func (s DescribeDDoSTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficRequest) SetSourceIp(v string) *DescribeDDoSTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetResourceGroupId(v string) *DescribeDDoSTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetStartTime(v int64) *DescribeDDoSTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetInterval(v int) *DescribeDDoSTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEndTime(v int64) *DescribeDDoSTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEip(v string) *DescribeDDoSTrafficRequest {
	s.Eip = &v
	return s
}

type DescribeDDoSTrafficResponse struct {
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DefenseInBytes    *int64                                          `json:"DefenseInBytes,omitempty" xml:"DefenseInBytes,omitempty" require:"true"`
	SourceInBytes     *int64                                          `json:"SourceInBytes,omitempty" xml:"SourceInBytes,omitempty" require:"true"`
	DDoSTrafficPoints []*DescribeDDoSTrafficResponseDDoSTrafficPoints `json:"DDoSTrafficPoints,omitempty" xml:"DDoSTrafficPoints,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDDoSTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponse) SetRequestId(v string) *DescribeDDoSTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetDefenseInBytes(v int64) *DescribeDDoSTrafficResponse {
	s.DefenseInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetSourceInBytes(v int64) *DescribeDDoSTrafficResponse {
	s.SourceInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetDDoSTrafficPoints(v []*DescribeDDoSTrafficResponseDDoSTrafficPoints) *DescribeDDoSTrafficResponse {
	s.DDoSTrafficPoints = v
	return s
}

type DescribeDDoSTrafficResponseDDoSTrafficPoints struct {
	Time            *int64 `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	DefenseMaxInBps *int64 `json:"DefenseMaxInBps,omitempty" xml:"DefenseMaxInBps,omitempty" require:"true"`
	SourceMaxInBps  *int64 `json:"SourceMaxInBps,omitempty" xml:"SourceMaxInBps,omitempty" require:"true"`
}

func (s DescribeDDoSTrafficResponseDDoSTrafficPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponseDDoSTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseDDoSTrafficPoints) SetTime(v int64) *DescribeDDoSTrafficResponseDDoSTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeDDoSTrafficResponseDDoSTrafficPoints) SetDefenseMaxInBps(v int64) *DescribeDDoSTrafficResponseDDoSTrafficPoints {
	s.DefenseMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseDDoSTrafficPoints) SetSourceMaxInBps(v int64) *DescribeDDoSTrafficResponseDDoSTrafficPoints {
	s.SourceMaxInBps = &v
	return s
}

type DescribeDDoSEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty" require:"true"`
	Offset          *int    `json:"Offset,omitempty" xml:"Offset,omitempty" require:"true"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) SetSourceIp(v string) *DescribeDDoSEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEip(v string) *DescribeDDoSEventsRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetOffset(v int) *DescribeDDoSEventsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageSize(v string) *DescribeDDoSEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeDDoSEventsResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                              `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Events    []*DescribeDDoSEventsResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDDoSEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponse) SetRequestId(v string) *DescribeDDoSEventsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSEventsResponse) SetTotal(v int64) *DescribeDDoSEventsResponse {
	s.Total = &v
	return s
}

func (s *DescribeDDoSEventsResponse) SetEvents(v []*DescribeDDoSEventsResponseEvents) *DescribeDDoSEventsResponse {
	s.Events = v
	return s
}

type DescribeDDoSEventsResponseEvents struct {
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Interval  *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeDDoSEventsResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseEvents) SetInterval(v int) *DescribeDDoSEventsResponseEvents {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSEventsResponseEvents) SetStatus(v string) *DescribeDDoSEventsResponseEvents {
	s.Status = &v
	return s
}

type DeleteLayer4RuleRequest struct {
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s DeleteLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleRequest) SetListeners(v string) *DeleteLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type DeleteLayer4RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponse) SetRequestId(v string) *DeleteLayer4RuleResponse {
	s.RequestId = &v
	return s
}

type CreateLayer4RuleRequest struct {
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s CreateLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequest) SetListeners(v string) *CreateLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type CreateLayer4RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponse) SetRequestId(v string) *CreateLayer4RuleResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer4RuleAttributeRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty" require:"true"`
	FrontendPort    *int    `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty" require:"true"`
}

func (s ConfigLayer4RuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeRequest) SetInstanceId(v string) *ConfigLayer4RuleAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetForwardProtocol(v string) *ConfigLayer4RuleAttributeRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetFrontendPort(v int) *ConfigLayer4RuleAttributeRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetConfig(v string) *ConfigLayer4RuleAttributeRequest {
	s.Config = &v
	return s
}

type ConfigLayer4RuleAttributeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer4RuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeResponse) SetRequestId(v string) *ConfigLayer4RuleAttributeResponse {
	s.RequestId = &v
	return s
}

type ConfigLayer4RuleRequest struct {
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty" require:"true"`
}

func (s ConfigLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleRequest) SetListeners(v string) *ConfigLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type ConfigLayer4RuleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleResponse) SetRequestId(v string) *ConfigLayer4RuleResponse {
	s.RequestId = &v
	return s
}

type ConfigHealthCheckRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty" require:"true"`
	FrontendPort    *int    `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty" require:"true"`
	HealthCheck     *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" require:"true"`
}

func (s ConfigHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckRequest) SetInstanceId(v string) *ConfigHealthCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetForwardProtocol(v string) *ConfigHealthCheckRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetFrontendPort(v int) *ConfigHealthCheckRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetHealthCheck(v string) *ConfigHealthCheckRequest {
	s.HealthCheck = &v
	return s
}

type ConfigHealthCheckResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigHealthCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckResponse) SetRequestId(v string) *ConfigHealthCheckResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddoscoo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ModifyFullLogTtlWithOptions(request *ModifyFullLogTtlRequest, runtime *util.RuntimeOptions) (_result *ModifyFullLogTtlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyFullLogTtl"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFullLogTtl(request *ModifyFullLogTtlRequest) (_result *ModifyFullLogTtlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.ModifyFullLogTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseValueAddedWithOptions(request *ReleaseValueAddedRequest, runtime *util.RuntimeOptions) (_result *ReleaseValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleaseValueAddedResponse{}
	_body, _err := client.DoRequest(tea.String("ReleaseValueAdded"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseValueAdded(request *ReleaseValueAddedRequest) (_result *ReleaseValueAddedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseValueAddedResponse{}
	_body, _err := client.ReleaseValueAddedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListValueAddedWithOptions(request *ListValueAddedRequest, runtime *util.RuntimeOptions) (_result *ListValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListValueAddedResponse{}
	_body, _err := client.DoRequest(tea.String("ListValueAdded"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListValueAdded(request *ListValueAddedRequest) (_result *ListValueAddedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListValueAddedResponse{}
	_body, _err := client.ListValueAddedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLayer7CustomPortsWithOptions(request *ListLayer7CustomPortsRequest, runtime *util.RuntimeOptions) (_result *ListLayer7CustomPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListLayer7CustomPortsResponse{}
	_body, _err := client.DoRequest(tea.String("ListLayer7CustomPorts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLayer7CustomPorts(request *ListLayer7CustomPortsRequest) (_result *ListLayer7CustomPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLayer7CustomPortsResponse{}
	_body, _err := client.ListLayer7CustomPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSimpleDomainsWithOptions(request *DescribeSimpleDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeSimpleDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSimpleDomainsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSimpleDomains"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSimpleDomains(request *DescribeSimpleDomainsRequest) (_result *DescribeSimpleDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSimpleDomainsResponse{}
	_body, _err := client.DescribeSimpleDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatisticsWithOptions(request *DescribeDefenseCountStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDefenseCountStatistics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatistics(request *DescribeDefenseCountStatisticsRequest) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DescribeDefenseCountStatisticsWithOptions(request, runtime)
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
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagKeys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeDomainQpsWithCacheWithOptions(request *DescribeDomainQpsWithCacheRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomainQpsWithCache"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithCache(request *DescribeDomainQpsWithCacheRequest) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DescribeDomainQpsWithCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatusWithOptions(request *DescribeLogStoreExistStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLogStoreExistStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatus(request *DescribeLogStoreExistStatusRequest) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DescribeLogStoreExistStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBatchSlsDispatchStatusWithOptions(request *DescribeBatchSlsDispatchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBatchSlsDispatchStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBatchSlsDispatchStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBatchSlsDispatchStatus(request *DescribeBatchSlsDispatchStatusRequest) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBatchSlsDispatchStatusResponse{}
	_body, _err := client.DescribeBatchSlsDispatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsEmptyCountWithOptions(request *DescribeSlsEmptyCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsEmptyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlsEmptyCountResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlsEmptyCount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsEmptyCount(request *DescribeSlsEmptyCountRequest) (_result *DescribeSlsEmptyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsEmptyCountResponse{}
	_body, _err := client.DescribeSlsEmptyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmptySlsLogstoreWithOptions(request *EmptySlsLogstoreRequest, runtime *util.RuntimeOptions) (_result *EmptySlsLogstoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.DoRequest(tea.String("EmptySlsLogstore"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmptySlsLogstore(request *EmptySlsLogstoreRequest) (_result *EmptySlsLogstoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.EmptySlsLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDomainSlsConfigWithOptions(request *CloseDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *CloseDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CloseDomainSlsConfigResponse{}
	_body, _err := client.DoRequest(tea.String("CloseDomainSlsConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDomainSlsConfig(request *CloseDomainSlsConfigRequest) (_result *CloseDomainSlsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDomainSlsConfigResponse{}
	_body, _err := client.CloseDomainSlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenDomainSlsConfigWithOptions(request *OpenDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *OpenDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &OpenDomainSlsConfigResponse{}
	_body, _err := client.DoRequest(tea.String("OpenDomainSlsConfig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenDomainSlsConfig(request *OpenDomainSlsConfigRequest) (_result *OpenDomainSlsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDomainSlsConfigResponse{}
	_body, _err := client.OpenDomainSlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsLogstoreInfoWithOptions(request *DescribeSlsLogstoreInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlsLogstoreInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsLogstoreInfo(request *DescribeSlsLogstoreInfoRequest) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DescribeSlsLogstoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlsAuthStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (_result *DescribeSlsAuthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DescribeSlsAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatusWithOptions(request *DescribeSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlsOpenStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatus(request *DescribeSlsOpenStatusRequest) (_result *DescribeSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DescribeSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSlsStatusWithOptions(request *DescribeDomainSlsStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSlsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainSlsStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomainSlsStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSlsStatus(request *DescribeDomainSlsStatusRequest) (_result *DescribeDomainSlsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSlsStatusResponse{}
	_body, _err := client.DescribeDomainSlsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigDomainAccessModeWithOptions(request *ConfigDomainAccessModeRequest, runtime *util.RuntimeOptions) (_result *ConfigDomainAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigDomainAccessModeResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigDomainAccessMode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigDomainAccessMode(request *ConfigDomainAccessModeRequest) (_result *ConfigDomainAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigDomainAccessModeResponse{}
	_body, _err := client.ConfigDomainAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAccessModeWithOptions(request *DescribeDomainAccessModeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainAccessModeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomainAccessMode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAccessMode(request *DescribeDomainAccessModeRequest) (_result *DescribeDomainAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAccessModeResponse{}
	_body, _err := client.DescribeDomainAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAsyncTaskWithOptions(request *DeleteAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAsyncTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAsyncTask(request *DeleteAsyncTaskRequest) (_result *DeleteAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DeleteAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAsyncTaskWithOptions(request *CreateAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAsyncTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncTask(request *CreateAsyncTaskRequest) (_result *CreateAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CreateAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsyncTaskWithOptions(request *ListAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *ListAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAsyncTaskResponse{}
	_body, _err := client.DoRequest(tea.String("ListAsyncTask"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsyncTask(request *ListAsyncTaskRequest) (_result *ListAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsyncTaskResponse{}
	_body, _err := client.ListAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLayer7CCRuleWithOptions(request *EnableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableLayer7CCRuleResponse{}
	_body, _err := client.DoRequest(tea.String("EnableLayer7CCRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLayer7CCRule(request *EnableLayer7CCRuleRequest) (_result *EnableLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLayer7CCRuleResponse{}
	_body, _err := client.EnableLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLayer7CCWithOptions(request *EnableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableLayer7CCResponse{}
	_body, _err := client.DoRequest(tea.String("EnableLayer7CC"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLayer7CC(request *EnableLayer7CCRequest) (_result *EnableLayer7CCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLayer7CCResponse{}
	_body, _err := client.EnableLayer7CCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLayer7CCRuleWithOptions(request *DisableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableLayer7CCRuleResponse{}
	_body, _err := client.DoRequest(tea.String("DisableLayer7CCRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLayer7CCRule(request *DisableLayer7CCRuleRequest) (_result *DisableLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLayer7CCRuleResponse{}
	_body, _err := client.DisableLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLayer7CCWithOptions(request *DisableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableLayer7CCResponse{}
	_body, _err := client.DoRequest(tea.String("DisableLayer7CC"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLayer7CC(request *DisableLayer7CCRequest) (_result *DisableLayer7CCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLayer7CCResponse{}
	_body, _err := client.DisableLayer7CCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribleLayer7InstanceRelationsWithOptions(request *DescribleLayer7InstanceRelationsRequest, runtime *util.RuntimeOptions) (_result *DescribleLayer7InstanceRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribleLayer7InstanceRelationsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribleLayer7InstanceRelations"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribleLayer7InstanceRelations(request *DescribleLayer7InstanceRelationsRequest) (_result *DescribleLayer7InstanceRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribleLayer7InstanceRelationsResponse{}
	_body, _err := client.DescribleLayer7InstanceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribleCertListWithOptions(request *DescribleCertListRequest, runtime *util.RuntimeOptions) (_result *DescribleCertListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribleCertListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribleCertList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribleCertList(request *DescribleCertListRequest) (_result *DescribleCertListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribleCertListResponse{}
	_body, _err := client.DescribleCertListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer7CCRulesWithOptions(request *DescribeLayer7CCRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer7CCRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLayer7CCRulesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLayer7CCRules"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer7CCRules(request *DescribeLayer7CCRulesRequest) (_result *DescribeLayer7CCRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer7CCRulesResponse{}
	_body, _err := client.DescribeLayer7CCRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomains"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithOptions(request *DescribeDomainQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainQpsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomainQps"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQps(request *DescribeDomainQpsRequest) (_result *DescribeDomainQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsResponse{}
	_body, _err := client.DescribeDomainQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAttackEventsWithOptions(request *DescribeDomainAttackEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDomainAttackEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAttackEvents(request *DescribeDomainAttackEventsRequest) (_result *DescribeDomainAttackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DescribeDomainAttackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackSourceCidrWithOptions(request *DescribeBackSourceCidrRequest, runtime *util.RuntimeOptions) (_result *DescribeBackSourceCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeBackSourceCidr"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackSourceCidr(request *DescribeBackSourceCidrRequest) (_result *DescribeBackSourceCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DescribeBackSourceCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer7RuleWithOptions(request *DeleteLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLayer7RuleResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLayer7Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer7Rule(request *DeleteLayer7RuleRequest) (_result *DeleteLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer7RuleResponse{}
	_body, _err := client.DeleteLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer7CCRuleWithOptions(request *DeleteLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLayer7CCRuleResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLayer7CCRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer7CCRule(request *DeleteLayer7CCRuleRequest) (_result *DeleteLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer7CCRuleResponse{}
	_body, _err := client.DeleteLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer7RuleWithOptions(request *CreateLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateLayer7RuleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateLayer7Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer7Rule(request *CreateLayer7RuleRequest) (_result *CreateLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer7RuleResponse{}
	_body, _err := client.CreateLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7RuleWithOptions(request *ConfigLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer7RuleResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer7Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7Rule(request *ConfigLayer7RuleRequest) (_result *ConfigLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7RuleResponse{}
	_body, _err := client.ConfigLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CertWithOptions(request *ConfigLayer7CertRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer7CertResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer7Cert"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7Cert(request *ConfigLayer7CertRequest) (_result *ConfigLayer7CertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CertResponse{}
	_body, _err := client.ConfigLayer7CertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CCTemplateWithOptions(request *ConfigLayer7CCTemplateRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer7CCTemplateResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer7CCTemplate"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7CCTemplate(request *ConfigLayer7CCTemplateRequest) (_result *ConfigLayer7CCTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CCTemplateResponse{}
	_body, _err := client.ConfigLayer7CCTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CCRuleWithOptions(request *ConfigLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer7CCRuleResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer7CCRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7CCRule(request *ConfigLayer7CCRuleRequest) (_result *ConfigLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CCRuleResponse{}
	_body, _err := client.ConfigLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7BlackWhiteListWithOptions(request *ConfigLayer7BlackWhiteListRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7BlackWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer7BlackWhiteListResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer7BlackWhiteList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7BlackWhiteList(request *ConfigLayer7BlackWhiteListRequest) (_result *ConfigLayer7BlackWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7BlackWhiteListResponse{}
	_body, _err := client.ConfigLayer7BlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLayer7CCRuleWithOptions(request *AddLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *AddLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddLayer7CCRuleResponse{}
	_body, _err := client.DoRequest(tea.String("AddLayer7CCRule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLayer7CCRule(request *AddLayer7CCRuleRequest) (_result *AddLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLayer7CCRuleResponse{}
	_body, _err := client.AddLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("ReleaseInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceRemarkWithOptions(request *ModifyInstanceRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyInstanceRemark"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (_result *ModifyInstanceRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.ModifyInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticBandWidthWithOptions(request *ModifyElasticBandWidthRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticBandWidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyElasticBandWidth"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticBandWidth(request *ModifyElasticBandWidthRequest) (_result *ModifyElasticBandWidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.ModifyElasticBandWidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOpEntities"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer4RulesWithOptions(request *DescribeLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLayer4Rules"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer4Rules(request *DescribeLayer4RulesRequest) (_result *DescribeLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DescribeLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer4RuleAttributesWithOptions(request *DescribeLayer4RuleAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RuleAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeLayer4RuleAttributesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeLayer4RuleAttributes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer4RuleAttributes(request *DescribeLayer4RuleAttributesRequest) (_result *DescribeLayer4RuleAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer4RuleAttributesResponse{}
	_body, _err := client.DescribeLayer4RuleAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpTrafficWithOptions(request *DescribeIpTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeIpTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeIpTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeIpTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpTraffic(request *DescribeIpTrafficRequest) (_result *DescribeIpTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpTrafficResponse{}
	_body, _err := client.DescribeIpTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceStatistics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (_result *DescribeInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DescribeInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceSpecs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDetailsWithOptions(request *DescribeInstanceDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceDetails"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (_result *DescribeInstanceDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DescribeInstanceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatusListWithOptions(request *DescribeHealthCheckStatusListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckStatusListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeHealthCheckStatusListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeHealthCheckStatusList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatusList(request *DescribeHealthCheckStatusListRequest) (_result *DescribeHealthCheckStatusListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckStatusListResponse{}
	_body, _err := client.DescribeHealthCheckStatusListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckListWithOptions(request *DescribeHealthCheckListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeHealthCheckList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckList(request *DescribeHealthCheckListRequest) (_result *DescribeHealthCheckListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DescribeHealthCheckListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpecWithOptions(request *DescribeElasticBandwidthSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeElasticBandwidthSpec"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpec(request *DescribeElasticBandwidthSpecRequest) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DescribeElasticBandwidthSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDoSTrafficWithOptions(request *DescribeDDoSTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDDoSTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDDoSTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDoSTraffic(request *DescribeDDoSTrafficRequest) (_result *DescribeDDoSTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSTrafficResponse{}
	_body, _err := client.DescribeDDoSTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDoSEventsWithOptions(request *DescribeDDoSEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDDoSEvents"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDoSEvents(request *DescribeDDoSEventsRequest) (_result *DescribeDDoSEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DescribeDDoSEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer4RuleWithOptions(request *DeleteLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLayer4Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer4Rule(request *DeleteLayer4RuleRequest) (_result *DeleteLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DeleteLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer4RuleWithOptions(request *CreateLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateLayer4Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer4Rule(request *CreateLayer4RuleRequest) (_result *CreateLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.CreateLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer4RuleAttributeWithOptions(request *ConfigLayer4RuleAttributeRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer4RuleAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer4RuleAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer4RuleAttribute(request *ConfigLayer4RuleAttributeRequest) (_result *ConfigLayer4RuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer4RuleAttributeResponse{}
	_body, _err := client.ConfigLayer4RuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer4RuleWithOptions(request *ConfigLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigLayer4RuleResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigLayer4Rule"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer4Rule(request *ConfigLayer4RuleRequest) (_result *ConfigLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer4RuleResponse{}
	_body, _err := client.ConfigLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigHealthCheckWithOptions(request *ConfigHealthCheckRequest, runtime *util.RuntimeOptions) (_result *ConfigHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigHealthCheckResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigHealthCheck"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-12-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigHealthCheck(request *ConfigHealthCheckRequest) (_result *ConfigHealthCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigHealthCheckResponse{}
	_body, _err := client.ConfigHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
