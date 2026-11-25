// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScheduledTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScheduledTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScheduledTasksResponseBody
	GetRequestId() *string
	SetScheduledTasks(v []*DescribeScheduledTasksResponseBodyScheduledTasks) *DescribeScheduledTasksResponseBody
	GetScheduledTasks() []*DescribeScheduledTasksResponseBodyScheduledTasks
	SetTotalCount(v int32) *DescribeScheduledTasksResponseBody
	GetTotalCount() *int32
}

type DescribeScheduledTasksResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the scheduled tasks.
	ScheduledTasks []*DescribeScheduledTasksResponseBodyScheduledTasks `json:"ScheduledTasks,omitempty" xml:"ScheduledTasks,omitempty" type:"Repeated"`
	// The total number of scheduled tasks.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScheduledTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScheduledTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScheduledTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduledTasksResponseBody) GetScheduledTasks() []*DescribeScheduledTasksResponseBodyScheduledTasks {
	return s.ScheduledTasks
}

func (s *DescribeScheduledTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScheduledTasksResponseBody) SetPageNumber(v int32) *DescribeScheduledTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetPageSize(v int32) *DescribeScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetRequestId(v string) *DescribeScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetScheduledTasks(v []*DescribeScheduledTasksResponseBodyScheduledTasks) *DescribeScheduledTasksResponseBody {
	s.ScheduledTasks = v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTotalCount(v int32) *DescribeScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) Validate() error {
	if s.ScheduledTasks != nil {
		for _, item := range s.ScheduledTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScheduledTasksResponseBodyScheduledTasks struct {
	// The description of the scheduled task.
	//
	// example:
	//
	// Test scheduled task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected number of instances in the scaling group. This parameter takes effect when the scaling method of the scheduled task is set to adjust the number of instances in the scaling group.
	//
	// example:
	//
	// 10
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The time period during which a retry is performed after the scheduled task fails to be triggered. Unit: seconds. Valid values: 0 to 21600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The time at which the scheduled task is triggered. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-08-18T10:52Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The maximum number of instances in the scaling group. This parameter takes effect when the scaling method of the scheduled task is set to adjust the number of instances in the scaling group.
	//
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum number of instances in the scaling group. This parameter takes effect when the scaling method of the scheduled task is set to adjust the number of instances in the scaling group.
	//
	// example:
	//
	// 0
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The end time of the recurring scheduled task. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-08-20T16:55Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// The recurrence type of the scheduled task.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The recurrence value of the scheduled task.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The ID of the scaling group that executes the scheduled task.
	//
	// example:
	//
	// asg-bp1bo5tca4m56nap****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling rule used when the scheduled task is executed. This parameter is returned only when a scaling rule is specified for the scheduled task.
	//
	// example:
	//
	// ari:acs:ess:cn-hangzhou:1406926474****:scalingrule/asr-bp1id5rhu8edp7kd****
	ScheduledAction *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	// The ID of the scheduled task.
	//
	// example:
	//
	// edRtShc57WGXdt8TlPbr****
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	// The name of the scheduled task.
	//
	// example:
	//
	// scheduled****
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	// Indicates whether the scheduled task is started. Valid values:
	//
	// - true: The task is started.
	//
	// - false: The task is stopped.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	TaskEnabled *bool `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	// The warm pool capacity configurations of the scaling group.
	WarmPoolCapacityOptions *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions `json:"WarmPoolCapacityOptions,omitempty" xml:"WarmPoolCapacityOptions,omitempty" type:"Struct"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetLaunchTime() *string {
	return s.LaunchTime
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetMaxValue() *int32 {
	return s.MaxValue
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetMinValue() *int32 {
	return s.MinValue
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetRecurrenceEndTime() *string {
	return s.RecurrenceEndTime
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetScheduledAction() *string {
	return s.ScheduledAction
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetScheduledTaskId() *string {
	return s.ScheduledTaskId
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetScheduledTaskName() *string {
	return s.ScheduledTaskName
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetTaskEnabled() *bool {
	return s.TaskEnabled
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) GetWarmPoolCapacityOptions() *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions {
	return s.WarmPoolCapacityOptions
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetDescription(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.Description = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetDesiredCapacity(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetLaunchExpirationTime(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetLaunchTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.LaunchTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetMaxValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.MaxValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetMinValue(v int32) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.MinValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceEndTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceEndTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceType(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetRecurrenceValue(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScalingGroupId(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledAction(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledAction = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTaskId(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTaskId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTaskName(v string) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTaskName = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetTaskEnabled(v bool) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.TaskEnabled = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetWarmPoolCapacityOptions(v *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.WarmPoolCapacityOptions = v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) Validate() error {
	if s.WarmPoolCapacityOptions != nil {
		if err := s.WarmPoolCapacityOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions struct {
	// The expected number of instances covered by the warm pool. The system maintains the sum of the current number of instances in the scaling group and the number of instances in the warm pool at this value.
	//
	// example:
	//
	// 5
	DesiredGroupPreparedCapacity *int32 `json:"DesiredGroupPreparedCapacity,omitempty" xml:"DesiredGroupPreparedCapacity,omitempty"`
	// The minimum number of instances in the warm pool.
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) GetDesiredGroupPreparedCapacity() *int32 {
	return s.DesiredGroupPreparedCapacity
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) SetDesiredGroupPreparedCapacity(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions {
	s.DesiredGroupPreparedCapacity = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) SetMinCapacity(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions {
	s.MinCapacity = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksWarmPoolCapacityOptions) Validate() error {
	return dara.Validate(s)
}
