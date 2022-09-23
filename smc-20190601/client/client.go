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

type CreateAccessTokenRequest struct {
	Count                *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TimeToLiveInDays     *string `json:"TimeToLiveInDays,omitempty" xml:"TimeToLiveInDays,omitempty"`
}

func (s CreateAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenRequest) SetCount(v string) *CreateAccessTokenRequest {
	s.Count = &v
	return s
}

func (s *CreateAccessTokenRequest) SetDescription(v string) *CreateAccessTokenRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessTokenRequest) SetName(v string) *CreateAccessTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateAccessTokenRequest) SetOwnerId(v int64) *CreateAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccessTokenRequest) SetResourceOwnerAccount(v string) *CreateAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccessTokenRequest) SetTimeToLiveInDays(v string) *CreateAccessTokenRequest {
	s.TimeToLiveInDays = &v
	return s
}

type CreateAccessTokenResponseBody struct {
	AccessTokenCode *string `json:"AccessTokenCode,omitempty" xml:"AccessTokenCode,omitempty"`
	AccessTokenId   *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenResponseBody) SetAccessTokenCode(v string) *CreateAccessTokenResponseBody {
	s.AccessTokenCode = &v
	return s
}

func (s *CreateAccessTokenResponseBody) SetAccessTokenId(v string) *CreateAccessTokenResponseBody {
	s.AccessTokenId = &v
	return s
}

func (s *CreateAccessTokenResponseBody) SetRequestId(v string) *CreateAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenResponse) SetHeaders(v map[string]*string) *CreateAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessTokenResponse) SetStatusCode(v int32) *CreateAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessTokenResponse) SetBody(v *CreateAccessTokenResponseBody) *CreateAccessTokenResponse {
	s.Body = v
	return s
}

type CreateReplicationJobRequest struct {
	ClientToken            *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContainerNamespace     *string                                      `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	ContainerRepository    *string                                      `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	ContainerTag           *string                                      `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	DataDisk               []*CreateReplicationJobRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	Description            *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Frequency              *int32                                       `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	ImageName              *string                                      `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId             *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRamRole        *string                                      `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	InstanceType           *string                                      `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	JobType                *int32                                       `json:"JobType,omitempty" xml:"JobType,omitempty"`
	LaunchTemplateId       *string                                      `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string                                      `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LicenseType            *string                                      `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	MaxNumberOfImageToKeep *int32                                       `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	Name                   *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NetMode                *int32                                       `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	OwnerId                *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId               *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicationParameters  *string                                      `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	ResourceOwnerAccount   *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RunOnce                *bool                                        `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	ScheduledStartTime     *string                                      `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	SourceId               *string                                      `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SystemDiskPart         []*CreateReplicationJobRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	SystemDiskSize         *int32                                       `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Tag                    []*CreateReplicationJobRequestTag            `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TargetType             *string                                      `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	VSwitchId              *string                                      `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ValidTime              *string                                      `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	VpcId                  *string                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequest) SetClientToken(v string) *CreateReplicationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateReplicationJobRequest) SetContainerNamespace(v string) *CreateReplicationJobRequest {
	s.ContainerNamespace = &v
	return s
}

func (s *CreateReplicationJobRequest) SetContainerRepository(v string) *CreateReplicationJobRequest {
	s.ContainerRepository = &v
	return s
}

func (s *CreateReplicationJobRequest) SetContainerTag(v string) *CreateReplicationJobRequest {
	s.ContainerTag = &v
	return s
}

func (s *CreateReplicationJobRequest) SetDataDisk(v []*CreateReplicationJobRequestDataDisk) *CreateReplicationJobRequest {
	s.DataDisk = v
	return s
}

func (s *CreateReplicationJobRequest) SetDescription(v string) *CreateReplicationJobRequest {
	s.Description = &v
	return s
}

func (s *CreateReplicationJobRequest) SetFrequency(v int32) *CreateReplicationJobRequest {
	s.Frequency = &v
	return s
}

func (s *CreateReplicationJobRequest) SetImageName(v string) *CreateReplicationJobRequest {
	s.ImageName = &v
	return s
}

func (s *CreateReplicationJobRequest) SetInstanceId(v string) *CreateReplicationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetInstanceRamRole(v string) *CreateReplicationJobRequest {
	s.InstanceRamRole = &v
	return s
}

func (s *CreateReplicationJobRequest) SetInstanceType(v string) *CreateReplicationJobRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetJobType(v int32) *CreateReplicationJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetLaunchTemplateId(v string) *CreateReplicationJobRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetLaunchTemplateVersion(v string) *CreateReplicationJobRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateReplicationJobRequest) SetLicenseType(v string) *CreateReplicationJobRequest {
	s.LicenseType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetMaxNumberOfImageToKeep(v int32) *CreateReplicationJobRequest {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *CreateReplicationJobRequest) SetName(v string) *CreateReplicationJobRequest {
	s.Name = &v
	return s
}

func (s *CreateReplicationJobRequest) SetNetMode(v int32) *CreateReplicationJobRequest {
	s.NetMode = &v
	return s
}

func (s *CreateReplicationJobRequest) SetOwnerId(v int64) *CreateReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetRegionId(v string) *CreateReplicationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetReplicationParameters(v string) *CreateReplicationJobRequest {
	s.ReplicationParameters = &v
	return s
}

func (s *CreateReplicationJobRequest) SetResourceOwnerAccount(v string) *CreateReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateReplicationJobRequest) SetRunOnce(v bool) *CreateReplicationJobRequest {
	s.RunOnce = &v
	return s
}

func (s *CreateReplicationJobRequest) SetScheduledStartTime(v string) *CreateReplicationJobRequest {
	s.ScheduledStartTime = &v
	return s
}

func (s *CreateReplicationJobRequest) SetSourceId(v string) *CreateReplicationJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetSystemDiskPart(v []*CreateReplicationJobRequestSystemDiskPart) *CreateReplicationJobRequest {
	s.SystemDiskPart = v
	return s
}

func (s *CreateReplicationJobRequest) SetSystemDiskSize(v int32) *CreateReplicationJobRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateReplicationJobRequest) SetTag(v []*CreateReplicationJobRequestTag) *CreateReplicationJobRequest {
	s.Tag = v
	return s
}

func (s *CreateReplicationJobRequest) SetTargetType(v string) *CreateReplicationJobRequest {
	s.TargetType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetVSwitchId(v string) *CreateReplicationJobRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetValidTime(v string) *CreateReplicationJobRequest {
	s.ValidTime = &v
	return s
}

func (s *CreateReplicationJobRequest) SetVpcId(v string) *CreateReplicationJobRequest {
	s.VpcId = &v
	return s
}

type CreateReplicationJobRequestDataDisk struct {
	Index *int32                                     `json:"Index,omitempty" xml:"Index,omitempty"`
	Part  []*CreateReplicationJobRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	Size  *int32                                     `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateReplicationJobRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDataDisk) SetIndex(v int32) *CreateReplicationJobRequestDataDisk {
	s.Index = &v
	return s
}

