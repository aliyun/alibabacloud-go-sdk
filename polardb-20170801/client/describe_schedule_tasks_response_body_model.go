// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeScheduleTasksResponseBodyData) *DescribeScheduleTasksResponseBody
	GetData() *DescribeScheduleTasksResponseBodyData
	SetMessage(v string) *DescribeScheduleTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeScheduleTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeScheduleTasksResponseBody
	GetSuccess() *bool
}

type DescribeScheduleTasksResponseBody struct {
	// The result set.
	Data *DescribeScheduleTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// > If the request is successful, **Successful*	- is returned. Otherwise, an error message is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 936C7025-27A5-4CB1-BB31-540E1F0CCA12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeScheduleTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBody) GetData() *DescribeScheduleTasksResponseBodyData {
	return s.Data
}

func (s *DescribeScheduleTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScheduleTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduleTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeScheduleTasksResponseBody) SetData(v *DescribeScheduleTasksResponseBodyData) *DescribeScheduleTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetMessage(v string) *DescribeScheduleTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetRequestId(v string) *DescribeScheduleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetSuccess(v bool) *DescribeScheduleTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeScheduleTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScheduleTasksResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Details of scheduled tasks.
	TimerInfos []*DescribeScheduleTasksResponseBodyDataTimerInfos `json:"TimerInfos,omitempty" xml:"TimerInfos,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeScheduleTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScheduleTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScheduleTasksResponseBodyData) GetTimerInfos() []*DescribeScheduleTasksResponseBodyDataTimerInfos {
	return s.TimerInfos
}

func (s *DescribeScheduleTasksResponseBodyData) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeScheduleTasksResponseBodyData) SetPageNumber(v int32) *DescribeScheduleTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetPageSize(v int32) *DescribeScheduleTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetTimerInfos(v []*DescribeScheduleTasksResponseBodyDataTimerInfos) *DescribeScheduleTasksResponseBodyData {
	s.TimerInfos = v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetTotalRecordCount(v int32) *DescribeScheduleTasksResponseBodyData {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) Validate() error {
	if s.TimerInfos != nil {
		for _, item := range s.TimerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScheduleTasksResponseBodyDataTimerInfos struct {
	// The task type.
	//
	// example:
	//
	// CreateDBNodes
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ID of the scheduled O\\&M task.
	//
	// example:
	//
	// 86293c29-a03d-4872-b625-***********
	CrontabJobId *string `json:"CrontabJobId,omitempty" xml:"CrontabJobId,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test_cluster
	DbClusterDescription *string `json:"DbClusterDescription,omitempty" xml:"DbClusterDescription,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// Running
	DbClusterStatus *string `json:"DbClusterStatus,omitempty" xml:"DbClusterStatus,omitempty"`
	// The order ID.
	//
	// > This parameter is returned only if `Action` is **CreateDBNodes*	- or **ModifyDBNodeClass**.
	//
	// example:
	//
	// 208161753******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The latest start time of the task. The time is in UTC.
	//
	// example:
	//
	// 2021-01-28T12:30Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The scheduled time for the transient switchover.
	//
	// example:
	//
	// 2021-01-28T12:30Z
	PlannedFlashingOffTime *string `json:"PlannedFlashingOffTime,omitempty" xml:"PlannedFlashingOffTime,omitempty"`
	// The scheduled start time of the task. The time is in UTC.
	//
	// example:
	//
	// 2021-01-28T12:00Z
	PlannedStartTime *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	// The estimated start time of the task. The time is in UTC.
	//
	// example:
	//
	// 2021-01-28T12:16Z
	PlannedTime *string `json:"PlannedTime,omitempty" xml:"PlannedTime,omitempty"`
	// The region ID of the task.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the scheduled task can be canceled. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	TaskCancel *bool `json:"TaskCancel,omitempty" xml:"TaskCancel,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 53879cdb-9a00-428e-acaf-ff4cff******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScheduleTasksResponseBodyDataTimerInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleTasksResponseBodyDataTimerInfos) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetCrontabJobId() *string {
	return s.CrontabJobId
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetDbClusterDescription() *string {
	return s.DbClusterDescription
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetDbClusterStatus() *string {
	return s.DbClusterStatus
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetPlannedFlashingOffTime() *string {
	return s.PlannedFlashingOffTime
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetPlannedTime() *string {
	return s.PlannedTime
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetRegion() *string {
	return s.Region
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetTaskCancel() *bool {
	return s.TaskCancel
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetAction(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Action = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetCrontabJobId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.CrontabJobId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDBClusterId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DBClusterId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDbClusterDescription(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DbClusterDescription = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDbClusterStatus(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DbClusterStatus = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetOrderId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.OrderId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedEndTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedEndTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedFlashingOffTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedFlashingOffTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedStartTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedStartTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetRegion(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Region = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetStatus(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Status = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetTaskCancel(v bool) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.TaskCancel = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetTaskId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.TaskId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) Validate() error {
	return dara.Validate(s)
}
