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
	SetWarmPoolCapacityOptions(v *CreateScheduledTaskRequestWarmPoolCapacityOptions) *CreateScheduledTaskRequest
	GetWarmPoolCapacityOptions() *CreateScheduledTaskRequestWarmPoolCapacityOptions
}

type CreateScheduledTaskRequest struct {
	// The description of the scheduled task. The description must be 2 to 200 characters in length.
	//
	// example:
	//
	// Test scheduled task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected number of instances in the scaling group when the scaling method of the scheduled task is set to specifying the number of instances in the scaling group.
	//
	// > The scaling group must support the expected number of instances. This means that DesiredCapacity must be specified when the scaling group is created.
	//
	// example:
	//
	// 10
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The period within which the scheduled task retries after a trigger failure. Unit: seconds. Valid values: 0 to 1800.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The time at which the scheduled task is triggered. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot specify a time that is later than 90 days from the creation date.
	//
	// - If RecurrenceType is specified, the time specified by this parameter is the default time for recurring execution.
	//
	// - If RecurrenceType is not specified, the task is executed only once at the specified date and time.
	//
	// example:
	//
	// 2014-08-17T16:52Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The maximum number of instances in the scaling group when the scaling method of the scheduled task is set to specifying the number of instances in the scaling group.
	//
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum number of instances in the scaling group when the scaling method of the scheduled task is set to specifying the number of instances in the scaling group.
	//
	// example:
	//
	// 0
	MinValue     *int32  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The end time of the scheduled task that is executed repeatedly. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format. The time must be in UTC. You cannot specify a time that is later than 365 days from the creation date.
	//
	// example:
	//
	// 2014-08-17T16:55Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// The recurrence type of the scheduled task. Valid values:
	//
	// - Daily: The scheduled task is executed once every specified number of days.
	//
	// - Weekly: The scheduled task is executed on specified days of each week.
	//
	// - Monthly: The scheduled task is executed on specified days of each month.
	//
	// - Cron: The scheduled task is executed based on the specified cron expression.
	//
	// You must specify both RecurrenceType and RecurrenceValue.
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
	// - If RecurrenceType is set to Cron, the value represents UTC time and supports a five-field expression for minutes, hours, days, months, and days of the week. The expression supports the following wildcards: commas (,), question marks (?), hyphens (-), asterisks (*), number signs (#), forward slashes (/), L, and W.
	//
	// You must specify both RecurrenceType and RecurrenceValue.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The ID of the region to which the scheduled task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group in which you want to modify the number of instances or adjust the warm pool capacity when the scheduled task is triggered.
	//
	// After you specify ScalingGroupId, the scaling method of the scheduled task is set to specifying the number of instances in the scaling group or adjusting the warm pool capacity.
	//
	// - To specify the number of instances in the scaling group, you must specify a value for at least one of the following parameters: MinValue, MaxValue, and DesiredCapacity.
	//
	// - To adjust the warm pool capacity, you must configure WarmPoolCapacityOptions.
	//
	// > You cannot specify both ScheduledAction and ScalingGroupId at the same time.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The unique identifier of the scaling rule to be executed when the scheduled task is triggered.
	//
	// After you specify ScheduledAction, the scaling method of the scheduled task is set to selecting an existing scaling rule. You cannot specify both ScheduledAction and ScalingGroupId at the same time.
	//
	// example:
	//
	// ari:acs:ess:cn-hangzhou:140692647****:scalingrule/asr-bp12tcnol686y1ik****
	ScheduledAction *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	// The name of the scheduled task. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter, digit, or Chinese character. The name must be unique within the same region and Alibaba Cloud account.
	//
	// Default value: the value of ScheduledTaskId.
	//
	// example:
	//
	// scheduled****
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	// Specifies whether to enable the scheduled task. Valid values:
	//
	// - true: enables the task.
	//
	// - false: disables the task.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	TaskEnabled *bool `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	// The warm pool capacity configurations of the scaling group. This parameter takes effect only when a warm pool is configured for the scaling group.
	//
	// > When you configure this parameter, you must specify ScalingGroupId.
	WarmPoolCapacityOptions *CreateScheduledTaskRequestWarmPoolCapacityOptions `json:"WarmPoolCapacityOptions,omitempty" xml:"WarmPoolCapacityOptions,omitempty" type:"Struct"`
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

func (s *CreateScheduledTaskRequest) GetWarmPoolCapacityOptions() *CreateScheduledTaskRequestWarmPoolCapacityOptions {
	return s.WarmPoolCapacityOptions
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

func (s *CreateScheduledTaskRequest) SetWarmPoolCapacityOptions(v *CreateScheduledTaskRequestWarmPoolCapacityOptions) *CreateScheduledTaskRequest {
	s.WarmPoolCapacityOptions = v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	if s.WarmPoolCapacityOptions != nil {
		if err := s.WarmPoolCapacityOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScheduledTaskRequestWarmPoolCapacityOptions struct {
	// The expected number of prepared instances to be covered. The system maintains the sum of the current number of instances in the scaling group and the number of instances in the warm pool at this value. Default value: the maximum number of instances in the scaling group, which is specified by MaxSize.
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

func (s CreateScheduledTaskRequestWarmPoolCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestWarmPoolCapacityOptions) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestWarmPoolCapacityOptions) GetDesiredGroupPreparedCapacity() *int32 {
	return s.DesiredGroupPreparedCapacity
}

func (s *CreateScheduledTaskRequestWarmPoolCapacityOptions) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *CreateScheduledTaskRequestWarmPoolCapacityOptions) SetDesiredGroupPreparedCapacity(v int32) *CreateScheduledTaskRequestWarmPoolCapacityOptions {
	s.DesiredGroupPreparedCapacity = &v
	return s
}

func (s *CreateScheduledTaskRequestWarmPoolCapacityOptions) SetMinCapacity(v int32) *CreateScheduledTaskRequestWarmPoolCapacityOptions {
	s.MinCapacity = &v
	return s
}

func (s *CreateScheduledTaskRequestWarmPoolCapacityOptions) Validate() error {
	return dara.Validate(s)
}
