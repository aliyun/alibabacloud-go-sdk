// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiagnosisTasksResponseBody
	GetRequestId() *string
	SetTaskList(v []*DescribeDiagnosisTasksResponseBodyTaskList) *DescribeDiagnosisTasksResponseBody
	GetTaskList() []*DescribeDiagnosisTasksResponseBodyTaskList
	SetTotalCount(v int32) *DescribeDiagnosisTasksResponseBody
	GetTotalCount() *int32
}

type DescribeDiagnosisTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F0983B43-B2EC-536A-8791-142B5CF1E9B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tasks.
	TaskList []*DescribeDiagnosisTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of tasks in the stage.
	//
	// example:
	//
	// 33
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisTasksResponseBody) GetTaskList() []*DescribeDiagnosisTasksResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeDiagnosisTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiagnosisTasksResponseBody) SetRequestId(v string) *DescribeDiagnosisTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBody) SetTaskList(v []*DescribeDiagnosisTasksResponseBodyTaskList) *DescribeDiagnosisTasksResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeDiagnosisTasksResponseBody) SetTotalCount(v int32) *DescribeDiagnosisTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosisTasksResponseBodyTaskList struct {
	// The compute time ratio, which can be used to determine whether the task is really time-consuming. This parameter can be calculated by using the following formula: OperatorCost/Drivers/ElapsedTime. A greater value indicates that the task was executed for computing for most of the task time. A less value indicates that the task was waiting for scheduling or blocked due to other reasons for most of the task time.
	//
	// example:
	//
	// 0.89
	ComputeTimeRatio *string `json:"ComputeTimeRatio,omitempty" xml:"ComputeTimeRatio,omitempty"`
	// The number of tasks that can be executed concurrently.
	//
	// example:
	//
	// 16
	Drivers *string `json:"Drivers,omitempty" xml:"Drivers,omitempty"`
	// The amount of time that elapsed from when the task was created to when the task was completed. Unit: milliseconds.
	//
	// example:
	//
	// 456
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The amount of input data in the task. Unit: bytes.
	//
	// example:
	//
	// 123
	InputDataSize *int64 `json:"InputDataSize,omitempty" xml:"InputDataSize,omitempty"`
	// The number of input rows in the task.
	//
	// example:
	//
	// 105
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The total amount of time that is consumed by all operators in the task on a node. This parameter can be used to determine whether long tails occur in computing. Unit: milliseconds.
	//
	// example:
	//
	// 3
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The amount of output data in the task. Unit: bytes.
	//
	// example:
	//
	// 123
	OutputDataSize *int64 `json:"OutputDataSize,omitempty" xml:"OutputDataSize,omitempty"`
	// The number of output rows in the task.
	//
	// example:
	//
	// 105
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The peak memory of the task. Unit: bytes.
	//
	// example:
	//
	// 234
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The queuing duration of the task. Unit: milliseconds.
	//
	// example:
	//
	// 12
	QueuedTime *string `json:"QueuedTime,omitempty" xml:"QueuedTime,omitempty"`
	// The amount of time that is consumed to scan data from a data source in the task. Unit: milliseconds.
	//
	// example:
	//
	// 0
	ScanCost *int64 `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	// The amount of scanned data in the task. Unit: bytes.
	//
	// example:
	//
	// 123
	ScanDataSize *int64 `json:"ScanDataSize,omitempty" xml:"ScanDataSize,omitempty"`
	// The number of rows that are scanned from a data source in the task.
	//
	// example:
	//
	// 0
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The final execution state of the task. Valid values:
	//
	// 	- FINISHED
	//
	// 	- CANCELED
	//
	// 	- ABORTED
	//
	// 	- FAILED
	//
	// example:
	//
	// FINISHED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timestamp when the task was created.
	//
	// example:
	//
	// 2022-12-12 00:00:12
	TaskCreateTime *int64 `json:"TaskCreateTime,omitempty" xml:"TaskCreateTime,omitempty"`
	// The timestamp when the task ends.
	//
	// example:
	//
	// 2022-12-22 00:00:00
	TaskEndTime *int64 `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	// The IP address of the host where the task was executed.
	//
	// example:
	//
	// 192.168.XX.XX
	TaskHost *string `json:"TaskHost,omitempty" xml:"TaskHost,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 22568****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDiagnosisTasksResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetComputeTimeRatio() *string {
	return s.ComputeTimeRatio
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetDrivers() *string {
	return s.Drivers
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetInputDataSize() *int64 {
	return s.InputDataSize
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetInputRows() *int64 {
	return s.InputRows
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetOperatorCost() *int64 {
	return s.OperatorCost
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetOutputDataSize() *int64 {
	return s.OutputDataSize
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetQueuedTime() *string {
	return s.QueuedTime
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetScanCost() *int64 {
	return s.ScanCost
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetScanDataSize() *int64 {
	return s.ScanDataSize
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetState() *string {
	return s.State
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetTaskCreateTime() *int64 {
	return s.TaskCreateTime
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetTaskEndTime() *int64 {
	return s.TaskEndTime
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetTaskHost() *string {
	return s.TaskHost
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetComputeTimeRatio(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ComputeTimeRatio = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetDrivers(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.Drivers = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetElapsedTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetInputDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.InputDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetInputRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.InputRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOperatorCost(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOutputDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OutputDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetOutputRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.OutputRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetPeakMemory(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetQueuedTime(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.QueuedTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanCost(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanDataSize(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanDataSize = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetScanRows(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetState(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.State = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskCreateTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskCreateTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskEndTime(v int64) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskEndTime = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskHost(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskHost = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) SetTaskId(v string) *DescribeDiagnosisTasksResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeDiagnosisTasksResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
