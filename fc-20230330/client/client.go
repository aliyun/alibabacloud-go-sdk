// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AccelerationInfo struct {
	// example:
	//
	// deprecated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AccelerationInfo) String() string {
	return tea.Prettify(s)
}

func (s AccelerationInfo) GoString() string {
	return s.String()
}

func (s *AccelerationInfo) SetStatus(v string) *AccelerationInfo {
	s.Status = &v
	return s
}

type Alias struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// example:
	//
	// prod
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// alias for pre env
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s Alias) String() string {
	return tea.Prettify(s)
}

func (s Alias) GoString() string {
	return s.String()
}

func (s *Alias) SetAdditionalVersionWeight(v map[string]*float32) *Alias {
	s.AdditionalVersionWeight = v
	return s
}

func (s *Alias) SetAliasName(v string) *Alias {
	s.AliasName = &v
	return s
}

func (s *Alias) SetCreatedTime(v string) *Alias {
	s.CreatedTime = &v
	return s
}

func (s *Alias) SetDescription(v string) *Alias {
	s.Description = &v
	return s
}

func (s *Alias) SetLastModifiedTime(v string) *Alias {
	s.LastModifiedTime = &v
	return s
}

func (s *Alias) SetVersionId(v string) *Alias {
	s.VersionId = &v
	return s
}

type AsyncConfig struct {
	// example:
	//
	// true
	AsyncTask *bool `json:"asyncTask,omitempty" xml:"asyncTask,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime       *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:1234/functions/my-func
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 3600
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// example:
	//
	// 3
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s AsyncConfig) String() string {
	return tea.Prettify(s)
}

func (s AsyncConfig) GoString() string {
	return s.String()
}

func (s *AsyncConfig) SetAsyncTask(v bool) *AsyncConfig {
	s.AsyncTask = &v
	return s
}

func (s *AsyncConfig) SetCreatedTime(v string) *AsyncConfig {
	s.CreatedTime = &v
	return s
}

func (s *AsyncConfig) SetDestinationConfig(v *DestinationConfig) *AsyncConfig {
	s.DestinationConfig = v
	return s
}

func (s *AsyncConfig) SetFunctionArn(v string) *AsyncConfig {
	s.FunctionArn = &v
	return s
}

func (s *AsyncConfig) SetLastModifiedTime(v string) *AsyncConfig {
	s.LastModifiedTime = &v
	return s
}

func (s *AsyncConfig) SetMaxAsyncEventAgeInSeconds(v int64) *AsyncConfig {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *AsyncConfig) SetMaxAsyncRetryAttempts(v int64) *AsyncConfig {
	s.MaxAsyncRetryAttempts = &v
	return s
}

type AsyncTask struct {
	// example:
	//
	// 3
	AlreadyRetriedTimes *int64 `json:"alreadyRetriedTimes,omitempty" xml:"alreadyRetriedTimes,omitempty"`
	// example:
	//
	// Succeeded
	DestinationStatus *string `json:"destinationStatus,omitempty" xml:"destinationStatus,omitempty"`
	// example:
	//
	// 1000
	DurationMs *int64 `json:"durationMs,omitempty" xml:"durationMs,omitempty"`
	// example:
	//
	// 1633449590000
	EndTime *int64            `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Events  []*AsyncTaskEvent `json:"events" xml:"events" type:"Repeated"`
	// example:
	//
	// acs:fc:cn-shanghai:1234/functions/my-func
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// D4-*******9FD1-882707E
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// prod
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// e026ae92-61e5-472f-b32d-1c9e3c4e****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// result
	ReturnPayload *string `json:"returnPayload,omitempty" xml:"returnPayload,omitempty"`
	// example:
	//
	// 1633449590000
	StartedTime *int64 `json:"startedTime,omitempty" xml:"startedTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// UnhandledInvocationError
	TaskErrorMessage *string `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// e026ae92-61e5-472f-b32d-1c9e3c4e****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// body
	TaskPayload *string `json:"taskPayload,omitempty" xml:"taskPayload,omitempty"`
}

func (s AsyncTask) String() string {
	return tea.Prettify(s)
}

func (s AsyncTask) GoString() string {
	return s.String()
}

func (s *AsyncTask) SetAlreadyRetriedTimes(v int64) *AsyncTask {
	s.AlreadyRetriedTimes = &v
	return s
}

func (s *AsyncTask) SetDestinationStatus(v string) *AsyncTask {
	s.DestinationStatus = &v
	return s
}

func (s *AsyncTask) SetDurationMs(v int64) *AsyncTask {
	s.DurationMs = &v
	return s
}

func (s *AsyncTask) SetEndTime(v int64) *AsyncTask {
	s.EndTime = &v
	return s
}

func (s *AsyncTask) SetEvents(v []*AsyncTaskEvent) *AsyncTask {
	s.Events = v
	return s
}

func (s *AsyncTask) SetFunctionArn(v string) *AsyncTask {
	s.FunctionArn = &v
	return s
}

func (s *AsyncTask) SetInstanceId(v string) *AsyncTask {
	s.InstanceId = &v
	return s
}

func (s *AsyncTask) SetQualifier(v string) *AsyncTask {
	s.Qualifier = &v
	return s
}

func (s *AsyncTask) SetRequestId(v string) *AsyncTask {
	s.RequestId = &v
	return s
}

func (s *AsyncTask) SetReturnPayload(v string) *AsyncTask {
	s.ReturnPayload = &v
	return s
}

func (s *AsyncTask) SetStartedTime(v int64) *AsyncTask {
	s.StartedTime = &v
	return s
}

func (s *AsyncTask) SetStatus(v string) *AsyncTask {
	s.Status = &v
	return s
}

func (s *AsyncTask) SetTaskErrorMessage(v string) *AsyncTask {
	s.TaskErrorMessage = &v
	return s
}

func (s *AsyncTask) SetTaskId(v string) *AsyncTask {
	s.TaskId = &v
	return s
}

func (s *AsyncTask) SetTaskPayload(v string) *AsyncTask {
	s.TaskPayload = &v
	return s
}

type AsyncTaskEvent struct {
	// example:
	//
	// body
	EventDetail *string `json:"eventDetail,omitempty" xml:"eventDetail,omitempty"`
	// example:
	//
	// 1
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1647420449721
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s AsyncTaskEvent) String() string {
	return tea.Prettify(s)
}

func (s AsyncTaskEvent) GoString() string {
	return s.String()
}

func (s *AsyncTaskEvent) SetEventDetail(v string) *AsyncTaskEvent {
	s.EventDetail = &v
	return s
}

func (s *AsyncTaskEvent) SetEventId(v int64) *AsyncTaskEvent {
	s.EventId = &v
	return s
}

func (s *AsyncTaskEvent) SetStatus(v string) *AsyncTaskEvent {
	s.Status = &v
	return s
}

func (s *AsyncTaskEvent) SetTimestamp(v int64) *AsyncTaskEvent {
	s.Timestamp = &v
	return s
}

type AuthConfig struct {
	// example:
	//
	// {}
	AuthInfo *string `json:"authInfo,omitempty" xml:"authInfo,omitempty"`
	// example:
	//
	// anonymous, function, jwt
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
}

func (s AuthConfig) String() string {
	return tea.Prettify(s)
}

func (s AuthConfig) GoString() string {
	return s.String()
}

func (s *AuthConfig) SetAuthInfo(v string) *AuthConfig {
	s.AuthInfo = &v
	return s
}

func (s *AuthConfig) SetAuthType(v string) *AuthConfig {
	s.AuthType = &v
	return s
}

type BatchWindow struct {
	// example:
	//
	// 100
	CountBasedWindow *int32 `json:"CountBasedWindow,omitempty" xml:"CountBasedWindow,omitempty"`
	// example:
	//
	// 10
	TimeBasedWindow *int32 `json:"TimeBasedWindow,omitempty" xml:"TimeBasedWindow,omitempty"`
}

func (s BatchWindow) String() string {
	return tea.Prettify(s)
}

func (s BatchWindow) GoString() string {
	return s.String()
}

func (s *BatchWindow) SetCountBasedWindow(v int32) *BatchWindow {
	s.CountBasedWindow = &v
	return s
}

func (s *BatchWindow) SetTimeBasedWindow(v int32) *BatchWindow {
	s.TimeBasedWindow = &v
	return s
}

type CDNTriggerConfig struct {
	// example:
	//
	// CdnDomainStarted
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// example:
	//
	// 1.0.0
	EventVersion *string              `json:"eventVersion,omitempty" xml:"eventVersion,omitempty"`
	Filter       map[string][]*string `json:"filter" xml:"filter"`
	// example:
	//
	// 缓存事件触发器
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
}

func (s CDNTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s CDNTriggerConfig) GoString() string {
	return s.String()
}

func (s *CDNTriggerConfig) SetEventName(v string) *CDNTriggerConfig {
	s.EventName = &v
	return s
}

func (s *CDNTriggerConfig) SetEventVersion(v string) *CDNTriggerConfig {
	s.EventVersion = &v
	return s
}

func (s *CDNTriggerConfig) SetFilter(v map[string][]*string) *CDNTriggerConfig {
	s.Filter = v
	return s
}

func (s *CDNTriggerConfig) SetNotes(v string) *CDNTriggerConfig {
	s.Notes = &v
	return s
}

type CertConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// my-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PEM format
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PEM format
	PrivateKey *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
}

func (s CertConfig) String() string {
	return tea.Prettify(s)
}

func (s CertConfig) GoString() string {
	return s.String()
}

func (s *CertConfig) SetCertName(v string) *CertConfig {
	s.CertName = &v
	return s
}

func (s *CertConfig) SetCertificate(v string) *CertConfig {
	s.Certificate = &v
	return s
}

func (s *CertConfig) SetPrivateKey(v string) *CertConfig {
	s.PrivateKey = &v
	return s
}

type ConcurrencyConfig struct {
	// example:
	//
	// acs:fc:cn-shanghai:123:functions/demo
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// 10
	ReservedConcurrency *int64 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s ConcurrencyConfig) String() string {
	return tea.Prettify(s)
}

func (s ConcurrencyConfig) GoString() string {
	return s.String()
}

func (s *ConcurrencyConfig) SetFunctionArn(v string) *ConcurrencyConfig {
	s.FunctionArn = &v
	return s
}

func (s *ConcurrencyConfig) SetReservedConcurrency(v int64) *ConcurrencyConfig {
	s.ReservedConcurrency = &v
	return s
}

type CreateAliasInput struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// This parameter is required.
	//
	// example:
	//
	// prod
	AliasName *string `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	// example:
	//
	// my alias
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s CreateAliasInput) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasInput) GoString() string {
	return s.String()
}

func (s *CreateAliasInput) SetAdditionalVersionWeight(v map[string]*float32) *CreateAliasInput {
	s.AdditionalVersionWeight = v
	return s
}

func (s *CreateAliasInput) SetAliasName(v string) *CreateAliasInput {
	s.AliasName = &v
	return s
}

func (s *CreateAliasInput) SetDescription(v string) *CreateAliasInput {
	s.Description = &v
	return s
}

func (s *CreateAliasInput) SetVersionId(v string) *CreateAliasInput {
	s.VersionId = &v
	return s
}

type CreateCustomDomainInput struct {
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// example:
	//
	// HTTP
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	TlsConfig   *TLSConfig   `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig   *WAFConfig   `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s CreateCustomDomainInput) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainInput) SetAuthConfig(v *AuthConfig) *CreateCustomDomainInput {
	s.AuthConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetCertConfig(v *CertConfig) *CreateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetDomainName(v string) *CreateCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainInput) SetProtocol(v string) *CreateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *CreateCustomDomainInput) SetRouteConfig(v *RouteConfig) *CreateCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetWafConfig(v *WAFConfig) *CreateCustomDomainInput {
	s.WafConfig = v
	return s
}

type CreateFunctionInput struct {
	Code *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Cpu                   *float32               `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS             *CustomDNS             `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig   *CustomRuntimeConfig   `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// example:
	//
	// my function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 512
	DiskSize             *int32             `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// This parameter is required.
	//
	// example:
	//
	// my-function-1
	FunctionName *string    `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig    *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 1
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// example:
	//
	// true
	InternetAccess *bool      `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers         []*string  `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig      *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// example:
	//
	// 512
	MemorySize     *int32          `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig      *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s CreateFunctionInput) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInput) GoString() string {
	return s.String()
}

func (s *CreateFunctionInput) SetCode(v *InputCodeLocation) *CreateFunctionInput {
	s.Code = v
	return s
}

func (s *CreateFunctionInput) SetCpu(v float32) *CreateFunctionInput {
	s.Cpu = &v
	return s
}

func (s *CreateFunctionInput) SetCustomContainerConfig(v *CustomContainerConfig) *CreateFunctionInput {
	s.CustomContainerConfig = v
	return s
}

func (s *CreateFunctionInput) SetCustomDNS(v *CustomDNS) *CreateFunctionInput {
	s.CustomDNS = v
	return s
}

func (s *CreateFunctionInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *CreateFunctionInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *CreateFunctionInput) SetDescription(v string) *CreateFunctionInput {
	s.Description = &v
	return s
}

func (s *CreateFunctionInput) SetDiskSize(v int32) *CreateFunctionInput {
	s.DiskSize = &v
	return s
}

func (s *CreateFunctionInput) SetEnvironmentVariables(v map[string]*string) *CreateFunctionInput {
	s.EnvironmentVariables = v
	return s
}

func (s *CreateFunctionInput) SetFunctionName(v string) *CreateFunctionInput {
	s.FunctionName = &v
	return s
}

func (s *CreateFunctionInput) SetGpuConfig(v *GPUConfig) *CreateFunctionInput {
	s.GpuConfig = v
	return s
}

func (s *CreateFunctionInput) SetHandler(v string) *CreateFunctionInput {
	s.Handler = &v
	return s
}

func (s *CreateFunctionInput) SetInstanceConcurrency(v int32) *CreateFunctionInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *CreateFunctionInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *CreateFunctionInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *CreateFunctionInput) SetInternetAccess(v bool) *CreateFunctionInput {
	s.InternetAccess = &v
	return s
}

func (s *CreateFunctionInput) SetLayers(v []*string) *CreateFunctionInput {
	s.Layers = v
	return s
}

func (s *CreateFunctionInput) SetLogConfig(v *LogConfig) *CreateFunctionInput {
	s.LogConfig = v
	return s
}

func (s *CreateFunctionInput) SetMemorySize(v int32) *CreateFunctionInput {
	s.MemorySize = &v
	return s
}

func (s *CreateFunctionInput) SetNasConfig(v *NASConfig) *CreateFunctionInput {
	s.NasConfig = v
	return s
}

func (s *CreateFunctionInput) SetOssMountConfig(v *OSSMountConfig) *CreateFunctionInput {
	s.OssMountConfig = v
	return s
}

func (s *CreateFunctionInput) SetRole(v string) *CreateFunctionInput {
	s.Role = &v
	return s
}

func (s *CreateFunctionInput) SetRuntime(v string) *CreateFunctionInput {
	s.Runtime = &v
	return s
}

func (s *CreateFunctionInput) SetTimeout(v int32) *CreateFunctionInput {
	s.Timeout = &v
	return s
}

func (s *CreateFunctionInput) SetTracingConfig(v *TracingConfig) *CreateFunctionInput {
	s.TracingConfig = v
	return s
}

func (s *CreateFunctionInput) SetVpcConfig(v *VPCConfig) *CreateFunctionInput {
	s.VpcConfig = v
	return s
}

type CreateLayerVersionInput struct {
	Code              *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	CompatibleRuntime []*string          `json:"compatibleRuntime" xml:"compatibleRuntime" type:"Repeated"`
	// example:
	//
	// my first layer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Apache
	License *string `json:"license,omitempty" xml:"license,omitempty"`
}

func (s CreateLayerVersionInput) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionInput) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionInput) SetCode(v *InputCodeLocation) *CreateLayerVersionInput {
	s.Code = v
	return s
}

func (s *CreateLayerVersionInput) SetCompatibleRuntime(v []*string) *CreateLayerVersionInput {
	s.CompatibleRuntime = v
	return s
}

func (s *CreateLayerVersionInput) SetDescription(v string) *CreateLayerVersionInput {
	s.Description = &v
	return s
}

func (s *CreateLayerVersionInput) SetLicense(v string) *CreateLayerVersionInput {
	s.License = &v
	return s
}

type CreateTriggerInput struct {
	// example:
	//
	// trigger for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// acs:ram::1234567890:role/fc-test
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:12345:mybucket
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"events":["oss:ObjectCreated:*"],"filter":{"key":{"prefix":"/prefix","suffix":".zip"}}}
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss_create_object_demo
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s CreateTriggerInput) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerInput) GoString() string {
	return s.String()
}

func (s *CreateTriggerInput) SetDescription(v string) *CreateTriggerInput {
	s.Description = &v
	return s
}

func (s *CreateTriggerInput) SetInvocationRole(v string) *CreateTriggerInput {
	s.InvocationRole = &v
	return s
}

func (s *CreateTriggerInput) SetQualifier(v string) *CreateTriggerInput {
	s.Qualifier = &v
	return s
}

func (s *CreateTriggerInput) SetSourceArn(v string) *CreateTriggerInput {
	s.SourceArn = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerConfig(v string) *CreateTriggerInput {
	s.TriggerConfig = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerName(v string) *CreateTriggerInput {
	s.TriggerName = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerType(v string) *CreateTriggerInput {
	s.TriggerType = &v
	return s
}

type CreateVpcBindingInput struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc-8vb8x8dggvr0axxxxxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateVpcBindingInput) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingInput) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingInput) SetVpcId(v string) *CreateVpcBindingInput {
	s.VpcId = &v
	return s
}

type CustomContainerConfig struct {
	AccelerationInfo *AccelerationInfo `json:"accelerationInfo,omitempty" xml:"accelerationInfo,omitempty"`
	// example:
	//
	// deprecated
	AccelerationType *string `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	// example:
	//
	// deprecated
	AcrInstanceId     *string                  `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	Command           []*string                `json:"command" xml:"command" type:"Repeated"`
	Entrypoint        []*string                `json:"entrypoint" xml:"entrypoint" type:"Repeated"`
	HealthCheckConfig *CustomHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/fc-demo/helloworld:v1
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
	// example:
	//
	// 9000
	Port           *int32          `json:"port,omitempty" xml:"port,omitempty"`
	RegistryConfig *RegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	// example:
	//
	// stand-sh-registry-vpc.cn-shanghai.cr.aliyuncs.com/fc-demo2/springboot-helloworld@sha256:68d1****0d64d6
	ResolvedImageUri *string `json:"resolvedImageUri,omitempty" xml:"resolvedImageUri,omitempty"`
}

func (s CustomContainerConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomContainerConfig) GoString() string {
	return s.String()
}

func (s *CustomContainerConfig) SetAccelerationInfo(v *AccelerationInfo) *CustomContainerConfig {
	s.AccelerationInfo = v
	return s
}

func (s *CustomContainerConfig) SetAccelerationType(v string) *CustomContainerConfig {
	s.AccelerationType = &v
	return s
}

func (s *CustomContainerConfig) SetAcrInstanceId(v string) *CustomContainerConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *CustomContainerConfig) SetCommand(v []*string) *CustomContainerConfig {
	s.Command = v
	return s
}

func (s *CustomContainerConfig) SetEntrypoint(v []*string) *CustomContainerConfig {
	s.Entrypoint = v
	return s
}

func (s *CustomContainerConfig) SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomContainerConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *CustomContainerConfig) SetImage(v string) *CustomContainerConfig {
	s.Image = &v
	return s
}

func (s *CustomContainerConfig) SetPort(v int32) *CustomContainerConfig {
	s.Port = &v
	return s
}

func (s *CustomContainerConfig) SetRegistryConfig(v *RegistryConfig) *CustomContainerConfig {
	s.RegistryConfig = v
	return s
}

func (s *CustomContainerConfig) SetResolvedImageUri(v string) *CustomContainerConfig {
	s.ResolvedImageUri = &v
	return s
}

type CustomDNS struct {
	DnsOptions  []*DNSOption `json:"dnsOptions" xml:"dnsOptions" type:"Repeated"`
	NameServers []*string    `json:"nameServers" xml:"nameServers" type:"Repeated"`
	Searches    []*string    `json:"searches" xml:"searches" type:"Repeated"`
}

func (s CustomDNS) String() string {
	return tea.Prettify(s)
}

func (s CustomDNS) GoString() string {
	return s.String()
}

func (s *CustomDNS) SetDnsOptions(v []*DNSOption) *CustomDNS {
	s.DnsOptions = v
	return s
}

func (s *CustomDNS) SetNameServers(v []*string) *CustomDNS {
	s.NameServers = v
	return s
}

func (s *CustomDNS) SetSearches(v []*string) *CustomDNS {
	s.Searches = v
	return s
}

type CustomDomain struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 2023-03-30
	ApiVersion *string     `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// example:
	//
	// 2023-03-30T08:02:19Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// example:
	//
	// 2023-03-30T08:02:19Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// HTTP
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// example:
	//
	// 1
	SubdomainCount *string    `json:"subdomainCount,omitempty" xml:"subdomainCount,omitempty"`
	TlsConfig      *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig      *WAFConfig `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s CustomDomain) String() string {
	return tea.Prettify(s)
}

func (s CustomDomain) GoString() string {
	return s.String()
}

func (s *CustomDomain) SetAccountId(v string) *CustomDomain {
	s.AccountId = &v
	return s
}

func (s *CustomDomain) SetApiVersion(v string) *CustomDomain {
	s.ApiVersion = &v
	return s
}

func (s *CustomDomain) SetAuthConfig(v *AuthConfig) *CustomDomain {
	s.AuthConfig = v
	return s
}

func (s *CustomDomain) SetCertConfig(v *CertConfig) *CustomDomain {
	s.CertConfig = v
	return s
}

func (s *CustomDomain) SetCreatedTime(v string) *CustomDomain {
	s.CreatedTime = &v
	return s
}

func (s *CustomDomain) SetDomainName(v string) *CustomDomain {
	s.DomainName = &v
	return s
}

func (s *CustomDomain) SetLastModifiedTime(v string) *CustomDomain {
	s.LastModifiedTime = &v
	return s
}

func (s *CustomDomain) SetProtocol(v string) *CustomDomain {
	s.Protocol = &v
	return s
}

func (s *CustomDomain) SetRouteConfig(v *RouteConfig) *CustomDomain {
	s.RouteConfig = v
	return s
}

func (s *CustomDomain) SetSubdomainCount(v string) *CustomDomain {
	s.SubdomainCount = &v
	return s
}

func (s *CustomDomain) SetTlsConfig(v *TLSConfig) *CustomDomain {
	s.TlsConfig = v
	return s
}

func (s *CustomDomain) SetWafConfig(v *WAFConfig) *CustomDomain {
	s.WafConfig = v
	return s
}

type CustomHealthCheckConfig struct {
	// example:
	//
	// 1
	FailureThreshold *int32 `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	// example:
	//
	// /ready
	HttpGetUrl *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	// example:
	//
	// 1
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	// example:
	//
	// 1
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	// example:
	//
	// 2
	SuccessThreshold *int32 `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	// example:
	//
	// 2
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
}

func (s CustomHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CustomHealthCheckConfig) SetFailureThreshold(v int32) *CustomHealthCheckConfig {
	s.FailureThreshold = &v
	return s
}

func (s *CustomHealthCheckConfig) SetHttpGetUrl(v string) *CustomHealthCheckConfig {
	s.HttpGetUrl = &v
	return s
}

func (s *CustomHealthCheckConfig) SetInitialDelaySeconds(v int32) *CustomHealthCheckConfig {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CustomHealthCheckConfig) SetPeriodSeconds(v int32) *CustomHealthCheckConfig {
	s.PeriodSeconds = &v
	return s
}

func (s *CustomHealthCheckConfig) SetSuccessThreshold(v int32) *CustomHealthCheckConfig {
	s.SuccessThreshold = &v
	return s
}

func (s *CustomHealthCheckConfig) SetTimeoutSeconds(v int32) *CustomHealthCheckConfig {
	s.TimeoutSeconds = &v
	return s
}

type CustomRuntimeConfig struct {
	Args              []*string                `json:"args" xml:"args" type:"Repeated"`
	Command           []*string                `json:"command" xml:"command" type:"Repeated"`
	HealthCheckConfig *CustomHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty"`
	// example:
	//
	// 9000
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s CustomRuntimeConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomRuntimeConfig) GoString() string {
	return s.String()
}

