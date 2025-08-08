// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployOllamaModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployOllamaModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeployOllamaModelInputConcurrencyConfig) *DeployOllamaModelInput
	GetConcurrencyConfig() *DeployOllamaModelInputConcurrencyConfig
	SetCpu(v float32) *DeployOllamaModelInput
	GetCpu() *float32
	SetDescription(v string) *DeployOllamaModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployOllamaModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployOllamaModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployOllamaModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetFeatureGates(v *DeployOllamaModelInputFeatureGates) *DeployOllamaModelInput
	GetFeatureGates() *DeployOllamaModelInputFeatureGates
	SetGpuConfig(v *DeployOllamaModelInputGpuConfig) *DeployOllamaModelInput
	GetGpuConfig() *DeployOllamaModelInputGpuConfig
	SetHttpTrigger(v *DeployOllamaModelInputHttpTrigger) *DeployOllamaModelInput
	GetHttpTrigger() *DeployOllamaModelInputHttpTrigger
	SetImageName(v string) *DeployOllamaModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeployOllamaModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeployOllamaModelInputLogConfig) *DeployOllamaModelInput
	GetLogConfig() *DeployOllamaModelInputLogConfig
	SetMemorySize(v int32) *DeployOllamaModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployOllamaModelInputModelConfig) *DeployOllamaModelInput
	GetModelConfig() *DeployOllamaModelInputModelConfig
	SetName(v string) *DeployOllamaModelInput
	GetName() *string
	SetNasConfig(v *DeployOllamaModelInputNasConfig) *DeployOllamaModelInput
	GetNasConfig() *DeployOllamaModelInputNasConfig
	SetOriginalName(v string) *DeployOllamaModelInput
	GetOriginalName() *string
	SetProjectName(v string) *DeployOllamaModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployOllamaModelInputProvisionConfig) *DeployOllamaModelInput
	GetProvisionConfig() *DeployOllamaModelInputProvisionConfig
	SetRegion(v string) *DeployOllamaModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployOllamaModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployOllamaModelInput
	GetRole() *string
	SetTimeout(v int32) *DeployOllamaModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployOllamaModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeployOllamaModelInputVpcConfig) *DeployOllamaModelInput
	GetVpcConfig() *DeployOllamaModelInputVpcConfig
}

type DeployOllamaModelInput struct {
	AccountID            *string                                  `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployOllamaModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                  `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                   `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	FeatureGates         *DeployOllamaModelInputFeatureGates      `json:"featureGates,omitempty" xml:"featureGates,omitempty" type:"Struct"`
	GpuConfig            *DeployOllamaModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployOllamaModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                  `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployOllamaModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployOllamaModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployOllamaModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployOllamaModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                          `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                           `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployOllamaModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployOllamaModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInput) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployOllamaModelInput) GetConcurrencyConfig() *DeployOllamaModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployOllamaModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployOllamaModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeployOllamaModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployOllamaModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployOllamaModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployOllamaModelInput) GetFeatureGates() *DeployOllamaModelInputFeatureGates {
	return s.FeatureGates
}

func (s *DeployOllamaModelInput) GetGpuConfig() *DeployOllamaModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployOllamaModelInput) GetHttpTrigger() *DeployOllamaModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployOllamaModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeployOllamaModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployOllamaModelInput) GetLogConfig() *DeployOllamaModelInputLogConfig {
	return s.LogConfig
}

func (s *DeployOllamaModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployOllamaModelInput) GetModelConfig() *DeployOllamaModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeployOllamaModelInput) GetName() *string {
	return s.Name
}

func (s *DeployOllamaModelInput) GetNasConfig() *DeployOllamaModelInputNasConfig {
	return s.NasConfig
}

func (s *DeployOllamaModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployOllamaModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployOllamaModelInput) GetProvisionConfig() *DeployOllamaModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployOllamaModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeployOllamaModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployOllamaModelInput) GetRole() *string {
	return s.Role
}

func (s *DeployOllamaModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployOllamaModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployOllamaModelInput) GetVpcConfig() *DeployOllamaModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployOllamaModelInput) SetAccountID(v string) *DeployOllamaModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployOllamaModelInput) SetConcurrencyConfig(v *DeployOllamaModelInputConcurrencyConfig) *DeployOllamaModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetCpu(v float32) *DeployOllamaModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployOllamaModelInput) SetDescription(v string) *DeployOllamaModelInput {
	s.Description = &v
	return s
}

func (s *DeployOllamaModelInput) SetDiskSize(v int32) *DeployOllamaModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployOllamaModelInput) SetEnvName(v string) *DeployOllamaModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployOllamaModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployOllamaModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployOllamaModelInput) SetFeatureGates(v *DeployOllamaModelInputFeatureGates) *DeployOllamaModelInput {
	s.FeatureGates = v
	return s
}

