// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheAnalysisReportListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDailyTasks(v *DescribeCacheAnalysisReportListResponseBodyDailyTasks) *DescribeCacheAnalysisReportListResponseBody
	GetDailyTasks() *DescribeCacheAnalysisReportListResponseBodyDailyTasks
	SetInstanceId(v string) *DescribeCacheAnalysisReportListResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeCacheAnalysisReportListResponseBody
	GetRequestId() *string
}

type DescribeCacheAnalysisReportListResponseBody struct {
	// The list of the offline key analysis reports.
	DailyTasks *DescribeCacheAnalysisReportListResponseBodyDailyTasks `json:"DailyTasks,omitempty" xml:"DailyTasks,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// example:
	//
	// 1041xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 743D0A03-52DE-4E6F-8D09-EC1414CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBody) GetDailyTasks() *DescribeCacheAnalysisReportListResponseBodyDailyTasks {
	return s.DailyTasks
}

func (s *DescribeCacheAnalysisReportListResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheAnalysisReportListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetDailyTasks(v *DescribeCacheAnalysisReportListResponseBodyDailyTasks) *DescribeCacheAnalysisReportListResponseBody {
	s.DailyTasks = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetInstanceId(v string) *DescribeCacheAnalysisReportListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBody) SetRequestId(v string) *DescribeCacheAnalysisReportListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasks struct {
	DailyTask []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask `json:"DailyTask,omitempty" xml:"DailyTask,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasks) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasks) GetDailyTask() []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask {
	return s.DailyTask
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasks) SetDailyTask(v []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) *DescribeCacheAnalysisReportListResponseBodyDailyTasks {
	s.DailyTask = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask struct {
	// The date when the offline key analytics task was performed.
	//
	// example:
	//
	// 2019-08-01Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Details about the offline key analysis reports.
	Tasks *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) GetDate() *string {
	return s.Date
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) GetTasks() *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks {
	return s.Tasks
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) SetDate(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask {
	s.Date = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) SetTasks(v *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask {
	s.Tasks = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTask) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks struct {
	Task []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) GetTask() []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	return s.Task
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) SetTask(v []*DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks {
	s.Task = v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask struct {
	// The ID of the child node in the cluster instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The start time of the offline key analytics task.
	//
	// example:
	//
	// 2019-08-01T19:08:49Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the offline key analytics task. Valid values:
	//
	// 	- **success**
	//
	// 	- **running**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 156465****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetNodeId(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.NodeId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetStartTime(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.StartTime = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetStatus(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) SetTaskId(v string) *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask {
	s.TaskId = &v
	return s
}

func (s *DescribeCacheAnalysisReportListResponseBodyDailyTasksDailyTaskTasksTask) Validate() error {
	return dara.Validate(s)
}
