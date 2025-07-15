// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAskUser(v string) *RunCustomHotTopicAnalysisRequest
	GetAskUser() *string
	SetForceAnalysisExistsTopic(v bool) *RunCustomHotTopicAnalysisRequest
	GetForceAnalysisExistsTopic() *bool
	SetPrompt(v string) *RunCustomHotTopicAnalysisRequest
	GetPrompt() *string
	SetSessionId(v string) *RunCustomHotTopicAnalysisRequest
	GetSessionId() *string
	SetTaskId(v string) *RunCustomHotTopicAnalysisRequest
	GetTaskId() *string
	SetUserBack(v string) *RunCustomHotTopicAnalysisRequest
	GetUserBack() *string
	SetWorkspaceId(v string) *RunCustomHotTopicAnalysisRequest
	GetWorkspaceId() *string
}

type RunCustomHotTopicAnalysisRequest struct {
	// example:
	//
	// 模型反问
	AskUser *string `json:"AskUser,omitempty" xml:"AskUser,omitempty"`
	// example:
	//
	// false
	ForceAnalysisExistsTopic *bool `json:"ForceAnalysisExistsTopic,omitempty" xml:"ForceAnalysisExistsTopic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 用户输入Prompt
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 用户针对模型反问的输入
	UserBack *string `json:"UserBack,omitempty" xml:"UserBack,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunCustomHotTopicAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisRequest) GetAskUser() *string {
	return s.AskUser
}

func (s *RunCustomHotTopicAnalysisRequest) GetForceAnalysisExistsTopic() *bool {
	return s.ForceAnalysisExistsTopic
}

func (s *RunCustomHotTopicAnalysisRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunCustomHotTopicAnalysisRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCustomHotTopicAnalysisRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCustomHotTopicAnalysisRequest) GetUserBack() *string {
	return s.UserBack
}

func (s *RunCustomHotTopicAnalysisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunCustomHotTopicAnalysisRequest) SetAskUser(v string) *RunCustomHotTopicAnalysisRequest {
	s.AskUser = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetForceAnalysisExistsTopic(v bool) *RunCustomHotTopicAnalysisRequest {
	s.ForceAnalysisExistsTopic = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetPrompt(v string) *RunCustomHotTopicAnalysisRequest {
	s.Prompt = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetSessionId(v string) *RunCustomHotTopicAnalysisRequest {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetTaskId(v string) *RunCustomHotTopicAnalysisRequest {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetUserBack(v string) *RunCustomHotTopicAnalysisRequest {
	s.UserBack = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) SetWorkspaceId(v string) *RunCustomHotTopicAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
