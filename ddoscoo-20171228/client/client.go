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

type AddLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s AddLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleRequest) SetSourceIp(v string) *AddLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetResourceGroupId(v string) *AddLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetDomain(v string) *AddLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetName(v string) *AddLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetAct(v string) *AddLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetCount(v int32) *AddLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetInterval(v int32) *AddLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetMode(v string) *AddLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetTtl(v int32) *AddLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetUri(v string) *AddLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

type AddLayer7CCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLayer7CCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleResponseBody) SetRequestId(v string) *AddLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

type AddLayer7CCRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleResponse) SetHeaders(v map[string]*string) *AddLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *AddLayer7CCRuleResponse) SetBody(v *AddLayer7CCRuleResponseBody) *AddLayer7CCRuleResponse {
	s.Body = v
	return s
}

type CloseDomainSlsConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CloseDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigRequest) SetSourceIp(v string) *CloseDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetLang(v string) *CloseDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetResourceGroupId(v string) *CloseDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetDomain(v string) *CloseDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

type CloseDomainSlsConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDomainSlsConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigResponseBody) SetRequestId(v string) *CloseDomainSlsConfigResponseBody {
	s.RequestId = &v
	return s
}

type CloseDomainSlsConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseDomainSlsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigResponse) SetHeaders(v map[string]*string) *CloseDomainSlsConfigResponse {
	s.Headers = v
	return s
}

func (s *CloseDomainSlsConfigResponse) SetBody(v *CloseDomainSlsConfigResponseBody) *CloseDomainSlsConfigResponse {
	s.Body = v
	return s
}

type ConfigDomainAccessModeRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	AccessMode *int32  `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
}

func (s ConfigDomainAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *ConfigDomainAccessModeRequest) SetSourceIp(v string) *ConfigDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigDomainAccessModeRequest) SetDomain(v string) *ConfigDomainAccessModeRequest {
	s.Domain = &v
	return s
}

func (s *ConfigDomainAccessModeRequest) SetAccessMode(v int32) *ConfigDomainAccessModeRequest {
	s.AccessMode = &v
	return s
}

type ConfigDomainAccessModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigDomainAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigDomainAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigDomainAccessModeResponseBody) SetRequestId(v string) *ConfigDomainAccessModeResponseBody {
	s.RequestId = &v
	return s
}

type ConfigDomainAccessModeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigDomainAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigDomainAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigDomainAccessModeResponse) GoString() string {
	return s.String()
}

func (s *ConfigDomainAccessModeResponse) SetHeaders(v map[string]*string) *ConfigDomainAccessModeResponse {
	s.Headers = v
	return s
}

func (s *ConfigDomainAccessModeResponse) SetBody(v *ConfigDomainAccessModeResponseBody) *ConfigDomainAccessModeResponse {
	s.Body = v
	return s
}

type ConfigHealthCheckRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	HealthCheck     *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
}

func (s ConfigHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckRequest) SetSourceIp(v string) *ConfigHealthCheckRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetInstanceId(v string) *ConfigHealthCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetForwardProtocol(v string) *ConfigHealthCheckRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetFrontendPort(v int32) *ConfigHealthCheckRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigHealthCheckRequest) SetHealthCheck(v string) *ConfigHealthCheckRequest {
	s.HealthCheck = &v
	return s
}

type ConfigHealthCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigHealthCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckResponseBody) SetRequestId(v string) *ConfigHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

type ConfigHealthCheckResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigHealthCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckResponse) SetHeaders(v map[string]*string) *ConfigHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *ConfigHealthCheckResponse) SetBody(v *ConfigHealthCheckResponseBody) *ConfigHealthCheckResponse {
	s.Body = v
	return s
}

type ConfigLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s ConfigLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleRequest) SetSourceIp(v string) *ConfigLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer4RuleRequest) SetListeners(v string) *ConfigLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type ConfigLayer4RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleResponseBody) SetRequestId(v string) *ConfigLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleResponse) SetHeaders(v map[string]*string) *ConfigLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RuleResponse) SetBody(v *ConfigLayer4RuleResponseBody) *ConfigLayer4RuleResponse {
	s.Body = v
	return s
}

type ConfigLayer4RuleAttributeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Config          *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ConfigLayer4RuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeRequest) SetSourceIp(v string) *ConfigLayer4RuleAttributeRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetInstanceId(v string) *ConfigLayer4RuleAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetForwardProtocol(v string) *ConfigLayer4RuleAttributeRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetFrontendPort(v int32) *ConfigLayer4RuleAttributeRequest {
	s.FrontendPort = &v
	return s
}

func (s *ConfigLayer4RuleAttributeRequest) SetConfig(v string) *ConfigLayer4RuleAttributeRequest {
	s.Config = &v
	return s
}

type ConfigLayer4RuleAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeResponseBody) SetRequestId(v string) *ConfigLayer4RuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer4RuleAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer4RuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer4RuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeResponse) SetHeaders(v map[string]*string) *ConfigLayer4RuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RuleAttributeResponse) SetBody(v *ConfigLayer4RuleAttributeResponseBody) *ConfigLayer4RuleAttributeResponse {
	s.Body = v
	return s
}

type ConfigLayer7BlackWhiteListRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	BlackList       []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	WhiteList       []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigLayer7BlackWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListRequest) SetSourceIp(v string) *ConfigLayer7BlackWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetResourceGroupId(v string) *ConfigLayer7BlackWhiteListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetDomain(v string) *ConfigLayer7BlackWhiteListRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetBlackList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.BlackList = v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetWhiteList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.WhiteList = v
	return s
}

type ConfigLayer7BlackWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7BlackWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListResponseBody) SetRequestId(v string) *ConfigLayer7BlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer7BlackWhiteListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer7BlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer7BlackWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListResponse) SetHeaders(v map[string]*string) *ConfigLayer7BlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7BlackWhiteListResponse) SetBody(v *ConfigLayer7BlackWhiteListResponseBody) *ConfigLayer7BlackWhiteListResponse {
	s.Body = v
	return s
}

type ConfigLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Act             *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode            *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Uri             *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ConfigLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleRequest) SetSourceIp(v string) *ConfigLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetResourceGroupId(v string) *ConfigLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetDomain(v string) *ConfigLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetName(v string) *ConfigLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetAct(v string) *ConfigLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetCount(v int32) *ConfigLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetInterval(v int32) *ConfigLayer7CCRuleRequest {
	s.Interval = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetMode(v string) *ConfigLayer7CCRuleRequest {
	s.Mode = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetTtl(v int32) *ConfigLayer7CCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetUri(v string) *ConfigLayer7CCRuleRequest {
	s.Uri = &v
	return s
}

type ConfigLayer7CCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7CCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleResponseBody) SetRequestId(v string) *ConfigLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer7CCRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleResponse) SetHeaders(v map[string]*string) *ConfigLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CCRuleResponse) SetBody(v *ConfigLayer7CCRuleResponseBody) *ConfigLayer7CCRuleResponse {
	s.Body = v
	return s
}

type ConfigLayer7CCTemplateRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Template        *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ConfigLayer7CCTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateRequest) SetSourceIp(v string) *ConfigLayer7CCTemplateRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetResourceGroupId(v string) *ConfigLayer7CCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetDomain(v string) *ConfigLayer7CCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetTemplate(v string) *ConfigLayer7CCTemplateRequest {
	s.Template = &v
	return s
}

type ConfigLayer7CCTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7CCTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateResponseBody) SetRequestId(v string) *ConfigLayer7CCTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer7CCTemplateResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer7CCTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer7CCTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateResponse) SetHeaders(v map[string]*string) *ConfigLayer7CCTemplateResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CCTemplateResponse) SetBody(v *ConfigLayer7CCTemplateResponseBody) *ConfigLayer7CCTemplateResponse {
	s.Body = v
	return s
}

type ConfigLayer7CertRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CertId          *int32  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cert            *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ConfigLayer7CertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertRequest) SetSourceIp(v string) *ConfigLayer7CertRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetResourceGroupId(v string) *ConfigLayer7CertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetDomain(v string) *ConfigLayer7CertRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertId(v int32) *ConfigLayer7CertRequest {
	s.CertId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertName(v string) *ConfigLayer7CertRequest {
	s.CertName = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCert(v string) *ConfigLayer7CertRequest {
	s.Cert = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetKey(v string) *ConfigLayer7CertRequest {
	s.Key = &v
	return s
}

type ConfigLayer7CertResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7CertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertResponseBody) SetRequestId(v string) *ConfigLayer7CertResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer7CertResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer7CertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer7CertResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertResponse) SetHeaders(v map[string]*string) *ConfigLayer7CertResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CertResponse) SetBody(v *ConfigLayer7CertResponseBody) *ConfigLayer7CertResponse {
	s.Body = v
	return s
}

type ConfigLayer7RuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyTypeList   *string   `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty"`
	RsType          *int32    `json:"RsType,omitempty" xml:"RsType,omitempty"`
	ProxyTypes      []*string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	RealServers     []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s ConfigLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleRequest) SetSourceIp(v string) *ConfigLayer7RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetResourceGroupId(v string) *ConfigLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetDomain(v string) *ConfigLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypeList(v string) *ConfigLayer7RuleRequest {
	s.ProxyTypeList = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRsType(v int32) *ConfigLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypes(v []*string) *ConfigLayer7RuleRequest {
	s.ProxyTypes = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRealServers(v []*string) *ConfigLayer7RuleRequest {
	s.RealServers = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetInstanceIds(v []*string) *ConfigLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

type ConfigLayer7RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer7RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleResponseBody) SetRequestId(v string) *ConfigLayer7RuleResponseBody {
	s.RequestId = &v
	return s
}

type ConfigLayer7RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleResponse) SetHeaders(v map[string]*string) *ConfigLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7RuleResponse) SetBody(v *ConfigLayer7RuleResponseBody) *ConfigLayer7RuleResponse {
	s.Body = v
	return s
}

type CreateAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskType        *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskParams      *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) SetSourceIp(v string) *CreateAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetLang(v string) *CreateAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int32) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

type CreateAsyncTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponseBody) SetRequestId(v string) *CreateAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateAsyncTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponse) SetHeaders(v map[string]*string) *CreateAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAsyncTaskResponse) SetBody(v *CreateAsyncTaskResponseBody) *CreateAsyncTaskResponse {
	s.Body = v
	return s
}

type CreateLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s CreateLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequest) SetSourceIp(v string) *CreateLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetListeners(v string) *CreateLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type CreateLayer4RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponseBody) SetRequestId(v string) *CreateLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleResponse) SetHeaders(v map[string]*string) *CreateLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer4RuleResponse) SetBody(v *CreateLayer4RuleResponseBody) *CreateLayer4RuleResponse {
	s.Body = v
	return s
}

type CreateLayer7RuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	RsType          *int32    `json:"RsType,omitempty" xml:"RsType,omitempty"`
	Rules           *string   `json:"Rules,omitempty" xml:"Rules,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleRequest) SetSourceIp(v string) *CreateLayer7RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetResourceGroupId(v string) *CreateLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetDomain(v string) *CreateLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRsType(v int32) *CreateLayer7RuleRequest {
	s.RsType = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetRules(v string) *CreateLayer7RuleRequest {
	s.Rules = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetInstanceIds(v []*string) *CreateLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

type CreateLayer7RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLayer7RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleResponseBody) SetRequestId(v string) *CreateLayer7RuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateLayer7RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleResponse) SetHeaders(v map[string]*string) *CreateLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *CreateLayer7RuleResponse) SetBody(v *CreateLayer7RuleResponseBody) *CreateLayer7RuleResponse {
	s.Body = v
	return s
}

type DeleteAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TaskId          *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskRequest) SetSourceIp(v string) *DeleteAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetLang(v string) *DeleteAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetResourceGroupId(v string) *DeleteAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteAsyncTaskRequest) SetTaskId(v int32) *DeleteAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteAsyncTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponseBody) SetRequestId(v string) *DeleteAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAsyncTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsyncTaskResponse) SetHeaders(v map[string]*string) *DeleteAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsyncTaskResponse) SetBody(v *DeleteAsyncTaskResponseBody) *DeleteAsyncTaskResponse {
	s.Body = v
	return s
}

type DeleteLayer4RuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DeleteLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleRequest) SetSourceIp(v string) *DeleteLayer4RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteLayer4RuleRequest) SetListeners(v string) *DeleteLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type DeleteLayer4RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayer4RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponseBody) SetRequestId(v string) *DeleteLayer4RuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLayer4RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLayer4RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleResponse) SetHeaders(v map[string]*string) *DeleteLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer4RuleResponse) SetBody(v *DeleteLayer4RuleResponseBody) *DeleteLayer4RuleResponse {
	s.Body = v
	return s
}

type DeleteLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleRequest) SetSourceIp(v string) *DeleteLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetResourceGroupId(v string) *DeleteLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetDomain(v string) *DeleteLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetName(v string) *DeleteLayer7CCRuleRequest {
	s.Name = &v
	return s
}

type DeleteLayer7CCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayer7CCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleResponseBody) SetRequestId(v string) *DeleteLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLayer7CCRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleResponse) SetHeaders(v map[string]*string) *DeleteLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer7CCRuleResponse) SetBody(v *DeleteLayer7CCRuleResponseBody) *DeleteLayer7CCRuleResponse {
	s.Body = v
	return s
}

type DeleteLayer7RuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleRequest) SetSourceIp(v string) *DeleteLayer7RuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteLayer7RuleRequest) SetResourceGroupId(v string) *DeleteLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteLayer7RuleRequest) SetDomain(v string) *DeleteLayer7RuleRequest {
	s.Domain = &v
	return s
}

type DeleteLayer7RuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayer7RuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleResponseBody) SetRequestId(v string) *DeleteLayer7RuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLayer7RuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLayer7RuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleResponse) SetHeaders(v map[string]*string) *DeleteLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLayer7RuleResponse) SetBody(v *DeleteLayer7RuleResponseBody) *DeleteLayer7RuleResponse {
	s.Body = v
	return s
}

type DescribeBackSourceCidrRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Line            *string `json:"Line,omitempty" xml:"Line,omitempty"`
}

func (s DescribeBackSourceCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrRequest) SetSourceIp(v string) *DescribeBackSourceCidrRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetResourceGroupId(v string) *DescribeBackSourceCidrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

type DescribeBackSourceCidrResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CidrList  []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
}

func (s DescribeBackSourceCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBody) SetRequestId(v string) *DescribeBackSourceCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) SetCidrList(v []*string) *DescribeBackSourceCidrResponseBody {
	s.CidrList = v
	return s
}

type DescribeBackSourceCidrResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackSourceCidrResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackSourceCidrResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponse) SetHeaders(v map[string]*string) *DescribeBackSourceCidrResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetBody(v *DescribeBackSourceCidrResponseBody) *DescribeBackSourceCidrResponse {
	s.Body = v
	return s
}

type DescribeBatchSlsDispatchStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetSourceIp(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetLang(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetResourceGroupId(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageNo(v int32) *DescribeBatchSlsDispatchStatusRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetPageSize(v int32) *DescribeBatchSlsDispatchStatusRequest {
	s.PageSize = &v
	return s
}

type DescribeBatchSlsDispatchStatusResponseBody struct {
	TotalCount          *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId           *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsConfigStatusList []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList `json:"SlsConfigStatusList,omitempty" xml:"SlsConfigStatusList,omitempty" type:"Repeated"`
}

func (s DescribeBatchSlsDispatchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetTotalCount(v int32) *DescribeBatchSlsDispatchStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetSlsConfigStatusList(v []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.SlsConfigStatusList = v
	return s
}

type DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) SetDomain(v string) *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList {
	s.Domain = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) SetEnable(v bool) *DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList {
	s.Enable = &v
	return s
}

type DescribeBatchSlsDispatchStatusResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBatchSlsDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBatchSlsDispatchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetHeaders(v map[string]*string) *DescribeBatchSlsDispatchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetBody(v *DescribeBatchSlsDispatchStatusResponseBody) *DescribeBatchSlsDispatchStatusResponse {
	s.Body = v
	return s
}

type DescribeDDoSEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) SetSourceIp(v string) *DescribeDDoSEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEip(v string) *DescribeDDoSEventsRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetOffset(v int32) *DescribeDDoSEventsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetPageSize(v string) *DescribeDDoSEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeDDoSEventsResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeDDoSEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	Total     *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDDoSEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBody) SetRequestId(v string) *DescribeDDoSEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetEvents(v []*DescribeDDoSEventsResponseBodyEvents) *DescribeDDoSEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetTotal(v int64) *DescribeDDoSEventsResponseBody {
	s.Total = &v
	return s
}

type DescribeDDoSEventsResponseBodyEvents struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeDDoSEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStatus(v string) *DescribeDDoSEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetInterval(v int32) *DescribeDDoSEventsResponseBodyEvents {
	s.Interval = &v
	return s
}

type DescribeDDoSEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDoSEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDoSEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponse) SetHeaders(v map[string]*string) *DescribeDDoSEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSEventsResponse) SetBody(v *DescribeDDoSEventsResponseBody) *DescribeDDoSEventsResponse {
	s.Body = v
	return s
}

type DescribeDDoSTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
}

func (s DescribeDDoSTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficRequest) SetSourceIp(v string) *DescribeDDoSTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetResourceGroupId(v string) *DescribeDDoSTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetStartTime(v int64) *DescribeDDoSTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetInterval(v int32) *DescribeDDoSTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEndTime(v int64) *DescribeDDoSTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEip(v string) *DescribeDDoSTrafficRequest {
	s.Eip = &v
	return s
}

type DescribeDDoSTrafficResponseBody struct {
	DefenseInBytes    *int64                                              `json:"DefenseInBytes,omitempty" xml:"DefenseInBytes,omitempty"`
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceInBytes     *int64                                              `json:"SourceInBytes,omitempty" xml:"SourceInBytes,omitempty"`
	DDoSTrafficPoints []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints `json:"DDoSTrafficPoints,omitempty" xml:"DDoSTrafficPoints,omitempty" type:"Repeated"`
}

func (s DescribeDDoSTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBody) SetDefenseInBytes(v int64) *DescribeDDoSTrafficResponseBody {
	s.DefenseInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetRequestId(v string) *DescribeDDoSTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetSourceInBytes(v int64) *DescribeDDoSTrafficResponseBody {
	s.SourceInBytes = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBody) SetDDoSTrafficPoints(v []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) *DescribeDDoSTrafficResponseBody {
	s.DDoSTrafficPoints = v
	return s
}

type DescribeDDoSTrafficResponseBodyDDoSTrafficPoints struct {
	Time            *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	SourceMaxInBps  *int64 `json:"SourceMaxInBps,omitempty" xml:"SourceMaxInBps,omitempty"`
	DefenseMaxInBps *int64 `json:"DefenseMaxInBps,omitempty" xml:"DefenseMaxInBps,omitempty"`
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetTime(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetSourceMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.SourceMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetDefenseMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.DefenseMaxInBps = &v
	return s
}

type DescribeDDoSTrafficResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDDoSTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDDoSTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponse) SetHeaders(v map[string]*string) *DescribeDDoSTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetBody(v *DescribeDDoSTrafficResponseBody) *DescribeDDoSTrafficResponse {
	s.Body = v
	return s
}

type DescribeDefenseCountStatisticsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDefenseCountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsRequest) SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeDefenseCountStatisticsResponseBody struct {
	RequestId              *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" type:"Struct"`
}

func (s DescribeDefenseCountStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetRequestId(v string) *DescribeDefenseCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) *DescribeDefenseCountStatisticsResponseBody {
	s.DefenseCountStatistics = v
	return s
}

type DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics struct {
	MaxUsableDefenseCountCurrentMonth    *int32 `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty"`
	FlowPackCountRemain                  *int32 `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty"`
	DefenseCountTotalUsageOfCurrentMonth *int32 `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetMaxUsableDefenseCountCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.MaxUsableDefenseCountCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetFlowPackCountRemain(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.FlowPackCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetDefenseCountTotalUsageOfCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.DefenseCountTotalUsageOfCurrentMonth = &v
	return s
}

type DescribeDefenseCountStatisticsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDefenseCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDefenseCountStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDefenseCountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetBody(v *DescribeDefenseCountStatisticsResponseBody) *DescribeDefenseCountStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDomainAccessModeRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
}

func (s DescribeDomainAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeRequest) SetSourceIp(v string) *DescribeDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAccessModeRequest) SetDomainList(v []*string) *DescribeDomainAccessModeRequest {
	s.DomainList = v
	return s
}

type DescribeDomainAccessModeResponseBody struct {
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainModeList []*DescribeDomainAccessModeResponseBodyDomainModeList `json:"DomainModeList,omitempty" xml:"DomainModeList,omitempty" type:"Repeated"`
}

func (s DescribeDomainAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseBody) SetRequestId(v string) *DescribeDomainAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAccessModeResponseBody) SetDomainModeList(v []*DescribeDomainAccessModeResponseBodyDomainModeList) *DescribeDomainAccessModeResponseBody {
	s.DomainModeList = v
	return s
}

type DescribeDomainAccessModeResponseBodyDomainModeList struct {
	AccessMode *int32  `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainAccessModeResponseBodyDomainModeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponseBodyDomainModeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) SetAccessMode(v int32) *DescribeDomainAccessModeResponseBodyDomainModeList {
	s.AccessMode = &v
	return s
}

func (s *DescribeDomainAccessModeResponseBodyDomainModeList) SetDomain(v string) *DescribeDomainAccessModeResponseBodyDomainModeList {
	s.Domain = &v
	return s
}

type DescribeDomainAccessModeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAccessModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponse) SetHeaders(v map[string]*string) *DescribeDomainAccessModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAccessModeResponse) SetBody(v *DescribeDomainAccessModeResponseBody) *DescribeDomainAccessModeResponse {
	s.Body = v
	return s
}

type DescribeDomainAttackEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainAttackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsRequest) SetSourceIp(v string) *DescribeDomainAttackEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetStartTime(v int64) *DescribeDomainAttackEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetEndTime(v int64) *DescribeDomainAttackEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetDomain(v string) *DescribeDomainAttackEventsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetOffset(v int32) *DescribeDomainAttackEventsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageSize(v string) *DescribeDomainAttackEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeDomainAttackEventsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeDomainAttackEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	Total     *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBody) SetRequestId(v string) *DescribeDomainAttackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetEvents(v []*DescribeDomainAttackEventsResponseBodyEvents) *DescribeDomainAttackEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDomainAttackEventsResponseBody) SetTotal(v int64) *DescribeDomainAttackEventsResponseBody {
	s.Total = &v
	return s
}

type DescribeDomainAttackEventsResponseBodyEvents struct {
	EndTime    *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	MaxQps     *int32 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	Duration   *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Finished   *bool  `json:"Finished,omitempty" xml:"Finished,omitempty"`
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
}

func (s DescribeDomainAttackEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetEndTime(v int64) *DescribeDomainAttackEventsResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetStartTime(v int64) *DescribeDomainAttackEventsResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetMaxQps(v int32) *DescribeDomainAttackEventsResponseBodyEvents {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetDuration(v int32) *DescribeDomainAttackEventsResponseBodyEvents {
	s.Duration = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetFinished(v bool) *DescribeDomainAttackEventsResponseBodyEvents {
	s.Finished = &v
	return s
}

func (s *DescribeDomainAttackEventsResponseBodyEvents) SetBlockCount(v int64) *DescribeDomainAttackEventsResponseBodyEvents {
	s.BlockCount = &v
	return s
}

type DescribeDomainAttackEventsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainAttackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainAttackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackEventsResponse) SetBody(v *DescribeDomainAttackEventsResponseBody) *DescribeDomainAttackEventsResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsRequest) SetSourceIp(v string) *DescribeDomainQpsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetResourceGroupId(v string) *DescribeDomainQpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetDomain(v string) *DescribeDomainQpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetStartTime(v int64) *DescribeDomainQpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsRequest) SetEndTime(v int64) *DescribeDomainQpsRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainQpsResponseBody struct {
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" type:"Repeated"`
	CcJsQps       []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" type:"Repeated"`
	Blocks        []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" type:"Repeated"`
	Totals        []*string `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Repeated"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CcBlockQps    []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" type:"Repeated"`
	Interval      *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" type:"Repeated"`
	CacheHits     []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsResponseBody) SetIpBlockQps(v []*string) *DescribeDomainQpsResponseBody {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetCcJsQps(v []*string) *DescribeDomainQpsResponseBody {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetBlocks(v []*string) *DescribeDomainQpsResponseBody {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetPreciseBlocks(v []*string) *DescribeDomainQpsResponseBody {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetRequestId(v string) *DescribeDomainQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetPreciseJsQps(v []*string) *DescribeDomainQpsResponseBody {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetTotals(v []*string) *DescribeDomainQpsResponseBody {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetStartTime(v int64) *DescribeDomainQpsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetCcBlockQps(v []*string) *DescribeDomainQpsResponseBody {
	s.CcBlockQps = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetInterval(v int32) *DescribeDomainQpsResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetRegionBlocks(v []*string) *DescribeDomainQpsResponseBody {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsResponseBody) SetCacheHits(v []*string) *DescribeDomainQpsResponseBody {
	s.CacheHits = v
	return s
}

type DescribeDomainQpsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsResponse) SetBody(v *DescribeDomainQpsResponseBody) *DescribeDomainQpsResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsWithCacheRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDomainQpsWithCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheRequest) SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetDomain(v string) *DescribeDomainQpsWithCacheRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.EndTime = &v
	return s
}

type DescribeDomainQpsWithCacheResponseBody struct {
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" type:"Repeated"`
	CcJsQps       []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" type:"Repeated"`
	Blocks        []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" type:"Repeated"`
	Totals        []*string `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Repeated"`
	StartTime     *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CcBlockQps    []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" type:"Repeated"`
	Interval      *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" type:"Repeated"`
	CacheHits     []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsWithCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRequestId(v string) *DescribeDomainQpsWithCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetTotals(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Totals = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetStartTime(v int64) *DescribeDomainQpsWithCacheResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetInterval(v int32) *DescribeDomainQpsWithCacheResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CacheHits = v
	return s
}

type DescribeDomainQpsWithCacheResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainQpsWithCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainQpsWithCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsWithCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetBody(v *DescribeDomainQpsWithCacheResponseBody) *DescribeDomainQpsWithCacheResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	SourceIp           *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId    *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain             *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	QueryDomainPattern *string   `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	Offset             *int32    `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize           *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceIds        []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetSourceIp(v string) *DescribeDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetQueryDomainPattern(v string) *DescribeDomainsRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeDomainsRequest) SetOffset(v int32) *DescribeDomainsRequest {
	s.Offset = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v string) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains   []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *DescribeDomainsResponseBody) SetTotal(v int64) *DescribeDomainsResponseBody {
	s.Total = &v
	return s
}

type DescribeDomainsResponseBodyDomains struct {
	Domain        *string                                            `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyTypeList []*DescribeDomainsResponseBodyDomainsProxyTypeList `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty" type:"Repeated"`
	WhiteList     []*string                                          `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
	BlackList     []*string                                          `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	RealServers   []*DescribeDomainsResponseBodyDomainsRealServers   `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	SslProtocols  *string                                            `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	CcTemplate    *string                                            `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	CcEnabled     *bool                                              `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	SslCiphers    *string                                            `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	CcRuleEnabled *bool                                              `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	CertName      *string                                            `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cname         *string                                            `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Http2Enable   *bool                                              `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v string) *DescribeDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetProxyTypeList(v []*DescribeDomainsResponseBodyDomainsProxyTypeList) *DescribeDomainsResponseBodyDomains {
	s.ProxyTypeList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetWhiteList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.WhiteList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetBlackList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.BlackList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetRealServers(v []*DescribeDomainsResponseBodyDomainsRealServers) *DescribeDomainsResponseBodyDomains {
	s.RealServers = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslProtocols(v string) *DescribeDomainsResponseBodyDomains {
	s.SslProtocols = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcTemplate(v string) *DescribeDomainsResponseBodyDomains {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslCiphers(v string) *DescribeDomainsResponseBodyDomains {
	s.SslCiphers = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcRuleEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCertName(v string) *DescribeDomainsResponseBodyDomains {
	s.CertName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCname(v string) *DescribeDomainsResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetHttp2Enable(v bool) *DescribeDomainsResponseBodyDomains {
	s.Http2Enable = &v
	return s
}

type DescribeDomainsResponseBodyDomainsProxyTypeList struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsProxyTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsProxyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) SetProxyPorts(v []*string) *DescribeDomainsResponseBodyDomainsProxyTypeList {
	s.ProxyPorts = v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsProxyTypeList) SetProxyType(v string) *DescribeDomainsResponseBodyDomainsProxyTypeList {
	s.ProxyType = &v
	return s
}

type DescribeDomainsResponseBodyDomainsRealServers struct {
	RsType     *int32  `json:"RsType,omitempty" xml:"RsType,omitempty"`
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsRealServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsRealServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRsType(v int32) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RsType = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRealServer(v string) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RealServer = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

type DescribeDomainSlsStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainSlsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusRequest) SetSourceIp(v string) *DescribeDomainSlsStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetLang(v string) *DescribeDomainSlsStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetResourceGroupId(v string) *DescribeDomainSlsStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainSlsStatusRequest) SetDomain(v string) *DescribeDomainSlsStatusRequest {
	s.Domain = &v
	return s
}

type DescribeDomainSlsStatusResponseBody struct {
	SlsProject  *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	SlsStatus   *bool   `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
}

func (s DescribeDomainSlsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsProject(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsStatus(v bool) *DescribeDomainSlsStatusResponseBody {
	s.SlsStatus = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetRequestId(v string) *DescribeDomainSlsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsLogstore(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsLogstore = &v
	return s
}

type DescribeDomainSlsStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSlsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSlsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponse) SetHeaders(v map[string]*string) *DescribeDomainSlsStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetBody(v *DescribeDomainSlsStatusResponseBody) *DescribeDomainSlsStatusResponse {
	s.Body = v
	return s
}

type DescribeElasticBandwidthSpecRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeElasticBandwidthSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecRequest) SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

type DescribeElasticBandwidthSpecResponseBody struct {
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" type:"Repeated"`
}

func (s DescribeElasticBandwidthSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetRequestId(v string) *DescribeElasticBandwidthSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponseBody {
	s.ElasticBandwidthSpec = v
	return s
}

type DescribeElasticBandwidthSpecResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeElasticBandwidthSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeElasticBandwidthSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponse) SetHeaders(v map[string]*string) *DescribeElasticBandwidthSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetBody(v *DescribeElasticBandwidthSpecResponseBody) *DescribeElasticBandwidthSpecResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DescribeHealthCheckListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListRequest) SetSourceIp(v string) *DescribeHealthCheckListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckListRequest) SetListeners(v string) *DescribeHealthCheckListRequest {
	s.Listeners = &v
	return s
}

type DescribeHealthCheckListResponseBody struct {
	Listeners []*DescribeHealthCheckListResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBody) SetListeners(v []*DescribeHealthCheckListResponseBodyListeners) *DescribeHealthCheckListResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeHealthCheckListResponseBody) SetRequestId(v string) *DescribeHealthCheckListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHealthCheckListResponseBodyListeners struct {
	FrontendPort *int32                                                   `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol     *string                                                  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId   *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HealthCheck  *DescribeHealthCheckListResponseBodyListenersHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
}

func (s DescribeHealthCheckListResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetFrontendPort(v int32) *DescribeHealthCheckListResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetProtocol(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetInstanceId(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetHealthCheck(v *DescribeHealthCheckListResponseBodyListenersHealthCheck) *DescribeHealthCheckListResponseBodyListeners {
	s.HealthCheck = v
	return s
}

type DescribeHealthCheckListResponseBodyListenersHealthCheck struct {
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Up       *int32  `json:"Up,omitempty" xml:"Up,omitempty"`
	Down     *int32  `json:"Down,omitempty" xml:"Down,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetTimeout(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetType(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetInterval(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUp(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDown(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetPort(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Uri = &v
	return s
}

type DescribeHealthCheckListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHealthCheckListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthCheckListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckListResponse) SetBody(v *DescribeHealthCheckListResponseBody) *DescribeHealthCheckListResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckStatusListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DescribeHealthCheckStatusListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListRequest) SetSourceIp(v string) *DescribeHealthCheckStatusListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckStatusListRequest) SetListeners(v string) *DescribeHealthCheckStatusListRequest {
	s.Listeners = &v
	return s
}

