// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployCustomContainerInput interface {
	dara.Model
	String() string
	GoString() string
	SetAccountID(v string) *DeployCustomContainerInput
	GetAccountID() *string
	SetAsyncInvokeConfig(v *DeployCustomContainerInputAsyncInvokeConfig) *DeployCustomContainerInput
	GetAsyncInvokeConfig() *DeployCustomContainerInputAsyncInvokeConfig
	SetConcurrencyConfig(v *DeployCustomContainerInputConcurrencyConfig) *DeployCustomContainerInput
	GetConcurrencyConfig() *DeployCustomContainerInputConcurrencyConfig
	SetCpu(v float32) *DeployCustomContainerInput
	GetCpu() *float32
	SetCustomContainerConfig(v *DeployCustomContainerInputCustomContainerConfig) *DeployCustomContainerInput
	GetCustomContainerConfig() *DeployCustomContainerInputCustomContainerConfig
	SetDescription(v string) *DeployCustomContainerInput
	GetDescription() *string
	SetDiskSize(v int32) *DeployCustomContainerInput
	GetDiskSize() *int32
	SetEnvName(v string) *DeployCustomContainerInput
	GetEnvName() *string
	SetEnvironmentVariables(v map[string]interface{}) *DeployCustomContainerInput
	GetEnvironmentVariables() map[string]interface{}
	SetFeatureGates(v *DeployCustomContainerInputFeatureGates) *DeployCustomContainerInput
	GetFeatureGates() *DeployCustomContainerInputFeatureGates
	SetGpuConfig(v *DeployCustomContainerInputGpuConfig) *DeployCustomContainerInput
	GetGpuConfig() *DeployCustomContainerInputGpuConfig
	SetHttpTrigger(v *DeployCustomContainerInputHttpTrigger) *DeployCustomContainerInput
	GetHttpTrigger() *DeployCustomContainerInputHttpTrigger
	SetLogConfig(v *DeployCustomContainerInputLogConfig) *DeployCustomContainerInput
	GetLogConfig() *DeployCustomContainerInputLogConfig
	SetMemorySize(v int32) *DeployCustomContainerInput
	GetMemorySize() *int32
	SetModelConfig(v *DeployCustomContainerInputModelConfig) *DeployCustomContainerInput
	GetModelConfig() *DeployCustomContainerInputModelConfig
	SetName(v string) *DeployCustomContainerInput
	GetName() *string
	SetNasConfig(v *DeployCustomContainerInputNasConfig) *DeployCustomContainerInput
	GetNasConfig() *DeployCustomContainerInputNasConfig
	SetOriginalName(v string) *DeployCustomContainerInput
	GetOriginalName() *string
	SetOssMountConfig(v *DeployCustomContainerInputOssMountConfig) *DeployCustomContainerInput
	GetOssMountConfig() *DeployCustomContainerInputOssMountConfig
	SetProjectName(v string) *DeployCustomContainerInput
	GetProjectName() *string
	SetProvisionConfig(v *DeployCustomContainerInputProvisionConfig) *DeployCustomContainerInput
	GetProvisionConfig() *DeployCustomContainerInputProvisionConfig
	SetRegion(v string) *DeployCustomContainerInput
	GetRegion() *string
	SetReportStatusURL(v string) *DeployCustomContainerInput
	GetReportStatusURL() *string
	SetRole(v string) *DeployCustomContainerInput
	GetRole() *string
	SetTimeout(v int32) *DeployCustomContainerInput
	GetTimeout() *int32
	SetTraceId(v string) *DeployCustomContainerInput
	GetTraceId() *string
	SetVpcConfig(v *DeployCustomContainerInputVpcConfig) *DeployCustomContainerInput
	GetVpcConfig() *DeployCustomContainerInputVpcConfig
}

