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

type CreateTrailRequest struct {
	EventRW             *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	IsOrganizationTrail *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	MnsTopicArn         *string `json:"MnsTopicArn,omitempty" xml:"MnsTopicArn,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName       *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix        *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	RoleName            *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SlsProjectArn       *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn     *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	TrailRegion         *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s CreateTrailRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrailRequest) GoString() string {
	return s.String()
}

func (s *CreateTrailRequest) SetEventRW(v string) *CreateTrailRequest {
	s.EventRW = &v
	return s
}

func (s *CreateTrailRequest) SetIsOrganizationTrail(v bool) *CreateTrailRequest {
	s.IsOrganizationTrail = &v
	return s
}

func (s *CreateTrailRequest) SetMnsTopicArn(v string) *CreateTrailRequest {
	s.MnsTopicArn = &v
	return s
}

func (s *CreateTrailRequest) SetName(v string) *CreateTrailRequest {
	s.Name = &v
	return s
}

func (s *CreateTrailRequest) SetOssBucketName(v string) *CreateTrailRequest {
	s.OssBucketName = &v
	return s
}

func (s *CreateTrailRequest) SetOssKeyPrefix(v string) *CreateTrailRequest {
	s.OssKeyPrefix = &v
	return s
}

func (s *CreateTrailRequest) SetRoleName(v string) *CreateTrailRequest {
	s.RoleName = &v
	return s
}

func (s *CreateTrailRequest) SetSlsProjectArn(v string) *CreateTrailRequest {
	s.SlsProjectArn = &v
	return s
}

func (s *CreateTrailRequest) SetSlsWriteRoleArn(v string) *CreateTrailRequest {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *CreateTrailRequest) SetTrailRegion(v string) *CreateTrailRequest {
	s.TrailRegion = &v
	return s
}

type CreateTrailResponseBody struct {
	EventRW         *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	HomeRegion      *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	MnsTopicArn     *string `json:"MnsTopicArn,omitempty" xml:"MnsTopicArn,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SlsProjectArn   *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	TrailRegion     *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s CreateTrailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrailResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrailResponseBody) SetEventRW(v string) *CreateTrailResponseBody {
	s.EventRW = &v
	return s
}

func (s *CreateTrailResponseBody) SetHomeRegion(v string) *CreateTrailResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *CreateTrailResponseBody) SetMnsTopicArn(v string) *CreateTrailResponseBody {
	s.MnsTopicArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetName(v string) *CreateTrailResponseBody {
	s.Name = &v
	return s
}

func (s *CreateTrailResponseBody) SetOssBucketName(v string) *CreateTrailResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *CreateTrailResponseBody) SetOssKeyPrefix(v string) *CreateTrailResponseBody {
	s.OssKeyPrefix = &v
	return s
}

func (s *CreateTrailResponseBody) SetRequestId(v string) *CreateTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrailResponseBody) SetRoleName(v string) *CreateTrailResponseBody {
	s.RoleName = &v
	return s
}

func (s *CreateTrailResponseBody) SetSlsProjectArn(v string) *CreateTrailResponseBody {
	s.SlsProjectArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetSlsWriteRoleArn(v string) *CreateTrailResponseBody {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetTrailRegion(v string) *CreateTrailResponseBody {
	s.TrailRegion = &v
	return s
}

type CreateTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrailResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrailResponse) GoString() string {
	return s.String()
}

func (s *CreateTrailResponse) SetHeaders(v map[string]*string) *CreateTrailResponse {
	s.Headers = v
	return s
}

func (s *CreateTrailResponse) SetStatusCode(v int32) *CreateTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrailResponse) SetBody(v *CreateTrailResponseBody) *CreateTrailResponse {
	s.Body = v
	return s
}

type DeleteTrailRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteTrailRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrailRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrailRequest) SetName(v string) *DeleteTrailRequest {
	s.Name = &v
	return s
}

type DeleteTrailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrailResponseBody) SetRequestId(v string) *DeleteTrailResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrailResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrailResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrailResponse) SetHeaders(v map[string]*string) *DeleteTrailResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrailResponse) SetStatusCode(v int32) *DeleteTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrailResponse) SetBody(v *DeleteTrailResponseBody) *DeleteTrailResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTrailsRequest struct {
	IncludeOrganizationTrail *bool   `json:"IncludeOrganizationTrail,omitempty" xml:"IncludeOrganizationTrail,omitempty"`
	IncludeShadowTrails      *bool   `json:"IncludeShadowTrails,omitempty" xml:"IncludeShadowTrails,omitempty"`
	NameList                 *string `json:"NameList,omitempty" xml:"NameList,omitempty"`
}