type DescribeHealthCheckStatusListResponseBody struct {
	HealthCheckStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList `json:"HealthCheckStatusList,omitempty" xml:"HealthCheckStatusList,omitempty" type:"Repeated"`
	RequestId             *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBody) SetHealthCheckStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) *DescribeHealthCheckStatusListResponseBody {
	s.HealthCheckStatusList = v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBody) SetRequestId(v string) *DescribeHealthCheckStatusListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList struct {
	Status               *string                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	FrontendPort         *int32                                                                                `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol             *string                                                                               `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId           *string                                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RealServerStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" type:"Repeated"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetFrontendPort(v int32) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetProtocol(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetInstanceId(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetRealServerStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.RealServerStatusList = v
	return s
}

type DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Status = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Address = &v
	return s
}

type DescribeHealthCheckStatusListResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHealthCheckStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthCheckStatusListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckStatusListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) SetBody(v *DescribeHealthCheckStatusListResponseBody) *DescribeHealthCheckStatusListResponse {
	s.Body = v
	return s
}

type DescribeInstanceDetailsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) SetSourceIp(v string) *DescribeInstanceDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceDetailsResponseBody struct {
	InstanceDetails []*DescribeInstanceDetailsResponseBodyInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBody) SetInstanceDetails(v []*DescribeInstanceDetailsResponseBodyInstanceDetails) *DescribeInstanceDetailsResponseBody {
	s.InstanceDetails = v
	return s
}

func (s *DescribeInstanceDetailsResponseBody) SetRequestId(v string) *DescribeInstanceDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceDetailsResponseBodyInstanceDetails struct {
	EipInfoList []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList `json:"EipInfoList,omitempty" xml:"EipInfoList,omitempty" type:"Repeated"`
	Line        *string                                                          `json:"Line,omitempty" xml:"Line,omitempty"`
	InstanceId  *string                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetEipInfoList(v []*DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.EipInfoList = v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.Line = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.InstanceId = &v
	return s
}

type DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Eip    *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetStatus(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetEip(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Eip = &v
	return s
}

type DescribeInstanceDetailsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetBody(v *DescribeInstanceDetailsResponseBody) *DescribeInstanceDetailsResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	SourceIp        *string                        `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     *string                        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	PageNo          *string                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *string                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Ip              *string                        `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Remark          *string                        `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Edition         *int32                         `json:"Edition,omitempty" xml:"Edition,omitempty"`
	Enabled         *int32                         `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ExpireStartTime *int64                         `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	ExpireEndTime   *int64                         `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	Status          []*int                         `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNo(v string) *DescribeInstancesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesRequest) SetRemark(v string) *DescribeInstancesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesRequest) SetEdition(v int32) *DescribeInstancesRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnabled(v int32) *DescribeInstancesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireStartTime(v int64) *DescribeInstancesRequest {
	s.ExpireStartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireEndTime(v int64) *DescribeInstancesRequest {
	s.ExpireEndTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v []*int) *DescribeInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotal(v int64) *DescribeInstancesResponseBody {
	s.Total = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DebtStatus *int32  `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	Edition    *int32  `json:"Edition,omitempty" xml:"Edition,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Enabled    *int32  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDebtStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEdition(v int32) *DescribeInstancesResponseBodyInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRemark(v string) *DescribeInstancesResponseBodyInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnabled(v int32) *DescribeInstancesResponseBodyInstances {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetGmtCreate(v int64) *DescribeInstancesResponseBodyInstances {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	BaseBandwidth    *int32  `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	QpsLimit         *int32  `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	BandwidthMbps    *int32  `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty"`
	DefenseCount     *int32  `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	SiteLimit        *int32  `json:"SiteLimit,omitempty" xml:"SiteLimit,omitempty"`
	PortLimit        *int32  `json:"PortLimit,omitempty" xml:"PortLimit,omitempty"`
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	FunctionVersion  *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DomainLimit      *int32  `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBaseBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetQpsLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.QpsLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBandwidthMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BandwidthMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetSiteLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.SiteLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPortLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PortLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetElasticBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetFunctionVersion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.FunctionVersion = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDomainLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DomainLimit = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = &v
	return s
}

type DescribeInstanceStatisticsResponseBody struct {
	InstanceStatistics []*DescribeInstanceStatisticsResponseBodyInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" type:"Repeated"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBody) SetInstanceStatistics(v []*DescribeInstanceStatisticsResponseBodyInstanceStatistics) *DescribeInstanceStatisticsResponseBody {
	s.InstanceStatistics = v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceStatisticsResponseBodyInstanceStatistics struct {
	DomainUsage       *int32  `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty"`
	DefenseCountUsage *int32  `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SiteUsage         *int32  `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
	PortUsage         *int32  `json:"PortUsage,omitempty" xml:"PortUsage,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDomainUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DomainUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDefenseCountUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DefenseCountUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetInstanceId(v string) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetSiteUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.SiteUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetPortUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.PortUsage = &v
	return s
}

type DescribeInstanceStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeIpTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Eip             *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	QueryProtocol   *string `json:"QueryProtocol,omitempty" xml:"QueryProtocol,omitempty"`
}

func (s DescribeIpTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficRequest) SetSourceIp(v string) *DescribeIpTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetResourceGroupId(v string) *DescribeIpTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetStartTime(v int64) *DescribeIpTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetInterval(v int32) *DescribeIpTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEndTime(v int64) *DescribeIpTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEip(v string) *DescribeIpTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetPort(v int32) *DescribeIpTrafficRequest {
	s.Port = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetQueryProtocol(v string) *DescribeIpTrafficRequest {
	s.QueryProtocol = &v
	return s
}

type DescribeIpTrafficResponseBody struct {
	MaxOutBps       *int64                                          `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvgInBps        *int64                                          `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	MaxInBps        *int64                                          `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	AvgOutBps       *int64                                          `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	IpTrafficPoints []*DescribeIpTrafficResponseBodyIpTrafficPoints `json:"IpTrafficPoints,omitempty" xml:"IpTrafficPoints,omitempty" type:"Repeated"`
}

func (s DescribeIpTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseBody) SetMaxOutBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetRequestId(v string) *DescribeIpTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetAvgInBps(v int64) *DescribeIpTrafficResponseBody {
	s.AvgInBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetMaxInBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetAvgOutBps(v int64) *DescribeIpTrafficResponseBody {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetIpTrafficPoints(v []*DescribeIpTrafficResponseBodyIpTrafficPoints) *DescribeIpTrafficResponseBody {
	s.IpTrafficPoints = v
	return s
}

type DescribeIpTrafficResponseBodyIpTrafficPoints struct {
	ActConns   *int32 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	Time       *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	InactConns *int32 `json:"InactConns,omitempty" xml:"InactConns,omitempty"`
	MaxInbps   *int64 `json:"MaxInbps,omitempty" xml:"MaxInbps,omitempty"`
	MaxOutbps  *int64 `json:"MaxOutbps,omitempty" xml:"MaxOutbps,omitempty"`
	Cps        *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
}

func (s DescribeIpTrafficResponseBodyIpTrafficPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponseBodyIpTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetActConns(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.ActConns = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetTime(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Time = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetInactConns(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.InactConns = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetMaxInbps(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.MaxInbps = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetMaxOutbps(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.MaxOutbps = &v
	return s
}

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetCps(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Cps = &v
	return s
}

type DescribeIpTrafficResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponse) SetHeaders(v map[string]*string) *DescribeIpTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpTrafficResponse) SetBody(v *DescribeIpTrafficResponseBody) *DescribeIpTrafficResponse {
	s.Body = v
	return s
}

type DescribeLayer4RuleAttributesRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DescribeLayer4RuleAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesRequest) SetSourceIp(v string) *DescribeLayer4RuleAttributesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RuleAttributesRequest) SetListeners(v string) *DescribeLayer4RuleAttributesRequest {
	s.Listeners = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBody struct {
	Listeners []*DescribeLayer4RuleAttributesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBody) SetListeners(v []*DescribeLayer4RuleAttributesResponseBodyListeners) *DescribeLayer4RuleAttributesResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBody) SetRequestId(v string) *DescribeLayer4RuleAttributesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListeners struct {
	FrontendPort *int32                                                   `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Protocol     *string                                                  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	InstanceId   *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Config       *DescribeLayer4RuleAttributesResponseBodyListenersConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetConfig(v *DescribeLayer4RuleAttributesResponseBodyListenersConfig) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Config = v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfig struct {
	Cc                 *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc         `json:"Cc,omitempty" xml:"Cc,omitempty" type:"Struct"`
	PayloadLen         *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" type:"Struct"`
	PersistenceTimeout *int32                                                             `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Sla                *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla        `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	Slimit             *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit     `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
	NodataConn         *string                                                            `json:"NodataConn,omitempty" xml:"NodataConn,omitempty"`
	Synproxy           *string                                                            `json:"Synproxy,omitempty" xml:"Synproxy,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfig) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetCc(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Cc = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetPayloadLen(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.PayloadLen = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetPersistenceTimeout(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSla(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Sla = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSlimit(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Slimit = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetNodataConn(v string) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.NodataConn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSynproxy(v string) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Synproxy = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigCc struct {
	Sblack []*DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack `json:"Sblack,omitempty" xml:"Sblack,omitempty" type:"Repeated"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) SetSblack(v []*DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc {
	s.Sblack = v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack struct {
	Type    *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	During  *int32 `json:"During,omitempty" xml:"During,omitempty"`
	Cnt     *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetType(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Type = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetExpires(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetDuring(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetCnt(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Cnt = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen struct {
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) SetMax(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen {
	s.Max = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) SetMin(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen {
	s.Min = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigSla struct {
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Maxconn       *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Maxconn = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit struct {
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	Pps           *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	Bps           *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Maxconn       *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	CpsMode       *int32 `json:"CpsMode,omitempty" xml:"CpsMode,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetPps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Pps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetBps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsMode(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsMode = &v
	return s
}

type DescribeLayer4RuleAttributesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLayer4RuleAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLayer4RuleAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponse) SetHeaders(v map[string]*string) *DescribeLayer4RuleAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) SetBody(v *DescribeLayer4RuleAttributesResponseBody) *DescribeLayer4RuleAttributesResponse {
	s.Body = v
	return s
}

type DescribeLayer4RulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	FrontendPort    *int32  `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesRequest) SetSourceIp(v string) *DescribeLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetInstanceId(v string) *DescribeLayer4RulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetForwardProtocol(v string) *DescribeLayer4RulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetFrontendPort(v int32) *DescribeLayer4RulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetOffset(v int32) *DescribeLayer4RulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetPageSize(v string) *DescribeLayer4RulesRequest {
	s.PageSize = &v
	return s
}

type DescribeLayer4RulesResponseBody struct {
	Listeners []*DescribeLayer4RulesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLayer4RulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBody) SetListeners(v []*DescribeLayer4RulesResponseBodyListeners) *DescribeLayer4RulesResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetRequestId(v string) *DescribeLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetTotal(v int64) *DescribeLayer4RulesResponseBody {
	s.Total = &v
	return s
}

type DescribeLayer4RulesResponseBodyListeners struct {
	FrontendPort *int32    `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	IsAutoCreate *bool     `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	Protocol     *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServers  []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BackendPort  *int32    `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
}

