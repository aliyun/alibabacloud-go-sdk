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
	// The HTTP status code. A `200` status code indicates a successful request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application.
	Data *GetRumAppInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned for a failed request.
	//
	// example:
	//
	// 内部错误，请联系管理员。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRumAppInfoResponseBodyData struct {
	// This parameter is deprecated. The legacy application configuration in the JSON format.
	//
	// example:
	//
	// {"apiRequestOfH5":300,"apiRequestOfOriginal":500,"coldStart":5000,"hotStart":3000,"staticResourceLoad":300,"stutter":1000,"viewLoadOfH5":1000,"viewLoadOfOriginal":2000}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The application group.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The application type. Valid values: `web`, `miniapp`, `ios`, and `android`. `web` indicates Web and H5 applications, `miniapp` indicates mini programs.
	//
	// example:
	//
	// web
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The region where the back-end application is deployed. This parameter is used for end-to-end tracing.
	//
	// example:
	//
	// cn-hangzhou
	BackendServiceTraceRegion *string `json:"BackendServiceTraceRegion,omitempty" xml:"BackendServiceTraceRegion,omitempty"`
	// The data collection configurations for mobile applications.
	BonreeSDKConfig *GetRumAppInfoResponseBodyDataBonreeSDKConfig `json:"BonreeSDKConfig,omitempty" xml:"BonreeSDKConfig,omitempty" type:"Struct"`
	// The SDK domain name.
	//
	// example:
	//
	// b59xxxxxxxx-sdk.rum.aliyuncs.com/v2/browser-sdk.js
	CdnDomain *string `json:"CdnDomain,omitempty" xml:"CdnDomain,omitempty"`
	// The creation time of the application. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1683353594000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// 门户首页。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint for reporting application data.
	//
	// example:
	//
	// xxxxxxxx-default-cn.rum.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Indicates whether the application is bookmarked. Valid values: `true` and `false`.
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
	// The application package name.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2vezare****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of service domain name configurations. This parameter is supported only for mobile applications.
	ServiceDomainConfigs []*GetRumAppInfoResponseBodyDataServiceDomainConfigs `json:"ServiceDomainConfigs,omitempty" xml:"ServiceDomainConfigs,omitempty" type:"Repeated"`
	// The name of the Log Service Logstore that is used to store application data.
	//
	// example:
	//
	// logstore-rum
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the Log Service project that is used to store application data.
	//
	// example:
	//
	// proj-xtrace-xxxxxxxxxxxxxxxxxxxxxxx-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The application status. Valid values: `created`, `running`, and `stopped`. `stopped` indicates that data reporting is stopped.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*GetRumAppInfoResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The application type. This parameter is a constant of `RUM`.
	//
	// example:
	//
	// RUM
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WebSDKConfigJson *string `json:"WebSDKConfigJson,omitempty" xml:"WebSDKConfigJson,omitempty"`
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

func (s *GetRumAppInfoResponseBodyData) GetWebSDKConfigJson() *string {
	return s.WebSDKConfigJson
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

func (s *GetRumAppInfoResponseBodyData) SetWebSDKConfigJson(v string) *GetRumAppInfoResponseBodyData {
	s.WebSDKConfigJson = &v
	return s
}

func (s *GetRumAppInfoResponseBodyData) Validate() error {
	if s.BonreeSDKConfig != nil {
		if err := s.BonreeSDKConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceDomainConfigs != nil {
		for _, item := range s.ServiceDomainConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRumAppInfoResponseBodyDataBonreeSDKConfig struct {
	// The feature switches for modules.
	ModuleConfig *GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig `json:"moduleConfig,omitempty" xml:"moduleConfig,omitempty" type:"Struct"`
	// The sampling configuration.
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
	if s.ModuleConfig != nil {
		if err := s.ModuleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SamplingConfig != nil {
		if err := s.SamplingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRumAppInfoResponseBodyDataBonreeSDKConfigModuleConfig struct {
	// The default configuration of the application.
	DefaultConfig map[string]*DataBonreeSDKConfigModuleConfigDefaultConfigValue `json:"defaultConfig,omitempty" xml:"defaultConfig,omitempty"`
	// The master switch.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The application version configurations.
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
	// The sampling rate, in parts per thousand. The value must be greater than 0 and less than or equal to 1,000.
	//
	// example:
	//
	// 500
	SamplingRate *int32 `json:"samplingRate,omitempty" xml:"samplingRate,omitempty"`
	// The sampling type. Only random session sampling is supported. You must set this parameter to `1`.
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
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The domain name or IP address.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of trace pass-through protocols. This parameter is required when trace tracking is enabled.
	PropagatorTypes []*string `json:"PropagatorTypes,omitempty" xml:"PropagatorTypes,omitempty" type:"Repeated"`
	// The trace sampling rate. Valid values: (0, 100].
	//
	// example:
	//
	// 100
	SamplingRate *int32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// Indicates whether to enable trace tracking. You must activate Application Real-Time Monitoring Service (ARMS) OpenTelemetry Edition to use this feature. Valid values:
	//
	// - `true`: enables trace tracking. If you set this parameter to true, a related header is inserted into the request for this domain name.
	//
	// - `false`: does not enable trace tracking.
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
	// The key of the tag.
	//
	// example:
	//
	// Label
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
