// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateToolInputV2 interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *CreateToolInputV2
	GetArtifactType() *string
	SetCodeConfiguration(v *CodeConfiguration) *CreateToolInputV2
	GetCodeConfiguration() *CodeConfiguration
	SetContainerConfiguration(v *ContainerConfiguration) *CreateToolInputV2
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *CreateToolInputV2
	GetCpu() *float32
	SetCreateMethod(v string) *CreateToolInputV2
	GetCreateMethod() *string
	SetCredentialName(v string) *CreateToolInputV2
	GetCredentialName() *string
	SetDescription(v string) *CreateToolInputV2
	GetDescription() *string
	SetEnvironmentVariables(v map[string]*string) *CreateToolInputV2
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *CreateToolInputV2
	GetExecutionRoleArn() *string
	SetLogConfiguration(v *LogConfiguration) *CreateToolInputV2
	GetLogConfiguration() *LogConfiguration
	SetMcpConfig(v *McpConfig) *CreateToolInputV2
	GetMcpConfig() *McpConfig
	SetMemory(v int32) *CreateToolInputV2
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *CreateToolInputV2
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *CreateToolInputV2
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssMountConfig(v *OSSMountConfig) *CreateToolInputV2
	GetOssMountConfig() *OSSMountConfig
	SetPort(v int32) *CreateToolInputV2
	GetPort() *int32
	SetProtocolSpec(v string) *CreateToolInputV2
	GetProtocolSpec() *string
	SetTimeout(v int32) *CreateToolInputV2
	GetTimeout() *int32
	SetToolName(v string) *CreateToolInputV2
	GetToolName() *string
	SetToolType(v string) *CreateToolInputV2
	GetToolType() *string
	SetWorkspaceId(v string) *CreateToolInputV2
	GetWorkspaceId() *string
}

type CreateToolInputV2 struct {
	// 工具部署的制品类型，支持：Code（代码包）、Container（容器镜像）
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// 代码包类型工具的配置信息，包括代码位置、入口函数、运行时等，仅在制品类型为 Code 时需要提供
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// 容器类型工具的配置信息，包括镜像地址、启动命令、端口映射等，仅在制品类型为 Container 时需要提供
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// 工具实例的 CPU 核心数，单位为核，支持小数，如 0.5 表示半核
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 工具的创建方式，必填项。支持：MCP_REMOTE（远程 MCP 服务器）、MCP_LOCAL_STDIO（本地 MCP 标准输入输出）、MCP_BUNDLE（MCP 打包部署）、CODE_PACKAGE（代码包部署）、OPENAPI_IMPORT（OpenAPI 导入）
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP_REMOTE
	CreateMethod *string `json:"createMethod,omitempty" xml:"createMethod,omitempty"`
	// 工具使用的凭证名称，用于访问需要认证的外部服务，需要提前在系统中创建凭证
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
	// 工具运行时的环境变量配置，键值对形式，用于传递配置信息到工具运行环境
	//
	// example:
	//
	// {"ENV": "production", "LOG_LEVEL": "info"}
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// 工具执行时使用的 RAM 角色 ARN，用于权限控制，格式如：acs:ram::123456789:role/AliyunFCDefaultRole
	//
	// example:
	//
	// acs:ram::123456789:role/AliyunFCDefaultRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 工具的日志配置，包括日志存储位置（SLS）、日志级别等
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// MCP 工具的配置信息，包括会话亲和性、代理配置、绑定配置等，仅在工具类型为 MCP 时需要提供
	McpConfig *McpConfig `json:"mcpConfig,omitempty" xml:"mcpConfig,omitempty"`
	// 工具实例的内存大小，单位为 MB，建议根据工具的实际需求配置
	//
	// example:
	//
	// 512
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// 文件存储 NAS 的配置信息，用于工具访问 NAS 文件系统，实现持久化存储
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// 工具的网络配置，包括 VPC、安全组、交换机等信息，用于配置工具的网络访问能力
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// 对象存储 OSS 的挂载配置，用于工具访问 OSS 存储桶中的文件
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// 工具服务监听的端口号，用于接收请求
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// 工具使用的协议规范定义，JSON 格式的字符串，定义工具的接口和交互方式
	ProtocolSpec *string `json:"protocolSpec,omitempty" xml:"protocolSpec,omitempty"`
	// 工具执行的超时时间，单位为秒，超过此时间工具调用将被终止
	//
	// example:
	//
	// 300
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 工具的名称，必填项，用于标识和引用工具，需要在工作空间内唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-tool
	ToolName *string `json:"toolName,omitempty" xml:"toolName,omitempty"`
	// 工具的类型，必填项。支持：MCP（Model Context Protocol 工具）、FUNCTIONCALL（函数调用工具）、SKILL（技能工具）
	//
	// This parameter is required.
	//
	// example:
	//
	// MCP
	ToolType *string `json:"toolType,omitempty" xml:"toolType,omitempty"`
	// 工具所属的工作空间标识符，可选项，不填则使用默认工作空间
	//
	// example:
	//
	// workspace-xyz789
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateToolInputV2) String() string {
	return dara.Prettify(s)
}

