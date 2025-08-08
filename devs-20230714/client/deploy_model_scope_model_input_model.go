// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployModelScopeModelInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployModelScopeModelInput
	GetAccountID() *string
	SetConcurrencyConfig(v *DeployModelScopeModelInputConcurrencyConfig) *DeployModelScopeModelInput
	GetConcurrencyConfig() *DeployModelScopeModelInputConcurrencyConfig
	SetCpu(v float32) *DeployModelScopeModelInput
	GetCpu() *float32
	SetDescription(v string) *DeployModelScopeModelInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployModelScopeModelInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployModelScopeModelInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployModelScopeModelInput
	GetEnvironmentVariables() map[string]interface{}
	SetGpuConfig(v *DeployModelScopeModelInputGpuConfig) *DeployModelScopeModelInput
	GetGpuConfig() *DeployModelScopeModelInputGpuConfig
	SetHttpTrigger(v *DeployModelScopeModelInputHttpTrigger) *DeployModelScopeModelInput
	GetHttpTrigger() *DeployModelScopeModelInputHttpTrigger
	SetImageName(v string) *DeployModelScopeModelInput
	GetImageName() *string
	SetInstanceConcurrency(v int32) *DeployModelScopeModelInput
	GetInstanceConcurrency() *int32
	SetLogConfig(v *DeployModelScopeModelInputLogConfig) *DeployModelScopeModelInput
	GetLogConfig() *DeployModelScopeModelInputLogConfig
	SetMemorySize(v int32) *DeployModelScopeModelInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployModelScopeModelInputModelConfig) *DeployModelScopeModelInput
	GetModelConfig() *DeployModelScopeModelInputModelConfig
	SetName(v string) *DeployModelScopeModelInput
	GetName() *string
	SetNasConfig(v *DeployModelScopeModelInputNasConfig) *DeployModelScopeModelInput
	GetNasConfig() *DeployModelScopeModelInputNasConfig
	SetOriginalName(v string) *DeployModelScopeModelInput
	GetOriginalName() *string
	SetProjectName(v string) *DeployModelScopeModelInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployModelScopeModelInputProvisionConfig) *DeployModelScopeModelInput
	GetProvisionConfig() *DeployModelScopeModelInputProvisionConfig
	SetRegion(v string) *DeployModelScopeModelInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployModelScopeModelInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployModelScopeModelInput
	GetRole() *string
	SetTimeout(v int32) *DeployModelScopeModelInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployModelScopeModelInput
	GetTraceId() *string
	SetVpcConfig(v *DeployModelScopeModelInputVpcConfig) *DeployModelScopeModelInput
	GetVpcConfig() *DeployModelScopeModelInputVpcConfig
}

type DeployModelScopeModelInput struct {
	AccountID            *string                                      `json:"accountID,omitempty" xml:"accountID,omitempty"`
	ConcurrencyConfig    *DeployModelScopeModelInputConcurrencyConfig `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                  *float32                                     `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Description          *string                                      `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize             *int32                                       `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName              *string                                      `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables map[string]interface{}                       `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	GpuConfig            *DeployModelScopeModelInputGpuConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger          *DeployModelScopeModelInputHttpTrigger       `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	ImageName            *string                                      `json:"imageName,omitempty" xml:"imageName,omitempty"`
	InstanceConcurrency  *int32                                       `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	LogConfig            *DeployModelScopeModelInputLogConfig         `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize           *int32                                       `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig          *DeployModelScopeModelInputModelConfig       `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployModelScopeModelInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                    `json:"originalName,omitempty" xml:"originalName,omitempty"`
	ProjectName     *string                                    `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployModelScopeModelInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                    `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                    `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                              `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                               `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployModelScopeModelInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployModelScopeModelInput) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInput) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployModelScopeModelInput) GetConcurrencyConfig() *DeployModelScopeModelInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployModelScopeModelInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployModelScopeModelInput) GetDescription() *string {
	return s.Description
}

func (s *DeployModelScopeModelInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployModelScopeModelInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployModelScopeModelInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployModelScopeModelInput) GetGpuConfig() *DeployModelScopeModelInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployModelScopeModelInput) GetHttpTrigger() *DeployModelScopeModelInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployModelScopeModelInput) GetImageName() *string {
	return s.ImageName
}

