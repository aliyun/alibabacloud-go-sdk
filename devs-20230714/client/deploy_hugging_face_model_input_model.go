// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHuggingFaceModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployHuggingFaceModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeployHuggingFaceModelInputConcurrencyConfig) *DeployHuggingFaceModelInput
	GetConcurrencyConfig() *DeployHuggingFaceModelInputConcurrencyConfig
	SetCpu(v float32) *DeployHuggingFaceModelInput
	GetCpu() *float32
	SetDescription(v string) *DeployHuggingFaceModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployHuggingFaceModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployHuggingFaceModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployHuggingFaceModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetFeatureGates(v *DeployHuggingFaceModelInputFeatureGates) *DeployHuggingFaceModelInput
	GetFeatureGates() *DeployHuggingFaceModelInputFeatureGates
	SetGpuConfig(v *DeployHuggingFaceModelInputGpuConfig) *DeployHuggingFaceModelInput
	GetGpuConfig() *DeployHuggingFaceModelInputGpuConfig
	SetHttpTrigger(v *DeployHuggingFaceModelInputHttpTrigger) *DeployHuggingFaceModelInput
	GetHttpTrigger() *DeployHuggingFaceModelInputHttpTrigger
	SetImageName(v string) *DeployHuggingFaceModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeployHuggingFaceModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeployHuggingFaceModelInputLogConfig) *DeployHuggingFaceModelInput
	GetLogConfig() *DeployHuggingFaceModelInputLogConfig
	SetMemorySize(v int32) *DeployHuggingFaceModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployHuggingFaceModelInputModelConfig) *DeployHuggingFaceModelInput
	GetModelConfig() *DeployHuggingFaceModelInputModelConfig
	SetName(v string) *DeployHuggingFaceModelInput
	GetName() *string
	SetNasConfig(v *DeployHuggingFaceModelInputNasConfig) *DeployHuggingFaceModelInput
	GetNasConfig() *DeployHuggingFaceModelInputNasConfig
	SetOriginalName(v string) *DeployHuggingFaceModelInput
	GetOriginalName() *string
	SetProjectName(v string) *DeployHuggingFaceModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployHuggingFaceModelInputProvisionConfig) *DeployHuggingFaceModelInput
	GetProvisionConfig() *DeployHuggingFaceModelInputProvisionConfig
	SetRegion(v string) *DeployHuggingFaceModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployHuggingFaceModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployHuggingFaceModelInput
	GetRole() *string
	SetTimeout(v int32) *DeployHuggingFaceModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployHuggingFaceModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeployHuggingFaceModelInputVpcConfig) *DeployHuggingFaceModelInput
	GetVpcConfig() *DeployHuggingFaceModelInputVpcConfig
}

type DeployHuggingFaceModelInput struct {
	AccountID            *string                                       `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployHuggingFaceModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                      `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                        `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                       `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                        `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	FeatureGates         *DeployHuggingFaceModelInputFeatureGates      `json:"featureGates,omitempty" xml:"featureGates,omitempty" type:"Struct"`
	GpuConfig            *DeployHuggingFaceModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployHuggingFaceModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                       `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                        `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployHuggingFaceModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                        `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployHuggingFaceModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployHuggingFaceModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                     `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                     `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployHuggingFaceModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                     `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                     `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                               `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                                `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployHuggingFaceModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployHuggingFaceModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInput) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployHuggingFaceModelInput) GetConcurrencyConfig() *DeployHuggingFaceModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployHuggingFaceModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployHuggingFaceModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeployHuggingFaceModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployHuggingFaceModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployHuggingFaceModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployHuggingFaceModelInput) GetFeatureGates() *DeployHuggingFaceModelInputFeatureGates {
	return s.FeatureGates
}

func (s *DeployHuggingFaceModelInput) GetGpuConfig() *DeployHuggingFaceModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployHuggingFaceModelInput) GetHttpTrigger() *DeployHuggingFaceModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployHuggingFaceModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeployHuggingFaceModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployHuggingFaceModelInput) GetLogConfig() *DeployHuggingFaceModelInputLogConfig {
	return s.LogConfig
}

func (s *DeployHuggingFaceModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployHuggingFaceModelInput) GetModelConfig() *DeployHuggingFaceModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeployHuggingFaceModelInput) GetName() *string {
	return s.Name
}

