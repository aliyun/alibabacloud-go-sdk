// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTools(v bool) *CreateCustomAgentShrinkRequest
	GetEnableTools() *bool
	SetName(v string) *CreateCustomAgentShrinkRequest
	GetName() *string
	SetSystemPrompt(v string) *CreateCustomAgentShrinkRequest
	GetSystemPrompt() *string
	SetToolsShrink(v string) *CreateCustomAgentShrinkRequest
	GetToolsShrink() *string
}

type CreateCustomAgentShrinkRequest struct {
	// example:
	//
	// true
	EnableTools *bool   `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	ToolsShrink  *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
}

func (s CreateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentShrinkRequest) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *CreateCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateCustomAgentShrinkRequest) GetToolsShrink() *string {
	return s.ToolsShrink
}

func (s *CreateCustomAgentShrinkRequest) SetEnableTools(v bool) *CreateCustomAgentShrinkRequest {
	s.EnableTools = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetName(v string) *CreateCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetSystemPrompt(v string) *CreateCustomAgentShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetToolsShrink(v string) *CreateCustomAgentShrinkRequest {
	s.ToolsShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
