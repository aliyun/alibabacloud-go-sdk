// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateAgentConfigRequest
	GetAgentId() *string
	SetConfigName(v string) *CreateAgentConfigRequest
	GetConfigName() *string
	SetContent(v string) *CreateAgentConfigRequest
	GetContent() *string
}

type CreateAgentConfigRequest struct {
	// The ID of the agent associated with the configuration. You can obtain the agent ID from ListAgents.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The name of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// Configuration 1
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// The content of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"oomkill":{"enable":false}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CreateAgentConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentConfigRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *CreateAgentConfigRequest) GetContent() *string {
	return s.Content
}

func (s *CreateAgentConfigRequest) SetAgentId(v string) *CreateAgentConfigRequest {
	s.AgentId = &v
	return s
}

func (s *CreateAgentConfigRequest) SetConfigName(v string) *CreateAgentConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateAgentConfigRequest) SetContent(v string) *CreateAgentConfigRequest {
	s.Content = &v
	return s
}

func (s *CreateAgentConfigRequest) Validate() error {
	return dara.Validate(s)
}