type DeployCustomContainerInput struct {
	AccountID             *string                                          `json:"accountID,omitempty" xml:"accountID,omitempty"`
	AsyncInvokeConfig     *DeployCustomContainerInputAsyncInvokeConfig     `json:"asyncInvokeConfig,omitempty" xml:"asyncInvokeConfig,omitempty" type:"Struct"`
	ConcurrencyConfig     *DeployCustomContainerInputConcurrencyConfig     `json:"concurrencyConfig,omitempty" xml:"concurrencyConfig,omitempty" type:"Struct"`
	Cpu                   *float32                                         `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *DeployCustomContainerInputCustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty" type:"Struct"`
	Description           *string                                          `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize              *int32                                           `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvName               *string                                          `json:"envName,omitempty" xml:"envName,omitempty"`
	EnvironmentVariables  map[string]interface{}                           `json:"environmentVariables,omitempty" xml:"environmentVariables,omitempty"`
	FeatureGates          *DeployCustomContainerInputFeatureGates          `json:"featureGates,omitempty" xml:"featureGates,omitempty" type:"Struct"`
	GpuConfig             *DeployCustomContainerInputGpuConfig             `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty" type:"Struct"`
	HttpTrigger           *DeployCustomContainerInputHttpTrigger           `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty" type:"Struct"`
	LogConfig             *DeployCustomContainerInputLogConfig             `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	MemorySize            *int32                                           `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	ModelConfig           *DeployCustomContainerInputModelConfig           `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Name            *string                                    `json:"name,omitempty" xml:"name,omitempty"`
	NasConfig       *DeployCustomContainerInputNasConfig       `json:"nasConfig,omitempty" xml:"nasConfig,omitempty" type:"Struct"`
	OriginalName    *string                                    `json:"originalName,omitempty" xml:"originalName,omitempty"`
	OssMountConfig  *DeployCustomContainerInputOssMountConfig  `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty" type:"Struct"`
	ProjectName     *string                                    `json:"projectName,omitempty" xml:"projectName,omitempty"`
	ProvisionConfig *DeployCustomContainerInputProvisionConfig `json:"provisionConfig,omitempty" xml:"provisionConfig,omitempty" type:"Struct"`
	Region          *string                                    `json:"region,omitempty" xml:"region,omitempty"`
	ReportStatusURL *string                                    `json:"reportStatusURL,omitempty" xml:"reportStatusURL,omitempty"`
	// This parameter is required.
	Role      *string                              `json:"role,omitempty" xml:"role,omitempty"`
	Timeout   *int32                               `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
	VpcConfig *DeployCustomContainerInputVpcConfig `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInput) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInput) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInput) GetAccountID() *string {
	return s.AccountID
}

func (s *DeployCustomContainerInput) GetAsyncInvokeConfig() *DeployCustomContainerInputAsyncInvokeConfig {
	return s.AsyncInvokeConfig
}

func (s *DeployCustomContainerInput) GetConcurrencyConfig() *DeployCustomContainerInputConcurrencyConfig {
	return s.ConcurrencyConfig
}

func (s *DeployCustomContainerInput) GetCpu() *float32 {
	return s.Cpu
}

func (s *DeployCustomContainerInput) GetCustomContainerConfig() *DeployCustomContainerInputCustomContainerConfig {
	return s.CustomContainerConfig
}

func (s *DeployCustomContainerInput) GetDescription() *string {
	return s.Description
}

func (s *DeployCustomContainerInput) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DeployCustomContainerInput) GetEnvName() *string {
	return s.EnvName
}

func (s *DeployCustomContainerInput) GetEnvironmentVariables() map[string]interface{} {
	return s.EnvironmentVariables
}

func (s *DeployCustomContainerInput) GetFeatureGates() *DeployCustomContainerInputFeatureGates {
	return s.FeatureGates
}

func (s *DeployCustomContainerInput) GetGpuConfig() *DeployCustomContainerInputGpuConfig {
	return s.GpuConfig
}

func (s *DeployCustomContainerInput) GetHttpTrigger() *DeployCustomContainerInputHttpTrigger {
	return s.HttpTrigger
}

func (s *DeployCustomContainerInput) GetLogConfig() *DeployCustomContainerInputLogConfig {
	return s.LogConfig
}

func (s *DeployCustomContainerInput) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *DeployCustomContainerInput) GetModelConfig() *DeployCustomContainerInputModelConfig {
	return s.ModelConfig
}

