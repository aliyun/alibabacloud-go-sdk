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
	SetArmsConfiguration(v *ArmsConfiguration) *UpdateAgentRuntimeInput
	GetArmsConfiguration() *ArmsConfiguration
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
	SetDisableOndemand(v bool) *UpdateAgentRuntimeInput
	GetDisableOndemand() *bool
	SetDisableSessionAffinity(v bool) *UpdateAgentRuntimeInput
	GetDisableSessionAffinity() *bool
	SetDiskSize(v int32) *UpdateAgentRuntimeInput
	GetDiskSize() *int32
	SetEdition(v string) *UpdateAgentRuntimeInput
	GetEdition() *string
	SetEnableSessionIsolation(v bool) *UpdateAgentRuntimeInput
	GetEnableSessionIsolation() *bool
	SetEnvironmentVariables(v map[string]*string) *UpdateAgentRuntimeInput
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *UpdateAgentRuntimeInput
	GetExecutionRoleArn() *string
	SetExternalAgentEndpointUrl(v string) *UpdateAgentRuntimeInput
	GetExternalAgentEndpointUrl() *string
	SetForceEvictInstances(v bool) *UpdateAgentRuntimeInput
	GetForceEvictInstances() *bool
	SetHeaderFieldName(v string) *UpdateAgentRuntimeInput
	GetHeaderFieldName() *string
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
	SetSessionAffinityType(v string) *UpdateAgentRuntimeInput
	GetSessionAffinityType() *string
	SetSessionConcurrencyLimitPerInstance(v int32) *UpdateAgentRuntimeInput
	GetSessionConcurrencyLimitPerInstance() *int32
	SetSessionIdleTimeoutSeconds(v int32) *UpdateAgentRuntimeInput
	GetSessionIdleTimeoutSeconds() *int32
	SetSystemTags(v []*string) *UpdateAgentRuntimeInput
	GetSystemTags() []*string
	SetWorkspaceId(v string) *UpdateAgentRuntimeInput
	GetWorkspaceId() *string
}

