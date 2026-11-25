// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyScheduledTaskRequest
	GetDescription() *string
	SetDesiredCapacity(v int32) *ModifyScheduledTaskRequest
	GetDesiredCapacity() *int32
	SetLaunchExpirationTime(v int32) *ModifyScheduledTaskRequest
	GetLaunchExpirationTime() *int32
	SetLaunchTime(v string) *ModifyScheduledTaskRequest
	GetLaunchTime() *string
	SetMaxValue(v int32) *ModifyScheduledTaskRequest
	GetMaxValue() *int32
	SetMinValue(v int32) *ModifyScheduledTaskRequest
	GetMinValue() *int32
	SetOwnerAccount(v string) *ModifyScheduledTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScheduledTaskRequest
	GetOwnerId() *int64
	SetRecurrenceEndTime(v string) *ModifyScheduledTaskRequest
	GetRecurrenceEndTime() *string
	SetRecurrenceType(v string) *ModifyScheduledTaskRequest
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *ModifyScheduledTaskRequest
	GetRecurrenceValue() *string
	SetRegionId(v string) *ModifyScheduledTaskRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyScheduledTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyScheduledTaskRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *ModifyScheduledTaskRequest
	GetScalingGroupId() *string
	SetScheduledAction(v string) *ModifyScheduledTaskRequest
	GetScheduledAction() *string
	SetScheduledTaskId(v string) *ModifyScheduledTaskRequest
	GetScheduledTaskId() *string
	SetScheduledTaskName(v string) *ModifyScheduledTaskRequest
	GetScheduledTaskName() *string
	SetTaskEnabled(v bool) *ModifyScheduledTaskRequest
	GetTaskEnabled() *bool
	SetWarmPoolCapacityOptions(v *ModifyScheduledTaskRequestWarmPoolCapacityOptions) *ModifyScheduledTaskRequest
	GetWarmPoolCapacityOptions() *ModifyScheduledTaskRequestWarmPoolCapacityOptions
}

