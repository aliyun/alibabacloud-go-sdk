// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GatewayOption struct {
	// 日志配置详情
	LogConfigDetails *GatewayOptionLogConfigDetails `json:"LogConfigDetails,omitempty" xml:"LogConfigDetails,omitempty" type:"Struct"`
	// xtrace config option
	TraceDetails *GatewayOptionTraceDetails `json:"TraceDetails,omitempty" xml:"TraceDetails,omitempty" type:"Struct"`
}

func (s GatewayOption) String() string {
	return tea.Prettify(s)
}

func (s GatewayOption) GoString() string {
	return s.String()
}

func (s *GatewayOption) SetLogConfigDetails(v *GatewayOptionLogConfigDetails) *GatewayOption {
	s.LogConfigDetails = v
	return s
}

func (s *GatewayOption) SetTraceDetails(v *GatewayOptionTraceDetails) *GatewayOption {
	s.TraceDetails = v
	return s
}

type GatewayOptionLogConfigDetails struct {
	// 是否开启日志投递
	LogEnabled *bool `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	// 投递的目标logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// 投递的目标project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GatewayOptionLogConfigDetails) String() string {
	return tea.Prettify(s)
}

func (s GatewayOptionLogConfigDetails) GoString() string {
	return s.String()
}

func (s *GatewayOptionLogConfigDetails) SetLogEnabled(v bool) *GatewayOptionLogConfigDetails {
	s.LogEnabled = &v
	return s
}

func (s *GatewayOptionLogConfigDetails) SetLogStoreName(v string) *GatewayOptionLogConfigDetails {
	s.LogStoreName = &v
	return s
}

func (s *GatewayOptionLogConfigDetails) SetProjectName(v string) *GatewayOptionLogConfigDetails {
	s.ProjectName = &v
	return s
}

type GatewayOptionTraceDetails struct {
	// trace 采样率
	Sample *int64 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// trace是否开启
	TraceEnabled *bool `json:"TraceEnabled,omitempty" xml:"TraceEnabled,omitempty"`
}

func (s GatewayOptionTraceDetails) String() string {
	return tea.Prettify(s)
}

func (s GatewayOptionTraceDetails) GoString() string {
	return s.String()
}

func (s *GatewayOptionTraceDetails) SetSample(v int64) *GatewayOptionTraceDetails {
	s.Sample = &v
	return s
}

func (s *GatewayOptionTraceDetails) SetTraceEnabled(v bool) *GatewayOptionTraceDetails {
	s.TraceEnabled = &v
	return s
}

type TrafficPolicy struct {
	// 负载均衡相关配置
	LoadBalancerSettings *TrafficPolicyLoadBalancerSettings `json:"LoadBalancerSettings,omitempty" xml:"LoadBalancerSettings,omitempty" type:"Struct"`
	// tls相关配置
	TlsSetting *TrafficPolicyTlsSetting `json:"TlsSetting,omitempty" xml:"TlsSetting,omitempty" type:"Struct"`
}

func (s TrafficPolicy) String() string {
	return tea.Prettify(s)
}

func (s TrafficPolicy) GoString() string {
	return s.String()
}

func (s *TrafficPolicy) SetLoadBalancerSettings(v *TrafficPolicyLoadBalancerSettings) *TrafficPolicy {
	s.LoadBalancerSettings = v
	return s
}

func (s *TrafficPolicy) SetTlsSetting(v *TrafficPolicyTlsSetting) *TrafficPolicy {
	s.TlsSetting = v
	return s
}

type TrafficPolicyLoadBalancerSettings struct {
	// 一致性hash相关配置
	ConsistentHashLBConfig *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig `json:"ConsistentHashLBConfig,omitempty" xml:"ConsistentHashLBConfig,omitempty" type:"Struct"`
	// 负载均衡类型，枚举类可为ROUND_ROBIN, LEAST_CONN,RANDOM, CONSISTENT_HASH
	LoadbalancerType *string `json:"LoadbalancerType,omitempty" xml:"LoadbalancerType,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettings) String() string {
	return tea.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettings) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettings) SetConsistentHashLBConfig(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) *TrafficPolicyLoadBalancerSettings {
	s.ConsistentHashLBConfig = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettings) SetLoadbalancerType(v string) *TrafficPolicyLoadBalancerSettings {
	s.LoadbalancerType = &v
	return s
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig struct {
	// HEADER, COOKIE, SOURCE_IP, QUERY_PARAMETER
	ConsistentHashLBType *string `json:"ConsistentHashLBType,omitempty" xml:"ConsistentHashLBType,omitempty"`
	// 使用cookie时配置
	HttpCookie *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie `json:"HttpCookie,omitempty" xml:"HttpCookie,omitempty" type:"Struct"`
	// 使用根据header和参数路由时生效
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) String() string {
	return tea.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetConsistentHashLBType(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ConsistentHashLBType = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetHttpCookie(v *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.HttpCookie = v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetParameterName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ParameterName = &v
	return s
}

type TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie struct {
	// cookie名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// cookie path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// cookie生命周期
	TTL *string `json:"TTL,omitempty" xml:"TTL,omitempty"`
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) String() string {
	return tea.Prettify(s)
}

func (s TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GoString() string {
	return s.String()
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetName(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Name = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetPath(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Path = &v
	return s
}

func (s *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetTTL(v string) *TrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.TTL = &v
	return s
}

type TrafficPolicyTlsSetting struct {
	// ca证书内容
	CaCertContent *string `json:"CaCertContent,omitempty" xml:"CaCertContent,omitempty"`
	// 使用的证书id，仅当为mutual时需要填写
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// 到后端服务些带
	Sni *string `json:"Sni,omitempty" xml:"Sni,omitempty"`
	// tls模式。为枚举类，可为NONE, SIMPLE, MUITUAL
	TlsMode *string `json:"TlsMode,omitempty" xml:"TlsMode,omitempty"`
}

func (s TrafficPolicyTlsSetting) String() string {
	return tea.Prettify(s)
}

func (s TrafficPolicyTlsSetting) GoString() string {
	return s.String()
}

func (s *TrafficPolicyTlsSetting) SetCaCertContent(v string) *TrafficPolicyTlsSetting {
	s.CaCertContent = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetCertId(v string) *TrafficPolicyTlsSetting {
	s.CertId = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetSni(v string) *TrafficPolicyTlsSetting {
	s.Sni = &v
	return s
}

func (s *TrafficPolicyTlsSetting) SetTlsMode(v string) *TrafficPolicyTlsSetting {
	s.TlsMode = &v
	return s
}

type AddMockRuleRequest struct {
	ConsumerAppIds  *string `json:"ConsumerAppIds,omitempty" xml:"ConsumerAppIds,omitempty"`
	DubboMockItems  *string `json:"DubboMockItems,omitempty" xml:"DubboMockItems,omitempty"`
	Enable          *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ExtraJson       *string `json:"ExtraJson,omitempty" xml:"ExtraJson,omitempty"`
	MockType        *int64  `json:"MockType,omitempty" xml:"MockType,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProviderAppId   *string `json:"ProviderAppId,omitempty" xml:"ProviderAppId,omitempty"`
	ProviderAppName *string `json:"ProviderAppName,omitempty" xml:"ProviderAppName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ScMockItems     *string `json:"ScMockItems,omitempty" xml:"ScMockItems,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddMockRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMockRuleRequest) GoString() string {
	return s.String()
}

func (s *AddMockRuleRequest) SetConsumerAppIds(v string) *AddMockRuleRequest {
	s.ConsumerAppIds = &v
	return s
}

func (s *AddMockRuleRequest) SetDubboMockItems(v string) *AddMockRuleRequest {
	s.DubboMockItems = &v
	return s
}

func (s *AddMockRuleRequest) SetEnable(v bool) *AddMockRuleRequest {
	s.Enable = &v
	return s
}

func (s *AddMockRuleRequest) SetExtraJson(v string) *AddMockRuleRequest {
	s.ExtraJson = &v
	return s
}

func (s *AddMockRuleRequest) SetMockType(v int64) *AddMockRuleRequest {
	s.MockType = &v
	return s
}

func (s *AddMockRuleRequest) SetName(v string) *AddMockRuleRequest {
	s.Name = &v
	return s
}

func (s *AddMockRuleRequest) SetProviderAppId(v string) *AddMockRuleRequest {
	s.ProviderAppId = &v
	return s
}

func (s *AddMockRuleRequest) SetProviderAppName(v string) *AddMockRuleRequest {
	s.ProviderAppName = &v
	return s
}

func (s *AddMockRuleRequest) SetRegion(v string) *AddMockRuleRequest {
	s.Region = &v
	return s
}

func (s *AddMockRuleRequest) SetScMockItems(v string) *AddMockRuleRequest {
	s.ScMockItems = &v
	return s
}

func (s *AddMockRuleRequest) SetSource(v string) *AddMockRuleRequest {
	s.Source = &v
	return s
}

type AddMockRuleResponseBody struct {
	Code           *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *AddMockRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMockRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMockRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponseBody) SetCode(v int32) *AddMockRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddMockRuleResponseBody) SetData(v *AddMockRuleResponseBodyData) *AddMockRuleResponseBody {
	s.Data = v
	return s
}

func (s *AddMockRuleResponseBody) SetHttpStatusCode(v int32) *AddMockRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddMockRuleResponseBody) SetMessage(v string) *AddMockRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddMockRuleResponseBody) SetRequestId(v string) *AddMockRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMockRuleResponseBody) SetSuccess(v bool) *AddMockRuleResponseBody {
	s.Success = &v
	return s
}

type AddMockRuleResponseBodyData struct {
	AccountId       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ConsumerAppId   *string `json:"ConsumerAppId,omitempty" xml:"ConsumerAppId,omitempty"`
	ConsumerAppName *string `json:"ConsumerAppName,omitempty" xml:"ConsumerAppName,omitempty"`
	Enable          *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	ExtraJson       *string `json:"ExtraJson,omitempty" xml:"ExtraJson,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MockType        *int64  `json:"MockType,omitempty" xml:"MockType,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId     *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	ProviderAppId   *string `json:"ProviderAppId,omitempty" xml:"ProviderAppId,omitempty"`
	ProviderAppName *string `json:"ProviderAppName,omitempty" xml:"ProviderAppName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ScMockItemJson  *string `json:"ScMockItemJson,omitempty" xml:"ScMockItemJson,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s AddMockRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddMockRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponseBodyData) SetAccountId(v string) *AddMockRuleResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetConsumerAppId(v string) *AddMockRuleResponseBodyData {
	s.ConsumerAppId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetConsumerAppName(v string) *AddMockRuleResponseBodyData {
	s.ConsumerAppName = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetEnable(v bool) *AddMockRuleResponseBodyData {
	s.Enable = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetExtraJson(v string) *AddMockRuleResponseBodyData {
	s.ExtraJson = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetId(v int64) *AddMockRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetMockType(v int64) *AddMockRuleResponseBodyData {
	s.MockType = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetName(v string) *AddMockRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetNamespaceId(v string) *AddMockRuleResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetProviderAppId(v string) *AddMockRuleResponseBodyData {
	s.ProviderAppId = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetProviderAppName(v string) *AddMockRuleResponseBodyData {
	s.ProviderAppName = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetRegion(v string) *AddMockRuleResponseBodyData {
	s.Region = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetScMockItemJson(v string) *AddMockRuleResponseBodyData {
	s.ScMockItemJson = &v
	return s
}

func (s *AddMockRuleResponseBodyData) SetSource(v string) *AddMockRuleResponseBodyData {
	s.Source = &v
	return s
}

type AddMockRuleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMockRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddMockRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMockRuleResponse) GoString() string {
	return s.String()
}

func (s *AddMockRuleResponse) SetHeaders(v map[string]*string) *AddMockRuleResponse {
	s.Headers = v
	return s
}

func (s *AddMockRuleResponse) SetBody(v *AddMockRuleResponseBody) *AddMockRuleResponse {
	s.Body = v
	return s
}

type AddServiceSourceRequest struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	GatewayId       *int64  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	Info1           *string `json:"Info1,omitempty" xml:"Info1,omitempty"`
	Info2           *string `json:"Info2,omitempty" xml:"Info2,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddServiceSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSourceRequest) GoString() string {
	return s.String()
}

func (s *AddServiceSourceRequest) SetAddress(v string) *AddServiceSourceRequest {
	s.Address = &v
	return s
}

func (s *AddServiceSourceRequest) SetGatewayId(v int64) *AddServiceSourceRequest {
	s.GatewayId = &v
	return s
}

func (s *AddServiceSourceRequest) SetGatewayUniqueId(v string) *AddServiceSourceRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddServiceSourceRequest) SetInfo1(v string) *AddServiceSourceRequest {
	s.Info1 = &v
	return s
}

func (s *AddServiceSourceRequest) SetInfo2(v string) *AddServiceSourceRequest {
	s.Info2 = &v
	return s
}

func (s *AddServiceSourceRequest) SetName(v string) *AddServiceSourceRequest {
	s.Name = &v
	return s
}

func (s *AddServiceSourceRequest) SetSource(v string) *AddServiceSourceRequest {
	s.Source = &v
	return s
}

func (s *AddServiceSourceRequest) SetType(v string) *AddServiceSourceRequest {
	s.Type = &v
	return s
}

type AddServiceSourceResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddServiceSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddServiceSourceResponseBody) SetCode(v int32) *AddServiceSourceResponseBody {
	s.Code = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetData(v int64) *AddServiceSourceResponseBody {
	s.Data = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetHttpStatusCode(v int32) *AddServiceSourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetMessage(v string) *AddServiceSourceResponseBody {
	s.Message = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetRequestId(v string) *AddServiceSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServiceSourceResponseBody) SetSuccess(v bool) *AddServiceSourceResponseBody {
	s.Success = &v
	return s
}

type AddServiceSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddServiceSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddServiceSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServiceSourceResponse) GoString() string {
	return s.String()
}

func (s *AddServiceSourceResponse) SetHeaders(v map[string]*string) *AddServiceSourceResponse {
	s.Headers = v
	return s
}

func (s *AddServiceSourceResponse) SetBody(v *AddServiceSourceResponseBody) *AddServiceSourceResponse {
	s.Body = v
	return s
}

type CloneNacosConfigRequest struct {
	Ids               *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OriginNamespaceId *string `json:"OriginNamespaceId,omitempty" xml:"OriginNamespaceId,omitempty"`
	Policy            *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	TargetNamespaceId *string `json:"TargetNamespaceId,omitempty" xml:"TargetNamespaceId,omitempty"`
}

func (s CloneNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigRequest) SetIds(v string) *CloneNacosConfigRequest {
	s.Ids = &v
	return s
}

func (s *CloneNacosConfigRequest) SetInstanceId(v string) *CloneNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneNacosConfigRequest) SetOriginNamespaceId(v string) *CloneNacosConfigRequest {
	s.OriginNamespaceId = &v
	return s
}

func (s *CloneNacosConfigRequest) SetPolicy(v string) *CloneNacosConfigRequest {
	s.Policy = &v
	return s
}

func (s *CloneNacosConfigRequest) SetTargetNamespaceId(v string) *CloneNacosConfigRequest {
	s.TargetNamespaceId = &v
	return s
}

type CloneNacosConfigResponseBody struct {
	Code           *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CloneNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloneNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBody) SetCode(v int32) *CloneNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetData(v *CloneNacosConfigResponseBodyData) *CloneNacosConfigResponseBody {
	s.Data = v
	return s
}

func (s *CloneNacosConfigResponseBody) SetDynamicMessage(v string) *CloneNacosConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetErrorCode(v string) *CloneNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetHttpStatusCode(v int32) *CloneNacosConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetMessage(v string) *CloneNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetRequestId(v string) *CloneNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneNacosConfigResponseBody) SetSuccess(v bool) *CloneNacosConfigResponseBody {
	s.Success = &v
	return s
}

type CloneNacosConfigResponseBodyData struct {
	FailData  []*CloneNacosConfigResponseBodyDataFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Repeated"`
	SkipCount *int32                                      `json:"SkipCount,omitempty" xml:"SkipCount,omitempty"`
	SkipData  []*CloneNacosConfigResponseBodyDataSkipData `json:"SkipData,omitempty" xml:"SkipData,omitempty" type:"Repeated"`
	SuccCount *int32                                      `json:"SuccCount,omitempty" xml:"SuccCount,omitempty"`
}

func (s CloneNacosConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyData) SetFailData(v []*CloneNacosConfigResponseBodyDataFailData) *CloneNacosConfigResponseBodyData {
	s.FailData = v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSkipCount(v int32) *CloneNacosConfigResponseBodyData {
	s.SkipCount = &v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSkipData(v []*CloneNacosConfigResponseBodyDataSkipData) *CloneNacosConfigResponseBodyData {
	s.SkipData = v
	return s
}

func (s *CloneNacosConfigResponseBodyData) SetSuccCount(v int32) *CloneNacosConfigResponseBodyData {
	s.SuccCount = &v
	return s
}

type CloneNacosConfigResponseBodyDataFailData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s CloneNacosConfigResponseBodyDataFailData) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigResponseBodyDataFailData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyDataFailData) SetDataId(v string) *CloneNacosConfigResponseBodyDataFailData {
	s.DataId = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataFailData) SetGroup(v string) *CloneNacosConfigResponseBodyDataFailData {
	s.Group = &v
	return s
}

type CloneNacosConfigResponseBodyDataSkipData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s CloneNacosConfigResponseBodyDataSkipData) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigResponseBodyDataSkipData) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponseBodyDataSkipData) SetDataId(v string) *CloneNacosConfigResponseBodyDataSkipData {
	s.DataId = &v
	return s
}

func (s *CloneNacosConfigResponseBodyDataSkipData) SetGroup(v string) *CloneNacosConfigResponseBodyDataSkipData {
	s.Group = &v
	return s
}

type CloneNacosConfigResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *CloneNacosConfigResponse) SetHeaders(v map[string]*string) *CloneNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *CloneNacosConfigResponse) SetBody(v *CloneNacosConfigResponseBody) *CloneNacosConfigResponse {
	s.Body = v
	return s
}

type CreateAlarmRuleRequest struct {
	Aggregates      *string                `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	AlarmAliasName  *string                `json:"AlarmAliasName,omitempty" xml:"AlarmAliasName,omitempty"`
	AlarmItem       *string                `json:"AlarmItem,omitempty" xml:"AlarmItem,omitempty"`
	AlertWay        map[string]interface{} `json:"AlertWay,omitempty" xml:"AlertWay,omitempty"`
	ContactGroupIds map[string]interface{} `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	InstanceId      *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NValue          *int32                 `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator        *string                `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value           *float32               `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRuleRequest) SetAggregates(v string) *CreateAlarmRuleRequest {
	s.Aggregates = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetAlarmAliasName(v string) *CreateAlarmRuleRequest {
	s.AlarmAliasName = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetAlarmItem(v string) *CreateAlarmRuleRequest {
	s.AlarmItem = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetAlertWay(v map[string]interface{}) *CreateAlarmRuleRequest {
	s.AlertWay = v
	return s
}

func (s *CreateAlarmRuleRequest) SetContactGroupIds(v map[string]interface{}) *CreateAlarmRuleRequest {
	s.ContactGroupIds = v
	return s
}

func (s *CreateAlarmRuleRequest) SetInstanceId(v string) *CreateAlarmRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetNValue(v int32) *CreateAlarmRuleRequest {
	s.NValue = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetOperator(v string) *CreateAlarmRuleRequest {
	s.Operator = &v
	return s
}

func (s *CreateAlarmRuleRequest) SetValue(v float32) *CreateAlarmRuleRequest {
	s.Value = &v
	return s
}

type CreateAlarmRuleShrinkRequest struct {
	Aggregates            *string  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	AlarmAliasName        *string  `json:"AlarmAliasName,omitempty" xml:"AlarmAliasName,omitempty"`
	AlarmItem             *string  `json:"AlarmItem,omitempty" xml:"AlarmItem,omitempty"`
	AlertWayShrink        *string  `json:"AlertWay,omitempty" xml:"AlertWay,omitempty"`
	ContactGroupIdsShrink *string  `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	InstanceId            *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NValue                *int32   `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator              *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value                 *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAlarmRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRuleShrinkRequest) SetAggregates(v string) *CreateAlarmRuleShrinkRequest {
	s.Aggregates = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetAlarmAliasName(v string) *CreateAlarmRuleShrinkRequest {
	s.AlarmAliasName = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetAlarmItem(v string) *CreateAlarmRuleShrinkRequest {
	s.AlarmItem = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetAlertWayShrink(v string) *CreateAlarmRuleShrinkRequest {
	s.AlertWayShrink = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetContactGroupIdsShrink(v string) *CreateAlarmRuleShrinkRequest {
	s.ContactGroupIdsShrink = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetInstanceId(v string) *CreateAlarmRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetNValue(v int32) *CreateAlarmRuleShrinkRequest {
	s.NValue = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetOperator(v string) *CreateAlarmRuleShrinkRequest {
	s.Operator = &v
	return s
}

func (s *CreateAlarmRuleShrinkRequest) SetValue(v float32) *CreateAlarmRuleShrinkRequest {
	s.Value = &v
	return s
}

type CreateAlarmRuleResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlarmRuleResponseBody) SetErrorCode(v string) *CreateAlarmRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAlarmRuleResponseBody) SetMessage(v string) *CreateAlarmRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlarmRuleResponseBody) SetRequestId(v string) *CreateAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlarmRuleResponseBody) SetSuccess(v bool) *CreateAlarmRuleResponseBody {
	s.Success = &v
	return s
}

type CreateAlarmRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAlarmRuleResponse) SetHeaders(v map[string]*string) *CreateAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAlarmRuleResponse) SetBody(v *CreateAlarmRuleResponseBody) *CreateAlarmRuleResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Language  *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetExtraInfo(v string) *CreateApplicationRequest {
	s.ExtraInfo = &v
	return s
}

func (s *CreateApplicationRequest) SetLanguage(v string) *CreateApplicationRequest {
	s.Language = &v
	return s
}

func (s *CreateApplicationRequest) SetRegion(v string) *CreateApplicationRequest {
	s.Region = &v
	return s
}

func (s *CreateApplicationRequest) SetSource(v string) *CreateApplicationRequest {
	s.Source = &v
	return s
}

type CreateApplicationResponseBody struct {
	Code           *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetCode(v int32) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationResponseBody) SetHttpStatusCode(v int32) *CreateApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetSuccess(v string) *CreateApplicationResponseBody {
	s.Success = &v
	return s
}

type CreateApplicationResponseBodyData struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtraInfo  *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Language   *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyData) SetAppId(v string) *CreateApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetAppName(v string) *CreateApplicationResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetCreateTime(v int64) *CreateApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetExtraInfo(v string) *CreateApplicationResponseBodyData {
	s.ExtraInfo = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetLanguage(v string) *CreateApplicationResponseBodyData {
	s.Language = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetLicenseKey(v string) *CreateApplicationResponseBodyData {
	s.LicenseKey = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetRegionId(v string) *CreateApplicationResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetSource(v string) *CreateApplicationResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetStatus(v int32) *CreateApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetUpdateTime(v int64) *CreateApplicationResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetUserId(v string) *CreateApplicationResponseBodyData {
	s.UserId = &v
	return s
}

type CreateApplicationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	ClusterType          *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ClusterVersion       *string `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	ConnectionType       *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	DiskCapacity         *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskType             *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	InstanceCount        *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// 用于区分基础/专业版本
	MseVersion              *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	NetType                 *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PrivateSlbSpecification *string `json:"PrivateSlbSpecification,omitempty" xml:"PrivateSlbSpecification,omitempty"`
	PubNetworkFlow          *string `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	PubSlbSpecification     *string `json:"PubSlbSpecification,omitempty" xml:"PubSlbSpecification,omitempty"`
	Region                  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestPars             *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClusterSpecification(v string) *CreateClusterRequest {
	s.ClusterSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetClusterVersion(v string) *CreateClusterRequest {
	s.ClusterVersion = &v
	return s
}

func (s *CreateClusterRequest) SetConnectionType(v string) *CreateClusterRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateClusterRequest) SetDiskCapacity(v int32) *CreateClusterRequest {
	s.DiskCapacity = &v
	return s
}

func (s *CreateClusterRequest) SetDiskType(v string) *CreateClusterRequest {
	s.DiskType = &v
	return s
}

func (s *CreateClusterRequest) SetInstanceCount(v int32) *CreateClusterRequest {
	s.InstanceCount = &v
	return s
}

func (s *CreateClusterRequest) SetMseVersion(v string) *CreateClusterRequest {
	s.MseVersion = &v
	return s
}

func (s *CreateClusterRequest) SetNetType(v string) *CreateClusterRequest {
	s.NetType = &v
	return s
}

func (s *CreateClusterRequest) SetPrivateSlbSpecification(v string) *CreateClusterRequest {
	s.PrivateSlbSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetPubNetworkFlow(v string) *CreateClusterRequest {
	s.PubNetworkFlow = &v
	return s
}

func (s *CreateClusterRequest) SetPubSlbSpecification(v string) *CreateClusterRequest {
	s.PubSlbSpecification = &v
	return s
}

func (s *CreateClusterRequest) SetRegion(v string) *CreateClusterRequest {
	s.Region = &v
	return s
}

func (s *CreateClusterRequest) SetRequestPars(v string) *CreateClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

type CreateClusterResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetErrorCode(v string) *CreateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetInstanceId(v string) *CreateClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateClusterResponseBody) SetMessage(v string) *CreateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateEngineNamespaceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Desc         *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceCount *int32  `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
}

func (s CreateEngineNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceRequest) SetClusterId(v string) *CreateEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetDesc(v string) *CreateEngineNamespaceRequest {
	s.Desc = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetInstanceId(v string) *CreateEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetName(v string) *CreateEngineNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateEngineNamespaceRequest) SetServiceCount(v int32) *CreateEngineNamespaceRequest {
	s.ServiceCount = &v
	return s
}

type CreateEngineNamespaceResponseBody struct {
	ClusterId *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Data      *CreateEngineNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEngineNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponseBody) SetClusterId(v string) *CreateEngineNamespaceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetData(v *CreateEngineNamespaceResponseBodyData) *CreateEngineNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetErrorCode(v string) *CreateEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetMessage(v string) *CreateEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetRequestId(v string) *CreateEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEngineNamespaceResponseBody) SetSuccess(v bool) *CreateEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

type CreateEngineNamespaceResponseBodyData struct {
	ConfigCount       *int32  `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceDesc     *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	Quota             *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	ServiceCount      *int32  `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEngineNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponseBodyData) SetConfigCount(v int32) *CreateEngineNamespaceResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespace(v string) *CreateEngineNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespaceDesc(v string) *CreateEngineNamespaceResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetNamespaceShowName(v string) *CreateEngineNamespaceResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetQuota(v int32) *CreateEngineNamespaceResponseBodyData {
	s.Quota = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetServiceCount(v int32) *CreateEngineNamespaceResponseBodyData {
	s.ServiceCount = &v
	return s
}

func (s *CreateEngineNamespaceResponseBodyData) SetType(v int32) *CreateEngineNamespaceResponseBodyData {
	s.Type = &v
	return s
}

type CreateEngineNamespaceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEngineNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateEngineNamespaceResponse) SetHeaders(v map[string]*string) *CreateEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateEngineNamespaceResponse) SetBody(v *CreateEngineNamespaceResponseBody) *CreateEngineNamespaceResponse {
	s.Body = v
	return s
}

type CreateGovernanceKubernetesClusterRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName    *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	K8sVersion     *string `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	NameSpaceInfos *string `json:"NameSpaceInfos,omitempty" xml:"NameSpaceInfos,omitempty"`
	PilotStartTime *int64  `json:"PilotStartTime,omitempty" xml:"PilotStartTime,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGovernanceKubernetesClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateGovernanceKubernetesClusterRequest) SetClusterId(v string) *CreateGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterRequest) SetClusterName(v string) *CreateGovernanceKubernetesClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterRequest) SetK8sVersion(v string) *CreateGovernanceKubernetesClusterRequest {
	s.K8sVersion = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterRequest) SetNameSpaceInfos(v string) *CreateGovernanceKubernetesClusterRequest {
	s.NameSpaceInfos = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterRequest) SetPilotStartTime(v int64) *CreateGovernanceKubernetesClusterRequest {
	s.PilotStartTime = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterRequest) SetRegionId(v string) *CreateGovernanceKubernetesClusterRequest {
	s.RegionId = &v
	return s
}

type CreateGovernanceKubernetesClusterResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGovernanceKubernetesClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetCode(v int32) *CreateGovernanceKubernetesClusterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetData(v int64) *CreateGovernanceKubernetesClusterResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetHttpStatusCode(v int32) *CreateGovernanceKubernetesClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetMessage(v string) *CreateGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *CreateGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponseBody) SetSuccess(v string) *CreateGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

type CreateGovernanceKubernetesClusterResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGovernanceKubernetesClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *CreateGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateGovernanceKubernetesClusterResponse) SetBody(v *CreateGovernanceKubernetesClusterResponseBody) *CreateGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

type CreateNacosConfigRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BetaIps     *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc        *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigRequest) SetAppName(v string) *CreateNacosConfigRequest {
	s.AppName = &v
	return s
}

func (s *CreateNacosConfigRequest) SetBetaIps(v string) *CreateNacosConfigRequest {
	s.BetaIps = &v
	return s
}

func (s *CreateNacosConfigRequest) SetContent(v string) *CreateNacosConfigRequest {
	s.Content = &v
	return s
}

func (s *CreateNacosConfigRequest) SetDataId(v string) *CreateNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetDesc(v string) *CreateNacosConfigRequest {
	s.Desc = &v
	return s
}

func (s *CreateNacosConfigRequest) SetGroup(v string) *CreateNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *CreateNacosConfigRequest) SetInstanceId(v string) *CreateNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetNamespaceId(v string) *CreateNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateNacosConfigRequest) SetTags(v string) *CreateNacosConfigRequest {
	s.Tags = &v
	return s
}

func (s *CreateNacosConfigRequest) SetType(v string) *CreateNacosConfigRequest {
	s.Type = &v
	return s
}

type CreateNacosConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigResponseBody) SetCode(v string) *CreateNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetErrorCode(v string) *CreateNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetHttpCode(v string) *CreateNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetMessage(v string) *CreateNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetRequestId(v string) *CreateNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNacosConfigResponseBody) SetSuccess(v bool) *CreateNacosConfigResponseBody {
	s.Success = &v
	return s
}

type CreateNacosConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateNacosConfigResponse) SetHeaders(v map[string]*string) *CreateNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateNacosConfigResponse) SetBody(v *CreateNacosConfigResponseBody) *CreateNacosConfigResponse {
	s.Body = v
	return s
}

type CreateZnodeRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateZnodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateZnodeRequest) GoString() string {
	return s.String()
}

func (s *CreateZnodeRequest) SetClusterId(v string) *CreateZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateZnodeRequest) SetData(v string) *CreateZnodeRequest {
	s.Data = &v
	return s
}

func (s *CreateZnodeRequest) SetPath(v string) *CreateZnodeRequest {
	s.Path = &v
	return s
}

type CreateZnodeResponseBody struct {
	Data      *CreateZnodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string                      `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateZnodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponseBody) SetData(v *CreateZnodeResponseBodyData) *CreateZnodeResponseBody {
	s.Data = v
	return s
}

func (s *CreateZnodeResponseBody) SetErrorCode(v string) *CreateZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateZnodeResponseBody) SetHttpCode(v string) *CreateZnodeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateZnodeResponseBody) SetMessage(v string) *CreateZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateZnodeResponseBody) SetRequestId(v string) *CreateZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateZnodeResponseBody) SetSuccess(v bool) *CreateZnodeResponseBody {
	s.Success = &v
	return s
}

type CreateZnodeResponseBodyData struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Dir  *bool   `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateZnodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateZnodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponseBodyData) SetData(v string) *CreateZnodeResponseBodyData {
	s.Data = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetDir(v bool) *CreateZnodeResponseBodyData {
	s.Dir = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetName(v string) *CreateZnodeResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateZnodeResponseBodyData) SetPath(v string) *CreateZnodeResponseBodyData {
	s.Path = &v
	return s
}

type CreateZnodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateZnodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateZnodeResponse) GoString() string {
	return s.String()
}

func (s *CreateZnodeResponse) SetHeaders(v map[string]*string) *CreateZnodeResponse {
	s.Headers = v
	return s
}

func (s *CreateZnodeResponse) SetBody(v *CreateZnodeResponseBody) *CreateZnodeResponse {
	s.Body = v
	return s
}

type DeleteAlarmRuleRequest struct {
	AlarmRuleId *string `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s DeleteAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRuleRequest) SetAlarmRuleId(v string) *DeleteAlarmRuleRequest {
	s.AlarmRuleId = &v
	return s
}

func (s *DeleteAlarmRuleRequest) SetRequestPars(v string) *DeleteAlarmRuleRequest {
	s.RequestPars = &v
	return s
}

type DeleteAlarmRuleResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRuleResponseBody) SetErrorCode(v string) *DeleteAlarmRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAlarmRuleResponseBody) SetHttpCode(v string) *DeleteAlarmRuleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteAlarmRuleResponseBody) SetMessage(v string) *DeleteAlarmRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlarmRuleResponseBody) SetRequestId(v string) *DeleteAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlarmRuleResponseBody) SetSuccess(v bool) *DeleteAlarmRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteAlarmRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRuleResponse) SetHeaders(v map[string]*string) *DeleteAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlarmRuleResponse) SetBody(v *DeleteAlarmRuleResponseBody) *DeleteAlarmRuleResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetInstanceId(v string) *DeleteClusterRequest {
	s.InstanceId = &v
	return s
}

type DeleteClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetErrorCode(v string) *DeleteClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteClusterResponseBody) SetHttpCode(v string) *DeleteClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetSuccess(v bool) *DeleteClusterResponseBody {
	s.Success = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteEngineNamespaceRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEngineNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceRequest) SetClusterId(v string) *DeleteEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) SetId(v string) *DeleteEngineNamespaceRequest {
	s.Id = &v
	return s
}

func (s *DeleteEngineNamespaceRequest) SetInstanceId(v string) *DeleteEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

type DeleteEngineNamespaceResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEngineNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceResponseBody) SetErrorCode(v string) *DeleteEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetHttpCode(v string) *DeleteEngineNamespaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetMessage(v string) *DeleteEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetRequestId(v string) *DeleteEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEngineNamespaceResponseBody) SetSuccess(v bool) *DeleteEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

type DeleteEngineNamespaceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEngineNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEngineNamespaceResponse) SetHeaders(v map[string]*string) *DeleteEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEngineNamespaceResponse) SetBody(v *DeleteEngineNamespaceResponseBody) *DeleteEngineNamespaceResponse {
	s.Body = v
	return s
}

type DeleteNacosConfigRequest struct {
	Beta        *bool   `json:"Beta,omitempty" xml:"Beta,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigRequest) SetBeta(v bool) *DeleteNacosConfigRequest {
	s.Beta = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetDataId(v string) *DeleteNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetGroup(v string) *DeleteNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetInstanceId(v string) *DeleteNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosConfigRequest) SetNamespaceId(v string) *DeleteNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

type DeleteNacosConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigResponseBody) SetCode(v string) *DeleteNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetErrorCode(v string) *DeleteNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetHttpCode(v string) *DeleteNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetMessage(v string) *DeleteNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetRequestId(v string) *DeleteNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosConfigResponseBody) SetSuccess(v bool) *DeleteNacosConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteNacosConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigResponse) SetHeaders(v map[string]*string) *DeleteNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosConfigResponse) SetBody(v *DeleteNacosConfigResponseBody) *DeleteNacosConfigResponse {
	s.Body = v
	return s
}

type DeleteNacosConfigsRequest struct {
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s DeleteNacosConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsRequest) SetIds(v string) *DeleteNacosConfigsRequest {
	s.Ids = &v
	return s
}

func (s *DeleteNacosConfigsRequest) SetInstanceId(v string) *DeleteNacosConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosConfigsRequest) SetNamespaceId(v string) *DeleteNacosConfigsRequest {
	s.NamespaceId = &v
	return s
}

type DeleteNacosConfigsResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNacosConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsResponseBody) SetCode(v int32) *DeleteNacosConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetErrorCode(v string) *DeleteNacosConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetHttpCode(v string) *DeleteNacosConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetMessage(v string) *DeleteNacosConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetRequestId(v string) *DeleteNacosConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosConfigsResponseBody) SetSuccess(v bool) *DeleteNacosConfigsResponseBody {
	s.Success = &v
	return s
}

type DeleteNacosConfigsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNacosConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNacosConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosConfigsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosConfigsResponse) SetHeaders(v map[string]*string) *DeleteNacosConfigsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosConfigsResponse) SetBody(v *DeleteNacosConfigsResponseBody) *DeleteNacosConfigsResponse {
	s.Body = v
	return s
}

type DeleteNacosServiceRequest struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DeleteNacosServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceRequest) SetGroupName(v string) *DeleteNacosServiceRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetInstanceId(v string) *DeleteNacosServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetNamespaceId(v string) *DeleteNacosServiceRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteNacosServiceRequest) SetServiceName(v string) *DeleteNacosServiceRequest {
	s.ServiceName = &v
	return s
}

type DeleteNacosServiceResponseBody struct {
	// 响应码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 删除服务的结果
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 响应信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNacosServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceResponseBody) SetCode(v int32) *DeleteNacosServiceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetData(v string) *DeleteNacosServiceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetHttpStatusCode(v int32) *DeleteNacosServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetMessage(v string) *DeleteNacosServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetRequestId(v string) *DeleteNacosServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosServiceResponseBody) SetSuccess(v bool) *DeleteNacosServiceResponseBody {
	s.Success = &v
	return s
}

type DeleteNacosServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNacosServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNacosServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNacosServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNacosServiceResponse) SetHeaders(v map[string]*string) *DeleteNacosServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNacosServiceResponse) SetBody(v *DeleteNacosServiceResponseBody) *DeleteNacosServiceResponse {
	s.Body = v
	return s
}

type DeleteZnodeRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s DeleteZnodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteZnodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteZnodeRequest) SetClusterId(v string) *DeleteZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteZnodeRequest) SetPath(v string) *DeleteZnodeRequest {
	s.Path = &v
	return s
}

func (s *DeleteZnodeRequest) SetRequestPars(v string) *DeleteZnodeRequest {
	s.RequestPars = &v
	return s
}

type DeleteZnodeResponseBody struct {
	Data      *DeleteZnodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string                      `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteZnodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponseBody) SetData(v *DeleteZnodeResponseBodyData) *DeleteZnodeResponseBody {
	s.Data = v
	return s
}

func (s *DeleteZnodeResponseBody) SetErrorCode(v string) *DeleteZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetHttpCode(v string) *DeleteZnodeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetMessage(v string) *DeleteZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetRequestId(v string) *DeleteZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZnodeResponseBody) SetSuccess(v bool) *DeleteZnodeResponseBody {
	s.Success = &v
	return s
}

type DeleteZnodeResponseBodyData struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Dir  *bool   `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DeleteZnodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteZnodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponseBodyData) SetData(v string) *DeleteZnodeResponseBodyData {
	s.Data = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetDir(v bool) *DeleteZnodeResponseBodyData {
	s.Dir = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetName(v string) *DeleteZnodeResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteZnodeResponseBodyData) SetPath(v string) *DeleteZnodeResponseBodyData {
	s.Path = &v
	return s
}

type DeleteZnodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteZnodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteZnodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteZnodeResponse) SetHeaders(v map[string]*string) *DeleteZnodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteZnodeResponse) SetBody(v *DeleteZnodeResponseBody) *DeleteZnodeResponse {
	s.Body = v
	return s
}

type ExportNacosConfigRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Ids         *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ExportNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *ExportNacosConfigRequest) SetAppName(v string) *ExportNacosConfigRequest {
	s.AppName = &v
	return s
}

func (s *ExportNacosConfigRequest) SetDataId(v string) *ExportNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *ExportNacosConfigRequest) SetGroup(v string) *ExportNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *ExportNacosConfigRequest) SetIds(v string) *ExportNacosConfigRequest {
	s.Ids = &v
	return s
}

func (s *ExportNacosConfigRequest) SetInstanceId(v string) *ExportNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ExportNacosConfigRequest) SetNamespaceId(v string) *ExportNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

type ExportNacosConfigResponseBody struct {
	Code           *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ExportNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ExportNacosConfigResponseBody) SetCode(v int32) *ExportNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetData(v *ExportNacosConfigResponseBodyData) *ExportNacosConfigResponseBody {
	s.Data = v
	return s
}

func (s *ExportNacosConfigResponseBody) SetDynamicMessage(v string) *ExportNacosConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetErrorCode(v string) *ExportNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetHttpStatusCode(v int32) *ExportNacosConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetMessage(v string) *ExportNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetRequestId(v string) *ExportNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportNacosConfigResponseBody) SetSuccess(v bool) *ExportNacosConfigResponseBody {
	s.Success = &v
	return s
}

type ExportNacosConfigResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ExportNacosConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExportNacosConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExportNacosConfigResponseBodyData) SetUrl(v string) *ExportNacosConfigResponseBodyData {
	s.Url = &v
	return s
}

type ExportNacosConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *ExportNacosConfigResponse) SetHeaders(v map[string]*string) *ExportNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *ExportNacosConfigResponse) SetBody(v *ExportNacosConfigResponseBody) *ExportNacosConfigResponse {
	s.Body = v
	return s
}

type GetEngineNamepaceRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEngineNamepaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEngineNamepaceRequest) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceRequest) SetClusterId(v string) *GetEngineNamepaceRequest {
	s.ClusterId = &v
	return s
}

func (s *GetEngineNamepaceRequest) SetId(v string) *GetEngineNamepaceRequest {
	s.Id = &v
	return s
}

func (s *GetEngineNamepaceRequest) SetInstanceId(v string) *GetEngineNamepaceRequest {
	s.InstanceId = &v
	return s
}

type GetEngineNamepaceResponseBody struct {
	ConfigCount       *string `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	ErrorCode         *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode          *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message           *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceDesc     *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	Quota             *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEngineNamepaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEngineNamepaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceResponseBody) SetConfigCount(v string) *GetEngineNamepaceResponseBody {
	s.ConfigCount = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetErrorCode(v string) *GetEngineNamepaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetHttpCode(v string) *GetEngineNamepaceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetMessage(v string) *GetEngineNamepaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespace(v string) *GetEngineNamepaceResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespaceDesc(v string) *GetEngineNamepaceResponseBody {
	s.NamespaceDesc = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetNamespaceShowName(v string) *GetEngineNamepaceResponseBody {
	s.NamespaceShowName = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetQuota(v string) *GetEngineNamepaceResponseBody {
	s.Quota = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetRequestId(v string) *GetEngineNamepaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetSuccess(v bool) *GetEngineNamepaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetEngineNamepaceResponseBody) SetType(v string) *GetEngineNamepaceResponseBody {
	s.Type = &v
	return s
}

type GetEngineNamepaceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEngineNamepaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEngineNamepaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEngineNamepaceResponse) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceResponse) SetHeaders(v map[string]*string) *GetEngineNamepaceResponse {
	s.Headers = v
	return s
}

func (s *GetEngineNamepaceResponse) SetBody(v *GetEngineNamepaceResponseBody) *GetEngineNamepaceResponse {
	s.Body = v
	return s
}

type GetGatewayRequest struct {
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s GetGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayRequest) SetGatewayUniqueId(v string) *GetGatewayRequest {
	s.GatewayUniqueId = &v
	return s
}

type GetGatewayResponseBody struct {
	Code           *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBody) SetCode(v int32) *GetGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResponseBody) SetData(v *GetGatewayResponseBodyData) *GetGatewayResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResponseBody) SetHttpStatusCode(v int32) *GetGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayResponseBody) SetMessage(v string) *GetGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResponseBody) SetRequestId(v string) *GetGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResponseBody) SetSuccess(v bool) *GetGatewayResponseBody {
	s.Success = &v
	return s
}

type GetGatewayResponseBodyData struct {
	ArmsOn           *bool                                       `json:"ArmsOn,omitempty" xml:"ArmsOn,omitempty"`
	ChargeType       *string                                     `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	EndDate          *string                                     `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GatewayType      *string                                     `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	GatewayUniqueId  *string                                     `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	GmtCreate        *string                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId       *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LogConfigDetails *GetGatewayResponseBodyDataLogConfigDetails `json:"LogConfigDetails,omitempty" xml:"LogConfigDetails,omitempty" type:"Struct"`
	Name             *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	PrimaryUser      *string                                     `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	Region           *string                                     `json:"Region,omitempty" xml:"Region,omitempty"`
	Replica          *int32                                      `json:"Replica,omitempty" xml:"Replica,omitempty"`
	SecurityGroup    *string                                     `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	Spec             *string                                     `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status           *int32                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Vpc              *string                                     `json:"Vpc,omitempty" xml:"Vpc,omitempty"`
	Vswitch          *string                                     `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	Vswitch2         *string                                     `json:"Vswitch2,omitempty" xml:"Vswitch2,omitempty"`
	XtraceDetails    *GetGatewayResponseBodyDataXtraceDetails    `json:"XtraceDetails,omitempty" xml:"XtraceDetails,omitempty" type:"Struct"`
}

func (s GetGatewayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyData) SetArmsOn(v bool) *GetGatewayResponseBodyData {
	s.ArmsOn = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetChargeType(v string) *GetGatewayResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetEndDate(v string) *GetGatewayResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayType(v string) *GetGatewayResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGmtCreate(v string) *GetGatewayResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetGmtModified(v string) *GetGatewayResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetId(v int64) *GetGatewayResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetInstanceId(v string) *GetGatewayResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetLogConfigDetails(v *GetGatewayResponseBodyDataLogConfigDetails) *GetGatewayResponseBodyData {
	s.LogConfigDetails = v
	return s
}

func (s *GetGatewayResponseBodyData) SetName(v string) *GetGatewayResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetPrimaryUser(v string) *GetGatewayResponseBodyData {
	s.PrimaryUser = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetRegion(v string) *GetGatewayResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetReplica(v int32) *GetGatewayResponseBodyData {
	s.Replica = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSecurityGroup(v string) *GetGatewayResponseBodyData {
	s.SecurityGroup = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetSpec(v string) *GetGatewayResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetStatus(v int32) *GetGatewayResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVpc(v string) *GetGatewayResponseBodyData {
	s.Vpc = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVswitch(v string) *GetGatewayResponseBodyData {
	s.Vswitch = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetVswitch2(v string) *GetGatewayResponseBodyData {
	s.Vswitch2 = &v
	return s
}

func (s *GetGatewayResponseBodyData) SetXtraceDetails(v *GetGatewayResponseBodyDataXtraceDetails) *GetGatewayResponseBodyData {
	s.XtraceDetails = v
	return s
}

type GetGatewayResponseBodyDataLogConfigDetails struct {
	LogEnabled   *bool   `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetGatewayResponseBodyDataLogConfigDetails) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataLogConfigDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetLogEnabled(v bool) *GetGatewayResponseBodyDataLogConfigDetails {
	s.LogEnabled = &v
	return s
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetLogStoreName(v string) *GetGatewayResponseBodyDataLogConfigDetails {
	s.LogStoreName = &v
	return s
}

func (s *GetGatewayResponseBodyDataLogConfigDetails) SetProjectName(v string) *GetGatewayResponseBodyDataLogConfigDetails {
	s.ProjectName = &v
	return s
}

type GetGatewayResponseBodyDataXtraceDetails struct {
	Sample  *int32 `json:"Sample,omitempty" xml:"Sample,omitempty"`
	TraceOn *bool  `json:"TraceOn,omitempty" xml:"TraceOn,omitempty"`
}

func (s GetGatewayResponseBodyDataXtraceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponseBodyDataXtraceDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayResponseBodyDataXtraceDetails) SetSample(v int32) *GetGatewayResponseBodyDataXtraceDetails {
	s.Sample = &v
	return s
}

func (s *GetGatewayResponseBodyDataXtraceDetails) SetTraceOn(v bool) *GetGatewayResponseBodyDataXtraceDetails {
	s.TraceOn = &v
	return s
}

type GetGatewayResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayResponse) SetHeaders(v map[string]*string) *GetGatewayResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayResponse) SetBody(v *GetGatewayResponseBody) *GetGatewayResponse {
	s.Body = v
	return s
}

type GetGatewayOptionRequest struct {
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
}

func (s GetGatewayOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayOptionRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionRequest) SetGatewayId(v int64) *GetGatewayOptionRequest {
	s.GatewayId = &v
	return s
}

