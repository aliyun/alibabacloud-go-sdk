// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgent interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Agent
	GetCreatedAt() *int64
	SetDescription(v string) *Agent
	GetDescription() *string
	SetMcpServerId(v string) *Agent
	GetMcpServerId() *string
	SetMcpServerName(v string) *Agent
	GetMcpServerName() *string
	SetMetadata(v *Metadata) *Agent
	GetMetadata() *Metadata
	SetName(v string) *Agent
	GetName() *string
	SetPrompt(v string) *Agent
	GetPrompt() *string
	SetUpdatedAt(v int64) *Agent
	GetUpdatedAt() *int64
}

type Agent struct {
	// The creation time in UNIX milliseconds.
	//
	// example:
	//
	// 1717200000000
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description of the agent.
	//
	// example:
	//
	// IoT device data query agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the MCP server associated with the agent.
	//
	// example:
	//
	// 8vHW4mDRxxxxx
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	// The name of the MCP server.
	//
	// example:
	//
	// data-analyze-agent
	McpServerName *string `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
	// The associated table metadata.
	Metadata *Metadata `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom prompt.
	//
	// example:
	//
	// You are an IoT data analytics assistant...
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The update time in UNIX milliseconds.
	//
	// example:
	//
	// 1717200000000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s Agent) String() string {
	return dara.Prettify(s)
}

func (s Agent) GoString() string {
	return s.String()
}

func (s *Agent) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Agent) GetDescription() *string {
	return s.Description
}

func (s *Agent) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *Agent) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *Agent) GetMetadata() *Metadata {
	return s.Metadata
}

func (s *Agent) GetName() *string {
	return s.Name
}

func (s *Agent) GetPrompt() *string {
	return s.Prompt
}

func (s *Agent) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Agent) SetCreatedAt(v int64) *Agent {
	s.CreatedAt = &v
	return s
}

func (s *Agent) SetDescription(v string) *Agent {
	s.Description = &v
	return s
}

func (s *Agent) SetMcpServerId(v string) *Agent {
	s.McpServerId = &v
	return s
}

func (s *Agent) SetMcpServerName(v string) *Agent {
	s.McpServerName = &v
	return s
}

func (s *Agent) SetMetadata(v *Metadata) *Agent {
	s.Metadata = v
	return s
}

func (s *Agent) SetName(v string) *Agent {
	s.Name = &v
	return s
}

func (s *Agent) SetPrompt(v string) *Agent {
	s.Prompt = &v
	return s
}

func (s *Agent) SetUpdatedAt(v int64) *Agent {
	s.UpdatedAt = &v
	return s
}

func (s *Agent) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