func (s *CustomRuntimeConfig) SetArgs(v []*string) *CustomRuntimeConfig {
	s.Args = v
	return s
}

func (s *CustomRuntimeConfig) SetCommand(v []*string) *CustomRuntimeConfig {
	s.Command = v
	return s
}

func (s *CustomRuntimeConfig) SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomRuntimeConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *CustomRuntimeConfig) SetPort(v int32) *CustomRuntimeConfig {
	s.Port = &v
	return s
}

type DNSOption struct {
	// example:
	//
	// ndots
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DNSOption) String() string {
	return tea.Prettify(s)
}

func (s DNSOption) GoString() string {
	return s.String()
}

func (s *DNSOption) SetName(v string) *DNSOption {
	s.Name = &v
	return s
}

func (s *DNSOption) SetValue(v string) *DNSOption {
	s.Value = &v
	return s
}

type DeadLetterQueue struct {
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
}

func (s DeadLetterQueue) String() string {
	return tea.Prettify(s)
}

func (s DeadLetterQueue) GoString() string {
	return s.String()
}

func (s *DeadLetterQueue) SetArn(v string) *DeadLetterQueue {
	s.Arn = &v
	return s
}

type DeliveryOption struct {
	// example:
	//
	// 2
	Concurrency *int64 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// example:
	//
	// RawData
	EventSchema *string `json:"eventSchema,omitempty" xml:"eventSchema,omitempty"`
}

func (s DeliveryOption) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOption) GoString() string {
	return s.String()
}

func (s *DeliveryOption) SetConcurrency(v int64) *DeliveryOption {
	s.Concurrency = &v
	return s
}

func (s *DeliveryOption) SetEventSchema(v string) *DeliveryOption {
	s.EventSchema = &v
	return s
}

type Destination struct {
	// example:
	//
	// acs:fc:cn-shanghai:xxx:functions/f1
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
}

func (s Destination) String() string {
	return tea.Prettify(s)
}

func (s Destination) GoString() string {
	return s.String()
}

func (s *Destination) SetDestination(v string) *Destination {
	s.Destination = &v
	return s
}

type DestinationConfig struct {
	OnFailure *Destination `json:"onFailure,omitempty" xml:"onFailure,omitempty"`
	OnSuccess *Destination `json:"onSuccess,omitempty" xml:"onSuccess,omitempty"`
}

func (s DestinationConfig) String() string {
	return tea.Prettify(s)
}

func (s DestinationConfig) GoString() string {
	return s.String()
}

func (s *DestinationConfig) SetOnFailure(v *Destination) *DestinationConfig {
	s.OnFailure = v
	return s
}

func (s *DestinationConfig) SetOnSuccess(v *Destination) *DestinationConfig {
	s.OnSuccess = v
	return s
}

type EqualRule struct {
	// This parameter is required.
	//
	// example:
	//
	// /old
	Match *string `json:"match,omitempty" xml:"match,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /new
	Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s EqualRule) String() string {
	return tea.Prettify(s)
}

func (s EqualRule) GoString() string {
	return s.String()
}

func (s *EqualRule) SetMatch(v string) *EqualRule {
	s.Match = &v
	return s
}

func (s *EqualRule) SetReplacement(v string) *EqualRule {
	s.Replacement = &v
	return s
}

type Error struct {
	// example:
	//
	// FunctionNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// function not found
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1-64e70cf1-5cbef92ea8fc8c42899cf5d1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s Error) String() string {
	return tea.Prettify(s)
}

func (s Error) GoString() string {
	return s.String()
}

func (s *Error) SetCode(v string) *Error {
	s.Code = &v
	return s
}

func (s *Error) SetMessage(v string) *Error {
	s.Message = &v
	return s
}

func (s *Error) SetRequestId(v string) *Error {
	s.RequestId = &v
	return s
}

type EventBridgeTriggerConfig struct {
	AsyncInvocationType *bool `json:"asyncInvocationType,omitempty" xml:"asyncInvocationType,omitempty"`
	// example:
	//
	// {}
	EventRuleFilterPattern *string            `json:"eventRuleFilterPattern,omitempty" xml:"eventRuleFilterPattern,omitempty"`
	EventSinkConfig        *EventSinkConfig   `json:"eventSinkConfig,omitempty" xml:"eventSinkConfig,omitempty"`
	EventSourceConfig      *EventSourceConfig `json:"eventSourceConfig,omitempty" xml:"eventSourceConfig,omitempty"`
	RunOptions             *RunOptions        `json:"runOptions,omitempty" xml:"runOptions,omitempty"`
	// example:
	//
	// true
	TriggerEnable *bool `json:"triggerEnable,omitempty" xml:"triggerEnable,omitempty"`
}

func (s EventBridgeTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s EventBridgeTriggerConfig) GoString() string {
	return s.String()
}

func (s *EventBridgeTriggerConfig) SetAsyncInvocationType(v bool) *EventBridgeTriggerConfig {
	s.AsyncInvocationType = &v
	return s
}

func (s *EventBridgeTriggerConfig) SetEventRuleFilterPattern(v string) *EventBridgeTriggerConfig {
	s.EventRuleFilterPattern = &v
	return s
}

func (s *EventBridgeTriggerConfig) SetEventSinkConfig(v *EventSinkConfig) *EventBridgeTriggerConfig {
	s.EventSinkConfig = v
	return s
}

func (s *EventBridgeTriggerConfig) SetEventSourceConfig(v *EventSourceConfig) *EventBridgeTriggerConfig {
	s.EventSourceConfig = v
	return s
}

func (s *EventBridgeTriggerConfig) SetRunOptions(v *RunOptions) *EventBridgeTriggerConfig {
	s.RunOptions = v
	return s
}

func (s *EventBridgeTriggerConfig) SetTriggerEnable(v bool) *EventBridgeTriggerConfig {
	s.TriggerEnable = &v
	return s
}

type EventSinkConfig struct {
	DeliveryOption *DeliveryOption `json:"deliveryOption,omitempty" xml:"deliveryOption,omitempty"`
}

func (s EventSinkConfig) String() string {
	return tea.Prettify(s)
}

func (s EventSinkConfig) GoString() string {
	return s.String()
}

func (s *EventSinkConfig) SetDeliveryOption(v *DeliveryOption) *EventSinkConfig {
	s.DeliveryOption = v
	return s
}

type EventSourceConfig struct {
	EventSourceParameters *EventSourceParameters `json:"eventSourceParameters,omitempty" xml:"eventSourceParameters,omitempty"`
	// example:
	//
	// MNS
	EventSourceType *string `json:"eventSourceType,omitempty" xml:"eventSourceType,omitempty"`
}

func (s EventSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s EventSourceConfig) GoString() string {
	return s.String()
}

func (s *EventSourceConfig) SetEventSourceParameters(v *EventSourceParameters) *EventSourceConfig {
	s.EventSourceParameters = v
	return s
}

func (s *EventSourceConfig) SetEventSourceType(v string) *EventSourceConfig {
	s.EventSourceType = &v
	return s
}

type EventSourceParameters struct {
	SourceDTSParameters      *SourceDTSParameters      `json:"sourceDTSParameters,omitempty" xml:"sourceDTSParameters,omitempty"`
	SourceKafkaParameters    *SourceKafkaParameters    `json:"sourceKafkaParameters,omitempty" xml:"sourceKafkaParameters,omitempty"`
	SourceMNSParameters      *SourceMNSParameters      `json:"sourceMNSParameters,omitempty" xml:"sourceMNSParameters,omitempty"`
	SourceMQTTParameters     *SourceMQTTParameters     `json:"sourceMQTTParameters,omitempty" xml:"sourceMQTTParameters,omitempty"`
	SourceRabbitMQParameters *SourceRabbitMQParameters `json:"sourceRabbitMQParameters,omitempty" xml:"sourceRabbitMQParameters,omitempty"`
	SourceRocketMQParameters *SourceRocketMQParameters `json:"sourceRocketMQParameters,omitempty" xml:"sourceRocketMQParameters,omitempty"`
}

func (s EventSourceParameters) String() string {
	return tea.Prettify(s)
}

func (s EventSourceParameters) GoString() string {
	return s.String()
}

func (s *EventSourceParameters) SetSourceDTSParameters(v *SourceDTSParameters) *EventSourceParameters {
	s.SourceDTSParameters = v
	return s
}

func (s *EventSourceParameters) SetSourceKafkaParameters(v *SourceKafkaParameters) *EventSourceParameters {
	s.SourceKafkaParameters = v
	return s
}

func (s *EventSourceParameters) SetSourceMNSParameters(v *SourceMNSParameters) *EventSourceParameters {
	s.SourceMNSParameters = v
	return s
}

func (s *EventSourceParameters) SetSourceMQTTParameters(v *SourceMQTTParameters) *EventSourceParameters {
	s.SourceMQTTParameters = v
	return s
}

func (s *EventSourceParameters) SetSourceRabbitMQParameters(v *SourceRabbitMQParameters) *EventSourceParameters {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *EventSourceParameters) SetSourceRocketMQParameters(v *SourceRocketMQParameters) *EventSourceParameters {
	s.SourceRocketMQParameters = v
	return s
}

type Filter struct {
	Key *Key `json:"key,omitempty" xml:"key,omitempty"`
}

func (s Filter) String() string {
	return tea.Prettify(s)
}

func (s Filter) GoString() string {
	return s.String()
}

func (s *Filter) SetKey(v *Key) *Filter {
	s.Key = v
	return s
}

type Function struct {
	// example:
	//
	// 2825179536350****
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// example:
	//
	// 412
	CodeSize *int64 `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	// example:
	//
	// 1
	Cpu *float32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// 2023-04-01T08:15:27Z
	CreatedTime           *string                `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS             *CustomDNS             `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig   *CustomRuntimeConfig   `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// example:
	//
	// my function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 512
	DiskSize             *int32             `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	// example:
	//
	// acs:fc:cn-shanghai:123:functions/functionName
	FunctionArn *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	// example:
	//
	// aa715851-1c20-4b89-a8fb-***
	FunctionId *string `json:"functionId,omitempty" xml:"functionId,omitempty"`
	// example:
	//
	// my-function-1
	FunctionName *string    `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig    *GPUConfig `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 1
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// example:
	//
	// true
	InternetAccess *bool `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	// example:
	//
	// 2023-05-01T08:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// InProgress
	LastUpdateStatus *string `json:"lastUpdateStatus,omitempty" xml:"lastUpdateStatus,omitempty"`
	// example:
	//
	// The system is currently processing the acceleration optimization for the image.
	LastUpdateStatusReason *string `json:"lastUpdateStatusReason,omitempty" xml:"lastUpdateStatusReason,omitempty"`
	// example:
	//
	// ImageOptimizing
	LastUpdateStatusReasonCode *string          `json:"lastUpdateStatusReasonCode,omitempty" xml:"lastUpdateStatusReasonCode,omitempty"`
	Layers                     []*FunctionLayer `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig                  *LogConfig       `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// example:
	//
	// 512
	MemorySize     *int32          `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig      *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// python3.10
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// Pending
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// Function creating
	StateReason *string `json:"stateReason,omitempty" xml:"stateReason,omitempty"`
	// example:
	//
	// Creating
	StateReasonCode *string `json:"stateReasonCode,omitempty" xml:"stateReasonCode,omitempty"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s Function) String() string {
	return tea.Prettify(s)
}

func (s Function) GoString() string {
	return s.String()
}

func (s *Function) SetCodeChecksum(v string) *Function {
	s.CodeChecksum = &v
	return s
}

func (s *Function) SetCodeSize(v int64) *Function {
	s.CodeSize = &v
	return s
}

func (s *Function) SetCpu(v float32) *Function {
	s.Cpu = &v
	return s
}

func (s *Function) SetCreatedTime(v string) *Function {
	s.CreatedTime = &v
	return s
}

func (s *Function) SetCustomContainerConfig(v *CustomContainerConfig) *Function {
	s.CustomContainerConfig = v
	return s
}

func (s *Function) SetCustomDNS(v *CustomDNS) *Function {
	s.CustomDNS = v
	return s
}

func (s *Function) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *Function {
	s.CustomRuntimeConfig = v
	return s
}

func (s *Function) SetDescription(v string) *Function {
	s.Description = &v
	return s
}

func (s *Function) SetDiskSize(v int32) *Function {
	s.DiskSize = &v
	return s
}

func (s *Function) SetEnvironmentVariables(v map[string]*string) *Function {
	s.EnvironmentVariables = v
	return s
}

func (s *Function) SetFunctionArn(v string) *Function {
	s.FunctionArn = &v
	return s
}

func (s *Function) SetFunctionId(v string) *Function {
	s.FunctionId = &v
	return s
}

func (s *Function) SetFunctionName(v string) *Function {
	s.FunctionName = &v
	return s
}

func (s *Function) SetGpuConfig(v *GPUConfig) *Function {
	s.GpuConfig = v
	return s
}

func (s *Function) SetHandler(v string) *Function {
	s.Handler = &v
	return s
}

func (s *Function) SetInstanceConcurrency(v int32) *Function {
	s.InstanceConcurrency = &v
	return s
}

func (s *Function) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *Function {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *Function) SetInternetAccess(v bool) *Function {
	s.InternetAccess = &v
	return s
}

func (s *Function) SetLastModifiedTime(v string) *Function {
	s.LastModifiedTime = &v
	return s
}

func (s *Function) SetLastUpdateStatus(v string) *Function {
	s.LastUpdateStatus = &v
	return s
}

func (s *Function) SetLastUpdateStatusReason(v string) *Function {
	s.LastUpdateStatusReason = &v
	return s
}

func (s *Function) SetLastUpdateStatusReasonCode(v string) *Function {
	s.LastUpdateStatusReasonCode = &v
	return s
}

func (s *Function) SetLayers(v []*FunctionLayer) *Function {
	s.Layers = v
	return s
}

func (s *Function) SetLogConfig(v *LogConfig) *Function {
	s.LogConfig = v
	return s
}

func (s *Function) SetMemorySize(v int32) *Function {
	s.MemorySize = &v
	return s
}

func (s *Function) SetNasConfig(v *NASConfig) *Function {
	s.NasConfig = v
	return s
}

func (s *Function) SetOssMountConfig(v *OSSMountConfig) *Function {
	s.OssMountConfig = v
	return s
}

func (s *Function) SetRole(v string) *Function {
	s.Role = &v
	return s
}

func (s *Function) SetRuntime(v string) *Function {
	s.Runtime = &v
	return s
}

func (s *Function) SetState(v string) *Function {
	s.State = &v
	return s
}

func (s *Function) SetStateReason(v string) *Function {
	s.StateReason = &v
	return s
}

func (s *Function) SetStateReasonCode(v string) *Function {
	s.StateReasonCode = &v
	return s
}

func (s *Function) SetTimeout(v int32) *Function {
	s.Timeout = &v
	return s
}

func (s *Function) SetTracingConfig(v *TracingConfig) *Function {
	s.TracingConfig = v
	return s
}

func (s *Function) SetVpcConfig(v *VPCConfig) *Function {
	s.VpcConfig = v
	return s
}

type FunctionLayer struct {
	// example:
	//
	// acs:fc:cn-beijing:186824xxxxxx:layers/fc_layer/versions/1
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// example:
	//
	// 421
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s FunctionLayer) String() string {
	return tea.Prettify(s)
}

func (s FunctionLayer) GoString() string {
	return s.String()
}

func (s *FunctionLayer) SetArn(v string) *FunctionLayer {
	s.Arn = &v
	return s
}

func (s *FunctionLayer) SetSize(v int64) *FunctionLayer {
	s.Size = &v
	return s
}

type GPUConfig struct {
	// example:
	//
	// 2048
	GpuMemorySize *int32 `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	// example:
	//
	// fc.gpu.ampere.1
	GpuType *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
}

