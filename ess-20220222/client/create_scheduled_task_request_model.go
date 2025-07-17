// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateScheduledTaskRequest
	GetDescription() *string
	SetDesiredCapacity(v int32) *CreateScheduledTaskRequest
	GetDesiredCapacity() *int32
	SetLaunchExpirationTime(v int32) *CreateScheduledTaskRequest
	GetLaunchExpirationTime() *int32
	SetLaunchTime(v string) *CreateScheduledTaskRequest
	GetLaunchTime() *string
	SetMaxValue(v int32) *CreateScheduledTaskRequest
	GetMaxValue() *int32
	SetMinValue(v int32) *CreateScheduledTaskRequest
	GetMinValue() *int32
	SetOwnerAccount(v string) *CreateScheduledTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateScheduledTaskRequest
	GetOwnerId() *int64
	SetRecurrenceEndTime(v string) *CreateScheduledTaskRequest
	GetRecurrenceEndTime() *string
	SetRecurrenceType(v string) *CreateScheduledTaskRequest
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *CreateScheduledTaskRequest
	GetRecurrenceValue() *string
	SetRegionId(v string) *CreateScheduledTaskRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateScheduledTaskRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *CreateScheduledTaskRequest
	GetScalingGroupId() *string
	SetScheduledAction(v string) *CreateScheduledTaskRequest
	GetScheduledAction() *string
	SetScheduledTaskName(v string) *CreateScheduledTaskRequest
	GetScheduledTaskName() *string
	SetTaskEnabled(v bool) *CreateScheduledTaskRequest
	GetTaskEnabled() *bool
}

type CreateScheduledTaskRequest struct {
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
	// The time period during which the failed scheduled task can be retried. Unit: seconds. Valid values: 0 to 1800.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The point in time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard. The time must be in UTC. You cannot trigger a scheduled task more than 90 days after its creation.
	//
	// 	- If you specify `RecurrenceType`, the scheduled task is repeatedly triggered at the point in time that is specified by LaunchTime.
	//
	// 	- If you do not specify `RecurrenceType`, the scheduled task is triggered only once at the time point that is specified by LaunchTime.
	//
	// example:
	//
	// 2014-08-17T16:52Z
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
	// The end time of the scheduled task. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	//
	// The time must be in UTC. You cannot enter a point in time that is later than 365 days from the point in time at which the scheduled task is created.
	//
	// example:
	//
	// 2014-08-17T16:55Z
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
	// You must specify the `RecurrenceType` and `RecurrenceValue` parameters at the same time.
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
	// You must specify both the `RecurrenceType` parameter and the `RecurrenceValue` parameter.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The region ID of the scheduled task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group. If you specify this parameter, the number of instances in the scaling group will be changed when the scheduled task is triggered.
	//
	// If you specify `ScalingGroupId`, you must also specify at least one of the following parameters: `MinValue`, `MaxValue`, and `DesiredCapacity`. You cannot specify `ScheduledAction` and `ScalingGroupId` at the same time.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling rule that you want to execute when the scheduled task is triggered. Specify the unique identifier of the scaling rule.
	//
	// If you specify `ScheduledAction`, the scheduled task uses an existing scaling rule. You cannot specify `ScheduledAction` and `ScalingGroupId` at the same time.
	//
	// example:
	//
	// ari:acs:ess:cn-hangzhou:140692647****:scalingrule/asr-bp12tcnol686y1ik****
	ScheduledAction *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	// The name of the scheduled task. The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit. The name of the scheduled task must be unique in the region and within the Alibaba Cloud account.
	//
	// If you do not specify this parameter, the value of the `ScheduledTaskId` parameter is used.
	//
	// example:
	//
	// scheduled****
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	// Specifies whether to enable the scheduled task.
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

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScheduledTaskRequest) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *CreateScheduledTaskRequest) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *CreateScheduledTaskRequest) GetLaunchTime() *string {
	return s.LaunchTime
}

func (s *CreateScheduledTaskRequest) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *CreateScheduledTaskRequest) GetMinValue() *int32 {
	return s.MinValue
}

func (s *CreateScheduledTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateScheduledTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateScheduledTaskRequest) GetRecurrenceEndTime() *string {
	return s.RecurrenceEndTime
}

func (s *CreateScheduledTaskRequest) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *CreateScheduledTaskRequest) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *CreateScheduledTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScheduledTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateScheduledTaskRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateScheduledTaskRequest) GetScheduledAction() *string {
	return s.ScheduledAction
}

func (s *CreateScheduledTaskRequest) GetScheduledTaskName() *string {
	return s.ScheduledTaskName
}

func (s *CreateScheduledTaskRequest) GetTaskEnabled() *bool {
	return s.TaskEnabled
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetDesiredCapacity(v int32) *CreateScheduledTaskRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchExpirationTime(v int32) *CreateScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchTime(v string) *CreateScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMaxValue(v int32) *CreateScheduledTaskRequest {
	s.MaxValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetMinValue(v int32) *CreateScheduledTaskRequest {
	s.MinValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerId(v int64) *CreateScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceEndTime(v string) *CreateScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceType(v string) *CreateScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceValue(v string) *CreateScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRegionId(v string) *CreateScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetResourceOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScalingGroupId(v string) *CreateScheduledTaskRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledAction(v string) *CreateScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledTaskName(v string) *CreateScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTaskEnabled(v bool) *CreateScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
