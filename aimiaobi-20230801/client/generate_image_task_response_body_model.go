// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateImageTaskResponseBody
	GetCode() *string
	SetData(v *GenerateImageTaskResponseBodyData) *GenerateImageTaskResponseBody
	GetData() *GenerateImageTaskResponseBodyData
	SetHttpStatusCode(v int32) *GenerateImageTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateImageTaskResponseBody
	GetSuccess() *bool
}

type GenerateImageTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateImageTaskResponseBody) GetData() *GenerateImageTaskResponseBodyData {
	return s.Data
}

func (s *GenerateImageTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateImageTaskResponseBody) SetCode(v string) *GenerateImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetData(v *GenerateImageTaskResponseBodyData) *GenerateImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *GenerateImageTaskResponseBody) SetHttpStatusCode(v int32) *GenerateImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetMessage(v string) *GenerateImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetRequestId(v string) *GenerateImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetSuccess(v bool) *GenerateImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateImageTaskResponseBodyData struct {
	TaskList []*GenerateImageTaskResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s GenerateImageTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyData) GetTaskList() []*GenerateImageTaskResponseBodyDataTaskList {
	return s.TaskList
}

func (s *GenerateImageTaskResponseBodyData) SetTaskList(v []*GenerateImageTaskResponseBodyDataTaskList) *GenerateImageTaskResponseBodyData {
	s.TaskList = v
	return s
}

func (s *GenerateImageTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GenerateImageTaskResponseBodyDataTaskList struct {
	// example:
	//
	// 一直忧伤的猫
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// e1be065b-adc3-435e-bd01-1c18c5ed75d3
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskResponseBodyDataTaskList) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyDataTaskList) GetContent() *string {
	return s.Content
}

func (s *GenerateImageTaskResponseBodyDataTaskList) GetId() *int64 {
	return s.Id
}

func (s *GenerateImageTaskResponseBodyDataTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *GenerateImageTaskResponseBodyDataTaskList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetContent(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetId(v int64) *GenerateImageTaskResponseBodyDataTaskList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskId(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskStatus(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskStatus = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) Validate() error {
	return dara.Validate(s)
}
