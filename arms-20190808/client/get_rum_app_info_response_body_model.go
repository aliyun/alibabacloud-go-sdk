// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetRumAppInfoResponseBody
	GetCode() *int32
	SetData(v *GetRumAppInfoResponseBodyData) *GetRumAppInfoResponseBody
	GetData() *GetRumAppInfoResponseBodyData
	SetHttpStatusCode(v string) *GetRumAppInfoResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetRumAppInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumAppInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRumAppInfoResponseBody
	GetSuccess() *bool
}

type GetRumAppInfoResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application details.
	Data *GetRumAppInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// StartTime is mandatory for this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRumAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetRumAppInfoResponseBody) GetData() *GetRumAppInfoResponseBodyData {
	return s.Data
}

func (s *GetRumAppInfoResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetRumAppInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumAppInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRumAppInfoResponseBody) SetCode(v int32) *GetRumAppInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumAppInfoResponseBody) SetData(v *GetRumAppInfoResponseBodyData) *GetRumAppInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetRumAppInfoResponseBody) SetHttpStatusCode(v string) *GetRumAppInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRumAppInfoResponseBody) SetMessage(v string) *GetRumAppInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumAppInfoResponseBody) SetRequestId(v string) *GetRumAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumAppInfoResponseBody) SetSuccess(v bool) *GetRumAppInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetRumAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyData struct {
	// The application configurations in the JSON format. This parameter is deprecated.
	//
	// example:
	//
	// {"apiRequestOfH5":300,"apiRequestOfOriginal":500,"coldStart":5000,"hotStart":3000,"staticResourceLoad":300,"stutter":1000,"viewLoadOfH5":1000,"viewLoadOfOriginal":2000}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The group to which the application belongs.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The application type. Valid values: web, miniapp, ios, and android.
	//
	// example:
	//
	// web
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The region where the backend is deployed.
	//
	// example:
	//
	// cn-hangzhou
	BackendServiceTraceRegion *string `json:"BackendServiceTraceRegion,omitempty" xml:"BackendServiceTraceRegion,omitempty"`
	// The collection configurations.
	BonreeSDKConfig *GetRumAppInfoResponseBodyDataBonreeSDKConfig `json:"BonreeSDKConfig,omitempty" xml:"BonreeSDKConfig,omitempty" type:"Struct"`
	// The domain name of the SDK.
	//
	// example:
	//
	// b59xxxxxxxx-sdk.rum.aliyuncs.com/v2/browser-sdk.js
	CdnDomain *string `json:"CdnDomain,omitempty" xml:"CdnDomain,omitempty"`
	// The time when the application was created. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1683353594000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Portal home page.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint that is used to report application data.
	//
	// example:
	//
	// xxxxxxxx-default-cn.rum.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Indicates whether the application is subscribed. Valid values: true and false.
	//
	// example:
	//
	// true
	IsSubscription *string `json:"IsSubscription,omitempty" xml:"IsSubscription,omitempty"`
	// The application name.
	//
	// example:
	//
	// tomcat-demo-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alias of the application.
	//
	// example:
	//
	// nickname
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The name of the application package.
	//
	// example:
	//
	// com.alibaba.rum
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The application ID.
	//
	// example:
	//
	// avccccefy0@24cccccbf384dc6
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2vezare****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of service domain configurations. Only mobile applications are supported.
	ServiceDomainConfigs []*GetRumAppInfoResponseBodyDataServiceDomainConfigs `json:"ServiceDomainConfigs,omitempty" xml:"ServiceDomainConfigs,omitempty" type:"Repeated"`
	// The name of the Simple Log Service Logstore that stores application data.
	//
	// example:
	//
	// logstore-rum
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the Simple Log Service project that stores application data.
	//
	// example:
	//
	// proj-xtrace-xxxxxxxxxxxxxxxxxxxxxxx-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The status of the application. Valid values: created, running, and stopped.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*GetRumAppInfoResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the application. Valid value: RUM.
	//
	// example:
	//
	// RUM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRumAppInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyData) GetAppConfig() *string {
	return s.AppConfig
}

