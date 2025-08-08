// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeploySGLangModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeploySGLangModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeploySGLangModelInputConcurrencyConfig) *DeploySGLangModelInput
	GetConcurrencyConfig() *DeploySGLangModelInputConcurrencyConfig
	SetCpu(v float32) *DeploySGLangModelInput
	GetCpu() *float32
	SetCustomContainerConfig(v *DeploySGLangModelInputCustomContainerConfig) *DeploySGLangModelInput
	GetCustomContainerConfig() *DeploySGLangModelInputCustomContainerConfig
	SetDescription(v string) *DeploySGLangModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeploySGLangModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeploySGLangModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeploySGLangModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetFeatureGates(v *DeploySGLangModelInputFeatureGates) *DeploySGLangModelInput
	GetFeatureGates() *DeploySGLangModelInputFeatureGates
	SetGpuConfig(v *DeploySGLangModelInputGpuConfig) *DeploySGLangModelInput
	GetGpuConfig() *DeploySGLangModelInputGpuConfig
	SetHttpTrigger(v *DeploySGLangModelInputHttpTrigger) *DeploySGLangModelInput
	GetHttpTrigger() *DeploySGLangModelInputHttpTrigger
	SetImageName(v string) *DeploySGLangModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeploySGLangModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeploySGLangModelInputLogConfig) *DeploySGLangModelInput
	GetLogConfig() *DeploySGLangModelInputLogConfig
	SetMemorySize(v int32) *DeploySGLangModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeploySGLangModelInputModelConfig) *DeploySGLangModelInput
	GetModelConfig() *DeploySGLangModelInputModelConfig
	SetName(v string) *DeploySGLangModelInput
	GetName() *string
	SetNasConfig(v *DeploySGLangModelInputNasConfig) *DeploySGLangModelInput
	GetNasConfig() *DeploySGLangModelInputNasConfig
	SetOriginalName(v string) *DeploySGLangModelInput
	GetOriginalName() *string
	SetOssMountConfig(v *DeploySGLangModelInputOssMountConfig) *DeploySGLangModelInput
	GetOssMountConfig() *DeploySGLangModelInputOssMountConfig
	SetProjectName(v string) *DeploySGLangModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeploySGLangModelInputProvisionConfig) *DeploySGLangModelInput
	GetProvisionConfig() *DeploySGLangModelInputProvisionConfig
	SetRegion(v string) *DeploySGLangModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeploySGLangModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeploySGLangModelInput
	GetRole() *string
	SetTimeout(v int32) *DeploySGLangModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeploySGLangModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeploySGLangModelInputVpcConfig) *DeploySGLangModelInput
	GetVpcConfig() *DeploySGLangModelInputVpcConfig
}

