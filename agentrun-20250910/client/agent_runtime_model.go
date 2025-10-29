// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeArn(v string) *AgentRuntime
	GetAgentRuntimeArn() *string
	SetAgentRuntimeId(v string) *AgentRuntime
	GetAgentRuntimeId() *string
	SetAgentRuntimeName(v string) *AgentRuntime
	GetAgentRuntimeName() *string
	SetAgentRuntimeVersion(v string) *AgentRuntime
	GetAgentRuntimeVersion() *string
	SetArtifactType(v string) *AgentRuntime
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *AgentRuntime
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *AgentRuntime
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *AgentRuntime
	GetCpu() *float32
	SetCreatedAt(v string) *AgentRuntime
	GetCreatedAt() *string
	SetDescription(v string) *AgentRuntime
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *AgentRuntime
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *AgentRuntime
	GetExecutionRoleArn() *string
	SetHealthCheckConfiguration(v *HealthCheckConfiguration) *AgentRuntime
	GetHealthCheckConfiguration() *HealthCheckConfiguration
	SetLastUpdatedAt(v string) *AgentRuntime
	GetLastUpdatedAt() *string
	SetLogConfiguration(v *LogConfiguration) *AgentRuntime
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int) *AgentRuntime
	GetMemory() *int
	SetNetworkConfiguration(v *NetworkConfiguration) *AgentRuntime
	GetNetworkConfiguration() *NetworkConfiguration
	SetPort(v int) *AgentRuntime
	GetPort() *int
	SetProtocolConfiguration(v *ProtocolConfiguration) *AgentRuntime
	GetProtocolConfiguration() *ProtocolConfiguration
	SetSessionConcurrencyLimitPerInstance(v int) *AgentRuntime
	GetSessionConcurrencyLimitPerInstance() *int
	SetSessionIdleTimeoutSeconds(v int32) *AgentRuntime
	GetSessionIdleTimeoutSeconds() *int32
	SetStatus(v string) *AgentRuntime
	GetStatus() *string
	SetStatusReason(v string) *AgentRuntime
	GetStatusReason() *string
}

type AgentRuntime struct {
	// 智能体运行时的全局唯一资源名称
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:1760720386195983:runtimes/7a1b6d39-9f8f-4ce2-b9c9-6db1b0b9e169
	AgentRuntimeArn *string `json:"agentRuntimeArn,omitempty" xml:"agentRuntimeArn,omitempty"`
	// 智能体运行时的唯一标识符
	//
	// example:
	//
	// ar-1234567890abcdef
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	// 智能体运行时的名称，用于标识和区分不同的运行时实例
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 智能体运行时的版本号，用于版本管理和回滚
	//
	// example:
	//
	// 1
	AgentRuntimeVersion *string `json:"agentRuntimeVersion,omitempty" xml:"agentRuntimeVersion,omitempty"`
	// 智能体运行时的部署类型，支持Code（代码模式）和Container（容器模式）
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// 当artifactType为Code时的代码配置信息
	//
	// example:
	//
	// {}
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// 当artifactType为Container时的容器配置信息
	//
	// example:
	//
	// {}
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// 智能体运行时分配的CPU资源，单位为核数
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 智能体运行时的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 智能体运行时的描述信息，说明该运行时的用途和功能
	//
	// example:
	//
	// AI agent runtime for customer service automation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 智能体运行时的环境变量配置
	//
	// example:
	//
	// ENV_VAR1=value1,ENV_VAR2=value2
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 为智能体运行时提供访问云服务权限的执行角色ARN
	//
	// example:
	//
	// acs:ram::1760720386195983:role/AgentRunExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 智能体运行时的健康检查配置，用于监控运行时实例的健康状态
	//
	// example:
	//
	// {}
	HealthCheckConfiguration *HealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" xml:"healthCheckConfiguration,omitempty"`
	// 智能体运行时最后一次更新的时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// SLS（简单日志服务）配置
	//
	// example:
	//
	// {}
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 智能体运行时分配的内存资源，单位为MB
	//
	// example:
	//
	// 2048
	Memory *int `json:"memory,omitempty" xml:"memory,omitempty"`
	// 智能体运行时的网络配置信息
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 智能体运行时监听的端口号
	//
	// example:
	//
	// 8080
	Port *int `json:"port,omitempty" xml:"port,omitempty"`
	// 智能体运行时的通信协议配置
	//
	// example:
	//
	// {}
	ProtocolConfiguration *ProtocolConfiguration `json:"protocolConfiguration,omitempty" xml:"protocolConfiguration,omitempty"`
	// 每个运行时实例允许的最大并发会话数
	//
	// example:
	//
	// 100
	SessionConcurrencyLimitPerInstance *int `json:"sessionConcurrencyLimitPerInstance,omitempty" xml:"sessionConcurrencyLimitPerInstance,omitempty"`
	// 会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// 智能体运行时的当前状态，如READY（就绪）、CREATING（创建中）、FAILED（失败）等
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 当前状态的原因说明（如适用）
	//
	// example:
	//
	// Runtime is ready for use
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
}

