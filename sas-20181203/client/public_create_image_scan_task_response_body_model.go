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
	// The data returned when the operation is successful.
	Data *PublicCreateImageScanTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
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
	// Indicates whether more scan tasks can be created. Valid values:
	//
	// - **true**: More scan tasks can be created.
	//
	// - **false**: No more scan tasks can be created.
	//
	// > By default, up to 10 scan tasks can exist at the same time. If the number of scan tasks exceeds 10, creating a scan task by calling this operation fails. Wait until an existing scan task is completed before creating a new scan task.
	//
	// example:
	//
	// true
	CanCreate *bool `json:"CanCreate,omitempty" xml:"CanCreate,omitempty"`
	// The timestamp when image information was collected, in milliseconds.
	//
	// example:
	//
	// 1644286364150
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The timestamp when the scan task started running, in milliseconds.
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
	// The progress percentage of the scan task.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution result of the scan task. Valid values:
	//
	// - **SUCCESS**: The scan task succeeded.
	//
	// - **TASK_NOT_SUPPORT_REGION**: The image is in a region that does not support scanning.
	//
	// > For the regions that support image security scanning, see the table of supported regions after the response parameters table in this topic.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the scan task. Valid values:
	//
	// - **INIT**: Initializing.
	//
	// - **PRE_ANALYZER**: Pre-analyzing.
	//
	// - **SUCCESS**: Succeeded.
	//
	// - **FAIL**: Failed.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the scan task.
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
