// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRunningTasksResponseBody
	GetRequestId() *string
	SetResult(v []*GetRunningTasksResponseBodyResult) *GetRunningTasksResponseBody
	GetResult() []*GetRunningTasksResponseBodyResult
	SetVendorRequestId(v string) *GetRunningTasksResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetRunningTasksResponseBody
	GetVendorType() *string
}

type GetRunningTasksResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// [{}]
	Result []*GetRunningTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetRunningTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunningTasksResponseBody) GetResult() []*GetRunningTasksResponseBodyResult {
	return s.Result
}

func (s *GetRunningTasksResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetRunningTasksResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetRunningTasksResponseBody) SetRequestId(v string) *GetRunningTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunningTasksResponseBody) SetResult(v []*GetRunningTasksResponseBodyResult) *GetRunningTasksResponseBody {
	s.Result = v
	return s
}

func (s *GetRunningTasksResponseBody) SetVendorRequestId(v string) *GetRunningTasksResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetRunningTasksResponseBody) SetVendorType(v string) *GetRunningTasksResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetRunningTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRunningTasksResponseBodyResult struct {
	// example:
	//
	// 2020-01-01
	ActiveTimeGMT *string `json:"ActiveTimeGMT,omitempty" xml:"ActiveTimeGMT,omitempty"`
	// example:
	//
	// act-xxaanfaf
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 123456
	ActualActionerId *string `json:"ActualActionerId,omitempty" xml:"ActualActionerId,omitempty"`
	// example:
	//
	// 2020-01-01
	CreateTimeGMT *string `json:"CreateTimeGMT,omitempty" xml:"CreateTimeGMT,omitempty"`
	// example:
	//
	// 2020-01-01
	FinishTimeGMT *string `json:"FinishTimeGMT,omitempty" xml:"FinishTimeGMT,omitempty"`
	// example:
	//
	// 123456
	OriginatorId *string `json:"OriginatorId,omitempty" xml:"OriginatorId,omitempty"`
	// example:
	//
	// instancexxxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// append task
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// title
	TitleInEnglish *string `json:"TitleInEnglish,omitempty" xml:"TitleInEnglish,omitempty"`
}

func (s GetRunningTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponseBodyResult) GetActiveTimeGMT() *string {
	return s.ActiveTimeGMT
}

func (s *GetRunningTasksResponseBodyResult) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetRunningTasksResponseBodyResult) GetActualActionerId() *string {
	return s.ActualActionerId
}

func (s *GetRunningTasksResponseBodyResult) GetCreateTimeGMT() *string {
	return s.CreateTimeGMT
}

func (s *GetRunningTasksResponseBodyResult) GetFinishTimeGMT() *string {
	return s.FinishTimeGMT
}

func (s *GetRunningTasksResponseBodyResult) GetOriginatorId() *string {
	return s.OriginatorId
}

func (s *GetRunningTasksResponseBodyResult) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetRunningTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetRunningTasksResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetRunningTasksResponseBodyResult) GetTaskType() *string {
	return s.TaskType
}

func (s *GetRunningTasksResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *GetRunningTasksResponseBodyResult) GetTitleInEnglish() *string {
	return s.TitleInEnglish
}

func (s *GetRunningTasksResponseBodyResult) SetActiveTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActivityId(v string) *GetRunningTasksResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActualActionerId(v string) *GetRunningTasksResponseBodyResult {
	s.ActualActionerId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetCreateTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetFinishTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetOriginatorId(v string) *GetRunningTasksResponseBodyResult {
	s.OriginatorId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetProcessInstanceId(v string) *GetRunningTasksResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetStatus(v string) *GetRunningTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskId(v string) *GetRunningTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskType(v string) *GetRunningTasksResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitle(v string) *GetRunningTasksResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitleInEnglish(v string) *GetRunningTasksResponseBodyResult {
	s.TitleInEnglish = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
