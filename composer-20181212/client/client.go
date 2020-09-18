// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListTemplatesRequest struct {
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetPageNumber(v int) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetTag(v string) *ListTemplatesRequest {
	s.Tag = &v
	return s
}

func (s *ListTemplatesRequest) SetLang(v string) *ListTemplatesRequest {
	s.Lang = &v
	return s
}

type ListTemplatesResponse struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Templates  []*ListTemplatesResponseTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" require:"true" type:"Repeated"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetRequestId(v string) *ListTemplatesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponse) SetTotalCount(v int) *ListTemplatesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponse) SetTemplates(v []*ListTemplatesResponseTemplates) *ListTemplatesResponse {
	s.Templates = v
	return s
}

type ListTemplatesResponseTemplates struct {
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty" require:"true"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty" require:"true"`
	TemplateTag         *string `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	TemplateConnector   *string `json:"TemplateConnector,omitempty" xml:"TemplateConnector,omitempty" require:"true"`
	TemplateSummary     *string `json:"TemplateSummary,omitempty" xml:"TemplateSummary,omitempty" require:"true"`
	TemplateSummaryEn   *string `json:"TemplateSummaryEn,omitempty" xml:"TemplateSummaryEn,omitempty" require:"true"`
	TemplateLocale      *string `json:"TemplateLocale,omitempty" xml:"TemplateLocale,omitempty" require:"true"`
	TemplateVersion     *int    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty" require:"true"`
	TemplateCreator     *string `json:"TemplateCreator,omitempty" xml:"TemplateCreator,omitempty" require:"true"`
	TemplateOverview    *string `json:"TemplateOverview,omitempty" xml:"TemplateOverview,omitempty" require:"true"`
}

