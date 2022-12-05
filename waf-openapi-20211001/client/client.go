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
)

type ClearMajorProtectionBlackIpRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ClearMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpRequest) SetInstanceId(v string) *ClearMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetRuleId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type ClearMajorProtectionBlackIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ClearMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type ClearMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClearMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClearMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ClearMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ClearMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearMajorProtectionBlackIpResponse) SetBody(v *ClearMajorProtectionBlackIpResponseBody) *ClearMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type CreateDefenseResourceGroupRequest struct {
	AddList     *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupRequest) SetAddList(v string) *CreateDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetDescription(v string) *CreateDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetGroupName(v string) *CreateDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetInstanceId(v string) *CreateDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

type CreateDefenseResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponseBody) SetRequestId(v string) *CreateDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *CreateDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetStatusCode(v int32) *CreateDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseResourceGroupResponse) SetBody(v *CreateDefenseResourceGroupResponseBody) *CreateDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type CreateDefenseRuleRequest struct {
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Rules        *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleRequest) SetDefenseScene(v string) *CreateDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetInstanceId(v string) *CreateDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetRules(v string) *CreateDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateDefenseRuleRequest) SetTemplateId(v int64) *CreateDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type CreateDefenseRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponseBody) SetRequestId(v string) *CreateDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseRuleResponse) SetHeaders(v map[string]*string) *CreateDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseRuleResponse) SetStatusCode(v int32) *CreateDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseRuleResponse) SetBody(v *CreateDefenseRuleResponseBody) *CreateDefenseRuleResponse {
	s.Body = v
	return s
}

type CreateDefenseTemplateRequest struct {
	DefenseScene   *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	TemplateStatus *int32  `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateRequest) SetDefenseScene(v string) *CreateDefenseTemplateRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetDescription(v string) *CreateDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetInstanceId(v string) *CreateDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateName(v string) *CreateDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateOrigin(v string) *CreateDefenseTemplateRequest {
	s.TemplateOrigin = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateStatus(v int32) *CreateDefenseTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateType(v string) *CreateDefenseTemplateRequest {
	s.TemplateType = &v
	return s
}

type CreateDefenseTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponseBody) SetRequestId(v string) *CreateDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefenseTemplateResponseBody) SetTemplateId(v int64) *CreateDefenseTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponse) SetHeaders(v map[string]*string) *CreateDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseTemplateResponse) SetStatusCode(v int32) *CreateDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseTemplateResponse) SetBody(v *CreateDefenseTemplateResponseBody) *CreateDefenseTemplateResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	AccessType *string                      `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Domain     *string                      `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Listen     *CreateDomainRequestListen   `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	Redirect   *CreateDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	RegionId   *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetAccessType(v string) *CreateDomainRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainRequest) SetListen(v *CreateDomainRequestListen) *CreateDomainRequest {
	s.Listen = v
	return s
}

func (s *CreateDomainRequest) SetRedirect(v *CreateDomainRequestRedirect) *CreateDomainRequest {
	s.Redirect = v
	return s
}

func (s *CreateDomainRequest) SetRegionId(v string) *CreateDomainRequest {
	s.RegionId = &v
	return s
}

type CreateDomainRequestListen struct {
	CertId             *string   `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CipherSuite        *int32    `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	CustomCiphers      []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	EnableTLSv3        *bool     `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	ExclusiveIp        *bool     `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	FocusHttps         *bool     `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	Http2Enabled       *bool     `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	HttpPorts          []*int32  `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	HttpsPorts         []*int32  `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	IPv6Enabled        *bool     `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	ProtectionResource *string   `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	TLSVersion         *string   `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	XffHeaderMode      *int32    `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	XffHeaders         []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s CreateDomainRequestListen) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestListen) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestListen) SetCertId(v string) *CreateDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *CreateDomainRequestListen) SetCipherSuite(v int32) *CreateDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *CreateDomainRequestListen) SetCustomCiphers(v []*string) *CreateDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *CreateDomainRequestListen) SetEnableTLSv3(v bool) *CreateDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *CreateDomainRequestListen) SetExclusiveIp(v bool) *CreateDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *CreateDomainRequestListen) SetFocusHttps(v bool) *CreateDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttp2Enabled(v bool) *CreateDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttpPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetHttpsPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetIPv6Enabled(v bool) *CreateDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetProtectionResource(v string) *CreateDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *CreateDomainRequestListen) SetTLSVersion(v string) *CreateDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaderMode(v int32) *CreateDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaders(v []*string) *CreateDomainRequestListen {
	s.XffHeaders = v
	return s
}