func (s GPUConfig) String() string {
	return tea.Prettify(s)
}

func (s GPUConfig) GoString() string {
	return s.String()
}

func (s *GPUConfig) SetGpuMemorySize(v int32) *GPUConfig {
	s.GpuMemorySize = &v
	return s
}

func (s *GPUConfig) SetGpuType(v string) *GPUConfig {
	s.GpuType = &v
	return s
}

type GetResourceTagsOutput struct {
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResouceType *string `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:****:functions/demo
	ResourceArn *string            `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	Tags        map[string]*string `json:"tags" xml:"tags"`
}

func (s GetResourceTagsOutput) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsOutput) GoString() string {
	return s.String()
}

func (s *GetResourceTagsOutput) SetResouceType(v string) *GetResourceTagsOutput {
	s.ResouceType = &v
	return s
}

func (s *GetResourceTagsOutput) SetResourceArn(v string) *GetResourceTagsOutput {
	s.ResourceArn = &v
	return s
}

func (s *GetResourceTagsOutput) SetTags(v map[string]*string) *GetResourceTagsOutput {
	s.Tags = v
	return s
}

type HTTPTrigger struct {
	// example:
	//
	// https://svc-func-xxxxxxxx.cn-hangzhou.fcapp.run
	UrlInternet *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
	// example:
	//
	// https://svc-func-xxxxxxxx.cn-hangzhou-vpc.fcapp.run
	UrlIntranet *string `json:"urlIntranet,omitempty" xml:"urlIntranet,omitempty"`
}

func (s HTTPTrigger) String() string {
	return tea.Prettify(s)
}

func (s HTTPTrigger) GoString() string {
	return s.String()
}

func (s *HTTPTrigger) SetUrlInternet(v string) *HTTPTrigger {
	s.UrlInternet = &v
	return s
}

func (s *HTTPTrigger) SetUrlIntranet(v string) *HTTPTrigger {
	s.UrlIntranet = &v
	return s
}

type HTTPTriggerConfig struct {
	// example:
	//
	// {"JWKS":{"foo":"bar"},"TokenLookup":"header:Authorization:Bearer,cookie:AuthorizationCookie","ClaimPassBy":"query:uid:uid,header:name:name"}
	AuthConfig *string `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// example:
	//
	// anonymous
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// example:
	//
	// true
	DisableURLInternet *bool     `json:"disableURLInternet,omitempty" xml:"disableURLInternet,omitempty"`
	Methods            []*string `json:"methods" xml:"methods" type:"Repeated"`
}

func (s HTTPTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s HTTPTriggerConfig) GoString() string {
	return s.String()
}

func (s *HTTPTriggerConfig) SetAuthConfig(v string) *HTTPTriggerConfig {
	s.AuthConfig = &v
	return s
}

func (s *HTTPTriggerConfig) SetAuthType(v string) *HTTPTriggerConfig {
	s.AuthType = &v
	return s
}

func (s *HTTPTriggerConfig) SetDisableURLInternet(v bool) *HTTPTriggerConfig {
	s.DisableURLInternet = &v
	return s
}

func (s *HTTPTriggerConfig) SetMethods(v []*string) *HTTPTriggerConfig {
	s.Methods = v
	return s
}

type InputCodeLocation struct {
	// example:
	//
	// 2825179536350****
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// demo-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// example:
	//
	// demo-object
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// example:
	//
	// UEsDBAoAAAAAANF
	ZipFile *string `json:"zipFile,omitempty" xml:"zipFile,omitempty"`
}

func (s InputCodeLocation) String() string {
	return tea.Prettify(s)
}

func (s InputCodeLocation) GoString() string {
	return s.String()
}

func (s *InputCodeLocation) SetChecksum(v string) *InputCodeLocation {
	s.Checksum = &v
	return s
}

func (s *InputCodeLocation) SetOssBucketName(v string) *InputCodeLocation {
	s.OssBucketName = &v
	return s
}

func (s *InputCodeLocation) SetOssObjectName(v string) *InputCodeLocation {
	s.OssObjectName = &v
	return s
}

func (s *InputCodeLocation) SetZipFile(v string) *InputCodeLocation {
	s.ZipFile = &v
	return s
}

type InstanceInfo struct {
	// example:
	//
	// 1ef6b6ff-7f7b-485e-ab49-501ac681****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// LATEST
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s InstanceInfo) String() string {
	return tea.Prettify(s)
}

func (s InstanceInfo) GoString() string {
	return s.String()
}

func (s *InstanceInfo) SetInstanceId(v string) *InstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *InstanceInfo) SetVersionId(v string) *InstanceInfo {
	s.VersionId = &v
	return s
}

type InstanceLifecycleConfig struct {
	Initializer *LifecycleHook `json:"initializer,omitempty" xml:"initializer,omitempty"`
	PreStop     *LifecycleHook `json:"preStop,omitempty" xml:"preStop,omitempty"`
}

func (s InstanceLifecycleConfig) String() string {
	return tea.Prettify(s)
}

func (s InstanceLifecycleConfig) GoString() string {
	return s.String()
}

func (s *InstanceLifecycleConfig) SetInitializer(v *LifecycleHook) *InstanceLifecycleConfig {
	s.Initializer = v
	return s
}

func (s *InstanceLifecycleConfig) SetPreStop(v *LifecycleHook) *InstanceLifecycleConfig {
	s.PreStop = v
	return s
}

type JobConfig struct {
	// example:
	//
	// 3
	MaxRetryTime *int32 `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	// example:
	//
	// 60
	TriggerInterval *int32 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s JobConfig) String() string {
	return tea.Prettify(s)
}

func (s JobConfig) GoString() string {
	return s.String()
}

func (s *JobConfig) SetMaxRetryTime(v int32) *JobConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *JobConfig) SetTriggerInterval(v int32) *JobConfig {
	s.TriggerInterval = &v
	return s
}

type Key struct {
	// example:
	//
	// serverless_
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// example:
	//
	// .zip
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
}

func (s Key) String() string {
	return tea.Prettify(s)
}

func (s Key) GoString() string {
	return s.String()
}

func (s *Key) SetPrefix(v string) *Key {
	s.Prefix = &v
	return s
}

func (s *Key) SetSuffix(v string) *Key {
	s.Suffix = &v
	return s
}

type Layer struct {
	// example:
	//
	// 0
	Acl  *string             `json:"acl,omitempty" xml:"acl,omitempty"`
	Code *OutputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 2825179536350****
	CodeChecksum *string `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	// example:
	//
	// 421
	CodeSize          *int64    `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	CompatibleRuntime []*string `json:"compatibleRuntime" xml:"compatibleRuntime" type:"Repeated"`
	// example:
	//
	// 2023-03-30T11:08:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// My first layer
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyLayer
	LayerName *string `json:"layerName,omitempty" xml:"layerName,omitempty"`
	// example:
	//
	// acs:fc:cn-beijing:186824xxxxxx:layers/fc_layer/versions/1
	LayerVersionArn *string `json:"layerVersionArn,omitempty" xml:"layerVersionArn,omitempty"`
	// example:
	//
	// Apache
	License *string `json:"license,omitempty" xml:"license,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s Layer) String() string {
	return tea.Prettify(s)
}

func (s Layer) GoString() string {
	return s.String()
}

func (s *Layer) SetAcl(v string) *Layer {
	s.Acl = &v
	return s
}

func (s *Layer) SetCode(v *OutputCodeLocation) *Layer {
	s.Code = v
	return s
}

func (s *Layer) SetCodeChecksum(v string) *Layer {
	s.CodeChecksum = &v
	return s
}

func (s *Layer) SetCodeSize(v int64) *Layer {
	s.CodeSize = &v
	return s
}

func (s *Layer) SetCompatibleRuntime(v []*string) *Layer {
	s.CompatibleRuntime = v
	return s
}

func (s *Layer) SetCreateTime(v string) *Layer {
	s.CreateTime = &v
	return s
}

func (s *Layer) SetDescription(v string) *Layer {
	s.Description = &v
	return s
}

func (s *Layer) SetLayerName(v string) *Layer {
	s.LayerName = &v
	return s
}

func (s *Layer) SetLayerVersionArn(v string) *Layer {
	s.LayerVersionArn = &v
	return s
}

func (s *Layer) SetLicense(v string) *Layer {
	s.License = &v
	return s
}

func (s *Layer) SetVersion(v int32) *Layer {
	s.Version = &v
	return s
}

type LifecycleHook struct {
	// example:
	//
	// index.initializer
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s LifecycleHook) String() string {
	return tea.Prettify(s)
}

func (s LifecycleHook) GoString() string {
	return s.String()
}

func (s *LifecycleHook) SetHandler(v string) *LifecycleHook {
	s.Handler = &v
	return s
}

func (s *LifecycleHook) SetTimeout(v int32) *LifecycleHook {
	s.Timeout = &v
	return s
}

type ListAliasesOutput struct {
	Aliases []*Alias `json:"aliases" xml:"aliases" type:"Repeated"`
	// example:
	//
	// test
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAliasesOutput) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesOutput) GoString() string {
	return s.String()
}

func (s *ListAliasesOutput) SetAliases(v []*Alias) *ListAliasesOutput {
	s.Aliases = v
	return s
}

func (s *ListAliasesOutput) SetNextToken(v string) *ListAliasesOutput {
	s.NextToken = &v
	return s
}

type ListAsyncInvokeConfigOutput struct {
	Configs []*AsyncConfig `json:"configs" xml:"configs" type:"Repeated"`
	// example:
	//
	// 8bj81uI8n****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAsyncInvokeConfigOutput) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncInvokeConfigOutput) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigOutput) SetConfigs(v []*AsyncConfig) *ListAsyncInvokeConfigOutput {
	s.Configs = v
	return s
}

func (s *ListAsyncInvokeConfigOutput) SetNextToken(v string) *ListAsyncInvokeConfigOutput {
	s.NextToken = &v
	return s
}

type ListAsyncTaskOutput struct {
	NextToken *string      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Tasks     []*AsyncTask `json:"tasks" xml:"tasks" type:"Repeated"`
}

func (s ListAsyncTaskOutput) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskOutput) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskOutput) SetNextToken(v string) *ListAsyncTaskOutput {
	s.NextToken = &v
	return s
}

func (s *ListAsyncTaskOutput) SetTasks(v []*AsyncTask) *ListAsyncTaskOutput {
	s.Tasks = v
	return s
}

type ListConcurrencyConfigsOutput struct {
	Configs []*ConcurrencyConfig `json:"configs" xml:"configs" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConcurrencyConfigsOutput) String() string {
	return tea.Prettify(s)
}

func (s ListConcurrencyConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsOutput) SetConfigs(v []*ConcurrencyConfig) *ListConcurrencyConfigsOutput {
	s.Configs = v
	return s
}

func (s *ListConcurrencyConfigsOutput) SetNextToken(v string) *ListConcurrencyConfigsOutput {
	s.NextToken = &v
	return s
}

type ListCustomDomainOutput struct {
	CustomDomains []*CustomDomain `json:"customDomains" xml:"customDomains" type:"Repeated"`
	// example:
	//
	// next_domain_name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCustomDomainOutput) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainOutput) GoString() string {
	return s.String()
}

func (s *ListCustomDomainOutput) SetCustomDomains(v []*CustomDomain) *ListCustomDomainOutput {
	s.CustomDomains = v
	return s
}

func (s *ListCustomDomainOutput) SetNextToken(v string) *ListCustomDomainOutput {
	s.NextToken = &v
	return s
}

type ListFunctionsOutput struct {
	Functions []*Function `json:"functions" xml:"functions" type:"Repeated"`
	// example:
	//
	// next_function_name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionsOutput) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsOutput) GoString() string {
	return s.String()
}

func (s *ListFunctionsOutput) SetFunctions(v []*Function) *ListFunctionsOutput {
	s.Functions = v
	return s
}

func (s *ListFunctionsOutput) SetNextToken(v string) *ListFunctionsOutput {
	s.NextToken = &v
	return s
}

type ListInstancesOutput struct {
	Instances []*InstanceInfo `json:"instances" xml:"instances" type:"Repeated"`
}

func (s ListInstancesOutput) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOutput) GoString() string {
	return s.String()
}

func (s *ListInstancesOutput) SetInstances(v []*InstanceInfo) *ListInstancesOutput {
	s.Instances = v
	return s
}

type ListLayerVersionOutput struct {
	Layers []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	// example:
	//
	// 10
	NextVersion *int32 `json:"nextVersion,omitempty" xml:"nextVersion,omitempty"`
}

func (s ListLayerVersionOutput) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionOutput) GoString() string {
	return s.String()
}

func (s *ListLayerVersionOutput) SetLayers(v []*Layer) *ListLayerVersionOutput {
	s.Layers = v
	return s
}

func (s *ListLayerVersionOutput) SetNextVersion(v int32) *ListLayerVersionOutput {
	s.NextVersion = &v
	return s
}

type ListLayersOutput struct {
	Layers []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	// example:
	//
	// next-layer-name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListLayersOutput) String() string {
	return tea.Prettify(s)
}

func (s ListLayersOutput) GoString() string {
	return s.String()
}

func (s *ListLayersOutput) SetLayers(v []*Layer) *ListLayersOutput {
	s.Layers = v
	return s
}

func (s *ListLayersOutput) SetNextToken(v string) *ListLayersOutput {
	s.NextToken = &v
	return s
}

type ListProvisionConfigsOutput struct {
	// example:
	//
	// next_token
	NextToken        *string            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProvisionConfigs []*ProvisionConfig `json:"provisionConfigs" xml:"provisionConfigs" type:"Repeated"`
}

func (s ListProvisionConfigsOutput) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsOutput) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsOutput) SetNextToken(v string) *ListProvisionConfigsOutput {
	s.NextToken = &v
	return s
}

func (s *ListProvisionConfigsOutput) SetProvisionConfigs(v []*ProvisionConfig) *ListProvisionConfigsOutput {
	s.ProvisionConfigs = v
	return s
}

type ListTagResourcesOutput struct {
	// example:
	//
	// next_token
	NextToken    *string        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*TagResource `json:"TagResources" xml:"TagResources" type:"Repeated"`
}

func (s ListTagResourcesOutput) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesOutput) GoString() string {
	return s.String()
}

func (s *ListTagResourcesOutput) SetNextToken(v string) *ListTagResourcesOutput {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesOutput) SetRequestId(v string) *ListTagResourcesOutput {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesOutput) SetTagResources(v []*TagResource) *ListTagResourcesOutput {
	s.TagResources = v
	return s
}

type ListTaggedResourcesOutput struct {
	// example:
	//
	// next_token
	NextToken *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Resources []*Resource `json:"resources" xml:"resources" type:"Repeated"`
}

func (s ListTaggedResourcesOutput) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesOutput) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesOutput) SetNextToken(v string) *ListTaggedResourcesOutput {
	s.NextToken = &v
	return s
}

func (s *ListTaggedResourcesOutput) SetResources(v []*Resource) *ListTaggedResourcesOutput {
	s.Resources = v
	return s
}

type ListTriggersOutput struct {
	// example:
	//
	// next_token
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Triggers  []*Trigger `json:"triggers" xml:"triggers" type:"Repeated"`
}

func (s ListTriggersOutput) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersOutput) GoString() string {
	return s.String()
}

func (s *ListTriggersOutput) SetNextToken(v string) *ListTriggersOutput {
	s.NextToken = &v
	return s
}

func (s *ListTriggersOutput) SetTriggers(v []*Trigger) *ListTriggersOutput {
	s.Triggers = v
	return s
}

type ListVersionsOutput struct {
	// example:
	//
	// FORWARD
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// 3
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Versions  []*Version `json:"versions" xml:"versions" type:"Repeated"`
}

func (s ListVersionsOutput) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsOutput) GoString() string {
	return s.String()
}

func (s *ListVersionsOutput) SetDirection(v string) *ListVersionsOutput {
	s.Direction = &v
	return s
}

func (s *ListVersionsOutput) SetNextToken(v string) *ListVersionsOutput {
	s.NextToken = &v
	return s
}

func (s *ListVersionsOutput) SetVersions(v []*Version) *ListVersionsOutput {
	s.Versions = v
	return s
}

type ListVpcBindingsOutput struct {
	VpcIds []*string `json:"vpcIds" xml:"vpcIds" type:"Repeated"`
}

func (s ListVpcBindingsOutput) String() string {
	return tea.Prettify(s)
}

func (s ListVpcBindingsOutput) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsOutput) SetVpcIds(v []*string) *ListVpcBindingsOutput {
	s.VpcIds = v
	return s
}

type LogConfig struct {
	// example:
	//
	// true
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// example:
	//
	// true
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s LogConfig) String() string {
	return tea.Prettify(s)
}

func (s LogConfig) GoString() string {
	return s.String()
}

func (s *LogConfig) SetEnableInstanceMetrics(v bool) *LogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *LogConfig) SetEnableRequestMetrics(v bool) *LogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *LogConfig) SetLogBeginRule(v string) *LogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *LogConfig) SetLogstore(v string) *LogConfig {
	s.Logstore = &v
	return s
}

func (s *LogConfig) SetProject(v string) *LogConfig {
	s.Project = &v
	return s
}

type MNSTopicTriggerConfig struct {
	// example:
	//
	// serverless
	FilterTag *string `json:"filterTag,omitempty" xml:"filterTag,omitempty"`
	// example:
	//
	// JSON
	NotifyContentFormat *string `json:"notifyContentFormat,omitempty" xml:"notifyContentFormat,omitempty"`
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"notifyStrategy,omitempty" xml:"notifyStrategy,omitempty"`
}

func (s MNSTopicTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s MNSTopicTriggerConfig) GoString() string {
	return s.String()
}

func (s *MNSTopicTriggerConfig) SetFilterTag(v string) *MNSTopicTriggerConfig {
	s.FilterTag = &v
	return s
}

func (s *MNSTopicTriggerConfig) SetNotifyContentFormat(v string) *MNSTopicTriggerConfig {
	s.NotifyContentFormat = &v
	return s
}

func (s *MNSTopicTriggerConfig) SetNotifyStrategy(v string) *MNSTopicTriggerConfig {
	s.NotifyStrategy = &v
	return s
}

type NASConfig struct {
	// example:
	//
	// 100
	GroupId     *int32            `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*NASMountConfig `json:"mountPoints" xml:"mountPoints" type:"Repeated"`
	// example:
	//
	// 100
	UserId *int32 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s NASConfig) String() string {
	return tea.Prettify(s)
}

func (s NASConfig) GoString() string {
	return s.String()
}

func (s *NASConfig) SetGroupId(v int32) *NASConfig {
	s.GroupId = &v
	return s
}

func (s *NASConfig) SetMountPoints(v []*NASMountConfig) *NASConfig {
	s.MountPoints = v
	return s
}

func (s *NASConfig) SetUserId(v int32) *NASConfig {
	s.UserId = &v
	return s
}

type NASMountConfig struct {
	EnableTLS *bool `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	// example:
	//
	// /home/test
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// ***-uni85.cn-hangzhou.nas.com:/
	ServerAddr *string `json:"serverAddr,omitempty" xml:"serverAddr,omitempty"`
}

func (s NASMountConfig) String() string {
	return tea.Prettify(s)
}

func (s NASMountConfig) GoString() string {
	return s.String()
}

func (s *NASMountConfig) SetEnableTLS(v bool) *NASMountConfig {
	s.EnableTLS = &v
	return s
}

func (s *NASMountConfig) SetMountDir(v string) *NASMountConfig {
	s.MountDir = &v
	return s
}

func (s *NASMountConfig) SetServerAddr(v string) *NASMountConfig {
	s.ServerAddr = &v
	return s
}

type OSSMountConfig struct {
	MountPoints []*OSSMountPoint `json:"mountPoints" xml:"mountPoints" type:"Repeated"`
}

func (s OSSMountConfig) String() string {
	return tea.Prettify(s)
}

func (s OSSMountConfig) GoString() string {
	return s.String()
}

func (s *OSSMountConfig) SetMountPoints(v []*OSSMountPoint) *OSSMountConfig {
	s.MountPoints = v
	return s
}

type OSSMountPoint struct {
	// example:
	//
	// my-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// /my-dir
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	// example:
	//
	// http://oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// /mnt/dir
	MountDir *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	// example:
	//
	// true
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
}

func (s OSSMountPoint) String() string {
	return tea.Prettify(s)
}

func (s OSSMountPoint) GoString() string {
	return s.String()
}

func (s *OSSMountPoint) SetBucketName(v string) *OSSMountPoint {
	s.BucketName = &v
	return s
}

func (s *OSSMountPoint) SetBucketPath(v string) *OSSMountPoint {
	s.BucketPath = &v
	return s
}

func (s *OSSMountPoint) SetEndpoint(v string) *OSSMountPoint {
	s.Endpoint = &v
	return s
}

func (s *OSSMountPoint) SetMountDir(v string) *OSSMountPoint {
	s.MountDir = &v
	return s
}

func (s *OSSMountPoint) SetReadOnly(v bool) *OSSMountPoint {
	s.ReadOnly = &v
	return s
}

type OSSTriggerConfig struct {
	Events []*string `json:"events" xml:"events" type:"Repeated"`
	Filter *Filter   `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s OSSTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s OSSTriggerConfig) GoString() string {
	return s.String()
}

