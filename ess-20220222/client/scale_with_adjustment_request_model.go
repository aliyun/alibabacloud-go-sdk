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
	// The metadata for the scaling activity.
	//
	// example:
	//
	// {"key":"value"}
	ActivityMetadata *string `json:"ActivityMetadata,omitempty" xml:"ActivityMetadata,omitempty"`
	// The method used to adjust the number of instances in a scaling activity. Valid values:
	//
	// - `QuantityChangeInCapacity`: Adds or removes a specified number of ECS instances.
	//
	// - `PercentChangeInCapacity`: Adds or removes a specified percentage of ECS instances.
	//
	// - `TotalCapacity`: Adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value for the scaling activity. A single adjustment cannot add or remove more than 1,000 ECS instances. The valid range depends on `AdjustmentType`:
	//
	// - `QuantityChangeInCapacity`: -1000 to 1000.
	//
	// - `PercentChangeInCapacity`: -100 to 10000.
	//
	// - `TotalCapacity`: 0 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// A client-generated token to ensure the idempotence of the request. This token must be a unique string of up to 64 ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// - `None`: Executes a standard scaling activity.
	//
	// - `PlanOnly`: Only performs elastic planning and returns the results in `PlanResult` without triggering the scaling activity. The results include details such as instance types, availability zones, billing methods, and the number of new instances.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The lifecycle hook context.
	LifecycleHookContext *ScaleWithAdjustmentRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	// The minimum number of instances to adjust in a scaling activity. This parameter takes effect only when `AdjustmentType` is set to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The parameters to override when scaling out an ECI scaling group.
	Overrides *ScaleWithAdjustmentRequestOverrides `json:"Overrides,omitempty" xml:"Overrides,omitempty" type:"Struct"`
	OwnerId   *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether the current scaling activity supports concurrency.
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
	// Specifies whether to execute the scaling activity synchronously. This parameter applies only to scaling groups that are configured with an expected number of instances. Valid values:
	//
	// - `true`: Synchronous execution. The scaling activity is triggered immediately.
	//
	// - `false`: Asynchronous execution. The call updates the expected number of instances without immediately triggering the scaling activity. The activity occurs when the system detects a discrepancy between the new expected number and the current number of instances.
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
	// - `true`: Disables all lifecycle hooks.
	//
	// - `false`: Does not disable lifecycle hooks.
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// A list of lifecycle hook IDs to ignore during the scaling activity.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
	LifecycleHookResult     *string   `json:"LifecycleHookResult,omitempty" xml:"LifecycleHookResult,omitempty"`
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
	// A list of container-specific overrides.
	ContainerOverrides []*ScaleWithAdjustmentRequestOverridesContainerOverrides `json:"ContainerOverrides,omitempty" xml:"ContainerOverrides,omitempty" type:"Repeated"`
	// The number of vCPUs for the instance. Unit: cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory size for the instance. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The user data for the ECS instance. It must be Base64-encoded, and the raw data cannot exceed 32 KB.
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
	// The arguments for the container\\"s startup command. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container\\"s startup command, specified as an array of strings. You can specify up to 20 strings, and each can be up to 256 characters long.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs for the container. Unit: cores.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Environment variables to set in the container.
	EnvironmentVars []*ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The memory size for the container. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container to override. The override takes effect only if this name matches a container name in the scaling configuration.
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
	// The name of the environment variable. It must be 1 to 128 characters long, cannot start with a digit, and can contain only letters (a-z, A-Z), digits (0-9), and underscores (_).
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable, up to 256 characters long.
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