func (s *DeployHuggingFaceModelInput) GetNasConfig() *DeployHuggingFaceModelInputNasConfig {
	return s.NasConfig
}

func (s *DeployHuggingFaceModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployHuggingFaceModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployHuggingFaceModelInput) GetProvisionConfig() *DeployHuggingFaceModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployHuggingFaceModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeployHuggingFaceModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployHuggingFaceModelInput) GetRole() *string {
	return s.Role
}

func (s *DeployHuggingFaceModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployHuggingFaceModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployHuggingFaceModelInput) GetVpcConfig() *DeployHuggingFaceModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployHuggingFaceModelInput) SetAccountID(v string) *DeployHuggingFaceModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetConcurrencyConfig(v *DeployHuggingFaceModelInputConcurrencyConfig) *DeployHuggingFaceModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetCpu(v float32) *DeployHuggingFaceModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetDescription(v string) *DeployHuggingFaceModelInput {
	s.Description = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetDiskSize(v int32) *DeployHuggingFaceModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetEnvName(v string) *DeployHuggingFaceModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployHuggingFaceModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetFeatureGates(v *DeployHuggingFaceModelInputFeatureGates) *DeployHuggingFaceModelInput {
	s.FeatureGates = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetGpuConfig(v *DeployHuggingFaceModelInputGpuConfig) *DeployHuggingFaceModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetHttpTrigger(v *DeployHuggingFaceModelInputHttpTrigger) *DeployHuggingFaceModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetImageName(v string) *DeployHuggingFaceModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetInstanceConcurrency(v int32) *DeployHuggingFaceModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetLogConfig(v *DeployHuggingFaceModelInputLogConfig) *DeployHuggingFaceModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetMemorySize(v int32) *DeployHuggingFaceModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetModelConfig(v *DeployHuggingFaceModelInputModelConfig) *DeployHuggingFaceModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetName(v string) *DeployHuggingFaceModelInput {
	s.Name = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetNasConfig(v *DeployHuggingFaceModelInputNasConfig) *DeployHuggingFaceModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetOriginalName(v string) *DeployHuggingFaceModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetProjectName(v string) *DeployHuggingFaceModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetProvisionConfig(v *DeployHuggingFaceModelInputProvisionConfig) *DeployHuggingFaceModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) SetRegion(v string) *DeployHuggingFaceModelInput {
	s.Region = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetReportStatusURL(v string) *DeployHuggingFaceModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetRole(v string) *DeployHuggingFaceModelInput {
	s.Role = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetTimeout(v int32) *DeployHuggingFaceModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetTraceId(v string) *DeployHuggingFaceModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployHuggingFaceModelInput) SetVpcConfig(v *DeployHuggingFaceModelInputVpcConfig) *DeployHuggingFaceModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeployHuggingFaceModelInput) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployHuggingFaceModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployHuggingFaceModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployHuggingFaceModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployHuggingFaceModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputFeatureGates struct {
	AsyncProvisionCheck               *bool `json:"asyncProvisionCheck,omitempty" xml:"asyncProvisionCheck,omitempty"`
	DisableRollbackOnProvisionFailure *bool `json:"disableRollbackOnProvisionFailure,omitempty" xml:"disableRollbackOnProvisionFailure,omitempty"`
}

func (s DeployHuggingFaceModelInputFeatureGates) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputFeatureGates) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputFeatureGates) GetAsyncProvisionCheck() *bool {
	return s.AsyncProvisionCheck
}

func (s *DeployHuggingFaceModelInputFeatureGates) GetDisableRollbackOnProvisionFailure() *bool {
	return s.DisableRollbackOnProvisionFailure
}

func (s *DeployHuggingFaceModelInputFeatureGates) SetAsyncProvisionCheck(v bool) *DeployHuggingFaceModelInputFeatureGates {
	s.AsyncProvisionCheck = &v
	return s
}

func (s *DeployHuggingFaceModelInputFeatureGates) SetDisableRollbackOnProvisionFailure(v bool) *DeployHuggingFaceModelInputFeatureGates {
	s.DisableRollbackOnProvisionFailure = &v
	return s
}