type CreateDomainRequestRedirect struct {
	Backends          []*string                                    `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	CnameEnabled      *bool                                        `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	ConnectTimeout    *int32                                       `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	FocusHttpBackend  *bool                                        `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	Keepalive         *bool                                        `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	KeepaliveRequests *int32                                       `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	KeepaliveTimeout  *int32                                       `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	Loadbalance       *string                                      `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	ReadTimeout       *int32                                       `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	RequestHeaders    []*CreateDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	Retry             *bool                                        `json:"Retry,omitempty" xml:"Retry,omitempty"`
	RoutingRules      *string                                      `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	SniEnabled        *bool                                        `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	SniHost           *string                                      `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	WriteTimeout      *int32                                       `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s CreateDomainRequestRedirect) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirect) SetBackends(v []*string) *CreateDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *CreateDomainRequestRedirect) SetCnameEnabled(v bool) *CreateDomainRequestRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetConnectTimeout(v int32) *CreateDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetFocusHttpBackend(v bool) *CreateDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepalive(v bool) *CreateDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveRequests(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveTimeout(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetLoadbalance(v string) *CreateDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetReadTimeout(v int32) *CreateDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRequestHeaders(v []*CreateDomainRequestRedirectRequestHeaders) *CreateDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *CreateDomainRequestRedirect) SetRetry(v bool) *CreateDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRoutingRules(v string) *CreateDomainRequestRedirect {
	s.RoutingRules = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniEnabled(v bool) *CreateDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniHost(v string) *CreateDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetWriteTimeout(v int32) *CreateDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

type CreateDomainRequestRedirectRequestHeaders struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDomainRequestRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetKey(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetValue(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

type CreateDomainShrinkRequest struct {
	AccessType     *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ListenShrink   *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDomainShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainShrinkRequest) SetAccessType(v string) *CreateDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetDomain(v string) *CreateDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetInstanceId(v string) *CreateDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetListenShrink(v string) *CreateDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRedirectShrink(v string) *CreateDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *CreateDomainShrinkRequest) SetRegionId(v string) *CreateDomainShrinkRequest {
	s.RegionId = &v
	return s
}

type CreateDomainResponseBody struct {
	DomainInfo *CreateDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetDomainInfo(v *CreateDomainResponseBodyDomainInfo) *CreateDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponseBodyDomainInfo struct {
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CreateDomainResponseBodyDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBodyDomainInfo) SetCname(v string) *CreateDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponseBodyDomainInfo) SetDomain(v string) *CreateDomainResponseBodyDomainInfo {
	s.Domain = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateMajorProtectionBlackIpRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList      *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpRequest) SetDescription(v string) *CreateMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *CreateMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetInstanceId(v string) *CreateMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetIpList(v string) *CreateMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRuleId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetTemplateId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type CreateMajorProtectionBlackIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponseBody) SetRequestId(v string) *CreateMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type CreateMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *CreateMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetStatusCode(v int32) *CreateMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMajorProtectionBlackIpResponse) SetBody(v *CreateMajorProtectionBlackIpResponseBody) *CreateMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type DeleteDefenseResourceGroupRequest struct {
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupRequest) SetGroupName(v string) *DeleteDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetInstanceId(v string) *DeleteDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteDefenseResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponseBody) SetRequestId(v string) *DeleteDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetStatusCode(v int32) *DeleteDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseResourceGroupResponse) SetBody(v *DeleteDefenseResourceGroupResponseBody) *DeleteDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteDefenseRuleRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleIds    *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleRequest) SetInstanceId(v string) *DeleteDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetRuleIds(v string) *DeleteDefenseRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetTemplateId(v int64) *DeleteDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type DeleteDefenseRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponseBody) SetRequestId(v string) *DeleteDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponse) SetHeaders(v map[string]*string) *DeleteDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseRuleResponse) SetStatusCode(v int32) *DeleteDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseRuleResponse) SetBody(v *DeleteDefenseRuleResponseBody) *DeleteDefenseRuleResponse {
	s.Body = v
	return s
}

type DeleteDefenseTemplateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateRequest) SetInstanceId(v string) *DeleteDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetTemplateId(v int64) *DeleteDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteDefenseTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponseBody) SetRequestId(v string) *DeleteDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponse) SetHeaders(v map[string]*string) *DeleteDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetStatusCode(v int32) *DeleteDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetBody(v *DeleteDefenseTemplateResponseBody) *DeleteDefenseTemplateResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetAccessType(v string) *DeleteDomainRequest {
	s.AccessType = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainRequest) SetInstanceId(v string) *DeleteDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDomainRequest) SetRegionId(v string) *DeleteDomainRequest {
	s.RegionId = &v
	return s
}

type DeleteDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteMajorProtectionBlackIpRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList     *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpRequest) SetInstanceId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetIpList(v string) *DeleteMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetRuleId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetTemplateId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type DeleteMajorProtectionBlackIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponseBody) SetRequestId(v string) *DeleteMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *DeleteMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetStatusCode(v int32) *DeleteMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponse) SetBody(v *DeleteMajorProtectionBlackIpResponseBody) *DeleteMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourceGroupRequest struct {
	GroupName  *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupRequest) SetGroupName(v string) *DescribeDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

type DescribeDefenseResourceGroupResponseBody struct {
	Group     *DescribeDefenseResourceGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBody) SetGroup(v *DescribeDefenseResourceGroupResponseBodyGroup) *DescribeDefenseResourceGroupResponseBody {
	s.Group = v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBody) SetRequestId(v string) *DescribeDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDefenseResourceGroupResponseBodyGroup struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ResourceList *string `json:"ResourceList,omitempty" xml:"ResourceList,omitempty"`
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetDescription(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtCreate(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGmtModified(v int64) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetGroupName(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponseBodyGroup) SetResourceList(v string) *DescribeDefenseResourceGroupResponseBodyGroup {
	s.ResourceList = &v
	return s
}

type DescribeDefenseResourceGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetStatusCode(v int32) *DescribeDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourceGroupResponse) SetBody(v *DescribeDefenseResourceGroupResponseBody) *DescribeDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type DescribeDefenseResourcesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query      *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeDefenseResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesRequest) SetInstanceId(v string) *DescribeDefenseResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageNumber(v int32) *DescribeDefenseResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageSize(v int32) *DescribeDefenseResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetQuery(v string) *DescribeDefenseResourcesRequest {
	s.Query = &v
	return s
}

type DescribeDefenseResourcesResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*DescribeDefenseResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	TotalCount *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBody) SetRequestId(v string) *DescribeDefenseResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetResources(v []*DescribeDefenseResourcesResponseBodyResources) *DescribeDefenseResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseResourcesResponseBody) SetTotalCount(v int64) *DescribeDefenseResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseResourcesResponseBodyResources struct {
	CustomHeaders  []*string              `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty" type:"Repeated"`
	Description    *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail         map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GmtCreate      *int64                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *int64                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Pattern        *string                `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Product        *string                `json:"Product,omitempty" xml:"Product,omitempty"`
	Resource       *string                `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceGroup  *string                `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	ResourceOrigin *string                `json:"ResourceOrigin,omitempty" xml:"ResourceOrigin,omitempty"`
	XffStatus      *int32                 `json:"XffStatus,omitempty" xml:"XffStatus,omitempty"`
}

func (s DescribeDefenseResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetCustomHeaders(v []*string) *DescribeDefenseResourcesResponseBodyResources {
	s.CustomHeaders = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDescription(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Description = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetDetail(v map[string]interface{}) *DescribeDefenseResourcesResponseBodyResources {
	s.Detail = v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtCreate(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetGmtModified(v int64) *DescribeDefenseResourcesResponseBodyResources {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetPattern(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Pattern = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetProduct(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Product = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResource(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceGroup(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetResourceOrigin(v string) *DescribeDefenseResourcesResponseBodyResources {
	s.ResourceOrigin = &v
	return s
}

func (s *DescribeDefenseResourcesResponseBodyResources) SetXffStatus(v int32) *DescribeDefenseResourcesResponseBodyResources {
	s.XffStatus = &v
	return s
}

type DescribeDefenseResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefenseResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesResponse) SetHeaders(v map[string]*string) *DescribeDefenseResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetStatusCode(v int32) *DescribeDefenseResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseResourcesResponse) SetBody(v *DescribeDefenseResourcesResponseBody) *DescribeDefenseResourcesResponse {
	s.Body = v
	return s
}

type DescribeDefenseRuleRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleRequest) SetInstanceId(v string) *DescribeDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetRuleId(v int64) *DescribeDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleRequest) SetTemplateId(v int64) *DescribeDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRuleResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rule      *DescribeDefenseRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s DescribeDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBody) SetRequestId(v string) *DescribeDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBody) SetRule(v *DescribeDefenseRuleResponseBodyRule) *DescribeDefenseRuleResponseBody {
	s.Rule = v
	return s
}

type DescribeDefenseRuleResponseBodyRule struct {
	Config        *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DefenseOrigin *string `json:"DefenseOrigin,omitempty" xml:"DefenseOrigin,omitempty"`
	DefenseScene  *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId        *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId    *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRuleResponseBodyRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponseBodyRule) SetConfig(v string) *DescribeDefenseRuleResponseBodyRule {
	s.Config = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseOrigin(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseOrigin = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetDefenseScene(v string) *DescribeDefenseRuleResponseBodyRule {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetGmtModified(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetRuleName(v string) *DescribeDefenseRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetStatus(v int32) *DescribeDefenseRuleResponseBodyRule {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRuleResponseBodyRule) SetTemplateId(v int64) *DescribeDefenseRuleResponseBodyRule {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRuleResponse) SetHeaders(v map[string]*string) *DescribeDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRuleResponse) SetStatusCode(v int32) *DescribeDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRuleResponse) SetBody(v *DescribeDefenseRuleResponseBody) *DescribeDefenseRuleResponse {
	s.Body = v
	return s
}

type DescribeDefenseRulesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query      *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RuleType   *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeDefenseRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesRequest) SetInstanceId(v string) *DescribeDefenseRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageNumber(v int32) *DescribeDefenseRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetPageSize(v int32) *DescribeDefenseRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetQuery(v string) *DescribeDefenseRulesRequest {
	s.Query = &v
	return s
}

func (s *DescribeDefenseRulesRequest) SetRuleType(v string) *DescribeDefenseRulesRequest {
	s.RuleType = &v
	return s
}

type DescribeDefenseRulesResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules      []*DescribeDefenseRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponseBody) SetRequestId(v string) *DescribeDefenseRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseRulesResponseBody) SetRules(v []*DescribeDefenseRulesResponseBodyRules) *DescribeDefenseRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeDefenseRulesResponseBody) SetTotalCount(v int64) *DescribeDefenseRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDefenseRulesResponseBodyRules struct {
	Config        *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DefenseOrigin *string `json:"DefenseOrigin,omitempty" xml:"DefenseOrigin,omitempty"`
	DefenseScene  *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleId        *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName      *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId    *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponseBodyRules) SetConfig(v string) *DescribeDefenseRulesResponseBodyRules {
	s.Config = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetDefenseOrigin(v string) *DescribeDefenseRulesResponseBodyRules {
	s.DefenseOrigin = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetDefenseScene(v string) *DescribeDefenseRulesResponseBodyRules {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetGmtModified(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetRuleId(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetRuleName(v string) *DescribeDefenseRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetStatus(v int32) *DescribeDefenseRulesResponseBodyRules {
	s.Status = &v
	return s
}

func (s *DescribeDefenseRulesResponseBodyRules) SetTemplateId(v int64) *DescribeDefenseRulesResponseBodyRules {
	s.TemplateId = &v
	return s
}

type DescribeDefenseRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefenseRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponse) SetHeaders(v map[string]*string) *DescribeDefenseRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRulesResponse) SetStatusCode(v int32) *DescribeDefenseRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRulesResponse) SetBody(v *DescribeDefenseRulesResponseBody) *DescribeDefenseRulesResponse {
	s.Body = v
	return s
}

type DescribeDefenseTemplateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateRequest) SetInstanceId(v string) *DescribeDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetTemplateId(v int64) *DescribeDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

type DescribeDefenseTemplateResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *DescribeDefenseTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s DescribeDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBody) SetRequestId(v string) *DescribeDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBody) SetTemplate(v *DescribeDefenseTemplateResponseBodyTemplate) *DescribeDefenseTemplateResponseBody {
	s.Template = v
	return s
}

type DescribeDefenseTemplateResponseBodyTemplate struct {
	DefenseScene   *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TemplateId     *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	TemplateStatus *int32  `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDefenseTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDefenseScene(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetDescription(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetGmtModified(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.GmtModified = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateId(v int64) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateName(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateOrigin(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateOrigin = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateStatus(v int32) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeDefenseTemplateResponseBodyTemplate) SetTemplateType(v string) *DescribeDefenseTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

type DescribeDefenseTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetStatusCode(v int32) *DescribeDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetBody(v *DescribeDefenseTemplateResponseBody) *DescribeDefenseTemplateResponse {
	s.Body = v
	return s
}

type DescribeDomainDetailRequest struct {
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailRequest) SetDomain(v string) *DescribeDomainDetailRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetInstanceId(v string) *DescribeDomainDetailRequest {
	s.InstanceId = &v
	return s
}

type DescribeDomainDetailResponseBody struct {
	Cname     *string                                   `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain    *string                                   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Listen    *DescribeDomainDetailResponseBodyListen   `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	Redirect  *DescribeDomainDetailResponseBodyRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int64                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBody) SetCname(v string) *DescribeDomainDetailResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomain(v string) *DescribeDomainDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetListen(v *DescribeDomainDetailResponseBodyListen) *DescribeDomainDetailResponseBody {
	s.Listen = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRedirect(v *DescribeDomainDetailResponseBodyRedirect) *DescribeDomainDetailResponseBody {
	s.Redirect = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRequestId(v string) *DescribeDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetStatus(v int64) *DescribeDomainDetailResponseBody {
	s.Status = &v
	return s
}

type DescribeDomainDetailResponseBodyListen struct {
	CertId             *int64    `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CipherSuite        *int64    `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	CustomCiphers      []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	EnableTLSv3        *bool     `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	ExclusiveIp        *bool     `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	FocusHttps         *bool     `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	Http2Enabled       *bool     `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	HttpPorts          []*int64  `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	HttpsPorts         []*int64  `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	IPv6Enabled        *bool     `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	ProtectionResource *string   `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	TLSVersion         *string   `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	XffHeaderMode      *int64    `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	XffHeaders         []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeDomainDetailResponseBodyListen) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyListen) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyListen) SetCertId(v int64) *DescribeDomainDetailResponseBodyListen {
	s.CertId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCipherSuite(v int64) *DescribeDomainDetailResponseBodyListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCustomCiphers(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetEnableTLSv3(v bool) *DescribeDomainDetailResponseBodyListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetExclusiveIp(v bool) *DescribeDomainDetailResponseBodyListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetFocusHttps(v bool) *DescribeDomainDetailResponseBodyListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttp2Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpsPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetIPv6Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetProtectionResource(v string) *DescribeDomainDetailResponseBodyListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetTLSVersion(v string) *DescribeDomainDetailResponseBodyListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaderMode(v int64) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaders(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaders = v
	return s
}

type DescribeDomainDetailResponseBodyRedirect struct {
	Backends          []*DescribeDomainDetailResponseBodyRedirectBackends       `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	ConnectTimeout    *int32                                                    `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	FocusHttpBackend  *bool                                                     `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	Keepalive         *bool                                                     `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	KeepaliveRequests *int32                                                    `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	KeepaliveTimeout  *int32                                                    `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	Loadbalance       *string                                                   `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	ReadTimeout       *int32                                                    `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	RequestHeaders    []*DescribeDomainDetailResponseBodyRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	Retry             *bool                                                     `json:"Retry,omitempty" xml:"Retry,omitempty"`
	SniEnabled        *bool                                                     `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	SniHost           *string                                                   `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	WriteTimeout      *int32                                                    `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirect) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirect) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackends(v []*DescribeDomainDetailResponseBodyRedirectBackends) *DescribeDomainDetailResponseBodyRedirect {
	s.Backends = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetConnectTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetFocusHttpBackend(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepalive(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveRequests(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetLoadbalance(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetReadTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRequestHeaders(v []*DescribeDomainDetailResponseBodyRedirectRequestHeaders) *DescribeDomainDetailResponseBodyRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRetry(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniEnabled(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniHost(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetWriteTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.WriteTimeout = &v
	return s
}

type DescribeDomainDetailResponseBodyRedirectBackends struct {
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectBackends) SetBackend(v string) *DescribeDomainDetailResponseBodyRedirectBackends {
	s.Backend = &v
	return s
}

type DescribeDomainDetailResponseBodyRedirectRequestHeaders struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetKey(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetValue(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Value = &v
	return s
}

type DescribeDomainDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailResponse) SetStatusCode(v int32) *DescribeDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainDetailResponse) SetBody(v *DescribeDomainDetailResponseBody) *DescribeDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	Backend    *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetBackend(v string) *DescribeDomainsRequest {
	s.Backend = &v
	return s
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceId(v string) *DescribeDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains    []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) SetDomains(v []*DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	Backeds     *DescribeDomainsResponseBodyDomainsBackeds     `json:"Backeds,omitempty" xml:"Backeds,omitempty" type:"Struct"`
	Cname       *string                                        `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain      *string                                        `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ListenPorts *DescribeDomainsResponseBodyDomainsListenPorts `json:"ListenPorts,omitempty" xml:"ListenPorts,omitempty" type:"Struct"`
	Status      *int32                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetBackeds(v *DescribeDomainsResponseBodyDomainsBackeds) *DescribeDomainsResponseBodyDomains {
	s.Backeds = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCname(v string) *DescribeDomainsResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v string) *DescribeDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetListenPorts(v *DescribeDomainsResponseBodyDomainsListenPorts) *DescribeDomainsResponseBodyDomains {
	s.ListenPorts = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetStatus(v int32) *DescribeDomainsResponseBodyDomains {
	s.Status = &v
	return s
}

type DescribeDomainsResponseBodyDomainsBackeds struct {
	Http  []*DescribeDomainsResponseBodyDomainsBackedsHttp  `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	Https []*DescribeDomainsResponseBodyDomainsBackedsHttps `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsBackeds) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackeds) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttp(v []*DescribeDomainsResponseBodyDomainsBackedsHttp) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsBackeds) SetHttps(v []*DescribeDomainsResponseBodyDomainsBackedsHttps) *DescribeDomainsResponseBodyDomainsBackeds {
	s.Https = v
	return s
}

type DescribeDomainsResponseBodyDomainsBackedsHttp struct {
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttp) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttp) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttp {
	s.Backend = &v
	return s
}

type DescribeDomainsResponseBodyDomainsBackedsHttps struct {
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsBackedsHttps) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsBackedsHttps) SetBackend(v string) *DescribeDomainsResponseBodyDomainsBackedsHttps {
	s.Backend = &v
	return s
}

type DescribeDomainsResponseBodyDomainsListenPorts struct {
	Http  []*int64 `json:"Http,omitempty" xml:"Http,omitempty" type:"Repeated"`
	Https []*int64 `json:"Https,omitempty" xml:"Https,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsListenPorts) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttp(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Http = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsListenPorts) SetHttps(v []*int64) *DescribeDomainsResponseBodyDomainsListenPorts {
	s.Https = v
	return s
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) SetHeaders(v map[string]*string) *DescribeDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsResponse) SetStatusCode(v int32) *DescribeDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type DescribeFlowChartRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowChartRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartRequest) SetEndTimestamp(v string) *DescribeFlowChartRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInstanceId(v string) *DescribeFlowChartRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowChartRequest) SetInterval(v string) *DescribeFlowChartRequest {
	s.Interval = &v
	return s
}

func (s *DescribeFlowChartRequest) SetResource(v string) *DescribeFlowChartRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowChartRequest) SetStartTimestamp(v string) *DescribeFlowChartRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowChartResponseBody struct {
	FlowChart []*DescribeFlowChartResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowChartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBody) SetFlowChart(v []*DescribeFlowChartResponseBodyFlowChart) *DescribeFlowChartResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribeFlowChartResponseBody) SetRequestId(v string) *DescribeFlowChartResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowChartResponseBodyFlowChart struct {
	AclCustomBlockSum     *int64  `json:"AclCustomBlockSum,omitempty" xml:"AclCustomBlockSum,omitempty"`
	AclCustomReportsSum   *int64  `json:"AclCustomReportsSum,omitempty" xml:"AclCustomReportsSum,omitempty"`
	AntiScanBlockSum      *int64  `json:"AntiScanBlockSum,omitempty" xml:"AntiScanBlockSum,omitempty"`
	AntibotBlockSum       *int64  `json:"AntibotBlockSum,omitempty" xml:"AntibotBlockSum,omitempty"`
	AntibotReportSum      *string `json:"AntibotReportSum,omitempty" xml:"AntibotReportSum,omitempty"`
	AntiscanReportsSum    *int64  `json:"AntiscanReportsSum,omitempty" xml:"AntiscanReportsSum,omitempty"`
	BlacklistBlockSum     *string `json:"BlacklistBlockSum,omitempty" xml:"BlacklistBlockSum,omitempty"`
	BlacklistReportsSum   *int64  `json:"BlacklistReportsSum,omitempty" xml:"BlacklistReportsSum,omitempty"`
	CcCustomBlockSum      *int64  `json:"CcCustomBlockSum,omitempty" xml:"CcCustomBlockSum,omitempty"`
	CcCustomReportsSum    *int64  `json:"CcCustomReportsSum,omitempty" xml:"CcCustomReportsSum,omitempty"`
	CcSystemBlocksSum     *int64  `json:"CcSystemBlocksSum,omitempty" xml:"CcSystemBlocksSum,omitempty"`
	CcSystemReportsSum    *int64  `json:"CcSystemReportsSum,omitempty" xml:"CcSystemReportsSum,omitempty"`
	Count                 *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	InBytes               *int64  `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	Index                 *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	MaxPv                 *int64  `json:"MaxPv,omitempty" xml:"MaxPv,omitempty"`
	OutBytes              *int64  `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	RegionBlockBlocksSum  *int64  `json:"RegionBlockBlocksSum,omitempty" xml:"RegionBlockBlocksSum,omitempty"`
	RegionBlockReportsSum *int64  `json:"RegionBlockReportsSum,omitempty" xml:"RegionBlockReportsSum,omitempty"`
	WafBlockSum           *int64  `json:"WafBlockSum,omitempty" xml:"WafBlockSum,omitempty"`
	WafReportSum          *string `json:"WafReportSum,omitempty" xml:"WafReportSum,omitempty"`
}

func (s DescribeFlowChartResponseBodyFlowChart) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiScanBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiScanBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiscanReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiscanReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistBlockSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCount(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetInBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.InBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetIndex(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetMaxPv(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.MaxPv = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetOutBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.OutBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.WafBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.WafReportSum = &v
	return s
}

type DescribeFlowChartResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFlowChartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowChartResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponse) SetHeaders(v map[string]*string) *DescribeFlowChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowChartResponse) SetStatusCode(v int32) *DescribeFlowChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowChartResponse) SetBody(v *DescribeFlowChartResponseBody) *DescribeFlowChartResponse {
	s.Body = v
	return s
}

type DescribeFlowTopResourceRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowTopResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceRequest) SetEndTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetInstanceId(v string) *DescribeFlowTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopResourceRequest) SetStartTimestamp(v string) *DescribeFlowTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowTopResourceResponseBody struct {
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopResource []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBody) SetRequestId(v string) *DescribeFlowTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeFlowTopResourceResponseBodyRuleHitsTopResource) *DescribeFlowTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

