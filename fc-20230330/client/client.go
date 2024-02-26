// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
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
	AliasName               *string             `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	CreatedTime             *string             `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description             *string             `json:"description,omitempty" xml:"description,omitempty"`
	LastModifiedTime        *string             `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	VersionId               *string             `json:"versionId,omitempty" xml:"versionId,omitempty"`
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
	CreatedTime               *string            `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DestinationConfig         *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	FunctionArn               *string            `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	LastModifiedTime          *string            `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	MaxAsyncEventAgeInSeconds *int64             `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	MaxAsyncRetryAttempts     *int64             `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s AsyncConfig) String() string {
	return tea.Prettify(s)
}

func (s AsyncConfig) GoString() string {
	return s.String()
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

type AuthConfig struct {
	AuthInfo *string `json:"authInfo,omitempty" xml:"authInfo,omitempty"`
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

type CertConfig struct {
	CertName    *string `json:"certName,omitempty" xml:"certName,omitempty"`
	Certificate *string `json:"certificate,omitempty" xml:"certificate,omitempty"`
	PrivateKey  *string `json:"privateKey,omitempty" xml:"privateKey,omitempty"`
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
	FunctionArn         *string `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	ReservedConcurrency *int64  `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
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
	AliasName               *string             `json:"aliasName,omitempty" xml:"aliasName,omitempty"`
	Description             *string             `json:"description,omitempty" xml:"description,omitempty"`
	VersionId               *string             `json:"versionId,omitempty" xml:"versionId,omitempty"`
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
	AuthConfig  *AuthConfig  `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig  *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	DomainName  *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
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
	Code                    *InputCodeLocation       `json:"code,omitempty" xml:"code,omitempty"`
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig   *CustomContainerConfig   `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS               *CustomDNS               `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig     *CustomRuntimeConfig     `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables" xml:"environmentVariables"`
	FunctionName            *string                  `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig               *GPUConfig               `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	Handler                 *string                  `json:"handler,omitempty" xml:"handler,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InternetAccess          *bool                    `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers                  []*string                `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig               *LogConfig               `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	MemorySize              *int32                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig               *NASConfig               `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig          *OSSMountConfig          `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	Role                    *string                  `json:"role,omitempty" xml:"role,omitempty"`
	Runtime                 *string                  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	Timeout                 *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig           *TracingConfig           `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig               *VPCConfig               `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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
	Description       *string            `json:"description,omitempty" xml:"description,omitempty"`
	License           *string            `json:"license,omitempty" xml:"license,omitempty"`
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
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	Qualifier      *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	SourceArn      *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	TriggerConfig  *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerName    *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	TriggerType    *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
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
	AccelerationInfo  *AccelerationInfo        `json:"accelerationInfo,omitempty" xml:"accelerationInfo,omitempty"`
	AccelerationType  *string                  `json:"accelerationType,omitempty" xml:"accelerationType,omitempty"`
	AcrInstanceId     *string                  `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	Command           []*string                `json:"command" xml:"command" type:"Repeated"`
	Entrypoint        []*string                `json:"entrypoint" xml:"entrypoint" type:"Repeated"`
	HealthCheckConfig *CustomHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty"`
	Image             *string                  `json:"image,omitempty" xml:"image,omitempty"`
	Port              *int32                   `json:"port,omitempty" xml:"port,omitempty"`
	ResolvedImageUri  *string                  `json:"resolvedImageUri,omitempty" xml:"resolvedImageUri,omitempty"`
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
	AccountId        *string      `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion       *string      `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	AuthConfig       *AuthConfig  `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig       *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	SubdomainCount   *string      `json:"subdomainCount,omitempty" xml:"subdomainCount,omitempty"`
	TlsConfig        *TLSConfig   `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig        *WAFConfig   `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
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
	FailureThreshold    *int32  `json:"failureThreshold,omitempty" xml:"failureThreshold,omitempty"`
	HttpGetUrl          *string `json:"httpGetUrl,omitempty" xml:"httpGetUrl,omitempty"`
	InitialDelaySeconds *int32  `json:"initialDelaySeconds,omitempty" xml:"initialDelaySeconds,omitempty"`
	PeriodSeconds       *int32  `json:"periodSeconds,omitempty" xml:"periodSeconds,omitempty"`
	SuccessThreshold    *int32  `json:"successThreshold,omitempty" xml:"successThreshold,omitempty"`
	TimeoutSeconds      *int32  `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
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
	Port              *int32                   `json:"port,omitempty" xml:"port,omitempty"`
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
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
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

type Destination struct {
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
	Match       *string `json:"match,omitempty" xml:"match,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

type Function struct {
	CodeChecksum               *string                  `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	CodeSize                   *int64                   `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	Cpu                        *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CreatedTime                *string                  `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	CustomContainerConfig      *CustomContainerConfig   `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS                  *CustomDNS               `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig        *CustomRuntimeConfig     `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	Description                *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                   *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables       map[string]*string       `json:"environmentVariables" xml:"environmentVariables"`
	FunctionArn                *string                  `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	FunctionId                 *string                  `json:"functionId,omitempty" xml:"functionId,omitempty"`
	FunctionName               *string                  `json:"functionName,omitempty" xml:"functionName,omitempty"`
	GpuConfig                  *GPUConfig               `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	Handler                    *string                  `json:"handler,omitempty" xml:"handler,omitempty"`
	InstanceConcurrency        *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig    *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InternetAccess             *bool                    `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	LastModifiedTime           *string                  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LastUpdateStatus           *string                  `json:"lastUpdateStatus,omitempty" xml:"lastUpdateStatus,omitempty"`
	LastUpdateStatusReason     *string                  `json:"lastUpdateStatusReason,omitempty" xml:"lastUpdateStatusReason,omitempty"`
	LastUpdateStatusReasonCode *string                  `json:"lastUpdateStatusReasonCode,omitempty" xml:"lastUpdateStatusReasonCode,omitempty"`
	Layers                     []*FunctionLayer         `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig                  *LogConfig               `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	MemorySize                 *int32                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig                  *NASConfig               `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig             *OSSMountConfig          `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	Role                       *string                  `json:"role,omitempty" xml:"role,omitempty"`
	Runtime                    *string                  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	State                      *string                  `json:"state,omitempty" xml:"state,omitempty"`
	StateReason                *string                  `json:"stateReason,omitempty" xml:"stateReason,omitempty"`
	StateReasonCode            *string                  `json:"stateReasonCode,omitempty" xml:"stateReasonCode,omitempty"`
	Timeout                    *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig              *TracingConfig           `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig                  *VPCConfig               `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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
	Arn  *string `json:"arn,omitempty" xml:"arn,omitempty"`
	Size *int64  `json:"size,omitempty" xml:"size,omitempty"`
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
	GpuMemorySize *int32  `json:"gpuMemorySize,omitempty" xml:"gpuMemorySize,omitempty"`
	GpuType       *string `json:"gpuType,omitempty" xml:"gpuType,omitempty"`
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
	ResouceType *string            `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
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
	UrlInternet *string `json:"urlInternet,omitempty" xml:"urlInternet,omitempty"`
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

type InputCodeLocation struct {
	Checksum      *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	ZipFile       *string `json:"zipFile,omitempty" xml:"zipFile,omitempty"`
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
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	VersionId  *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
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

type Layer struct {
	Acl               *string             `json:"acl,omitempty" xml:"acl,omitempty"`
	Code              *OutputCodeLocation `json:"code,omitempty" xml:"code,omitempty"`
	CodeChecksum      *string             `json:"codeChecksum,omitempty" xml:"codeChecksum,omitempty"`
	CodeSize          *int64              `json:"codeSize,omitempty" xml:"codeSize,omitempty"`
	CompatibleRuntime []*string           `json:"compatibleRuntime" xml:"compatibleRuntime" type:"Repeated"`
	CreateTime        *string             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description       *string             `json:"description,omitempty" xml:"description,omitempty"`
	LayerName         *string             `json:"layerName,omitempty" xml:"layerName,omitempty"`
	LayerVersionArn   *string             `json:"layerVersionArn,omitempty" xml:"layerVersionArn,omitempty"`
	License           *string             `json:"license,omitempty" xml:"license,omitempty"`
	Version           *int32              `json:"version,omitempty" xml:"version,omitempty"`
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
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
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
	Aliases   []*Alias `json:"aliases" xml:"aliases" type:"Repeated"`
	NextToken *string  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Configs   []*AsyncConfig `json:"configs" xml:"configs" type:"Repeated"`
	NextToken *string        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

type ListConcurrencyConfigsOutput struct {
	Configs   []*ConcurrencyConfig `json:"configs" xml:"configs" type:"Repeated"`
	NextToken *string              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	NextToken     *string         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	NextToken *string     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Layers      []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	NextVersion *int32   `json:"nextVersion,omitempty" xml:"nextVersion,omitempty"`
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
	Layers    []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	NextToken *string  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Direction *string    `json:"direction,omitempty" xml:"direction,omitempty"`
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
	EnableInstanceMetrics *bool   `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableRequestMetrics  *bool   `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	LogBeginRule          *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	Logstore              *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	Project               *string `json:"project,omitempty" xml:"project,omitempty"`
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

type NASConfig struct {
	GroupId     *int32            `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MountPoints []*NASMountConfig `json:"mountPoints" xml:"mountPoints" type:"Repeated"`
	UserId      *int32            `json:"userId,omitempty" xml:"userId,omitempty"`
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
	EnableTLS  *bool   `json:"enableTLS,omitempty" xml:"enableTLS,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
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
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	BucketPath *string `json:"bucketPath,omitempty" xml:"bucketPath,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	MountDir   *string `json:"mountDir,omitempty" xml:"mountDir,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
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

type OutputCodeLocation struct {
	Location       *string `json:"location,omitempty" xml:"location,omitempty"`
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
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
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
	FunctionName  *string        `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Methods       []*string      `json:"methods" xml:"methods" type:"Repeated"`
	Path          *string        `json:"path,omitempty" xml:"path,omitempty"`
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
	AlwaysAllocateCPU      *bool                   `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	Current                *int64                  `json:"current,omitempty" xml:"current,omitempty"`
	CurrentError           *string                 `json:"currentError,omitempty" xml:"currentError,omitempty"`
	FunctionArn            *string                 `json:"functionArn,omitempty" xml:"functionArn,omitempty"`
	ScheduledActions       []*ScheduledAction      `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
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
	DestinationConfig         *DestinationConfig `json:"destinationConfig,omitempty" xml:"destinationConfig,omitempty"`
	MaxAsyncEventAgeInSeconds *int64             `json:"maxAsyncEventAgeInSeconds,omitempty" xml:"maxAsyncEventAgeInSeconds,omitempty"`
	MaxAsyncRetryAttempts     *int64             `json:"maxAsyncRetryAttempts,omitempty" xml:"maxAsyncRetryAttempts,omitempty"`
}

func (s PutAsyncInvokeConfigInput) String() string {
	return tea.Prettify(s)
}

func (s PutAsyncInvokeConfigInput) GoString() string {
	return s.String()
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
	AlwaysAllocateCPU      *bool                   `json:"alwaysAllocateCPU,omitempty" xml:"alwaysAllocateCPU,omitempty"`
	ScheduledActions       []*ScheduledAction      `json:"scheduledActions" xml:"scheduledActions" type:"Repeated"`
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
	Match       *string `json:"match,omitempty" xml:"match,omitempty"`
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

type Resource struct {
	ResouceType *string            `json:"resouceType,omitempty" xml:"resouceType,omitempty"`
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

type ScheduledAction struct {
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int64  `json:"target,omitempty" xml:"target,omitempty"`
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

type TLSConfig struct {
	CipherSuites []*string `json:"cipherSuites" xml:"cipherSuites" type:"Repeated"`
	MaxVersion   *string   `json:"maxVersion,omitempty" xml:"maxVersion,omitempty"`
	MinVersion   *string   `json:"minVersion,omitempty" xml:"minVersion,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	ResourceArn *string            `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	Tags        map[string]*string `json:"tags" xml:"tags"`
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
	ResourceId   []*string `json:"ResourceId" xml:"ResourceId" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*Tag    `json:"Tag" xml:"Tag" type:"Repeated"`
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
	EndTime      *string  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MaxCapacity  *int64   `json:"maxCapacity,omitempty" xml:"maxCapacity,omitempty"`
	MetricTarget *float32 `json:"metricTarget,omitempty" xml:"metricTarget,omitempty"`
	MetricType   *string  `json:"metricType,omitempty" xml:"metricType,omitempty"`
	MinCapacity  *int64   `json:"minCapacity,omitempty" xml:"minCapacity,omitempty"`
	Name         *string  `json:"name,omitempty" xml:"name,omitempty"`
	StartTime    *string  `json:"startTime,omitempty" xml:"startTime,omitempty"`
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

type TracingConfig struct {
	Params map[string]*string `json:"params" xml:"params"`
	Type   *string            `json:"type,omitempty" xml:"type,omitempty"`
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
	Description             *string             `json:"description,omitempty" xml:"description,omitempty"`
	VersionId               *string             `json:"versionId,omitempty" xml:"versionId,omitempty"`
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
	AuthConfig  *AuthConfig  `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig  *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
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
	Code                    *InputCodeLocation       `json:"code,omitempty" xml:"code,omitempty"`
	Cpu                     *float32                 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	CustomContainerConfig   *CustomContainerConfig   `json:"customContainerConfig,omitempty" xml:"customContainerConfig,omitempty"`
	CustomDNS               *CustomDNS               `json:"customDNS,omitempty" xml:"customDNS,omitempty"`
	CustomRuntimeConfig     *CustomRuntimeConfig     `json:"customRuntimeConfig,omitempty" xml:"customRuntimeConfig,omitempty"`
	Description             *string                  `json:"description,omitempty" xml:"description,omitempty"`
	DiskSize                *int32                   `json:"diskSize,omitempty" xml:"diskSize,omitempty"`
	EnvironmentVariables    map[string]*string       `json:"environmentVariables" xml:"environmentVariables"`
	GpuConfig               *GPUConfig               `json:"gpuConfig,omitempty" xml:"gpuConfig,omitempty"`
	Handler                 *string                  `json:"handler,omitempty" xml:"handler,omitempty"`
	InstanceConcurrency     *int32                   `json:"instanceConcurrency,omitempty" xml:"instanceConcurrency,omitempty"`
	InstanceLifecycleConfig *InstanceLifecycleConfig `json:"instanceLifecycleConfig,omitempty" xml:"instanceLifecycleConfig,omitempty"`
	InternetAccess          *bool                    `json:"internetAccess,omitempty" xml:"internetAccess,omitempty"`
	Layers                  []*string                `json:"layers" xml:"layers" type:"Repeated"`
	LogConfig               *LogConfig               `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	MemorySize              *int32                   `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	NasConfig               *NASConfig               `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig          *OSSMountConfig          `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	Role                    *string                  `json:"role,omitempty" xml:"role,omitempty"`
	Runtime                 *string                  `json:"runtime,omitempty" xml:"runtime,omitempty"`
	Timeout                 *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	TracingConfig           *TracingConfig           `json:"tracingConfig,omitempty" xml:"tracingConfig,omitempty"`
	VpcConfig               *VPCConfig               `json:"vpcConfig,omitempty" xml:"vpcConfig,omitempty"`
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
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	Qualifier      *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	TriggerConfig  *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
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
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds" xml:"vSwitchIds" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
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
	CreatedTime      *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	VersionId        *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
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
	Match       *string `json:"match,omitempty" xml:"match,omitempty"`
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

type GetResourceTagsRequest struct {
	// The resource ID.
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
}

func (s GetResourceTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTagsRequest) SetArn(v string) *GetResourceTagsRequest {
	s.Arn = &v
	return s
}

type GetResourceTagsResponse struct {
	Headers    map[string]*string     `json:"headers" xml:"headers"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceTagsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTagsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTagsResponse) SetHeaders(v map[string]*string) *GetResourceTagsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTagsResponse) SetStatusCode(v int32) *GetResourceTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTagsResponse) SetBody(v *GetResourceTagsOutput) *GetResourceTagsResponse {
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
	CommonHeaders     map[string]*string `json:"commonHeaders" xml:"commonHeaders"`
	XFcInvocationType *string            `json:"x-fc-invocation-type,omitempty" xml:"x-fc-invocation-type,omitempty"`
	XFcLogType        *string            `json:"x-fc-log-type,omitempty" xml:"x-fc-log-type,omitempty"`
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

func (s *InvokeFunctionHeaders) SetXFcInvocationType(v string) *InvokeFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcLogType(v string) *InvokeFunctionHeaders {
	s.XFcLogType = &v
	return s
}

type InvokeFunctionRequest struct {
	Body      io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	Qualifier *string   `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
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
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Prefix    *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
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
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Limit        *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

type ListConcurrencyConfigsRequest struct {
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Limit        *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Prefix    *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
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
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
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
	// The number of functions to return.
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// A pagination token.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// A prefix of function names.
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
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
	Qualifier     *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	WithAllActive *bool   `json:"withAllActive,omitempty" xml:"withAllActive,omitempty"`
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
	Limit        *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
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
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Official  *string `json:"official,omitempty" xml:"official,omitempty"`
	Prefix    *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	Public    *string `json:"public,omitempty" xml:"public,omitempty"`
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
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// Number of provisioned configurations to return.
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// A pagination token.
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

type ListTaggedResourcesRequest struct {
	// The number of resources to return.
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The type of the resource quantity. ALIYUN::FC::FUNCTION indicates functions in Function Compute 3.0. ALIYUN::FC::SERVICE indicates the original service of Function Compute. The default value is ALIYUN::FC:FUNCTION.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListTaggedResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesRequest) SetLimit(v int32) *ListTaggedResourcesRequest {
	s.Limit = &v
	return s
}

func (s *ListTaggedResourcesRequest) SetNextToken(v string) *ListTaggedResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaggedResourcesRequest) SetResourceType(v string) *ListTaggedResourcesRequest {
	s.ResourceType = &v
	return s
}

type ListTaggedResourcesResponse struct {
	Headers    map[string]*string         `json:"headers" xml:"headers"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaggedResourcesOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaggedResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaggedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTaggedResourcesResponse) SetHeaders(v map[string]*string) *ListTaggedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTaggedResourcesResponse) SetStatusCode(v int32) *ListTaggedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaggedResourcesResponse) SetBody(v *ListTaggedResourcesOutput) *ListTaggedResourcesResponse {
	s.Body = v
	return s
}

type ListTriggersRequest struct {
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Prefix    *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
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
	Body      *PutAsyncInvokeConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	Qualifier *string                    `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
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
	Public *string `json:"public,omitempty" xml:"public,omitempty"`
}

func (s PutLayerACLRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLayerACLRequest) GoString() string {
	return s.String()
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
	Body      *PutProvisionConfigInput `json:"body,omitempty" xml:"body,omitempty"`
	Qualifier *string                  `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
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

type TagResourceRequest struct {
	// The configuration of the resource tag.
	Body *TagResourceInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) SetBody(v *TagResourceInput) *TagResourceRequest {
	s.Body = v
	return s
}

type TagResourceResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s TagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourceResponse) GoString() string {
	return s.String()
}

func (s *TagResourceResponse) SetHeaders(v map[string]*string) *TagResourceResponse {
	s.Headers = v
	return s
}

func (s *TagResourceResponse) SetStatusCode(v int32) *TagResourceResponse {
	s.StatusCode = &v
	return s
}

type UntagResourceRequest struct {
	// Specifies whether to delete all tags.
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource ID.
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// The tag name.
	TagKeys *string `json:"tagKeys,omitempty" xml:"tagKeys,omitempty"`
}

func (s UntagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) SetAll(v bool) *UntagResourceRequest {
	s.All = &v
	return s
}

func (s *UntagResourceRequest) SetArn(v string) *UntagResourceRequest {
	s.Arn = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v string) *UntagResourceRequest {
	s.TagKeys = &v
	return s
}

type UntagResourceResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UntagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceResponse) GoString() string {
	return s.String()
}

func (s *UntagResourceResponse) SetHeaders(v map[string]*string) *UntagResourceResponse {
	s.Headers = v
	return s
}

func (s *UntagResourceResponse) SetStatusCode(v int32) *UntagResourceResponse {
	s.StatusCode = &v
	return s
}

type UpdateAliasRequest struct {
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

/**
 * @deprecated
 *
 * @param request GetResourceTagsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetResourceTagsResponse
 */
// Deprecated
func (client *Client) GetResourceTagsWithOptions(request *GetResourceTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Arn)) {
		query["arn"] = request.Arn
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceTags"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetResourceTagsRequest
 * @return GetResourceTagsResponse
 */
// Deprecated
func (client *Client) GetResourceTags(request *GetResourceTagsRequest) (_result *GetResourceTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceTagsResponse{}
	_body, _err := client.GetResourceTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
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

/**
 * @deprecated
 *
 * @param request ListTaggedResourcesRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTaggedResourcesResponse
 */
// Deprecated
func (client *Client) ListTaggedResourcesWithOptions(request *ListTaggedResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTaggedResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaggedResources"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaggedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ListTaggedResourcesRequest
 * @return ListTaggedResourcesResponse
 */
// Deprecated
func (client *Client) ListTaggedResources(request *ListTaggedResourcesRequest) (_result *ListTaggedResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTaggedResourcesResponse{}
	_body, _err := client.ListTaggedResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

func (client *Client) PutLayerACLWithOptions(layerName *string, request *PutLayerACLRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutLayerACLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

/**
 * @deprecated
 *
 * @param request TagResourceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourceResponse
 */
// Deprecated
func (client *Client) TagResourceWithOptions(request *TagResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResource"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &TagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request TagResourceRequest
 * @return TagResourceResponse
 */
// Deprecated
func (client *Client) TagResource(request *TagResourceRequest) (_result *TagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourceResponse{}
	_body, _err := client.TagResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request UntagResourceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourceResponse
 */
// Deprecated
func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.Arn)) {
		query["arn"] = request.Arn
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		query["tagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResource"),
		Version:     tea.String("2023-03-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/2023-03-30/tag"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UntagResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request UntagResourceRequest
 * @return UntagResourceResponse
 */
// Deprecated
func (client *Client) UntagResource(request *UntagResourceRequest) (_result *UntagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourceResponse{}
	_body, _err := client.UntagResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
