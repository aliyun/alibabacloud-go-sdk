// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListAclsRequest struct {
	// 访问控制策略Id
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// 访问控制策略名称
	AclNames []*string `json:"AclNames,omitempty" xml:"AclNames,omitempty" type:"Repeated"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclsRequest) GoString() string {
	return s.String()
}

func (s *ListAclsRequest) SetAclIds(v []*string) *ListAclsRequest {
	s.AclIds = v
	return s
}

func (s *ListAclsRequest) SetAclNames(v []*string) *ListAclsRequest {
	s.AclNames = v
	return s
}

func (s *ListAclsRequest) SetResourceGroupId(v string) *ListAclsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAclsRequest) SetNextToken(v string) *ListAclsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclsRequest) SetMaxResults(v int32) *ListAclsRequest {
	s.MaxResults = &v
	return s
}

type ListAclsResponseBody struct {
	// 访问控制列表
	Acls []*ListAclsResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBody) SetAcls(v []*ListAclsResponseBodyAcls) *ListAclsResponseBody {
	s.Acls = v
	return s
}

func (s *ListAclsResponseBody) SetMaxResults(v int32) *ListAclsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclsResponseBody) SetNextToken(v string) *ListAclsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclsResponseBody) SetRequestId(v string) *ListAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclsResponseBody) SetTotalCount(v int32) *ListAclsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAclsResponseBodyAcls struct {
	// 访问控制策略id
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 访问控制策略名称
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// 状态
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// IP版本
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListAclsResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *ListAclsResponseBodyAcls) SetAclId(v string) *ListAclsResponseBodyAcls {
	s.AclId = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclName(v string) *ListAclsResponseBodyAcls {
	s.AclName = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAclStatus(v string) *ListAclsResponseBodyAcls {
	s.AclStatus = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetAddressIPVersion(v string) *ListAclsResponseBodyAcls {
	s.AddressIPVersion = &v
	return s
}

func (s *ListAclsResponseBodyAcls) SetResourceGroupId(v string) *ListAclsResponseBodyAcls {
	s.ResourceGroupId = &v
	return s
}

type ListAclsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclsResponse) GoString() string {
	return s.String()
}

func (s *ListAclsResponse) SetHeaders(v map[string]*string) *ListAclsResponse {
	s.Headers = v
	return s
}

func (s *ListAclsResponse) SetBody(v *ListAclsResponseBody) *ListAclsResponse {
	s.Body = v
	return s
}

type CreateHealthCheckTemplateRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 名称
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// 端口号
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 时间间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// uri
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// 协议
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 健康阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 不健康阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateHealthCheckTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateRequest) SetClientToken(v string) *CreateHealthCheckTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTemplateName(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckConnectPort(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHost(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckCodes(v []*string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckInterval(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckTimeout(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckPath(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckMethod(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckProtocol(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthCheckHttpVersion(v string) *CreateHealthCheckTemplateRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetHealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetUnhealthyThreshold(v int32) *CreateHealthCheckTemplateRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateHealthCheckTemplateRequest) SetDryRun(v bool) *CreateHealthCheckTemplateRequest {
	s.DryRun = &v
	return s
}

type CreateHealthCheckTemplateResponseBody struct {
	// 健康检查模板ID
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHealthCheckTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponseBody) SetHealthCheckTemplateId(v string) *CreateHealthCheckTemplateResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *CreateHealthCheckTemplateResponseBody) SetRequestId(v string) *CreateHealthCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateHealthCheckTemplateResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHealthCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHealthCheckTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHealthCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckTemplateResponse) SetHeaders(v map[string]*string) *CreateHealthCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateHealthCheckTemplateResponse) SetBody(v *CreateHealthCheckTemplateResponseBody) *CreateHealthCheckTemplateResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源实例Id
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// 标签列表
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
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

type ListTagResourcesResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 标签值列表
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetMaxResults(v int32) *ListTagResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 标签Key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// 标签Value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type GetListenerAttributeRequest struct {
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

type GetListenerAttributeResponseBody struct {
	// ACL相关配置信息
	AclConfig *GetListenerAttributeResponseBodyAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// 监听默认服务器证书列表，N当前取值范围为1
	Certificates []*GetListenerAttributeResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// 是否开启双向认证
	CaEnabled *bool `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// 默认动作
	DefaultActions []*GetListenerAttributeResponseBodyDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// 是否开启Gzip压缩
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// 是否开启HTTP/2特性
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// 连接空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 监听状态
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// 负载均衡标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 监听访问日志相关配置
	LogConfig *GetListenerAttributeResponseBodyLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// HTTPS启用QUIC时相关属性
	QuicConfig *GetListenerAttributeResponseBodyQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求超时时间
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// 安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// XForward字段相关的配置
	XForwardedForConfig *GetListenerAttributeResponseBodyXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) SetAclConfig(v *GetListenerAttributeResponseBodyAclConfig) *GetListenerAttributeResponseBody {
	s.AclConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCertificates(v []*GetListenerAttributeResponseBodyCertificates) *GetListenerAttributeResponseBody {
	s.Certificates = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetCaEnabled(v bool) *GetListenerAttributeResponseBody {
	s.CaEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetDefaultActions(v []*GetListenerAttributeResponseBodyDefaultActions) *GetListenerAttributeResponseBody {
	s.DefaultActions = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetGzipEnabled(v bool) *GetListenerAttributeResponseBody {
	s.GzipEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetHttp2Enabled(v bool) *GetListenerAttributeResponseBody {
	s.Http2Enabled = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetIdleTimeout(v int32) *GetListenerAttributeResponseBody {
	s.IdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerPort(v int32) *GetListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerProtocol(v string) *GetListenerAttributeResponseBody {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLogConfig(v *GetListenerAttributeResponseBodyLogConfig) *GetListenerAttributeResponseBody {
	s.LogConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetQuicConfig(v *GetListenerAttributeResponseBodyQuicConfig) *GetListenerAttributeResponseBody {
	s.QuicConfig = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestTimeout(v int32) *GetListenerAttributeResponseBody {
	s.RequestTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSecurityPolicyId(v string) *GetListenerAttributeResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetXForwardedForConfig(v *GetListenerAttributeResponseBodyXForwardedForConfig) *GetListenerAttributeResponseBody {
	s.XForwardedForConfig = v
	return s
}

type GetListenerAttributeResponseBodyAclConfig struct {
	// 监听绑定的访问策略组
	AclRelations []*GetListenerAttributeResponseBodyAclConfigAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	// 访问控制类型
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclRelations(v []*GetListenerAttributeResponseBodyAclConfigAclRelations) *GetListenerAttributeResponseBodyAclConfig {
	s.AclRelations = v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfig) SetAclType(v string) *GetListenerAttributeResponseBodyAclConfig {
	s.AclType = &v
	return s
}

type GetListenerAttributeResponseBodyAclConfigAclRelations struct {
	// ACL标识
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// ACL与监听关联的状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyAclConfigAclRelations) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetAclId(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.AclId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyAclConfigAclRelations) SetStatus(v string) *GetListenerAttributeResponseBodyAclConfigAclRelations {
	s.Status = &v
	return s
}

type GetListenerAttributeResponseBodyCertificates struct {
	// 正式标识
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetListenerAttributeResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyCertificates) SetCertificateId(v string) *GetListenerAttributeResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

type GetListenerAttributeResponseBodyDefaultActions struct {
	// 转发到服务器组
	ForwardGroupConfig *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActions) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetForwardGroupConfig(v *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) *GetListenerAttributeResponseBodyDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *GetListenerAttributeResponseBodyDefaultActions) SetType(v string) *GetListenerAttributeResponseBodyDefaultActions {
	s.Type = &v
	return s
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig struct {
	// 服务器组列表
	ServerGroupTuples []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *GetListenerAttributeResponseBodyDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type GetListenerAttributeResponseBodyLogConfig struct {
	// 访问日志是否开启携带自定义Header
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// 访问日志Xtrace相关的配置
	AccessLogTracingConfig *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s GetListenerAttributeResponseBodyLogConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfig) SetAccessLogTracingConfig(v *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) *GetListenerAttributeResponseBodyLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

type GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig struct {
	// Xtrace功能状态
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// Xtrace功能状态
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// xtrace的类型
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig) SetTracingType(v string) *GetListenerAttributeResponseBodyLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type GetListenerAttributeResponseBodyQuicConfig struct {
	// 需要关联的QUIC监听ID，HTTPS监听时有效，QuicUpgradeEnabled为true时必选
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// 是否开启quic升级，HTTPS监听时有效
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyQuicConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicListenerId(v string) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyQuicConfig) SetQuicUpgradeEnabled(v bool) *GetListenerAttributeResponseBodyQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type GetListenerAttributeResponseBodyXForwardedForConfig struct {
	// 自定义HEADER头名称，只有当XForwardedForClientCertClientVerifyEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-clientverify  头字段获取对访问负载均衡实例客户端证书的校验结果。HTTPS监听有效。
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertFingerprintEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-fingerprint 头字段获取访问负载均衡实例客户端证书的指纹取值，HTTPS监听有效。
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertIssuerDNEnabled的值为‘On’的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// 是否通过 X-Forwarded-Clientcert-issuerdn 头字段获取访问负载均衡实例客户端证书的发行者信息。HTTPS监听有效。
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertSubjectDNEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-subjectdn  头字段获取访问负载均衡实例客户端证书的所有者信息。HTTPS监听有效。
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// 是否通过X-Forwarded-Client-Port 头字段获取访问负载均衡实例客户端的端口。HTTPS监听有效。
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// 是否开启通过X-Forwarded-For头字段获取来访者真实 IP
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// 是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// 是否通过SLB-ID头字段获取负载均衡实例ID。
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// 是否通过X-Forwarded-Port 头字段获取负载均衡实例的监听端口。HTTPS监听有效。
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *GetListenerAttributeResponseBodyXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type GetListenerAttributeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponse) SetHeaders(v map[string]*string) *GetListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetListenerAttributeResponse) SetBody(v *GetListenerAttributeResponseBody) *GetListenerAttributeResponse {
	s.Body = v
	return s
}

type UpdateHealthCheckTemplateAttributeRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 名称
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// 端口号
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 时间间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// uri
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// 协议
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 健康阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 不健康阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 健康检查模板ID
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetClientToken(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateName(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckConnectPort(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHost(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckCodes(v []*string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckInterval(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTimeout(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTimeout = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckPath(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckMethod(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckProtocol(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckHttpVersion(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetUnhealthyThreshold(v int32) *UpdateHealthCheckTemplateAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetDryRun(v bool) *UpdateHealthCheckTemplateAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *UpdateHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

type UpdateHealthCheckTemplateAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *UpdateHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateHealthCheckTemplateAttributeResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateHealthCheckTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *UpdateHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateHealthCheckTemplateAttributeResponse) SetBody(v *UpdateHealthCheckTemplateAttributeResponseBody) *UpdateHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

type ListServerGroupServersRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 服务器组id
	ServerGroupId *string   `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	ServerIds     []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
}

func (s ListServerGroupServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

type ListServerGroupServersResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 后端服务器
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) SetMaxResults(v int32) *ListServerGroupServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetNextToken(v string) *ListServerGroupServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetTotalCount(v int32) *ListServerGroupServersResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupServersResponseBodyServers struct {
	// 描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 后端服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 权重
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) SetDescription(v string) *ListServerGroupServersResponseBodyServers {
	s.Description = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerIp(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerType(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetStatus(v string) *ListServerGroupServersResponseBodyServers {
	s.Status = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetWeight(v int32) *ListServerGroupServersResponseBodyServers {
	s.Weight = &v
	return s
}

type ListServerGroupServersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServerGroupServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponse) SetHeaders(v map[string]*string) *ListServerGroupServersResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupServersResponse) SetBody(v *ListServerGroupServersResponseBody) *ListServerGroupServersResponse {
	s.Body = v
	return s
}

type RemoveEntriesFromAclRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 访问控制策略Id
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 条目列表
	Entries []*string `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
}

func (s RemoveEntriesFromAclRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclRequest) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclRequest) SetClientToken(v string) *RemoveEntriesFromAclRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetDryRun(v bool) *RemoveEntriesFromAclRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetAclId(v string) *RemoveEntriesFromAclRequest {
	s.AclId = &v
	return s
}

func (s *RemoveEntriesFromAclRequest) SetEntries(v []*string) *RemoveEntriesFromAclRequest {
	s.Entries = v
	return s
}

type RemoveEntriesFromAclResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveEntriesFromAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponseBody) SetJobId(v string) *RemoveEntriesFromAclResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveEntriesFromAclResponseBody) SetRequestId(v string) *RemoveEntriesFromAclResponseBody {
	s.RequestId = &v
	return s
}

type RemoveEntriesFromAclResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveEntriesFromAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEntriesFromAclResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEntriesFromAclResponse) GoString() string {
	return s.String()
}

func (s *RemoveEntriesFromAclResponse) SetHeaders(v map[string]*string) *RemoveEntriesFromAclResponse {
	s.Headers = v
	return s
}

func (s *RemoveEntriesFromAclResponse) SetBody(v *RemoveEntriesFromAclResponseBody) *RemoveEntriesFromAclResponse {
	s.Body = v
	return s
}

type DeleteSecurityPolicyRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 安全策略Id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s DeleteSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyRequest) SetClientToken(v string) *DeleteSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetDryRun(v bool) *DeleteSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

type DeleteSecurityPolicyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBody) SetRequestId(v string) *DeleteSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponse) SetHeaders(v map[string]*string) *DeleteSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityPolicyResponse) SetBody(v *DeleteSecurityPolicyResponseBody) *DeleteSecurityPolicyResponse {
	s.Body = v
	return s
}

type DisableLoadBalancerAccessLogRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DisableLoadBalancerAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogRequest) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogRequest) SetClientToken(v string) *DisableLoadBalancerAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableLoadBalancerAccessLogRequest) SetDryRun(v bool) *DisableLoadBalancerAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *DisableLoadBalancerAccessLogRequest) SetLoadBalancerId(v string) *DisableLoadBalancerAccessLogRequest {
	s.LoadBalancerId = &v
	return s
}

type DisableLoadBalancerAccessLogResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLoadBalancerAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *DisableLoadBalancerAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type DisableLoadBalancerAccessLogResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLoadBalancerAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponse) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *DisableLoadBalancerAccessLogResponse {
	s.Headers = v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) SetBody(v *DisableLoadBalancerAccessLogResponseBody) *DisableLoadBalancerAccessLogResponse {
	s.Body = v
	return s
}

type ListServerGroupsRequest struct {
	// 服务器组Id列表
	ServerGroupIds []*string `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	// 服务器组名称
	ServerGroupNames []*string `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// VpcId
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务器组
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) SetMaxResults(v int32) *ListServerGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetNextToken(v string) *ListServerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBody) SetTotalCount(v int32) *ListServerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupsResponseBodyServerGroups struct {
	// 健康检查配置
	HealthCheckConfig *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// 服务器组协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 调度策略
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 服务器组Id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 服务器组名称
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// 服务器组状态
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// 会话保持配置
	StickySessionConfig *ListServerGroupsResponseBodyServerGroupsStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// 服务器组所在VpcId
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheckConfig(v *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheckConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetStickySessionConfig(v *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) *ListServerGroupsResponseBodyServerGroups {
	s.StickySessionConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsHealthCheckConfig struct {
	// 端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 是否启用健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 间隔时间
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// uri
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 协议
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 不健康阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHost(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckCodes(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHttpVersion(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckMethod(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckPath(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckProtocol(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsStickySessionConfig struct {
	// Cookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// Cookie超时时间
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// 是否开启会话保持
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// 会话保持类型
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetCookie(v string) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetCookieTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetStickySessionEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) SetStickySessionType(v string) *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type ListServerGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponse) SetHeaders(v map[string]*string) *ListServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupsResponse) SetBody(v *ListServerGroupsResponseBody) *ListServerGroupsResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 指定资源Id
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 资源组id
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type EnableDeletionProtectionRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例id
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s EnableDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionRequest) SetClientToken(v string) *EnableDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableDeletionProtectionRequest) SetDryRun(v bool) *EnableDeletionProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *EnableDeletionProtectionRequest) SetResourceId(v string) *EnableDeletionProtectionRequest {
	s.ResourceId = &v
	return s
}

type EnableDeletionProtectionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionResponseBody) SetRequestId(v string) *EnableDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type EnableDeletionProtectionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *EnableDeletionProtectionResponse) SetHeaders(v map[string]*string) *EnableDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *EnableDeletionProtectionResponse) SetBody(v *EnableDeletionProtectionResponseBody) *EnableDeletionProtectionResponse {
	s.Body = v
	return s
}

type UpdateListenerAttributeRequest struct {
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 请求超时时间
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// 连接空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 是否开启Gzip压缩
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// 是否开启HTTP/2特性
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// 安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// XForward字段相关的配置
	XForwardedForConfig *UpdateListenerAttributeRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
	// HTTPS启用QUIC时相关属性
	QuicConfig *UpdateListenerAttributeRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// 监听默认服务器证书列表，N当前取值范围为1
	Certificates []*UpdateListenerAttributeRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// 监听默认动作
	DefaultActions []*UpdateListenerAttributeRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRequestTimeout(v int32) *UpdateListenerAttributeRequest {
	s.RequestTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetGzipEnabled(v bool) *UpdateListenerAttributeRequest {
	s.GzipEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetHttp2Enabled(v bool) *UpdateListenerAttributeRequest {
	s.Http2Enabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetXForwardedForConfig(v *UpdateListenerAttributeRequestXForwardedForConfig) *UpdateListenerAttributeRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetQuicConfig(v *UpdateListenerAttributeRequestQuicConfig) *UpdateListenerAttributeRequest {
	s.QuicConfig = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificates(v []*UpdateListenerAttributeRequestCertificates) *UpdateListenerAttributeRequest {
	s.Certificates = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDefaultActions(v []*UpdateListenerAttributeRequestDefaultActions) *UpdateListenerAttributeRequest {
	s.DefaultActions = v
	return s
}

type UpdateListenerAttributeRequestXForwardedForConfig struct {
	// 自定义HEADER头名称，只有当XForwardedForClientCertClientVerifyEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-clientverify  头字段获取对访问负载均衡实例客户端证书的校验结果。HTTPS监听有效。
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertFingerprintEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-fingerprint 头字段获取访问负载均衡实例客户端证书的指纹取值，HTTPS监听有效。
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertIssuerDNEnabled的值为‘On’的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// 是否通过 X-Forwarded-Clientcert-issuerdn 头字段获取访问负载均衡实例客户端证书的发行者信息。HTTPS监听有效。
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertSubjectDNEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-subjectdn  头字段获取访问负载均衡实例客户端证书的所有者信息。HTTPS监听有效。
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// 是否通过X-Forwarded-Client-Port 头字段获取访问负载均衡实例客户端的端口。HTTPS监听有效。
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// 是否开启通过X-Forwarded-For头字段获取来访者真实 IP
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// 是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// 是否通过SLB-ID头字段获取负载均衡实例ID。
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// 是否通过X-Forwarded-Port 头字段获取负载均衡实例的监听端口。HTTPS监听有效。
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *UpdateListenerAttributeRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type UpdateListenerAttributeRequestQuicConfig struct {
	// 需要关联的QUIC监听ID，HTTPS监听时有效，QuicUpgradeEnabled为true时必选
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// 是否开启quic升级，HTTPS监听时有效
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s UpdateListenerAttributeRequestQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicListenerId(v string) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *UpdateListenerAttributeRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type UpdateListenerAttributeRequestCertificates struct {
	// 正式标识
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s UpdateListenerAttributeRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestCertificates) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestCertificates) SetCertificateId(v string) *UpdateListenerAttributeRequestCertificates {
	s.CertificateId = &v
	return s
}

type UpdateListenerAttributeRequestDefaultActions struct {
	// 转发组
	ForwardGroupConfig *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetForwardGroupConfig(v *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) *UpdateListenerAttributeRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateListenerAttributeRequestDefaultActions) SetType(v string) *UpdateListenerAttributeRequestDefaultActions {
	s.Type = &v
	return s
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig struct {
	// 服务器组列表
	ServerGroupTuples []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateListenerAttributeRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type UpdateListenerAttributeResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) SetJobId(v string) *UpdateListenerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponse) SetHeaders(v map[string]*string) *UpdateListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerAttributeResponse) SetBody(v *UpdateListenerAttributeResponseBody) *UpdateListenerAttributeResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 负载均衡实例的专有网络ID。
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 负载均衡的地址类型
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// 是否开启删除保护
	DeletionProtectionEnabled *bool `json:"DeletionProtectionEnabled,omitempty" xml:"DeletionProtectionEnabled,omitempty"`
	// 负载均衡修改保护相关信息
	ModificationProtectionConfig *CreateLoadBalancerRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// 可用区及交换机映射列表
	ZoneMappings []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
	// 地址模式
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// 资源组
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 负载均衡的版本
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// 计费相关配置信息
	LoadBalancerBillingConfig *CreateLoadBalancerRequestLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeletionProtectionEnabled(v bool) *CreateLoadBalancerRequest {
	s.DeletionProtectionEnabled = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest {
	s.ModificationProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressAllocatedMode(v string) *CreateLoadBalancerRequest {
	s.AddressAllocatedMode = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerEdition(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest {
	s.LoadBalancerBillingConfig = v
	return s
}

type CreateLoadBalancerRequestModificationProtectionConfig struct {
	// 设置修改保护状态的原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 负载均衡修改保护状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetStatus(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

type CreateLoadBalancerRequestZoneMappings struct {
	// 交换机标识
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 可用区
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateLoadBalancerRequestLoadBalancerBillingConfig struct {
	// 实例的计费类型
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	// 负载均衡实例标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerResponse) SetBody(v *CreateLoadBalancerResponseBody) *CreateLoadBalancerResponse {
	s.Body = v
	return s
}

type DeleteHealthCheckTemplatesRequest struct {
	// 健康检查模板Id列表
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
	// 幂等token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteHealthCheckTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *DeleteHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetClientToken(v string) *DeleteHealthCheckTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHealthCheckTemplatesRequest) SetDryRun(v bool) *DeleteHealthCheckTemplatesRequest {
	s.DryRun = &v
	return s
}

type DeleteHealthCheckTemplatesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHealthCheckTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponseBody) SetRequestId(v string) *DeleteHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHealthCheckTemplatesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHealthCheckTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *DeleteHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteHealthCheckTemplatesResponse) SetBody(v *DeleteHealthCheckTemplatesResponseBody) *DeleteHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

type UpdateServerGroupAttributeRequest struct {
	// Acl名称
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// 调度策略
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 健康检查配置
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// 会话保持配置
	StickySessionConfig *UpdateServerGroupAttributeRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// 服务器组Id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupName(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupName = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetScheduler(v string) *UpdateServerGroupAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetDryRun(v bool) *UpdateServerGroupAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetStickySessionConfig(v *UpdateServerGroupAttributeRequestStickySessionConfig) *UpdateServerGroupAttributeRequest {
	s.StickySessionConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupId = &v
	return s
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// 健康检查端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 是否启用健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 健康检查域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 健康检查正常的状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 健康检查HTTP协议版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 健康检查间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// 健康检查Path
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 健康检查协议类型
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 健康检查超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康检查成功判定阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查不成功判定阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHost(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckMethod(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckPath(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type UpdateServerGroupAttributeRequestStickySessionConfig struct {
	// 服务器上配置的Cookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// 服务器上配置的Cookie
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// 是否启用会话保持
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// 会话保持类型
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookie(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookieTimeout(v int32) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionEnabled(v bool) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionType(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type UpdateServerGroupAttributeResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) SetJobId(v string) *UpdateServerGroupAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServerGroupAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetBody(v *UpdateServerGroupAttributeResponseBody) *UpdateServerGroupAttributeResponse {
	s.Body = v
	return s
}

type CreateServerGroupRequest struct {
	// Acl名称
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// VpcId
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 调度策略
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 后端协议类型
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 健康检查配置
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// 会话保持配置
	StickySessionConfig *CreateServerGroupRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s CreateServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequest) SetServerGroupName(v string) *CreateServerGroupRequest {
	s.ServerGroupName = &v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerGroupRequest) SetScheduler(v string) *CreateServerGroupRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateServerGroupRequest) SetProtocol(v string) *CreateServerGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateServerGroupRequest) SetResourceGroupId(v string) *CreateServerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerGroupRequest) SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetStickySessionConfig(v *CreateServerGroupRequestStickySessionConfig) *CreateServerGroupRequest {
	s.StickySessionConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetDryRun(v bool) *CreateServerGroupRequest {
	s.DryRun = &v
	return s
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// 健康检查端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 是否启用健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 健康检查域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 健康检查正常的状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 健康检查HTTP协议版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 健康检查间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// 健康检查Path
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 健康检查协议类型
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 健康检查超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康检查成功判定阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查不成功判定阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateServerGroupRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHost(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckMethod(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckPath(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type CreateServerGroupRequestStickySessionConfig struct {
	// 服务器上配置的Cookie
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// 服务器上配置的Cookie
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// 是否启用会话保持
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// 会话保持类型
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s CreateServerGroupRequestStickySessionConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookie(v string) *CreateServerGroupRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookieTimeout(v int32) *CreateServerGroupRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionEnabled(v bool) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionType(v string) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

type CreateServerGroupResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务器组id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBody) SetJobId(v string) *CreateServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetRequestId(v string) *CreateServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetServerGroupId(v string) *CreateServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type CreateServerGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponse) SetHeaders(v map[string]*string) *CreateServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateServerGroupResponse) SetBody(v *CreateServerGroupResponseBody) *CreateServerGroupResponse {
	s.Body = v
	return s
}

type AssociateAclsWithListenerRequest struct {
	// 访问控制策略Id
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// 监听Id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 绑定类型
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// 预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AssociateAclsWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerRequest) SetAclIds(v []*string) *AssociateAclsWithListenerRequest {
	s.AclIds = v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetListenerId(v string) *AssociateAclsWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetAclType(v string) *AssociateAclsWithListenerRequest {
	s.AclType = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetDryRun(v bool) *AssociateAclsWithListenerRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAclsWithListenerRequest) SetClientToken(v string) *AssociateAclsWithListenerRequest {
	s.ClientToken = &v
	return s
}

type AssociateAclsWithListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAclsWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponseBody) SetJobId(v string) *AssociateAclsWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAclsWithListenerResponseBody) SetRequestId(v string) *AssociateAclsWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAclsWithListenerResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateAclsWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAclsWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAclsWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAclsWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAclsWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAclsWithListenerResponse) SetBody(v *AssociateAclsWithListenerResponseBody) *AssociateAclsWithListenerResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 转发规则标识
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetClientToken(v string) *DeleteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRuleRequest) SetDryRun(v bool) *DeleteRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteRuleResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetJobId(v string) *DeleteRuleResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，取值1-100，用户传入为空时，默认为20。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 转发规则ID列表，N最大支持20
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// 监听ID列表
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetNextToken(v string) *ListRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRulesRequest) SetMaxResults(v int32) *ListRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRulesRequest) SetRuleIds(v []*string) *ListRulesRequest {
	s.RuleIds = v
	return s
}

func (s *ListRulesRequest) SetListenerIds(v []*string) *ListRulesRequest {
	s.ListenerIds = v
	return s
}

type ListRulesResponseBody struct {
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 用来表示当前调用返回读取到的位置，空代表数据已经读取完毕。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 转发规则列表
	Rules []*ListRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// 本次请求条件下的数据总量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetMaxResults(v int32) *ListRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRulesResponseBody) SetNextToken(v string) *ListRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetRules(v []*ListRulesResponseBodyRules) *ListRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListRulesResponseBody) SetTotalCount(v int32) *ListRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListRulesResponseBodyRules struct {
	// 监听ID
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 转发规则优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 转发规则动作
	RuleActions []*ListRulesResponseBodyRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// 转发规则条件
	RuleConditions []*ListRulesResponseBodyRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// 转发规则标识
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 转发规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 转发规则状态
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
}

func (s ListRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRules) SetListenerId(v string) *ListRulesResponseBodyRules {
	s.ListenerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetLoadBalancerId(v string) *ListRulesResponseBodyRules {
	s.LoadBalancerId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetPriority(v int32) *ListRulesResponseBodyRules {
	s.Priority = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleActions(v []*ListRulesResponseBodyRulesRuleActions) *ListRulesResponseBodyRules {
	s.RuleActions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleConditions(v []*ListRulesResponseBodyRulesRuleConditions) *ListRulesResponseBodyRules {
	s.RuleConditions = v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleId(v string) *ListRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleName(v string) *ListRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *ListRulesResponseBodyRules) SetRuleStatus(v string) *ListRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

type ListRulesResponseBodyRulesRuleActions struct {
	// 返回固定内容动作配置
	FixedResponseConfig *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// 转发组动作配置
	ForwardGroupConfig *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 插入头部动作配置
	InsertHeaderConfig *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// 优先级
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 重定向动作配置
	RedirectConfig *ListRulesResponseBodyRulesRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// 内部重定向动作配置
	RewriteConfig *ListRulesResponseBodyRulesRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// 转发规则动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActions) SetFixedResponseConfig(v *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) *ListRulesResponseBodyRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetForwardGroupConfig(v *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) *ListRulesResponseBodyRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetInsertHeaderConfig(v *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) *ListRulesResponseBodyRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetOrder(v int32) *ListRulesResponseBodyRulesRuleActions {
	s.Order = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRedirectConfig(v *ListRulesResponseBodyRulesRuleActionsRedirectConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetRewriteConfig(v *ListRulesResponseBodyRulesRuleActionsRewriteConfig) *ListRulesResponseBodyRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActions) SetType(v string) *ListRulesResponseBodyRulesRuleActions {
	s.Type = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsFixedResponseConfig struct {
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 内容类型
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// HTTP响应码
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContent(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetContentType(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfig struct {
	// 转发到的目的服务器组列表
	ServerGroupTuples []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组标识
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListRulesResponseBodyRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig struct {
	// HTTP标头
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 取值类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValue(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *ListRulesResponseBodyRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsRedirectConfig struct {
	// 要跳转的主机地址
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 跳转方式
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 要跳转的路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 要跳转的端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 要跳转的协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 要跳转的查询字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetHttpCode(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetPort(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetProtocol(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRedirectConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type ListRulesResponseBodyRulesRuleActionsRewriteConfig struct {
	// 主机名
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 查询
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetHost(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetPath(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleActionsRewriteConfig) SetQuery(v string) *ListRulesResponseBodyRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditions struct {
	// Cookie条件配置
	CookieConfig *ListRulesResponseBodyRulesRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// HTTP标头条件配置
	HeaderConfig *ListRulesResponseBodyRulesRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// 主机名条件配置
	HostConfig *ListRulesResponseBodyRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// HTTP请求方法条件配置
	MethodConfig *ListRulesResponseBodyRulesRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// 路径条件配置
	PathConfig *ListRulesResponseBodyRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// 查询字符串条件配置
	QueryStringConfig *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// 条件类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetCookieConfig(v *ListRulesResponseBodyRulesRuleConditionsCookieConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHeaderConfig(v *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetHostConfig(v *ListRulesResponseBodyRulesRuleConditionsHostConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetMethodConfig(v *ListRulesResponseBodyRulesRuleConditionsMethodConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetPathConfig(v *ListRulesResponseBodyRulesRuleConditionsPathConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetQueryStringConfig(v *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) *ListRulesResponseBodyRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditions) SetType(v string) *ListRulesResponseBodyRulesRuleConditions {
	s.Type = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfig struct {
	// Cookie键值对列表
	Values []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) *ListRulesResponseBodyRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsCookieConfigValues struct {
	// Cookie条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Cookie条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsHeaderConfig struct {
	// HTTP标头键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头值列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsHeaderConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsHostConfig struct {
	// 主机名列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsHostConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsMethodConfig struct {
	// HTTP请求方法列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsMethodConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsPathConfig struct {
	// 路径条件列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsPathConfig) SetValues(v []*string) *ListRulesResponseBodyRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfig struct {
	// 查询字符串条件键值对列表
	Values []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig) SetValues(v []*ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues struct {
	// 查询字符串条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 查询字符串条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *ListRulesResponseBodyRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type ListRulesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type DescribeZonesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 可用区列表
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// 可用区名称
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 可用区id
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DeleteServerGroupRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 服务器组id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 是否DryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

type DeleteServerGroupResponseBody struct {
	// job
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBody) SetJobId(v string) *DeleteServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetRequestId(v string) *DeleteServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponse) SetHeaders(v map[string]*string) *DeleteServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerGroupResponse) SetBody(v *DeleteServerGroupResponseBody) *DeleteServerGroupResponse {
	s.Body = v
	return s
}

type AddServersToServerGroupRequest struct {
	// 后端服务器Id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 后端服务器
	Servers []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AddServersToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

type AddServersToServerGroupRequestServers struct {
	// 描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 后端端口号
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 后端服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 后端服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 后端服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 后端服务器权重
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s AddServersToServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequestServers) SetDescription(v string) *AddServersToServerGroupRequestServers {
	s.Description = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetPort(v int32) *AddServersToServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerId(v string) *AddServersToServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerIp(v string) *AddServersToServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerType(v string) *AddServersToServerGroupRequestServers {
	s.ServerType = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetWeight(v int32) *AddServersToServerGroupRequestServers {
	s.Weight = &v
	return s
}

type AddServersToServerGroupResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServersToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) SetJobId(v string) *AddServersToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddServersToServerGroupResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddServersToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddServersToServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponse) SetHeaders(v map[string]*string) *AddServersToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *AddServersToServerGroupResponse) SetBody(v *AddServersToServerGroupResponseBody) *AddServersToServerGroupResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerRequest struct {
	// 实例id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetJobId(v string) *DeleteLoadBalancerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerResponse) SetBody(v *DeleteLoadBalancerResponseBody) *DeleteLoadBalancerResponse {
	s.Body = v
	return s
}

type UpdateListenerLogConfigRequest struct {
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 是否开启携带自定义Header
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// 访问日志xtrace字段相关的配置
	AccessLogTracingConfig *UpdateListenerLogConfigRequestAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s UpdateListenerLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequest) SetListenerId(v string) *UpdateListenerLogConfigRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetClientToken(v string) *UpdateListenerLogConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetDryRun(v bool) *UpdateListenerLogConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *UpdateListenerLogConfigRequest {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequest) SetAccessLogTracingConfig(v *UpdateListenerLogConfigRequestAccessLogTracingConfig) *UpdateListenerLogConfigRequest {
	s.AccessLogTracingConfig = v
	return s
}

type UpdateListenerLogConfigRequestAccessLogTracingConfig struct {
	// Xtrace功能状态
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// xtrace的采样率
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// xtrace的类型
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigRequestAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingEnabled(v bool) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingSample(v int32) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *UpdateListenerLogConfigRequestAccessLogTracingConfig) SetTracingType(v string) *UpdateListenerLogConfigRequestAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type UpdateListenerLogConfigResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponseBody) SetJobId(v string) *UpdateListenerLogConfigResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateListenerLogConfigResponseBody) SetRequestId(v string) *UpdateListenerLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerLogConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateListenerLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateListenerLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerLogConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerLogConfigResponse) SetHeaders(v map[string]*string) *UpdateListenerLogConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerLogConfigResponse) SetBody(v *UpdateListenerLogConfigResponseBody) *UpdateListenerLogConfigResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源实例Id
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// 标签列表
	Tag []*UnTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// 标签键列表
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetTag(v []*UnTagResourcesRequestTag) *UnTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesRequestTag struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UnTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequestTag) SetKey(v string) *UnTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *UnTagResourcesRequestTag) SetValue(v string) *UnTagResourcesRequestTag {
	s.Value = &v
	return s
}

type UnTagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type ListSystemSecurityPoliciesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 安全策略
	SecurityPolicies []*ListSystemSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBody) SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

type ListSystemSecurityPoliciesResponseBodySecurityPolicies struct {
	// 加密套件
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// 安全策略Id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// 协议版本
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

type ListSystemSecurityPoliciesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSystemSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSystemSecurityPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSystemSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetBody(v *ListSystemSecurityPoliciesResponseBody) *ListSystemSecurityPoliciesResponse {
	s.Body = v
	return s
}

type AddEntriesToAclRequest struct {
	// AclId
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 条目信息列表
	AclEntries []*AddEntriesToAclRequestAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AddEntriesToAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequest) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequest) SetAclId(v string) *AddEntriesToAclRequest {
	s.AclId = &v
	return s
}

func (s *AddEntriesToAclRequest) SetAclEntries(v []*AddEntriesToAclRequestAclEntries) *AddEntriesToAclRequest {
	s.AclEntries = v
	return s
}

func (s *AddEntriesToAclRequest) SetClientToken(v string) *AddEntriesToAclRequest {
	s.ClientToken = &v
	return s
}

func (s *AddEntriesToAclRequest) SetDryRun(v bool) *AddEntriesToAclRequest {
	s.DryRun = &v
	return s
}

type AddEntriesToAclRequestAclEntries struct {
	// 描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 条目
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s AddEntriesToAclRequestAclEntries) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclRequestAclEntries) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclRequestAclEntries) SetDescription(v string) *AddEntriesToAclRequestAclEntries {
	s.Description = &v
	return s
}

func (s *AddEntriesToAclRequestAclEntries) SetEntry(v string) *AddEntriesToAclRequestAclEntries {
	s.Entry = &v
	return s
}

type AddEntriesToAclResponseBody struct {
	// job
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEntriesToAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponseBody) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponseBody) SetJobId(v string) *AddEntriesToAclResponseBody {
	s.JobId = &v
	return s
}

func (s *AddEntriesToAclResponseBody) SetRequestId(v string) *AddEntriesToAclResponseBody {
	s.RequestId = &v
	return s
}

type AddEntriesToAclResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEntriesToAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEntriesToAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEntriesToAclResponse) GoString() string {
	return s.String()
}

func (s *AddEntriesToAclResponse) SetHeaders(v map[string]*string) *AddEntriesToAclResponse {
	s.Headers = v
	return s
}

func (s *AddEntriesToAclResponse) SetBody(v *AddEntriesToAclResponseBody) *AddEntriesToAclResponse {
	s.Body = v
	return s
}

type DeleteAclRequest struct {
	// 访问控制策略id
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// DryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRequest) SetAclId(v string) *DeleteAclRequest {
	s.AclId = &v
	return s
}

func (s *DeleteAclRequest) SetClientToken(v string) *DeleteAclRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAclRequest) SetDryRun(v bool) *DeleteAclRequest {
	s.DryRun = &v
	return s
}

type DeleteAclResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) SetJobId(v string) *DeleteAclResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAclResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

type DissociateAdditionalCertificatesFromListenerRequest struct {
	// 监听Id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 证书列表
	Certificates []*DissociateAdditionalCertificatesFromListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetListenerId(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetCertificates(v []*DissociateAdditionalCertificatesFromListenerRequestCertificates) *DissociateAdditionalCertificatesFromListenerRequest {
	s.Certificates = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetClientToken(v string) *DissociateAdditionalCertificatesFromListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerRequest) SetDryRun(v bool) *DissociateAdditionalCertificatesFromListenerRequest {
	s.DryRun = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerRequestCertificates struct {
	// 证书Id
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerRequestCertificates) SetCertificateId(v string) *DissociateAdditionalCertificatesFromListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetJobId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponseBody) SetRequestId(v string) *DissociateAdditionalCertificatesFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAdditionalCertificatesFromListenerResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateAdditionalCertificatesFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAdditionalCertificatesFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAdditionalCertificatesFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAdditionalCertificatesFromListenerResponse) SetBody(v *DissociateAdditionalCertificatesFromListenerResponseBody) *DissociateAdditionalCertificatesFromListenerResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerEditionRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例Id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 实例版本
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
}

func (s UpdateLoadBalancerEditionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionRequest) SetClientToken(v string) *UpdateLoadBalancerEditionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetDryRun(v bool) *UpdateLoadBalancerEditionRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerEditionRequest) SetLoadBalancerEdition(v string) *UpdateLoadBalancerEditionRequest {
	s.LoadBalancerEdition = &v
	return s
}

type UpdateLoadBalancerEditionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerEditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponseBody) SetRequestId(v string) *UpdateLoadBalancerEditionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerEditionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLoadBalancerEditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerEditionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerEditionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) SetBody(v *UpdateLoadBalancerEditionResponseBody) *UpdateLoadBalancerEditionResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，取值1-100，用户传入为空时，默认为20。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 监听ID列表，N最大支持20
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// 实例ID列表，N最大支持20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
	return s
}

type ListListenersResponseBody struct {
	// 监听列表
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 用来表示当前调用返回读取到的位置，空代表数据已经读取完毕。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	// 默认动作
	DefaultActions []*ListListenersResponseBodyListenersDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
	// 是否开启Gzip压缩
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// 是否开启HTTP/2特性
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// 连接空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 监听状态
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// 负载均衡标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 监听访问日志相关配置
	LogConfig *ListListenersResponseBodyListenersLogConfig `json:"LogConfig,omitempty" xml:"LogConfig,omitempty" type:"Struct"`
	// HTTPS启用QUIC时相关属性
	QuicConfig *ListListenersResponseBodyListenersQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// 请求超时时间
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// 安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// XForward字段相关的配置
	XForwardedForConfig *ListListenersResponseBodyListenersXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetDefaultActions(v []*ListListenersResponseBodyListenersDefaultActions) *ListListenersResponseBodyListeners {
	s.DefaultActions = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetGzipEnabled(v bool) *ListListenersResponseBodyListeners {
	s.GzipEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetHttp2Enabled(v bool) *ListListenersResponseBodyListeners {
	s.Http2Enabled = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerPort(v int32) *ListListenersResponseBodyListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerProtocol(v string) *ListListenersResponseBodyListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLogConfig(v *ListListenersResponseBodyListenersLogConfig) *ListListenersResponseBodyListeners {
	s.LogConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetQuicConfig(v *ListListenersResponseBodyListenersQuicConfig) *ListListenersResponseBodyListeners {
	s.QuicConfig = v
	return s
}

func (s *ListListenersResponseBodyListeners) SetRequestTimeout(v int32) *ListListenersResponseBodyListeners {
	s.RequestTimeout = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetXForwardedForConfig(v *ListListenersResponseBodyListenersXForwardedForConfig) *ListListenersResponseBodyListeners {
	s.XForwardedForConfig = v
	return s
}

type ListListenersResponseBodyListenersDefaultActions struct {
	// 转发到服务器组
	ForwardGroupConfig *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActions) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetForwardGroupConfig(v *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) *ListListenersResponseBodyListenersDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *ListListenersResponseBodyListenersDefaultActions) SetType(v string) *ListListenersResponseBodyListenersDefaultActions {
	s.Type = &v
	return s
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig struct {
	// 服务器组列表
	ServerGroupTuples []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *ListListenersResponseBodyListenersDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type ListListenersResponseBodyListenersLogConfig struct {
	// 访问日志是否开启携带自定义Header
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// 访问日志Xtrace相关的配置
	AccessLogTracingConfig *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig `json:"AccessLogTracingConfig,omitempty" xml:"AccessLogTracingConfig,omitempty" type:"Struct"`
}

func (s ListListenersResponseBodyListenersLogConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfig) SetAccessLogTracingConfig(v *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) *ListListenersResponseBodyListenersLogConfig {
	s.AccessLogTracingConfig = v
	return s
}

type ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig struct {
	// Xtrace功能状态
	TracingEnabled *bool `json:"TracingEnabled,omitempty" xml:"TracingEnabled,omitempty"`
	// Xtrace功能状态
	TracingSample *int32 `json:"TracingSample,omitempty" xml:"TracingSample,omitempty"`
	// xtrace的类型
	TracingType *string `json:"TracingType,omitempty" xml:"TracingType,omitempty"`
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingEnabled(v bool) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingSample(v int32) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingSample = &v
	return s
}

func (s *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig) SetTracingType(v string) *ListListenersResponseBodyListenersLogConfigAccessLogTracingConfig {
	s.TracingType = &v
	return s
}

type ListListenersResponseBodyListenersQuicConfig struct {
	// 需要关联的QUIC监听ID，HTTPS监听时有效，QuicUpgradeEnabled为true时必选
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// 是否开启quic升级，HTTPS监听时有效
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersQuicConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicListenerId(v string) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListenersQuicConfig) SetQuicUpgradeEnabled(v bool) *ListListenersResponseBodyListenersQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type ListListenersResponseBodyListenersXForwardedForConfig struct {
	// 自定义HEADER头名称，只有当XForwardedForClientCertClientVerifyEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-clientverify  头字段获取对访问负载均衡实例客户端证书的校验结果。HTTPS监听有效。
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertFingerprintEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-fingerprint 头字段获取访问负载均衡实例客户端证书的指纹取值，HTTPS监听有效。
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertIssuerDNEnabled的值为‘On’的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// 是否通过 X-Forwarded-Clientcert-issuerdn 头字段获取访问负载均衡实例客户端证书的发行者信息。HTTPS监听有效。
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertSubjectDNEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-subjectdn  头字段获取访问负载均衡实例客户端证书的所有者信息。HTTPS监听有效。
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// 是否通过X-Forwarded-Client-Port 头字段获取访问负载均衡实例客户端的端口。HTTPS监听有效。
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// 是否开启通过X-Forwarded-For头字段获取来访者真实 IP
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// 是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// 是否通过SLB-ID头字段获取负载均衡实例ID。
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// 是否通过X-Forwarded-Port 头字段获取负载均衡实例的监听端口。HTTPS监听有效。
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *ListListenersResponseBodyListenersXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *ListListenersResponseBodyListenersXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type ListListenersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponse) GoString() string {
	return s.String()
}

func (s *ListListenersResponse) SetHeaders(v map[string]*string) *ListListenersResponse {
	s.Headers = v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type ListAclRelationsRequest struct {
	// 访问控制策略Id
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
}

func (s ListAclRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListAclRelationsRequest) SetAclIds(v []*string) *ListAclRelationsRequest {
	s.AclIds = v
	return s
}

type ListAclRelationsResponseBody struct {
	// 访问控制列表
	AclRelations []*ListAclRelationsResponseBodyAclRelations `json:"AclRelations,omitempty" xml:"AclRelations,omitempty" type:"Repeated"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAclRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBody) SetAclRelations(v []*ListAclRelationsResponseBodyAclRelations) *ListAclRelationsResponseBody {
	s.AclRelations = v
	return s
}

func (s *ListAclRelationsResponseBody) SetRequestId(v string) *ListAclRelationsResponseBody {
	s.RequestId = &v
	return s
}

type ListAclRelationsResponseBodyAclRelations struct {
	// 访问控制策略id
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 关联的监听
	RelatedListeners []*ListAclRelationsResponseBodyAclRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
}

func (s ListAclRelationsResponseBodyAclRelations) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelations) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelations) SetAclId(v string) *ListAclRelationsResponseBodyAclRelations {
	s.AclId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelations) SetRelatedListeners(v []*ListAclRelationsResponseBodyAclRelationsRelatedListeners) *ListAclRelationsResponseBodyAclRelations {
	s.RelatedListeners = v
	return s
}

type ListAclRelationsResponseBodyAclRelationsRelatedListeners struct {
	// 监听ID
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 关联状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponseBodyAclRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerPort(v int32) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetListenerProtocol(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetLoadBalancerId(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListAclRelationsResponseBodyAclRelationsRelatedListeners) SetStatus(v string) *ListAclRelationsResponseBodyAclRelationsRelatedListeners {
	s.Status = &v
	return s
}

type ListAclRelationsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAclRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListAclRelationsResponse) SetHeaders(v map[string]*string) *ListAclRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListAclRelationsResponse) SetBody(v *ListAclRelationsResponseBody) *ListAclRelationsResponse {
	s.Body = v
	return s
}

type DeleteRulesRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 同一个监听下的转发规则标识列表
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRulesRequest) SetClientToken(v string) *DeleteRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRulesRequest) SetDryRun(v bool) *DeleteRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRulesRequest) SetRuleIds(v []*string) *DeleteRulesRequest {
	s.RuleIds = v
	return s
}

type DeleteRulesResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponseBody) SetJobId(v string) *DeleteRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteRulesResponseBody) SetRequestId(v string) *DeleteRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRulesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponse) SetHeaders(v map[string]*string) *DeleteRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRulesResponse) SetBody(v *DeleteRulesResponseBody) *DeleteRulesResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
	// 监听id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

type DeleteListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetJobId(v string) *DeleteListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListenerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponse) SetHeaders(v map[string]*string) *DeleteListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 标签类型
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 具体的标签Key
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetKeyword(v string) *ListTagKeysRequest {
	s.Keyword = &v
	return s
}

type ListTagKeysResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 标签Key列表
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	// 标签类型
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签Key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetCategory(v string) *ListTagKeysResponseBodyTagKeys {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequest struct {
	// 监听Id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 证书列表
	Certificates []*AssociateAdditionalCertificatesWithListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequest) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetListenerId(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetCertificates(v []*AssociateAdditionalCertificatesWithListenerRequestCertificates) *AssociateAdditionalCertificatesWithListenerRequest {
	s.Certificates = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetClientToken(v string) *AssociateAdditionalCertificatesWithListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerRequest) SetDryRun(v bool) *AssociateAdditionalCertificatesWithListenerRequest {
	s.DryRun = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerRequestCertificates struct {
	// 证书Id
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerRequestCertificates) SetCertificateId(v string) *AssociateAdditionalCertificatesWithListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetJobId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponseBody) SetRequestId(v string) *AssociateAdditionalCertificatesWithListenerResponseBody {
	s.RequestId = &v
	return s
}

type AssociateAdditionalCertificatesWithListenerResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssociateAdditionalCertificatesWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssociateAdditionalCertificatesWithListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateAdditionalCertificatesWithListenerResponse) GoString() string {
	return s.String()
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetHeaders(v map[string]*string) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Headers = v
	return s
}

func (s *AssociateAdditionalCertificatesWithListenerResponse) SetBody(v *AssociateAdditionalCertificatesWithListenerResponseBody) *AssociateAdditionalCertificatesWithListenerResponse {
	s.Body = v
	return s
}

type ListHealthCheckTemplatesRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 健康检查模板ID列表
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitempty" xml:"HealthCheckTemplateIds,omitempty" type:"Repeated"`
	// 健康检查模板名称列表
	HealthCheckTemplateNames []*string `json:"HealthCheckTemplateNames,omitempty" xml:"HealthCheckTemplateNames,omitempty" type:"Repeated"`
}

func (s ListHealthCheckTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesRequest) SetNextToken(v string) *ListHealthCheckTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetMaxResults(v int32) *ListHealthCheckTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateIds(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateIds = v
	return s
}

func (s *ListHealthCheckTemplatesRequest) SetHealthCheckTemplateNames(v []*string) *ListHealthCheckTemplatesRequest {
	s.HealthCheckTemplateNames = v
	return s
}

type ListHealthCheckTemplatesResponseBody struct {
	// 健康检查模板
	HealthCheckTemplates []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates `json:"HealthCheckTemplates,omitempty" xml:"HealthCheckTemplates,omitempty" type:"Repeated"`
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBody) SetHealthCheckTemplates(v []*ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) *ListHealthCheckTemplatesResponseBody {
	s.HealthCheckTemplates = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetMaxResults(v int32) *ListHealthCheckTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetNextToken(v string) *ListHealthCheckTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetRequestId(v string) *ListHealthCheckTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBody) SetTotalCount(v int32) *ListHealthCheckTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListHealthCheckTemplatesResponseBodyHealthCheckTemplates struct {
	// 端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 间隔时间
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// uri
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 协议
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 健康检查模板Id
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// 名称
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// 超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 不健康阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckConnectPort(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHost(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHost = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckCodes(v []*string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckCodes = v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckHttpVersion(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckInterval(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckMethod(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckMethod = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckPath(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckPath = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckProtocol(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateId(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTemplateName(v string) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthCheckTimeout(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthCheckTimeout = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetHealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.HealthyThreshold = &v
	return s
}

func (s *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates) SetUnhealthyThreshold(v int32) *ListHealthCheckTemplatesResponseBodyHealthCheckTemplates {
	s.UnhealthyThreshold = &v
	return s
}

type ListHealthCheckTemplatesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHealthCheckTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHealthCheckTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHealthCheckTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListHealthCheckTemplatesResponse) SetHeaders(v map[string]*string) *ListHealthCheckTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListHealthCheckTemplatesResponse) SetBody(v *ListHealthCheckTemplatesResponseBody) *ListHealthCheckTemplatesResponse {
	s.Body = v
	return s
}

type GetLoadBalancerAttributeRequest struct {
	// 实例标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

type GetLoadBalancerAttributeResponseBody struct {
	// 访问日志属性
	AccessLogConfig *GetLoadBalancerAttributeResponseBodyAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// 地址分配方式
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// 地址类型
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 带宽包ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// DNS域名
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// 负载均衡删除保护相关信息
	DeletionProtectionConfig *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// 计费相关属性
	LoadBalancerBillingConfig *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// 实例业务状态
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// 负载均衡的版本
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// 负载均衡标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// 锁定原因
	LoadBalancerOperationLocks []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
	// 实例状态
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// 负载均衡修改保护相关信息
	ModificationProtectionConfig *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// 地域
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 企业资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 标签列表
	Tags []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Vpc网络ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 负载均衡的可用区资源
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) SetAccessLogConfig(v *GetLoadBalancerAttributeResponseBodyAccessLogConfig) *GetLoadBalancerAttributeResponseBody {
	s.AccessLogConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressAllocatedMode(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressAllocatedMode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDNSName(v string) *GetLoadBalancerAttributeResponseBody {
	s.DNSName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.DeletionProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBussinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerEdition(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerEdition = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerOperationLocks(v []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.ModificationProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRegionId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

type GetLoadBalancerAttributeResponseBodyAccessLogConfig struct {
	// 访问日志投递的logProject
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// 删除保护开启时间
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogProject(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogStore(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogStore = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig struct {
	// 删除保护状态
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 删除保护开启时间
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabled(v bool) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig struct {
	// 实例的计费类型
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) SetPayType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks struct {
	// 锁定原因
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// 锁定类型
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyModificationProtectionConfig struct {
	// 设置修改保护状态的原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 负载均衡修改保护状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Status = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyTags struct {
	// 实例的标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 实例的标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// 固定VIP模式下，负载均衡在此可用区中的地址列表
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// 交换机标识
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 可用区标识
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// IP地址
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAddress(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Address = &v
	return s
}

type GetLoadBalancerAttributeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *GetLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetBody(v *GetLoadBalancerAttributeResponseBody) *GetLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源实例Id
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 具体的标签Key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceId(v string) *ListTagValuesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

type ListTagValuesResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 标签值列表
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *ListTagValuesResponseBody) SetTotalCount(v int32) *ListTagValuesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListLoadBalancersRequest struct {
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量，此参数为可选参数，取值1-100，用户传入为空时，默认为20。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 可用区ID
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// 实例状态
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// 实例业务状态
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// 实例ID列表，N最大支持20
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// 实例Name列表，N最大支持10
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// vpcId列表
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// tag列表
	Tag []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// 负载均衡的地址类型
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 付费类型
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneId(v string) *ListLoadBalancersRequest {
	s.ZoneId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerNames = v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetAddressType(v string) *ListLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetPayType(v string) *ListLoadBalancersRequest {
	s.PayType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

type ListLoadBalancersRequestTag struct {
	// 实例的标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 实例的标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequestTag) SetKey(v string) *ListLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersRequestTag) SetValue(v string) *ListLoadBalancersRequestTag {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBody struct {
	// 实例列表
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 用来表示当前调用返回读取到的位置，空代表数据已经读取完毕。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMaxResults(v int32) *ListLoadBalancersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetNextToken(v string) *ListLoadBalancersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// 访问日志属性
	AccessLogConfig *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// 地址模式
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// 地址类型
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 带宽包ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// 资源创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// DNS域名
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// 负载均衡删除保护相关信息
	DeletionProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// 计费相关属性
	LoadBalancerBillingConfig *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// 实例业务状态
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// 负载均衡的版本
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// 负载均衡标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// 锁定的原因
	LoadBalancerOperationLocks []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
	// 实例状态
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// 负载均衡修改保护相关信息
	ModificationProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// 企业资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 标签列表
	Tags []*ListLoadBalancersResponseBodyLoadBalancersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Vpc网络ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAccessLogConfig(v *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AccessLogConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressAllocatedMode(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressAllocatedMode = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetBandwidthPackageId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDNSName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDeletionProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DeletionProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBillingConfig(v *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerEdition(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerEdition = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerOperationLocks(v []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetModificationProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ModificationProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTags(v []*ListLoadBalancersResponseBodyLoadBalancersTags) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Tags = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.VpcId = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig struct {
	// 访问日志投递的logProject
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// 删除保护开启时间
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogProject(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogStore(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogStore = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig struct {
	// 删除保护状态
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 删除保护开启时间
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabledTime(v string) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig struct {
	// 实例的计费类型
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) SetPayType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks struct {
	// 锁定的原因
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// 锁定的类型
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockReason(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig struct {
	// 设置修改保护状态的原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 负载均衡修改保护状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetReason(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Status = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersTags struct {
	// 实例的标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 实例的标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetKey(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetValue(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Value = &v
	return s
}

type ListLoadBalancersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponse) SetHeaders(v map[string]*string) *ListLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancersResponse) SetBody(v *ListLoadBalancersResponseBody) *ListLoadBalancersResponse {
	s.Body = v
	return s
}

type EnableLoadBalancerAccessLogRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 日志Project
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// 日志Store
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s EnableLoadBalancerAccessLogRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogRequest) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogRequest) SetClientToken(v string) *EnableLoadBalancerAccessLogRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetDryRun(v bool) *EnableLoadBalancerAccessLogRequest {
	s.DryRun = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLoadBalancerId(v string) *EnableLoadBalancerAccessLogRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogProject(v string) *EnableLoadBalancerAccessLogRequest {
	s.LogProject = &v
	return s
}

func (s *EnableLoadBalancerAccessLogRequest) SetLogStore(v string) *EnableLoadBalancerAccessLogRequest {
	s.LogStore = &v
	return s
}

type EnableLoadBalancerAccessLogResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLoadBalancerAccessLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogResponseBody) SetRequestId(v string) *EnableLoadBalancerAccessLogResponseBody {
	s.RequestId = &v
	return s
}

type EnableLoadBalancerAccessLogResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLoadBalancerAccessLogResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLoadBalancerAccessLogResponse) GoString() string {
	return s.String()
}

func (s *EnableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *EnableLoadBalancerAccessLogResponse {
	s.Headers = v
	return s
}

func (s *EnableLoadBalancerAccessLogResponse) SetBody(v *EnableLoadBalancerAccessLogResponseBody) *EnableLoadBalancerAccessLogResponse {
	s.Body = v
	return s
}

type DissociateAclsFromListenerRequest struct {
	// 访问控制策略Id
	AclIds []*string `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Repeated"`
	// 监听Id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DissociateAclsFromListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerRequest) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerRequest) SetAclIds(v []*string) *DissociateAclsFromListenerRequest {
	s.AclIds = v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetListenerId(v string) *DissociateAclsFromListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetDryRun(v bool) *DissociateAclsFromListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateAclsFromListenerRequest) SetClientToken(v string) *DissociateAclsFromListenerRequest {
	s.ClientToken = &v
	return s
}

type DissociateAclsFromListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateAclsFromListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponseBody) SetJobId(v string) *DissociateAclsFromListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *DissociateAclsFromListenerResponseBody) SetRequestId(v string) *DissociateAclsFromListenerResponseBody {
	s.RequestId = &v
	return s
}

type DissociateAclsFromListenerResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DissociateAclsFromListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DissociateAclsFromListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateAclsFromListenerResponse) GoString() string {
	return s.String()
}

func (s *DissociateAclsFromListenerResponse) SetHeaders(v map[string]*string) *DissociateAclsFromListenerResponse {
	s.Headers = v
	return s
}

func (s *DissociateAclsFromListenerResponse) SetBody(v *DissociateAclsFromListenerResponseBody) *DissociateAclsFromListenerResponse {
	s.Body = v
	return s
}

type ListListenerCertificatesRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 监听Id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 证书类型
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetCertificateType(v string) *ListListenerCertificatesRequest {
	s.CertificateType = &v
	return s
}

type ListListenerCertificatesResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 监听SSL证书列表
	Certificates []*ListListenerCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
}

func (s ListListenerCertificatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBody) SetMaxResults(v int32) *ListListenerCertificatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetNextToken(v string) *ListListenerCertificatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetRequestId(v string) *ListListenerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetTotalCount(v int32) *ListListenerCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody {
	s.Certificates = v
	return s
}

type ListListenerCertificatesResponseBodyCertificates struct {
	// 证书Id
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// 是否为默认证书
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 证书状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 证书类型
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
}

func (s ListListenerCertificatesResponseBodyCertificates) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetStatus(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Status = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateType(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateType = &v
	return s
}

type ListListenerCertificatesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListListenerCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListListenerCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenerCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponse) SetHeaders(v map[string]*string) *ListListenerCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListListenerCertificatesResponse) SetBody(v *ListListenerCertificatesResponseBody) *ListListenerCertificatesResponse {
	s.Body = v
	return s
}

type CreateAclRequest struct {
	// Acl名称
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 资源组Id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) SetAclName(v string) *CreateAclRequest {
	s.AclName = &v
	return s
}

func (s *CreateAclRequest) SetClientToken(v string) *CreateAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAclRequest) SetDryRun(v bool) *CreateAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAclRequest) SetResourceGroupId(v string) *CreateAclRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateAclResponseBody struct {
	// AclId
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) SetAclId(v string) *CreateAclResponseBody {
	s.AclId = &v
	return s
}

func (s *CreateAclResponseBody) SetJobId(v string) *CreateAclResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

type CreateAclResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAclResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponse) GoString() string {
	return s.String()
}

func (s *CreateAclResponse) SetHeaders(v map[string]*string) *CreateAclResponse {
	s.Headers = v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

type GetHealthCheckTemplateAttributeRequest struct {
	// 健康检查模板Id
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
}

func (s GetHealthCheckTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeRequest) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeRequest {
	s.HealthCheckTemplateId = &v
	return s
}

type GetHealthCheckTemplateAttributeResponseBody struct {
	// 端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 域名
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// 状态码
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// 版本
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// 健康检查模板Id
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitempty" xml:"HealthCheckTemplateId,omitempty"`
	// 间隔时间
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 方法
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// uri
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// 协议
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// 名称
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitempty" xml:"HealthCheckTemplateName,omitempty"`
	// 超时时间
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// 健康阈值
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 不健康阈值
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s GetHealthCheckTemplateAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckConnectPort(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHost(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHost = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckCodes(v []*string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckCodes = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckHttpVersion(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckInterval(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckMethod(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckPath(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckProtocol(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTemplateName(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTemplateName = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthCheckTimeout(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthCheckTimeout = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetHealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetRequestId(v string) *GetHealthCheckTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponseBody) SetUnhealthyThreshold(v int32) *GetHealthCheckTemplateAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

type GetHealthCheckTemplateAttributeResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHealthCheckTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHealthCheckTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthCheckTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetHealthCheckTemplateAttributeResponse) SetHeaders(v map[string]*string) *GetHealthCheckTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetHealthCheckTemplateAttributeResponse) SetBody(v *GetHealthCheckTemplateAttributeResponseBody) *GetHealthCheckTemplateAttributeResponse {
	s.Body = v
	return s
}

type UpdateSecurityPolicyAttributeRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 安全策略id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// TLS版本
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
	// 加密套件
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// 安全策略名称
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
}

func (s UpdateSecurityPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeRequest) SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetTLSVersions(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.TLSVersions = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyName(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyName = &v
	return s
}

type UpdateSecurityPolicyAttributeResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSecurityPolicyAttributeResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSecurityPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSecurityPolicyAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateSecurityPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponse) SetBody(v *UpdateSecurityPolicyAttributeResponseBody) *UpdateSecurityPolicyAttributeResponse {
	s.Body = v
	return s
}

type ListSecurityPolicyRelationsRequest struct {
	// 安全策略id
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRelationsRequest {
	s.SecurityPolicyIds = v
	return s
}

type ListSecurityPolicyRelationsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 安全策略关联关系
	SecrityPolicyRelations []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations `json:"SecrityPolicyRelations,omitempty" xml:"SecrityPolicyRelations,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBody) SetRequestId(v string) *ListSecurityPolicyRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBody) SetSecrityPolicyRelations(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) *ListSecurityPolicyRelationsResponseBody {
	s.SecrityPolicyRelations = v
	return s
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations struct {
	// 关联的监听列表
	RelatedListeners []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners `json:"RelatedListeners,omitempty" xml:"RelatedListeners,omitempty" type:"Repeated"`
	// 安全策略id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetRelatedListeners(v []*ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.RelatedListeners = v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations) SetSecurityPolicyId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelations {
	s.SecurityPolicyId = &v
	return s
}

type ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners struct {
	// 监听id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 实例id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerId = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerPort(v int64) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetListenerProtocol(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners) SetLoadBalancerId(v string) *ListSecurityPolicyRelationsResponseBodySecrityPolicyRelationsRelatedListeners {
	s.LoadBalancerId = &v
	return s
}

type ListSecurityPolicyRelationsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityPolicyRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPolicyRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRelationsResponse) SetHeaders(v map[string]*string) *ListSecurityPolicyRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPolicyRelationsResponse) SetBody(v *ListSecurityPolicyRelationsResponseBody) *ListSecurityPolicyRelationsResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAttributeRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// 负载均衡修改保护相关信息
	ModificationProtectionConfig *UpdateLoadBalancerAttributeRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetModificationProtectionConfig(v *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) *UpdateLoadBalancerAttributeRequest {
	s.ModificationProtectionConfig = v
	return s
}

type UpdateLoadBalancerAttributeRequestModificationProtectionConfig struct {
	// 设置修改保护状态的原因
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 负载均衡修改保护状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetReason(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequestModificationProtectionConfig) SetStatus(v string) *UpdateLoadBalancerAttributeRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

type UpdateLoadBalancerAttributeResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetJobId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAttributeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetBody(v *UpdateLoadBalancerAttributeResponseBody) *UpdateLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type ListAsynJobsRequest struct {
	// 任务ID列表
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// 资源类型类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源实例ID列表
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// 操作接口名
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 任务过滤时间范围-开始时间
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 任务过滤时间范围-结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 下一个查询开始Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAsynJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsRequest) GoString() string {
	return s.String()
}

func (s *ListAsynJobsRequest) SetJobIds(v []*string) *ListAsynJobsRequest {
	s.JobIds = v
	return s
}

func (s *ListAsynJobsRequest) SetResourceType(v string) *ListAsynJobsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListAsynJobsRequest) SetResourceIds(v []*string) *ListAsynJobsRequest {
	s.ResourceIds = v
	return s
}

func (s *ListAsynJobsRequest) SetApiName(v string) *ListAsynJobsRequest {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsRequest) SetBeginTime(v int64) *ListAsynJobsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetEndTime(v int64) *ListAsynJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAsynJobsRequest) SetNextToken(v string) *ListAsynJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsRequest) SetMaxResults(v int64) *ListAsynJobsRequest {
	s.MaxResults = &v
	return s
}

type ListAsynJobsResponseBody struct {
	// 任务列表
	Jobs []*ListAsynJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// 本次查询返回记录数量
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAsynJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBody) SetJobs(v []*ListAsynJobsResponseBodyJobs) *ListAsynJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListAsynJobsResponseBody) SetMaxResults(v int64) *ListAsynJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetNextToken(v string) *ListAsynJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetRequestId(v string) *ListAsynJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsynJobsResponseBody) SetTotalCount(v int64) *ListAsynJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAsynJobsResponseBodyJobs struct {
	// openapi名称
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 任务开始时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 如果Status为失败，则为错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 如果Status为失败，则为错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 任务ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 任务结束时间戳
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 操作类型
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// 关联的资源实例ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 关联的资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAsynJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponseBodyJobs) SetApiName(v string) *ListAsynJobsResponseBodyJobs {
	s.ApiName = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetCreateTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorCode(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetErrorMessage(v string) *ListAsynJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetId(v string) *ListAsynJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetModifyTime(v int64) *ListAsynJobsResponseBodyJobs {
	s.ModifyTime = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetOperateType(v string) *ListAsynJobsResponseBodyJobs {
	s.OperateType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceId(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceId = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetResourceType(v string) *ListAsynJobsResponseBodyJobs {
	s.ResourceType = &v
	return s
}

func (s *ListAsynJobsResponseBodyJobs) SetStatus(v string) *ListAsynJobsResponseBodyJobs {
	s.Status = &v
	return s
}

type ListAsynJobsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAsynJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAsynJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsynJobsResponse) GoString() string {
	return s.String()
}

func (s *ListAsynJobsResponse) SetHeaders(v map[string]*string) *ListAsynJobsResponse {
	s.Headers = v
	return s
}

func (s *ListAsynJobsResponse) SetBody(v *ListAsynJobsResponseBody) *ListAsynJobsResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 转发规则优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 转发规则动作
	RuleActions []*CreateRuleRequestRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// 转发规则条件
	RuleConditions []*CreateRuleRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// 转发规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetListenerId(v string) *CreateRuleRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRuleRequest) SetClientToken(v string) *CreateRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRuleRequest) SetDryRun(v bool) *CreateRuleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRuleRequest) SetPriority(v int32) *CreateRuleRequest {
	s.Priority = &v
	return s
}

func (s *CreateRuleRequest) SetRuleActions(v []*CreateRuleRequestRuleActions) *CreateRuleRequest {
	s.RuleActions = v
	return s
}

func (s *CreateRuleRequest) SetRuleConditions(v []*CreateRuleRequestRuleConditions) *CreateRuleRequest {
	s.RuleConditions = v
	return s
}

func (s *CreateRuleRequest) SetRuleName(v string) *CreateRuleRequest {
	s.RuleName = &v
	return s
}

type CreateRuleRequestRuleActions struct {
	// 返回固定内容动作配置
	FixedResponseConfig *CreateRuleRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// 转发组动作配置
	ForwardGroupConfig *CreateRuleRequestRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 插入头部动作配置
	InsertHeaderConfig *CreateRuleRequestRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// 优先级
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 重定向动作配置
	RedirectConfig *CreateRuleRequestRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// 内部重定向动作配置
	RewriteConfig *CreateRuleRequestRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// 转发规则动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActions) SetFixedResponseConfig(v *CreateRuleRequestRuleActionsFixedResponseConfig) *CreateRuleRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetForwardGroupConfig(v *CreateRuleRequestRuleActionsForwardGroupConfig) *CreateRuleRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetInsertHeaderConfig(v *CreateRuleRequestRuleActionsInsertHeaderConfig) *CreateRuleRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetOrder(v int32) *CreateRuleRequestRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRedirectConfig(v *CreateRuleRequestRuleActionsRedirectConfig) *CreateRuleRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetRewriteConfig(v *CreateRuleRequestRuleActionsRewriteConfig) *CreateRuleRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRuleRequestRuleActions) SetType(v string) *CreateRuleRequestRuleActions {
	s.Type = &v
	return s
}

type CreateRuleRequestRuleActionsFixedResponseConfig struct {
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 内容类型
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// HTTP响应码
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s CreateRuleRequestRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContent(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRuleRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type CreateRuleRequestRuleActionsForwardGroupConfig struct {
	// 转发到的目的服务器组列表
	ServerGroupTuples []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) *CreateRuleRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组标识
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRuleRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateRuleRequestRuleActionsInsertHeaderConfig struct {
	// HTTP标头
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 取值类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRuleRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRuleRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type CreateRuleRequestRuleActionsRedirectConfig struct {
	// 要跳转的主机地址
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 跳转方式
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 要跳转的路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 要跳转的端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 要跳转的协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 要跳转的查询字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHost(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPath(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetPort(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetProtocol(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRedirectConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type CreateRuleRequestRuleActionsRewriteConfig struct {
	// 主机名
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 查询
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRuleRequestRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetHost(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetPath(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRuleRequestRuleActionsRewriteConfig) SetQuery(v string) *CreateRuleRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type CreateRuleRequestRuleConditions struct {
	// Cookie条件配置
	CookieConfig *CreateRuleRequestRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// HTTP标头条件配置
	HeaderConfig *CreateRuleRequestRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// 主机名条件配置
	HostConfig *CreateRuleRequestRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// HTTP请求方法条件配置
	MethodConfig *CreateRuleRequestRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// 路径条件配置
	PathConfig *CreateRuleRequestRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// 查询字符串条件配置
	QueryStringConfig *CreateRuleRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// 条件类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRuleRequestRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditions) SetCookieConfig(v *CreateRuleRequestRuleConditionsCookieConfig) *CreateRuleRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHeaderConfig(v *CreateRuleRequestRuleConditionsHeaderConfig) *CreateRuleRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetHostConfig(v *CreateRuleRequestRuleConditionsHostConfig) *CreateRuleRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetMethodConfig(v *CreateRuleRequestRuleConditionsMethodConfig) *CreateRuleRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetPathConfig(v *CreateRuleRequestRuleConditionsPathConfig) *CreateRuleRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetQueryStringConfig(v *CreateRuleRequestRuleConditionsQueryStringConfig) *CreateRuleRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRuleRequestRuleConditions) SetType(v string) *CreateRuleRequestRuleConditions {
	s.Type = &v
	return s
}

type CreateRuleRequestRuleConditionsCookieConfig struct {
	// Cookie键值对列表
	Values []*CreateRuleRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfig) SetValues(v []*CreateRuleRequestRuleConditionsCookieConfigValues) *CreateRuleRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsCookieConfigValues struct {
	// Cookie条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Cookie条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsCookieConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type CreateRuleRequestRuleConditionsHeaderConfig struct {
	// HTTP标头键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头值列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetKey(v string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsHostConfig struct {
	// 主机名列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsHostConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsMethodConfig struct {
	// HTTP请求方法列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsMethodConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsPathConfig struct {
	// 路径条件列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsPathConfig) SetValues(v []*string) *CreateRuleRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsQueryStringConfig struct {
	// 查询字符串条件键值对列表
	Values []*CreateRuleRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfig) SetValues(v []*CreateRuleRequestRuleConditionsQueryStringConfigValues) *CreateRuleRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type CreateRuleRequestRuleConditionsQueryStringConfigValues struct {
	// 查询字符串条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 查询字符串条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRuleRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRuleRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type CreateRuleResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 转发规则标识
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetJobId(v string) *CreateRuleResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRuleId(v string) *CreateRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type RemoveServersFromServerGroupRequest struct {
	// 后端服务器Id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 后端服务器
	Servers []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

type RemoveServersFromServerGroupRequestServers struct {
	// 后端端口号
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 后端服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 后端服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 后端服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s RemoveServersFromServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequestServers) SetPort(v int32) *RemoveServersFromServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerId(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerIp(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerType(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerType = &v
	return s
}

type RemoveServersFromServerGroupResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) SetJobId(v string) *RemoveServersFromServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveServersFromServerGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveServersFromServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveServersFromServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponse) SetHeaders(v map[string]*string) *RemoveServersFromServerGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetBody(v *RemoveServersFromServerGroupResponseBody) *RemoveServersFromServerGroupResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// 语言
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// Region列表
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// 名称
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// endpoint
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// RegionId
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type CreateSecurityPolicyRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// tls版本
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 安全策略名称
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// 加密套件
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
}

func (s CreateSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequest) SetClientToken(v string) *CreateSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetDryRun(v bool) *CreateSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetTLSVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TLSVersions = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetResourceGroupId(v string) *CreateSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetSecurityPolicyName(v string) *CreateSecurityPolicyRequest {
	s.SecurityPolicyName = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetCiphers(v []*string) *CreateSecurityPolicyRequest {
	s.Ciphers = v
	return s
}

type CreateSecurityPolicyResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 安全策略id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s CreateSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBody) SetRequestId(v string) *CreateSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBody {
	s.SecurityPolicyId = &v
	return s
}

type CreateSecurityPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponse) SetHeaders(v map[string]*string) *CreateSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityPolicyResponse) SetBody(v *CreateSecurityPolicyResponseBody) *CreateSecurityPolicyResponse {
	s.Body = v
	return s
}

type UpdateRuleAttributeRequest struct {
	// 转发规则标识
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 转发规则条件
	RuleConditions []*UpdateRuleAttributeRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// 转发规则动作
	RuleActions []*UpdateRuleAttributeRequestRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// 转发规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 转发规则优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequest) SetRuleId(v string) *UpdateRuleAttributeRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetClientToken(v string) *UpdateRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetDryRun(v bool) *UpdateRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleConditions(v []*UpdateRuleAttributeRequestRuleConditions) *UpdateRuleAttributeRequest {
	s.RuleConditions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleActions(v []*UpdateRuleAttributeRequestRuleActions) *UpdateRuleAttributeRequest {
	s.RuleActions = v
	return s
}

func (s *UpdateRuleAttributeRequest) SetRuleName(v string) *UpdateRuleAttributeRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRuleAttributeRequest) SetPriority(v int32) *UpdateRuleAttributeRequest {
	s.Priority = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditions struct {
	// Cookie条件配置
	CookieConfig *UpdateRuleAttributeRequestRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// HTTP标头条件配置
	HeaderConfig *UpdateRuleAttributeRequestRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// 主机名条件配置
	HostConfig *UpdateRuleAttributeRequestRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// HTTP请求方法条件配置
	MethodConfig *UpdateRuleAttributeRequestRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// 路径条件配置
	PathConfig *UpdateRuleAttributeRequestRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// 查询字符串条件配置
	QueryStringConfig *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// 条件类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetCookieConfig(v *UpdateRuleAttributeRequestRuleConditionsCookieConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHeaderConfig(v *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetHostConfig(v *UpdateRuleAttributeRequestRuleConditionsHostConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.HostConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetMethodConfig(v *UpdateRuleAttributeRequestRuleConditionsMethodConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetPathConfig(v *UpdateRuleAttributeRequestRuleConditionsPathConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.PathConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetQueryStringConfig(v *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) *UpdateRuleAttributeRequestRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditions) SetType(v string) *UpdateRuleAttributeRequestRuleConditions {
	s.Type = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfig struct {
	// Cookie键值对列表
	Values []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) *UpdateRuleAttributeRequestRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsCookieConfigValues struct {
	// Cookie条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Cookie条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsHeaderConfig struct {
	// HTTP标头键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头值列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsHeaderConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsHostConfig struct {
	// 主机名列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsHostConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsHostConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsMethodConfig struct {
	// HTTP请求方法列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsMethodConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsPathConfig struct {
	// 路径条件列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsPathConfig) SetValues(v []*string) *UpdateRuleAttributeRequestRuleConditionsPathConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfig struct {
	// 查询字符串条件键值对列表
	Values []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig) SetValues(v []*UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues struct {
	// 查询字符串条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 查询字符串条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetKey(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues) SetValue(v string) *UpdateRuleAttributeRequestRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type UpdateRuleAttributeRequestRuleActions struct {
	// 返回固定内容动作配置
	FixedResponseConfig *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// 转发组动作配置
	ForwardGroupConfig *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 插入头部动作配置
	InsertHeaderConfig *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// 优先级
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 重定向动作配置
	RedirectConfig *UpdateRuleAttributeRequestRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// 内部重定向动作配置
	RewriteConfig *UpdateRuleAttributeRequestRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// 转发规则动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActions) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActions) SetFixedResponseConfig(v *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) *UpdateRuleAttributeRequestRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetForwardGroupConfig(v *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) *UpdateRuleAttributeRequestRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetInsertHeaderConfig(v *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) *UpdateRuleAttributeRequestRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetOrder(v int32) *UpdateRuleAttributeRequestRuleActions {
	s.Order = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRedirectConfig(v *UpdateRuleAttributeRequestRuleActionsRedirectConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetRewriteConfig(v *UpdateRuleAttributeRequestRuleActionsRewriteConfig) *UpdateRuleAttributeRequestRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActions) SetType(v string) *UpdateRuleAttributeRequestRuleActions {
	s.Type = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsFixedResponseConfig struct {
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 内容类型
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// HTTP响应码
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContent(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetContentType(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfig struct {
	// 转发到的目的服务器组列表
	ServerGroupTuples []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组标识
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *UpdateRuleAttributeRequestRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig struct {
	// HTTP标头
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 取值类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetKey(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValue(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig) SetValueType(v string) *UpdateRuleAttributeRequestRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsRedirectConfig struct {
	// 要跳转的主机地址
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 跳转方式
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 要跳转的路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 要跳转的端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 要跳转的协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 要跳转的查询字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetHttpCode(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetPort(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetProtocol(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRedirectConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type UpdateRuleAttributeRequestRuleActionsRewriteConfig struct {
	// 主机名
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 查询
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeRequestRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetHost(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetPath(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *UpdateRuleAttributeRequestRuleActionsRewriteConfig) SetQuery(v string) *UpdateRuleAttributeRequestRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type UpdateRuleAttributeResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponseBody) SetJobId(v string) *UpdateRuleAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateRuleAttributeResponseBody) SetRequestId(v string) *UpdateRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRuleAttributeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleAttributeResponse) SetBody(v *UpdateRuleAttributeResponseBody) *UpdateRuleAttributeResponse {
	s.Body = v
	return s
}

type ListSecurityPoliciesRequest struct {
	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 安全策略名称
	SecurityPolicyNames []*string `json:"SecurityPolicyNames,omitempty" xml:"SecurityPolicyNames,omitempty" type:"Repeated"`
	// 安全策略id
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesRequest) SetResourceGroupId(v string) *ListSecurityPoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetNextToken(v string) *ListSecurityPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetMaxResults(v int32) *ListSecurityPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyNames(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyNames = v
	return s
}

func (s *ListSecurityPoliciesRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPoliciesRequest {
	s.SecurityPolicyIds = v
	return s
}

type ListSecurityPoliciesResponseBody struct {
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 安全策略
	SecurityPolicies []*ListSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBody) SetMaxResults(v int32) *ListSecurityPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetNextToken(v string) *ListSecurityPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetRequestId(v string) *ListSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSecurityPoliciesResponseBodySecurityPolicies) *ListSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSecurityPoliciesResponseBody) SetTotalCount(v int32) *ListSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecurityPoliciesResponseBodySecurityPolicies struct {
	// 加密套件
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 安全策略id
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// 安全策略名称
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// 状态
	SecurityPolicyStatus *string `json:"SecurityPolicyStatus,omitempty" xml:"SecurityPolicyStatus,omitempty"`
	// TLS策略
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetResourceGroupId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyName(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyStatus(v string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyStatus = &v
	return s
}

func (s *ListSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

type ListSecurityPoliciesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPoliciesResponse) SetBody(v *ListSecurityPoliciesResponseBody) *ListSecurityPoliciesResponse {
	s.Body = v
	return s
}

type UpdateAclAttributeRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// AclId
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// 访问控制策略名称
	AclName *string `json:"AclName,omitempty" xml:"AclName,omitempty"`
	// 是否预校验请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateAclAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeRequest) SetClientToken(v string) *UpdateAclAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclId(v string) *UpdateAclAttributeRequest {
	s.AclId = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetAclName(v string) *UpdateAclAttributeRequest {
	s.AclName = &v
	return s
}

func (s *UpdateAclAttributeRequest) SetDryRun(v bool) *UpdateAclAttributeRequest {
	s.DryRun = &v
	return s
}

type UpdateAclAttributeResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAclAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponseBody) SetRequestId(v string) *UpdateAclAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAclAttributeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAclAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponse) SetHeaders(v map[string]*string) *UpdateAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclAttributeResponse) SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse {
	s.Body = v
	return s
}

type DisableDeletionProtectionRequest struct {
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 实例id
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DisableDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionRequest) SetClientToken(v string) *DisableDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetDryRun(v bool) *DisableDeletionProtectionRequest {
	s.DryRun = &v
	return s
}

func (s *DisableDeletionProtectionRequest) SetResourceId(v string) *DisableDeletionProtectionRequest {
	s.ResourceId = &v
	return s
}

type DisableDeletionProtectionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionResponseBody) SetRequestId(v string) *DisableDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type DisableDeletionProtectionResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *DisableDeletionProtectionResponse) SetHeaders(v map[string]*string) *DisableDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *DisableDeletionProtectionResponse) SetBody(v *DisableDeletionProtectionResponseBody) *DisableDeletionProtectionResponse {
	s.Body = v
	return s
}

type StopListenerRequest struct {
	// 监听id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s StopListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StopListenerRequest) GoString() string {
	return s.String()
}

func (s *StopListenerRequest) SetListenerId(v string) *StopListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *StopListenerRequest) SetClientToken(v string) *StopListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StopListenerRequest) SetDryRun(v bool) *StopListenerRequest {
	s.DryRun = &v
	return s
}

type StopListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StopListenerResponseBody) SetJobId(v string) *StopListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StopListenerResponseBody) SetRequestId(v string) *StopListenerResponseBody {
	s.RequestId = &v
	return s
}

type StopListenerResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StopListenerResponse) GoString() string {
	return s.String()
}

func (s *StopListenerResponse) SetHeaders(v map[string]*string) *StopListenerResponse {
	s.Headers = v
	return s
}

func (s *StopListenerResponse) SetBody(v *StopListenerResponseBody) *StopListenerResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
	// 负载均衡标识
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 请求超时时间
	RequestTimeout *int32 `json:"RequestTimeout,omitempty" xml:"RequestTimeout,omitempty"`
	// 连接空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 是否开启Gzip压缩
	GzipEnabled *bool `json:"GzipEnabled,omitempty" xml:"GzipEnabled,omitempty"`
	// 是否开启HTTP/2特性
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// 安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// XForward字段相关的配置
	XForwardedForConfig *CreateListenerRequestXForwardedForConfig `json:"XForwardedForConfig,omitempty" xml:"XForwardedForConfig,omitempty" type:"Struct"`
	// HTTPS启用QUIC时相关属性
	QuicConfig *CreateListenerRequestQuicConfig `json:"QuicConfig,omitempty" xml:"QuicConfig,omitempty" type:"Struct"`
	// 监听默认服务器证书列表，N当前取值范围为1
	Certificates []*CreateListenerRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// 监听默认动作
	DefaultActions []*CreateListenerRequestDefaultActions `json:"DefaultActions,omitempty" xml:"DefaultActions,omitempty" type:"Repeated"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetListenerProtocol(v string) *CreateListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetListenerPort(v int32) *CreateListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetRequestTimeout(v int32) *CreateListenerRequest {
	s.RequestTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetGzipEnabled(v bool) *CreateListenerRequest {
	s.GzipEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetHttp2Enabled(v bool) *CreateListenerRequest {
	s.Http2Enabled = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetXForwardedForConfig(v *CreateListenerRequestXForwardedForConfig) *CreateListenerRequest {
	s.XForwardedForConfig = v
	return s
}

func (s *CreateListenerRequest) SetQuicConfig(v *CreateListenerRequestQuicConfig) *CreateListenerRequest {
	s.QuicConfig = v
	return s
}

func (s *CreateListenerRequest) SetCertificates(v []*CreateListenerRequestCertificates) *CreateListenerRequest {
	s.Certificates = v
	return s
}

func (s *CreateListenerRequest) SetDefaultActions(v []*CreateListenerRequestDefaultActions) *CreateListenerRequest {
	s.DefaultActions = v
	return s
}

type CreateListenerRequestXForwardedForConfig struct {
	// 自定义HEADER头名称，只有当XForwardedForClientCertClientVerifyEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertClientVerifyAlias *string `json:"XForwardedForClientCertClientVerifyAlias,omitempty" xml:"XForwardedForClientCertClientVerifyAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-clientverify  头字段获取对访问负载均衡实例客户端证书的校验结果。HTTPS监听有效。
	XForwardedForClientCertClientVerifyEnabled *bool `json:"XForwardedForClientCertClientVerifyEnabled,omitempty" xml:"XForwardedForClientCertClientVerifyEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertFingerprintEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertFingerprintAlias *string `json:"XForwardedForClientCertFingerprintAlias,omitempty" xml:"XForwardedForClientCertFingerprintAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-fingerprint 头字段获取访问负载均衡实例客户端证书的指纹取值，HTTPS监听有效。
	XForwardedForClientCertFingerprintEnabled *bool `json:"XForwardedForClientCertFingerprintEnabled,omitempty" xml:"XForwardedForClientCertFingerprintEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertIssuerDNEnabled的值为‘On’的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertIssuerDNAlias *string `json:"XForwardedForClientCertIssuerDNAlias,omitempty" xml:"XForwardedForClientCertIssuerDNAlias,omitempty"`
	// 是否通过 X-Forwarded-Clientcert-issuerdn 头字段获取访问负载均衡实例客户端证书的发行者信息。HTTPS监听有效。
	XForwardedForClientCertIssuerDNEnabled *bool `json:"XForwardedForClientCertIssuerDNEnabled,omitempty" xml:"XForwardedForClientCertIssuerDNEnabled,omitempty"`
	// 自定义HEADER头名称，只有当XForwardedForClientCertSubjectDNEnabled的值为true的时候，此值才会生效；否则该值不会生效。HTTPS监听有效
	XForwardedForClientCertSubjectDNAlias *string `json:"XForwardedForClientCertSubjectDNAlias,omitempty" xml:"XForwardedForClientCertSubjectDNAlias,omitempty"`
	// 是否通过X-Forwarded-Clientcert-subjectdn  头字段获取访问负载均衡实例客户端证书的所有者信息。HTTPS监听有效。
	XForwardedForClientCertSubjectDNEnabled *bool `json:"XForwardedForClientCertSubjectDNEnabled,omitempty" xml:"XForwardedForClientCertSubjectDNEnabled,omitempty"`
	// 是否通过X-Forwarded-Client-Port 头字段获取访问负载均衡实例客户端的端口。HTTPS监听有效。
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitempty" xml:"XForwardedForClientSrcPortEnabled,omitempty"`
	// 是否开启通过X-Forwarded-For头字段获取来访者真实 IP
	XForwardedForEnabled *bool `json:"XForwardedForEnabled,omitempty" xml:"XForwardedForEnabled,omitempty"`
	// 是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitempty" xml:"XForwardedForProtoEnabled,omitempty"`
	// 是否通过SLB-ID头字段获取负载均衡实例ID。
	XForwardedForSLBIdEnabled *bool `json:"XForwardedForSLBIdEnabled,omitempty" xml:"XForwardedForSLBIdEnabled,omitempty"`
	// 是否通过X-Forwarded-Port 头字段获取负载均衡实例的监听端口。HTTPS监听有效。
	XForwardedForSLBPortEnabled *bool `json:"XForwardedForSLBPortEnabled,omitempty" xml:"XForwardedForSLBPortEnabled,omitempty"`
}

func (s CreateListenerRequestXForwardedForConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestXForwardedForConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertClientVerifyEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertClientVerifyEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertFingerprintEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertFingerprintEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertIssuerDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertIssuerDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNAlias(v string) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNAlias = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientCertSubjectDNEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientCertSubjectDNEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForClientSrcPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForClientSrcPortEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForProtoEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForProtoEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBIdEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBIdEnabled = &v
	return s
}

func (s *CreateListenerRequestXForwardedForConfig) SetXForwardedForSLBPortEnabled(v bool) *CreateListenerRequestXForwardedForConfig {
	s.XForwardedForSLBPortEnabled = &v
	return s
}

type CreateListenerRequestQuicConfig struct {
	// 需要关联的QUIC监听ID，HTTPS监听时有效，QuicUpgradeEnabled为true时必选
	QuicListenerId *string `json:"QuicListenerId,omitempty" xml:"QuicListenerId,omitempty"`
	// 是否开启quic升级，HTTPS监听时有效
	QuicUpgradeEnabled *bool `json:"QuicUpgradeEnabled,omitempty" xml:"QuicUpgradeEnabled,omitempty"`
}

func (s CreateListenerRequestQuicConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestQuicConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestQuicConfig) SetQuicListenerId(v string) *CreateListenerRequestQuicConfig {
	s.QuicListenerId = &v
	return s
}

func (s *CreateListenerRequestQuicConfig) SetQuicUpgradeEnabled(v bool) *CreateListenerRequestQuicConfig {
	s.QuicUpgradeEnabled = &v
	return s
}

type CreateListenerRequestCertificates struct {
	// 正式标识
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CreateListenerRequestCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestCertificates) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestCertificates) SetCertificateId(v string) *CreateListenerRequestCertificates {
	s.CertificateId = &v
	return s
}

type CreateListenerRequestDefaultActions struct {
	// 转发组
	ForwardGroupConfig *CreateListenerRequestDefaultActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateListenerRequestDefaultActions) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActions) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActions) SetForwardGroupConfig(v *CreateListenerRequestDefaultActionsForwardGroupConfig) *CreateListenerRequestDefaultActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateListenerRequestDefaultActions) SetType(v string) *CreateListenerRequestDefaultActions {
	s.Type = &v
	return s
}

type CreateListenerRequestDefaultActionsForwardGroupConfig struct {
	// 服务器组列表
	ServerGroupTuples []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) *CreateListenerRequestDefaultActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateListenerRequestDefaultActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateListenerResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetJobId(v string) *CreateListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateListenerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateListenerResponse) SetHeaders(v map[string]*string) *CreateListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源实例Id
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// 标签列表
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UpdateServerGroupServersAttributeRequest struct {
	// 后端服务器Id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 后端服务器
	Servers []*UpdateServerGroupServersAttributeRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// dryRun
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s UpdateServerGroupServersAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest {
	s.Servers = v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetClientToken(v string) *UpdateServerGroupServersAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest {
	s.DryRun = &v
	return s
}

type UpdateServerGroupServersAttributeRequestServers struct {
	// 后端服务器描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 后端端口号
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 后端服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 后端服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 后端服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 后端服务器权重
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateServerGroupServersAttributeRequestServers) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequestServers) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetDescription(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.Description = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetPort(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Port = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerId(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerIp(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerIp = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetServerType(v string) *UpdateServerGroupServersAttributeRequestServers {
	s.ServerType = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequestServers) SetWeight(v int32) *UpdateServerGroupServersAttributeRequestServers {
	s.Weight = &v
	return s
}

type UpdateServerGroupServersAttributeResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupServersAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetJobId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServerGroupServersAttributeResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServerGroupServersAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServerGroupServersAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupServersAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupServersAttributeResponse) SetBody(v *UpdateServerGroupServersAttributeResponseBody) *UpdateServerGroupServersAttributeResponse {
	s.Body = v
	return s
}

type StartListenerRequest struct {
	// 监听id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s StartListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s StartListenerRequest) GoString() string {
	return s.String()
}

func (s *StartListenerRequest) SetListenerId(v string) *StartListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *StartListenerRequest) SetClientToken(v string) *StartListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *StartListenerRequest) SetDryRun(v bool) *StartListenerRequest {
	s.DryRun = &v
	return s
}

type StartListenerResponseBody struct {
	// 异步任务id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponseBody) GoString() string {
	return s.String()
}

func (s *StartListenerResponseBody) SetJobId(v string) *StartListenerResponseBody {
	s.JobId = &v
	return s
}

func (s *StartListenerResponseBody) SetRequestId(v string) *StartListenerResponseBody {
	s.RequestId = &v
	return s
}

type StartListenerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartListenerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartListenerResponse) GoString() string {
	return s.String()
}

func (s *StartListenerResponse) SetHeaders(v map[string]*string) *StartListenerResponse {
	s.Headers = v
	return s
}

func (s *StartListenerResponse) SetBody(v *StartListenerResponseBody) *StartListenerResponse {
	s.Body = v
	return s
}

type CreateRulesRequest struct {
	// 监听标识
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 幂等标识
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	//  是否只预检此次请求
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 转发规则列表
	Rules []*CreateRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateRulesRequest) SetListenerId(v string) *CreateRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateRulesRequest) SetClientToken(v string) *CreateRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRulesRequest) SetDryRun(v bool) *CreateRulesRequest {
	s.DryRun = &v
	return s
}

func (s *CreateRulesRequest) SetRules(v []*CreateRulesRequestRules) *CreateRulesRequest {
	s.Rules = v
	return s
}

type CreateRulesRequestRules struct {
	// 转发规则优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 转发规则动作
	RuleActions []*CreateRulesRequestRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// 转发规则条件
	RuleConditions []*CreateRulesRequestRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// 转发规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateRulesRequestRules) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRules) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRules) SetPriority(v int32) *CreateRulesRequestRules {
	s.Priority = &v
	return s
}

func (s *CreateRulesRequestRules) SetRuleActions(v []*CreateRulesRequestRulesRuleActions) *CreateRulesRequestRules {
	s.RuleActions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleConditions(v []*CreateRulesRequestRulesRuleConditions) *CreateRulesRequestRules {
	s.RuleConditions = v
	return s
}

func (s *CreateRulesRequestRules) SetRuleName(v string) *CreateRulesRequestRules {
	s.RuleName = &v
	return s
}

type CreateRulesRequestRulesRuleActions struct {
	// 返回固定内容动作配置
	FixedResponseConfig *CreateRulesRequestRulesRuleActionsFixedResponseConfig `json:"FixedResponseConfig,omitempty" xml:"FixedResponseConfig,omitempty" type:"Struct"`
	// 转发组动作配置
	ForwardGroupConfig *CreateRulesRequestRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// 插入头部动作配置
	InsertHeaderConfig *CreateRulesRequestRulesRuleActionsInsertHeaderConfig `json:"InsertHeaderConfig,omitempty" xml:"InsertHeaderConfig,omitempty" type:"Struct"`
	// 优先级
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// 重定向动作配置
	RedirectConfig *CreateRulesRequestRulesRuleActionsRedirectConfig `json:"RedirectConfig,omitempty" xml:"RedirectConfig,omitempty" type:"Struct"`
	// 内部重定向动作配置
	RewriteConfig *CreateRulesRequestRulesRuleActionsRewriteConfig `json:"RewriteConfig,omitempty" xml:"RewriteConfig,omitempty" type:"Struct"`
	// 转发规则动作类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleActions) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActions) SetFixedResponseConfig(v *CreateRulesRequestRulesRuleActionsFixedResponseConfig) *CreateRulesRequestRulesRuleActions {
	s.FixedResponseConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetForwardGroupConfig(v *CreateRulesRequestRulesRuleActionsForwardGroupConfig) *CreateRulesRequestRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetInsertHeaderConfig(v *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) *CreateRulesRequestRulesRuleActions {
	s.InsertHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetOrder(v int32) *CreateRulesRequestRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRedirectConfig(v *CreateRulesRequestRulesRuleActionsRedirectConfig) *CreateRulesRequestRulesRuleActions {
	s.RedirectConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetRewriteConfig(v *CreateRulesRequestRulesRuleActionsRewriteConfig) *CreateRulesRequestRulesRuleActions {
	s.RewriteConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleActions) SetType(v string) *CreateRulesRequestRulesRuleActions {
	s.Type = &v
	return s
}

type CreateRulesRequestRulesRuleActionsFixedResponseConfig struct {
	// 内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 内容类型
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// HTTP响应码
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsFixedResponseConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContent(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.Content = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetContentType(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.ContentType = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsFixedResponseConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsFixedResponseConfig {
	s.HttpCode = &v
	return s
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfig struct {
	// 转发到的目的服务器组列表
	ServerGroupTuples []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateRulesRequestRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

type CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// 服务器组标识
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples) SetServerGroupId(v string) *CreateRulesRequestRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.ServerGroupId = &v
	return s
}

type CreateRulesRequestRulesRuleActionsInsertHeaderConfig struct {
	// HTTP标头
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头内容
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// 取值类型
	ValueType *string `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsInsertHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValue(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.Value = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsInsertHeaderConfig) SetValueType(v string) *CreateRulesRequestRulesRuleActionsInsertHeaderConfig {
	s.ValueType = &v
	return s
}

type CreateRulesRequestRulesRuleActionsRedirectConfig struct {
	// 要跳转的主机地址
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 跳转方式
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 要跳转的路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 要跳转的端口
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 要跳转的协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 要跳转的查询字符串
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRedirectConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetHttpCode(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.HttpCode = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetPort(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Port = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetProtocol(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Protocol = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRedirectConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRedirectConfig {
	s.Query = &v
	return s
}

type CreateRulesRequestRulesRuleActionsRewriteConfig struct {
	// 主机名
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 查询
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleActionsRewriteConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetHost(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Host = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetPath(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Path = &v
	return s
}

func (s *CreateRulesRequestRulesRuleActionsRewriteConfig) SetQuery(v string) *CreateRulesRequestRulesRuleActionsRewriteConfig {
	s.Query = &v
	return s
}

type CreateRulesRequestRulesRuleConditions struct {
	// Cookie条件配置
	CookieConfig *CreateRulesRequestRulesRuleConditionsCookieConfig `json:"CookieConfig,omitempty" xml:"CookieConfig,omitempty" type:"Struct"`
	// HTTP标头条件配置
	HeaderConfig *CreateRulesRequestRulesRuleConditionsHeaderConfig `json:"HeaderConfig,omitempty" xml:"HeaderConfig,omitempty" type:"Struct"`
	// 主机名条件配置
	HostConfig *CreateRulesRequestRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	// HTTP请求方法条件配置
	MethodConfig *CreateRulesRequestRulesRuleConditionsMethodConfig `json:"MethodConfig,omitempty" xml:"MethodConfig,omitempty" type:"Struct"`
	// 路径条件配置
	PathConfig *CreateRulesRequestRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	// 查询字符串条件配置
	QueryStringConfig *CreateRulesRequestRulesRuleConditionsQueryStringConfig `json:"QueryStringConfig,omitempty" xml:"QueryStringConfig,omitempty" type:"Struct"`
	// 返回HTTP标头
	ResponseHeaderConfig *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig `json:"ResponseHeaderConfig,omitempty" xml:"ResponseHeaderConfig,omitempty" type:"Struct"`
	// 条件类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditions) SetCookieConfig(v *CreateRulesRequestRulesRuleConditionsCookieConfig) *CreateRulesRequestRulesRuleConditions {
	s.CookieConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHeaderConfig(v *CreateRulesRequestRulesRuleConditionsHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.HeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetHostConfig(v *CreateRulesRequestRulesRuleConditionsHostConfig) *CreateRulesRequestRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetMethodConfig(v *CreateRulesRequestRulesRuleConditionsMethodConfig) *CreateRulesRequestRulesRuleConditions {
	s.MethodConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetPathConfig(v *CreateRulesRequestRulesRuleConditionsPathConfig) *CreateRulesRequestRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetQueryStringConfig(v *CreateRulesRequestRulesRuleConditionsQueryStringConfig) *CreateRulesRequestRulesRuleConditions {
	s.QueryStringConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetResponseHeaderConfig(v *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) *CreateRulesRequestRulesRuleConditions {
	s.ResponseHeaderConfig = v
	return s
}

func (s *CreateRulesRequestRulesRuleConditions) SetType(v string) *CreateRulesRequestRulesRuleConditions {
	s.Type = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsCookieConfig struct {
	// Cookie键值对列表
	Values []*CreateRulesRequestRulesRuleConditionsCookieConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsCookieConfigValues) *CreateRulesRequestRulesRuleConditionsCookieConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsCookieConfigValues struct {
	// Cookie条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Cookie条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsCookieConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsCookieConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsCookieConfigValues {
	s.Value = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsHeaderConfig struct {
	// HTTP标头键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// HTTP标头值列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHeaderConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsHostConfig struct {
	// 主机名列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsMethodConfig struct {
	// HTTP请求方法列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsMethodConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsMethodConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsMethodConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsPathConfig struct {
	// 路径条件列表
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsQueryStringConfig struct {
	// 查询字符串条件键值对列表
	Values []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfig) SetValues(v []*CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) *CreateRulesRequestRulesRuleConditionsQueryStringConfig {
	s.Values = v
	return s
}

type CreateRulesRequestRulesRuleConditionsQueryStringConfigValues struct {
	// 查询字符串条件键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 查询字符串条件值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetKey(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues) SetValue(v string) *CreateRulesRequestRulesRuleConditionsQueryStringConfigValues {
	s.Value = &v
	return s
}

type CreateRulesRequestRulesRuleConditionsResponseHeaderConfig struct {
	// 返回HTTP标头键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 返回HTTP标头值
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) GoString() string {
	return s.String()
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetKey(v string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Key = &v
	return s
}

func (s *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig) SetValues(v []*string) *CreateRulesRequestRulesRuleConditionsResponseHeaderConfig {
	s.Values = v
	return s
}

type CreateRulesResponseBody struct {
	// 异步任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 转发规则标识列表
	RuleIds []*CreateRulesResponseBodyRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s CreateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBody) SetJobId(v string) *CreateRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRequestId(v string) *CreateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRulesResponseBody) SetRuleIds(v []*CreateRulesResponseBodyRuleIds) *CreateRulesResponseBody {
	s.RuleIds = v
	return s
}

type CreateRulesResponseBodyRuleIds struct {
	// 转发规则标识
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRulesResponseBodyRuleIds) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponseBodyRuleIds) GoString() string {
	return s.String()
}

func (s *CreateRulesResponseBodyRuleIds) SetRuleId(v string) *CreateRulesResponseBodyRuleIds {
	s.RuleId = &v
	return s
}

type CreateRulesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateRulesResponse) SetHeaders(v map[string]*string) *CreateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateRulesResponse) SetBody(v *CreateRulesResponseBody) *CreateRulesResponse {
	s.Body = v
	return s
}

type ListAclEntriesRequest struct {
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 访问控制策略Id
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
}

func (s ListAclEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListAclEntriesRequest) SetNextToken(v string) *ListAclEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAclEntriesRequest) SetMaxResults(v int32) *ListAclEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesRequest) SetAclId(v string) *ListAclEntriesRequest {
	s.AclId = &v
	return s
}

type ListAclEntriesResponseBody struct {
	// 访问控制列表
	AclEntries []*ListAclEntriesResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// 本次查询返回记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 分页查询标识
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAclEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBody) SetAclEntries(v []*ListAclEntriesResponseBodyAclEntries) *ListAclEntriesResponseBody {
	s.AclEntries = v
	return s
}

func (s *ListAclEntriesResponseBody) SetMaxResults(v int32) *ListAclEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetNextToken(v string) *ListAclEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetRequestId(v string) *ListAclEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclEntriesResponseBody) SetTotalCount(v int32) *ListAclEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAclEntriesResponseBodyAclEntries struct {
	// 描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IP条目
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAclEntriesResponseBodyAclEntries) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponseBodyAclEntries) SetDescription(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Description = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetEntry(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Entry = &v
	return s
}

func (s *ListAclEntriesResponseBodyAclEntries) SetStatus(v string) *ListAclEntriesResponseBodyAclEntries {
	s.Status = &v
	return s
}

type ListAclEntriesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAclEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListAclEntriesResponse) SetHeaders(v map[string]*string) *ListAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListAclEntriesResponse) SetBody(v *ListAclEntriesResponseBody) *ListAclEntriesResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-beijing":     tea.String("alb.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("alb.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":    tea.String("alb.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":    tea.String("alb.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    tea.String("alb.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":    tea.String("alb.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1": tea.String("alb.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": tea.String("alb.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-5": tea.String("alb.ap-southeast-5.aliyuncs.com"),
		"us-east-1":      tea.String("alb.us-east-1.aliyuncs.com"),
		"eu-central-1":   tea.String("alb.eu-central-1.aliyuncs.com"),
		"ap-south-1":     tea.String("alb.ap-south-1.aliyuncs.com"),
		"ap-northeast-1": tea.String("alb.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-3": tea.String("alb.ap-southeast-3.aliyuncs.com"),
		"cn-chengdu":     tea.String("alb.cn-chengdu.aliyuncs.com"),
		"cn-huhehaote":   tea.String("alb.cn-huhehaote.aliyuncs.com"),
		"cn-qingdao":     tea.String("alb.cn-qingdao.aliyuncs.com"),
		"cn-wulanchabu":  tea.String("alb.cn-wulanchabu.aliyuncs.com"),
		"eu-west-1":      tea.String("alb.eu-west-1.aliyuncs.com"),
		"us-west-1":      tea.String("alb.us-west-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListAclsWithOptions(request *ListAclsRequest, runtime *util.RuntimeOptions) (_result *ListAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAclsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAcls"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAcls(request *ListAclsRequest) (_result *ListAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclsResponse{}
	_body, _err := client.ListAclsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHealthCheckTemplateWithOptions(request *CreateHealthCheckTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateHealthCheckTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHealthCheckTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHealthCheckTemplate"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHealthCheckTemplate(request *CreateHealthCheckTemplateRequest) (_result *CreateHealthCheckTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHealthCheckTemplateResponse{}
	_body, _err := client.CreateHealthCheckTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetListenerAttributeWithOptions(request *GetListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetListenerAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetListenerAttribute(request *GetListenerAttributeRequest) (_result *GetListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.GetListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateHealthCheckTemplateAttributeWithOptions(request *UpdateHealthCheckTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateHealthCheckTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateHealthCheckTemplateAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateHealthCheckTemplateAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateHealthCheckTemplateAttribute(request *UpdateHealthCheckTemplateAttributeRequest) (_result *UpdateHealthCheckTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateHealthCheckTemplateAttributeResponse{}
	_body, _err := client.UpdateHealthCheckTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupServersWithOptions(request *ListServerGroupServersRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServerGroupServers"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroupServers(request *ListServerGroupServersRequest) (_result *ListServerGroupServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.ListServerGroupServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEntriesFromAclWithOptions(request *RemoveEntriesFromAclRequest, runtime *util.RuntimeOptions) (_result *RemoveEntriesFromAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveEntriesFromAcl"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEntriesFromAcl(request *RemoveEntriesFromAclRequest) (_result *RemoveEntriesFromAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEntriesFromAclResponse{}
	_body, _err := client.RemoveEntriesFromAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityPolicyWithOptions(request *DeleteSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecurityPolicy"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityPolicy(request *DeleteSecurityPolicyRequest) (_result *DeleteSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.DeleteSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLoadBalancerAccessLogWithOptions(request *DisableLoadBalancerAccessLogRequest, runtime *util.RuntimeOptions) (_result *DisableLoadBalancerAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableLoadBalancerAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableLoadBalancerAccessLog"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLoadBalancerAccessLog(request *DisableLoadBalancerAccessLogRequest) (_result *DisableLoadBalancerAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLoadBalancerAccessLogResponse{}
	_body, _err := client.DisableLoadBalancerAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServerGroupsWithOptions(request *ListServerGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServerGroups"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServerGroups(request *ListServerGroupsRequest) (_result *ListServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.ListServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableDeletionProtectionWithOptions(request *EnableDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *EnableDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableDeletionProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableDeletionProtection"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableDeletionProtection(request *EnableDeletionProtectionRequest) (_result *EnableDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDeletionProtectionResponse{}
	_body, _err := client.EnableDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateListenerAttributeWithOptions(request *UpdateListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateListenerAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateListenerAttribute(request *UpdateListenerAttributeRequest) (_result *UpdateListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.UpdateListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLoadBalancer"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHealthCheckTemplatesWithOptions(request *DeleteHealthCheckTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteHealthCheckTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHealthCheckTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHealthCheckTemplates"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHealthCheckTemplates(request *DeleteHealthCheckTemplatesRequest) (_result *DeleteHealthCheckTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHealthCheckTemplatesResponse{}
	_body, _err := client.DeleteHealthCheckTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServerGroupAttributeWithOptions(request *UpdateServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServerGroupAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServerGroupAttribute(request *UpdateServerGroupAttributeRequest) (_result *UpdateServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.UpdateServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServerGroupWithOptions(request *CreateServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServerGroup"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServerGroup(request *CreateServerGroupRequest) (_result *CreateServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CreateServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateAclsWithListenerWithOptions(request *AssociateAclsWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAclsWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateAclsWithListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAclsWithListener(request *AssociateAclsWithListenerRequest) (_result *AssociateAclsWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAclsWithListenerResponse{}
	_body, _err := client.AssociateAclsWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRule"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRules"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeZones"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones() (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServerGroupWithOptions(request *DeleteServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteServerGroup"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServerGroup(request *DeleteServerGroupRequest) (_result *DeleteServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.DeleteServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddServersToServerGroupWithOptions(request *AddServersToServerGroupRequest, runtime *util.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddServersToServerGroup"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddServersToServerGroup(request *AddServersToServerGroupRequest) (_result *AddServersToServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.AddServersToServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLoadBalancer"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateListenerLogConfigWithOptions(request *UpdateListenerLogConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateListenerLogConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateListenerLogConfig"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateListenerLogConfig(request *UpdateListenerLogConfigRequest) (_result *UpdateListenerLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerLogConfigResponse{}
	_body, _err := client.UpdateListenerLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnTagResources"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSystemSecurityPoliciesWithOptions(runtime *util.RuntimeOptions) (_result *ListSystemSecurityPoliciesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSystemSecurityPolicies"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSystemSecurityPolicies() (_result *ListSystemSecurityPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSystemSecurityPoliciesResponse{}
	_body, _err := client.ListSystemSecurityPoliciesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEntriesToAclWithOptions(request *AddEntriesToAclRequest, runtime *util.RuntimeOptions) (_result *AddEntriesToAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEntriesToAcl"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEntriesToAcl(request *AddEntriesToAclRequest) (_result *AddEntriesToAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEntriesToAclResponse{}
	_body, _err := client.AddEntriesToAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAcl"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAcl(request *DeleteAclRequest) (_result *DeleteAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateAdditionalCertificatesFromListenerWithOptions(request *DissociateAdditionalCertificatesFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateAdditionalCertificatesFromListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateAdditionalCertificatesFromListener(request *DissociateAdditionalCertificatesFromListenerRequest) (_result *DissociateAdditionalCertificatesFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAdditionalCertificatesFromListenerResponse{}
	_body, _err := client.DissociateAdditionalCertificatesFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoadBalancerEditionWithOptions(request *UpdateLoadBalancerEditionRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerEditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLoadBalancerEditionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLoadBalancerEdition"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoadBalancerEdition(request *UpdateLoadBalancerEditionRequest) (_result *UpdateLoadBalancerEditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerEditionResponse{}
	_body, _err := client.UpdateLoadBalancerEditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListListeners"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclRelationsWithOptions(request *ListAclRelationsRequest, runtime *util.RuntimeOptions) (_result *ListAclRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAclRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAclRelations"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAclRelations(request *ListAclRelationsRequest) (_result *ListAclRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclRelationsResponse{}
	_body, _err := client.ListAclRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRulesWithOptions(request *DeleteRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRules"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRules(request *DeleteRulesRequest) (_result *DeleteRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRulesResponse{}
	_body, _err := client.DeleteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateAdditionalCertificatesWithListenerWithOptions(request *AssociateAdditionalCertificatesWithListenerRequest, runtime *util.RuntimeOptions) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssociateAdditionalCertificatesWithListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateAdditionalCertificatesWithListener(request *AssociateAdditionalCertificatesWithListenerRequest) (_result *AssociateAdditionalCertificatesWithListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateAdditionalCertificatesWithListenerResponse{}
	_body, _err := client.AssociateAdditionalCertificatesWithListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHealthCheckTemplatesWithOptions(request *ListHealthCheckTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListHealthCheckTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHealthCheckTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHealthCheckTemplates"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHealthCheckTemplates(request *ListHealthCheckTemplatesRequest) (_result *ListHealthCheckTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHealthCheckTemplatesResponse{}
	_body, _err := client.ListHealthCheckTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoadBalancerAttributeWithOptions(request *GetLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLoadBalancerAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoadBalancerAttribute(request *GetLoadBalancerAttributeRequest) (_result *GetLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.GetLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLoadBalancers"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (_result *ListLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.ListLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLoadBalancerAccessLogWithOptions(request *EnableLoadBalancerAccessLogRequest, runtime *util.RuntimeOptions) (_result *EnableLoadBalancerAccessLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableLoadBalancerAccessLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableLoadBalancerAccessLog"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLoadBalancerAccessLog(request *EnableLoadBalancerAccessLogRequest) (_result *EnableLoadBalancerAccessLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLoadBalancerAccessLogResponse{}
	_body, _err := client.EnableLoadBalancerAccessLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DissociateAclsFromListenerWithOptions(request *DissociateAclsFromListenerRequest, runtime *util.RuntimeOptions) (_result *DissociateAclsFromListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DissociateAclsFromListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DissociateAclsFromListener(request *DissociateAclsFromListenerRequest) (_result *DissociateAclsFromListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateAclsFromListenerResponse{}
	_body, _err := client.DissociateAclsFromListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListListenerCertificatesWithOptions(request *ListListenerCertificatesRequest, runtime *util.RuntimeOptions) (_result *ListListenerCertificatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListListenerCertificates"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (_result *ListListenerCertificatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenerCertificatesResponse{}
	_body, _err := client.ListListenerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *util.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAcl"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAcl(request *CreateAclRequest) (_result *CreateAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAclResponse{}
	_body, _err := client.CreateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHealthCheckTemplateAttributeWithOptions(request *GetHealthCheckTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *GetHealthCheckTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHealthCheckTemplateAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHealthCheckTemplateAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHealthCheckTemplateAttribute(request *GetHealthCheckTemplateAttributeRequest) (_result *GetHealthCheckTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHealthCheckTemplateAttributeResponse{}
	_body, _err := client.GetHealthCheckTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSecurityPolicyAttributeWithOptions(request *UpdateSecurityPolicyAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSecurityPolicyAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSecurityPolicyAttribute(request *UpdateSecurityPolicyAttributeRequest) (_result *UpdateSecurityPolicyAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.UpdateSecurityPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityPolicyRelationsWithOptions(request *ListSecurityPolicyRelationsRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPolicyRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecurityPolicyRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecurityPolicyRelations"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicyRelations(request *ListSecurityPolicyRelationsRequest) (_result *ListSecurityPolicyRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPolicyRelationsResponse{}
	_body, _err := client.ListSecurityPolicyRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoadBalancerAttributeWithOptions(request *UpdateLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLoadBalancerAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoadBalancerAttribute(request *UpdateLoadBalancerAttributeRequest) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.UpdateLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsynJobsWithOptions(request *ListAsynJobsRequest, runtime *util.RuntimeOptions) (_result *ListAsynJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAsynJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAsynJobs"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsynJobs(request *ListAsynJobsRequest) (_result *ListAsynJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsynJobsResponse{}
	_body, _err := client.ListAsynJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRule"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveServersFromServerGroupWithOptions(request *RemoveServersFromServerGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveServersFromServerGroup"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveServersFromServerGroup(request *RemoveServersFromServerGroupRequest) (_result *RemoveServersFromServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.RemoveServersFromServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityPolicyWithOptions(request *CreateSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSecurityPolicy"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityPolicy(request *CreateSecurityPolicyRequest) (_result *CreateSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CreateSecurityPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRuleAttributeWithOptions(request *UpdateRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRuleAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRuleAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRuleAttribute(request *UpdateRuleAttributeRequest) (_result *UpdateRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleAttributeResponse{}
	_body, _err := client.UpdateRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityPoliciesWithOptions(request *ListSecurityPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSecurityPoliciesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecurityPolicies"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicies(request *ListSecurityPoliciesRequest) (_result *ListSecurityPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPoliciesResponse{}
	_body, _err := client.ListSecurityPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAclAttributeWithOptions(request *UpdateAclAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateAclAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAclAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAclAttribute(request *UpdateAclAttributeRequest) (_result *UpdateAclAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAclAttributeResponse{}
	_body, _err := client.UpdateAclAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableDeletionProtectionWithOptions(request *DisableDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *DisableDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableDeletionProtectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableDeletionProtection"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableDeletionProtection(request *DisableDeletionProtectionRequest) (_result *DisableDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDeletionProtectionResponse{}
	_body, _err := client.DisableDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopListenerWithOptions(request *StopListenerRequest, runtime *util.RuntimeOptions) (_result *StopListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopListener(request *StopListenerRequest) (_result *StopListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopListenerResponse{}
	_body, _err := client.StopListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServerGroupServersAttributeWithOptions(request *UpdateServerGroupServersAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServerGroupServersAttribute"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServerGroupServersAttribute(request *UpdateServerGroupServersAttributeRequest) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.UpdateServerGroupServersAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartListenerWithOptions(request *StartListenerRequest, runtime *util.RuntimeOptions) (_result *StartListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartListenerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartListener"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartListener(request *StartListenerRequest) (_result *StartListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartListenerResponse{}
	_body, _err := client.StartListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRulesWithOptions(request *CreateRulesRequest, runtime *util.RuntimeOptions) (_result *CreateRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRules"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRules(request *CreateRulesRequest) (_result *CreateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRulesResponse{}
	_body, _err := client.CreateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAclEntriesWithOptions(request *ListAclEntriesRequest, runtime *util.RuntimeOptions) (_result *ListAclEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAclEntriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAclEntries"), tea.String("2020-06-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAclEntries(request *ListAclEntriesRequest) (_result *ListAclEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAclEntriesResponse{}
	_body, _err := client.ListAclEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
