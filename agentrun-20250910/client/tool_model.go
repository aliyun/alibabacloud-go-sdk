// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTool interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *Tool
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *Tool
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *Tool
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *Tool
	GetCpu() *float32
	SetCreateMethod(v string) *Tool
	GetCreateMethod() *string
	SetCreatedAt(v string) *Tool
	GetCreatedAt() *string
	SetCredentialName(v string) *Tool
	GetCredentialName() *string
	SetDescription(v string) *Tool
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *Tool
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *Tool
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *Tool
	GetLogConfiguration() *LogConfiguration
	SetMcpConfig(v *McpConfig) *Tool
	GetMcpConfig() *McpConfig
	SetMemory(v int32) *Tool
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *Tool
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *Tool
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssMountConfig(v *OSSMountConfig) *Tool
	GetOssMountConfig() *OSSMountConfig
	SetPort(v int32) *Tool
	GetPort() *int32
	SetProtocolSpec(v string) *Tool
	GetProtocolSpec() *string
	SetStatus(v string) *Tool
	GetStatus() *string
	SetStatusReason(v string) *Tool
	GetStatusReason() *string
	SetTimeout(v int32) *Tool
	GetTimeout() *int32
	SetToolId(v string) *Tool
	GetToolId() *string
	SetToolName(v string) *Tool
	GetToolName() *string
	SetToolType(v string) *Tool
	GetToolType() *string
	SetUpdatedAt(v string) *Tool
	GetUpdatedAt() *string
	SetWorkspaceId(v string) *Tool
	GetWorkspaceId() *string
}

type Tool struct {
	// 工具部署的制品类型，支持：Code（代码包）、Container（容器镜像）
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// 代码包类型工具的配置信息，包括代码位置、入口函数等
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// 容器类型工具的配置信息，包括镜像地址、启动命令等
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// 工具实例的 CPU 核心数，单位为核
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 工具的创建方式，支持：MCP_REMOTE（远程 MCP 服务器）、MCP_LOCAL_STDIO（本地 MCP 标准输入输出）、MCP_BUNDLE（MCP 打包部署）、CODE_PACKAGE（代码包部署）、OPENAPI_IMPORT（OpenAPI 导入）
	//
	// example:
	//
	// MCP_REMOTE
	CreateMethod *string `json:"createMethod,omitempty" xml:"createMethod,omitempty"`
	// 工具的创建时间，ISO 8601 格式
	//
	// example:
	//
	// 2025-09-10T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 工具使用的凭证名称，用于访问需要认证的外部服务
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// 工具的详细描述信息，说明工具的功能和用途
	//
	// example:
	//
	// 这是一个用于处理文档的 MCP 工具
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工具运行时的环境变量配置，键值对形式
	//
	// example:
	//
	// {"ENV": "production", "LOG_LEVEL": "info"}
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// 工具执行时使用的 RAM 角色 ARN，用于权限控制
	//
	// example:
	//
	// acs:ram::123456789:role/AliyunFCDefaultRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 工具的日志配置，包括日志存储位置和日志级别等
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// MCP 工具的配置信息，包括会话亲和性、代理配置等
	McpConfig *McpConfig `json:"mcpConfig,omitempty" xml:"mcpConfig,omitempty"`
	// 工具实例的内存大小，单位为 MB
	//
	// example:
	//
	// 512
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 文件存储 NAS 的配置信息，用于工具访问 NAS 文件系统
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 工具的网络配置，包括 VPC、安全组等信息
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 对象存储 OSS 的挂载配置，用于工具访问 OSS 存储
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// 工具服务监听的端口号
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// 工具使用的协议规范定义，JSON 格式的字符串
	ProtocolSpec *string `json:"protocolSpec,omitempty" xml:"protocolSpec,omitempty"`
	// 工具的当前运行状态，如：Running（运行中）、Stopped（已停止）、Failed（失败）等
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 工具状态的详细原因说明，特别是在失败或异常状态时提供错误信息
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// 工具执行的超时时间，单位为秒
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 工具的唯一标识符，由系统自动生成
	//
	// example:
	//
	// tool-abc123
	ToolId *string `json:"toolId,omitempty" xml:"toolId,omitempty"`
	// 工具的名称，用于标识和引用工具
	//
	// example:
	//
	// my-mcp-tool
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
	// 工具的类型，支持：MCP（Model Context Protocol 工具）、FUNCTIONCALL（函数调用工具）、SKILL（技能工具）
	//
	// example:
	//
	// MCP
	ToolType *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
	// 工具的最后更新时间，ISO 8601 格式
	//
	// example:
	//
	// 2025-09-10T12:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// 工具所属的工作空间标识符
	//
	// example:
	//
	// workspace-xyz789
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s Tool) String() string {
	return dara.Prettify(s)
}