type UpdateAgentRuntimeInput struct {
	// The name of the agent runtime.
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 应用实时监控服务（ARMS）的配置信息
	//
	// example:
	//
	// {}
	ArmsConfiguration *ArmsConfiguration `json:"armsConfiguration,omitempty" xml:"armsConfiguration,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code configuration.
	//
	// example:
	//
	// {}
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// The container configuration.
	//
	// example:
	//
	// {}
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// The number of CPU cores.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the credential that the agent runtime uses to authenticate requests.
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// The description of the agent runtime.
	//
	// example:
	//
	// 更新后的智能体运行时描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable on-demand elasticity. Set to true to disable. Default: false.
	//
	// example:
	//
	// false
	DisableOndemand *bool `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// Specifies whether to disable session affinity. Set to true to disable. Default: false.
	//
	// example:
	//
	// false
	DisableSessionAffinity *bool `json:"disableSessionAffinity,omitempty" xml:"disableSessionAffinity,omitempty"`
	// The disk size in gigabytes (GB).
	DiskSize *int32  `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	Edition  *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// Specifies whether to enable session isolation. If enabled, each session runs in an isolated environment.
	//
	// example:
	//
	// false
	EnableSessionIsolation *bool `json:"enableSessionIsolation,omitempty" xml:"enableSessionIsolation,omitempty"`
	// Environment variables for the agent runtime.
	//
	// example:
	//
	// ENV_VAR1=value1,ENV_VAR2=value2
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The execution role ARN that grants the agent runtime permissions to access cloud services.
	//
	// example:
	//
	// acs:ram::1760720386195983:role/AgentRunExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The endpoint URL for an externally registered agent. The platform uses this URL to connect to an agent service deployed outside the platform.
	//
	// example:
	//
	// https://external-agent.example.com/api
	ExternalAgentEndpointUrl *string `json:"externalAgentEndpointUrl,omitempty" xml:"externalAgentEndpointUrl,omitempty"`
	// Specifies whether to perform a best-effort eviction of active Function Compute (FC) sessions when the configuration is updated. This helps the new settings take effect faster.
	//
	// example:
	//
	// true
	ForceEvictInstances *bool `json:"forceEvictInstances,omitempty" xml:"forceEvictInstances,omitempty"`
	// The name of the request header used for session affinity when sessionAffinityType is set to "HEADER_FIELD".
	//
	// example:
	//
	// x-agentrun-session-id
	HeaderFieldName *string `json:"headerFieldName,omitempty" xml:"headerFieldName,omitempty"`
	// The health check configuration for monitoring the health of agent runtime instances.
	//
	// example:
	//
	// {}
	HealthCheckConfiguration *HealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" xml:"healthCheckConfiguration,omitempty"`
	// The configuration for Simple Log Service (SLS).
	//
	// example:
	//
	// {}
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// The amount of memory in megabytes (MB).
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"memory,omitempty" xml:"memory,omitempty"`
	// Configuration for mounting a NAS file system to the agent runtime.
	//
	// example:
	//
	// {}
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The network configuration.
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// Configuration for mounting an OSS bucket to the agent runtime.
	//
	// example:
	//
	// {}
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The port on which the agent service listens.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol configuration.
	//
	// example:
	//
	// {}
	ProtocolConfiguration *ProtocolConfiguration `json:"protocolConfiguration,omitempty" xml:"protocolConfiguration,omitempty"`
	// The session affinity mode. Valid values: NONE (disables session affinity), HEADER_FIELD (routes requests based on a request header), and GENERATED_COOKIE (routes requests using a cookie generated by Function Compute (FC)). The value COOKIE is an alias for GENERATED_COOKIE.
	//
	// example:
	//
	// GENERATED_COOKIE
	SessionAffinityType *string `json:"sessionAffinityType,omitempty" xml:"sessionAffinityType,omitempty"`
	// The maximum number of concurrent sessions allowed per runtime instance.
	//
	// example:
	//
	// 100
	SessionConcurrencyLimitPerInstance *int32 `json:"sessionConcurrencyLimitPerInstance,omitempty" xml:"sessionConcurrencyLimitPerInstance,omitempty"`
	// The idle timeout for a session, in seconds. If an instance remains idle longer than this timeout after receiving no requests, the session expires.
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The system tags for the agent runtime, used for resource classification and management.
	//
	// example:
	//
	// system-tag-1,system-tag-2
	SystemTags []*string `json:"systemTags" xml:"systemTags" type:"Repeated"`
	// The ID of the workspace.
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *UpdateAgentRuntimeInput) GetArmsConfiguration() *ArmsConfiguration {
	return s.ArmsConfiguration
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

func (s *UpdateAgentRuntimeInput) GetDisableOndemand() *bool {
	return s.DisableOndemand
}

func (s *UpdateAgentRuntimeInput) GetDisableSessionAffinity() *bool {
	return s.DisableSessionAffinity
}

func (s *UpdateAgentRuntimeInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpdateAgentRuntimeInput) GetEdition() *string {
	return s.Edition
}

func (s *UpdateAgentRuntimeInput) GetEnableSessionIsolation() *bool {
	return s.EnableSessionIsolation
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

func (s *UpdateAgentRuntimeInput) GetForceEvictInstances() *bool {
	return s.ForceEvictInstances
}

func (s *UpdateAgentRuntimeInput) GetHeaderFieldName() *string {
	return s.HeaderFieldName
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

func (s *UpdateAgentRuntimeInput) GetSessionAffinityType() *string {
	return s.SessionAffinityType
}

func (s *UpdateAgentRuntimeInput) GetSessionConcurrencyLimitPerInstance() *int32 {
	return s.SessionConcurrencyLimitPerInstance
}

func (s *UpdateAgentRuntimeInput) GetSessionIdleTimeoutSeconds() *int32 {
	return s.SessionIdleTimeoutSeconds
}

func (s *UpdateAgentRuntimeInput) GetSystemTags() []*string {
	return s.SystemTags
}

func (s *UpdateAgentRuntimeInput) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateAgentRuntimeInput) SetAgentRuntimeName(v string) *UpdateAgentRuntimeInput {
	s.AgentRuntimeName = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetArmsConfiguration(v *ArmsConfiguration) *UpdateAgentRuntimeInput {
	s.ArmsConfiguration = v
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

func (s *UpdateAgentRuntimeInput) SetDisableOndemand(v bool) *UpdateAgentRuntimeInput {
	s.DisableOndemand = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetDisableSessionAffinity(v bool) *UpdateAgentRuntimeInput {
	s.DisableSessionAffinity = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetDiskSize(v int32) *UpdateAgentRuntimeInput {
	s.DiskSize = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetEdition(v string) *UpdateAgentRuntimeInput {
	s.Edition = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetEnableSessionIsolation(v bool) *UpdateAgentRuntimeInput {
	s.EnableSessionIsolation = &v
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

func (s *UpdateAgentRuntimeInput) SetForceEvictInstances(v bool) *UpdateAgentRuntimeInput {
	s.ForceEvictInstances = &v
	return s
}

func (s *UpdateAgentRuntimeInput) SetHeaderFieldName(v string) *UpdateAgentRuntimeInput {
	s.HeaderFieldName = &v
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

func (s *UpdateAgentRuntimeInput) SetSessionAffinityType(v string) *UpdateAgentRuntimeInput {
	s.SessionAffinityType = &v
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

func (s *UpdateAgentRuntimeInput) SetSystemTags(v []*string) *UpdateAgentRuntimeInput {
	s.SystemTags = v
	return s
}

func (s *UpdateAgentRuntimeInput) SetWorkspaceId(v string) *UpdateAgentRuntimeInput {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateAgentRuntimeInput) Validate() error {
	if s.ArmsConfiguration != nil {
		if err := s.ArmsConfiguration.Validate(); err != nil {
			return err
		}
	}
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
