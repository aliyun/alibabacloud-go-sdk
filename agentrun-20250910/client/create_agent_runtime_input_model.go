// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *CreateAgentRuntimeInput
	GetAgentRuntimeName() *string
	SetArtifactType(v string) *CreateAgentRuntimeInput
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *CreateAgentRuntimeInput
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *CreateAgentRuntimeInput
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *CreateAgentRuntimeInput
	GetCpu() *float32
	SetCredentialId(v string) *CreateAgentRuntimeInput
	GetCredentialId() *string
	SetCredentialName(v string) *CreateAgentRuntimeInput
	GetCredentialName() *string
	SetDescription(v string) *CreateAgentRuntimeInput
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *CreateAgentRuntimeInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *CreateAgentRuntimeInput
	GetExecutionRoleArn() *string
	SetHealthCheckConfiguration(v *HealthCheckConfiguration) *CreateAgentRuntimeInput
	GetHealthCheckConfiguration() *HealthCheckConfiguration
	SetLogConfiguration(v *LogConfiguration) *CreateAgentRuntimeInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *CreateAgentRuntimeInput
	GetMemory() *int32
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateAgentRuntimeInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetPort(v int32) *CreateAgentRuntimeInput
	GetPort() *int32
	SetProtocolConfiguration(v *ProtocolConfiguration) *CreateAgentRuntimeInput
	GetProtocolConfiguration() *ProtocolConfiguration
	SetSessionConcurrencyLimitPerInstance(v int32) *CreateAgentRuntimeInput
	GetSessionConcurrencyLimitPerInstance() *int32
	SetSessionIdleTimeoutSeconds(v int32) *CreateAgentRuntimeInput
	GetSessionIdleTimeoutSeconds() *int32
}

type CreateAgentRuntimeInput struct {
	// 智能体运行时的唯一标识名称，用于区分不同的智能体运行时实例
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 指定智能体运行时的部署类型，支持Code（代码模式）和Container（容器模式）
	//
	// This parameter is required.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// 当artifactType为Code时的代码配置信息，包括代码源、入口文件等
	//
	// example:
	//
	// {}
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// 当artifactType为Container时的容器配置信息，包括镜像地址、启动命令等
	//
	// example:
	//
	// {}
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// 为智能体运行时分配的CPU资源，单位为核数
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// Deprecated
	//
	// 用于访问外部服务的凭证ID，智能体运行时将使用此凭证进行身份验证
	//
	// example:
	//
	// cred-1234567890abcdef
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// 用于访问智能体的凭证名称，访问智能体运行时将使用此凭证进行身份验证
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// 智能体运行时的描述信息，用于说明该运行时的用途和功能
	//
	// example:
	//
	// AI agent runtime for customer service automation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 智能体运行时的环境变量配置，用于在运行时传递配置参数
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
	// SLS（简单日志服务）配置
	//
	// example:
	//
	// {}
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 为智能体运行时分配的内存资源，单位为MB
	//
	// This parameter is required.
	//
	// example:
	//
	// 2048
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 智能体运行时的网络配置，包括VPC、安全组等网络访问设置
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 智能体运行时监听的端口号，用于接收外部请求
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// 智能体运行时的通信协议配置，定义运行时如何与外部系统交互
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
	SessionConcurrencyLimitPerInstance *int32 `json:"sessionConcurrencyLimitPerInstance,omitempty" xml:"sessionConcurrencyLimitPerInstance,omitempty"`
	// 会话的空闲超时时间，单位为秒。实例没有会话请求后处于空闲状态，空闲态为闲置计费模式，超过此超时时间后会话自动过期，不可继续使用
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
}

func (s CreateAgentRuntimeInput) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeInput) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeInput) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *CreateAgentRuntimeInput) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateAgentRuntimeInput) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *CreateAgentRuntimeInput) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *CreateAgentRuntimeInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateAgentRuntimeInput) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateAgentRuntimeInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateAgentRuntimeInput) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRuntimeInput) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateAgentRuntimeInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateAgentRuntimeInput) GetHealthCheckConfiguration() *HealthCheckConfiguration {
	return s.HealthCheckConfiguration
}

func (s *CreateAgentRuntimeInput) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *CreateAgentRuntimeInput) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateAgentRuntimeInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateAgentRuntimeInput) GetPort() *int32 {
	return s.Port
}

func (s *CreateAgentRuntimeInput) GetProtocolConfiguration() *ProtocolConfiguration {
	return s.ProtocolConfiguration
}

func (s *CreateAgentRuntimeInput) GetSessionConcurrencyLimitPerInstance() *int32 {
	return s.SessionConcurrencyLimitPerInstance
}

func (s *CreateAgentRuntimeInput) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *CreateAgentRuntimeInput) SetAgentRuntimeName(v string) *CreateAgentRuntimeInput {
	s.AgentRuntimeName = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetArtifactType(v string) *CreateAgentRuntimeInput {
	s.ArtifactType = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetCodeConfiguration(v *CodeConfiguration) *CreateAgentRuntimeInput {
	s.CodeConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetContainerConfiguration(v *ContainerConfiguration) *CreateAgentRuntimeInput {
	s.ContainerConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetCpu(v float32) *CreateAgentRuntimeInput {
	s.Cpu = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetCredentialId(v string) *CreateAgentRuntimeInput {
	s.CredentialId = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetCredentialName(v string) *CreateAgentRuntimeInput {
	s.CredentialName = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetDescription(v string) *CreateAgentRuntimeInput {
	s.Description = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetEnvironmentVariables(v map[string]*string) *CreateAgentRuntimeInput {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateAgentRuntimeInput) SetExecutionRoleArn(v string) *CreateAgentRuntimeInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetHealthCheckConfiguration(v *HealthCheckConfiguration) *CreateAgentRuntimeInput {
	s.HealthCheckConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetLogConfiguration(v *LogConfiguration) *CreateAgentRuntimeInput {
	s.LogConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetMemory(v int32) *CreateAgentRuntimeInput {
	s.Memory = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetNetworkConfiguration(v *NetworkConfiguration) *CreateAgentRuntimeInput {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetPort(v int32) *CreateAgentRuntimeInput {
	s.Port = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetProtocolConfiguration(v *ProtocolConfiguration) *CreateAgentRuntimeInput {
	s.ProtocolConfiguration = v
	return s
}

func (s *CreateAgentRuntimeInput) SetSessionConcurrencyLimitPerInstance(v int32) *CreateAgentRuntimeInput {
	s.SessionConcurrencyLimitPerInstance = &v
	return s
}

func (s *CreateAgentRuntimeInput) SetSessionIdleTimeoutSeconds(v int32) *CreateAgentRuntimeInput {
	s.SessionIdleTimeoutSeconds = &v
	return s
}

func (s *CreateAgentRuntimeInput) Validate() error {
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
