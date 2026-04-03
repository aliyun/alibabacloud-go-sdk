// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateToolInputV2 interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *UpdateToolInputV2
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *UpdateToolInputV2
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *UpdateToolInputV2
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *UpdateToolInputV2
	GetCpu() *float32
	SetCreateMethod(v string) *UpdateToolInputV2
	GetCreateMethod() *string
	SetCredentialName(v string) *UpdateToolInputV2
	GetCredentialName() *string
	SetDescription(v string) *UpdateToolInputV2
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *UpdateToolInputV2
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateToolInputV2
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *UpdateToolInputV2
	GetLogConfiguration() *LogConfiguration
	SetMcpConfig(v *McpConfig) *UpdateToolInputV2
	GetMcpConfig() *McpConfig
	SetMemory(v int32) *UpdateToolInputV2
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *UpdateToolInputV2
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *UpdateToolInputV2
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssMountConfig(v *OSSMountConfig) *UpdateToolInputV2
	GetOssMountConfig() *OSSMountConfig
	SetPort(v int32) *UpdateToolInputV2
	GetPort() *int32
	SetProtocolSpec(v string) *UpdateToolInputV2
	GetProtocolSpec() *string
	SetTimeout(v int32) *UpdateToolInputV2
	GetTimeout() *int32
	SetWorkspaceId(v string) *UpdateToolInputV2
	GetWorkspaceId() *string
}

type UpdateToolInputV2 struct {
	// 更新工具部署的制品类型，支持：Code（代码包）、Container（容器镜像）
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// 更新代码包类型工具的配置信息
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// 更新容器类型工具的配置信息
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// 更新工具实例的 CPU 核心数，单位为核
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 更新工具的创建方式。支持：MCP_REMOTE、MCP_LOCAL_STDIO、MCP_BUNDLE、CODE_PACKAGE、OPENAPI_IMPORT
	//
	// example:
	//
	// MCP_REMOTE
	CreateMethod *string `json:"createMethod,omitempty" xml:"createMethod,omitempty"`
	// 更新工具使用的凭证名称
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// 更新工具的描述信息
	//
	// example:
	//
	// 更新后的工具描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 更新工具运行时的环境变量配置
	//
	// example:
	//
	// {"ENV": "production"}
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// 更新工具执行时使用的 RAM 角色 ARN
	//
	// example:
	//
	// acs:ram::123456789:role/AliyunFCDefaultRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 更新工具的日志配置
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// 更新 MCP 工具的配置信息，包括会话亲和性、代理配置等
	McpConfig *McpConfig `json:"mcpConfig,omitempty" xml:"mcpConfig,omitempty"`
	// 更新工具实例的内存大小，单位为 MB
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 更新文件存储 NAS 的配置信息
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 更新工具的网络配置
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 更新对象存储 OSS 的挂载配置
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// 更新工具服务监听的端口号
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// 更新工具使用的协议规范定义
	ProtocolSpec *string `json:"protocolSpec,omitempty" xml:"protocolSpec,omitempty"`
	// 更新工具执行的超时时间，单位为秒
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 更新工具所属的工作空间标识符
	//
	// example:
	//
	// workspace-xyz789
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateToolInputV2) String() string {
	return dara.Prettify(s)
}

func (s UpdateToolInputV2) GoString() string {
	return s.String()
}

func (s *UpdateToolInputV2) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateToolInputV2) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *UpdateToolInputV2) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *UpdateToolInputV2) GetCpu() *float32 {
	return s.Cpu
}

func (s *UpdateToolInputV2) GetCreateMethod() *string {
	return s.CreateMethod
}

func (s *UpdateToolInputV2) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateToolInputV2) GetDescription() *string {
	return s.Description
}

func (s *UpdateToolInputV2) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *UpdateToolInputV2) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateToolInputV2) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *UpdateToolInputV2) GetMcpConfig() *McpConfig {
	return s.McpConfig
}

func (s *UpdateToolInputV2) GetMemory() *int32 {
	return s.Memory
}

func (s *UpdateToolInputV2) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateToolInputV2) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateToolInputV2) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *UpdateToolInputV2) GetPort() *int32 {
	return s.Port
}

func (s *UpdateToolInputV2) GetProtocolSpec() *string {
	return s.ProtocolSpec
}

func (s *UpdateToolInputV2) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateToolInputV2) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateToolInputV2) SetArtifactType(v string) *UpdateToolInputV2 {
	s.ArtifactType = &v
	return s
}

func (s *UpdateToolInputV2) SetCodeConfiguration(v *CodeConfiguration) *UpdateToolInputV2 {
	s.CodeConfiguration = v
	return s
}

func (s *UpdateToolInputV2) SetContainerConfiguration(v *ContainerConfiguration) *UpdateToolInputV2 {
	s.ContainerConfiguration = v
	return s
}

func (s *UpdateToolInputV2) SetCpu(v float32) *UpdateToolInputV2 {
	s.Cpu = &v
	return s
}

func (s *UpdateToolInputV2) SetCreateMethod(v string) *UpdateToolInputV2 {
	s.CreateMethod = &v
	return s
}

func (s *UpdateToolInputV2) SetCredentialName(v string) *UpdateToolInputV2 {
	s.CredentialName = &v
	return s
}

func (s *UpdateToolInputV2) SetDescription(v string) *UpdateToolInputV2 {
	s.Description = &v
	return s
}

func (s *UpdateToolInputV2) SetEnvironmentVariables(v map[string]*string) *UpdateToolInputV2 {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateToolInputV2) SetExecutionRoleArn(v string) *UpdateToolInputV2 {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateToolInputV2) SetLogConfiguration(v *LogConfiguration) *UpdateToolInputV2 {
	s.LogConfiguration = v
	return s
}

func (s *UpdateToolInputV2) SetMcpConfig(v *McpConfig) *UpdateToolInputV2 {
	s.McpConfig = v
	return s
}

func (s *UpdateToolInputV2) SetMemory(v int32) *UpdateToolInputV2 {
	s.Memory = &v
	return s
}

func (s *UpdateToolInputV2) SetNasConfig(v *NASConfig) *UpdateToolInputV2 {
	s.NasConfig = v
	return s
}

func (s *UpdateToolInputV2) SetNetworkConfiguration(v *NetworkConfiguration) *UpdateToolInputV2 {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateToolInputV2) SetOssMountConfig(v *OSSMountConfig) *UpdateToolInputV2 {
	s.OssMountConfig = v
	return s
}

func (s *UpdateToolInputV2) SetPort(v int32) *UpdateToolInputV2 {
	s.Port = &v
	return s
}

func (s *UpdateToolInputV2) SetProtocolSpec(v string) *UpdateToolInputV2 {
	s.ProtocolSpec = &v
	return s
}

func (s *UpdateToolInputV2) SetTimeout(v int32) *UpdateToolInputV2 {
	s.Timeout = &v
	return s
}

func (s *UpdateToolInputV2) SetWorkspaceId(v string) *UpdateToolInputV2 {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateToolInputV2) Validate() error {
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
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.McpConfig != nil {
		if err := s.McpConfig.Validate(); err != nil {
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
	return nil
}
