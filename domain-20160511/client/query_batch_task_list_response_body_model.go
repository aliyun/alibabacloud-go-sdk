// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryBatchTaskListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryBatchTaskListResponseBodyData) *QueryBatchTaskListResponseBody
	GetData() *QueryBatchTaskListResponseBodyData
	SetNextPage(v bool) *QueryBatchTaskListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryBatchTaskListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryBatchTaskListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryBatchTaskListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryBatchTaskListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryBatchTaskListResponseBody
	GetTotalPageNum() *int32
}

type QueryBatchTaskListResponseBody struct {
	CurrentPageNum *int32                              `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryBatchTaskListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                               `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                               `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                              `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                              `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryBatchTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryBatchTaskListResponseBody) GetData() *QueryBatchTaskListResponseBodyData {
	return s.Data
}

func (s *QueryBatchTaskListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryBatchTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBatchTaskListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryBatchTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBatchTaskListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryBatchTaskListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryBatchTaskListResponseBody) SetCurrentPageNum(v int32) *QueryBatchTaskListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetData(v *QueryBatchTaskListResponseBodyData) *QueryBatchTaskListResponseBody {
	s.Data = v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetNextPage(v bool) *QueryBatchTaskListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetPageSize(v int32) *QueryBatchTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetPrePage(v bool) *QueryBatchTaskListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetRequestId(v string) *QueryBatchTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetTotalItemNum(v int32) *QueryBatchTaskListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) SetTotalPageNum(v int32) *QueryBatchTaskListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryBatchTaskListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBatchTaskListResponseBodyData struct {
	TaskInfo []*QueryBatchTaskListResponseBodyDataTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
}

func (s QueryBatchTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskListResponseBodyData) GetTaskInfo() []*QueryBatchTaskListResponseBodyDataTaskInfo {
	return s.TaskInfo
}

func (s *QueryBatchTaskListResponseBodyData) SetTaskInfo(v []*QueryBatchTaskListResponseBodyDataTaskInfo) *QueryBatchTaskListResponseBodyData {
	s.TaskInfo = v
	return s
}

func (s *QueryBatchTaskListResponseBodyData) Validate() error {
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

type QueryBatchTaskListResponseBodyDataTaskInfo struct {
	Clientip   *string `json:"Clientip,omitempty" xml:"Clientip,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskNo     *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskNum    *int32  `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType   *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryBatchTaskListResponseBodyDataTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskListResponseBodyDataTaskInfo) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetClientip() *string {
	return s.Clientip
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetClientip(v string) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.Clientip = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetCreateTime(v string) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetTaskNo(v string) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.TaskNo = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetTaskNum(v int32) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.TaskNum = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetTaskStatus(v string) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) SetTaskType(v string) *QueryBatchTaskListResponseBodyDataTaskInfo {
	s.TaskType = &v
	return s
}

func (s *QueryBatchTaskListResponseBodyDataTaskInfo) Validate() error {
	return dara.Validate(s)
}
