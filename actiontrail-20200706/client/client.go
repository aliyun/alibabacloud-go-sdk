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

type CreateDeliveryHistoryJobRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	TrailName   *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s CreateDeliveryHistoryJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobRequest) SetClientToken(v string) *CreateDeliveryHistoryJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDeliveryHistoryJobRequest) SetTrailName(v string) *CreateDeliveryHistoryJobRequest {
	s.TrailName = &v
	return s
}

type CreateDeliveryHistoryJobResponseBody struct {
	JobId     *int32  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeliveryHistoryJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobResponseBody) SetJobId(v int32) *CreateDeliveryHistoryJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDeliveryHistoryJobResponseBody) SetRequestId(v string) *CreateDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDeliveryHistoryJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *CreateDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryHistoryJobResponse) SetStatusCode(v int32) *CreateDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryHistoryJobResponse) SetBody(v *CreateDeliveryHistoryJobResponseBody) *CreateDeliveryHistoryJobResponse {
	s.Body = v
	return s
}

type CreateTrailRequest struct {
	EventRW             *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	IsOrganizationTrail *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName       *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix        *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	OssWriteRoleArn     *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
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

func (s *CreateTrailRequest) SetOssWriteRoleArn(v string) *CreateTrailRequest {
	s.OssWriteRoleArn = &v
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
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateTrailResponseBody) SetOssWriteRoleArn(v string) *CreateTrailResponseBody {
	s.OssWriteRoleArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetRequestId(v string) *CreateTrailResponseBody {
	s.RequestId = &v
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

type DeleteDeliveryHistoryJobRequest struct {
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteDeliveryHistoryJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobRequest) SetJobId(v int32) *DeleteDeliveryHistoryJobRequest {
	s.JobId = &v
	return s
}

type DeleteDeliveryHistoryJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeliveryHistoryJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobResponseBody) SetRequestId(v string) *DeleteDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeliveryHistoryJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *DeleteDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeliveryHistoryJobResponse) SetStatusCode(v int32) *DeleteDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeliveryHistoryJobResponse) SetBody(v *DeleteDeliveryHistoryJobResponseBody) *DeleteDeliveryHistoryJobResponse {
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

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
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
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrganizationId      *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	OssBucketLocation   *string `json:"OssBucketLocation,omitempty" xml:"OssBucketLocation,omitempty"`
	OssBucketName       *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix        *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	OssWriteRoleArn     *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	Region              *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SlsProjectArn       *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	SlsWriteRoleArn     *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	StartLoggingTime    *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopLoggingTime     *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
	TrailArn            *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
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

func (s *DescribeTrailsResponseBodyTrailList) SetName(v string) *DescribeTrailsResponseBodyTrailList {
	s.Name = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOrganizationId(v string) *DescribeTrailsResponseBodyTrailList {
	s.OrganizationId = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetOssBucketLocation(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssBucketLocation = &v
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

func (s *DescribeTrailsResponseBodyTrailList) SetOssWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.OssWriteRoleArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetRegion(v string) *DescribeTrailsResponseBodyTrailList {
	s.Region = &v
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

func (s *DescribeTrailsResponseBodyTrailList) SetTrailArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.TrailArn = &v
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

type GetAccessKeyLastUsedEventsRequest struct {
	AccessKey   *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedEventsRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetNextToken(v string) *GetAccessKeyLastUsedEventsRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetPageSize(v string) *GetAccessKeyLastUsedEventsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsRequest) SetServiceName(v string) *GetAccessKeyLastUsedEventsRequest {
	s.ServiceName = &v
	return s
}

type GetAccessKeyLastUsedEventsResponseBody struct {
	Events    []*GetAccessKeyLastUsedEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	NextToken *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetEvents(v []*GetAccessKeyLastUsedEventsResponseBodyEvents) *GetAccessKeyLastUsedEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyLastUsedEventsResponseBodyEvents struct {
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	EventName     *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UsedTimestamp *int64  `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetDetail(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetEventName(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.EventName = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetSource(v string) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponseBodyEvents) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedEventsResponseBodyEvents {
	s.UsedTimestamp = &v
	return s
}

type GetAccessKeyLastUsedEventsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedEventsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedEventsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedEventsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedEventsResponse) SetBody(v *GetAccessKeyLastUsedEventsResponseBody) *GetAccessKeyLastUsedEventsResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedInfoRequest struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
}

func (s GetAccessKeyLastUsedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoRequest) SetAccessKey(v string) *GetAccessKeyLastUsedInfoRequest {
	s.AccessKey = &v
	return s
}

type GetAccessKeyLastUsedInfoResponseBody struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UsedTimestamp *int64  `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetAccessKeyLastUsedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccessKeyId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccountId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetAccountType(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.AccountType = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetDetail(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetOwnerId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceName(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceNameCn(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceNameCn = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetServiceNameEn(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.ServiceNameEn = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetSource(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedInfoResponseBody {
	s.UsedTimestamp = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponseBody) SetUserName(v string) *GetAccessKeyLastUsedInfoResponseBody {
	s.UserName = &v
	return s
}

type GetAccessKeyLastUsedInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedInfoResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedInfoResponse) SetBody(v *GetAccessKeyLastUsedInfoResponseBody) *GetAccessKeyLastUsedInfoResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedIpsRequest struct {
	AccessKey   *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedIpsRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetNextToken(v string) *GetAccessKeyLastUsedIpsRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetPageSize(v string) *GetAccessKeyLastUsedIpsRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsRequest) SetServiceName(v string) *GetAccessKeyLastUsedIpsRequest {
	s.ServiceName = &v
	return s
}

type GetAccessKeyLastUsedIpsResponseBody struct {
	Ips       []*GetAccessKeyLastUsedIpsResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetIps(v []*GetAccessKeyLastUsedIpsResponseBodyIps) *GetAccessKeyLastUsedIpsResponseBody {
	s.Ips = v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedIpsResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyLastUsedIpsResponseBodyIps struct {
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UsedTimestamp *int64  `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedIpsResponseBodyIps) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponseBodyIps) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetDetail(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetIp(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Ip = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetSource(v string) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponseBodyIps) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedIpsResponseBodyIps {
	s.UsedTimestamp = &v
	return s
}

type GetAccessKeyLastUsedIpsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedIpsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedIpsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedIpsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedIpsResponse) SetBody(v *GetAccessKeyLastUsedIpsResponseBody) *GetAccessKeyLastUsedIpsResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedProductsRequest struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
}

func (s GetAccessKeyLastUsedProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsRequest) SetAccessKey(v string) *GetAccessKeyLastUsedProductsRequest {
	s.AccessKey = &v
	return s
}

type GetAccessKeyLastUsedProductsResponseBody struct {
	Products  []*GetAccessKeyLastUsedProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponseBody) SetProducts(v []*GetAccessKeyLastUsedProductsResponseBodyProducts) *GetAccessKeyLastUsedProductsResponseBody {
	s.Products = v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedProductsResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyLastUsedProductsResponseBodyProducts struct {
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UsedTimestamp *int64  `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedProductsResponseBodyProducts) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetDetail(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceName(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceName = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceNameCn(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceNameCn = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetServiceNameEn(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.ServiceNameEn = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetSource(v string) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponseBodyProducts) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedProductsResponseBodyProducts {
	s.UsedTimestamp = &v
	return s
}

type GetAccessKeyLastUsedProductsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedProductsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponse) SetBody(v *GetAccessKeyLastUsedProductsResponseBody) *GetAccessKeyLastUsedProductsResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedResourcesRequest struct {
	AccessKey   *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetAccessKey(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.AccessKey = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetNextToken(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetPageSize(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesRequest) SetServiceName(v string) *GetAccessKeyLastUsedResourcesRequest {
	s.ServiceName = &v
	return s
}

type GetAccessKeyLastUsedResourcesResponseBody struct {
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*GetAccessKeyLastUsedResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s GetAccessKeyLastUsedResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetNextToken(v string) *GetAccessKeyLastUsedResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBody) SetResources(v []*GetAccessKeyLastUsedResourcesResponseBodyResources) *GetAccessKeyLastUsedResourcesResponseBody {
	s.Resources = v
	return s
}

type GetAccessKeyLastUsedResourcesResponseBodyResources struct {
	Detail        *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	ResourceName  *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UsedTimestamp *int64  `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
}

func (s GetAccessKeyLastUsedResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetDetail(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.Detail = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetResourceName(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetResourceType(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetSource(v string) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.Source = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponseBodyResources) SetUsedTimestamp(v int64) *GetAccessKeyLastUsedResourcesResponseBodyResources {
	s.UsedTimestamp = &v
	return s
}

type GetAccessKeyLastUsedResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccessKeyLastUsedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedResourcesResponse) SetBody(v *GetAccessKeyLastUsedResourcesResponseBody) *GetAccessKeyLastUsedResourcesResponse {
	s.Body = v
	return s
}

type GetDeliveryHistoryJobRequest struct {
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDeliveryHistoryJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryHistoryJobRequest) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobRequest) SetJobId(v int64) *GetDeliveryHistoryJobRequest {
	s.JobId = &v
	return s
}

type GetDeliveryHistoryJobResponseBody struct {
	CreatedTime *string                                    `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	EndTime     *string                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HomeRegion  *string                                    `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	JobId       *int64                                     `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobStatus   *int32                                     `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime   *string                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status      []*GetDeliveryHistoryJobResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	TrailName   *string                                    `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	UpdatedTime *string                                    `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetDeliveryHistoryJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponseBody) SetCreatedTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetEndTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetHomeRegion(v string) *GetDeliveryHistoryJobResponseBody {
	s.HomeRegion = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetJobId(v int64) *GetDeliveryHistoryJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetJobStatus(v int32) *GetDeliveryHistoryJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetRequestId(v string) *GetDeliveryHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetStartTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetStatus(v []*GetDeliveryHistoryJobResponseBodyStatus) *GetDeliveryHistoryJobResponseBody {
	s.Status = v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetTrailName(v string) *GetDeliveryHistoryJobResponseBody {
	s.TrailName = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBody) SetUpdatedTime(v string) *GetDeliveryHistoryJobResponseBody {
	s.UpdatedTime = &v
	return s
}

type GetDeliveryHistoryJobResponseBodyStatus struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeliveryHistoryJobResponseBodyStatus) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryHistoryJobResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) SetRegion(v string) *GetDeliveryHistoryJobResponseBodyStatus {
	s.Region = &v
	return s
}

func (s *GetDeliveryHistoryJobResponseBodyStatus) SetStatus(v int32) *GetDeliveryHistoryJobResponseBodyStatus {
	s.Status = &v
	return s
}

type GetDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeliveryHistoryJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *GetDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryHistoryJobResponse) SetStatusCode(v int32) *GetDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryHistoryJobResponse) SetBody(v *GetDeliveryHistoryJobResponseBody) *GetDeliveryHistoryJobResponse {
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
	IsLogging                     *bool   `json:"IsLogging,omitempty" xml:"IsLogging,omitempty"`
	LatestDeliveryError           *string `json:"LatestDeliveryError,omitempty" xml:"LatestDeliveryError,omitempty"`
	LatestDeliveryLogServiceError *string `json:"LatestDeliveryLogServiceError,omitempty" xml:"LatestDeliveryLogServiceError,omitempty"`
	LatestDeliveryLogServiceTime  *string `json:"LatestDeliveryLogServiceTime,omitempty" xml:"LatestDeliveryLogServiceTime,omitempty"`
	LatestDeliveryTime            *string `json:"LatestDeliveryTime,omitempty" xml:"LatestDeliveryTime,omitempty"`
	OssBucketStatus               *bool   `json:"OssBucketStatus,omitempty" xml:"OssBucketStatus,omitempty"`
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogStoreStatus             *bool   `json:"SlsLogStoreStatus,omitempty" xml:"SlsLogStoreStatus,omitempty"`
	StartLoggingTime              *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	StopLoggingTime               *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
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

func (s *GetTrailStatusResponseBody) SetLatestDeliveryLogServiceError(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryLogServiceError = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryLogServiceTime(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryLogServiceTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetLatestDeliveryTime(v string) *GetTrailStatusResponseBody {
	s.LatestDeliveryTime = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetOssBucketStatus(v bool) *GetTrailStatusResponseBody {
	s.OssBucketStatus = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetRequestId(v string) *GetTrailStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrailStatusResponseBody) SetSlsLogStoreStatus(v bool) *GetTrailStatusResponseBody {
	s.SlsLogStoreStatus = &v
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

type ListDeliveryHistoryJobsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeliveryHistoryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsRequest) SetPageNumber(v int32) *ListDeliveryHistoryJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeliveryHistoryJobsRequest) SetPageSize(v int32) *ListDeliveryHistoryJobsRequest {
	s.PageSize = &v
	return s
}

type ListDeliveryHistoryJobsResponseBody struct {
	DeliveryHistoryJobs []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs `json:"DeliveryHistoryJobs,omitempty" xml:"DeliveryHistoryJobs,omitempty" type:"Repeated"`
	PageNumber          *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeliveryHistoryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponseBody) SetDeliveryHistoryJobs(v []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) *ListDeliveryHistoryJobsResponseBody {
	s.DeliveryHistoryJobs = v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetPageNumber(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetPageSize(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetRequestId(v string) *ListDeliveryHistoryJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBody) SetTotalCount(v int32) *ListDeliveryHistoryJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HomeRegion  *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	JobId       *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobStatus   *int32  `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrailName   *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetCreatedTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.CreatedTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetEndTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.EndTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetHomeRegion(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.HomeRegion = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetJobId(v int64) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.JobId = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetJobStatus(v int32) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetStartTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.StartTime = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetTrailName(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.TrailName = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs) SetUpdatedTime(v string) *ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs {
	s.UpdatedTime = &v
	return s
}

type ListDeliveryHistoryJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeliveryHistoryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeliveryHistoryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeliveryHistoryJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryHistoryJobsResponse) SetHeaders(v map[string]*string) *ListDeliveryHistoryJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetStatusCode(v int32) *ListDeliveryHistoryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryHistoryJobsResponse) SetBody(v *ListDeliveryHistoryJobsResponseBody) *ListDeliveryHistoryJobsResponse {
	s.Body = v
	return s
}

type LookupEventsRequest struct {
	Direction       *string                               `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EndTime         *string                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LookupAttribute []*LookupEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	MaxResults      *string                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StartTime       *string                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LookupEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsRequest) GoString() string {
	return s.String()
}

func (s *LookupEventsRequest) SetDirection(v string) *LookupEventsRequest {
	s.Direction = &v
	return s
}

func (s *LookupEventsRequest) SetEndTime(v string) *LookupEventsRequest {
	s.EndTime = &v
	return s
}

func (s *LookupEventsRequest) SetLookupAttribute(v []*LookupEventsRequestLookupAttribute) *LookupEventsRequest {
	s.LookupAttribute = v
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

func (s *LookupEventsRequest) SetStartTime(v string) *LookupEventsRequest {
	s.StartTime = &v
	return s
}

type LookupEventsRequestLookupAttribute struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LookupEventsRequestLookupAttribute) String() string {
	return tea.Prettify(s)
}

func (s LookupEventsRequestLookupAttribute) GoString() string {
	return s.String()
}

func (s *LookupEventsRequestLookupAttribute) SetKey(v string) *LookupEventsRequestLookupAttribute {
	s.Key = &v
	return s
}

func (s *LookupEventsRequestLookupAttribute) SetValue(v string) *LookupEventsRequestLookupAttribute {
	s.Value = &v
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
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
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

func (s *UpdateTrailRequest) SetOssWriteRoleArn(v string) *UpdateTrailRequest {
	s.OssWriteRoleArn = &v
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
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucketName   *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssKeyPrefix    *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *UpdateTrailResponseBody) SetOssWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.OssWriteRoleArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetRequestId(v string) *UpdateTrailResponseBody {
	s.RequestId = &v
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

func (client *Client) CreateDeliveryHistoryJobWithOptions(request *CreateDeliveryHistoryJobRequest, runtime *util.RuntimeOptions) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.TrailName)) {
		query["TrailName"] = request.TrailName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeliveryHistoryJob"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDeliveryHistoryJob(request *CreateDeliveryHistoryJobRequest) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CreateDeliveryHistoryJobWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssKeyPrefix)) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OssWriteRoleArn)) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
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
		Version:     tea.String("2020-07-06"),
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

func (client *Client) DeleteDeliveryHistoryJobWithOptions(request *DeleteDeliveryHistoryJobRequest, runtime *util.RuntimeOptions) (_result *DeleteDeliveryHistoryJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeliveryHistoryJob"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeliveryHistoryJob(request *DeleteDeliveryHistoryJobRequest) (_result *DeleteDeliveryHistoryJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.DeleteDeliveryHistoryJobWithOptions(request, runtime)
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
		Version:     tea.String("2020-07-06"),
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-07-06"),
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
		Version:     tea.String("2020-07-06"),
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

func (client *Client) GetAccessKeyLastUsedEventsWithOptions(request *GetAccessKeyLastUsedEventsRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsedEvents"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedEvents(request *GetAccessKeyLastUsedEventsRequest) (_result *GetAccessKeyLastUsedEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedEventsResponse{}
	_body, _err := client.GetAccessKeyLastUsedEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedInfoWithOptions(request *GetAccessKeyLastUsedInfoRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsedInfo"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedInfo(request *GetAccessKeyLastUsedInfoRequest) (_result *GetAccessKeyLastUsedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedInfoResponse{}
	_body, _err := client.GetAccessKeyLastUsedInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedIpsWithOptions(request *GetAccessKeyLastUsedIpsRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsedIps"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedIps(request *GetAccessKeyLastUsedIpsRequest) (_result *GetAccessKeyLastUsedIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedIpsResponse{}
	_body, _err := client.GetAccessKeyLastUsedIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedProductsWithOptions(request *GetAccessKeyLastUsedProductsRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsedProducts"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedProducts(request *GetAccessKeyLastUsedProductsRequest) (_result *GetAccessKeyLastUsedProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedProductsResponse{}
	_body, _err := client.GetAccessKeyLastUsedProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedResourcesWithOptions(request *GetAccessKeyLastUsedResourcesRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsedResources"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedResources(request *GetAccessKeyLastUsedResourcesRequest) (_result *GetAccessKeyLastUsedResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResourcesResponse{}
	_body, _err := client.GetAccessKeyLastUsedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeliveryHistoryJobWithOptions(request *GetDeliveryHistoryJobRequest, runtime *util.RuntimeOptions) (_result *GetDeliveryHistoryJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeliveryHistoryJob"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeliveryHistoryJob(request *GetDeliveryHistoryJobRequest) (_result *GetDeliveryHistoryJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.GetDeliveryHistoryJobWithOptions(request, runtime)
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
		Version:     tea.String("2020-07-06"),
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

func (client *Client) ListDeliveryHistoryJobsWithOptions(request *ListDeliveryHistoryJobsRequest, runtime *util.RuntimeOptions) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeliveryHistoryJobs"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeliveryHistoryJobs(request *ListDeliveryHistoryJobsRequest) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.ListDeliveryHistoryJobsWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.LookupAttribute)) {
		query["LookupAttribute"] = request.LookupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LookupEvents"),
		Version:     tea.String("2020-07-06"),
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartLogging"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
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
		Version:     tea.String("2020-07-06"),
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

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssKeyPrefix)) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OssWriteRoleArn)) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
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
		Version:     tea.String("2020-07-06"),
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
