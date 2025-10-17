// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTasksResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeTasksResponseBody
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeTasksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeTasksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeTasksResponseBody
	GetStartTime() *string
	SetTasks(v *DescribeTasksResponseBodyTasks) *DescribeTasksResponseBody
	GetTasks() *DescribeTasksResponseBodyTasks
	SetTotalRecordCount(v int32) *DescribeTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeTasksResponseBody struct {
	// The ID of the cluster for which the task was created.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2020-12-02T03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4352AD99-9FF5-41A6-A319-068089******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2020-11-30T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The details of the task.
	Tasks *DescribeTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTasksResponseBody) GetEndTime() *string {
	return s.EndTime
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

func (s *DescribeTasksResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksResponseBody) GetTasks() *DescribeTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeTasksResponseBody) SetDBClusterId(v string) *DescribeTasksResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetEndTime(v string) *DescribeTasksResponseBody {
	s.EndTime = &v
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

func (s *DescribeTasksResponseBody) SetStartTime(v string) *DescribeTasksResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTasks(v *DescribeTasksResponseBodyTasks) *DescribeTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTasksResponseBodyTasks struct {
	Task []*DescribeTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasks) GetTask() []*DescribeTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeTasksResponseBodyTasks) SetTask(v []*DescribeTasksResponseBodyTasksTask) *DescribeTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeTasksResponseBodyTasks) Validate() error {
	if s.Task != nil {
		for _, item := range s.Task {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyTasksTask struct {
	// The time when the task was started. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-02T02:39:15Z
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The name of the current step.
	//
	// example:
	//
	// create_instance
	CurrentStepName *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	// The database name.
	//
	// >  This parameter is returned for only the tasks that involve database operations.
	//
	// example:
	//
	// test
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The estimated end time of the task. In most cases, this parameter is empty.
	//
	// example:
	//
	// null
	ExpectedFinishTime *string `json:"ExpectedFinishTime,omitempty" xml:"ExpectedFinishTime,omitempty"`
	// The time when the task was completed. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-02T02:40:15Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The progress of the task in percentage.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The description of the task progress. If no progress description is provided for the task, this parameter is empty.
	//
	// example:
	//
	// null
	ProgressInfo *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	// The estimated remaining duration of the task. Unit: seconds.
	//
	// example:
	//
	// 1767
	Remain *int32 `json:"Remain,omitempty" xml:"Remain,omitempty"`
	// The progress of the subtasks. For example, the value `1/4` indicates that the task consists of four subtasks and the first subtask is in progress.
	//
	// example:
	//
	// 1/4
	StepProgressInfo *string `json:"StepProgressInfo,omitempty" xml:"StepProgressInfo,omitempty"`
	// The details of the subtasks.
	//
	// example:
	//
	// [{\\"remain\\":0,\\"name\\":\\"init_task\\",\\"progress\\":100},{\\"remain\\":1764,\\"name\\":\\"create_instance\\",\\"progress\\":0},{\\"remain\\":1,\\"name\\":\\"init_cluster\\",\\"progress\\":0},{\\"remain\\":2,\\"name\\":\\"create_backup\\",\\"progress\\":0}]
	StepsInfo *string `json:"StepsInfo,omitempty" xml:"StepsInfo,omitempty"`
	// The API operation that is used by the task. Example: `CreateDBInstance`.
	//
	// example:
	//
	// CreateDBInstance
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The error code that is returned when an error occurs.
	//
	// >  This parameter is returned only when the task is in the **Stop*	- state.
	//
	// example:
	//
	// null
	TaskErrorCode *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	// The error message that is returned when an error occurs.
	//
	// >  This parameter is returned only when the task is in the **Stop*	- state.
	//
	// example:
	//
	// null
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 111111111
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasksTask) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTasksResponseBodyTasksTask) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *DescribeTasksResponseBodyTasksTask) GetDBName() *string {
	return s.DBName
}

func (s *DescribeTasksResponseBodyTasksTask) GetExpectedFinishTime() *string {
	return s.ExpectedFinishTime
}

func (s *DescribeTasksResponseBodyTasksTask) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeTasksResponseBodyTasksTask) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeTasksResponseBodyTasksTask) GetProgressInfo() *string {
	return s.ProgressInfo
}

func (s *DescribeTasksResponseBodyTasksTask) GetRemain() *int32 {
	return s.Remain
}

func (s *DescribeTasksResponseBodyTasksTask) GetStepProgressInfo() *string {
	return s.StepProgressInfo
}

func (s *DescribeTasksResponseBodyTasksTask) GetStepsInfo() *string {
	return s.StepsInfo
}

func (s *DescribeTasksResponseBodyTasksTask) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksResponseBodyTasksTask) GetTaskErrorCode() *string {
	return s.TaskErrorCode
}

func (s *DescribeTasksResponseBodyTasksTask) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *DescribeTasksResponseBodyTasksTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyTasksTask) SetBeginTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetCurrentStepName(v string) *DescribeTasksResponseBodyTasksTask {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetDBName(v string) *DescribeTasksResponseBodyTasksTask {
	s.DBName = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetExpectedFinishTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.ExpectedFinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetFinishTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetProgressInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetRemain(v int32) *DescribeTasksResponseBodyTasksTask {
	s.Remain = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetStepProgressInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.StepProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetStepsInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.StepsInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskAction(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskErrorCode(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