func (s *GetRumAppInfoResponseBodyData) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetRumAppInfoResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *GetRumAppInfoResponseBodyData) GetBackendServiceTraceRegion() *string {
	return s.BackendServiceTraceRegion
}

func (s *GetRumAppInfoResponseBodyData) GetBonreeSDKConfig() *GetRumAppInfoResponseBodyDataBonreeSDKConfig {
	return s.BonreeSDKConfig
}

func (s *GetRumAppInfoResponseBodyData) GetCdnDomain() *string {
	return s.CdnDomain
}

func (s *GetRumAppInfoResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRumAppInfoResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetRumAppInfoResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetRumAppInfoResponseBodyData) GetIsSubscription() *string {
	return s.IsSubscription
}

func (s *GetRumAppInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetRumAppInfoResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *GetRumAppInfoResponseBodyData) GetPackageName() *string {
	return s.PackageName
}

func (s *GetRumAppInfoResponseBodyData) GetPid() *string {
	return s.Pid
}

func (s *GetRumAppInfoResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumAppInfoResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRumAppInfoResponseBodyData) GetServiceDomainConfigs() []*GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	return s.ServiceDomainConfigs
}

func (s *GetRumAppInfoResponseBodyData) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *GetRumAppInfoResponseBodyData) GetSlsProject() *string {
	return s.SlsProject
}

func (s *GetRumAppInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRumAppInfoResponseBodyData) GetTags() []*GetRumAppInfoResponseBodyDataTags {
	return s.Tags
}

