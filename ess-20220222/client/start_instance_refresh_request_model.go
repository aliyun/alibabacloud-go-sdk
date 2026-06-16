// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpointPauseTime(v int32) *StartInstanceRefreshRequest
	GetCheckpointPauseTime() *int32
	SetCheckpoints(v []*StartInstanceRefreshRequestCheckpoints) *StartInstanceRefreshRequest
	GetCheckpoints() []*StartInstanceRefreshRequestCheckpoints
	SetClientToken(v string) *StartInstanceRefreshRequest
	GetClientToken() *string
	SetDesiredConfiguration(v *StartInstanceRefreshRequestDesiredConfiguration) *StartInstanceRefreshRequest
	GetDesiredConfiguration() *StartInstanceRefreshRequestDesiredConfiguration
	SetMaxHealthyPercentage(v int32) *StartInstanceRefreshRequest
	GetMaxHealthyPercentage() *int32
	SetMinHealthyPercentage(v int32) *StartInstanceRefreshRequest
	GetMinHealthyPercentage() *int32
	SetOwnerId(v int64) *StartInstanceRefreshRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartInstanceRefreshRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartInstanceRefreshRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *StartInstanceRefreshRequest
	GetScalingGroupId() *string
	SetSkipMatching(v bool) *StartInstanceRefreshRequest
	GetSkipMatching() *bool
	SetStrategy(v string) *StartInstanceRefreshRequest
	GetStrategy() *string
}

