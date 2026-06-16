// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRefreshesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTasks(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) *DescribeInstanceRefreshesResponseBody
	GetInstanceRefreshTasks() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks
	SetMaxResults(v int32) *DescribeInstanceRefreshesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceRefreshesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeInstanceRefreshesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceRefreshesResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceRefreshesResponseBody struct {
	// The list of instance refresh tasks.
	InstanceRefreshTasks []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks `json:"InstanceRefreshTasks,omitempty" xml:"InstanceRefreshTasks,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instance refresh tasks.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBody) GetInstanceRefreshTasks() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	return s.InstanceRefreshTasks
}

func (s *DescribeInstanceRefreshesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceRefreshesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceRefreshesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRefreshesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceRefreshesResponseBody) SetInstanceRefreshTasks(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) *DescribeInstanceRefreshesResponseBody {
	s.InstanceRefreshTasks = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBody) SetMaxResults(v int32) *DescribeInstanceRefreshesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBody) SetNextToken(v string) *DescribeInstanceRefreshesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBody) SetRequestId(v string) *DescribeInstanceRefreshesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBody) SetTotalCount(v int32) *DescribeInstanceRefreshesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBody) Validate() error {
	if s.InstanceRefreshTasks != nil {
		for _, item := range s.InstanceRefreshTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks struct {
	// The duration for which the task pauses when a checkpoint is reached. Unit: minutes.
	//
	// example:
	//
	// 30
	CheckpointPauseTime *int32 `json:"CheckpointPauseTime,omitempty" xml:"CheckpointPauseTime,omitempty"`
	// The checkpoints for the refresh task. A checkpoint specifies that the task automatically pauses for CheckpointPauseTime minutes when the proportion of new instances reaches the specified value during the instance refresh.
	Checkpoints []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints `json:"Checkpoints,omitempty" xml:"Checkpoints,omitempty" type:"Repeated"`
	// The desired configuration for the instance refresh.
	DesiredConfiguration *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration `json:"DesiredConfiguration,omitempty" xml:"DesiredConfiguration,omitempty" type:"Struct"`
	// The failure reason when the instance refresh task fails.
	//
	// example:
	//
	// The task exceeded its maximum run time of one week. So the task failed.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the instance refresh task.
	//
	// example:
	//
	// 2024-08-22T02:09:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The capacity that has been refreshed.
	//
	// example:
	//
	// 10
	FinishedUpdateCapacity *int32 `json:"FinishedUpdateCapacity,omitempty" xml:"FinishedUpdateCapacity,omitempty"`
	// The ID of the instance refresh task.
	//
	// example:
	//
	// ir-1adfa123****
	InstanceRefreshTaskId *string `json:"InstanceRefreshTaskId,omitempty" xml:"InstanceRefreshTaskId,omitempty"`
	// The maximum percentage by which the number of instances in the scaling group can exceed the scaling group capacity during the instance refresh.
	//
	// example:
	//
	// 120
	MaxHealthyPercentage *int32 `json:"MaxHealthyPercentage,omitempty" xml:"MaxHealthyPercentage,omitempty"`
	// The minimum percentage of instances that must remain in service in the scaling group during the instance refresh.
	//
	// example:
	//
	// 80
	MinHealthyPercentage *int32 `json:"MinHealthyPercentage,omitempty" xml:"MinHealthyPercentage,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp16pbfcr8j9*****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// Indicates whether instances that already match the desired configuration are skipped.
	//
	// > The system determines whether an instance matches based on the ID of the desired scaling configuration, not by comparing individual configuration items.
	//
	// Valid values:
	//
	// - true: Skipped. When the instance refresh task starts, the system checks the configuration of each instance. Instances that were already created with the desired configuration are not refreshed.
	//
	// - false: Not skipped. After the instance refresh task starts, all instances in the scaling group are refreshed.
	//
	// example:
	//
	// true
	SkipMatching *bool `json:"SkipMatching,omitempty" xml:"SkipMatching,omitempty"`
	// The start time of the instance refresh task.
	//
	// example:
	//
	// 2024-08-22T01:09:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The current status of the instance refresh task. Valid values:
	//
	// - Pending: The instance refresh task is created and waiting to be scheduled.
	//
	// - InProgress: The instance refresh task is in progress.
	//
	// - Paused: The instance refresh task is paused.
	//
	// - CheckpointPause: The instance refresh task is paused because the task progress reached a checkpoint (`Checkpoint.Percentage`).
	//
	// - Failed: The instance refresh task failed.
	//
	// - Successful: The instance refresh task succeeded.
	//
	// - Cancelling: The instance refresh task is being canceled.
	//
	// - Cancelled: The instance refresh task is canceled.
	//
	// - RollbackInProgress: The instance refresh task is being rolled back.
	//
	// - RollbackSuccessful: The instance refresh task is rolled back.
	//
	// - RollbackFailed: The rollback of the instance refresh task failed.
	//
	// example:
	//
	// InProgress
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The total capacity that needs to be refreshed.
	//
	// example:
	//
	// 20
	TotalNeedUpdateCapacity *int32 `json:"TotalNeedUpdateCapacity,omitempty" xml:"TotalNeedUpdateCapacity,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetCheckpointPauseTime() *int32 {
	return s.CheckpointPauseTime
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetCheckpoints() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints {
	return s.Checkpoints
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetDesiredConfiguration() *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	return s.DesiredConfiguration
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetDetail() *string {
	return s.Detail
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetFinishedUpdateCapacity() *int32 {
	return s.FinishedUpdateCapacity
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetMaxHealthyPercentage() *int32 {
	return s.MaxHealthyPercentage
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetMinHealthyPercentage() *int32 {
	return s.MinHealthyPercentage
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetSkipMatching() *bool {
	return s.SkipMatching
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) GetTotalNeedUpdateCapacity() *int32 {
	return s.TotalNeedUpdateCapacity
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetCheckpointPauseTime(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.CheckpointPauseTime = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetCheckpoints(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.Checkpoints = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetDesiredConfiguration(v *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.DesiredConfiguration = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetDetail(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.Detail = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetEndTime(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetFinishedUpdateCapacity(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.FinishedUpdateCapacity = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetInstanceRefreshTaskId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetMaxHealthyPercentage(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.MaxHealthyPercentage = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetMinHealthyPercentage(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.MinHealthyPercentage = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetRegionId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetScalingGroupId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetSkipMatching(v bool) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.SkipMatching = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetStartTime(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetStatus(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.Status = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetStrategy(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.Strategy = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) SetTotalNeedUpdateCapacity(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks {
	s.TotalNeedUpdateCapacity = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasks) Validate() error {
	if s.Checkpoints != nil {
		for _, item := range s.Checkpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DesiredConfiguration != nil {
		if err := s.DesiredConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints struct {
	// The percentage of new instances relative to the total instances in the scaling group. The task automatically pauses when this percentage is reached.
	//
	// example:
	//
	// 60
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) SetPercentage(v int32) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints {
	s.Percentage = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksCheckpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration struct {
	// The list of containers included in the instance.
	Containers []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The ID of the image file used for automatic creation of instances.
	//
	// example:
	//
	// m-uf6g5noisr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the launch template from which the scaling group obtains launch configuration information.
	//
	// example:
	//
	// lt-2ze5x4mp*****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The instance type information that overrides the launch template.
	LaunchTemplateOverrides []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template. Valid values:
	//
	// - A fixed template version number.
	//
	// - Default: always uses the default version of the template.
	//
	// - Latest: always uses the latest version of the template.
	//
	// example:
	//
	// Latest
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-wz91ibkhfor****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetContainers() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	return s.Containers
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetLaunchTemplateOverrides() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetContainers(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.Containers = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetImageId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.ImageId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetLaunchTemplateId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetLaunchTemplateOverrides(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetLaunchTemplateVersion(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) SetScalingConfigurationId(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfiguration) Validate() error {
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LaunchTemplateOverrides != nil {
		for _, item := range s.LaunchTemplateOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers struct {
	// The arguments for the container startup commands.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container startup commands.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The environment variable information.
	EnvironmentVars []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The container image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The custom container name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GetArgs() []*string {
	return s.Args
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GetCommands() []*string {
	return s.Commands
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GetEnvironmentVars() []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) SetArgs(v []*string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	s.Args = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) SetCommands(v []*string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	s.Commands = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) SetEnvironmentVars(v []*DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) SetImage(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	s.Image = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) SetName(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers {
	s.Name = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainers) Validate() error {
	if s.EnvironmentVars != nil {
		for _, item := range s.EnvironmentVars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars struct {
	// >This parameter is not available for use.
	//
	// example:
	//
	// fieldPath
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) SetKey(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) SetValue(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationContainersEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides struct {
	// The instance type that overrides the instance type specified in the launch template.
	//
	// example:
	//
	// ecs.sn1ne.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) SetInstanceType(v string) *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceRefreshesResponseBodyInstanceRefreshTasksDesiredConfigurationLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}
