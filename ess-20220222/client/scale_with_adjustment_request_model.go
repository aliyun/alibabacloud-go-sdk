// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleWithAdjustmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityMetadata(v string) *ScaleWithAdjustmentRequest
	GetActivityMetadata() *string
	SetAdjustmentType(v string) *ScaleWithAdjustmentRequest
	GetAdjustmentType() *string
	SetAdjustmentValue(v int32) *ScaleWithAdjustmentRequest
	GetAdjustmentValue() *int32
	SetClientToken(v string) *ScaleWithAdjustmentRequest
	GetClientToken() *string
	SetExecutionMode(v string) *ScaleWithAdjustmentRequest
	GetExecutionMode() *string
	SetLifecycleHookContext(v *ScaleWithAdjustmentRequestLifecycleHookContext) *ScaleWithAdjustmentRequest
	GetLifecycleHookContext() *ScaleWithAdjustmentRequestLifecycleHookContext
	SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentRequest
	GetMinAdjustmentMagnitude() *int32
	SetOverrides(v *ScaleWithAdjustmentRequestOverrides) *ScaleWithAdjustmentRequest
	GetOverrides() *ScaleWithAdjustmentRequestOverrides
	SetOwnerId(v int64) *ScaleWithAdjustmentRequest
	GetOwnerId() *int64
	SetParallelTask(v bool) *ScaleWithAdjustmentRequest
	GetParallelTask() *bool
	SetResourceOwnerAccount(v string) *ScaleWithAdjustmentRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ScaleWithAdjustmentRequest
	GetScalingGroupId() *string
	SetSyncActivity(v bool) *ScaleWithAdjustmentRequest
	GetSyncActivity() *bool
}

