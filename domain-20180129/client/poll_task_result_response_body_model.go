// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *PollTaskResultResponseBody
	GetCurrentPageNum() *int32
	SetData(v *PollTaskResultResponseBodyData) *PollTaskResultResponseBody
	GetData() *PollTaskResultResponseBodyData
	SetNextPage(v bool) *PollTaskResultResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *PollTaskResultResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *PollTaskResultResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *PollTaskResultResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *PollTaskResultResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *PollTaskResultResponseBody
	GetTotalPageNum() *int32
}

type PollTaskResultResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                          `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *PollTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// false
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// E879DC07-38EE-4408-9F33-73B30CD965CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 10
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s PollTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PollTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *PollTaskResultResponseBody) GetData() *PollTaskResultResponseBodyData {
	return s.Data
}

func (s *PollTaskResultResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *PollTaskResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PollTaskResultResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *PollTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PollTaskResultResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *PollTaskResultResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *PollTaskResultResponseBody) SetCurrentPageNum(v int32) *PollTaskResultResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *PollTaskResultResponseBody) SetData(v *PollTaskResultResponseBodyData) *PollTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *PollTaskResultResponseBody) SetNextPage(v bool) *PollTaskResultResponseBody {
	s.NextPage = &v
	return s
}

func (s *PollTaskResultResponseBody) SetPageSize(v int32) *PollTaskResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *PollTaskResultResponseBody) SetPrePage(v bool) *PollTaskResultResponseBody {
	s.PrePage = &v
	return s
}

func (s *PollTaskResultResponseBody) SetRequestId(v string) *PollTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *PollTaskResultResponseBody) SetTotalItemNum(v int32) *PollTaskResultResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *PollTaskResultResponseBody) SetTotalPageNum(v int32) *PollTaskResultResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *PollTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PollTaskResultResponseBodyData struct {
	TaskDetail []*PollTaskResultResponseBodyDataTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Repeated"`
}

func (s PollTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PollTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBodyData) GetTaskDetail() []*PollTaskResultResponseBodyDataTaskDetail {
	return s.TaskDetail
}

func (s *PollTaskResultResponseBodyData) SetTaskDetail(v []*PollTaskResultResponseBodyDataTaskDetail) *PollTaskResultResponseBodyData {
	s.TaskDetail = v
	return s
}

func (s *PollTaskResultResponseBodyData) Validate() error {
	if s.TaskDetail != nil {
		for _, item := range s.TaskDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PollTaskResultResponseBodyDataTaskDetail struct {
	// example:
	//
	// 2018-03-26 15:08:20
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// The operation is successful.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// S201817141000000
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 15fee9d10d514bada66bd08c5723c583
	TaskDetailNo *string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty"`
	// example:
	//
	// b95bc334-f7d8-4f39-8a62-4c4302a243d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// test
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// EXECUTE_SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 2
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CHG_DNS
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
	// example:
	//
	// 0
	TryCount *int32 `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	// example:
	//
	// 2018-03-26 15:22:18
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s PollTaskResultResponseBodyDataTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s PollTaskResultResponseBodyDataTaskDetail) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskDetailNo() *string {
	return s.TaskDetailNo
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskNo() *string {
	return s.TaskNo
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskResult() *string {
	return s.TaskResult
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskType() *string {
	return s.TaskType
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetTryCount() *int32 {
	return s.TryCount
}

func (s *PollTaskResultResponseBodyDataTaskDetail) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetCreateTime(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetDomainName(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.DomainName = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetErrorMsg(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.ErrorMsg = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetInstanceId(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.InstanceId = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskDetailNo(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskDetailNo = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskNo(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskNo = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskResult(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskResult = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskStatus(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskStatusCode(v int32) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskStatusCode = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskType(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskType = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTaskTypeDescription(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.TaskTypeDescription = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetTryCount(v int32) *PollTaskResultResponseBodyDataTaskDetail {
	s.TryCount = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) SetUpdateTime(v string) *PollTaskResultResponseBodyDataTaskDetail {
	s.UpdateTime = &v
	return s
}

func (s *PollTaskResultResponseBodyDataTaskDetail) Validate() error {
	return dara.Validate(s)
}
