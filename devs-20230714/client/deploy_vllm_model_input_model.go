// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployVllmModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployVllmModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeployVllmModelInputConcurrencyConfig) *DeployVllmModelInput
	GetConcurrencyConfig() *DeployVllmModelInputConcurrencyConfig
	SetCpu(v float32) *DeployVllmModelInput
	GetCpu() *float32
	SetCustomContainerConfig(v *DeployVllmModelInputCustomContainerConfig) *DeployVllmModelInput
	GetCustomContainerConfig() *DeployVllmModelInputCustomContainerConfig
	SetDescription(v string) *DeployVllmModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployVllmModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployVllmModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployVllmModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetFeatureGates(v *DeployVllmModelInputFeatureGates) *DeployVllmModelInput
	GetFeatureGates() *DeployVllmModelInputFeatureGates
	SetGpuConfig(v *DeployVllmModelInputGpuConfig) *DeployVllmModelInput
	GetGpuConfig() *DeployVllmModelInputGpuConfig
	SetHttpTrigger(v *DeployVllmModelInputHttpTrigger) *DeployVllmModelInput
	GetHttpTrigger() *DeployVllmModelInputHttpTrigger
	SetImageName(v string) *DeployVllmModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeployVllmModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeployVllmModelInputLogConfig) *DeployVllmModelInput
	GetLogConfig() *DeployVllmModelInputLogConfig
	SetMemorySize(v int32) *DeployVllmModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployVllmModelInputModelConfig) *DeployVllmModelInput
	GetModelConfig() *DeployVllmModelInputModelConfig
	SetName(v string) *DeployVllmModelInput
	GetName() *string
	SetNasConfig(v *DeployVllmModelInputNasConfig) *DeployVllmModelInput
	GetNasConfig() *DeployVllmModelInputNasConfig
	SetOriginalName(v string) *DeployVllmModelInput
	GetOriginalName() *string
	SetOssMountConfig(v *DeployVllmModelInputOssMountConfig) *DeployVllmModelInput
	GetOssMountConfig() *DeployVllmModelInputOssMountConfig
	SetProjectName(v string) *DeployVllmModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployVllmModelInputProvisionConfig) *DeployVllmModelInput
	GetProvisionConfig() *DeployVllmModelInputProvisionConfig
	SetRegion(v string) *DeployVllmModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployVllmModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployVllmModelInput
	GetRole() *string
	SetTimeout(v int32) *DeployVllmModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployVllmModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeployVllmModelInputVpcConfig) *DeployVllmModelInput
	GetVpcConfig() *DeployVllmModelInputVpcConfig
}