type GetGatewayOptionResponseBody struct {
	Code           *int32         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GatewayOption `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionResponseBody) SetCode(v int32) *GetGatewayOptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetData(v *GatewayOption) *GetGatewayOptionResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayOptionResponseBody) SetHttpStatusCode(v int32) *GetGatewayOptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetMessage(v string) *GetGatewayOptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetRequestId(v string) *GetGatewayOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayOptionResponseBody) SetSuccess(v bool) *GetGatewayOptionResponseBody {
	s.Success = &v
	return s
}

type GetGatewayOptionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGatewayOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGatewayOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGatewayOptionResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayOptionResponse) SetHeaders(v map[string]*string) *GetGatewayOptionResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayOptionResponse) SetBody(v *GetGatewayOptionResponseBody) *GetGatewayOptionResponse {
	s.Body = v
	return s
}

type GetGovernanceKubernetesClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGovernanceKubernetesClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterRequest) SetClusterId(v string) *GetGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterRequest) SetRegionId(v string) *GetGovernanceKubernetesClusterRequest {
	s.RegionId = &v
	return s
}

type GetGovernanceKubernetesClusterResponseBody struct {
	Code           *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetGovernanceKubernetesClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetCode(v int32) *GetGovernanceKubernetesClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetData(v *GetGovernanceKubernetesClusterResponseBodyData) *GetGovernanceKubernetesClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetHttpStatusCode(v int32) *GetGovernanceKubernetesClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetMessage(v string) *GetGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *GetGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBody) SetSuccess(v bool) *GetGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

type GetGovernanceKubernetesClusterResponseBodyData struct {
	ClusterId      *string                                                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName    *string                                                     `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	K8sVersion     *string                                                     `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	NamespaceInfos *string                                                     `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	Namespaces     []*GetGovernanceKubernetesClusterResponseBodyDataNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	PilotStartTime *string                                                     `json:"PilotStartTime,omitempty" xml:"PilotStartTime,omitempty"`
	Region         *string                                                     `json:"Region,omitempty" xml:"Region,omitempty"`
	UpdateTime     *string                                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetClusterId(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetClusterName(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetK8sVersion(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.K8sVersion = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetNamespaceInfos(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.NamespaceInfos = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetNamespaces(v []*GetGovernanceKubernetesClusterResponseBodyDataNamespaces) *GetGovernanceKubernetesClusterResponseBodyData {
	s.Namespaces = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetPilotStartTime(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.PilotStartTime = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetRegion(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyData) SetUpdateTime(v string) *GetGovernanceKubernetesClusterResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetGovernanceKubernetesClusterResponseBodyDataNamespaces struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponseBodyDataNamespaces) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponseBodyDataNamespaces) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) SetName(v string) *GetGovernanceKubernetesClusterResponseBodyDataNamespaces {
	s.Name = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponseBodyDataNamespaces) SetTags(v string) *GetGovernanceKubernetesClusterResponseBodyDataNamespaces {
	s.Tags = &v
	return s
}

type GetGovernanceKubernetesClusterResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGovernanceKubernetesClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *GetGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponse) SetBody(v *GetGovernanceKubernetesClusterResponseBody) *GetGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

type GetGovernanceKubernetesClusterListRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetGovernanceKubernetesClusterListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterListRequest) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterListRequest) SetClusterId(v string) *GetGovernanceKubernetesClusterListRequest {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListRequest) SetClusterName(v string) *GetGovernanceKubernetesClusterListRequest {
	s.ClusterName = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListRequest) SetPageNumber(v int32) *GetGovernanceKubernetesClusterListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListRequest) SetPageSize(v int32) *GetGovernanceKubernetesClusterListRequest {
	s.PageSize = &v
	return s
}

type GetGovernanceKubernetesClusterListResponseBody struct {
	Code           *int32                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetGovernanceKubernetesClusterListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGovernanceKubernetesClusterListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetCode(v int32) *GetGovernanceKubernetesClusterListResponseBody {
	s.Code = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetData(v *GetGovernanceKubernetesClusterListResponseBodyData) *GetGovernanceKubernetesClusterListResponseBody {
	s.Data = v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetHttpStatusCode(v int32) *GetGovernanceKubernetesClusterListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetMessage(v string) *GetGovernanceKubernetesClusterListResponseBody {
	s.Message = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetRequestId(v string) *GetGovernanceKubernetesClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBody) SetSuccess(v bool) *GetGovernanceKubernetesClusterListResponseBody {
	s.Success = &v
	return s
}

type GetGovernanceKubernetesClusterListResponseBodyData struct {
	PageNumber *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result     []*GetGovernanceKubernetesClusterListResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalSize  *int32                                                      `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetGovernanceKubernetesClusterListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterListResponseBodyData) SetPageNumber(v int32) *GetGovernanceKubernetesClusterListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyData) SetPageSize(v int32) *GetGovernanceKubernetesClusterListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyData) SetResult(v []*GetGovernanceKubernetesClusterListResponseBodyDataResult) *GetGovernanceKubernetesClusterListResponseBodyData {
	s.Result = v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyData) SetTotalSize(v int32) *GetGovernanceKubernetesClusterListResponseBodyData {
	s.TotalSize = &v
	return s
}

type GetGovernanceKubernetesClusterListResponseBodyDataResult struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName    *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	K8sVersion     *string `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	NamespaceInfos *string `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	PilotStartTime *string `json:"PilotStartTime,omitempty" xml:"PilotStartTime,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetGovernanceKubernetesClusterListResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterListResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetClusterId(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetClusterName(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.ClusterName = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetK8sVersion(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.K8sVersion = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetNamespaceInfos(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.NamespaceInfos = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetPilotStartTime(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.PilotStartTime = &v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponseBodyDataResult) SetRegion(v string) *GetGovernanceKubernetesClusterListResponseBodyDataResult {
	s.Region = &v
	return s
}

type GetGovernanceKubernetesClusterListResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGovernanceKubernetesClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGovernanceKubernetesClusterListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceKubernetesClusterListResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterListResponse) SetHeaders(v map[string]*string) *GetGovernanceKubernetesClusterListResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceKubernetesClusterListResponse) SetBody(v *GetGovernanceKubernetesClusterListResponseBody) *GetGovernanceKubernetesClusterListResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	// 集群版本
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetVersionCode(v string) *GetImageRequest {
	s.VersionCode = &v
	return s
}

type GetImageResponseBody struct {
	Data      *GetImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetData(v *GetImageResponseBodyData) *GetImageResponseBody {
	s.Data = v
	return s
}

func (s *GetImageResponseBody) SetErrorCode(v string) *GetImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetImageResponseBody) SetHttpCode(v string) *GetImageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetImageResponseBody) SetMessage(v string) *GetImageResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

type GetImageResponseBodyData struct {
	// 当前集群镜像版本的4位全名
	CurrentVersionFullShowName *string `json:"CurrentVersionFullShowName,omitempty" xml:"CurrentVersionFullShowName,omitempty"`
	// 可升级的最大版本变更日志url
	MaxVersionChangelogUrl *string `json:"MaxVersionChangelogUrl,omitempty" xml:"MaxVersionChangelogUrl,omitempty"`
	// 可升级的增量版本Code
	MaxVersionCode *string `json:"MaxVersionCode,omitempty" xml:"MaxVersionCode,omitempty"`
	// 可升级的增量版本全名
	MaxVersionFullShowName *string `json:"MaxVersionFullShowName,omitempty" xml:"MaxVersionFullShowName,omitempty"`
}

func (s GetImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyData) SetCurrentVersionFullShowName(v string) *GetImageResponseBodyData {
	s.CurrentVersionFullShowName = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionChangelogUrl(v string) *GetImageResponseBodyData {
	s.MaxVersionChangelogUrl = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionCode(v string) *GetImageResponseBodyData {
	s.MaxVersionCode = &v
	return s
}

func (s *GetImageResponseBodyData) SetMaxVersionFullShowName(v string) *GetImageResponseBodyData {
	s.MaxVersionFullShowName = &v
	return s
}

type GetImageResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetImportFileUrlRequest struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetImportFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlRequest) SetContentType(v string) *GetImportFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetImportFileUrlRequest) SetInstanceId(v string) *GetImportFileUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetImportFileUrlRequest) SetNamespaceId(v string) *GetImportFileUrlRequest {
	s.NamespaceId = &v
	return s
}

type GetImportFileUrlResponseBody struct {
	Code           *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetImportFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImportFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImportFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponseBody) SetCode(v int32) *GetImportFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetData(v *GetImportFileUrlResponseBodyData) *GetImportFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetImportFileUrlResponseBody) SetDynamicMessage(v string) *GetImportFileUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetErrorCode(v string) *GetImportFileUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetHttpStatusCode(v int32) *GetImportFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetMessage(v string) *GetImportFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetRequestId(v string) *GetImportFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImportFileUrlResponseBody) SetSuccess(v bool) *GetImportFileUrlResponseBody {
	s.Success = &v
	return s
}

type GetImportFileUrlResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImportFileUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetImportFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponseBodyData) SetUrl(v string) *GetImportFileUrlResponseBodyData {
	s.Url = &v
	return s
}

type GetImportFileUrlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImportFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImportFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImportFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetImportFileUrlResponse) SetHeaders(v map[string]*string) *GetImportFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetImportFileUrlResponse) SetBody(v *GetImportFileUrlResponseBody) *GetImportFileUrlResponse {
	s.Body = v
	return s
}

type GetMseFeatureSwitchResponseBody struct {
	ErrorCode *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMseFeatureSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMseFeatureSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetMseFeatureSwitchResponseBody) SetErrorCode(v string) *GetMseFeatureSwitchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetMessage(v string) *GetMseFeatureSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetRequestId(v string) *GetMseFeatureSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetResult(v map[string]interface{}) *GetMseFeatureSwitchResponseBody {
	s.Result = v
	return s
}

func (s *GetMseFeatureSwitchResponseBody) SetSuccess(v bool) *GetMseFeatureSwitchResponseBody {
	s.Success = &v
	return s
}

type GetMseFeatureSwitchResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMseFeatureSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMseFeatureSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMseFeatureSwitchResponse) GoString() string {
	return s.String()
}

func (s *GetMseFeatureSwitchResponse) SetHeaders(v map[string]*string) *GetMseFeatureSwitchResponse {
	s.Headers = v
	return s
}

func (s *GetMseFeatureSwitchResponse) SetBody(v *GetMseFeatureSwitchResponseBody) *GetMseFeatureSwitchResponse {
	s.Body = v
	return s
}

type GetNacosConfigRequest struct {
	Beta        *bool   `json:"Beta,omitempty" xml:"Beta,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNacosConfigRequest) SetBeta(v bool) *GetNacosConfigRequest {
	s.Beta = &v
	return s
}

func (s *GetNacosConfigRequest) SetDataId(v string) *GetNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *GetNacosConfigRequest) SetGroup(v string) *GetNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *GetNacosConfigRequest) SetInstanceId(v string) *GetNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNacosConfigRequest) SetNamespaceId(v string) *GetNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

type GetNacosConfigResponseBody struct {
	Configuration *GetNacosConfigResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	ErrorCode     *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message       *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponseBody) SetConfiguration(v *GetNacosConfigResponseBodyConfiguration) *GetNacosConfigResponseBody {
	s.Configuration = v
	return s
}

func (s *GetNacosConfigResponseBody) SetErrorCode(v string) *GetNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetMessage(v string) *GetNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetRequestId(v string) *GetNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacosConfigResponseBody) SetSuccess(v bool) *GetNacosConfigResponseBody {
	s.Success = &v
	return s
}

type GetNacosConfigResponseBodyConfiguration struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BetaIps          *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc             *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	Group            *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetNacosConfigResponseBodyConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetNacosConfigResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponseBodyConfiguration) SetAppName(v string) *GetNacosConfigResponseBodyConfiguration {
	s.AppName = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetBetaIps(v string) *GetNacosConfigResponseBodyConfiguration {
	s.BetaIps = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetContent(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Content = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetDataId(v string) *GetNacosConfigResponseBodyConfiguration {
	s.DataId = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetDesc(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Desc = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetEncryptedDataKey(v string) *GetNacosConfigResponseBodyConfiguration {
	s.EncryptedDataKey = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetGroup(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Group = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetMd5(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Md5 = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetTags(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Tags = &v
	return s
}

func (s *GetNacosConfigResponseBodyConfiguration) SetType(v string) *GetNacosConfigResponseBodyConfiguration {
	s.Type = &v
	return s
}

type GetNacosConfigResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNacosConfigResponse) SetHeaders(v map[string]*string) *GetNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNacosConfigResponse) SetBody(v *GetNacosConfigResponseBody) *GetNacosConfigResponse {
	s.Body = v
	return s
}

type GetNacosHistoryConfigRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Nid         *string `json:"Nid,omitempty" xml:"Nid,omitempty"`
}

func (s GetNacosHistoryConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNacosHistoryConfigRequest) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigRequest) SetDataId(v string) *GetNacosHistoryConfigRequest {
	s.DataId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetGroup(v string) *GetNacosHistoryConfigRequest {
	s.Group = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetInstanceId(v string) *GetNacosHistoryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetNamespaceId(v string) *GetNacosHistoryConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNacosHistoryConfigRequest) SetNid(v string) *GetNacosHistoryConfigRequest {
	s.Nid = &v
	return s
}

type GetNacosHistoryConfigResponseBody struct {
	Configuration *GetNacosHistoryConfigResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	ErrorCode     *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message       *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNacosHistoryConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNacosHistoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponseBody) SetConfiguration(v *GetNacosHistoryConfigResponseBodyConfiguration) *GetNacosHistoryConfigResponseBody {
	s.Configuration = v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetErrorCode(v string) *GetNacosHistoryConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetMessage(v string) *GetNacosHistoryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetRequestId(v string) *GetNacosHistoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetSuccess(v bool) *GetNacosHistoryConfigResponseBody {
	s.Success = &v
	return s
}

type GetNacosHistoryConfigResponseBodyConfiguration struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	Group            *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	OpType           *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s GetNacosHistoryConfigResponseBodyConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetNacosHistoryConfigResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetAppName(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.AppName = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetContent(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Content = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetDataId(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.DataId = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetEncryptedDataKey(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.EncryptedDataKey = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetGroup(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Group = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetMd5(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Md5 = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetOpType(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.OpType = &v
	return s
}

type GetNacosHistoryConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNacosHistoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNacosHistoryConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNacosHistoryConfigResponse) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponse) SetHeaders(v map[string]*string) *GetNacosHistoryConfigResponse {
	s.Headers = v
	return s
}

func (s *GetNacosHistoryConfigResponse) SetBody(v *GetNacosHistoryConfigResponseBody) *GetNacosHistoryConfigResponse {
	s.Body = v
	return s
}

type GetOverviewRequest struct {
	Period *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetOverviewRequest) SetPeriod(v int32) *GetOverviewRequest {
	s.Period = &v
	return s
}

func (s *GetOverviewRequest) SetRegion(v string) *GetOverviewRequest {
	s.Region = &v
	return s
}

type GetOverviewResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetOverviewResponseBody) SetCode(v int32) *GetOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *GetOverviewResponseBody) SetData(v string) *GetOverviewResponseBody {
	s.Data = &v
	return s
}

func (s *GetOverviewResponseBody) SetHttpStatusCode(v int32) *GetOverviewResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOverviewResponseBody) SetMessage(v string) *GetOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetOverviewResponseBody) SetRequestId(v string) *GetOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOverviewResponseBody) SetSuccess(v string) *GetOverviewResponseBody {
	s.Success = &v
	return s
}

type GetOverviewResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOverviewResponse) GoString() string {
	return s.String()
}

func (s *GetOverviewResponse) SetHeaders(v map[string]*string) *GetOverviewResponse {
	s.Headers = v
	return s
}

func (s *GetOverviewResponse) SetBody(v *GetOverviewResponseBody) *GetOverviewResponse {
	s.Body = v
	return s
}

type ImportNacosConfigRequest struct {
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Policy      *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ImportNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigRequest) SetFileUrl(v string) *ImportNacosConfigRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportNacosConfigRequest) SetInstanceId(v string) *ImportNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportNacosConfigRequest) SetNamespaceId(v string) *ImportNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *ImportNacosConfigRequest) SetPolicy(v string) *ImportNacosConfigRequest {
	s.Policy = &v
	return s
}

type ImportNacosConfigResponseBody struct {
	Code           *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ImportNacosConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBody) SetCode(v int32) *ImportNacosConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetData(v *ImportNacosConfigResponseBodyData) *ImportNacosConfigResponseBody {
	s.Data = v
	return s
}

func (s *ImportNacosConfigResponseBody) SetDynamicMessage(v string) *ImportNacosConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetErrorCode(v string) *ImportNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetHttpStatusCode(v int32) *ImportNacosConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetMessage(v string) *ImportNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetRequestId(v string) *ImportNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNacosConfigResponseBody) SetSuccess(v bool) *ImportNacosConfigResponseBody {
	s.Success = &v
	return s
}

type ImportNacosConfigResponseBodyData struct {
	FailData  []*ImportNacosConfigResponseBodyDataFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Repeated"`
	SkipCount *int32                                       `json:"SkipCount,omitempty" xml:"SkipCount,omitempty"`
	SkipData  []*ImportNacosConfigResponseBodyDataSkipData `json:"SkipData,omitempty" xml:"SkipData,omitempty" type:"Repeated"`
	SuccCount *int32                                       `json:"SuccCount,omitempty" xml:"SuccCount,omitempty"`
}

func (s ImportNacosConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyData) SetFailData(v []*ImportNacosConfigResponseBodyDataFailData) *ImportNacosConfigResponseBodyData {
	s.FailData = v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSkipCount(v int32) *ImportNacosConfigResponseBodyData {
	s.SkipCount = &v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSkipData(v []*ImportNacosConfigResponseBodyDataSkipData) *ImportNacosConfigResponseBodyData {
	s.SkipData = v
	return s
}

func (s *ImportNacosConfigResponseBodyData) SetSuccCount(v int32) *ImportNacosConfigResponseBodyData {
	s.SuccCount = &v
	return s
}

type ImportNacosConfigResponseBodyDataFailData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s ImportNacosConfigResponseBodyDataFailData) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigResponseBodyDataFailData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyDataFailData) SetDataId(v string) *ImportNacosConfigResponseBodyDataFailData {
	s.DataId = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataFailData) SetGroup(v string) *ImportNacosConfigResponseBodyDataFailData {
	s.Group = &v
	return s
}

type ImportNacosConfigResponseBodyDataSkipData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s ImportNacosConfigResponseBodyDataSkipData) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigResponseBodyDataSkipData) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponseBodyDataSkipData) SetDataId(v string) *ImportNacosConfigResponseBodyDataSkipData {
	s.DataId = &v
	return s
}

func (s *ImportNacosConfigResponseBodyDataSkipData) SetGroup(v string) *ImportNacosConfigResponseBodyDataSkipData {
	s.Group = &v
	return s
}

type ImportNacosConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigResponse) SetHeaders(v map[string]*string) *ImportNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *ImportNacosConfigResponse) SetBody(v *ImportNacosConfigResponseBody) *ImportNacosConfigResponse {
	s.Body = v
	return s
}

type ListAlarmContactGroupsRequest struct {
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListAlarmContactGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmContactGroupsRequest) SetPageNum(v int32) *ListAlarmContactGroupsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlarmContactGroupsRequest) SetPageSize(v int32) *ListAlarmContactGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmContactGroupsRequest) SetRequestPars(v string) *ListAlarmContactGroupsRequest {
	s.RequestPars = &v
	return s
}

type ListAlarmContactGroupsResponseBody struct {
	Data       []*ListAlarmContactGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmContactGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmContactGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmContactGroupsResponseBody) SetData(v []*ListAlarmContactGroupsResponseBodyData) *ListAlarmContactGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetErrorCode(v string) *ListAlarmContactGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetHttpCode(v string) *ListAlarmContactGroupsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetMessage(v string) *ListAlarmContactGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetPageNumber(v int32) *ListAlarmContactGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetPageSize(v int32) *ListAlarmContactGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetRequestId(v string) *ListAlarmContactGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetSuccess(v bool) *ListAlarmContactGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBody) SetTotalCount(v int32) *ListAlarmContactGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlarmContactGroupsResponseBodyData struct {
	ContactGroupId   *string `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
}

func (s ListAlarmContactGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmContactGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmContactGroupsResponseBodyData) SetContactGroupId(v string) *ListAlarmContactGroupsResponseBodyData {
	s.ContactGroupId = &v
	return s
}

func (s *ListAlarmContactGroupsResponseBodyData) SetContactGroupName(v string) *ListAlarmContactGroupsResponseBodyData {
	s.ContactGroupName = &v
	return s
}

type ListAlarmContactGroupsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmContactGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmContactGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmContactGroupsResponse) SetHeaders(v map[string]*string) *ListAlarmContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmContactGroupsResponse) SetBody(v *ListAlarmContactGroupsResponseBody) *ListAlarmContactGroupsResponse {
	s.Body = v
	return s
}

type ListAlarmHistoriesRequest struct {
	AlarmMseType *string `json:"AlarmMseType,omitempty" xml:"AlarmMseType,omitempty"`
	AlarmName    *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars  *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAlarmHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesRequest) SetAlarmMseType(v string) *ListAlarmHistoriesRequest {
	s.AlarmMseType = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetAlarmName(v string) *ListAlarmHistoriesRequest {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetEndTime(v int64) *ListAlarmHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetPageNum(v int32) *ListAlarmHistoriesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetPageSize(v int32) *ListAlarmHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetRequestPars(v string) *ListAlarmHistoriesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAlarmHistoriesRequest) SetStartTime(v int64) *ListAlarmHistoriesRequest {
	s.StartTime = &v
	return s
}

type ListAlarmHistoriesResponseBody struct {
	Data       []*ListAlarmHistoriesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                               `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBody) SetData(v []*ListAlarmHistoriesResponseBodyData) *ListAlarmHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetErrorCode(v string) *ListAlarmHistoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetHttpCode(v string) *ListAlarmHistoriesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetMessage(v string) *ListAlarmHistoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetPageNumber(v int32) *ListAlarmHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetPageSize(v int32) *ListAlarmHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetRequestId(v string) *ListAlarmHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetSuccess(v bool) *ListAlarmHistoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmHistoriesResponseBody) SetTotalCount(v int32) *ListAlarmHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlarmHistoriesResponseBodyData struct {
	AlarmContent  *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmDingDing *string `json:"AlarmDingDing,omitempty" xml:"AlarmDingDing,omitempty"`
	AlarmEmail    *string `json:"AlarmEmail,omitempty" xml:"AlarmEmail,omitempty"`
	AlarmName     *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	AlarmPhone    *string `json:"AlarmPhone,omitempty" xml:"AlarmPhone,omitempty"`
	AlarmTime     *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
}

func (s ListAlarmHistoriesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmContent(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmContent = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmDingDing(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmDingDing = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmEmail(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmEmail = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmName(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmPhone(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmPhone = &v
	return s
}

func (s *ListAlarmHistoriesResponseBodyData) SetAlarmTime(v string) *ListAlarmHistoriesResponseBodyData {
	s.AlarmTime = &v
	return s
}

type ListAlarmHistoriesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmHistoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmHistoriesResponse) SetHeaders(v map[string]*string) *ListAlarmHistoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmHistoriesResponse) SetBody(v *ListAlarmHistoriesResponseBody) *ListAlarmHistoriesResponse {
	s.Body = v
	return s
}

type ListAlarmItemsRequest struct {
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListAlarmItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmItemsRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmItemsRequest) SetRequestPars(v string) *ListAlarmItemsRequest {
	s.RequestPars = &v
	return s
}

type ListAlarmItemsResponseBody struct {
	Data       []*ListAlarmItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                           `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmItemsResponseBody) SetData(v []*ListAlarmItemsResponseBodyData) *ListAlarmItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmItemsResponseBody) SetErrorCode(v string) *ListAlarmItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetHttpCode(v string) *ListAlarmItemsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetMessage(v string) *ListAlarmItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetPageNumber(v int32) *ListAlarmItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetPageSize(v int32) *ListAlarmItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetRequestId(v string) *ListAlarmItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetSuccess(v bool) *ListAlarmItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmItemsResponseBody) SetTotalCount(v int32) *ListAlarmItemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlarmItemsResponseBodyData struct {
	AlarmCode   *string `json:"AlarmCode,omitempty" xml:"AlarmCode,omitempty"`
	AlarmDesc   *string `json:"AlarmDesc,omitempty" xml:"AlarmDesc,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s ListAlarmItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmItemsResponseBodyData) SetAlarmCode(v string) *ListAlarmItemsResponseBodyData {
	s.AlarmCode = &v
	return s
}

func (s *ListAlarmItemsResponseBodyData) SetAlarmDesc(v string) *ListAlarmItemsResponseBodyData {
	s.AlarmDesc = &v
	return s
}

func (s *ListAlarmItemsResponseBodyData) SetClusterType(v string) *ListAlarmItemsResponseBodyData {
	s.ClusterType = &v
	return s
}

type ListAlarmItemsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmItemsResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmItemsResponse) SetHeaders(v map[string]*string) *ListAlarmItemsResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmItemsResponse) SetBody(v *ListAlarmItemsResponseBody) *ListAlarmItemsResponse {
	s.Body = v
	return s
}

type ListAlarmRulesRequest struct {
	AlarmMseType *string `json:"AlarmMseType,omitempty" xml:"AlarmMseType,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars  *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListAlarmRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmRulesRequest) SetAlarmMseType(v string) *ListAlarmRulesRequest {
	s.AlarmMseType = &v
	return s
}

func (s *ListAlarmRulesRequest) SetPageNum(v int32) *ListAlarmRulesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAlarmRulesRequest) SetPageSize(v int32) *ListAlarmRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlarmRulesRequest) SetRequestPars(v string) *ListAlarmRulesRequest {
	s.RequestPars = &v
	return s
}

type ListAlarmRulesResponseBody struct {
	Data       []*ListAlarmRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                           `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlarmRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmRulesResponseBody) SetData(v []*ListAlarmRulesResponseBodyData) *ListAlarmRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmRulesResponseBody) SetErrorCode(v string) *ListAlarmRulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetHttpCode(v string) *ListAlarmRulesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetMessage(v string) *ListAlarmRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetPageNumber(v int32) *ListAlarmRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetPageSize(v int32) *ListAlarmRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetRequestId(v string) *ListAlarmRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetSuccess(v bool) *ListAlarmRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmRulesResponseBody) SetTotalCount(v int32) *ListAlarmRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAlarmRulesResponseBodyData struct {
	AlarmName       *string `json:"AlarmName,omitempty" xml:"AlarmName,omitempty"`
	AlarmRuleDetail *string `json:"AlarmRuleDetail,omitempty" xml:"AlarmRuleDetail,omitempty"`
	AlarmRuleId     *string `json:"AlarmRuleId,omitempty" xml:"AlarmRuleId,omitempty"`
	AlarmStatus     *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListAlarmRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmRulesResponseBodyData) SetAlarmName(v string) *ListAlarmRulesResponseBodyData {
	s.AlarmName = &v
	return s
}

