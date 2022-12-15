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

type LingJieApiInvokeCount struct {
	ApiName        *string                         `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Metrics        []*LingJieApiInvokeCountMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	ProjectId      *string                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Region         *string                         `json:"Region,omitempty" xml:"Region,omitempty"`
	StatisticsUnit *string                         `json:"StatisticsUnit,omitempty" xml:"StatisticsUnit,omitempty"`
	TimeUnit       *string                         `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s LingJieApiInvokeCount) String() string {
	return tea.Prettify(s)
}

func (s LingJieApiInvokeCount) GoString() string {
	return s.String()
}

func (s *LingJieApiInvokeCount) SetApiName(v string) *LingJieApiInvokeCount {
	s.ApiName = &v
	return s
}

func (s *LingJieApiInvokeCount) SetMetrics(v []*LingJieApiInvokeCountMetrics) *LingJieApiInvokeCount {
	s.Metrics = v
	return s
}

func (s *LingJieApiInvokeCount) SetProjectId(v string) *LingJieApiInvokeCount {
	s.ProjectId = &v
	return s
}

func (s *LingJieApiInvokeCount) SetRegion(v string) *LingJieApiInvokeCount {
	s.Region = &v
	return s
}

func (s *LingJieApiInvokeCount) SetStatisticsUnit(v string) *LingJieApiInvokeCount {
	s.StatisticsUnit = &v
	return s
}

func (s *LingJieApiInvokeCount) SetTimeUnit(v string) *LingJieApiInvokeCount {
	s.TimeUnit = &v
	return s
}

type LingJieApiInvokeCountMetrics struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Day   *string `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour  *string `json:"Hour,omitempty" xml:"Hour,omitempty"`
}

func (s LingJieApiInvokeCountMetrics) String() string {
	return tea.Prettify(s)
}

func (s LingJieApiInvokeCountMetrics) GoString() string {
	return s.String()
}

func (s *LingJieApiInvokeCountMetrics) SetCount(v int32) *LingJieApiInvokeCountMetrics {
	s.Count = &v
	return s
}

func (s *LingJieApiInvokeCountMetrics) SetDay(v string) *LingJieApiInvokeCountMetrics {
	s.Day = &v
	return s
}

func (s *LingJieApiInvokeCountMetrics) SetHour(v string) *LingJieApiInvokeCountMetrics {
	s.Hour = &v
	return s
}

type LingJieApiInvokeQps struct {
	ApiName        *string                       `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Metrics        []*LingJieApiInvokeQpsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	ProjectId      *string                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Region         *string                       `json:"Region,omitempty" xml:"Region,omitempty"`
	StatisticsUnit *string                       `json:"StatisticsUnit,omitempty" xml:"StatisticsUnit,omitempty"`
	TimeUnit       *string                       `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s LingJieApiInvokeQps) String() string {
	return tea.Prettify(s)
}

func (s LingJieApiInvokeQps) GoString() string {
	return s.String()
}

func (s *LingJieApiInvokeQps) SetApiName(v string) *LingJieApiInvokeQps {
	s.ApiName = &v
	return s
}

func (s *LingJieApiInvokeQps) SetMetrics(v []*LingJieApiInvokeQpsMetrics) *LingJieApiInvokeQps {
	s.Metrics = v
	return s
}

func (s *LingJieApiInvokeQps) SetProjectId(v string) *LingJieApiInvokeQps {
	s.ProjectId = &v
	return s
}

func (s *LingJieApiInvokeQps) SetRegion(v string) *LingJieApiInvokeQps {
	s.Region = &v
	return s
}

func (s *LingJieApiInvokeQps) SetStatisticsUnit(v string) *LingJieApiInvokeQps {
	s.StatisticsUnit = &v
	return s
}

func (s *LingJieApiInvokeQps) SetTimeUnit(v string) *LingJieApiInvokeQps {
	s.TimeUnit = &v
	return s
}

type LingJieApiInvokeQpsMetrics struct {
	Day    *string `json:"Day,omitempty" xml:"Day,omitempty"`
	Hour   *string `json:"Hour,omitempty" xml:"Hour,omitempty"`
	MaxQps *int32  `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
}

func (s LingJieApiInvokeQpsMetrics) String() string {
	return tea.Prettify(s)
}