func (s *DeployCustomContainerInput) GetName() *string {
	return s.Name
}

func (s *DeployCustomContainerInput) GetNasConfig() *DeployCustomContainerInputNasConfig {
	return s.NasConfig
}

func (s *DeployCustomContainerInput) GetOriginalName() *string {
	return s.OriginalName
}

func (s *DeployCustomContainerInput) GetOssMountConfig() *DeployCustomContainerInputOssMountConfig {
	return s.OssMountConfig
}

func (s *DeployCustomContainerInput) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeployCustomContainerInput) GetProvisionConfig() *DeployCustomContainerInputProvisionConfig {
	return s.ProvisionConfig
}

func (s *DeployCustomContainerInput) GetRegion() *string {
	return s.Region
}

func (s *DeployCustomContainerInput) GetReportStatusURL() *string {
	return s.ReportStatusURL
}

func (s *DeployCustomContainerInput) GetRole() *string {
	return s.Role
}

func (s *DeployCustomContainerInput) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployCustomContainerInput) GetTraceId() *string {
	return s.TraceId
}

func (s *DeployCustomContainerInput) GetVpcConfig() *DeployCustomContainerInputVpcConfig {
	return s.VpcConfig
}

func (s *DeployCustomContainerInput) SetAccountID(v string) *DeployCustomContainerInput {
	s.AccountID = &v
	return s
}