func (s ListTemplatesResponseTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseTemplates) SetTemplateId(v string) *ListTemplatesResponseTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateName(v string) *ListTemplatesResponseTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateDescription(v string) *ListTemplatesResponseTemplates {
	s.TemplateDescription = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateTag(v string) *ListTemplatesResponseTemplates {
	s.TemplateTag = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetCreateTime(v string) *ListTemplatesResponseTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetUpdateTime(v string) *ListTemplatesResponseTemplates {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateConnector(v string) *ListTemplatesResponseTemplates {
	s.TemplateConnector = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateSummary(v string) *ListTemplatesResponseTemplates {
	s.TemplateSummary = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateSummaryEn(v string) *ListTemplatesResponseTemplates {
	s.TemplateSummaryEn = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateLocale(v string) *ListTemplatesResponseTemplates {
	s.TemplateLocale = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateVersion(v int) *ListTemplatesResponseTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateCreator(v string) *ListTemplatesResponseTemplates {
	s.TemplateCreator = &v
	return s
}

func (s *ListTemplatesResponseTemplates) SetTemplateOverview(v string) *ListTemplatesResponseTemplates {
	s.TemplateOverview = &v
	return s
}

type GetTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetTemplateResponse struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty" require:"true"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty" require:"true"`
	TemplateTag         *string `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" require:"true"`
	Definition          *string `json:"Definition,omitempty" xml:"Definition,omitempty" require:"true"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	TemplateConnector   *string `json:"TemplateConnector,omitempty" xml:"TemplateConnector,omitempty" require:"true"`
	TemplateSummary     *string `json:"TemplateSummary,omitempty" xml:"TemplateSummary,omitempty" require:"true"`
	TemplateSummaryEn   *string `json:"TemplateSummaryEn,omitempty" xml:"TemplateSummaryEn,omitempty" require:"true"`
	TemplateLocale      *string `json:"TemplateLocale,omitempty" xml:"TemplateLocale,omitempty" require:"true"`
	TemplateVersion     *int    `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty" require:"true"`
	TemplateOverview    *string `json:"TemplateOverview,omitempty" xml:"TemplateOverview,omitempty" require:"true"`
	TemplateCreator     *string `json:"TemplateCreator,omitempty" xml:"TemplateCreator,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetRequestId(v string) *GetTemplateResponse {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponse) SetRegionId(v string) *GetTemplateResponse {
	s.RegionId = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateId(v string) *GetTemplateResponse {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateName(v string) *GetTemplateResponse {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateDescription(v string) *GetTemplateResponse {
	s.TemplateDescription = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateTag(v string) *GetTemplateResponse {
	s.TemplateTag = &v
	return s
}

func (s *GetTemplateResponse) SetDefinition(v string) *GetTemplateResponse {
	s.Definition = &v
	return s
}

func (s *GetTemplateResponse) SetCreateTime(v string) *GetTemplateResponse {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponse) SetUpdateTime(v string) *GetTemplateResponse {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateConnector(v string) *GetTemplateResponse {
	s.TemplateConnector = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateSummary(v string) *GetTemplateResponse {
	s.TemplateSummary = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateSummaryEn(v string) *GetTemplateResponse {
	s.TemplateSummaryEn = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateLocale(v string) *GetTemplateResponse {
	s.TemplateLocale = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateVersion(v int) *GetTemplateResponse {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateOverview(v string) *GetTemplateResponse {
	s.TemplateOverview = &v
	return s
}

func (s *GetTemplateResponse) SetTemplateCreator(v string) *GetTemplateResponse {
	s.TemplateCreator = &v
	return s
}

type GroupInvokeFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	GroupKey    *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty" require:"true"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
	TotalCount  *int    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
}

func (s GroupInvokeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupInvokeFlowRequest) GoString() string {
	return s.String()
}

func (s *GroupInvokeFlowRequest) SetFlowId(v string) *GroupInvokeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetGroupKey(v string) *GroupInvokeFlowRequest {
	s.GroupKey = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetData(v string) *GroupInvokeFlowRequest {
	s.Data = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetClientToken(v string) *GroupInvokeFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetTotalCount(v int) *GroupInvokeFlowRequest {
	s.TotalCount = &v
	return s
}

type GroupInvokeFlowResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	GroupInvocationId *string `json:"GroupInvocationId,omitempty" xml:"GroupInvocationId,omitempty" require:"true"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	CurrentCount      *int    `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty" require:"true"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s GroupInvokeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupInvokeFlowResponse) GoString() string {
	return s.String()
}

func (s *GroupInvokeFlowResponse) SetRequestId(v string) *GroupInvokeFlowResponse {
	s.RequestId = &v
	return s
}

func (s *GroupInvokeFlowResponse) SetGroupInvocationId(v string) *GroupInvokeFlowResponse {
	s.GroupInvocationId = &v
	return s
}

func (s *GroupInvokeFlowResponse) SetSuccess(v bool) *GroupInvokeFlowResponse {
	s.Success = &v
	return s
}

func (s *GroupInvokeFlowResponse) SetCurrentCount(v int) *GroupInvokeFlowResponse {
	s.CurrentCount = &v
	return s
}

func (s *GroupInvokeFlowResponse) SetStatus(v string) *GroupInvokeFlowResponse {
	s.Status = &v
	return s
}

type ListTagResourcesRequest struct {
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
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

func (s *ListTagResourcesRequest) SetMaxResults(v int) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
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
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TotalCount   *int                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TagResources []*ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Repeated"`
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

func (s *ListTagResourcesResponse) SetTotalCount(v int) *ListTagResourcesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v []*ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetTagKey(v string) *ListTagResourcesResponseTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetTagValue(v string) *ListTagResourcesResponseTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetResourceId(v string) *ListTagResourcesResponseTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseTagResources) SetResourceType(v string) *ListTagResourcesResponseTagResources {
	s.ResourceType = &v
	return s
}

type TagResourcesRequest struct {
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetSuccess(v bool) *TagResourcesResponse {
	s.Success = &v
	return s
}

type UntagResourcesRequest struct {
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetSuccess(v bool) *UntagResourcesResponse {
	s.Success = &v
	return s
}

type GetVersionRequest struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	VersionId *int    `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
}

func (s GetVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVersionRequest) GoString() string {
	return s.String()
}

func (s *GetVersionRequest) SetFlowId(v string) *GetVersionRequest {
	s.FlowId = &v
	return s
}

func (s *GetVersionRequest) SetVersionId(v int) *GetVersionRequest {
	s.VersionId = &v
	return s
}

type GetVersionResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowId             *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	VersionName        *string `json:"VersionName,omitempty" xml:"VersionName,omitempty" require:"true"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty" require:"true"`
	Definition         *string `json:"Definition,omitempty" xml:"Definition,omitempty" require:"true"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	VersionId          *int    `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	VersionStatus      *string `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty" require:"true"`
}

func (s GetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetVersionResponse) SetRequestId(v string) *GetVersionResponse {
	s.RequestId = &v
	return s
}

func (s *GetVersionResponse) SetFlowId(v string) *GetVersionResponse {
	s.FlowId = &v
	return s
}

func (s *GetVersionResponse) SetRegionId(v string) *GetVersionResponse {
	s.RegionId = &v
	return s
}

func (s *GetVersionResponse) SetVersionName(v string) *GetVersionResponse {
	s.VersionName = &v
	return s
}

func (s *GetVersionResponse) SetVersionDescription(v string) *GetVersionResponse {
	s.VersionDescription = &v
	return s
}

func (s *GetVersionResponse) SetDefinition(v string) *GetVersionResponse {
	s.Definition = &v
	return s
}

func (s *GetVersionResponse) SetCreateTime(v string) *GetVersionResponse {
	s.CreateTime = &v
	return s
}

func (s *GetVersionResponse) SetUpdateTime(v string) *GetVersionResponse {
	s.UpdateTime = &v
	return s
}

func (s *GetVersionResponse) SetVersionId(v int) *GetVersionResponse {
	s.VersionId = &v
	return s
}

func (s *GetVersionResponse) SetVersionStatus(v string) *GetVersionResponse {
	s.VersionStatus = &v
	return s
}

type ListVersionsRequest struct {
	FlowId     *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) SetFlowId(v string) *ListVersionsRequest {
	s.FlowId = &v
	return s
}

func (s *ListVersionsRequest) SetPageNumber(v int) *ListVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVersionsRequest) SetPageSize(v int) *ListVersionsRequest {
	s.PageSize = &v
	return s
}

type ListVersionsResponse struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Versions   []*ListVersionsResponseVersions `json:"Versions,omitempty" xml:"Versions,omitempty" require:"true" type:"Repeated"`
}

func (s ListVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionsResponse) SetRequestId(v string) *ListVersionsResponse {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponse) SetTotalCount(v int) *ListVersionsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListVersionsResponse) SetVersions(v []*ListVersionsResponseVersions) *ListVersionsResponse {
	s.Versions = v
	return s
}

type ListVersionsResponseVersions struct {
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	FlowId        *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	VersionName   *int    `json:"VersionName,omitempty" xml:"VersionName,omitempty" require:"true"`
	VersionStatus *int    `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty" require:"true"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
}

func (s ListVersionsResponseVersions) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseVersions) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseVersions) SetVersionId(v string) *ListVersionsResponseVersions {
	s.VersionId = &v
	return s
}

func (s *ListVersionsResponseVersions) SetFlowId(v string) *ListVersionsResponseVersions {
	s.FlowId = &v
	return s
}

func (s *ListVersionsResponseVersions) SetVersionName(v int) *ListVersionsResponseVersions {
	s.VersionName = &v
	return s
}

func (s *ListVersionsResponseVersions) SetVersionStatus(v int) *ListVersionsResponseVersions {
	s.VersionStatus = &v
	return s
}

func (s *ListVersionsResponseVersions) SetCreateTime(v string) *ListVersionsResponseVersions {
	s.CreateTime = &v
	return s
}

func (s *ListVersionsResponseVersions) SetUpdateTime(v string) *ListVersionsResponseVersions {
	s.UpdateTime = &v
	return s
}

type UpdateFlowRequest struct {
	FlowId          *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	Definition      *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
}

func (s UpdateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) SetFlowId(v string) *UpdateFlowRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowName(v string) *UpdateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowDescription(v string) *UpdateFlowRequest {
	s.FlowDescription = &v
	return s
}

func (s *UpdateFlowRequest) SetDefinition(v string) *UpdateFlowRequest {
	s.Definition = &v
	return s
}

type UpdateFlowResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	CurrentVersionId *int    `json:"CurrentVersionId,omitempty" xml:"CurrentVersionId,omitempty" require:"true"`
}

func (s UpdateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponse) SetRequestId(v string) *UpdateFlowResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponse) SetSuccess(v bool) *UpdateFlowResponse {
	s.Success = &v
	return s
}

func (s *UpdateFlowResponse) SetCurrentVersionId(v int) *UpdateFlowResponse {
	s.CurrentVersionId = &v
	return s
}

type CloneFlowRequest struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CloneFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowRequest) SetFlowId(v string) *CloneFlowRequest {
	s.FlowId = &v
	return s
}

func (s *CloneFlowRequest) SetVersionId(v string) *CloneFlowRequest {
	s.VersionId = &v
	return s
}

type CloneFlowResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s CloneFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowResponse) SetRequestId(v string) *CloneFlowResponse {
	s.RequestId = &v
	return s
}

func (s *CloneFlowResponse) SetFlowId(v string) *CloneFlowResponse {
	s.FlowId = &v
	return s
}

type GetFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s GetFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) SetFlowId(v string) *GetFlowRequest {
	s.FlowId = &v
	return s
}

type GetFlowResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowId           *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	FlowName         *string `json:"FlowName,omitempty" xml:"FlowName,omitempty" require:"true"`
	FlowDescription  *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty" require:"true"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	CurrentVersionId *int    `json:"CurrentVersionId,omitempty" xml:"CurrentVersionId,omitempty" require:"true"`
	FlowStatus       *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty" require:"true"`
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty" require:"true"`
	TemplateId       *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
	FlowSource       *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty" require:"true"`
	FlowEditMode     *string `json:"FlowEditMode,omitempty" xml:"FlowEditMode,omitempty" require:"true"`
}

func (s GetFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponse) GoString() string {
	return s.String()
}

func (s *GetFlowResponse) SetRequestId(v string) *GetFlowResponse {
	s.RequestId = &v
	return s
}

func (s *GetFlowResponse) SetFlowId(v string) *GetFlowResponse {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponse) SetRegionId(v string) *GetFlowResponse {
	s.RegionId = &v
	return s
}

func (s *GetFlowResponse) SetFlowName(v string) *GetFlowResponse {
	s.FlowName = &v
	return s
}

func (s *GetFlowResponse) SetFlowDescription(v string) *GetFlowResponse {
	s.FlowDescription = &v
	return s
}

func (s *GetFlowResponse) SetCreateTime(v string) *GetFlowResponse {
	s.CreateTime = &v
	return s
}

func (s *GetFlowResponse) SetUpdateTime(v string) *GetFlowResponse {
	s.UpdateTime = &v
	return s
}

func (s *GetFlowResponse) SetCurrentVersionId(v int) *GetFlowResponse {
	s.CurrentVersionId = &v
	return s
}

func (s *GetFlowResponse) SetFlowStatus(v string) *GetFlowResponse {
	s.FlowStatus = &v
	return s
}

func (s *GetFlowResponse) SetDefinition(v string) *GetFlowResponse {
	s.Definition = &v
	return s
}

func (s *GetFlowResponse) SetTemplateId(v string) *GetFlowResponse {
	s.TemplateId = &v
	return s
}

func (s *GetFlowResponse) SetFlowSource(v string) *GetFlowResponse {
	s.FlowSource = &v
	return s
}

func (s *GetFlowResponse) SetFlowEditMode(v string) *GetFlowResponse {
	s.FlowEditMode = &v
	return s
}

type ListFlowsRequest struct {
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FlowName   *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetPageSize(v int) *ListFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowsRequest) SetPageNumber(v int) *ListFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsRequest) SetFlowName(v string) *ListFlowsRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowsRequest) SetFilter(v string) *ListFlowsRequest {
	s.Filter = &v
	return s
}

type ListFlowsResponse struct {
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Flows      []*ListFlowsResponseFlows `json:"Flows,omitempty" xml:"Flows,omitempty" require:"true" type:"Repeated"`
}

func (s ListFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowsResponse) SetRequestId(v string) *ListFlowsResponse {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponse) SetTotalCount(v int) *ListFlowsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListFlowsResponse) SetFlows(v []*ListFlowsResponseFlows) *ListFlowsResponse {
	s.Flows = v
	return s
}

type ListFlowsResponseFlows struct {
	FlowId          *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty" require:"true"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty" require:"true"`
	VersionId       *int    `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	FlowStatus      *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty" require:"true"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" require:"true"`
	FlowSource      *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty" require:"true"`
	FlowEditMode    *string `json:"FlowEditMode,omitempty" xml:"FlowEditMode,omitempty" require:"true"`
}

func (s ListFlowsResponseFlows) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseFlows) SetFlowId(v string) *ListFlowsResponseFlows {
	s.FlowId = &v
	return s
}