func (s *ListAlarmRulesResponseBodyData) SetAlarmRuleDetail(v string) *ListAlarmRulesResponseBodyData {
	s.AlarmRuleDetail = &v
	return s
}

func (s *ListAlarmRulesResponseBodyData) SetAlarmRuleId(v string) *ListAlarmRulesResponseBodyData {
	s.AlarmRuleId = &v
	return s
}

func (s *ListAlarmRulesResponseBodyData) SetAlarmStatus(v string) *ListAlarmRulesResponseBodyData {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmRulesResponseBodyData) SetCreateTime(v string) *ListAlarmRulesResponseBodyData {
	s.CreateTime = &v
	return s
}

type ListAlarmRulesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlarmRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlarmRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlarmRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAlarmRulesResponse) SetHeaders(v map[string]*string) *ListAlarmRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAlarmRulesResponse) SetBody(v *ListAlarmRulesResponseBody) *ListAlarmRulesResponse {
	s.Body = v
	return s
}

type ListAnsInstancesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesRequest) SetClusterId(v string) *ListAnsInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsInstancesRequest) SetClusterName(v string) *ListAnsInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListAnsInstancesRequest) SetGroupName(v string) *ListAnsInstancesRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsInstancesRequest) SetNamespaceId(v string) *ListAnsInstancesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsInstancesRequest) SetPageNum(v int32) *ListAnsInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsInstancesRequest) SetPageSize(v int32) *ListAnsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsInstancesRequest) SetRequestPars(v string) *ListAnsInstancesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAnsInstancesRequest) SetServiceName(v string) *ListAnsInstancesRequest {
	s.ServiceName = &v
	return s
}

type ListAnsInstancesResponseBody struct {
	Data       []*ListAnsInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                             `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponseBody) SetData(v []*ListAnsInstancesResponseBodyData) *ListAnsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsInstancesResponseBody) SetErrorCode(v string) *ListAnsInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetHttpCode(v string) *ListAnsInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetMessage(v string) *ListAnsInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetPageNumber(v int32) *ListAnsInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetPageSize(v int32) *ListAnsInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetRequestId(v string) *ListAnsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetSuccess(v bool) *ListAnsInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnsInstancesResponseBody) SetTotalCount(v int32) *ListAnsInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAnsInstancesResponseBodyData struct {
	App                       *string                `json:"App,omitempty" xml:"App,omitempty"`
	ClusterName               *string                `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DatumKey                  *string                `json:"DatumKey,omitempty" xml:"DatumKey,omitempty"`
	DefaultKey                *string                `json:"DefaultKey,omitempty" xml:"DefaultKey,omitempty"`
	Enabled                   *bool                  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Ephemeral                 *bool                  `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	FailCount                 *int32                 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	Healthy                   *bool                  `json:"Healthy,omitempty" xml:"Healthy,omitempty"`
	InstanceHeartBeatInterval *int32                 `json:"InstanceHeartBeatInterval,omitempty" xml:"InstanceHeartBeatInterval,omitempty"`
	InstanceHeartBeatTimeOut  *int32                 `json:"InstanceHeartBeatTimeOut,omitempty" xml:"InstanceHeartBeatTimeOut,omitempty"`
	InstanceId                *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip                        *string                `json:"Ip,omitempty" xml:"Ip,omitempty"`
	IpDeleteTimeout           *int32                 `json:"IpDeleteTimeout,omitempty" xml:"IpDeleteTimeout,omitempty"`
	LastBeat                  *int64                 `json:"LastBeat,omitempty" xml:"LastBeat,omitempty"`
	Marked                    *bool                  `json:"Marked,omitempty" xml:"Marked,omitempty"`
	Metadata                  map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	OkCount                   *int32                 `json:"OkCount,omitempty" xml:"OkCount,omitempty"`
	Port                      *int32                 `json:"Port,omitempty" xml:"Port,omitempty"`
	ServiceName               *string                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Weight                    *int32                 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListAnsInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAnsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponseBodyData) SetApp(v string) *ListAnsInstancesResponseBodyData {
	s.App = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetClusterName(v string) *ListAnsInstancesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetDatumKey(v string) *ListAnsInstancesResponseBodyData {
	s.DatumKey = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetDefaultKey(v string) *ListAnsInstancesResponseBodyData {
	s.DefaultKey = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetEnabled(v bool) *ListAnsInstancesResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetEphemeral(v bool) *ListAnsInstancesResponseBodyData {
	s.Ephemeral = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetFailCount(v int32) *ListAnsInstancesResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetHealthy(v bool) *ListAnsInstancesResponseBodyData {
	s.Healthy = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceHeartBeatInterval(v int32) *ListAnsInstancesResponseBodyData {
	s.InstanceHeartBeatInterval = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceHeartBeatTimeOut(v int32) *ListAnsInstancesResponseBodyData {
	s.InstanceHeartBeatTimeOut = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetInstanceId(v string) *ListAnsInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetIp(v string) *ListAnsInstancesResponseBodyData {
	s.Ip = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetIpDeleteTimeout(v int32) *ListAnsInstancesResponseBodyData {
	s.IpDeleteTimeout = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetLastBeat(v int64) *ListAnsInstancesResponseBodyData {
	s.LastBeat = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetMarked(v bool) *ListAnsInstancesResponseBodyData {
	s.Marked = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetMetadata(v map[string]interface{}) *ListAnsInstancesResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetOkCount(v int32) *ListAnsInstancesResponseBodyData {
	s.OkCount = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetPort(v int32) *ListAnsInstancesResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetServiceName(v string) *ListAnsInstancesResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListAnsInstancesResponseBodyData) SetWeight(v int32) *ListAnsInstancesResponseBodyData {
	s.Weight = &v
	return s
}

type ListAnsInstancesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAnsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAnsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesResponse) SetHeaders(v map[string]*string) *ListAnsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAnsInstancesResponse) SetBody(v *ListAnsInstancesResponseBody) *ListAnsInstancesResponse {
	s.Body = v
	return s
}

type ListAnsServiceClustersRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsServiceClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServiceClustersRequest) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersRequest) SetClusterId(v string) *ListAnsServiceClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetClusterName(v string) *ListAnsServiceClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetGroupName(v string) *ListAnsServiceClustersRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetNamespaceId(v string) *ListAnsServiceClustersRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetPageNum(v int32) *ListAnsServiceClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetPageSize(v int32) *ListAnsServiceClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetRequestPars(v string) *ListAnsServiceClustersRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetServiceName(v string) *ListAnsServiceClustersRequest {
	s.ServiceName = &v
	return s
}

type ListAnsServiceClustersResponseBody struct {
	Data      *ListAnsServiceClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAnsServiceClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServiceClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBody) SetData(v *ListAnsServiceClustersResponseBodyData) *ListAnsServiceClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetErrorCode(v string) *ListAnsServiceClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetMessage(v string) *ListAnsServiceClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetRequestId(v string) *ListAnsServiceClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsServiceClustersResponseBody) SetSuccess(v bool) *ListAnsServiceClustersResponseBody {
	s.Success = &v
	return s
}

type ListAnsServiceClustersResponseBodyData struct {
	Clusters         []*ListAnsServiceClustersResponseBodyDataClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	GroupName        *string                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Metadata         map[string]interface{}                            `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Name             *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	ProtectThreshold *float32                                          `json:"ProtectThreshold,omitempty" xml:"ProtectThreshold,omitempty"`
	SelectorType     *string                                           `json:"SelectorType,omitempty" xml:"SelectorType,omitempty"`
}

func (s ListAnsServiceClustersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServiceClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBodyData) SetClusters(v []*ListAnsServiceClustersResponseBodyDataClusters) *ListAnsServiceClustersResponseBodyData {
	s.Clusters = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetGroupName(v string) *ListAnsServiceClustersResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetMetadata(v map[string]interface{}) *ListAnsServiceClustersResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetName(v string) *ListAnsServiceClustersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetProtectThreshold(v float32) *ListAnsServiceClustersResponseBodyData {
	s.ProtectThreshold = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyData) SetSelectorType(v string) *ListAnsServiceClustersResponseBodyData {
	s.SelectorType = &v
	return s
}

type ListAnsServiceClustersResponseBodyDataClusters struct {
	DefaultCheckPort  *int32                 `json:"DefaultCheckPort,omitempty" xml:"DefaultCheckPort,omitempty"`
	DefaultPort       *int32                 `json:"DefaultPort,omitempty" xml:"DefaultPort,omitempty"`
	HealthCheckerType *string                `json:"HealthCheckerType,omitempty" xml:"HealthCheckerType,omitempty"`
	Metadata          map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Name              *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceName       *string                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	UseIPPort4Check   *bool                  `json:"UseIPPort4Check,omitempty" xml:"UseIPPort4Check,omitempty"`
}

func (s ListAnsServiceClustersResponseBodyDataClusters) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServiceClustersResponseBodyDataClusters) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetDefaultCheckPort(v int32) *ListAnsServiceClustersResponseBodyDataClusters {
	s.DefaultCheckPort = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetDefaultPort(v int32) *ListAnsServiceClustersResponseBodyDataClusters {
	s.DefaultPort = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetHealthCheckerType(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.HealthCheckerType = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetMetadata(v map[string]interface{}) *ListAnsServiceClustersResponseBodyDataClusters {
	s.Metadata = v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetName(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.Name = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetServiceName(v string) *ListAnsServiceClustersResponseBodyDataClusters {
	s.ServiceName = &v
	return s
}

func (s *ListAnsServiceClustersResponseBodyDataClusters) SetUseIPPort4Check(v bool) *ListAnsServiceClustersResponseBodyDataClusters {
	s.UseIPPort4Check = &v
	return s
}

type ListAnsServiceClustersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAnsServiceClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAnsServiceClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServiceClustersResponse) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersResponse) SetHeaders(v map[string]*string) *ListAnsServiceClustersResponse {
	s.Headers = v
	return s
}

func (s *ListAnsServiceClustersResponse) SetBody(v *ListAnsServiceClustersResponseBody) *ListAnsServiceClustersResponse {
	s.Body = v
	return s
}

type ListAnsServicesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HasIpCount  *string `json:"HasIpCount,omitempty" xml:"HasIpCount,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServicesRequest) GoString() string {
	return s.String()
}

func (s *ListAnsServicesRequest) SetClusterId(v string) *ListAnsServicesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsServicesRequest) SetGroupName(v string) *ListAnsServicesRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsServicesRequest) SetHasIpCount(v string) *ListAnsServicesRequest {
	s.HasIpCount = &v
	return s
}

func (s *ListAnsServicesRequest) SetInstanceId(v string) *ListAnsServicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAnsServicesRequest) SetNamespaceId(v string) *ListAnsServicesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsServicesRequest) SetPageNum(v int32) *ListAnsServicesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsServicesRequest) SetPageSize(v int32) *ListAnsServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsServicesRequest) SetRequestPars(v string) *ListAnsServicesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAnsServicesRequest) SetServiceName(v string) *ListAnsServicesRequest {
	s.ServiceName = &v
	return s
}

type ListAnsServicesResponseBody struct {
	Data       []*ListAnsServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnsServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponseBody) SetData(v []*ListAnsServicesResponseBodyData) *ListAnsServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListAnsServicesResponseBody) SetErrorCode(v string) *ListAnsServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetHttpCode(v string) *ListAnsServicesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetMessage(v string) *ListAnsServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetPageNumber(v int32) *ListAnsServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetPageSize(v int32) *ListAnsServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetRequestId(v string) *ListAnsServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetSuccess(v bool) *ListAnsServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnsServicesResponseBody) SetTotalCount(v int32) *ListAnsServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAnsServicesResponseBodyData struct {
	ClusterCount         *int32  `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HealthyInstanceCount *int32  `json:"HealthyInstanceCount,omitempty" xml:"HealthyInstanceCount,omitempty"`
	IpCount              *int32  `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAnsServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponseBodyData) SetClusterCount(v int32) *ListAnsServicesResponseBodyData {
	s.ClusterCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetGroupName(v string) *ListAnsServicesResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetHealthyInstanceCount(v int32) *ListAnsServicesResponseBodyData {
	s.HealthyInstanceCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetIpCount(v int32) *ListAnsServicesResponseBodyData {
	s.IpCount = &v
	return s
}

func (s *ListAnsServicesResponseBodyData) SetName(v string) *ListAnsServicesResponseBodyData {
	s.Name = &v
	return s
}

type ListAnsServicesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAnsServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAnsServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAnsServicesResponse) GoString() string {
	return s.String()
}

func (s *ListAnsServicesResponse) SetHeaders(v map[string]*string) *ListAnsServicesResponse {
	s.Headers = v
	return s
}

func (s *ListAnsServicesResponse) SetBody(v *ListAnsServicesResponseBody) *ListAnsServicesResponse {
	s.Body = v
	return s
}

type ListClusterConnectionTypesResponseBody struct {
	Code           *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListClusterConnectionTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DynamicMessage *string                                       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterConnectionTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterConnectionTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponseBody) SetCode(v int32) *ListClusterConnectionTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetData(v []*ListClusterConnectionTypesResponseBodyData) *ListClusterConnectionTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetDynamicMessage(v string) *ListClusterConnectionTypesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetErrorCode(v string) *ListClusterConnectionTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetHttpStatusCode(v int32) *ListClusterConnectionTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetMessage(v string) *ListClusterConnectionTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetRequestId(v string) *ListClusterConnectionTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterConnectionTypesResponseBody) SetSuccess(v bool) *ListClusterConnectionTypesResponseBody {
	s.Success = &v
	return s
}

type ListClusterConnectionTypesResponseBodyData struct {
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterConnectionTypesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterConnectionTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponseBodyData) SetShowName(v string) *ListClusterConnectionTypesResponseBodyData {
	s.ShowName = &v
	return s
}

type ListClusterConnectionTypesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterConnectionTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterConnectionTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterConnectionTypesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponse) SetHeaders(v map[string]*string) *ListClusterConnectionTypesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterConnectionTypesResponse) SetBody(v *ListClusterConnectionTypesResponseBody) *ListClusterConnectionTypesResponse {
	s.Body = v
	return s
}

type ListClusterTypesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterTypesRequest) SetRegionId(v string) *ListClusterTypesRequest {
	s.RegionId = &v
	return s
}

type ListClusterTypesResponseBody struct {
	Code           *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListClusterTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponseBody) SetCode(v int32) *ListClusterTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetData(v []*ListClusterTypesResponseBodyData) *ListClusterTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterTypesResponseBody) SetDynamicMessage(v string) *ListClusterTypesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetErrorCode(v string) *ListClusterTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetHttpStatusCode(v int32) *ListClusterTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetMessage(v string) *ListClusterTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetRequestId(v string) *ListClusterTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTypesResponseBody) SetSuccess(v bool) *ListClusterTypesResponseBody {
	s.Success = &v
	return s
}

type ListClusterTypesResponseBodyData struct {
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterTypesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponseBodyData) SetShowName(v string) *ListClusterTypesResponseBodyData {
	s.ShowName = &v
	return s
}

type ListClusterTypesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponse) SetHeaders(v map[string]*string) *ListClusterTypesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTypesResponse) SetBody(v *ListClusterTypesResponseBody) *ListClusterTypesResponse {
	s.Body = v
	return s
}

type ListClusterVersionsRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s ListClusterVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsRequest) SetClusterType(v string) *ListClusterVersionsRequest {
	s.ClusterType = &v
	return s
}

type ListClusterVersionsResponseBody struct {
	Code           *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListClusterVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	DynamicMessage *string                                `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponseBody) SetCode(v int32) *ListClusterVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetData(v []*ListClusterVersionsResponseBodyData) *ListClusterVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterVersionsResponseBody) SetDynamicMessage(v string) *ListClusterVersionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetErrorCode(v string) *ListClusterVersionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetHttpStatusCode(v int32) *ListClusterVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetMessage(v string) *ListClusterVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetRequestId(v string) *ListClusterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterVersionsResponseBody) SetSuccess(v bool) *ListClusterVersionsResponseBody {
	s.Success = &v
	return s
}

type ListClusterVersionsResponseBodyData struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ShowName    *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s ListClusterVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponseBodyData) SetClusterType(v string) *ListClusterVersionsResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClusterVersionsResponseBodyData) SetCode(v string) *ListClusterVersionsResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListClusterVersionsResponseBodyData) SetShowName(v string) *ListClusterVersionsResponseBodyData {
	s.ShowName = &v
	return s
}

type ListClusterVersionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponse) SetHeaders(v map[string]*string) *ListClusterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterVersionsResponse) SetBody(v *ListClusterVersionsResponseBody) *ListClusterVersionsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestPars      *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetClusterAliasName(v string) *ListClustersRequest {
	s.ClusterAliasName = &v
	return s
}

func (s *ListClustersRequest) SetPageNum(v int32) *ListClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) SetRequestPars(v string) *ListClustersRequest {
	s.RequestPars = &v
	return s
}

type ListClustersResponseBody struct {
	Data       []*ListClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                         `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetErrorCode(v string) *ListClustersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClustersResponseBody) SetHttpCode(v string) *ListClustersResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetSuccess(v bool) *ListClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

type ListClustersResponseBodyData struct {
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CanUpdate        *bool   `json:"CanUpdate,omitempty" xml:"CanUpdate,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	ClusterName      *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType      *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndDate          *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InitStatus       *string `json:"InitStatus,omitempty" xml:"InitStatus,omitempty"`
	InstanceCount    *int64  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetAddress  *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	InternetDomain   *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	IntranetAddress  *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	IntranetDomain   *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) SetAppVersion(v string) *ListClustersResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCanUpdate(v bool) *ListClustersResponseBodyData {
	s.CanUpdate = &v
	return s
}

func (s *ListClustersResponseBodyData) SetChargeType(v string) *ListClustersResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterAliasName(v string) *ListClustersResponseBodyData {
	s.ClusterAliasName = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterName(v string) *ListClustersResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterType(v string) *ListClustersResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCreateTime(v string) *ListClustersResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyData) SetEndDate(v string) *ListClustersResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInitStatus(v string) *ListClustersResponseBodyData {
	s.InitStatus = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInstanceCount(v int64) *ListClustersResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInstanceId(v string) *ListClustersResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInternetAddress(v string) *ListClustersResponseBodyData {
	s.InternetAddress = &v
	return s
}

func (s *ListClustersResponseBodyData) SetInternetDomain(v string) *ListClustersResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *ListClustersResponseBodyData) SetIntranetAddress(v string) *ListClustersResponseBodyData {
	s.IntranetAddress = &v
	return s
}

func (s *ListClustersResponseBodyData) SetIntranetDomain(v string) *ListClustersResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *ListClustersResponseBodyData) SetVersionCode(v string) *ListClustersResponseBodyData {
	s.VersionCode = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListEngineNamespacesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEngineNamespacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEngineNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesRequest) SetInstanceId(v string) *ListEngineNamespacesRequest {
	s.InstanceId = &v
	return s
}

type ListEngineNamespacesResponseBody struct {
	Data       []*ListEngineNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                                 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEngineNamespacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEngineNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponseBody) SetData(v []*ListEngineNamespacesResponseBodyData) *ListEngineNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetErrorCode(v string) *ListEngineNamespacesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetHttpCode(v string) *ListEngineNamespacesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetMessage(v string) *ListEngineNamespacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetPageNumber(v int32) *ListEngineNamespacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetPageSize(v int32) *ListEngineNamespacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetRequestId(v string) *ListEngineNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetSuccess(v bool) *ListEngineNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEngineNamespacesResponseBody) SetTotalCount(v int32) *ListEngineNamespacesResponseBody {
	s.TotalCount = &v
	return s
}

type ListEngineNamespacesResponseBodyData struct {
	ConfigCount       *int32  `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceDesc     *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	Quota             *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	ServiceCount      *string `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEngineNamespacesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEngineNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponseBodyData) SetConfigCount(v int32) *ListEngineNamespacesResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespace(v string) *ListEngineNamespacesResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespaceDesc(v string) *ListEngineNamespacesResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetNamespaceShowName(v string) *ListEngineNamespacesResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetQuota(v int32) *ListEngineNamespacesResponseBodyData {
	s.Quota = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetServiceCount(v string) *ListEngineNamespacesResponseBodyData {
	s.ServiceCount = &v
	return s
}

func (s *ListEngineNamespacesResponseBodyData) SetType(v int32) *ListEngineNamespacesResponseBodyData {
	s.Type = &v
	return s
}

type ListEngineNamespacesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEngineNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEngineNamespacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEngineNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesResponse) SetHeaders(v map[string]*string) *ListEngineNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListEngineNamespacesResponse) SetBody(v *ListEngineNamespacesResponseBody) *ListEngineNamespacesResponse {
	s.Body = v
	return s
}

type ListEurekaInstancesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListEurekaInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesRequest) SetClusterId(v string) *ListEurekaInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetPageNum(v int32) *ListEurekaInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetPageSize(v int32) *ListEurekaInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetRequestPars(v string) *ListEurekaInstancesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListEurekaInstancesRequest) SetServiceName(v string) *ListEurekaInstancesRequest {
	s.ServiceName = &v
	return s
}

