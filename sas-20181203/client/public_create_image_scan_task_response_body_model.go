// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicCreateImageScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PublicCreateImageScanTaskResponseBodyData) *PublicCreateImageScanTaskResponseBody
	GetData() *PublicCreateImageScanTaskResponseBodyData
	SetRequestId(v string) *PublicCreateImageScanTaskResponseBody
	GetRequestId() *string
}

type PublicCreateImageScanTaskResponseBody struct {
	// The data returned if the call is successful.
	Data *PublicCreateImageScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F9353221-40F4-5F98-B73C-2803DC804033
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublicCreateImageScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublicCreateImageScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PublicCreateImageScanTaskResponseBody) GetData() *PublicCreateImageScanTaskResponseBodyData {
	return s.Data
}

func (s *PublicCreateImageScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublicCreateImageScanTaskResponseBody) SetData(v *PublicCreateImageScanTaskResponseBodyData) *PublicCreateImageScanTaskResponseBody {
	s.Data = v
	return s
}

func (s *PublicCreateImageScanTaskResponseBody) SetRequestId(v string) *PublicCreateImageScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublicCreateImageScanTaskResponseBodyData struct {
	// Indicates whether you can create more image scan tasks. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// > By default, a maximum of 10 image scan tasks can be running at the same time. If 10 image scan tasks are running, you cannot create an image scan task by calling this operation. You must wait for at least one of the 10 existing image scan tasks to complete before you can create an image scan task.
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
	// The progress of the image scan task in percentage.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result of the image scan task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **TASK_NOT_SUPPORT_REGION**: The images are deployed in a region that is not supported by container image scan.
	//
	// > For more information about the regions supported by container image scan, see the "Regions supported by container image scan" section in this topic.
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
	// 	- **FAIL**: The task fails.
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

func (s PublicCreateImageScanTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublicCreateImageScanTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetCanCreate() *bool {
	return s.CanCreate
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetExecTime() *int64 {
	return s.ExecTime
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PublicCreateImageScanTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetCanCreate(v bool) *PublicCreateImageScanTaskResponseBodyData {
	s.CanCreate = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetCollectTime(v int64) *PublicCreateImageScanTaskResponseBodyData {
	s.CollectTime = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetExecTime(v int64) *PublicCreateImageScanTaskResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetFinishCount(v int32) *PublicCreateImageScanTaskResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetProgress(v int32) *PublicCreateImageScanTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetResult(v string) *PublicCreateImageScanTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetStatus(v string) *PublicCreateImageScanTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetTaskId(v string) *PublicCreateImageScanTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) SetTotalCount(v int32) *PublicCreateImageScanTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *PublicCreateImageScanTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