type StartInstanceRefreshRequest struct {
	// The duration for which the task is paused when a checkpoint is reached.
	//
	// - Unit: minutes.
	//
	// - Valid values: 1 to 2880.
	//
	//  - Default value: 60.
	//
	// example:
	//
	// 10
	CheckpointPauseTime *int32 `json:"CheckpointPauseTime,omitempty" xml:"CheckpointPauseTime,omitempty"`
	// The checkpoints for the refresh task. When the percentage of new instances reaches a specified value during the instance refresh, the task is automatically paused for CheckpointPauseTime minutes.
	Checkpoints []*StartInstanceRefreshRequestCheckpoints `json:"Checkpoints,omitempty" xml:"Checkpoints,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The desired configuration for the instance refresh.
	//
	// > - You cannot specify ScalingConfigurationId, ImageId, LaunchTemplateId, and Containers at the same time. If this parameter is left empty, the currently active configuration of the scaling group is used for the refresh.
	//
	// > - After the instance refresh task is completed, the active scaling configuration of the scaling group is updated to this configuration.
	DesiredConfiguration *StartInstanceRefreshRequestDesiredConfiguration `json:"DesiredConfiguration,omitempty" xml:"DesiredConfiguration,omitempty" type:"Struct"`
	// The maximum percentage by which the number of instances in the scaling group can exceed the scaling group capacity during the instance refresh. Valid values: 100 to 200.
	//
	// Default value: 120.
	//
	// > When MinHealthyPercentage = MaxHealthyPercentage = 100, one instance is refreshed at a time.
	//
	// example:
	//
	// 100
	MaxHealthyPercentage *int32 `json:"MaxHealthyPercentage,omitempty" xml:"MaxHealthyPercentage,omitempty"`
	// The minimum percentage of instances that must remain in service in the scaling group during the instance refresh. Valid values: 0 to 100.
	//
	// Default value: 80.
	//
	// example:
	//
	// 80
	MinHealthyPercentage *int32 `json:"MinHealthyPercentage,omitempty" xml:"MinHealthyPercentage,omitempty"`
	OwnerId              *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// Specifies whether to skip instances that already match the desired configuration.
	//
	// > The system determines whether an instance matches based on the ID of the desired scaling configuration, not by comparing individual configuration items.
	//
	// Valid values:
	//
	// - true: Instances that were already created with the desired configuration are skipped.
	//
	// - false: All instances in the scaling group are refreshed when the instance refresh task starts.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	SkipMatching *bool   `json:"SkipMatching,omitempty" xml:"SkipMatching,omitempty"`
	Strategy     *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s StartInstanceRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequest) GetCheckpointPauseTime() *int32 {
	return s.CheckpointPauseTime
}

func (s *StartInstanceRefreshRequest) GetCheckpoints() []*StartInstanceRefreshRequestCheckpoints {
	return s.Checkpoints
}

func (s *StartInstanceRefreshRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartInstanceRefreshRequest) GetDesiredConfiguration() *StartInstanceRefreshRequestDesiredConfiguration {
	return s.DesiredConfiguration
}

func (s *StartInstanceRefreshRequest) GetMaxHealthyPercentage() *int32 {
	return s.MaxHealthyPercentage
}

func (s *StartInstanceRefreshRequest) GetMinHealthyPercentage() *int32 {
	return s.MinHealthyPercentage
}

func (s *StartInstanceRefreshRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartInstanceRefreshRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstanceRefreshRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartInstanceRefreshRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *StartInstanceRefreshRequest) GetSkipMatching() *bool {
	return s.SkipMatching
}

func (s *StartInstanceRefreshRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *StartInstanceRefreshRequest) SetCheckpointPauseTime(v int32) *StartInstanceRefreshRequest {
	s.CheckpointPauseTime = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetCheckpoints(v []*StartInstanceRefreshRequestCheckpoints) *StartInstanceRefreshRequest {
	s.Checkpoints = v
	return s
}

func (s *StartInstanceRefreshRequest) SetClientToken(v string) *StartInstanceRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetDesiredConfiguration(v *StartInstanceRefreshRequestDesiredConfiguration) *StartInstanceRefreshRequest {
	s.DesiredConfiguration = v
	return s
}

func (s *StartInstanceRefreshRequest) SetMaxHealthyPercentage(v int32) *StartInstanceRefreshRequest {
	s.MaxHealthyPercentage = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetMinHealthyPercentage(v int32) *StartInstanceRefreshRequest {
	s.MinHealthyPercentage = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetOwnerId(v int64) *StartInstanceRefreshRequest {
	s.OwnerId = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetRegionId(v string) *StartInstanceRefreshRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetResourceOwnerAccount(v string) *StartInstanceRefreshRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetScalingGroupId(v string) *StartInstanceRefreshRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetSkipMatching(v bool) *StartInstanceRefreshRequest {
	s.SkipMatching = &v
	return s
}

func (s *StartInstanceRefreshRequest) SetStrategy(v string) *StartInstanceRefreshRequest {
	s.Strategy = &v
	return s
}

func (s *StartInstanceRefreshRequest) Validate() error {
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

type StartInstanceRefreshRequestCheckpoints struct {
	// The percentage of new instances relative to the total number of instances in the scaling group. The task is automatically paused when this percentage is reached. Valid values: 1 to 100 (%).
	//
	// > The values must be specified in ascending order, and the last value must be 100.
	//
	// example:
	//
	// 20
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s StartInstanceRefreshRequestCheckpoints) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequestCheckpoints) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequestCheckpoints) GetPercentage() *int32 {
	return s.Percentage
}

func (s *StartInstanceRefreshRequestCheckpoints) SetPercentage(v int32) *StartInstanceRefreshRequestCheckpoints {
	s.Percentage = &v
	return s
}

func (s *StartInstanceRefreshRequestCheckpoints) Validate() error {
	return dara.Validate(s)
}

type StartInstanceRefreshRequestDesiredConfiguration struct {
	// The list of containers included in the instance.
	//
	// > - This parameter is supported only for Elastic Container Instance (ECI) scaling groups.
	//
	// > - Only the container configurations that match `Container.Name` in the current scaling configuration container list are refreshed.
	Containers []*StartInstanceRefreshRequestDesiredConfigurationContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The image ID.
	//
	//
	//
	// > - After the instance refresh task is completed, the image in the currently active configuration of the scaling group is updated to this image.
	//
	// > - This parameter is not supported when the instance configuration source of the scaling group is a launch template.
	//
	// example:
	//
	// m-2ze8cqacj7opnf***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the launch template from which the scaling group obtains launch configuration information.
	//
	// example:
	//
	// lt-2ze2qli30u***
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The instance type information that overrides the launch template.
	LaunchTemplateOverrides []*StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template. Valid values:
	//
	// - A fixed template version number.
	//
	// - Default: always uses the default version of the template.
	//
	// - Latest: always uses the latest version of the template.
	//
	// > When the version is set to Default or Latest, the instance refresh task does not support rollback.
	//
	// example:
	//
	// 8
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-2zed7lqn4ts4****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s StartInstanceRefreshRequestDesiredConfiguration) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequestDesiredConfiguration) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetContainers() []*StartInstanceRefreshRequestDesiredConfigurationContainers {
	return s.Containers
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetImageId() *string {
	return s.ImageId
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetLaunchTemplateOverrides() []*StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetContainers(v []*StartInstanceRefreshRequestDesiredConfigurationContainers) *StartInstanceRefreshRequestDesiredConfiguration {
	s.Containers = v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetImageId(v string) *StartInstanceRefreshRequestDesiredConfiguration {
	s.ImageId = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetLaunchTemplateId(v string) *StartInstanceRefreshRequestDesiredConfiguration {
	s.LaunchTemplateId = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetLaunchTemplateOverrides(v []*StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) *StartInstanceRefreshRequestDesiredConfiguration {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetLaunchTemplateVersion(v string) *StartInstanceRefreshRequestDesiredConfiguration {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) SetScalingConfigurationId(v string) *StartInstanceRefreshRequestDesiredConfiguration {
	s.ScalingConfigurationId = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfiguration) Validate() error {
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

type StartInstanceRefreshRequestDesiredConfigurationContainers struct {
	// The arguments of the container startup command. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands of the container. You can specify up to 20 commands. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The environment variable information.
	EnvironmentVars []*StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
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
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StartInstanceRefreshRequestDesiredConfigurationContainers) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequestDesiredConfigurationContainers) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) GetArgs() []*string {
	return s.Args
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) GetCommands() []*string {
	return s.Commands
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) GetEnvironmentVars() []*StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) GetImage() *string {
	return s.Image
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) GetName() *string {
	return s.Name
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) SetArgs(v []*string) *StartInstanceRefreshRequestDesiredConfigurationContainers {
	s.Args = v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) SetCommands(v []*string) *StartInstanceRefreshRequestDesiredConfigurationContainers {
	s.Commands = v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) SetEnvironmentVars(v []*StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) *StartInstanceRefreshRequestDesiredConfigurationContainers {
	s.EnvironmentVars = v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) SetImage(v string) *StartInstanceRefreshRequestDesiredConfigurationContainers {
	s.Image = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) SetName(v string) *StartInstanceRefreshRequestDesiredConfigurationContainers {
	s.Name = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainers) Validate() error {
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

type StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars struct {
	// > This parameter is not available for use.
	//
	// example:
	//
	// fieldPath
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. The name must be 1 to 128 characters in length and can contain digits, letters, and underscores (_). It cannot start with a digit.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be 0 to 256 characters in length.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) SetFieldRefFieldPath(v string) *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) SetKey(v string) *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) SetValue(v string) *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides struct {
	// The instance type that overrides the instance type specified in the launch template.
	//
	// > This parameter takes effect only when the LaunchTemplateId parameter specifies a launch template.
	//
	// example:
	//
	// ecs.c5.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) SetInstanceType(v string) *StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}