type ListEurekaInstancesResponseBody struct {
	Data       []*ListEurekaInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEurekaInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponseBody) SetData(v []*ListEurekaInstancesResponseBodyData) *ListEurekaInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetErrorCode(v string) *ListEurekaInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetHttpCode(v string) *ListEurekaInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetMessage(v string) *ListEurekaInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetPageNumber(v int32) *ListEurekaInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetPageSize(v int32) *ListEurekaInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetRequestId(v string) *ListEurekaInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetSuccess(v bool) *ListEurekaInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEurekaInstancesResponseBody) SetTotalCount(v int32) *ListEurekaInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListEurekaInstancesResponseBodyData struct {
	App                   *string                `json:"App,omitempty" xml:"App,omitempty"`
	DurationInSecs        *int32                 `json:"DurationInSecs,omitempty" xml:"DurationInSecs,omitempty"`
	HomePageUrl           *string                `json:"HomePageUrl,omitempty" xml:"HomePageUrl,omitempty"`
	HostName              *string                `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId            *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpAddr                *string                `json:"IpAddr,omitempty" xml:"IpAddr,omitempty"`
	LastDirtyTimestamp    *int64                 `json:"LastDirtyTimestamp,omitempty" xml:"LastDirtyTimestamp,omitempty"`
	LastUpdatedTimestamp  *int64                 `json:"LastUpdatedTimestamp,omitempty" xml:"LastUpdatedTimestamp,omitempty"`
	Metadata              map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Port                  *int32                 `json:"Port,omitempty" xml:"Port,omitempty"`
	RenewalIntervalInSecs *int32                 `json:"RenewalIntervalInSecs,omitempty" xml:"RenewalIntervalInSecs,omitempty"`
	SecurePort            *int32                 `json:"SecurePort,omitempty" xml:"SecurePort,omitempty"`
	Status                *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	VipAddress            *string                `json:"VipAddress,omitempty" xml:"VipAddress,omitempty"`
}

func (s ListEurekaInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponseBodyData) SetApp(v string) *ListEurekaInstancesResponseBodyData {
	s.App = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetDurationInSecs(v int32) *ListEurekaInstancesResponseBodyData {
	s.DurationInSecs = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetHomePageUrl(v string) *ListEurekaInstancesResponseBodyData {
	s.HomePageUrl = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetHostName(v string) *ListEurekaInstancesResponseBodyData {
	s.HostName = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetInstanceId(v string) *ListEurekaInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetIpAddr(v string) *ListEurekaInstancesResponseBodyData {
	s.IpAddr = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetLastDirtyTimestamp(v int64) *ListEurekaInstancesResponseBodyData {
	s.LastDirtyTimestamp = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetLastUpdatedTimestamp(v int64) *ListEurekaInstancesResponseBodyData {
	s.LastUpdatedTimestamp = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetMetadata(v map[string]interface{}) *ListEurekaInstancesResponseBodyData {
	s.Metadata = v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetPort(v int32) *ListEurekaInstancesResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetRenewalIntervalInSecs(v int32) *ListEurekaInstancesResponseBodyData {
	s.RenewalIntervalInSecs = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetSecurePort(v int32) *ListEurekaInstancesResponseBodyData {
	s.SecurePort = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetStatus(v string) *ListEurekaInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListEurekaInstancesResponseBodyData) SetVipAddress(v string) *ListEurekaInstancesResponseBodyData {
	s.VipAddress = &v
	return s
}

type ListEurekaInstancesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEurekaInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEurekaInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponse) SetHeaders(v map[string]*string) *ListEurekaInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEurekaInstancesResponse) SetBody(v *ListEurekaInstancesResponseBody) *ListEurekaInstancesResponse {
	s.Body = v
	return s
}

type ListEurekaServicesRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListEurekaServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaServicesRequest) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesRequest) SetClusterId(v string) *ListEurekaServicesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListEurekaServicesRequest) SetPageNum(v int32) *ListEurekaServicesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEurekaServicesRequest) SetPageSize(v int32) *ListEurekaServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEurekaServicesRequest) SetRegionId(v string) *ListEurekaServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEurekaServicesRequest) SetRequestPars(v string) *ListEurekaServicesRequest {
	s.RequestPars = &v
	return s
}

type ListEurekaServicesResponseBody struct {
	Data       []*ListEurekaServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                               `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEurekaServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponseBody) SetData(v []*ListEurekaServicesResponseBodyData) *ListEurekaServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListEurekaServicesResponseBody) SetErrorCode(v string) *ListEurekaServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetHttpCode(v string) *ListEurekaServicesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetMessage(v string) *ListEurekaServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetPageNumber(v int32) *ListEurekaServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetPageSize(v int32) *ListEurekaServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetRequestId(v string) *ListEurekaServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetSuccess(v bool) *ListEurekaServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEurekaServicesResponseBody) SetTotalCount(v int32) *ListEurekaServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListEurekaServicesResponseBodyData struct {
	InstancesId []*string `json:"InstancesId,omitempty" xml:"InstancesId,omitempty" type:"Repeated"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	UpStatus    *string   `json:"UpStatus,omitempty" xml:"UpStatus,omitempty"`
}

func (s ListEurekaServicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponseBodyData) SetInstancesId(v []*string) *ListEurekaServicesResponseBodyData {
	s.InstancesId = v
	return s
}

func (s *ListEurekaServicesResponseBodyData) SetName(v string) *ListEurekaServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListEurekaServicesResponseBodyData) SetUpStatus(v string) *ListEurekaServicesResponseBodyData {
	s.UpStatus = &v
	return s
}

type ListEurekaServicesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEurekaServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEurekaServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEurekaServicesResponse) GoString() string {
	return s.String()
}

func (s *ListEurekaServicesResponse) SetHeaders(v map[string]*string) *ListEurekaServicesResponse {
	s.Headers = v
	return s
}

func (s *ListEurekaServicesResponse) SetBody(v *ListEurekaServicesResponseBody) *ListEurekaServicesResponse {
	s.Body = v
	return s
}

type ListGatewayRequest struct {
	DescSort     *bool                           `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	FilterParams *ListGatewayRequestFilterParams `json:"FilterParams,omitempty" xml:"FilterParams,omitempty" type:"Struct"`
	OrderItem    *string                         `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	PageNumber   *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayRequest) SetDescSort(v bool) *ListGatewayRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayRequest) SetFilterParams(v *ListGatewayRequestFilterParams) *ListGatewayRequest {
	s.FilterParams = v
	return s
}

func (s *ListGatewayRequest) SetOrderItem(v string) *ListGatewayRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayRequest) SetPageNumber(v int32) *ListGatewayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRequest) SetPageSize(v int32) *ListGatewayRequest {
	s.PageSize = &v
	return s
}

type ListGatewayRequestFilterParams struct {
	GatewayType     *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Vpc             *string `json:"Vpc,omitempty" xml:"Vpc,omitempty"`
}

func (s ListGatewayRequestFilterParams) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayRequestFilterParams) GoString() string {
	return s.String()
}

func (s *ListGatewayRequestFilterParams) SetGatewayType(v string) *ListGatewayRequestFilterParams {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetGatewayUniqueId(v string) *ListGatewayRequestFilterParams {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetInstanceId(v string) *ListGatewayRequestFilterParams {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetName(v string) *ListGatewayRequestFilterParams {
	s.Name = &v
	return s
}

func (s *ListGatewayRequestFilterParams) SetVpc(v string) *ListGatewayRequestFilterParams {
	s.Vpc = &v
	return s
}

type ListGatewayShrinkRequest struct {
	DescSort           *bool   `json:"DescSort,omitempty" xml:"DescSort,omitempty"`
	FilterParamsShrink *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
	OrderItem          *string `json:"OrderItem,omitempty" xml:"OrderItem,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGatewayShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayShrinkRequest) SetDescSort(v bool) *ListGatewayShrinkRequest {
	s.DescSort = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetFilterParamsShrink(v string) *ListGatewayShrinkRequest {
	s.FilterParamsShrink = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetOrderItem(v string) *ListGatewayShrinkRequest {
	s.OrderItem = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetPageNumber(v int32) *ListGatewayShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayShrinkRequest) SetPageSize(v int32) *ListGatewayShrinkRequest {
	s.PageSize = &v
	return s
}

type ListGatewayResponseBody struct {
	Code           *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBody) SetCode(v int32) *ListGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayResponseBody) SetData(v *ListGatewayResponseBodyData) *ListGatewayResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayResponseBody) SetHttpStatusCode(v int32) *ListGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayResponseBody) SetMessage(v string) *ListGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayResponseBody) SetRequestId(v string) *ListGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayResponseBody) SetSuccess(v bool) *ListGatewayResponseBody {
	s.Success = &v
	return s
}

type ListGatewayResponseBodyData struct {
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result     []*ListGatewayResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalSize  *int64                               `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyData) SetPageNumber(v int32) *ListGatewayResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPageSize(v int32) *ListGatewayResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetResult(v []*ListGatewayResponseBodyDataResult) *ListGatewayResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayResponseBodyData) SetTotalSize(v int64) *ListGatewayResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListGatewayResponseBodyDataResult struct {
	AhasOn          *bool                                           `json:"AhasOn,omitempty" xml:"AhasOn,omitempty"`
	ArmsOn          *bool                                           `json:"ArmsOn,omitempty" xml:"ArmsOn,omitempty"`
	ChargeType      *string                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CurrentVersion  *string                                         `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	EndDate         *string                                         `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GatewayType     *string                                         `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	GatewayUniqueId *string                                         `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	GmtCreate       *string                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string                                         `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id              *int64                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InternetSlb     []*ListGatewayResponseBodyDataResultInternetSlb `json:"InternetSlb,omitempty" xml:"InternetSlb,omitempty" type:"Repeated"`
	LatestVersion   *string                                         `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	MustUpgrade     *bool                                           `json:"MustUpgrade,omitempty" xml:"MustUpgrade,omitempty"`
	Name            *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PrimaryUser     *string                                         `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	Region          *string                                         `json:"Region,omitempty" xml:"Region,omitempty"`
	Replica         *int32                                          `json:"Replica,omitempty" xml:"Replica,omitempty"`
	Slb             []*ListGatewayResponseBodyDataResultSlb         `json:"Slb,omitempty" xml:"Slb,omitempty" type:"Repeated"`
	Spec            *string                                         `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status          *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDesc      *string                                         `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	Tag             *string                                         `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Upgrade         *bool                                           `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	Vswitch2        *string                                         `json:"Vswitch2,omitempty" xml:"Vswitch2,omitempty"`
}

