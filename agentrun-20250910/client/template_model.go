// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAllowAnonymousManage(v bool) *Template
	GetAllowAnonymousManage() *bool
	SetContainerConfiguration(v *ContainerConfiguration) *Template
	GetContainerConfiguration() *ContainerConfiguration
	SetCpu(v float32) *Template
	GetCpu() *float32
	SetCreatedAt(v string) *Template
	GetCreatedAt() *string
	SetCredentialConfiguration(v *CredentialConfiguration) *Template
	GetCredentialConfiguration() *CredentialConfiguration
	SetDescription(v string) *Template
	GetDescription() *string
	SetDiskSize(v int32) *Template
	GetDiskSize() *int32
	SetEnableAgent(v bool) *Template
	GetEnableAgent() *bool
	SetEnvironmentVariables(v map[string]*string) *Template
	GetEnvironmentVariables() map[string]*string
	SetExecutionRoleArn(v string) *Template
	GetExecutionRoleArn() *string
	SetLastUpdatedAt(v string) *Template
	GetLastUpdatedAt() *string
	SetLogConfiguration(v *LogConfiguration) *Template
	GetLogConfiguration() *LogConfiguration
	SetMcpOptions(v *TemplateMcpOptions) *Template
	GetMcpOptions() *TemplateMcpOptions
	SetMcpState(v *TemplateMcpState) *Template
	GetMcpState() *TemplateMcpState
	SetMemory(v int32) *Template
	GetMemory() *int32
	SetNasConfig(v *NASConfig) *Template
	GetNasConfig() *NASConfig
	SetNetworkConfiguration(v *NetworkConfiguration) *Template
	GetNetworkConfiguration() *NetworkConfiguration
	SetOssConfiguration(v []*OssConfiguration) *Template
	GetOssConfiguration() []*OssConfiguration
	SetResourceName(v string) *Template
	GetResourceName() *string
	SetSandboxIdleTimeoutInSeconds(v string) *Template
	GetSandboxIdleTimeoutInSeconds() *string
	SetSandboxTTLInSeconds(v string) *Template
	GetSandboxTTLInSeconds() *string
	SetStatus(v string) *Template
	GetStatus() *string
	SetStatusReason(v string) *Template
	GetStatusReason() *string
	SetTemplateArn(v string) *Template
	GetTemplateArn() *string
	SetTemplateConfiguration(v map[string]interface{}) *Template
	GetTemplateConfiguration() map[string]interface{}
	SetTemplateId(v string) *Template
	GetTemplateId() *string
	SetTemplateName(v string) *Template
	GetTemplateName() *string
	SetTemplateType(v string) *Template
	GetTemplateType() *string
	SetTemplateVersion(v string) *Template
	GetTemplateVersion() *string
}

type Template struct {
	AllowAnonymousManage   *bool                   `json:"allowAnonymousManage,omitempty" xml:"allowAnonymousManage,omitempty"`
	ContainerConfiguration *ContainerConfiguration `json:"containerConfiguration,omitempty" xml:"containerConfiguration,omitempty"`
	// This parameter is required.
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CreatedAt               *string                  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialConfiguration *CredentialConfiguration `json:"credentialConfiguration,omitempty" xml:"credentialConfiguration,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnableAgent             *bool                    `json:"enableAgent,omitempty" xml:"enableAgent,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	ExecutionRoleArn        *string                  `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	LastUpdatedAt           *string                  `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	LogConfiguration        *LogConfiguration        `json:"logConfiguration,omitempty" xml:"logConfiguration,omitempty"`
	McpOptions              *TemplateMcpOptions      `json:"mcpOptions,omitempty" xml:"mcpOptions,omitempty" type:"Struct"`
	McpState                *TemplateMcpState        `json:"mcpState,omitempty" xml:"mcpState,omitempty" type:"Struct"`
	// This parameter is required.
	Memory                      *int32                `json:"memory,omitempty" xml:"memory,omitempty"`
	NasConfig                   *NASConfig            `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	NetworkConfiguration        *NetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty"`
	OssConfiguration            []*OssConfiguration   `json:"ossConfiguration,omitempty" xml:"ossConfiguration,omitempty" type:"Repeated"`
	ResourceName                *string               `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	SandboxIdleTimeoutInSeconds *string               `json:"sandboxIdleTimeoutInSeconds,omitempty" xml:"sandboxIdleTimeoutInSeconds,omitempty"`
	// Deprecated
	SandboxTTLInSeconds   *string                `json:"sandboxTTLInSeconds,omitempty" xml:"sandboxTTLInSeconds,omitempty"`
	Status                *string                `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason          *string                `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	TemplateArn           *string                `json:"templateArn,omitempty" xml:"templateArn,omitempty"`
	TemplateConfiguration map[string]interface{} `json:"templateConfiguration,omitempty" xml:"templateConfiguration,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	TemplateName    *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateType    *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	TemplateVersion *string `json:"templateVersion,omitempty" xml:"templateVersion,omitempty"`
}

