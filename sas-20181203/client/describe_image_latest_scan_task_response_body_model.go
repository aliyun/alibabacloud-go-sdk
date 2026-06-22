// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageLatestScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeImageLatestScanTaskResponseBody
	GetRequestId() *string
	SetTask(v []*DescribeImageLatestScanTaskResponseBodyTask) *DescribeImageLatestScanTaskResponseBody
	GetTask() []*DescribeImageLatestScanTaskResponseBodyTask
}

type DescribeImageLatestScanTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task []*DescribeImageLatestScanTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeImageLatestScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageLatestScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageLatestScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageLatestScanTaskResponseBody) GetTask() []*DescribeImageLatestScanTaskResponseBodyTask {
	return s.Task
}

func (s *DescribeImageLatestScanTaskResponseBody) SetRequestId(v string) *DescribeImageLatestScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBody) SetTask(v []*DescribeImageLatestScanTaskResponseBodyTask) *DescribeImageLatestScanTaskResponseBody {
	s.Task = v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBody) Validate() error {
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

type DescribeImageLatestScanTaskResponseBodyTask struct {
	// The time when the task was created. Format: yyyy-MM-ddTHH:mm:ss.
	//
	// example:
	//
	// 2022-12-20 11:59:05
	Create *string `json:"Create,omitempty" xml:"Create,omitempty"`
	// The number of completed image tasks.
	//
	// example:
	//
	// 100
	Finish *int32 `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The time when the task ended. This parameter is returned only when the task status is Finished. Otherwise, an empty value is returned.
	//
	// example:
	//
	// 1669693430977
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 9755662
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the task was last modified. Format: yyyy-MM-ddTHH:mm:ss.
	//
	// example:
	//
	// 2022-12-20 12:00:05
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The task name.
	//
	// example:
	//
	// IMAGE_SCAN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The creation method. Valid values:
	//
	// - **console_batch**: console
	//
	// - **openapi**: API.
	//
	// example:
	//
	// console_batch
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 1668614400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **PROCESSING**: The task is being executed.
	//
	// - **START**: The task is starting.
	//
	// - **MESSAGE_SEND**: The scan is being distributed.
	//
	// - **PRE_ANALYZER**: The image is being pre-checked.
	//
	// - **SUCCESS**: The task is executed.
	//
	// - **FAIL**: The task failed.
	//
	// - **TIMOUT**: The task timed out.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The digest of the target image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d50****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The scan target type. Valid values:
	//
	// - **IMAGE**: image.
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// 0a960b9a48b788a8689154b032bf****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// - **IMAGE_SCAN**: image scan.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeImageLatestScanTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageLatestScanTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetCreate() *string {
	return s.Create
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetFinish() *int32 {
	return s.Finish
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetId() *int64 {
	return s.Id
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetModified() *string {
	return s.Modified
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetSource() *string {
	return s.Source
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetTarget() *string {
	return s.Target
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetCreate(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Create = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetFinish(v int32) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Finish = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetFinishTime(v int64) *DescribeImageLatestScanTaskResponseBodyTask {
	s.FinishTime = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetId(v int64) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Id = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetModified(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Modified = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetName(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetSource(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Source = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetStartTime(v int64) *DescribeImageLatestScanTaskResponseBodyTask {
	s.StartTime = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetStatus(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetTarget(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.Target = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetTargetType(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.TargetType = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetTaskId(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) SetTaskType(v string) *DescribeImageLatestScanTaskResponseBodyTask {
	s.TaskType = &v
	return s
}

func (s *DescribeImageLatestScanTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
