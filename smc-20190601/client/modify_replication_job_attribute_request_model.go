// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReplicationJobAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerNamespace(v string) *ModifyReplicationJobAttributeRequest
	GetContainerNamespace() *string
	SetContainerRepository(v string) *ModifyReplicationJobAttributeRequest
	GetContainerRepository() *string
	SetContainerTag(v string) *ModifyReplicationJobAttributeRequest
	GetContainerTag() *string
	SetDataDisk(v []*ModifyReplicationJobAttributeRequestDataDisk) *ModifyReplicationJobAttributeRequest
	GetDataDisk() []*ModifyReplicationJobAttributeRequestDataDisk
	SetDescription(v string) *ModifyReplicationJobAttributeRequest
	GetDescription() *string
	SetFrequency(v int32) *ModifyReplicationJobAttributeRequest
	GetFrequency() *int32
	SetImageName(v string) *ModifyReplicationJobAttributeRequest
	GetImageName() *string
	SetInstanceId(v string) *ModifyReplicationJobAttributeRequest
	GetInstanceId() *string
	SetInstanceRamRole(v string) *ModifyReplicationJobAttributeRequest
	GetInstanceRamRole() *string
	SetInstanceType(v string) *ModifyReplicationJobAttributeRequest
	GetInstanceType() *string
	SetJobId(v string) *ModifyReplicationJobAttributeRequest
	GetJobId() *string
	SetLaunchTemplateId(v string) *ModifyReplicationJobAttributeRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersion(v string) *ModifyReplicationJobAttributeRequest
	GetLaunchTemplateVersion() *string
	SetMaxNumberOfImageToKeep(v int32) *ModifyReplicationJobAttributeRequest
	GetMaxNumberOfImageToKeep() *int32
	SetName(v string) *ModifyReplicationJobAttributeRequest
	GetName() *string
	SetNetMode(v int32) *ModifyReplicationJobAttributeRequest
	GetNetMode() *int32
	SetOwnerId(v int64) *ModifyReplicationJobAttributeRequest
	GetOwnerId() *int64
	SetReplicationParameters(v string) *ModifyReplicationJobAttributeRequest
	GetReplicationParameters() *string
	SetResourceOwnerAccount(v string) *ModifyReplicationJobAttributeRequest
	GetResourceOwnerAccount() *string
	SetScheduledStartTime(v string) *ModifyReplicationJobAttributeRequest
	GetScheduledStartTime() *string
	SetSystemDiskPart(v []*ModifyReplicationJobAttributeRequestSystemDiskPart) *ModifyReplicationJobAttributeRequest
	GetSystemDiskPart() []*ModifyReplicationJobAttributeRequestSystemDiskPart
	SetSystemDiskSize(v int32) *ModifyReplicationJobAttributeRequest
	GetSystemDiskSize() *int32
	SetTargetType(v string) *ModifyReplicationJobAttributeRequest
	GetTargetType() *string
	SetVSwitchId(v string) *ModifyReplicationJobAttributeRequest
	GetVSwitchId() *string
	SetValidTime(v string) *ModifyReplicationJobAttributeRequest
	GetValidTime() *string
	SetVpcId(v string) *ModifyReplicationJobAttributeRequest
	GetVpcId() *string
}

