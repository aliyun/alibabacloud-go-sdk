// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateMmAppShrinkRequest
	GetAppName() *string
	SetBindingConfigShrink(v string) *CreateMmAppShrinkRequest
	GetBindingConfigShrink() *string
	SetConversationConfigShrink(v string) *CreateMmAppShrinkRequest
	GetConversationConfigShrink() *string
	SetModelConfigShrink(v string) *CreateMmAppShrinkRequest
	GetModelConfigShrink() *string
	SetPrompt(v string) *CreateMmAppShrinkRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *CreateMmAppShrinkRequest
	GetWorkspaceId() *string
}

type CreateMmAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 多模态xxx
	AppName                  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfigShrink      *string `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty"`
	ConversationConfigShrink *string `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty"`
	ModelConfigShrink        *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// 提示词
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMmAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMmAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateMmAppShrinkRequest) GetBindingConfigShrink() *string {
	return s.BindingConfigShrink
}

func (s *CreateMmAppShrinkRequest) GetConversationConfigShrink() *string {
	return s.ConversationConfigShrink
}

func (s *CreateMmAppShrinkRequest) GetModelConfigShrink() *string {
	return s.ModelConfigShrink
}

func (s *CreateMmAppShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateMmAppShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMmAppShrinkRequest) SetAppName(v string) *CreateMmAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateMmAppShrinkRequest) SetBindingConfigShrink(v string) *CreateMmAppShrinkRequest {
	s.BindingConfigShrink = &v
	return s
}

func (s *CreateMmAppShrinkRequest) SetConversationConfigShrink(v string) *CreateMmAppShrinkRequest {
	s.ConversationConfigShrink = &v
	return s
}

func (s *CreateMmAppShrinkRequest) SetModelConfigShrink(v string) *CreateMmAppShrinkRequest {
	s.ModelConfigShrink = &v
	return s
}

func (s *CreateMmAppShrinkRequest) SetPrompt(v string) *CreateMmAppShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CreateMmAppShrinkRequest) SetWorkspaceId(v string) *CreateMmAppShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMmAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