type DeployVllmModelInput struct {
	AccountID             *string                                    `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig     *DeployVllmModelInputConcurrencyConfig     `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                   *float32                                   `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *DeployVllmModelInputCustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty" type:"Struct"`
	Description           *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize              *int32                                     `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName               *string                                    `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables  map[string]interface{}                     `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	FeatureGates          *DeployVllmModelInputFeatureGates          `json:"featureGates,omitempty" xml:"featureGates,omitempty" type:"Struct"`
	GpuConfig             *DeployVllmModelInputGpuConfig             `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger           *DeployVllmModelInputHttpTrigger           `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName             *string                                    `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency   *int32                                     `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig             *DeployVllmModelInputLogConfig             `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize            *int32                                     `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig           *DeployVllmModelInputModelConfig           `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                              `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployVllmModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                              `json:"originalName,omitempty" xml:"originalName,omitempty"`
	OssMountConfig  *DeployVllmModelInputOssMountConfig  `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty" type:"Struct"`
	ProjectName     *string                              `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployVllmModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                              `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                              `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                        `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                        `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployVllmModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployVllmModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInput) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployVllmModelInput) GetConcurrencyConfig() *DeployVllmModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployVllmModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployVllmModelInput) GetCustomContainerConfig() *DeployVllmModelInputCustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *DeployVllmModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeployVllmModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployVllmModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployVllmModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployVllmModelInput) GetFeatureGates() *DeployVllmModelInputFeatureGates {
	return s.FeatureGates
}

func (s *DeployVllmModelInput) GetGpuConfig() *DeployVllmModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployVllmModelInput) GetHttpTrigger() *DeployVllmModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployVllmModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeployVllmModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployVllmModelInput) GetLogConfig() *DeployVllmModelInputLogConfig {
	return s.LogConfig
}

func (s *DeployVllmModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployVllmModelInput) GetModelConfig() *DeployVllmModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeployVllmModelInput) GetName() *string {
	return s.Name
}

func (s *DeployVllmModelInput) GetNasConfig() *DeployVllmModelInputNasConfig {
	return s.NasConfig
}

func (s *DeployVllmModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployVllmModelInput) GetOssMountConfig() *DeployVllmModelInputOssMountConfig {
	return s.OssMountConfig
}

func (s *DeployVllmModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployVllmModelInput) GetProvisionConfig() *DeployVllmModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployVllmModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeployVllmModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployVllmModelInput) GetRole() *string {
	return s.Role
}

func (s *DeployVllmModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployVllmModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployVllmModelInput) GetVpcConfig() *DeployVllmModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployVllmModelInput) SetAccountID(v string) *DeployVllmModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployVllmModelInput) SetConcurrencyConfig(v *DeployVllmModelInputConcurrencyConfig) *DeployVllmModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployVllmModelInput) SetCpu(v float32) *DeployVllmModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployVllmModelInput) SetCustomContainerConfig(v *DeployVllmModelInputCustomContainerConfig) *DeployVllmModelInput {
	s.CustomContainerConfig = v
	return s
}

func (s *DeployVllmModelInput) SetDescription(v string) *DeployVllmModelInput {
	s.Description = &v
	return s
}

func (s *DeployVllmModelInput) SetDiskSize(v int32) *DeployVllmModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployVllmModelInput) SetEnvName(v string) *DeployVllmModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployVllmModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployVllmModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployVllmModelInput) SetFeatureGates(v *DeployVllmModelInputFeatureGates) *DeployVllmModelInput {
	s.FeatureGates = v
	return s
}

func (s *DeployVllmModelInput) SetGpuConfig(v *DeployVllmModelInputGpuConfig) *DeployVllmModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployVllmModelInput) SetHttpTrigger(v *DeployVllmModelInputHttpTrigger) *DeployVllmModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployVllmModelInput) SetImageName(v string) *DeployVllmModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployVllmModelInput) SetInstanceConcurrency(v int32) *DeployVllmModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployVllmModelInput) SetLogConfig(v *DeployVllmModelInputLogConfig) *DeployVllmModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployVllmModelInput) SetMemorySize(v int32) *DeployVllmModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployVllmModelInput) SetModelConfig(v *DeployVllmModelInputModelConfig) *DeployVllmModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployVllmModelInput) SetName(v string) *DeployVllmModelInput {
	s.Name = &v
	return s
}

func (s *DeployVllmModelInput) SetNasConfig(v *DeployVllmModelInputNasConfig) *DeployVllmModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployVllmModelInput) SetOriginalName(v string) *DeployVllmModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployVllmModelInput) SetOssMountConfig(v *DeployVllmModelInputOssMountConfig) *DeployVllmModelInput {
	s.OssMountConfig = v
	return s
}

func (s *DeployVllmModelInput) SetProjectName(v string) *DeployVllmModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployVllmModelInput) SetProvisionConfig(v *DeployVllmModelInputProvisionConfig) *DeployVllmModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployVllmModelInput) SetRegion(v string) *DeployVllmModelInput {
	s.Region = &v
	return s
}

func (s *DeployVllmModelInput) SetReportStatusURL(v string) *DeployVllmModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployVllmModelInput) SetRole(v string) *DeployVllmModelInput {
	s.Role = &v
	return s
}

func (s *DeployVllmModelInput) SetTimeout(v int32) *DeployVllmModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployVllmModelInput) SetTraceId(v string) *DeployVllmModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployVllmModelInput) SetVpcConfig(v *DeployVllmModelInputVpcConfig) *DeployVllmModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeployVllmModelInput) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployVllmModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployVllmModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployVllmModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployVllmModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputCustomContainerConfig struct {
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s DeployVllmModelInputCustomContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputCustomContainerConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputCustomContainerConfig) GetRole() *string {
	return s.Role
}

func (s *DeployVllmModelInputCustomContainerConfig) SetRole(v string) *DeployVllmModelInputCustomContainerConfig {
	s.Role = &v
	return s
}

func (s *DeployVllmModelInputCustomContainerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputFeatureGates struct {
	AsyncProvisionCheck               *bool `json:"asyncProvisionCheck,omitempty" xml:"asyncProvisionCheck,omitempty"`
	DisableRollbackOnProvisionFailure *bool `json:"disableRollbackOnProvisionFailure,omitempty" xml:"disableRollbackOnProvisionFailure,omitempty"`
}

func (s DeployVllmModelInputFeatureGates) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputFeatureGates) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputFeatureGates) GetAsyncProvisionCheck() *bool {
	return s.AsyncProvisionCheck
}

func (s *DeployVllmModelInputFeatureGates) GetDisableRollbackOnProvisionFailure() *bool {
	return s.DisableRollbackOnProvisionFailure
}

func (s *DeployVllmModelInputFeatureGates) SetAsyncProvisionCheck(v bool) *DeployVllmModelInputFeatureGates {
	s.AsyncProvisionCheck = &v
	return s
}

func (s *DeployVllmModelInputFeatureGates) SetDisableRollbackOnProvisionFailure(v bool) *DeployVllmModelInputFeatureGates {
	s.DisableRollbackOnProvisionFailure = &v
	return s
}

func (s *DeployVllmModelInputFeatureGates) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployVllmModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeployVllmModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployVllmModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployVllmModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployVllmModelInputGpuConfig) SetGpuType(v string) *DeployVllmModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployVllmModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputHttpTrigger struct {
	Qualifier     *string                                       `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployVllmModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployVllmModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployVllmModelInputHttpTrigger) GetTriggerConfig() *DeployVllmModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployVllmModelInputHttpTrigger) SetQualifier(v string) *DeployVllmModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployVllmModelInputHttpTrigger) SetTriggerConfig(v *DeployVllmModelInputHttpTriggerTriggerConfig) *DeployVllmModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployVllmModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployVllmModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployVllmModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployVllmModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployVllmModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployVllmModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployVllmModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployVllmModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployVllmModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployVllmModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployVllmModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployVllmModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetLogBeginRule(v string) *DeployVllmModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetLogstore(v string) *DeployVllmModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) SetProject(v string) *DeployVllmModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployVllmModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputModelConfig struct {
	FmkVllmConfig *DeployVllmModelInputModelConfigFmkVllmConfig `json:"fmkVllmConfig,omitempty" xml:"fmkVllmConfig,omitempty" type:"Struct"`
	Framework     *string                                       `json:"framework,omitempty" xml:"framework,omitempty"`
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

func (s DeployVllmModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputModelConfig) GetFmkVllmConfig() *DeployVllmModelInputModelConfigFmkVllmConfig {
	return s.FmkVllmConfig
}

func (s *DeployVllmModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployVllmModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployVllmModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployVllmModelInputModelConfig) GetSkipDownload() *bool {
	return s.SkipDownload
}

func (s *DeployVllmModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployVllmModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployVllmModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployVllmModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployVllmModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployVllmModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployVllmModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployVllmModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployVllmModelInputModelConfig) GetWithPPU() *bool {
	return s.WithPPU
}

func (s *DeployVllmModelInputModelConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DeployVllmModelInputModelConfig) SetFmkVllmConfig(v *DeployVllmModelInputModelConfigFmkVllmConfig) *DeployVllmModelInputModelConfig {
	s.FmkVllmConfig = v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetFramework(v string) *DeployVllmModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployVllmModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetPrefix(v string) *DeployVllmModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSkipDownload(v bool) *DeployVllmModelInputModelConfig {
	s.SkipDownload = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSourceType(v string) *DeployVllmModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployVllmModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssBucket(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssPath(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSrcOssRegion(v string) *DeployVllmModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetSyncStrategy(v string) *DeployVllmModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetWithPPU(v bool) *DeployVllmModelInputModelConfig {
	s.WithPPU = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) SetWorkingDir(v string) *DeployVllmModelInputModelConfig {
	s.WorkingDir = &v
	return s
}

func (s *DeployVllmModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputModelConfigFmkVllmConfig struct {
	ApiKey                    *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	BlockSize                 *int32   `json:"blockSize,omitempty" xml:"blockSize,omitempty"`
	ChatTemplate              *string  `json:"chatTemplate,omitempty" xml:"chatTemplate,omitempty"`
	Dtype                     *string  `json:"dtype,omitempty" xml:"dtype,omitempty"`
	FullTextPostfix           *string  `json:"fullTextPostfix,omitempty" xml:"fullTextPostfix,omitempty"`
	GpuMemoryUtilization      *float32 `json:"gpuMemoryUtilization,omitempty" xml:"gpuMemoryUtilization,omitempty"`
	LoadFormat                *string  `json:"loadFormat,omitempty" xml:"loadFormat,omitempty"`
	MaxModelLen               *int32   `json:"maxModelLen,omitempty" xml:"maxModelLen,omitempty"`
	MaxParallelLoadingWorkers *int32   `json:"maxParallelLoadingWorkers,omitempty" xml:"maxParallelLoadingWorkers,omitempty"`
	Quantization              *string  `json:"quantization,omitempty" xml:"quantization,omitempty"`
	ServedModelName           *string  `json:"servedModelName,omitempty" xml:"servedModelName,omitempty"`
	SwapSpace                 *int32   `json:"swapSpace,omitempty" xml:"swapSpace,omitempty"`
}

func (s DeployVllmModelInputModelConfigFmkVllmConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputModelConfigFmkVllmConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetBlockSize() *int32 {
	return s.BlockSize
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetChatTemplate() *string {
	return s.ChatTemplate
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetDtype() *string {
	return s.Dtype
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetFullTextPostfix() *string {
	return s.FullTextPostfix
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetGpuMemoryUtilization() *float32 {
	return s.GpuMemoryUtilization
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetMaxModelLen() *int32 {
	return s.MaxModelLen
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetMaxParallelLoadingWorkers() *int32 {
	return s.MaxParallelLoadingWorkers
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetQuantization() *string {
	return s.Quantization
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetServedModelName() *string {
	return s.ServedModelName
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) GetSwapSpace() *int32 {
	return s.SwapSpace
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetApiKey(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ApiKey = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetBlockSize(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.BlockSize = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetChatTemplate(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ChatTemplate = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetDtype(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.Dtype = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetFullTextPostfix(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.FullTextPostfix = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetGpuMemoryUtilization(v float32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.GpuMemoryUtilization = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetLoadFormat(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.LoadFormat = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetMaxModelLen(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.MaxModelLen = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetMaxParallelLoadingWorkers(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.MaxParallelLoadingWorkers = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetQuantization(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.Quantization = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetServedModelName(v string) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.ServedModelName = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) SetSwapSpace(v int32) *DeployVllmModelInputModelConfigFmkVllmConfig {
	s.SwapSpace = &v
	return s
}

func (s *DeployVllmModelInputModelConfigFmkVllmConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputNasConfig struct {
	GroupId     *int32                                      `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*DeployVllmModelInputNasConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                                      `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployVllmModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeployVllmModelInputNasConfig) GetMountPoints() []*DeployVllmModelInputNasConfigMountPoints {
	return s.MountPoints
}

func (s *DeployVllmModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeployVllmModelInputNasConfig) SetGroupId(v int32) *DeployVllmModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployVllmModelInputNasConfig) SetMountPoints(v []*DeployVllmModelInputNasConfigMountPoints) *DeployVllmModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployVllmModelInputNasConfig) SetUserId(v int32) *DeployVllmModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployVllmModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputNasConfigMountPoints struct {
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s DeployVllmModelInputNasConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputNasConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputNasConfigMountPoints) GetEnableTLS() *bool {
	return s.EnableTLS
}

func (s *DeployVllmModelInputNasConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeployVllmModelInputNasConfigMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetEnableTLS(v bool) *DeployVllmModelInputNasConfigMountPoints {
	s.EnableTLS = &v
	return s
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetMountDir(v string) *DeployVllmModelInputNasConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployVllmModelInputNasConfigMountPoints) SetServerAddr(v string) *DeployVllmModelInputNasConfigMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *DeployVllmModelInputNasConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputOssMountConfig struct {
	MountPoints []*DeployVllmModelInputOssMountConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s DeployVllmModelInputOssMountConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputOssMountConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputOssMountConfig) GetMountPoints() []*DeployVllmModelInputOssMountConfigMountPoints {
	return s.MountPoints
}

func (s *DeployVllmModelInputOssMountConfig) SetMountPoints(v []*DeployVllmModelInputOssMountConfigMountPoints) *DeployVllmModelInputOssMountConfig {
	s.MountPoints = v
	return s
}

func (s *DeployVllmModelInputOssMountConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputOssMountConfigMountPoints struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DeployVllmModelInputOssMountConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputOssMountConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) SetBucketName(v string) *DeployVllmModelInputOssMountConfigMountPoints {
	s.BucketName = &v
	return s
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) SetBucketPath(v string) *DeployVllmModelInputOssMountConfigMountPoints {
	s.BucketPath = &v
	return s
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) SetEndpoint(v string) *DeployVllmModelInputOssMountConfigMountPoints {
	s.Endpoint = &v
	return s
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) SetMountDir(v string) *DeployVllmModelInputOssMountConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) SetReadOnly(v bool) *DeployVllmModelInputOssMountConfigMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *DeployVllmModelInputOssMountConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                  `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployVllmModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                 `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployVllmModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployVllmModelInputProvisionConfig) GetScheduledActions() []*DeployVllmModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployVllmModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeployVllmModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployVllmModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfig) SetScheduledActions(v []*DeployVllmModelInputProvisionConfigScheduledActions) *DeployVllmModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployVllmModelInputProvisionConfig) SetTarget(v int32) *DeployVllmModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployVllmModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetName(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployVllmModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployVllmModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployVllmModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployVllmModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployVllmModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployVllmModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployVllmModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployVllmModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployVllmModelInputVpcConfig) SetSecurityGroupId(v string) *DeployVllmModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployVllmModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployVllmModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployVllmModelInputVpcConfig) SetVpcId(v string) *DeployVllmModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployVllmModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