func (s *DeployOllamaModelInput) SetGpuConfig(v *DeployOllamaModelInputGpuConfig) *DeployOllamaModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetHttpTrigger(v *DeployOllamaModelInputHttpTrigger) *DeployOllamaModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployOllamaModelInput) SetImageName(v string) *DeployOllamaModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployOllamaModelInput) SetInstanceConcurrency(v int32) *DeployOllamaModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployOllamaModelInput) SetLogConfig(v *DeployOllamaModelInputLogConfig) *DeployOllamaModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetMemorySize(v int32) *DeployOllamaModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployOllamaModelInput) SetModelConfig(v *DeployOllamaModelInputModelConfig) *DeployOllamaModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetName(v string) *DeployOllamaModelInput {
	s.Name = &v
	return s
}

func (s *DeployOllamaModelInput) SetNasConfig(v *DeployOllamaModelInputNasConfig) *DeployOllamaModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetOriginalName(v string) *DeployOllamaModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployOllamaModelInput) SetProjectName(v string) *DeployOllamaModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployOllamaModelInput) SetProvisionConfig(v *DeployOllamaModelInputProvisionConfig) *DeployOllamaModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployOllamaModelInput) SetRegion(v string) *DeployOllamaModelInput {
	s.Region = &v
	return s
}

func (s *DeployOllamaModelInput) SetReportStatusURL(v string) *DeployOllamaModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployOllamaModelInput) SetRole(v string) *DeployOllamaModelInput {
	s.Role = &v
	return s
}

func (s *DeployOllamaModelInput) SetTimeout(v int32) *DeployOllamaModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployOllamaModelInput) SetTraceId(v string) *DeployOllamaModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployOllamaModelInput) SetVpcConfig(v *DeployOllamaModelInputVpcConfig) *DeployOllamaModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeployOllamaModelInput) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployOllamaModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployOllamaModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployOllamaModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployOllamaModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputFeatureGates struct {
	AsyncProvisionCheck               *bool `json:"asyncProvisionCheck,omitempty" xml:"asyncProvisionCheck,omitempty"`
	DisableRollbackOnProvisionFailure *bool `json:"disableRollbackOnProvisionFailure,omitempty" xml:"disableRollbackOnProvisionFailure,omitempty"`
}

func (s DeployOllamaModelInputFeatureGates) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputFeatureGates) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputFeatureGates) GetAsyncProvisionCheck() *bool {
	return s.AsyncProvisionCheck
}

func (s *DeployOllamaModelInputFeatureGates) GetDisableRollbackOnProvisionFailure() *bool {
	return s.DisableRollbackOnProvisionFailure
}

func (s *DeployOllamaModelInputFeatureGates) SetAsyncProvisionCheck(v bool) *DeployOllamaModelInputFeatureGates {
	s.AsyncProvisionCheck = &v
	return s
}

func (s *DeployOllamaModelInputFeatureGates) SetDisableRollbackOnProvisionFailure(v bool) *DeployOllamaModelInputFeatureGates {
	s.DisableRollbackOnProvisionFailure = &v
	return s
}