func (s *DeployCustomContainerInput) SetAsyncInvokeConfig(v *DeployCustomContainerInputAsyncInvokeConfig) *DeployCustomContainerInput {
	s.AsyncInvokeConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetConcurrencyConfig(v *DeployCustomContainerInputConcurrencyConfig) *DeployCustomContainerInput {
	s.ConcurrencyConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetCpu(v float32) *DeployCustomContainerInput {
	s.Cpu = &v
	return s
}

func (s *DeployCustomContainerInput) SetCustomContainerConfig(v *DeployCustomContainerInputCustomContainerConfig) *DeployCustomContainerInput {
	s.CustomContainerConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetDescription(v string) *DeployCustomContainerInput {
	s.Description = &v
	return s
}

func (s *DeployCustomContainerInput) SetDiskSize(v int32) *DeployCustomContainerInput {
	s.DiskSize = &v
	return s
}

func (s *DeployCustomContainerInput) SetEnvName(v string) *DeployCustomContainerInput {
	s.EnvName = &v
	return s
}

func (s *DeployCustomContainerInput) SetEnvironmentVariables(v map[string]interface{}) *DeployCustomContainerInput {
	s.EnvironmentVariables = v
	return s
}

func (s *DeployCustomContainerInput) SetFeatureGates(v *DeployCustomContainerInputFeatureGates) *DeployCustomContainerInput {
	s.FeatureGates = v
	return s
}

func (s *DeployCustomContainerInput) SetGpuConfig(v *DeployCustomContainerInputGpuConfig) *DeployCustomContainerInput {
	s.GpuConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetHttpTrigger(v *DeployCustomContainerInputHttpTrigger) *DeployCustomContainerInput {
	s.HttpTrigger = v
	return s
}

func (s *DeployCustomContainerInput) SetLogConfig(v *DeployCustomContainerInputLogConfig) *DeployCustomContainerInput {
	s.LogConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetMemorySize(v int32) *DeployCustomContainerInput {
	s.MemorySize = &v
	return s
}

func (s *DeployCustomContainerInput) SetModelConfig(v *DeployCustomContainerInputModelConfig) *DeployCustomContainerInput {
	s.ModelConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetName(v string) *DeployCustomContainerInput {
	s.Name = &v
	return s
}

func (s *DeployCustomContainerInput) SetNasConfig(v *DeployCustomContainerInputNasConfig) *DeployCustomContainerInput {
	s.NasConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetOriginalName(v string) *DeployCustomContainerInput {
	s.OriginalName = &v
	return s
}

func (s *DeployCustomContainerInput) SetOssMountConfig(v *DeployCustomContainerInputOssMountConfig) *DeployCustomContainerInput {
	s.OssMountConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetProjectName(v string) *DeployCustomContainerInput {
	s.ProjectName = &v
	return s
}

func (s *DeployCustomContainerInput) SetProvisionConfig(v *DeployCustomContainerInputProvisionConfig) *DeployCustomContainerInput {
	s.ProvisionConfig = v
	return s
}

func (s *DeployCustomContainerInput) SetRegion(v string) *DeployCustomContainerInput {
	s.Region = &v
	return s
}

func (s *DeployCustomContainerInput) SetReportStatusURL(v string) *DeployCustomContainerInput {
	s.ReportStatusURL = &v
	return s
}

func (s *DeployCustomContainerInput) SetRole(v string) *DeployCustomContainerInput {
	s.Role = &v
	return s
}

func (s *DeployCustomContainerInput) SetTimeout(v int32) *DeployCustomContainerInput {
	s.Timeout = &v
	return s
}

func (s *DeployCustomContainerInput) SetTraceId(v string) *DeployCustomContainerInput {
	s.TraceId = &v
	return s
}

func (s *DeployCustomContainerInput) SetVpcConfig(v *DeployCustomContainerInputVpcConfig) *DeployCustomContainerInput {
	s.VpcConfig = v
	return s
}

func (s *DeployCustomContainerInput) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputAsyncInvokeConfig struct {
	AsyncTask                 *bool                                                         `json:"asyncTask,omitempty" xml:"asyncTask,omitempty"`
	DestinationConfig         *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty" type:"Struct"`
	MaxAsyncEventAgeInSeconds *int64                                                        `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	MaxAsyncRetryAttempts     *int64                                                        `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s DeployCustomContainerInputAsyncInvokeConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputAsyncInvokeConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) GetAsyncTask() *bool {
	return s.AsyncTask
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) GetDestinationConfig() *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig {
	return s.DestinationConfig
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) GetMaxAsyncEventAgeInSeconds() *int64 {
	return s.MaxAsyncEventAgeInSeconds
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) GetMaxAsyncRetryAttempts() *int64 {
	return s.MaxAsyncRetryAttempts
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) SetAsyncTask(v bool) *DeployCustomContainerInputAsyncInvokeConfig {
	s.AsyncTask = &v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) SetDestinationConfig(v *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) *DeployCustomContainerInputAsyncInvokeConfig {
	s.DestinationConfig = v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) SetMaxAsyncEventAgeInSeconds(v int64) *DeployCustomContainerInputAsyncInvokeConfig {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) SetMaxAsyncRetryAttempts(v int64) *DeployCustomContainerInputAsyncInvokeConfig {
	s.MaxAsyncRetryAttempts = &v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputAsyncInvokeConfigDestinationConfig struct {
	OnFailure *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure `json:"onFailure,omitempty" xml:"onFailure,omitempty" type:"Struct"`
	OnSuccess *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess `json:"onSuccess,omitempty" xml:"onSuccess,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) GetOnFailure() *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure {
	return s.OnFailure
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) GetOnSuccess() *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess {
	return s.OnSuccess
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) SetOnFailure(v *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig {
	s.OnFailure = v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) SetOnSuccess(v *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig {
	s.OnSuccess = v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure struct {
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) GetDestination() *string {
	return s.Destination
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) SetDestination(v string) *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure {
	s.Destination = &v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnFailure) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess struct {
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) GetDestination() *string {
	return s.Destination
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) SetDestination(v string) *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess {
	s.Destination = &v
	return s
}

func (s *DeployCustomContainerInputAsyncInvokeConfigDestinationConfigOnSuccess) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputConcurrencyConfig struct {
	ReservedConcurrency *int32 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s DeployCustomContainerInputConcurrencyConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputConcurrencyConfig) GetReservedConcurrency() *int32 {
	return s.ReservedConcurrency
}

func (s *DeployCustomContainerInputConcurrencyConfig) SetReservedConcurrency(v int32) *DeployCustomContainerInputConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

func (s *DeployCustomContainerInputConcurrencyConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputCustomContainerConfig struct {
	Command                 []*string                                                               `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	Entrypoint              []*string                                                               `json:"entrypoint,omitempty" xml:"entrypoint,omitempty" type:"Repeated"`
	HealthCheckConfig       *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig       `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty" type:"Struct"`
	Image                   *string                                                                 `json:"image,omitempty" xml:"image,omitempty"`
	InstanceConcurrency     *int32                                                                  `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty" type:"Struct"`
	Port                    *int32                                                                  `json:"port,omitempty" xml:"port,omitempty"`
	Role                    *string                                                                 `json:"role,omitempty" xml:"role,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetCommand() []*string {
	return s.Command
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetEntrypoint() []*string {
	return s.Entrypoint
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetHealthCheckConfig() *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetImage() *string {
	return s.Image
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetInstanceConcurrency() *int32 {
	return s.InstanceConcurrency
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetInstanceLifecycleConfig() *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig {
	return s.InstanceLifecycleConfig
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetPort() *int32 {
	return s.Port
}

func (s *DeployCustomContainerInputCustomContainerConfig) GetRole() *string {
	return s.Role
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetCommand(v []*string) *DeployCustomContainerInputCustomContainerConfig {
	s.Command = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetEntrypoint(v []*string) *DeployCustomContainerInputCustomContainerConfig {
	s.Entrypoint = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetHealthCheckConfig(v *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) *DeployCustomContainerInputCustomContainerConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetImage(v string) *DeployCustomContainerInputCustomContainerConfig {
	s.Image = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetInstanceConcurrency(v int32) *DeployCustomContainerInputCustomContainerConfig {
	s.InstanceConcurrency = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetInstanceLifecycleConfig(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) *DeployCustomContainerInputCustomContainerConfig {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetPort(v int32) *DeployCustomContainerInputCustomContainerConfig {
	s.Port = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) SetRole(v string) *DeployCustomContainerInputCustomContainerConfig {
	s.Role = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputCustomContainerConfigHealthCheckConfig struct {
	FailureThreshold    *int32  `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	HttpGetUrl          *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	InitialDelaySeconds *int32  `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int32  `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	SuccessThreshold    *int32  `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	TimeoutSeconds      *int64  `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetHttpGetUrl() *string {
	return s.HttpGetUrl
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) GetTimeoutSeconds() *int64 {
	return s.TimeoutSeconds
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetFailureThreshold(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.FailureThreshold = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetHttpGetUrl(v string) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.HttpGetUrl = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetInitialDelaySeconds(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetPeriodSeconds(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.PeriodSeconds = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetSuccessThreshold(v int32) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.SuccessThreshold = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) SetTimeoutSeconds(v int64) *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig struct {
	Initializer *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer `json:"initializer,omitempty" xml:"initializer,omitempty" type:"Struct"`
	PreStop     *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop     `json:"preStop,omitempty" xml:"preStop,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) GetInitializer() *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer {
	return s.Initializer
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) GetPreStop() *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop {
	return s.PreStop
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) SetInitializer(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig {
	s.Initializer = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) SetPreStop(v *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig {
	s.PreStop = v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer struct {
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) GetHandler() *string {
	return s.Handler
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) SetHandler(v string) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer {
	s.Handler = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) SetTimeout(v int32) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer {
	s.Timeout = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigInitializer) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop struct {
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) GetHandler() *string {
	return s.Handler
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) SetHandler(v string) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop {
	s.Handler = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) SetTimeout(v int32) *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop {
	s.Timeout = &v
	return s
}

func (s *DeployCustomContainerInputCustomContainerConfigInstanceLifecycleConfigPreStop) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputFeatureGates struct {
	AsyncProvisionCheck               *bool `json:"asyncProvisionCheck,omitempty" xml:"asyncProvisionCheck,omitempty"`
	DisableRollbackOnProvisionFailure *bool `json:"disableRollbackOnProvisionFailure,omitempty" xml:"disableRollbackOnProvisionFailure,omitempty"`
}

func (s DeployCustomContainerInputFeatureGates) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputFeatureGates) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputFeatureGates) GetAsyncProvisionCheck() *bool {
	return s.AsyncProvisionCheck
}

func (s *DeployCustomContainerInputFeatureGates) GetDisableRollbackOnProvisionFailure() *bool {
	return s.DisableRollbackOnProvisionFailure
}

func (s *DeployCustomContainerInputFeatureGates) SetAsyncProvisionCheck(v bool) *DeployCustomContainerInputFeatureGates {
	s.AsyncProvisionCheck = &v
	return s
}

func (s *DeployCustomContainerInputFeatureGates) SetDisableRollbackOnProvisionFailure(v bool) *DeployCustomContainerInputFeatureGates {
	s.DisableRollbackOnProvisionFailure = &v
	return s
}

func (s *DeployCustomContainerInputFeatureGates) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputGpuConfig struct {
	GpuMemorySize *int64  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s DeployCustomContainerInputGpuConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputGpuConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputGpuConfig) GetGpuMemorySize() *int64 {
	return s.GpuMemorySize
}

func (s *DeployCustomContainerInputGpuConfig) GetGpuType() *string {
	return s.GpuType
}

func (s *DeployCustomContainerInputGpuConfig) SetGpuMemorySize(v int64) *DeployCustomContainerInputGpuConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *DeployCustomContainerInputGpuConfig) SetGpuType(v string) *DeployCustomContainerInputGpuConfig {
	s.GpuType = &v
	return s
}

func (s *DeployCustomContainerInputGpuConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputHttpTrigger struct {
	Qualifier     *string                                             `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig *DeployCustomContainerInputHttpTriggerTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s DeployCustomContainerInputHttpTrigger) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputHttpTrigger) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputHttpTrigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *DeployCustomContainerInputHttpTrigger) GetTriggerConfig() *DeployCustomContainerInputHttpTriggerTriggerConfig {
	return s.TriggerConfig
}

func (s *DeployCustomContainerInputHttpTrigger) SetQualifier(v string) *DeployCustomContainerInputHttpTrigger {
	s.Qualifier = &v
	return s
}

func (s *DeployCustomContainerInputHttpTrigger) SetTriggerConfig(v *DeployCustomContainerInputHttpTriggerTriggerConfig) *DeployCustomContainerInputHttpTrigger {
	s.TriggerConfig = v
	return s
}

func (s *DeployCustomContainerInputHttpTrigger) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputHttpTriggerTriggerConfig struct {
	AuthConfig         *string   `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	AuthType           *string   `json:"authType,omitempty" xml:"authType,omitempty"`
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	DsableURLInternet  *bool     `json:"dsableURLInternet,omitempty" xml:"dsableURLInternet,omitempty"`
	Methods            []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
}

func (s DeployCustomContainerInputHttpTriggerTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputHttpTriggerTriggerConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) GetDisableURLInternet() *bool {
	return s.DisableURLInternet
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) GetDsableURLInternet() *bool {
	return s.DsableURLInternet
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) GetMethods() []*string {
	return s.Methods
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetAuthConfig(v string) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetAuthType(v string) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetDisableURLInternet(v bool) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetDsableURLInternet(v bool) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.DsableURLInternet = &v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) SetMethods(v []*string) *DeployCustomContainerInputHttpTriggerTriggerConfig {
	s.Methods = v
	return s
}

func (s *DeployCustomContainerInputHttpTriggerTriggerConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputLogConfig struct {
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s DeployCustomContainerInputLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputLogConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputLogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *DeployCustomContainerInputLogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *DeployCustomContainerInputLogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *DeployCustomContainerInputLogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *DeployCustomContainerInputLogConfig) GetProject() *string {
	return s.Project
}

func (s *DeployCustomContainerInputLogConfig) SetEnableInstanceMetrics(v bool) *DeployCustomContainerInputLogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetEnableRequestMetrics(v bool) *DeployCustomContainerInputLogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetLogBeginRule(v string) *DeployCustomContainerInputLogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetLogstore(v string) *DeployCustomContainerInputLogConfig {
	s.Logstore = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) SetProject(v string) *DeployCustomContainerInputLogConfig {
	s.Project = &v
	return s
}

func (s *DeployCustomContainerInputLogConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputModelConfig struct {
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
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

func (s DeployCustomContainerInputModelConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputModelConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *DeployCustomContainerInputModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *DeployCustomContainerInputModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DeployCustomContainerInputModelConfig) GetSkipDownload() *bool {
	return s.SkipDownload
}

func (s *DeployCustomContainerInputModelConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *DeployCustomContainerInputModelConfig) GetSrcModelScopeModelID() *string {
	return s.SrcModelScopeModelID
}

func (s *DeployCustomContainerInputModelConfig) GetSrcModelScopeModelRevision() *string {
	return s.SrcModelScopeModelRevision
}

func (s *DeployCustomContainerInputModelConfig) GetSrcModelScopeToken() *string {
	return s.SrcModelScopeToken
}

func (s *DeployCustomContainerInputModelConfig) GetSrcOssBucket() *string {
	return s.SrcOssBucket
}

func (s *DeployCustomContainerInputModelConfig) GetSrcOssPath() *string {
	return s.SrcOssPath
}

func (s *DeployCustomContainerInputModelConfig) GetSrcOssRegion() *string {
	return s.SrcOssRegion
}

func (s *DeployCustomContainerInputModelConfig) GetSyncStrategy() *string {
	return s.SyncStrategy
}

func (s *DeployCustomContainerInputModelConfig) GetWithPPU() *bool {
	return s.WithPPU
}

func (s *DeployCustomContainerInputModelConfig) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DeployCustomContainerInputModelConfig) SetFramework(v string) *DeployCustomContainerInputModelConfig {
	s.Framework = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetMultiModelConfig(v []*ModelConfig) *DeployCustomContainerInputModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetPrefix(v string) *DeployCustomContainerInputModelConfig {
	s.Prefix = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSkipDownload(v bool) *DeployCustomContainerInputModelConfig {
	s.SkipDownload = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSourceType(v string) *DeployCustomContainerInputModelConfig {
	s.SourceType = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeModelID(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeModelID = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeModelRevision(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeModelRevision = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcModelScopeToken(v string) *DeployCustomContainerInputModelConfig {
	s.SrcModelScopeToken = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssBucket(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssBucket = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssPath(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssPath = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSrcOssRegion(v string) *DeployCustomContainerInputModelConfig {
	s.SrcOssRegion = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetSyncStrategy(v string) *DeployCustomContainerInputModelConfig {
	s.SyncStrategy = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetWithPPU(v bool) *DeployCustomContainerInputModelConfig {
	s.WithPPU = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) SetWorkingDir(v string) *DeployCustomContainerInputModelConfig {
	s.WorkingDir = &v
	return s
}

func (s *DeployCustomContainerInputModelConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputNasConfig struct {
	GroupId     *int64    `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*string `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
	UserId      *int64    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeployCustomContainerInputNasConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputNasConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputNasConfig) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeployCustomContainerInputNasConfig) GetMountPoints() []*string {
	return s.MountPoints
}

func (s *DeployCustomContainerInputNasConfig) GetUserId() *int64 {
	return s.UserId
}

func (s *DeployCustomContainerInputNasConfig) SetGroupId(v int64) *DeployCustomContainerInputNasConfig {
	s.GroupId = &v
	return s
}

func (s *DeployCustomContainerInputNasConfig) SetMountPoints(v []*string) *DeployCustomContainerInputNasConfig {
	s.MountPoints = v
	return s
}

func (s *DeployCustomContainerInputNasConfig) SetUserId(v int64) *DeployCustomContainerInputNasConfig {
	s.UserId = &v
	return s
}

func (s *DeployCustomContainerInputNasConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputOssMountConfig struct {
	MountPoints []*DeployCustomContainerInputOssMountConfigMountPoints `json:"mountPoints,omitempty" xml:"mountPoints,omitempty" type:"Repeated"`
}

func (s DeployCustomContainerInputOssMountConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputOssMountConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputOssMountConfig) GetMountPoints() []*DeployCustomContainerInputOssMountConfigMountPoints {
	return s.MountPoints
}

func (s *DeployCustomContainerInputOssMountConfig) SetMountPoints(v []*DeployCustomContainerInputOssMountConfigMountPoints) *DeployCustomContainerInputOssMountConfig {
	s.MountPoints = v
	return s
}

func (s *DeployCustomContainerInputOssMountConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputOssMountConfigMountPoints struct {
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s DeployCustomContainerInputOssMountConfigMountPoints) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputOssMountConfigMountPoints) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) GetBucketName() *string {
	return s.BucketName
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) GetBucketPath() *string {
	return s.BucketPath
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) GetMountDir() *string {
	return s.MountDir
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) SetBucketName(v string) *DeployCustomContainerInputOssMountConfigMountPoints {
	s.BucketName = &v
	return s
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) SetBucketPath(v string) *DeployCustomContainerInputOssMountConfigMountPoints {
	s.BucketPath = &v
	return s
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) SetEndpoint(v string) *DeployCustomContainerInputOssMountConfigMountPoints {
	s.Endpoint = &v
	return s
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) SetMountDir(v string) *DeployCustomContainerInputOssMountConfigMountPoints {
	s.MountDir = &v
	return s
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) SetReadOnly(v bool) *DeployCustomContainerInputOssMountConfigMountPoints {
	s.ReadOnly = &v
	return s
}

func (s *DeployCustomContainerInputOssMountConfigMountPoints) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputProvisionConfig struct {
	AlwaysAllocateGPU *bool                                                        `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*DeployCustomContainerInputProvisionConfigScheduledActions `json:"scheduledActions,omitempty" xml:"scheduledActions,omitempty" type:"Repeated"`
	Target            *int64                                                       `json:"target,omitempty" xml:"target,omitempty"`
}

func (s DeployCustomContainerInputProvisionConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputProvisionConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputProvisionConfig) GetAlwaysAllocateGPU() *bool {
	return s.AlwaysAllocateGPU
}

func (s *DeployCustomContainerInputProvisionConfig) GetScheduledActions() []*DeployCustomContainerInputProvisionConfigScheduledActions {
	return s.ScheduledActions
}

func (s *DeployCustomContainerInputProvisionConfig) GetTarget() *int64 {
	return s.Target
}

func (s *DeployCustomContainerInputProvisionConfig) SetAlwaysAllocateGPU(v bool) *DeployCustomContainerInputProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfig) SetScheduledActions(v []*DeployCustomContainerInputProvisionConfigScheduledActions) *DeployCustomContainerInputProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *DeployCustomContainerInputProvisionConfig) SetTarget(v int64) *DeployCustomContainerInputProvisionConfig {
	s.Target = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfig) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputProvisionConfigScheduledActions struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int32  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s DeployCustomContainerInputProvisionConfigScheduledActions) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputProvisionConfigScheduledActions) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetEndTime() *string {
	return s.EndTime
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetName() *string {
	return s.Name
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetStartTime() *string {
	return s.StartTime
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetTarget() *int32 {
	return s.Target
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetEndTime(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.EndTime = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetName(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.Name = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetScheduleExpression(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.ScheduleExpression = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetStartTime(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.StartTime = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetTarget(v int32) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.Target = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) SetTimeZone(v string) *DeployCustomContainerInputProvisionConfigScheduledActions {
	s.TimeZone = &v
	return s
}

func (s *DeployCustomContainerInputProvisionConfigScheduledActions) Validate() error {
	return dara.Validate(s)
}

type DeployCustomContainerInputVpcConfig struct {
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s DeployCustomContainerInputVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployCustomContainerInputVpcConfig) GoString() string {
	return s.String()
}

func (s *DeployCustomContainerInputVpcConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeployCustomContainerInputVpcConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DeployCustomContainerInputVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DeployCustomContainerInputVpcConfig) SetSecurityGroupId(v string) *DeployCustomContainerInputVpcConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *DeployCustomContainerInputVpcConfig) SetVSwitchIds(v []*string) *DeployCustomContainerInputVpcConfig {
	s.VSwitchIds = v
	return s
}

func (s *DeployCustomContainerInputVpcConfig) SetVpcId(v string) *DeployCustomContainerInputVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DeployCustomContainerInputVpcConfig) Validate() error {
	return dara.Validate(s)
}