func (s DescribeTrailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrailsRequest) SetIncludeOrganizationTrail(v bool) *DescribeTrailsRequest {
	s.IncludeOrganizationTrail = &v
	return s
}

func (s *DescribeTrailsRequest) SetIncludeShadowTrails(v bool) *DescribeTrailsRequest {
	s.IncludeShadowTrails = &v
	return s
}

func (s *DescribeTrailsRequest) SetNameList(v string) *DescribeTrailsRequest {
	s.NameList = &v
	return s
}

type DescribeTrailsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrailList []*DescribeTrailsResponseBodyTrailList `json:"TrailList,omitempty" xml:"TrailList,omitempty" type:"Repeated"`
}

func (s DescribeTrailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponseBody) SetRequestId(v string) *DescribeTrailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrailsResponseBody) SetTrailList(v []*DescribeTrailsResponseBodyTrailList) *DescribeTrailsResponseBody {
	s.TrailList = v
	return s
}

type DescribeTrailsResponseBodyTrailList struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EventRW             *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	HomeRegion          *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	IsOrganizationTrail *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	MnsTopicArn         *string `json:"MnsTopicArn,omitempty" xml:"MnsTopicArn,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName       *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix        *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	RoleName            *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SlsProjectArn       *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn     *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	StartLoggingTime    *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopLoggingTime     *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
	TrailRegion         *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTrailsResponseBodyTrailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrailsResponseBodyTrailList) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponseBodyTrailList) SetCreateTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.CreateTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetEventRW(v string) *DescribeTrailsResponseBodyTrailList {
	s.EventRW = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetHomeRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.HomeRegion = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetIsOrganizationTrail(v bool) *DescribeTrailsResponseBodyTrailList {
	s.IsOrganizationTrail = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetMnsTopicArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.MnsTopicArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetName(v string) *DescribeTrailsResponseBodyTrailList {
	s.Name = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssBucketName(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssBucketName = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssKeyPrefix(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssKeyPrefix = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetRoleName(v string) *DescribeTrailsResponseBodyTrailList {
	s.RoleName = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetSlsProjectArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.SlsProjectArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetSlsWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStartLoggingTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.StartLoggingTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStatus(v string) *DescribeTrailsResponseBodyTrailList {
	s.Status = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetStopLoggingTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.StopLoggingTime = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetTrailRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.TrailRegion = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetUpdateTime(v string) *DescribeTrailsResponseBodyTrailList {
	s.UpdateTime = &v
	return s
}

type DescribeTrailsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTrailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrailsResponse) SetHeaders(v map[string]*string) *DescribeTrailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrailsResponse) SetStatusCode(v int32) *DescribeTrailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrailsResponse) SetBody(v *DescribeTrailsResponseBody) *DescribeTrailsResponse {
	s.Body = v
	return s
}

type GetTrailStatusRequest struct {
	IsOrganizationTrail *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTrailStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrailStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTrailStatusRequest) SetIsOrganizationTrail(v bool) *GetTrailStatusRequest {
	s.IsOrganizationTrail = &v
	return s
}

func (s *GetTrailStatusRequest) SetName(v string) *GetTrailStatusRequest {
	s.Name = &v
	return s
}

type GetTrailStatusResponseBody struct {
	IsLogging           *bool   `json:"IsLogging,omitempty" xml:"IsLogging,omitempty"`
	LatestDeliveryError *string `json:"LatestDeliveryError,omitempty" xml:"LatestDeliveryError,omitempty"`
	LatestDeliveryTime  *string `json:"LatestDeliveryTime,omitempty" xml:"LatestDeliveryTime,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartLoggingTime    *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	StopLoggingTime     *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
}

