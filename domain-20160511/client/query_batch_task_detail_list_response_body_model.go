// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBatchTaskDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *QueryBatchTaskDetailListResponseBody
	GetCurrentPageNum() *int32
	SetData(v *QueryBatchTaskDetailListResponseBodyData) *QueryBatchTaskDetailListResponseBody
	GetData() *QueryBatchTaskDetailListResponseBodyData
	SetNextPage(v bool) *QueryBatchTaskDetailListResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *QueryBatchTaskDetailListResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *QueryBatchTaskDetailListResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *QueryBatchTaskDetailListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *QueryBatchTaskDetailListResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *QueryBatchTaskDetailListResponseBody
	GetTotalPageNum() *int32
}

type QueryBatchTaskDetailListResponseBody struct {
	CurrentPageNum *int32                                    `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryBatchTaskDetailListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	NextPage       *bool                                     `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	PageSize       *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrePage        *bool                                     `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                    `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                    `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryBatchTaskDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskDetailListResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryBatchTaskDetailListResponseBody) GetData() *QueryBatchTaskDetailListResponseBodyData {
	return s.Data
}

func (s *QueryBatchTaskDetailListResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *QueryBatchTaskDetailListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBatchTaskDetailListResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *QueryBatchTaskDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBatchTaskDetailListResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryBatchTaskDetailListResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryBatchTaskDetailListResponseBody) SetCurrentPageNum(v int32) *QueryBatchTaskDetailListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetData(v *QueryBatchTaskDetailListResponseBodyData) *QueryBatchTaskDetailListResponseBody {
	s.Data = v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetNextPage(v bool) *QueryBatchTaskDetailListResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetPageSize(v int32) *QueryBatchTaskDetailListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetPrePage(v bool) *QueryBatchTaskDetailListResponseBody {
	s.PrePage = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetRequestId(v string) *QueryBatchTaskDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetTotalItemNum(v int32) *QueryBatchTaskDetailListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) SetTotalPageNum(v int32) *QueryBatchTaskDetailListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBatchTaskDetailListResponseBodyData struct {
	TaskDetail []*QueryBatchTaskDetailListResponseBodyDataTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Repeated"`
}

func (s QueryBatchTaskDetailListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskDetailListResponseBodyData) GetTaskDetail() []*QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	return s.TaskDetail
}

func (s *QueryBatchTaskDetailListResponseBodyData) SetTaskDetail(v []*QueryBatchTaskDetailListResponseBodyDataTaskDetail) *QueryBatchTaskDetailListResponseBodyData {
	s.TaskDetail = v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyData) Validate() error {
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

type QueryBatchTaskDetailListResponseBodyDataTaskDetail struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	TaskNo     *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType   *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TryCount   *int32  `json:"TryCount,omitempty" xml:"TryCount,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryBatchTaskDetailListResponseBodyDataTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryBatchTaskDetailListResponseBodyDataTaskDetail) GoString() string {
	return s.String()
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetTaskNo() *string {
	return s.TaskNo
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetTryCount() *int32 {
	return s.TryCount
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetDomainName(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.DomainName = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetErrorMsg(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.ErrorMsg = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetTaskNo(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.TaskNo = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetTaskStatus(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetTaskType(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.TaskType = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetTryCount(v int32) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.TryCount = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) SetUpdateTime(v string) *QueryBatchTaskDetailListResponseBodyDataTaskDetail {
	s.UpdateTime = &v
	return s
}

func (s *QueryBatchTaskDetailListResponseBodyDataTaskDetail) Validate() error {
	return dara.Validate(s)
}