type ScaleWithAdjustmentRequest struct {
	// The metadata of the scaling activity.
	//
	// example:
	//
	// {"key":"value"}
	ActivityMetadata *string `json:"ActivityMetadata,omitempty" xml:"ActivityMetadata,omitempty"`
	// The adjustment method for the scaling activity. Valid values:
	//
	// - QuantityChangeInCapacity: Increases or decreases a specified number of ECS instances.
	//
	// - PercentChangeInCapacity: Increases or decreases a specified percentage of ECS instances.
	//
	// - TotalCapacity: Adjusts the number of ECS instances in the current scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value for the scaling activity. In all cases, the number of ECS instances adjusted in a single operation cannot exceed 1,000. Otherwise, the adjustment fails. Valid values vary based on the adjustment method:
	//
	// - QuantityChangeInCapacity: -1,000 to 1,000.
	//
	// - PercentChangeInCapacity: -100 to 10,000.
	//
	// - TotalCapacity: 0 to 2,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// - None: Not specified. Normal scaling is performed.
	//
	// - PlanOnly: Does not trigger scaling. Only elastic planning is executed. The planning results, including instance types, zone IDs, billing methods, and the number of instances to be created, are returned in PlanResult.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The context information of the lifecycle hook.
	LifecycleHookContext *ScaleWithAdjustmentRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	// The minimum number of instances to be adjusted in the scaling activity. This parameter takes effect only when `AdjustmentType` is set to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The override parameters specified for scaling out an Elastic Container Instance (ECI) scaling group.
	Overrides *ScaleWithAdjustmentRequestOverrides `json:"Overrides,omitempty" xml:"Overrides,omitempty" type:"Struct"`
	OwnerId   *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether the current scaling task supports concurrent execution.
	//
	// example:
	//
	// false
	ParallelTask         *bool   `json:"ParallelTask,omitempty" xml:"ParallelTask,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-j6c1o397427hyjdc****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// Specifies whether to synchronously execute the scaling activity. This parameter is valid only for scaling groups that have the expected number of instances configured. Valid values:
	//
	// - true: Synchronous process. When the scaling rule is executed, the scale-out or scale-in activity of the scaling group is directly triggered.
	//
	// - false: Asynchronous process. When the expected number of instances in the scaling group is modified, the actual scale-out or scale-in activity of the scaling group is not immediately triggered. The system automatically detects the difference between the expected number of instances and the current number of instances in the scaling group, and then triggers the scale-out or scale-in activity.
	//
	// > For more information about the expected number of instances, see [Expected number of instances](https://help.aliyun.com/document_detail/146231.html).
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SyncActivity *bool `json:"SyncActivity,omitempty" xml:"SyncActivity,omitempty"`
}

func (s ScaleWithAdjustmentRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentRequest) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequest) GetActivityMetadata() *string {
	return s.ActivityMetadata
}

func (s *ScaleWithAdjustmentRequest) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *ScaleWithAdjustmentRequest) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ScaleWithAdjustmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ScaleWithAdjustmentRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *ScaleWithAdjustmentRequest) GetLifecycleHookContext() *ScaleWithAdjustmentRequestLifecycleHookContext {
	return s.LifecycleHookContext
}

func (s *ScaleWithAdjustmentRequest) GetMinAdjustmentMagnitude() *int32 {
	return s.MinAdjustmentMagnitude
}

func (s *ScaleWithAdjustmentRequest) GetOverrides() *ScaleWithAdjustmentRequestOverrides {
	return s.Overrides
}

func (s *ScaleWithAdjustmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ScaleWithAdjustmentRequest) GetParallelTask() *bool {
	return s.ParallelTask
}

func (s *ScaleWithAdjustmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ScaleWithAdjustmentRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ScaleWithAdjustmentRequest) GetSyncActivity() *bool {
	return s.SyncActivity
}

func (s *ScaleWithAdjustmentRequest) SetActivityMetadata(v string) *ScaleWithAdjustmentRequest {
	s.ActivityMetadata = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentType(v string) *ScaleWithAdjustmentRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetAdjustmentValue(v int32) *ScaleWithAdjustmentRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetClientToken(v string) *ScaleWithAdjustmentRequest {
	s.ClientToken = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetExecutionMode(v string) *ScaleWithAdjustmentRequest {
	s.ExecutionMode = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetLifecycleHookContext(v *ScaleWithAdjustmentRequestLifecycleHookContext) *ScaleWithAdjustmentRequest {
	s.LifecycleHookContext = v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetOverrides(v *ScaleWithAdjustmentRequestOverrides) *ScaleWithAdjustmentRequest {
	s.Overrides = v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetOwnerId(v int64) *ScaleWithAdjustmentRequest {
	s.OwnerId = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetParallelTask(v bool) *ScaleWithAdjustmentRequest {
	s.ParallelTask = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetResourceOwnerAccount(v string) *ScaleWithAdjustmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetScalingGroupId(v string) *ScaleWithAdjustmentRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) SetSyncActivity(v bool) *ScaleWithAdjustmentRequest {
	s.SyncActivity = &v
	return s
}

func (s *ScaleWithAdjustmentRequest) Validate() error {
	if s.LifecycleHookContext != nil {
		if err := s.LifecycleHookContext.Validate(); err != nil {
			return err
		}
	}
	if s.Overrides != nil {
		if err := s.Overrides.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ScaleWithAdjustmentRequestLifecycleHookContext struct {
	// Specifies whether to disable all lifecycle hooks for the scaling activity. Valid values:
	//
	// - true: Disables all lifecycle hooks.
	//
	// - false: Does not disable lifecycle hooks.
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// The list of IDs of the lifecycle hooks that are disabled for the scaling activity.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
	// The default execution result of the lifecycle hook in this scaling activity. If specified, this value overrides the original default execution result of the hook.
	//
	// example:
	//
	// CONTINUE
	LifecycleHookResult *string `json:"LifecycleHookResult,omitempty" xml:"LifecycleHookResult,omitempty"`
}

func (s ScaleWithAdjustmentRequestLifecycleHookContext) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentRequestLifecycleHookContext) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) GetDisableLifecycleHook() *bool {
	return s.DisableLifecycleHook
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) GetIgnoredLifecycleHookIds() []*string {
	return s.IgnoredLifecycleHookIds
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) GetLifecycleHookResult() *string {
	return s.LifecycleHookResult
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) SetDisableLifecycleHook(v bool) *ScaleWithAdjustmentRequestLifecycleHookContext {
	s.DisableLifecycleHook = &v
	return s
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) SetIgnoredLifecycleHookIds(v []*string) *ScaleWithAdjustmentRequestLifecycleHookContext {
	s.IgnoredLifecycleHookIds = v
	return s
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) SetLifecycleHookResult(v string) *ScaleWithAdjustmentRequestLifecycleHookContext {
	s.LifecycleHookResult = &v
	return s
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) Validate() error {
	return dara.Validate(s)
}

type ScaleWithAdjustmentRequestOverrides struct {
	// The override parameters for the container list.
	ContainerOverrides []*ScaleWithAdjustmentRequestOverridesContainerOverrides `json:"ContainerOverrides,omitempty" xml:"ContainerOverrides,omitempty" type:"Repeated"`
	// The number of vCPUs at the instance level. Unit: cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory size at the instance level. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The user data of the ECS instance. The user data must be encoded in Base64. The maximum size of the raw data before encoding is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ScaleWithAdjustmentRequestOverrides) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentRequestOverrides) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequestOverrides) GetContainerOverrides() []*ScaleWithAdjustmentRequestOverridesContainerOverrides {
	return s.ContainerOverrides
}

func (s *ScaleWithAdjustmentRequestOverrides) GetCpu() *float32 {
	return s.Cpu
}

func (s *ScaleWithAdjustmentRequestOverrides) GetMemory() *float32 {
	return s.Memory
}

func (s *ScaleWithAdjustmentRequestOverrides) GetUserData() *string {
	return s.UserData
}

func (s *ScaleWithAdjustmentRequestOverrides) SetContainerOverrides(v []*ScaleWithAdjustmentRequestOverridesContainerOverrides) *ScaleWithAdjustmentRequestOverrides {
	s.ContainerOverrides = v
	return s
}

func (s *ScaleWithAdjustmentRequestOverrides) SetCpu(v float32) *ScaleWithAdjustmentRequestOverrides {
	s.Cpu = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverrides) SetMemory(v float32) *ScaleWithAdjustmentRequestOverrides {
	s.Memory = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverrides) SetUserData(v string) *ScaleWithAdjustmentRequestOverrides {
	s.UserData = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverrides) Validate() error {
	if s.ContainerOverrides != nil {
		for _, item := range s.ContainerOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ScaleWithAdjustmentRequestOverridesContainerOverrides struct {
	// The arguments for the container startup command. A maximum of 10 arguments can be specified.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands for the container. A maximum of 20 commands can be specified. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs for the container. Unit: cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The list of environment variables.
	EnvironmentVars []*ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The memory size of the container. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container. When you need to override container parameters, you must specify the container name. The container parameters can be overridden only when the container name matches the container name in the scaling configuration.
	//
	// example:
	//
	// container-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ScaleWithAdjustmentRequestOverridesContainerOverrides) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentRequestOverridesContainerOverrides) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetArgs() []*string {
	return s.Args
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetCommands() []*string {
	return s.Commands
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetCpu() *float32 {
	return s.Cpu
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetEnvironmentVars() []*ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars {
	return s.EnvironmentVars
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetMemory() *float32 {
	return s.Memory
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) GetName() *string {
	return s.Name
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetArgs(v []*string) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.Args = v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetCommands(v []*string) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.Commands = v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetCpu(v float32) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.Cpu = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetEnvironmentVars(v []*ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.EnvironmentVars = v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetMemory(v float32) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.Memory = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) SetName(v string) *ScaleWithAdjustmentRequestOverridesContainerOverrides {
	s.Name = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverrides) Validate() error {
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

type ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars struct {
	// The name of the environment variable. The name must be 1 to 128 characters in length. It can contain letters, digits, and underscores (_), and cannot start with a digit.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be 0 to 256 characters in length.
	//
	// example:
	//
	// /usr/local/tomcat
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) SetKey(v string) *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars {
	s.Key = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) SetValue(v string) *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars {
	s.Value = &v
	return s
}

func (s *ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars) Validate() error {
	return dara.Validate(s)
}
