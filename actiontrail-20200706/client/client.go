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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can contain only ASCII characters and can be up to 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the trail for which you want to create a historical event delivery task.
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
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
	// The ID of the historical event delivery task.
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
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
	// The read/write type of the events to be delivered. Valid values:
	//
	// *   Write: write events. It is the default value.
	// *   Read: read events.
	// *   All: read and write events.
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// Specifies whether to create a multi-account trail. Valid values:
	//
	// *   true: creates a multi-account trail.
	// *   false (default): creates a single-account trail.
	IsOrganizationTrail    *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	MaxComputeProjectArn   *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail to be created.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (\_).
	//
	// > The name must be unique within your Alibaba Cloud account.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket to which events are to be delivered.
	//
	// The name must be 3 to 63 characters in length. The name must start with a lowercase letter or a digit and can contain lowercase letters, digits, and hyphens (-).
	//
	// > You must specify at least one of the OssBucketName and SlsProjectArn parameters.
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the log files to be stored in the destination OSS bucket. This parameter can be left empty.
	//
	// The prefix must be 6 to 32 characters in length. The prefix must start with a letter and can contain letters, digits, hyphens (-), forward slashes (/), and underscores (\_).
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	//
	// *   If you do not specify this parameter, ActionTrail creates a service-linked role to create the required resources. For more information, see [Manage the service-linked role](~~169244~~).
	// *   If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](~~207462~~).
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ARN of the Log Service project to which events are to be delivered.
	//
	// > You must specify at least one of the OssBucketName and SlsProjectArn parameters.
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the Log Service project.
	//
	// *   If you do not specify this parameter, ActionTrail creates a service-linked role to create the corresponding resource. For more information, see [Manage the service-linked role](~~169244~~).
	// *   If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](~~207462~~).
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The one or more regions from which the trail delivers events.
	//
	// The default value is All, which indicates that the trail delivers events from all regions.
	//
	// You can also specify specific regions. You can call the [DescribeRegions](~~213597~~) operation to query all the supported regions.
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
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