func (s *CreateReplicationJobRequestDataDisk) SetPart(v []*CreateReplicationJobRequestDataDiskPart) *CreateReplicationJobRequestDataDisk {
	s.Part = v
	return s
}

func (s *CreateReplicationJobRequestDataDisk) SetSize(v int32) *CreateReplicationJobRequestDataDisk {
	s.Size = &v
	return s
}

type CreateReplicationJobRequestDataDiskPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestDataDiskPart) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDataDiskPart) SetBlock(v bool) *CreateReplicationJobRequestDataDiskPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestDataDiskPart) SetDevice(v string) *CreateReplicationJobRequestDataDiskPart {
	s.Device = &v
	return s
}

func (s *CreateReplicationJobRequestDataDiskPart) SetSizeBytes(v int64) *CreateReplicationJobRequestDataDiskPart {
	s.SizeBytes = &v
	return s
}

type CreateReplicationJobRequestSystemDiskPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetBlock(v bool) *CreateReplicationJobRequestSystemDiskPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetDevice(v string) *CreateReplicationJobRequestSystemDiskPart {
	s.Device = &v
	return s
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetSizeBytes(v int64) *CreateReplicationJobRequestSystemDiskPart {
	s.SizeBytes = &v
	return s
}

type CreateReplicationJobRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateReplicationJobRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestTag) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestTag) SetKey(v string) *CreateReplicationJobRequestTag {
	s.Key = &v
	return s
}

func (s *CreateReplicationJobRequestTag) SetValue(v string) *CreateReplicationJobRequestTag {
	s.Value = &v
	return s
}

type CreateReplicationJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobResponseBody) SetJobId(v string) *CreateReplicationJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateReplicationJobResponseBody) SetRequestId(v string) *CreateReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateReplicationJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateReplicationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobResponse) SetHeaders(v map[string]*string) *CreateReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateReplicationJobResponse) SetStatusCode(v int32) *CreateReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReplicationJobResponse) SetBody(v *CreateReplicationJobResponseBody) *CreateReplicationJobResponse {
	s.Body = v
	return s
}

type CutOverReplicationJobRequest struct {
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SyncData             *bool   `json:"SyncData,omitempty" xml:"SyncData,omitempty"`
}

func (s CutOverReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CutOverReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobRequest) SetJobId(v string) *CutOverReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetOwnerId(v int64) *CutOverReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetResourceOwnerAccount(v string) *CutOverReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetSyncData(v bool) *CutOverReplicationJobRequest {
	s.SyncData = &v
	return s
}

type CutOverReplicationJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CutOverReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CutOverReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobResponseBody) SetRequestId(v string) *CutOverReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

type CutOverReplicationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CutOverReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CutOverReplicationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CutOverReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobResponse) SetHeaders(v map[string]*string) *CutOverReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *CutOverReplicationJobResponse) SetStatusCode(v int32) *CutOverReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CutOverReplicationJobResponse) SetBody(v *CutOverReplicationJobResponseBody) *CutOverReplicationJobResponse {
	s.Body = v
	return s
}

type DeleteAccessTokenRequest struct {
	AccessTokenId        *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenRequest) SetAccessTokenId(v string) *DeleteAccessTokenRequest {
	s.AccessTokenId = &v
	return s
}