func (s ListGatewayResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResult) SetAhasOn(v bool) *ListGatewayResponseBodyDataResult {
	s.AhasOn = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetArmsOn(v bool) *ListGatewayResponseBodyDataResult {
	s.ArmsOn = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetChargeType(v string) *ListGatewayResponseBodyDataResult {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetCurrentVersion(v string) *ListGatewayResponseBodyDataResult {
	s.CurrentVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetEndDate(v string) *ListGatewayResponseBodyDataResult {
	s.EndDate = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayType(v string) *ListGatewayResponseBodyDataResult {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGmtModified(v string) *ListGatewayResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetId(v int64) *ListGatewayResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetInstanceId(v string) *ListGatewayResponseBodyDataResult {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetInternetSlb(v []*ListGatewayResponseBodyDataResultInternetSlb) *ListGatewayResponseBodyDataResult {
	s.InternetSlb = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetLatestVersion(v string) *ListGatewayResponseBodyDataResult {
	s.LatestVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetMustUpgrade(v bool) *ListGatewayResponseBodyDataResult {
	s.MustUpgrade = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetName(v string) *ListGatewayResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetPrimaryUser(v string) *ListGatewayResponseBodyDataResult {
	s.PrimaryUser = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetRegion(v string) *ListGatewayResponseBodyDataResult {
	s.Region = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetReplica(v int32) *ListGatewayResponseBodyDataResult {
	s.Replica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetSlb(v []*ListGatewayResponseBodyDataResultSlb) *ListGatewayResponseBodyDataResult {
	s.Slb = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetSpec(v string) *ListGatewayResponseBodyDataResult {
	s.Spec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetStatus(v int32) *ListGatewayResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetStatusDesc(v string) *ListGatewayResponseBodyDataResult {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetTag(v string) *ListGatewayResponseBodyDataResult {
	s.Tag = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetUpgrade(v bool) *ListGatewayResponseBodyDataResult {
	s.Upgrade = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetVswitch2(v string) *ListGatewayResponseBodyDataResult {
	s.Vswitch2 = &v
	return s
}

type ListGatewayResponseBodyDataResultInternetSlb struct {
	GatewaySlbMode      *string `json:"GatewaySlbMode,omitempty" xml:"GatewaySlbMode,omitempty"`
	GatewaySlbStatus    *string `json:"GatewaySlbStatus,omitempty" xml:"GatewaySlbStatus,omitempty"`
	InternetNetworkFlow *string `json:"InternetNetworkFlow,omitempty" xml:"InternetNetworkFlow,omitempty"`
	SlbId               *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp               *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SlbPort             *string `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	SlbSpec             *string `json:"SlbSpec,omitempty" xml:"SlbSpec,omitempty"`
	StatusDesc          *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayResponseBodyDataResultInternetSlb) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultInternetSlb) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetGatewaySlbMode(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.GatewaySlbMode = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetGatewaySlbStatus(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.GatewaySlbStatus = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetInternetNetworkFlow(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.InternetNetworkFlow = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbId(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbIp(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbIp = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbPort(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbPort = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbSpec(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbSpec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetStatusDesc(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetType(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.Type = &v
	return s
}

type ListGatewayResponseBodyDataResultSlb struct {
	GatewaySlbMode   *string `json:"GatewaySlbMode,omitempty" xml:"GatewaySlbMode,omitempty"`
	GatewaySlbStatus *string `json:"GatewaySlbStatus,omitempty" xml:"GatewaySlbStatus,omitempty"`
	SlbId            *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp            *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SlbPort          *string `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	SlbSpec          *string `json:"SlbSpec,omitempty" xml:"SlbSpec,omitempty"`
	StatusDesc       *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayResponseBodyDataResultSlb) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultSlb) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultSlb) SetGatewaySlbMode(v string) *ListGatewayResponseBodyDataResultSlb {
	s.GatewaySlbMode = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetGatewaySlbStatus(v string) *ListGatewayResponseBodyDataResultSlb {
	s.GatewaySlbStatus = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbId(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbIp(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbIp = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbPort(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbPort = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbSpec(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbSpec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetStatusDesc(v string) *ListGatewayResponseBodyDataResultSlb {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetType(v string) *ListGatewayResponseBodyDataResultSlb {
	s.Type = &v
	return s
}

type ListGatewayResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGatewayResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGatewayResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayResponse) SetHeaders(v map[string]*string) *ListGatewayResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayResponse) SetBody(v *ListGatewayResponseBody) *ListGatewayResponse {
	s.Body = v
	return s
}

type ListListenersByConfigRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListListenersByConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByConfigRequest) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigRequest) SetDataId(v string) *ListListenersByConfigRequest {
	s.DataId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetGroup(v string) *ListListenersByConfigRequest {
	s.Group = &v
	return s
}

func (s *ListListenersByConfigRequest) SetInstanceId(v string) *ListListenersByConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetNamespaceId(v string) *ListListenersByConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByConfigRequest) SetRequestPars(v string) *ListListenersByConfigRequest {
	s.RequestPars = &v
	return s
}

type ListListenersByConfigResponseBody struct {
	ErrorCode  *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                                       `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Listeners  []*ListListenersByConfigResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersByConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponseBody) SetErrorCode(v string) *ListListenersByConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetHttpCode(v string) *ListListenersByConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetListeners(v []*ListListenersByConfigResponseBodyListeners) *ListListenersByConfigResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersByConfigResponseBody) SetMessage(v string) *ListListenersByConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetPageNumber(v int32) *ListListenersByConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetPageSize(v int32) *ListListenersByConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetRequestId(v string) *ListListenersByConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetSuccess(v bool) *ListListenersByConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListListenersByConfigResponseBody) SetTotalCount(v int32) *ListListenersByConfigResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersByConfigResponseBodyListeners struct {
	Ip  *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s ListListenersByConfigResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByConfigResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponseBodyListeners) SetIp(v string) *ListListenersByConfigResponseBodyListeners {
	s.Ip = &v
	return s
}

func (s *ListListenersByConfigResponseBodyListeners) SetMd5(v string) *ListListenersByConfigResponseBodyListeners {
	s.Md5 = &v
	return s
}

type ListListenersByConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListListenersByConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersByConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByConfigResponse) GoString() string {
	return s.String()
}

func (s *ListListenersByConfigResponse) SetHeaders(v map[string]*string) *ListListenersByConfigResponse {
	s.Headers = v
	return s
}

func (s *ListListenersByConfigResponse) SetBody(v *ListListenersByConfigResponseBody) *ListListenersByConfigResponse {
	s.Body = v
	return s
}

type ListListenersByIpRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListListenersByIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByIpRequest) GoString() string {
	return s.String()
}

func (s *ListListenersByIpRequest) SetInstanceId(v string) *ListListenersByIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ListListenersByIpRequest) SetIp(v string) *ListListenersByIpRequest {
	s.Ip = &v
	return s
}

func (s *ListListenersByIpRequest) SetNamespaceId(v string) *ListListenersByIpRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListListenersByIpRequest) SetRequestPars(v string) *ListListenersByIpRequest {
	s.RequestPars = &v
	return s
}

type ListListenersByIpResponseBody struct {
	ErrorCode  *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode   *string                                   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Listeners  []*ListListenersByIpResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersByIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponseBody) SetErrorCode(v string) *ListListenersByIpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetHttpCode(v string) *ListListenersByIpResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetListeners(v []*ListListenersByIpResponseBodyListeners) *ListListenersByIpResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersByIpResponseBody) SetMessage(v string) *ListListenersByIpResponseBody {
	s.Message = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetPageNumber(v int32) *ListListenersByIpResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetPageSize(v int32) *ListListenersByIpResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetRequestId(v string) *ListListenersByIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetSuccess(v bool) *ListListenersByIpResponseBody {
	s.Success = &v
	return s
}

func (s *ListListenersByIpResponseBody) SetTotalCount(v int32) *ListListenersByIpResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersByIpResponseBodyListeners struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group  *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Md5    *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s ListListenersByIpResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByIpResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponseBodyListeners) SetDataId(v string) *ListListenersByIpResponseBodyListeners {
	s.DataId = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) SetGroup(v string) *ListListenersByIpResponseBodyListeners {
	s.Group = &v
	return s
}

func (s *ListListenersByIpResponseBodyListeners) SetMd5(v string) *ListListenersByIpResponseBodyListeners {
	s.Md5 = &v
	return s
}

type ListListenersByIpResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListListenersByIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersByIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersByIpResponse) GoString() string {
	return s.String()
}

func (s *ListListenersByIpResponse) SetHeaders(v map[string]*string) *ListListenersByIpResponse {
	s.Headers = v
	return s
}

func (s *ListListenersByIpResponse) SetBody(v *ListListenersByIpResponseBody) *ListListenersByIpResponse {
	s.Body = v
	return s
}

type ListNacosConfigsRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListNacosConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNacosConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsRequest) SetAppName(v string) *ListNacosConfigsRequest {
	s.AppName = &v
	return s
}

func (s *ListNacosConfigsRequest) SetDataId(v string) *ListNacosConfigsRequest {
	s.DataId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetGroup(v string) *ListNacosConfigsRequest {
	s.Group = &v
	return s
}

func (s *ListNacosConfigsRequest) SetInstanceId(v string) *ListNacosConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetNamespaceId(v string) *ListNacosConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetPageNum(v int32) *ListNacosConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListNacosConfigsRequest) SetPageSize(v int32) *ListNacosConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacosConfigsRequest) SetRegionId(v string) *ListNacosConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNacosConfigsRequest) SetRequestPars(v string) *ListNacosConfigsRequest {
	s.RequestPars = &v
	return s
}

func (s *ListNacosConfigsRequest) SetTags(v string) *ListNacosConfigsRequest {
	s.Tags = &v
	return s
}

type ListNacosConfigsResponseBody struct {
	Code           *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Configurations []*ListNacosConfigsResponseBodyConfigurations `json:"Configurations,omitempty" xml:"Configurations,omitempty" type:"Repeated"`
	ErrorCode      *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode       *string                                       `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNacosConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNacosConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponseBody) SetCode(v int32) *ListNacosConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetConfigurations(v []*ListNacosConfigsResponseBodyConfigurations) *ListNacosConfigsResponseBody {
	s.Configurations = v
	return s
}

func (s *ListNacosConfigsResponseBody) SetErrorCode(v string) *ListNacosConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetHttpCode(v string) *ListNacosConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetMessage(v string) *ListNacosConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetPageNumber(v int32) *ListNacosConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetPageSize(v int32) *ListNacosConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetRequestId(v string) *ListNacosConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetSuccess(v bool) *ListNacosConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNacosConfigsResponseBody) SetTotalCount(v int32) *ListNacosConfigsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNacosConfigsResponseBodyConfigurations struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DataId  *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNacosConfigsResponseBodyConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListNacosConfigsResponseBodyConfigurations) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetAppName(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.AppName = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetDataId(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.DataId = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetGroup(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.Group = &v
	return s
}

func (s *ListNacosConfigsResponseBodyConfigurations) SetId(v string) *ListNacosConfigsResponseBodyConfigurations {
	s.Id = &v
	return s
}

type ListNacosConfigsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNacosConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNacosConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNacosConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListNacosConfigsResponse) SetHeaders(v map[string]*string) *ListNacosConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListNacosConfigsResponse) SetBody(v *ListNacosConfigsResponseBody) *ListNacosConfigsResponse {
	s.Body = v
	return s
}

type ListNacosHistoryConfigsRequest struct {
	DataId      *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group       *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListNacosHistoryConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNacosHistoryConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsRequest) SetDataId(v string) *ListNacosHistoryConfigsRequest {
	s.DataId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetGroup(v string) *ListNacosHistoryConfigsRequest {
	s.Group = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetInstanceId(v string) *ListNacosHistoryConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetNamespaceId(v string) *ListNacosHistoryConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetPageNum(v int32) *ListNacosHistoryConfigsRequest {
	s.PageNum = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetPageSize(v int32) *ListNacosHistoryConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetRegionId(v string) *ListNacosHistoryConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNacosHistoryConfigsRequest) SetRequestPars(v string) *ListNacosHistoryConfigsRequest {
	s.RequestPars = &v
	return s
}

type ListNacosHistoryConfigsResponseBody struct {
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HistoryItems []*ListNacosHistoryConfigsResponseBodyHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
	HttpCode     *string                                            `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber   *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount   *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNacosHistoryConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNacosHistoryConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponseBody) SetErrorCode(v string) *ListNacosHistoryConfigsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetHistoryItems(v []*ListNacosHistoryConfigsResponseBodyHistoryItems) *ListNacosHistoryConfigsResponseBody {
	s.HistoryItems = v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetHttpCode(v string) *ListNacosHistoryConfigsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetMessage(v string) *ListNacosHistoryConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetPageNumber(v int32) *ListNacosHistoryConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetPageSize(v int32) *ListNacosHistoryConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetRequestId(v string) *ListNacosHistoryConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetSuccess(v bool) *ListNacosHistoryConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBody) SetTotalCount(v int32) *ListNacosHistoryConfigsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNacosHistoryConfigsResponseBodyHistoryItems struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Group            *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifiedTime *int64  `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	OpType           *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s ListNacosHistoryConfigsResponseBodyHistoryItems) String() string {
	return tea.Prettify(s)
}

func (s ListNacosHistoryConfigsResponseBodyHistoryItems) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetAppName(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.AppName = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetDataId(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.DataId = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetGroup(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.Group = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetId(v int64) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.Id = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetLastModifiedTime(v int64) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.LastModifiedTime = &v
	return s
}

func (s *ListNacosHistoryConfigsResponseBodyHistoryItems) SetOpType(v string) *ListNacosHistoryConfigsResponseBodyHistoryItems {
	s.OpType = &v
	return s
}

type ListNacosHistoryConfigsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNacosHistoryConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNacosHistoryConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNacosHistoryConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListNacosHistoryConfigsResponse) SetHeaders(v map[string]*string) *ListNacosHistoryConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListNacosHistoryConfigsResponse) SetBody(v *ListNacosHistoryConfigsResponseBody) *ListNacosHistoryConfigsResponse {
	s.Body = v
	return s
}

type ListZnodeChildrenRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Path      *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListZnodeChildrenRequest) String() string {
	return tea.Prettify(s)
}

func (s ListZnodeChildrenRequest) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenRequest) SetClusterId(v string) *ListZnodeChildrenRequest {
	s.ClusterId = &v
	return s
}

func (s *ListZnodeChildrenRequest) SetPath(v string) *ListZnodeChildrenRequest {
	s.Path = &v
	return s
}

type ListZnodeChildrenResponseBody struct {
	Data      []*ListZnodeChildrenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListZnodeChildrenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZnodeChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponseBody) SetData(v []*ListZnodeChildrenResponseBodyData) *ListZnodeChildrenResponseBody {
	s.Data = v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetErrorCode(v string) *ListZnodeChildrenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetMessage(v string) *ListZnodeChildrenResponseBody {
	s.Message = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetRequestId(v string) *ListZnodeChildrenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZnodeChildrenResponseBody) SetSuccess(v bool) *ListZnodeChildrenResponseBody {
	s.Success = &v
	return s
}

type ListZnodeChildrenResponseBodyData struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Dir  *bool   `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListZnodeChildrenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListZnodeChildrenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponseBodyData) SetData(v string) *ListZnodeChildrenResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetDir(v bool) *ListZnodeChildrenResponseBodyData {
	s.Dir = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetName(v string) *ListZnodeChildrenResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListZnodeChildrenResponseBodyData) SetPath(v string) *ListZnodeChildrenResponseBodyData {
	s.Path = &v
	return s
}

type ListZnodeChildrenResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListZnodeChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListZnodeChildrenResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZnodeChildrenResponse) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenResponse) SetHeaders(v map[string]*string) *ListZnodeChildrenResponse {
	s.Headers = v
	return s
}

func (s *ListZnodeChildrenResponse) SetBody(v *ListZnodeChildrenResponseBody) *ListZnodeChildrenResponse {
	s.Body = v
	return s
}

type ModifyGovernanceKubernetesClusterRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NamespaceInfos *string `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetClusterId(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetNamespaceInfos(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.NamespaceInfos = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetRegionId(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.RegionId = &v
	return s
}

type ModifyGovernanceKubernetesClusterResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetData(v bool) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetHttpStatusCode(v int32) *ModifyGovernanceKubernetesClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetMessage(v string) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetRequestId(v string) *ModifyGovernanceKubernetesClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponseBody) SetSuccess(v bool) *ModifyGovernanceKubernetesClusterResponseBody {
	s.Success = &v
	return s
}

type ModifyGovernanceKubernetesClusterResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGovernanceKubernetesClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *ModifyGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponse) SetBody(v *ModifyGovernanceKubernetesClusterResponseBody) *ModifyGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

type QueryBusinessLocationsResponseBody struct {
	Data      []*QueryBusinessLocationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBusinessLocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBusinessLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponseBody) SetData(v []*QueryBusinessLocationsResponseBodyData) *QueryBusinessLocationsResponseBody {
	s.Data = v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetErrorCode(v string) *QueryBusinessLocationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetMessage(v string) *QueryBusinessLocationsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetRequestId(v string) *QueryBusinessLocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBusinessLocationsResponseBody) SetSuccess(v string) *QueryBusinessLocationsResponseBody {
	s.Success = &v
	return s
}

type QueryBusinessLocationsResponseBodyData struct {
	CnName           *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DistrictCnName   *string `json:"DistrictCnName,omitempty" xml:"DistrictCnName,omitempty"`
	DistrictEnName   *string `json:"DistrictEnName,omitempty" xml:"DistrictEnName,omitempty"`
	DistrictId       *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	DistrictOrdering *int32  `json:"DistrictOrdering,omitempty" xml:"DistrictOrdering,omitempty"`
	DistrictShowName *string `json:"DistrictShowName,omitempty" xml:"DistrictShowName,omitempty"`
	EnDescription    *string `json:"EnDescription,omitempty" xml:"EnDescription,omitempty"`
	EnName           *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Ordering         *int32  `json:"Ordering,omitempty" xml:"Ordering,omitempty"`
	ShowName         *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryBusinessLocationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBusinessLocationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponseBodyData) SetCnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.CnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDescription(v string) *QueryBusinessLocationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictCnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictCnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictEnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictEnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictId(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictId = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictOrdering(v int32) *QueryBusinessLocationsResponseBodyData {
	s.DistrictOrdering = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetDistrictShowName(v string) *QueryBusinessLocationsResponseBodyData {
	s.DistrictShowName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetEnDescription(v string) *QueryBusinessLocationsResponseBodyData {
	s.EnDescription = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetEnName(v string) *QueryBusinessLocationsResponseBodyData {
	s.EnName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetName(v string) *QueryBusinessLocationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetOrdering(v int32) *QueryBusinessLocationsResponseBodyData {
	s.Ordering = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetShowName(v string) *QueryBusinessLocationsResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *QueryBusinessLocationsResponseBodyData) SetType(v string) *QueryBusinessLocationsResponseBodyData {
	s.Type = &v
	return s
}

type QueryBusinessLocationsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBusinessLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBusinessLocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBusinessLocationsResponse) GoString() string {
	return s.String()
}

func (s *QueryBusinessLocationsResponse) SetHeaders(v map[string]*string) *QueryBusinessLocationsResponse {
	s.Headers = v
	return s
}

func (s *QueryBusinessLocationsResponse) SetBody(v *QueryBusinessLocationsResponseBody) *QueryBusinessLocationsResponse {
	s.Body = v
	return s
}

type QueryClusterDetailRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s QueryClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailRequest) SetInstanceId(v string) *QueryClusterDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterDetailRequest) SetOrderId(v string) *QueryClusterDetailRequest {
	s.OrderId = &v
	return s
}

type QueryClusterDetailResponseBody struct {
	Data      *QueryClusterDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBody) SetData(v *QueryClusterDetailResponseBodyData) *QueryClusterDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterDetailResponseBody) SetErrorCode(v string) *QueryClusterDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetMessage(v string) *QueryClusterDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetRequestId(v string) *QueryClusterDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterDetailResponseBody) SetSuccess(v bool) *QueryClusterDetailResponseBody {
	s.Success = &v
	return s
}

type QueryClusterDetailResponseBodyData struct {
	AclEntryList         *string                                             `json:"AclEntryList,omitempty" xml:"AclEntryList,omitempty"`
	AclId                *string                                             `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AppVersion           *string                                             `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ChargeType           *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterAliasName     *string                                             `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	ClusterName          *string                                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterSpecification *string                                             `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	ClusterType          *string                                             `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ClusterVersion       *string                                             `json:"ClusterVersion,omitempty" xml:"ClusterVersion,omitempty"`
	ConnectionType       *string                                             `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	Cpu                  *int32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime           *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DiskCapacity         *int64                                              `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskType             *string                                             `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HealthStatus         *string                                             `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InitCostTime         *int64                                              `json:"InitCostTime,omitempty" xml:"InitCostTime,omitempty"`
	InitStatus           *string                                             `json:"InitStatus,omitempty" xml:"InitStatus,omitempty"`
	InstanceCount        *int32                                              `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceId           *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceModels       []*QueryClusterDetailResponseBodyDataInstanceModels `json:"InstanceModels,omitempty" xml:"InstanceModels,omitempty" type:"Repeated"`
	InternetAddress      *string                                             `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	InternetDomain       *string                                             `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	InternetPort         *string                                             `json:"InternetPort,omitempty" xml:"InternetPort,omitempty"`
	IntranetAddress      *string                                             `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	IntranetDomain       *string                                             `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	IntranetPort         *string                                             `json:"IntranetPort,omitempty" xml:"IntranetPort,omitempty"`
	MemoryCapacity       *int64                                              `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
	MseVersion           *string                                             `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	NetType              *string                                             `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PayInfo              *string                                             `json:"PayInfo,omitempty" xml:"PayInfo,omitempty"`
	PubNetworkFlow       *string                                             `json:"PubNetworkFlow,omitempty" xml:"PubNetworkFlow,omitempty"`
	RegionId             *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId            *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryClusterDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyData) SetAclEntryList(v string) *QueryClusterDetailResponseBodyData {
	s.AclEntryList = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetAclId(v string) *QueryClusterDetailResponseBodyData {
	s.AclId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetAppVersion(v string) *QueryClusterDetailResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetChargeType(v string) *QueryClusterDetailResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterAliasName(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterAliasName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterName(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterSpecification(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterSpecification = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterType(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetClusterVersion(v string) *QueryClusterDetailResponseBodyData {
	s.ClusterVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetConnectionType(v string) *QueryClusterDetailResponseBodyData {
	s.ConnectionType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetCpu(v int32) *QueryClusterDetailResponseBodyData {
	s.Cpu = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetCreateTime(v string) *QueryClusterDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetDiskCapacity(v int64) *QueryClusterDetailResponseBodyData {
	s.DiskCapacity = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetDiskType(v string) *QueryClusterDetailResponseBodyData {
	s.DiskType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetHealthStatus(v string) *QueryClusterDetailResponseBodyData {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInitCostTime(v int64) *QueryClusterDetailResponseBodyData {
	s.InitCostTime = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInitStatus(v string) *QueryClusterDetailResponseBodyData {
	s.InitStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceCount(v int32) *QueryClusterDetailResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceId(v string) *QueryClusterDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInstanceModels(v []*QueryClusterDetailResponseBodyDataInstanceModels) *QueryClusterDetailResponseBodyData {
	s.InstanceModels = v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetAddress(v string) *QueryClusterDetailResponseBodyData {
	s.InternetAddress = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetDomain(v string) *QueryClusterDetailResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetInternetPort(v string) *QueryClusterDetailResponseBodyData {
	s.InternetPort = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetAddress(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetAddress = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetDomain(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetIntranetPort(v string) *QueryClusterDetailResponseBodyData {
	s.IntranetPort = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetMemoryCapacity(v int64) *QueryClusterDetailResponseBodyData {
	s.MemoryCapacity = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetMseVersion(v string) *QueryClusterDetailResponseBodyData {
	s.MseVersion = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetNetType(v string) *QueryClusterDetailResponseBodyData {
	s.NetType = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetPayInfo(v string) *QueryClusterDetailResponseBodyData {
	s.PayInfo = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetPubNetworkFlow(v string) *QueryClusterDetailResponseBodyData {
	s.PubNetworkFlow = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetRegionId(v string) *QueryClusterDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetVSwitchId(v string) *QueryClusterDetailResponseBodyData {
	s.VSwitchId = &v
	return s
}

func (s *QueryClusterDetailResponseBodyData) SetVpcId(v string) *QueryClusterDetailResponseBodyData {
	s.VpcId = &v
	return s
}

type QueryClusterDetailResponseBodyDataInstanceModels struct {
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	HealthStatus      *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InternetIp        *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PodName           *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	Role              *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SingleTunnelVip   *string `json:"SingleTunnelVip,omitempty" xml:"SingleTunnelVip,omitempty"`
}

func (s QueryClusterDetailResponseBodyDataInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponseBodyDataInstanceModels) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetCreationTimestamp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.CreationTimestamp = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetHealthStatus(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.HealthStatus = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetInternetIp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.InternetIp = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetIp(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.Ip = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetPodName(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.PodName = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetRole(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.Role = &v
	return s
}

func (s *QueryClusterDetailResponseBodyDataInstanceModels) SetSingleTunnelVip(v string) *QueryClusterDetailResponseBodyDataInstanceModels {
	s.SingleTunnelVip = &v
	return s
}

type QueryClusterDetailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponse) SetHeaders(v map[string]*string) *QueryClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterDetailResponse) SetBody(v *QueryClusterDetailResponseBody) *QueryClusterDetailResponse {
	s.Body = v
	return s
}

type QueryClusterDiskSpecificationRequest struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
}

func (s QueryClusterDiskSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDiskSpecificationRequest) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationRequest) SetClusterType(v string) *QueryClusterDiskSpecificationRequest {
	s.ClusterType = &v
	return s
}

type QueryClusterDiskSpecificationResponseBody struct {
	Code           *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryClusterDiskSpecificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicMessage *string                                        `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterDiskSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponseBody) SetCode(v int32) *QueryClusterDiskSpecificationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetData(v *QueryClusterDiskSpecificationResponseBodyData) *QueryClusterDiskSpecificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetDynamicMessage(v string) *QueryClusterDiskSpecificationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetErrorCode(v string) *QueryClusterDiskSpecificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetHttpStatusCode(v int32) *QueryClusterDiskSpecificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetMessage(v string) *QueryClusterDiskSpecificationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetRequestId(v string) *QueryClusterDiskSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetSuccess(v bool) *QueryClusterDiskSpecificationResponseBody {
	s.Success = &v
	return s
}

type QueryClusterDiskSpecificationResponseBodyData struct {
	Max  *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min  *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s QueryClusterDiskSpecificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetMax(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Max = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetMin(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Min = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetStep(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Step = &v
	return s
}

type QueryClusterDiskSpecificationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClusterDiskSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClusterDiskSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponse) SetHeaders(v map[string]*string) *QueryClusterDiskSpecificationResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterDiskSpecificationResponse) SetBody(v *QueryClusterDiskSpecificationResponseBody) *QueryClusterDiskSpecificationResponse {
	s.Body = v
	return s
}

type QueryClusterSpecificationResponseBody struct {
	Code           *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*QueryClusterSpecificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode      *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterSpecificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponseBody) SetCode(v int32) *QueryClusterSpecificationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetData(v []*QueryClusterSpecificationResponseBodyData) *QueryClusterSpecificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetErrorCode(v string) *QueryClusterSpecificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetHttpStatusCode(v int32) *QueryClusterSpecificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetMessage(v string) *QueryClusterSpecificationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetRequestId(v string) *QueryClusterSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetSuccess(v bool) *QueryClusterSpecificationResponseBody {
	s.Success = &v
	return s
}

type QueryClusterSpecificationResponseBodyData struct {
	ClusterSpecificationName *string `json:"ClusterSpecificationName,omitempty" xml:"ClusterSpecificationName,omitempty"`
	CpuCapacity              *string `json:"CpuCapacity,omitempty" xml:"CpuCapacity,omitempty"`
	DiskCapacity             *string `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	InstanceCount            *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	MaxCon                   *string `json:"MaxCon,omitempty" xml:"MaxCon,omitempty"`
	MaxTps                   *string `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	MemoryCapacity           *string `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
}

func (s QueryClusterSpecificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterSpecificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponseBodyData) SetClusterSpecificationName(v string) *QueryClusterSpecificationResponseBodyData {
	s.ClusterSpecificationName = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetCpuCapacity(v string) *QueryClusterSpecificationResponseBodyData {
	s.CpuCapacity = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetDiskCapacity(v string) *QueryClusterSpecificationResponseBodyData {
	s.DiskCapacity = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetInstanceCount(v string) *QueryClusterSpecificationResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetMaxCon(v string) *QueryClusterSpecificationResponseBodyData {
	s.MaxCon = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetMaxTps(v string) *QueryClusterSpecificationResponseBodyData {
	s.MaxTps = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetMemoryCapacity(v string) *QueryClusterSpecificationResponseBodyData {
	s.MemoryCapacity = &v
	return s
}

type QueryClusterSpecificationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryClusterSpecificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryClusterSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryClusterSpecificationResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponse) SetHeaders(v map[string]*string) *QueryClusterSpecificationResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterSpecificationResponse) SetBody(v *QueryClusterSpecificationResponseBody) *QueryClusterSpecificationResponse {
	s.Body = v
	return s
}

type QueryConfigRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ConfigType  *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s QueryConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryConfigRequest) SetClusterId(v string) *QueryConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryConfigRequest) SetConfigType(v string) *QueryConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *QueryConfigRequest) SetInstanceId(v string) *QueryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConfigRequest) SetRequestPars(v string) *QueryConfigRequest {
	s.RequestPars = &v
	return s
}

type QueryConfigResponseBody struct {
	Code           *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBody) SetCode(v int32) *QueryConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConfigResponseBody) SetData(v *QueryConfigResponseBodyData) *QueryConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryConfigResponseBody) SetHttpStatusCode(v int32) *QueryConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryConfigResponseBody) SetMessage(v string) *QueryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConfigResponseBody) SetRequestId(v string) *QueryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConfigResponseBody) SetSuccess(v bool) *QueryConfigResponseBody {
	s.Success = &v
	return s
}

type QueryConfigResponseBodyData struct {
	AutopurgePurgeInterval   *string `json:"AutopurgePurgeInterval,omitempty" xml:"AutopurgePurgeInterval,omitempty"`
	AutopurgeSnapRetainCount *string `json:"AutopurgeSnapRetainCount,omitempty" xml:"AutopurgeSnapRetainCount,omitempty"`
	ClusterName              *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ConfigAuthEnabled        *bool   `json:"ConfigAuthEnabled,omitempty" xml:"ConfigAuthEnabled,omitempty"`
	ConfigAuthSupported      *bool   `json:"ConfigAuthSupported,omitempty" xml:"ConfigAuthSupported,omitempty"`
	ConfigSecretEnabled      *bool   `json:"ConfigSecretEnabled,omitempty" xml:"ConfigSecretEnabled,omitempty"`
	ConfigSecretSupported    *bool   `json:"ConfigSecretSupported,omitempty" xml:"ConfigSecretSupported,omitempty"`
	InitLimit                *string `json:"InitLimit,omitempty" xml:"InitLimit,omitempty"`
	JuteMaxbuffer            *string `json:"JuteMaxbuffer,omitempty" xml:"JuteMaxbuffer,omitempty"`
	JvmFlagsCustom           *string `json:"JvmFlagsCustom,omitempty" xml:"JvmFlagsCustom,omitempty"`
	MCPEnabled               *bool   `json:"MCPEnabled,omitempty" xml:"MCPEnabled,omitempty"`
	MCPSupported             *bool   `json:"MCPSupported,omitempty" xml:"MCPSupported,omitempty"`
	MaxClientCnxns           *string `json:"MaxClientCnxns,omitempty" xml:"MaxClientCnxns,omitempty"`
	OpenSuperAcl             *bool   `json:"OpenSuperAcl,omitempty" xml:"OpenSuperAcl,omitempty"`
	PassWord                 *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	RestartFlag              *bool   `json:"RestartFlag,omitempty" xml:"RestartFlag,omitempty"`
	SyncLimit                *string `json:"SyncLimit,omitempty" xml:"SyncLimit,omitempty"`
	TickTime                 *string `json:"TickTime,omitempty" xml:"TickTime,omitempty"`
	UserName                 *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryConfigResponseBodyData) SetAutopurgePurgeInterval(v string) *QueryConfigResponseBodyData {
	s.AutopurgePurgeInterval = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetAutopurgeSnapRetainCount(v string) *QueryConfigResponseBodyData {
	s.AutopurgeSnapRetainCount = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetClusterName(v string) *QueryConfigResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigAuthEnabled(v bool) *QueryConfigResponseBodyData {
	s.ConfigAuthEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigAuthSupported(v bool) *QueryConfigResponseBodyData {
	s.ConfigAuthSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigSecretEnabled(v bool) *QueryConfigResponseBodyData {
	s.ConfigSecretEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetConfigSecretSupported(v bool) *QueryConfigResponseBodyData {
	s.ConfigSecretSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetInitLimit(v string) *QueryConfigResponseBodyData {
	s.InitLimit = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetJuteMaxbuffer(v string) *QueryConfigResponseBodyData {
	s.JuteMaxbuffer = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetJvmFlagsCustom(v string) *QueryConfigResponseBodyData {
	s.JvmFlagsCustom = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMCPEnabled(v bool) *QueryConfigResponseBodyData {
	s.MCPEnabled = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMCPSupported(v bool) *QueryConfigResponseBodyData {
	s.MCPSupported = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetMaxClientCnxns(v string) *QueryConfigResponseBodyData {
	s.MaxClientCnxns = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetOpenSuperAcl(v bool) *QueryConfigResponseBodyData {
	s.OpenSuperAcl = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetPassWord(v string) *QueryConfigResponseBodyData {
	s.PassWord = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetRestartFlag(v bool) *QueryConfigResponseBodyData {
	s.RestartFlag = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetSyncLimit(v string) *QueryConfigResponseBodyData {
	s.SyncLimit = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetTickTime(v string) *QueryConfigResponseBodyData {
	s.TickTime = &v
	return s
}

func (s *QueryConfigResponseBodyData) SetUserName(v string) *QueryConfigResponseBodyData {
	s.UserName = &v
	return s
}

type QueryConfigResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryConfigResponse) SetHeaders(v map[string]*string) *QueryConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryConfigResponse) SetBody(v *QueryConfigResponseBody) *QueryConfigResponse {
	s.Body = v
	return s
}

type QueryGatewayRegionResponseBody struct {
	Code           *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGatewayRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGatewayRegionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGatewayRegionResponseBody) SetCode(v int32) *QueryGatewayRegionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetData(v []*string) *QueryGatewayRegionResponseBody {
	s.Data = v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetHttpStatusCode(v int32) *QueryGatewayRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetMessage(v string) *QueryGatewayRegionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetRequestId(v string) *QueryGatewayRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGatewayRegionResponseBody) SetSuccess(v bool) *QueryGatewayRegionResponseBody {
	s.Success = &v
	return s
}

type QueryGatewayRegionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGatewayRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGatewayRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGatewayRegionResponse) GoString() string {
	return s.String()
}

func (s *QueryGatewayRegionResponse) SetHeaders(v map[string]*string) *QueryGatewayRegionResponse {
	s.Headers = v
	return s
}

func (s *QueryGatewayRegionResponse) SetBody(v *QueryGatewayRegionResponseBody) *QueryGatewayRegionResponse {
	s.Body = v
	return s
}

type QueryGatewayTypeResponseBody struct {
	Code           *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGatewayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGatewayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGatewayTypeResponseBody) SetCode(v int32) *QueryGatewayTypeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetData(v []*string) *QueryGatewayTypeResponseBody {
	s.Data = v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetHttpStatusCode(v int32) *QueryGatewayTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetMessage(v string) *QueryGatewayTypeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetRequestId(v string) *QueryGatewayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGatewayTypeResponseBody) SetSuccess(v bool) *QueryGatewayTypeResponseBody {
	s.Success = &v
	return s
}

type QueryGatewayTypeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGatewayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGatewayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGatewayTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryGatewayTypeResponse) SetHeaders(v map[string]*string) *QueryGatewayTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryGatewayTypeResponse) SetBody(v *QueryGatewayTypeResponseBody) *QueryGatewayTypeResponse {
	s.Body = v
	return s
}

type QueryMonitorRequest struct {
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Step        *int64  `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s QueryMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorRequest) GoString() string {
	return s.String()
}

func (s *QueryMonitorRequest) SetEndTime(v int64) *QueryMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMonitorRequest) SetInstanceId(v string) *QueryMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMonitorRequest) SetMonitorType(v string) *QueryMonitorRequest {
	s.MonitorType = &v
	return s
}

func (s *QueryMonitorRequest) SetRequestPars(v string) *QueryMonitorRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryMonitorRequest) SetStartTime(v int64) *QueryMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMonitorRequest) SetStep(v int64) *QueryMonitorRequest {
	s.Step = &v
	return s
}

type QueryMonitorResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonitorResponseBody) SetData(v string) *QueryMonitorResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMonitorResponseBody) SetErrorCode(v string) *QueryMonitorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMonitorResponseBody) SetMessage(v string) *QueryMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMonitorResponseBody) SetRequestId(v string) *QueryMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonitorResponseBody) SetSuccess(v bool) *QueryMonitorResponseBody {
	s.Success = &v
	return s
}

type QueryMonitorResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonitorResponse) GoString() string {
	return s.String()
}

func (s *QueryMonitorResponse) SetHeaders(v map[string]*string) *QueryMonitorResponse {
	s.Headers = v
	return s
}

func (s *QueryMonitorResponse) SetBody(v *QueryMonitorResponseBody) *QueryMonitorResponse {
	s.Body = v
	return s
}

type QuerySlbSpecResponseBody struct {
	Code           *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*QuerySlbSpecResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySlbSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySlbSpecResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponseBody) SetCode(v int32) *QuerySlbSpecResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetData(v []*QuerySlbSpecResponseBodyData) *QuerySlbSpecResponseBody {
	s.Data = v
	return s
}

func (s *QuerySlbSpecResponseBody) SetHttpStatusCode(v int32) *QuerySlbSpecResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetMessage(v string) *QuerySlbSpecResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetRequestId(v string) *QuerySlbSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySlbSpecResponseBody) SetSuccess(v bool) *QuerySlbSpecResponseBody {
	s.Success = &v
	return s
}

type QuerySlbSpecResponseBodyData struct {
	Id                     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxConnection          *string `json:"MaxConnection,omitempty" xml:"MaxConnection,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewConnectionPerSecond *string `json:"NewConnectionPerSecond,omitempty" xml:"NewConnectionPerSecond,omitempty"`
	Qps                    *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	Spec                   *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s QuerySlbSpecResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySlbSpecResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponseBodyData) SetId(v int32) *QuerySlbSpecResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetMaxConnection(v string) *QuerySlbSpecResponseBodyData {
	s.MaxConnection = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetName(v string) *QuerySlbSpecResponseBodyData {
	s.Name = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetNewConnectionPerSecond(v string) *QuerySlbSpecResponseBodyData {
	s.NewConnectionPerSecond = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetQps(v string) *QuerySlbSpecResponseBodyData {
	s.Qps = &v
	return s
}

func (s *QuerySlbSpecResponseBodyData) SetSpec(v string) *QuerySlbSpecResponseBodyData {
	s.Spec = &v
	return s
}

type QuerySlbSpecResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySlbSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySlbSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySlbSpecResponse) GoString() string {
	return s.String()
}

func (s *QuerySlbSpecResponse) SetHeaders(v map[string]*string) *QuerySlbSpecResponse {
	s.Headers = v
	return s
}

func (s *QuerySlbSpecResponse) SetBody(v *QuerySlbSpecResponseBody) *QuerySlbSpecResponse {
	s.Body = v
	return s
}

type QueryZnodeDetailRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s QueryZnodeDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryZnodeDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailRequest) SetClusterId(v string) *QueryZnodeDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetPath(v string) *QueryZnodeDetailRequest {
	s.Path = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetRequestPars(v string) *QueryZnodeDetailRequest {
	s.RequestPars = &v
	return s
}

type QueryZnodeDetailResponseBody struct {
	Data      *QueryZnodeDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryZnodeDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryZnodeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponseBody) SetData(v *QueryZnodeDetailResponseBodyData) *QueryZnodeDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetErrorCode(v string) *QueryZnodeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetMessage(v string) *QueryZnodeDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetRequestId(v string) *QueryZnodeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryZnodeDetailResponseBody) SetSuccess(v string) *QueryZnodeDetailResponseBody {
	s.Success = &v
	return s
}

type QueryZnodeDetailResponseBodyData struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Dir  *bool   `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s QueryZnodeDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryZnodeDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponseBodyData) SetData(v string) *QueryZnodeDetailResponseBodyData {
	s.Data = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetDir(v bool) *QueryZnodeDetailResponseBodyData {
	s.Dir = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetName(v string) *QueryZnodeDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryZnodeDetailResponseBodyData) SetPath(v string) *QueryZnodeDetailResponseBodyData {
	s.Path = &v
	return s
}

type QueryZnodeDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryZnodeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryZnodeDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryZnodeDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponse) SetHeaders(v map[string]*string) *QueryZnodeDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryZnodeDetailResponse) SetBody(v *QueryZnodeDetailResponseBody) *QueryZnodeDetailResponse {
	s.Body = v
	return s
}

type RestartClusterRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PodNameList *string `json:"PodNameList,omitempty" xml:"PodNameList,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s RestartClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartClusterRequest) SetClusterId(v string) *RestartClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RestartClusterRequest) SetInstanceId(v string) *RestartClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartClusterRequest) SetPodNameList(v string) *RestartClusterRequest {
	s.PodNameList = &v
	return s
}

func (s *RestartClusterRequest) SetRequestPars(v string) *RestartClusterRequest {
	s.RequestPars = &v
	return s
}

type RestartClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartClusterResponseBody) SetErrorCode(v string) *RestartClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartClusterResponseBody) SetMessage(v string) *RestartClusterResponseBody {
	s.Message = &v
	return s
}

func (s *RestartClusterResponseBody) SetRequestId(v string) *RestartClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartClusterResponseBody) SetSuccess(v bool) *RestartClusterResponseBody {
	s.Success = &v
	return s
}

type RestartClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartClusterResponse) GoString() string {
	return s.String()
}

func (s *RestartClusterResponse) SetHeaders(v map[string]*string) *RestartClusterResponse {
	s.Headers = v
	return s
}

func (s *RestartClusterResponse) SetBody(v *RestartClusterResponseBody) *RestartClusterResponse {
	s.Body = v
	return s
}

type RetryClusterRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s RetryClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RetryClusterRequest) GoString() string {
	return s.String()
}

func (s *RetryClusterRequest) SetInstanceId(v string) *RetryClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *RetryClusterRequest) SetRequestPars(v string) *RetryClusterRequest {
	s.RequestPars = &v
	return s
}

type RetryClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetryClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RetryClusterResponseBody) SetErrorCode(v string) *RetryClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryClusterResponseBody) SetMessage(v string) *RetryClusterResponseBody {
	s.Message = &v
	return s
}

func (s *RetryClusterResponseBody) SetRequestId(v string) *RetryClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryClusterResponseBody) SetSuccess(v bool) *RetryClusterResponseBody {
	s.Success = &v
	return s
}

type RetryClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetryClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetryClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RetryClusterResponse) GoString() string {
	return s.String()
}

func (s *RetryClusterResponse) SetHeaders(v map[string]*string) *RetryClusterResponse {
	s.Headers = v
	return s
}

func (s *RetryClusterResponse) SetBody(v *RetryClusterResponseBody) *RetryClusterResponse {
	s.Body = v
	return s
}

type ScalingClusterRequest struct {
	ClusterSpecification *string `json:"ClusterSpecification,omitempty" xml:"ClusterSpecification,omitempty"`
	Cpu                  *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	InstanceCount        *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemoryCapacity       *int64  `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
}

func (s ScalingClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ScalingClusterRequest) GoString() string {
	return s.String()
}

func (s *ScalingClusterRequest) SetClusterSpecification(v string) *ScalingClusterRequest {
	s.ClusterSpecification = &v
	return s
}

func (s *ScalingClusterRequest) SetCpu(v int32) *ScalingClusterRequest {
	s.Cpu = &v
	return s
}

func (s *ScalingClusterRequest) SetInstanceCount(v int32) *ScalingClusterRequest {
	s.InstanceCount = &v
	return s
}

func (s *ScalingClusterRequest) SetInstanceId(v string) *ScalingClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *ScalingClusterRequest) SetMemoryCapacity(v int64) *ScalingClusterRequest {
	s.MemoryCapacity = &v
	return s
}

type ScalingClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScalingClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScalingClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ScalingClusterResponseBody) SetErrorCode(v string) *ScalingClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ScalingClusterResponseBody) SetMessage(v string) *ScalingClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ScalingClusterResponseBody) SetRequestId(v string) *ScalingClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScalingClusterResponseBody) SetSuccess(v bool) *ScalingClusterResponseBody {
	s.Success = &v
	return s
}

type ScalingClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ScalingClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ScalingClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ScalingClusterResponse) GoString() string {
	return s.String()
}