type DeploySGLangModelInput struct {
	AccountID             *string                                      `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig     *DeploySGLangModelInputConcurrencyConfig     `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                   *float32                                     `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *DeploySGLangModelInputCustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty" type:"Struct"`
	Description           *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize              *int32                                       `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName               *string                                      `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables  map[string]interface{}                       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	FeatureGates          *DeploySGLangModelInputFeatureGates          `json:"featureGates,omitempty" xml:"featureGates,omitempty" type:"Struct"`
	GpuConfig             *DeploySGLangModelInputGpuConfig             `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger           *DeploySGLangModelInputHttpTrigger           `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName             *string                                      `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency   *int32                                       `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig             *DeploySGLangModelInputLogConfig             `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize            *int32                                       `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig           *DeploySGLangModelInputModelConfig           `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeploySGLangModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                `json:"originalName,omitempty" xml:"originalName,omitempty"`
	OssMountConfig  *DeploySGLangModelInputOssMountConfig  `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty" type:"Struct"`
	ProjectName     *string                                `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeploySGLangModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                          `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                           `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeploySGLangModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeploySGLangModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInput) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeploySGLangModelInput) GetConcurrencyConfig() *DeploySGLangModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeploySGLangModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeploySGLangModelInput) GetCustomContainerConfig() *DeploySGLangModelInputCustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *DeploySGLangModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeploySGLangModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeploySGLangModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeploySGLangModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeploySGLangModelInput) GetFeatureGates() *DeploySGLangModelInputFeatureGates {
	return s.FeatureGates
}

func (s *DeploySGLangModelInput) GetGpuConfig() *DeploySGLangModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeploySGLangModelInput) GetHttpTrigger() *DeploySGLangModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeploySGLangModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeploySGLangModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeploySGLangModelInput) GetLogConfig() *DeploySGLangModelInputLogConfig {
	return s.LogConfig
}

func (s *DeploySGLangModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeploySGLangModelInput) GetModelConfig() *DeploySGLangModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeploySGLangModelInput) GetName() *string {
	return s.Name
}

func (s *DeploySGLangModelInput) GetNasConfig() *DeploySGLangModelInputNasConfig {
	return s.NasConfig
}

func (s *DeploySGLangModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeploySGLangModelInput) GetOssMountConfig() *DeploySGLangModelInputOssMountConfig {
	return s.OssMountConfig
}

func (s *DeploySGLangModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeploySGLangModelInput) GetProvisionConfig() *DeploySGLangModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeploySGLangModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeploySGLangModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeploySGLangModelInput) GetRole() *string {
	return s.Role
}

func (s *DeploySGLangModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeploySGLangModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeploySGLangModelInput) GetVpcConfig() *DeploySGLangModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeploySGLangModelInput) SetAccountID(v string) *DeploySGLangModelInput {
	s.AccountID = &v
	return s
}

func (s *DeploySGLangModelInput) SetConcurrencyConfig(v *DeploySGLangModelInputConcurrencyConfig) *DeploySGLangModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetCpu(v float32) *DeploySGLangModelInput {
	s.Cpu = &v
	return s
}

func (s *DeploySGLangModelInput) SetCustomContainerConfig(v *DeploySGLangModelInputCustomContainerConfig) *DeploySGLangModelInput {
	s.CustomContainerConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetDescription(v string) *DeploySGLangModelInput {
	s.Description = &v
	return s
}

func (s *DeploySGLangModelInput) SetDiskSize(v int32) *DeploySGLangModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeploySGLangModelInput) SetEnvName(v string) *DeploySGLangModelInput {
	s.EnvName = &v
	return s
}

func (s *DeploySGLangModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeploySGLangModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeploySGLangModelInput) SetFeatureGates(v *DeploySGLangModelInputFeatureGates) *DeploySGLangModelInput {
	s.FeatureGates = v
	return s
}

func (s *DeploySGLangModelInput) SetGpuConfig(v *DeploySGLangModelInputGpuConfig) *DeploySGLangModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetHttpTrigger(v *DeploySGLangModelInputHttpTrigger) *DeploySGLangModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeploySGLangModelInput) SetImageName(v string) *DeploySGLangModelInput {
	s.ImageName = &v
	return s
}

func (s *DeploySGLangModelInput) SetInstanceConcurrency(v int32) *DeploySGLangModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeploySGLangModelInput) SetLogConfig(v *DeploySGLangModelInputLogConfig) *DeploySGLangModelInput {
	s.LogConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetMemorySize(v int32) *DeploySGLangModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeploySGLangModelInput) SetModelConfig(v *DeploySGLangModelInputModelConfig) *DeploySGLangModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetName(v string) *DeploySGLangModelInput {
	s.Name = &v
	return s
}

func (s *DeploySGLangModelInput) SetNasConfig(v *DeploySGLangModelInputNasConfig) *DeploySGLangModelInput {
	s.NasConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetOriginalName(v string) *DeploySGLangModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeploySGLangModelInput) SetOssMountConfig(v *DeploySGLangModelInputOssMountConfig) *DeploySGLangModelInput {
	s.OssMountConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetProjectName(v string) *DeploySGLangModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeploySGLangModelInput) SetProvisionConfig(v *DeploySGLangModelInputProvisionConfig) *DeploySGLangModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeploySGLangModelInput) SetRegion(v string) *DeploySGLangModelInput {
	s.Region = &v
	return s
}

func (s *DeploySGLangModelInput) SetReportStatusURL(v string) *DeploySGLangModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeploySGLangModelInput) SetRole(v string) *DeploySGLangModelInput {
	s.Role = &v
	return s
}

func (s *DeploySGLangModelInput) SetTimeout(v int32) *DeploySGLangModelInput {
	s.Timeout = &v
	return s
}

func (s *DeploySGLangModelInput) SetTraceId(v string) *DeploySGLangModelInput {
	s.TraceId = &v
	return s
}

func (s *DeploySGLangModelInput) SetVpcConfig(v *DeploySGLangModelInputVpcConfig) *DeploySGLangModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeploySGLangModelInput) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeploySGLangModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeploySGLangModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeploySGLangModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeploySGLangModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputCustomContainerConfig struct {
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s DeploySGLangModelInputCustomContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputCustomContainerConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputCustomContainerConfig) GetRole() *string {
	return s.Role
}

func (s *DeploySGLangModelInputCustomContainerConfig) SetRole(v string) *DeploySGLangModelInputCustomContainerConfig {
	s.Role = &v
	return s
}

func (s *DeploySGLangModelInputCustomContainerConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputFeatureGates struct {
	AsyncProvisionCheck               *bool `json:"asyncProvisionCheck,omitempty" xml:"asyncProvisionCheck,omitempty"`
	DisableRollbackOnProvisionFailure *bool `json:"disableRollbackOnProvisionFailure,omitempty" xml:"disableRollbackOnProvisionFailure,omitempty"`
}

func (s DeploySGLangModelInputFeatureGates) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputFeatureGates) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputFeatureGates) GetAsyncProvisionCheck() *bool {
	return s.AsyncProvisionCheck
}

func (s *DeploySGLangModelInputFeatureGates) GetDisableRollbackOnProvisionFailure() *bool {
	return s.DisableRollbackOnProvisionFailure
}

func (s *DeploySGLangModelInputFeatureGates) SetAsyncProvisionCheck(v bool) *DeploySGLangModelInputFeatureGates {
	s.AsyncProvisionCheck = &v
	return s
}

func (s *DeploySGLangModelInputFeatureGates) SetDisableRollbackOnProvisionFailure(v bool) *DeploySGLangModelInputFeatureGates {
	s.DisableRollbackOnProvisionFailure = &v
	return s
}

func (s *DeploySGLangModelInputFeatureGates) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeploySGLangModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeploySGLangModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeploySGLangModelInputGpuConfig) SetGpuMemorySize(v int32) *DeploySGLangModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeploySGLangModelInputGpuConfig) SetGpuType(v string) *DeploySGLangModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeploySGLangModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputHttpTrigger struct {
	Qualifier     *string                                         `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeploySGLangModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeploySGLangModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeploySGLangModelInputHttpTrigger) GetTriggerConfig() *DeploySGLangModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeploySGLangModelInputHttpTrigger) SetQualifier(v string) *DeploySGLangModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeploySGLangModelInputHttpTrigger) SetTriggerConfig(v *DeploySGLangModelInputHttpTriggerTriggerConfig) *DeploySGLangModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeploySGLangModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeploySGLangModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeploySGLangModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeploySGLangModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeploySGLangModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeploySGLangModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeploySGLangModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeploySGLangModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeploySGLangModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeploySGLangModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeploySGLangModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeploySGLangModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeploySGLangModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeploySGLangModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeploySGLangModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeploySGLangModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeploySGLangModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeploySGLangModelInputLogConfig) SetLogBeginRule(v string) *DeploySGLangModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeploySGLangModelInputLogConfig) SetLogstore(v string) *DeploySGLangModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeploySGLangModelInputLogConfig) SetProject(v string) *DeploySGLangModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeploySGLangModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputModelConfig struct {
	FmkSGLangConfig *DeploySGLangModelInputModelConfigFmkSGLangConfig `json:"fmkSGLangConfig,omitempty" xml:"fmkSGLangConfig,omitempty" type:"Struct"`
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

func (s DeploySGLangModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputModelConfig) GetFmkSGLangConfig() *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	return s.FmkSGLangConfig
}

