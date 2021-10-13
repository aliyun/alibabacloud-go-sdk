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

type CreateResourceRequest struct {
	Body    *string `json:"body,omitempty" xml:"body,omitempty"`
	IsAsync *bool   `json:"isAsync,omitempty" xml:"isAsync,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetBody(v string) *CreateResourceRequest {
	s.Body = &v
	return s
}

func (s *CreateResourceRequest) SetIsAsync(v bool) *CreateResourceRequest {
	s.IsAsync = &v
	return s
}

type CreateResourceResponseBody struct {
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源id
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetTaskId(v string) *CreateResourceResponseBody {
	s.TaskId = &v
	return s
}

type CreateResourceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	IsAsync  *bool   `json:"isAsync,omitempty" xml:"isAsync,omitempty"`
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetIsAsync(v bool) *DeleteResourceRequest {
	s.IsAsync = &v
	return s
}

func (s *DeleteResourceRequest) SetRegionId(v string) *DeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DeleteResourceResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetTaskId(v string) *DeleteResourceResponseBody {
	s.TaskId = &v
	return s
}

type DeleteResourceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetRegionId(v string) *GetResourceRequest {
	s.RegionId = &v
	return s
}

type GetResourceResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Resource  *GetResourceResponseBodyResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody {
	s.Resource = v
	return s
}

type GetResourceResponseBodyResource struct {
	ProductCode        *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	RegionId           *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceAttributes *string `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
	ResourceId         *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceTypeCode   *string `json:"resourceTypeCode,omitempty" xml:"resourceTypeCode,omitempty"`
}

