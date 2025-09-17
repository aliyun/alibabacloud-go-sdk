// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateReplicationJobRequest
	GetClientToken() *string
	SetContainerNamespace(v string) *CreateReplicationJobRequest
	GetContainerNamespace() *string
	SetContainerRepository(v string) *CreateReplicationJobRequest
	GetContainerRepository() *string
	SetContainerTag(v string) *CreateReplicationJobRequest
	GetContainerTag() *string
	SetDataDisk(v []*CreateReplicationJobRequestDataDisk) *CreateReplicationJobRequest
	GetDataDisk() []*CreateReplicationJobRequestDataDisk
	SetDescription(v string) *CreateReplicationJobRequest
	GetDescription() *string
	SetDisks(v *CreateReplicationJobRequestDisks) *CreateReplicationJobRequest
	GetDisks() *CreateReplicationJobRequestDisks
	SetFrequency(v int32) *CreateReplicationJobRequest
	GetFrequency() *int32
	SetImageName(v string) *CreateReplicationJobRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateReplicationJobRequest
	GetInstanceId() *string
	SetInstanceRamRole(v string) *CreateReplicationJobRequest
	GetInstanceRamRole() *string
	SetInstanceType(v string) *CreateReplicationJobRequest
	GetInstanceType() *string
	SetJobType(v int32) *CreateReplicationJobRequest
	GetJobType() *int32
	SetLaunchTemplateId(v string) *CreateReplicationJobRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersion(v string) *CreateReplicationJobRequest
	GetLaunchTemplateVersion() *string
	SetLicenseType(v string) *CreateReplicationJobRequest
	GetLicenseType() *string
	SetMaxNumberOfImageToKeep(v int32) *CreateReplicationJobRequest
	GetMaxNumberOfImageToKeep() *int32
	SetName(v string) *CreateReplicationJobRequest
	GetName() *string
	SetNetMode(v int32) *CreateReplicationJobRequest
	GetNetMode() *int32
	SetOwnerId(v int64) *CreateReplicationJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateReplicationJobRequest
	GetRegionId() *string
	SetReplicationParameters(v string) *CreateReplicationJobRequest
	GetReplicationParameters() *string
	SetResourceGroupId(v string) *CreateReplicationJobRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateReplicationJobRequest
	GetResourceOwnerAccount() *string
	SetRunOnce(v bool) *CreateReplicationJobRequest
	GetRunOnce() *bool
	SetScheduledStartTime(v string) *CreateReplicationJobRequest
	GetScheduledStartTime() *string
	SetSourceId(v string) *CreateReplicationJobRequest
	GetSourceId() *string
	SetSystemDiskPart(v []*CreateReplicationJobRequestSystemDiskPart) *CreateReplicationJobRequest
	GetSystemDiskPart() []*CreateReplicationJobRequestSystemDiskPart
	SetSystemDiskSize(v int32) *CreateReplicationJobRequest
	GetSystemDiskSize() *int32
	SetTag(v []*CreateReplicationJobRequestTag) *CreateReplicationJobRequest
	GetTag() []*CreateReplicationJobRequestTag
	SetTargetType(v string) *CreateReplicationJobRequest
	GetTargetType() *string
	SetVSwitchId(v string) *CreateReplicationJobRequest
	GetVSwitchId() *string
	SetValidTime(v string) *CreateReplicationJobRequest
	GetValidTime() *string
	SetVpcId(v string) *CreateReplicationJobRequest
	GetVpcId() *string
}

type CreateReplicationJobRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The namespace of the destination Docker container image. For more information about Docker container images, see [Terms](https://help.aliyun.com/document_detail/60744.html).
	//
	// example:
	//
	// testNamespace
	ContainerNamespace *string `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	// The repository that stores the destination Docker container image. For more information about Docker container images, see [Terms](https://help.aliyun.com/document_detail/60744.html).
	//
	// example:
	//
	// testRepository
	ContainerRepository *string `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	// The tag of the destination Docker container image. For more information about Docker container images, see [Terms](https://help.aliyun.com/document_detail/60744.html).
	//
	// example:
	//
	// CentOS:v1
	ContainerTag *string `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	// The data disks.
	DataDisk []*CreateReplicationJobRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The description of the migration job.
	//
	// The description must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (_), and hyphens (-). The description must start with a letter, but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This_is_a_migration_task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the disk.
	Disks *CreateReplicationJobRequestDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The interval at which SMC synchronizes incremental data to Alibaba Cloud. Unit: hour. Valid values: 1 to 168.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// 12
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The name of the destination image. The name must meet the following requirements:
	//
	// 	- The name must be unique within an Alibaba Cloud region.
	//
	// 	- The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with `http://` or `https://`.
	//
	// >  If you specify an image name that already exists in the destination region, the migration job ID is appended to the image name as a suffix. Example: ImageName_j-2zexxxxxxxxxxxxx.
	//
	// example:
	//
	// testAliCloudImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the destination ECS instance.
	//
	// example:
	//
	// i-bp1f1dvfto1sigz5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is assigned to the instance.
	//
	// example:
	//
	// SMCAdmin
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The type of the intermediate instance.
	//
	// You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to obtain the ECS instance types.
	//
	// 	- If you specify this parameter, SMC creates an intermediate instance of the specified instance type. If the specified instance type is unavailable, you cannot create the migration job.
	//
	// 	- If you do not specify this parameter, SMC selects an available instance type in a specific order to create an intermediate instance. For more information, see [SMC FAQ](https://help.aliyun.com/document_detail/121707.html).
	//
	// example:
	//
	// ecs.c6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The type of the migration job. Valid values:
	//
	// 	- 0: server migration.
	//
	// 	- 1: operating system migration.
	//
	// 	- 2: cross-zone migration.
	//
	// 	- 3: agentless migration for a VMware VM.
	//
	// example:
	//
	// 0
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the launch template.
	//
	// example:
	//
	// lt-bp16jovvln1cgaaq****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version number of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The license type. Valid values:
	//
	// 	- An empty value specifies no license.
	//
	// 	- A value of BYOL specifies Bring Your Own License (BYOL).
	//
	// For more information, see [SMC FAQ](https://help.aliyun.com/document_detail/121707.html).
	//
	// example:
	//
	// BYOL
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The maximum number of images retained for the incremental migration job. Valid values: 1 to 10.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// 10
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job. The name must meet the following requirements:
	//
	// 	- The name must be unique.
	//
	// 	- The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter, but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testMigrationTaskName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network mode for data transmission. Valid values:
	//
	// 	- 0: Data is transmitted over the Internet. Make sure that the source server can access the Internet.
	//
	// 	- 2: Data is transmitted over a VPC. If you specify this value, you must specify the VSwitchId parameter. You do not need to specify the VpcId parameter because the value of the VpcId parameter can be retrieved based on the value of the VSwitchId parameter.
	//
	// Default value: 0
	//
	// example:
	//
	// 0
	NetMode *int32 `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the Alibaba Cloud region to which you want to migrate the source server.
	//
	// For example, if you want to migrate the source server to the China (Hangzhou) region, set this parameter to `cn-hangzhou`. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameters of the replication driver. The parameters must be specified as key-value pairs in the JSON format. The keys are fixed for each type of replication driver. The JSON string can be up to 2,048 characters in length.
	//
	// A replication driver is a tool that is used to migrate a source server to an intermediate instance. The parameters vary based on the replication driver type. If you use a Server Migration Tool (SMT) driver, you can specify the following parameters:
	//
	// 	- bandwidth_limit: the maximum bandwidth for data transmission.
	//
	// 	- compress_level: the compression ratio of data to be transmitted.
	//
	// 	- checksum: specifies whether to enable checksum verification.
	//
	// For more information about replication drivers, see the response parameter `SourceServers.ReplicationDriver` of the [DescribeSourceServers](https://help.aliyun.com/document_detail/121818.html) operation.
	//
	// example:
	//
	// {"bandwidth_limit":0,"compress_level":1,"checksum":true}
	ReplicationParameters *string `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// Specifies whether to disable incremental migration for the source server. Valid values:
	//
	// 	- true: creates a migration job that runs only once. This is the default value. Incremental data of the source server is not synchronized.
	//
	// 	- false: creates an incremental migration job. In this case, you must specify the `Frequency` parameter. SMC synchronizes incremental data of the source server to Alibaba Cloud at the specified frequency. You can use an incremental migration job to synchronize incremental data from the source server to Alibaba Cloud without the need to interrupt your business. A full data image is generated for the source server when the job is running.
	//
	// >  You can specify this parameter only when you create a migration job. The parameter value cannot be changed after the migration job is created.
	//
	// example:
	//
	// true
	RunOnce *bool `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	// The time when you want to run the migration job. The time must meet the following requirements:
	//
	// 	- The time must be specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. For example, 2018-01-01T12:00:00Z specifies 20:00:00 on January 1, 2018 (UTC+8).
	//
	// 	- The value must be within 30 days after the current time.
	//
	// >  If you do not specify this parameter, you must manually start the migration job after the job is created. You can call the [StartReplicationJob](https://help.aliyun.com/document_detail/121823.html) operation to start the migration job.
	//
	// example:
	//
	// 2019-06-04T13:35:00Z
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The ID of the source server.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp1e2fsl57knvuug****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The information about system disk partitions.
	SystemDiskPart []*CreateReplicationJobRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	// The system disk size of the destination ECS instance. Unit: GiB. Valid values: 20 to 2048.
	//
	// >  The value must be greater than the used space of the system disk on the source server. For example, if the total size of the source disk is 500 GiB and the used space is 100 GiB, the value of this parameter must be greater than 100 GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags.
	Tag []*CreateReplicationJobRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of destination to which you want to migrate the source server. Valid values:
	//
	// 	- Image: After the migration job is complete, SMC generates an Elastic Compute Service (ECS) image for the source server.
	//
	// 	- ContainerImage: After the migration job is complete, SMC generates a Docker container image for the source server.
	//
	// 	- TargetInstance: After the migration job is completed, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the `InstanceId` parameter.
	//
	// example:
	//
	// Image
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	//
	// You must set this parameter if you use a VPC to migrate data.
	//
	// example:
	//
	// vsw-bp1ddbrxdlrcbim46****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The time when the migration job expires. You can schedule the migration job to expire 7 to 90 days after the job is created.
	//
	// 	- The time must be specified in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. For example, 2018-01-01T12:00:00Z specifies 20:00:00 on January 1, 2018 (UTC+8).
	//
	// 	- If you do not specify this parameter, the migration job does not expire.
	//
	// 	- After a migration job expires, the job state changes to Expired. SMC retains the migration job for seven days after the job expires. After the job is retained for seven days, SMC deletes the migration job.
	//
	// By default, a migration job is valid for 30 days after it is created.
	//
	// example:
	//
	// 2019-06-04T13:35:00Z
	ValidTime *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	// The ID of a VPC for which you have configured an Express Connect circuit or a VPN gateway.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateReplicationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateReplicationJobRequest) GetContainerNamespace() *string {
	return s.ContainerNamespace
}

func (s *CreateReplicationJobRequest) GetContainerRepository() *string {
	return s.ContainerRepository
}

func (s *CreateReplicationJobRequest) GetContainerTag() *string {
	return s.ContainerTag
}

func (s *CreateReplicationJobRequest) GetDataDisk() []*CreateReplicationJobRequestDataDisk {
	return s.DataDisk
}

func (s *CreateReplicationJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateReplicationJobRequest) GetDisks() *CreateReplicationJobRequestDisks {
	return s.Disks
}

func (s *CreateReplicationJobRequest) GetFrequency() *int32 {
	return s.Frequency
}

func (s *CreateReplicationJobRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateReplicationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateReplicationJobRequest) GetInstanceRamRole() *string {
	return s.InstanceRamRole
}

func (s *CreateReplicationJobRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateReplicationJobRequest) GetJobType() *int32 {
	return s.JobType
}

func (s *CreateReplicationJobRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateReplicationJobRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *CreateReplicationJobRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *CreateReplicationJobRequest) GetMaxNumberOfImageToKeep() *int32 {
	return s.MaxNumberOfImageToKeep
}

func (s *CreateReplicationJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateReplicationJobRequest) GetNetMode() *int32 {
	return s.NetMode
}

func (s *CreateReplicationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateReplicationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateReplicationJobRequest) GetReplicationParameters() *string {
	return s.ReplicationParameters
}

func (s *CreateReplicationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateReplicationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateReplicationJobRequest) GetRunOnce() *bool {
	return s.RunOnce
}

func (s *CreateReplicationJobRequest) GetScheduledStartTime() *string {
	return s.ScheduledStartTime
}

func (s *CreateReplicationJobRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateReplicationJobRequest) GetSystemDiskPart() []*CreateReplicationJobRequestSystemDiskPart {
	return s.SystemDiskPart
}

func (s *CreateReplicationJobRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateReplicationJobRequest) GetTag() []*CreateReplicationJobRequestTag {
	return s.Tag
}

func (s *CreateReplicationJobRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateReplicationJobRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateReplicationJobRequest) GetValidTime() *string {
	return s.ValidTime
}

func (s *CreateReplicationJobRequest) GetVpcId() *string {
	return s.VpcId
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

func (s *CreateReplicationJobRequest) SetDisks(v *CreateReplicationJobRequestDisks) *CreateReplicationJobRequest {
	s.Disks = v
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

func (s *CreateReplicationJobRequest) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDataDisk struct {
	// The index of data disk N on the destination ECS instance. Data disks on a destination ECS instance are arranged in a sequential order that starts from 1. Valid values: 1 to 16.
	//
	// >  To create a destination data disk for a source server, make sure that the source server has data disks.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The data disk partitions.
	Part []*CreateReplicationJobRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the data disk on the destination ECS instance. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The size of a destination data disk must be larger than the size of data in the source data disk. For example, if the size of the source data disk is 500 GiB and the used space is 100 GiB, you must set this parameter to a value greater than 100.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateReplicationJobRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDataDisk) GetIndex() *int32 {
	return s.Index
}

func (s *CreateReplicationJobRequestDataDisk) GetPart() []*CreateReplicationJobRequestDataDiskPart {
	return s.Part
}

func (s *CreateReplicationJobRequestDataDisk) GetSize() *int32 {
	return s.Size
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

func (s *CreateReplicationJobRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDataDiskPart struct {
	// Specifies whether to enable block replication for partition N in the destination data disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of partition N in the destination data disk. The partitions in the destination data disk are arranged in the same sequential order as those in the source data disk.
	//
	// >  You must set both the DataDisk.N.Part.N.Device and `DataDisk.N.Part.N.SizeBytes` parameters or leave both parameters empty.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination data disk. Unit: bytes. The default value is equal to the corresponding partition size of the source data disk.
	//
	// >
	//
	// 	- The total size of all partitions in a destination data disk cannot exceed the size of the destination data disk.
	//
	// 	- You must set both the `DataDisk.N.Part.N.Device` and DataDisk.N.Part.N.SizeBytes parameters or leave both parameters empty.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestDataDiskPart) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDataDiskPart) GetBlock() *bool {
	return s.Block
}

func (s *CreateReplicationJobRequestDataDiskPart) GetDevice() *string {
	return s.Device
}

func (s *CreateReplicationJobRequestDataDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *CreateReplicationJobRequestDataDiskPart) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDisks struct {
	// The information about the data disk partition.
	Data []*CreateReplicationJobRequestDisksData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The information about the system disk.
	System *CreateReplicationJobRequestDisksSystem `json:"System,omitempty" xml:"System,omitempty" type:"Struct"`
}

func (s CreateReplicationJobRequestDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDisks) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDisks) GetData() []*CreateReplicationJobRequestDisksData {
	return s.Data
}

func (s *CreateReplicationJobRequestDisks) GetSystem() *CreateReplicationJobRequestDisksSystem {
	return s.System
}

func (s *CreateReplicationJobRequestDisks) SetData(v []*CreateReplicationJobRequestDisksData) *CreateReplicationJobRequestDisks {
	s.Data = v
	return s
}

func (s *CreateReplicationJobRequestDisks) SetSystem(v *CreateReplicationJobRequestDisksSystem) *CreateReplicationJobRequestDisks {
	s.System = v
	return s
}

func (s *CreateReplicationJobRequestDisks) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDisksData struct {
	// The ID of the data disk.
	//
	// example:
	//
	// d-2ze8hyowhdgd6ou2m5z6
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether the data disk uses LVM. Valid values:
	//
	// 	- true: Use LVM.
	//
	// 	- false: Not use LVM.
	LVM *bool `json:"LVM,omitempty" xml:"LVM,omitempty"`
	// The information about the data disk partition.
	Part []*CreateReplicationJobRequestDisksDataPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the data disk of the migration source. Unit: GiB.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateReplicationJobRequestDisksData) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDisksData) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDisksData) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateReplicationJobRequestDisksData) GetLVM() *bool {
	return s.LVM
}

func (s *CreateReplicationJobRequestDisksData) GetPart() []*CreateReplicationJobRequestDisksDataPart {
	return s.Part
}

func (s *CreateReplicationJobRequestDisksData) GetSize() *int32 {
	return s.Size
}

func (s *CreateReplicationJobRequestDisksData) SetDiskId(v string) *CreateReplicationJobRequestDisksData {
	s.DiskId = &v
	return s
}

func (s *CreateReplicationJobRequestDisksData) SetLVM(v bool) *CreateReplicationJobRequestDisksData {
	s.LVM = &v
	return s
}

func (s *CreateReplicationJobRequestDisksData) SetPart(v []*CreateReplicationJobRequestDisksDataPart) *CreateReplicationJobRequestDisksData {
	s.Part = v
	return s
}

func (s *CreateReplicationJobRequestDisksData) SetSize(v int32) *CreateReplicationJobRequestDisksData {
	s.Size = &v
	return s
}

func (s *CreateReplicationJobRequestDisksData) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDisksDataPart struct {
	// Whether block replication is enabled for the data disk partition. Valid values:
	//
	// 	- true: Block replication is enabled for the data disk partition.
	//
	// 	- false: Block replication is disabled for the data disk partition.
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The path of the data disk partition.
	//
	// example:
	//
	// /home/date
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestDisksDataPart) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDisksDataPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDisksDataPart) GetBlock() *bool {
	return s.Block
}

func (s *CreateReplicationJobRequestDisksDataPart) GetPath() *string {
	return s.Path
}

func (s *CreateReplicationJobRequestDisksDataPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateReplicationJobRequestDisksDataPart) SetBlock(v bool) *CreateReplicationJobRequestDisksDataPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestDisksDataPart) SetPath(v string) *CreateReplicationJobRequestDisksDataPart {
	s.Path = &v
	return s
}

func (s *CreateReplicationJobRequestDisksDataPart) SetSizeBytes(v int64) *CreateReplicationJobRequestDisksDataPart {
	s.SizeBytes = &v
	return s
}

func (s *CreateReplicationJobRequestDisksDataPart) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDisksSystem struct {
	// Specifies whether to use LVM. Valid values:
	//
	// 	- true: Use LVM.
	//
	// 	- false: Not use LVM.
	//
	// LVM is not supported:
	//
	// 	- If your source server runs Windows, LVM is not supported.
	//
	// 	- The system disk does not have a boot partition, and LVM is not supported.
	//
	// After LVM is enabled, this feature does not take effect in the following scenarios:
	//
	// 	- LVM2 is not supported on your source server and the software package is not installed.
	//
	// 	- Your source server runs Debian with a kernel version of 3.x or earlier and XFS file systems are mounted.
	//
	// example:
	//
	// true
	LVM *bool `json:"LVM,omitempty" xml:"LVM,omitempty"`
	// The information about the system disk partition.
	Part []*CreateReplicationJobRequestDisksSystemPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the source system disk. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The parameter value must be greater than the actual used space of the data disk on the source server. For example, if the size of the source disk is 500 GiB but the actual used space is 100 GiB, you must set this parameter to a value greater than 100 GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateReplicationJobRequestDisksSystem) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDisksSystem) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDisksSystem) GetLVM() *bool {
	return s.LVM
}

func (s *CreateReplicationJobRequestDisksSystem) GetPart() []*CreateReplicationJobRequestDisksSystemPart {
	return s.Part
}

func (s *CreateReplicationJobRequestDisksSystem) GetSize() *int32 {
	return s.Size
}

func (s *CreateReplicationJobRequestDisksSystem) SetLVM(v bool) *CreateReplicationJobRequestDisksSystem {
	s.LVM = &v
	return s
}

func (s *CreateReplicationJobRequestDisksSystem) SetPart(v []*CreateReplicationJobRequestDisksSystemPart) *CreateReplicationJobRequestDisksSystem {
	s.Part = v
	return s
}

func (s *CreateReplicationJobRequestDisksSystem) SetSize(v int32) *CreateReplicationJobRequestDisksSystem {
	s.Size = &v
	return s
}

func (s *CreateReplicationJobRequestDisksSystem) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestDisksSystemPart struct {
	// Specifies whether block replication is enabled for the system disk partition.
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The path of the system disk partition.
	//
	// example:
	//
	// /boot
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the system disk partition. Unit: bytes.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestDisksSystemPart) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestDisksSystemPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestDisksSystemPart) GetBlock() *bool {
	return s.Block
}

func (s *CreateReplicationJobRequestDisksSystemPart) GetPath() *string {
	return s.Path
}

func (s *CreateReplicationJobRequestDisksSystemPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateReplicationJobRequestDisksSystemPart) SetBlock(v bool) *CreateReplicationJobRequestDisksSystemPart {
	s.Block = &v
	return s
}

func (s *CreateReplicationJobRequestDisksSystemPart) SetPath(v string) *CreateReplicationJobRequestDisksSystemPart {
	s.Path = &v
	return s
}

func (s *CreateReplicationJobRequestDisksSystemPart) SetSizeBytes(v int64) *CreateReplicationJobRequestDisksSystemPart {
	s.SizeBytes = &v
	return s
}

func (s *CreateReplicationJobRequestDisksSystemPart) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestSystemDiskPart struct {
	// Specifies whether to enable block replication for partition N in the destination system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination system disk. The partitions in the destination system disk are arranged in the same sequential order as those in the source system disk.
	//
	// >  You must set both the SystemDiskPart.N.Device and `SystemDiskPart.N.SizeBytes` parameters or leave both parameters empty.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the partition N in the destination system disk. Unit: bytes. The default value is equal to the partition size of the source system disk.
	//
	// >
	//
	// 	- The total size of all partitions in the destination system disk cannot exceed the size of the destination system disk.
	//
	// 	- You must set both the `SystemDiskPart.N.Device` and SystemDiskPart.N.SizeBytes parameters or leave both parameters empty.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateReplicationJobRequestSystemDiskPart) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestSystemDiskPart) GetBlock() *bool {
	return s.Block
}

func (s *CreateReplicationJobRequestSystemDiskPart) GetDevice() *string {
	return s.Device
}

func (s *CreateReplicationJobRequestSystemDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *CreateReplicationJobRequestSystemDiskPart) Validate() error {
	return dara.Validate(s)
}

type CreateReplicationJobRequestTag struct {
	// The key of the tag for the migration job. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag for the migration job. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateReplicationJobRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationJobRequestTag) GoString() string {
	return s.String()
}

func (s *CreateReplicationJobRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateReplicationJobRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateReplicationJobRequestTag) SetKey(v string) *CreateReplicationJobRequestTag {
	s.Key = &v
	return s
}

func (s *CreateReplicationJobRequestTag) SetValue(v string) *CreateReplicationJobRequestTag {
	s.Value = &v
	return s
}

func (s *CreateReplicationJobRequestTag) Validate() error {
	return dara.Validate(s)
}