func (s LingJieApiInvokeQpsMetrics) GoString() string {
	return s.String()
}

func (s *LingJieApiInvokeQpsMetrics) SetDay(v string) *LingJieApiInvokeQpsMetrics {
	s.Day = &v
	return s
}

func (s *LingJieApiInvokeQpsMetrics) SetHour(v string) *LingJieApiInvokeQpsMetrics {
	s.Hour = &v
	return s
}

func (s *LingJieApiInvokeQpsMetrics) SetMaxQps(v int32) *LingJieApiInvokeQpsMetrics {
	s.MaxQps = &v
	return s
}

type LingJieOpenStatus struct {
	CommoditiesOpenStatus []*LingJieOpenStatusCommoditiesOpenStatus `json:"CommoditiesOpenStatus,omitempty" xml:"CommoditiesOpenStatus,omitempty" type:"Repeated"`
	OpenTime              *string                                   `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	ProductConsoleUrl     *string                                   `json:"ProductConsoleUrl,omitempty" xml:"ProductConsoleUrl,omitempty"`
	ProductEnabled        *bool                                     `json:"ProductEnabled,omitempty" xml:"ProductEnabled,omitempty"`
	ProductMonitorUrl     *string                                   `json:"ProductMonitorUrl,omitempty" xml:"ProductMonitorUrl,omitempty"`
	ProductWikiUrl        *string                                   `json:"ProductWikiUrl,omitempty" xml:"ProductWikiUrl,omitempty"`
}

func (s LingJieOpenStatus) String() string {
	return tea.Prettify(s)
}

func (s LingJieOpenStatus) GoString() string {
	return s.String()
}

func (s *LingJieOpenStatus) SetCommoditiesOpenStatus(v []*LingJieOpenStatusCommoditiesOpenStatus) *LingJieOpenStatus {
	s.CommoditiesOpenStatus = v
	return s
}

func (s *LingJieOpenStatus) SetOpenTime(v string) *LingJieOpenStatus {
	s.OpenTime = &v
	return s
}

func (s *LingJieOpenStatus) SetProductConsoleUrl(v string) *LingJieOpenStatus {
	s.ProductConsoleUrl = &v
	return s
}

func (s *LingJieOpenStatus) SetProductEnabled(v bool) *LingJieOpenStatus {
	s.ProductEnabled = &v
	return s
}

func (s *LingJieOpenStatus) SetProductMonitorUrl(v string) *LingJieOpenStatus {
	s.ProductMonitorUrl = &v
	return s
}

func (s *LingJieOpenStatus) SetProductWikiUrl(v string) *LingJieOpenStatus {
	s.ProductWikiUrl = &v
	return s
}

type LingJieOpenStatusCommoditiesOpenStatus struct {
	CnName     *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	Commodity  *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	Describe   *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	DetailPage *string `json:"DetailPage,omitempty" xml:"DetailPage,omitempty"`
	Open       *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	OpenTime   *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s LingJieOpenStatusCommoditiesOpenStatus) String() string {
	return tea.Prettify(s)
}

func (s LingJieOpenStatusCommoditiesOpenStatus) GoString() string {
	return s.String()
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetCnName(v string) *LingJieOpenStatusCommoditiesOpenStatus {
	s.CnName = &v
	return s
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetCommodity(v string) *LingJieOpenStatusCommoditiesOpenStatus {
	s.Commodity = &v
	return s
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetDescribe(v string) *LingJieOpenStatusCommoditiesOpenStatus {
	s.Describe = &v
	return s
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetDetailPage(v string) *LingJieOpenStatusCommoditiesOpenStatus {
	s.DetailPage = &v
	return s
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetOpen(v bool) *LingJieOpenStatusCommoditiesOpenStatus {
	s.Open = &v
	return s
}

func (s *LingJieOpenStatusCommoditiesOpenStatus) SetOpenTime(v string) *LingJieOpenStatusCommoditiesOpenStatus {
	s.OpenTime = &v
	return s
}

type LingJieOpenStatusDetail struct {
	ApiList                      []*LingJieOpenStatusDetailApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	CommodityConcurrentLimitTips *string                           `json:"CommodityConcurrentLimitTips,omitempty" xml:"CommodityConcurrentLimitTips,omitempty"`
	CommodityHasOpen             *bool                             `json:"CommodityHasOpen,omitempty" xml:"CommodityHasOpen,omitempty"`
	CommodityOpenUrl             *string                           `json:"CommodityOpenUrl,omitempty" xml:"CommodityOpenUrl,omitempty"`
	CommodityUsageTips           *string                           `json:"CommodityUsageTips,omitempty" xml:"CommodityUsageTips,omitempty"`
	UserStopMode                 *bool                             `json:"UserStopMode,omitempty" xml:"UserStopMode,omitempty"`
}

func (s LingJieOpenStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s LingJieOpenStatusDetail) GoString() string {
	return s.String()
}

func (s *LingJieOpenStatusDetail) SetApiList(v []*LingJieOpenStatusDetailApiList) *LingJieOpenStatusDetail {
	s.ApiList = v
	return s
}

func (s *LingJieOpenStatusDetail) SetCommodityConcurrentLimitTips(v string) *LingJieOpenStatusDetail {
	s.CommodityConcurrentLimitTips = &v
	return s
}

func (s *LingJieOpenStatusDetail) SetCommodityHasOpen(v bool) *LingJieOpenStatusDetail {
	s.CommodityHasOpen = &v
	return s
}

func (s *LingJieOpenStatusDetail) SetCommodityOpenUrl(v string) *LingJieOpenStatusDetail {
	s.CommodityOpenUrl = &v
	return s
}

func (s *LingJieOpenStatusDetail) SetCommodityUsageTips(v string) *LingJieOpenStatusDetail {
	s.CommodityUsageTips = &v
	return s
}

func (s *LingJieOpenStatusDetail) SetUserStopMode(v bool) *LingJieOpenStatusDetail {
	s.UserStopMode = &v
	return s
}

type LingJieOpenStatusDetailApiList struct {
	CnName          *string                                         `json:"CnName,omitempty" xml:"CnName,omitempty"`
	ConcurrentLimit *string                                         `json:"ConcurrentLimit,omitempty" xml:"ConcurrentLimit,omitempty"`
	EnName          *string                                         `json:"EnName,omitempty" xml:"EnName,omitempty"`
	MoreOperations  []*LingJieOpenStatusDetailApiListMoreOperations `json:"MoreOperations,omitempty" xml:"MoreOperations,omitempty" type:"Repeated"`
	StatisticsUnit  *string                                         `json:"StatisticsUnit,omitempty" xml:"StatisticsUnit,omitempty"`
	Status          *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Usage           *string                                         `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s LingJieOpenStatusDetailApiList) String() string {
	return tea.Prettify(s)
}

