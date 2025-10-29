// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeVersion interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeArn(v string) *AgentRuntimeVersion
	GetAgentRuntimeArn() *string
	SetAgentRuntimeId(v string) *AgentRuntimeVersion
	GetAgentRuntimeId() *string
	SetAgentRuntimeName(v string) *AgentRuntimeVersion
	GetAgentRuntimeName() *string
	SetAgentRuntimeVersion(v string) *AgentRuntimeVersion
	GetAgentRuntimeVersion() *string
	SetDescription(v string) *AgentRuntimeVersion
	GetDescription() *string
	SetLastUpdatedAt(v string) *AgentRuntimeVersion
	GetLastUpdatedAt() *string
}

type AgentRuntimeVersion struct {
	// 智能体运行时的ARN
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:1760720386195983:runtimes/7a1b6d39-9f8f-4ce2-b9c9-6db1b0b9e169
	AgentRuntimeArn *string `json:"agentRuntimeArn,omitempty" xml:"agentRuntimeArn,omitempty"`
	// 智能体运行时的ID
	//
	// example:
	//
	// ar-1234567890abcdef
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	// 智能体运行时的名称
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 已发布版本的版本号
	//
	// example:
	//
	// LATEST
	AgentRuntimeVersion *string `json:"agentRuntimeVersion,omitempty" xml:"agentRuntimeVersion,omitempty"`
	// 此版本的描述
	//
	// example:
	//
	// Initial release with basic functionality
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 最后更新的时间戳
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
}

func (s AgentRuntimeVersion) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeVersion) GoString() string {
	return s.String()
}

func (s *AgentRuntimeVersion) GetAgentRuntimeArn() *string {
	return s.AgentRuntimeArn
}

func (s *AgentRuntimeVersion) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *AgentRuntimeVersion) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *AgentRuntimeVersion) GetAgentRuntimeVersion() *string {
	return s.AgentRuntimeVersion
}

func (s *AgentRuntimeVersion) GetDescription() *string {
	return s.Description
}

func (s *AgentRuntimeVersion) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *AgentRuntimeVersion) SetAgentRuntimeArn(v string) *AgentRuntimeVersion {
	s.AgentRuntimeArn = &v
	return s
}

func (s *AgentRuntimeVersion) SetAgentRuntimeId(v string) *AgentRuntimeVersion {
	s.AgentRuntimeId = &v
	return s
}

func (s *AgentRuntimeVersion) SetAgentRuntimeName(v string) *AgentRuntimeVersion {
	s.AgentRuntimeName = &v
	return s
}

func (s *AgentRuntimeVersion) SetAgentRuntimeVersion(v string) *AgentRuntimeVersion {
	s.AgentRuntimeVersion = &v
	return s
}

func (s *AgentRuntimeVersion) SetDescription(v string) *AgentRuntimeVersion {
	s.Description = &v
	return s
}

func (s *AgentRuntimeVersion) SetLastUpdatedAt(v string) *AgentRuntimeVersion {
	s.LastUpdatedAt = &v
	return s
}

func (s *AgentRuntimeVersion) Validate() error {
	return dara.Validate(s)
}
