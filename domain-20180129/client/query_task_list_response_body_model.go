// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryTaskListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody
	GetData() *QueryTaskListResponseBodyData
	SetNextPage(v bool) *QueryTaskListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryTaskListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryTaskListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryTaskListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryTaskListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryTaskListResponseBody
	GetTotalPageNum() *int32
}

type QueryTaskListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// true
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// example:
	//
	// 8D7D294A-8E99-481F-B64C-017EFC793059
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 43
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 22
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryTaskListResponseBody) GetData() *QueryTaskListResponseBodyData {
	return s.Data
}

func (s *QueryTaskListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTaskListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTaskListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryTaskListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryTaskListResponseBody) SetCurrentPageNum(v int32) *QueryTaskListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetData(v *QueryTaskListResponseBodyData) *QueryTaskListResponseBody {
	s.Data = v
	return s
}

func (s *QueryTaskListResponseBody) SetNextPage(v bool) *QueryTaskListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryTaskListResponseBody) SetPageSize(v int32) *QueryTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTaskListResponseBody) SetPrePage(v bool) *QueryTaskListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryTaskListResponseBody) SetRequestId(v string) *QueryTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalItemNum(v int32) *QueryTaskListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryTaskListResponseBody) SetTotalPageNum(v int32) *QueryTaskListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryTaskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTaskListResponseBodyData struct {
	TaskInfo []*QueryTaskListResponseBodyDataTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
}

func (s QueryTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyData) GetTaskInfo() []*QueryTaskListResponseBodyDataTaskInfo {
	return s.TaskInfo
}

func (s *QueryTaskListResponseBodyData) SetTaskInfo(v []*QueryTaskListResponseBodyDataTaskInfo) *QueryTaskListResponseBodyData {
	s.TaskInfo = v
	return s
}

func (s *QueryTaskListResponseBodyData) Validate() error {
	if s.TaskInfo != nil {
		for _, item := range s.TaskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTaskListResponseBodyDataTaskInfo struct {
	// example:
	//
	// 127.0.0.1
	Clientip *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	// example:
	//
	// Dec 26,2017 11:00:54
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskBizType *string `json:"TaskBizType,omitempty" xml:"TaskBizType,omitempty"`
	// example:
	//
	// INIT
	TaskCancelStatus *string `json:"TaskCancelStatus,omitempty" xml:"TaskCancelStatus,omitempty"`
	// example:
	//
	// 0
	TaskCancelStatusCode *int32 `json:"TaskCancelStatusCode,omitempty" xml:"TaskCancelStatusCode,omitempty"`
	// example:
	//
	// 8b1cd755-4928-4b02-adee-e5d41d7b1939
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	// example:
	//
	// 1
	TaskNum    *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TaskStatusCode *int32 `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	// example:
	//
	// CREATE_DNSHOST
	TaskType            *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeCode        *int32  `json:"TaskTypeCode,omitempty" xml:"TaskTypeCode,omitempty"`
	TaskTypeDescription *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
}

func (s QueryTaskListResponseBodyDataTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskListResponseBodyDataTaskInfo) GoString() string {
	return s.String()
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetClientip() *string {
	return s.Clientip
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskBizType() *string {
	return s.TaskBizType
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskCancelStatus() *string {
	return s.TaskCancelStatus
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskCancelStatusCode() *int32 {
	return s.TaskCancelStatusCode
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskResult() *string {
	return s.TaskResult
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskTypeCode() *int32 {
	return s.TaskTypeCode
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskTypeDescription() *string {
	return s.TaskTypeDescription
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetClientip(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.Clientip = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetCreateTime(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskBizType(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskBizType = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskCancelStatus(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskCancelStatus = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskCancelStatusCode(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskCancelStatusCode = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskNo(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskNo = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskNum(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskNum = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskResult(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskResult = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskStatus(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskStatusCode(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskStatusCode = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskType(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskType = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskTypeCode(v int32) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskTypeCode = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskTypeDescription(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) Validate() error {
	return dara.Validate(s)
}