type ModifyScheduledTaskRequest struct {
	// The description of the scheduled task. The description must be 2 to 200 characters in length.
	//
	// example:
	//
	// Test scheduled task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected number of instances in the scaling group. This parameter is required when the scaling method of the scheduled task is set to specify the number of instances in the scaling group.
	//
	// > The scaling group must support the expected number of instances. This means that DesiredCapacity must be specified when the scaling group is created.
	//
	// example:
	//
	// 10
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The retry period during which Auto Scaling continues to attempt to execute the scheduled task after the triggered operation fails. Unit: seconds. Valid values: 0 to 1800.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The time at which the scheduled task is executed. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot specify a time that is later than 90 days from the date of modification.
	//
	// - If RecurrenceType is specified, the time specified by this parameter is the default time for recurring execution.
	//
	// - If RecurrenceType is not specified, the task is executed only once at the specified date and time.
	//
	// example:
	//
	// 2014-08-18T10:52Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The maximum number of instances in the scaling group. This parameter is required when the scaling method of the scheduled task is set to specify the number of instances in the scaling group.
	//
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum number of instances in the scaling group. This parameter is required when the scaling method of the scheduled task is set to specify the number of instances in the scaling group.
	//
	// example:
	//
	// 0
	MinValue     *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The end time of the scheduled task that is executed repeatedly. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot specify a time that is later than 365 days from the date of modification.
	//
	// example:
	//
	// 2014-08-20T16:55Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// The recurrence type of the scheduled task. Valid values:
	//
	//
	//
	// - Daily: The scheduled task is executed once every specified number of days.
	//
	// - Weekly: The scheduled task is executed on specified days of each week.
	//
	// - Monthly: The scheduled task is executed on specified days of each month.
	//
	// - Cron: The scheduled task is executed based on the specified cron expression.
	//
	// After modification, both RecurrenceType and RecurrenceValue must be valid.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The recurrence value of the scheduled task.
	//
	// - If RecurrenceType is set to Daily, you can specify only one value. Valid values: 1 to 31.
	//
	// - If RecurrenceType is set to Weekly, you can specify multiple values separated by commas (,). For example, the values for Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday are 0,1,2,3,4,5,6.
	//
	// - If RecurrenceType is set to Monthly, the format is A-B. Valid values of A and B: 1 to 31. B must be greater than or equal to A.
	//
	// - If RecurrenceType is set to Cron, the value specifies a cron expression in UTC. The expression supports five fields: minute, hour, day, month, and week. It supports wildcards such as commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), L, and W.
	//
	// After modification, both RecurrenceType and RecurrenceValue must be valid.
	//
	// example:
	//
	// 2
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group in which you want to modify the number of instances or adjust the warm pool capacity when the scheduled task is triggered.
	//
	// After you specify ScalingGroupId, the scaling method of the scheduled task is set to specify the number of instances in the scaling group or adjust the warm pool capacity.
	//
	// - Specify the number of instances in the scaling group: You must specify at least one of the MinValue, MaxValue, and DesiredCapacity parameters.
	//
	// - Adjust the warm pool capacity: You must specify WarmPoolCapacityOptions.
	//
	// > You cannot specify both ScheduledAction and ScalingGroupId at the same time.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The ID of the scaling rule to execute. This parameter is available only when the scaling method is set to execute an existing scaling rule. You can view the ID on the details page of the scaling rule.
	//
	// > You cannot specify both ScheduledAction and ScalingGroupId at the same time.
	//
	// example:
	//
	// ari:acs:ess:cn-hangzhou:14069264****:scalingrule/asr-bp12tcnol686y1ik****
	ScheduledAction *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	// The ID of the scheduled task.
	//
	// This parameter is required.
	//
	// example:
	//
	// edRtShc57WGXdt8TlPbr****
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	// The name of the scheduled task. The name must be 2 to 64 characters in length and can contain periods (.), underscores (_), and hyphens (-). It must start with a letter, digit, or Chinese character. The name must be unique within the same region and Alibaba Cloud account.
	//
	// example:
	//
	// scheduled****
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	// Specifies whether to enable the scheduled task. Valid values:
	//
	// - true: enables the scheduled task.
	//
	// - false: disables the scheduled task.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	TaskEnabled *bool `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	// The warm pool capacity configuration of the scaling group. This parameter is valid only when a warm pool is configured for the scaling group.
	//
	// > You must specify ScalingGroupId when you specify this parameter.
	WarmPoolCapacityOptions *ModifyScheduledTaskRequestWarmPoolCapacityOptions `json:"WarmPoolCapacityOptions,omitempty" xml:"WarmPoolCapacityOptions,omitempty" type:"Struct"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyScheduledTaskRequest) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *ModifyScheduledTaskRequest) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *ModifyScheduledTaskRequest) GetLaunchTime() *string {
	return s.LaunchTime
}

func (s *ModifyScheduledTaskRequest) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *ModifyScheduledTaskRequest) GetMinValue() *int32 {
	return s.MinValue
}

func (s *ModifyScheduledTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScheduledTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScheduledTaskRequest) GetRecurrenceEndTime() *string {
	return s.RecurrenceEndTime
}

func (s *ModifyScheduledTaskRequest) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ModifyScheduledTaskRequest) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ModifyScheduledTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyScheduledTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScheduledTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyScheduledTaskRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyScheduledTaskRequest) GetScheduledAction() *string {
	return s.ScheduledAction
}

func (s *ModifyScheduledTaskRequest) GetScheduledTaskId() *string {
	return s.ScheduledTaskId
}

func (s *ModifyScheduledTaskRequest) GetScheduledTaskName() *string {
	return s.ScheduledTaskName
}

func (s *ModifyScheduledTaskRequest) GetTaskEnabled() *bool {
	return s.TaskEnabled
}

func (s *ModifyScheduledTaskRequest) GetWarmPoolCapacityOptions() *ModifyScheduledTaskRequestWarmPoolCapacityOptions {
	return s.WarmPoolCapacityOptions
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetDesiredCapacity(v int32) *ModifyScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchExpirationTime(v int32) *ModifyScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchTime(v string) *ModifyScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMaxValue(v int32) *ModifyScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetMinValue(v int32) *ModifyScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceEndTime(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceType(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceValue(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRegionId(v string) *ModifyScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScalingGroupId(v string) *ModifyScheduledTaskRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledAction(v string) *ModifyScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskName(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTaskEnabled(v bool) *ModifyScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetWarmPoolCapacityOptions(v *ModifyScheduledTaskRequestWarmPoolCapacityOptions) *ModifyScheduledTaskRequest {
	s.WarmPoolCapacityOptions = v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	if s.WarmPoolCapacityOptions != nil {
		if err := s.WarmPoolCapacityOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyScheduledTaskRequestWarmPoolCapacityOptions struct {
	// The expected number of prepared instances. The system maintains the sum of the current number of instances in the scaling group and the number of instances in the warm pool at this value. Default value: the maximum number of instances in the scaling group, which is specified by MaxSize.
	//
	// example:
	//
	// 5
	DesiredGroupPreparedCapacity *int32 `json:"DesiredGroupPreparedCapacity,omitempty" xml:"DesiredGroupPreparedCapacity,omitempty"`
	// The minimum number of instances in the warm pool. The system always retains at least this number of prepared instances in the warm pool to ensure that instances can be quickly added from the warm pool during scale-out. Default value: 0.
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s ModifyScheduledTaskRequestWarmPoolCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequestWarmPoolCapacityOptions) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequestWarmPoolCapacityOptions) GetDesiredGroupPreparedCapacity() *int32 {
	return s.DesiredGroupPreparedCapacity
}

func (s *ModifyScheduledTaskRequestWarmPoolCapacityOptions) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *ModifyScheduledTaskRequestWarmPoolCapacityOptions) SetDesiredGroupPreparedCapacity(v int32) *ModifyScheduledTaskRequestWarmPoolCapacityOptions {
	s.DesiredGroupPreparedCapacity = &v
	return s
}

func (s *ModifyScheduledTaskRequestWarmPoolCapacityOptions) SetMinCapacity(v int32) *ModifyScheduledTaskRequestWarmPoolCapacityOptions {
	s.MinCapacity = &v
	return s
}

func (s *ModifyScheduledTaskRequestWarmPoolCapacityOptions) Validate() error {
	return dara.Validate(s)
}
