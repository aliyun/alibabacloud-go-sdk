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
}

type ModifyScheduledTaskRequest struct {
	// The description of the scheduled task. The description must be 2 to 200 characters in length.
	//
	// example:
	//
	// Test scheduled task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected number of instances in the scaling group if you specify the ScalingGroupId parameter.
	//
	// > You must specify the `DesiredCapacity` parameter when you create a scaling group.
	//
	// example:
	//
	// 10
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The time period during which the failed scheduled task is retried. Unit: seconds. Valid values: 0 to 1800.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The point in time at which the scheduled task is triggered. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot enter a time point later than 90 days from the point in time at which the scheduled task is modified.
	//
	// 	- If you specify the `RecurrenceType` parameter, the task is repeatedly executed at the time point that is specified by the LaunchTime parameter.
	//
	// 	- If you do not specify the `RecurrenceType` parameter, the task is executed only once at the point in time that is specified by the LaunchTime parameter.
	//
	// example:
	//
	// 2014-08-18T10:52Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The maximum number of instances in the scaling group if you specify the ScalingGroupId parameter.
	//
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum number of instances in the scaling group if you specify the ScalingGroupId parameter.
	//
	// example:
	//
	// 0
	MinValue     *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The end time of the scheduled task. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot enter a point in time that is later than 365 days from the point in time at which the scheduled task is modified.
	//
	// example:
	//
	// 2014-08-20T16:55Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// The interval at which the scheduled task is repeated. Valid values:
	//
	// 	- Daily: The scheduled task is executed once every specified number of days.
	//
	// 	- Weekly: The scheduled task is executed on each specified day of the week.
	//
	// 	- Monthly: The scheduled task is executed on each specified day of the month.
	//
	// 	- Cron: The scheduled task is executed based on the specified cron expression.
	//
	// After you modify the scheduled task, the values that you specify for the `RecurrenceType` and `RecurrenceValue` parameters must be valid at the same time.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The number of recurrences of the scheduled task.
	//
	// 	- If you set the `RecurrenceType` parameter to `Daily`, you can specify only one value for this parameter. Valid values: 1 to 31.
	//
	// 	- If you set the `RecurrenceType` parameter to `Weekly`, you can specify multiple values for this parameter. Separate the values with commas (,). The values that correspond to Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday are 0, 1, 2, 3, 4, 5, and 6.``
	//
	// 	- If you set the `RecurrenceType` parameter to `Monthly`, you can specify two values in the `A-B` format for this parameter. Valid values of A and B: 1 to 31. B must be greater than or equal to A.
	//
	// 	- If you set the `RecurrenceType` parameter to `Cron`, you can specify a cron expression. A cron expression is written in UTC time and consists of the following fields: minute, hour, day, month, and week. The expression can contain the letters L and W and the following wildcard characters: commas (,), question marks (?), hyphens (-), asterisks (\\*), number signs (#), and forward slashes (/).
	//
	// After you modify the scheduled task, the values that you specify for the `RecurrenceType` and `RecurrenceValue` parameters must be valid at the same time.
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
	// The ID of the scaling group whose number of instances must be modified when the scheduled task is triggered. If you specify the `ScalingGroupId` parameter for a scheduled task, you must specify the minimum, maximum, or expected numbers of instances for a scaling group in the scheduled task. That is, you must specify at least one of the `MinValue`, `MaxValue`, and `DesiredCapacity` parameters.
	//
	// > You cannot specify the `ScheduledAction` and `ScalingGroupId` parameters at the same time.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling rule that you want to execute when the scheduled task is triggered. Enter the ID of the scaling rule. To obtain the ID of a scaling rule, go to the scaling rule details page. If you specify `ScheduledAction`, you must select an existing scaling rule for the scheduled task.
	//
	// >  You can specify only one of `ScheduledAction` and `ScalingGroupId`.
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
	// The name of the scheduled task. The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter or a digit. The name of the scheduled task must be unique in the region and within the Alibaba Cloud account.
	//
	// example:
	//
	// scheduled****
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	// Specifies whether to enable the scheduled task. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	TaskEnabled *bool `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
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

func (s *ModifyScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
