// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoClipsTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAutoClipsTaskResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListAutoClipsTaskResponseBody
	GetCurrent() *int32
	SetData(v []*ListAutoClipsTaskResponseBodyData) *ListAutoClipsTaskResponseBody
	GetData() []*ListAutoClipsTaskResponseBodyData
	SetHttpStatusCode(v int32) *ListAutoClipsTaskResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListAutoClipsTaskResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAutoClipsTaskResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAutoClipsTaskResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAutoClipsTaskResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListAutoClipsTaskResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListAutoClipsTaskResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAutoClipsTaskResponseBody
	GetTotal() *int32
	SetTotalCount(v int32) *ListAutoClipsTaskResponseBody
	GetTotalCount() *int32
}

type ListAutoClipsTaskResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                               `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListAutoClipsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// null
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// null
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoClipsTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoClipsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoClipsTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAutoClipsTaskResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAutoClipsTaskResponseBody) GetData() []*ListAutoClipsTaskResponseBodyData {
	return s.Data
}

func (s *ListAutoClipsTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAutoClipsTaskResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoClipsTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAutoClipsTaskResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoClipsTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoClipsTaskResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListAutoClipsTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutoClipsTaskResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAutoClipsTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAutoClipsTaskResponseBody) SetCode(v string) *ListAutoClipsTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetCurrent(v int32) *ListAutoClipsTaskResponseBody {
	s.Current = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetData(v []*ListAutoClipsTaskResponseBodyData) *ListAutoClipsTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetHttpStatusCode(v int32) *ListAutoClipsTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetMaxResults(v int32) *ListAutoClipsTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetMessage(v string) *ListAutoClipsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetNextToken(v string) *ListAutoClipsTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetRequestId(v string) *ListAutoClipsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetSize(v int32) *ListAutoClipsTaskResponseBody {
	s.Size = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetSuccess(v bool) *ListAutoClipsTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetTotal(v int32) *ListAutoClipsTaskResponseBody {
	s.Total = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) SetTotalCount(v int32) *ListAutoClipsTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAutoClipsTaskResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoClipsTaskResponseBodyData struct {
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// upload
	//
	// example:
	//
	// upload
	//
	// clips
	//
	// generate
	TaskStep *string `json:"TaskStep,omitempty" xml:"TaskStep,omitempty"`
	// example:
	//
	// type001
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListAutoClipsTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutoClipsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoClipsTaskResponseBodyData) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListAutoClipsTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAutoClipsTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAutoClipsTaskResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAutoClipsTaskResponseBodyData) GetTaskStep() *string {
	return s.TaskStep
}

func (s *ListAutoClipsTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAutoClipsTaskResponseBodyData) SetCreateTimeStart(v string) *ListAutoClipsTaskResponseBodyData {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) SetTaskId(v string) *ListAutoClipsTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) SetTaskName(v string) *ListAutoClipsTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) SetTaskStatus(v int32) *ListAutoClipsTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) SetTaskStep(v string) *ListAutoClipsTaskResponseBodyData {
	s.TaskStep = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) SetTaskType(v string) *ListAutoClipsTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListAutoClipsTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