func (s *ScalingClusterResponse) SetHeaders(v map[string]*string) *ScalingClusterResponse {
	s.Headers = v
	return s
}

func (s *ScalingClusterResponse) SetBody(v *ScalingClusterResponseBody) *ScalingClusterResponse {
	s.Body = v
	return s
}

type UpdateAclRequest struct {
	AclEntryList *string `json:"AclEntryList,omitempty" xml:"AclEntryList,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclRequest) SetAclEntryList(v string) *UpdateAclRequest {
	s.AclEntryList = &v
	return s
}

func (s *UpdateAclRequest) SetInstanceId(v string) *UpdateAclRequest {
	s.InstanceId = &v
	return s
}

type UpdateAclResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclResponseBody) SetErrorCode(v string) *UpdateAclResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAclResponseBody) SetMessage(v string) *UpdateAclResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAclResponseBody) SetRequestId(v string) *UpdateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAclResponseBody) SetSuccess(v bool) *UpdateAclResponseBody {
	s.Success = &v
	return s
}

type UpdateAclResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAclResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclResponse) SetHeaders(v map[string]*string) *UpdateAclResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclResponse) SetBody(v *UpdateAclResponseBody) *UpdateAclResponse {
	s.Body = v
	return s
}

type UpdateClusterRequest struct {
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestPars      *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) SetClusterAliasName(v string) *UpdateClusterRequest {
	s.ClusterAliasName = &v
	return s
}

func (s *UpdateClusterRequest) SetInstanceId(v string) *UpdateClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateClusterRequest) SetRequestPars(v string) *UpdateClusterRequest {
	s.RequestPars = &v
	return s
}

type UpdateClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) SetErrorCode(v string) *UpdateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateClusterResponseBody) SetMessage(v string) *UpdateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

type UpdateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponse) SetHeaders(v map[string]*string) *UpdateClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterResponse) SetBody(v *UpdateClusterResponseBody) *UpdateClusterResponse {
	s.Body = v
	return s
}

type UpdateConfigRequest struct {
	AutopurgePurgeInterval   *string `json:"AutopurgePurgeInterval,omitempty" xml:"AutopurgePurgeInterval,omitempty"`
	AutopurgeSnapRetainCount *string `json:"AutopurgeSnapRetainCount,omitempty" xml:"AutopurgeSnapRetainCount,omitempty"`
	ClusterId                *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ConfigAuthEnabled        *bool   `json:"ConfigAuthEnabled,omitempty" xml:"ConfigAuthEnabled,omitempty"`
	ConfigSecretEnabled      *bool   `json:"ConfigSecretEnabled,omitempty" xml:"ConfigSecretEnabled,omitempty"`
	ConfigType               *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	InitLimit                *string `json:"InitLimit,omitempty" xml:"InitLimit,omitempty"`
	InstanceId               *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JuteMaxbuffer            *string `json:"JuteMaxbuffer,omitempty" xml:"JuteMaxbuffer,omitempty"`
	MCPEnabled               *bool   `json:"MCPEnabled,omitempty" xml:"MCPEnabled,omitempty"`
	MaxClientCnxns           *string `json:"MaxClientCnxns,omitempty" xml:"MaxClientCnxns,omitempty"`
	OpenSuperAcl             *string `json:"OpenSuperAcl,omitempty" xml:"OpenSuperAcl,omitempty"`
	PassWord                 *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	RequestPars              *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	SyncLimit                *string `json:"SyncLimit,omitempty" xml:"SyncLimit,omitempty"`
	TickTime                 *string `json:"TickTime,omitempty" xml:"TickTime,omitempty"`
	UserName                 *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) SetAutopurgePurgeInterval(v string) *UpdateConfigRequest {
	s.AutopurgePurgeInterval = &v
	return s
}

func (s *UpdateConfigRequest) SetAutopurgeSnapRetainCount(v string) *UpdateConfigRequest {
	s.AutopurgeSnapRetainCount = &v
	return s
}

func (s *UpdateConfigRequest) SetClusterId(v string) *UpdateConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigAuthEnabled(v bool) *UpdateConfigRequest {
	s.ConfigAuthEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigSecretEnabled(v bool) *UpdateConfigRequest {
	s.ConfigSecretEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetConfigType(v string) *UpdateConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateConfigRequest) SetInitLimit(v string) *UpdateConfigRequest {
	s.InitLimit = &v
	return s
}

func (s *UpdateConfigRequest) SetInstanceId(v string) *UpdateConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigRequest) SetJuteMaxbuffer(v string) *UpdateConfigRequest {
	s.JuteMaxbuffer = &v
	return s
}

func (s *UpdateConfigRequest) SetMCPEnabled(v bool) *UpdateConfigRequest {
	s.MCPEnabled = &v
	return s
}

func (s *UpdateConfigRequest) SetMaxClientCnxns(v string) *UpdateConfigRequest {
	s.MaxClientCnxns = &v
	return s
}

func (s *UpdateConfigRequest) SetOpenSuperAcl(v string) *UpdateConfigRequest {
	s.OpenSuperAcl = &v
	return s
}

func (s *UpdateConfigRequest) SetPassWord(v string) *UpdateConfigRequest {
	s.PassWord = &v
	return s
}

func (s *UpdateConfigRequest) SetRequestPars(v string) *UpdateConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateConfigRequest) SetSyncLimit(v string) *UpdateConfigRequest {
	s.SyncLimit = &v
	return s
}

func (s *UpdateConfigRequest) SetTickTime(v string) *UpdateConfigRequest {
	s.TickTime = &v
	return s
}

func (s *UpdateConfigRequest) SetUserName(v string) *UpdateConfigRequest {
	s.UserName = &v
	return s
}

type UpdateConfigResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponseBody) SetCode(v int32) *UpdateConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigResponseBody) SetHttpStatusCode(v int32) *UpdateConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConfigResponseBody) SetMessage(v string) *UpdateConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigResponseBody) SetRequestId(v string) *UpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigResponseBody) SetSuccess(v bool) *UpdateConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateConfigResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponse) SetHeaders(v map[string]*string) *UpdateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigResponse) SetBody(v *UpdateConfigResponseBody) *UpdateConfigResponse {
	s.Body = v
	return s
}

type UpdateEngineNamespaceRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Desc         *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceCount *int32  `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
}

func (s UpdateEngineNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceRequest) SetClusterId(v string) *UpdateEngineNamespaceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetDesc(v string) *UpdateEngineNamespaceRequest {
	s.Desc = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetId(v string) *UpdateEngineNamespaceRequest {
	s.Id = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetInstanceId(v string) *UpdateEngineNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetName(v string) *UpdateEngineNamespaceRequest {
	s.Name = &v
	return s
}

func (s *UpdateEngineNamespaceRequest) SetServiceCount(v int32) *UpdateEngineNamespaceRequest {
	s.ServiceCount = &v
	return s
}

type UpdateEngineNamespaceResponseBody struct {
	Data      *UpdateEngineNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEngineNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponseBody) SetData(v *UpdateEngineNamespaceResponseBodyData) *UpdateEngineNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetErrorCode(v string) *UpdateEngineNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetMessage(v string) *UpdateEngineNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetRequestId(v string) *UpdateEngineNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBody) SetSuccess(v bool) *UpdateEngineNamespaceResponseBody {
	s.Success = &v
	return s
}

type UpdateEngineNamespaceResponseBodyData struct {
	ConfigCount       *int32  `json:"ConfigCount,omitempty" xml:"ConfigCount,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	NamespaceDesc     *string `json:"NamespaceDesc,omitempty" xml:"NamespaceDesc,omitempty"`
	NamespaceShowName *string `json:"NamespaceShowName,omitempty" xml:"NamespaceShowName,omitempty"`
	Quota             *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateEngineNamespaceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponseBodyData) SetConfigCount(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.ConfigCount = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespace(v string) *UpdateEngineNamespaceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespaceDesc(v string) *UpdateEngineNamespaceResponseBodyData {
	s.NamespaceDesc = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetNamespaceShowName(v string) *UpdateEngineNamespaceResponseBodyData {
	s.NamespaceShowName = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetQuota(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.Quota = &v
	return s
}

func (s *UpdateEngineNamespaceResponseBodyData) SetType(v int32) *UpdateEngineNamespaceResponseBodyData {
	s.Type = &v
	return s
}

type UpdateEngineNamespaceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEngineNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEngineNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEngineNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateEngineNamespaceResponse) SetHeaders(v map[string]*string) *UpdateEngineNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateEngineNamespaceResponse) SetBody(v *UpdateEngineNamespaceResponseBody) *UpdateEngineNamespaceResponse {
	s.Body = v
	return s
}

type UpdateGatewayNameRequest struct {
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateGatewayNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameRequest) SetGatewayUniqueId(v string) *UpdateGatewayNameRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayNameRequest) SetName(v string) *UpdateGatewayNameRequest {
	s.Name = &v
	return s
}

type UpdateGatewayNameResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponseBody) SetCode(v int32) *UpdateGatewayNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetData(v string) *UpdateGatewayNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetMessage(v string) *UpdateGatewayNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetRequestId(v string) *UpdateGatewayNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetSuccess(v bool) *UpdateGatewayNameResponseBody {
	s.Success = &v
	return s
}

type UpdateGatewayNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGatewayNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponse) SetHeaders(v map[string]*string) *UpdateGatewayNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayNameResponse) SetBody(v *UpdateGatewayNameResponseBody) *UpdateGatewayNameResponse {
	s.Body = v
	return s
}

type UpdateGatewayOptionRequest struct {
	GatewayId     *int64         `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayOption *GatewayOption `json:"GatewayOption,omitempty" xml:"GatewayOption,omitempty"`
}

func (s UpdateGatewayOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionRequest) SetGatewayId(v int64) *UpdateGatewayOptionRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayOptionRequest) SetGatewayOption(v *GatewayOption) *UpdateGatewayOptionRequest {
	s.GatewayOption = v
	return s
}

type UpdateGatewayOptionShrinkRequest struct {
	GatewayId           *int64  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayOptionShrink *string `json:"GatewayOption,omitempty" xml:"GatewayOption,omitempty"`
}

func (s UpdateGatewayOptionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayOptionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionShrinkRequest) SetGatewayId(v int64) *UpdateGatewayOptionShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayOptionShrinkRequest) SetGatewayOptionShrink(v string) *UpdateGatewayOptionShrinkRequest {
	s.GatewayOptionShrink = &v
	return s
}

type UpdateGatewayOptionResponseBody struct {
	Code           *int32         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GatewayOption `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string        `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionResponseBody) SetCode(v int32) *UpdateGatewayOptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetData(v *GatewayOption) *UpdateGatewayOptionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayOptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetMessage(v string) *UpdateGatewayOptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetRequestId(v string) *UpdateGatewayOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayOptionResponseBody) SetSuccess(v bool) *UpdateGatewayOptionResponseBody {
	s.Success = &v
	return s
}

type UpdateGatewayOptionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGatewayOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayOptionResponse) SetHeaders(v map[string]*string) *UpdateGatewayOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayOptionResponse) SetBody(v *UpdateGatewayOptionResponseBody) *UpdateGatewayOptionResponse {
	s.Body = v
	return s
}

type UpdateGatewayRouteHTTPRewriteRequest struct {
	GatewayId       *int64  `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	HttpRewriteJSON *string `json:"HttpRewriteJSON,omitempty" xml:"HttpRewriteJSON,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteHTTPRewriteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetGatewayId(v int64) *UpdateGatewayRouteHTTPRewriteRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetHttpRewriteJSON(v string) *UpdateGatewayRouteHTTPRewriteRequest {
	s.HttpRewriteJSON = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteRequest) SetId(v int64) *UpdateGatewayRouteHTTPRewriteRequest {
	s.Id = &v
	return s
}

type UpdateGatewayRouteHTTPRewriteResponseBody struct {
	Code           *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayRouteHTTPRewriteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetData(v int64) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetMessage(v string) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetRequestId(v string) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponseBody) SetSuccess(v bool) *UpdateGatewayRouteHTTPRewriteResponseBody {
	s.Success = &v
	return s
}

type UpdateGatewayRouteHTTPRewriteResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateGatewayRouteHTTPRewriteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGatewayRouteHTTPRewriteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGatewayRouteHTTPRewriteResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) SetHeaders(v map[string]*string) *UpdateGatewayRouteHTTPRewriteResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayRouteHTTPRewriteResponse) SetBody(v *UpdateGatewayRouteHTTPRewriteResponseBody) *UpdateGatewayRouteHTTPRewriteResponse {
	s.Body = v
	return s
}

type UpdateImageRequest struct {
	// 目标集群的id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 想修改的镜像版本code
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) SetClusterId(v string) *UpdateImageRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateImageRequest) SetVersionCode(v string) *UpdateImageRequest {
	s.VersionCode = &v
	return s
}

type UpdateImageResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetErrorCode(v string) *UpdateImageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateImageResponseBody) SetMessage(v string) *UpdateImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

type UpdateImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageResponse) SetHeaders(v map[string]*string) *UpdateImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
	s.Body = v
	return s
}