type DescribeFlowTopResourceResponseBodyRuleHitsTopResource struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeFlowTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

type DescribeFlowTopResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFlowTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowTopResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopResourceResponse) SetHeaders(v map[string]*string) *DescribeFlowTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetStatusCode(v int32) *DescribeFlowTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopResourceResponse) SetBody(v *DescribeFlowTopResourceResponseBody) *DescribeFlowTopResourceResponse {
	s.Body = v
	return s
}

type DescribeFlowTopUrlRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFlowTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlRequest) SetEndTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetInstanceId(v string) *DescribeFlowTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetResource(v string) *DescribeFlowTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeFlowTopUrlRequest) SetStartTimestamp(v string) *DescribeFlowTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeFlowTopUrlResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopUrl []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeFlowTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBody) SetRequestId(v string) *DescribeFlowTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) *DescribeFlowTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

type DescribeFlowTopUrlResponseBodyRuleHitsTopUrl struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeFlowTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

type DescribeFlowTopUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFlowTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTopUrlResponse) SetHeaders(v map[string]*string) *DescribeFlowTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetStatusCode(v int32) *DescribeFlowTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowTopUrlResponse) SetBody(v *DescribeFlowTopUrlResponseBody) *DescribeFlowTopUrlResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetRegionId(v string) *DescribeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetResourceGroupId(v string) *DescribeInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceResponseBody struct {
	Details    *DescribeInstanceResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Struct"`
	Edition    *string                              `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EndTime    *int64                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InDebt     *string                              `json:"InDebt,omitempty" xml:"InDebt,omitempty"`
	InstanceId *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PayType    *string                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId   *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime  *int64                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status     *int32                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetDetails(v *DescribeInstanceResponseBodyDetails) *DescribeInstanceResponseBody {
	s.Details = v
	return s
}

func (s *DescribeInstanceResponseBody) SetEdition(v string) *DescribeInstanceResponseBody {
	s.Edition = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEndTime(v int64) *DescribeInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInDebt(v string) *DescribeInstanceResponseBody {
	s.InDebt = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStartTime(v int64) *DescribeInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v int32) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

type DescribeInstanceResponseBodyDetails struct {
	AclRuleMaxIpCount                    *int64  `json:"AclRuleMaxIpCount,omitempty" xml:"AclRuleMaxIpCount,omitempty"`
	AntiScan                             *bool   `json:"AntiScan,omitempty" xml:"AntiScan,omitempty"`
	AntiScanTemplateMaxCount             *int64  `json:"AntiScanTemplateMaxCount,omitempty" xml:"AntiScanTemplateMaxCount,omitempty"`
	BackendMaxCount                      *int64  `json:"BackendMaxCount,omitempty" xml:"BackendMaxCount,omitempty"`
	BaseWafGroup                         *bool   `json:"BaseWafGroup,omitempty" xml:"BaseWafGroup,omitempty"`
	BaseWafGroupRuleInTemplateMaxCount   *int64  `json:"BaseWafGroupRuleInTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleInTemplateMaxCount,omitempty"`
	BaseWafGroupRuleTemplateMaxCount     *int64  `json:"BaseWafGroupRuleTemplateMaxCount,omitempty" xml:"BaseWafGroupRuleTemplateMaxCount,omitempty"`
	Bot                                  *bool   `json:"Bot,omitempty" xml:"Bot,omitempty"`
	BotApp                               *string `json:"BotApp,omitempty" xml:"BotApp,omitempty"`
	BotTemplateMaxCount                  *int64  `json:"BotTemplateMaxCount,omitempty" xml:"BotTemplateMaxCount,omitempty"`
	BotWeb                               *string `json:"BotWeb,omitempty" xml:"BotWeb,omitempty"`
	CnameResourceMaxCount                *int64  `json:"CnameResourceMaxCount,omitempty" xml:"CnameResourceMaxCount,omitempty"`
	CustomResponse                       *bool   `json:"CustomResponse,omitempty" xml:"CustomResponse,omitempty"`
	CustomResponseRuleInTemplateMaxCount *int64  `json:"CustomResponseRuleInTemplateMaxCount,omitempty" xml:"CustomResponseRuleInTemplateMaxCount,omitempty"`
	CustomResponseTemplateMaxCount       *int64  `json:"CustomResponseTemplateMaxCount,omitempty" xml:"CustomResponseTemplateMaxCount,omitempty"`
	CustomRule                           *bool   `json:"CustomRule,omitempty" xml:"CustomRule,omitempty"`
	CustomRuleAction                     *string `json:"CustomRuleAction,omitempty" xml:"CustomRuleAction,omitempty"`
	CustomRuleCondition                  *string `json:"CustomRuleCondition,omitempty" xml:"CustomRuleCondition,omitempty"`
	CustomRuleInTemplateMaxCount         *int64  `json:"CustomRuleInTemplateMaxCount,omitempty" xml:"CustomRuleInTemplateMaxCount,omitempty"`
	CustomRuleRatelimitor                *string `json:"CustomRuleRatelimitor,omitempty" xml:"CustomRuleRatelimitor,omitempty"`
	CustomRuleTemplateMaxCount           *int64  `json:"CustomRuleTemplateMaxCount,omitempty" xml:"CustomRuleTemplateMaxCount,omitempty"`
	DefenseGroupMaxCount                 *int64  `json:"DefenseGroupMaxCount,omitempty" xml:"DefenseGroupMaxCount,omitempty"`
	DefenseObjectInGroupMaxCount         *int64  `json:"DefenseObjectInGroupMaxCount,omitempty" xml:"DefenseObjectInGroupMaxCount,omitempty"`
	DefenseObjectInTemplateMaxCount      *int64  `json:"DefenseObjectInTemplateMaxCount,omitempty" xml:"DefenseObjectInTemplateMaxCount,omitempty"`
	DefenseObjectMaxCount                *int64  `json:"DefenseObjectMaxCount,omitempty" xml:"DefenseObjectMaxCount,omitempty"`
	Dlp                                  *bool   `json:"Dlp,omitempty" xml:"Dlp,omitempty"`
	DlpRuleInTemplateMaxCount            *int64  `json:"DlpRuleInTemplateMaxCount,omitempty" xml:"DlpRuleInTemplateMaxCount,omitempty"`
	DlpTemplateMaxCount                  *int64  `json:"DlpTemplateMaxCount,omitempty" xml:"DlpTemplateMaxCount,omitempty"`
	ExclusiveIp                          *bool   `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	Gslb                                 *bool   `json:"Gslb,omitempty" xml:"Gslb,omitempty"`
	HttpPorts                            *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	HttpsPorts                           *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	IpBlacklist                          *bool   `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	IpBlacklistIpInRuleMaxCount          *int64  `json:"IpBlacklistIpInRuleMaxCount,omitempty" xml:"IpBlacklistIpInRuleMaxCount,omitempty"`
	IpBlacklistRuleInTemplateMaxCount    *int64  `json:"IpBlacklistRuleInTemplateMaxCount,omitempty" xml:"IpBlacklistRuleInTemplateMaxCount,omitempty"`
	IpBlacklistTemplateMaxCount          *int64  `json:"IpBlacklistTemplateMaxCount,omitempty" xml:"IpBlacklistTemplateMaxCount,omitempty"`
	Ipv6                                 *bool   `json:"Ipv6,omitempty" xml:"Ipv6,omitempty"`
	LogService                           *bool   `json:"LogService,omitempty" xml:"LogService,omitempty"`
	MajorProtection                      *bool   `json:"MajorProtection,omitempty" xml:"MajorProtection,omitempty"`
	MajorProtectionTemplateMaxCount      *int64  `json:"MajorProtectionTemplateMaxCount,omitempty" xml:"MajorProtectionTemplateMaxCount,omitempty"`
	Tamperproof                          *bool   `json:"Tamperproof,omitempty" xml:"Tamperproof,omitempty"`
	TamperproofRuleInTemplateMaxCount    *int64  `json:"TamperproofRuleInTemplateMaxCount,omitempty" xml:"TamperproofRuleInTemplateMaxCount,omitempty"`
	TamperproofTemplateMaxCount          *int64  `json:"TamperproofTemplateMaxCount,omitempty" xml:"TamperproofTemplateMaxCount,omitempty"`
	VastIpBlacklistInFileMaxCount        *int64  `json:"VastIpBlacklistInFileMaxCount,omitempty" xml:"VastIpBlacklistInFileMaxCount,omitempty"`
	VastIpBlacklistInOperationMaxCount   *int64  `json:"VastIpBlacklistInOperationMaxCount,omitempty" xml:"VastIpBlacklistInOperationMaxCount,omitempty"`
	VastIpBlacklistMaxCount              *int64  `json:"VastIpBlacklistMaxCount,omitempty" xml:"VastIpBlacklistMaxCount,omitempty"`
	Whitelist                            *bool   `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	WhitelistLogical                     *string `json:"WhitelistLogical,omitempty" xml:"WhitelistLogical,omitempty"`
	WhitelistRuleCondition               *string `json:"WhitelistRuleCondition,omitempty" xml:"WhitelistRuleCondition,omitempty"`
	WhitelistRuleInTemplateMaxCount      *int64  `json:"WhitelistRuleInTemplateMaxCount,omitempty" xml:"WhitelistRuleInTemplateMaxCount,omitempty"`
	WhitelistTemplateMaxCount            *int64  `json:"WhitelistTemplateMaxCount,omitempty" xml:"WhitelistTemplateMaxCount,omitempty"`
}