func (s CreateToolInputV2) GoString() string {
	return s.String()
}

func (s *CreateToolInputV2) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateToolInputV2) GetCodeConfiguration() *CodeConfiguration {
	return s.CodeConfiguration
}

func (s *CreateToolInputV2) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *CreateToolInputV2) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateToolInputV2) GetCreateMethod() *string {
	return s.CreateMethod
}

func (s *CreateToolInputV2) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateToolInputV2) GetDescription() *string {
	return s.Description
}

func (s *CreateToolInputV2) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *CreateToolInputV2) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateToolInputV2) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *CreateToolInputV2) GetMcpConfig() *McpConfig {
	return s.McpConfig
}

func (s *CreateToolInputV2) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateToolInputV2) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateToolInputV2) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateToolInputV2) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateToolInputV2) GetPort() *int32 {
	return s.Port
}

func (s *CreateToolInputV2) GetProtocolSpec() *string {
	return s.ProtocolSpec
}

func (s *CreateToolInputV2) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateToolInputV2) GetToolName() *string {
	return s.ToolName
}

func (s *CreateToolInputV2) GetToolType() *string {
	return s.ToolType
}

func (s *CreateToolInputV2) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateToolInputV2) SetArtifactType(v string) *CreateToolInputV2 {
	s.ArtifactType = &v
	return s
}

func (s *CreateToolInputV2) SetCodeConfiguration(v *CodeConfiguration) *CreateToolInputV2 {
	s.CodeConfiguration = v
	return s
}

func (s *CreateToolInputV2) SetContainerConfiguration(v *ContainerConfiguration) *CreateToolInputV2 {
	s.ContainerConfiguration = v
	return s
}

func (s *CreateToolInputV2) SetCpu(v float32) *CreateToolInputV2 {
	s.Cpu = &v
	return s
}

func (s *CreateToolInputV2) SetCreateMethod(v string) *CreateToolInputV2 {
	s.CreateMethod = &v
	return s
}

func (s *CreateToolInputV2) SetCredentialName(v string) *CreateToolInputV2 {
	s.CredentialName = &v
	return s
}

func (s *CreateToolInputV2) SetDescription(v string) *CreateToolInputV2 {
	s.Description = &v
	return s
}

func (s *CreateToolInputV2) SetEnvironmentVariables(v map[string]*string) *CreateToolInputV2 {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateToolInputV2) SetExecutionRoleArn(v string) *CreateToolInputV2 {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateToolInputV2) SetLogConfiguration(v *LogConfiguration) *CreateToolInputV2 {
	s.LogConfiguration = v
	return s
}

func (s *CreateToolInputV2) SetMcpConfig(v *McpConfig) *CreateToolInputV2 {
	s.McpConfig = v
	return s
}

func (s *CreateToolInputV2) SetMemory(v int32) *CreateToolInputV2 {
	s.Memory = &v
	return s
}

func (s *CreateToolInputV2) SetNasConfig(v *NASConfig) *CreateToolInputV2 {
	s.NasConfig = v
	return s
}

func (s *CreateToolInputV2) SetNetworkConfiguration(v *NetworkConfiguration) *CreateToolInputV2 {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateToolInputV2) SetOssMountConfig(v *OSSMountConfig) *CreateToolInputV2 {
	s.OssMountConfig = v
	return s
}

func (s *CreateToolInputV2) SetPort(v int32) *CreateToolInputV2 {
	s.Port = &v
	return s
}

func (s *CreateToolInputV2) SetProtocolSpec(v string) *CreateToolInputV2 {
	s.ProtocolSpec = &v
	return s
}

func (s *CreateToolInputV2) SetTimeout(v int32) *CreateToolInputV2 {
	s.Timeout = &v
	return s
}

func (s *CreateToolInputV2) SetToolName(v string) *CreateToolInputV2 {
	s.ToolName = &v
	return s
}

func (s *CreateToolInputV2) SetToolType(v string) *CreateToolInputV2 {
	s.ToolType = &v
	return s
}

func (s *CreateToolInputV2) SetWorkspaceId(v string) *CreateToolInputV2 {
	s.WorkspaceId = &v
	return s
}

func (s *CreateToolInputV2) Validate() error {
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
