// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialoguesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDialoguesResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListDialoguesResponseBody
	GetCurrent() *int32
	SetData(v []*ListDialoguesResponseBodyData) *ListDialoguesResponseBody
	GetData() []*ListDialoguesResponseBodyData
	SetHttpStatusCode(v int32) *ListDialoguesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDialoguesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDialoguesResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListDialoguesResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListDialoguesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListDialoguesResponseBody
	GetTotal() *int32
}

type ListDialoguesResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                           `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListDialoguesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDialoguesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDialoguesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDialoguesResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListDialoguesResponseBody) GetData() []*ListDialoguesResponseBodyData {
	return s.Data
}

func (s *ListDialoguesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDialoguesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDialoguesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDialoguesResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListDialoguesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDialoguesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDialoguesResponseBody) SetCode(v string) *ListDialoguesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialoguesResponseBody) SetCurrent(v int32) *ListDialoguesResponseBody {
	s.Current = &v
	return s
}

func (s *ListDialoguesResponseBody) SetData(v []*ListDialoguesResponseBodyData) *ListDialoguesResponseBody {
	s.Data = v
	return s
}

func (s *ListDialoguesResponseBody) SetHttpStatusCode(v int32) *ListDialoguesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDialoguesResponseBody) SetMessage(v string) *ListDialoguesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialoguesResponseBody) SetRequestId(v string) *ListDialoguesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSize(v int32) *ListDialoguesResponseBody {
	s.Size = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSuccess(v bool) *ListDialoguesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDialoguesResponseBody) SetTotal(v int32) *ListDialoguesResponseBody {
	s.Total = &v
	return s
}

func (s *ListDialoguesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDialoguesResponseBodyData struct {
	// example:
	//
	// xx
	Bot *string `json:"Bot,omitempty" xml:"Bot,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 2
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// example:
	//
	// xx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// x
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDialoguesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDialoguesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBodyData) GetBot() *string {
	return s.Bot
}

func (s *ListDialoguesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDialoguesResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDialoguesResponseBodyData) GetDialogueType() *int32 {
	return s.DialogueType
}

func (s *ListDialoguesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDialoguesResponseBodyData) GetUser() *string {
	return s.User
}

func (s *ListDialoguesResponseBodyData) SetBot(v string) *ListDialoguesResponseBodyData {
	s.Bot = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateTime(v string) *ListDialoguesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateUser(v string) *ListDialoguesResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetDialogueType(v int32) *ListDialoguesResponseBodyData {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetTaskId(v string) *ListDialoguesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetUser(v string) *ListDialoguesResponseBodyData {
	s.User = &v
	return s
}

func (s *ListDialoguesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
