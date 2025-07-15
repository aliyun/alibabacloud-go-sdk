// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSearchTasksResponseBody
	GetCode() *string
	SetData(v []*ListSearchTasksResponseBodyData) *ListSearchTasksResponseBody
	GetData() []*ListSearchTasksResponseBodyData
	SetHttpStatusCode(v int32) *ListSearchTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSearchTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListSearchTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSearchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSearchTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSearchTasksResponseBody
	GetTotalCount() *int32
}

type ListSearchTasksResponseBody struct {
	// example:
	//
	// NoData
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListSearchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSearchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSearchTasksResponseBody) GetData() []*ListSearchTasksResponseBodyData {
	return s.Data
}

func (s *ListSearchTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSearchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSearchTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSearchTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSearchTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSearchTasksResponseBody) SetCode(v string) *ListSearchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetData(v []*ListSearchTasksResponseBodyData) *ListSearchTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListSearchTasksResponseBody) SetHttpStatusCode(v int32) *ListSearchTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetMessage(v string) *ListSearchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetPageNumber(v int32) *ListSearchTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetPageSize(v int32) *ListSearchTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetRequestId(v string) *ListSearchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetSuccess(v bool) *ListSearchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListSearchTasksResponseBody) SetTotalCount(v int32) *ListSearchTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSearchTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSearchTasksResponseBodyData struct {
	// example:
	//
	// 2024-11-25 11:40:50
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 24
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// xxx
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 2024-11-25 11:40:50
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// xxxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListSearchTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSearchTasksResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSearchTasksResponseBodyData) GetDialogueType() *int32 {
	return s.DialogueType
}

func (s *ListSearchTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListSearchTasksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListSearchTasksResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSearchTasksResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *ListSearchTasksResponseBodyData) SetCreateTime(v string) *ListSearchTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) SetDialogueType(v int32) *ListSearchTasksResponseBodyData {
	s.DialogueType = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) SetTaskId(v string) *ListSearchTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) SetTaskName(v string) *ListSearchTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) SetUpdateTime(v string) *ListSearchTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) SetUsername(v string) *ListSearchTasksResponseBodyData {
	s.Username = &v
	return s
}

func (s *ListSearchTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
