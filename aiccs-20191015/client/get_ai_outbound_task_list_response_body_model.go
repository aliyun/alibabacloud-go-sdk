// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAiOutboundTaskListResponseBody
	GetCode() *string
	SetData(v *GetAiOutboundTaskListResponseBodyData) *GetAiOutboundTaskListResponseBody
	GetData() *GetAiOutboundTaskListResponseBodyData
	SetMessage(v string) *GetAiOutboundTaskListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAiOutboundTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAiOutboundTaskListResponseBody
	GetSuccess() *bool
}

type GetAiOutboundTaskListResponseBody struct {
	// Status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Job list.
	Data *GetAiOutboundTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAiOutboundTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAiOutboundTaskListResponseBody) GetData() *GetAiOutboundTaskListResponseBodyData {
	return s.Data
}

func (s *GetAiOutboundTaskListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAiOutboundTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiOutboundTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAiOutboundTaskListResponseBody) SetCode(v string) *GetAiOutboundTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBody) SetData(v *GetAiOutboundTaskListResponseBodyData) *GetAiOutboundTaskListResponseBody {
	s.Data = v
	return s
}

func (s *GetAiOutboundTaskListResponseBody) SetMessage(v string) *GetAiOutboundTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBody) SetRequestId(v string) *GetAiOutboundTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBody) SetSuccess(v bool) *GetAiOutboundTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiOutboundTaskListResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Indicates whether there is a next page. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	HasNextPage *bool `json:"HasNextPage,omitempty" xml:"HasNextPage,omitempty"`
	// Job information.
	List []*GetAiOutboundTaskListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of data entries.
	//
	// example:
	//
	// 10
	TotalResults *int32 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s GetAiOutboundTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskListResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAiOutboundTaskListResponseBodyData) GetHasNextPage() *bool {
	return s.HasNextPage
}

func (s *GetAiOutboundTaskListResponseBodyData) GetList() []*GetAiOutboundTaskListResponseBodyDataList {
	return s.List
}

func (s *GetAiOutboundTaskListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAiOutboundTaskListResponseBodyData) GetTotalResults() *int32 {
	return s.TotalResults
}

func (s *GetAiOutboundTaskListResponseBodyData) SetCurrentPage(v int32) *GetAiOutboundTaskListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyData) SetHasNextPage(v bool) *GetAiOutboundTaskListResponseBodyData {
	s.HasNextPage = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyData) SetList(v []*GetAiOutboundTaskListResponseBodyDataList) *GetAiOutboundTaskListResponseBodyData {
	s.List = v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyData) SetPageSize(v int32) *GetAiOutboundTaskListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyData) SetTotalResults(v int32) *GetAiOutboundTaskListResponseBodyData {
	s.TotalResults = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiOutboundTaskListResponseBodyDataList struct {
	// Concurrent outbound calls.
	//
	// > Parameter specific to auto dialing.
	//
	// example:
	//
	// 1
	ConcurrentRate *int32 `json:"ConcurrentRate,omitempty" xml:"ConcurrentRate,omitempty"`
	// The deadline for job creation. UNIX timestamp format, in milliseconds.
	//
	// example:
	//
	// 1615083365000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Job description.
	//
	// example:
	//
	// 房产销售
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of completed tasks.
	//
	// example:
	//
	// 70
	FinishCount *int32 `json:"FinishCount,omitempty" xml:"FinishCount,omitempty"`
	// Job completion rate.
	//
	// example:
	//
	// 0.7
	FinishRate *float32 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// Skill group ID (for predictive dialing) or IVR ID (for auto dialing).
	//
	// example:
	//
	// 2468****
	HandlerId *int64 `json:"HandlerId,omitempty" xml:"HandlerId,omitempty"`
	// Skill group name or IVR name.
	//
	// example:
	//
	// 热线技能组
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	// Job name.
	//
	// example:
	//
	// xxxx外呼任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Task Status. Valid values:
	//
	// - **0**: Not started.
	//
	// - **1**: In progress.
	//
	// - **2**: System paused.
	//
	// - **3**: Manually paused.
	//
	// - **4**: Completed.
	//
	// - **5**: Stopped.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task status description.
	//
	// example:
	//
	// 未开始
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 1763****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Total number of jobs.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAiOutboundTaskListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetConcurrentRate() *int32 {
	return s.ConcurrentRate
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetFinishCount() *int32 {
	return s.FinishCount
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetFinishRate() *float32 {
	return s.FinishRate
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetHandlerId() *int64 {
	return s.HandlerId
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetHandlerName() *string {
	return s.HandlerName
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskListResponseBodyDataList) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetConcurrentRate(v int32) *GetAiOutboundTaskListResponseBodyDataList {
	s.ConcurrentRate = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetCreateTime(v int64) *GetAiOutboundTaskListResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetDescription(v string) *GetAiOutboundTaskListResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetFinishCount(v int32) *GetAiOutboundTaskListResponseBodyDataList {
	s.FinishCount = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetFinishRate(v float32) *GetAiOutboundTaskListResponseBodyDataList {
	s.FinishRate = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetHandlerId(v int64) *GetAiOutboundTaskListResponseBodyDataList {
	s.HandlerId = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetHandlerName(v string) *GetAiOutboundTaskListResponseBodyDataList {
	s.HandlerName = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetName(v string) *GetAiOutboundTaskListResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetStatus(v int32) *GetAiOutboundTaskListResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetStatusDesc(v string) *GetAiOutboundTaskListResponseBodyDataList {
	s.StatusDesc = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetTaskId(v int64) *GetAiOutboundTaskListResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) SetTotalCount(v int32) *GetAiOutboundTaskListResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *GetAiOutboundTaskListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
