// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicSyncAndCreateImageScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PublicSyncAndCreateImageScanTaskResponseBodyData) *PublicSyncAndCreateImageScanTaskResponseBody
	GetData() *PublicSyncAndCreateImageScanTaskResponseBodyData
	SetRequestId(v string) *PublicSyncAndCreateImageScanTaskResponseBody
	GetRequestId() *string
}

type PublicSyncAndCreateImageScanTaskResponseBody struct {
	// The data returned if the call is successful.
	Data *PublicSyncAndCreateImageScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F9353221-40F4-5F98-B73C-2803DC804033
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublicSyncAndCreateImageScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublicSyncAndCreateImageScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PublicSyncAndCreateImageScanTaskResponseBody) GetData() *PublicSyncAndCreateImageScanTaskResponseBodyData {
	return s.Data
}

func (s *PublicSyncAndCreateImageScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublicSyncAndCreateImageScanTaskResponseBody) SetData(v *PublicSyncAndCreateImageScanTaskResponseBodyData) *PublicSyncAndCreateImageScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBody) SetRequestId(v string) *PublicSyncAndCreateImageScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type PublicSyncAndCreateImageScanTaskResponseBodyData struct {
	// Indicates whether you can create more image scan tasks. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  By default, a maximum of 10 image scan tasks can be running at the same time. If 10 image scan tasks are running, you cannot create an image scan task by calling this operation. You must wait for at least one of the 10 existing image scan tasks to complete before you can create an image scan task.
	//
	// example:
	//
	// true
	CanCreate *bool `json:"CanCreate,omitempty" xml:"CanCreate,omitempty"`
	// The timestamp when the image information was collected. Unit: milliseconds.
	//
	// example:
	//
	// 1644286364150
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The timestamp when the image scan task started to run. Unit: milliseconds.
	//
	// example:
	//
	// 1644286364150
	ExecTime *int64 `json:"ExecTime,omitempty" xml:"ExecTime,omitempty"`
	// The number of images that have been scanned.
	//
	// example:
	//
	// 5
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The progress of the image scan task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result of the image scan task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **TASK_NOT_SUPPORT_REGION**: The image is deployed in a region that is not supported by container image scan.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the image scan task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **PRE_ANALYZER**: The task is being pre-processed.
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **FAIL**: The task failed.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// a410bb3e68c217a3368bc0238c66886d
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of images to scan.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PublicSyncAndCreateImageScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublicSyncAndCreateImageScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetCanCreate() *bool {
	return s.CanCreate
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetExecTime() *int64 {
	return s.ExecTime
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetCanCreate(v bool) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.CanCreate = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetCollectTime(v int64) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.CollectTime = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetExecTime(v int64) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetFinishCount(v int32) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetProgress(v int32) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetResult(v string) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetStatus(v string) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetTaskId(v string) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) SetTotalCount(v int32) *PublicSyncAndCreateImageScanTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
