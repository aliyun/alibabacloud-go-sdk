// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOnceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanCreate(v bool) *GenerateOnceTaskResponseBody
	GetCanCreate() *bool
	SetCollectTime(v int64) *GenerateOnceTaskResponseBody
	GetCollectTime() *int64
	SetFinishCount(v int32) *GenerateOnceTaskResponseBody
	GetFinishCount() *int32
	SetLastTask(v string) *GenerateOnceTaskResponseBody
	GetLastTask() *string
	SetRequestId(v string) *GenerateOnceTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *GenerateOnceTaskResponseBody
	GetTaskId() *string
	SetTotalCount(v int32) *GenerateOnceTaskResponseBody
	GetTotalCount() *int32
}

type GenerateOnceTaskResponseBody struct {
	// Indicates whether you can create more scan tasks. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > By default, a maximum of 10 scan tasks can be running at the same time. If 10 image scan tasks are running, you cannot create a scan task by calling this operation. You must wait for at least one of the 10 existing scan tasks to complete before you can create a scan task.
	//
	// example:
	//
	// true
	CanCreate *bool `json:"CanCreate,omitempty" xml:"CanCreate,omitempty"`
	// The collection time.
	//
	// example:
	//
	// 1670307567000
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The number of scan tasks that are complete.
	//
	// example:
	//
	// 61
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The ID of the last scan task.
	//
	// example:
	//
	// 38730bb078f4a1461d4ed283994c****
	LastTask *string `json:"LastTask,omitempty" xml:"LastTask,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8BB6B8FA-39E8-5654-A309-8EED13B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// 38730bb078f4a1461d4ed283994c****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of scan tasks.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GenerateOnceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateOnceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOnceTaskResponseBody) GetCanCreate() *bool {
	return s.CanCreate
}

func (s *GenerateOnceTaskResponseBody) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *GenerateOnceTaskResponseBody) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GenerateOnceTaskResponseBody) GetLastTask() *string {
	return s.LastTask
}

func (s *GenerateOnceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateOnceTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GenerateOnceTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GenerateOnceTaskResponseBody) SetCanCreate(v bool) *GenerateOnceTaskResponseBody {
	s.CanCreate = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetCollectTime(v int64) *GenerateOnceTaskResponseBody {
	s.CollectTime = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetFinishCount(v int32) *GenerateOnceTaskResponseBody {
	s.FinishCount = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetLastTask(v string) *GenerateOnceTaskResponseBody {
	s.LastTask = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetRequestId(v string) *GenerateOnceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetTaskId(v string) *GenerateOnceTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) SetTotalCount(v int32) *GenerateOnceTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GenerateOnceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
