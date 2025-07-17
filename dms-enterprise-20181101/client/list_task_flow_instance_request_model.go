// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListTaskFlowInstanceRequest
	GetDagId() *int64
	SetPageIndex(v int32) *ListTaskFlowInstanceRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListTaskFlowInstanceRequest
	GetPageSize() *int32
	SetStartTimeBegin(v string) *ListTaskFlowInstanceRequest
	GetStartTimeBegin() *string
	SetStartTimeEnd(v string) *ListTaskFlowInstanceRequest
	GetStartTimeEnd() *string
	SetStatus(v int32) *ListTaskFlowInstanceRequest
	GetStatus() *int32
	SetTid(v int64) *ListTaskFlowInstanceRequest
	GetTid() *int64
	SetTriggerType(v int32) *ListTaskFlowInstanceRequest
	GetTriggerType() *int32
	SetUseBizDate(v bool) *ListTaskFlowInstanceRequest
	GetUseBizDate() *bool
}

type ListTaskFlowInstanceRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query the execution records of the task flow. Specify the time in the yyyy-MM-DD format.
	//
	// example:
	//
	// 2022-01-07
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// The end of the time range to query the execution records of the task flow. Specify the time in the yyyy-MM-DD format.
	//
	// example:
	//
	// 2022-04-08
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// The running status of the task node. Valid values:
	//
	// - **0**: Waiting for scheduling
	//
	// - **1**: Running
	//
	// - **2**: Suspend
	//
	// - **3**: Failed to run
	//
	// - **4**: Run successfully
	//
	// - **5**: Completed
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The mode in which the task flow is triggered. Valid values:
	//
	// 	- **0**: The task flow is automatically triggered based on periodic scheduling.
	//
	// 	- **1**: The task flow is manually triggered.
	//
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// Adjust filter conditions:
	//
	// - true: StartTimeBegin and StartTimeEnd are the time range for filtering services.
	//
	// - false: StartTimeBegin and StartTimeEnd are the time range for the task to run.
	//
	// example:
	//
	// true
	UseBizDate *bool `json:"UseBizDate,omitempty" xml:"UseBizDate,omitempty"`
}

func (s ListTaskFlowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowInstanceRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListTaskFlowInstanceRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListTaskFlowInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskFlowInstanceRequest) GetStartTimeBegin() *string {
	return s.StartTimeBegin
}

func (s *ListTaskFlowInstanceRequest) GetStartTimeEnd() *string {
	return s.StartTimeEnd
}

func (s *ListTaskFlowInstanceRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListTaskFlowInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowInstanceRequest) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *ListTaskFlowInstanceRequest) GetUseBizDate() *bool {
	return s.UseBizDate
}

func (s *ListTaskFlowInstanceRequest) SetDagId(v int64) *ListTaskFlowInstanceRequest {
	s.DagId = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetPageIndex(v int32) *ListTaskFlowInstanceRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetPageSize(v int32) *ListTaskFlowInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetStartTimeBegin(v string) *ListTaskFlowInstanceRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetStartTimeEnd(v string) *ListTaskFlowInstanceRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetStatus(v int32) *ListTaskFlowInstanceRequest {
	s.Status = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetTid(v int64) *ListTaskFlowInstanceRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetTriggerType(v int32) *ListTaskFlowInstanceRequest {
	s.TriggerType = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) SetUseBizDate(v bool) *ListTaskFlowInstanceRequest {
	s.UseBizDate = &v
	return s
}

func (s *ListTaskFlowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