func (s DescribeInstanceResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyDetails) SetAclRuleMaxIpCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AclRuleMaxIpCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScan(v bool) *DescribeInstanceResponseBodyDetails {
	s.AntiScan = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetAntiScanTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.AntiScanTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBackendMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BackendMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroup(v bool) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroup = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBaseWafGroupRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BaseWafGroupRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBot(v bool) *DescribeInstanceResponseBodyDetails {
	s.Bot = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotApp(v string) *DescribeInstanceResponseBodyDetails {
	s.BotApp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.BotTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetBotWeb(v string) *DescribeInstanceResponseBodyDetails {
	s.BotWeb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCnameResourceMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CnameResourceMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponse(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomResponse = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomResponseTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomResponseTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRule(v bool) *DescribeInstanceResponseBodyDetails {
	s.CustomRule = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleAction(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleAction = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleRatelimitor(v string) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleRatelimitor = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetCustomRuleTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.CustomRuleTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInGroupMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInGroupMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDefenseObjectMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DefenseObjectMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlp(v bool) *DescribeInstanceResponseBodyDetails {
	s.Dlp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetDlpTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.DlpTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetExclusiveIp(v bool) *DescribeInstanceResponseBodyDetails {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetGslb(v bool) *DescribeInstanceResponseBodyDetails {
	s.Gslb = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetHttpsPorts(v string) *DescribeInstanceResponseBodyDetails {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklist(v bool) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistIpInRuleMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistIpInRuleMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpBlacklistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.IpBlacklistTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetIpv6(v bool) *DescribeInstanceResponseBodyDetails {
	s.Ipv6 = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetLogService(v bool) *DescribeInstanceResponseBodyDetails {
	s.LogService = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtection(v bool) *DescribeInstanceResponseBodyDetails {
	s.MajorProtection = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetMajorProtectionTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.MajorProtectionTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproof(v bool) *DescribeInstanceResponseBodyDetails {
	s.Tamperproof = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetTamperproofTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.TamperproofTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInFileMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInFileMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistInOperationMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistInOperationMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetVastIpBlacklistMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.VastIpBlacklistMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelist(v bool) *DescribeInstanceResponseBodyDetails {
	s.Whitelist = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistLogical(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistLogical = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleCondition(v string) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleCondition = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistRuleInTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistRuleInTemplateMaxCount = &v
	return s
}

func (s *DescribeInstanceResponseBodyDetails) SetWhitelistTemplateMaxCount(v int64) *DescribeInstanceResponseBodyDetails {
	s.WhitelistTemplateMaxCount = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeMajorProtectionBlackIpsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpLike     *string `json:"IpLike,omitempty" xml:"IpLike,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetInstanceId(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetIpLike(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.IpLike = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetOrderBy(v string) *DescribeMajorProtectionBlackIpsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageNumber(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetPageSize(v int32) *DescribeMajorProtectionBlackIpsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsRequest) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsRequest {
	s.TemplateId = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponseBody struct {
	IpList     []*DescribeMajorProtectionBlackIpsResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetIpList(v []*DescribeMajorProtectionBlackIpsResponseBodyIpList) *DescribeMajorProtectionBlackIpsResponseBody {
	s.IpList = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetRequestId(v string) *DescribeMajorProtectionBlackIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBody) SetTotalCount(v int64) *DescribeMajorProtectionBlackIpsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponseBodyIpList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetDescription(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Description = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetExpiredTime(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetGmtModified(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.GmtModified = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetIp(v string) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetRuleId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.RuleId = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponseBodyIpList) SetTemplateId(v int64) *DescribeMajorProtectionBlackIpsResponseBodyIpList {
	s.TemplateId = &v
	return s
}

type DescribeMajorProtectionBlackIpsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMajorProtectionBlackIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMajorProtectionBlackIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetHeaders(v map[string]*string) *DescribeMajorProtectionBlackIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetStatusCode(v int32) *DescribeMajorProtectionBlackIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetBody(v *DescribeMajorProtectionBlackIpsResponseBody) *DescribeMajorProtectionBlackIpsResponse {
	s.Body = v
	return s
}

type DescribePeakTrendRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribePeakTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendRequest) SetEndTimestamp(v string) *DescribePeakTrendRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInstanceId(v string) *DescribePeakTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInterval(v string) *DescribePeakTrendRequest {
	s.Interval = &v
	return s
}

func (s *DescribePeakTrendRequest) SetResource(v string) *DescribePeakTrendRequest {
	s.Resource = &v
	return s
}

func (s *DescribePeakTrendRequest) SetStartTimestamp(v string) *DescribePeakTrendRequest {
	s.StartTimestamp = &v
	return s
}

type DescribePeakTrendResponseBody struct {
	FlowChart []*DescribePeakTrendResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePeakTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBody) SetFlowChart(v []*DescribePeakTrendResponseBodyFlowChart) *DescribePeakTrendResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribePeakTrendResponseBody) SetRequestId(v string) *DescribePeakTrendResponseBody {
	s.RequestId = &v
	return s
}

type DescribePeakTrendResponseBodyFlowChart struct {
	AclSum      *int64 `json:"AclSum,omitempty" xml:"AclSum,omitempty"`
	AntiScanSum *int64 `json:"AntiScanSum,omitempty" xml:"AntiScanSum,omitempty"`
	CcSum       *int64 `json:"CcSum,omitempty" xml:"CcSum,omitempty"`
	Count       *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	Index       *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	WafSum      *int64 `json:"WafSum,omitempty" xml:"WafSum,omitempty"`
}

func (s DescribePeakTrendResponseBodyFlowChart) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAclSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AclSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAntiScanSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AntiScanSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCcSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.CcSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCount(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetIndex(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetWafSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.WafSum = &v
	return s
}

type DescribePeakTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePeakTrendResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePeakTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePeakTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponse) SetHeaders(v map[string]*string) *DescribePeakTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribePeakTrendResponse) SetStatusCode(v int32) *DescribePeakTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePeakTrendResponse) SetBody(v *DescribePeakTrendResponseBody) *DescribePeakTrendResponse {
	s.Body = v
	return s
}

type DescribeResourceLogStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resources  *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeResourceLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusRequest) SetInstanceId(v string) *DescribeResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceLogStatusRequest) SetResources(v string) *DescribeResourceLogStatusRequest {
	s.Resources = &v
	return s
}

type DescribeResourceLogStatusResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeResourceLogStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s DescribeResourceLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBody) SetRequestId(v string) *DescribeResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBody) SetResult(v []*DescribeResourceLogStatusResponseBodyResult) *DescribeResourceLogStatusResponseBody {
	s.Result = v
	return s
}

type DescribeResourceLogStatusResponseBodyResult struct {
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Status   *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetResource(v string) *DescribeResourceLogStatusResponseBodyResult {
	s.Resource = &v
	return s
}

func (s *DescribeResourceLogStatusResponseBodyResult) SetStatus(v bool) *DescribeResourceLogStatusResponseBodyResult {
	s.Status = &v
	return s
}

type DescribeResourceLogStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogStatusResponse) SetHeaders(v map[string]*string) *DescribeResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetStatusCode(v int32) *DescribeResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogStatusResponse) SetBody(v *DescribeResourceLogStatusResponseBody) *DescribeResourceLogStatusResponse {
	s.Body = v
	return s
}

type DescribeResourcePortRequest struct {
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
}

func (s DescribeResourcePortRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortRequest) SetInstanceId(v string) *DescribeResourcePortRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetResourceInstanceId(v string) *DescribeResourcePortRequest {
	s.ResourceInstanceId = &v
	return s
}

type DescribeResourcePortResponseBody struct {
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePorts []*string `json:"ResourcePorts,omitempty" xml:"ResourcePorts,omitempty" type:"Repeated"`
}

func (s DescribeResourcePortResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponseBody) SetRequestId(v string) *DescribeResourcePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePortResponseBody) SetResourcePorts(v []*string) *DescribeResourcePortResponseBody {
	s.ResourcePorts = v
	return s
}

type DescribeResourcePortResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourcePortResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePortResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePortResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponse) SetHeaders(v map[string]*string) *DescribeResourcePortResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePortResponse) SetStatusCode(v int32) *DescribeResourcePortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePortResponse) SetBody(v *DescribeResourcePortResponseBody) *DescribeResourcePortResponse {
	s.Body = v
	return s
}

type DescribeResponseCodeTrendGraphRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResponseCodeTrendGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphRequest) SetEndTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInstanceId(v string) *DescribeResponseCodeTrendGraphRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetInterval(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Interval = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetResource(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Resource = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetStartTimestamp(v string) *DescribeResponseCodeTrendGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphRequest) SetType(v string) *DescribeResponseCodeTrendGraphRequest {
	s.Type = &v
	return s
}

type DescribeResponseCodeTrendGraphResponseBody struct {
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseCodes []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes `json:"ResponseCodes,omitempty" xml:"ResponseCodes,omitempty" type:"Repeated"`
}

func (s DescribeResponseCodeTrendGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetRequestId(v string) *DescribeResponseCodeTrendGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetResponseCodes(v []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes) *DescribeResponseCodeTrendGraphResponseBody {
	s.ResponseCodes = v
	return s
}

type DescribeResponseCodeTrendGraphResponseBodyResponseCodes struct {
	Code302Pv *int64 `json:"302Pv,omitempty" xml:"302Pv,omitempty"`
	Code405Pv *int64 `json:"405Pv,omitempty" xml:"405Pv,omitempty"`
	Code499Pv *int64 `json:"499Pv,omitempty" xml:"499Pv,omitempty"`
	Code5xxPv *int64 `json:"5xxPv,omitempty" xml:"5xxPv,omitempty"`
	Index     *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode302Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code302Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode405Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code405Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode499Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code499Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode5xxPv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code5xxPv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetIndex(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Index = &v
	return s
}

type DescribeResponseCodeTrendGraphResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResponseCodeTrendGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResponseCodeTrendGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponse) SetHeaders(v map[string]*string) *DescribeResponseCodeTrendGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetStatusCode(v int32) *DescribeResponseCodeTrendGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponse) SetBody(v *DescribeResponseCodeTrendGraphResponseBody) *DescribeResponseCodeTrendGraphResponse {
	s.Body = v
	return s
}

type DescribeRuleGroupsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchType  *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeRuleGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsRequest) SetInstanceId(v string) *DescribeRuleGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageNumber(v int32) *DescribeRuleGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetPageSize(v int32) *DescribeRuleGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchType(v string) *DescribeRuleGroupsRequest {
	s.SearchType = &v
	return s
}

func (s *DescribeRuleGroupsRequest) SetSearchValue(v string) *DescribeRuleGroupsRequest {
	s.SearchValue = &v
	return s
}

type DescribeRuleGroupsResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleGroups []*DescribeRuleGroupsResponseBodyRuleGroups `json:"RuleGroups,omitempty" xml:"RuleGroups,omitempty" type:"Repeated"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRuleGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBody) SetRequestId(v string) *DescribeRuleGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetRuleGroups(v []*DescribeRuleGroupsResponseBodyRuleGroups) *DescribeRuleGroupsResponseBody {
	s.RuleGroups = v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetTotalCount(v int64) *DescribeRuleGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRuleGroupsResponseBodyRuleGroups struct {
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RuleGroupId    *int64  `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	RuleGroupName  *string `json:"RuleGroupName,omitempty" xml:"RuleGroupName,omitempty"`
	RuleTotalCount *int32  `json:"RuleTotalCount,omitempty" xml:"RuleTotalCount,omitempty"`
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetGmtModified(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupName(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupName = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleTotalCount(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleTotalCount = &v
	return s
}

type DescribeRuleGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponse) SetHeaders(v map[string]*string) *DescribeRuleGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleGroupsResponse) SetStatusCode(v int32) *DescribeRuleGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleGroupsResponse) SetBody(v *DescribeRuleGroupsResponseBody) *DescribeRuleGroupsResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopClientIpRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleType       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopClientIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetInstanceId(v string) *DescribeRuleHitsTopClientIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetResource(v string) *DescribeRuleHitsTopClientIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetRuleType(v string) *DescribeRuleHitsTopClientIpRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopClientIpRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopClientIpResponseBody struct {
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopClientIp []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp `json:"RuleHitsTopClientIp,omitempty" xml:"RuleHitsTopClientIp,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopClientIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRequestId(v string) *DescribeRuleHitsTopClientIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBody) SetRuleHitsTopClientIp(v []*DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) *DescribeRuleHitsTopClientIpResponseBody {
	s.RuleHitsTopClientIp = v
	return s
}

type DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp struct {
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetClientIp(v string) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.ClientIp = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp) SetCount(v int64) *DescribeRuleHitsTopClientIpResponseBodyRuleHitsTopClientIp {
	s.Count = &v
	return s
}

type DescribeRuleHitsTopClientIpResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopClientIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopClientIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopClientIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopClientIpResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopClientIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetStatusCode(v int32) *DescribeRuleHitsTopClientIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopClientIpResponse) SetBody(v *DescribeRuleHitsTopClientIpResponseBody) *DescribeRuleHitsTopClientIpResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopResourceRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleType       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetInstanceId(v string) *DescribeRuleHitsTopResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetRuleType(v string) *DescribeRuleHitsTopResourceRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopResourceRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopResourceRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopResourceResponseBody struct {
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopResource []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource `json:"RuleHitsTopResource,omitempty" xml:"RuleHitsTopResource,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRequestId(v string) *DescribeRuleHitsTopResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBody) SetRuleHitsTopResource(v []*DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) *DescribeRuleHitsTopResourceResponseBody {
	s.RuleHitsTopResource = v
	return s
}

type DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetCount(v int64) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource) SetResource(v string) *DescribeRuleHitsTopResourceResponseBodyRuleHitsTopResource {
	s.Resource = &v
	return s
}

type DescribeRuleHitsTopResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetStatusCode(v int32) *DescribeRuleHitsTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetBody(v *DescribeRuleHitsTopResourceResponseBody) *DescribeRuleHitsTopResourceResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopRuleIdRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleType       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetInstanceId(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetResource(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetRuleType(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopRuleIdRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopRuleIdResponseBody struct {
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopRuleId []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId `json:"RuleHitsTopRuleId,omitempty" xml:"RuleHitsTopRuleId,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopRuleIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRequestId(v string) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBody) SetRuleHitsTopRuleId(v []*DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) *DescribeRuleHitsTopRuleIdResponseBody {
	s.RuleHitsTopRuleId = v
	return s
}

type DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetCount(v int64) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetResource(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId) SetRuleId(v string) *DescribeRuleHitsTopRuleIdResponseBodyRuleHitsTopRuleId {
	s.RuleId = &v
	return s
}

type DescribeRuleHitsTopRuleIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopRuleIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopRuleIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopRuleIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetStatusCode(v int32) *DescribeRuleHitsTopRuleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetBody(v *DescribeRuleHitsTopRuleIdResponseBody) *DescribeRuleHitsTopRuleIdResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopTuleTypeRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetInstanceId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetResource(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopTuleTypeResponseBody struct {
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopTuleType []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType `json:"RuleHitsTopTuleType,omitempty" xml:"RuleHitsTopTuleType,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRequestId(v string) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBody) SetRuleHitsTopTuleType(v []*DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) *DescribeRuleHitsTopTuleTypeResponseBody {
	s.RuleHitsTopTuleType = v
	return s
}

type DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType struct {
	Count    *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetCount(v int64) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType) SetRuleType(v string) *DescribeRuleHitsTopTuleTypeResponseBodyRuleHitsTopTuleType {
	s.RuleType = &v
	return s
}

type DescribeRuleHitsTopTuleTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopTuleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopTuleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopTuleTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetStatusCode(v int32) *DescribeRuleHitsTopTuleTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetBody(v *DescribeRuleHitsTopTuleTypeResponseBody) *DescribeRuleHitsTopTuleTypeResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopUaRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetInstanceId(v string) *DescribeRuleHitsTopUaRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetResource(v string) *DescribeRuleHitsTopUaRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopUaResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopUa []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa `json:"RuleHitsTopUa,omitempty" xml:"RuleHitsTopUa,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBody) SetRuleHitsTopUa(v []*DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) *DescribeRuleHitsTopUaResponseBody {
	s.RuleHitsTopUa = v
	return s
}

type DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Ua    *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetCount(v int64) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa) SetUa(v string) *DescribeRuleHitsTopUaResponseBodyRuleHitsTopUa {
	s.Ua = &v
	return s
}

type DescribeRuleHitsTopUaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopUaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopUaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetBody(v *DescribeRuleHitsTopUaResponseBody) *DescribeRuleHitsTopUaResponse {
	s.Body = v
	return s
}

type DescribeRuleHitsTopUrlRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	RuleType       *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetInstanceId(v string) *DescribeRuleHitsTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetResource(v string) *DescribeRuleHitsTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetRuleType(v string) *DescribeRuleHitsTopUrlRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeRuleHitsTopUrlResponseBody struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleHitsTopUrl []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl `json:"RuleHitsTopUrl,omitempty" xml:"RuleHitsTopUrl,omitempty" type:"Repeated"`
}

func (s DescribeRuleHitsTopUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRequestId(v string) *DescribeRuleHitsTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBody) SetRuleHitsTopUrl(v []*DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) *DescribeRuleHitsTopUrlResponseBody {
	s.RuleHitsTopUrl = v
	return s
}

type DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetCount(v int64) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Count = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl) SetUrl(v string) *DescribeRuleHitsTopUrlResponseBodyRuleHitsTopUrl {
	s.Url = &v
	return s
}

type DescribeRuleHitsTopUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRuleHitsTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRuleHitsTopUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetBody(v *DescribeRuleHitsTopUrlResponseBody) *DescribeRuleHitsTopUrlResponse {
	s.Body = v
	return s
}

type DescribeTemplateResourcesRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesRequest) SetInstanceId(v string) *DescribeTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetResourceType(v string) *DescribeTemplateResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTemplateResourcesRequest) SetTemplateId(v int64) *DescribeTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

type DescribeTemplateResourcesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeTemplateResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponseBody) SetRequestId(v string) *DescribeTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetResources(v []*string) *DescribeTemplateResourcesResponseBody {
	s.Resources = v
	return s
}

type DescribeTemplateResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTemplateResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponse) SetHeaders(v map[string]*string) *DescribeTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetStatusCode(v int32) *DescribeTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResourcesResponse) SetBody(v *DescribeTemplateResourcesResponseBody) *DescribeTemplateResourcesResponse {
	s.Body = v
	return s
}

type DescribeVisitTopIpRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeVisitTopIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpRequest) SetEndTimestamp(v string) *DescribeVisitTopIpRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetInstanceId(v string) *DescribeVisitTopIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetResource(v string) *DescribeVisitTopIpRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitTopIpRequest) SetStartTimestamp(v string) *DescribeVisitTopIpRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeVisitTopIpResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopIp     []*DescribeVisitTopIpResponseBodyTopIp `json:"TopIp,omitempty" xml:"TopIp,omitempty" type:"Repeated"`
}

func (s DescribeVisitTopIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBody) SetRequestId(v string) *DescribeVisitTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitTopIpResponseBody) SetTopIp(v []*DescribeVisitTopIpResponseBodyTopIp) *DescribeVisitTopIpResponseBody {
	s.TopIp = v
	return s
}

type DescribeVisitTopIpResponseBodyTopIp struct {
	Area  *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Ip    *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Isp   *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeVisitTopIpResponseBodyTopIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponseBodyTopIp) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetArea(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Area = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetCount(v int64) *DescribeVisitTopIpResponseBodyTopIp {
	s.Count = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Ip = &v
	return s
}

func (s *DescribeVisitTopIpResponseBodyTopIp) SetIsp(v string) *DescribeVisitTopIpResponseBodyTopIp {
	s.Isp = &v
	return s
}

type DescribeVisitTopIpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVisitTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVisitTopIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitTopIpResponse) SetHeaders(v map[string]*string) *DescribeVisitTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitTopIpResponse) SetStatusCode(v int32) *DescribeVisitTopIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitTopIpResponse) SetBody(v *DescribeVisitTopIpResponseBody) *DescribeVisitTopIpResponse {
	s.Body = v
	return s
}

type DescribeVisitUasRequest struct {
	EndTimestamp   *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeVisitUasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasRequest) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasRequest) SetEndTimestamp(v string) *DescribeVisitUasRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeVisitUasRequest) SetInstanceId(v string) *DescribeVisitUasRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVisitUasRequest) SetResource(v string) *DescribeVisitUasRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVisitUasRequest) SetStartTimestamp(v string) *DescribeVisitUasRequest {
	s.StartTimestamp = &v
	return s
}

type DescribeVisitUasResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uas       []*DescribeVisitUasResponseBodyUas `json:"Uas,omitempty" xml:"Uas,omitempty" type:"Repeated"`
}

func (s DescribeVisitUasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBody) SetRequestId(v string) *DescribeVisitUasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVisitUasResponseBody) SetUas(v []*DescribeVisitUasResponseBodyUas) *DescribeVisitUasResponseBody {
	s.Uas = v
	return s
}

type DescribeVisitUasResponseBodyUas struct {
	Count *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Ua    *string `json:"Ua,omitempty" xml:"Ua,omitempty"`
}

func (s DescribeVisitUasResponseBodyUas) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponseBodyUas) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponseBodyUas) SetCount(v int64) *DescribeVisitUasResponseBodyUas {
	s.Count = &v
	return s
}

func (s *DescribeVisitUasResponseBodyUas) SetUa(v string) *DescribeVisitUasResponseBodyUas {
	s.Ua = &v
	return s
}

type DescribeVisitUasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVisitUasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVisitUasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVisitUasResponse) GoString() string {
	return s.String()
}

func (s *DescribeVisitUasResponse) SetHeaders(v map[string]*string) *DescribeVisitUasResponse {
	s.Headers = v
	return s
}

func (s *DescribeVisitUasResponse) SetStatusCode(v int32) *DescribeVisitUasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVisitUasResponse) SetBody(v *DescribeVisitUasResponseBody) *DescribeVisitUasResponse {
	s.Body = v
	return s
}

type DescribeWafSourceIpSegmentRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeWafSourceIpSegmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentRequest) SetInstanceId(v string) *DescribeWafSourceIpSegmentRequest {
	s.InstanceId = &v
	return s
}

type DescribeWafSourceIpSegmentResponseBody struct {
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WafSourceIp *DescribeWafSourceIpSegmentResponseBodyWafSourceIp `json:"WafSourceIp,omitempty" xml:"WafSourceIp,omitempty" type:"Struct"`
}

func (s DescribeWafSourceIpSegmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetRequestId(v string) *DescribeWafSourceIpSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBody) SetWafSourceIp(v *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) *DescribeWafSourceIpSegmentResponseBody {
	s.WafSourceIp = v
	return s
}

type DescribeWafSourceIpSegmentResponseBodyWafSourceIp struct {
	IPv4 []*string `json:"IPv4,omitempty" xml:"IPv4,omitempty" type:"Repeated"`
	IPv6 []*string `json:"IPv6,omitempty" xml:"IPv6,omitempty" type:"Repeated"`
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponseBodyWafSourceIp) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv4(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv4 = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponseBodyWafSourceIp) SetIPv6(v []*string) *DescribeWafSourceIpSegmentResponseBodyWafSourceIp {
	s.IPv6 = v
	return s
}

type DescribeWafSourceIpSegmentResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWafSourceIpSegmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWafSourceIpSegmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWafSourceIpSegmentResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafSourceIpSegmentResponse) SetHeaders(v map[string]*string) *DescribeWafSourceIpSegmentResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetStatusCode(v int32) *DescribeWafSourceIpSegmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafSourceIpSegmentResponse) SetBody(v *DescribeWafSourceIpSegmentResponseBody) *DescribeWafSourceIpSegmentResponse {
	s.Body = v
	return s
}