func (s *DeleteAccessTokenRequest) SetOwnerId(v int64) *DeleteAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccessTokenRequest) SetResourceOwnerAccount(v string) *DeleteAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DeleteAccessTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponseBody) SetRequestId(v string) *DeleteAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessTokenResponse) SetHeaders(v map[string]*string) *DeleteAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessTokenResponse) SetStatusCode(v int32) *DeleteAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessTokenResponse) SetBody(v *DeleteAccessTokenResponseBody) *DeleteAccessTokenResponse {
	s.Body = v
	return s
}

type DeleteReplicationJobRequest struct {
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeleteReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobRequest) SetJobId(v string) *DeleteReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetOwnerId(v int64) *DeleteReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetResourceOwnerAccount(v string) *DeleteReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DeleteReplicationJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobResponseBody) SetRequestId(v string) *DeleteReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteReplicationJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteReplicationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobResponse) SetHeaders(v map[string]*string) *DeleteReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteReplicationJobResponse) SetStatusCode(v int32) *DeleteReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReplicationJobResponse) SetBody(v *DeleteReplicationJobResponseBody) *DeleteReplicationJobResponse {
	s.Body = v
	return s
}

type DeleteSourceServerRequest struct {
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DeleteSourceServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerRequest) SetForce(v bool) *DeleteSourceServerRequest {
	s.Force = &v
	return s
}

func (s *DeleteSourceServerRequest) SetOwnerId(v int64) *DeleteSourceServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSourceServerRequest) SetResourceOwnerAccount(v string) *DeleteSourceServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSourceServerRequest) SetSourceId(v string) *DeleteSourceServerRequest {
	s.SourceId = &v
	return s
}

type DeleteSourceServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSourceServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerResponseBody) SetRequestId(v string) *DeleteSourceServerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSourceServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSourceServerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerResponse) SetHeaders(v map[string]*string) *DeleteSourceServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceServerResponse) SetStatusCode(v int32) *DeleteSourceServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceServerResponse) SetBody(v *DeleteSourceServerResponseBody) *DeleteSourceServerResponse {
	s.Body = v
	return s
}

