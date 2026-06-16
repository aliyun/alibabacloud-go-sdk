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
	SetCredentialName(v string) *AgentRuntime
	GetCredentialName() *string
	SetDescription(v string) *AgentRuntime
	GetDescription() *string
	SetDisableOndemand(v bool) *AgentRuntime
	GetDisableOndemand() *bool
	SetDisableSessionAffinity(v bool) *AgentRuntime
	GetDisableSessionAffinity() *bool
	SetDiskSize(v int) *AgentRuntime
	GetDiskSize() *int
	SetEdition(v string) *AgentRuntime
	GetEdition() *string
	SetEnableSessionIsolation(v bool) *AgentRuntime
	GetEnableSessionIsolation() *bool
	SetEnvironmentVariables(v map[string]*string) *AgentRuntime
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *AgentRuntime
	GetExecutionRoleArn() *string
	SetExternalAgentEndpointUrl(v string) *AgentRuntime
	GetExternalAgentEndpointUrl() *string
	SetHeaderFieldName(v string) *AgentRuntime
	GetHeaderFieldName() *string
	SetHealthCheckConfiguration(v *HealthCheckConfiguration) *AgentRuntime
	GetHealthCheckConfiguration() *HealthCheckConfiguration
	SetLastUpdatedAt(v string) *AgentRuntime
	GetLastUpdatedAt() *string
	SetLogConfiguration(v *LogConfiguration) *AgentRuntime
	GetLogConfiguration() *LogConfiguration
	SetMemory(v int) *AgentRuntime
	GetMemory() *int
	SetNasConfig(v *NASConfig) *AgentRuntime
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *AgentRuntime
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssMountConfig(v *OSSMountConfig) *AgentRuntime
	GetOssMountConfig() *OSSMountConfig
	SetPort(v int) *AgentRuntime
	GetPort() *int
	SetProtocolConfiguration(v *ProtocolConfiguration) *AgentRuntime
	GetProtocolConfiguration() *ProtocolConfiguration
	SetResourceGroupId(v string) *AgentRuntime
	GetResourceGroupId() *string
	SetSessionAffinityType(v string) *AgentRuntime
	GetSessionAffinityType() *string
	SetSessionConcurrencyLimitPerInstance(v int) *AgentRuntime
	GetSessionConcurrencyLimitPerInstance() *int
	SetSessionIdleTimeoutSeconds(v int32) *AgentRuntime
	GetSessionIdleTimeoutSeconds() *int32
	SetStatus(v string) *AgentRuntime
	GetStatus() *string
	SetStatusReason(v string) *AgentRuntime
	GetStatusReason() *string
	SetSystemTags(v []*string) *AgentRuntime
	GetSystemTags() []*string
	SetWorkspaceId(v string) *AgentRuntime
	GetWorkspaceId() *string
}

