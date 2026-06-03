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
	CurrentPageNum *int32                         `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                          `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                          `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                         `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                         `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
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
	Clientip             *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskCancelStatus     *string `json:"TaskCancelStatus,omitempty" xml:"TaskCancelStatus,omitempty"`
	TaskCancelStatusCode *int32  `json:"TaskCancelStatusCode,omitempty" xml:"TaskCancelStatusCode,omitempty"`
	TaskNo               *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum              *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus           *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusCode       *int32  `json:"TaskStatusCode,omitempty" xml:"TaskStatusCode,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeDescription  *string `json:"TaskTypeDescription,omitempty" xml:"TaskTypeDescription,omitempty"`
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

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskStatusCode() *int32 {
	return s.TaskStatusCode
}

func (s *QueryTaskListResponseBodyDataTaskInfo) GetTaskType() *string {
	return s.TaskType
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

func (s *QueryTaskListResponseBodyDataTaskInfo) SetTaskTypeDescription(v string) *QueryTaskListResponseBodyDataTaskInfo {
	s.TaskTypeDescription = &v
	return s
}

func (s *QueryTaskListResponseBodyDataTaskInfo) Validate() error {
	return dara.Validate(s)
}