func (s GetTrailStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrailStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrailStatusResponseBody) SetIsLogging(v bool) *GetTrailStatusResponseBody {
	s.IsLogging = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryError(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryError = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryTime(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetRequestId(v string) *GetTrailStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetStartLoggingTime(v string) *GetTrailStatusResponseBody {
	s.StartLoggingTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetStopLoggingTime(v string) *GetTrailStatusResponseBody {
	s.StopLoggingTime = &v
	return s
}

type GetTrailStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrailStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrailStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTrailStatusResponse) SetHeaders(v map[string]*string) *GetTrailStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTrailStatusResponse) SetStatusCode(v int32) *GetTrailStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrailStatusResponse) SetBody(v *GetTrailStatusResponseBody) *GetTrailStatusResponse {
	s.Body = v
	return s
}

type LookupEventsRequest struct {
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Event            *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventAccessKeyId *string `json:"EventAccessKeyId,omitempty" xml:"EventAccessKeyId,omitempty"`
	EventName        *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventRW          *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	EventType        *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	MaxResults       *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Request          *string `json:"Request,omitempty" xml:"Request,omitempty"`
	ResourceName     *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceName      *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	User             *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s LookupEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupEventsRequest) SetEndTime(v string) *LookupEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupEventsRequest) SetEvent(v string) *LookupEventsRequest {
	s.Event = &v
	return s
}

func (s *LookupEventsRequest) SetEventAccessKeyId(v string) *LookupEventsRequest {
	s.EventAccessKeyId = &v
	return s
}

func (s *LookupEventsRequest) SetEventName(v string) *LookupEventsRequest {
	s.EventName = &v
	return s
}

func (s *LookupEventsRequest) SetEventRW(v string) *LookupEventsRequest {
	s.EventRW = &v
	return s
}

func (s *LookupEventsRequest) SetEventType(v string) *LookupEventsRequest {
	s.EventType = &v
	return s
}

func (s *LookupEventsRequest) SetMaxResults(v string) *LookupEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *LookupEventsRequest) SetNextToken(v string) *LookupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *LookupEventsRequest) SetRequest(v string) *LookupEventsRequest {
	s.Request = &v
	return s
}

func (s *LookupEventsRequest) SetResourceName(v string) *LookupEventsRequest {
	s.ResourceName = &v
	return s
}

func (s *LookupEventsRequest) SetResourceType(v string) *LookupEventsRequest {
	s.ResourceType = &v
	return s
}

func (s *LookupEventsRequest) SetServiceName(v string) *LookupEventsRequest {
	s.ServiceName = &v
	return s
}

func (s *LookupEventsRequest) SetStartTime(v string) *LookupEventsRequest {
	s.StartTime = &v
	return s
}

func (s *LookupEventsRequest) SetUser(v string) *LookupEventsRequest {
	s.User = &v
	return s
}

type LookupEventsResponseBody struct {
	EndTime   *string                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Events    []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	NextToken *string                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime *string                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsResponseBody) GoString() string {
	return s.String()
}

func (s *LookupEventsResponseBody) SetEndTime(v string) *LookupEventsResponseBody {
	s.EndTime = &v
	return s
}

func (s *LookupEventsResponseBody) SetEvents(v []map[string]interface{}) *LookupEventsResponseBody {
	s.Events = v
	return s
}

func (s *LookupEventsResponseBody) SetNextToken(v string) *LookupEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *LookupEventsResponseBody) SetRequestId(v string) *LookupEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *LookupEventsResponseBody) SetStartTime(v string) *LookupEventsResponseBody {
	s.StartTime = &v
	return s
}

type LookupEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LookupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LookupEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsResponse) GoString() string {
	return s.String()
}

func (s *LookupEventsResponse) SetHeaders(v map[string]*string) *LookupEventsResponse {
	s.Headers = v
	return s
}

func (s *LookupEventsResponse) SetStatusCode(v int32) *LookupEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupEventsResponse) SetBody(v *LookupEventsResponseBody) *LookupEventsResponse {
	s.Body = v
	return s
}

type StartLoggingRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StartLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s StartLoggingRequest) GoString() string {
	return s.String()
}

func (s *StartLoggingRequest) SetName(v string) *StartLoggingRequest {
	s.Name = &v
	return s
}

type StartLoggingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *StartLoggingResponseBody) SetRequestId(v string) *StartLoggingResponseBody {
	s.RequestId = &v
	return s
}

type StartLoggingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s StartLoggingResponse) GoString() string {
	return s.String()
}

func (s *StartLoggingResponse) SetHeaders(v map[string]*string) *StartLoggingResponse {
	s.Headers = v
	return s
}

func (s *StartLoggingResponse) SetStatusCode(v int32) *StartLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLoggingResponse) SetBody(v *StartLoggingResponseBody) *StartLoggingResponse {
	s.Body = v
	return s
}

type StopLoggingRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StopLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s StopLoggingRequest) GoString() string {
	return s.String()
}

func (s *StopLoggingRequest) SetName(v string) *StopLoggingRequest {
	s.Name = &v
	return s
}

type StopLoggingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *StopLoggingResponseBody) SetRequestId(v string) *StopLoggingResponseBody {
	s.RequestId = &v
	return s
}

type StopLoggingResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s StopLoggingResponse) GoString() string {
	return s.String()
}

func (s *StopLoggingResponse) SetHeaders(v map[string]*string) *StopLoggingResponse {
	s.Headers = v
	return s
}

func (s *StopLoggingResponse) SetStatusCode(v int32) *StopLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLoggingResponse) SetBody(v *StopLoggingResponseBody) *StopLoggingResponse {
	s.Body = v
	return s
}

type UpdateTrailRequest struct {
	EventRW         *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	MnsTopicArn     *string `json:"MnsTopicArn,omitempty" xml:"MnsTopicArn,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SlsProjectArn   *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	TrailRegion     *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s UpdateTrailRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrailRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrailRequest) SetEventRW(v string) *UpdateTrailRequest {
	s.EventRW = &v
	return s
}

func (s *UpdateTrailRequest) SetMnsTopicArn(v string) *UpdateTrailRequest {
	s.MnsTopicArn = &v
	return s
}

func (s *UpdateTrailRequest) SetName(v string) *UpdateTrailRequest {
	s.Name = &v
	return s
}

func (s *UpdateTrailRequest) SetOssBucketName(v string) *UpdateTrailRequest {
	s.OssBucketName = &v
	return s
}

func (s *UpdateTrailRequest) SetOssKeyPrefix(v string) *UpdateTrailRequest {
	s.OssKeyPrefix = &v
	return s
}

func (s *UpdateTrailRequest) SetRoleName(v string) *UpdateTrailRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateTrailRequest) SetSlsProjectArn(v string) *UpdateTrailRequest {
	s.SlsProjectArn = &v
	return s
}

func (s *UpdateTrailRequest) SetSlsWriteRoleArn(v string) *UpdateTrailRequest {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *UpdateTrailRequest) SetTrailRegion(v string) *UpdateTrailRequest {
	s.TrailRegion = &v
	return s
}

type UpdateTrailResponseBody struct {
	EventRW         *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	HomeRegion      *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	MnsTopicArn     *string `json:"MnsTopicArn,omitempty" xml:"MnsTopicArn,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	SlsProjectArn   *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	TrailRegion     *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
}

func (s UpdateTrailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrailResponseBody) SetEventRW(v string) *UpdateTrailResponseBody {
	s.EventRW = &v
	return s
}

func (s *UpdateTrailResponseBody) SetHomeRegion(v string) *UpdateTrailResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *UpdateTrailResponseBody) SetMnsTopicArn(v string) *UpdateTrailResponseBody {
	s.MnsTopicArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetName(v string) *UpdateTrailResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateTrailResponseBody) SetOssBucketName(v string) *UpdateTrailResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *UpdateTrailResponseBody) SetOssKeyPrefix(v string) *UpdateTrailResponseBody {
	s.OssKeyPrefix = &v
	return s
}

func (s *UpdateTrailResponseBody) SetRequestId(v string) *UpdateTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrailResponseBody) SetRoleName(v string) *UpdateTrailResponseBody {
	s.RoleName = &v
	return s
}

