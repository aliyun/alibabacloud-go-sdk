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
	// The maximum number of times that the activation code can be used to import the information of migration sources. Valid values: 1 to 1000.
	//
	// Default value: 100.
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the activation code.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the activation code. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain digits, colons (:), underscores (\_), and hyphens (-).
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The validity period of the activation code. The activation code can no longer be used to import the information of migration sources after the code expires. Unit: day. Valid values: 1 to 90.
	//
	// Default value: 30.
	TimeToLiveInDays *string `json:"TimeToLiveInDays,omitempty" xml:"TimeToLiveInDays,omitempty"`
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
	// The value of the activation code. The value is returned only once after the CreateAccessToken operation is called and cannot be subsequently queried. Make sure that you properly save the returned value.
	AccessTokenCode *string `json:"AccessTokenCode,omitempty" xml:"AccessTokenCode,omitempty"`
	// The ID of the activation code.
	AccessTokenId *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The namespace of the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerNamespace *string `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	// The repository that stores the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerRepository *string `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	// The tag of the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerTag *string `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	// The data disks.
	DataDisk []*CreateReplicationJobRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The description of the migration job.
	//
	// The description must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (\_), and hyphens (-). The description must start with a letter, but cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The interval at which SMC synchronizes incremental data to Alibaba Cloud. Unit: hour. Valid values: 1 to 168.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// By default, this parameter is empty.
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The name of the destination image. The name must meet the following requirements:
	//
	// *   The name must be unique within an Alibaba Cloud region.
	// *   The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (\_), and hyphens (-). The name must start with a letter, but cannot start with `http://` or `https://`.
	//
	// >  If you specify an image name that already exists in the destination region, the migration job ID is appended to the image name as a suffix. Example: ImageName_j-2zexxxxxxxxxxxxx.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the destination ECS instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is assigned to the instance.
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The type of the intermediate instance.
	//
	// You can call the [DescribeInstanceTypes](~~25620~~) operation to query the ECS instance types.
	//
	// *   If you specify this parameter, SMC creates an intermediate instance of the specified instance type. If the specified instance type is unavailable, the migration job fails to be created.
	// *   If you do not specify this parameter, SMC selects an available instance type in a specific order to create an intermediate instance. For more information, see the "How does SMC create an intermediate instance?" section of the SMC FAQ topic.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The type of the migration job. Valid values:
	//
	// *   0: server migration.
	// *   1: operating system migration.
	// *   2: cross-zone migration.
	// *   3: agentless migration for a VMware VM.
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the launch template.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version number of the launch template.
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The license type. Valid values:
	//
	// *   An empty value specifies no license.
	// *   A value of BYOL specifies Bring Your Own License (BYOL).
	//
	// For more information, see [SMC FAQ](~~121707~~).
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The maximum number of images retained for the incremental migration job. Valid values: 1 to 10.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// By default, this parameter is empty.
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job. The name must meet the following requirements:
	//
	// *   The name must be unique.
	// *   The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (\_), and hyphens (-). The name must start with a letter, but cannot start with `http://` or `https://`.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network mode for data transmission. Valid values:
	//
	// *   0: Data is transmitted over the Internet. Make sure that the source server can access the Internet.
	// *   2: Data is transmitted over a VPC. If you specify this value, you must specify the VSwitchId parameter. You do not need to specify the VpcId parameter because the value of the VpcId parameter can be retrieved based on the value of the VSwitchId parameter.
	//
	// Default value: 0
	NetMode *int32 `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Alibaba Cloud region to which you want to migrate the source server.
	//
	// For example, if you want to migrate the source server to the China (Hangzhou) region, set this parameter to `cn-hangzhou`. You can call the [DescribeRegions](~~25609~~) operation to query the latest regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameters of the replication driver. The parameters must be specified as key-value pairs in the JSON format. The keys are fixed for each type of replication driver. The JSON string can be up to 2,048 characters in length.
	//
	// A replication driver is a tool that is used to migrate a source server to an intermediate instance. The parameters vary based on the replication driver type. If you use a Server Migration Tool (SMT) driver, you can specify the following parameters:
	//
	// *   bandwidth_limit: the maximum bandwidth for data transmission.
	// *   compress_level: the compression ratio of data to be transmitted.
	// *   checksum: specifies whether to enable checksum verification.
	//
	// For more information about replication drivers, see the response parameter `SourceServers.ReplicationDriver` of the [DescribeSourceServers](~~121818~~) operation.
	ReplicationParameters *string `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// Specifies whether to disable incremental migration for the source server. Valid values:
	//
	// *   true: creates a migration job that runs only once. This is the default value. Incremental data of the source server is not synchronized.
	// *   false: creates an incremental migration job. In this case, you must specify the `Frequency` parameter. SMC synchronizes incremental data of the source server to Alibaba Cloud at the specified frequency. You can use an incremental migration job to synchronize incremental data from the source server to Alibaba Cloud without the need to interrupt your business. A full data image is generated for the source server when the job is running.
	//
	// >  You can specify this parameter only when you create a migration job. The parameter value cannot be changed after the migration job is created.
	RunOnce *bool `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	// The time when you want to run the migration job. The time must meet the following requirements:
	//
	// *   The time must be specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. For example, 2018-01-01T12:00:00Z specifies 20:00:00 on January 1, 2018 (UTC+8).
	// *   The value must be within 30 days after the current time.
	//
	// >  If you do not specify this parameter, you must manually start the migration job after the job is created. You can call the [StartReplicationJob](~~121823~~) operation to start the migration job.
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The ID of the source server.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The information about system disk partitions.
	SystemDiskPart []*CreateReplicationJobRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	// The system disk size of the destination ECS instance. Unit: GiB. Valid values: 20 to 2048.
	//
	// >  The value must be greater than the used space of the system disk on the source server. For example, if the total size of the source disk is 500 GiB and the used space is 100 GiB, the value of this parameter must be greater than 100 GiB.
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags.
	Tag []*CreateReplicationJobRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of destination to which you want to migrate the source server. Valid values:
	//
	// *   Image: After the migration job is complete, SMC generates an Elastic Compute Service (ECS) image for the source server.
	// *   ContainerImage: After the migration job is complete, SMC generates a Docker container image for the source server.
	// *   TargetInstance: After the migration job is completed, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the `InstanceId` parameter.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	//
	// You must set this parameter if you use a VPC to migrate data.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The time when the migration job expires. You can schedule the migration job to expire 7 to 90 days after the job is created.
	//
	// *   The time must be specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. For example, 2018-01-01T12:00:00Z specifies 20:00:00 on January 1, 2018 (UTC+8).
	// *   If you do not specify this parameter, the migration job does not expire.
	// *   After a migration job expires, the job state changes to Expired. SMC retains the migration job for seven days after the job expires. After the job is retained for seven days, SMC deletes the migration job.
	//
	// By default, a migration job is valid for 30 days after it is created.
	ValidTime *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	// The ID of a VPC for which you have configured an Express Connect circuit or a VPN gateway.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *CreateReplicationJobRequest) SetResourceGroupId(v string) *CreateReplicationJobRequest {
	s.ResourceGroupId = &v
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
	// The index of data disk N on the destination ECS instance. Data disks on a destination ECS instance are arranged in a sequential order that starts from 1. Valid values: 1 to 16.
	//
	// >  To create a destination data disk for a source server, make sure that the source server has data disks.
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The data disk partitions.
	Part []*CreateReplicationJobRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the data disk on the destination ECS instance. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The size of a destination data disk must be larger than the size of data in the source data disk. For example, if the size of the source data disk is 500 GiB and the used space is 100 GiB, you must set this parameter to a value greater than 100.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// Specifies whether to enable block replication for partition N in the destination data disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of partition N in the destination data disk. The partitions in the destination data disk are arranged in the same sequential order as those in the source data disk.
	//
	// >  You must set both the DataDisk.N.Part.N.Device and `DataDisk.N.Part.N.SizeBytes` parameters or leave both parameters empty.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination data disk. Unit: bytes. The default value is equal to the corresponding partition size of the source data disk.
	//
	// >
	//
	// *   The total size of all partitions in a destination data disk cannot exceed the size of the destination data disk.
	//
	// *   You must set both the `DataDisk.N.Part.N.Device` and DataDisk.N.Part.N.SizeBytes parameters or leave both parameters empty.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// Specifies whether to enable block replication for partition N in the destination system disk. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination system disk. The partitions in the destination system disk are arranged in the same sequential order as those in the source system disk.
	//
	// >  You must set both the SystemDiskPart.N.Device and `SystemDiskPart.N.SizeBytes` parameters or leave both parameters empty.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the partition N in the destination system disk. Unit: bytes. The default value is equal to the partition size of the source system disk.
	//
	// >
	//
	// *   The total size of all partitions in the destination system disk cannot exceed the size of the destination system disk.
	//
	// *   You must set both the `SystemDiskPart.N.Device` and SystemDiskPart.N.SizeBytes parameters or leave both parameters empty.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// The key of the tag for the migration job. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag for the migration job. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
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
	// The ID of the migration job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
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
	// The ID of the incremental migration job.
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// Specifies whether to migrate full data for the last time. Valid Values:
	//
	// *   true: migrates full data for the last time.
	// *   false: does not migrate full data for the last time.
	//
	// Default value: false.
	SyncData *bool `json:"SyncData,omitempty" xml:"SyncData,omitempty"`
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
	// The request ID.
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
	// The ID of the activation code.
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
	// The request ID.
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
	// The migration job ID.
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
	// The request ID.
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
	// Specifies whether to forcibly delete the migration source. Valid values:
	//
	// *   true: forcibly deletes the migration source and the migration job created for the migration source, and releases the intermediate resources of the migration job.
	// *   false: does not delete the migration source if a migration job is created for the migration source.
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The migration source ID.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
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
	// The request ID.
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
	// The business status of the migration job. Valid values:
	//
	// *   Preparing: The migration is being prepared.
	// *   Syncing: Data is being synchronized.
	// *   Processing: The migration is in progress.
	// *   Cleaning: Intermediate resources are being released.
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The IDs of the destination Elastic Compute Service (ECS) instances.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The IDs of the migration jobs. You can specify a maximum of 50 IDs.
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	// The type of the migration job. Valid values:
	//
	// *   0: server migration.
	// *   1: operating system migration.
	// *   2: cross-zone migration.
	// *   3: agentless migration for a VMware VM.
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the migration job.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Alibaba Cloud region to which you want to migrate the source server.
	//
	// For example, if you want to migrate a source server to the China (Hangzhou) region, set this parameter to `cn-hangzhou`. You can call the [DescribeRegions](~~25609~~) operation to query the latest regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The IDs of the source servers. You can specify a maximum of 50 IDs.
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The status of the migration job. Valid values:
	//
	// *   Ready: The migration job is not started.
	// *   Running: The migration job is running.
	// *   Stopped: The migration job is paused.
	// *   InError: An error occurs in the migration job.
	// *   Finished: The migration job is complete.
	// *   Waiting: The migration job is waiting to run.
	// *   Expired: The migration job has expired.
	// *   Deleting: The migration job is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about tags that are attached to the SMC resource.
	Tag []*DescribeReplicationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeReplicationJobsRequest) SetResourceGroupId(v string) *DescribeReplicationJobsRequest {
	s.ResourceGroupId = &v
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

func (s *DescribeReplicationJobsRequest) SetTag(v []*DescribeReplicationJobsRequestTag) *DescribeReplicationJobsRequest {
	s.Tag = v
	return s
}

type DescribeReplicationJobsRequestTag struct {
	// The key of the tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReplicationJobsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsRequestTag) SetKey(v string) *DescribeReplicationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeReplicationJobsRequestTag) SetValue(v string) *DescribeReplicationJobsRequestTag {
	s.Value = &v
	return s
}

type DescribeReplicationJobsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of migration jobs.
	ReplicationJobs *DescribeReplicationJobsResponseBodyReplicationJobs `json:"ReplicationJobs,omitempty" xml:"ReplicationJobs,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of migration jobs returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The business status of the migration job. Valid values:
	//
	// *   Preparing: The migration is being prepared.
	// *   Syncing: Data is being synchronized.
	// *   Processing: The migration is in progress.
	// *   Cleaning: Intermediate resources are being released.
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The namespace of the destination Docker container image.
	ContainerNamespace *string `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	// The repository that stores the destination Docker container image.
	ContainerRepository *string `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	// The tag of the destination Docker container image.
	ContainerTag *string `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	// The time when the migration job was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data disks on the destination ECS instance.
	DataDisks *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	// The description of the migration job.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the migration job was complete. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code returned if an error occurred in the migration job.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The interval at which the incremental migration job runs. Unit: hour. Valid values: 1 to 168.
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ID of the destination image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the destination image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the destination ECS instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is assigned to the instance.
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The instance type of the intermediate instance.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the migration job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the migration job. Valid values:
	//
	// *   0: server migration.
	// *   1: operating system migration.
	// *   2: cross-zone migration.
	// *   3: agentless migration for a VMware VM.
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the launch template.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The versions of the launch template.
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The type of license for the migration job. Valid values:
	//
	// *   An empty value indicates no license.
	// *   A value of BYOL indicates Bring Your Own License (BYOL).
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The maximum number of images retained for the incremental migration job. Valid values: 1 to 10.
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of network used for the migration.
	NetMode *int32 `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	// The progress of the migration job.
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the Alibaba Cloud region to which the source server is migrated.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution records of the migration job.
	ReplicationJobRuns *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns `json:"ReplicationJobRuns,omitempty" xml:"ReplicationJobRuns,omitempty" type:"Struct"`
	// The string of key-value pairs configured for the replication driver.
	ReplicationParameters *string `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether incremental migration is disabled for the source server. Valid values:
	//
	// *   true: Incremental migration is disabled. A migration job runs only once after the job is created.
	// *   false: Incremental migration is enabled. For an incremental migration job, SMC synchronizes incremental data to Alibaba Cloud at the interval specified by the `Frequency` parameter.
	RunOnce *bool `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	// The time when the migration job is scheduled to run. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. The time must meet the following requirements:
	//
	// *   The value must be within 30 days after the current time.
	// *   If you do not specify this parameter, you must manually start the migration job after the migration job is created. You can call the [StartReplicationJob](~~121823~~) operation to start the migration job.
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The ID of the source server.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The time when the migration job was started. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the migration job. Valid values:
	//
	// *   Ready: The migration job is not started.
	// *   Running: The migration job is running.
	// *   Stopped: The migration job is paused.
	// *   InError: An error occurs in the migration job.
	// *   Finished: The migration job is complete.
	// *   Waiting: The migration job is waiting to run.
	// *   Expired: The migration job has expired.
	// *   Deleting: The migration job is being deleted.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information about the migration job.
	StatusInfo *string `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	// The system disk partitions.
	SystemDiskParts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	// The size of the system disk of the destination ECS instance.
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The information about tags that are attached to the SMC resource.
	Tags *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of destination to which the source server is migrated. Valid values:
	//
	// *   Image: After the migration job is complete, SMC generates an ECS image for the source server.
	// *   ContainerImage: After the migration job is complete, SMC generates a Docker container image for the source server.
	// *   TargetInstance: After the migration job is complete, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the InstanceId parameter.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the intermediate instance.
	TransitionInstanceId *string `json:"TransitionInstanceId,omitempty" xml:"TransitionInstanceId,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The time when the migration job expired. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	ValidTime *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	// The ID of a virtual private cloud (VPC) for which you have configured an Express Connect circuit or a VPN gateway.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetResourceGroupId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.ResourceGroupId = &v
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetTags(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Tags = v
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
	// The index number of the data disk.
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The data disk partitions.
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The size of the data disk. Unit: GiB.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// Indicates whether block replication is enabled for the data disk partition.
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of the data disk partition.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// The time when the migration job ended. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the destination image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The time when the migration job was started. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The method used to run the migration job. Valid values:
	//
	// *   Manual: The migration job was manually started.
	// *   Schedule: The migration job was started at a scheduled time or at a specific interval.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// Indicates whether block replication is enabled for the system disk partition.
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of the system disk partition.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the system disk partition. Unit: bytes.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags struct {
	Tag []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) SetTag(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags {
	s.Tag = v
	return s
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag struct {
	// The key of the tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) SetKey(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) SetValue(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag {
	s.Value = &v
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
	// The migration job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the migration source. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain digits, colons (:), underscores (\_), and hyphens (-).
	//
	// Default value: null.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The migration source IDs. You can specify multiple IDs.
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The state of the migration source. Valid Values:
	//
	// *   Unavailable: The migration source is inactive, or an error occurs in the migration source.
	// *   Available: The migration source is active.
	// *   InUse: The migration source is being migrated.
	// *   Deleting: The migration source is being deleted from Server Migration Center (SMC).
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag.
	Tag []*DescribeSourceServersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeSourceServersRequest) SetResourceGroupId(v string) *DescribeSourceServersRequest {
	s.ResourceGroupId = &v
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

func (s *DescribeSourceServersRequest) SetTag(v []*DescribeSourceServersRequestTag) *DescribeSourceServersRequest {
	s.Tag = v
	return s
}

type DescribeSourceServersRequestTag struct {
	// The key of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSourceServersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersRequestTag) SetKey(v string) *DescribeSourceServersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSourceServersRequestTag) SetValue(v string) *DescribeSourceServersRequestTag {
	s.Value = &v
	return s
}

type DescribeSourceServersResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the migration source.
	SourceServers *DescribeSourceServersResponseBodySourceServers `json:"SourceServers,omitempty" xml:"SourceServers,omitempty" type:"Struct"`
	// The total number of migration sources returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The version number of the SMC client.
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The system architecture of the migration source.
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the migration source was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data disk on the migration source.
	DataDisks *DescribeSourceServersResponseBodySourceServersSourceServerDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	// The description of the migration source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error code of the migration source.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The interval at which heartbeats are sent from the SMC client. Unit: seconds.
	HeartbeatRate *int32 `json:"HeartbeatRate,omitempty" xml:"HeartbeatRate,omitempty"`
	// The ID of the last migration job.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The kernel level of the migration source.
	KernelLevel *int32 `json:"KernelLevel,omitempty" xml:"KernelLevel,omitempty"`
	// The name of the migration source.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system of the migration source.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The replication driver used for migration. Default value: SMT.
	ReplicationDriver *string `json:"ReplicationDriver,omitempty" xml:"ReplicationDriver,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the migration source.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The state of the migration source.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status information of the migration source. This parameter is returned if the migration source is in the Unavailable state. The value of this parameter consists of key-value pairs in the JSON format. Sample keys:
	//
	//     error_code: The error code.
	//     error_msg: The error message.
	StatusInfo *string `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	// The information about the system disk partition.
	SystemDiskParts *DescribeSourceServersResponseBodySourceServersSourceServerSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	// The system disk size of the migration source. Unit: GiB.
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The system information of the migration source. The value of this parameter consists of key-value pairs in the JSON format. The key-value pairs are extensible and have fixed keys. The JSON string does not exceed 1 KB in size. Sample keys:
	//
	//     agent_mode: The migration mode.
	//     agent_type: The migration type.
	//     client_type: The client type.
	//     hostname: The host name.
	//     ipv4: The IPv4 address.
	//     ipv6: The IPv6 address.
	//     .cores: The number of CPU cores.
	//     cpu_usage: The CPU utilization.
	//     memory: The memory size.
	//     memory_usage: The memory usage.
	SystemInfo *string `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
	// The information about the tags.
	Tags *DescribeSourceServersResponseBodySourceServersSourceServerTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetResourceGroupId(v string) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.ResourceGroupId = &v
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

func (s *DescribeSourceServersResponseBodySourceServersSourceServer) SetTags(v *DescribeSourceServersResponseBodySourceServersSourceServerTags) *DescribeSourceServersResponseBodySourceServersSourceServer {
	s.Tags = v
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
	// The index number of the data disk.
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The information about the data disk partition.
	Parts *DescribeSourceServersResponseBodySourceServersSourceServerDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The path of the data disk.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk. Unit: GiB.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// Indicates whether block replication is enabled for the data disk partition.
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The device ID of the data disk partition.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Indicates whether the data disk partition must be selected.
	Need *bool `json:"Need,omitempty" xml:"Need,omitempty"`
	// The path of the data disk partition.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk partition. Unit: byte.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// Indicates whether block replication is enabled for the system disk partition.
	CanBlock *bool `json:"CanBlock,omitempty" xml:"CanBlock,omitempty"`
	// The device ID of the system disk partition.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Indicates whether the system disk partition must be selected.
	Need *bool `json:"Need,omitempty" xml:"Need,omitempty"`
	// The path of the system disk partition.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the system disk partition. Unit: byte.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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

type DescribeSourceServersResponseBodySourceServersSourceServerTags struct {
	Tag []*DescribeSourceServersResponseBodySourceServersSourceServerTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTags) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTags) SetTag(v []*DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) *DescribeSourceServersResponseBodySourceServersSourceServerTags {
	s.Tag = v
	return s
}

type DescribeSourceServersResponseBodySourceServersSourceServerTagsTag struct {
	// The key of tag N that is added to the SMC resource. Valid values of N: 1 to 20
	//
	// You cannot specify empty strings as tag keys. It can be up to 64 characters in length and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20
	//
	// The tag key can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) SetKey(v string) *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag) SetValue(v string) *DescribeSourceServersResponseBodySourceServersSourceServerTagsTag {
	s.Value = &v
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
	// The ID of the activation code.
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
	// The request ID.
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
	// The information about activation codes.
	AccessTokenId []*string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty" type:"Repeated"`
	// The name of the activation code.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The status of the activation code. Valid values:
	//
	// *   activated
	// *   unactivated
	// *   expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The activation codes returned.
	AccessTokens *ListAccessTokensResponseBodyAccessTokens `json:"AccessTokens,omitempty" xml:"AccessTokens,omitempty" type:"Struct"`
	// The number of entries per page. Valid values:
	//
	// *   10
	// *   20
	// *   50
	//
	// Default value: 20.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of migration sources returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the activation code.
	AccessTokenId *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	// The maximum number of times that the activation code can be used. Valid values: 1 to 1000.
	//
	// Default value: 100.
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the activation code was created. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the activation code.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the activation code.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of migration sources whose information has been imported to Server Migration Center (SMC) by using the activation code.
	RegisteredCount *string `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	// The status of the activation code. Valid values:
	//
	// *   activated
	// *   unactivated
	// *   expired
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The validity period of the activation code. Unit: day. Valid values: 1 to 90. Default value: 30.
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of SMC resources. SMC resources include migration sources and migration jobs. You can specify a maximum of 50 SMC resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the SMC resource. Valid values:
	//
	// *   sourceserver: migration source.
	// *   replicationjob: migration job.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that are attached to SMC resources.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of tag N. The tag key must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// Tag.N is used for exact match of SMC resources to which the tag is attached. Tag N consists of a key-value pair.
	//
	// *   Tag keys and values are case-sensitive.
	// *   If you set only the Tag.N.Key parameter, all resources to which the specified tags are attached are returned.
	// *   If you set only the Tag.N.Value parameter, the error message InvalidParameter.TagValue is returned.
	// *   If you specify multiple tag key-value pairs at a time, only SMC resources that match all tag key-value pairs are returned.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. The value must be 1 to 64 characters in length. Valid values of N: 1 to 20.
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If NextToken is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about SMC resources and tags, such as the IDs, types, and tag key-value pairs of the resources.
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
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is attached to the resource.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is attached to the resource.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	// The namespace of the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerNamespace *string `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	// The repository that stores the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerRepository *string `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	// The tag of the destination Docker container image. For more information about Docker container images, see [Terms](~~60744~~).
	ContainerTag *string `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	// The information about the data disk.
	DataDisk []*ModifyReplicationJobAttributeRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The description of the migration job.
	//
	// The description must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (\_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The interval at which an incremental migration job runs. Unit: hour. Valid values: 1 to 168.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The name of the destination image. The name must meet the following requirements:
	//
	// *   The name must be unique within an Alibaba Cloud region.
	// *   The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (\_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// >  If an image whose name is the same as that of the destination image already exists in the current region when the migration job is in progress, the system adds the migration job ID to the end of the image name by default. Example: ImageName-JobId.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The destination instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is attached to the intermediate instance.
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The type of the intermediate instance.
	//
	// You can call the [DescribeInstanceTypes](~~25620~~) operation to query the ECS instance types.
	//
	// *   If you specify this parameter, SMC creates an intermediate instance of the specified instance type. If the specified instance type is unavailable, you cannot create the migration job.
	// *   If you do not specify this parameter, SMC selects an available instance type in a specific order to create an intermediate instance. For more information,
	//
	// see the "How does SMC create an intermediate instance?" section of the "FAQ" topic.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The migration job ID.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The launch template ID.
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version number of the launch template.
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum number of images that are retained for an incremental migration job. Valid values: 1 to 10.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job. The name must meet the following requirements:
	//
	// *   The name must be unique.
	// *   The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (\_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The time when the migration job is executed. SMC starts the migration job at the specified time.
	//
	// Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. For example, 2018-01-01T12:00:00Z indicates 20:00:00 on January 1, 2018 (UTC+8).
	//
	// >  If ScheduledStartTime is left empty, SMC does not automatically start the migration job. In this case, you must call the [StartReplicationJob](~~121823~~) operation to start the migration job.
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The partition information of the system disk.
	SystemDiskPart []*ModifyReplicationJobAttributeRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	// The system disk size of the destination ECS instance. Unit: GiB. Valid values: 20 to 500.
	//
	// >  The size of a destination data disk must be greater than the size of data in the source data disk. For example, if the source data disk has 500 GiB of storage space and 100 GiB of data, you must set this parameter to a value greater than 100.
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The type of destination to which the source server is migrated. You can modify the value only before the migration job starts. Valid values:
	//
	// *   Image: After the migration job is complete, Server Migration Center (SMC) generates a destination Elastic Compute Service (ECS) image for the source server. You can use the image to create an ECS instance.
	// *   ContainerImage: After the migration job is complete, SMC generates a container image for the source server. You can use the container image in Container Registry.
	// *   TargetInstance: After the migration job is complete, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the `InstanceId` parameter.
	//
	// >
	//
	// *   The value of this parameter is not case-sensitive.
	//
	// *   SMC does not allow you to migrate Windows servers or servers that run operating systems on the ARM architecture to Container Registry.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The time when the migration job expires. You can schedule the migration job to expire 7 to 90 days after the job is created.
	//
	// *   This parameter can be modified only if the migration job is in the Ready, Running, Stopped, InError, or Waiting state.
	// *   Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC. For example, 2018-01-01T12:00:00Z indicates 20:00:00 on January 1, 2018 (UTC+8).
	// *   If you do not specify this parameter, the migration job does not expire.
	// *   After a migration job expires, the job state changes to Expired. SMC retains the migration job for seven days after the job expires. After the job is retained for seven days, SMC deletes the migration job.
	//
	// By default, a migration job is valid for 30 days after it is created.
	ValidTime *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
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
	// The index of data disk N on the destination ECS instance. Valid values of N: 1 to 16.
	//
	// Data disks on a destination ECS instance are arranged in a sequential order that starts from 1.
	//
	// >  You can create a destination data disk only for a source server that has data disks.
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The information about partitions.
	Part []*ModifyReplicationJobAttributeRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the data disk on the destination ECS instance. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The size of a destination data disk must be greater than the size of data in the source data disk. For example, if the source data disk has 500 GiB of storage space and 100 GiB of data, you must set this parameter to a value greater than 100.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// Specifies whether to enable block replication for partition N in the destination data disk. Valid values:
	//
	// *   true
	// *   false
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination data disk.
	//
	// >  The partitions in the destination data disk are arranged in the same sequential order as those in the source data disk.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination data disk. Unit: bytes. The default value is equal to the corresponding size of the partition in the source data disk.
	//
	// >  The total size of all partitions in the destination data disk cannot exceed the size of the destination data disk.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// Specifies whether to enable block replication for partition N in the destination system disk. Valid values:
	//
	// *   true
	// *   false
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination system disk.
	//
	// >  The partitions in the destination system disk are arranged in the same sequential order as those in the source system disk.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination system disk. Unit: bytes. The default value is equal to the partition size of the source system disk.
	//
	// >  The total size of all partitions in the destination system disk cannot exceed the size of the destination system disk.
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
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
	// The request ID.
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
	// The description of the migration source. The description can be up to 256 characters in length and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the migration source. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (\_), and hyphens (-).
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The migration source ID.
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
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
	// The request ID.
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
	// The migration job ID.
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
	// The request ID.
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
	// The migration job ID.
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
	// The request ID.
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of N SMC resources. SMC resources include migration sources and jobs. Valid values of N: 1 to 50.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the SMC resource. Valid values:
	//
	// *   sourceserver: migration source.
	// *   replicationjob: migration job.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of tag N to be added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to be added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
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
	// The request ID.
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
	// Specifies whether to remove all tags that are added to the specified SMC resource. This parameter is valid only if you do not set `TagKey.N`. Valid values:
	//
	// *   true: removes all tags that are added to the specified SMC resource. If no tags are added to the specified SMC resource, no operation is performed.
	// *   false: does not remove tags that are added to the specified SMC resource.
	//
	// Default value: false.
	All     *bool  `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of N SMC resources. SMC resources include migration sources and jobs. Valid values of N: 1 to 50.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the SMC resource. Valid values:
	//
	// *   sourceserver: migration source.
	// *   replicationjob: migration job.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N that is added to the SMC resource. Tag keys are case-sensitive. Valid values of N: 1 to 20.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	// The request ID.
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

/**
 * If you want to import the information of migration sources by using an activation code, you can call this operation to create one.
 *
 * @param request CreateAccessTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAccessTokenResponse
 */
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

/**
 * If you want to import the information of migration sources by using an activation code, you can call this operation to create one.
 *
 * @param request CreateAccessTokenRequest
 * @return CreateAccessTokenResponse
 */
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

/**
 * ## Usage notes
 * *   You can create migration jobs only for source servers that are in the Available state.
 * *   Each source server can be associated with only one migration job that is in the Ready, Running, Stopped, Waiting, InError, or Expired state.
 * *   You can create a maximum of 1,000 migration jobs within each Alibaba Cloud account.
 * *   If you migrate a source server to an image, you must specify the ImageName, SystemDiskSize, and DataDisk parameters.
 * *   If you use a virtual private cloud (VPC) to migrate data, the VSwitchId parameter is required and the VpcId parameter is optional.
 * *   Server Migration Center (SMC) allows you to migrate source servers to Docker container images. This allows you to migrate containerized applications in a cost-effective way.
 *
 * @param request CreateReplicationJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateReplicationJobResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

/**
 * ## Usage notes
 * *   You can create migration jobs only for source servers that are in the Available state.
 * *   Each source server can be associated with only one migration job that is in the Ready, Running, Stopped, Waiting, InError, or Expired state.
 * *   You can create a maximum of 1,000 migration jobs within each Alibaba Cloud account.
 * *   If you migrate a source server to an image, you must specify the ImageName, SystemDiskSize, and DataDisk parameters.
 * *   If you use a virtual private cloud (VPC) to migrate data, the VSwitchId parameter is required and the VpcId parameter is optional.
 * *   Server Migration Center (SMC) allows you to migrate source servers to Docker container images. This allows you to migrate containerized applications in a cost-effective way.
 *
 * @param request CreateReplicationJobRequest
 * @return CreateReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * *   The incremental migration job must be in the Waiting state.
 * *   After you call this operation, the incremental migration job no longer periodically runs. In the meantime, Server Migration Center (SMC) determines whether to perform a full data migration for the last time based on the value of the `SyncData` parameter. If you set the SyncData parameter to `false`, SMC releases intermediate resources without data migration before the migration job is complete. If you set the SyncData parameter to `true`, SMC performs a full data migration and releases intermediate resources before the migration job is complete.
 *
 * @param request CutOverReplicationJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CutOverReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * *   The incremental migration job must be in the Waiting state.
 * *   After you call this operation, the incremental migration job no longer periodically runs. In the meantime, Server Migration Center (SMC) determines whether to perform a full data migration for the last time based on the value of the `SyncData` parameter. If you set the SyncData parameter to `false`, SMC releases intermediate resources without data migration before the migration job is complete. If you set the SyncData parameter to `true`, SMC performs a full data migration and releases intermediate resources before the migration job is complete.
 *
 * @param request CutOverReplicationJobRequest
 * @return CutOverReplicationJobResponse
 */
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

/**
 * You can call this operation to delete an activation code if you no longer need to import the information about migration sources by using the activation code or if the activation code has expired.
 *
 * @param request DeleteAccessTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAccessTokenResponse
 */
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

/**
 * You can call this operation to delete an activation code if you no longer need to import the information about migration sources by using the activation code or if the activation code has expired.
 *
 * @param request DeleteAccessTokenRequest
 * @return DeleteAccessTokenResponse
 */
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

/**
 * ## Usage notes
 * *   Deleted migration jobs cannot be restored.
 * *   After a migration job is deleted, associated resources, such as the intermediate instance, are automatically released.
 *
 * @param request DeleteReplicationJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * *   Deleted migration jobs cannot be restored.
 * *   After a migration job is deleted, associated resources, such as the intermediate instance, are automatically released.
 *
 * @param request DeleteReplicationJobRequest
 * @return DeleteReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * *   If a migration job is created for the migration source and the migration job is in the Running state, the migration source cannot be deleted.
 * *   If a migration job is created for the migration source but the migration job is not in the Running state, you can set the `Force` parameter to true to delete the migration source.
 *
 * @param request DeleteSourceServerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSourceServerResponse
 */
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

/**
 * ## Usage notes
 * *   If a migration job is created for the migration source and the migration job is in the Running state, the migration source cannot be deleted.
 * *   If a migration job is created for the migration source but the migration job is not in the Running state, you can set the `Force` parameter to true to delete the migration source.
 *
 * @param request DeleteSourceServerRequest
 * @return DeleteSourceServerResponse
 */
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

/**
 * ## Usage notes
 * *   You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are used as filter conditions.
 * *   Server Migration Center (SMC) allows you to migrate source servers to Docker container images. You can use SMC to migrate containerized applications in a cost-effective way. For more information, see [Terms](~~60744~~).
 *
 * @param request DescribeReplicationJobsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeReplicationJobsResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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

/**
 * ## Usage notes
 * *   You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are used as filter conditions.
 * *   Server Migration Center (SMC) allows you to migrate source servers to Docker container images. You can use SMC to migrate containerized applications in a cost-effective way. For more information, see [Terms](~~60744~~).
 *
 * @param request DescribeReplicationJobsRequest
 * @return DescribeReplicationJobsResponse
 */
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

/**
 * ## Usage notes
 * You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are used as filter conditions.
 *
 * @param request DescribeSourceServersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSourceServersResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
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

/**
 * ## Usage notes
 * You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are used as filter conditions.
 *
 * @param request DescribeSourceServersRequest
 * @return DescribeSourceServersResponse
 */
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

/**
 * To prevent an activation code from being leaked, you can call this operation to disable the activation code. Disabled activation codes can no longer be used to import the information about migration sources. However, migration sources whose information has been imported are not affected.
 *
 * @param request DisableAccessTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DisableAccessTokenResponse
 */
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

/**
 * To prevent an activation code from being leaked, you can call this operation to disable the activation code. Disabled activation codes can no longer be used to import the information about migration sources. However, migration sources whose information has been imported are not affected.
 *
 * @param request DisableAccessTokenRequest
 * @return DisableAccessTokenResponse
 */
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

/**
 * You can call this operation to query activation codes and the usage details of the activation codes. An activation code can be in the activated, unactivated, or expired state.
 *
 * @param request ListAccessTokensRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAccessTokensResponse
 */
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

/**
 * You can call this operation to query activation codes and the usage details of the activation codes. An activation code can be in the activated, unactivated, or expired state.
 *
 * @param request ListAccessTokensRequest
 * @return ListAccessTokensResponse
 */
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

/**
 * ## Usage notes
 * Before you modify the parameters of a migration job, take note of the following information:
 * *   The `Name` and `Description` parameters can be modified during the lifecycle of the migration job.
 * *   The `Frequency` and `MaxNumberOfImageToKeep` parameters can be modified only before the migration job is executed or when the migration job is in the `Waiting` state.
 * *   Other parameters can be modified only before the migration job is executed.
 *
 * @param request ModifyReplicationJobAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyReplicationJobAttributeResponse
 */
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

/**
 * ## Usage notes
 * Before you modify the parameters of a migration job, take note of the following information:
 * *   The `Name` and `Description` parameters can be modified during the lifecycle of the migration job.
 * *   The `Frequency` and `MaxNumberOfImageToKeep` parameters can be modified only before the migration job is executed or when the migration job is in the `Waiting` state.
 * *   Other parameters can be modified only before the migration job is executed.
 *
 * @param request ModifyReplicationJobAttributeRequest
 * @return ModifyReplicationJobAttributeResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation regardless of the status of the migration source.
 *
 * @param request ModifySourceServerAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifySourceServerAttributeResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation regardless of the status of the migration source.
 *
 * @param request ModifySourceServerAttributeRequest
 * @return ModifySourceServerAttributeResponse
 */
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

/**
 * ## Usage notes
 * This operation can only be used to start the migration jobs that are in the Ready, Stopped, or InError state.
 *
 * @param request StartReplicationJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * This operation can only be used to start the migration jobs that are in the Ready, Stopped, or InError state.
 *
 * @param request StartReplicationJobRequest
 * @return StartReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation to pause only a migration job whose primary status is Running and business status is Syncing.
 *
 * @param request StopReplicationJobRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation to pause only a migration job whose primary status is Running and business status is Syncing.
 *
 * @param request StopReplicationJobRequest
 * @return StopReplicationJobResponse
 */
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

/**
 * ## Usage notes
 * Up to 20 tags can be added to each SMC resource.
 * Before you add tags to an SMC resource, Alibaba Cloud checks the number of the tags that have been added to the resource. If the maximum number is reached, an error message is returned.
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
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

/**
 * ## Usage notes
 * Up to 20 tags can be added to each SMC resource.
 * Before you add tags to an SMC resource, Alibaba Cloud checks the number of the tags that have been added to the resource. If the maximum number is reached, an error message is returned.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
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

/**
 * You can call this operation to remove tags that are added to one or more SMC resources and deletes the tags if the tags are no longer used.
 *
 * @param request UntagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
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

/**
 * You can call this operation to remove tags that are added to one or more SMC resources and deletes the tags if the tags are no longer used.
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
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
