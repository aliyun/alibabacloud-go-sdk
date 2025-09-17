// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeReplicationJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReplicationJobsResponseBody
	GetPageSize() *int32
	SetReplicationJobs(v *DescribeReplicationJobsResponseBodyReplicationJobs) *DescribeReplicationJobsResponseBody
	GetReplicationJobs() *DescribeReplicationJobsResponseBodyReplicationJobs
	SetRequestId(v string) *DescribeReplicationJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeReplicationJobsResponseBody
	GetTotalCount() *int32
}

type DescribeReplicationJobsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The details of migration jobs.
	ReplicationJobs *DescribeReplicationJobsResponseBodyReplicationJobs `json:"ReplicationJobs,omitempty" xml:"ReplicationJobs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6E1187E8-843A-4850-B97E-2F17F00D48F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of migration jobs returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeReplicationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReplicationJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReplicationJobsResponseBody) GetReplicationJobs() *DescribeReplicationJobsResponseBodyReplicationJobs {
	return s.ReplicationJobs
}

func (s *DescribeReplicationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReplicationJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *DescribeReplicationJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobs struct {
	ReplicationJob []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob `json:"ReplicationJob,omitempty" xml:"ReplicationJob,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobs) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobs) GetReplicationJob() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	return s.ReplicationJob
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobs) SetReplicationJob(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) *DescribeReplicationJobsResponseBodyReplicationJobs {
	s.ReplicationJob = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobs) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob struct {
	// The business status of the migration job. Valid values:
	//
	// 	- Preparing: The migration is being prepared.
	//
	// 	- Syncing: Data is being synchronized.
	//
	// 	- Processing: The migration is in progress.
	//
	// 	- Cleaning: Intermediate resources are being released.
	//
	// example:
	//
	// Preparing
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The namespace of the destination Docker container image.
	//
	// example:
	//
	// testNamespace
	ContainerNamespace *string `json:"ContainerNamespace,omitempty" xml:"ContainerNamespace,omitempty"`
	// The repository that stores the destination Docker container image.
	//
	// example:
	//
	// testRepository
	ContainerRepository *string `json:"ContainerRepository,omitempty" xml:"ContainerRepository,omitempty"`
	// The tag of the destination Docker container image.
	//
	// example:
	//
	// CentOS:v1
	ContainerTag *string `json:"ContainerTag,omitempty" xml:"ContainerTag,omitempty"`
	// The time when the migration job was created.
	//
	// example:
	//
	// 2014-07-24T13:00:52Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data disks on the destination ECS instance.
	DataDisks *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	// The description of the migration job.
	//
	// example:
	//
	// This is my migration task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information about the disk.
	Disks *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The time when the migration job was complete. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	//
	// example:
	//
	// 2019-06-04T16:00:52Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code returned if an error occurred in the migration job.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The interval at which the incremental migration job runs. Unit: hour. Valid values: 1 to 168.
	//
	// example:
	//
	// 15
	Frequency *int32 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ID of the destination image.
	//
	// example:
	//
	// m-o6w3gy99qf89rkga****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the destination image.
	//
	// example:
	//
	// testAliCloudImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the destination ECS instance.
	//
	// example:
	//
	// i-bp1ff25rzvnul6kr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Resource Access Management (RAM) role that is assigned to the instance.
	//
	// example:
	//
	// SMCAdmin
	InstanceRamRole *string `json:"InstanceRamRole,omitempty" xml:"InstanceRamRole,omitempty"`
	// The instance type of the intermediate instance.
	//
	// example:
	//
	// ecs.sn1ne.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the migration job.
	//
	// example:
	//
	// j-bp19vlwm0tyigbmj****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
	// lt-launchtemplateid
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The versions of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The type of license for the migration job. Valid values:
	//
	// 	- An empty value indicates no license.
	//
	// 	- A value of BYOL indicates Bring Your Own License (BYOL).
	//
	// example:
	//
	// BYOL
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The maximum number of images retained for the incremental migration job. Valid values: 1 to 10.
	//
	// example:
	//
	// 8
	MaxNumberOfImageToKeep *int32 `json:"MaxNumberOfImageToKeep,omitempty" xml:"MaxNumberOfImageToKeep,omitempty"`
	// The name of the migration job.
	//
	// example:
	//
	// testMigrationTaskName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of network used for the migration.
	//
	// example:
	//
	// 0
	NetMode *int32 `json:"NetMode,omitempty" xml:"NetMode,omitempty"`
	// The progress of the migration job.
	//
	// example:
	//
	// 100
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the Alibaba Cloud region to which the source server is migrated.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution records of the migration job.
	ReplicationJobRuns *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns `json:"ReplicationJobRuns,omitempty" xml:"ReplicationJobRuns,omitempty" type:"Struct"`
	// The string of key-value pairs configured for the replication driver.
	//
	// example:
	//
	// BandWidthLimit:0
	ReplicationParameters *string `json:"ReplicationParameters,omitempty" xml:"ReplicationParameters,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether incremental migration is disabled for the source server. Valid values:
	//
	// 	- true: Incremental migration is disabled. A migration job runs only once after the job is created.
	//
	// 	- false: Incremental migration is enabled. For an incremental migration job, SMC synchronizes incremental data to Alibaba Cloud at the interval specified by the `Frequency` parameter.
	//
	// example:
	//
	// true
	RunOnce *bool `json:"RunOnce,omitempty" xml:"RunOnce,omitempty"`
	// The time when the migration job is scheduled to run. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. The time must meet the following requirements:
	//
	// 	- The value must be within 30 days after the current time.
	//
	// 	- If you do not specify this parameter, you must manually start the migration job after the migration job is created. You can call the [StartReplicationJob](https://help.aliyun.com/document_detail/121823.html) operation to start the migration job.
	//
	// example:
	//
	// 2019-06-04T13:35:00Z
	ScheduledStartTime *string `json:"ScheduledStartTime,omitempty" xml:"ScheduledStartTime,omitempty"`
	// The ID of the source server.
	//
	// example:
	//
	// s-bp1e2fsl57knvuug****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The time when the migration job was started. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	//
	// example:
	//
	// 2019-06-04T14:40:52Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the migration job. Valid values:
	//
	// 	- Ready: The migration job is not started.
	//
	// 	- Running: The migration job is running.
	//
	// 	- Stopped: The migration job is paused.
	//
	// 	- InError: An error occurs in the migration job.
	//
	// 	- Finished: The migration job is complete.
	//
	// 	- Waiting: The migration job is waiting to run.
	//
	// 	- Expired: The migration job has expired.
	//
	// 	- Deleting: The migration job is being deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information about the migration job.
	//
	// example:
	//
	// statusinfo
	StatusInfo *string `json:"StatusInfo,omitempty" xml:"StatusInfo,omitempty"`
	// The system disk partitions.
	SystemDiskParts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts `json:"SystemDiskParts,omitempty" xml:"SystemDiskParts,omitempty" type:"Struct"`
	// The size of the system disk of the destination ECS instance.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The information about tags that are attached to the SMC resource.
	Tags *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of destination to which the source server is migrated. Valid values:
	//
	// 	- Image: After the migration job is complete, SMC generates an ECS image for the source server.
	//
	// 	- ContainerImage: After the migration job is complete, SMC generates a Docker container image for the source server.
	//
	// 	- TargetInstance: After the migration job is complete, SMC migrates the source server to the destination instance. If you set this parameter to TargetInstance, you must set the InstanceId parameter.
	//
	// example:
	//
	// Image
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the intermediate instance.
	//
	// example:
	//
	// i-bp1ff25rzvnul6kr****
	TransitionInstanceId *string `json:"TransitionInstanceId,omitempty" xml:"TransitionInstanceId,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	//
	// example:
	//
	// vsw-bp1ddbrxdlrcbim46****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The time when the migration job expired. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	//
	// example:
	//
	// 2019-06-08T14:40:52Z
	ValidTime *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	// The ID of a virtual private cloud (VPC) for which you have configured an Express Connect circuit or a VPN gateway.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetContainerNamespace() *string {
	return s.ContainerNamespace
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetContainerRepository() *string {
	return s.ContainerRepository
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetContainerTag() *string {
	return s.ContainerTag
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetDataDisks() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks {
	return s.DataDisks
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetDescription() *string {
	return s.Description
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetDisks() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks {
	return s.Disks
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetFrequency() *int32 {
	return s.Frequency
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetInstanceRamRole() *string {
	return s.InstanceRamRole
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetJobId() *string {
	return s.JobId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetJobType() *int32 {
	return s.JobType
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetLicenseType() *string {
	return s.LicenseType
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetMaxNumberOfImageToKeep() *int32 {
	return s.MaxNumberOfImageToKeep
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetName() *string {
	return s.Name
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetNetMode() *int32 {
	return s.NetMode
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetReplicationJobRuns() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns {
	return s.ReplicationJobRuns
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetReplicationParameters() *string {
	return s.ReplicationParameters
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetRunOnce() *bool {
	return s.RunOnce
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetScheduledStartTime() *string {
	return s.ScheduledStartTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetStatusInfo() *string {
	return s.StatusInfo
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetSystemDiskParts() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts {
	return s.SystemDiskParts
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetTags() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags {
	return s.Tags
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetTransitionInstanceId() *string {
	return s.TransitionInstanceId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetValidTime() *string {
	return s.ValidTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) GetWorkgroupId() *string {
	return s.WorkgroupId
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetDisks(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.Disks = v
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) SetWorkgroupId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob {
	s.WorkgroupId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJob) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks struct {
	DataDisk []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) GetDataDisk() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk {
	return s.DataDisk
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) SetDataDisk(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks {
	s.DataDisk = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk struct {
	// The index number of the data disk.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The data disk partitions.
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) GetParts() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts {
	return s.Parts
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) GetSize() *int32 {
	return s.Size
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts struct {
	Part []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) GetPart() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart {
	return s.Part
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) SetPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts {
	s.Part = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskParts) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart struct {
	// Indicates whether block replication is enabled for the data disk partition.
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of the data disk partition.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GetBlock() *bool {
	return s.Block
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GetDevice() *string {
	return s.Device
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDataDisksDataDiskPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks struct {
	// The information about the data disk.
	Data *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The information about the system disk.
	System *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem `json:"System,omitempty" xml:"System,omitempty" type:"Struct"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) GetData() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData {
	return s.Data
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) GetSystem() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem {
	return s.System
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) SetData(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks {
	s.Data = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) SetSystem(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks {
	s.System = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData struct {
	Data []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) GetData() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData {
	return s.Data
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) SetData(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData {
	s.Data = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksData) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData struct {
	// The ID of the data disk.
	//
	// example:
	//
	// d-2zeh4twm100qskw7z41z
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to use LVM. Valid values:
	//
	// 	- true: Use LVM.
	//
	// 	- false: Not use LVM.
	//
	// example:
	//
	// false
	LVM *bool `json:"LVM,omitempty" xml:"LVM,omitempty"`
	// The information about the data disk partition.
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The size of a data disk on the destination ECS instance. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The size of a destination data disk must be larger than the size of data in the corresponding source data disk. For example, if the size of the source disk is 500 GiB but the actual used space is 100 GiB, you must set this parameter to a value greater than 100 GiB.
	//
	// example:
	//
	// 22548578304
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) GetLVM() *bool {
	return s.LVM
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) GetParts() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts {
	return s.Parts
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) GetSize() *int32 {
	return s.Size
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) SetDiskId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData {
	s.DiskId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) SetLVM(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData {
	s.LVM = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) SetParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData {
	s.Parts = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) SetSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData {
	s.Size = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataData) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts struct {
	Part []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) GetPart() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart {
	return s.Part
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) SetPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts {
	s.Part = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataParts) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart struct {
	// Whether block replication is enabled for the data disk partition. Valid values:
	//
	// 	- true: Block replication is enabled for the data disk partition.
	//
	// 	- false: Block replication is disabled for the data disk partition.
	//
	// example:
	//
	// false
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The path of the data disk partition.
	//
	// example:
	//
	// /home/data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The size of the data disk partition. Unit: bytes.
	//
	// example:
	//
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) GetBlock() *bool {
	return s.Block
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) GetPath() *string {
	return s.Path
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) SetPath(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksDataDataPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem struct {
	// The ID of the system disk.
	//
	// example:
	//
	// d-2zeh4twm100qskw7z41z
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to use LVM. Valid values:
	//
	// 	- true: Use LVM.
	//
	// 	- false: Not use LVM.
	//
	// example:
	//
	// false
	LVM *bool `json:"LVM,omitempty" xml:"LVM,omitempty"`
	// The information about the system disk partition.
	Parts *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Struct"`
	// The size of the source system disk. Unit: GiB. Valid values: 20 to 32768.
	//
	// >  The parameter value must be greater than the actual used space of the data disk on the source server. For example, if the size of the source disk is 500 GiB but the actual used space is 100 GiB, you must set this parameter to a value greater than 100 GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) GetLVM() *bool {
	return s.LVM
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) GetParts() *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts {
	return s.Parts
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) GetSize() *int32 {
	return s.Size
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) SetDiskId(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem {
	s.DiskId = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) SetLVM(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem {
	s.LVM = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) SetParts(v *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem {
	s.Parts = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) SetSize(v int32) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem {
	s.Size = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystem) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts struct {
	Part []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) GetPart() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart {
	return s.Part
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) SetPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts {
	s.Part = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemParts) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart struct {
	// Specifies whether block replication is enabled for the system disk partition. Valid values:
	//
	// 	- true: Block replication is enabled for the system disk partition.
	//
	// 	- false: Block replication is disabled for the system disk partition.
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
	// 21474836480
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) GetBlock() *bool {
	return s.Block
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) GetPath() *string {
	return s.Path
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) SetBlock(v bool) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart {
	s.Block = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) SetPath(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart {
	s.Path = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) SetSizeBytes(v int64) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart {
	s.SizeBytes = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobDisksSystemPartsPart) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns struct {
	ReplicationJobRun []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun `json:"ReplicationJobRun,omitempty" xml:"ReplicationJobRun,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) GetReplicationJobRun() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun {
	return s.ReplicationJobRun
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) SetReplicationJobRun(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns {
	s.ReplicationJobRun = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRuns) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun struct {
	// The time when the migration job ended. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	//
	// example:
	//
	// 2019-10-04T13:35:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the destination image.
	//
	// example:
	//
	// m-o6w3gy99qf89rkga****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The time when the migration job was started. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the SMC console is in the format of UTC+8.
	//
	// example:
	//
	// 2019-10-01T13:35:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The method used to run the migration job. Valid values:
	//
	// 	- Manual: The migration job was manually started.
	//
	// 	- Schedule: The migration job was started at a scheduled time or at a specific interval.
	//
	// example:
	//
	// Schedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) GetType() *string {
	return s.Type
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobReplicationJobRunsReplicationJobRun) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts struct {
	SystemDiskPart []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart `json:"SystemDiskPart,omitempty" xml:"SystemDiskPart,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) GetSystemDiskPart() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart {
	return s.SystemDiskPart
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) SetSystemDiskPart(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts {
	s.SystemDiskPart = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskParts) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart struct {
	// Indicates whether block replication is enabled for the system disk partition.
	//
	// example:
	//
	// true
	Block *bool `json:"Block,omitempty" xml:"Block,omitempty"`
	// The device ID of the system disk partition.
	//
	// example:
	//
	// 0_1
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of the system disk partition. Unit: bytes.
	//
	// example:
	//
	// 254803968
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GetBlock() *bool {
	return s.Block
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GetDevice() *string {
	return s.Device
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) GetSizeBytes() *int64 {
	return s.SizeBytes
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

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobSystemDiskPartsSystemDiskPart) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags struct {
	Tag []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) GetTag() []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag {
	return s.Tag
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) SetTag(v []*DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags {
	s.Tag = v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTags) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag struct {
	// The key of the tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) SetKey(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) SetValue(v string) *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeReplicationJobsResponseBodyReplicationJobsReplicationJobTagsTag) Validate() error {
	return dara.Validate(s)
}