func (s *DeployModelScopeModelInput) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployModelScopeModelInput) GetLogConfig() *DeployModelScopeModelInputLogConfig {
	return s.LogConfig
}

func (s *DeployModelScopeModelInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployModelScopeModelInput) GetModelConfig() *DeployModelScopeModelInputModelConfig {
	return s.ModelConfig
}

func (s *DeployModelScopeModelInput) GetName() *string {
	return s.Name
}

func (s *DeployModelScopeModelInput) GetNasConfig() *DeployModelScopeModelInputNasConfig {
	return s.NasConfig
}

func (s *DeployModelScopeModelInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployModelScopeModelInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployModelScopeModelInput) GetProvisionConfig() *DeployModelScopeModelInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployModelScopeModelInput) GetRegion() *string {
	return s.Region
}

func (s *DeployModelScopeModelInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployModelScopeModelInput) GetRole() *string {
	return s.Role
}

func (s *DeployModelScopeModelInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployModelScopeModelInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployModelScopeModelInput) GetVpcConfig() *DeployModelScopeModelInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployModelScopeModelInput) SetAccountID(v string) *DeployModelScopeModelInput {
	s.AccountID = &v
	return s
}

func (s *DeployModelScopeModelInput) SetConcurrencyConfig(v *DeployModelScopeModelInputConcurrencyConfig) *DeployModelScopeModelInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetCpu(v float32) *DeployModelScopeModelInput {
	s.Cpu = &v
	return s
}

func (s *DeployModelScopeModelInput) SetDescription(v string) *DeployModelScopeModelInput {
	s.Description = &v
	return s
}

func (s *DeployModelScopeModelInput) SetDiskSize(v int32) *DeployModelScopeModelInput {
	s.DiskSize = &v
	return s
}

