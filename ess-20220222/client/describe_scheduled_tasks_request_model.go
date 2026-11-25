// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeScheduledTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScheduledTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScheduledTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScheduledTasksRequest
	GetPageSize() *int32
	SetRecurrenceType(v string) *DescribeScheduledTasksRequest
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *DescribeScheduledTasksRequest
	GetRecurrenceValue() *string
	SetRegionId(v string) *DescribeScheduledTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScheduledTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScheduledTasksRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DescribeScheduledTasksRequest
	GetScalingGroupId() *string
	SetScheduledActions(v []*string) *DescribeScheduledTasksRequest
	GetScheduledActions() []*string
	SetScheduledTaskIds(v []*string) *DescribeScheduledTasksRequest
	GetScheduledTaskIds() []*string
	SetScheduledTaskNames(v []*string) *DescribeScheduledTasksRequest
	GetScheduledTaskNames() []*string
	SetTaskEnabled(v bool) *DescribeScheduledTasksRequest
	GetTaskEnabled() *bool
	SetTaskName(v string) *DescribeScheduledTasksRequest
	GetTaskName() *string
}

type DescribeScheduledTasksRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the scheduled task list. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The recurrence type of the scheduled task. Valid values:
	//
	// - Daily: The scheduled task is executed once every specified number of days.
	//
	// - Weekly: The scheduled task is executed on specified days of each week.
	//
	// - Monthly: The scheduled task is executed on specified days of each month.
	//
	// - Cron: The scheduled task is executed based on a specified cron expression.
	//
	// example:
	//
	// Weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The recurrence value of the scheduled task.
	//
	// This parameter is supported only when RecurrenceType is set to Weekly. You can specify multiple values separated by commas (,). For example, the values for Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday are 0, 1, 2, 3, 4, 5, and 6.
	//
	// example:
	//
	// 1,2,3
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The region ID of the scaling group to which the scheduled task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group that executes the scheduled task.
	//
	// example:
	//
	// asg-bp1bo5tca4m56nap****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scaling rules that are used when the scheduled tasks are triggered. This parameter is valid only when a scaling rule is specified for the scheduled task.
	ScheduledActions []*string `json:"ScheduledActions,omitempty" xml:"ScheduledActions,omitempty" type:"Repeated"`
	// The IDs of the scheduled tasks to be queried.
	ScheduledTaskIds []*string `json:"ScheduledTaskIds,omitempty" xml:"ScheduledTaskIds,omitempty" type:"Repeated"`
	// The names of the scheduled tasks to be queried.
	ScheduledTaskNames []*string `json:"ScheduledTaskNames,omitempty" xml:"ScheduledTaskNames,omitempty" type:"Repeated"`
	// Specifies whether the scheduled task is started. Valid values:
	//
	// - true: The task is started.
	//
	// - false: The task is stopped.
	//
	// example:
	//
	// true
	TaskEnabled *bool `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
	// The name of the task. Fuzzy match by keyword is supported.
	//
	// example:
	//
	// scheduled****
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScheduledTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScheduledTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScheduledTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScheduledTasksRequest) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribeScheduledTasksRequest) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *DescribeScheduledTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScheduledTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScheduledTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScheduledTasksRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScheduledTasksRequest) GetScheduledActions() []*string {
	return s.ScheduledActions
}

func (s *DescribeScheduledTasksRequest) GetScheduledTaskIds() []*string {
	return s.ScheduledTaskIds
}

func (s *DescribeScheduledTasksRequest) GetScheduledTaskNames() []*string {
	return s.ScheduledTaskNames
}

func (s *DescribeScheduledTasksRequest) GetTaskEnabled() *bool {
	return s.TaskEnabled
}

func (s *DescribeScheduledTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeScheduledTasksRequest) SetOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageNumber(v int32) *DescribeScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageSize(v int32) *DescribeScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetRecurrenceType(v string) *DescribeScheduledTasksRequest {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetRecurrenceValue(v string) *DescribeScheduledTasksRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetRegionId(v string) *DescribeScheduledTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScalingGroupId(v string) *DescribeScheduledTasksRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledActions(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledActions = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskIds(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskIds = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskNames(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskNames = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetTaskEnabled(v bool) *DescribeScheduledTasksRequest {
	s.TaskEnabled = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetTaskName(v string) *DescribeScheduledTasksRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
