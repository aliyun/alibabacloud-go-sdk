// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTensorRtModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployTensorRtModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeployTensorRtModelInputConcurrencyConfig) *DeployTensorRtModelInput
	GetConcurrencyConfig() *DeployTensorRtModelInputConcurrencyConfig
	SetCpu(v float32) *DeployTensorRtModelInput
	GetCpu() *float32
	SetDescription(v string) *DeployTensorRtModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployTensorRtModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployTensorRtModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployTensorRtModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetGpuConfig(v *DeployTensorRtModelInputGpuConfig) *DeployTensorRtModelInput
	GetGpuConfig() *DeployTensorRtModelInputGpuConfig
	SetHttpTrigger(v *DeployTensorRtModelInputHttpTrigger) *DeployTensorRtModelInput
	GetHttpTrigger() *DeployTensorRtModelInputHttpTrigger
	SetImageName(v string) *DeployTensorRtModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeployTensorRtModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeployTensorRtModelInputLogConfig) *DeployTensorRtModelInput
	GetLogConfig() *DeployTensorRtModelInputLogConfig
	SetMemorySize(v int32) *DeployTensorRtModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployTensorRtModelInputModelConfig) *DeployTensorRtModelInput
	GetModelConfig() *DeployTensorRtModelInputModelConfig
	SetName(v string) *DeployTensorRtModelInput
	GetName() *string
	SetNasConfig(v *DeployTensorRtModelInputNasConfig) *DeployTensorRtModelInput
	GetNasConfig() *DeployTensorRtModelInputNasConfig
	SetOriginalName(v string) *DeployTensorRtModelInput
	GetOriginalName() *string
	SetProjectName(v string) *DeployTensorRtModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployTensorRtModelInputProvisionConfig) *DeployTensorRtModelInput
	GetProvisionConfig() *DeployTensorRtModelInputProvisionConfig
	SetRegion(v string) *DeployTensorRtModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployTensorRtModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployTensorRtModelInput
	GetRole() *string
	SetTimeout(v int32) *DeployTensorRtModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployTensorRtModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeployTensorRtModelInputVpcConfig) *DeployTensorRtModelInput
	GetVpcConfig() *DeployTensorRtModelInputVpcConfig
}

type DeployTensorRtModelInput struct {
	AccountID            *string                                    `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployTensorRtModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                   `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                     `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                    `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                     `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployTensorRtModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployTensorRtModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                    `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                     `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployTensorRtModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                     `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployTensorRtModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployTensorRtModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                  `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                  `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployTensorRtModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                  `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                  `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                            `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                             `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                            `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployTensorRtModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployTensorRtModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInput) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployTensorRtModelInput) GetConcurrencyConfig() *DeployTensorRtModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployTensorRtModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployTensorRtModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeployTensorRtModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployTensorRtModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployTensorRtModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployTensorRtModelInput) GetGpuConfig() *DeployTensorRtModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployTensorRtModelInput) GetHttpTrigger() *DeployTensorRtModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployTensorRtModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeployTensorRtModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployTensorRtModelInput) GetLogConfig() *DeployTensorRtModelInputLogConfig {
	return s.LogConfig
}

func (s *DeployTensorRtModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployTensorRtModelInput) GetModelConfig() *DeployTensorRtModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeployTensorRtModelInput) GetName() *string {
	return s.Name
}

func (s *DeployTensorRtModelInput) GetNasConfig() *DeployTensorRtModelInputNasConfig {
	return s.NasConfig
}

func (s *DeployTensorRtModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployTensorRtModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployTensorRtModelInput) GetProvisionConfig() *DeployTensorRtModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployTensorRtModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeployTensorRtModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployTensorRtModelInput) GetRole() *string {
	return s.Role
}

func (s *DeployTensorRtModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployTensorRtModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployTensorRtModelInput) GetVpcConfig() *DeployTensorRtModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployTensorRtModelInput) SetAccountID(v string) *DeployTensorRtModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployTensorRtModelInput) SetConcurrencyConfig(v *DeployTensorRtModelInputConcurrencyConfig) *DeployTensorRtModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetCpu(v float32) *DeployTensorRtModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployTensorRtModelInput) SetDescription(v string) *DeployTensorRtModelInput {
	s.Description = &v
	return s
}

func (s *DeployTensorRtModelInput) SetDiskSize(v int32) *DeployTensorRtModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployTensorRtModelInput) SetEnvName(v string) *DeployTensorRtModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployTensorRtModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployTensorRtModelInput) SetGpuConfig(v *DeployTensorRtModelInputGpuConfig) *DeployTensorRtModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetHttpTrigger(v *DeployTensorRtModelInputHttpTrigger) *DeployTensorRtModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployTensorRtModelInput) SetImageName(v string) *DeployTensorRtModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetInstanceConcurrency(v int32) *DeployTensorRtModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployTensorRtModelInput) SetLogConfig(v *DeployTensorRtModelInputLogConfig) *DeployTensorRtModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetMemorySize(v int32) *DeployTensorRtModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployTensorRtModelInput) SetModelConfig(v *DeployTensorRtModelInputModelConfig) *DeployTensorRtModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetName(v string) *DeployTensorRtModelInput {
	s.Name = &v
	return s
}

func (s *DeployTensorRtModelInput) SetNasConfig(v *DeployTensorRtModelInputNasConfig) *DeployTensorRtModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetOriginalName(v string) *DeployTensorRtModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetProjectName(v string) *DeployTensorRtModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployTensorRtModelInput) SetProvisionConfig(v *DeployTensorRtModelInputProvisionConfig) *DeployTensorRtModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployTensorRtModelInput) SetRegion(v string) *DeployTensorRtModelInput {
	s.Region = &v
	return s
}

func (s *DeployTensorRtModelInput) SetReportStatusURL(v string) *DeployTensorRtModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployTensorRtModelInput) SetRole(v string) *DeployTensorRtModelInput {
	s.Role = &v
	return s
}

func (s *DeployTensorRtModelInput) SetTimeout(v int32) *DeployTensorRtModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployTensorRtModelInput) SetTraceId(v string) *DeployTensorRtModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployTensorRtModelInput) SetVpcConfig(v *DeployTensorRtModelInputVpcConfig) *DeployTensorRtModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeployTensorRtModelInput) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployTensorRtModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployTensorRtModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployTensorRtModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployTensorRtModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployTensorRtModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeployTensorRtModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployTensorRtModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployTensorRtModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployTensorRtModelInputGpuConfig) SetGpuType(v string) *DeployTensorRtModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployTensorRtModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputHttpTrigger struct {
	Qualifier     *string                                           `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployTensorRtModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployTensorRtModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployTensorRtModelInputHttpTrigger) GetTriggerConfig() *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployTensorRtModelInputHttpTrigger) SetQualifier(v string) *DeployTensorRtModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTrigger) SetTriggerConfig(v *DeployTensorRtModelInputHttpTriggerTriggerConfig) *DeployTensorRtModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployTensorRtModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployTensorRtModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployTensorRtModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployTensorRtModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployTensorRtModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployTensorRtModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployTensorRtModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployTensorRtModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployTensorRtModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployTensorRtModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployTensorRtModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployTensorRtModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetLogBeginRule(v string) *DeployTensorRtModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetLogstore(v string) *DeployTensorRtModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) SetProject(v string) *DeployTensorRtModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployTensorRtModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputModelConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
	SyncStrategy               *string        `json:"syncStrategy,omitempty" xml:"syncStrategy,omitempty"`
}