func (s *OSSTriggerConfig) SetEvents(v []*string) *OSSTriggerConfig {
	s.Events = v
	return s
}

func (s *OSSTriggerConfig) SetFilter(v *Filter) *OSSTriggerConfig {
	s.Filter = v
	return s
}

type OutputCodeLocation struct {
	// example:
	//
	// https://xyz.oss-cn-shanghai.aliyuncs.com/xxx/xxx/xxx
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// OSS
	RepositoryType *string `json:"repositoryType,omitempty" xml:"repositoryType,omitempty"`
}

func (s OutputCodeLocation) String() string {
	return tea.Prettify(s)
}

func (s OutputCodeLocation) GoString() string {
	return s.String()
}

func (s *OutputCodeLocation) SetLocation(v string) *OutputCodeLocation {
	s.Location = &v
	return s
}

func (s *OutputCodeLocation) SetRepositoryType(v string) *OutputCodeLocation {
	s.RepositoryType = &v
	return s
}

type OutputFuncCode struct {
	// example:
	//
	// 1234567890
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// example:
	//
	// http://func-code.oss-cn-shanghai.aliyuncs.com/1a2b3c4d5e6f
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s OutputFuncCode) String() string {
	return tea.Prettify(s)
}

func (s OutputFuncCode) GoString() string {
	return s.String()
}

func (s *OutputFuncCode) SetChecksum(v string) *OutputFuncCode {
	s.Checksum = &v
	return s
}

func (s *OutputFuncCode) SetUrl(v string) *OutputFuncCode {
	s.Url = &v
	return s
}

type PathConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// myFunction
	FunctionName *string   `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Methods      []*string `json:"methods" xml:"methods" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// /api/*
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// myAlias
	Qualifier     *string        `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	RewriteConfig *RewriteConfig `json:"rewriteConfig,omitempty" xml:"rewriteConfig,omitempty"`
}

func (s PathConfig) String() string {
	return tea.Prettify(s)
}

func (s PathConfig) GoString() string {
	return s.String()
}

func (s *PathConfig) SetFunctionName(v string) *PathConfig {
	s.FunctionName = &v
	return s
}

func (s *PathConfig) SetMethods(v []*string) *PathConfig {
	s.Methods = v
	return s
}

func (s *PathConfig) SetPath(v string) *PathConfig {
	s.Path = &v
	return s
}

func (s *PathConfig) SetQualifier(v string) *PathConfig {
	s.Qualifier = &v
	return s
}

func (s *PathConfig) SetRewriteConfig(v *RewriteConfig) *PathConfig {
	s.RewriteConfig = v
	return s
}

type ProvisionConfig struct {
	// example:
	//
	// true
	AlwaysAllocateCPU *bool `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	// example:
	//
	// true
	AlwaysAllocateGPU *bool `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// image not found
	CurrentError *string `json:"currentError,omitempty" xml:"currentError,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:124:functions/myFunction
	FunctionArn      *string            `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	ScheduledActions []*ScheduledAction `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
	// example:
	//
	// 5
	Target                 *int64                  `json:"target,omitempty" xml:"target,omitempty"`
	TargetTrackingPolicies []*TargetTrackingPolicy `json:"targetTrackingPolicies" xml:"targetTrackingPolicies" type:"Repeated"`
}

func (s ProvisionConfig) String() string {
	return tea.Prettify(s)
}

func (s ProvisionConfig) GoString() string {
	return s.String()
}

func (s *ProvisionConfig) SetAlwaysAllocateCPU(v bool) *ProvisionConfig {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *ProvisionConfig) SetAlwaysAllocateGPU(v bool) *ProvisionConfig {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *ProvisionConfig) SetCurrent(v int64) *ProvisionConfig {
	s.Current = &v
	return s
}

func (s *ProvisionConfig) SetCurrentError(v string) *ProvisionConfig {
	s.CurrentError = &v
	return s
}

func (s *ProvisionConfig) SetFunctionArn(v string) *ProvisionConfig {
	s.FunctionArn = &v
	return s
}

func (s *ProvisionConfig) SetScheduledActions(v []*ScheduledAction) *ProvisionConfig {
	s.ScheduledActions = v
	return s
}

func (s *ProvisionConfig) SetTarget(v int64) *ProvisionConfig {
	s.Target = &v
	return s
}

func (s *ProvisionConfig) SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *ProvisionConfig {
	s.TargetTrackingPolicies = v
	return s
}

type PublishVersionInput struct {
	// example:
	//
	// my version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s PublishVersionInput) String() string {
	return tea.Prettify(s)
}

func (s PublishVersionInput) GoString() string {
	return s.String()
}

func (s *PublishVersionInput) SetDescription(v string) *PublishVersionInput {
	s.Description = &v
	return s
}

type PutAsyncInvokeConfigInput struct {
	// example:
	//
	// true
	AsyncTask         *bool              `json:"asyncTask,omitempty" xml:"asyncTask,omitempty"`
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	// example:
	//
	// 300
	MaxAsyncEventAgeInSeconds *int64 `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	// example:
	//
	// 3
	MaxAsyncRetryAttempts *int64 `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s PutAsyncInvokeConfigInput) String() string {
	return tea.Prettify(s)
}

func (s PutAsyncInvokeConfigInput) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigInput) SetAsyncTask(v bool) *PutAsyncInvokeConfigInput {
	s.AsyncTask = &v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetDestinationConfig(v *DestinationConfig) *PutAsyncInvokeConfigInput {
	s.DestinationConfig = v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetMaxAsyncEventAgeInSeconds(v int64) *PutAsyncInvokeConfigInput {
	s.MaxAsyncEventAgeInSeconds = &v
	return s
}

func (s *PutAsyncInvokeConfigInput) SetMaxAsyncRetryAttempts(v int64) *PutAsyncInvokeConfigInput {
	s.MaxAsyncRetryAttempts = &v
	return s
}

type PutConcurrencyInput struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	ReservedConcurrency *int64 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s PutConcurrencyInput) String() string {
	return tea.Prettify(s)
}

func (s PutConcurrencyInput) GoString() string {
	return s.String()
}

func (s *PutConcurrencyInput) SetReservedConcurrency(v int64) *PutConcurrencyInput {
	s.ReservedConcurrency = &v
	return s
}

type PutProvisionConfigInput struct {
	// example:
	//
	// true
	AlwaysAllocateCPU *bool `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	// example:
	//
	// true
	AlwaysAllocateGPU *bool              `json:"alwaysAllocateGPU,omitempty" xml:"alwaysAllocateGPU,omitempty"`
	ScheduledActions  []*ScheduledAction `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Target                 *int64                  `json:"target,omitempty" xml:"target,omitempty"`
	TargetTrackingPolicies []*TargetTrackingPolicy `json:"targetTrackingPolicies" xml:"targetTrackingPolicies" type:"Repeated"`
}

func (s PutProvisionConfigInput) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigInput) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigInput) SetAlwaysAllocateCPU(v bool) *PutProvisionConfigInput {
	s.AlwaysAllocateCPU = &v
	return s
}

func (s *PutProvisionConfigInput) SetAlwaysAllocateGPU(v bool) *PutProvisionConfigInput {
	s.AlwaysAllocateGPU = &v
	return s
}

func (s *PutProvisionConfigInput) SetScheduledActions(v []*ScheduledAction) *PutProvisionConfigInput {
	s.ScheduledActions = v
	return s
}

func (s *PutProvisionConfigInput) SetTarget(v int64) *PutProvisionConfigInput {
	s.Target = &v
	return s
}

func (s *PutProvisionConfigInput) SetTargetTrackingPolicies(v []*TargetTrackingPolicy) *PutProvisionConfigInput {
	s.TargetTrackingPolicies = v
	return s
}

type RegexRule struct {
	// This parameter is required.
	//
	// example:
	//
	// ^/api/.+?/(.*)
	Match *string `json:"match,omitempty" xml:"match,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /api/v2/$1
	Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s RegexRule) String() string {
	return tea.Prettify(s)
}

func (s RegexRule) GoString() string {
	return s.String()
}

func (s *RegexRule) SetMatch(v string) *RegexRule {
	s.Match = &v
	return s
}

func (s *RegexRule) SetReplacement(v string) *RegexRule {
	s.Replacement = &v
	return s
}

type RegistryAuthConfig struct {
	// example:
	//
	// abc***
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// example:
	//
	// admin
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s RegistryAuthConfig) String() string {
	return tea.Prettify(s)
}

func (s RegistryAuthConfig) GoString() string {
	return s.String()
}

func (s *RegistryAuthConfig) SetPassword(v string) *RegistryAuthConfig {
	s.Password = &v
	return s
}

func (s *RegistryAuthConfig) SetUserName(v string) *RegistryAuthConfig {
	s.UserName = &v
	return s
}

type RegistryCertConfig struct {
	Insecure *bool `json:"insecure,omitempty" xml:"insecure,omitempty"`
	// example:
	//
	// cm9vdF9jYV9jZXJ0
	RootCaCertBase64 *string `json:"rootCaCertBase64,omitempty" xml:"rootCaCertBase64,omitempty"`
}

func (s RegistryCertConfig) String() string {
	return tea.Prettify(s)
}

func (s RegistryCertConfig) GoString() string {
	return s.String()
}

func (s *RegistryCertConfig) SetInsecure(v bool) *RegistryCertConfig {
	s.Insecure = &v
	return s
}

func (s *RegistryCertConfig) SetRootCaCertBase64(v string) *RegistryCertConfig {
	s.RootCaCertBase64 = &v
	return s
}

type RegistryConfig struct {
	AuthConfig    *RegistryAuthConfig    `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig    *RegistryCertConfig    `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	NetworkConfig *RegistryNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty"`
}

func (s RegistryConfig) String() string {
	return tea.Prettify(s)
}

func (s RegistryConfig) GoString() string {
	return s.String()
}

func (s *RegistryConfig) SetAuthConfig(v *RegistryAuthConfig) *RegistryConfig {
	s.AuthConfig = v
	return s
}

func (s *RegistryConfig) SetCertConfig(v *RegistryCertConfig) *RegistryConfig {
	s.CertConfig = v
	return s
}

func (s *RegistryConfig) SetNetworkConfig(v *RegistryNetworkConfig) *RegistryConfig {
	s.NetworkConfig = v
	return s
}

type RegistryNetworkConfig struct {
	// example:
	//
	// sg-xxxxxxxxxxxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vsw-xxxxxxxxxxxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// vpc-xxxxxxxxxxxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s RegistryNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s RegistryNetworkConfig) GoString() string {
	return s.String()
}

func (s *RegistryNetworkConfig) SetSecurityGroupId(v string) *RegistryNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *RegistryNetworkConfig) SetVSwitchId(v string) *RegistryNetworkConfig {
	s.VSwitchId = &v
	return s
}

func (s *RegistryNetworkConfig) SetVpcId(v string) *RegistryNetworkConfig {
	s.VpcId = &v
	return s
}

type Resource struct {
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResouceType *string `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
	// example:
	//
	// acs:fc:cn-shanghai:****:functions/demo
	ResourceArn *string            `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	Tags        map[string]*string `json:"tags" xml:"tags"`
}

func (s Resource) String() string {
	return tea.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) SetResouceType(v string) *Resource {
	s.ResouceType = &v
	return s
}

func (s *Resource) SetResourceArn(v string) *Resource {
	s.ResourceArn = &v
	return s
}

func (s *Resource) SetTags(v map[string]*string) *Resource {
	s.Tags = v
	return s
}

type RetryStrategy struct {
	// example:
	//
	// BACKOFF_RETRY
	PushRetryStrategy *string `json:"PushRetryStrategy,omitempty" xml:"PushRetryStrategy,omitempty"`
}

func (s RetryStrategy) String() string {
	return tea.Prettify(s)
}

func (s RetryStrategy) GoString() string {
	return s.String()
}

func (s *RetryStrategy) SetPushRetryStrategy(v string) *RetryStrategy {
	s.PushRetryStrategy = &v
	return s
}

type RewriteConfig struct {
	EqualRules    []*EqualRule    `json:"equalRules" xml:"equalRules" type:"Repeated"`
	RegexRules    []*RegexRule    `json:"regexRules" xml:"regexRules" type:"Repeated"`
	WildcardRules []*WildcardRule `json:"wildcardRules" xml:"wildcardRules" type:"Repeated"`
}

func (s RewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s RewriteConfig) GoString() string {
	return s.String()
}

func (s *RewriteConfig) SetEqualRules(v []*EqualRule) *RewriteConfig {
	s.EqualRules = v
	return s
}

func (s *RewriteConfig) SetRegexRules(v []*RegexRule) *RewriteConfig {
	s.RegexRules = v
	return s
}

func (s *RewriteConfig) SetWildcardRules(v []*WildcardRule) *RewriteConfig {
	s.WildcardRules = v
	return s
}

type RouteConfig struct {
	Routes []*PathConfig `json:"routes" xml:"routes" type:"Repeated"`
}

func (s RouteConfig) String() string {
	return tea.Prettify(s)
}

func (s RouteConfig) GoString() string {
	return s.String()
}

func (s *RouteConfig) SetRoutes(v []*PathConfig) *RouteConfig {
	s.Routes = v
	return s
}

type RunOptions struct {
	BatchWindow     *BatchWindow     `json:"batchWindow,omitempty" xml:"batchWindow,omitempty"`
	DeadLetterQueue *DeadLetterQueue `json:"deadLetterQueue,omitempty" xml:"deadLetterQueue,omitempty"`
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"errorsTolerance,omitempty" xml:"errorsTolerance,omitempty"`
	// example:
	//
	// event-streaming
	Mode          *string        `json:"mode,omitempty" xml:"mode,omitempty"`
	RetryStrategy *RetryStrategy `json:"retryStrategy,omitempty" xml:"retryStrategy,omitempty"`
}

func (s RunOptions) String() string {
	return tea.Prettify(s)
}

func (s RunOptions) GoString() string {
	return s.String()
}

func (s *RunOptions) SetBatchWindow(v *BatchWindow) *RunOptions {
	s.BatchWindow = v
	return s
}

func (s *RunOptions) SetDeadLetterQueue(v *DeadLetterQueue) *RunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *RunOptions) SetErrorsTolerance(v string) *RunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *RunOptions) SetMode(v string) *RunOptions {
	s.Mode = &v
	return s
}

func (s *RunOptions) SetRetryStrategy(v *RetryStrategy) *RunOptions {
	s.RetryStrategy = v
	return s
}

type SLSTriggerConfig struct {
	// example:
	//
	// true
	Enable            *bool                `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionParameter map[string]*string   `json:"functionParameter" xml:"functionParameter"`
	JobConfig         *JobConfig           `json:"jobConfig,omitempty" xml:"jobConfig,omitempty"`
	LogConfig         *SLSTriggerLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *SourceConfig        `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
}

func (s SLSTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s SLSTriggerConfig) GoString() string {
	return s.String()
}

func (s *SLSTriggerConfig) SetEnable(v bool) *SLSTriggerConfig {
	s.Enable = &v
	return s
}

func (s *SLSTriggerConfig) SetFunctionParameter(v map[string]*string) *SLSTriggerConfig {
	s.FunctionParameter = v
	return s
}

func (s *SLSTriggerConfig) SetJobConfig(v *JobConfig) *SLSTriggerConfig {
	s.JobConfig = v
	return s
}

func (s *SLSTriggerConfig) SetLogConfig(v *SLSTriggerLogConfig) *SLSTriggerConfig {
	s.LogConfig = v
	return s
}

func (s *SLSTriggerConfig) SetSourceConfig(v *SourceConfig) *SLSTriggerConfig {
	s.SourceConfig = v
	return s
}

type SLSTriggerLogConfig struct {
	// example:
	//
	// my-sls-logstore-name
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// my-sls-project-name
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s SLSTriggerLogConfig) String() string {
	return tea.Prettify(s)
}

func (s SLSTriggerLogConfig) GoString() string {
	return s.String()
}

func (s *SLSTriggerLogConfig) SetLogstore(v string) *SLSTriggerLogConfig {
	s.Logstore = &v
	return s
}

func (s *SLSTriggerLogConfig) SetProject(v string) *SLSTriggerLogConfig {
	s.Project = &v
	return s
}

type ScheduledAction struct {
	// example:
	//
	// 2024-03-10T10:10:10
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cron(0 0 22 	- 	- *)
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	// example:
	//
	// 2023-03-10T10:10:10
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScheduledAction) String() string {
	return tea.Prettify(s)
}

func (s ScheduledAction) GoString() string {
	return s.String()
}

func (s *ScheduledAction) SetEndTime(v string) *ScheduledAction {
	s.EndTime = &v
	return s
}

func (s *ScheduledAction) SetName(v string) *ScheduledAction {
	s.Name = &v
	return s
}

func (s *ScheduledAction) SetScheduleExpression(v string) *ScheduledAction {
	s.ScheduleExpression = &v
	return s
}

func (s *ScheduledAction) SetStartTime(v string) *ScheduledAction {
	s.StartTime = &v
	return s
}

func (s *ScheduledAction) SetTarget(v int64) *ScheduledAction {
	s.Target = &v
	return s
}

func (s *ScheduledAction) SetTimeZone(v string) *ScheduledAction {
	s.TimeZone = &v
	return s
}

type SourceConfig struct {
	// example:
	//
	// my-sls-logstore-name
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// 1704790317
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SourceConfig) String() string {
	return tea.Prettify(s)
}

func (s SourceConfig) GoString() string {
	return s.String()
}

func (s *SourceConfig) SetLogstore(v string) *SourceConfig {
	s.Logstore = &v
	return s
}

func (s *SourceConfig) SetStartTime(v int64) *SourceConfig {
	s.StartTime = &v
	return s
}

type SourceDTSParameters struct {
	// example:
	//
	// dts-cn-shanghai-vpc.com:18003
	BrokerUrl *string `json:"BrokerUrl,omitempty" xml:"BrokerUrl,omitempty"`
	// example:
	//
	// 1677340805
	InitCheckPoint *int32 `json:"InitCheckPoint,omitempty" xml:"InitCheckPoint,omitempty"`
	// example:
	//
	// dtsTest123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// dtse34j22j025a****
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// example:
	//
	// e34z2gm325q****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// cn_shanghai_vpc_rm_uf6398ykj0218****_dts_trigger_upgrade_from_old_version2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// dts_trigger
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s SourceDTSParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceDTSParameters) GoString() string {
	return s.String()
}

func (s *SourceDTSParameters) SetBrokerUrl(v string) *SourceDTSParameters {
	s.BrokerUrl = &v
	return s
}

func (s *SourceDTSParameters) SetInitCheckPoint(v int32) *SourceDTSParameters {
	s.InitCheckPoint = &v
	return s
}

func (s *SourceDTSParameters) SetPassword(v string) *SourceDTSParameters {
	s.Password = &v
	return s
}

func (s *SourceDTSParameters) SetRegionId(v string) *SourceDTSParameters {
	s.RegionId = &v
	return s
}

func (s *SourceDTSParameters) SetSid(v string) *SourceDTSParameters {
	s.Sid = &v
	return s
}

func (s *SourceDTSParameters) SetTaskId(v string) *SourceDTSParameters {
	s.TaskId = &v
	return s
}

func (s *SourceDTSParameters) SetTopic(v string) *SourceDTSParameters {
	s.Topic = &v
	return s
}

func (s *SourceDTSParameters) SetUsername(v string) *SourceDTSParameters {
	s.Username = &v
	return s
}

type SourceKafkaParameters struct {
	// example:
	//
	// DEFAULT_GROUP
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// example:
	//
	// r-8vb64581862c****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-bp1iv19sp1msc7zot4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// popvip_center_robot_order
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// vsw-bp179l3llg3jjxwrq72****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-8vblalsi0vbhizr77****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SourceKafkaParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *SourceKafkaParameters) SetConsumerGroup(v string) *SourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *SourceKafkaParameters) SetInstanceId(v string) *SourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceKafkaParameters) SetNetwork(v string) *SourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *SourceKafkaParameters) SetOffsetReset(v string) *SourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *SourceKafkaParameters) SetRegionId(v string) *SourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *SourceKafkaParameters) SetSecurityGroupId(v string) *SourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SourceKafkaParameters) SetTopic(v string) *SourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *SourceKafkaParameters) SetVSwitchIds(v string) *SourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SourceKafkaParameters) SetVpcId(v string) *SourceKafkaParameters {
	s.VpcId = &v
	return s
}

type SourceMNSParameters struct {
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SourceMNSParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceMNSParameters) GoString() string {
	return s.String()
}

func (s *SourceMNSParameters) SetIsBase64Decode(v bool) *SourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *SourceMNSParameters) SetQueueName(v string) *SourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *SourceMNSParameters) SetRegionId(v string) *SourceMNSParameters {
	s.RegionId = &v
	return s
}

type SourceMQTTParameters struct {
	// example:
	//
	// mqtt-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// testTopic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SourceMQTTParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceMQTTParameters) GoString() string {
	return s.String()
}

func (s *SourceMQTTParameters) SetInstanceId(v string) *SourceMQTTParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceMQTTParameters) SetRegionId(v string) *SourceMQTTParameters {
	s.RegionId = &v
	return s
}

func (s *SourceMQTTParameters) SetTopic(v string) *SourceMQTTParameters {
	s.Topic = &v
	return s
}

type SourceRabbitMQParameters struct {
	// example:
	//
	// amqp-cn-nif22u74****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// demo
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s SourceRabbitMQParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *SourceRabbitMQParameters) SetInstanceId(v string) *SourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRabbitMQParameters) SetQueueName(v string) *SourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *SourceRabbitMQParameters) SetRegionId(v string) *SourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *SourceRabbitMQParameters) SetVirtualHostName(v string) *SourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

type SourceRocketMQParameters struct {
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// Tag
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// example:
	//
	// GID_group1
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// PrivateNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// example:
	//
	// 123
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// example:
	//
	// sg-hp35r2hc3a3sv8q2****
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// example:
	//
	// Cloud_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 6W0xz2uPfiwp****
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// example:
	//
	// vpc-uf6of9452b2pba82c****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// example:
	//
	// CONSUME_FROM_TIMESTAMP
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 1636597951964
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// myTopic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s SourceRocketMQParameters) String() string {
	return tea.Prettify(s)
}

func (s SourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *SourceRocketMQParameters) SetAuthType(v string) *SourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *SourceRocketMQParameters) SetFilterType(v string) *SourceRocketMQParameters {
	s.FilterType = &v
	return s
}

func (s *SourceRocketMQParameters) SetGroupID(v string) *SourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceEndpoint(v string) *SourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceId(v string) *SourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceNetwork(v string) *SourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstancePassword(v string) *SourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *SourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceType(v string) *SourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceUsername(v string) *SourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceVSwitchIds(v string) *SourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *SourceRocketMQParameters) SetInstanceVpcId(v string) *SourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *SourceRocketMQParameters) SetOffset(v string) *SourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *SourceRocketMQParameters) SetRegionId(v string) *SourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *SourceRocketMQParameters) SetTag(v string) *SourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *SourceRocketMQParameters) SetTimestamp(v int32) *SourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *SourceRocketMQParameters) SetTopic(v string) *SourceRocketMQParameters {
	s.Topic = &v
	return s
}

type TLSConfig struct {
	// This parameter is required.
	CipherSuites []*string `json:"cipherSuites" xml:"cipherSuites" type:"Repeated"`
	// example:
	//
	// TLSv1.3
	MaxVersion *string `json:"maxVersion,omitempty" xml:"maxVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TLSv1.0
	MinVersion *string `json:"minVersion,omitempty" xml:"minVersion,omitempty"`
}

