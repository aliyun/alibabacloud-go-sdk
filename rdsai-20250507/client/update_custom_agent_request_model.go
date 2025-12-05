// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *UpdateCustomAgentRequest
	GetCustomAgentId() *string
	SetEnableTools(v bool) *UpdateCustomAgentRequest
	GetEnableTools() *bool
	SetName(v string) *UpdateCustomAgentRequest
	GetName() *string
	SetSystemPrompt(v string) *UpdateCustomAgentRequest
	GetSystemPrompt() *string
	SetTools(v []*string) *UpdateCustomAgentRequest
	GetTools() []*string
}

type UpdateCustomAgentRequest struct {
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
	EnableTools  *bool     `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	SystemPrompt *string   `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools        []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s UpdateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *UpdateCustomAgentRequest) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *UpdateCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateCustomAgentRequest) GetTools() []*string {
	return s.Tools
}

func (s *UpdateCustomAgentRequest) SetCustomAgentId(v string) *UpdateCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *UpdateCustomAgentRequest) SetEnableTools(v bool) *UpdateCustomAgentRequest {
	s.EnableTools = &v
	return s
}

func (s *UpdateCustomAgentRequest) SetName(v string) *UpdateCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomAgentRequest) SetSystemPrompt(v string) *UpdateCustomAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateCustomAgentRequest) SetTools(v []*string) *UpdateCustomAgentRequest {
	s.Tools = v
	return s
}

func (s *UpdateCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