func (s GetResourceResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResource) SetProductCode(v string) *GetResourceResponseBodyResource {
	s.ProductCode = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetRegionId(v string) *GetResourceResponseBodyResource {
	s.RegionId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetResourceAttributes(v string) *GetResourceResponseBodyResource {
	s.ResourceAttributes = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetResourceId(v string) *GetResourceResponseBodyResource {
	s.ResourceId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetResourceTypeCode(v string) *GetResourceResponseBodyResource {
	s.ResourceTypeCode = &v
	return s
}

type GetResourceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetTaskResponseBody struct {
	// Id of the request
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponseBodyTask struct {
	FailedReason *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	ResourceId   *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) SetFailedReason(v string) *GetTaskResponseBodyTask {
	s.FailedReason = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourceId(v string) *GetTaskResponseBodyTask {
	s.ResourceId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

type GetTaskResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetFilter(v map[string]interface{}) *ListDataSourcesRequest {
	s.Filter = v
	return s
}

type ListDataSourcesShrinkRequest struct {
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) SetFilterShrink(v string) *ListDataSourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

type ListDataSourcesResponseBody struct {
	DataSources []*ListDataSourcesResponseBodyDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourcesResponseBodyDataSources struct {
	Id                   *string `json:"id,omitempty" xml:"id,omitempty"`
	DataSourceAttributes *string `json:"dataSourceAttributes,omitempty" xml:"dataSourceAttributes,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSources) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSources) SetId(v string) *ListDataSourcesResponseBodyDataSources {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) SetDataSourceAttributes(v string) *ListDataSourcesResponseBodyDataSources {
	s.DataSourceAttributes = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetMaxResults(v int64) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

type ListProductsResponseBody struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Products  []*ListProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetMaxResults(v int64) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int64) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductsResponseBodyProducts struct {
	ProductCode *string                                      `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName *ListProductsResponseBodyProductsProductName `json:"productName,omitempty" xml:"productName,omitempty" type:"Struct"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) SetProductCode(v string) *ListProductsResponseBodyProducts {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductName(v *ListProductsResponseBodyProductsProductName) *ListProductsResponseBodyProducts {
	s.ProductName = v
	return s
}

type ListProductsResponseBodyProductsProductName struct {
	ZhCN *string `json:"zh_CN,omitempty" xml:"zh_CN,omitempty"`
	EnUS *string `json:"en_US,omitempty" xml:"en_US,omitempty"`
}

func (s ListProductsResponseBodyProductsProductName) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProductsProductName) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductsProductName) SetZhCN(v string) *ListProductsResponseBodyProductsProductName {
	s.ZhCN = &v
	return s
}

func (s *ListProductsResponseBodyProductsProductName) SetEnUS(v string) *ListProductsResponseBodyProductsProductName {
	s.EnUS = &v
	return s
}

type ListProductsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListResourceTypesRequest struct {
	MaxResults        *int64    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken         *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceTypeCodes []*string `json:"resourceTypeCodes,omitempty" xml:"resourceTypeCodes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetMaxResults(v int64) *ListResourceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesRequest) SetNextToken(v string) *ListResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesRequest) SetResourceTypeCodes(v []*string) *ListResourceTypesRequest {
	s.ResourceTypeCodes = v
	return s
}

type ListResourceTypesShrinkRequest struct {
	MaxResults              *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken               *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceTypeCodesShrink *string `json:"resourceTypeCodes,omitempty" xml:"resourceTypeCodes,omitempty"`
}

func (s ListResourceTypesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesShrinkRequest) SetMaxResults(v int64) *ListResourceTypesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetNextToken(v string) *ListResourceTypesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetResourceTypeCodesShrink(v string) *ListResourceTypesShrinkRequest {
	s.ResourceTypeCodesShrink = &v
	return s
}

type ListResourceTypesResponseBody struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	RequestId     *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetMaxResults(v int64) *ListResourceTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetNextToken(v string) *ListResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) SetTotalCount(v int64) *ListResourceTypesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypes struct {
	ProductCode        *string                                                        `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ResourceTypeCode   *string                                                        `json:"resourceTypeCode,omitempty" xml:"resourceTypeCode,omitempty"`
	Info               *ListResourceTypesResponseBodyResourceTypesInfo                `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	IdentityDefinition *ListResourceTypesResponseBodyResourceTypesIdentityDefinition  `json:"identityDefinition,omitempty" xml:"identityDefinition,omitempty" type:"Struct"`
	StatusDefinition   []*ListResourceTypesResponseBodyResourceTypesStatusDefinition  `json:"statusDefinition,omitempty" xml:"statusDefinition,omitempty" type:"Repeated"`
	ResourceProperties *string                                                        `json:"resourceProperties,omitempty" xml:"resourceProperties,omitempty"`
	ResourceRelations  []*ListResourceTypesResponseBodyResourceTypesResourceRelations `json:"resourceRelations,omitempty" xml:"resourceRelations,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductCode(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductCode = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceTypeCode(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceTypeCode = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetInfo(v *ListResourceTypesResponseBodyResourceTypesInfo) *ListResourceTypesResponseBodyResourceTypes {
	s.Info = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetIdentityDefinition(v *ListResourceTypesResponseBodyResourceTypesIdentityDefinition) *ListResourceTypesResponseBodyResourceTypes {
	s.IdentityDefinition = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetStatusDefinition(v []*ListResourceTypesResponseBodyResourceTypesStatusDefinition) *ListResourceTypesResponseBodyResourceTypes {
	s.StatusDefinition = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceProperties(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceProperties = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceRelations(v []*ListResourceTypesResponseBodyResourceTypesResourceRelations) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceRelations = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesInfo struct {
	// 资源类型的中文名称，如实例
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 资源分类 枚举:normal(普通资源)/singleton(单例资源)/virtual(虚拟资源)/readonly(只读资源)
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// 交付级别 枚举:center(中心化部署级别)/region(地域部署级别)/zone(可用区部署级别)
	DeliveryScope *string `json:"deliveryScope,omitempty" xml:"deliveryScope,omitempty"`
	// 付费形式  枚举:paid(付费)/free(免费)
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 允许资源展示的站点  枚举:china(中国站)/intl(国际站)/japan(日本站)
	AvailableSites []*string `json:"availableSites,omitempty" xml:"availableSites,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetTitle(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Title = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetCategory(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Category = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDeliveryScope(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.DeliveryScope = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetChargeType(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.ChargeType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetAvailableSites(v []*string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.AvailableSites = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesIdentityDefinition struct {
	// uniqueKey的字段列表，有顺序
	UniqueKeyFields []*string `json:"uniqueKeyFields,omitempty" xml:"uniqueKeyFields,omitempty" type:"Repeated"`
	// 备选Id字段列表，有顺序
	SecondUniqueKeyFields []*string `json:"secondUniqueKeyFields,omitempty" xml:"secondUniqueKeyFields,omitempty" type:"Repeated"`
	// 资源ARN
	ArnPattern *string `json:"arnPattern,omitempty" xml:"arnPattern,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesIdentityDefinition) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesIdentityDefinition) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesIdentityDefinition) SetUniqueKeyFields(v []*string) *ListResourceTypesResponseBodyResourceTypesIdentityDefinition {
	s.UniqueKeyFields = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesIdentityDefinition) SetSecondUniqueKeyFields(v []*string) *ListResourceTypesResponseBodyResourceTypesIdentityDefinition {
	s.SecondUniqueKeyFields = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesIdentityDefinition) SetArnPattern(v string) *ListResourceTypesResponseBodyResourceTypesIdentityDefinition {
	s.ArnPattern = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypesStatusDefinition struct {
	// 状态code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 资源状态分类，必须对代表资源创建后的初始状态进行initial标识。枚举:initial(初始状态)
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesStatusDefinition) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesStatusDefinition) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesStatusDefinition) SetCode(v string) *ListResourceTypesResponseBodyResourceTypesStatusDefinition {
	s.Code = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesStatusDefinition) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypesStatusDefinition {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesStatusDefinition) SetType(v string) *ListResourceTypesResponseBodyResourceTypesStatusDefinition {
	s.Type = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypesResourceRelations struct {
	// 云产品B
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// 资源类型B
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 资源关系  枚举:relevance(关联关系)/dependency(依赖关系)/childParent(子父关系)
	Relation *string `json:"relation,omitempty" xml:"relation,omitempty"`
	// 资源关系描述 枚举：枚举:关联关系/依赖关系/子父关系
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesResourceRelations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesResourceRelations) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesResourceRelations) SetProduct(v string) *ListResourceTypesResponseBodyResourceTypesResourceRelations {
	s.Product = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesResourceRelations) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypesResourceRelations {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesResourceRelations) SetRelation(v string) *ListResourceTypesResponseBodyResourceTypesResourceRelations {
	s.Relation = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesResourceRelations) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypesResourceRelations {
	s.Description = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	Filter    map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	IsReload  *bool                  `json:"isReload,omitempty" xml:"isReload,omitempty"`
	PageNum   *int32                 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize  *int32                 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RegionIds []*string              `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetFilter(v map[string]interface{}) *ListResourcesRequest {
	s.Filter = v
	return s
}

func (s *ListResourcesRequest) SetIsReload(v bool) *ListResourcesRequest {
	s.IsReload = &v
	return s
}

func (s *ListResourcesRequest) SetPageNum(v int32) *ListResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetRegionIds(v []*string) *ListResourcesRequest {
	s.RegionIds = v
	return s
}

type ListResourcesShrinkRequest struct {
	FilterShrink    *string `json:"filter,omitempty" xml:"filter,omitempty"`
	IsReload        *bool   `json:"isReload,omitempty" xml:"isReload,omitempty"`
	PageNum         *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RegionIdsShrink *string `json:"regionIds,omitempty" xml:"regionIds,omitempty"`
}

func (s ListResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesShrinkRequest) SetFilterShrink(v string) *ListResourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListResourcesShrinkRequest) SetIsReload(v bool) *ListResourcesShrinkRequest {
	s.IsReload = &v
	return s
}

func (s *ListResourcesShrinkRequest) SetPageNum(v int32) *ListResourcesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListResourcesShrinkRequest) SetPageSize(v int32) *ListResourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesShrinkRequest) SetRegionIdsShrink(v string) *ListResourcesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

type ListResourcesResponseBody struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	RequestId  *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Resources  []*ListResourcesResponseBodyResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPageNum(v int32) *ListResourcesResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyResources struct {
	ProductCode        *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ResourceTypeCode   *string `json:"resourceTypeCode,omitempty" xml:"resourceTypeCode,omitempty"`
	RegionId           *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceId         *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceAttributes *string `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) SetProductCode(v string) *ListResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceTypeCode(v string) *ListResourcesResponseBodyResources {
	s.ResourceTypeCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetRegionId(v string) *ListResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceId(v string) *ListResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceAttributes(v string) *ListResourcesResponseBodyResources {
	s.ResourceAttributes = &v
	return s
}

type ListResourcesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ReloadResourcesRequest struct {
	RegionIds []*string `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
}

func (s ReloadResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReloadResourcesRequest) GoString() string {
	return s.String()
}

func (s *ReloadResourcesRequest) SetRegionIds(v []*string) *ReloadResourcesRequest {
	s.RegionIds = v
	return s
}

type ReloadResourcesShrinkRequest struct {
	RegionIdsShrink *string `json:"regionIds,omitempty" xml:"regionIds,omitempty"`
}

func (s ReloadResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReloadResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReloadResourcesShrinkRequest) SetRegionIdsShrink(v string) *ReloadResourcesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

type ReloadResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ReloadResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReloadResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ReloadResourcesResponseBody) SetRequestId(v string) *ReloadResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReloadResourcesResponseBody) SetTaskId(v string) *ReloadResourcesResponseBody {
	s.TaskId = &v
	return s
}

type ReloadResourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReloadResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReloadResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReloadResourcesResponse) GoString() string {
	return s.String()
}

func (s *ReloadResourcesResponse) SetHeaders(v map[string]*string) *ReloadResourcesResponse {
	s.Headers = v
	return s
}

func (s *ReloadResourcesResponse) SetBody(v *ReloadResourcesResponseBody) *ReloadResourcesResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	Body    *string `json:"body,omitempty" xml:"body,omitempty"`
	IsAsync *bool   `json:"isAsync,omitempty" xml:"isAsync,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetBody(v string) *UpdateResourceRequest {
	s.Body = &v
	return s
}

func (s *UpdateResourceRequest) SetIsAsync(v bool) *UpdateResourceRequest {
	s.IsAsync = &v
	return s
}

type UpdateResourceResponseBody struct {
	// 请求id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetTaskId(v string) *UpdateResourceResponseBody {
	s.TaskId = &v
	return s
}

type UpdateResourceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iacservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateResource(productCode *string, provider *string, resourceTypeCode *string, request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(productCode, provider, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceWithOptions(productCode *string, provider *string, resourceTypeCode *string, request *CreateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		query["isAsync"] = request.IsAsync
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResource(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(productCode, provider, resourceId, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceWithOptions(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *DeleteResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		query["isAsync"] = request.IsAsync
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources/{resourceId}"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResource(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(productCode, provider, resourceId, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceWithOptions(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *GetResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources/{resourceId}"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(taskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	taskId = openapiutil.GetEncodeParam(taskId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tasks/" + tea.StringValue(taskId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSources(attributeName *string, productCode *string, provider *string, resourceTypeCode *string, request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(attributeName, productCode, provider, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourcesWithOptions(attributeName *string, productCode *string, provider *string, resourceTypeCode *string, tmpReq *ListDataSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	attributeName = openapiutil.GetEncodeParam(attributeName)
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["filter"] = request.FilterShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/dataSources/{attributeName}"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProducts(provider *string, request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(provider, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsWithOptions(provider *string, request *ListProductsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	provider = openapiutil.GetEncodeParam(provider)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypes(productCode *string, provider *string, request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(productCode, provider, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypesWithOptions(productCode *string, provider *string, tmpReq *ListResourceTypesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	request := &ListResourceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypeCodes)) {
		request.ResourceTypeCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypeCodes, tea.String("resourceTypeCodes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeCodesShrink)) {
		query["resourceTypeCodes"] = request.ResourceTypeCodesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ListResources(productCode *string, provider *string, resourceTypeCode *string, request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(productCode, provider, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesWithOptions(productCode *string, provider *string, resourceTypeCode *string, tmpReq *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	request := &ListResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("filter"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("regionIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IsReload)) {
		query["isReload"] = request.IsReload
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["regionIds"] = request.RegionIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReloadResources(productCode *string, provider *string, resourceTypeCode *string, request *ReloadResourcesRequest) (_result *ReloadResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReloadResourcesResponse{}
	_body, _err := client.ReloadResourcesWithOptions(productCode, provider, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReloadResourcesWithOptions(productCode *string, provider *string, resourceTypeCode *string, tmpReq *ReloadResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReloadResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	request := &ReloadResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("regionIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["regionIds"] = request.RegionIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReloadResources"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources/operation/reload"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReloadResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResource(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(productCode, provider, resourceId, resourceTypeCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceWithOptions(productCode *string, provider *string, resourceId *string, resourceTypeCode *string, request *UpdateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	productCode = openapiutil.GetEncodeParam(productCode)
	provider = openapiutil.GetEncodeParam(provider)
	resourceId = openapiutil.GetEncodeParam(resourceId)
	resourceTypeCode = openapiutil.GetEncodeParam(resourceTypeCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		query["isAsync"] = request.IsAsync
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2021-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(provider) + "/products/{productCode}/resourceTypes/{resourceTypeCode}/resources/{resourceId}"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
