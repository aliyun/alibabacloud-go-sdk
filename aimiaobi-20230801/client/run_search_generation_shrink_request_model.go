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
	// Context.
	AgentContextShrink *string `json:"AgentContext,omitempty" xml:"AgentContext,omitempty"`
	// Session configuration.
	//
	// example:
	//
	// xxx
	ChatConfigShrink *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// Image URL. Used for image search and hybrid text-and-image (prompt) search generation.
	//
	// example:
	//
	// http://xxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Model ID:
	//
	// - quanmiao-max: Quanmiao-Max
	//
	// - quanmiao-plus: Quanmiao-Plus
	//
	// example:
	//
	// quanmiao-max
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// Original session identifier. Usually empty. When non-empty, it indicates that the current conversation is based on the referenced session. The system loads parameters and search results from that session and replaces the generated result. Use this for re-generation, changing data sources, or adding new agents.
	//
	// example:
	//
	// xxx
	OriginalSessionId *string `json:"OriginalSessionId,omitempty" xml:"OriginalSessionId,omitempty"`
	// Search query.
	//
	// example:
	//
	// 杭州亚运会吉祥物是什么
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Unique identifier for the session task.
	//
	// > By default, you do not need to provide a TaskId. The system generates one automatically. If you specify the same TaskId in subsequent requests, those tasks are grouped into the same conversation.
	//
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// ID of the Alibaba Cloud Model Studio workspace. To learn how to obtain this ID, see [How to use workspaces](https://help.aliyun.com/document_detail/2782167.html).
	//
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