func (s *ListFlowsResponseFlows) SetRegionId(v string) *ListFlowsResponseFlows {
	s.RegionId = &v
	return s
}

func (s *ListFlowsResponseFlows) SetFlowName(v string) *ListFlowsResponseFlows {
	s.FlowName = &v
	return s
}

func (s *ListFlowsResponseFlows) SetFlowDescription(v string) *ListFlowsResponseFlows {
	s.FlowDescription = &v
	return s
}

func (s *ListFlowsResponseFlows) SetVersionId(v int) *ListFlowsResponseFlows {
	s.VersionId = &v
	return s
}

func (s *ListFlowsResponseFlows) SetCreateTime(v string) *ListFlowsResponseFlows {
	s.CreateTime = &v
	return s
}

func (s *ListFlowsResponseFlows) SetUpdateTime(v string) *ListFlowsResponseFlows {
	s.UpdateTime = &v
	return s
}

func (s *ListFlowsResponseFlows) SetFlowStatus(v string) *ListFlowsResponseFlows {
	s.FlowStatus = &v
	return s
}

func (s *ListFlowsResponseFlows) SetTemplateId(v string) *ListFlowsResponseFlows {
	s.TemplateId = &v
	return s
}

func (s *ListFlowsResponseFlows) SetFlowSource(v string) *ListFlowsResponseFlows {
	s.FlowSource = &v
	return s
}

