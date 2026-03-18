// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAITeacherExpansionPracticeTaskGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AITeacherExpansionPracticeTaskGenerateResponseBodyData) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetData() *AITeacherExpansionPracticeTaskGenerateResponseBodyData
	SetErrCode(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AITeacherExpansionPracticeTaskGenerateResponseBody
	GetSuccess() *bool
}

type AITeacherExpansionPracticeTaskGenerateResponseBody struct {
	Data *AITeacherExpansionPracticeTaskGenerateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetData() *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	return s.Data
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetData(v *AITeacherExpansionPracticeTaskGenerateResponseBodyData) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Data = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrCode(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetErrMessage(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetHttpStatusCode(v int32) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetRequestId(v string) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) SetSuccess(v bool) *AITeacherExpansionPracticeTaskGenerateResponseBody {
	s.Success = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyData struct {
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	BackgroundDescription *string                                                        `json:"backgroundDescription,omitempty" xml:"backgroundDescription,omitempty"`
	RoleSet               *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet `json:"roleSet,omitempty" xml:"roleSet,omitempty" type:"Struct"`
	// example:
	//
	// Hey Jamie, do you know what a travel blogger does?
	StartSentence *string                                                              `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	TaskContent   []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent `json:"taskContent,omitempty" xml:"taskContent,omitempty" type:"Repeated"`
	// example:
	//
	// textbook_dialogue
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) GetBackgroundDescription() *string {
	return s.BackgroundDescription
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) GetRoleSet() *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	return s.RoleSet
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) GetStartSentence() *string {
	return s.StartSentence
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) GetTaskContent() []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	return s.TaskContent
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetBackgroundDescription(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.BackgroundDescription = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetRoleSet(v *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.RoleSet = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetStartSentence(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.StartSentence = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskContent(v []*AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskContent = v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) SetTaskType(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyData) Validate() error {
	if s.RoleSet != nil {
		if err := s.RoleSet.Validate(); err != nil {
			return err
		}
	}
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

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet struct {
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) GetAssistant() *string {
	return s.Assistant
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) GetUser() *string {
	return s.User
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet {
	s.User = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataRoleSet) Validate() error {
	return dara.Validate(s)
}

type AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent struct {
	// example:
	//
	// Why might some people think dog walking is a great job?
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// They think it\\"s great because they won\\"t be stuck in an office.
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) String() string {
	return dara.Prettify(s)
}

func (s AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) GoString() string {
	return s.String()
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) GetAssistant() *string {
	return s.Assistant
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) GetUser() *string {
	return s.User
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetAssistant(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.Assistant = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) SetUser(v string) *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent {
	s.User = &v
	return s
}

func (s *AITeacherExpansionPracticeTaskGenerateResponseBodyDataTaskContent) Validate() error {
	return dara.Validate(s)
}
