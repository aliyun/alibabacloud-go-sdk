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
}

type StartInstanceRefreshRequest struct {
	// The duration of the pause when the refresh task checkpoint is entered.
	//
	// 	- Unit: minutes
	//
	// 	- Valid values: 1 to 2880.
	//
	// 	- Default: 60.
	//
	// example:
	//
	// 10
	CheckpointPauseTime *int32 `json:"CheckpointPauseTime,omitempty" xml:"CheckpointPauseTime,omitempty"`
	// Refresh Task Checkpoint: specifies that the task is automatically suspended for CheckpointPauseTime minutes when the proportion of new instances reaches the specified value during instance refresh.
	Checkpoints []*StartInstanceRefreshRequestCheckpoints `json:"Checkpoints,omitempty" xml:"Checkpoints,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The desired configurations of the instance refresh task.
	//
	// >
	//
	// 	- ScalingConfigurationId, ImageId, LaunchTemplateId, and Containers cannot be set at the same time. If you do not specify this parameter, the scaling group is refreshed based on the configurations that are in effect.
	//
	// 	- After the instance refresh task is complete, the scaling group uses the scaling configuration specified by this parameter.
	DesiredConfiguration *StartInstanceRefreshRequestDesiredConfiguration `json:"DesiredConfiguration,omitempty" xml:"DesiredConfiguration,omitempty" type:"Struct"`
	// The ratio of instances that can exceed the upper limit of the scaling group capacity to all instances in the scaling group during instance refresh. Valid values: 100 to 200. Default value: 120.
	//
	// >  If you set MinHealthyPercentage and MaxHealthyPercentage to 100, Auto Scaling refreshes the configurations of one instance each time the instance refresh task starts.
	//
	// example:
	//
	// 100
	MaxHealthyPercentage *int32 `json:"MaxHealthyPercentage,omitempty" xml:"MaxHealthyPercentage,omitempty"`
	// The ratio of instances that are in the In Service state to all instances in the scaling group during instance refresh. Valid values: 0 to 100. Default value: 80.
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
	// Specifies whether to skip instances that match the desired scaling configuration.
	//
	// >  The system determines the match based on the ID of the desired scaling configuration rather than individual configuration items.
	//
	// Valid values:
	//
	// 	- true: skips instances that match the desired scaling configuration. When you initiate an instance refresh task, the system checks the configurations of all instances. The refresh operation is skipped for instances created based on the desired scaling configuration.
	//
	// 	- false: does not skip instances that match the desired scaling configuration. When an instance refresh task is initiated, all instances in the scaling group at the time of initiation are refreshed.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	SkipMatching *bool `json:"SkipMatching,omitempty" xml:"SkipMatching,omitempty"`
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
	// The percentage of new instances in the scaling group to the total number of instances. When this percentage is reached, the task is automatically suspended. Valid values: 1 to 100 (%).
	//
	// >  Requires a small to large setting, and the last progress percentage needs to be 100.
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
	// The containers in the elastic container instance.
	//
	// >
	//
	// 	- This parameter supports only scaling groups of the ECI type.
	//
	// 	- Only the containers in the scaling configuration list that are the same as those in the `Container.Name` are refreshed.
	Containers []*StartInstanceRefreshRequestDesiredConfigurationContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The image ID.
	//
	// >
	//
	// 	- After the instance refresh task is complete, the active scaling configuration uses the image specified by this parameter.
	//
	// 	- If the instance configuration source of the scaling group is a launch template, you cannot specify this parameter.
	//
	// example:
	//
	// m-2ze8cqacj7opnf***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the launch template that you want to enable in the scaling group.
	//
	// example:
	//
	// lt-2ze2qli30u***
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The information about the instance types that are extended in the launch template.
	LaunchTemplateOverrides []*StartInstanceRefreshRequestDesiredConfigurationLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version number of the launch template. Valid value:
	//
	// 	- A fixed template version number.
	//
	// 	- Default: the default version of the template.
	//
	// 	- Latest: the latest version of the template.
	//
	// >  If you set the version to Default or Latest, the instance refresh task cannot be rolled back.
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
	// The argument that corresponds to the startup command of the container. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container startup commands. You can specify up to 20 commands. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The environment variables.
	EnvironmentVars []*StartInstanceRefreshRequestDesiredConfigurationContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The image in the container.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The custom name of the container.
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
	// >  This parameter is unavailable for use.
	//
	// example:
	//
	// fieldPath
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. It can be 1 to 128 characters in length. Format requirement:[0-9a-zA-Z], and underscores, cannot start with a number.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value must be 0 to 256 bits in length.
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
	// The instance type specified by using this parameter overwrites the instance type of the launch template.
	//
	// >  This parameter takes effect only if you specify LaunchTemplateId.
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