func (s Tool) GoString() string {
	return s.String()
}

func (s *Tool) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *Tool) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *Tool) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *Tool) GetCpu() *float32 {
	return s.Cpu
}

func (s *Tool) GetCreateMethod() *string {
	return s.CreateMethod
}

func (s *Tool) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Tool) GetCredentialName() *string {
	return s.CredentialName
}

func (s *Tool) GetDescription() *string {
	return s.Description
}

func (s *Tool) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *Tool) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *Tool) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *Tool) GetMcpConfig() *McpConfig {
	return s.McpConfig
}

func (s *Tool) GetMemory() *int32 {
	return s.Memory
}

func (s *Tool) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Tool) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *Tool) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *Tool) GetPort() *int32 {
	return s.Port
}

func (s *Tool) GetProtocolSpec() *string {
	return s.ProtocolSpec
}

func (s *Tool) GetStatus() *string {
	return s.Status
}

func (s *Tool) GetStatusReason() *string {
	return s.StatusReason
}

func (s *Tool) GetTimeout() *int32 {
	return s.Timeout
}

func (s *Tool) GetToolId() *string {
	return s.ToolId
}

func (s *Tool) GetToolName() *string {
	return s.ToolName
}

func (s *Tool) GetToolType() *string {
	return s.ToolType
}

func (s *Tool) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Tool) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Tool) SetArtifactType(v string) *Tool {
	s.ArtifactType = &v
	return s
}

func (s *Tool) SetCodeConfiguration(v *CodeConfiguration) *Tool {
	s.CodeConfiguration = v
	return s
}

func (s *Tool) SetContainerConfiguration(v *ContainerConfiguration) *Tool {
	s.ContainerConfiguration = v
	return s
}

func (s *Tool) SetCpu(v float32) *Tool {
	s.Cpu = &v
	return s
}

func (s *Tool) SetCreateMethod(v string) *Tool {
	s.CreateMethod = &v
	return s
}

func (s *Tool) SetCreatedAt(v string) *Tool {
	s.CreatedAt = &v
	return s
}

func (s *Tool) SetCredentialName(v string) *Tool {
	s.CredentialName = &v
	return s
}

func (s *Tool) SetDescription(v string) *Tool {
	s.Description = &v
	return s
}

func (s *Tool) SetEnvironmentVariables(v map[string]*string) *Tool {
	s.EnvironmentVariables = v
	return s
}

func (s *Tool) SetExecutionRoleArn(v string) *Tool {
	s.ExecutionRoleArn = &v
	return s
}

func (s *Tool) SetLogConfiguration(v *LogConfiguration) *Tool {
	s.LogConfiguration = v
	return s
}

func (s *Tool) SetMcpConfig(v *McpConfig) *Tool {
	s.McpConfig = v
	return s
}

func (s *Tool) SetMemory(v int32) *Tool {
	s.Memory = &v
	return s
}

func (s *Tool) SetNasConfig(v *NASConfig) *Tool {
	s.NasConfig = v
	return s
}

func (s *Tool) SetNetworkConfiguration(v *NetworkConfiguration) *Tool {
	s.NetworkConfiguration = v
	return s
}

func (s *Tool) SetOssMountConfig(v *OSSMountConfig) *Tool {
	s.OssMountConfig = v
	return s
}

func (s *Tool) SetPort(v int32) *Tool {
	s.Port = &v
	return s
}

func (s *Tool) SetProtocolSpec(v string) *Tool {
	s.ProtocolSpec = &v
	return s
}

func (s *Tool) SetStatus(v string) *Tool {
	s.Status = &v
	return s
}

func (s *Tool) SetStatusReason(v string) *Tool {
	s.StatusReason = &v
	return s
}

func (s *Tool) SetTimeout(v int32) *Tool {
	s.Timeout = &v
	return s
}

func (s *Tool) SetToolId(v string) *Tool {
	s.ToolId = &v
	return s
}

func (s *Tool) SetToolName(v string) *Tool {
	s.ToolName = &v
	return s
}

func (s *Tool) SetToolType(v string) *Tool {
	s.ToolType = &v
	return s
}

func (s *Tool) SetUpdatedAt(v string) *Tool {
	s.UpdatedAt = &v
	return s
}

func (s *Tool) SetWorkspaceId(v string) *Tool {
	s.WorkspaceId = &v
	return s
}

func (s *Tool) Validate() error {
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