func (s DeployTensorRtModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployTensorRtModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployTensorRtModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployTensorRtModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployTensorRtModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployTensorRtModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployTensorRtModelInputModelConfig) SetFramework(v string) *DeployTensorRtModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployTensorRtModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetPrefix(v string) *DeployTensorRtModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSourceType(v string) *DeployTensorRtModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssBucket(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssPath(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSrcOssRegion(v string) *DeployTensorRtModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) SetSyncStrategy(v string) *DeployTensorRtModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployTensorRtModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputNasConfig struct {
	GroupId     *int32                                          `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*DeployTensorRtModelInputNasConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                                          `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployTensorRtModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeployTensorRtModelInputNasConfig) GetMountPoints() []*DeployTensorRtModelInputNasConfigMountPoints {
	return s.MountPoints
}

func (s *DeployTensorRtModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeployTensorRtModelInputNasConfig) SetGroupId(v int32) *DeployTensorRtModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfig) SetMountPoints(v []*DeployTensorRtModelInputNasConfigMountPoints) *DeployTensorRtModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployTensorRtModelInputNasConfig) SetUserId(v int32) *DeployTensorRtModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputNasConfigMountPoints struct {
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s DeployTensorRtModelInputNasConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputNasConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) GetEnableTLS() *bool {
	return s.EnableTLS
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetEnableTLS(v bool) *DeployTensorRtModelInputNasConfigMountPoints {
	s.EnableTLS = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetMountDir(v string) *DeployTensorRtModelInputNasConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) SetServerAddr(v string) *DeployTensorRtModelInputNasConfigMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *DeployTensorRtModelInputNasConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                      `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployTensorRtModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                     `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployTensorRtModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployTensorRtModelInputProvisionConfig) GetScheduledActions() []*DeployTensorRtModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployTensorRtModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeployTensorRtModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployTensorRtModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfig) SetScheduledActions(v []*DeployTensorRtModelInputProvisionConfigScheduledActions) *DeployTensorRtModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfig) SetTarget(v int32) *DeployTensorRtModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployTensorRtModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetName(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployTensorRtModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployTensorRtModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployTensorRtModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployTensorRtModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployTensorRtModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployTensorRtModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployTensorRtModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployTensorRtModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployTensorRtModelInputVpcConfig) SetSecurityGroupId(v string) *DeployTensorRtModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployTensorRtModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployTensorRtModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployTensorRtModelInputVpcConfig) SetVpcId(v string) *DeployTensorRtModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployTensorRtModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
