// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeTasksResponseBodyItems) *DescribeTasksResponseBody
	GetItems() *DescribeTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeTasksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeTasksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeTasksResponseBody struct {
	// The details of the task execution.
	Items *DescribeTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A103039D-B1B2-4C57-B989-7D7C0DA95426
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 40
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetItems() *DescribeTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeTasksResponseBody) SetItems(v *DescribeTasksResponseBodyItems) *DescribeTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageRecordCount(v int32) *DescribeTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTasksResponseBodyItems struct {
	TaskProgressInfo []*DescribeTasksResponseBodyItemsTaskProgressInfo `json:"TaskProgressInfo,omitempty" xml:"TaskProgressInfo,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyItems) GetTaskProgressInfo() []*DescribeTasksResponseBodyItemsTaskProgressInfo {
	return s.TaskProgressInfo
}

func (s *DescribeTasksResponseBodyItems) SetTaskProgressInfo(v []*DescribeTasksResponseBodyItemsTaskProgressInfo) *DescribeTasksResponseBodyItems {
	s.TaskProgressInfo = v
	return s
}

func (s *DescribeTasksResponseBodyItems) Validate() error {
	if s.TaskProgressInfo != nil {
		for _, item := range s.TaskProgressInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyItemsTaskProgressInfo struct {
	// The start time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-11-20T01:00Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The name of the subtask.
	//
	// example:
	//
	// create_instance
	CurrentStepName *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	// The name of the database. If the task involves a database, the database name is returned.
	//
	// example:
	//
	// DBtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The estimated end time of the task.
	//
	// >  In most cases, this parameter is empty.
	//
	// example:
	//
	// null
	ExpectedFinishTime *string `json:"ExpectedFinishTime,omitempty" xml:"ExpectedFinishTime,omitempty"`
	// The end time of the task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-11-20T02:00Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task progress in percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The description of the task progress.
	//
	// >  If no progress description is provided for the task, this parameter is empty.
	//
	// example:
	//
	// null
	ProgressInfo *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	// The estimated remaining time of the task. Unit: seconds.
	//
	// >  If the task is not running, this parameter is not returned or the returned value is **0**.
	//
	// example:
	//
	// 60
	Remain *int32 `json:"Remain,omitempty" xml:"Remain,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The progress of the subtask. For example, a value of `1/4` indicates that the task consists of four subtasks and the first subtask is in progress.
	//
	// example:
	//
	// 1/4
	StepProgressInfo *string `json:"StepProgressInfo,omitempty" xml:"StepProgressInfo,omitempty"`
	// The details of the subtasks.
	//
	// example:
	//
	// null
	StepsInfo *string `json:"StepsInfo,omitempty" xml:"StepsInfo,omitempty"`
	// The operation that is used by the task, such as **CreateDBInstance**.
	//
	// example:
	//
	// CreateDBInstance
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The error code that is returned when an error occurs.
	//
	// >  This parameter is returned only when an error occurs.
	//
	// example:
	//
	// null
	TaskErrorCode *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	// The error message that is returned when an error occurs.
	//
	// >  This parameter is returned only when an error occurs.
	//
	// example:
	//
	// null
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// The task ID. You can use one of the following methods to obtain the task ID:
	//
	// example:
	//
	// 3472xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksResponseBodyItemsTaskProgressInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyItemsTaskProgressInfo) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetDBName() *string {
	return s.DBName
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetExpectedFinishTime() *string {
	return s.ExpectedFinishTime
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetProgress() *string {
	return s.Progress
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetProgressInfo() *string {
	return s.ProgressInfo
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetRemain() *int32 {
	return s.Remain
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetStepProgressInfo() *string {
	return s.StepProgressInfo
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetStepsInfo() *string {
	return s.StepsInfo
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetTaskErrorCode() *string {
	return s.TaskErrorCode
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetBeginTime(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetCurrentStepName(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetDBName(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.DBName = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetExpectedFinishTime(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.ExpectedFinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetFinishTime(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetProgress(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetProgressInfo(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetRemain(v int32) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.Remain = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetStatus(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.Status = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetStepProgressInfo(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.StepProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetStepsInfo(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.StepsInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetTaskAction(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetTaskErrorCode(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) SetTaskId(v string) *DescribeTasksResponseBodyItemsTaskProgressInfo {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyItemsTaskProgressInfo) Validate() error {
	return dara.Validate(s)
}