type ModifyDefenseResourceGroupRequest struct {
	AddList     *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	DeleteList  *string `json:"DeleteList,omitempty" xml:"DeleteList,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyDefenseResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupRequest) SetAddList(v string) *ModifyDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDeleteList(v string) *ModifyDefenseResourceGroupRequest {
	s.DeleteList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDescription(v string) *ModifyDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetGroupName(v string) *ModifyDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetInstanceId(v string) *ModifyDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

type ModifyDefenseResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponseBody) SetRequestId(v string) *ModifyDefenseResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseResourceGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupResponse) SetHeaders(v map[string]*string) *ModifyDefenseResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetStatusCode(v int32) *ModifyDefenseResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseResourceGroupResponse) SetBody(v *ModifyDefenseResourceGroupResponseBody) *ModifyDefenseResourceGroupResponse {
	s.Body = v
	return s
}

type ModifyDefenseRuleRequest struct {
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Rules        *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleRequest) SetDefenseScene(v string) *ModifyDefenseRuleRequest {
	s.DefenseScene = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetInstanceId(v string) *ModifyDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetRules(v string) *ModifyDefenseRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyDefenseRuleRequest) SetTemplateId(v int64) *ModifyDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

type ModifyDefenseRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponseBody) SetRequestId(v string) *ModifyDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleResponse) SetStatusCode(v int32) *ModifyDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleResponse) SetBody(v *ModifyDefenseRuleResponseBody) *ModifyDefenseRuleResponse {
	s.Body = v
	return s
}

type ModifyDefenseRuleStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleStatus *int32  `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusRequest) SetInstanceId(v string) *ModifyDefenseRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleId(v int64) *ModifyDefenseRuleStatusRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleStatus(v int32) *ModifyDefenseRuleStatusRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetTemplateId(v int64) *ModifyDefenseRuleStatusRequest {
	s.TemplateId = &v
	return s
}

type ModifyDefenseRuleStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseRuleStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponseBody) SetRequestId(v string) *ModifyDefenseRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseRuleStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseRuleStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetStatusCode(v int32) *ModifyDefenseRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetBody(v *ModifyDefenseRuleStatusResponseBody) *ModifyDefenseRuleStatusResponse {
	s.Body = v
	return s
}

type ModifyDefenseTemplateRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyDefenseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateRequest) SetDescription(v string) *ModifyDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetInstanceId(v string) *ModifyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateId(v int64) *ModifyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateName(v string) *ModifyDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

type ModifyDefenseTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponseBody) SetRequestId(v string) *ModifyDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetStatusCode(v int32) *ModifyDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetBody(v *ModifyDefenseTemplateResponseBody) *ModifyDefenseTemplateResponse {
	s.Body = v
	return s
}

type ModifyDefenseTemplateStatusRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId     *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateStatus *int32  `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s ModifyDefenseTemplateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusRequest) SetInstanceId(v string) *ModifyDefenseTemplateStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateId(v int64) *ModifyDefenseTemplateStatusRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateStatusRequest) SetTemplateStatus(v int32) *ModifyDefenseTemplateStatusRequest {
	s.TemplateStatus = &v
	return s
}

type ModifyDefenseTemplateStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefenseTemplateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponseBody) SetRequestId(v string) *ModifyDefenseTemplateStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefenseTemplateStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDefenseTemplateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDefenseTemplateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetStatusCode(v int32) *ModifyDefenseTemplateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetBody(v *ModifyDefenseTemplateStatusResponseBody) *ModifyDefenseTemplateStatusResponse {
	s.Body = v
	return s
}

type ModifyDomainRequest struct {
	AccessType *string                      `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Domain     *string                      `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId *string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Listen     *ModifyDomainRequestListen   `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	Redirect   *ModifyDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	RegionId   *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequest) SetAccessType(v string) *ModifyDomainRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainRequest) SetDomain(v string) *ModifyDomainRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainRequest) SetInstanceId(v string) *ModifyDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainRequest) SetListen(v *ModifyDomainRequestListen) *ModifyDomainRequest {
	s.Listen = v
	return s
}

func (s *ModifyDomainRequest) SetRedirect(v *ModifyDomainRequestRedirect) *ModifyDomainRequest {
	s.Redirect = v
	return s
}

func (s *ModifyDomainRequest) SetRegionId(v string) *ModifyDomainRequest {
	s.RegionId = &v
	return s
}

type ModifyDomainRequestListen struct {
	CertId             *string   `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CipherSuite        *int32    `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	CustomCiphers      []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	EnableTLSv3        *bool     `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	ExclusiveIp        *bool     `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	FocusHttps         *bool     `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	Http2Enabled       *bool     `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	HttpPorts          []*int32  `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	HttpsPorts         []*int32  `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	IPv6Enabled        *bool     `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	ProtectionResource *string   `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	TLSVersion         *string   `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	XffHeaderMode      *int32    `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	XffHeaders         []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s ModifyDomainRequestListen) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestListen) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestListen) SetCertId(v string) *ModifyDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCipherSuite(v int32) *ModifyDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCustomCiphers(v []*string) *ModifyDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *ModifyDomainRequestListen) SetEnableTLSv3(v bool) *ModifyDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyDomainRequestListen) SetExclusiveIp(v bool) *ModifyDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *ModifyDomainRequestListen) SetFocusHttps(v bool) *ModifyDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttp2Enabled(v bool) *ModifyDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpsPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetIPv6Enabled(v bool) *ModifyDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetProtectionResource(v string) *ModifyDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *ModifyDomainRequestListen) SetTLSVersion(v string) *ModifyDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaderMode(v int32) *ModifyDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaders(v []*string) *ModifyDomainRequestListen {
	s.XffHeaders = v
	return s
}

