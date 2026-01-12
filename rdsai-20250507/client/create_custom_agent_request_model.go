// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableTools(v bool) *CreateCustomAgentRequest
	GetEnableTools() *bool
	SetName(v string) *CreateCustomAgentRequest
	GetName() *string
	SetSystemPrompt(v string) *CreateCustomAgentRequest
	GetSystemPrompt() *string
	SetTools(v []*string) *CreateCustomAgentRequest
	GetTools() []*string
}

type CreateCustomAgentRequest struct {
	// The system prompts.
	//
	// example:
	//
	// true
	EnableTools *bool `json:"EnableTools,omitempty" xml:"EnableTools,omitempty"`
	// The operation that you want to perform. Set the value to **CreateCustomAgent**.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the dedicated agent.
	//
	// This parameter is required.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// Specifies whether to enable tools.
	Tools []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
}

func (s CreateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequest) GetEnableTools() *bool {
	return s.EnableTools
}

func (s *CreateCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateCustomAgentRequest) GetTools() []*string {
	return s.Tools
}

func (s *CreateCustomAgentRequest) SetEnableTools(v bool) *CreateCustomAgentRequest {
	s.EnableTools = &v
	return s
}

func (s *CreateCustomAgentRequest) SetName(v string) *CreateCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequest) SetSystemPrompt(v string) *CreateCustomAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateCustomAgentRequest) SetTools(v []*string) *CreateCustomAgentRequest {
	s.Tools = v
	return s
}

func (s *CreateCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
