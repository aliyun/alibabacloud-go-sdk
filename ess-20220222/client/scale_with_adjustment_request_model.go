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
	// The type of the scaling policy. Valid values:
	//
	// 	- QuantityChangeInCapacity: adds the specified number of ECS instances to or removes the specified number of ECS instances from the scaling group.
	//
	// 	- PercentChangeInCapacity: adds the specified percentage of ECS instances to or removes the specified percentage of ECS instances from the scaling group.
	//
	// 	- TotalCapacity: adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The number of instances in each adjustment. The number of ECS instances in each adjustment cannot exceed 1,000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to QuantityChangeInCapacity: -1000 to 1000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to PercentChangeInCapacity: -100 to 10000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to TotalCapacity: 0 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- None: If this is not specified, auto scaling is performed.
	//
	// 	- PlanOnly: Scaling is not triggered. Only elastic planning is performed. The planning result is returned in PlanResult, including the instance type, zone ID, billing type, and number of created instances.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The context of the lifecycle hook.
	LifecycleHookContext *ScaleWithAdjustmentRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	// The minimum number of instances allowed in each adjustment. This parameter takes effect only if you set the `AdjustmentType` parameter to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The overrides that allow you to adjust the scaling group of the Elastic Container Instance type during a scale-out event.
	Overrides *ScaleWithAdjustmentRequestOverrides `json:"Overrides,omitempty" xml:"Overrides,omitempty" type:"Struct"`
	OwnerId   *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Whether the current scale-out task supports concurrency.
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
	// Specifies whether to trigger the scaling task in a synchronous manner. This parameter takes effect only on scaling groups for which you specified an expected number of instances. Valid Values:
	//
	// 	- true: triggers the scaling task in a synchronous manner. A scaling activity is triggered at the time when the scaling rule is executed.
	//
	// 	- false: does not trigger the scaling task in a synchronous manner. After you change the expected number of instances for the scaling group, Auto Scaling checks whether the total number of instances in the scaling group matches the new expected number and determines whether to trigger the scaling activity based on the check result.
	//
	// >  For more information, see [Expected number of instances](https://help.aliyun.com/document_detail/146231.html).
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
	return dara.Validate(s)
}

type ScaleWithAdjustmentRequestLifecycleHookContext struct {
	// Specifies whether to disable the Lifecycle Hook feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// The IDs of the lifecycle hooks that you want to disable.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
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

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) SetDisableLifecycleHook(v bool) *ScaleWithAdjustmentRequestLifecycleHookContext {
	s.DisableLifecycleHook = &v
	return s
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) SetIgnoredLifecycleHookIds(v []*string) *ScaleWithAdjustmentRequestLifecycleHookContext {
	s.IgnoredLifecycleHookIds = v
	return s
}

func (s *ScaleWithAdjustmentRequestLifecycleHookContext) Validate() error {
	return dara.Validate(s)
}

type ScaleWithAdjustmentRequestOverrides struct {
	// The list of parameters that you want to use to override specific configurations for containers.
	ContainerOverrides []*ScaleWithAdjustmentRequestOverridesContainerOverrides `json:"ContainerOverrides,omitempty" xml:"ContainerOverrides,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the instance. Unit: vCPUs.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory size that you want to allocate to the instance. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
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
	return dara.Validate(s)
}

type ScaleWithAdjustmentRequestOverridesContainerOverrides struct {
	// The argument that corresponds to the startup command of the container. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container startup commands. You can specify up to 20 commands. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs that you want to allocate to the container. Unit: vCPUs.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The information about the environment variables.
	EnvironmentVars []*ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The memory size that you want to allocate to the container. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of container N. If you specify ContainerOverrides, you must also specify Name. ContainerOverrides takes effect only when the container name specified by Name matches that specified in the scaling configuration.
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
	return dara.Validate(s)
}

type ScaleWithAdjustmentRequestOverridesContainerOverridesEnvironmentVars struct {
	// The name of the environment variable. The name must be 1 to 128 characters in length. Format requirement: `[0-9a-zA-Z]` and underscores (_). It cannot start with a digit.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be up to 256 characters in length.
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