type ModifyDomainRequestRedirect struct {
	Backends          []*string                                    `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	ConnectTimeout    *int32                                       `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	FocusHttpBackend  *bool                                        `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	Keepalive         *bool                                        `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	KeepaliveRequests *int32                                       `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	KeepaliveTimeout  *int32                                       `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	Loadbalance       *string                                      `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	ReadTimeout       *int32                                       `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	RequestHeaders    []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	Retry             *bool                                        `json:"Retry,omitempty" xml:"Retry,omitempty"`
	SniEnabled        *bool                                        `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	SniHost           *string                                      `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	WriteTimeout      *int32                                       `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s ModifyDomainRequestRedirect) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirect) SetBackends(v []*string) *ModifyDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetConnectTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetFocusHttpBackend(v bool) *ModifyDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepalive(v bool) *ModifyDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveRequests(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveTimeout(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetLoadbalance(v string) *ModifyDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetReadTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRequestHeaders(v []*ModifyDomainRequestRedirectRequestHeaders) *ModifyDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRetry(v bool) *ModifyDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniEnabled(v bool) *ModifyDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniHost(v string) *ModifyDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetWriteTimeout(v int32) *ModifyDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

type ModifyDomainRequestRedirectRequestHeaders struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDomainRequestRedirectRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetKey(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetValue(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

type ModifyDomainShrinkRequest struct {
	AccessType     *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ListenShrink   *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainShrinkRequest) SetAccessType(v string) *ModifyDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetDomain(v string) *ModifyDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetInstanceId(v string) *ModifyDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetListenShrink(v string) *ModifyDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRedirectShrink(v string) *ModifyDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRegionId(v string) *ModifyDomainShrinkRequest {
	s.RegionId = &v
	return s
}

type ModifyDomainResponseBody struct {
	DomainInfo *ModifyDomainResponseBodyDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBody) SetDomainInfo(v *ModifyDomainResponseBodyDomainInfo) *ModifyDomainResponseBody {
	s.DomainInfo = v
	return s
}

func (s *ModifyDomainResponseBody) SetRequestId(v string) *ModifyDomainResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainResponseBodyDomainInfo struct {
	Cname  *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s ModifyDomainResponseBodyDomainInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponseBodyDomainInfo) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponseBodyDomainInfo) SetCname(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Cname = &v
	return s
}

func (s *ModifyDomainResponseBodyDomainInfo) SetDomain(v string) *ModifyDomainResponseBodyDomainInfo {
	s.Domain = &v
	return s
}

type ModifyDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainResponse) SetHeaders(v map[string]*string) *ModifyDomainResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainResponse) SetStatusCode(v int32) *ModifyDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDomainResponse) SetBody(v *ModifyDomainResponseBody) *ModifyDomainResponse {
	s.Body = v
	return s
}

type ModifyMajorProtectionBlackIpRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpList      *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpRequest) SetDescription(v string) *ModifyMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetInstanceId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetIpList(v string) *ModifyMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetRuleId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

type ModifyMajorProtectionBlackIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ModifyMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMajorProtectionBlackIpResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMajorProtectionBlackIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMajorProtectionBlackIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponse) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponse) SetHeaders(v map[string]*string) *ModifyMajorProtectionBlackIpResponse {
	s.Headers = v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetStatusCode(v int32) *ModifyMajorProtectionBlackIpResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponse) SetBody(v *ModifyMajorProtectionBlackIpResponseBody) *ModifyMajorProtectionBlackIpResponse {
	s.Body = v
	return s
}

type ModifyResourceLogStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resource   *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Status     *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusRequest) SetInstanceId(v string) *ModifyResourceLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetResource(v string) *ModifyResourceLogStatusRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogStatusRequest) SetStatus(v bool) *ModifyResourceLogStatusRequest {
	s.Status = &v
	return s
}

type ModifyResourceLogStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponseBody) SetRequestId(v string) *ModifyResourceLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyResourceLogStatusResponseBody) SetStatus(v bool) *ModifyResourceLogStatusResponseBody {
	s.Status = &v
	return s
}

type ModifyResourceLogStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyResourceLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyResourceLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyResourceLogStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogStatusResponse) SetHeaders(v map[string]*string) *ModifyResourceLogStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetStatusCode(v int32) *ModifyResourceLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceLogStatusResponse) SetBody(v *ModifyResourceLogStatusResponseBody) *ModifyResourceLogStatusResponse {
	s.Body = v
	return s
}

type ModifyTemplateResourcesRequest struct {
	BindResourceGroups   []*string `json:"BindResourceGroups,omitempty" xml:"BindResourceGroups,omitempty" type:"Repeated"`
	BindResources        []*string `json:"BindResources,omitempty" xml:"BindResources,omitempty" type:"Repeated"`
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId           *int64    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	UnbindResources      []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
}

func (s ModifyTemplateResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesRequest) SetBindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetBindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.BindResources = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetInstanceId(v string) *ModifyTemplateResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetTemplateId(v int64) *ModifyTemplateResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResourceGroups(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResourceGroups = v
	return s
}

func (s *ModifyTemplateResourcesRequest) SetUnbindResources(v []*string) *ModifyTemplateResourcesRequest {
	s.UnbindResources = v
	return s
}

type ModifyTemplateResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTemplateResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponseBody) SetRequestId(v string) *ModifyTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTemplateResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTemplateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTemplateResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTemplateResourcesResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResourcesResponse) SetHeaders(v map[string]*string) *ModifyTemplateResourcesResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetStatusCode(v int32) *ModifyTemplateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateResourcesResponse) SetBody(v *ModifyTemplateResourcesResponseBody) *ModifyTemplateResourcesResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-heyuan":             tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-wulanchabu":         tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("waf-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ClearMajorProtectionBlackIpWithOptions(request *ClearMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearMajorProtectionBlackIp(request *ClearMajorProtectionBlackIpRequest) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.ClearMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseResourceGroupWithOptions(request *CreateDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddList)) {
		query["AddList"] = request.AddList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseResourceGroup(request *CreateDefenseResourceGroupRequest) (_result *CreateDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CreateDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseRuleWithOptions(request *CreateDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseRule(request *CreateDefenseRuleRequest) (_result *CreateDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CreateDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefenseTemplateWithOptions(request *CreateDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateOrigin)) {
		query["TemplateOrigin"] = request.TemplateOrigin
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStatus)) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDefenseTemplate(request *CreateDefenseTemplateRequest) (_result *CreateDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CreateDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(tmpReq *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Listen)) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, tea.String("Listen"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Redirect)) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, tea.String("Redirect"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenShrink)) {
		query["Listen"] = request.ListenShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectShrink)) {
		query["Redirect"] = request.RedirectShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMajorProtectionBlackIpWithOptions(request *CreateMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMajorProtectionBlackIp(request *CreateMajorProtectionBlackIpRequest) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CreateMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseResourceGroupWithOptions(request *DeleteDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseResourceGroup(request *DeleteDefenseResourceGroupRequest) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.DeleteDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseRuleWithOptions(request *DeleteDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIds)) {
		query["RuleIds"] = request.RuleIds
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseRule(request *DeleteDefenseRuleRequest) (_result *DeleteDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.DeleteDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDefenseTemplateWithOptions(request *DeleteDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDefenseTemplate(request *DeleteDefenseTemplateRequest) (_result *DeleteDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.DeleteDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMajorProtectionBlackIpWithOptions(request *DeleteMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMajorProtectionBlackIp(request *DeleteMajorProtectionBlackIpRequest) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.DeleteMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroupWithOptions(request *DescribeDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResourceGroup(request *DescribeDefenseResourceGroupRequest) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.DescribeDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseResourcesWithOptions(request *DescribeDefenseResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseResources(request *DescribeDefenseResourcesRequest) (_result *DescribeDefenseResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.DescribeDefenseResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseRuleWithOptions(request *DescribeDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseRule(request *DescribeDefenseRuleRequest) (_result *DescribeDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.DescribeDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseRulesWithOptions(request *DescribeDefenseRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseRules"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseRules(request *DescribeDefenseRulesRequest) (_result *DescribeDefenseRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.DescribeDefenseRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseTemplateWithOptions(request *DescribeDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseTemplate(request *DescribeDefenseTemplateRequest) (_result *DescribeDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.DescribeDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDetailWithOptions(request *DescribeDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainDetail"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDetail(request *DescribeDomainDetailRequest) (_result *DescribeDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.DescribeDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Backend)) {
		query["Backend"] = request.Backend
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomains"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowChartWithOptions(request *DescribeFlowChartRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowChartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowChart"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowChart(request *DescribeFlowChartRequest) (_result *DescribeFlowChartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.DescribeFlowChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowTopResourceWithOptions(request *DescribeFlowTopResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowTopResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowTopResource"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowTopResource(request *DescribeFlowTopResourceRequest) (_result *DescribeFlowTopResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.DescribeFlowTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowTopUrlWithOptions(request *DescribeFlowTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlowTopUrl"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowTopUrl(request *DescribeFlowTopUrlRequest) (_result *DescribeFlowTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.DescribeFlowTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMajorProtectionBlackIpsWithOptions(request *DescribeMajorProtectionBlackIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpLike)) {
		query["IpLike"] = request.IpLike
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMajorProtectionBlackIps"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMajorProtectionBlackIps(request *DescribeMajorProtectionBlackIpsRequest) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.DescribeMajorProtectionBlackIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePeakTrendWithOptions(request *DescribePeakTrendRequest, runtime *util.RuntimeOptions) (_result *DescribePeakTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePeakTrend"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePeakTrend(request *DescribePeakTrendRequest) (_result *DescribePeakTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.DescribePeakTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceLogStatusWithOptions(request *DescribeResourceLogStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceLogStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceLogStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceLogStatus(request *DescribeResourceLogStatusRequest) (_result *DescribeResourceLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.DescribeResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePortWithOptions(request *DescribeResourcePortRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePortResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceInstanceId)) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourcePort"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePort(request *DescribeResourcePortRequest) (_result *DescribeResourcePortResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.DescribeResourcePortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResponseCodeTrendGraphWithOptions(request *DescribeResponseCodeTrendGraphRequest, runtime *util.RuntimeOptions) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResponseCodeTrendGraph"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResponseCodeTrendGraph(request *DescribeResponseCodeTrendGraphRequest) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.DescribeResponseCodeTrendGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleGroupsWithOptions(request *DescribeRuleGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleGroups"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleGroups(request *DescribeRuleGroupsRequest) (_result *DescribeRuleGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.DescribeRuleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopClientIpWithOptions(request *DescribeRuleHitsTopClientIpRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopClientIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopClientIp(request *DescribeRuleHitsTopClientIpRequest) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.DescribeRuleHitsTopClientIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopResourceWithOptions(request *DescribeRuleHitsTopResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopResource"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopResource(request *DescribeRuleHitsTopResourceRequest) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.DescribeRuleHitsTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopRuleIdWithOptions(request *DescribeRuleHitsTopRuleIdRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopRuleId"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopRuleId(request *DescribeRuleHitsTopRuleIdRequest) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.DescribeRuleHitsTopRuleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopTuleTypeWithOptions(request *DescribeRuleHitsTopTuleTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopTuleType"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopTuleType(request *DescribeRuleHitsTopTuleTypeRequest) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.DescribeRuleHitsTopTuleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUaWithOptions(request *DescribeRuleHitsTopUaRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopUa"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUa(request *DescribeRuleHitsTopUaRequest) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.DescribeRuleHitsTopUaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUrlWithOptions(request *DescribeRuleHitsTopUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRuleHitsTopUrl"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRuleHitsTopUrl(request *DescribeRuleHitsTopUrlRequest) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.DescribeRuleHitsTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateResourcesWithOptions(request *DescribeTemplateResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTemplateResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateResources(request *DescribeTemplateResourcesRequest) (_result *DescribeTemplateResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.DescribeTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVisitTopIpWithOptions(request *DescribeVisitTopIpRequest, runtime *util.RuntimeOptions) (_result *DescribeVisitTopIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVisitTopIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVisitTopIp(request *DescribeVisitTopIpRequest) (_result *DescribeVisitTopIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.DescribeVisitTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVisitUasWithOptions(request *DescribeVisitUasRequest, runtime *util.RuntimeOptions) (_result *DescribeVisitUasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVisitUas"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVisitUas(request *DescribeVisitUasRequest) (_result *DescribeVisitUasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.DescribeVisitUasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWafSourceIpSegmentWithOptions(request *DescribeWafSourceIpSegmentRequest, runtime *util.RuntimeOptions) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWafSourceIpSegment"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWafSourceIpSegment(request *DescribeWafSourceIpSegmentRequest) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.DescribeWafSourceIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseResourceGroupWithOptions(request *ModifyDefenseResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddList)) {
		query["AddList"] = request.AddList
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteList)) {
		query["DeleteList"] = request.DeleteList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseResourceGroup"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseResourceGroup(request *ModifyDefenseResourceGroupRequest) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.ModifyDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseRuleWithOptions(request *ModifyDefenseRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefenseScene)) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseRule"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseRule(request *ModifyDefenseRuleRequest) (_result *ModifyDefenseRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.ModifyDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseRuleStatusWithOptions(request *ModifyDefenseRuleStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseRuleStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleStatus)) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseRuleStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseRuleStatus(request *ModifyDefenseRuleStatusRequest) (_result *ModifyDefenseRuleStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.ModifyDefenseRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateWithOptions(request *ModifyDefenseTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseTemplate"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseTemplate(request *ModifyDefenseTemplateRequest) (_result *ModifyDefenseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.ModifyDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateStatusWithOptions(request *ModifyDefenseTemplateStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStatus)) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefenseTemplateStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefenseTemplateStatus(request *ModifyDefenseTemplateStatusRequest) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.ModifyDefenseTemplateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainWithOptions(tmpReq *ModifyDomainRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Listen)) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, tea.String("Listen"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Redirect)) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, tea.String("Redirect"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessType)) {
		query["AccessType"] = request.AccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ListenShrink)) {
		query["Listen"] = request.ListenShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectShrink)) {
		query["Redirect"] = request.RedirectShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDomain"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomain(request *ModifyDomainRequest) (_result *ModifyDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainResponse{}
	_body, _err := client.ModifyDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMajorProtectionBlackIpWithOptions(request *ModifyMajorProtectionBlackIpRequest, runtime *util.RuntimeOptions) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMajorProtectionBlackIp"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMajorProtectionBlackIp(request *ModifyMajorProtectionBlackIpRequest) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.ModifyMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyResourceLogStatusWithOptions(request *ModifyResourceLogStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyResourceLogStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyResourceLogStatus"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyResourceLogStatus(request *ModifyResourceLogStatusRequest) (_result *ModifyResourceLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.ModifyResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTemplateResourcesWithOptions(request *ModifyTemplateResourcesRequest, runtime *util.RuntimeOptions) (_result *ModifyTemplateResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BindResourceGroups)) {
		query["BindResourceGroups"] = request.BindResourceGroups
	}

	if !tea.BoolValue(util.IsUnset(request.BindResources)) {
		query["BindResources"] = request.BindResources
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindResourceGroups)) {
		query["UnbindResourceGroups"] = request.UnbindResourceGroups
	}

	if !tea.BoolValue(util.IsUnset(request.UnbindResources)) {
		query["UnbindResources"] = request.UnbindResources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTemplateResources"),
		Version:     tea.String("2021-10-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTemplateResources(request *ModifyTemplateResourcesRequest) (_result *ModifyTemplateResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.ModifyTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
