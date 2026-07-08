// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStepByStepWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginSessionId(v string) *RunStepByStepWritingShrinkRequest
	GetOriginSessionId() *string
	SetPrompt(v string) *RunStepByStepWritingShrinkRequest
	GetPrompt() *string
	SetReferenceDataShrink(v string) *RunStepByStepWritingShrinkRequest
	GetReferenceDataShrink() *string
	SetSessionId(v string) *RunStepByStepWritingShrinkRequest
	GetSessionId() *string
	SetTaskId(v string) *RunStepByStepWritingShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunStepByStepWritingShrinkRequest
	GetWorkspaceId() *string
	SetWritingConfigShrink(v string) *RunStepByStepWritingShrinkRequest
	GetWritingConfigShrink() *string
}

type RunStepByStepWritingShrinkRequest struct {
	// The ID of the original conversation when regenerating content.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// The prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The reference article data for writing.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// The ID of a single-turn conversation.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The task ID. You can reuse the same task ID for a multi-turn conversation.
	//
	// > By default, you do not need to specify this parameter. The system automatically generates a task ID. If you specify the same TaskId for subsequent tasks, the tasks are considered part of the same conversation group.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Obtain a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The writing configuration.
	WritingConfigShrink *string `json:"WritingConfig,omitempty" xml:"WritingConfig,omitempty"`
}

func (s RunStepByStepWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStepByStepWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStepByStepWritingShrinkRequest) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunStepByStepWritingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunStepByStepWritingShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunStepByStepWritingShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunStepByStepWritingShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunStepByStepWritingShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunStepByStepWritingShrinkRequest) GetWritingConfigShrink() *string {
	return s.WritingConfigShrink
}

func (s *RunStepByStepWritingShrinkRequest) SetOriginSessionId(v string) *RunStepByStepWritingShrinkRequest {
	s.OriginSessionId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetPrompt(v string) *RunStepByStepWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetReferenceDataShrink(v string) *RunStepByStepWritingShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetSessionId(v string) *RunStepByStepWritingShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetTaskId(v string) *RunStepByStepWritingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetWorkspaceId(v string) *RunStepByStepWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) SetWritingConfigShrink(v string) *RunStepByStepWritingShrinkRequest {
	s.WritingConfigShrink = &v
	return s
}

func (s *RunStepByStepWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