type DescribeReplicationJobsRequest struct {
	BusinessStatus       *string   `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	JobId                []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	JobType              *int32    `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeReplicationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsRequest) SetBusinessStatus(v string) *DescribeReplicationJobsRequest {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetInstanceId(v []*string) *DescribeReplicationJobsRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetJobId(v []*string) *DescribeReplicationJobsRequest {
	s.JobId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetJobType(v int32) *DescribeReplicationJobsRequest {
	s.JobType = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetName(v string) *DescribeReplicationJobsRequest {
	s.Name = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetOwnerId(v int64) *DescribeReplicationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetPageNumber(v int32) *DescribeReplicationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetPageSize(v int32) *DescribeReplicationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetRegionId(v string) *DescribeReplicationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetResourceOwnerAccount(v string) *DescribeReplicationJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetSourceId(v []*string) *DescribeReplicationJobsRequest {
	s.SourceId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetStatus(v string) *DescribeReplicationJobsRequest {
	s.Status = &v
	return s
}

type DescribeReplicationJobsResponseBody struct {
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReplicationJobs *DescribeReplicationJobsResponseBodyReplicationJobs `json:"ReplicationJobs,omitempty" xml:"ReplicationJobs,omitempty" type:"Struct"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeReplicationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBody) SetPageNumber(v int32) *DescribeReplicationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetPageSize(v int32) *DescribeReplicationJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetReplicationJobs(v *DescribeReplicationJobsResponseBodyReplicationJobs) *DescribeReplicationJobsResponseBody {
	s.ReplicationJobs = v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetRequestId(v string) *DescribeReplicationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetTotalCount(v int32) *DescribeReplicationJobsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobs struct {
	ReplicationJob []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob `json:"ReplicationJob,omitempty" xml:"ReplicationJob,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobs) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobs) SetReplicationJob(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) *DescribeReplicationJobsResponseBodyReplicationJobs {
	s.ReplicationJob = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob struct {
	BusinessStatus         *string                                                                             `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ContainerNamespace     *string                                                                             `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	ContainerRepository    *string                                                                             `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	ContainerTag           *string                                                                             `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	CreationTime           *string                                                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDisks              *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks          `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	Description            *string                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                *string                                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorCode              *string                                                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Frequency              *int32                                                                              `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	ImageId                *string                                                                             `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName              *string                                                                             `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId             *string                                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRamRole        *string                                                                             `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	InstanceType           *string                                                                             `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	JobId                  *string                                                                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType                *int32                                                                              `json:"JobType,omitempty" xml:"JobType,omitempty"`
	LaunchTemplateId       *string                                                                             `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string                                                                             `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	LicenseType            *string                                                                             `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	MaxNumberOfImageToKeep *int32                                                                              `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	Name                   *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	NetMode                *int32                                                                              `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	Progress               *float32                                                                            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionId               *string                                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicationJobRuns     *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns `json:"ReplicationJobRuns,omitempty" xml:"ReplicationJobRuns,omitempty" type:"Struct"`
	ReplicationParameters  *string                                                                             `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	RunOnce                *bool                                                                               `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	ScheduledStartTime     *string                                                                             `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	SourceId               *string                                                                             `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	StartTime              *string                                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                 *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusInfo             *string                                                                             `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	SystemDiskParts        *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts    `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	SystemDiskSize         *int32                                                                              `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	TargetType             *string                                                                             `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TransitionInstanceId   *string                                                                             `json:"TransitionInstanceId,omitempty" xml:"TransitionInstanceId,omitempty"`
	VSwitchId              *string                                                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ValidTime              *string                                                                             `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	VpcId                  *string                                                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetBusinessStatus(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerNamespace(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerNamespace = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerRepository(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerRepository = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerTag(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerTag = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetCreationTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.CreationTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetDataDisks(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.DataDisks = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetDescription(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Description = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetEndTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.EndTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetErrorCode(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ErrorCode = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetFrequency(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Frequency = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetImageId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ImageId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetImageName(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ImageName = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceRamRole(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceRamRole = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetJobId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.JobId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetJobType(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.JobType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLaunchTemplateId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLaunchTemplateVersion(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLicenseType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LicenseType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetMaxNumberOfImageToKeep(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetName(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Name = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetNetMode(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.NetMode = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetProgress(v float32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Progress = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetRegionId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetReplicationJobRuns(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ReplicationJobRuns = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetReplicationParameters(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ReplicationParameters = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetRunOnce(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.RunOnce = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetScheduledStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ScheduledStartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSourceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SourceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.StartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStatus(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Status = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStatusInfo(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.StatusInfo = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSystemDiskParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SystemDiskParts = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSystemDiskSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetTargetType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.TargetType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetTransitionInstanceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.TransitionInstanceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetVSwitchId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.VSwitchId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetValidTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ValidTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetVpcId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.VpcId = &v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks struct {
	DataDisk []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) SetDataDisk(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks {
	s.DataDisk = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk struct {
	Index *int32                                                                                  `json:"Index,omitempty" xml:"Index,omitempty"`
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	Size  *int32                                                                                  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) SetIndex(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	s.Index = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) SetParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	s.Parts = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) SetSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	s.Size = &v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts struct {
	Part []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) SetPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts {
	s.Part = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetDevice(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.Device = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.SizeBytes = &v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns struct {
	ReplicationJobRun []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun `json:"ReplicationJobRun,omitempty" xml:"ReplicationJobRun,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) SetReplicationJobRun(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns {
	s.ReplicationJobRun = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetEndTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.EndTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetImageId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.ImageId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.StartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.Type = &v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts struct {
	SystemDiskPart []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) SetSystemDiskPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts {
	s.SystemDiskPart = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetDevice(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.Device = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.SizeBytes = &v
	return s
}

type DescribeReplicationJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeReplicationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeReplicationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponse) SetHeaders(v map[string]*string) *DescribeReplicationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeReplicationJobsResponse) SetStatusCode(v int32) *DescribeReplicationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReplicationJobsResponse) SetBody(v *DescribeReplicationJobsResponseBody) *DescribeReplicationJobsResponse {
	s.Body = v
	return s
}

type DescribeSourceServersRequest struct {
	JobId                *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	State                *string   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSourceServersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersRequest) SetJobId(v string) *DescribeSourceServersRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetName(v string) *DescribeSourceServersRequest {
	s.Name = &v
	return s
}

func (s *DescribeSourceServersRequest) SetOwnerId(v int64) *DescribeSourceServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetPageNumber(v int32) *DescribeSourceServersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSourceServersRequest) SetPageSize(v int32) *DescribeSourceServersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSourceServersRequest) SetResourceOwnerAccount(v string) *DescribeSourceServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSourceServersRequest) SetSourceId(v []*string) *DescribeSourceServersRequest {
	s.SourceId = v
	return s
}

func (s *DescribeSourceServersRequest) SetState(v string) *DescribeSourceServersRequest {
	s.State = &v
	return s
}

type DescribeSourceServersResponseBody struct {
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceServers *DescribeSourceServersResponseBodySourceServers `json:"SourceServers,omitempty" xml:"SourceServers,omitempty" type:"Struct"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSourceServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBody) SetPageNumber(v int32) *DescribeSourceServersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetPageSize(v int32) *DescribeSourceServersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetRequestId(v string) *DescribeSourceServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSourceServersResponseBody) SetSourceServers(v *DescribeSourceServersResponseBodySourceServers) *DescribeSourceServersResponseBody {
	s.SourceServers = v
	return s
}

func (s *DescribeSourceServersResponseBody) SetTotalCount(v int32) *DescribeSourceServersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSourceServersResponseBodySourceServers struct {
	SourceServer []*DescribeSourceServersResponseBodySourceServersSourceServer `json:"SourceServer,omitempty" xml:"SourceServer,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServers) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServers) SetSourceServer(v []*DescribeSourceServersResponseBodySourceServersSourceServer) *DescribeSourceServersResponseBodySourceServers {
	s.SourceServer = v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServer struct {
	AgentVersion      *string                                                                    `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	Architecture      *string                                                                    `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	CreationTime      *string                                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDisks         *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks       `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	Description       *string                                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCode         *string                                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HeartbeatRate     *int32                                                                     `json:"HeartbeatRate,omitempty" xml:"HeartbeatRate,omitempty"`
	JobId             *string                                                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	KernelLevel       *int32                                                                     `json:"KernelLevel,omitempty" xml:"KernelLevel,omitempty"`
	Name              *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Platform          *string                                                                    `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ReplicationDriver *string                                                                    `json:"ReplicationDriver,omitempty" xml:"ReplicationDriver,omitempty"`
	SourceId          *string                                                                    `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	State             *string                                                                    `json:"State,omitempty" xml:"State,omitempty"`
	StatusInfo        *string                                                                    `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	SystemDiskParts   *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	SystemDiskSize    *int32                                                                     `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemInfo        *string                                                                    `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetAgentVersion(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.AgentVersion = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetArchitecture(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Architecture = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetCreationTime(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.CreationTime = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDataDisks(v *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.DataDisks = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDescription(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Description = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetErrorCode(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetHeartbeatRate(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.HeartbeatRate = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetJobId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.JobId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetKernelLevel(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.KernelLevel = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetName(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Name = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetPlatform(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Platform = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetReplicationDriver(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ReplicationDriver = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSourceId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SourceId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetState(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.State = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetStatusInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.StatusInfo = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskParts(v *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskParts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemInfo = &v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisks struct {
	DataDisk []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) SetDataDisk(v []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks {
	s.DataDisk = v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk struct {
	Index *int32                                                                            `json:"Index,omitempty" xml:"Index,omitempty"`
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	Path  *string                                                                           `json:"Path,omitempty" xml:"Path,omitempty"`
	Size  *int32                                                                            `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetIndex(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Index = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetParts(v *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Parts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Size = &v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts struct {
	Part []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts) SetPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts {
	s.Part = v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart struct {
	CanBlock  *bool   `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Need      *bool   `json:"Need,omitempty" xml:"Need,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.SizeBytes = &v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts struct {
	SystemDiskPart []*DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) SetSystemDiskPart(v []*DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts {
	s.SystemDiskPart = v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart struct {
	CanBlock  *bool   `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Need      *bool   `json:"Need,omitempty" xml:"Need,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetCanBlock(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.CanBlock = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Path = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.SizeBytes = &v
	return s
}

type DescribeSourceServersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSourceServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSourceServersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponse) SetHeaders(v map[string]*string) *DescribeSourceServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSourceServersResponse) SetStatusCode(v int32) *DescribeSourceServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSourceServersResponse) SetBody(v *DescribeSourceServersResponseBody) *DescribeSourceServersResponse {
	s.Body = v
	return s
}

type DisableAccessTokenRequest struct {
	AccessTokenId        *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DisableAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenRequest) SetAccessTokenId(v string) *DisableAccessTokenRequest {
	s.AccessTokenId = &v
	return s
}

func (s *DisableAccessTokenRequest) SetOwnerId(v int64) *DisableAccessTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAccessTokenRequest) SetResourceOwnerAccount(v string) *DisableAccessTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DisableAccessTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAccessTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponseBody) SetRequestId(v string) *DisableAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

type DisableAccessTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponse) SetHeaders(v map[string]*string) *DisableAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableAccessTokenResponse) SetStatusCode(v int32) *DisableAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAccessTokenResponse) SetBody(v *DisableAccessTokenResponseBody) *DisableAccessTokenResponse {
	s.Body = v
	return s
}

type ListAccessTokensRequest struct {
	AccessTokenId        []*string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty" type:"Repeated"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAccessTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensRequest) GoString() string {
	return s.String()
}

func (s *ListAccessTokensRequest) SetAccessTokenId(v []*string) *ListAccessTokensRequest {
	s.AccessTokenId = v
	return s
}

func (s *ListAccessTokensRequest) SetName(v string) *ListAccessTokensRequest {
	s.Name = &v
	return s
}

func (s *ListAccessTokensRequest) SetOwnerId(v int64) *ListAccessTokensRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAccessTokensRequest) SetResourceOwnerAccount(v string) *ListAccessTokensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAccessTokensRequest) SetStatus(v string) *ListAccessTokensRequest {
	s.Status = &v
	return s
}

type ListAccessTokensResponseBody struct {
	AccessTokens *ListAccessTokensResponseBodyAccessTokens `json:"AccessTokens,omitempty" xml:"AccessTokens,omitempty" type:"Struct"`
	PageNumber   *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessTokensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBody) SetAccessTokens(v *ListAccessTokensResponseBodyAccessTokens) *ListAccessTokensResponseBody {
	s.AccessTokens = v
	return s
}

func (s *ListAccessTokensResponseBody) SetPageNumber(v int32) *ListAccessTokensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetPageSize(v int32) *ListAccessTokensResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetRequestId(v string) *ListAccessTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetTotalCount(v int32) *ListAccessTokensResponseBody {
	s.TotalCount = &v
	return s
}

type ListAccessTokensResponseBodyAccessTokens struct {
	AccessToken []*ListAccessTokensResponseBodyAccessTokensAccessToken `json:"AccessToken,omitempty" xml:"AccessToken,omitempty" type:"Repeated"`
}

func (s ListAccessTokensResponseBodyAccessTokens) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponseBodyAccessTokens) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBodyAccessTokens) SetAccessToken(v []*ListAccessTokensResponseBodyAccessTokensAccessToken) *ListAccessTokensResponseBodyAccessTokens {
	s.AccessToken = v
	return s
}

type ListAccessTokensResponseBodyAccessTokensAccessToken struct {
	AccessTokenId    *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	Count            *string `json:"Count,omitempty" xml:"Count,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegisteredCount  *string `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeToLiveInDays *string `json:"TimeToLiveInDays,omitempty" xml:"TimeToLiveInDays,omitempty"`
}

func (s ListAccessTokensResponseBodyAccessTokensAccessToken) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponseBodyAccessTokensAccessToken) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetAccessTokenId(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.AccessTokenId = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetCount(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Count = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetCreationTime(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.CreationTime = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetDescription(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Description = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetName(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Name = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetRegisteredCount(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.RegisteredCount = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetStatus(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Status = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetTimeToLiveInDays(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.TimeToLiveInDays = &v
	return s
}

type ListAccessTokensResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAccessTokensResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccessTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessTokensResponse) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponse) SetHeaders(v map[string]*string) *ListAccessTokensResponse {
	s.Headers = v
	return s
}

func (s *ListAccessTokensResponse) SetStatusCode(v int32) *ListAccessTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessTokensResponse) SetBody(v *ListAccessTokensResponseBody) *ListAccessTokensResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
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

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyReplicationJobAttributeRequest struct {
	ContainerNamespace     *string                                               `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	ContainerRepository    *string                                               `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	ContainerTag           *string                                               `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	DataDisk               []*ModifyReplicationJobAttributeRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	Description            *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Frequency              *int32                                                `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	ImageName              *string                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId             *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRamRole        *string                                               `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	InstanceType           *string                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	JobId                  *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	LaunchTemplateId       *string                                               `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string                                               `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	MaxNumberOfImageToKeep *int32                                                `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	Name                   *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId                *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScheduledStartTime     *string                                               `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	SystemDiskPart         []*ModifyReplicationJobAttributeRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	SystemDiskSize         *int32                                                `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	TargetType             *string                                               `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	ValidTime              *string                                               `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s ModifyReplicationJobAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequest) SetContainerNamespace(v string) *ModifyReplicationJobAttributeRequest {
	s.ContainerNamespace = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetContainerRepository(v string) *ModifyReplicationJobAttributeRequest {
	s.ContainerRepository = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetContainerTag(v string) *ModifyReplicationJobAttributeRequest {
	s.ContainerTag = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetDataDisk(v []*ModifyReplicationJobAttributeRequestDataDisk) *ModifyReplicationJobAttributeRequest {
	s.DataDisk = v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetDescription(v string) *ModifyReplicationJobAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetFrequency(v int32) *ModifyReplicationJobAttributeRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetImageName(v string) *ModifyReplicationJobAttributeRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetInstanceId(v string) *ModifyReplicationJobAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetInstanceRamRole(v string) *ModifyReplicationJobAttributeRequest {
	s.InstanceRamRole = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetInstanceType(v string) *ModifyReplicationJobAttributeRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetJobId(v string) *ModifyReplicationJobAttributeRequest {
	s.JobId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetLaunchTemplateId(v string) *ModifyReplicationJobAttributeRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetLaunchTemplateVersion(v string) *ModifyReplicationJobAttributeRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetMaxNumberOfImageToKeep(v int32) *ModifyReplicationJobAttributeRequest {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetName(v string) *ModifyReplicationJobAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetOwnerId(v int64) *ModifyReplicationJobAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetResourceOwnerAccount(v string) *ModifyReplicationJobAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetScheduledStartTime(v string) *ModifyReplicationJobAttributeRequest {
	s.ScheduledStartTime = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetSystemDiskPart(v []*ModifyReplicationJobAttributeRequestSystemDiskPart) *ModifyReplicationJobAttributeRequest {
	s.SystemDiskPart = v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetSystemDiskSize(v int32) *ModifyReplicationJobAttributeRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetTargetType(v string) *ModifyReplicationJobAttributeRequest {
	s.TargetType = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetValidTime(v string) *ModifyReplicationJobAttributeRequest {
	s.ValidTime = &v
	return s
}

type ModifyReplicationJobAttributeRequestDataDisk struct {
	Index *int32                                              `json:"Index,omitempty" xml:"Index,omitempty"`
	Part  []*ModifyReplicationJobAttributeRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	Size  *int32                                              `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestDataDisk) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) SetIndex(v int32) *ModifyReplicationJobAttributeRequestDataDisk {
	s.Index = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) SetPart(v []*ModifyReplicationJobAttributeRequestDataDiskPart) *ModifyReplicationJobAttributeRequestDataDisk {
	s.Part = v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) SetSize(v int32) *ModifyReplicationJobAttributeRequestDataDisk {
	s.Size = &v
	return s
}

type ModifyReplicationJobAttributeRequestDataDiskPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetBlock(v bool) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.Block = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetDevice(v string) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.Device = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetSizeBytes(v int64) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.SizeBytes = &v
	return s
}

type ModifyReplicationJobAttributeRequestSystemDiskPart struct {
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetBlock(v bool) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.Block = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetDevice(v string) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.Device = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetSizeBytes(v int64) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.SizeBytes = &v
	return s
}

type ModifyReplicationJobAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReplicationJobAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeResponseBody) SetRequestId(v string) *ModifyReplicationJobAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyReplicationJobAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyReplicationJobAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyReplicationJobAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeResponse) SetHeaders(v map[string]*string) *ModifyReplicationJobAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyReplicationJobAttributeResponse) SetStatusCode(v int32) *ModifyReplicationJobAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReplicationJobAttributeResponse) SetBody(v *ModifyReplicationJobAttributeResponseBody) *ModifyReplicationJobAttributeResponse {
	s.Body = v
	return s
}

type ModifySourceServerAttributeRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s ModifySourceServerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySourceServerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeRequest) SetDescription(v string) *ModifySourceServerAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetName(v string) *ModifySourceServerAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetOwnerId(v int64) *ModifySourceServerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetResourceOwnerAccount(v string) *ModifySourceServerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetSourceId(v string) *ModifySourceServerAttributeRequest {
	s.SourceId = &v
	return s
}

type ModifySourceServerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySourceServerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySourceServerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeResponseBody) SetRequestId(v string) *ModifySourceServerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifySourceServerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySourceServerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySourceServerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySourceServerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySourceServerAttributeResponse) SetHeaders(v map[string]*string) *ModifySourceServerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySourceServerAttributeResponse) SetStatusCode(v int32) *ModifySourceServerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySourceServerAttributeResponse) SetBody(v *ModifySourceServerAttributeResponseBody) *ModifySourceServerAttributeResponse {
	s.Body = v
	return s
}

type StartReplicationJobRequest struct {
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s StartReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StartReplicationJobRequest) SetJobId(v string) *StartReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *StartReplicationJobRequest) SetOwnerId(v int64) *StartReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartReplicationJobRequest) SetResourceOwnerAccount(v string) *StartReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type StartReplicationJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartReplicationJobResponseBody) SetRequestId(v string) *StartReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

type StartReplicationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartReplicationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *StartReplicationJobResponse) SetHeaders(v map[string]*string) *StartReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *StartReplicationJobResponse) SetStatusCode(v int32) *StartReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartReplicationJobResponse) SetBody(v *StartReplicationJobResponseBody) *StartReplicationJobResponse {
	s.Body = v
	return s
}

type StopReplicationJobRequest struct {
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s StopReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StopReplicationJobRequest) SetJobId(v string) *StopReplicationJobRequest {
	s.JobId = &v
	return s
}

func (s *StopReplicationJobRequest) SetOwnerId(v int64) *StopReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopReplicationJobRequest) SetResourceOwnerAccount(v string) *StopReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type StopReplicationJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopReplicationJobResponseBody) SetRequestId(v string) *StopReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

type StopReplicationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopReplicationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopReplicationJobResponse) GoString() string {
	return s.String()
}

func (s *StopReplicationJobResponse) SetHeaders(v map[string]*string) *StopReplicationJobResponse {
	s.Headers = v
	return s
}

func (s *StopReplicationJobResponse) SetStatusCode(v int32) *StopReplicationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopReplicationJobResponse) SetBody(v *StopReplicationJobResponseBody) *StopReplicationJobResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("smc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAccessTokenWithOptions(request *CreateAccessTokenRequest, runtime *util.RuntimeOptions) (_result *CreateAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TimeToLiveInDays)) {
		query["TimeToLiveInDays"] = request.TimeToLiveInDays
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessToken"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessToken(request *CreateAccessTokenRequest) (_result *CreateAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessTokenResponse{}
	_body, _err := client.CreateAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReplicationJobWithOptions(request *CreateReplicationJobRequest, runtime *util.RuntimeOptions) (_result *CreateReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerNamespace)) {
		query["ContainerNamespace"] = request.ContainerNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerRepository)) {
		query["ContainerRepository"] = request.ContainerRepository
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerTag)) {
		query["ContainerTag"] = request.ContainerTag
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		query["DataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		query["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRamRole)) {
		query["InstanceRamRole"] = request.InstanceRamRole
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateId)) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateVersion)) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseType)) {
		query["LicenseType"] = request.LicenseType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxNumberOfImageToKeep)) {
		query["MaxNumberOfImageToKeep"] = request.MaxNumberOfImageToKeep
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NetMode)) {
		query["NetMode"] = request.NetMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicationParameters)) {
		query["ReplicationParameters"] = request.ReplicationParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.RunOnce)) {
		query["RunOnce"] = request.RunOnce
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledStartTime)) {
		query["ScheduledStartTime"] = request.ScheduledStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskPart)) {
		query["SystemDiskPart"] = request.SystemDiskPart
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskSize)) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ValidTime)) {
		query["ValidTime"] = request.ValidTime
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReplicationJob"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReplicationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReplicationJob(request *CreateReplicationJobRequest) (_result *CreateReplicationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateReplicationJobResponse{}
	_body, _err := client.CreateReplicationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CutOverReplicationJobWithOptions(request *CutOverReplicationJobRequest, runtime *util.RuntimeOptions) (_result *CutOverReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SyncData)) {
		query["SyncData"] = request.SyncData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CutOverReplicationJob"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CutOverReplicationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CutOverReplicationJob(request *CutOverReplicationJobRequest) (_result *CutOverReplicationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CutOverReplicationJobResponse{}
	_body, _err := client.CutOverReplicationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessTokenWithOptions(request *DeleteAccessTokenRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenId)) {
		query["AccessTokenId"] = request.AccessTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessToken"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessToken(request *DeleteAccessTokenRequest) (_result *DeleteAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessTokenResponse{}
	_body, _err := client.DeleteAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteReplicationJobWithOptions(request *DeleteReplicationJobRequest, runtime *util.RuntimeOptions) (_result *DeleteReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReplicationJob"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReplicationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteReplicationJob(request *DeleteReplicationJobRequest) (_result *DeleteReplicationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteReplicationJobResponse{}
	_body, _err := client.DeleteReplicationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSourceServerWithOptions(request *DeleteSourceServerRequest, runtime *util.RuntimeOptions) (_result *DeleteSourceServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSourceServer"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSourceServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSourceServer(request *DeleteSourceServerRequest) (_result *DeleteSourceServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSourceServerResponse{}
	_body, _err := client.DeleteSourceServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeReplicationJobsWithOptions(request *DescribeReplicationJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeReplicationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessStatus)) {
		query["BusinessStatus"] = request.BusinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeReplicationJobs"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeReplicationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeReplicationJobs(request *DescribeReplicationJobsRequest) (_result *DescribeReplicationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeReplicationJobsResponse{}
	_body, _err := client.DescribeReplicationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSourceServersWithOptions(request *DescribeSourceServersRequest, runtime *util.RuntimeOptions) (_result *DescribeSourceServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSourceServers"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSourceServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSourceServers(request *DescribeSourceServersRequest) (_result *DescribeSourceServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSourceServersResponse{}
	_body, _err := client.DescribeSourceServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAccessTokenWithOptions(request *DisableAccessTokenRequest, runtime *util.RuntimeOptions) (_result *DisableAccessTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenId)) {
		query["AccessTokenId"] = request.AccessTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAccessToken"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAccessToken(request *DisableAccessTokenRequest) (_result *DisableAccessTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAccessTokenResponse{}
	_body, _err := client.DisableAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessTokensWithOptions(request *ListAccessTokensRequest, runtime *util.RuntimeOptions) (_result *ListAccessTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenId)) {
		query["AccessTokenId"] = request.AccessTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessTokens"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessTokens(request *ListAccessTokensRequest) (_result *ListAccessTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessTokensResponse{}
	_body, _err := client.ListAccessTokensWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyReplicationJobAttributeWithOptions(request *ModifyReplicationJobAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyReplicationJobAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerNamespace)) {
		query["ContainerNamespace"] = request.ContainerNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerRepository)) {
		query["ContainerRepository"] = request.ContainerRepository
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerTag)) {
		query["ContainerTag"] = request.ContainerTag
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		query["DataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		query["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRamRole)) {
		query["InstanceRamRole"] = request.InstanceRamRole
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateId)) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTemplateVersion)) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.MaxNumberOfImageToKeep)) {
		query["MaxNumberOfImageToKeep"] = request.MaxNumberOfImageToKeep
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledStartTime)) {
		query["ScheduledStartTime"] = request.ScheduledStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskPart)) {
		query["SystemDiskPart"] = request.SystemDiskPart
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDiskSize)) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.ValidTime)) {
		query["ValidTime"] = request.ValidTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyReplicationJobAttribute"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyReplicationJobAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyReplicationJobAttribute(request *ModifyReplicationJobAttributeRequest) (_result *ModifyReplicationJobAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyReplicationJobAttributeResponse{}
	_body, _err := client.ModifyReplicationJobAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySourceServerAttributeWithOptions(request *ModifySourceServerAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifySourceServerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySourceServerAttribute"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySourceServerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySourceServerAttribute(request *ModifySourceServerAttributeRequest) (_result *ModifySourceServerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySourceServerAttributeResponse{}
	_body, _err := client.ModifySourceServerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartReplicationJobWithOptions(request *StartReplicationJobRequest, runtime *util.RuntimeOptions) (_result *StartReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartReplicationJob"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartReplicationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartReplicationJob(request *StartReplicationJobRequest) (_result *StartReplicationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartReplicationJobResponse{}
	_body, _err := client.StartReplicationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopReplicationJobWithOptions(request *StopReplicationJobRequest, runtime *util.RuntimeOptions) (_result *StopReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopReplicationJob"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopReplicationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopReplicationJob(request *StopReplicationJobRequest) (_result *StopReplicationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopReplicationJobResponse{}
	_body, _err := client.StopReplicationJobWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