type ModifyReplicationJobAttributeRequest struct {
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
	// The information about the data disk.
	DataDisk []*ModifyReplicationJobAttributeRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The description of the migration job.
	//
	// The description must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This_is_my_migration_task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The interval at which an incremental migration job runs. Unit: hour. Valid values: 1 to 168.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// example:
	//
	// 10
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The name of the destination image. The name must meet the following requirements:
	//
	// 	- The name must be unique within an Alibaba Cloud region.
	//
	// 	- The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// >  If an image whose name is the same as that of the destination image already exists in the current region when the migration job is in progress, the system adds the migration job ID to the end of the image name by default. Example: ImageName-JobId.
	//
	// example:
	//
	// testAliCloudImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// i-bp1f1dvfto1sigz5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is attached to the intermediate instance.
	//
	// example:
	//
	// SMCAdmin
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The type of the intermediate instance.
	//
	// You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the ECS instance types.
	//
	// 	- If you specify this parameter, SMC creates an intermediate instance of the specified instance type. If the specified instance type is unavailable, you cannot create the migration job.
	//
	// 	- If you do not specify this parameter, SMC selects an available instance type in a specific order to create an intermediate instance. For more information,
	//
	// see the "How does SMC create an intermediate instance?" section of the "FAQ" topic.
	//
	// example:
	//
	// ecs.c5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The migration job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// j-bp19vlwm0tyigbmj****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The launch template ID.
	//
	// example:
	//
	// lt-bp16jovvln1cgaaq****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version number of the launch template.
	//
	// example:
	//
	// Latest
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum number of images that are retained for an incremental migration job. Valid values: 1 to 10.
	//
	// This parameter is required if you set the `RunOnce` parameter to false.
	//
	// example:
	//
	// 5
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job. The name must meet the following requirements:
	//
	// 	- The name must be unique.
	//
	// 	- The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testMigrationTaskName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network mode for data transmission. Valid values:
	//
	// 	- 0: Data is transmitted over the Internet. Make sure that the source server can access the Internet.
	//
	// 	- 2: Data is transmitted over a virtual private cloud (VPC). If you set this parameter to 2, you must specify the VSwitchId parameter. You can leave the VpcId parameter empty, the VPC ID can be queried by calling an operation.
	//
	// example:
	//
	// 0
	NetMode *int32 `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameters of the replication driver. The parameters are fixed key-value pairs of the JSON format. The value can be up to 2,048 characters in length.
	//
	// A replication driver is a tool that is used to replicate the data of a source server to an intermediate instance. The parameters vary based on the replication driver type. If you use a Server Migration Tool (SMT) driver, you can set the following parameters:
	//
	// 	- bandwidth_limit: the maximum bandwidth for data transmission.
	//
	// 	- compress_level: the compression ratio of data to be transmitted.
	//
	// 	- checksum: specifies whether to enable checksum verification
	//
	// For more information about the replication driver, see the response parameter `SourceServers.ReplicationDriver` of the [DescribeSourceServers](https://help.aliyun.com/document_detail/2402126.html) operation.
	//
	// example:
	//
	// {"bandwidth_limit":0,"compress_level":1,"checksum":true}
	ReplicationParameters *string `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The time when the migration job is executed. SMC starts the migration job at the specified time.
	//
	// Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC. For example, 2018-01-01T12:00:00Z indicates 20:00:00 on January 1, 2018 (UTC+8).
	//
	// >  If ScheduledStartTime is left empty, SMC does not automatically start the migration job. In this case, you must call the [StartReplicationJob](https://help.aliyun.com/document_detail/121823.html) operation to start the migration job.
	//
	// example:
	//
	// 2019-06-04T13:35:00Z
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The partition information of the system disk.
	SystemDiskPart []*ModifyReplicationJobAttributeRequestSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
	// The system disk size of the destination ECS instance. Unit: GiB. Valid values: 20 to 500.
	//
	// >  The size of a destination data disk must be greater than the size of data in the source data disk. For example, if the source data disk has 500 GiB of storage space and 100 GiB of data, you must set this parameter to a value greater than 100.
	//
	// example:
	//
	// 50
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The type of destination to which the source server is migrated. You can modify the value only before the migration job starts. Valid values:
	//
	// 	- Image: After the migration job is complete, Server Migration Center (SMC) generates a destination Elastic Compute Service (ECS) image for the source server. You can use the image to create an ECS instance.
	//
	// 	- ContainerImage: After the migration job is complete, SMC generates a container image for the source server. You can use the container image in Container Registry.
	//
	// 	- TargetInstance: After the migration job is complete, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the `InstanceId` parameter.
	//
	// >
	//
	// 	- The value of this parameter is not case-sensitive.
	//
	// 	- SMC does not allow you to migrate Windows servers or servers that run operating systems on the ARM architecture to Container Registry.
	//
	// example:
	//
	// Image
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// example:
	//
	// vsw-bp1ddbrxdlrcbim46****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The time when the migration job expires. You can schedule the migration job to expire 7 to 90 days after the job is created.
	//
	// 	- This parameter can be modified only if the migration job is in the Ready, Running, Stopped, InError, or Waiting state.
	//
	// 	- Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC. For example, 2018-01-01T12:00:00Z indicates 20:00:00 on January 1, 2018 (UTC+8).
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
	// The ID of the VPC for which an Express Connect circuit or VPN gateway is configured.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyReplicationJobAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequest) GetContainerNamespace() *string {
	return s.ContainerNamespace
}

func (s *ModifyReplicationJobAttributeRequest) GetContainerRepository() *string {
	return s.ContainerRepository
}

func (s *ModifyReplicationJobAttributeRequest) GetContainerTag() *string {
	return s.ContainerTag
}

func (s *ModifyReplicationJobAttributeRequest) GetDataDisk() []*ModifyReplicationJobAttributeRequestDataDisk {
	return s.DataDisk
}

func (s *ModifyReplicationJobAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyReplicationJobAttributeRequest) GetFrequency() *int32 {
	return s.Frequency
}

func (s *ModifyReplicationJobAttributeRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyReplicationJobAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyReplicationJobAttributeRequest) GetInstanceRamRole() *string {
	return s.InstanceRamRole
}

func (s *ModifyReplicationJobAttributeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyReplicationJobAttributeRequest) GetJobId() *string {
	return s.JobId
}

func (s *ModifyReplicationJobAttributeRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *ModifyReplicationJobAttributeRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *ModifyReplicationJobAttributeRequest) GetMaxNumberOfImageToKeep() *int32 {
	return s.MaxNumberOfImageToKeep
}

func (s *ModifyReplicationJobAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyReplicationJobAttributeRequest) GetNetMode() *int32 {
	return s.NetMode
}

func (s *ModifyReplicationJobAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReplicationJobAttributeRequest) GetReplicationParameters() *string {
	return s.ReplicationParameters
}

func (s *ModifyReplicationJobAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReplicationJobAttributeRequest) GetScheduledStartTime() *string {
	return s.ScheduledStartTime
}

func (s *ModifyReplicationJobAttributeRequest) GetSystemDiskPart() []*ModifyReplicationJobAttributeRequestSystemDiskPart {
	return s.SystemDiskPart
}

func (s *ModifyReplicationJobAttributeRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ModifyReplicationJobAttributeRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ModifyReplicationJobAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyReplicationJobAttributeRequest) GetValidTime() *string {
	return s.ValidTime
}

func (s *ModifyReplicationJobAttributeRequest) GetVpcId() *string {
	return s.VpcId
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

func (s *ModifyReplicationJobAttributeRequest) SetNetMode(v int32) *ModifyReplicationJobAttributeRequest {
	s.NetMode = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetOwnerId(v int64) *ModifyReplicationJobAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetReplicationParameters(v string) *ModifyReplicationJobAttributeRequest {
	s.ReplicationParameters = &v
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

func (s *ModifyReplicationJobAttributeRequest) SetVSwitchId(v string) *ModifyReplicationJobAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetValidTime(v string) *ModifyReplicationJobAttributeRequest {
	s.ValidTime = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) SetVpcId(v string) *ModifyReplicationJobAttributeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyReplicationJobAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyReplicationJobAttributeRequestDataDisk struct {
	// The index of data disk N on the destination ECS instance. Valid values of N: 1 to 16.
	//
	// Data disks on a destination ECS instance are arranged in a sequential order that starts from 1.
	//
	// >  You can create a destination data disk only for a source server that has data disks.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The information about partitions.
	Part []*ModifyReplicationJobAttributeRequestDataDiskPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	// The size of the data disk on the destination ECS instance. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The size of a destination data disk must be greater than the size of data in the source data disk. For example, if the source data disk has 500 GiB of storage space and 100 GiB of data, you must set this parameter to a value greater than 100.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestDataDisk) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) GetIndex() *int32 {
	return s.Index
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) GetPart() []*ModifyReplicationJobAttributeRequestDataDiskPart {
	return s.Part
}

func (s *ModifyReplicationJobAttributeRequestDataDisk) GetSize() *int32 {
	return s.Size
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

func (s *ModifyReplicationJobAttributeRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type ModifyReplicationJobAttributeRequestDataDiskPart struct {
	// Specifies whether to enable block replication for partition N in the destination data disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination data disk.
	//
	// >  The partitions in the destination data disk are arranged in the same sequential order as those in the source data disk.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination data disk. Unit: bytes. The default value is equal to the corresponding size of the partition in the source data disk.
	//
	// >  The total size of all partitions in the destination data disk cannot exceed the size of the destination data disk.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestDataDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) GetBlock() *bool {
	return s.Block
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) GetDevice() *string {
	return s.Device
}

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *ModifyReplicationJobAttributeRequestDataDiskPart) Validate() error {
	return dara.Validate(s)
}

type ModifyReplicationJobAttributeRequestSystemDiskPart struct {
	// Specifies whether to enable block replication for partition N in the destination system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The ID of partition N in the destination system disk.
	//
	// >  The partitions in the destination system disk are arranged in the same sequential order as those in the source system disk.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of partition N in the destination system disk. Unit: bytes. The default value is equal to the partition size of the source system disk.
	//
	// >  The total size of all partitions in the destination system disk cannot exceed the size of the destination system disk.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) String() string {
	return dara.Prettify(s)
}

func (s ModifyReplicationJobAttributeRequestSystemDiskPart) GoString() string {
	return s.String()
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) GetBlock() *bool {
	return s.Block
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) GetDevice() *string {
	return s.Device
}

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *ModifyReplicationJobAttributeRequestSystemDiskPart) Validate() error {
	return dara.Validate(s)
}