func (s LingJieOpenStatusDetailApiList) GoString() string {
	return s.String()
}

func (s *LingJieOpenStatusDetailApiList) SetCnName(v string) *LingJieOpenStatusDetailApiList {
	s.CnName = &v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetConcurrentLimit(v string) *LingJieOpenStatusDetailApiList {
	s.ConcurrentLimit = &v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetEnName(v string) *LingJieOpenStatusDetailApiList {
	s.EnName = &v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetMoreOperations(v []*LingJieOpenStatusDetailApiListMoreOperations) *LingJieOpenStatusDetailApiList {
	s.MoreOperations = v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetStatisticsUnit(v string) *LingJieOpenStatusDetailApiList {
	s.StatisticsUnit = &v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetStatus(v string) *LingJieOpenStatusDetailApiList {
	s.Status = &v
	return s
}

func (s *LingJieOpenStatusDetailApiList) SetUsage(v string) *LingJieOpenStatusDetailApiList {
	s.Usage = &v
	return s
}

type LingJieOpenStatusDetailApiListMoreOperations struct {
	ClickUrl  *string `json:"ClickUrl,omitempty" xml:"ClickUrl,omitempty"`
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s LingJieOpenStatusDetailApiListMoreOperations) String() string {
	return tea.Prettify(s)
}

func (s LingJieOpenStatusDetailApiListMoreOperations) GoString() string {
	return s.String()
}

func (s *LingJieOpenStatusDetailApiListMoreOperations) SetClickUrl(v string) *LingJieOpenStatusDetailApiListMoreOperations {
	s.ClickUrl = &v
	return s
}

func (s *LingJieOpenStatusDetailApiListMoreOperations) SetOperation(v string) *LingJieOpenStatusDetailApiListMoreOperations {
	s.Operation = &v
	return s
}

type News struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Last        *bool   `json:"Last,omitempty" xml:"Last,omitempty"`
	Pic         *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	Sort        *int32  `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s News) String() string {
	return tea.Prettify(s)
}

func (s News) GoString() string {
	return s.String()
}

func (s *News) SetContent(v string) *News {
	s.Content = &v
	return s
}

func (s *News) SetGmtCreate(v string) *News {
	s.GmtCreate = &v
	return s
}

func (s *News) SetGmtModified(v string) *News {
	s.GmtModified = &v
	return s
}

func (s *News) SetId(v int64) *News {
	s.Id = &v
	return s
}

func (s *News) SetLast(v bool) *News {
	s.Last = &v
	return s
}

func (s *News) SetPic(v string) *News {
	s.Pic = &v
	return s
}

func (s *News) SetSort(v int32) *News {
	s.Sort = &v
	return s
}

func (s *News) SetStatus(v int32) *News {
	s.Status = &v
	return s
}

func (s *News) SetTags(v string) *News {
	s.Tags = &v
	return s
}

func (s *News) SetTitle(v string) *News {
	s.Title = &v
	return s
}

func (s *News) SetType(v int32) *News {
	s.Type = &v
	return s
}

func (s *News) SetUrl(v string) *News {
	s.Url = &v
	return s
}

type UserQpsDetail struct {
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Day     *string `json:"Day,omitempty" xml:"Day,omitempty"`
	Qps     *int32  `json:"Qps,omitempty" xml:"Qps,omitempty"`
	QpsRule *string `json:"QpsRule,omitempty" xml:"QpsRule,omitempty"`
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UserQpsDetail) String() string {
	return tea.Prettify(s)
}

func (s UserQpsDetail) GoString() string {
	return s.String()
}

func (s *UserQpsDetail) SetApiName(v string) *UserQpsDetail {
	s.ApiName = &v
	return s
}

func (s *UserQpsDetail) SetDay(v string) *UserQpsDetail {
	s.Day = &v
	return s
}

func (s *UserQpsDetail) SetQps(v int32) *UserQpsDetail {
	s.Qps = &v
	return s
}

func (s *UserQpsDetail) SetQpsRule(v string) *UserQpsDetail {
	s.QpsRule = &v
	return s
}

func (s *UserQpsDetail) SetStatus(v int32) *UserQpsDetail {
	s.Status = &v
	return s
}

type CreateSDKInstanceRequest struct {
	BundleId  *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	Pk        *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Platform  *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Recipe    *string `json:"Recipe,omitempty" xml:"Recipe,omitempty"`
	ValidFrom *int64  `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	ValidTo   *int64  `json:"ValidTo,omitempty" xml:"ValidTo,omitempty"`
}

func (s CreateSDKInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSDKInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSDKInstanceRequest) SetBundleId(v string) *CreateSDKInstanceRequest {
	s.BundleId = &v
	return s
}

func (s *CreateSDKInstanceRequest) SetPk(v string) *CreateSDKInstanceRequest {
	s.Pk = &v
	return s
}

func (s *CreateSDKInstanceRequest) SetPlatform(v string) *CreateSDKInstanceRequest {
	s.Platform = &v
	return s
}

func (s *CreateSDKInstanceRequest) SetRecipe(v string) *CreateSDKInstanceRequest {
	s.Recipe = &v
	return s
}

func (s *CreateSDKInstanceRequest) SetValidFrom(v int64) *CreateSDKInstanceRequest {
	s.ValidFrom = &v
	return s
}

func (s *CreateSDKInstanceRequest) SetValidTo(v int64) *CreateSDKInstanceRequest {
	s.ValidTo = &v
	return s
}

type CreateSDKInstanceResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Ok           *bool   `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSDKInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSDKInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSDKInstanceResponseBody) SetCode(v string) *CreateSDKInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSDKInstanceResponseBody) SetData(v string) *CreateSDKInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSDKInstanceResponseBody) SetErrorMessage(v string) *CreateSDKInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSDKInstanceResponseBody) SetHttpCode(v int32) *CreateSDKInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateSDKInstanceResponseBody) SetOk(v bool) *CreateSDKInstanceResponseBody {
	s.Ok = &v
	return s
}

func (s *CreateSDKInstanceResponseBody) SetRequestId(v string) *CreateSDKInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateSDKInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSDKInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSDKInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSDKInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSDKInstanceResponse) SetHeaders(v map[string]*string) *CreateSDKInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSDKInstanceResponse) SetStatusCode(v int32) *CreateSDKInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSDKInstanceResponse) SetBody(v *CreateSDKInstanceResponseBody) *CreateSDKInstanceResponse {
	s.Body = v
	return s
}

type GetSDKInstanceDebugInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSDKInstanceDebugInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSDKInstanceDebugInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSDKInstanceDebugInfoRequest) SetInstanceId(v string) *GetSDKInstanceDebugInfoRequest {
	s.InstanceId = &v
	return s
}

type GetSDKInstanceDebugInfoResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetSDKInstanceDebugInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *int32                                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Ok           *bool                                    `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSDKInstanceDebugInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSDKInstanceDebugInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetCode(v string) *GetSDKInstanceDebugInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetData(v *GetSDKInstanceDebugInfoResponseBodyData) *GetSDKInstanceDebugInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetErrorMessage(v string) *GetSDKInstanceDebugInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetHttpCode(v int32) *GetSDKInstanceDebugInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetOk(v bool) *GetSDKInstanceDebugInfoResponseBody {
	s.Ok = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBody) SetRequestId(v string) *GetSDKInstanceDebugInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetSDKInstanceDebugInfoResponseBodyData struct {
	BundleId      *string                                          `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	CreatedAt     *string                                          `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Events        []*GetSDKInstanceDebugInfoResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	InstanceId    *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LatestBuild   *string                                          `json:"LatestBuild,omitempty" xml:"LatestBuild,omitempty"`
	Platform      *string                                          `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Recipe        *string                                          `json:"Recipe,omitempty" xml:"Recipe,omitempty"`
	Status        *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdatedAt     *string                                          `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	ValidFromDate *string                                          `json:"ValidFromDate,omitempty" xml:"ValidFromDate,omitempty"`
	ValidToDate   *string                                          `json:"ValidToDate,omitempty" xml:"ValidToDate,omitempty"`
}

func (s GetSDKInstanceDebugInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSDKInstanceDebugInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetBundleId(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.BundleId = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetCreatedAt(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetEvents(v []*GetSDKInstanceDebugInfoResponseBodyDataEvents) *GetSDKInstanceDebugInfoResponseBodyData {
	s.Events = v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetInstanceId(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetLatestBuild(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.LatestBuild = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetPlatform(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetRecipe(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.Recipe = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetStatus(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetUpdatedAt(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetValidFromDate(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.ValidFromDate = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyData) SetValidToDate(v string) *GetSDKInstanceDebugInfoResponseBodyData {
	s.ValidToDate = &v
	return s
}

type GetSDKInstanceDebugInfoResponseBodyDataEvents struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s GetSDKInstanceDebugInfoResponseBodyDataEvents) String() string {
	return tea.Prettify(s)
}

func (s GetSDKInstanceDebugInfoResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *GetSDKInstanceDebugInfoResponseBodyDataEvents) SetContent(v string) *GetSDKInstanceDebugInfoResponseBodyDataEvents {
	s.Content = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyDataEvents) SetId(v int64) *GetSDKInstanceDebugInfoResponseBodyDataEvents {
	s.Id = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyDataEvents) SetSuccess(v bool) *GetSDKInstanceDebugInfoResponseBodyDataEvents {
	s.Success = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponseBodyDataEvents) SetTimeStamp(v string) *GetSDKInstanceDebugInfoResponseBodyDataEvents {
	s.TimeStamp = &v
	return s
}

type GetSDKInstanceDebugInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSDKInstanceDebugInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSDKInstanceDebugInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSDKInstanceDebugInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSDKInstanceDebugInfoResponse) SetHeaders(v map[string]*string) *GetSDKInstanceDebugInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSDKInstanceDebugInfoResponse) SetStatusCode(v int32) *GetSDKInstanceDebugInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSDKInstanceDebugInfoResponse) SetBody(v *GetSDKInstanceDebugInfoResponseBody) *GetSDKInstanceDebugInfoResponse {
	s.Body = v
	return s
}

type QuerySDKInstancesRequest struct {
	CodeList          *string `json:"CodeList,omitempty" xml:"CodeList,omitempty"`
	CreatedAfterDate  *string `json:"CreatedAfterDate,omitempty" xml:"CreatedAfterDate,omitempty"`
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pk                *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QuerySDKInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySDKInstancesRequest) GoString() string {
	return s.String()
}

func (s *QuerySDKInstancesRequest) SetCodeList(v string) *QuerySDKInstancesRequest {
	s.CodeList = &v
	return s
}

func (s *QuerySDKInstancesRequest) SetCreatedAfterDate(v string) *QuerySDKInstancesRequest {
	s.CreatedAfterDate = &v
	return s
}

func (s *QuerySDKInstancesRequest) SetCreatedBeforeDate(v string) *QuerySDKInstancesRequest {
	s.CreatedBeforeDate = &v
	return s
}

func (s *QuerySDKInstancesRequest) SetPageNumber(v int32) *QuerySDKInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *QuerySDKInstancesRequest) SetPageSize(v int32) *QuerySDKInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySDKInstancesRequest) SetPk(v string) *QuerySDKInstancesRequest {
	s.Pk = &v
	return s
}

type QuerySDKInstancesResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QuerySDKInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *int32                             `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Ok           *bool                              `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySDKInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySDKInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySDKInstancesResponseBody) SetCode(v string) *QuerySDKInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySDKInstancesResponseBody) SetData(v *QuerySDKInstancesResponseBodyData) *QuerySDKInstancesResponseBody {
	s.Data = v
	return s
}

func (s *QuerySDKInstancesResponseBody) SetErrorMessage(v string) *QuerySDKInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QuerySDKInstancesResponseBody) SetHttpCode(v int32) *QuerySDKInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QuerySDKInstancesResponseBody) SetOk(v bool) *QuerySDKInstancesResponseBody {
	s.Ok = &v
	return s
}

func (s *QuerySDKInstancesResponseBody) SetRequestId(v string) *QuerySDKInstancesResponseBody {
	s.RequestId = &v
	return s
}

type QuerySDKInstancesResponseBodyData struct {
	Content    []*QuerySDKInstancesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySDKInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySDKInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySDKInstancesResponseBodyData) SetContent(v []*QuerySDKInstancesResponseBodyDataContent) *QuerySDKInstancesResponseBodyData {
	s.Content = v
	return s
}

func (s *QuerySDKInstancesResponseBodyData) SetPageNumber(v int32) *QuerySDKInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyData) SetPageSize(v int32) *QuerySDKInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyData) SetTotalCount(v int32) *QuerySDKInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

type QuerySDKInstancesResponseBodyDataContent struct {
	BundleId      *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	CreatedAt     *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LatestBuild   *string `json:"LatestBuild,omitempty" xml:"LatestBuild,omitempty"`
	Platform      *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Recipe        *string `json:"Recipe,omitempty" xml:"Recipe,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdatedAt     *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ValidFromDate *string `json:"ValidFromDate,omitempty" xml:"ValidFromDate,omitempty"`
	ValidToDate   *string `json:"ValidToDate,omitempty" xml:"ValidToDate,omitempty"`
}

func (s QuerySDKInstancesResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s QuerySDKInstancesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetBundleId(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.BundleId = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetCreatedAt(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.CreatedAt = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetInstanceId(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.InstanceId = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetLatestBuild(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.LatestBuild = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetPlatform(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.Platform = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetRecipe(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.Recipe = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetStatus(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetUpdatedAt(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.UpdatedAt = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetUserId(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.UserId = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetValidFromDate(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.ValidFromDate = &v
	return s
}

func (s *QuerySDKInstancesResponseBodyDataContent) SetValidToDate(v string) *QuerySDKInstancesResponseBodyDataContent {
	s.ValidToDate = &v
	return s
}

type QuerySDKInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySDKInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySDKInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySDKInstancesResponse) GoString() string {
	return s.String()
}

func (s *QuerySDKInstancesResponse) SetHeaders(v map[string]*string) *QuerySDKInstancesResponse {
	s.Headers = v
	return s
}

func (s *QuerySDKInstancesResponse) SetStatusCode(v int32) *QuerySDKInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySDKInstancesResponse) SetBody(v *QuerySDKInstancesResponseBody) *QuerySDKInstancesResponse {
	s.Body = v
	return s
}

type StartSDKInstanceProductionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartSDKInstanceProductionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSDKInstanceProductionRequest) GoString() string {
	return s.String()
}

func (s *StartSDKInstanceProductionRequest) SetInstanceId(v string) *StartSDKInstanceProductionRequest {
	s.InstanceId = &v
	return s
}

type StartSDKInstanceProductionResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Ok           *bool   `json:"Ok,omitempty" xml:"Ok,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartSDKInstanceProductionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSDKInstanceProductionResponseBody) GoString() string {
	return s.String()
}

func (s *StartSDKInstanceProductionResponseBody) SetCode(v string) *StartSDKInstanceProductionResponseBody {
	s.Code = &v
	return s
}

func (s *StartSDKInstanceProductionResponseBody) SetErrorMessage(v string) *StartSDKInstanceProductionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSDKInstanceProductionResponseBody) SetHttpCode(v int32) *StartSDKInstanceProductionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartSDKInstanceProductionResponseBody) SetOk(v bool) *StartSDKInstanceProductionResponseBody {
	s.Ok = &v
	return s
}