func (s *DeployOllamaModelInputFeatureGates) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployOllamaModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeployOllamaModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployOllamaModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployOllamaModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployOllamaModelInputGpuConfig) SetGpuType(v string) *DeployOllamaModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployOllamaModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputHttpTrigger struct {
	Qualifier     *string                                         `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployOllamaModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployOllamaModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployOllamaModelInputHttpTrigger) GetTriggerConfig() *DeployOllamaModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployOllamaModelInputHttpTrigger) SetQualifier(v string) *DeployOllamaModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployOllamaModelInputHttpTrigger) SetTriggerConfig(v *DeployOllamaModelInputHttpTriggerTriggerConfig) *DeployOllamaModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployOllamaModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployOllamaModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployOllamaModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployOllamaModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployOllamaModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployOllamaModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployOllamaModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployOllamaModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployOllamaModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployOllamaModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployOllamaModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployOllamaModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetLogBeginRule(v string) *DeployOllamaModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetLogstore(v string) *DeployOllamaModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) SetProject(v string) *DeployOllamaModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployOllamaModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputModelConfig struct {
	FmkOllamaConfig *DeployOllamaModelInputModelConfigFmkOllamaConfig `json:"fmkOllamaConfig,omitempty" xml:"fmkOllamaConfig,omitempty" type:"Struct"`
	Framework       *string                                           `json:"framework,omitempty" xml:"framework,omitempty"`
	// if can be null:
	// true
	MultiModelConfig           []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Prefix                     *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SkipDownload               *bool          `json:"skipDownload,omitempty" xml:"skipDownload,omitempty"`
	SourceType                 *string        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SrcModelScopeModelID       *string        `json:"srcModelScopeModelID,omitempty" xml:"srcModelScopeModelID,omitempty"`
	SrcModelScopeModelRevision *string        `json:"srcModelScopeModelRevision,omitempty" xml:"srcModelScopeModelRevision,omitempty"`
	SrcModelScopeToken         *string        `json:"srcModelScopeToken,omitempty" xml:"srcModelScopeToken,omitempty"`
	SrcOssBucket               *string        `json:"srcOssBucket,omitempty" xml:"srcOssBucket,omitempty"`
	SrcOssPath                 *string        `json:"srcOssPath,omitempty" xml:"srcOssPath,omitempty"`
	SrcOssRegion               *string        `json:"srcOssRegion,omitempty" xml:"srcOssRegion,omitempty"`
	SyncStrategy               *string        `json:"syncStrategy,omitempty" xml:"syncStrategy,omitempty"`
	WithPPU                    *bool          `json:"withPPU,omitempty" xml:"withPPU,omitempty"`
	WorkingDir                 *string        `json:"workingDir,omitempty" xml:"workingDir,omitempty"`
}

func (s DeployOllamaModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputModelConfig) GetFmkOllamaConfig() *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	return s.FmkOllamaConfig
}

func (s *DeployOllamaModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployOllamaModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployOllamaModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployOllamaModelInputModelConfig) GetSkipDownload() *bool {
	return s.SkipDownload
}

func (s *DeployOllamaModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployOllamaModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployOllamaModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployOllamaModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployOllamaModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployOllamaModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployOllamaModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployOllamaModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployOllamaModelInputModelConfig) GetWithPPU() *bool {
	return s.WithPPU
}

func (s *DeployOllamaModelInputModelConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DeployOllamaModelInputModelConfig) SetFmkOllamaConfig(v *DeployOllamaModelInputModelConfigFmkOllamaConfig) *DeployOllamaModelInputModelConfig {
	s.FmkOllamaConfig = v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetFramework(v string) *DeployOllamaModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployOllamaModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetPrefix(v string) *DeployOllamaModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSkipDownload(v bool) *DeployOllamaModelInputModelConfig {
	s.SkipDownload = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSourceType(v string) *DeployOllamaModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployOllamaModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssBucket(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssPath(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSrcOssRegion(v string) *DeployOllamaModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetSyncStrategy(v string) *DeployOllamaModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetWithPPU(v bool) *DeployOllamaModelInputModelConfig {
	s.WithPPU = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) SetWorkingDir(v string) *DeployOllamaModelInputModelConfig {
	s.WorkingDir = &v
	return s
}

func (s *DeployOllamaModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputModelConfigFmkOllamaConfig struct {
	MinP                           *float32 `json:"minP,omitempty" xml:"minP,omitempty"`
	Mirostat                       *int32   `json:"mirostat,omitempty" xml:"mirostat,omitempty"`
	MirostatEta                    *float32 `json:"mirostatEta,omitempty" xml:"mirostatEta,omitempty"`
	MirostatTau                    *float32 `json:"mirostatTau,omitempty" xml:"mirostatTau,omitempty"`
	ModelName                      *string  `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelfileAdapter               *string  `json:"modelfileAdapter,omitempty" xml:"modelfileAdapter,omitempty"`
	ModelfileAdditionalFromsString *string  `json:"modelfileAdditionalFromsString,omitempty" xml:"modelfileAdditionalFromsString,omitempty"`
	ModelfileFullTextPostfix       *string  `json:"modelfileFullTextPostfix,omitempty" xml:"modelfileFullTextPostfix,omitempty"`
	ModelfileParams                *string  `json:"modelfileParams,omitempty" xml:"modelfileParams,omitempty"`
	ModelfileSystem                *string  `json:"modelfileSystem,omitempty" xml:"modelfileSystem,omitempty"`
	ModelfileTemplate              *string  `json:"modelfileTemplate,omitempty" xml:"modelfileTemplate,omitempty"`
	NumCtx                         *int32   `json:"numCtx,omitempty" xml:"numCtx,omitempty"`
	NumPredict                     *int32   `json:"numPredict,omitempty" xml:"numPredict,omitempty"`
	Quantize                       *string  `json:"quantize,omitempty" xml:"quantize,omitempty"`
	RepeatLastN                    *int32   `json:"repeatLastN,omitempty" xml:"repeatLastN,omitempty"`
	RepeatPenalty                  *float32 `json:"repeatPenalty,omitempty" xml:"repeatPenalty,omitempty"`
	Seed                           *int32   `json:"seed,omitempty" xml:"seed,omitempty"`
	SingleModelFile                *string  `json:"singleModelFile,omitempty" xml:"singleModelFile,omitempty"`
	SplitedModelStartFile          *string  `json:"splitedModelStartFile,omitempty" xml:"splitedModelStartFile,omitempty"`
	Stop                           *string  `json:"stop,omitempty" xml:"stop,omitempty"`
	Stream                         *bool    `json:"stream,omitempty" xml:"stream,omitempty"`
	Temperature                    *float32 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	TfsZ                           *float32 `json:"tfsZ,omitempty" xml:"tfsZ,omitempty"`
	TopK                           *int32   `json:"topK,omitempty" xml:"topK,omitempty"`
	TopP                           *float32 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s DeployOllamaModelInputModelConfigFmkOllamaConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputModelConfigFmkOllamaConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetMinP() *float32 {
	return s.MinP
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetMirostat() *int32 {
	return s.Mirostat
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetMirostatEta() *float32 {
	return s.MirostatEta
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetMirostatTau() *float32 {
	return s.MirostatTau
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelName() *string {
	return s.ModelName
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileAdapter() *string {
	return s.ModelfileAdapter
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileAdditionalFromsString() *string {
	return s.ModelfileAdditionalFromsString
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileFullTextPostfix() *string {
	return s.ModelfileFullTextPostfix
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileParams() *string {
	return s.ModelfileParams
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileSystem() *string {
	return s.ModelfileSystem
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetModelfileTemplate() *string {
	return s.ModelfileTemplate
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetNumCtx() *int32 {
	return s.NumCtx
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetNumPredict() *int32 {
	return s.NumPredict
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetQuantize() *string {
	return s.Quantize
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetRepeatLastN() *int32 {
	return s.RepeatLastN
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetRepeatPenalty() *float32 {
	return s.RepeatPenalty
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetSeed() *int32 {
	return s.Seed
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetSingleModelFile() *string {
	return s.SingleModelFile
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetSplitedModelStartFile() *string {
	return s.SplitedModelStartFile
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetStop() *string {
	return s.Stop
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetStream() *bool {
	return s.Stream
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetTfsZ() *float32 {
	return s.TfsZ
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMinP(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MinP = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostat(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Mirostat = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostatEta(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MirostatEta = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetMirostatTau(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.MirostatTau = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelName(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelName = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileAdapter(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileAdapter = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileAdditionalFromsString(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileAdditionalFromsString = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileFullTextPostfix(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileFullTextPostfix = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileParams(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileParams = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileSystem(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileSystem = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetModelfileTemplate(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.ModelfileTemplate = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetNumCtx(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.NumCtx = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetNumPredict(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.NumPredict = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetQuantize(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Quantize = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetRepeatLastN(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.RepeatLastN = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetRepeatPenalty(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.RepeatPenalty = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSeed(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Seed = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSingleModelFile(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.SingleModelFile = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetSplitedModelStartFile(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.SplitedModelStartFile = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetStop(v string) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Stop = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetStream(v bool) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Stream = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTemperature(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.Temperature = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTfsZ(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TfsZ = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTopK(v int32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TopK = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) SetTopP(v float32) *DeployOllamaModelInputModelConfigFmkOllamaConfig {
	s.TopP = &v
	return s
}

func (s *DeployOllamaModelInputModelConfigFmkOllamaConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployOllamaModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeployOllamaModelInputNasConfig) GetMountPoints() []*string {
	return s.MountPoints
}

func (s *DeployOllamaModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeployOllamaModelInputNasConfig) SetGroupId(v int32) *DeployOllamaModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployOllamaModelInputNasConfig) SetMountPoints(v []*string) *DeployOllamaModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployOllamaModelInputNasConfig) SetUserId(v int32) *DeployOllamaModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployOllamaModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                    `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployOllamaModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                   `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployOllamaModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployOllamaModelInputProvisionConfig) GetScheduledActions() []*DeployOllamaModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployOllamaModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeployOllamaModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployOllamaModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfig) SetScheduledActions(v []*DeployOllamaModelInputProvisionConfigScheduledActions) *DeployOllamaModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployOllamaModelInputProvisionConfig) SetTarget(v int32) *DeployOllamaModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployOllamaModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetName(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployOllamaModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployOllamaModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployOllamaModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployOllamaModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployOllamaModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployOllamaModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployOllamaModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployOllamaModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployOllamaModelInputVpcConfig) SetSecurityGroupId(v string) *DeployOllamaModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployOllamaModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployOllamaModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployOllamaModelInputVpcConfig) SetVpcId(v string) *DeployOllamaModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployOllamaModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
