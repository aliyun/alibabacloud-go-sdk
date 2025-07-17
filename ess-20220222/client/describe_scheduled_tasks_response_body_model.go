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
	// The page number of the returned page.
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information collection of the scheduled tasks.
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
	return dara.Validate(s)
}

type DescribeScheduledTasksResponseBodyScheduledTasks struct {
	// The description of the scheduled task.
	//
	// example:
	//
	// Test scheduled task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expected number of instances in the scaling group. If you set `Scaling Method` to `Configure Number of Instances in Scaling Group`, you can specify this parameter.
	//
	// example:
	//
	// 10
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// The time window during which the scheduled task can be retried. Unit: seconds. Valid values: 0 to 21600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The point in time at which the scheduled task is triggered.
	//
	// example:
	//
	// 2014-08-18T10:52Z
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The maximum number of instances that must be contained in the scaling group. If you set `Scaling Method` to `Configure Number of Instances in Scaling Group`, you can specify this parameter.
	//
	// example:
	//
	// 10
	MaxValue *int32 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum number of instances that must be contained in the scaling group. If you set `Scaling Method` to `Configure Number of Instances in Scaling Group`, you can specify this parameter.
	//
	// example:
	//
	// 0
	MinValue *int32 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The end time of the recurrence of the scheduled task.
	//
	// example:
	//
	// 2014-08-20T16:55Z
	RecurrenceEndTime *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	// The recurring interval of the scheduled task.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The frequency of recurrence of the scheduled task.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The ID of the scaling group to which the scheduled task belongs.
	//
	// example:
	//
	// asg-bp1bo5tca4m56nap****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling rule of the scheduled task. A value is returned for this parameter only after you specify ScheduledActions.
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
	// Indicates whether the scheduled task is enabled.
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

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) Validate() error {
	return dara.Validate(s)
}
