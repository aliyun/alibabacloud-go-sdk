// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLPlanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSQLPlanTaskResponseBody
	GetRequestId() *string
	SetTaskList(v []*DescribeSQLPlanTaskResponseBodyTaskList) *DescribeSQLPlanTaskResponseBody
	GetTaskList() []*DescribeSQLPlanTaskResponseBodyTaskList
}

type DescribeSQLPlanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried task.
	TaskList []*DescribeSQLPlanTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLPlanTaskResponseBody) GetTaskList() []*DescribeSQLPlanTaskResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeSQLPlanTaskResponseBody) SetRequestId(v string) *DescribeSQLPlanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBody) SetTaskList(v []*DescribeSQLPlanTaskResponseBodyTaskList) *DescribeSQLPlanTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeSQLPlanTaskResponseBody) Validate() error {
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

type DescribeSQLPlanTaskResponseBodyTaskList struct {
	// The time elapsed from when the task was created to when the task was complete. Unit: milliseconds.
	//
	// example:
	//
	// 10
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The number of input rows in the task.
	//
	// example:
	//
	// 105
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The amount of input data in the task. Unit: bytes.
	//
	// example:
	//
	// 3763
	InputSize *int64 `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	// The total amount of time consumed by all operators in the task on a specific node. This parameter can be used to determine whether long tails occur in computing. Unit: milliseconds.
	//
	// example:
	//
	// 3
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The number of output rows in the task.
	//
	// example:
	//
	// 105
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The amount of output data in the task. Unit: bytes.
	//
	// example:
	//
	// 945
	OutputSize *int64 `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The peak memory of the task on a specific node. Unit: bytes.
	//
	// example:
	//
	// 898576
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The amount of time consumed to scan data from a data source in the task. Unit: milliseconds.
	//
	// example:
	//
	// 0
	ScanCost *int64 `json:"ScanCost,omitempty" xml:"ScanCost,omitempty"`
	// The number of rows scanned from a data source in the task.
	//
	// example:
	//
	// 0
	ScanRows *int64 `json:"ScanRows,omitempty" xml:"ScanRows,omitempty"`
	// The amount of data scanned from a data source in the task. Unit: bytes.
	//
	// example:
	//
	// 0
	ScanSize *int64 `json:"ScanSize,omitempty" xml:"ScanSize,omitempty"`
	// The final execution status of the task. Valid values:
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
	// The task ID.
	//
	// example:
	//
	// 198877623
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLPlanTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetInputRows() *int64 {
	return s.InputRows
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetInputSize() *int64 {
	return s.InputSize
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetOperatorCost() *int64 {
	return s.OperatorCost
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetOutputSize() *int64 {
	return s.OutputSize
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetScanCost() *int64 {
	return s.ScanCost
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetScanRows() *int64 {
	return s.ScanRows
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetScanSize() *int64 {
	return s.ScanSize
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetState() *string {
	return s.State
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetElapsedTime(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetInputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.InputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOperatorCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OperatorCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetOutputSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.OutputSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetPeakMemory(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.PeakMemory = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanCost(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanCost = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanRows(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanRows = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetScanSize(v int64) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.ScanSize = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetState(v string) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.State = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) SetTaskId(v int32) *DescribeSQLPlanTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeSQLPlanTaskResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}
