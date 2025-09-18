// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppShrinkRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMmAppShrinkRequest
	GetAppName() *string
	SetBindingConfigShrink(v string) *UpdateMmAppShrinkRequest
	GetBindingConfigShrink() *string
	SetConversationConfigShrink(v string) *UpdateMmAppShrinkRequest
	GetConversationConfigShrink() *string
	SetModelConfigShrink(v string) *UpdateMmAppShrinkRequest
	GetModelConfigShrink() *string
	SetPrompt(v string) *UpdateMmAppShrinkRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *UpdateMmAppShrinkRequest
	GetWorkspaceId() *string
}

type UpdateMmAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多模态应用xxxxx
	AppName                  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfigShrink      *string `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty"`
	ConversationConfigShrink *string `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty"`
	ModelConfigShrink        *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	// example:
	//
	// 提示词，不超过8000字符
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMmAppShrinkRequest) GetBindingConfigShrink() *string {
	return s.BindingConfigShrink
}

func (s *UpdateMmAppShrinkRequest) GetConversationConfigShrink() *string {
	return s.ConversationConfigShrink
}

func (s *UpdateMmAppShrinkRequest) GetModelConfigShrink() *string {
	return s.ModelConfigShrink
}

func (s *UpdateMmAppShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateMmAppShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppShrinkRequest) SetAppId(v string) *UpdateMmAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetAppName(v string) *UpdateMmAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetBindingConfigShrink(v string) *UpdateMmAppShrinkRequest {
	s.BindingConfigShrink = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetConversationConfigShrink(v string) *UpdateMmAppShrinkRequest {
	s.ConversationConfigShrink = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetModelConfigShrink(v string) *UpdateMmAppShrinkRequest {
	s.ModelConfigShrink = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetPrompt(v string) *UpdateMmAppShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) SetWorkspaceId(v string) *UpdateMmAppShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