func (s Template) String() string {
	return dara.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) GetAllowAnonymousManage() *bool {
	return s.AllowAnonymousManage
}

func (s *Template) GetContainerConfiguration() *ContainerConfiguration {
	return s.ContainerConfiguration
}

func (s *Template) GetCpu() *float32 {
	return s.Cpu
}

func (s *Template) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Template) GetCredentialConfiguration() *CredentialConfiguration {
	return s.CredentialConfiguration
}

func (s *Template) GetDescription() *string {
	return s.Description
}

func (s *Template) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *Template) GetEnableAgent() *bool {
	return s.EnableAgent
}

func (s *Template) GetEnvironmentVariables() map[string]*string {
	return s.EnvironmentVariables
}

func (s *Template) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *Template) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *Template) GetLogConfiguration() *LogConfiguration {
	return s.LogConfiguration
}

func (s *Template) GetMcpOptions() *TemplateMcpOptions {
	return s.McpOptions
}

func (s *Template) GetMcpState() *TemplateMcpState {
	return s.McpState
}

func (s *Template) GetMemory() *int32 {
	return s.Memory
}

func (s *Template) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Template) GetNetworkConfiguration() *NetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *Template) GetOssConfiguration() []*OssConfiguration {
	return s.OssConfiguration
}

func (s *Template) GetResourceName() *string {
	return s.ResourceName
}

func (s *Template) GetSandboxIdleTimeoutInSeconds() *string {
	return s.SandboxIdleTimeoutInSeconds
}

func (s *Template) GetSandboxTTLInSeconds() *string {
	return s.SandboxTTLInSeconds
}

func (s *Template) GetStatus() *string {
	return s.Status
}

func (s *Template) GetStatusReason() *string {
	return s.StatusReason
}

func (s *Template) GetTemplateArn() *string {
	return s.TemplateArn
}

func (s *Template) GetTemplateConfiguration() map[string]interface{} {
	return s.TemplateConfiguration
}

func (s *Template) GetTemplateId() *string {
	return s.TemplateId
}

func (s *Template) GetTemplateName() *string {
	return s.TemplateName
}

func (s *Template) GetTemplateType() *string {
	return s.TemplateType
}

func (s *Template) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *Template) SetAllowAnonymousManage(v bool) *Template {
	s.AllowAnonymousManage = &v
	return s
}

func (s *Template) SetContainerConfiguration(v *ContainerConfiguration) *Template {
	s.ContainerConfiguration = v
	return s
}

func (s *Template) SetCpu(v float32) *Template {
	s.Cpu = &v
	return s
}

func (s *Template) SetCreatedAt(v string) *Template {
	s.CreatedAt = &v
	return s
}

func (s *Template) SetCredentialConfiguration(v *CredentialConfiguration) *Template {
	s.CredentialConfiguration = v
	return s
}

func (s *Template) SetDescription(v string) *Template {
	s.Description = &v
	return s
}

func (s *Template) SetDiskSize(v int32) *Template {
	s.DiskSize = &v
	return s
}

func (s *Template) SetEnableAgent(v bool) *Template {
	s.EnableAgent = &v
	return s
}

func (s *Template) SetEnvironmentVariables(v map[string]*string) *Template {
	s.EnvironmentVariables = v
	return s
}

func (s *Template) SetExecutionRoleArn(v string) *Template {
	s.ExecutionRoleArn = &v
	return s
}

func (s *Template) SetLastUpdatedAt(v string) *Template {
	s.LastUpdatedAt = &v
	return s
}

func (s *Template) SetLogConfiguration(v *LogConfiguration) *Template {
	s.LogConfiguration = v
	return s
}

func (s *Template) SetMcpOptions(v *TemplateMcpOptions) *Template {
	s.McpOptions = v
	return s
}

