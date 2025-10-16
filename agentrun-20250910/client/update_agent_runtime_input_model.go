// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *UpdateAgentRuntimeInput
	GetAgentRuntimeName() *string
	SetArtifactType(v string) *UpdateAgentRuntimeInput
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *UpdateAgentRuntimeInput
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *UpdateAgentRuntimeInput
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *UpdateAgentRuntimeInput
	GetCpu() *float32
	SetDescription(v string) *UpdateAgentRuntimeInput
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *UpdateAgentRuntimeInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateAgentRuntimeInput
	GetExecutionRoleArn() *string
	SetHealthCheckConfiguration(v *HealthCheckConfiguration) *UpdateAgentRuntimeInput
	GetHealthCheckConfiguration() *HealthCheckConfiguration
	SetLogConfiguration(v *LogConfiguration) *UpdateAgentRuntimeInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *UpdateAgentRuntimeInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateAgentRuntimeInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetPort(v int32) *UpdateAgentRuntimeInput
	GetPort() *int32
	SetProtocolConfiguration(v *ProtocolConfiguration) *UpdateAgentRuntimeInput
	GetProtocolConfiguration() *ProtocolConfiguration
	SetSessionConcurrencyLimitPerInstance(v int32) *UpdateAgentRuntimeInput
	GetSessionConcurrencyLimitPerInstance() *int32
	SetSessionIdleTimeoutSeconds(v int32) *UpdateAgentRuntimeInput
	GetSessionIdleTimeoutSeconds() *int32
}

type UpdateAgentRuntimeInput struct {
	AgentRuntimeName       *string                 `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	ArtifactType           *string                 `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	CodeConfiguration      *CodeConfiguration      `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// This parameter is required.
	Cpu                  *float32           `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string            `json:"description,omitempty" xml:"description,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	// 为智能体运行时提供访问云服务权限的执行角色ARN
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 智能体运行时的健康检查配置，用于监控运行时实例的健康状态
	HealthCheckConfiguration *HealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" xml:"healthCheckConfiguration,omitempty"`
	// SLS（简单日志服务）配置
	LogConfiguration      *LogConfiguration      `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	Memory                *int32                 `json:"memory,omitempty" xml:"memory,omitempty"`
	NetworkConfiguration  *NetworkConfiguration  `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	Port                  *int32                 `json:"port,omitempty" xml:"port,omitempty"`
	ProtocolConfiguration *ProtocolConfiguration `json:"protocolConfiguration,omitempty" xml:"protocolConfiguration,omitempty"`
	// 每个运行时实例允许的最大并发会话数
	SessionConcurrencyLimitPerInstance *int32 `json:"sessionConcurrencyLimitPerInstance,omitempty" xml:"sessionConcurrencyLimitPerInstance,omitempty"`
	// 会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
}

func (s UpdateAgentRuntimeInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeInput) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeInput) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *UpdateAgentRuntimeInput) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateAgentRuntimeInput) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *UpdateAgentRuntimeInput) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateAgentRuntimeInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateAgentRuntimeInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRuntimeInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateAgentRuntimeInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateAgentRuntimeInput) GetHealthCheckConfiguration() *HealthCheckConfiguration {
	return s.HealthCheckConfiguration
}

func (s *UpdateAgentRuntimeInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateAgentRuntimeInput) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateAgentRuntimeInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateAgentRuntimeInput) GetPort() *int32 {
	return s.Port
}

func (s *UpdateAgentRuntimeInput) GetProtocolConfiguration() *ProtocolConfiguration {
	return s.ProtocolConfiguration
}

func (s *UpdateAgentRuntimeInput) GetSessionConcurrencyLimitPerInstance() *int32 {
	return s.SessionConcurrencyLimitPerInstance
}

func (s *UpdateAgentRuntimeInput) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *UpdateAgentRuntimeInput) SetAgentRuntimeName(v string) *UpdateAgentRuntimeInput {
	s.AgentRuntimeName = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetArtifactType(v string) *UpdateAgentRuntimeInput {
	s.ArtifactType = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetCodeConfiguration(v *CodeConfiguration) *UpdateAgentRuntimeInput {
	s.CodeConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetContainerConfiguration(v *ContainerConfiguration) *UpdateAgentRuntimeInput {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetCpu(v float32) *UpdateAgentRuntimeInput {
	s.Cpu = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetDescription(v string) *UpdateAgentRuntimeInput {
	s.Description = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetEnvironmentVariables(v map[string]*string) *UpdateAgentRuntimeInput {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetExecutionRoleArn(v string) *UpdateAgentRuntimeInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetHealthCheckConfiguration(v *HealthCheckConfiguration) *UpdateAgentRuntimeInput {
	s.HealthCheckConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetLogConfiguration(v *LogConfiguration) *UpdateAgentRuntimeInput {
	s.LogConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetMemory(v int32) *UpdateAgentRuntimeInput {
	s.Memory = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateAgentRuntimeInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetPort(v int32) *UpdateAgentRuntimeInput {
	s.Port = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetProtocolConfiguration(v *ProtocolConfiguration) *UpdateAgentRuntimeInput {
	s.ProtocolConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetSessionConcurrencyLimitPerInstance(v int32) *UpdateAgentRuntimeInput {
	s.SessionConcurrencyLimitPerInstance = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetSessionIdleTimeoutSeconds(v int32) *UpdateAgentRuntimeInput {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *UpdateAgentRuntimeInput) Validate() error {
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