func (s *ListFlowsResponseFlows) SetFlowEditMode(v string) *ListFlowsResponseFlows {
	s.FlowEditMode = &v
	return s
}

type DeleteFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetFlowId(v string) *DeleteFlowRequest {
	s.FlowId = &v
	return s
}

type DeleteFlowResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s DeleteFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) SetRequestId(v string) *DeleteFlowResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowResponse) SetSuccess(v bool) *DeleteFlowResponse {
	s.Success = &v
	return s
}

type DisableFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s DisableFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRequest) GoString() string {
	return s.String()
}

func (s *DisableFlowRequest) SetFlowId(v string) *DisableFlowRequest {
	s.FlowId = &v
	return s
}

type DisableFlowResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty" require:"true"`
}

func (s DisableFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowResponse) GoString() string {
	return s.String()
}

func (s *DisableFlowResponse) SetRequestId(v string) *DisableFlowResponse {
	s.RequestId = &v
	return s
}

func (s *DisableFlowResponse) SetSuccess(v bool) *DisableFlowResponse {
	s.Success = &v
	return s
}

func (s *DisableFlowResponse) SetFlowStatus(v string) *DisableFlowResponse {
	s.FlowStatus = &v
	return s
}

type EnableFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s EnableFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRequest) GoString() string {
	return s.String()
}