func (s *StartSDKInstanceProductionResponseBody) SetRequestId(v string) *StartSDKInstanceProductionResponseBody {
	s.RequestId = &v
	return s
}

type StartSDKInstanceProductionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartSDKInstanceProductionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartSDKInstanceProductionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSDKInstanceProductionResponse) GoString() string {
	return s.String()
}

func (s *StartSDKInstanceProductionResponse) SetHeaders(v map[string]*string) *StartSDKInstanceProductionResponse {
	s.Headers = v
	return s
}

func (s *StartSDKInstanceProductionResponse) SetStatusCode(v int32) *StartSDKInstanceProductionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSDKInstanceProductionResponse) SetBody(v *StartSDKInstanceProductionResponseBody) *StartSDKInstanceProductionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("viapi-oxs-cross"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateSDKInstanceWithOptions(request *CreateSDKInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateSDKInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Recipe)) {
		query["Recipe"] = request.Recipe
	}

	if !tea.BoolValue(util.IsUnset(request.ValidFrom)) {
		query["ValidFrom"] = request.ValidFrom
	}

	if !tea.BoolValue(util.IsUnset(request.ValidTo)) {
		query["ValidTo"] = request.ValidTo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSDKInstance"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSDKInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSDKInstance(request *CreateSDKInstanceRequest) (_result *CreateSDKInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSDKInstanceResponse{}
	_body, _err := client.CreateSDKInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSDKInstanceDebugInfoWithOptions(request *GetSDKInstanceDebugInfoRequest, runtime *util.RuntimeOptions) (_result *GetSDKInstanceDebugInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSDKInstanceDebugInfo"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSDKInstanceDebugInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSDKInstanceDebugInfo(request *GetSDKInstanceDebugInfoRequest) (_result *GetSDKInstanceDebugInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSDKInstanceDebugInfoResponse{}
	_body, _err := client.GetSDKInstanceDebugInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySDKInstancesWithOptions(request *QuerySDKInstancesRequest, runtime *util.RuntimeOptions) (_result *QuerySDKInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeList)) {
		query["CodeList"] = request.CodeList
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedAfterDate)) {
		query["CreatedAfterDate"] = request.CreatedAfterDate
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedBeforeDate)) {
		query["CreatedBeforeDate"] = request.CreatedBeforeDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pk)) {
		query["Pk"] = request.Pk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySDKInstances"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySDKInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySDKInstances(request *QuerySDKInstancesRequest) (_result *QuerySDKInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySDKInstancesResponse{}
	_body, _err := client.QuerySDKInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSDKInstanceProductionWithOptions(request *StartSDKInstanceProductionRequest, runtime *util.RuntimeOptions) (_result *StartSDKInstanceProductionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSDKInstanceProduction"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSDKInstanceProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSDKInstanceProduction(request *StartSDKInstanceProductionRequest) (_result *StartSDKInstanceProductionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSDKInstanceProductionResponse{}
	_body, _err := client.StartSDKInstanceProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
