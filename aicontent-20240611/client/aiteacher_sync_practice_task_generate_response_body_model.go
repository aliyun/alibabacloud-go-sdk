// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherSyncPracticeTaskGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AITeacherSyncPracticeTaskGenerateResponseBodyData) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetData() *AITeacherSyncPracticeTaskGenerateResponseBodyData
	SetErrCode(v string) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AITeacherSyncPracticeTaskGenerateResponseBody
	GetSuccess() *bool
}

type AITeacherSyncPracticeTaskGenerateResponseBody struct {
	Data *AITeacherSyncPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetData() *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	return s.Data
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetData(v *AITeacherSyncPracticeTaskGenerateResponseBodyData) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherSyncPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AITeacherSyncPracticeTaskGenerateResponseBodyData struct {
	TaskContent []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_question_answering
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) GetTaskContent() []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	return s.TaskContent
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyData) Validate() error {
	if s.TaskContent != nil {
		for _, item := range s.TaskContent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return dara.Prettify(s)
}

func (s AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) GetAssistant() *string {
	return s.Assistant
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) GetUser() *string {
	return s.User
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

func (s *AITeacherSyncPracticeTaskGenerateResponseBodyDataTaskContent) Validate() error {
	return dara.Validate(s)
}
