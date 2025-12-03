// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *UpdateCustomAgentShrinkRequest
	GetApiId() *string
	SetCustomAgentId(v string) *UpdateCustomAgentShrinkRequest
	GetCustomAgentId() *string
	SetEnableTools(v bool) *UpdateCustomAgentShrinkRequest
	GetEnableTools() *bool
	SetName(v string) *UpdateCustomAgentShrinkRequest
	GetName() *string
	SetSystemPrompt(v string) *UpdateCustomAgentShrinkRequest
	GetSystemPrompt() *string
	SetToolsShrink(v string) *UpdateCustomAgentShrinkRequest
	GetToolsShrink() *string
}

type UpdateCustomAgentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNA****
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// AgentId。
	//
	// This parameter is required.
	//
	// example:
	//
	// ebe44453-3b41-4c74-94d1-01d088d7****
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// true
	EnableTools  *bool   `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	ToolsShrink  *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
}

func (s UpdateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentShrinkRequest) GetApiId() *string {
	return s.ApiId
}

func (s *UpdateCustomAgentShrinkRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *UpdateCustomAgentShrinkRequest) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *UpdateCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateCustomAgentShrinkRequest) GetToolsShrink() *string {
	return s.ToolsShrink
}

func (s *UpdateCustomAgentShrinkRequest) SetApiId(v string) *UpdateCustomAgentShrinkRequest {
	s.ApiId = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetCustomAgentId(v string) *UpdateCustomAgentShrinkRequest {
	s.CustomAgentId = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetEnableTools(v bool) *UpdateCustomAgentShrinkRequest {
	s.EnableTools = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetName(v string) *UpdateCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetSystemPrompt(v string) *UpdateCustomAgentShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) SetToolsShrink(v string) *UpdateCustomAgentShrinkRequest {
	s.ToolsShrink = &v
	return s
}

func (s *UpdateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
