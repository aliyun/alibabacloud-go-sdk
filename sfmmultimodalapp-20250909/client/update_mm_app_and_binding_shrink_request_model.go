// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppAndBindingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppAndBindingShrinkRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMmAppAndBindingShrinkRequest
	GetAppName() *string
	SetBindingConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest
	GetBindingConfigShrink() *string
	SetConversationConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest
	GetConversationConfigShrink() *string
	SetMemoryConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest
	GetMemoryConfigShrink() *string
	SetModelConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest
	GetModelConfigShrink() *string
	SetPrompt(v string) *UpdateMmAppAndBindingShrinkRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *UpdateMmAppAndBindingShrinkRequest
	GetWorkspaceId() *string
}

type UpdateMmAppAndBindingShrinkRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	AppName                  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfigShrink      *string `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty"`
	ConversationConfigShrink *string `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty"`
	MemoryConfigShrink       *string `json:"MemoryConfig,omitempty" xml:"MemoryConfig,omitempty"`
	ModelConfigShrink        *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	Prompt                   *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppAndBindingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetBindingConfigShrink() *string {
	return s.BindingConfigShrink
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetConversationConfigShrink() *string {
	return s.ConversationConfigShrink
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetMemoryConfigShrink() *string {
	return s.MemoryConfigShrink
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetModelConfigShrink() *string {
	return s.ModelConfigShrink
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateMmAppAndBindingShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetAppId(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetAppName(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetBindingConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.BindingConfigShrink = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetConversationConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.ConversationConfigShrink = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetMemoryConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.MemoryConfigShrink = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetModelConfigShrink(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.ModelConfigShrink = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetPrompt(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) SetWorkspaceId(v string) *UpdateMmAppAndBindingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppAndBindingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