func (s TLSConfig) String() string {
	return tea.Prettify(s)
}

func (s TLSConfig) GoString() string {
	return s.String()
}

func (s *TLSConfig) SetCipherSuites(v []*string) *TLSConfig {
	s.CipherSuites = v
	return s
}

func (s *TLSConfig) SetMaxVersion(v string) *TLSConfig {
	s.MaxVersion = &v
	return s
}

func (s *TLSConfig) SetMinVersion(v string) *TLSConfig {
	s.MinVersion = &v
	return s
}

type Tag struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Tag) String() string {
	return tea.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

type TagResource struct {
	// example:
	//
	// acs:fc:cn-shanghai:****:functions/demo
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ALIYUN::FC::FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// key1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s TagResource) String() string {
	return tea.Prettify(s)
}

func (s TagResource) GoString() string {
	return s.String()
}

func (s *TagResource) SetResourceId(v string) *TagResource {
	s.ResourceId = &v
	return s
}

func (s *TagResource) SetResourceType(v string) *TagResource {
	s.ResourceType = &v
	return s
}

func (s *TagResource) SetTagKey(v string) *TagResource {
	s.TagKey = &v
	return s
}

func (s *TagResource) SetTagValue(v string) *TagResource {
	s.TagValue = &v
	return s
}

type TagResourceInput struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:fc:cn-shanghai:xxx:functions/f1
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// This parameter is required.
	Tags map[string]*string `json:"tags" xml:"tags"`
}

func (s TagResourceInput) String() string {
	return tea.Prettify(s)
}

func (s TagResourceInput) GoString() string {
	return s.String()
}

func (s *TagResourceInput) SetResourceArn(v string) *TagResourceInput {
	s.ResourceArn = &v
	return s
}

func (s *TagResourceInput) SetTags(v map[string]*string) *TagResourceInput {
	s.Tags = v
	return s
}

type TagResourcesInput struct {
	// This parameter is required.
	ResourceId []*string `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	// example:
	//
	// FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tag []*Tag `json:"Tag" xml:"Tag" type:"Repeated"`
}

func (s TagResourcesInput) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesInput) GoString() string {
	return s.String()
}

func (s *TagResourcesInput) SetResourceId(v []*string) *TagResourcesInput {
	s.ResourceId = v
	return s
}

func (s *TagResourcesInput) SetResourceType(v string) *TagResourcesInput {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesInput) SetTag(v []*Tag) *TagResourcesInput {
	s.Tag = v
	return s
}

type TargetTrackingPolicy struct {
	// example:
	//
	// 2024-03-10T10:10:10
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxCapacity *int64 `json:"maxCapacity,omitempty" xml:"maxCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.6
	MetricTarget *float32 `json:"metricTarget,omitempty" xml:"metricTarget,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CPUUtilization
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinCapacity *int64 `json:"minCapacity,omitempty" xml:"minCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2023-03-10T10:10:10
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s TargetTrackingPolicy) String() string {
	return tea.Prettify(s)
}

func (s TargetTrackingPolicy) GoString() string {
	return s.String()
}

func (s *TargetTrackingPolicy) SetEndTime(v string) *TargetTrackingPolicy {
	s.EndTime = &v
	return s
}

func (s *TargetTrackingPolicy) SetMaxCapacity(v int64) *TargetTrackingPolicy {
	s.MaxCapacity = &v
	return s
}

func (s *TargetTrackingPolicy) SetMetricTarget(v float32) *TargetTrackingPolicy {
	s.MetricTarget = &v
	return s
}

func (s *TargetTrackingPolicy) SetMetricType(v string) *TargetTrackingPolicy {
	s.MetricType = &v
	return s
}

func (s *TargetTrackingPolicy) SetMinCapacity(v int64) *TargetTrackingPolicy {
	s.MinCapacity = &v
	return s
}

func (s *TargetTrackingPolicy) SetName(v string) *TargetTrackingPolicy {
	s.Name = &v
	return s
}

func (s *TargetTrackingPolicy) SetStartTime(v string) *TargetTrackingPolicy {
	s.StartTime = &v
	return s
}

func (s *TargetTrackingPolicy) SetTimeZone(v string) *TargetTrackingPolicy {
	s.TimeZone = &v
	return s
}

type TimerTriggerConfig struct {
	// example:
	//
	// 0 0 4 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// {"workflowInstanceId":"39639"}
	Payload *string `json:"payload,omitempty" xml:"payload,omitempty"`
}

func (s TimerTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s TimerTriggerConfig) GoString() string {
	return s.String()
}

func (s *TimerTriggerConfig) SetCronExpression(v string) *TimerTriggerConfig {
	s.CronExpression = &v
	return s
}

func (s *TimerTriggerConfig) SetEnable(v bool) *TimerTriggerConfig {
	s.Enable = &v
	return s
}

func (s *TimerTriggerConfig) SetPayload(v string) *TimerTriggerConfig {
	s.Payload = &v
	return s
}

type TracingConfig struct {
	Params map[string]*string `json:"params" xml:"params"`
	// example:
	//
	// Jaeger
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TracingConfig) String() string {
	return tea.Prettify(s)
}

func (s TracingConfig) GoString() string {
	return s.String()
}

func (s *TracingConfig) SetParams(v map[string]*string) *TracingConfig {
	s.Params = v
	return s
}

func (s *TracingConfig) SetType(v string) *TracingConfig {
	s.Type = &v
	return s
}

type Trigger struct {
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description      *string      `json:"description,omitempty" xml:"description,omitempty"`
	HttpTrigger      *HTTPTrigger `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty"`
	InvocationRole   *string      `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Qualifier        *string      `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	SourceArn        *string      `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	Status           *string      `json:"status,omitempty" xml:"status,omitempty"`
	TargetArn        *string      `json:"targetArn,omitempty" xml:"targetArn,omitempty"`
	TriggerConfig    *string      `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId        *string      `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	TriggerName      *string      `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	TriggerType      *string      `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s Trigger) String() string {
	return tea.Prettify(s)
}

func (s Trigger) GoString() string {
	return s.String()
}

func (s *Trigger) SetCreatedTime(v string) *Trigger {
	s.CreatedTime = &v
	return s
}

func (s *Trigger) SetDescription(v string) *Trigger {
	s.Description = &v
	return s
}

func (s *Trigger) SetHttpTrigger(v *HTTPTrigger) *Trigger {
	s.HttpTrigger = v
	return s
}

func (s *Trigger) SetInvocationRole(v string) *Trigger {
	s.InvocationRole = &v
	return s
}

func (s *Trigger) SetLastModifiedTime(v string) *Trigger {
	s.LastModifiedTime = &v
	return s
}

func (s *Trigger) SetQualifier(v string) *Trigger {
	s.Qualifier = &v
	return s
}

func (s *Trigger) SetSourceArn(v string) *Trigger {
	s.SourceArn = &v
	return s
}

func (s *Trigger) SetStatus(v string) *Trigger {
	s.Status = &v
	return s
}

func (s *Trigger) SetTargetArn(v string) *Trigger {
	s.TargetArn = &v
	return s
}

func (s *Trigger) SetTriggerConfig(v string) *Trigger {
	s.TriggerConfig = &v
	return s
}

func (s *Trigger) SetTriggerId(v string) *Trigger {
	s.TriggerId = &v
	return s
}

func (s *Trigger) SetTriggerName(v string) *Trigger {
	s.TriggerName = &v
	return s
}

func (s *Trigger) SetTriggerType(v string) *Trigger {
	s.TriggerType = &v
	return s
}

type UpdateAliasInput struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight" xml:"additionalVersionWeight"`
	// example:
	//
	// my alias
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s UpdateAliasInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasInput) GoString() string {
	return s.String()
}

func (s *UpdateAliasInput) SetAdditionalVersionWeight(v map[string]*float32) *UpdateAliasInput {
	s.AdditionalVersionWeight = v
	return s
}

func (s *UpdateAliasInput) SetDescription(v string) *UpdateAliasInput {
	s.Description = &v
	return s
}

func (s *UpdateAliasInput) SetVersionId(v string) *UpdateAliasInput {
	s.VersionId = &v
	return s
}

type UpdateCustomDomainInput struct {
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// example:
	//
	// HTTP
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	TlsConfig   *TLSConfig   `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig   *WAFConfig   `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s UpdateCustomDomainInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainInput) SetAuthConfig(v *AuthConfig) *UpdateCustomDomainInput {
	s.AuthConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetCertConfig(v *CertConfig) *UpdateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetProtocol(v string) *UpdateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *UpdateCustomDomainInput) SetRouteConfig(v *RouteConfig) *UpdateCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetWafConfig(v *WAFConfig) *UpdateCustomDomainInput {
	s.WafConfig = v
	return s
}

type UpdateFunctionInput struct {
	Code *InputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Cpu                   *float32               `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig *CustomContainerConfig `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS             *CustomDNS             `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig   *CustomRuntimeConfig   `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	// example:
	//
	// my function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 512
	DiskSize             *int32             `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables map[string]*string `json:"environmentVariables" xml:"environmentVariables"`
	GpuConfig            *GPUConfig         `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	// example:
	//
	// index.handler
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 1
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	// example:
	//
	// true
	InternetAccess *bool      `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers         []*string  `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig      *LogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	// example:
	//
	// 512
	MemorySize     *int32          `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig      *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// example:
	//
	// 60
	Timeout       *int32         `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig *TracingConfig `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig     *VPCConfig     `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
}

func (s UpdateFunctionInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInput) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInput) SetCode(v *InputCodeLocation) *UpdateFunctionInput {
	s.Code = v
	return s
}

func (s *UpdateFunctionInput) SetCpu(v float32) *UpdateFunctionInput {
	s.Cpu = &v
	return s
}

func (s *UpdateFunctionInput) SetCustomContainerConfig(v *CustomContainerConfig) *UpdateFunctionInput {
	s.CustomContainerConfig = v
	return s
}

func (s *UpdateFunctionInput) SetCustomDNS(v *CustomDNS) *UpdateFunctionInput {
	s.CustomDNS = v
	return s
}

func (s *UpdateFunctionInput) SetCustomRuntimeConfig(v *CustomRuntimeConfig) *UpdateFunctionInput {
	s.CustomRuntimeConfig = v
	return s
}

func (s *UpdateFunctionInput) SetDescription(v string) *UpdateFunctionInput {
	s.Description = &v
	return s
}

func (s *UpdateFunctionInput) SetDiskSize(v int32) *UpdateFunctionInput {
	s.DiskSize = &v
	return s
}

func (s *UpdateFunctionInput) SetEnvironmentVariables(v map[string]*string) *UpdateFunctionInput {
	s.EnvironmentVariables = v
	return s
}

func (s *UpdateFunctionInput) SetGpuConfig(v *GPUConfig) *UpdateFunctionInput {
	s.GpuConfig = v
	return s
}

func (s *UpdateFunctionInput) SetHandler(v string) *UpdateFunctionInput {
	s.Handler = &v
	return s
}

func (s *UpdateFunctionInput) SetInstanceConcurrency(v int32) *UpdateFunctionInput {
	s.InstanceConcurrency = &v
	return s
}

func (s *UpdateFunctionInput) SetInstanceLifecycleConfig(v *InstanceLifecycleConfig) *UpdateFunctionInput {
	s.InstanceLifecycleConfig = v
	return s
}

func (s *UpdateFunctionInput) SetInternetAccess(v bool) *UpdateFunctionInput {
	s.InternetAccess = &v
	return s
}

func (s *UpdateFunctionInput) SetLayers(v []*string) *UpdateFunctionInput {
	s.Layers = v
	return s
}

func (s *UpdateFunctionInput) SetLogConfig(v *LogConfig) *UpdateFunctionInput {
	s.LogConfig = v
	return s
}

func (s *UpdateFunctionInput) SetMemorySize(v int32) *UpdateFunctionInput {
	s.MemorySize = &v
	return s
}

func (s *UpdateFunctionInput) SetNasConfig(v *NASConfig) *UpdateFunctionInput {
	s.NasConfig = v
	return s
}

func (s *UpdateFunctionInput) SetOssMountConfig(v *OSSMountConfig) *UpdateFunctionInput {
	s.OssMountConfig = v
	return s
}

func (s *UpdateFunctionInput) SetRole(v string) *UpdateFunctionInput {
	s.Role = &v
	return s
}

func (s *UpdateFunctionInput) SetRuntime(v string) *UpdateFunctionInput {
	s.Runtime = &v
	return s
}

func (s *UpdateFunctionInput) SetTimeout(v int32) *UpdateFunctionInput {
	s.Timeout = &v
	return s
}

func (s *UpdateFunctionInput) SetTracingConfig(v *TracingConfig) *UpdateFunctionInput {
	s.TracingConfig = v
	return s
}

func (s *UpdateFunctionInput) SetVpcConfig(v *VPCConfig) *UpdateFunctionInput {
	s.VpcConfig = v
	return s
}

type UpdateTriggerInput struct {
	// example:
	//
	// trigger for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// acs:ram::1234567890:role/fc-test
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// {"events":["oss:ObjectCreated:*"],"filter":{"key":{"prefix":"/prefix","suffix":".zip"}}}
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateTriggerInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerInput) GoString() string {
	return s.String()
}

func (s *UpdateTriggerInput) SetDescription(v string) *UpdateTriggerInput {
	s.Description = &v
	return s
}

func (s *UpdateTriggerInput) SetInvocationRole(v string) *UpdateTriggerInput {
	s.InvocationRole = &v
	return s
}

func (s *UpdateTriggerInput) SetQualifier(v string) *UpdateTriggerInput {
	s.Qualifier = &v
	return s
}

func (s *UpdateTriggerInput) SetTriggerConfig(v string) *UpdateTriggerInput {
	s.TriggerConfig = &v
	return s
}

type VPCConfig struct {
	// example:
	//
	// sg-bp18hj1wtxgy3b0***
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds" xml:"vSwitchIds" type:"Repeated"`
	// example:
	//
	// vpc-***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s VPCConfig) String() string {
	return tea.Prettify(s)
}

func (s VPCConfig) GoString() string {
	return s.String()
}

func (s *VPCConfig) SetSecurityGroupId(v string) *VPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *VPCConfig) SetVSwitchIds(v []*string) *VPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *VPCConfig) SetVpcId(v string) *VPCConfig {
	s.VpcId = &v
	return s
}

type Version struct {
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// my version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z07:00
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// 1
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s Version) String() string {
	return tea.Prettify(s)
}

func (s Version) GoString() string {
	return s.String()
}

func (s *Version) SetCreatedTime(v string) *Version {
	s.CreatedTime = &v
	return s
}

func (s *Version) SetDescription(v string) *Version {
	s.Description = &v
	return s
}

func (s *Version) SetLastModifiedTime(v string) *Version {
	s.LastModifiedTime = &v
	return s
}

func (s *Version) SetVersionId(v string) *Version {
	s.VersionId = &v
	return s
}

type WAFConfig struct {
	// example:
	//
	// true
	EnableWAF *bool `json:"enableWAF,omitempty" xml:"enableWAF,omitempty"`
}

func (s WAFConfig) String() string {
	return tea.Prettify(s)
}

func (s WAFConfig) GoString() string {
	return s.String()
}

func (s *WAFConfig) SetEnableWAF(v bool) *WAFConfig {
	s.EnableWAF = &v
	return s
}