func (s AgentRuntime) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntime) GoString() string {
	return s.String()
}

func (s *AgentRuntime) GetAgentRuntimeArn() *string {
	return s.AgentRuntimeArn
}

func (s *AgentRuntime) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *AgentRuntime) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *AgentRuntime) GetAgentRuntimeVersion() *string {
	return s.AgentRuntimeVersion
}

func (s *AgentRuntime) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *AgentRuntime) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *AgentRuntime) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *AgentRuntime) GetCpu() *float32 {
	return s.Cpu
}

func (s *AgentRuntime) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AgentRuntime) GetDescription() *string {
	return s.Description
}

func (s *AgentRuntime) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *AgentRuntime) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *AgentRuntime) GetHealthCheckConfiguration() *HealthCheckConfiguration {
	return s.HealthCheckConfiguration
}

func (s *AgentRuntime) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *AgentRuntime) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *AgentRuntime) GetMemory() *int {
	return s.Memory
}

func (s *AgentRuntime) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *AgentRuntime) GetPort() *int {
	return s.Port
}

func (s *AgentRuntime) GetProtocolConfiguration() *ProtocolConfiguration {
	return s.ProtocolConfiguration
}

func (s *AgentRuntime) GetSessionConcurrencyLimitPerInstance() *int {
	return s.SessionConcurrencyLimitPerInstance
}

func (s *AgentRuntime) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *AgentRuntime) GetStatus() *string {
	return s.Status
}

func (s *AgentRuntime) GetStatusReason() *string {
	return s.StatusReason
}

func (s *AgentRuntime) SetAgentRuntimeArn(v string) *AgentRuntime {
	s.AgentRuntimeArn = &v
	return s
}

func (s *AgentRuntime) SetAgentRuntimeId(v string) *AgentRuntime {
	s.AgentRuntimeId = &v
	return s
}

func (s *AgentRuntime) SetAgentRuntimeName(v string) *AgentRuntime {
	s.AgentRuntimeName = &v
	return s
}

func (s *AgentRuntime) SetAgentRuntimeVersion(v string) *AgentRuntime {
	s.AgentRuntimeVersion = &v
	return s
}

func (s *AgentRuntime) SetArtifactType(v string) *AgentRuntime {
	s.ArtifactType = &v
	return s
}

func (s *AgentRuntime) SetCodeConfiguration(v *CodeConfiguration) *AgentRuntime {
	s.CodeConfiguration = v
	return s
}

func (s *AgentRuntime) SetContainerConfiguration(v *ContainerConfiguration) *AgentRuntime {
	s.ContainerConfiguration = v
	return s
}

func (s *AgentRuntime) SetCpu(v float32) *AgentRuntime {
	s.Cpu = &v
	return s
}

func (s *AgentRuntime) SetCreatedAt(v string) *AgentRuntime {
	s.CreatedAt = &v
	return s
}

func (s *AgentRuntime) SetDescription(v string) *AgentRuntime {
	s.Description = &v
	return s
}

func (s *AgentRuntime) SetEnvironmentVariables(v map[string]*string) *AgentRuntime {
	s.EnvironmentVariables = v
	return s
}

func (s *AgentRuntime) SetExecutionRoleArn(v string) *AgentRuntime {
	s.ExecutionRoleArn = &v
	return s
}

func (s *AgentRuntime) SetHealthCheckConfiguration(v *HealthCheckConfiguration) *AgentRuntime {
	s.HealthCheckConfiguration = v
	return s
}

func (s *AgentRuntime) SetLastUpdatedAt(v string) *AgentRuntime {
	s.LastUpdatedAt = &v
	return s
}

func (s *AgentRuntime) SetLogConfiguration(v *LogConfiguration) *AgentRuntime {
	s.LogConfiguration = v
	return s
}

func (s *AgentRuntime) SetMemory(v int) *AgentRuntime {
	s.Memory = &v
	return s
}

func (s *AgentRuntime) SetNetworkConfiguration(v *NetworkConfiguration) *AgentRuntime {
	s.NetworkConfiguration = v
	return s
}

func (s *AgentRuntime) SetPort(v int) *AgentRuntime {
	s.Port = &v
	return s
}

func (s *AgentRuntime) SetProtocolConfiguration(v *ProtocolConfiguration) *AgentRuntime {
	s.ProtocolConfiguration = v
	return s
}

func (s *AgentRuntime) SetSessionConcurrencyLimitPerInstance(v int) *AgentRuntime {
	s.SessionConcurrencyLimitPerInstance = &v
	return s
}

func (s *AgentRuntime) SetSessionIdleTimeoutSeconds(v int32) *AgentRuntime {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *AgentRuntime) SetStatus(v string) *AgentRuntime {
	s.Status = &v
	return s
}

func (s *AgentRuntime) SetStatusReason(v string) *AgentRuntime {
	s.StatusReason = &v
	return s
}

func (s *AgentRuntime) Validate() error {
	if s.CodeConfiguration != nil {
		if err := s.CodeConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckConfiguration != nil {
		if err := s.HealthCheckConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ProtocolConfiguration != nil {
		if err := s.ProtocolConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