func (s *UpdateTrailResponseBody) SetSlsProjectArn(v string) *UpdateTrailResponseBody {
	s.SlsProjectArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetSlsWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.SlsWriteRoleArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetTrailRegion(v string) *UpdateTrailResponseBody {
	s.TrailRegion = &v
	return s
}

type UpdateTrailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTrailResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrailResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrailResponse) SetHeaders(v map[string]*string) *UpdateTrailResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrailResponse) SetStatusCode(v int32) *UpdateTrailResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrailResponse) SetBody(v *UpdateTrailResponseBody) *UpdateTrailResponse {
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
		"ap-northeast-2-pop":          tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("actiontrail.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-edge-1":                   tea.String("actiontrail.aliyuncs.com"),
		"cn-fujian":                   tea.String("actiontrail.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("actiontrail.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("actiontrail.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-wuhan":                    tea.String("actiontrail.aliyuncs.com"),
		"cn-yushanfang":               tea.String("actiontrail.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("actiontrail.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("actiontrail.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("actiontrail.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("actiontrail.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("actiontrail"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateTrailWithOptions(request *CreateTrailRequest, runtime *util.RuntimeOptions) (_result *CreateTrailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventRW)) {
		query["EventRW"] = request.EventRW
	}

	if !tea.BoolValue(util.IsUnset(request.IsOrganizationTrail)) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !tea.BoolValue(util.IsUnset(request.MnsTopicArn)) {
		query["MnsTopicArn"] = request.MnsTopicArn
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssKeyPrefix)) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProjectArn)) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !tea.BoolValue(util.IsUnset(request.SlsWriteRoleArn)) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.TrailRegion)) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrail"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrail(request *CreateTrailRequest) (_result *CreateTrailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrailResponse{}
	_body, _err := client.CreateTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrailWithOptions(request *DeleteTrailRequest, runtime *util.RuntimeOptions) (_result *DeleteTrailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrail"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrail(request *DeleteTrailRequest) (_result *DeleteTrailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrailResponse{}
	_body, _err := client.DeleteTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrailsWithOptions(request *DescribeTrailsRequest, runtime *util.RuntimeOptions) (_result *DescribeTrailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeOrganizationTrail)) {
		query["IncludeOrganizationTrail"] = request.IncludeOrganizationTrail
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeShadowTrails)) {
		query["IncludeShadowTrails"] = request.IncludeShadowTrails
	}

	if !tea.BoolValue(util.IsUnset(request.NameList)) {
		query["NameList"] = request.NameList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrails"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrails(request *DescribeTrailsRequest) (_result *DescribeTrailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.DescribeTrailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrailStatusWithOptions(request *GetTrailStatusRequest, runtime *util.RuntimeOptions) (_result *GetTrailStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsOrganizationTrail)) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrailStatus"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrailStatus(request *GetTrailStatusRequest) (_result *GetTrailStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.GetTrailStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LookupEventsWithOptions(request *LookupEventsRequest, runtime *util.RuntimeOptions) (_result *LookupEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Event)) {
		query["Event"] = request.Event
	}

	if !tea.BoolValue(util.IsUnset(request.EventAccessKeyId)) {
		query["EventAccessKeyId"] = request.EventAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventRW)) {
		query["EventRW"] = request.EventRW
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Request)) {
		query["Request"] = request.Request
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LookupEvents"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LookupEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LookupEvents(request *LookupEventsRequest) (_result *LookupEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LookupEventsResponse{}
	_body, _err := client.LookupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartLoggingWithOptions(request *StartLoggingRequest, runtime *util.RuntimeOptions) (_result *StartLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartLogging"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartLogging(request *StartLoggingRequest) (_result *StartLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartLoggingResponse{}
	_body, _err := client.StartLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopLoggingWithOptions(request *StopLoggingRequest, runtime *util.RuntimeOptions) (_result *StopLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopLogging"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopLogging(request *StopLoggingRequest) (_result *StopLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopLoggingResponse{}
	_body, _err := client.StopLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTrailWithOptions(request *UpdateTrailRequest, runtime *util.RuntimeOptions) (_result *UpdateTrailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventRW)) {
		query["EventRW"] = request.EventRW
	}

	if !tea.BoolValue(util.IsUnset(request.MnsTopicArn)) {
		query["MnsTopicArn"] = request.MnsTopicArn
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssKeyPrefix)) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProjectArn)) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !tea.BoolValue(util.IsUnset(request.SlsWriteRoleArn)) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.TrailRegion)) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrail"),
		Version:     tea.String("2017-12-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTrail(request *UpdateTrailRequest) (_result *UpdateTrailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTrailResponse{}
	_body, _err := client.UpdateTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