type UpdateNacosConfigRequest struct {
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BetaIps          *string `json:"BetaIps,omitempty" xml:"BetaIps,omitempty"`
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Desc             *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	Group            *string `json:"Group,omitempty" xml:"Group,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	NamespaceId      *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateNacosConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigRequest) SetAppName(v string) *UpdateNacosConfigRequest {
	s.AppName = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetBetaIps(v string) *UpdateNacosConfigRequest {
	s.BetaIps = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetContent(v string) *UpdateNacosConfigRequest {
	s.Content = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetDataId(v string) *UpdateNacosConfigRequest {
	s.DataId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetDesc(v string) *UpdateNacosConfigRequest {
	s.Desc = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetEncryptedDataKey(v string) *UpdateNacosConfigRequest {
	s.EncryptedDataKey = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetGroup(v string) *UpdateNacosConfigRequest {
	s.Group = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetInstanceId(v string) *UpdateNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetMd5(v string) *UpdateNacosConfigRequest {
	s.Md5 = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetNamespaceId(v string) *UpdateNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetTags(v string) *UpdateNacosConfigRequest {
	s.Tags = &v
	return s
}

func (s *UpdateNacosConfigRequest) SetType(v string) *UpdateNacosConfigRequest {
	s.Type = &v
	return s
}

type UpdateNacosConfigResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNacosConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigResponseBody) SetErrorCode(v string) *UpdateNacosConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetHttpCode(v string) *UpdateNacosConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetMessage(v string) *UpdateNacosConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetRequestId(v string) *UpdateNacosConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosConfigResponseBody) SetSuccess(v bool) *UpdateNacosConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateNacosConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateNacosConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNacosConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosConfigResponse) SetHeaders(v map[string]*string) *UpdateNacosConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosConfigResponse) SetBody(v *UpdateNacosConfigResponseBody) *UpdateNacosConfigResponse {
	s.Body = v
	return s
}

type UpdateNacosInstanceRequest struct {
	// Nacos集群名
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 服务禁用标志
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 临时节点标志
	Ephemeral *bool `json:"Ephemeral,omitempty" xml:"Ephemeral,omitempty"`
	// 分组名
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Nacos实例ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 节点元数据
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// 命名空间id
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Nacos实例端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// 权重
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateNacosInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceRequest) SetClusterName(v string) *UpdateNacosInstanceRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetEnabled(v bool) *UpdateNacosInstanceRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetEphemeral(v bool) *UpdateNacosInstanceRequest {
	s.Ephemeral = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetGroupName(v string) *UpdateNacosInstanceRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetInstanceId(v string) *UpdateNacosInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetIp(v string) *UpdateNacosInstanceRequest {
	s.Ip = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetMetadata(v string) *UpdateNacosInstanceRequest {
	s.Metadata = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetNamespaceId(v string) *UpdateNacosInstanceRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetPort(v int32) *UpdateNacosInstanceRequest {
	s.Port = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetServiceName(v string) *UpdateNacosInstanceRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateNacosInstanceRequest) SetWeight(v string) *UpdateNacosInstanceRequest {
	s.Weight = &v
	return s
}

type UpdateNacosInstanceResponseBody struct {
	// 响应码
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 修改结果
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// http状态码
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 响应信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 成功标志
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNacosInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceResponseBody) SetCode(v int32) *UpdateNacosInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetData(v string) *UpdateNacosInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateNacosInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetMessage(v string) *UpdateNacosInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetRequestId(v string) *UpdateNacosInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacosInstanceResponseBody) SetSuccess(v bool) *UpdateNacosInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateNacosInstanceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateNacosInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateNacosInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacosInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosInstanceResponse) SetHeaders(v map[string]*string) *UpdateNacosInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosInstanceResponse) SetBody(v *UpdateNacosInstanceResponseBody) *UpdateNacosInstanceResponse {
	s.Body = v
	return s
}

type UpdateZnodeRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s UpdateZnodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateZnodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateZnodeRequest) SetClusterId(v string) *UpdateZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateZnodeRequest) SetData(v string) *UpdateZnodeRequest {
	s.Data = &v
	return s
}

func (s *UpdateZnodeRequest) SetPath(v string) *UpdateZnodeRequest {
	s.Path = &v
	return s
}

func (s *UpdateZnodeRequest) SetRequestPars(v string) *UpdateZnodeRequest {
	s.RequestPars = &v
	return s
}

type UpdateZnodeResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateZnodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateZnodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZnodeResponseBody) SetErrorCode(v string) *UpdateZnodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetMessage(v string) *UpdateZnodeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetRequestId(v string) *UpdateZnodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZnodeResponseBody) SetSuccess(v bool) *UpdateZnodeResponseBody {
	s.Success = &v
	return s
}

type UpdateZnodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateZnodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateZnodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateZnodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateZnodeResponse) SetHeaders(v map[string]*string) *UpdateZnodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateZnodeResponse) SetBody(v *UpdateZnodeResponseBody) *UpdateZnodeResponse {
	s.Body = v
	return s
}

type UpgradeClusterRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestPars    *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
}

func (s UpgradeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequest) SetInstanceId(v string) *UpgradeClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeClusterRequest) SetRequestPars(v string) *UpgradeClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *UpgradeClusterRequest) SetUpgradeVersion(v string) *UpgradeClusterRequest {
	s.UpgradeVersion = &v
	return s
}

type UpgradeClusterResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpCode  *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponseBody) SetErrorCode(v string) *UpgradeClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetHttpCode(v string) *UpgradeClusterResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetMessage(v string) *UpgradeClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetRequestId(v string) *UpgradeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeClusterResponseBody) SetSuccess(v bool) *UpgradeClusterResponseBody {
	s.Success = &v
	return s
}

type UpgradeClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClusterResponse) SetHeaders(v map[string]*string) *UpgradeClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClusterResponse) SetBody(v *UpgradeClusterResponseBody) *UpgradeClusterResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mse"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddMockRuleWithOptions(request *AddMockRuleRequest, runtime *util.RuntimeOptions) (_result *AddMockRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ConsumerAppIds"] = request.ConsumerAppIds
	query["DubboMockItems"] = request.DubboMockItems
	query["Enable"] = request.Enable
	query["ExtraJson"] = request.ExtraJson
	query["MockType"] = request.MockType
	query["Name"] = request.Name
	query["ProviderAppId"] = request.ProviderAppId
	query["ProviderAppName"] = request.ProviderAppName
	query["Region"] = request.Region
	query["ScMockItems"] = request.ScMockItems
	query["Source"] = request.Source
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMockRule"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMockRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMockRule(request *AddMockRuleRequest) (_result *AddMockRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMockRuleResponse{}
	_body, _err := client.AddMockRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddServiceSourceWithOptions(request *AddServiceSourceRequest, runtime *util.RuntimeOptions) (_result *AddServiceSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Address"] = request.Address
	query["GatewayId"] = request.GatewayId
	query["GatewayUniqueId"] = request.GatewayUniqueId
	query["Info1"] = request.Info1
	query["Info2"] = request.Info2
	query["Name"] = request.Name
	query["Source"] = request.Source
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServiceSource"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddServiceSource(request *AddServiceSourceRequest) (_result *AddServiceSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServiceSourceResponse{}
	_body, _err := client.AddServiceSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneNacosConfigWithOptions(request *CloneNacosConfigRequest, runtime *util.RuntimeOptions) (_result *CloneNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Ids"] = request.Ids
	query["InstanceId"] = request.InstanceId
	query["OriginNamespaceId"] = request.OriginNamespaceId
	query["Policy"] = request.Policy
	query["TargetNamespaceId"] = request.TargetNamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneNacosConfig(request *CloneNacosConfigRequest) (_result *CloneNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneNacosConfigResponse{}
	_body, _err := client.CloneNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlarmRuleWithOptions(tmpReq *CreateAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlertWay)) {
		request.AlertWayShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertWay, tea.String("AlertWay"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ContactGroupIds)) {
		request.ContactGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactGroupIds, tea.String("ContactGroupIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["Aggregates"] = request.Aggregates
	query["AlarmAliasName"] = request.AlarmAliasName
	query["AlarmItem"] = request.AlarmItem
	query["AlertWay"] = request.AlertWayShrink
	query["ContactGroupIds"] = request.ContactGroupIdsShrink
	query["InstanceId"] = request.InstanceId
	query["NValue"] = request.NValue
	query["Operator"] = request.Operator
	query["Value"] = request.Value
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlarmRule"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlarmRule(request *CreateAlarmRuleRequest) (_result *CreateAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlarmRuleResponse{}
	_body, _err := client.CreateAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["ExtraInfo"] = request.ExtraInfo
	query["Language"] = request.Language
	query["Region"] = request.Region
	query["Source"] = request.Source
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterSpecification"] = request.ClusterSpecification
	query["ClusterType"] = request.ClusterType
	query["ClusterVersion"] = request.ClusterVersion
	query["ConnectionType"] = request.ConnectionType
	query["DiskCapacity"] = request.DiskCapacity
	query["DiskType"] = request.DiskType
	query["InstanceCount"] = request.InstanceCount
	query["MseVersion"] = request.MseVersion
	query["NetType"] = request.NetType
	query["PrivateSlbSpecification"] = request.PrivateSlbSpecification
	query["PubNetworkFlow"] = request.PubNetworkFlow
	query["PubSlbSpecification"] = request.PubSlbSpecification
	query["Region"] = request.Region
	query["RequestPars"] = request.RequestPars
	query["VSwitchId"] = request.VSwitchId
	query["VpcId"] = request.VpcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEngineNamespaceWithOptions(request *CreateEngineNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateEngineNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Desc"] = request.Desc
	query["InstanceId"] = request.InstanceId
	query["Name"] = request.Name
	query["ServiceCount"] = request.ServiceCount
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEngineNamespace"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEngineNamespace(request *CreateEngineNamespaceRequest) (_result *CreateEngineNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEngineNamespaceResponse{}
	_body, _err := client.CreateEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGovernanceKubernetesClusterWithOptions(request *CreateGovernanceKubernetesClusterRequest, runtime *util.RuntimeOptions) (_result *CreateGovernanceKubernetesClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ClusterName"] = request.ClusterName
	query["K8sVersion"] = request.K8sVersion
	query["NameSpaceInfos"] = request.NameSpaceInfos
	query["PilotStartTime"] = request.PilotStartTime
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGovernanceKubernetesCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGovernanceKubernetesCluster(request *CreateGovernanceKubernetesClusterRequest) (_result *CreateGovernanceKubernetesClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGovernanceKubernetesClusterResponse{}
	_body, _err := client.CreateGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateNacosConfigWithOptions(request *CreateNacosConfigRequest, runtime *util.RuntimeOptions) (_result *CreateNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["BetaIps"] = request.BetaIps
	query["Content"] = request.Content
	query["DataId"] = request.DataId
	query["Desc"] = request.Desc
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["Tags"] = request.Tags
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateNacosConfig(request *CreateNacosConfigRequest) (_result *CreateNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNacosConfigResponse{}
	_body, _err := client.CreateNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateZnodeWithOptions(request *CreateZnodeRequest, runtime *util.RuntimeOptions) (_result *CreateZnodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Data"] = request.Data
	query["Path"] = request.Path
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateZnode"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateZnode(request *CreateZnodeRequest) (_result *CreateZnodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateZnodeResponse{}
	_body, _err := client.CreateZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlarmRuleWithOptions(request *DeleteAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AlarmRuleId"] = request.AlarmRuleId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlarmRule"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlarmRule(request *DeleteAlarmRuleRequest) (_result *DeleteAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlarmRuleResponse{}
	_body, _err := client.DeleteAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEngineNamespaceWithOptions(request *DeleteEngineNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteEngineNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Id"] = request.Id
	query["InstanceId"] = request.InstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEngineNamespace"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEngineNamespace(request *DeleteEngineNamespaceRequest) (_result *DeleteEngineNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEngineNamespaceResponse{}
	_body, _err := client.DeleteEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNacosConfigWithOptions(request *DeleteNacosConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Beta"] = request.Beta
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNacosConfig(request *DeleteNacosConfigRequest) (_result *DeleteNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNacosConfigResponse{}
	_body, _err := client.DeleteNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNacosConfigsWithOptions(request *DeleteNacosConfigsRequest, runtime *util.RuntimeOptions) (_result *DeleteNacosConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Ids"] = request.Ids
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNacosConfigs"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNacosConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNacosConfigs(request *DeleteNacosConfigsRequest) (_result *DeleteNacosConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNacosConfigsResponse{}
	_body, _err := client.DeleteNacosConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNacosServiceWithOptions(request *DeleteNacosServiceRequest, runtime *util.RuntimeOptions) (_result *DeleteNacosServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupName"] = request.GroupName
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["ServiceName"] = request.ServiceName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNacosService"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNacosServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNacosService(request *DeleteNacosServiceRequest) (_result *DeleteNacosServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNacosServiceResponse{}
	_body, _err := client.DeleteNacosServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteZnodeWithOptions(request *DeleteZnodeRequest, runtime *util.RuntimeOptions) (_result *DeleteZnodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Path"] = request.Path
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteZnode"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteZnode(request *DeleteZnodeRequest) (_result *DeleteZnodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteZnodeResponse{}
	_body, _err := client.DeleteZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportNacosConfigWithOptions(request *ExportNacosConfigRequest, runtime *util.RuntimeOptions) (_result *ExportNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["Ids"] = request.Ids
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportNacosConfig(request *ExportNacosConfigRequest) (_result *ExportNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportNacosConfigResponse{}
	_body, _err := client.ExportNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEngineNamepaceWithOptions(request *GetEngineNamepaceRequest, runtime *util.RuntimeOptions) (_result *GetEngineNamepaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Id"] = request.Id
	query["InstanceId"] = request.InstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEngineNamepace"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEngineNamepaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEngineNamepace(request *GetEngineNamepaceRequest) (_result *GetEngineNamepaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEngineNamepaceResponse{}
	_body, _err := client.GetEngineNamepaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayWithOptions(request *GetGatewayRequest, runtime *util.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGateway"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGateway(request *GetGatewayRequest) (_result *GetGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayResponse{}
	_body, _err := client.GetGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGatewayOptionWithOptions(request *GetGatewayOptionRequest, runtime *util.RuntimeOptions) (_result *GetGatewayOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GatewayId"] = request.GatewayId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGatewayOption"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGatewayOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGatewayOption(request *GetGatewayOptionRequest) (_result *GetGatewayOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGatewayOptionResponse{}
	_body, _err := client.GetGatewayOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGovernanceKubernetesClusterWithOptions(request *GetGovernanceKubernetesClusterRequest, runtime *util.RuntimeOptions) (_result *GetGovernanceKubernetesClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGovernanceKubernetesCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGovernanceKubernetesCluster(request *GetGovernanceKubernetesClusterRequest) (_result *GetGovernanceKubernetesClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGovernanceKubernetesClusterResponse{}
	_body, _err := client.GetGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGovernanceKubernetesClusterListWithOptions(request *GetGovernanceKubernetesClusterListRequest, runtime *util.RuntimeOptions) (_result *GetGovernanceKubernetesClusterListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["ClusterName"] = request.ClusterName
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGovernanceKubernetesClusterList"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGovernanceKubernetesClusterListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGovernanceKubernetesClusterList(request *GetGovernanceKubernetesClusterListRequest) (_result *GetGovernanceKubernetesClusterListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGovernanceKubernetesClusterListResponse{}
	_body, _err := client.GetGovernanceKubernetesClusterListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["VersionCode"] = request.VersionCode
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImportFileUrlWithOptions(request *GetImportFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetImportFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ContentType"] = request.ContentType
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImportFileUrl"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImportFileUrl(request *GetImportFileUrlRequest) (_result *GetImportFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImportFileUrlResponse{}
	_body, _err := client.GetImportFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMseFeatureSwitchWithOptions(runtime *util.RuntimeOptions) (_result *GetMseFeatureSwitchResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetMseFeatureSwitch"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMseFeatureSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMseFeatureSwitch() (_result *GetMseFeatureSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMseFeatureSwitchResponse{}
	_body, _err := client.GetMseFeatureSwitchWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNacosConfigWithOptions(request *GetNacosConfigRequest, runtime *util.RuntimeOptions) (_result *GetNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Beta"] = request.Beta
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNacosConfig(request *GetNacosConfigRequest) (_result *GetNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNacosConfigResponse{}
	_body, _err := client.GetNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNacosHistoryConfigWithOptions(request *GetNacosHistoryConfigRequest, runtime *util.RuntimeOptions) (_result *GetNacosHistoryConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["Nid"] = request.Nid
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNacosHistoryConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNacosHistoryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNacosHistoryConfig(request *GetNacosHistoryConfigRequest) (_result *GetNacosHistoryConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNacosHistoryConfigResponse{}
	_body, _err := client.GetNacosHistoryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOverviewWithOptions(request *GetOverviewRequest, runtime *util.RuntimeOptions) (_result *GetOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Period"] = request.Period
	query["Region"] = request.Region
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOverview"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOverview(request *GetOverviewRequest) (_result *GetOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOverviewResponse{}
	_body, _err := client.GetOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportNacosConfigWithOptions(request *ImportNacosConfigRequest, runtime *util.RuntimeOptions) (_result *ImportNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FileUrl"] = request.FileUrl
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["Policy"] = request.Policy
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportNacosConfig(request *ImportNacosConfigRequest) (_result *ImportNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportNacosConfigResponse{}
	_body, _err := client.ImportNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmContactGroupsWithOptions(request *ListAlarmContactGroupsRequest, runtime *util.RuntimeOptions) (_result *ListAlarmContactGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmContactGroups"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmContactGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmContactGroups(request *ListAlarmContactGroupsRequest) (_result *ListAlarmContactGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmContactGroupsResponse{}
	_body, _err := client.ListAlarmContactGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmHistoriesWithOptions(request *ListAlarmHistoriesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmHistories"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmHistories(request *ListAlarmHistoriesRequest) (_result *ListAlarmHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.ListAlarmHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmItemsWithOptions(request *ListAlarmItemsRequest, runtime *util.RuntimeOptions) (_result *ListAlarmItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmItems"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmItems(request *ListAlarmItemsRequest) (_result *ListAlarmItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmItemsResponse{}
	_body, _err := client.ListAlarmItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlarmRulesWithOptions(request *ListAlarmRulesRequest, runtime *util.RuntimeOptions) (_result *ListAlarmRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlarmRules"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlarmRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlarmRules(request *ListAlarmRulesRequest) (_result *ListAlarmRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlarmRulesResponse{}
	_body, _err := client.ListAlarmRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAnsInstancesWithOptions(request *ListAnsInstancesRequest, runtime *util.RuntimeOptions) (_result *ListAnsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnsInstances"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAnsInstances(request *ListAnsInstancesRequest) (_result *ListAnsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAnsInstancesResponse{}
	_body, _err := client.ListAnsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAnsServiceClustersWithOptions(request *ListAnsServiceClustersRequest, runtime *util.RuntimeOptions) (_result *ListAnsServiceClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnsServiceClusters"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnsServiceClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAnsServiceClusters(request *ListAnsServiceClustersRequest) (_result *ListAnsServiceClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAnsServiceClustersResponse{}
	_body, _err := client.ListAnsServiceClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAnsServicesWithOptions(request *ListAnsServicesRequest, runtime *util.RuntimeOptions) (_result *ListAnsServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAnsServices"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAnsServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAnsServices(request *ListAnsServicesRequest) (_result *ListAnsServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAnsServicesResponse{}
	_body, _err := client.ListAnsServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterConnectionTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListClusterConnectionTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListClusterConnectionTypes"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterConnectionTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterConnectionTypes() (_result *ListClusterConnectionTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterConnectionTypesResponse{}
	_body, _err := client.ListClusterConnectionTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterTypesWithOptions(request *ListClusterTypesRequest, runtime *util.RuntimeOptions) (_result *ListClusterTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterTypes"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterTypes(request *ListClusterTypesRequest) (_result *ListClusterTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterTypesResponse{}
	_body, _err := client.ListClusterTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterVersionsWithOptions(request *ListClusterVersionsRequest, runtime *util.RuntimeOptions) (_result *ListClusterVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterType"] = request.ClusterType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterVersions"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterVersions(request *ListClusterVersionsRequest) (_result *ListClusterVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterVersionsResponse{}
	_body, _err := client.ListClusterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEngineNamespacesWithOptions(request *ListEngineNamespacesRequest, runtime *util.RuntimeOptions) (_result *ListEngineNamespacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEngineNamespaces"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEngineNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEngineNamespaces(request *ListEngineNamespacesRequest) (_result *ListEngineNamespacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEngineNamespacesResponse{}
	_body, _err := client.ListEngineNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEurekaInstancesWithOptions(request *ListEurekaInstancesRequest, runtime *util.RuntimeOptions) (_result *ListEurekaInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEurekaInstances"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEurekaInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEurekaInstances(request *ListEurekaInstancesRequest) (_result *ListEurekaInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEurekaInstancesResponse{}
	_body, _err := client.ListEurekaInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEurekaServicesWithOptions(request *ListEurekaServicesRequest, runtime *util.RuntimeOptions) (_result *ListEurekaServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEurekaServices"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEurekaServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEurekaServices(request *ListEurekaServicesRequest) (_result *ListEurekaServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEurekaServicesResponse{}
	_body, _err := client.ListEurekaServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGatewayWithOptions(tmpReq *ListGatewayRequest, runtime *util.RuntimeOptions) (_result *ListGatewayResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGatewayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.FilterParams))) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.FilterParams), tea.String("FilterParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["DescSort"] = request.DescSort
	query["FilterParams"] = request.FilterParamsShrink
	query["OrderItem"] = request.OrderItem
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGateway"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGateway(request *ListGatewayRequest) (_result *ListGatewayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGatewayResponse{}
	_body, _err := client.ListGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersByConfigWithOptions(request *ListListenersByConfigRequest, runtime *util.RuntimeOptions) (_result *ListListenersByConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenersByConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersByConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenersByConfig(request *ListListenersByConfigRequest) (_result *ListListenersByConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersByConfigResponse{}
	_body, _err := client.ListListenersByConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersByIpWithOptions(request *ListListenersByIpRequest, runtime *util.RuntimeOptions) (_result *ListListenersByIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["Ip"] = request.Ip
	query["NamespaceId"] = request.NamespaceId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListenersByIp"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersByIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenersByIp(request *ListListenersByIpRequest) (_result *ListListenersByIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersByIpResponse{}
	_body, _err := client.ListListenersByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNacosConfigsWithOptions(request *ListNacosConfigsRequest, runtime *util.RuntimeOptions) (_result *ListNacosConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["PageNum"] = request.PageNum
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["RequestPars"] = request.RequestPars
	query["Tags"] = request.Tags
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNacosConfigs"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNacosConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNacosConfigs(request *ListNacosConfigsRequest) (_result *ListNacosConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNacosConfigsResponse{}
	_body, _err := client.ListNacosConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNacosHistoryConfigsWithOptions(request *ListNacosHistoryConfigsRequest, runtime *util.RuntimeOptions) (_result *ListNacosHistoryConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DataId"] = request.DataId
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["NamespaceId"] = request.NamespaceId
	query["PageNum"] = request.PageNum
	query["PageSize"] = request.PageSize
	query["RegionId"] = request.RegionId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNacosHistoryConfigs"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNacosHistoryConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNacosHistoryConfigs(request *ListNacosHistoryConfigsRequest) (_result *ListNacosHistoryConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNacosHistoryConfigsResponse{}
	_body, _err := client.ListNacosHistoryConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListZnodeChildrenWithOptions(request *ListZnodeChildrenRequest, runtime *util.RuntimeOptions) (_result *ListZnodeChildrenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListZnodeChildren"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListZnodeChildrenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListZnodeChildren(request *ListZnodeChildrenRequest) (_result *ListZnodeChildrenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListZnodeChildrenResponse{}
	_body, _err := client.ListZnodeChildrenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGovernanceKubernetesClusterWithOptions(request *ModifyGovernanceKubernetesClusterRequest, runtime *util.RuntimeOptions) (_result *ModifyGovernanceKubernetesClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["NamespaceInfos"] = request.NamespaceInfos
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGovernanceKubernetesCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGovernanceKubernetesCluster(request *ModifyGovernanceKubernetesClusterRequest) (_result *ModifyGovernanceKubernetesClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGovernanceKubernetesClusterResponse{}
	_body, _err := client.ModifyGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBusinessLocationsWithOptions(runtime *util.RuntimeOptions) (_result *QueryBusinessLocationsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryBusinessLocations"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBusinessLocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBusinessLocations() (_result *QueryBusinessLocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBusinessLocationsResponse{}
	_body, _err := client.QueryBusinessLocationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClusterDetailWithOptions(request *QueryClusterDetailRequest, runtime *util.RuntimeOptions) (_result *QueryClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["OrderId"] = request.OrderId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryClusterDetail"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClusterDetail(request *QueryClusterDetailRequest) (_result *QueryClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.QueryClusterDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClusterDiskSpecificationWithOptions(request *QueryClusterDiskSpecificationRequest, runtime *util.RuntimeOptions) (_result *QueryClusterDiskSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterType"] = request.ClusterType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryClusterDiskSpecification"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClusterDiskSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClusterDiskSpecification(request *QueryClusterDiskSpecificationRequest) (_result *QueryClusterDiskSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryClusterDiskSpecificationResponse{}
	_body, _err := client.QueryClusterDiskSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryClusterSpecificationWithOptions(runtime *util.RuntimeOptions) (_result *QueryClusterSpecificationResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryClusterSpecification"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryClusterSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryClusterSpecification() (_result *QueryClusterSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryClusterSpecificationResponse{}
	_body, _err := client.QueryClusterSpecificationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConfigWithOptions(request *QueryConfigRequest, runtime *util.RuntimeOptions) (_result *QueryConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConfig(request *QueryConfigRequest) (_result *QueryConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryConfigResponse{}
	_body, _err := client.QueryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGatewayRegionWithOptions(runtime *util.RuntimeOptions) (_result *QueryGatewayRegionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryGatewayRegion"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGatewayRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGatewayRegion() (_result *QueryGatewayRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGatewayRegionResponse{}
	_body, _err := client.QueryGatewayRegionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGatewayTypeWithOptions(runtime *util.RuntimeOptions) (_result *QueryGatewayTypeResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryGatewayType"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGatewayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGatewayType() (_result *QueryGatewayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGatewayTypeResponse{}
	_body, _err := client.QueryGatewayTypeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonitorWithOptions(request *QueryMonitorRequest, runtime *util.RuntimeOptions) (_result *QueryMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMonitor"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonitor(request *QueryMonitorRequest) (_result *QueryMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonitorResponse{}
	_body, _err := client.QueryMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySlbSpecWithOptions(runtime *util.RuntimeOptions) (_result *QuerySlbSpecResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QuerySlbSpec"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySlbSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySlbSpec() (_result *QuerySlbSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySlbSpecResponse{}
	_body, _err := client.QuerySlbSpecWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryZnodeDetailWithOptions(request *QueryZnodeDetailRequest, runtime *util.RuntimeOptions) (_result *QueryZnodeDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryZnodeDetail"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryZnodeDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryZnodeDetail(request *QueryZnodeDetailRequest) (_result *QueryZnodeDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryZnodeDetailResponse{}
	_body, _err := client.QueryZnodeDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartClusterWithOptions(request *RestartClusterRequest, runtime *util.RuntimeOptions) (_result *RestartClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["InstanceId"] = request.InstanceId
	query["PodNameList"] = request.PodNameList
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartCluster(request *RestartClusterRequest) (_result *RestartClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartClusterResponse{}
	_body, _err := client.RestartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetryClusterWithOptions(request *RetryClusterRequest, runtime *util.RuntimeOptions) (_result *RetryClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RetryCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetryClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetryCluster(request *RetryClusterRequest) (_result *RetryClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetryClusterResponse{}
	_body, _err := client.RetryClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScalingClusterWithOptions(request *ScalingClusterRequest, runtime *util.RuntimeOptions) (_result *ScalingClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterSpecification"] = request.ClusterSpecification
	query["Cpu"] = request.Cpu
	query["InstanceCount"] = request.InstanceCount
	query["InstanceId"] = request.InstanceId
	query["MemoryCapacity"] = request.MemoryCapacity
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ScalingCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ScalingClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScalingCluster(request *ScalingClusterRequest) (_result *ScalingClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScalingClusterResponse{}
	_body, _err := client.ScalingClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAclWithOptions(request *UpdateAclRequest, runtime *util.RuntimeOptions) (_result *UpdateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AclEntryList"] = request.AclEntryList
	query["InstanceId"] = request.InstanceId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAcl"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAcl(request *UpdateAclRequest) (_result *UpdateAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAclResponse{}
	_body, _err := client.UpdateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterWithOptions(request *UpdateClusterRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterAliasName"] = request.ClusterAliasName
	query["InstanceId"] = request.InstanceId
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigWithOptions(request *UpdateConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutopurgePurgeInterval"] = request.AutopurgePurgeInterval
	query["AutopurgeSnapRetainCount"] = request.AutopurgeSnapRetainCount
	query["ClusterId"] = request.ClusterId
	query["ConfigAuthEnabled"] = request.ConfigAuthEnabled
	query["ConfigSecretEnabled"] = request.ConfigSecretEnabled
	query["ConfigType"] = request.ConfigType
	query["InitLimit"] = request.InitLimit
	query["InstanceId"] = request.InstanceId
	query["JuteMaxbuffer"] = request.JuteMaxbuffer
	query["MCPEnabled"] = request.MCPEnabled
	query["MaxClientCnxns"] = request.MaxClientCnxns
	query["PassWord"] = request.PassWord
	query["RequestPars"] = request.RequestPars
	query["SyncLimit"] = request.SyncLimit
	query["TickTime"] = request.TickTime
	query["UserName"] = request.UserName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfig(request *UpdateConfigRequest) (_result *UpdateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigResponse{}
	_body, _err := client.UpdateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEngineNamespaceWithOptions(request *UpdateEngineNamespaceRequest, runtime *util.RuntimeOptions) (_result *UpdateEngineNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Desc"] = request.Desc
	query["Id"] = request.Id
	query["InstanceId"] = request.InstanceId
	query["Name"] = request.Name
	query["ServiceCount"] = request.ServiceCount
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEngineNamespace"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEngineNamespace(request *UpdateEngineNamespaceRequest) (_result *UpdateEngineNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEngineNamespaceResponse{}
	_body, _err := client.UpdateEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayNameWithOptions(request *UpdateGatewayNameRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayName"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayName(request *UpdateGatewayNameRequest) (_result *UpdateGatewayNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.UpdateGatewayNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayOptionWithOptions(tmpReq *UpdateGatewayOptionRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayOptionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGatewayOptionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.GatewayOption))) {
		request.GatewayOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.GatewayOption), tea.String("GatewayOption"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["GatewayId"] = request.GatewayId
	query["GatewayOption"] = request.GatewayOptionShrink
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayOption"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayOption(request *UpdateGatewayOptionRequest) (_result *UpdateGatewayOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayOptionResponse{}
	_body, _err := client.UpdateGatewayOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGatewayRouteHTTPRewriteWithOptions(request *UpdateGatewayRouteHTTPRewriteRequest, runtime *util.RuntimeOptions) (_result *UpdateGatewayRouteHTTPRewriteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GatewayId"] = request.GatewayId
	query["HttpRewriteJSON"] = request.HttpRewriteJSON
	query["Id"] = request.Id
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGatewayRouteHTTPRewrite"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGatewayRouteHTTPRewriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGatewayRouteHTTPRewrite(request *UpdateGatewayRouteHTTPRewriteRequest) (_result *UpdateGatewayRouteHTTPRewriteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGatewayRouteHTTPRewriteResponse{}
	_body, _err := client.UpdateGatewayRouteHTTPRewriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageWithOptions(request *UpdateImageRequest, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["VersionCode"] = request.VersionCode
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImage"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNacosConfigWithOptions(request *UpdateNacosConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateNacosConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppName"] = request.AppName
	query["BetaIps"] = request.BetaIps
	query["Content"] = request.Content
	query["DataId"] = request.DataId
	query["Desc"] = request.Desc
	query["EncryptedDataKey"] = request.EncryptedDataKey
	query["Group"] = request.Group
	query["InstanceId"] = request.InstanceId
	query["Md5"] = request.Md5
	query["NamespaceId"] = request.NamespaceId
	query["Tags"] = request.Tags
	query["Type"] = request.Type
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNacosConfig"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNacosConfig(request *UpdateNacosConfigRequest) (_result *UpdateNacosConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNacosConfigResponse{}
	_body, _err := client.UpdateNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNacosInstanceWithOptions(request *UpdateNacosInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateNacosInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterName"] = request.ClusterName
	query["Enabled"] = request.Enabled
	query["Ephemeral"] = request.Ephemeral
	query["GroupName"] = request.GroupName
	query["InstanceId"] = request.InstanceId
	query["Ip"] = request.Ip
	query["NamespaceId"] = request.NamespaceId
	query["Port"] = request.Port
	query["ServiceName"] = request.ServiceName
	query["Weight"] = request.Weight
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNacosInstance"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNacosInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNacosInstance(request *UpdateNacosInstanceRequest) (_result *UpdateNacosInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNacosInstanceResponse{}
	_body, _err := client.UpdateNacosInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateZnodeWithOptions(request *UpdateZnodeRequest, runtime *util.RuntimeOptions) (_result *UpdateZnodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClusterId"] = request.ClusterId
	query["Data"] = request.Data
	query["Path"] = request.Path
	query["RequestPars"] = request.RequestPars
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateZnode"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateZnode(request *UpdateZnodeRequest) (_result *UpdateZnodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateZnodeResponse{}
	_body, _err := client.UpdateZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClusterWithOptions(request *UpgradeClusterRequest, runtime *util.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["RequestPars"] = request.RequestPars
	query["UpgradeVersion"] = request.UpgradeVersion
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeCluster"),
		Version:     tea.String("2019-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeCluster(request *UpgradeClusterRequest) (_result *UpgradeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.UpgradeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