func (s *Template) SetMcpState(v *TemplateMcpState) *Template {
	s.McpState = v
	return s
}

func (s *Template) SetMemory(v int32) *Template {
	s.Memory = &v
	return s
}

func (s *Template) SetNasConfig(v *NASConfig) *Template {
	s.NasConfig = v
	return s
}

func (s *Template) SetNetworkConfiguration(v *NetworkConfiguration) *Template {
	s.NetworkConfiguration = v
	return s
}

func (s *Template) SetOssConfiguration(v []*OssConfiguration) *Template {
	s.OssConfiguration = v
	return s
}

func (s *Template) SetResourceName(v string) *Template {
	s.ResourceName = &v
	return s
}

func (s *Template) SetSandboxIdleTimeoutInSeconds(v string) *Template {
	s.SandboxIdleTimeoutInSeconds = &v
	return s
}

func (s *Template) SetSandboxTTLInSeconds(v string) *Template {
	s.SandboxTTLInSeconds = &v
	return s
}

func (s *Template) SetStatus(v string) *Template {
	s.Status = &v
	return s
}

func (s *Template) SetStatusReason(v string) *Template {
	s.StatusReason = &v
	return s
}

func (s *Template) SetTemplateArn(v string) *Template {
	s.TemplateArn = &v
	return s
}

func (s *Template) SetTemplateConfiguration(v map[string]interface{}) *Template {
	s.TemplateConfiguration = v
	return s
}

func (s *Template) SetTemplateId(v string) *Template {
	s.TemplateId = &v
	return s
}

func (s *Template) SetTemplateName(v string) *Template {
	s.TemplateName = &v
	return s
}

func (s *Template) SetTemplateType(v string) *Template {
	s.TemplateType = &v
	return s
}

func (s *Template) SetTemplateVersion(v string) *Template {
	s.TemplateVersion = &v
	return s
}

func (s *Template) Validate() error {
	if s.ContainerConfiguration != nil {
		if err := s.ContainerConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.CredentialConfiguration != nil {
		if err := s.CredentialConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.McpOptions != nil {
		if err := s.McpOptions.Validate(); err != nil {
			return err
		}
	}
	if s.McpState != nil {
		if err := s.McpState.Validate(); err != nil {
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
	if s.OssConfiguration != nil {
		for _, item := range s.OssConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TemplateMcpOptions struct {
	EnabledTools []*string `json:"enabledTools,omitempty" xml:"enabledTools,omitempty" type:"Repeated"`
	// example:
	//
	// streamable-http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s TemplateMcpOptions) String() string {
	return dara.Prettify(s)
}

func (s TemplateMcpOptions) GoString() string {
	return s.String()
}

func (s *TemplateMcpOptions) GetEnabledTools() []*string {
	return s.EnabledTools
}

func (s *TemplateMcpOptions) GetTransport() *string {
	return s.Transport
}

func (s *TemplateMcpOptions) SetEnabledTools(v []*string) *TemplateMcpOptions {
	s.EnabledTools = v
	return s
}

func (s *TemplateMcpOptions) SetTransport(v string) *TemplateMcpOptions {
	s.Transport = &v
	return s
}

func (s *TemplateMcpOptions) Validate() error {
	return dara.Validate(s)
}

type TemplateMcpState struct {
	AccessEndpoint *string `json:"accessEndpoint,omitempty" xml:"accessEndpoint,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason   *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
}

func (s TemplateMcpState) String() string {
	return dara.Prettify(s)
}

func (s TemplateMcpState) GoString() string {
	return s.String()
}

func (s *TemplateMcpState) GetAccessEndpoint() *string {
	return s.AccessEndpoint
}

func (s *TemplateMcpState) GetStatus() *string {
	return s.Status
}

func (s *TemplateMcpState) GetStatusReason() *string {
	return s.StatusReason
}

func (s *TemplateMcpState) SetAccessEndpoint(v string) *TemplateMcpState {
	s.AccessEndpoint = &v
	return s
}

func (s *TemplateMcpState) SetStatus(v string) *TemplateMcpState {
	s.Status = &v
	return s
}

func (s *TemplateMcpState) SetStatusReason(v string) *TemplateMcpState {
	s.StatusReason = &v
	return s
}

func (s *TemplateMcpState) Validate() error {
	return dara.Validate(s)
}
