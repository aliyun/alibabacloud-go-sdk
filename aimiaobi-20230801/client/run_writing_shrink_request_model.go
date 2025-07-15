// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginSessionId(v string) *RunWritingShrinkRequest
	GetOriginSessionId() *string
	SetPrompt(v string) *RunWritingShrinkRequest
	GetPrompt() *string
	SetReferenceDataShrink(v string) *RunWritingShrinkRequest
	GetReferenceDataShrink() *string
	SetSessionId(v string) *RunWritingShrinkRequest
	GetSessionId() *string
	SetTaskId(v string) *RunWritingShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunWritingShrinkRequest
	GetWorkspaceId() *string
	SetWritingConfigShrink(v string) *RunWritingShrinkRequest
	GetWritingConfigShrink() *string
}

type RunWritingShrinkRequest struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt              *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingConfigShrink *string `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty"`
}

func (s RunWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunWritingShrinkRequest) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunWritingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWritingShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunWritingShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWritingShrinkRequest) GetWritingConfigShrink() *string {
	return s.WritingConfigShrink
}

func (s *RunWritingShrinkRequest) SetOriginSessionId(v string) *RunWritingShrinkRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetPrompt(v string) *RunWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunWritingShrinkRequest) SetReferenceDataShrink(v string) *RunWritingShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunWritingShrinkRequest) SetSessionId(v string) *RunWritingShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetTaskId(v string) *RunWritingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetWorkspaceId(v string) *RunWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWritingShrinkRequest) SetWritingConfigShrink(v string) *RunWritingShrinkRequest {
	s.WritingConfigShrink = &v
	return s
}

func (s *RunWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
