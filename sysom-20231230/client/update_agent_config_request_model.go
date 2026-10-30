// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdateAgentConfigRequest
	GetAgentId() *string
	SetConfigId(v string) *UpdateAgentConfigRequest
	GetConfigId() *string
	SetContent(v string) *UpdateAgentConfigRequest
	GetContent() *string
}

type UpdateAgentConfigRequest struct {
	// The ID of the component whose configuration is to be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the component configuration to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// f0078fbb-4213-11f0-a19b-00163e4ae208
	ConfigId *string `json:"configId,omitempty" xml:"configId,omitempty"`
	// The new content of the component configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"oomkill":{"enable":false}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateAgentConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateAgentConfigRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateAgentConfigRequest) SetAgentId(v string) *UpdateAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentConfigRequest) SetConfigId(v string) *UpdateAgentConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateAgentConfigRequest) SetContent(v string) *UpdateAgentConfigRequest {
	s.Content = &v
	return s
}

func (s *UpdateAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