func (s *EnableFlowRequest) SetFlowId(v string) *EnableFlowRequest {
	s.FlowId = &v
	return s
}

type EnableFlowResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty" require:"true"`
}

func (s EnableFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowResponse) GoString() string {
	return s.String()
}

func (s *EnableFlowResponse) SetRequestId(v string) *EnableFlowResponse {
	s.RequestId = &v
	return s
}

func (s *EnableFlowResponse) SetSuccess(v bool) *EnableFlowResponse {
	s.Success = &v
	return s
}

func (s *EnableFlowResponse) SetFlowStatus(v string) *EnableFlowResponse {
	s.FlowStatus = &v
	return s
}

type CreateFlowRequest struct {
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty" require:"true"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	Definition      *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	FlowSource      *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetFlowName(v string) *CreateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowRequest) SetFlowDescription(v string) *CreateFlowRequest {
	s.FlowDescription = &v
	return s
}

func (s *CreateFlowRequest) SetDefinition(v string) *CreateFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateFlowRequest) SetTemplateId(v string) *CreateFlowRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateFlowRequest) SetFlowSource(v string) *CreateFlowRequest {
	s.FlowSource = &v
	return s
}

type CreateFlowResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
}

func (s CreateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) SetRequestId(v string) *CreateFlowResponse {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponse) SetFlowId(v string) *CreateFlowResponse {
	s.FlowId = &v
	return s
}

type InvokeFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty" require:"true"`
	Parameters  *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s InvokeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowRequest) GoString() string {
	return s.String()
}

func (s *InvokeFlowRequest) SetFlowId(v string) *InvokeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *InvokeFlowRequest) SetParameters(v string) *InvokeFlowRequest {
	s.Parameters = &v
	return s
}

func (s *InvokeFlowRequest) SetData(v string) *InvokeFlowRequest {
	s.Data = &v
	return s
}

func (s *InvokeFlowRequest) SetClientToken(v string) *InvokeFlowRequest {
	s.ClientToken = &v
	return s
}

type InvokeFlowResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty" require:"true"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s InvokeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowResponse) GoString() string {
	return s.String()
}

func (s *InvokeFlowResponse) SetRequestId(v string) *InvokeFlowResponse {
	s.RequestId = &v
	return s
}

func (s *InvokeFlowResponse) SetInvocationId(v string) *InvokeFlowResponse {
	s.InvocationId = &v
	return s
}

func (s *InvokeFlowResponse) SetSuccess(v bool) *InvokeFlowResponse {
	s.Success = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("composer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTemplates"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRequest(tea.String("GetTemplate"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupInvokeFlowWithOptions(request *GroupInvokeFlowRequest, runtime *util.RuntimeOptions) (_result *GroupInvokeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GroupInvokeFlowResponse{}
	_body, _err := client.DoRequest(tea.String("GroupInvokeFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupInvokeFlow(request *GroupInvokeFlowRequest) (_result *GroupInvokeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GroupInvokeFlowResponse{}
	_body, _err := client.GroupInvokeFlowWithOptions(request, runtime)
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
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetVersionWithOptions(request *GetVersionRequest, runtime *util.RuntimeOptions) (_result *GetVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetVersionResponse{}
	_body, _err := client.DoRequest(tea.String("GetVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVersion(request *GetVersionRequest) (_result *GetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVersionResponse{}
	_body, _err := client.GetVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVersionsWithOptions(request *ListVersionsRequest, runtime *util.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListVersionsResponse{}
	_body, _err := client.DoRequest(tea.String("ListVersions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVersions(request *ListVersionsRequest) (_result *ListVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVersionsResponse{}
	_body, _err := client.ListVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlowWithOptions(request *UpdateFlowRequest, runtime *util.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlow(request *UpdateFlowRequest) (_result *UpdateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlowResponse{}
	_body, _err := client.UpdateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneFlowWithOptions(request *CloneFlowRequest, runtime *util.RuntimeOptions) (_result *CloneFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CloneFlowResponse{}
	_body, _err := client.DoRequest(tea.String("CloneFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlow(request *CloneFlowRequest) (_result *CloneFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowResponse{}
	_body, _err := client.CloneFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowWithOptions(request *GetFlowRequest, runtime *util.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetFlowResponse{}
	_body, _err := client.DoRequest(tea.String("GetFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlow(request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *util.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.DoRequest(tea.String("ListFlows"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableFlowWithOptions(request *DisableFlowRequest, runtime *util.RuntimeOptions) (_result *DisableFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableFlowResponse{}
	_body, _err := client.DoRequest(tea.String("DisableFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableFlow(request *DisableFlowRequest) (_result *DisableFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableFlowResponse{}
	_body, _err := client.DisableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableFlowWithOptions(request *EnableFlowRequest, runtime *util.RuntimeOptions) (_result *EnableFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableFlowResponse{}
	_body, _err := client.DoRequest(tea.String("EnableFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableFlow(request *EnableFlowRequest) (_result *EnableFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFlowResponse{}
	_body, _err := client.EnableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.DoRequest(tea.String("CreateFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeFlowWithOptions(request *InvokeFlowRequest, runtime *util.RuntimeOptions) (_result *InvokeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InvokeFlowResponse{}
	_body, _err := client.DoRequest(tea.String("InvokeFlow"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-12-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeFlow(request *InvokeFlowRequest) (_result *InvokeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeFlowResponse{}
	_body, _err := client.InvokeFlowWithOptions(request, runtime)
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