func (s *DeployModelScopeModelInput) SetEnvName(v string) *DeployModelScopeModelInput {
	s.EnvName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetEnvironmentVariables(v map[string]interface{}) *DeployModelScopeModelInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployModelScopeModelInput) SetGpuConfig(v *DeployModelScopeModelInputGpuConfig) *DeployModelScopeModelInput {
	s.GpuConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetHttpTrigger(v *DeployModelScopeModelInputHttpTrigger) *DeployModelScopeModelInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployModelScopeModelInput) SetImageName(v string) *DeployModelScopeModelInput {
	s.ImageName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetInstanceConcurrency(v int32) *DeployModelScopeModelInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployModelScopeModelInput) SetLogConfig(v *DeployModelScopeModelInputLogConfig) *DeployModelScopeModelInput {
	s.LogConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetMemorySize(v int32) *DeployModelScopeModelInput {
	s.MemorySize = &v
	return s
}

func (s *DeployModelScopeModelInput) SetModelConfig(v *DeployModelScopeModelInputModelConfig) *DeployModelScopeModelInput {
	s.ModelConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetName(v string) *DeployModelScopeModelInput {
	s.Name = &v
	return s
}

func (s *DeployModelScopeModelInput) SetNasConfig(v *DeployModelScopeModelInputNasConfig) *DeployModelScopeModelInput {
	s.NasConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetOriginalName(v string) *DeployModelScopeModelInput {
	s.OriginalName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetProjectName(v string) *DeployModelScopeModelInput {
	s.ProjectName = &v
	return s
}

func (s *DeployModelScopeModelInput) SetProvisionConfig(v *DeployModelScopeModelInputProvisionConfig) *DeployModelScopeModelInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployModelScopeModelInput) SetRegion(v string) *DeployModelScopeModelInput {
	s.Region = &v
	return s
}

func (s *DeployModelScopeModelInput) SetReportStatusURL(v string) *DeployModelScopeModelInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployModelScopeModelInput) SetRole(v string) *DeployModelScopeModelInput {
	s.Role = &v
	return s
}

func (s *DeployModelScopeModelInput) SetTimeout(v int32) *DeployModelScopeModelInput {
	s.Timeout = &v
	return s
}

func (s *DeployModelScopeModelInput) SetTraceId(v string) *DeployModelScopeModelInput {
	s.TraceId = &v
	return s
}

func (s *DeployModelScopeModelInput) SetVpcConfig(v *DeployModelScopeModelInputVpcConfig) *DeployModelScopeModelInput {
	s.VpcConfig = v
	return s
}

func (s *DeployModelScopeModelInput) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployModelScopeModelInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployModelScopeModelInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployModelScopeModelInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployModelScopeModelInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputGpuConfig struct {
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployModelScopeModelInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputGpuConfig) GetGpuMemorySize() *int32 {
	return s.GpuMemorySize
}

func (s *DeployModelScopeModelInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployModelScopeModelInputGpuConfig) SetGpuMemorySize(v int32) *DeployModelScopeModelInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployModelScopeModelInputGpuConfig) SetGpuType(v string) *DeployModelScopeModelInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployModelScopeModelInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputHttpTrigger struct {
	Qualifier     *string                                             `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployModelScopeModelInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployModelScopeModelInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployModelScopeModelInputHttpTrigger) GetTriggerConfig() *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployModelScopeModelInputHttpTrigger) SetQualifier(v string) *DeployModelScopeModelInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTrigger) SetTriggerConfig(v *DeployModelScopeModelInputHttpTriggerTriggerConfig) *DeployModelScopeModelInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployModelScopeModelInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployModelScopeModelInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployModelScopeModelInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployModelScopeModelInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployModelScopeModelInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployModelScopeModelInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployModelScopeModelInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployModelScopeModelInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployModelScopeModelInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployModelScopeModelInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployModelScopeModelInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetEnableRequestMetrics(v bool) *DeployModelScopeModelInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetLogBeginRule(v string) *DeployModelScopeModelInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetLogstore(v string) *DeployModelScopeModelInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) SetProject(v string) *DeployModelScopeModelInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployModelScopeModelInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputModelConfig struct {
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

func (s DeployModelScopeModelInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployModelScopeModelInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployModelScopeModelInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployModelScopeModelInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployModelScopeModelInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployModelScopeModelInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployModelScopeModelInputModelConfig) SetFramework(v string) *DeployModelScopeModelInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployModelScopeModelInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetPrefix(v string) *DeployModelScopeModelInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSourceType(v string) *DeployModelScopeModelInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeModelID(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcModelScopeToken(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssBucket(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssPath(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSrcOssRegion(v string) *DeployModelScopeModelInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) SetSyncStrategy(v string) *DeployModelScopeModelInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployModelScopeModelInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputNasConfig struct {
	GroupId     *int32    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int32    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployModelScopeModelInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputNasConfig) GetGroupId() *int32 {
	return s.GroupId
}

func (s *DeployModelScopeModelInputNasConfig) GetMountPoints() []*string {
	return s.MountPoints
}

func (s *DeployModelScopeModelInputNasConfig) GetUserId() *int32 {
	return s.UserId
}

func (s *DeployModelScopeModelInputNasConfig) SetGroupId(v int32) *DeployModelScopeModelInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployModelScopeModelInputNasConfig) SetMountPoints(v []*string) *DeployModelScopeModelInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployModelScopeModelInputNasConfig) SetUserId(v int32) *DeployModelScopeModelInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployModelScopeModelInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                        `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployModelScopeModelInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int32                                                       `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployModelScopeModelInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployModelScopeModelInputProvisionConfig) GetScheduledActions() []*DeployModelScopeModelInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployModelScopeModelInputProvisionConfig) GetTarget() *int32 {
	return s.Target
}

func (s *DeployModelScopeModelInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployModelScopeModelInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfig) SetScheduledActions(v []*DeployModelScopeModelInputProvisionConfigScheduledActions) *DeployModelScopeModelInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfig) SetTarget(v int32) *DeployModelScopeModelInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployModelScopeModelInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetName(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployModelScopeModelInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployModelScopeModelInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployModelScopeModelInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployModelScopeModelInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployModelScopeModelInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployModelScopeModelInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployModelScopeModelInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployModelScopeModelInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployModelScopeModelInputVpcConfig) SetSecurityGroupId(v string) *DeployModelScopeModelInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployModelScopeModelInputVpcConfig) SetVSwitchIds(v []*string) *DeployModelScopeModelInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployModelScopeModelInputVpcConfig) SetVpcId(v string) *DeployModelScopeModelInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployModelScopeModelInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
