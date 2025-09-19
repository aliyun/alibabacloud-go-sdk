// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateContainerScanTaskResponseBodyData) *CreateContainerScanTaskResponseBody
	GetData() *CreateContainerScanTaskResponseBodyData
	SetHttpStatusCode(v int32) *CreateContainerScanTaskResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateContainerScanTaskResponseBody
	GetRequestId() *string
}

type CreateContainerScanTaskResponseBody struct {
	// The data returned.
	Data *CreateContainerScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9F4E6157-9600-5588-86B9-38F09067****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContainerScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskResponseBody) GetData() *CreateContainerScanTaskResponseBodyData {
	return s.Data
}

func (s *CreateContainerScanTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateContainerScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContainerScanTaskResponseBody) SetData(v *CreateContainerScanTaskResponseBodyData) *CreateContainerScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateContainerScanTaskResponseBody) SetHttpStatusCode(v int32) *CreateContainerScanTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateContainerScanTaskResponseBody) SetRequestId(v string) *CreateContainerScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContainerScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateContainerScanTaskResponseBodyData struct {
	// Indicates whether you can create more scan tasks.
	//
	// example:
	//
	// true
	CanCreate *bool `json:"CanCreate,omitempty" xml:"CanCreate,omitempty"`
	// The collection time.
	//
	// example:
	//
	// 1644286364150
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The execution time of the task.
	//
	// example:
	//
	// 1644286364150
	ExecTime *int64 `json:"ExecTime,omitempty" xml:"ExecTime,omitempty"`
	// The number of scan tasks that are complete.
	//
	// example:
	//
	// 33
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The progress of the task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution result of the task.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// fc98d58eb56f699d49bf7ebbd6d7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of scan tasks.
	//
	// example:
	//
	// 62
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateContainerScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskResponseBodyData) GetCanCreate() *bool {
	return s.CanCreate
}

func (s *CreateContainerScanTaskResponseBodyData) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *CreateContainerScanTaskResponseBodyData) GetExecTime() *int64 {
	return s.ExecTime
}

func (s *CreateContainerScanTaskResponseBodyData) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *CreateContainerScanTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateContainerScanTaskResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *CreateContainerScanTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateContainerScanTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateContainerScanTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateContainerScanTaskResponseBodyData) SetCanCreate(v bool) *CreateContainerScanTaskResponseBodyData {
	s.CanCreate = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetCollectTime(v int64) *CreateContainerScanTaskResponseBodyData {
	s.CollectTime = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetExecTime(v int64) *CreateContainerScanTaskResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetFinishCount(v int32) *CreateContainerScanTaskResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetProgress(v int32) *CreateContainerScanTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetResult(v string) *CreateContainerScanTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetStatus(v string) *CreateContainerScanTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetTaskId(v string) *CreateContainerScanTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) SetTotalCount(v int32) *CreateContainerScanTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CreateContainerScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