func (s *DeployHuggingFaceModelInputFeatureGates) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployHuggingFaceModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeployHuggingFaceModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployHuggingFaceModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployHuggingFaceModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployHuggingFaceModelInputGpuConfig) SetGpuType(v string) *DeployHuggingFaceModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployHuggingFaceModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputHttpTrigger struct {
	Qualifier     *string                                              `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployHuggingFaceModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployHuggingFaceModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployHuggingFaceModelInputHttpTrigger) GetTriggerConfig() *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployHuggingFaceModelInputHttpTrigger) SetQualifier(v string) *DeployHuggingFaceModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTrigger) SetTriggerConfig(v *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) *DeployHuggingFaceModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployHuggingFaceModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployHuggingFaceModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployHuggingFaceModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployHuggingFaceModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployHuggingFaceModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployHuggingFaceModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployHuggingFaceModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployHuggingFaceModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployHuggingFaceModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployHuggingFaceModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployHuggingFaceModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetLogBeginRule(v string) *DeployHuggingFaceModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetLogstore(v string) *DeployHuggingFaceModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) SetProject(v string) *DeployHuggingFaceModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployHuggingFaceModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputModelConfig struct {
	FmkHuggingFaceConfig *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig `json:"fmkHuggingFaceConfig,omitempty" xml:"fmkHuggingFaceConfig,omitempty" type:"Struct"`
	Framework            *string                                                     `json:"framework,omitempty" xml:"framework,omitempty"`
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

func (s DeployHuggingFaceModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputModelConfig) GetFmkHuggingFaceConfig() *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig {
	return s.FmkHuggingFaceConfig
}

func (s *DeployHuggingFaceModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployHuggingFaceModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployHuggingFaceModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSkipDownload() *bool {
	return s.SkipDownload
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployHuggingFaceModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployHuggingFaceModelInputModelConfig) GetWithPPU() *bool {
	return s.WithPPU
}

func (s *DeployHuggingFaceModelInputModelConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DeployHuggingFaceModelInputModelConfig) SetFmkHuggingFaceConfig(v *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) *DeployHuggingFaceModelInputModelConfig {
	s.FmkHuggingFaceConfig = v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetFramework(v string) *DeployHuggingFaceModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployHuggingFaceModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetPrefix(v string) *DeployHuggingFaceModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSkipDownload(v bool) *DeployHuggingFaceModelInputModelConfig {
	s.SkipDownload = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSourceType(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssBucket(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssPath(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSrcOssRegion(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetSyncStrategy(v string) *DeployHuggingFaceModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetWithPPU(v bool) *DeployHuggingFaceModelInputModelConfig {
	s.WithPPU = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) SetWorkingDir(v string) *DeployHuggingFaceModelInputModelConfig {
	s.WorkingDir = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	Task      *string `json:"task,omitempty" xml:"task,omitempty"`
}

func (s DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) GetTask() *string {
	return s.Task
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) SetFramework(v string) *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig {
	s.Framework = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) SetTask(v string) *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig {
	s.Task = &v
	return s
}

func (s *DeployHuggingFaceModelInputModelConfigFmkHuggingFaceConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployHuggingFaceModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeployHuggingFaceModelInputNasConfig) GetMountPoints() []*string {
	return s.MountPoints
}

func (s *DeployHuggingFaceModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeployHuggingFaceModelInputNasConfig) SetGroupId(v int32) *DeployHuggingFaceModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployHuggingFaceModelInputNasConfig) SetMountPoints(v []*string) *DeployHuggingFaceModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployHuggingFaceModelInputNasConfig) SetUserId(v int32) *DeployHuggingFaceModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployHuggingFaceModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                         `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployHuggingFaceModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                        `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployHuggingFaceModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployHuggingFaceModelInputProvisionConfig) GetScheduledActions() []*DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployHuggingFaceModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployHuggingFaceModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetScheduledActions(v []*DeployHuggingFaceModelInputProvisionConfigScheduledActions) *DeployHuggingFaceModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfig) SetTarget(v int32) *DeployHuggingFaceModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployHuggingFaceModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetName(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployHuggingFaceModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployHuggingFaceModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployHuggingFaceModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployHuggingFaceModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHuggingFaceModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployHuggingFaceModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployHuggingFaceModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployHuggingFaceModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetSecurityGroupId(v string) *DeployHuggingFaceModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployHuggingFaceModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployHuggingFaceModelInputVpcConfig) SetVpcId(v string) *DeployHuggingFaceModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployHuggingFaceModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