func (s *GetRumAppInfoResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetRumAppInfoResponseBodyData) SetAppConfig(v string) *GetRumAppInfoResponseBodyData {
	s.AppConfig = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetAppGroup(v string) *GetRumAppInfoResponseBodyData {
	s.AppGroup = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetAppType(v string) *GetRumAppInfoResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetBackendServiceTraceRegion(v string) *GetRumAppInfoResponseBodyData {
	s.BackendServiceTraceRegion = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetBonreeSDKConfig(v *GetRumAppInfoResponseBodyDataBonreeSDKConfig) *GetRumAppInfoResponseBodyData {
	s.BonreeSDKConfig = v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetCdnDomain(v string) *GetRumAppInfoResponseBodyData {
	s.CdnDomain = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetCreateTime(v string) *GetRumAppInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetDescription(v string) *GetRumAppInfoResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetEndpoint(v string) *GetRumAppInfoResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetIsSubscription(v string) *GetRumAppInfoResponseBodyData {
	s.IsSubscription = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetName(v string) *GetRumAppInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetNickName(v string) *GetRumAppInfoResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetPackageName(v string) *GetRumAppInfoResponseBodyData {
	s.PackageName = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetPid(v string) *GetRumAppInfoResponseBodyData {
	s.Pid = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetRegionId(v string) *GetRumAppInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetResourceGroupId(v string) *GetRumAppInfoResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetServiceDomainConfigs(v []*GetRumAppInfoResponseBodyDataServiceDomainConfigs) *GetRumAppInfoResponseBodyData {
	s.ServiceDomainConfigs = v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetSlsLogstore(v string) *GetRumAppInfoResponseBodyData {
	s.SlsLogstore = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetSlsProject(v string) *GetRumAppInfoResponseBodyData {
	s.SlsProject = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetStatus(v string) *GetRumAppInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetTags(v []*GetRumAppInfoResponseBodyDataTags) *GetRumAppInfoResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetRumAppInfoResponseBodyData) SetType(v string) *GetRumAppInfoResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyDataBonreeSDKConfig struct {
	// The module configuration.
	ModuleConfig *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig `json:"moduleConfig,omitempty" xml:"moduleConfig,omitempty" type:"Struct"`
	// Sampling configuration.
	SamplingConfig *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig `json:"samplingConfig,omitempty" xml:"samplingConfig,omitempty" type:"Struct"`
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfig) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfig) GetModuleConfig() *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig {
	return s.ModuleConfig
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfig) GetSamplingConfig() *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig {
	return s.SamplingConfig
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfig) SetModuleConfig(v *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) *GetRumAppInfoResponseBodyDataBonreeSDKConfig {
	s.ModuleConfig = v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfig) SetSamplingConfig(v *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) *GetRumAppInfoResponseBodyDataBonreeSDKConfig {
	s.SamplingConfig = v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfig) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig struct {
	// The default configuration of the application.
	DefaultConfig map[string]*DataBonreeSDKConfigModuleConfigDefaultConfigValue `json:"defaultConfig,omitempty" xml:"defaultConfig,omitempty"`
	// Indicates whether the configuration is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The version configurations of the application.
	VersionConfigs map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValue `json:"versionConfigs,omitempty" xml:"versionConfigs,omitempty"`
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) GetDefaultConfig() map[string]*DataBonreeSDKConfigModuleConfigDefaultConfigValue {
	return s.DefaultConfig
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) GetEnable() *bool {
	return s.Enable
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) GetVersionConfigs() map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValue {
	return s.VersionConfigs
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) SetDefaultConfig(v map[string]*DataBonreeSDKConfigModuleConfigDefaultConfigValue) *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig {
	s.DefaultConfig = v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) SetEnable(v bool) *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig {
	s.Enable = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) SetVersionConfigs(v map[string]*DataBonreeSDKConfigModuleConfigVersionConfigsValue) *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig {
	s.VersionConfigs = v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig struct {
	// Sampling rate: between (0, 1000], a thousandth.
	//
	// example:
	//
	// 500
	SamplingRate *int32 `json:"samplingRate,omitempty" xml:"samplingRate,omitempty"`
	// Sampling type, currently only session random sampling is supported, that is, fixed transmission: 1.
	//
	// example:
	//
	// 1
	SamplingType *int32 `json:"samplingType,omitempty" xml:"samplingType,omitempty"`
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) GetSamplingType() *int32 {
	return s.SamplingType
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) SetSamplingRate(v int32) *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig {
	s.SamplingRate = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) SetSamplingType(v int32) *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig {
	s.SamplingType = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataBonreeSDKConfigSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyDataServiceDomainConfigs struct {
	// The description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name or IP address.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The trace propagation protocols. This parameter is required if the tracing analysis feature is enabled.
	PropagatorTypes []*string `json:"PropagatorTypes,omitempty" xml:"PropagatorTypes,omitempty" type:"Repeated"`
	// The sampling rate of a trace. Valid values: (0, 100].
	//
	// example:
	//
	// 100
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// Indicates whether the tracing analysis feature is enabled. To enable the tracing analysis feature, you must activate Managed Service for OpenTelemetry. Valid values:
	//
	// 	- `true`: enables the tracing analysis feature. If you enable the tracing analysis feature, related headers are inserted into requests for the domain name.
	//
	// 	- `false`: disables the tracing analysis feature.
	//
	// example:
	//
	// true
	Tracing *bool `json:"Tracing,omitempty" xml:"Tracing,omitempty"`
}

func (s GetRumAppInfoResponseBodyDataServiceDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyDataServiceDomainConfigs) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) GetDescription() *string {
	return s.Description
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) GetDomain() *string {
	return s.Domain
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) GetPropagatorTypes() []*string {
	return s.PropagatorTypes
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) GetTracing() *bool {
	return s.Tracing
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) SetDescription(v string) *GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	s.Description = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) SetDomain(v string) *GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	s.Domain = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) SetPropagatorTypes(v []*string) *GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	s.PropagatorTypes = v
	return s
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) SetSamplingRate(v int32) *GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	s.SamplingRate = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) SetTracing(v bool) *GetRumAppInfoResponseBodyDataServiceDomainConfigs {
	s.Tracing = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataServiceDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type GetRumAppInfoResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// Label
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRumAppInfoResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppInfoResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetRumAppInfoResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetRumAppInfoResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetRumAppInfoResponseBodyDataTags) SetKey(v string) *GetRumAppInfoResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataTags) SetValue(v string) *GetRumAppInfoResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetRumAppInfoResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