type WildcardRule struct {
	// This parameter is required.
	//
	// example:
	//
	// /api/*
	Match *string `json:"match,omitempty" xml:"match,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /$1
	Replacement *string `json:"replacement,omitempty" xml:"replacement,omitempty"`
}

func (s WildcardRule) String() string {
	return tea.Prettify(s)
}

func (s WildcardRule) GoString() string {
	return s.String()
}

func (s *WildcardRule) SetMatch(v string) *WildcardRule {
	s.Match = &v
	return s
}

func (s *WildcardRule) SetReplacement(v string) *WildcardRule {
	s.Replacement = &v
	return s
}

type CreateAliasRequest struct {
	// The request parameters for creating an alias.
	//
	// This parameter is required.
	Body *CreateAliasInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) SetBody(v *CreateAliasInput) *CreateAliasRequest {
	s.Body = v
	return s
}

type CreateAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateAliasResponse) SetHeaders(v map[string]*string) *CreateAliasResponse {
	s.Headers = v
	return s
}

func (s *CreateAliasResponse) SetStatusCode(v int32) *CreateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAliasResponse) SetBody(v *Alias) *CreateAliasResponse {
	s.Body = v
	return s
}

type CreateCustomDomainRequest struct {
	// The information about the custom domain name.
	//
	// This parameter is required.
	Body *CreateCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainRequest) SetBody(v *CreateCustomDomainInput) *CreateCustomDomainRequest {
	s.Body = v
	return s
}

type CreateCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomain      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainResponse) SetHeaders(v map[string]*string) *CreateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomDomainResponse) SetStatusCode(v int32) *CreateCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomDomainResponse) SetBody(v *CustomDomain) *CreateCustomDomainResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	// The information about function configurations.
	//
	// This parameter is required.
	Body *CreateFunctionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetBody(v *CreateFunctionInput) *CreateFunctionRequest {
	s.Body = v
	return s
}

type CreateFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Function          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *Function) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateLayerVersionRequest struct {
	// The information about layer configurations.
	//
	// This parameter is required.
	Body *CreateLayerVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionRequest) SetBody(v *CreateLayerVersionInput) *CreateLayerVersionRequest {
	s.Body = v
	return s
}

type CreateLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionResponse) SetHeaders(v map[string]*string) *CreateLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateLayerVersionResponse) SetStatusCode(v int32) *CreateLayerVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayerVersionResponse) SetBody(v *Layer) *CreateLayerVersionResponse {
	s.Body = v
	return s
}

type CreateTriggerRequest struct {
	// The trigger configurations.
	//
	// This parameter is required.
	Body *CreateTriggerInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetBody(v *CreateTriggerInput) *CreateTriggerRequest {
	s.Body = v
	return s
}

type CreateTriggerResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Trigger           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetStatusCode(v int32) *CreateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *Trigger) *CreateTriggerResponse {
	s.Body = v
	return s
}

type CreateVpcBindingRequest struct {
	// The configurations of the virtual private cloud (VPC) binding.
	//
	// This parameter is required.
	Body *CreateVpcBindingInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingRequest) SetBody(v *CreateVpcBindingInput) *CreateVpcBindingRequest {
	s.Body = v
	return s
}

type CreateVpcBindingResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateVpcBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingResponse) SetHeaders(v map[string]*string) *CreateVpcBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcBindingResponse) SetStatusCode(v int32) *CreateVpcBindingResponse {
	s.StatusCode = &v
	return s
}

type DeleteAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponse) SetHeaders(v map[string]*string) *DeleteAliasResponse {
	s.Headers = v
	return s
}

func (s *DeleteAliasResponse) SetStatusCode(v int32) *DeleteAliasResponse {
	s.StatusCode = &v
	return s
}

type DeleteAsyncInvokeConfigRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncInvokeConfigRequest) SetQualifier(v string) *DeleteAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type DeleteAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *DeleteAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsyncInvokeConfigResponse) SetStatusCode(v int32) *DeleteAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

type DeleteConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteConcurrencyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteConcurrencyConfigResponse) SetHeaders(v map[string]*string) *DeleteConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteConcurrencyConfigResponse) SetStatusCode(v int32) *DeleteConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

type DeleteCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainResponse) SetHeaders(v map[string]*string) *DeleteCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomDomainResponse) SetStatusCode(v int32) *DeleteCustomDomainResponse {
	s.StatusCode = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

type DeleteFunctionVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteFunctionVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionVersionResponse) SetHeaders(v map[string]*string) *DeleteFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionVersionResponse) SetStatusCode(v int32) *DeleteFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

type DeleteLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayerVersionResponse) SetHeaders(v map[string]*string) *DeleteLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayerVersionResponse) SetStatusCode(v int32) *DeleteLayerVersionResponse {
	s.StatusCode = &v
	return s
}

type DeleteProvisionConfigRequest struct {
	// The function alias or LATEST.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s DeleteProvisionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteProvisionConfigRequest) SetQualifier(v string) *DeleteProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

type DeleteProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteProvisionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteProvisionConfigResponse) SetHeaders(v map[string]*string) *DeleteProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteProvisionConfigResponse) SetStatusCode(v int32) *DeleteProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

type DeleteTriggerResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponse) SetHeaders(v map[string]*string) *DeleteTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteTriggerResponse) SetStatusCode(v int32) *DeleteTriggerResponse {
	s.StatusCode = &v
	return s
}

type DeleteVpcBindingResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteVpcBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcBindingResponse) SetHeaders(v map[string]*string) *DeleteVpcBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcBindingResponse) SetStatusCode(v int32) *DeleteVpcBindingResponse {
	s.StatusCode = &v
	return s
}

type GetAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAliasResponse) GoString() string {
	return s.String()
}

func (s *GetAliasResponse) SetHeaders(v map[string]*string) *GetAliasResponse {
	s.Headers = v
	return s
}

func (s *GetAliasResponse) SetStatusCode(v int32) *GetAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliasResponse) SetBody(v *Alias) *GetAliasResponse {
	s.Body = v
	return s
}

type GetAsyncInvokeConfigRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncInvokeConfigRequest) SetQualifier(v string) *GetAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type GetAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncConfig       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *GetAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncInvokeConfigResponse) SetStatusCode(v int32) *GetAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncInvokeConfigResponse) SetBody(v *AsyncConfig) *GetAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

type GetAsyncTaskRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) SetQualifier(v string) *GetAsyncTaskRequest {
	s.Qualifier = &v
	return s
}

type GetAsyncTaskResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncTask         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponse) SetHeaders(v map[string]*string) *GetAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskResponse) SetStatusCode(v int32) *GetAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTaskResponse) SetBody(v *AsyncTask) *GetAsyncTaskResponse {
	s.Body = v
	return s
}

type GetConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConcurrencyConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConcurrencyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *GetConcurrencyConfigResponse) SetHeaders(v map[string]*string) *GetConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *GetConcurrencyConfigResponse) SetStatusCode(v int32) *GetConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConcurrencyConfigResponse) SetBody(v *ConcurrencyConfig) *GetConcurrencyConfigResponse {
	s.Body = v
	return s
}

type GetCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomain      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *GetCustomDomainResponse) SetHeaders(v map[string]*string) *GetCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *GetCustomDomainResponse) SetStatusCode(v int32) *GetCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomDomainResponse) SetBody(v *CustomDomain) *GetCustomDomainResponse {
	s.Body = v
	return s
}

type GetFunctionRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetQualifier(v string) *GetFunctionRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Function          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *Function) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetFunctionCodeRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetFunctionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeRequest) SetQualifier(v string) *GetFunctionCodeRequest {
	s.Qualifier = &v
	return s
}

type GetFunctionCodeResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OutputFuncCode    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCodeResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeResponse) SetHeaders(v map[string]*string) *GetFunctionCodeResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCodeResponse) SetStatusCode(v int32) *GetFunctionCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionCodeResponse) SetBody(v *OutputFuncCode) *GetFunctionCodeResponse {
	s.Body = v
	return s
}

type GetLayerVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionResponse) SetHeaders(v map[string]*string) *GetLayerVersionResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionResponse) SetStatusCode(v int32) *GetLayerVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerVersionResponse) SetBody(v *Layer) *GetLayerVersionResponse {
	s.Body = v
	return s
}

type GetLayerVersionByArnResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Layer             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayerVersionByArnResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLayerVersionByArnResponse) GoString() string {
	return s.String()
}

func (s *GetLayerVersionByArnResponse) SetHeaders(v map[string]*string) *GetLayerVersionByArnResponse {
	s.Headers = v
	return s
}

func (s *GetLayerVersionByArnResponse) SetStatusCode(v int32) *GetLayerVersionByArnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayerVersionByArnResponse) SetBody(v *Layer) *GetLayerVersionByArnResponse {
	s.Body = v
	return s
}

type GetProvisionConfigRequest struct {
	// The function alias or LATEST.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetProvisionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigRequest) SetQualifier(v string) *GetProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

type GetProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionConfig   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionConfigResponse) SetHeaders(v map[string]*string) *GetProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionConfigResponse) SetStatusCode(v int32) *GetProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionConfigResponse) SetBody(v *ProvisionConfig) *GetProvisionConfigResponse {
	s.Body = v
	return s
}

type GetTriggerResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Trigger           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetTriggerResponse) SetHeaders(v map[string]*string) *GetTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetTriggerResponse) SetStatusCode(v int32) *GetTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTriggerResponse) SetBody(v *Trigger) *GetTriggerResponse {
	s.Body = v
	return s
}

type InvokeFunctionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders" xml:"commonHeaders"`
	// The ID of the asynchronous task. You must enable the asynchronous task feature in advance.
	//
	// >  If you use an SDK to invoke a function, we recommend that you specify a business-related ID to facilitate subsequent operations. For example, a video processing function can use video file names as invocation IDs. This way, you can easily check whether a video is successfully processed or terminated before it is processed. The ID can start only with letters or underscores. An ID can contain *letters, digits (0 - 9), underscores*, and hyphens (-). It can be up to 128 characters in length. If you do not specify the ID of the asynchronous invocation, the system automatically generates an ID.
	//
	// example:
	//
	// test-id
	XFcAsyncTaskId *string `json:"x-fc-async-task-id,omitempty" xml:"x-fc-async-task-id,omitempty"`
	// The type of function invocation. Valid values: Sync and Async.
	//
	// example:
	//
	// Sync
	XFcInvocationType *string `json:"x-fc-invocation-type,omitempty" xml:"x-fc-invocation-type,omitempty"`
	// The log type of function invocation. Valid values: None and Tail.
	//
	// example:
	//
	// Tail
	XFcLogType *string `json:"x-fc-log-type,omitempty" xml:"x-fc-log-type,omitempty"`
}

func (s InvokeFunctionHeaders) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionHeaders) GoString() string {
	return s.String()
}

func (s *InvokeFunctionHeaders) SetCommonHeaders(v map[string]*string) *InvokeFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcAsyncTaskId(v string) *InvokeFunctionHeaders {
	s.XFcAsyncTaskId = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcInvocationType(v string) *InvokeFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcLogType(v string) *InvokeFunctionHeaders {
	s.XFcLogType = &v
	return s
}

type InvokeFunctionRequest struct {
	// The request parameters of function invocation.
	//
	// example:
	//
	// event
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s InvokeFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionRequest) GoString() string {
	return s.String()
}

func (s *InvokeFunctionRequest) SetBody(v io.Reader) *InvokeFunctionRequest {
	s.Body = v
	return s
}

func (s *InvokeFunctionRequest) SetQualifier(v string) *InvokeFunctionRequest {
	s.Qualifier = &v
	return s
}

type InvokeFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeFunctionResponse) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponse) SetHeaders(v map[string]*string) *InvokeFunctionResponse {
	s.Headers = v
	return s
}

func (s *InvokeFunctionResponse) SetStatusCode(v int32) *InvokeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeFunctionResponse) SetBody(v io.Reader) *InvokeFunctionResponse {
	s.Body = v
	return s
}

type ListAliasesRequest struct {
	// The number of aliases returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The alias prefix.
	//
	// example:
	//
	// my-alias
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) SetLimit(v int32) *ListAliasesRequest {
	s.Limit = &v
	return s
}

func (s *ListAliasesRequest) SetNextToken(v string) *ListAliasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAliasesRequest) SetPrefix(v string) *ListAliasesRequest {
	s.Prefix = &v
	return s
}

type ListAliasesResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliasesOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesResponse) SetHeaders(v map[string]*string) *ListAliasesResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesResponse) SetStatusCode(v int32) *ListAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliasesResponse) SetBody(v *ListAliasesOutput) *ListAliasesResponse {
	s.Body = v
	return s
}

type ListAsyncInvokeConfigsRequest struct {
	// The function name. If you do not configure this parameter, the asynchronous invocation configurations of all functions are displayed.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The maximum number of entries to be returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The paging information. This parameter specifies the start point of the query.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAsyncInvokeConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncInvokeConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigsRequest) SetFunctionName(v string) *ListAsyncInvokeConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListAsyncInvokeConfigsRequest) SetLimit(v int32) *ListAsyncInvokeConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListAsyncInvokeConfigsRequest) SetNextToken(v string) *ListAsyncInvokeConfigsRequest {
	s.NextToken = &v
	return s
}

type ListAsyncInvokeConfigsResponse struct {
	Headers    map[string]*string           `json:"headers" xml:"headers"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncInvokeConfigOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncInvokeConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncInvokeConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncInvokeConfigsResponse) SetHeaders(v map[string]*string) *ListAsyncInvokeConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncInvokeConfigsResponse) SetStatusCode(v int32) *ListAsyncInvokeConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncInvokeConfigsResponse) SetBody(v *ListAsyncInvokeConfigOutput) *ListAsyncInvokeConfigsResponse {
	s.Body = v
	return s
}

type ListAsyncTasksRequest struct {
	// Specifies whether to return input parameters of the asynchronous tasks. Valid values:
	//
	// 	- true: returns the `invocationPayload` parameter in the response.
	//
	// 	- false: does not return the `invocationPayload` parameter in the response.
	//
	// >  The `invocationPayload` parameter indicates the input parameters of an asynchronous task.
	//
	// example:
	//
	// true
	IncludePayload *bool `json:"includePayload,omitempty" xml:"includePayload,omitempty"`
	// The number of asynchronous tasks to return. Valid values: [1,100]. Default value: 50.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID prefix of asynchronous tasks. If this parameter is specified, a list of asynchronous tasks whose IDs match the prefix is returned.
	//
	// example:
	//
	// job-
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The order in which the returned asynchronous tasks are sorted.
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// example:
	//
	// asc
	SortOrderByTime *string `json:"sortOrderByTime,omitempty" xml:"sortOrderByTime,omitempty"`
	// The start time of the period during which the asynchronous tasks are initiated.
	//
	// example:
	//
	// 1640966400000
	StartedTimeBegin *int64 `json:"startedTimeBegin,omitempty" xml:"startedTimeBegin,omitempty"`
	// The end time of the period during which the asynchronous tasks are initiated.
	//
	// example:
	//
	// 1640966400000
	StartedTimeEnd *int64 `json:"startedTimeEnd,omitempty" xml:"startedTimeEnd,omitempty"`
	// The state of asynchronous tasks. The following items list the states of an asynchronous task:
	//
	// 	- Enqueued: The asynchronous invocation is enqueued and is waiting to be executed.
	//
	// 	- Succeeded: The invocation is successful.
	//
	// 	- Failed: The invocation fails.
	//
	// 	- Running: The invocation is being executed.
	//
	// 	- Stopped: The invocation is terminated.
	//
	// 	- Stopping: The invocation is being terminated.
	//
	// 	- Invalid: The invocation is invalid and not executed due to specific reasons. For example, the function is deleted.
	//
	// 	- Expired: The maximum validity period of messages is specified for asynchronous invocation. The invocation is discarded and not executed because the specified maximum validity period has elapsed.
	//
	// 	- Retrying: The asynchronous invocation is being retried due to an execution error.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAsyncTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) SetIncludePayload(v bool) *ListAsyncTasksRequest {
	s.IncludePayload = &v
	return s
}

func (s *ListAsyncTasksRequest) SetLimit(v int32) *ListAsyncTasksRequest {
	s.Limit = &v
	return s
}

func (s *ListAsyncTasksRequest) SetNextToken(v string) *ListAsyncTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsyncTasksRequest) SetPrefix(v string) *ListAsyncTasksRequest {
	s.Prefix = &v
	return s
}

func (s *ListAsyncTasksRequest) SetQualifier(v string) *ListAsyncTasksRequest {
	s.Qualifier = &v
	return s
}

func (s *ListAsyncTasksRequest) SetSortOrderByTime(v string) *ListAsyncTasksRequest {
	s.SortOrderByTime = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStartedTimeBegin(v int64) *ListAsyncTasksRequest {
	s.StartedTimeBegin = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStartedTimeEnd(v int64) *ListAsyncTasksRequest {
	s.StartedTimeEnd = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStatus(v string) *ListAsyncTasksRequest {
	s.Status = &v
	return s
}

type ListAsyncTasksResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTaskOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponse) SetHeaders(v map[string]*string) *ListAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTasksResponse) SetStatusCode(v int32) *ListAsyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTasksResponse) SetBody(v *ListAsyncTaskOutput) *ListAsyncTasksResponse {
	s.Body = v
	return s
}

type ListConcurrencyConfigsRequest struct {
	// The function name. If you leave this parameter empty, the concurrency configurations of all functions are returned.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConcurrencyConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConcurrencyConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsRequest) SetFunctionName(v string) *ListConcurrencyConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListConcurrencyConfigsRequest) SetLimit(v int32) *ListConcurrencyConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListConcurrencyConfigsRequest) SetNextToken(v string) *ListConcurrencyConfigsRequest {
	s.NextToken = &v
	return s
}

type ListConcurrencyConfigsResponse struct {
	Headers    map[string]*string            `json:"headers" xml:"headers"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConcurrencyConfigsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConcurrencyConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConcurrencyConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConcurrencyConfigsResponse) SetHeaders(v map[string]*string) *ListConcurrencyConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConcurrencyConfigsResponse) SetStatusCode(v int32) *ListConcurrencyConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConcurrencyConfigsResponse) SetBody(v *ListConcurrencyConfigsOutput) *ListConcurrencyConfigsResponse {
	s.Body = v
	return s
}

type ListCustomDomainsRequest struct {
	// The number of custom domain names returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The domain name prefix.
	//
	// example:
	//
	// foo
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListCustomDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsRequest) SetLimit(v int32) *ListCustomDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListCustomDomainsRequest) SetNextToken(v string) *ListCustomDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainsRequest) SetPrefix(v string) *ListCustomDomainsRequest {
	s.Prefix = &v
	return s
}

type ListCustomDomainsResponse struct {
	Headers    map[string]*string      `json:"headers" xml:"headers"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomDomainOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsResponse) SetHeaders(v map[string]*string) *ListCustomDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomDomainsResponse) SetStatusCode(v int32) *ListCustomDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomDomainsResponse) SetBody(v *ListCustomDomainOutput) *ListCustomDomainsResponse {
	s.Body = v
	return s
}

type ListFunctionVersionsRequest struct {
	// The sorting mode of function versions. Valid values: BACKWARD and FORWARD.
	//
	// example:
	//
	// BACKWARD
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The number of function versions that are returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFunctionVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionVersionsRequest) SetDirection(v string) *ListFunctionVersionsRequest {
	s.Direction = &v
	return s
}

func (s *ListFunctionVersionsRequest) SetLimit(v int32) *ListFunctionVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionVersionsRequest) SetNextToken(v string) *ListFunctionVersionsRequest {
	s.NextToken = &v
	return s
}

type ListFunctionVersionsResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionVersionsResponse) SetHeaders(v map[string]*string) *ListFunctionVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionVersionsResponse) SetStatusCode(v int32) *ListFunctionVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionVersionsResponse) SetBody(v *ListVersionsOutput) *ListFunctionVersionsResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// example:
	//
	// v3
	FcVersion *string `json:"fcVersion,omitempty" xml:"fcVersion,omitempty"`
	// The number of functions to return. The minimum value is 1 and the maximum value is 100.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The prefix of the function name.
	//
	// example:
	//
	// my-func
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetFcVersion(v string) *ListFunctionsRequest {
	s.FcVersion = &v
	return s
}

func (s *ListFunctionsRequest) SetLimit(v int32) *ListFunctionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFunctionsRequest) SetNextToken(v string) *ListFunctionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsOutput) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The function version or alias.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// Specifies whether to list all instances. Valid values: true and false.
	//
	// example:
	//
	// true
	WithAllActive *bool `json:"withAllActive,omitempty" xml:"withAllActive,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetQualifier(v string) *ListInstancesRequest {
	s.Qualifier = &v
	return s
}

