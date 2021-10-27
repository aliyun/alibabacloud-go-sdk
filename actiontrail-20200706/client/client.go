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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 地域名称
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 地域链接地址
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTrailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTrailsResponse) SetBody(v *DescribeTrailsResponseBody) *DescribeTrailsResponse {
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTrailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeliveryHistoryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LookupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTrailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDeliveryHistoryJob"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTrailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTrail"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDeliveryHistoryJob"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTrailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTrail"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTrails"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetDeliveryHistoryJobWithOptions(request *GetDeliveryHistoryJobRequest, runtime *util.RuntimeOptions) (_result *GetDeliveryHistoryJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeliveryHistoryJob"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTrailStatus"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDeliveryHistoryJobs"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LookupEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LookupEvents"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartLoggingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartLogging"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
		Query: query,
	}
	_result = &StopLoggingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopLogging"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTrailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTrail"), tea.String("2020-07-06"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