func (s *DeploySGLangModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeploySGLangModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeploySGLangModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeploySGLangModelInputModelConfig) GetSkipDownload() *bool {
	return s.SkipDownload
}

func (s *DeploySGLangModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeploySGLangModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeploySGLangModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeploySGLangModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeploySGLangModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeploySGLangModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeploySGLangModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeploySGLangModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeploySGLangModelInputModelConfig) GetWithPPU() *bool {
	return s.WithPPU
}

func (s *DeploySGLangModelInputModelConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DeploySGLangModelInputModelConfig) SetFmkSGLangConfig(v *DeploySGLangModelInputModelConfigFmkSGLangConfig) *DeploySGLangModelInputModelConfig {
	s.FmkSGLangConfig = v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetFramework(v string) *DeploySGLangModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeploySGLangModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetPrefix(v string) *DeploySGLangModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSkipDownload(v bool) *DeploySGLangModelInputModelConfig {
	s.SkipDownload = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSourceType(v string) *DeploySGLangModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeploySGLangModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeploySGLangModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcModelScopeToken(v string) *DeploySGLangModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcOssBucket(v string) *DeploySGLangModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcOssPath(v string) *DeploySGLangModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSrcOssRegion(v string) *DeploySGLangModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetSyncStrategy(v string) *DeploySGLangModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetWithPPU(v bool) *DeploySGLangModelInputModelConfig {
	s.WithPPU = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) SetWorkingDir(v string) *DeploySGLangModelInputModelConfig {
	s.WorkingDir = &v
	return s
}

func (s *DeploySGLangModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputModelConfigFmkSGLangConfig struct {
	ApiKey             *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	ChatTemplate       *string  `json:"chatTemplate,omitempty" xml:"chatTemplate,omitempty"`
	Dtype              *string  `json:"dtype,omitempty" xml:"dtype,omitempty"`
	FullTextPostfix    *string  `json:"fullTextPostfix,omitempty" xml:"fullTextPostfix,omitempty"`
	LoadFormat         *string  `json:"loadFormat,omitempty" xml:"loadFormat,omitempty"`
	MaxRunningRequests *int32   `json:"maxRunningRequests,omitempty" xml:"maxRunningRequests,omitempty"`
	MaxTotalTokens     *int32   `json:"maxTotalTokens,omitempty" xml:"maxTotalTokens,omitempty"`
	MemFractionStatic  *float32 `json:"memFractionStatic,omitempty" xml:"memFractionStatic,omitempty"`
	Quantization       *string  `json:"quantization,omitempty" xml:"quantization,omitempty"`
	ServedModelName    *string  `json:"servedModelName,omitempty" xml:"servedModelName,omitempty"`
}

func (s DeploySGLangModelInputModelConfigFmkSGLangConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputModelConfigFmkSGLangConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetChatTemplate() *string {
	return s.ChatTemplate
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetDtype() *string {
	return s.Dtype
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetFullTextPostfix() *string {
	return s.FullTextPostfix
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetLoadFormat() *string {
	return s.LoadFormat
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetMaxRunningRequests() *int32 {
	return s.MaxRunningRequests
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetMaxTotalTokens() *int32 {
	return s.MaxTotalTokens
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetMemFractionStatic() *float32 {
	return s.MemFractionStatic
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetQuantization() *string {
	return s.Quantization
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) GetServedModelName() *string {
	return s.ServedModelName
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetApiKey(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.ApiKey = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetChatTemplate(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.ChatTemplate = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetDtype(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.Dtype = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetFullTextPostfix(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.FullTextPostfix = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetLoadFormat(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.LoadFormat = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetMaxRunningRequests(v int32) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.MaxRunningRequests = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetMaxTotalTokens(v int32) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.MaxTotalTokens = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetMemFractionStatic(v float32) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.MemFractionStatic = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetQuantization(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.Quantization = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) SetServedModelName(v string) *DeploySGLangModelInputModelConfigFmkSGLangConfig {
	s.ServedModelName = &v
	return s
}

func (s *DeploySGLangModelInputModelConfigFmkSGLangConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputNasConfig struct {
	GroupId     *int32                                        `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*DeploySGLangModelInputNasConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeploySGLangModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeploySGLangModelInputNasConfig) GetMountPoints() []*DeploySGLangModelInputNasConfigMountPoints {
	return s.MountPoints
}

func (s *DeploySGLangModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeploySGLangModelInputNasConfig) SetGroupId(v int32) *DeploySGLangModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeploySGLangModelInputNasConfig) SetMountPoints(v []*DeploySGLangModelInputNasConfigMountPoints) *DeploySGLangModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeploySGLangModelInputNasConfig) SetUserId(v int32) *DeploySGLangModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeploySGLangModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputNasConfigMountPoints struct {
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s DeploySGLangModelInputNasConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputNasConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputNasConfigMountPoints) GetEnableTLS() *bool {
	return s.EnableTLS
}

func (s *DeploySGLangModelInputNasConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeploySGLangModelInputNasConfigMountPoints) GetServerAddr() *string {
	return s.ServerAddr
}

func (s *DeploySGLangModelInputNasConfigMountPoints) SetEnableTLS(v bool) *DeploySGLangModelInputNasConfigMountPoints {
	s.EnableTLS = &v
	return s
}

func (s *DeploySGLangModelInputNasConfigMountPoints) SetMountDir(v string) *DeploySGLangModelInputNasConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeploySGLangModelInputNasConfigMountPoints) SetServerAddr(v string) *DeploySGLangModelInputNasConfigMountPoints {
	s.ServerAddr = &v
	return s
}

func (s *DeploySGLangModelInputNasConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputOssMountConfig struct {
	MountPoints []*DeploySGLangModelInputOssMountConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s DeploySGLangModelInputOssMountConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputOssMountConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputOssMountConfig) GetMountPoints() []*DeploySGLangModelInputOssMountConfigMountPoints {
	return s.MountPoints
}

func (s *DeploySGLangModelInputOssMountConfig) SetMountPoints(v []*DeploySGLangModelInputOssMountConfigMountPoints) *DeploySGLangModelInputOssMountConfig {
	s.MountPoints = v
	return s
}

func (s *DeploySGLangModelInputOssMountConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputOssMountConfigMountPoints struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DeploySGLangModelInputOssMountConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputOssMountConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) SetBucketName(v string) *DeploySGLangModelInputOssMountConfigMountPoints {
	s.BucketName = &v
	return s
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) SetBucketPath(v string) *DeploySGLangModelInputOssMountConfigMountPoints {
	s.BucketPath = &v
	return s
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) SetEndpoint(v string) *DeploySGLangModelInputOssMountConfigMountPoints {
	s.Endpoint = &v
	return s
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) SetMountDir(v string) *DeploySGLangModelInputOssMountConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) SetReadOnly(v bool) *DeploySGLangModelInputOssMountConfigMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *DeploySGLangModelInputOssMountConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                    `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeploySGLangModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                   `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeploySGLangModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeploySGLangModelInputProvisionConfig) GetScheduledActions() []*DeploySGLangModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeploySGLangModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeploySGLangModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeploySGLangModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfig) SetScheduledActions(v []*DeploySGLangModelInputProvisionConfigScheduledActions) *DeploySGLangModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeploySGLangModelInputProvisionConfig) SetTarget(v int32) *DeploySGLangModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeploySGLangModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetName(v string) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeploySGLangModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeploySGLangModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeploySGLangModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeploySGLangModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploySGLangModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeploySGLangModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeploySGLangModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeploySGLangModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeploySGLangModelInputVpcConfig) SetSecurityGroupId(v string) *DeploySGLangModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeploySGLangModelInputVpcConfig) SetVSwitchIds(v []*string) *DeploySGLangModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeploySGLangModelInputVpcConfig) SetVpcId(v string) *DeploySGLangModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeploySGLangModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