func (s *ListInstancesRequest) SetWithAllActive(v bool) *ListInstancesRequest {
	s.WithAllActive = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesOutput) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListLayerVersionsRequest struct {
	// The number of versions to be returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The initial version of the layer.
	//
	// example:
	//
	// 1
	StartVersion *string `json:"startVersion,omitempty" xml:"startVersion,omitempty"`
}

func (s ListLayerVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsRequest) SetLimit(v int32) *ListLayerVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListLayerVersionsRequest) SetStartVersion(v string) *ListLayerVersionsRequest {
	s.StartVersion = &v
	return s
}

type ListLayerVersionsResponse struct {
	Headers    map[string]*string      `json:"headers" xml:"headers"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayerVersionOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayerVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayerVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsResponse) SetHeaders(v map[string]*string) *ListLayerVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListLayerVersionsResponse) SetStatusCode(v int32) *ListLayerVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayerVersionsResponse) SetBody(v *ListLayerVersionOutput) *ListLayerVersionsResponse {
	s.Body = v
	return s
}

type ListLayersRequest struct {
	// The number of layers that are returned
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Specifies whether the layer is official. Valid values: true and false.
	//
	// example:
	//
	// true
	Official *string `json:"official,omitempty" xml:"official,omitempty"`
	// The name prefix of the layer.
	//
	// example:
	//
	// my-layer
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// Specifies whether the layer is public. Valid values: true and false.
	//
	// example:
	//
	// true
	Public *string `json:"public,omitempty" xml:"public,omitempty"`
}

func (s ListLayersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayersRequest) GoString() string {
	return s.String()
}

func (s *ListLayersRequest) SetLimit(v int32) *ListLayersRequest {
	s.Limit = &v
	return s
}

func (s *ListLayersRequest) SetNextToken(v string) *ListLayersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLayersRequest) SetOfficial(v string) *ListLayersRequest {
	s.Official = &v
	return s
}

func (s *ListLayersRequest) SetPrefix(v string) *ListLayersRequest {
	s.Prefix = &v
	return s
}

func (s *ListLayersRequest) SetPublic(v string) *ListLayersRequest {
	s.Public = &v
	return s
}

type ListLayersResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayersOutput  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLayersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayersResponse) GoString() string {
	return s.String()
}

func (s *ListLayersResponse) SetHeaders(v map[string]*string) *ListLayersResponse {
	s.Headers = v
	return s
}

func (s *ListLayersResponse) SetStatusCode(v int32) *ListLayersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayersResponse) SetBody(v *ListLayersOutput) *ListLayersResponse {
	s.Body = v
	return s
}

type ListProvisionConfigsRequest struct {
	// The name of the function. If this parameter is not specified, the provisioned configurations of all functions are listed.
	//
	// example:
	//
	// my-func
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// Number of provisioned configurations to return.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProvisionConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsRequest) SetFunctionName(v string) *ListProvisionConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetLimit(v int32) *ListProvisionConfigsRequest {
	s.Limit = &v
	return s
}

func (s *ListProvisionConfigsRequest) SetNextToken(v string) *ListProvisionConfigsRequest {
	s.NextToken = &v
	return s
}

type ListProvisionConfigsResponse struct {
	Headers    map[string]*string          `json:"headers" xml:"headers"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionConfigsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionConfigsResponse) SetHeaders(v map[string]*string) *ListProvisionConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionConfigsResponse) SetStatusCode(v int32) *ListProvisionConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionConfigsResponse) SetBody(v *ListProvisionConfigsOutput) *ListProvisionConfigsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The number of resources to return.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs.
	ResourceId []*string `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::FC:FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// You can query up to 20 tags at a time.
	Tag []*ListTagResourcesRequestTag `json:"Tag" xml:"Tag" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetLimit(v int32) *ListTagResourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// The tag key can be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and can be an empty string.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	// The number of resources to return.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs.
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::FC:FUNCTION
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// You can query up to 20 tags at a time.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetLimit(v int32) *ListTagResourcesShrinkRequest {
	s.Limit = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string      `json:"headers" xml:"headers"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesOutput) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTriggersRequest struct {
	// The number of triggers returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The trigger name prefix.
	//
	// example:
	//
	// my-trigger
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) SetLimit(v int32) *ListTriggersRequest {
	s.Limit = &v
	return s
}

func (s *ListTriggersRequest) SetNextToken(v string) *ListTriggersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTriggersRequest) SetPrefix(v string) *ListTriggersRequest {
	s.Prefix = &v
	return s
}

type ListTriggersResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTriggersOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListTriggersResponse) SetHeaders(v map[string]*string) *ListTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListTriggersResponse) SetStatusCode(v int32) *ListTriggersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTriggersResponse) SetBody(v *ListTriggersOutput) *ListTriggersResponse {
	s.Body = v
	return s
}

type ListVpcBindingsResponse struct {
	Headers    map[string]*string     `json:"headers" xml:"headers"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcBindingsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcBindingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsResponse) SetHeaders(v map[string]*string) *ListVpcBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcBindingsResponse) SetStatusCode(v int32) *ListVpcBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcBindingsResponse) SetBody(v *ListVpcBindingsOutput) *ListVpcBindingsResponse {
	s.Body = v
	return s
}

type PublishFunctionVersionRequest struct {
	// The information about the function version.
	//
	// This parameter is required.
	Body *PublishVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFunctionVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishFunctionVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishFunctionVersionRequest) SetBody(v *PublishVersionInput) *PublishFunctionVersionRequest {
	s.Body = v
	return s
}

type PublishFunctionVersionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Version           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishFunctionVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishFunctionVersionResponse) SetHeaders(v map[string]*string) *PublishFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishFunctionVersionResponse) SetStatusCode(v int32) *PublishFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishFunctionVersionResponse) SetBody(v *Version) *PublishFunctionVersionResponse {
	s.Body = v
	return s
}

type PutAsyncInvokeConfigRequest struct {
	// The asynchronous invocation configurations.
	//
	// This parameter is required.
	Body *PutAsyncInvokeConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutAsyncInvokeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutAsyncInvokeConfigRequest) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigRequest) SetBody(v *PutAsyncInvokeConfigInput) *PutAsyncInvokeConfigRequest {
	s.Body = v
	return s
}

func (s *PutAsyncInvokeConfigRequest) SetQualifier(v string) *PutAsyncInvokeConfigRequest {
	s.Qualifier = &v
	return s
}

type PutAsyncInvokeConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncConfig       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAsyncInvokeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutAsyncInvokeConfigResponse) GoString() string {
	return s.String()
}

func (s *PutAsyncInvokeConfigResponse) SetHeaders(v map[string]*string) *PutAsyncInvokeConfigResponse {
	s.Headers = v
	return s
}

func (s *PutAsyncInvokeConfigResponse) SetStatusCode(v int32) *PutAsyncInvokeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAsyncInvokeConfigResponse) SetBody(v *AsyncConfig) *PutAsyncInvokeConfigResponse {
	s.Body = v
	return s
}

type PutConcurrencyConfigRequest struct {
	// The concurrency configurations.
	//
	// This parameter is required.
	Body *PutConcurrencyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutConcurrencyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutConcurrencyConfigRequest) GoString() string {
	return s.String()
}

func (s *PutConcurrencyConfigRequest) SetBody(v *PutConcurrencyInput) *PutConcurrencyConfigRequest {
	s.Body = v
	return s
}

type PutConcurrencyConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConcurrencyConfig `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutConcurrencyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutConcurrencyConfigResponse) GoString() string {
	return s.String()
}

func (s *PutConcurrencyConfigResponse) SetHeaders(v map[string]*string) *PutConcurrencyConfigResponse {
	s.Headers = v
	return s
}

func (s *PutConcurrencyConfigResponse) SetStatusCode(v int32) *PutConcurrencyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutConcurrencyConfigResponse) SetBody(v *ConcurrencyConfig) *PutConcurrencyConfigResponse {
	s.Body = v
	return s
}

type PutLayerACLRequest struct {
	// Specify the access permission of the layer. A value of 1 indicates public and a value of 0 indicates private. The default value is 0.
	//
	// example:
	//
	// 1
	Acl *string `json:"acl,omitempty" xml:"acl,omitempty"`
	// Deprecated
	//
	// Specify whether the layer is a public layer. Valid values: true and false.
	//
	// example:
	//
	// true
	Public *string `json:"public,omitempty" xml:"public,omitempty"`
}

func (s PutLayerACLRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLayerACLRequest) GoString() string {
	return s.String()
}

func (s *PutLayerACLRequest) SetAcl(v string) *PutLayerACLRequest {
	s.Acl = &v
	return s
}

func (s *PutLayerACLRequest) SetPublic(v string) *PutLayerACLRequest {
	s.Public = &v
	return s
}

type PutLayerACLResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutLayerACLResponse) String() string {
	return tea.Prettify(s)
}

func (s PutLayerACLResponse) GoString() string {
	return s.String()
}

func (s *PutLayerACLResponse) SetHeaders(v map[string]*string) *PutLayerACLResponse {
	s.Headers = v
	return s
}

func (s *PutLayerACLResponse) SetStatusCode(v int32) *PutLayerACLResponse {
	s.StatusCode = &v
	return s
}

type PutProvisionConfigRequest struct {
	// The information about the provisioned configuration.
	//
	// This parameter is required.
	Body *PutProvisionConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	// The function alias or LATEST.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s PutProvisionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigRequest) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigRequest) SetBody(v *PutProvisionConfigInput) *PutProvisionConfigRequest {
	s.Body = v
	return s
}

func (s *PutProvisionConfigRequest) SetQualifier(v string) *PutProvisionConfigRequest {
	s.Qualifier = &v
	return s
}

type PutProvisionConfigResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionConfig   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutProvisionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutProvisionConfigResponse) GoString() string {
	return s.String()
}

func (s *PutProvisionConfigResponse) SetHeaders(v map[string]*string) *PutProvisionConfigResponse {
	s.Headers = v
	return s
}

func (s *PutProvisionConfigResponse) SetStatusCode(v int32) *PutProvisionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutProvisionConfigResponse) SetBody(v *ProvisionConfig) *PutProvisionConfigResponse {
	s.Body = v
	return s
}

type StopAsyncTaskRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s StopAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAsyncTaskRequest) SetQualifier(v string) *StopAsyncTaskRequest {
	s.Qualifier = &v
	return s
}

type StopAsyncTaskResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s StopAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *StopAsyncTaskResponse) SetHeaders(v map[string]*string) *StopAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *StopAsyncTaskResponse) SetStatusCode(v int32) *StopAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

type TagResourcesRequest struct {
	// The configuration of the resource tag.
	//
	// This parameter is required.
	Body *TagResourcesInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetBody(v *TagResourcesInput) *TagResourcesRequest {
	s.Body = v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to delete all tags.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource identifiers.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// function
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag to remove. You can specify a maximum of 50 tags.
	TagKey []*string `json:"TagKey" xml:"TagKey" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to delete all tags.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The resource identifiers.
	//
	// This parameter is required.
	ResourceIdShrink *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// function
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag to remove. You can specify a maximum of 50 tags.
	TagKeyShrink *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeyShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeyShrink = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UpdateAliasRequest struct {
	// The alias information to be updated.
	//
	// This parameter is required.
	Body *UpdateAliasInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetBody(v *UpdateAliasInput) *UpdateAliasRequest {
	s.Body = v
	return s
}

type UpdateAliasResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Alias             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetStatusCode(v int32) *UpdateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *Alias) *UpdateAliasResponse {
	s.Body = v
	return s
}

type UpdateCustomDomainRequest struct {
	// The information about the custom domain name.
	//
	// This parameter is required.
	Body *UpdateCustomDomainInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainRequest) SetBody(v *UpdateCustomDomainInput) *UpdateCustomDomainRequest {
	s.Body = v
	return s
}

type UpdateCustomDomainResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomDomain      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainResponse) SetHeaders(v map[string]*string) *UpdateCustomDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomDomainResponse) SetStatusCode(v int32) *UpdateCustomDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomDomainResponse) SetBody(v *CustomDomain) *UpdateCustomDomainResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	// The function information
	//
	// This parameter is required.
	Body *UpdateFunctionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetBody(v *UpdateFunctionInput) *UpdateFunctionRequest {
	s.Body = v
	return s
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Function          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *Function) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateTriggerRequest struct {
	// The trigger configurations.
	//
	// This parameter is required.
	Body *UpdateTriggerInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTriggerRequest) SetBody(v *UpdateTriggerInput) *UpdateTriggerRequest {
	s.Body = v
	return s
}

type UpdateTriggerResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Trigger           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTriggerResponse) SetHeaders(v map[string]*string) *UpdateTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTriggerResponse) SetStatusCode(v int32) *UpdateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTriggerResponse) SetBody(v *Trigger) *UpdateTriggerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("fc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建函数别名。
//
// @param request - CreateAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAliasResponse
func (client *Client) CreateAliasWithOptions(functionName *string, request *CreateAliasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlias"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/aliases"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建函数别名。
//
// @param request - CreateAliasRequest
//
// @return CreateAliasResponse
func (client *Client) CreateAlias(functionName *string, request *CreateAliasRequest) (_result *CreateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAliasResponse{}
	_body, _err := client.CreateAliasWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom domain name.
//
// Description:
//
// If you wish to access applications or functions created in Function Compute via a fixed domain name in a production environment, or if you need to address the forced download behavior when accessing HTTP triggers, you can achieve this by binding a custom domain to your application or function.
//
// @param request - CreateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomainWithOptions(request *CreateCustomDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomDomain"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/custom-domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom domain name.
//
// Description:
//
// If you wish to access applications or functions created in Function Compute via a fixed domain name in a production environment, or if you need to address the forced download behavior when accessing HTTP triggers, you can achieve this by binding a custom domain to your application or function.
//
// @param request - CreateCustomDomainRequest
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomain(request *CreateCustomDomainRequest) (_result *CreateCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CreateCustomDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a function.
//
// Description:
//
// In Function Compute, a function serves as the smallest unit of resource scheduling and execution, typically referring to a piece of code written by users that can execute independently in response to specific events or requests.
//
// @param request - CreateFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a function.
//
// Description:
//
// In Function Compute, a function serves as the smallest unit of resource scheduling and execution, typically referring to a piece of code written by users that can execute independently in response to specific events or requests.
//
// @param request - CreateFunctionRequest
//
// @return CreateFunctionResponse
func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建层版本。
//
// @param request - CreateLayerVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayerVersionResponse
func (client *Client) CreateLayerVersionWithOptions(layerName *string, request *CreateLayerVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLayerVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayerVersion"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers/" + tea.StringValue(openapiutil.GetEncodeParam(layerName)) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建层版本。
//
// @param request - CreateLayerVersionRequest
//
// @return CreateLayerVersionResponse
func (client *Client) CreateLayerVersion(layerName *string, request *CreateLayerVersionRequest) (_result *CreateLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLayerVersionResponse{}
	_body, _err := client.CreateLayerVersionWithOptions(layerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建函数触发器。
//
// @param request - CreateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTriggerResponse
func (client *Client) CreateTriggerWithOptions(functionName *string, request *CreateTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrigger"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/triggers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建函数触发器。
//
// @param request - CreateTriggerRequest
//
// @return CreateTriggerResponse
func (client *Client) CreateTrigger(functionName *string, request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a VPC connection.
//
// @param request - CreateVpcBindingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcBindingResponse
func (client *Client) CreateVpcBindingWithOptions(functionName *string, request *CreateVpcBindingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateVpcBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcBinding"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/vpc-bindings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateVpcBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VPC connection.
//
// @param request - CreateVpcBindingRequest
//
// @return CreateVpcBindingResponse
func (client *Client) CreateVpcBinding(functionName *string, request *CreateVpcBindingRequest) (_result *CreateVpcBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVpcBindingResponse{}
	_body, _err := client.CreateVpcBindingWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alias.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAliasResponse
func (client *Client) DeleteAliasWithOptions(functionName *string, aliasName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlias"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/aliases/" + tea.StringValue(openapiutil.GetEncodeParam(aliasName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alias.
//
// @return DeleteAliasResponse
func (client *Client) DeleteAlias(functionName *string, aliasName *string) (_result *DeleteAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAliasResponse{}
	_body, _err := client.DeleteAliasWithOptions(functionName, aliasName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an asynchronous invocation configuration.
//
// @param request - DeleteAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsyncInvokeConfigResponse
func (client *Client) DeleteAsyncInvokeConfigWithOptions(functionName *string, request *DeleteAsyncInvokeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAsyncInvokeConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-invoke-config"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an asynchronous invocation configuration.
//
// @param request - DeleteAsyncInvokeConfigRequest
//
// @return DeleteAsyncInvokeConfigResponse
func (client *Client) DeleteAsyncInvokeConfig(functionName *string, request *DeleteAsyncInvokeConfigRequest) (_result *DeleteAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAsyncInvokeConfigResponse{}
	_body, _err := client.DeleteAsyncInvokeConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a concurrency configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConcurrencyConfigResponse
func (client *Client) DeleteConcurrencyConfigWithOptions(functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConcurrencyConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConcurrencyConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/concurrency"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteConcurrencyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a concurrency configuration.
//
// @return DeleteConcurrencyConfigResponse
func (client *Client) DeleteConcurrencyConfig(functionName *string) (_result *DeleteConcurrencyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConcurrencyConfigResponse{}
	_body, _err := client.DeleteConcurrencyConfigWithOptions(functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomainWithOptions(domainName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomDomain"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/custom-domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name.
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomain(domainName *string) (_result *DeleteCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.DeleteCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a function.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithOptions(functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a function.
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunction(functionName *string) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a function version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionVersionResponse
func (client *Client) DeleteFunctionVersionWithOptions(functionName *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionVersion"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteFunctionVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a function version.
//
// @return DeleteFunctionVersionResponse
func (client *Client) DeleteFunctionVersion(functionName *string, versionId *string) (_result *DeleteFunctionVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionVersionResponse{}
	_body, _err := client.DeleteFunctionVersionWithOptions(functionName, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a layer version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayerVersionResponse
func (client *Client) DeleteLayerVersionWithOptions(layerName *string, version *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLayerVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayerVersion"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers/" + tea.StringValue(openapiutil.GetEncodeParam(layerName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(version))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a layer version.
//
// @return DeleteLayerVersionResponse
func (client *Client) DeleteLayerVersion(layerName *string, version *string) (_result *DeleteLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLayerVersionResponse{}
	_body, _err := client.DeleteLayerVersionWithOptions(layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a provisioned configuration.
//
// @param request - DeleteProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProvisionConfigResponse
func (client *Client) DeleteProvisionConfigWithOptions(functionName *string, request *DeleteProvisionConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProvisionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProvisionConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/provision-config"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteProvisionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a provisioned configuration.
//
// @param request - DeleteProvisionConfigRequest
//
// @return DeleteProvisionConfigResponse
func (client *Client) DeleteProvisionConfig(functionName *string, request *DeleteProvisionConfigRequest) (_result *DeleteProvisionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProvisionConfigResponse{}
	_body, _err := client.DeleteProvisionConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a trigger.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTriggerWithOptions(functionName *string, triggerName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTriggerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrigger"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/triggers/" + tea.StringValue(openapiutil.GetEncodeParam(triggerName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trigger.
//
// @return DeleteTriggerResponse
func (client *Client) DeleteTrigger(functionName *string, triggerName *string) (_result *DeleteTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTriggerResponse{}
	_body, _err := client.DeleteTriggerWithOptions(functionName, triggerName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specified policy group for a VPC firewall.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcBindingResponse
func (client *Client) DeleteVpcBindingWithOptions(functionName *string, vpcId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVpcBindingResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcBinding"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/vpc-bindings/" + tea.StringValue(openapiutil.GetEncodeParam(vpcId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteVpcBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specified policy group for a VPC firewall.
//
// @return DeleteVpcBindingResponse
func (client *Client) DeleteVpcBinding(functionName *string, vpcId *string) (_result *DeleteVpcBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVpcBindingResponse{}
	_body, _err := client.DeleteVpcBindingWithOptions(functionName, vpcId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an alias.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliasResponse
func (client *Client) GetAliasWithOptions(functionName *string, aliasName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAliasResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlias"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/aliases/" + tea.StringValue(openapiutil.GetEncodeParam(aliasName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an alias.
//
// @return GetAliasResponse
func (client *Client) GetAlias(functionName *string, aliasName *string) (_result *GetAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAliasResponse{}
	_body, _err := client.GetAliasWithOptions(functionName, aliasName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets asynchronous invocation configurations of a function.
//
// @param request - GetAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncInvokeConfigResponse
func (client *Client) GetAsyncInvokeConfigWithOptions(functionName *string, request *GetAsyncInvokeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncInvokeConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-invoke-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets asynchronous invocation configurations of a function.
//
// @param request - GetAsyncInvokeConfigRequest
//
// @return GetAsyncInvokeConfigResponse
func (client *Client) GetAsyncInvokeConfig(functionName *string, request *GetAsyncInvokeConfigRequest) (_result *GetAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAsyncInvokeConfigResponse{}
	_body, _err := client.GetAsyncInvokeConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithOptions(functionName *string, taskId *string, request *GetAsyncTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncTask"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTask(functionName *string, taskId *string, request *GetAsyncTaskRequest) (_result *GetAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.GetAsyncTaskWithOptions(functionName, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a concurrency configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConcurrencyConfigResponse
func (client *Client) GetConcurrencyConfigWithOptions(functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConcurrencyConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConcurrencyConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/concurrency"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConcurrencyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a concurrency configuration.
//
// @return GetConcurrencyConfigResponse
func (client *Client) GetConcurrencyConfig(functionName *string) (_result *GetConcurrencyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConcurrencyConfigResponse{}
	_body, _err := client.GetConcurrencyConfigWithOptions(functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a custom domain name.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomainWithOptions(domainName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomDomain"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/custom-domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a custom domain name.
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomain(domainName *string) (_result *GetCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomDomainResponse{}
	_body, _err := client.GetCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a function.
//
// @param request - GetFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithOptions(functionName *string, request *GetFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a function.
//
// @param request - GetFunctionRequest
//
// @return GetFunctionResponse
func (client *Client) GetFunction(functionName *string, request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a code package of a function.
//
// @param request - GetFunctionCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionCodeResponse
func (client *Client) GetFunctionCodeWithOptions(functionName *string, request *GetFunctionCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionCode"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/code"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a code package of a function.
//
// @param request - GetFunctionCodeRequest
//
// @return GetFunctionCodeResponse
func (client *Client) GetFunctionCode(functionName *string, request *GetFunctionCodeRequest) (_result *GetFunctionCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionCodeResponse{}
	_body, _err := client.GetFunctionCodeWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries versions of a layer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerVersionResponse
func (client *Client) GetLayerVersionWithOptions(layerName *string, version *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLayerVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayerVersion"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers/" + tea.StringValue(openapiutil.GetEncodeParam(layerName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(version))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLayerVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of a layer.
//
// @return GetLayerVersionResponse
func (client *Client) GetLayerVersion(layerName *string, version *string) (_result *GetLayerVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLayerVersionResponse{}
	_body, _err := client.GetLayerVersionWithOptions(layerName, version, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain version information of a layer by using ARNs.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLayerVersionByArnResponse
func (client *Client) GetLayerVersionByArnWithOptions(arn *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLayerVersionByArnResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLayerVersionByArn"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layerarn/" + tea.StringValue(openapiutil.GetEncodeParam(arn))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLayerVersionByArnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain version information of a layer by using ARNs.
//
// @return GetLayerVersionByArnResponse
func (client *Client) GetLayerVersionByArn(arn *string) (_result *GetLayerVersionByArnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLayerVersionByArnResponse{}
	_body, _err := client.GetLayerVersionByArnWithOptions(arn, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries provisioned configurations.
//
// @param request - GetProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProvisionConfigResponse
func (client *Client) GetProvisionConfigWithOptions(functionName *string, request *GetProvisionConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProvisionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProvisionConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/provision-config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProvisionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries provisioned configurations.
//
// @param request - GetProvisionConfigRequest
//
// @return GetProvisionConfigResponse
func (client *Client) GetProvisionConfig(functionName *string, request *GetProvisionConfigRequest) (_result *GetProvisionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProvisionConfigResponse{}
	_body, _err := client.GetProvisionConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a trigger.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTriggerResponse
func (client *Client) GetTriggerWithOptions(functionName *string, triggerName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTriggerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrigger"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/triggers/" + tea.StringValue(openapiutil.GetEncodeParam(triggerName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a trigger.
//
// @return GetTriggerResponse
func (client *Client) GetTrigger(functionName *string, triggerName *string) (_result *GetTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTriggerResponse{}
	_body, _err := client.GetTriggerWithOptions(functionName, triggerName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes a function.
//
// @param request - InvokeFunctionRequest
//
// @param headers - InvokeFunctionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeFunctionResponse
func (client *Client) InvokeFunctionWithOptions(functionName *string, request *InvokeFunctionRequest, headers *InvokeFunctionHeaders, runtime *util.RuntimeOptions) (_result *InvokeFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcAsyncTaskId)) {
		realHeaders["x-fc-async-task-id"] = util.ToJSONString(headers.XFcAsyncTaskId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcInvocationType)) {
		realHeaders["x-fc-invocation-type"] = util.ToJSONString(headers.XFcInvocationType)
	}

	if !tea.BoolValue(util.IsUnset(headers.XFcLogType)) {
		realHeaders["x-fc-log-type"] = util.ToJSONString(headers.XFcLogType)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
		Stream:  request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeFunction"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/invocations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("binary"),
	}
	res := &InvokeFunctionResponse{}
	callApiTmp, err := client.CallApi(params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp, _err := util.AssertAsMap(callApiTmp)
	if _err != nil {
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(tmp["body"])) {
		respBody, _err := util.AssertAsReadable(tmp["body"])
		if _err != nil {
			return _result, _err
		}

		res.Body = respBody
	}

	if !tea.BoolValue(util.IsUnset(tmp["headers"])) {
		respHeaders, _err := util.AssertAsMap(tmp["headers"])
		if _err != nil {
			return _result, _err
		}

		res.Headers = util.StringifyMapValue(respHeaders)
	}

	if !tea.BoolValue(util.IsUnset(tmp["statusCode"])) {
		statusCode, _err := util.AssertAsInteger(tmp["statusCode"])
		if _err != nil {
			return _result, _err
		}

		res.StatusCode = tea.ToInt32(statusCode)
	}

	_result = res
	return _result, _err
}

// Summary:
//
// Invokes a function.
//
// @param request - InvokeFunctionRequest
//
// @return InvokeFunctionResponse
func (client *Client) InvokeFunction(functionName *string, request *InvokeFunctionRequest) (_result *InvokeFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InvokeFunctionHeaders{}
	_result = &InvokeFunctionResponse{}
	_body, _err := client.InvokeFunctionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries aliases.
//
// @param request - ListAliasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliasesResponse
func (client *Client) ListAliasesWithOptions(functionName *string, request *ListAliasesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAliases"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/aliases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries aliases.
//
// @param request - ListAliasesRequest
//
// @return ListAliasesResponse
func (client *Client) ListAliases(functionName *string, request *ListAliasesRequest) (_result *ListAliasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAliasesResponse{}
	_body, _err := client.ListAliasesWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all asynchronous configurations of a function.
//
// @param request - ListAsyncInvokeConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncInvokeConfigsResponse
func (client *Client) ListAsyncInvokeConfigsWithOptions(request *ListAsyncInvokeConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAsyncInvokeConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["functionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsyncInvokeConfigs"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/async-invoke-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsyncInvokeConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all asynchronous configurations of a function.
//
// @param request - ListAsyncInvokeConfigsRequest
//
// @return ListAsyncInvokeConfigsResponse
func (client *Client) ListAsyncInvokeConfigs(request *ListAsyncInvokeConfigsRequest) (_result *ListAsyncInvokeConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAsyncInvokeConfigsResponse{}
	_body, _err := client.ListAsyncInvokeConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists asynchronous tasks.
//
// @param request - ListAsyncTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasksWithOptions(functionName *string, request *ListAsyncTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludePayload)) {
		query["includePayload"] = request.IncludePayload
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrderByTime)) {
		query["sortOrderByTime"] = request.SortOrderByTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartedTimeBegin)) {
		query["startedTimeBegin"] = request.StartedTimeBegin
	}

	if !tea.BoolValue(util.IsUnset(request.StartedTimeEnd)) {
		query["startedTimeEnd"] = request.StartedTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsyncTasks"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists asynchronous tasks.
//
// @param request - ListAsyncTasksRequest
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasks(functionName *string, request *ListAsyncTasksRequest) (_result *ListAsyncTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.ListAsyncTasksWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出函数并发度配置。
//
// @param request - ListConcurrencyConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConcurrencyConfigsResponse
func (client *Client) ListConcurrencyConfigsWithOptions(request *ListConcurrencyConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConcurrencyConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["functionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConcurrencyConfigs"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/concurrency-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConcurrencyConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出函数并发度配置。
//
// @param request - ListConcurrencyConfigsRequest
//
// @return ListConcurrencyConfigsResponse
func (client *Client) ListConcurrencyConfigs(request *ListConcurrencyConfigsRequest) (_result *ListConcurrencyConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConcurrencyConfigsResponse{}
	_body, _err := client.ListConcurrencyConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom domain names.
//
// @param request - ListCustomDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomainsWithOptions(request *ListCustomDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCustomDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomDomains"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/custom-domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom domain names.
//
// @param request - ListCustomDomainsRequest
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomains(request *ListCustomDomainsRequest) (_result *ListCustomDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.ListCustomDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries versions of a function.
//
// @param request - ListFunctionVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionVersionsResponse
func (client *Client) ListFunctionVersionsWithOptions(functionName *string, request *ListFunctionVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionVersions"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of a function.
//
// @param request - ListFunctionVersionsRequest
//
// @return ListFunctionVersionsResponse
func (client *Client) ListFunctionVersions(functionName *string, request *ListFunctionVersionsRequest) (_result *ListFunctionVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionVersionsResponse{}
	_body, _err := client.ListFunctionVersionsWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出函数。
//
// @param request - ListFunctionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FcVersion)) {
		query["fcVersion"] = request.FcVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出函数。
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of function instances.
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(functionName *string, request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	if !tea.BoolValue(util.IsUnset(request.WithAllActive)) {
		query["withAllActive"] = request.WithAllActive
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of function instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(functionName *string, request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a list of layer versions.
//
// @param request - ListLayerVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayerVersionsResponse
func (client *Client) ListLayerVersionsWithOptions(layerName *string, request *ListLayerVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLayerVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.StartVersion)) {
		query["startVersion"] = request.StartVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayerVersions"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers/" + tea.StringValue(openapiutil.GetEncodeParam(layerName)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLayerVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a list of layer versions.
//
// @param request - ListLayerVersionsRequest
//
// @return ListLayerVersionsResponse
func (client *Client) ListLayerVersions(layerName *string, request *ListLayerVersionsRequest) (_result *ListLayerVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLayerVersionsResponse{}
	_body, _err := client.ListLayerVersionsWithOptions(layerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a list of layers.
//
// @param request - ListLayersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayersResponse
func (client *Client) ListLayersWithOptions(request *ListLayersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLayersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Official)) {
		query["official"] = request.Official
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.Public)) {
		query["public"] = request.Public
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayers"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLayersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a list of layers.
//
// @param request - ListLayersRequest
//
// @return ListLayersResponse
func (client *Client) ListLayers(request *ListLayersRequest) (_result *ListLayersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLayersResponse{}
	_body, _err := client.ListLayersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of provisioned configurations.
//
// @param request - ListProvisionConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProvisionConfigsResponse
func (client *Client) ListProvisionConfigsWithOptions(request *ListProvisionConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProvisionConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["functionName"] = request.FunctionName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProvisionConfigs"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/provision-configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProvisionConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of provisioned configurations.
//
// @param request - ListProvisionConfigsRequest
//
// @return ListProvisionConfigsResponse
func (client *Client) ListProvisionConfigs(request *ListProvisionConfigsRequest) (_result *ListProvisionConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProvisionConfigsResponse{}
	_body, _err := client.ListProvisionConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all tagged resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("ResourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tags-v2"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all tagged resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the triggers of a function.
//
// @param request - ListTriggersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTriggersResponse
func (client *Client) ListTriggersWithOptions(functionName *string, request *ListTriggersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTriggersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTriggers"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/triggers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTriggersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the triggers of a function.
//
// @param request - ListTriggersRequest
//
// @return ListTriggersResponse
func (client *Client) ListTriggers(functionName *string, request *ListTriggersRequest) (_result *ListTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTriggersResponse{}
	_body, _err := client.ListTriggersWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of existing VPC connections.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcBindingsResponse
func (client *Client) ListVpcBindingsWithOptions(functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVpcBindingsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListVpcBindings"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/vpc-bindings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVpcBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of existing VPC connections.
//
// @return ListVpcBindingsResponse
func (client *Client) ListVpcBindings(functionName *string) (_result *ListVpcBindingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVpcBindingsResponse{}
	_body, _err := client.ListVpcBindingsWithOptions(functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a function version.
//
// @param request - PublishFunctionVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFunctionVersionResponse
func (client *Client) PublishFunctionVersionWithOptions(functionName *string, request *PublishFunctionVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishFunctionVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishFunctionVersion"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishFunctionVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a function version.
//
// @param request - PublishFunctionVersionRequest
//
// @return PublishFunctionVersionResponse
func (client *Client) PublishFunctionVersion(functionName *string, request *PublishFunctionVersionRequest) (_result *PublishFunctionVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishFunctionVersionResponse{}
	_body, _err := client.PublishFunctionVersionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an asynchronous invocation configuration for a function.
//
// @param request - PutAsyncInvokeConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutAsyncInvokeConfigResponse
func (client *Client) PutAsyncInvokeConfigWithOptions(functionName *string, request *PutAsyncInvokeConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutAsyncInvokeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutAsyncInvokeConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-invoke-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutAsyncInvokeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an asynchronous invocation configuration for a function.
//
// @param request - PutAsyncInvokeConfigRequest
//
// @return PutAsyncInvokeConfigResponse
func (client *Client) PutAsyncInvokeConfig(functionName *string, request *PutAsyncInvokeConfigRequest) (_result *PutAsyncInvokeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutAsyncInvokeConfigResponse{}
	_body, _err := client.PutAsyncInvokeConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures concurrency of a function.
//
// @param request - PutConcurrencyConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutConcurrencyConfigResponse
func (client *Client) PutConcurrencyConfigWithOptions(functionName *string, request *PutConcurrencyConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutConcurrencyConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutConcurrencyConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/concurrency"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutConcurrencyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures concurrency of a function.
//
// @param request - PutConcurrencyConfigRequest
//
// @return PutConcurrencyConfigResponse
func (client *Client) PutConcurrencyConfig(functionName *string, request *PutConcurrencyConfigRequest) (_result *PutConcurrencyConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutConcurrencyConfigResponse{}
	_body, _err := client.PutConcurrencyConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies permissions of a layer.
//
// @param request - PutLayerACLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutLayerACLResponse
func (client *Client) PutLayerACLWithOptions(layerName *string, request *PutLayerACLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutLayerACLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Acl)) {
		query["acl"] = request.Acl
	}

	if !tea.BoolValue(util.IsUnset(request.Public)) {
		query["public"] = request.Public
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutLayerACL"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/layers/" + tea.StringValue(openapiutil.GetEncodeParam(layerName)) + "/acl"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &PutLayerACLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies permissions of a layer.
//
// @param request - PutLayerACLRequest
//
// @return PutLayerACLResponse
func (client *Client) PutLayerACL(layerName *string, request *PutLayerACLRequest) (_result *PutLayerACLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutLayerACLResponse{}
	_body, _err := client.PutLayerACLWithOptions(layerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates provisioned configurations.
//
// @param request - PutProvisionConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutProvisionConfigResponse
func (client *Client) PutProvisionConfigWithOptions(functionName *string, request *PutProvisionConfigRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutProvisionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutProvisionConfig"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/provision-config"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutProvisionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates provisioned configurations.
//
// @param request - PutProvisionConfigRequest
//
// @return PutProvisionConfigResponse
func (client *Client) PutProvisionConfig(functionName *string, request *PutProvisionConfigRequest) (_result *PutProvisionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutProvisionConfigResponse{}
	_body, _err := client.PutProvisionConfigWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an asynchronous task.
//
// @param request - StopAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAsyncTaskResponse
func (client *Client) StopAsyncTaskWithOptions(functionName *string, taskId *string, request *StopAsyncTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Qualifier)) {
		query["qualifier"] = request.Qualifier
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAsyncTask"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/async-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &StopAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an asynchronous task.
//
// @param request - StopAsyncTaskRequest
//
// @return StopAsyncTaskResponse
func (client *Client) StopAsyncTask(functionName *string, taskId *string, request *StopAsyncTaskRequest) (_result *StopAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAsyncTaskResponse{}
	_body, _err := client.StopAsyncTaskWithOptions(functionName, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to a resource.
//
// Description:
//
// Tags are used to identify resources. Tags allow you to categorize, search for, and aggregate resources that have the same characteristics from different dimensions. This facilitates resource management. For more information, see [Tag overview](https://help.aliyun.com/document_detail/156983.html).
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tags-v2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to a resource.
//
// Description:
//
// Tags are used to identify resources. Tags allow you to categorize, search for, and aggregate resources that have the same characteristics from different dimensions. This facilitates resource management. For more information, see [Tag overview](https://help.aliyun.com/document_detail/156983.html).
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("ResourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKey)) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, tea.String("TagKey"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyShrink)) {
		query["TagKey"] = request.TagKeyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tags-v2"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an alias.
//
// @param request - UpdateAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithOptions(functionName *string, aliasName *string, request *UpdateAliasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlias"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/aliases/" + tea.StringValue(openapiutil.GetEncodeParam(aliasName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alias.
//
// @param request - UpdateAliasRequest
//
// @return UpdateAliasResponse
func (client *Client) UpdateAlias(functionName *string, aliasName *string, request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(functionName, aliasName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a custom domain name.
//
// @param request - UpdateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomainWithOptions(domainName *string, request *UpdateCustomDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCustomDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomDomain"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/custom-domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a custom domain name.
//
// @param request - UpdateCustomDomainRequest
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomain(domainName *string, request *UpdateCustomDomainRequest) (_result *UpdateCustomDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.UpdateCustomDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a function.
//
// @param request - UpdateFunctionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithOptions(functionName *string, request *UpdateFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a function.
//
// @param request - UpdateFunctionRequest
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunction(functionName *string, request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a trigger.
//
// @param request - UpdateTriggerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTriggerWithOptions(functionName *string, triggerName *string, request *UpdateTriggerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTrigger"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/triggers/" + tea.StringValue(openapiutil.GetEncodeParam(triggerName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a trigger.
//
// @param request - UpdateTriggerRequest
//
// @return UpdateTriggerResponse
func (client *Client) UpdateTrigger(functionName *string, triggerName *string, request *UpdateTriggerRequest) (_result *UpdateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTriggerResponse{}
	_body, _err := client.UpdateTriggerWithOptions(functionName, triggerName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