func (s DescribeLayer4RulesResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetIsAutoCreate(v bool) *DescribeLayer4RulesResponseBodyListeners {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetRealServers(v []*string) *DescribeLayer4RulesResponseBodyListeners {
	s.RealServers = v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetBackendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.BackendPort = &v
	return s
}

type DescribeLayer4RulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLayer4RulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponse) SetHeaders(v map[string]*string) *DescribeLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RulesResponse) SetBody(v *DescribeLayer4RulesResponseBody) *DescribeLayer4RulesResponse {
	s.Body = v
	return s
}

type DescribeLayer7CCRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Offset          *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLayer7CCRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesRequest) SetSourceIp(v string) *DescribeLayer7CCRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetResourceGroupId(v string) *DescribeLayer7CCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetDomain(v string) *DescribeLayer7CCRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetOffset(v int32) *DescribeLayer7CCRulesRequest {
	s.Offset = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetPageSize(v string) *DescribeLayer7CCRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeLayer7CCRulesResponseBody struct {
	Layer7CCRules []*DescribeLayer7CCRulesResponseBodyLayer7CCRules `json:"Layer7CCRules,omitempty" xml:"Layer7CCRules,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total         *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLayer7CCRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseBody) SetLayer7CCRules(v []*DescribeLayer7CCRulesResponseBodyLayer7CCRules) *DescribeLayer7CCRulesResponseBody {
	s.Layer7CCRules = v
	return s
}

func (s *DescribeLayer7CCRulesResponseBody) SetRequestId(v string) *DescribeLayer7CCRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBody) SetTotal(v int64) *DescribeLayer7CCRulesResponseBody {
	s.Total = &v
	return s
}

type DescribeLayer7CCRulesResponseBodyLayer7CCRules struct {
	Ttl      *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Act      *string `json:"Act,omitempty" xml:"Act,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetTtl(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetAct(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Act = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetInterval(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Interval = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetMode(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Mode = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetName(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Name = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetUri(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Uri = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetCount(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Count = &v
	return s
}

type DescribeLayer7CCRulesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLayer7CCRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLayer7CCRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponse) SetHeaders(v map[string]*string) *DescribeLayer7CCRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetBody(v *DescribeLayer7CCRulesResponseBody) *DescribeLayer7CCRulesResponse {
	s.Body = v
	return s
}

type DescribeLogStoreExistStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeLogStoreExistStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) SetSourceIp(v string) *DescribeLogStoreExistStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetLang(v string) *DescribeLogStoreExistStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeLogStoreExistStatusResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExistStatus *bool   `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
}

func (s DescribeLogStoreExistStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponseBody) SetRequestId(v string) *DescribeLogStoreExistStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponseBody) SetExistStatus(v bool) *DescribeLogStoreExistStatusResponseBody {
	s.ExistStatus = &v
	return s
}

type DescribeLogStoreExistStatusResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogStoreExistStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogStoreExistStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponse) SetHeaders(v map[string]*string) *DescribeLogStoreExistStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetBody(v *DescribeLogStoreExistStatusResponseBody) *DescribeLogStoreExistStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	EntityType      *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject    *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int32) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageNo(v int32) *DescribeOpEntitiesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotal(v int64) *DescribeOpEntitiesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	EntityType   *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAction     *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribeSimpleDomainsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeSimpleDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsRequest) SetSourceIp(v string) *DescribeSimpleDomainsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetLang(v string) *DescribeSimpleDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetResourceGroupId(v string) *DescribeSimpleDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSimpleDomainsRequest) SetInstanceIds(v []*string) *DescribeSimpleDomainsRequest {
	s.InstanceIds = v
	return s
}

type DescribeSimpleDomainsResponseBody struct {
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
}

func (s DescribeSimpleDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponseBody) SetRequestId(v string) *DescribeSimpleDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimpleDomainsResponseBody) SetDomainList(v []*string) *DescribeSimpleDomainsResponseBody {
	s.DomainList = v
	return s
}

type DescribeSimpleDomainsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSimpleDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSimpleDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponse) SetHeaders(v map[string]*string) *DescribeSimpleDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimpleDomainsResponse) SetBody(v *DescribeSimpleDomainsResponseBody) *DescribeSimpleDomainsResponse {
	s.Body = v
	return s
}

type DescribeSlsAuthStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) SetSourceIp(v string) *DescribeSlsAuthStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetLang(v string) *DescribeSlsAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsAuthStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsAuthStatus *bool   `json:"SlsAuthStatus,omitempty" xml:"SlsAuthStatus,omitempty"`
}

func (s DescribeSlsAuthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponseBody) SetRequestId(v string) *DescribeSlsAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsAuthStatusResponseBody) SetSlsAuthStatus(v bool) *DescribeSlsAuthStatusResponseBody {
	s.SlsAuthStatus = &v
	return s
}

type DescribeSlsAuthStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsAuthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse {
	s.Body = v
	return s
}

type DescribeSlsEmptyCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsEmptyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountRequest) SetSourceIp(v string) *DescribeSlsEmptyCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetLang(v string) *DescribeSlsEmptyCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetResourceGroupId(v string) *DescribeSlsEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsEmptyCountResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableCount *int32  `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
}

func (s DescribeSlsEmptyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponseBody) SetRequestId(v string) *DescribeSlsEmptyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsEmptyCountResponseBody) SetAvailableCount(v int32) *DescribeSlsEmptyCountResponseBody {
	s.AvailableCount = &v
	return s
}

type DescribeSlsEmptyCountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsEmptyCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsEmptyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponse) SetHeaders(v map[string]*string) *DescribeSlsEmptyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsEmptyCountResponse) SetBody(v *DescribeSlsEmptyCountResponseBody) *DescribeSlsEmptyCountResponse {
	s.Body = v
	return s
}

type DescribeSlsLogstoreInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetLang(v string) *DescribeSlsLogstoreInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsLogstoreInfoResponseBody struct {
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quota     *int64  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	LogStore  *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Used      *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	Ttl       *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeSlsLogstoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetProject(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetRequestId(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetQuota(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetLogStore(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetUsed(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetTtl(v int32) *DescribeSlsLogstoreInfoResponseBody {
	s.Ttl = &v
	return s
}

type DescribeSlsLogstoreInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsLogstoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsLogstoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponse) SetHeaders(v map[string]*string) *DescribeSlsLogstoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetBody(v *DescribeSlsLogstoreInfoResponseBody) *DescribeSlsLogstoreInfoResponse {
	s.Body = v
	return s
}

type DescribeSlsOpenStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) SetSourceIp(v string) *DescribeSlsOpenStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetLang(v string) *DescribeSlsOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeSlsOpenStatusResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsOpenStatus *bool   `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s DescribeSlsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponseBody) SetRequestId(v string) *DescribeSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *DescribeSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

type DescribeSlsOpenStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetBody(v *DescribeSlsOpenStatusResponseBody) *DescribeSlsOpenStatusResponse {
	s.Body = v
	return s
}

type DescribleCertListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribleCertListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListRequest) GoString() string {
	return s.String()
}

func (s *DescribleCertListRequest) SetSourceIp(v string) *DescribleCertListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleCertListRequest) SetResourceGroupId(v string) *DescribleCertListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleCertListRequest) SetDomain(v string) *DescribleCertListRequest {
	s.Domain = &v
	return s
}

type DescribleCertListResponseBody struct {
	CertList  []*DescribleCertListResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribleCertListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseBody) SetCertList(v []*DescribleCertListResponseBodyCertList) *DescribleCertListResponseBody {
	s.CertList = v
	return s
}

func (s *DescribleCertListResponseBody) SetRequestId(v string) *DescribleCertListResponseBody {
	s.RequestId = &v
	return s
}

type DescribleCertListResponseBodyCertList struct {
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DomainRelated *bool   `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Issuer        *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Common        *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Id            *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribleCertListResponseBodyCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponseBodyCertList) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseBodyCertList) SetEndDate(v string) *DescribleCertListResponseBodyCertList {
	s.EndDate = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetDomainRelated(v bool) *DescribleCertListResponseBodyCertList {
	s.DomainRelated = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetStartDate(v string) *DescribleCertListResponseBodyCertList {
	s.StartDate = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetIssuer(v string) *DescribleCertListResponseBodyCertList {
	s.Issuer = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetName(v string) *DescribleCertListResponseBodyCertList {
	s.Name = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetCommon(v string) *DescribleCertListResponseBodyCertList {
	s.Common = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetId(v int32) *DescribleCertListResponseBodyCertList {
	s.Id = &v
	return s
}

type DescribleCertListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribleCertListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribleCertListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponse) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponse) SetHeaders(v map[string]*string) *DescribleCertListResponse {
	s.Headers = v
	return s
}

func (s *DescribleCertListResponse) SetBody(v *DescribleCertListResponseBody) *DescribleCertListResponse {
	s.Body = v
	return s
}

type DescribleLayer7InstanceRelationsRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DomainList      []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsRequest) SetSourceIp(v string) *DescribleLayer7InstanceRelationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetResourceGroupId(v string) *DescribleLayer7InstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetDomainList(v []*string) *DescribleLayer7InstanceRelationsRequest {
	s.DomainList = v
	return s
}

type DescribleLayer7InstanceRelationsResponseBody struct {
	RequestId               *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Layer7InstanceRelations []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations `json:"Layer7InstanceRelations,omitempty" xml:"Layer7InstanceRelations,omitempty" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetRequestId(v string) *DescribleLayer7InstanceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetLayer7InstanceRelations(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) *DescribleLayer7InstanceRelationsResponseBody {
	s.Layer7InstanceRelations = v
	return s
}

type DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations struct {
	Domain          *string                                                                               `json:"Domain,omitempty" xml:"Domain,omitempty"`
	InstanceDetails []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) SetDomain(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations {
	s.Domain = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) SetInstanceDetails(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations {
	s.InstanceDetails = v
	return s
}

type DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails struct {
	EipList         []*string `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	FunctionVersion *string   `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetEipList(v []*string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.EipList = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetFunctionVersion(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.FunctionVersion = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetInstanceId(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.InstanceId = &v
	return s
}

type DescribleLayer7InstanceRelationsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribleLayer7InstanceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribleLayer7InstanceRelationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponse) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponse) SetHeaders(v map[string]*string) *DescribleLayer7InstanceRelationsResponse {
	s.Headers = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) SetBody(v *DescribleLayer7InstanceRelationsResponseBody) *DescribleLayer7InstanceRelationsResponse {
	s.Body = v
	return s
}

type DisableLayer7CCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRequest) SetSourceIp(v string) *DisableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRequest) SetResourceGroupId(v string) *DisableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRequest) SetDomain(v string) *DisableLayer7CCRequest {
	s.Domain = &v
	return s
}

type DisableLayer7CCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLayer7CCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCResponseBody) SetRequestId(v string) *DisableLayer7CCResponseBody {
	s.RequestId = &v
	return s
}

type DisableLayer7CCResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLayer7CCResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCResponse) SetHeaders(v map[string]*string) *DisableLayer7CCResponse {
	s.Headers = v
	return s
}

func (s *DisableLayer7CCResponse) SetBody(v *DisableLayer7CCResponseBody) *DisableLayer7CCResponse {
	s.Body = v
	return s
}

type DisableLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DisableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleRequest) SetSourceIp(v string) *DisableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetResourceGroupId(v string) *DisableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetDomain(v string) *DisableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

type DisableLayer7CCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableLayer7CCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleResponseBody) SetRequestId(v string) *DisableLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableLayer7CCRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleResponse) SetHeaders(v map[string]*string) *DisableLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableLayer7CCRuleResponse) SetBody(v *DisableLayer7CCRuleResponseBody) *DisableLayer7CCRuleResponse {
	s.Body = v
	return s
}

type EmptySlsLogstoreRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreRequest) SetSourceIp(v string) *EmptySlsLogstoreRequest {
	s.SourceIp = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetLang(v string) *EmptySlsLogstoreRequest {
	s.Lang = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
	s.ResourceGroupId = &v
	return s
}

type EmptySlsLogstoreResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmptySlsLogstoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreResponseBody) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreResponseBody) SetRequestId(v string) *EmptySlsLogstoreResponseBody {
	s.RequestId = &v
	return s
}

type EmptySlsLogstoreResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EmptySlsLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EmptySlsLogstoreResponse) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreResponse) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreResponse) SetHeaders(v map[string]*string) *EmptySlsLogstoreResponse {
	s.Headers = v
	return s
}

func (s *EmptySlsLogstoreResponse) SetBody(v *EmptySlsLogstoreResponseBody) *EmptySlsLogstoreResponse {
	s.Body = v
	return s
}

type EnableLayer7CCRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRequest) SetSourceIp(v string) *EnableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableLayer7CCRequest) SetResourceGroupId(v string) *EnableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRequest) SetDomain(v string) *EnableLayer7CCRequest {
	s.Domain = &v
	return s
}

type EnableLayer7CCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLayer7CCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCResponseBody) SetRequestId(v string) *EnableLayer7CCResponseBody {
	s.RequestId = &v
	return s
}

type EnableLayer7CCResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLayer7CCResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCResponse) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCResponse) SetHeaders(v map[string]*string) *EnableLayer7CCResponse {
	s.Headers = v
	return s
}

func (s *EnableLayer7CCResponse) SetBody(v *EnableLayer7CCResponseBody) *EnableLayer7CCResponse {
	s.Body = v
	return s
}

type EnableLayer7CCRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s EnableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleRequest) SetSourceIp(v string) *EnableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetResourceGroupId(v string) *EnableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetDomain(v string) *EnableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

type EnableLayer7CCRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableLayer7CCRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleResponseBody) SetRequestId(v string) *EnableLayer7CCRuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableLayer7CCRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableLayer7CCRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleResponse) SetHeaders(v map[string]*string) *EnableLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableLayer7CCRuleResponse) SetBody(v *EnableLayer7CCRuleResponseBody) *EnableLayer7CCRuleResponse {
	s.Body = v
	return s
}

type ListAsyncTaskRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskRequest) SetSourceIp(v string) *ListAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *ListAsyncTaskRequest) SetLang(v string) *ListAsyncTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListAsyncTaskRequest) SetResourceGroupId(v string) *ListAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageNo(v int32) *ListAsyncTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListAsyncTaskRequest) SetPageSize(v int32) *ListAsyncTaskRequest {
	s.PageSize = &v
	return s
}

type ListAsyncTaskResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                 `json:"Total,omitempty" xml:"Total,omitempty"`
	AsyncTasks []*ListAsyncTaskResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
}

func (s ListAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseBody) SetRequestId(v string) *ListAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTaskResponseBody) SetTotal(v int32) *ListAsyncTaskResponseBody {
	s.Total = &v
	return s
}

func (s *ListAsyncTaskResponseBody) SetAsyncTasks(v []*ListAsyncTaskResponseBodyAsyncTasks) *ListAsyncTaskResponseBody {
	s.AsyncTasks = v
	return s
}

type ListAsyncTaskResponseBodyAsyncTasks struct {
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskType   *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskId     *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAsyncTaskResponseBodyAsyncTasks) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponseBodyAsyncTasks) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetEndTime(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.EndTime = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskType(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetStartTime(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskParams(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskStatus(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskResult(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskId(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskId = &v
	return s
}

type ListAsyncTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponse) SetHeaders(v map[string]*string) *ListAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTaskResponse) SetBody(v *ListAsyncTaskResponseBody) *ListAsyncTaskResponse {
	s.Body = v
	return s
}

type ListLayer7CustomPortsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListLayer7CustomPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsRequest) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsRequest) SetSourceIp(v string) *ListLayer7CustomPortsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetLang(v string) *ListLayer7CustomPortsRequest {
	s.Lang = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetResourceGroupId(v string) *ListLayer7CustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListLayer7CustomPortsResponseBody struct {
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Layer7CustomPorts []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts `json:"Layer7CustomPorts,omitempty" xml:"Layer7CustomPorts,omitempty" type:"Repeated"`
}

func (s ListLayer7CustomPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBody) SetRequestId(v string) *ListLayer7CustomPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLayer7CustomPortsResponseBody) SetLayer7CustomPorts(v []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts) *ListLayer7CustomPortsResponseBody {
	s.Layer7CustomPorts = v
	return s
}

type ListLayer7CustomPortsResponseBodyLayer7CustomPorts struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	ProxyType  *string   `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetProxyPorts(v []*string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.ProxyPorts = v
	return s
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetProxyType(v string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.ProxyType = &v
	return s
}

type ListLayer7CustomPortsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLayer7CustomPortsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLayer7CustomPortsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponse) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponse) SetHeaders(v map[string]*string) *ListLayer7CustomPortsResponse {
	s.Headers = v
	return s
}

func (s *ListLayer7CustomPortsResponse) SetBody(v *ListLayer7CustomPortsResponseBody) *ListLayer7CustomPortsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetSourceIp(v string) *ListTagKeysRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

type ListTagKeysResponseBody struct {
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
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

type ListTagResourcesRequest struct {
	SourceIp        *string                       `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken       *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId      []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetSourceIp(v string) *ListTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
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

type ListValueAddedRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ListValueAddedRequest) SetSourceIp(v string) *ListValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ListValueAddedRequest) SetResourceGroupId(v string) *ListValueAddedRequest {
	s.ResourceGroupId = &v
	return s
}

type ListValueAddedResponseBody struct {
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ValueAddedList []*ListValueAddedResponseBodyValueAddedList `json:"ValueAddedList,omitempty" xml:"ValueAddedList,omitempty" type:"Repeated"`
}

func (s ListValueAddedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseBody) SetRequestId(v string) *ListValueAddedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListValueAddedResponseBody) SetValueAddedList(v []*ListValueAddedResponseBodyValueAddedList) *ListValueAddedResponseBody {
	s.ValueAddedList = v
	return s
}

type ListValueAddedResponseBodyValueAddedList struct {
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LogSize    *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListValueAddedResponseBodyValueAddedList) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponseBodyValueAddedList) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseBodyValueAddedList) SetStatus(v int32) *ListValueAddedResponseBodyValueAddedList {
	s.Status = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetExpireTime(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.ExpireTime = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetLogSize(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.LogSize = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetGmtCreate(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.GmtCreate = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetInstanceId(v string) *ListValueAddedResponseBodyValueAddedList {
	s.InstanceId = &v
	return s
}

type ListValueAddedResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListValueAddedResponse) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponse) SetHeaders(v map[string]*string) *ListValueAddedResponse {
	s.Headers = v
	return s
}

func (s *ListValueAddedResponse) SetBody(v *ListValueAddedResponseBody) *ListValueAddedResponse {
	s.Body = v
	return s
}

type ModifyElasticBandWidthRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) SetSourceIp(v string) *ModifyElasticBandWidthRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

type ModifyElasticBandWidthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBandWidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponseBody) SetRequestId(v string) *ModifyElasticBandWidthResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticBandWidthResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyElasticBandWidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElasticBandWidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponse) SetHeaders(v map[string]*string) *ModifyElasticBandWidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBandWidthResponse) SetBody(v *ModifyElasticBandWidthResponseBody) *ModifyElasticBandWidthResponse {
	s.Body = v
	return s
}

type ModifyFullLogTtlRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) SetSourceIp(v string) *ModifyFullLogTtlRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetLang(v string) *ModifyFullLogTtlRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int32) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyFullLogTtlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFullLogTtlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponseBody) SetRequestId(v string) *ModifyFullLogTtlResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFullLogTtlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFullLogTtlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFullLogTtlResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlResponse) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlResponse) SetHeaders(v map[string]*string) *ModifyFullLogTtlResponse {
	s.Headers = v
	return s
}

func (s *ModifyFullLogTtlResponse) SetBody(v *ModifyFullLogTtlResponseBody) *ModifyFullLogTtlResponse {
	s.Body = v
	return s
}

type ModifyInstanceRemarkRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyInstanceRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkRequest) SetSourceIp(v string) *ModifyInstanceRemarkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

type ModifyInstanceRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponseBody) SetRequestId(v string) *ModifyInstanceRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceRemarkResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponse) SetHeaders(v map[string]*string) *ModifyInstanceRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceRemarkResponse) SetBody(v *ModifyInstanceRemarkResponseBody) *ModifyInstanceRemarkResponse {
	s.Body = v
	return s
}

type OpenDomainSlsConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s OpenDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigRequest) SetSourceIp(v string) *OpenDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetLang(v string) *OpenDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetResourceGroupId(v string) *OpenDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetDomain(v string) *OpenDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

type OpenDomainSlsConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenDomainSlsConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigResponseBody) SetRequestId(v string) *OpenDomainSlsConfigResponseBody {
	s.RequestId = &v
	return s
}

type OpenDomainSlsConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenDomainSlsConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigResponse) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigResponse) SetHeaders(v map[string]*string) *OpenDomainSlsConfigResponse {
	s.Headers = v
	return s
}

func (s *OpenDomainSlsConfigResponse) SetBody(v *OpenDomainSlsConfigResponseBody) *OpenDomainSlsConfigResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetSourceIp(v string) *ReleaseInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type ReleaseValueAddedRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedRequest) SetSourceIp(v string) *ReleaseValueAddedRequest {
	s.SourceIp = &v
	return s
}

func (s *ReleaseValueAddedRequest) SetInstanceId(v string) *ReleaseValueAddedRequest {
	s.InstanceId = &v
	return s
}

type ReleaseValueAddedResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseValueAddedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedResponseBody) SetRequestId(v string) *ReleaseValueAddedResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseValueAddedResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseValueAddedResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedResponse) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedResponse) SetHeaders(v map[string]*string) *ReleaseValueAddedResponse {
	s.Headers = v
	return s
}

func (s *ReleaseValueAddedResponse) SetBody(v *ReleaseValueAddedResponseBody) *ReleaseValueAddedResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	SourceIp        *string                   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId      []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetSourceIp(v string) *TagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

type UntagResourcesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetSourceIp(v string) *UntagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddoscoo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddLayer7CCRuleWithOptions(request *AddLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *AddLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddLayer7CCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddLayer7CCRule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLayer7CCRule(request *AddLayer7CCRuleRequest) (_result *AddLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLayer7CCRuleResponse{}
	_body, _err := client.AddLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDomainSlsConfigWithOptions(request *CloseDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *CloseDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseDomainSlsConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseDomainSlsConfig"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDomainSlsConfig(request *CloseDomainSlsConfigRequest) (_result *CloseDomainSlsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDomainSlsConfigResponse{}
	_body, _err := client.CloseDomainSlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigDomainAccessModeWithOptions(request *ConfigDomainAccessModeRequest, runtime *util.RuntimeOptions) (_result *ConfigDomainAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigDomainAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigDomainAccessMode"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigDomainAccessMode(request *ConfigDomainAccessModeRequest) (_result *ConfigDomainAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigDomainAccessModeResponse{}
	_body, _err := client.ConfigDomainAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigHealthCheckWithOptions(request *ConfigHealthCheckRequest, runtime *util.RuntimeOptions) (_result *ConfigHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigHealthCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigHealthCheck"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigHealthCheck(request *ConfigHealthCheckRequest) (_result *ConfigHealthCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigHealthCheckResponse{}
	_body, _err := client.ConfigHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer4RuleWithOptions(request *ConfigLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer4Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer4Rule(request *ConfigLayer4RuleRequest) (_result *ConfigLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer4RuleResponse{}
	_body, _err := client.ConfigLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer4RuleAttributeWithOptions(request *ConfigLayer4RuleAttributeRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer4RuleAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer4RuleAttribute"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer4RuleAttribute(request *ConfigLayer4RuleAttributeRequest) (_result *ConfigLayer4RuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer4RuleAttributeResponse{}
	_body, _err := client.ConfigLayer4RuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7BlackWhiteListWithOptions(request *ConfigLayer7BlackWhiteListRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7BlackWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer7BlackWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer7BlackWhiteList"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7BlackWhiteList(request *ConfigLayer7BlackWhiteListRequest) (_result *ConfigLayer7BlackWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7BlackWhiteListResponse{}
	_body, _err := client.ConfigLayer7BlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CCRuleWithOptions(request *ConfigLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer7CCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer7CCRule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7CCRule(request *ConfigLayer7CCRuleRequest) (_result *ConfigLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CCRuleResponse{}
	_body, _err := client.ConfigLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CCTemplateWithOptions(request *ConfigLayer7CCTemplateRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer7CCTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer7CCTemplate"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7CCTemplate(request *ConfigLayer7CCTemplateRequest) (_result *ConfigLayer7CCTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CCTemplateResponse{}
	_body, _err := client.ConfigLayer7CCTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7CertWithOptions(request *ConfigLayer7CertRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer7CertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer7Cert"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7Cert(request *ConfigLayer7CertRequest) (_result *ConfigLayer7CertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7CertResponse{}
	_body, _err := client.ConfigLayer7CertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigLayer7RuleWithOptions(request *ConfigLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigLayer7RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigLayer7Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigLayer7Rule(request *ConfigLayer7RuleRequest) (_result *ConfigLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigLayer7RuleResponse{}
	_body, _err := client.ConfigLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAsyncTaskWithOptions(request *CreateAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAsyncTask"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAsyncTask(request *CreateAsyncTaskRequest) (_result *CreateAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CreateAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer4RuleWithOptions(request *CreateLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLayer4Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer4Rule(request *CreateLayer4RuleRequest) (_result *CreateLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer4RuleResponse{}
	_body, _err := client.CreateLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLayer7RuleWithOptions(request *CreateLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLayer7RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLayer7Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLayer7Rule(request *CreateLayer7RuleRequest) (_result *CreateLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLayer7RuleResponse{}
	_body, _err := client.CreateLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAsyncTaskWithOptions(request *DeleteAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAsyncTask"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAsyncTask(request *DeleteAsyncTaskRequest) (_result *DeleteAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DeleteAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer4RuleWithOptions(request *DeleteLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLayer4Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer4Rule(request *DeleteLayer4RuleRequest) (_result *DeleteLayer4RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer4RuleResponse{}
	_body, _err := client.DeleteLayer4RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer7CCRuleWithOptions(request *DeleteLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLayer7CCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLayer7CCRule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer7CCRule(request *DeleteLayer7CCRuleRequest) (_result *DeleteLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer7CCRuleResponse{}
	_body, _err := client.DeleteLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLayer7RuleWithOptions(request *DeleteLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLayer7RuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLayer7Rule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLayer7Rule(request *DeleteLayer7RuleRequest) (_result *DeleteLayer7RuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLayer7RuleResponse{}
	_body, _err := client.DeleteLayer7RuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackSourceCidrWithOptions(request *DescribeBackSourceCidrRequest, runtime *util.RuntimeOptions) (_result *DescribeBackSourceCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackSourceCidr"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackSourceCidr(request *DescribeBackSourceCidrRequest) (_result *DescribeBackSourceCidrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DescribeBackSourceCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBatchSlsDispatchStatusWithOptions(request *DescribeBatchSlsDispatchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBatchSlsDispatchStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBatchSlsDispatchStatus"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBatchSlsDispatchStatus(request *DescribeBatchSlsDispatchStatusRequest) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBatchSlsDispatchStatusResponse{}
	_body, _err := client.DescribeBatchSlsDispatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDoSEventsWithOptions(request *DescribeDDoSEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDoSEvents"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDoSEvents(request *DescribeDDoSEventsRequest) (_result *DescribeDDoSEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DescribeDDoSEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDDoSTrafficWithOptions(request *DescribeDDoSTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDDoSTrafficResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDDoSTraffic"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDDoSTraffic(request *DescribeDDoSTrafficRequest) (_result *DescribeDDoSTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDDoSTrafficResponse{}
	_body, _err := client.DescribeDDoSTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatisticsWithOptions(request *DescribeDefenseCountStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDefenseCountStatistics"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefenseCountStatistics(request *DescribeDefenseCountStatisticsRequest) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DescribeDefenseCountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAccessModeWithOptions(request *DescribeDomainAccessModeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainAccessModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainAccessMode"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAccessMode(request *DescribeDomainAccessModeRequest) (_result *DescribeDomainAccessModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAccessModeResponse{}
	_body, _err := client.DescribeDomainAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainAttackEventsWithOptions(request *DescribeDomainAttackEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainAttackEvents"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainAttackEvents(request *DescribeDomainAttackEventsRequest) (_result *DescribeDomainAttackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DescribeDomainAttackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithOptions(request *DescribeDomainQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQps"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQps(request *DescribeDomainQpsRequest) (_result *DescribeDomainQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsResponse{}
	_body, _err := client.DescribeDomainQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithCacheWithOptions(request *DescribeDomainQpsWithCacheRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainQpsWithCache"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainQpsWithCache(request *DescribeDomainQpsWithCacheRequest) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsWithCacheResponse{}
	_body, _err := client.DescribeDomainQpsWithCacheWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomains"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeDomainSlsStatusWithOptions(request *DescribeDomainSlsStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSlsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSlsStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSlsStatus"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSlsStatus(request *DescribeDomainSlsStatusRequest) (_result *DescribeDomainSlsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSlsStatusResponse{}
	_body, _err := client.DescribeDomainSlsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpecWithOptions(request *DescribeElasticBandwidthSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeElasticBandwidthSpec"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeElasticBandwidthSpec(request *DescribeElasticBandwidthSpecRequest) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DescribeElasticBandwidthSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckListWithOptions(request *DescribeHealthCheckListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHealthCheckList"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckList(request *DescribeHealthCheckListRequest) (_result *DescribeHealthCheckListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DescribeHealthCheckListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatusListWithOptions(request *DescribeHealthCheckStatusListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckStatusListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHealthCheckStatusListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHealthCheckStatusList"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckStatusList(request *DescribeHealthCheckStatusListRequest) (_result *DescribeHealthCheckStatusListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckStatusListResponse{}
	_body, _err := client.DescribeHealthCheckStatusListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceDetailsWithOptions(request *DescribeInstanceDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceDetails"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (_result *DescribeInstanceDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DescribeInstanceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSpecs"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStatistics"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (_result *DescribeInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DescribeInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpTrafficWithOptions(request *DescribeIpTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeIpTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpTrafficResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpTraffic"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpTraffic(request *DescribeIpTrafficRequest) (_result *DescribeIpTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpTrafficResponse{}
	_body, _err := client.DescribeIpTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer4RuleAttributesWithOptions(request *DescribeLayer4RuleAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RuleAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLayer4RuleAttributesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLayer4RuleAttributes"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer4RuleAttributes(request *DescribeLayer4RuleAttributesRequest) (_result *DescribeLayer4RuleAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer4RuleAttributesResponse{}
	_body, _err := client.DescribeLayer4RuleAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer4RulesWithOptions(request *DescribeLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLayer4Rules"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer4Rules(request *DescribeLayer4RulesRequest) (_result *DescribeLayer4RulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer4RulesResponse{}
	_body, _err := client.DescribeLayer4RulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLayer7CCRulesWithOptions(request *DescribeLayer7CCRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer7CCRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLayer7CCRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLayer7CCRules"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLayer7CCRules(request *DescribeLayer7CCRulesRequest) (_result *DescribeLayer7CCRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLayer7CCRulesResponse{}
	_body, _err := client.DescribeLayer7CCRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatusWithOptions(request *DescribeLogStoreExistStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogStoreExistStatus"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogStoreExistStatus(request *DescribeLogStoreExistStatusRequest) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DescribeLogStoreExistStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOpEntities"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSimpleDomainsWithOptions(request *DescribeSimpleDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeSimpleDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSimpleDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSimpleDomains"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSimpleDomains(request *DescribeSimpleDomainsRequest) (_result *DescribeSimpleDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSimpleDomainsResponse{}
	_body, _err := client.DescribeSimpleDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsAuthStatus"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (_result *DescribeSlsAuthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DescribeSlsAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsEmptyCountWithOptions(request *DescribeSlsEmptyCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsEmptyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsEmptyCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsEmptyCount"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsEmptyCount(request *DescribeSlsEmptyCountRequest) (_result *DescribeSlsEmptyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsEmptyCountResponse{}
	_body, _err := client.DescribeSlsEmptyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsLogstoreInfoWithOptions(request *DescribeSlsLogstoreInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsLogstoreInfo"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsLogstoreInfo(request *DescribeSlsLogstoreInfoRequest) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DescribeSlsLogstoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatusWithOptions(request *DescribeSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlsOpenStatus"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlsOpenStatus(request *DescribeSlsOpenStatusRequest) (_result *DescribeSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DescribeSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribleCertListWithOptions(request *DescribleCertListRequest, runtime *util.RuntimeOptions) (_result *DescribleCertListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribleCertListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribleCertList"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribleCertList(request *DescribleCertListRequest) (_result *DescribleCertListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribleCertListResponse{}
	_body, _err := client.DescribleCertListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribleLayer7InstanceRelationsWithOptions(request *DescribleLayer7InstanceRelationsRequest, runtime *util.RuntimeOptions) (_result *DescribleLayer7InstanceRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribleLayer7InstanceRelationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribleLayer7InstanceRelations"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribleLayer7InstanceRelations(request *DescribleLayer7InstanceRelationsRequest) (_result *DescribleLayer7InstanceRelationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribleLayer7InstanceRelationsResponse{}
	_body, _err := client.DescribleLayer7InstanceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLayer7CCWithOptions(request *DisableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableLayer7CCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableLayer7CC"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLayer7CC(request *DisableLayer7CCRequest) (_result *DisableLayer7CCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLayer7CCResponse{}
	_body, _err := client.DisableLayer7CCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableLayer7CCRuleWithOptions(request *DisableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableLayer7CCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableLayer7CCRule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableLayer7CCRule(request *DisableLayer7CCRuleRequest) (_result *DisableLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableLayer7CCRuleResponse{}
	_body, _err := client.DisableLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EmptySlsLogstoreWithOptions(request *EmptySlsLogstoreRequest, runtime *util.RuntimeOptions) (_result *EmptySlsLogstoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EmptySlsLogstore"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EmptySlsLogstore(request *EmptySlsLogstoreRequest) (_result *EmptySlsLogstoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.EmptySlsLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLayer7CCWithOptions(request *EnableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableLayer7CCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableLayer7CC"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLayer7CC(request *EnableLayer7CCRequest) (_result *EnableLayer7CCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLayer7CCResponse{}
	_body, _err := client.EnableLayer7CCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableLayer7CCRuleWithOptions(request *EnableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableLayer7CCRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableLayer7CCRule"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableLayer7CCRule(request *EnableLayer7CCRuleRequest) (_result *EnableLayer7CCRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableLayer7CCRuleResponse{}
	_body, _err := client.EnableLayer7CCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsyncTaskWithOptions(request *ListAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *ListAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAsyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAsyncTask"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsyncTask(request *ListAsyncTaskRequest) (_result *ListAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsyncTaskResponse{}
	_body, _err := client.ListAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLayer7CustomPortsWithOptions(request *ListLayer7CustomPortsRequest, runtime *util.RuntimeOptions) (_result *ListLayer7CustomPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLayer7CustomPortsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLayer7CustomPorts"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLayer7CustomPorts(request *ListLayer7CustomPortsRequest) (_result *ListLayer7CustomPortsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLayer7CustomPortsResponse{}
	_body, _err := client.ListLayer7CustomPortsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListValueAddedWithOptions(request *ListValueAddedRequest, runtime *util.RuntimeOptions) (_result *ListValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListValueAddedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListValueAdded"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListValueAdded(request *ListValueAddedRequest) (_result *ListValueAddedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListValueAddedResponse{}
	_body, _err := client.ListValueAddedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticBandWidthWithOptions(request *ModifyElasticBandWidthRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticBandWidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyElasticBandWidth"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticBandWidth(request *ModifyElasticBandWidthRequest) (_result *ModifyElasticBandWidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.ModifyElasticBandWidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFullLogTtlWithOptions(request *ModifyFullLogTtlRequest, runtime *util.RuntimeOptions) (_result *ModifyFullLogTtlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFullLogTtl"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFullLogTtl(request *ModifyFullLogTtlRequest) (_result *ModifyFullLogTtlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.ModifyFullLogTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceRemarkWithOptions(request *ModifyInstanceRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceRemark"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (_result *ModifyInstanceRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.ModifyInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenDomainSlsConfigWithOptions(request *OpenDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *OpenDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OpenDomainSlsConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenDomainSlsConfig"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenDomainSlsConfig(request *OpenDomainSlsConfigRequest) (_result *OpenDomainSlsConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenDomainSlsConfigResponse{}
	_body, _err := client.OpenDomainSlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseInstance"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseValueAddedWithOptions(request *ReleaseValueAddedRequest, runtime *util.RuntimeOptions) (_result *ReleaseValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseValueAddedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseValueAdded"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseValueAdded(request *ReleaseValueAddedRequest) (_result *ReleaseValueAddedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseValueAddedResponse{}
	_body, _err := client.ReleaseValueAddedWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2017-12-28"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
