// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContainerScanTaskByAppNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateContainerScanTaskByAppNameResponseBodyData) *CreateContainerScanTaskByAppNameResponseBody
	GetData() *CreateContainerScanTaskByAppNameResponseBodyData
	SetHttpStatusCode(v int32) *CreateContainerScanTaskByAppNameResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateContainerScanTaskByAppNameResponseBody
	GetRequestId() *string
}

type CreateContainerScanTaskByAppNameResponseBody struct {
	// The data returned if the request was successful.
	Data *CreateContainerScanTaskByAppNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContainerScanTaskByAppNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskByAppNameResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskByAppNameResponseBody) GetData() *CreateContainerScanTaskByAppNameResponseBodyData {
	return s.Data
}

func (s *CreateContainerScanTaskByAppNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateContainerScanTaskByAppNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContainerScanTaskByAppNameResponseBody) SetData(v *CreateContainerScanTaskByAppNameResponseBodyData) *CreateContainerScanTaskByAppNameResponseBody {
	s.Data = v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBody) SetHttpStatusCode(v int32) *CreateContainerScanTaskByAppNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBody) SetRequestId(v string) *CreateContainerScanTaskByAppNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContainerScanTaskByAppNameResponseBodyData struct {
	// Indicates whether you can create more scan tasks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CanCreate *bool `json:"CanCreate,omitempty" xml:"CanCreate,omitempty"`
	// The timestamp generated when the image information was collected. Unit: milliseconds.
	//
	// example:
	//
	// 1644286364150
	CollectTime *int64 `json:"CollectTime,omitempty" xml:"CollectTime,omitempty"`
	// The timestamp generated when the scan task started. Unit: milliseconds.
	//
	// example:
	//
	// 1644286364150
	ExecTime *int64 `json:"ExecTime,omitempty" xml:"ExecTime,omitempty"`
	// The number of container applications that are scanned.
	//
	// example:
	//
	// 5
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// The progress of the scan task in percentage.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The result of the scan task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **TASK_NOT_SUPPORT_REGION**: The task is not supported in the region where the image is deployed.
	//
	// >
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the scan task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **PRE_ANALYZER**: The task is being pre-processed.
	//
	// 	- **SUCCESS**: The task succeeds.
	//
	// 	- **FAIL**: The task fails.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// fc98d58eb56f699d49bf7ebbd6d7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of container applications that you want to scan.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateContainerScanTaskByAppNameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContainerScanTaskByAppNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetCanCreate() *bool {
	return s.CanCreate
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetCollectTime() *int64 {
	return s.CollectTime
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetExecTime() *int64 {
	return s.ExecTime
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetCanCreate(v bool) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.CanCreate = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetCollectTime(v int64) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.CollectTime = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetExecTime(v int64) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetFinishCount(v int32) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.FinishCount = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetProgress(v int32) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetResult(v string) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.Result = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetStatus(v string) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetTaskId(v string) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) SetTotalCount(v int32) *CreateContainerScanTaskByAppNameResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CreateContainerScanTaskByAppNameResponseBodyData) Validate() error {
	return dara.Validate(s)
}
