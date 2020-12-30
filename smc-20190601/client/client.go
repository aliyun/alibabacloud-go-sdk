// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateReplicationJobRequest struct {
	OwnerId                *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	RegionId               *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken            *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Name                   *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Description            *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	SourceId               *string                                      `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	TargetType             *string                                      `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	ScheduledStartTime     *string                                      `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	ValidTime              *string                                      `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	ImageName              *string                                      `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId             *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SystemDiskSize         *int32                                       `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	VpcId                  *string                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId              *string                                      `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ReplicationParameters  *string                                      `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	NetMode                *int32                                       `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	RunOnce                *bool                                        `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	Frequency              *int32                                       `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	MaxNumberOfImageToKeep *int32                                       `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	InstanceType           *string                                      `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LaunchTemplateId       *string                                      `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string                                      `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	InstanceRamRole        *string                                      `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	ContainerNamespace     *string                                      `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	ContainerRepository    *string                                      `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	ContainerTag           *string                                      `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	LicenseType            *string                                      `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	DataDisk               []*CreateReplicationJobRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	Tag                    []*CreateReplicationJobRequestTag            `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	SystemDiskPart         []*CreateReplicationJobRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
}

func (s CreateReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequest) SetOwnerId(v int64) *CreateReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetResourceOwnerAccount(v string) *CreateReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateReplicationJobRequest) SetRegionId(v string) *CreateReplicationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetClientToken(v string) *CreateReplicationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateReplicationJobRequest) SetName(v string) *CreateReplicationJobRequest {
	s.Name = &v
	return s
}

func (s *CreateReplicationJobRequest) SetDescription(v string) *CreateReplicationJobRequest {
	s.Description = &v
	return s
}

func (s *CreateReplicationJobRequest) SetSourceId(v string) *CreateReplicationJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetTargetType(v string) *CreateReplicationJobRequest {
	s.TargetType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetScheduledStartTime(v string) *CreateReplicationJobRequest {
	s.ScheduledStartTime = &v
	return s
}

func (s *CreateReplicationJobRequest) SetValidTime(v string) *CreateReplicationJobRequest {
	s.ValidTime = &v
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

func (s *CreateReplicationJobRequest) SetSystemDiskSize(v int32) *CreateReplicationJobRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateReplicationJobRequest) SetVpcId(v string) *CreateReplicationJobRequest {
	s.VpcId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetVSwitchId(v string) *CreateReplicationJobRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateReplicationJobRequest) SetReplicationParameters(v string) *CreateReplicationJobRequest {
	s.ReplicationParameters = &v
	return s
}

func (s *CreateReplicationJobRequest) SetNetMode(v int32) *CreateReplicationJobRequest {
	s.NetMode = &v
	return s
}

func (s *CreateReplicationJobRequest) SetRunOnce(v bool) *CreateReplicationJobRequest {
	s.RunOnce = &v
	return s
}

func (s *CreateReplicationJobRequest) SetFrequency(v int32) *CreateReplicationJobRequest {
	s.Frequency = &v
	return s
}

func (s *CreateReplicationJobRequest) SetMaxNumberOfImageToKeep(v int32) *CreateReplicationJobRequest {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *CreateReplicationJobRequest) SetInstanceType(v string) *CreateReplicationJobRequest {
	s.InstanceType = &v
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

func (s *CreateReplicationJobRequest) SetInstanceRamRole(v string) *CreateReplicationJobRequest {
	s.InstanceRamRole = &v
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

func (s *CreateReplicationJobRequest) SetLicenseType(v string) *CreateReplicationJobRequest {
	s.LicenseType = &v
	return s
}

func (s *CreateReplicationJobRequest) SetDataDisk(v []*CreateReplicationJobRequestDataDisk) *CreateReplicationJobRequest {
	s.DataDisk = v
	return s
}

func (s *CreateReplicationJobRequest) SetTag(v []*CreateReplicationJobRequestTag) *CreateReplicationJobRequest {
	s.Tag = v
	return s
}

func (s *CreateReplicationJobRequest) SetSystemDiskPart(v []*CreateReplicationJobRequestSystemDiskPart) *CreateReplicationJobRequest {
	s.SystemDiskPart = v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s CreateReplicationJobRequestDataDiskPart) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDataDiskPart) SetSizeBytes(v int64) *CreateReplicationJobRequestDataDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *CreateReplicationJobRequestDataDiskPart) SetBlock(v bool) *CreateReplicationJobRequestDataDiskPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestDataDiskPart) SetDevice(v string) *CreateReplicationJobRequestDataDiskPart {
	s.Device = &v
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

type CreateReplicationJobRequestSystemDiskPart struct {
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s CreateReplicationJobRequestSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetSizeBytes(v int64) *CreateReplicationJobRequestSystemDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetBlock(v bool) *CreateReplicationJobRequestSystemDiskPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestSystemDiskPart) SetDevice(v string) *CreateReplicationJobRequestSystemDiskPart {
	s.Device = &v
	return s
}

type CreateReplicationJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateReplicationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReplicationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobResponseBody) SetRequestId(v string) *CreateReplicationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReplicationJobResponseBody) SetJobId(v string) *CreateReplicationJobResponseBody {
	s.JobId = &v
	return s
}

type CreateReplicationJobResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateReplicationJobResponse) SetBody(v *CreateReplicationJobResponseBody) *CreateReplicationJobResponse {
	s.Body = v
	return s
}

type CutOverReplicationJobRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SyncData             *bool   `json:"SyncData,omitempty" xml:"SyncData,omitempty"`
}

func (s CutOverReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CutOverReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CutOverReplicationJobRequest) SetOwnerId(v int64) *CutOverReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetResourceOwnerAccount(v string) *CutOverReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CutOverReplicationJobRequest) SetJobId(v string) *CutOverReplicationJobRequest {
	s.JobId = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CutOverReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CutOverReplicationJobResponse) SetBody(v *CutOverReplicationJobResponseBody) *CutOverReplicationJobResponse {
	s.Body = v
	return s
}

type DeleteReplicationJobRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteReplicationJobRequest) SetOwnerId(v int64) *DeleteReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetResourceOwnerAccount(v string) *DeleteReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteReplicationJobRequest) SetJobId(v string) *DeleteReplicationJobRequest {
	s.JobId = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteReplicationJobResponse) SetBody(v *DeleteReplicationJobResponseBody) *DeleteReplicationJobResponse {
	s.Body = v
	return s
}

type DeleteSourceServerRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteSourceServerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceServerRequest) GoString() string {
	return s.String()
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

func (s *DeleteSourceServerRequest) SetForce(v bool) *DeleteSourceServerRequest {
	s.Force = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSourceServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSourceServerResponse) SetBody(v *DeleteSourceServerResponseBody) *DeleteSourceServerResponse {
	s.Body = v
	return s
}

type DescribeReplicationJobsRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	BusinessStatus       *string   `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceId             []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	JobId                []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsRequest) SetOwnerId(v int64) *DescribeReplicationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetResourceOwnerAccount(v string) *DescribeReplicationJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetName(v string) *DescribeReplicationJobsRequest {
	s.Name = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetRegionId(v string) *DescribeReplicationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetStatus(v string) *DescribeReplicationJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetBusinessStatus(v string) *DescribeReplicationJobsRequest {
	s.BusinessStatus = &v
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

func (s *DescribeReplicationJobsRequest) SetSourceId(v []*string) *DescribeReplicationJobsRequest {
	s.SourceId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetJobId(v []*string) *DescribeReplicationJobsRequest {
	s.JobId = v
	return s
}

type DescribeReplicationJobsResponseBody struct {
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ReplicationJobs *DescribeReplicationJobsResponseBodyReplicationJobs `json:"ReplicationJobs,omitempty" xml:"ReplicationJobs,omitempty" type:"Struct"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeReplicationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBody) SetTotalCount(v int32) *DescribeReplicationJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetReplicationJobs(v *DescribeReplicationJobsResponseBodyReplicationJobs) *DescribeReplicationJobsResponseBody {
	s.ReplicationJobs = v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetPageSize(v int32) *DescribeReplicationJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetRequestId(v string) *DescribeReplicationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBody) SetPageNumber(v int32) *DescribeReplicationJobsResponseBody {
	s.PageNumber = &v
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
	Frequency              *int32                                                                              `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	VpcId                  *string                                                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	CreationTime           *string                                                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                 *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ScheduledStartTime     *string                                                                             `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	MaxNumberOfImageToKeep *int32                                                                              `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	ContainerNamespace     *string                                                                             `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	DataDisks              *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks          `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	StatusInfo             *string                                                                             `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	InstanceRamRole        *string                                                                             `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	SystemDiskSize         *int32                                                                              `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Description            *string                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	ReplicationParameters  *string                                                                             `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	ErrorCode              *string                                                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ValidTime              *string                                                                             `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	NetMode                *int32                                                                              `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	ContainerTag           *string                                                                             `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	LicenseType            *string                                                                             `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	Name                   *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId                *string                                                                             `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Progress               *float32                                                                            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RunOnce                *bool                                                                               `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	LaunchTemplateId       *string                                                                             `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	ContainerRepository    *string                                                                             `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	InstanceId             *string                                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SystemDiskParts        *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts    `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	InstanceType           *string                                                                             `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SourceId               *string                                                                             `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	LaunchTemplateVersion  *string                                                                             `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	RegionId               *string                                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TransitionInstanceId   *string                                                                             `json:"TransitionInstanceId,omitempty" xml:"TransitionInstanceId,omitempty"`
	EndTime                *string                                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime              *string                                                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VSwitchId              *string                                                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	JobId                  *string                                                                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ImageName              *string                                                                             `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	BusinessStatus         *string                                                                             `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ReplicationJobRuns     *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns `json:"ReplicationJobRuns,omitempty" xml:"ReplicationJobRuns,omitempty" type:"Struct"`
	TargetType             *string                                                                             `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetFrequency(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Frequency = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetVpcId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.VpcId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetCreationTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.CreationTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStatus(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Status = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetScheduledStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ScheduledStartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetMaxNumberOfImageToKeep(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerNamespace(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerNamespace = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetDataDisks(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.DataDisks = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStatusInfo(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.StatusInfo = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceRamRole(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceRamRole = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSystemDiskSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetDescription(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Description = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetReplicationParameters(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ReplicationParameters = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetErrorCode(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ErrorCode = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetValidTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ValidTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetNetMode(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.NetMode = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerTag(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerTag = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLicenseType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LicenseType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetName(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Name = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetImageId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ImageId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetProgress(v float32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Progress = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetRunOnce(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.RunOnce = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLaunchTemplateId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetContainerRepository(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ContainerRepository = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSystemDiskParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SystemDiskParts = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetInstanceType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.InstanceType = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetSourceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.SourceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetLaunchTemplateVersion(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetRegionId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetTransitionInstanceId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.TransitionInstanceId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetEndTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.EndTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.StartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetVSwitchId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.VSwitchId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetJobId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.JobId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetImageName(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ImageName = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetBusinessStatus(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetReplicationJobRuns(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ReplicationJobRuns = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetTargetType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.TargetType = &v
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
	Size  *int32                                                                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) SetSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) SetParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	s.Parts = v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) SetDevice(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	s.Device = &v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) SetDevice(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	s.Device = &v
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
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetType(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.Type = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetStartTime(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.StartTime = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) SetImageId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	s.ImageId = &v
	return s
}

type DescribeReplicationJobsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeReplicationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeReplicationJobsResponse) SetBody(v *DescribeReplicationJobsResponseBody) *DescribeReplicationJobsResponse {
	s.Body = v
	return s
}

type DescribeSourceServersRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	State                *string   `json:"State,omitempty" xml:"State,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceId             []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersRequest) SetOwnerId(v int64) *DescribeSourceServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetResourceOwnerAccount(v string) *DescribeSourceServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSourceServersRequest) SetJobId(v string) *DescribeSourceServersRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetState(v string) *DescribeSourceServersRequest {
	s.State = &v
	return s
}

func (s *DescribeSourceServersRequest) SetName(v string) *DescribeSourceServersRequest {
	s.Name = &v
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

func (s *DescribeSourceServersRequest) SetSourceId(v []*string) *DescribeSourceServersRequest {
	s.SourceId = v
	return s
}

type DescribeSourceServersResponseBody struct {
	SourceServers *DescribeSourceServersResponseBodySourceServers `json:"SourceServers,omitempty" xml:"SourceServers,omitempty" type:"Struct"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSourceServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBody) SetSourceServers(v *DescribeSourceServersResponseBodySourceServers) *DescribeSourceServersResponseBody {
	s.SourceServers = v
	return s
}

func (s *DescribeSourceServersResponseBody) SetTotalCount(v int32) *DescribeSourceServersResponseBody {
	s.TotalCount = &v
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

func (s *DescribeSourceServersResponseBody) SetPageNumber(v int32) *DescribeSourceServersResponseBody {
	s.PageNumber = &v
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
	CreationTime      *string                                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	HeartbeatRate     *int32                                                                     `json:"HeartbeatRate,omitempty" xml:"HeartbeatRate,omitempty"`
	State             *string                                                                    `json:"State,omitempty" xml:"State,omitempty"`
	DataDisks         *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks       `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	SystemDiskParts   *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	KernelLevel       *int32                                                                     `json:"KernelLevel,omitempty" xml:"KernelLevel,omitempty"`
	SourceId          *string                                                                    `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	AgentVersion      *string                                                                    `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	StatusInfo        *string                                                                    `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	SystemDiskSize    *int32                                                                     `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Description       *string                                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCode         *string                                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId             *string                                                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Platform          *string                                                                    `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ReplicationDriver *string                                                                    `json:"ReplicationDriver,omitempty" xml:"ReplicationDriver,omitempty"`
	Name              *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	SystemInfo        *string                                                                    `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
	Architecture      *string                                                                    `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServer) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetCreationTime(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.CreationTime = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetHeartbeatRate(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.HeartbeatRate = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetState(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.State = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetDataDisks(v *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.DataDisks = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskParts(v *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskParts = v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetKernelLevel(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.KernelLevel = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSourceId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SourceId = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetAgentVersion(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.AgentVersion = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetStatusInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.StatusInfo = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemDiskSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemDiskSize = &v
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetJobId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.JobId = &v
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetName(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Name = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetSystemInfo(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.SystemInfo = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetArchitecture(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Architecture = &v
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
	Size  *int32                                                                            `json:"Size,omitempty" xml:"Size,omitempty"`
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	Path  *string                                                                           `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk) SetSize(v int32) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDisk {
	s.Size = &v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Need      *bool   `json:"Need,omitempty" xml:"Need,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskPartsPart {
	s.Path = &v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Need      *bool   `json:"Need,omitempty" xml:"Need,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetSizeBytes(v int64) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetNeed(v bool) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Need = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetDevice(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Device = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart) SetPath(v string) *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskPartsSystemDiskPart {
	s.Path = &v
	return s
}

type DescribeSourceServersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSourceServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSourceServersResponse) SetBody(v *DescribeSourceServersResponseBody) *DescribeSourceServersResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
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

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyReplicationJobAttributeRequest struct {
	OwnerId                *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                  *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name                   *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Description            *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	TargetType             *string                                               `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	ScheduledStartTime     *string                                               `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	ImageName              *string                                               `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId             *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SystemDiskSize         *int32                                                `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Frequency              *int32                                                `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	MaxNumberOfImageToKeep *int32                                                `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	InstanceType           *string                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LaunchTemplateId       *string                                               `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateVersion  *string                                               `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	InstanceRamRole        *string                                               `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	ContainerNamespace     *string                                               `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	ContainerRepository    *string                                               `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	ContainerTag           *string                                               `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	ValidTime              *string                                               `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	SystemDiskPart         []*ModifyReplicationJobAttributeRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	DataDisk               []*ModifyReplicationJobAttributeRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s ModifyReplicationJobAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequest) SetOwnerId(v int64) *ModifyReplicationJobAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetResourceOwnerAccount(v string) *ModifyReplicationJobAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetJobId(v string) *ModifyReplicationJobAttributeRequest {
	s.JobId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetName(v string) *ModifyReplicationJobAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetDescription(v string) *ModifyReplicationJobAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetTargetType(v string) *ModifyReplicationJobAttributeRequest {
	s.TargetType = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetScheduledStartTime(v string) *ModifyReplicationJobAttributeRequest {
	s.ScheduledStartTime = &v
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

func (s *ModifyReplicationJobAttributeRequest) SetSystemDiskSize(v int32) *ModifyReplicationJobAttributeRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetFrequency(v int32) *ModifyReplicationJobAttributeRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetMaxNumberOfImageToKeep(v int32) *ModifyReplicationJobAttributeRequest {
	s.MaxNumberOfImageToKeep = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetInstanceType(v string) *ModifyReplicationJobAttributeRequest {
	s.InstanceType = &v
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

func (s *ModifyReplicationJobAttributeRequest) SetInstanceRamRole(v string) *ModifyReplicationJobAttributeRequest {
	s.InstanceRamRole = &v
	return s
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

func (s *ModifyReplicationJobAttributeRequest) SetValidTime(v string) *ModifyReplicationJobAttributeRequest {
	s.ValidTime = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetSystemDiskPart(v []*ModifyReplicationJobAttributeRequestSystemDiskPart) *ModifyReplicationJobAttributeRequest {
	s.SystemDiskPart = v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetDataDisk(v []*ModifyReplicationJobAttributeRequestDataDisk) *ModifyReplicationJobAttributeRequest {
	s.DataDisk = v
	return s
}

type ModifyReplicationJobAttributeRequestSystemDiskPart struct {
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetSizeBytes(v int64) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetBlock(v bool) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.Block = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) SetDevice(v string) *ModifyReplicationJobAttributeRequestSystemDiskPart {
	s.Device = &v
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
	SizeBytes *int64  `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
	Block     *bool   `json:"Block,omitempty" xml:"Block,omitempty"`
	Device    *string `json:"Device,omitempty" xml:"Device,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) String() string {
	return tea.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetSizeBytes(v int64) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.SizeBytes = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetBlock(v bool) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.Block = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) SetDevice(v string) *ModifyReplicationJobAttributeRequestDataDiskPart {
	s.Device = &v
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
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyReplicationJobAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyReplicationJobAttributeResponse) SetBody(v *ModifyReplicationJobAttributeResponseBody) *ModifyReplicationJobAttributeResponse {
	s.Body = v
	return s
}

type ModifySourceServerAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SourceId             *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifySourceServerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySourceServerAttributeRequest) GoString() string {
	return s.String()
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

func (s *ModifySourceServerAttributeRequest) SetName(v string) *ModifySourceServerAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifySourceServerAttributeRequest) SetDescription(v string) *ModifySourceServerAttributeRequest {
	s.Description = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySourceServerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySourceServerAttributeResponse) SetBody(v *ModifySourceServerAttributeResponseBody) *ModifySourceServerAttributeResponse {
	s.Body = v
	return s
}

type StartReplicationJobRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StartReplicationJobRequest) SetOwnerId(v int64) *StartReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartReplicationJobRequest) SetResourceOwnerAccount(v string) *StartReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartReplicationJobRequest) SetJobId(v string) *StartReplicationJobRequest {
	s.JobId = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartReplicationJobResponse) SetBody(v *StartReplicationJobResponseBody) *StartReplicationJobResponse {
	s.Body = v
	return s
}

type StopReplicationJobRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StopReplicationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *StopReplicationJobRequest) SetOwnerId(v int64) *StopReplicationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopReplicationJobRequest) SetResourceOwnerAccount(v string) *StopReplicationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopReplicationJobRequest) SetJobId(v string) *StopReplicationJobRequest {
	s.JobId = &v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopReplicationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopReplicationJobResponse) SetBody(v *StopReplicationJobResponseBody) *StopReplicationJobResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
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

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
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

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) CreateReplicationJobWithOptions(request *CreateReplicationJobRequest, runtime *util.RuntimeOptions) (_result *CreateReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateReplicationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateReplicationJob"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CutOverReplicationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CutOverReplicationJob"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteReplicationJobWithOptions(request *DeleteReplicationJobRequest, runtime *util.RuntimeOptions) (_result *DeleteReplicationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteReplicationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteReplicationJob"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSourceServerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSourceServer"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeReplicationJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeReplicationJobs"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSourceServersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSourceServers"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyReplicationJobAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyReplicationJobAttribute"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySourceServerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySourceServerAttribute"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartReplicationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartReplicationJob"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopReplicationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopReplicationJob"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-06-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