func (s *CreateTrailRequest) SetMaxComputeProjectArn(v string) *CreateTrailRequest {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *CreateTrailRequest) SetMaxComputeWriteRoleArn(v string) *CreateTrailRequest {
	s.MaxComputeWriteRoleArn = &v
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
	// The read/write type of the events to be delivered.
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The home region of the trail.
	HomeRegion             *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	MaxComputeProjectArn   *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket to which events are to be delivered.
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the log files to be stored in the destination OSS bucket.
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The ARN of the service-linked role that is assumed by ActionTrail to deliver events to the destination OSS bucket.
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the Log Service project to which events are to be delivered.
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the service-linked role that is assumed by ActionTrail to deliver events to the destination Log Service project.
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The one or more regions from which the trail delivers events.
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
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

func (s *CreateTrailResponseBody) SetMaxComputeProjectArn(v string) *CreateTrailResponseBody {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *CreateTrailResponseBody) SetMaxComputeWriteRoleArn(v string) *CreateTrailResponseBody {
	s.MaxComputeWriteRoleArn = &v
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
	// The ID of the historical event delivery task to be deleted.
	//
	// You can call the [ListDeliveryHistoryJobs](~~188101~~) operation to query task IDs.
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
	// The ID of the request.
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
	// The name of the trail that you want to delete.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (\_).
	//
	// > The name must be unique within your Alibaba Cloud account.
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
	// The ID of the request.
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
	// The language in which the region names are returned. Valid values:
	//
	// - zh-CN: Chinese.
	// - en-US: English. It is the default value.
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
	// The regions returned.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The name of the region.
	//
	// > If the AcceptLanguage parameter is set to zh-CN, the Chinese name of the region is returned. If the AcceptLanguage parameter is set to zh-US or left empty, the English name of the region is returned.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of ActionTrail in the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
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
	// Specifies whether to query the information about multi-account trails. Valid values:
	//
	// *   true
	// *   false (default)
	IncludeOrganizationTrail *bool `json:"IncludeOrganizationTrail,omitempty" xml:"IncludeOrganizationTrail,omitempty"`
	// Specifies whether to return the information about shadow trails. Valid values:
	//
	// *   false: Do not return the information about shadow trails. It is the default value.
	// *   true: Return the information about shadow trails.
	IncludeShadowTrails *bool `json:"IncludeShadowTrails,omitempty" xml:"IncludeShadowTrails,omitempty"`
	// The names of the trails whose information you want to query. Separate multiple trail names with commas (,).
	NameList *string `json:"NameList,omitempty" xml:"NameList,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of returned trails.
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
	// The time when the trail was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The read/write type of the events that are delivered. Valid values:
	//
	// *   Write: write events. This is the default value.
	// *   Read: read events.
	// *   All: read and write events.
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The home region of the trail.
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// Indicates whether the trail is a multi-account trail. Valid values:
	//
	// *   false (default)
	// *   true
	IsOrganizationTrail    *bool   `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	MaxComputeProjectArn   *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource directory.
	//
	// >  This parameter is returned only when the trail is a multi-account trail.
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The region where the OSS bucket resides.
	OssBucketLocation *string `json:"OssBucketLocation,omitempty" xml:"OssBucketLocation,omitempty"`
	// The name of the OSS bucket to which events are delivered.
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the files that are stored in the Object Storage Service (OSS) bucket.
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The region where the trail resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ARN of the Log Service project to which events are delivered.
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the Log Service project.
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The time when the trail was last enabled.
	StartLoggingTime *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	// The status of the trail. Valid values:
	//
	// *   Disable: disabled.
	// *   Enable: enabled.
	// *   Fresh: The trail is created but is not enabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the trail was last disabled.
	StopLoggingTime *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
	// The ARN of the trail.
	TrailArn *string `json:"TrailArn,omitempty" xml:"TrailArn,omitempty"`
	// The region of the trail.
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
	// The time when the configurations of the trail were last updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *DescribeTrailsResponseBodyTrailList) SetMaxComputeProjectArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *DescribeTrailsResponseBodyTrailList) SetMaxComputeWriteRoleArn(v string) *DescribeTrailsResponseBodyTrailList {
	s.MaxComputeWriteRoleArn = &v
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
	// The AccessKey ID.
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The token that determines the start point of the query.
	//
	// > The request parameters must be the same as those of the last request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 20.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud service. For more information about the Alibaba Cloud services supported by ActionTrail, see [Supported Alibaba Cloud services](~~28829~~).
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
	// The list of returned events.
	Events []*GetAccessKeyLastUsedEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The token that determines the start point of the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// An array that consists of the details about the event.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The event source.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the event was generated.
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
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
	// The AccessKey secret.
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
	// The AccessKey ID.
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The ID of the Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The type of the account to which the AccessKey pair belongs.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The details about the event.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the account to which the AccessKey pair belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud service that was last accessed.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The Chinese name of the Alibaba Cloud service that was last accessed.
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	// The English name of the Alibaba Cloud service that was last accessed.
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	// The event source.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the AccessKey pair was last called.
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
	// The name of the account to which the AccessKey pair belongs.
	//
	// If the value of the AccountType parameter is root-account, the value of the UserName parameter is root. If the value of the AccountType parameter is ram-user, the value of the UserName parameter is the name of a RAM user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	// The AccessKey ID.
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The token that determines the start point of the query.
	//
	// > The request parameters must be the same as those of the last request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 20.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud service. For more information about the Alibaba Cloud services supported by ActionTrail, see [Supported Alibaba Cloud services](~~28829~~).
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
	// The list of returned IP addresses.
	Ips []*GetAccessKeyLastUsedIpsResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The token that determines the start point of the query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// An array that consists of the details about the event.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The event source.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the IP address was used.
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
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
	// The AccessKey ID.
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
	// The list of returned Alibaba Cloud services.
	Products []*GetAccessKeyLastUsedProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The event details.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The Alibaba Cloud service.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The Chinese name of the Alibaba Cloud service.
	ServiceNameCn *string `json:"ServiceNameCn,omitempty" xml:"ServiceNameCn,omitempty"`
	// The English name of the Alibaba Cloud service.
	ServiceNameEn *string `json:"ServiceNameEn,omitempty" xml:"ServiceNameEn,omitempty"`
	// The event source.
	//
	// Valid values:
	//
	// *   Internal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     other events
	//
	//     <!-- -->
	//
	// *   ManagementEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     management events
	//
	//     <!-- -->
	//
	// *   DataEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     data events
	//
	//     <!-- -->
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Unit: millisecond.
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
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
	// The AccessKey ID.
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// > The request parameters must be the same as those of the last request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// *   Valid values: 0 to 100.
	// *   Default value: 20.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud service. For more information about the Alibaba Cloud services supported by ActionTrail, see [Supported Alibaba Cloud services](~~28829~~).
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
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned resources.
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
	// The event details.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The resource name.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The event source.
	//
	// Valid values:
	//
	// *   Internal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     other events
	//
	//     <!-- -->
	//
	// *   ManagementEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     management events
	//
	//     <!-- -->
	//
	// *   DataEvent
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     data events
	//
	//     <!-- -->
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The timestamp when the resource was used. Unit: millisecond.
	UsedTimestamp *int64 `json:"UsedTimestamp,omitempty" xml:"UsedTimestamp,omitempty"`
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
	// The ID of the historical event delivery task.
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
	// The time when the task was created.
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the task ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The home region of the trail.
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The ID of the task.
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task status. Valid values:
	//
	// *   0: The task is initializing.
	// *   1: The task is delivering historical events.
	// *   2: The task is complete.
	// *   3: The task fails.
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the task started.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of task statuses in each region.
	Status []*GetDeliveryHistoryJobResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The name of the trail based on which the task delivers events.
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// The time when the task was updated.
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
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
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The task status in each region. Valid values:
	//
	// *   0: The task is initializing.
	// *   1: The task is delivering historical events.
	// *   2: The task is complete.
	// *   3: The task fails.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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

type GetGlobalEventsStorageRegionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The region where global events are stored.
	//
	// Valid values:
	//
	// *   ap-southeast-1
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the Singapore region
	//
	//     <!-- -->
	//
	// *   cn-hangzhou
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the China (Hangzhou) region
	//
	//     <!-- -->
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s GetGlobalEventsStorageRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalEventsStorageRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalEventsStorageRegionResponseBody) SetRequestId(v string) *GetGlobalEventsStorageRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalEventsStorageRegionResponseBody) SetStorageRegion(v string) *GetGlobalEventsStorageRegionResponseBody {
	s.StorageRegion = &v
	return s
}

type GetGlobalEventsStorageRegionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGlobalEventsStorageRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGlobalEventsStorageRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalEventsStorageRegionResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalEventsStorageRegionResponse) SetHeaders(v map[string]*string) *GetGlobalEventsStorageRegionResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalEventsStorageRegionResponse) SetStatusCode(v int32) *GetGlobalEventsStorageRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGlobalEventsStorageRegionResponse) SetBody(v *GetGlobalEventsStorageRegionResponseBody) *GetGlobalEventsStorageRegionResponse {
	s.Body = v
	return s
}

type GetTrailStatusRequest struct {
	// Specifies whether to query the status of a multi-account trail. Valid values:
	//
	// *   true: Query the status of a multi-account trail.
	// *   false: Query the status of a single-account trail. It is the default value.
	IsOrganizationTrail *bool `json:"IsOrganizationTrail,omitempty" xml:"IsOrganizationTrail,omitempty"`
	// The name of the trail.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (\_).
	//
	// > The name must be unique within your Alibaba Cloud account.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// Indicates whether logging is enabled for the trail. Valid values:
	//
	// *   true
	// *   false
	IsLogging *bool `json:"IsLogging,omitempty" xml:"IsLogging,omitempty"`
	// The log of the last failed delivery.
	LatestDeliveryError *string `json:"LatestDeliveryError,omitempty" xml:"LatestDeliveryError,omitempty"`
	// The log of the last failed delivery to Log Service.
	LatestDeliveryLogServiceError *string `json:"LatestDeliveryLogServiceError,omitempty" xml:"LatestDeliveryLogServiceError,omitempty"`
	// The most recent time when an event was delivered to Log Service.
	LatestDeliveryLogServiceTime *string `json:"LatestDeliveryLogServiceTime,omitempty" xml:"LatestDeliveryLogServiceTime,omitempty"`
	// The most recent time when an event was delivered by the trail.
	LatestDeliveryTime *string `json:"LatestDeliveryTime,omitempty" xml:"LatestDeliveryTime,omitempty"`
	// Indicates whether the destination Object Storage Service (OSS) bucket is available. Valid values:
	//
	// *   true
	// *   false
	OssBucketStatus *bool `json:"OssBucketStatus,omitempty" xml:"OssBucketStatus,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the destination Log Service Logstore is available. Valid values:
	//
	// *   true
	// *   false
	SlsLogStoreStatus *bool `json:"SlsLogStoreStatus,omitempty" xml:"SlsLogStoreStatus,omitempty"`
	// The time when logging was last enabled for the trail.
	StartLoggingTime *string `json:"StartLoggingTime,omitempty" xml:"StartLoggingTime,omitempty"`
	// The time when logging was last disabled for the trail.
	StopLoggingTime *string `json:"StopLoggingTime,omitempty" xml:"StopLoggingTime,omitempty"`
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
	// The page number.
	//
	// *   Pages start from page 1.
	// *   Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// *   Valid values: 1 to 100.
	// *   Default value: 20.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The list of historical event delivery tasks.
	DeliveryHistoryJobs []*ListDeliveryHistoryJobsResponseBodyDeliveryHistoryJobs `json:"DeliveryHistoryJobs,omitempty" xml:"DeliveryHistoryJobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of historical event delivery tasks returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The time when the task was created.
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the task ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The home region of the trail.
	HomeRegion *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	// The task ID.
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task status. Valid values:
	//
	// *   0: The task is initializing.
	// *   1: The task is delivering historical events.
	// *   2: The task is complete.
	// *   3: The task fails.
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The time when the task started.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the trail.
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
	// The time when the task was updated.
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
	// The order in which details of events are to be retrieved. Valid values:
	//
	// *   FORWARD: ascending order.
	// *   BACKWARD: descending order. This is the default value.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. The default time is the current time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Query conditions.
	LookupAttribute []*LookupEventsRequestLookupAttribute `json:"LookupAttribute,omitempty" xml:"LookupAttribute,omitempty" type:"Repeated"`
	// The maximum number of entries to be returned.
	//
	// Valid values: 0 to 50.
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to request the next page of query results.
	//
	// > The request parameters must be the same as those of the last request.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query. The default time is seven days prior to the current time. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The key of the query condition. Valid values:
	//
	// *  ServiceName: the name of a specific Alibaba Cloud service.
	// *  EventName: the name of a specific event.
	// *  User: the name of the RAM user who calls a specific operation.
	// *  EventId: the ID of a specific event.
	// *  ResourceType: the type of resources.
	// *   ResourceName: the name of a specific resource.
	// *   EventRW: the read/write type of events.
	// *  EventAccessKeyId: the AccessKey ID used in events.
	//
	// > You can use only one query condition for each query.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the query condition. Valid values:
	//
	// *   When the LookupAttribute.N.Key parameter is set to ServiceName, you can set this parameter to a value such as `Ecs`.
	// *   When the LookupAttribute.N.Key parameter is set to EventName, you can set this parameter to a value such as `ConsoleSignin`.
	// *   When the LookupAttribute.N.Key parameter is set to User, you can set this parameter to a value such as `Alice`.
	// *   When the LookupAttribute.N.Key parameter is set to EventId, you can set this parameter to a value such as `B702AFA3-FD4B-40E3-88E4-C0752FAA****`.
	// *   When the LookupAttribute.N.Key parameter is set to ResourceType, you can set this parameter to a value such as `ACS::ECS::Instance`.
	// *   When the LookupAttribute.N.Key parameter is set to ResourceName, you can set this parameter to a value such as `i-bp14664y88udkt45****`.
	// *   When the LookupAttribute.N.Key parameter is set to EventRW, you can set this parameter to `Read` or `Write`.
	// *   When the LookupAttribute.N.Key parameter is set to EventAccessKeyId, you can set this parameter to a value such as `LTAI4FoDkCf4DU1bic1V****`.
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
	// The end of the time range when event details were queried.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The returned event details.
	//
	// For more information about the fields in an event log, see [ActionTrail event log reference](~~28819~~).
	Events []map[string]interface{} `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The token used to return the next page of query results.
	//
	// > This parameter is not returned if no more results are to be returned.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range when event details were queried.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The name of the trail that you want to enable.
	//
	// The name must be 6 to 36 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (\_). It must start with a lowercase letter.
	//
	// > The name must be unique within your Alibaba Cloud account.
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
	// The ID of the request.
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
	// The name of the trail that you want to disable.
	//
	// The name must be 6 to 36 characters in length, and can contain lowercase letters, digits, hyphens (-), and underscores (\_). It must start with a lowercase letter.
	//
	// > The name must be unique within your Alibaba Cloud account.
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
	// The ID of the request.
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

type UpdateGlobalEventsStorageRegionRequest struct {
	// The region where you want to store global events.
	//
	// Valid values:
	//
	// *   ap-southeast-1
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the Singapore region
	//
	//     <!-- -->
	//
	// *   cn-hangzhou
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     the China (Hangzhou) region
	//
	//     <!-- -->
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s UpdateGlobalEventsStorageRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionRequest) SetStorageRegion(v string) *UpdateGlobalEventsStorageRegionRequest {
	s.StorageRegion = &v
	return s
}

type UpdateGlobalEventsStorageRegionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGlobalEventsStorageRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionResponseBody) SetRequestId(v string) *UpdateGlobalEventsStorageRegionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGlobalEventsStorageRegionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGlobalEventsStorageRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGlobalEventsStorageRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetHeaders(v map[string]*string) *UpdateGlobalEventsStorageRegionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetStatusCode(v int32) *UpdateGlobalEventsStorageRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponse) SetBody(v *UpdateGlobalEventsStorageRegionResponseBody) *UpdateGlobalEventsStorageRegionResponse {
	s.Body = v
	return s
}

type UpdateTrailRequest struct {
	// The read/write type of the events to be delivered. Valid values:
	//
	// *   Write: write events. It is the default value.
	// *   Read: read events.
	// *   All: read and write events.
	EventRW                *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	MaxComputeProjectArn   *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail whose configurations you want to update.
	//
	// The name must be 6 to 36 characters in length and can contain lowercase letters, digits, hyphens (-), and underscores (\_). It must start with a lowercase letter.
	//
	// >  The name must be unique within an Alibaba Cloud account.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Object Storage Service (OSS) bucket to which you want to deliver events.
	//
	// The name must be 3 to 63 characters in length. The name must start with a lowercase letter or a digit and can contain lowercase letters, digits, and hyphens (-).
	//
	// >  Make sure that the bucket exists before you update the configuration of the trail.
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the files that are stored in the OSS bucket.
	//
	// The prefix must be 6 to 32 characters in length. The prefix must start with a letter and can contain letters, digits, hyphens (-), forward slashes (/), and underscores (\_).
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	//
	// *   If you do not specify this parameter, ActionTrail creates a service-linked role to create the required resources. For more information, see [Manage the service-linked role](~~169244~~).
	// *   If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](~~207462~~).
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ARN of the Log Service project to which you want to deliver events.
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the Log Service project.
	//
	// *   If you do not specify this parameter, ActionTrail creates a service-linked role to create the corresponding resource. For more information, see [Manage the service-linked role](~~169244~~).
	// *   If you specify this parameter, you must grant the permissions of the service-linked role that is assumed by ActionTrail to the RAM role before you can deliver events to your Alibaba Cloud account. If you need to deliver events to other Alibaba Cloud accounts, you must attach the permission policy that is used to grant permissions related to event delivery to the RAM role. For more information about how to deliver events across Alibaba Cloud accounts, see [Deliver events across Alibaba Cloud accounts](~~207462~~).
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The region of the trail.
	//
	// *   The default value is All, which indicates that the trail delivers events from all regions.
	//
	// You can also specify specific regions. You can call the [DescribeRegions](~~213597~~) operation to query all the supported regions.
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
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

func (s *UpdateTrailRequest) SetMaxComputeProjectArn(v string) *UpdateTrailRequest {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *UpdateTrailRequest) SetMaxComputeWriteRoleArn(v string) *UpdateTrailRequest {
	s.MaxComputeWriteRoleArn = &v
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
	// The read/write type of the events to be delivered.
	EventRW *string `json:"EventRW,omitempty" xml:"EventRW,omitempty"`
	// The home region of the trail.
	HomeRegion             *string `json:"HomeRegion,omitempty" xml:"HomeRegion,omitempty"`
	MaxComputeProjectArn   *string `json:"MaxComputeProjectArn,omitempty" xml:"MaxComputeProjectArn,omitempty"`
	MaxComputeWriteRoleArn *string `json:"MaxComputeWriteRoleArn,omitempty" xml:"MaxComputeWriteRoleArn,omitempty"`
	// The name of the trail.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the OSS bucket.
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The prefix of the log files to be stored in the destination OSS bucket.
	OssKeyPrefix *string `json:"OssKeyPrefix,omitempty" xml:"OssKeyPrefix,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail to deliver events to the OSS bucket.
	OssWriteRoleArn *string `json:"OssWriteRoleArn,omitempty" xml:"OssWriteRoleArn,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the Log Service project to which events are to be delivered.
	SlsProjectArn *string `json:"SlsProjectArn,omitempty" xml:"SlsProjectArn,omitempty"`
	// The ARN of the RAM role that is assumed by ActionTrail is to deliver events to the Log Service project.
	SlsWriteRoleArn *string `json:"SlsWriteRoleArn,omitempty" xml:"SlsWriteRoleArn,omitempty"`
	// The one or more regions from which the trail delivers events.
	TrailRegion *string `json:"TrailRegion,omitempty" xml:"TrailRegion,omitempty"`
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

func (s *UpdateTrailResponseBody) SetMaxComputeProjectArn(v string) *UpdateTrailResponseBody {
	s.MaxComputeProjectArn = &v
	return s
}

func (s *UpdateTrailResponseBody) SetMaxComputeWriteRoleArn(v string) *UpdateTrailResponseBody {
	s.MaxComputeWriteRoleArn = &v
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

/**
 * Take note of the following limits:
 * - You must have created and configured a single-account trail to deliver events to Log Service by calling the [CreateTrail](~~212313~~) operation.
 * - Only one historical event delivery task can be running at a time within an Alibaba Cloud account.
 * This topic shows you how to create a historical event delivery task for a sample trail named `trail-name`.
 *
 * @param request CreateDeliveryHistoryJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDeliveryHistoryJobResponse
 */
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

/**
 * Take note of the following limits:
 * - You must have created and configured a single-account trail to deliver events to Log Service by calling the [CreateTrail](~~212313~~) operation.
 * - Only one historical event delivery task can be running at a time within an Alibaba Cloud account.
 * This topic shows you how to create a historical event delivery task for a sample trail named `trail-name`.
 *
 * @param request CreateDeliveryHistoryJobRequest
 * @return CreateDeliveryHistoryJobResponse
 */
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

/**
 * You can create a trail to deliver events to Log Service, Object Storage Service (OSS), or both. Before you call this operation to create a trail, make sure that the following requirements are met:
 * *   Deliver events to Log Service: A project is created in Log Service.
 * **
 * **Description** After you create a trail to deliver events to Log Service, a Logstore whose name is in the `actiontrail_<Trail name>` format is automatically created and optimally configured for subsequent auditing. Indexes and a dashboard are created for the Logstore to facilitate event queries. You cannot manually write data to the Logstore. This ensures data accuracy. You do not need to create a Logstore in advance.
 * *   Deliver events to OSS: A bucket is created in OSS. This topic provides an example on how to call the API operation to create a single-account trail named `trail-test` to deliver events to an OSS bucket named `audit-log`.
 *
 * @param request CreateTrailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTrailResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.MaxComputeProjectArn)) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !tea.BoolValue(util.IsUnset(request.MaxComputeWriteRoleArn)) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
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

/**
 * You can create a trail to deliver events to Log Service, Object Storage Service (OSS), or both. Before you call this operation to create a trail, make sure that the following requirements are met:
 * *   Deliver events to Log Service: A project is created in Log Service.
 * **
 * **Description** After you create a trail to deliver events to Log Service, a Logstore whose name is in the `actiontrail_<Trail name>` format is automatically created and optimally configured for subsequent auditing. Indexes and a dashboard are created for the Logstore to facilitate event queries. You cannot manually write data to the Logstore. This ensures data accuracy. You do not need to create a Logstore in advance.
 * *   Deliver events to OSS: A bucket is created in OSS. This topic provides an example on how to call the API operation to create a single-account trail named `trail-test` to deliver events to an OSS bucket named `audit-log`.
 *
 * @param request CreateTrailRequest
 * @return CreateTrailResponse
 */
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

/**
 * This topic describes how to delete a sample historical event delivery task whose ID is `16602`.
 *
 * @param request DeleteDeliveryHistoryJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDeliveryHistoryJobResponse
 */
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

/**
 * This topic describes how to delete a sample historical event delivery task whose ID is `16602`.
 *
 * @param request DeleteDeliveryHistoryJobRequest
 * @return DeleteDeliveryHistoryJobResponse
 */
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

/**
 * This topic describes how to delete a sample trail named `trail-test`.
 *
 * @param request DeleteTrailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTrailResponse
 */
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

/**
 * This topic describes how to delete a sample trail named `trail-test`.
 *
 * @param request DeleteTrailRequest
 * @return DeleteTrailResponse
 */
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

/**
 * For more information, see [Regions and zones](~~40654~~).
 *
 * @param request DescribeRegionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRegionsResponse
 */
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

/**
 * For more information, see [Regions and zones](~~40654~~).
 *
 * @param request DescribeRegionsRequest
 * @return DescribeRegionsResponse
 */
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

/**
 * This topic shows you how to query the information about the single-account trails within an Alibaba Cloud account. In this example, the information about a trail named `test-4` is returned.
 *
 * @param request DescribeTrailsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeTrailsResponse
 */
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

/**
 * This topic shows you how to query the information about the single-account trails within an Alibaba Cloud account. In this example, the information about a trail named `test-4` is returned.
 *
 * @param request DescribeTrailsRequest
 * @return DescribeTrailsResponse
 */
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

/**
 * You can call this operation to query only the information about the most recent events that are generated within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. For more information about supported events, see [Alibaba Cloud services and events that are supported by the AccessKey pair audit feature](~~419214~~). Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedEventsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessKeyLastUsedEventsResponse
 */
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

/**
 * You can call this operation to query only the information about the most recent events that are generated within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. For more information about supported events, see [Alibaba Cloud services and events that are supported by the AccessKey pair audit feature](~~419214~~). Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedEventsRequest
 * @return GetAccessKeyLastUsedEventsResponse
 */
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

/**
 * You can call this operation to query only the information about the most recent call of a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessKeyLastUsedInfoResponse
 */
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

/**
 * You can call this operation to query only the information about the most recent call of a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedInfoRequest
 * @return GetAccessKeyLastUsedInfoResponse
 */
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

/**
 * You can call this operation to query only the information about the IP addresses that are most recently used within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedIpsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessKeyLastUsedIpsResponse
 */
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

/**
 * You can call this operation to query only the information about the IP addresses that are most recently used within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedIpsRequest
 * @return GetAccessKeyLastUsedIpsResponse
 */
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

/**
 * You can call this operation to query only the information about Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedProductsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessKeyLastUsedProductsResponse
 */
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

/**
 * You can call this operation to query only the information about Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedProductsRequest
 * @return GetAccessKeyLastUsedProductsResponse
 */
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

/**
 * You can call this operation to query only the information about resources that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAccessKeyLastUsedResourcesResponse
 */
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

/**
 * You can call this operation to query only the information about resources that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
 *
 * @param request GetAccessKeyLastUsedResourcesRequest
 * @return GetAccessKeyLastUsedResourcesResponse
 */
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

/**
 * This topic describes how to query the details of a historical event delivery tasks created within your Alibaba Cloud account. In this example, the details of a historical event delivery task whose ID is `16602` are returned. The sample response shows that this task is used to deliver the historical events recorded by the trail named `trail-name` to Log Service and the task is complete.
 *
 * @param request GetDeliveryHistoryJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDeliveryHistoryJobResponse
 */
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

/**
 * This topic describes how to query the details of a historical event delivery tasks created within your Alibaba Cloud account. In this example, the details of a historical event delivery task whose ID is `16602` are returned. The sample response shows that this task is used to deliver the historical events recorded by the trail named `trail-name` to Log Service and the task is complete.
 *
 * @param request GetDeliveryHistoryJobRequest
 * @return GetDeliveryHistoryJobResponse
 */
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

/**
 * By default, global events are stored in the Singapore region.
 * To obtain the permissions to call the API operation, you must submit a ticket.
 *
 * @param request GetGlobalEventsStorageRegionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetGlobalEventsStorageRegionResponse
 */
func (client *Client) GetGlobalEventsStorageRegionWithOptions(runtime *util.RuntimeOptions) (_result *GetGlobalEventsStorageRegionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetGlobalEventsStorageRegion"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGlobalEventsStorageRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * By default, global events are stored in the Singapore region.
 * To obtain the permissions to call the API operation, you must submit a ticket.
 *
 * @return GetGlobalEventsStorageRegionResponse
 */
func (client *Client) GetGlobalEventsStorageRegion() (_result *GetGlobalEventsStorageRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGlobalEventsStorageRegionResponse{}
	_body, _err := client.GetGlobalEventsStorageRegionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic describes how to query the status of a sample single-account trail named `trail-test`.
 *
 * @param request GetTrailStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTrailStatusResponse
 */
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

/**
 * This topic describes how to query the status of a sample single-account trail named `trail-test`.
 *
 * @param request GetTrailStatusRequest
 * @return GetTrailStatusResponse
 */
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

/**
 * This topic describes how to query the historical event delivery tasks created within your Alibaba Cloud account. In this example, a historical event delivery task whose ID is `16602` is returned. This task is used to deliver historical events for the trail named `trail-name` to Log Service.
 *
 * @param request ListDeliveryHistoryJobsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDeliveryHistoryJobsResponse
 */
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

/**
 * This topic describes how to query the historical event delivery tasks created within your Alibaba Cloud account. In this example, a historical event delivery task whose ID is `16602` is returned. This task is used to deliver historical events for the trail named `trail-name` to Log Service.
 *
 * @param request ListDeliveryHistoryJobsRequest
 * @return ListDeliveryHistoryJobsResponse
 */
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

/**
 * When you call this operation to query event details, you can query the event details at most twice per second.
 * > Do not frequently call this operation. You can create a trail to deliver events to Log Service. Then, you can query event details in near real time by using the real-time log consumption feature of Log Service. For more information, see [Create a single-account trail](~~28810~~), [Create a multi-account trail](~~160661~~), and [Overview](~~28997~~).
 *
 * @param request LookupEventsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LookupEventsResponse
 */
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

/**
 * When you call this operation to query event details, you can query the event details at most twice per second.
 * > Do not frequently call this operation. You can create a trail to deliver events to Log Service. Then, you can query event details in near real time by using the real-time log consumption feature of Log Service. For more information, see [Create a single-account trail](~~28810~~), [Create a multi-account trail](~~160661~~), and [Overview](~~28997~~).
 *
 * @param request LookupEventsRequest
 * @return LookupEventsResponse
 */
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

/**
 * This topic describes how to enable logging for a sample trail named `trail-test`.
 *
 * @param request StartLoggingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartLoggingResponse
 */
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

/**
 * This topic describes how to enable logging for a sample trail named `trail-test`.
 *
 * @param request StartLoggingRequest
 * @return StartLoggingResponse
 */
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

/**
 * This topic describes how to disable logging for a sample trail named `trail-test`.
 *
 * @param request StopLoggingRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopLoggingResponse
 */
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

/**
 * This topic describes how to disable logging for a sample trail named `trail-test`.
 *
 * @param request StopLoggingRequest
 * @return StopLoggingResponse
 */
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

/**
 * By default, global events are stored in the Singapore region.
 * *   To obtain the permissions to call the API operation, you must submit a ticket.
 * *   Only the China (Hangzhou) region (cn-hangzhou) and the Singapore region (ap-southeast-1) are supported.
 *
 * @param request UpdateGlobalEventsStorageRegionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateGlobalEventsStorageRegionResponse
 */
func (client *Client) UpdateGlobalEventsStorageRegionWithOptions(request *UpdateGlobalEventsStorageRegionRequest, runtime *util.RuntimeOptions) (_result *UpdateGlobalEventsStorageRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StorageRegion)) {
		query["StorageRegion"] = request.StorageRegion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGlobalEventsStorageRegion"),
		Version:     tea.String("2020-07-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGlobalEventsStorageRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * By default, global events are stored in the Singapore region.
 * *   To obtain the permissions to call the API operation, you must submit a ticket.
 * *   Only the China (Hangzhou) region (cn-hangzhou) and the Singapore region (ap-southeast-1) are supported.
 *
 * @param request UpdateGlobalEventsStorageRegionRequest
 * @return UpdateGlobalEventsStorageRegionResponse
 */
func (client *Client) UpdateGlobalEventsStorageRegion(request *UpdateGlobalEventsStorageRegionRequest) (_result *UpdateGlobalEventsStorageRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGlobalEventsStorageRegionResponse{}
	_body, _err := client.UpdateGlobalEventsStorageRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic shows you how to change the destination Object Storage Service (OSS) bucket of a sample trail named `trail-test` to `audit-log`.
 *
 * @param request UpdateTrailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateTrailResponse
 */
func (client *Client) UpdateTrailWithOptions(request *UpdateTrailRequest, runtime *util.RuntimeOptions) (_result *UpdateTrailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventRW)) {
		query["EventRW"] = request.EventRW
	}

	if !tea.BoolValue(util.IsUnset(request.MaxComputeProjectArn)) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !tea.BoolValue(util.IsUnset(request.MaxComputeWriteRoleArn)) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
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

/**
 * This topic shows you how to change the destination Object Storage Service (OSS) bucket of a sample trail named `trail-test` to `audit-log`.
 *
 * @param request UpdateTrailRequest
 * @return UpdateTrailResponse
 */
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
