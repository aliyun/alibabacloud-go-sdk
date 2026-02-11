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
	SetCredentialName(v string) *UpdateAgentRuntimeInput
	GetCredentialName() *string
	SetDescription(v string) *UpdateAgentRuntimeInput
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *UpdateAgentRuntimeInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateAgentRuntimeInput
	GetExecutionRoleArn() *string
	SetExternalAgentEndpointUrl(v string) *UpdateAgentRuntimeInput
	GetExternalAgentEndpointUrl() *string
	SetHealthCheckConfiguration(v *HealthCheckConfiguration) *UpdateAgentRuntimeInput
	GetHealthCheckConfiguration() *HealthCheckConfiguration
	SetLogConfiguration(v *LogConfiguration) *UpdateAgentRuntimeInput
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int32) *UpdateAgentRuntimeInput
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *UpdateAgentRuntimeInput
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateAgentRuntimeInput
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssMountConfig(v *OSSMountConfig) *UpdateAgentRuntimeInput
	GetOssMountConfig() *OSSMountConfig
	SetPort(v int32) *UpdateAgentRuntimeInput
	GetPort() *int32
	SetProtocolConfiguration(v *ProtocolConfiguration) *UpdateAgentRuntimeInput
	GetProtocolConfiguration() *ProtocolConfiguration
	SetSessionConcurrencyLimitPerInstance(v int32) *UpdateAgentRuntimeInput
	GetSessionConcurrencyLimitPerInstance() *int32
	SetSessionIdleTimeoutSeconds(v int32) *UpdateAgentRuntimeInput
	GetSessionIdleTimeoutSeconds() *int32
	SetWorkspaceId(v string) *UpdateAgentRuntimeInput
	GetWorkspaceId() *string
}

type UpdateAgentRuntimeInput struct {
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 用于访问智能体的凭证名称，访问智能体运行时将使用此凭证进行身份验证
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// example:
	//
	// 更新后的智能体运行时描述
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
	// 外部注册类型的智能体访问端点地址，用于连接已部署在外部的智能体服务
	//
	// example:
	//
	// https://external-agent.example.com/api
	ExternalAgentEndpointUrl *string `json:"externalAgentEndpointUrl,omitempty" xml:"externalAgentEndpointUrl,omitempty"`
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
	// example:
	//
	// 1024
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 文件存储NAS的配置信息，用于挂载NAS文件系统到智能体运行时
	//
	// example:
	//
	// {}
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 智能体运行时的网络配置，包括VPC、安全组等网络访问设置
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 对象存储OSS的挂载配置信息，用于挂载OSS存储桶到智能体运行时
	//
	// example:
	//
	// {}
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
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
	SessionIdleTimeoutSeconds *int32  `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	WorkspaceId               *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *UpdateAgentRuntimeInput) GetCredentialName() *string {
	return s.CredentialName
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

func (s *UpdateAgentRuntimeInput) GetExternalAgentEndpointUrl() *string {
	return s.ExternalAgentEndpointUrl
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

func (s *UpdateAgentRuntimeInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateAgentRuntimeInput) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateAgentRuntimeInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
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

func (s *UpdateAgentRuntimeInput) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *UpdateAgentRuntimeInput) SetCredentialName(v string) *UpdateAgentRuntimeInput {
	s.CredentialName = &v
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

func (s *UpdateAgentRuntimeInput) SetExternalAgentEndpointUrl(v string) *UpdateAgentRuntimeInput {
	s.ExternalAgentEndpointUrl = &v
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

func (s *UpdateAgentRuntimeInput) SetNasConfig(v *NASConfig) *UpdateAgentRuntimeInput {
	s.NasConfig = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateAgentRuntimeInput {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetOssMountConfig(v *OSSMountConfig) *UpdateAgentRuntimeInput {
	s.OssMountConfig = v
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

func (s *UpdateAgentRuntimeInput) SetWorkspaceId(v string) *UpdateAgentRuntimeInput {
	s.WorkspaceId = &v
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
	if s.NasConfig != nil {
		if err := s.NasConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.OssMountConfig != nil {
		if err := s.OssMountConfig.Validate(); err != nil {
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