type AgentRuntime struct {
	// The Alibaba Cloud Resource Name (ARN) of the agent runtime.
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:1760720386195983:runtimes/7a1b6d39-9f8f-4ce2-b9c9-6db1b0b9e169
	AgentRuntimeArn *string `json:"agentRuntimeArn,omitempty" xml:"agentRuntimeArn,omitempty"`
	// The unique identifier of the agent runtime.
	//
	// example:
	//
	// ar-1234567890abcdef
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	// The name of the agent runtime.
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// The version number of the agent runtime.
	//
	// example:
	//
	// 1
	AgentRuntimeVersion *string `json:"agentRuntimeVersion,omitempty" xml:"agentRuntimeVersion,omitempty"`
	// The deployment type of the agent runtime. Valid values: `Code` and `Container`.
	//
	// example:
	//
	// Code
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The code configuration details. This parameter is applicable when `artifactType` is set to `Code`.
	//
	// example:
	//
	// {}
	CodeConfiguration *CodeConfiguration `json:"codeConfiguration,omitempty" xml:"codeConfiguration,omitempty"`
	// The container configuration details. This parameter is applicable when `artifactType` is set to `Container`.
	//
	// example:
	//
	// {}
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// The amount of CPU allocated to the agent runtime, in vCPUs.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The creation time of the agent runtime, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The name of the credential used to authenticate requests to the agent runtime.
	//
	// example:
	//
	// my-credential
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// The description of the agent runtime.
	//
	// example:
	//
	// AI agent runtime for customer service automation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable on-demand elasticity. Default: `false`.
	//
	// example:
	//
	// false
	DisableOndemand *bool `json:"disableOndemand,omitempty" xml:"disableOndemand,omitempty"`
	// Specifies whether to disable session affinity. Default: `false`.
	//
	// example:
	//
	// false
	DisableSessionAffinity *bool `json:"disableSessionAffinity,omitempty" xml:"disableSessionAffinity,omitempty"`
	// The disk size.
	DiskSize *int `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	// The edition of the agent runtime.
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// Specifies whether to enable session isolation. If enabled, each session runs in an isolated environment.
	//
	// example:
	//
	// false
	EnableSessionIsolation *bool `json:"enableSessionIsolation,omitempty" xml:"enableSessionIsolation,omitempty"`
	// The environment variables for the agent runtime.
	//
	// example:
	//
	// ENV_VAR1=value1,ENV_VAR2=value2
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// The ARN of the execution role that grants the agent runtime permission to access cloud services.
	//
	// example:
	//
	// acs:ram::1760720386195983:role/AgentRunExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The endpoint URL of an externally deployed agent service.
	//
	// example:
	//
	// https://external-agent.example.com/api
	ExternalAgentEndpointUrl *string `json:"externalAgentEndpointUrl,omitempty" xml:"externalAgentEndpointUrl,omitempty"`
	// The name of the request header used for session affinity when `sessionAffinityType` is `HEADER_FIELD`.
	//
	// example:
	//
	// x-agentrun-session-id
	HeaderFieldName *string `json:"headerFieldName,omitempty" xml:"headerFieldName,omitempty"`
	// The health check configuration.
	//
	// example:
	//
	// {}
	HealthCheckConfiguration *HealthCheckConfiguration `json:"healthCheckConfiguration,omitempty" xml:"healthCheckConfiguration,omitempty"`
	// The last update time of the agent runtime, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The Log Service configuration.
	//
	// example:
	//
	// {}
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	// The amount of memory allocated to the agent runtime, in MB.
	//
	// example:
	//
	// 2048
	Memory *int `json:"memory,omitempty" xml:"memory,omitempty"`
	// The NAS file system configuration.
	//
	// example:
	//
	// {}
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The network configuration of the agent runtime.
	//
	// example:
	//
	// {}
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	// The OSS bucket mount configuration.
	//
	// example:
	//
	// {}
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The port on which the agent runtime listens.
	//
	// example:
	//
	// 8080
	Port *int `json:"port,omitempty" xml:"port,omitempty"`
	// The communication protocol configuration for the agent runtime.
	//
	// example:
	//
	// {}
	ProtocolConfiguration *ProtocolConfiguration `json:"protocolConfiguration,omitempty" xml:"protocolConfiguration,omitempty"`
	// Deprecated
	//
	// The ID of the resource group to which the agent runtime belongs.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The session affinity mode. Valid values: `NONE`, `HEADER_FIELD`, and `GENERATED_COOKIE`. `COOKIE` is a compatibility alias for `GENERATED_COOKIE`.
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
	SessionConcurrencyLimitPerInstance *int `json:"sessionConcurrencyLimitPerInstance,omitempty" xml:"sessionConcurrencyLimitPerInstance,omitempty"`
	// The idle timeout period for a session, in seconds. After this period of inactivity, the session expires and can no longer be used.
	//
	// example:
	//
	// 3600
	SessionIdleTimeoutSeconds *int32 `json:"sessionIdleTimeoutSeconds,omitempty" xml:"sessionIdleTimeoutSeconds,omitempty"`
	// The current status of the agent runtime. Possible values: `READY`, `CREATING`, and `FAILED`.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the current status.
	//
	// example:
	//
	// Runtime is ready for use
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The system tags for the agent runtime.
	//
	// example:
	//
	// system-tag-1,system-tag-2
	SystemTags []*string `json:"systemTags" xml:"systemTags" type:"Repeated"`
	// The ID of the workspace for the agent runtime.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
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

func (s *AgentRuntime) GetCredentialName() *string {
	return s.CredentialName
}

func (s *AgentRuntime) GetDescription() *string {
	return s.Description
}

func (s *AgentRuntime) GetDisableOndemand() *bool {
	return s.DisableOndemand
}

func (s *AgentRuntime) GetDisableSessionAffinity() *bool {
	return s.DisableSessionAffinity
}

func (s *AgentRuntime) GetDiskSize() *int {
	return s.DiskSize
}

func (s *AgentRuntime) GetEdition() *string {
	return s.Edition
}

func (s *AgentRuntime) GetEnableSessionIsolation() *bool {
	return s.EnableSessionIsolation
}

func (s *AgentRuntime) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *AgentRuntime) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *AgentRuntime) GetExternalAgentEndpointUrl() *string {
	return s.ExternalAgentEndpointUrl
}

func (s *AgentRuntime) GetHeaderFieldName() *string {
	return s.HeaderFieldName
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

func (s *AgentRuntime) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *AgentRuntime) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *AgentRuntime) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *AgentRuntime) GetPort() *int {
	return s.Port
}

func (s *AgentRuntime) GetProtocolConfiguration() *ProtocolConfiguration {
	return s.ProtocolConfiguration
}

func (s *AgentRuntime) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AgentRuntime) GetSessionAffinityType() *string {
	return s.SessionAffinityType
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

func (s *AgentRuntime) GetSystemTags() []*string {
	return s.SystemTags
}

func (s *AgentRuntime) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *AgentRuntime) SetCredentialName(v string) *AgentRuntime {
	s.CredentialName = &v
	return s
}

func (s *AgentRuntime) SetDescription(v string) *AgentRuntime {
	s.Description = &v
	return s
}

func (s *AgentRuntime) SetDisableOndemand(v bool) *AgentRuntime {
	s.DisableOndemand = &v
	return s
}

func (s *AgentRuntime) SetDisableSessionAffinity(v bool) *AgentRuntime {
	s.DisableSessionAffinity = &v
	return s
}

func (s *AgentRuntime) SetDiskSize(v int) *AgentRuntime {
	s.DiskSize = &v
	return s
}

func (s *AgentRuntime) SetEdition(v string) *AgentRuntime {
	s.Edition = &v
	return s
}

func (s *AgentRuntime) SetEnableSessionIsolation(v bool) *AgentRuntime {
	s.EnableSessionIsolation = &v
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

func (s *AgentRuntime) SetExternalAgentEndpointUrl(v string) *AgentRuntime {
	s.ExternalAgentEndpointUrl = &v
	return s
}

func (s *AgentRuntime) SetHeaderFieldName(v string) *AgentRuntime {
	s.HeaderFieldName = &v
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

func (s *AgentRuntime) SetNasConfig(v *NASConfig) *AgentRuntime {
	s.NasConfig = v
	return s
}

func (s *AgentRuntime) SetNetworkConfiguration(v *NetworkConfiguration) *AgentRuntime {
	s.NetworkConfiguration = v
	return s
}

func (s *AgentRuntime) SetOssMountConfig(v *OSSMountConfig) *AgentRuntime {
	s.OssMountConfig = v
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

func (s *AgentRuntime) SetResourceGroupId(v string) *AgentRuntime {
	s.ResourceGroupId = &v
	return s
}

func (s *AgentRuntime) SetSessionAffinityType(v string) *AgentRuntime {
	s.SessionAffinityType = &v
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

func (s *AgentRuntime) SetSystemTags(v []*string) *AgentRuntime {
	s.SystemTags = v
	return s
}

func (s *AgentRuntime) SetWorkspaceId(v string) *AgentRuntime {
	s.WorkspaceId = &v
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
