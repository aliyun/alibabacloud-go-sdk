// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentContextShrink(v string) *RunSearchGenerationShrinkRequest
	GetAgentContextShrink() *string
	SetChatConfigShrink(v string) *RunSearchGenerationShrinkRequest
	GetChatConfigShrink() *string
	SetFileUrl(v string) *RunSearchGenerationShrinkRequest
	GetFileUrl() *string
	SetModelId(v string) *RunSearchGenerationShrinkRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunSearchGenerationShrinkRequest
	GetOriginalSessionId() *string
	SetPrompt(v string) *RunSearchGenerationShrinkRequest
	GetPrompt() *string
	SetTaskId(v string) *RunSearchGenerationShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunSearchGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunSearchGenerationShrinkRequest struct {
	AgentContextShrink *string `json:"AgentContext,omitempty" xml:"AgentContext,omitempty"`
	// example:
	//
	// xxx
	ChatConfigShrink *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// example:
	//
	// http://xxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// qwen-max-latest
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// xxx
	OriginalSessionId *string `json:"OriginalSessionId,omitempty" xml:"OriginalSessionId,omitempty"`
	// example:
	//
	// xxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunSearchGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationShrinkRequest) GetAgentContextShrink() *string {
	return s.AgentContextShrink
}

func (s *RunSearchGenerationShrinkRequest) GetChatConfigShrink() *string {
	return s.ChatConfigShrink
}

func (s *RunSearchGenerationShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunSearchGenerationShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunSearchGenerationShrinkRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunSearchGenerationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunSearchGenerationShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunSearchGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunSearchGenerationShrinkRequest) SetAgentContextShrink(v string) *RunSearchGenerationShrinkRequest {
	s.AgentContextShrink = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetChatConfigShrink(v string) *RunSearchGenerationShrinkRequest {
	s.ChatConfigShrink = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetFileUrl(v string) *RunSearchGenerationShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetModelId(v string) *RunSearchGenerationShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetOriginalSessionId(v string) *RunSearchGenerationShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetPrompt(v string) *RunSearchGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetTaskId(v string) *RunSearchGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) SetWorkspaceId(v string) *RunSearchGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunSearchGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
