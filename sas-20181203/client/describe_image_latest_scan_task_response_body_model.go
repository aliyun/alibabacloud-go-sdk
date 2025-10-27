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
	// The ID of the request.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about the task.
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
	// The time when the task was created. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-12-20 11:59:05
	Create *string `json:"Create,omitempty" xml:"Create,omitempty"`
	// The number of images that are scanned.
	//
	// example:
	//
	// 100
	Finish *int32 `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// The end time of the task. A value is returned only when the task is in the Finished state. Otherwise, the returned value is empty.
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
	// The time when the task was last modified. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2022-12-20 12:00:05
	Modified *string `json:"Modified,omitempty" xml:"Modified,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// IMAGE_SCAN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The method in which the task was created. A task can be created in the Security Center console or by calling an API operation. Valid values:
	//
	// 	- **console_batch**: The task was created in the Security Center console.
	//
	// 	- **openapi**: The task was created by calling an API operation.
	//
	// example:
	//
	// console_batch
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 1668614400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid value:
	//
	// 	- **PROCESSING**: The task is running.
	//
	// 	- **START**: The task is being started.
	//
	// 	- **MESSAGE_SEND**: The scan task is sent.
	//
	// 	- **PRE_ANALYZER**: The image is in precheck.
	//
	// 	- **SUCCESS**: The task was successful.
	//
	// 	- **FAIL**: The task failed.
	//
	// 	- **TIMOUT**: The task timed out.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The digest value of the image.
	//
	// example:
	//
	// 8f0fbdb41d3d1ade4ffdf21558443f4c03342010563bb8c43ccc09594d50****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the scanned asset. Valid value:
	//
	// 	- **IMAGE**
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
	// The type of the task. Valid value:
	//
	// 	- **IMAGE_SCAN**
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
