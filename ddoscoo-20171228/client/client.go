// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// match
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s AddLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleRequest) SetAct(v string) *AddLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetCount(v int32) *AddLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetDomain(v string) *AddLayer7CCRuleRequest {
	s.Domain = &v
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

func (s *AddLayer7CCRuleRequest) SetName(v string) *AddLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *AddLayer7CCRuleRequest) SetResourceGroupId(v string) *AddLayer7CCRuleRequest {
	s.ResourceGroupId = &v
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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddLayer7CCRuleResponse) SetStatusCode(v int32) *AddLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLayer7CCRuleResponse) SetBody(v *AddLayer7CCRuleResponseBody) *AddLayer7CCRuleResponse {
	s.Body = v
	return s
}

type CloseDomainSlsConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CloseDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigRequest) SetDomain(v string) *CloseDomainSlsConfigRequest {
	s.Domain = &v
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

func (s *CloseDomainSlsConfigRequest) SetSourceIp(v string) *CloseDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

type CloseDomainSlsConfigResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CloseDomainSlsConfigResponse) SetStatusCode(v int32) *CloseDomainSlsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDomainSlsConfigResponse) SetBody(v *CloseDomainSlsConfigResponseBody) *CloseDomainSlsConfigResponse {
	s.Body = v
	return s
}

type ConfigHealthCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"Type":"tcp","Timeout":10,"Port":80,"Interval":10,"Up":10,"Down":40}"}
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigHealthCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigHealthCheckRequest) GoString() string {
	return s.String()
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

func (s *ConfigHealthCheckRequest) SetInstanceId(v string) *ConfigHealthCheckRequest {
	s.InstanceId = &v
	return s
}

type ConfigHealthCheckResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigHealthCheckResponse) SetStatusCode(v int32) *ConfigHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigHealthCheckResponse) SetBody(v *ConfigHealthCheckResponseBody) *ConfigHealthCheckResponse {
	s.Body = v
	return s
}

type ConfigLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners   *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable *int64  `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
}

func (s ConfigLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleRequest) SetListeners(v string) *ConfigLayer4RuleRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RuleRequest) SetProxyEnable(v int64) *ConfigLayer4RuleRequest {
	s.ProxyEnable = &v
	return s
}

type ConfigLayer4RuleResponseBody struct {
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer4RuleResponse) SetStatusCode(v int32) *ConfigLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RuleResponse) SetBody(v *ConfigLayer4RuleResponseBody) *ConfigLayer4RuleResponse {
	s.Body = v
	return s
}

type ConfigLayer4RuleAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"Slimit":{"CpsEnable":1,"MaxconnEnable":1,"Cps":1,"Maxconn":1},"Sla":{"CpsEnable":1,"MaxconnEnable":1,"Cps":100,"Maxconn":1000},"PayloadLen":{"Min":0,"Max":6000}}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TCP
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigLayer4RuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer4RuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeRequest) SetConfig(v string) *ConfigLayer4RuleAttributeRequest {
	s.Config = &v
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

func (s *ConfigLayer4RuleAttributeRequest) SetInstanceId(v string) *ConfigLayer4RuleAttributeRequest {
	s.InstanceId = &v
	return s
}

type ConfigLayer4RuleAttributeResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer4RuleAttributeResponse) SetStatusCode(v int32) *ConfigLayer4RuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RuleAttributeResponse) SetBody(v *ConfigLayer4RuleAttributeResponseBody) *ConfigLayer4RuleAttributeResponse {
	s.Body = v
	return s
}

type ConfigLayer7BlackWhiteListRequest struct {
	// example:
	//
	// 1.1.1.1
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigLayer7BlackWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7BlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7BlackWhiteListRequest) SetBlackList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.BlackList = v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetDomain(v string) *ConfigLayer7BlackWhiteListRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetResourceGroupId(v string) *ConfigLayer7BlackWhiteListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListRequest) SetWhiteList(v []*string) *ConfigLayer7BlackWhiteListRequest {
	s.WhiteList = v
	return s
}

type ConfigLayer7BlackWhiteListResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7BlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer7BlackWhiteListResponse) SetStatusCode(v int32) *ConfigLayer7BlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7BlackWhiteListResponse) SetBody(v *ConfigLayer7BlackWhiteListResponseBody) *ConfigLayer7BlackWhiteListResponse {
	s.Body = v
	return s
}

type ConfigLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// match
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ConfigLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleRequest) SetAct(v string) *ConfigLayer7CCRuleRequest {
	s.Act = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetCount(v int32) *ConfigLayer7CCRuleRequest {
	s.Count = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetDomain(v string) *ConfigLayer7CCRuleRequest {
	s.Domain = &v
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

func (s *ConfigLayer7CCRuleRequest) SetName(v string) *ConfigLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *ConfigLayer7CCRuleRequest) SetResourceGroupId(v string) *ConfigLayer7CCRuleRequest {
	s.ResourceGroupId = &v
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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer7CCRuleResponse) SetStatusCode(v int32) *ConfigLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CCRuleResponse) SetBody(v *ConfigLayer7CCRuleResponseBody) *ConfigLayer7CCRuleResponse {
	s.Body = v
	return s
}

type ConfigLayer7CCTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ConfigLayer7CCTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCTemplateRequest) SetDomain(v string) *ConfigLayer7CCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetResourceGroupId(v string) *ConfigLayer7CCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CCTemplateRequest) SetTemplate(v string) *ConfigLayer7CCTemplateRequest {
	s.Template = &v
	return s
}

type ConfigLayer7CCTemplateResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CCTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer7CCTemplateResponse) SetStatusCode(v int32) *ConfigLayer7CCTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CCTemplateResponse) SetBody(v *ConfigLayer7CCTemplateResponseBody) *ConfigLayer7CCTemplateResponse {
	s.Body = v
	return s
}

type ConfigLayer7CertRequest struct {
	// example:
	//
	// xx
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// example:
	//
	// 1234
	CertId         *int32  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// testCertName
	CertName   *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// xx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ConfigLayer7CertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7CertRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertRequest) SetCert(v string) *ConfigLayer7CertRequest {
	s.Cert = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertId(v int32) *ConfigLayer7CertRequest {
	s.CertId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertIdentifier(v string) *ConfigLayer7CertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertName(v string) *ConfigLayer7CertRequest {
	s.CertName = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertRegion(v string) *ConfigLayer7CertRequest {
	s.CertRegion = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetDomain(v string) *ConfigLayer7CertRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetKey(v string) *ConfigLayer7CertRequest {
	s.Key = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetResourceGroupId(v string) *ConfigLayer7CertRequest {
	s.ResourceGroupId = &v
	return s
}

type ConfigLayer7CertResponseBody struct {
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer7CertResponse) SetStatusCode(v int32) *ConfigLayer7CertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CertResponse) SetBody(v *ConfigLayer7CertResponseBody) *ConfigLayer7CertResponse {
	s.Body = v
	return s
}

type ConfigLayer7RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// [{"ProxyPorts":[80,8080],"ProxyType":"http"},{"ProxyPorts":[443],"ProxyType":"https"}]rts\\":[443],\\"ProxyType\\":\\"https\\"}]
	ProxyTypeList *string `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty"`
	// example:
	//
	// [{"ProxyPorts":[80,8080],"ProxyType":"http"},{"ProxyPorts":[443],"ProxyType":"https"}]rts\\":[443],\\"ProxyType\\":\\"https\\"}]
	ProxyTypes []*string `json:"ProxyTypes,omitempty" xml:"ProxyTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s ConfigLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleRequest) SetDomain(v string) *ConfigLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetInstanceIds(v []*string) *ConfigLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

func (s *ConfigLayer7RuleRequest) SetProxyTypeList(v string) *ConfigLayer7RuleRequest {
	s.ProxyTypeList = &v
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

func (s *ConfigLayer7RuleRequest) SetResourceGroupId(v string) *ConfigLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7RuleRequest) SetRsType(v int32) *ConfigLayer7RuleRequest {
	s.RsType = &v
	return s
}

type ConfigLayer7RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ConfigLayer7RuleResponse) SetStatusCode(v int32) *ConfigLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7RuleResponse) SetBody(v *ConfigLayer7RuleResponseBody) *ConfigLayer7RuleResponse {
	s.Body = v
	return s
}

type CreateAsyncTaskRequest struct {
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"timestamp": 1530276554, "instanceId": "ddoscoo-woieuroi234"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) SetResourceGroupId(v string) *CreateAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskParams(v string) *CreateAsyncTaskRequest {
	s.TaskParams = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetTaskType(v int32) *CreateAsyncTaskRequest {
	s.TaskType = &v
	return s
}

type CreateAsyncTaskResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAsyncTaskResponse) SetStatusCode(v int32) *CreateAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsyncTaskResponse) SetBody(v *CreateAsyncTaskResponseBody) *CreateAsyncTaskResponse {
	s.Body = v
	return s
}

type CreateLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners   *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable *int32  `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
}

func (s CreateLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequest) SetListeners(v string) *CreateLayer4RuleRequest {
	s.Listeners = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetProxyEnable(v int32) *CreateLayer4RuleRequest {
	s.ProxyEnable = &v
	return s
}

type CreateLayer4RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateLayer4RuleResponse) SetStatusCode(v int32) *CreateLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayer4RuleResponse) SetBody(v *CreateLayer4RuleResponseBody) *CreateLayer4RuleResponse {
	s.Body = v
	return s
}

type CreateLayer7RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"ProxyRules":[{"ProxyPort":443,"RealServers":["1.1.1.1:443"]}],"ProxyType":"https"},{"ProxyRules":[{"ProxyPort":80,"RealServers":["1.1.1.1:80"]}],"ProxyType":"http"}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s CreateLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer7RuleRequest) SetDomain(v string) *CreateLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateLayer7RuleRequest) SetInstanceIds(v []*string) *CreateLayer7RuleRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateLayer7RuleRequest) SetResourceGroupId(v string) *CreateLayer7RuleRequest {
	s.ResourceGroupId = &v
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

type CreateLayer7RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateLayer7RuleResponse) SetStatusCode(v int32) *CreateLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLayer7RuleResponse) SetBody(v *CreateLayer7RuleResponseBody) *CreateLayer7RuleResponse {
	s.Body = v
	return s
}

type DeleteAsyncTaskRequest struct {
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAsyncTaskRequest) GoString() string {
	return s.String()
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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAsyncTaskResponse) SetStatusCode(v int32) *DeleteAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsyncTaskResponse) SetBody(v *DeleteAsyncTaskResponseBody) *DeleteAsyncTaskResponse {
	s.Body = v
	return s
}

type DeleteLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"InstanceId":"0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc","Protocol":"tcp","FrontendPort":80}
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DeleteLayer4RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer4RuleRequest) SetListeners(v string) *DeleteLayer4RuleRequest {
	s.Listeners = &v
	return s
}

type DeleteLayer4RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteLayer4RuleResponse) SetStatusCode(v int32) *DeleteLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer4RuleResponse) SetBody(v *DeleteLayer4RuleResponseBody) *DeleteLayer4RuleResponse {
	s.Body = v
	return s
}

type DeleteLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7CCRuleRequest) SetDomain(v string) *DeleteLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetName(v string) *DeleteLayer7CCRuleRequest {
	s.Name = &v
	return s
}

func (s *DeleteLayer7CCRuleRequest) SetResourceGroupId(v string) *DeleteLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteLayer7CCRuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteLayer7CCRuleResponse) SetStatusCode(v int32) *DeleteLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer7CCRuleResponse) SetBody(v *DeleteLayer7CCRuleResponseBody) *DeleteLayer7CCRuleResponse {
	s.Body = v
	return s
}

type DeleteLayer7RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteLayer7RuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLayer7RuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLayer7RuleRequest) SetDomain(v string) *DeleteLayer7RuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteLayer7RuleRequest) SetResourceGroupId(v string) *DeleteLayer7RuleRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteLayer7RuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteLayer7RuleResponse) SetStatusCode(v int32) *DeleteLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLayer7RuleResponse) SetBody(v *DeleteLayer7RuleResponseBody) *DeleteLayer7RuleResponse {
	s.Body = v
	return s
}

type DescribeBackSourceCidrRequest struct {
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeBackSourceCidrRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrRequest) SetIpVersion(v string) *DescribeBackSourceCidrRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetResourceGroupId(v string) *DescribeBackSourceCidrRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetSourceIp(v string) *DescribeBackSourceCidrRequest {
	s.SourceIp = &v
	return s
}

type DescribeBackSourceCidrResponseBody struct {
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackSourceCidrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBody) SetCidrList(v []*string) *DescribeBackSourceCidrResponseBody {
	s.CidrList = v
	return s
}

func (s *DescribeBackSourceCidrResponseBody) SetRequestId(v string) *DescribeBackSourceCidrResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBackSourceCidrResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackSourceCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeBackSourceCidrResponse) SetStatusCode(v int32) *DescribeBackSourceCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetBody(v *DescribeBackSourceCidrResponseBody) *DescribeBackSourceCidrResponse {
	s.Body = v
	return s
}

type DescribeBatchSlsDispatchStatusRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetLang(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.Lang = &v
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

func (s *DescribeBatchSlsDispatchStatusRequest) SetResourceGroupId(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusRequest) SetSourceIp(v string) *DescribeBatchSlsDispatchStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeBatchSlsDispatchStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId           *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlsConfigStatusList []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList `json:"SlsConfigStatusList,omitempty" xml:"SlsConfigStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchSlsDispatchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBatchSlsDispatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetRequestId(v string) *DescribeBatchSlsDispatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetSlsConfigStatusList(v []*DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList) *DescribeBatchSlsDispatchStatusResponseBody {
	s.SlsConfigStatusList = v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponseBody) SetTotalCount(v int32) *DescribeBatchSlsDispatchStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBatchSlsDispatchStatusResponseBodySlsConfigStatusList struct {
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchSlsDispatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeBatchSlsDispatchStatusResponse) SetStatusCode(v int32) *DescribeBatchSlsDispatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchSlsDispatchStatusResponse) SetBody(v *DescribeBatchSlsDispatchStatusResponseBody) *DescribeBatchSlsDispatchStatusResponse {
	s.Body = v
	return s
}

type DescribeDDoSEventsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457324
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457398
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsRequest) SetEip(v string) *DescribeDDoSEventsRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetEndTime(v int64) *DescribeDDoSEventsRequest {
	s.EndTime = &v
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

func (s *DescribeDDoSEventsRequest) SetResourceGroupId(v string) *DescribeDDoSEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetSourceIp(v string) *DescribeDDoSEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSEventsRequest) SetStartTime(v int64) *DescribeDDoSEventsRequest {
	s.StartTime = &v
	return s
}

type DescribeDDoSEventsResponseBody struct {
	Events []*DescribeDDoSEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDDoSEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBody) SetEvents(v []*DescribeDDoSEventsResponseBodyEvents) *DescribeDDoSEventsResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetRequestId(v string) *DescribeDDoSEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSEventsResponseBody) SetTotal(v int64) *DescribeDDoSEventsResponseBody {
	s.Total = &v
	return s
}

type DescribeDDoSEventsResponseBodyEvents struct {
	// example:
	//
	// 3289457398
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 3289457324
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// blackhole_start
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDDoSEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetEndTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetInterval(v int32) *DescribeDDoSEventsResponseBodyEvents {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStartTime(v int64) *DescribeDDoSEventsResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSEventsResponseBodyEvents) SetStatus(v string) *DescribeDDoSEventsResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeDDoSEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDDoSEventsResponse) SetStatusCode(v int32) *DescribeDDoSEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSEventsResponse) SetBody(v *DescribeDDoSEventsResponseBody) *DescribeDDoSEventsResponse {
	s.Body = v
	return s
}

type DescribeDDoSTrafficRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457398
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3289457324
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDDoSTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficRequest) SetEip(v string) *DescribeDDoSTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetEndTime(v int64) *DescribeDDoSTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetInterval(v int32) *DescribeDDoSTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetResourceGroupId(v string) *DescribeDDoSTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetSourceIp(v string) *DescribeDDoSTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDDoSTrafficRequest) SetStartTime(v int64) *DescribeDDoSTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeDDoSTrafficResponseBody struct {
	DDoSTrafficPoints []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints `json:"DDoSTrafficPoints,omitempty" xml:"DDoSTrafficPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 23482234
	DefenseInBytes *int64 `json:"DefenseInBytes,omitempty" xml:"DefenseInBytes,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 19284762
	SourceInBytes *int64 `json:"SourceInBytes,omitempty" xml:"SourceInBytes,omitempty"`
}

func (s DescribeDDoSTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBody) SetDDoSTrafficPoints(v []*DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) *DescribeDDoSTrafficResponseBody {
	s.DDoSTrafficPoints = v
	return s
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

type DescribeDDoSTrafficResponseBodyDDoSTrafficPoints struct {
	// example:
	//
	// 129867
	DefenseMaxInBps *int64 `json:"DefenseMaxInBps,omitempty" xml:"DefenseMaxInBps,omitempty"`
	// example:
	//
	// 129867
	SourceMaxInBps *int64 `json:"SourceMaxInBps,omitempty" xml:"SourceMaxInBps,omitempty"`
	// example:
	//
	// 234082304
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) GoString() string {
	return s.String()
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetDefenseMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.DefenseMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetSourceMaxInBps(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.SourceMaxInBps = &v
	return s
}

func (s *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints) SetTime(v int64) *DescribeDDoSTrafficResponseBodyDDoSTrafficPoints {
	s.Time = &v
	return s
}

type DescribeDDoSTrafficResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDDoSTrafficResponse) SetStatusCode(v int32) *DescribeDDoSTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSTrafficResponse) SetBody(v *DescribeDDoSTrafficResponseBody) *DescribeDDoSTrafficResponse {
	s.Body = v
	return s
}

type DescribeDefenseCountStatisticsRequest struct {
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDefenseCountStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsRequest) SetResourceGroupId(v string) *DescribeDefenseCountStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsRequest) SetSourceIp(v string) *DescribeDefenseCountStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeDefenseCountStatisticsResponseBody struct {
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" type:"Struct"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) *DescribeDefenseCountStatisticsResponseBody {
	s.DefenseCountStatistics = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetRequestId(v string) *DescribeDefenseCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics struct {
	// example:
	//
	// 0
	DefenseCountTotalUsageOfCurrentMonth *int32 `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty"`
	// example:
	//
	// 10
	FlowPackCountRemain *int32 `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty"`
	// example:
	//
	// 0
	MaxUsableDefenseCountCurrentMonth *int32 `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetDefenseCountTotalUsageOfCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.DefenseCountTotalUsageOfCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetFlowPackCountRemain(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.FlowPackCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetMaxUsableDefenseCountCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.MaxUsableDefenseCountCurrentMonth = &v
	return s
}

type DescribeDefenseCountStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseCountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDefenseCountStatisticsResponse) SetStatusCode(v int32) *DescribeDefenseCountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponse) SetBody(v *DescribeDefenseCountStatisticsResponseBody) *DescribeDefenseCountStatisticsResponse {
	s.Body = v
	return s
}

type DescribeDomainAccessModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainAccessModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeRequest) SetDomainList(v []*string) *DescribeDomainAccessModeRequest {
	s.DomainList = v
	return s
}

func (s *DescribeDomainAccessModeRequest) SetSourceIp(v string) *DescribeDomainAccessModeRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainAccessModeResponseBody struct {
	DomainModeList []*DescribeDomainAccessModeResponseBodyDomainModeList `json:"DomainModeList,omitempty" xml:"DomainModeList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAccessModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAccessModeResponseBody) SetDomainModeList(v []*DescribeDomainAccessModeResponseBodyDomainModeList) *DescribeDomainAccessModeResponseBody {
	s.DomainModeList = v
	return s
}

func (s *DescribeDomainAccessModeResponseBody) SetRequestId(v string) *DescribeDomainAccessModeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainAccessModeResponseBodyDomainModeList struct {
	// example:
	//
	// 1
	AccessMode *int32 `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainAccessModeResponse) SetStatusCode(v int32) *DescribeDomainAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAccessModeResponse) SetBody(v *DescribeDomainAccessModeResponseBody) *DescribeDomainAccessModeResponse {
	s.Body = v
	return s
}

type DescribeDomainAttackEventListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1668740400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1681966800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListRequest) SetDomain(v string) *DescribeDomainAttackEventListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetEndTime(v int64) *DescribeDomainAttackEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetPageNo(v int32) *DescribeDomainAttackEventListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetPageSize(v int32) *DescribeDomainAttackEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventListRequest) SetStartTime(v int64) *DescribeDomainAttackEventListRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainAttackEventListResponseBody struct {
	DataList []*DescribeDomainAttackEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDomainAttackEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponseBody) SetDataList(v []*DescribeDomainAttackEventListResponseBodyDataList) *DescribeDomainAttackEventListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDomainAttackEventListResponseBody) SetRequestId(v string) *DescribeDomainAttackEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBody) SetTotal(v int64) *DescribeDomainAttackEventListResponseBody {
	s.Total = &v
	return s
}

type DescribeDomainAttackEventListResponseBodyDataList struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1670918400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 300
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// example:
	//
	// 1666083600
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackEventListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetDomain(v string) *DescribeDomainAttackEventListResponseBodyDataList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetEndTime(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetMaxQps(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainAttackEventListResponseBodyDataList) SetStartTime(v int64) *DescribeDomainAttackEventListResponseBodyDataList {
	s.StartTime = &v
	return s
}

type DescribeDomainAttackEventListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAttackEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAttackEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventListResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackEventListResponse) SetStatusCode(v int32) *DescribeDomainAttackEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAttackEventListResponse) SetBody(v *DescribeDomainAttackEventListResponseBody) *DescribeDomainAttackEventListResponse {
	s.Body = v
	return s
}

type DescribeDomainAttackMaxQpsRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1667801940
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1657562370
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAttackMaxQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsRequest) SetDomain(v string) *DescribeDomainAttackMaxQpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetEndTime(v int64) *DescribeDomainAttackMaxQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetResourceGroupId(v string) *DescribeDomainAttackMaxQpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsRequest) SetStartTime(v int64) *DescribeDomainAttackMaxQpsRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainAttackMaxQpsResponseBody struct {
	// example:
	//
	// 613
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// example:
	//
	// 62F9BD81-8BCA-5B23-A3CB-3FB7CEB7A4CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainAttackMaxQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsResponseBody) SetQps(v string) *DescribeDomainAttackMaxQpsResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponseBody) SetRequestId(v string) *DescribeDomainAttackMaxQpsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainAttackMaxQpsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainAttackMaxQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainAttackMaxQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainAttackMaxQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackMaxQpsResponse) SetHeaders(v map[string]*string) *DescribeDomainAttackMaxQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponse) SetStatusCode(v int32) *DescribeDomainAttackMaxQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainAttackMaxQpsResponse) SetBody(v *DescribeDomainAttackMaxQpsResponseBody) *DescribeDomainAttackMaxQpsResponse {
	s.Body = v
	return s
}

type DescribeDomainOverviewRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1651809600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1619798400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewRequest) SetDomain(v string) *DescribeDomainOverviewRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetEndTime(v int64) *DescribeDomainOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetResourceGroupId(v string) *DescribeDomainOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetStartTime(v int64) *DescribeDomainOverviewRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainOverviewResponseBody struct {
	// example:
	//
	// 0
	MaxHttp *int64 `json:"MaxHttp,omitempty" xml:"MaxHttp,omitempty"`
	// example:
	//
	// 0
	MaxHttps *int64 `json:"MaxHttps,omitempty" xml:"MaxHttps,omitempty"`
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttp(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttp = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetMaxHttps(v int64) *DescribeDomainOverviewResponseBody {
	s.MaxHttps = &v
	return s
}

func (s *DescribeDomainOverviewResponseBody) SetRequestId(v string) *DescribeDomainOverviewResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainOverviewResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponse) SetHeaders(v map[string]*string) *DescribeDomainOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainOverviewResponse) SetStatusCode(v int32) *DescribeDomainOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainOverviewResponse) SetBody(v *DescribeDomainOverviewResponseBody) *DescribeDomainOverviewResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1657162260
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1800
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1672362360
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListRequest) SetDomain(v string) *DescribeDomainQpsListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetEndTime(v int64) *DescribeDomainQpsListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetInterval(v int64) *DescribeDomainQpsListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetResourceGroupId(v string) *DescribeDomainQpsListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsListRequest) SetStartTime(v int64) *DescribeDomainQpsListRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainQpsListResponseBody struct {
	DataList  []*DescribeDomainQpsListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainQpsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponseBody) SetDataList(v []*DescribeDomainQpsListResponseBodyDataList) *DescribeDomainQpsListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDomainQpsListResponseBody) SetRequestId(v string) *DescribeDomainQpsListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainQpsListResponseBodyDataList struct {
	AttackQps    *int64 `json:"AttackQps,omitempty" xml:"AttackQps,omitempty"`
	CacheHits    *int64 `json:"CacheHits,omitempty" xml:"CacheHits,omitempty"`
	Index        *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	MaxAttackQps *int64 `json:"MaxAttackQps,omitempty" xml:"MaxAttackQps,omitempty"`
	MaxNormalQps *int64 `json:"MaxNormalQps,omitempty" xml:"MaxNormalQps,omitempty"`
	MaxQps       *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	TotalCount   *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalQps     *int64 `json:"TotalQps,omitempty" xml:"TotalQps,omitempty"`
}

func (s DescribeDomainQpsListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetAttackQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.AttackQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetCacheHits(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.CacheHits = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetIndex(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxAttackQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxAttackQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxNormalQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxNormalQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetMaxQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.MaxQps = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetTotalCount(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainQpsListResponseBodyDataList) SetTotalQps(v int64) *DescribeDomainQpsListResponseBodyDataList {
	s.TotalQps = &v
	return s
}

type DescribeDomainQpsListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQpsListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsListResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsListResponse) SetStatusCode(v int32) *DescribeDomainQpsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsListResponse) SetBody(v *DescribeDomainQpsListResponseBody) *DescribeDomainQpsListResponse {
	s.Body = v
	return s
}

type DescribeDomainQpsWithCacheRequest struct {
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1577796336
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1577794536
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainQpsWithCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheRequest) SetDomain(v string) *DescribeDomainQpsWithCacheRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetEndTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetResourceGroupId(v string) *DescribeDomainQpsWithCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetSourceIp(v string) *DescribeDomainQpsWithCacheRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainQpsWithCacheRequest) SetStartTime(v int64) *DescribeDomainQpsWithCacheRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainQpsWithCacheResponseBody struct {
	Blocks     []*string `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	CacheHits  []*string `json:"CacheHits,omitempty" xml:"CacheHits,omitempty" type:"Repeated"`
	CcBlockQps []*string `json:"CcBlockQps,omitempty" xml:"CcBlockQps,omitempty" type:"Repeated"`
	CcJsQps    []*string `json:"CcJsQps,omitempty" xml:"CcJsQps,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	Interval      *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IpBlockQps    []*string `json:"IpBlockQps,omitempty" xml:"IpBlockQps,omitempty" type:"Repeated"`
	PreciseBlocks []*string `json:"PreciseBlocks,omitempty" xml:"PreciseBlocks,omitempty" type:"Repeated"`
	PreciseJsQps  []*string `json:"PreciseJsQps,omitempty" xml:"PreciseJsQps,omitempty" type:"Repeated"`
	RegionBlocks  []*string `json:"RegionBlocks,omitempty" xml:"RegionBlocks,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1577794500
	StartTime *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Totals    []*string `json:"Totals,omitempty" xml:"Totals,omitempty" type:"Repeated"`
}

func (s DescribeDomainQpsWithCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Blocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCacheHits(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CacheHits = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetCcJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.CcJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetInterval(v int32) *DescribeDomainQpsWithCacheResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetIpBlockQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.IpBlockQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetPreciseJsQps(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.PreciseJsQps = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRegionBlocks(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.RegionBlocks = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetRequestId(v string) *DescribeDomainQpsWithCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetStartTime(v int64) *DescribeDomainQpsWithCacheResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponseBody) SetTotals(v []*string) *DescribeDomainQpsWithCacheResponseBody {
	s.Totals = v
	return s
}

type DescribeDomainQpsWithCacheResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsWithCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainQpsWithCacheResponse) SetStatusCode(v int32) *DescribeDomainQpsWithCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetBody(v *DescribeDomainQpsWithCacheResponseBody) *DescribeDomainQpsWithCacheResponse {
	s.Body = v
	return s
}

type DescribeDomainSlsStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSlsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusRequest) SetDomain(v string) *DescribeDomainSlsStatusRequest {
	s.Domain = &v
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

func (s *DescribeDomainSlsStatusRequest) SetSourceIp(v string) *DescribeDomainSlsStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainSlsStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ddoscoo-logstore
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// example:
	//
	// ddoscoo-project-xxxx-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// example:
	//
	// true
	SlsStatus *bool `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
}

func (s DescribeDomainSlsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSlsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSlsStatusResponseBody) SetRequestId(v string) *DescribeDomainSlsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsLogstore(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsProject(v string) *DescribeDomainSlsStatusResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeDomainSlsStatusResponseBody) SetSlsStatus(v bool) *DescribeDomainSlsStatusResponseBody {
	s.SlsStatus = &v
	return s
}

type DescribeDomainSlsStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSlsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeDomainSlsStatusResponse) SetStatusCode(v int32) *DescribeDomainSlsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSlsStatusResponse) SetBody(v *DescribeDomainSlsStatusResponseBody) *DescribeDomainSlsStatusResponse {
	s.Body = v
	return s
}

type DescribeDomainStatusCodeListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1800
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// upstream
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainStatusCodeListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListRequest) SetDomain(v string) *DescribeDomainStatusCodeListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetEndTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetInterval(v int64) *DescribeDomainStatusCodeListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetQueryType(v string) *DescribeDomainStatusCodeListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetStartTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.StartTime = &v
	return s
}

type DescribeDomainStatusCodeListResponseBody struct {
	// example:
	//
	// 3B63C0DD-8AC5-44B2-95D6-064CA9296B9C
	RequestId      *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusCodeList []*DescribeDomainStatusCodeListResponseBodyStatusCodeList `json:"StatusCodeList,omitempty" xml:"StatusCodeList,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatusCodeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBody) SetRequestId(v string) *DescribeDomainStatusCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBody) SetStatusCodeList(v []*DescribeDomainStatusCodeListResponseBodyStatusCodeList) *DescribeDomainStatusCodeListResponseBody {
	s.StatusCodeList = v
	return s
}

type DescribeDomainStatusCodeListResponseBodyStatusCodeList struct {
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 320
	Status200 *int64 `json:"Status200,omitempty" xml:"Status200,omitempty"`
	// example:
	//
	// 5776
	Status2XX *int64 `json:"Status2XX,omitempty" xml:"Status2XX,omitempty"`
	// example:
	//
	// 0
	Status3XX *int64 `json:"Status3XX,omitempty" xml:"Status3XX,omitempty"`
	// example:
	//
	// 0
	Status403 *int64 `json:"Status403,omitempty" xml:"Status403,omitempty"`
	// example:
	//
	// 34
	Status404 *int64 `json:"Status404,omitempty" xml:"Status404,omitempty"`
	// example:
	//
	// 11
	Status405 *int64 `json:"Status405,omitempty" xml:"Status405,omitempty"`
	Status410 *int64 `json:"Status410,omitempty" xml:"Status410,omitempty"`
	Status499 *int64 `json:"Status499,omitempty" xml:"Status499,omitempty"`
	// example:
	//
	// 168
	Status4XX *int64 `json:"Status4XX,omitempty" xml:"Status4XX,omitempty"`
	// example:
	//
	// 0
	Status501 *int64 `json:"Status501,omitempty" xml:"Status501,omitempty"`
	// example:
	//
	// 3
	Status502 *int64 `json:"Status502,omitempty" xml:"Status502,omitempty"`
	// example:
	//
	// 0
	Status503 *int64 `json:"Status503,omitempty" xml:"Status503,omitempty"`
	// example:
	//
	// 0
	Status504 *int64 `json:"Status504,omitempty" xml:"Status504,omitempty"`
	// example:
	//
	// 7
	Status5XX *int64 `json:"Status5XX,omitempty" xml:"Status5XX,omitempty"`
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponseBodyStatusCodeList) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetIndex(v int32) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Index = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus200(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status200 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus2XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status2XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus3XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status3XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus403(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status403 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus404(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status404 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus405(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status405 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus410(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status410 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus499(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status499 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus4XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status4XX = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus501(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status501 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus502(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status502 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus503(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status503 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus504(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status504 = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponseBodyStatusCodeList) SetStatus5XX(v int64) *DescribeDomainStatusCodeListResponseBodyStatusCodeList {
	s.Status5XX = &v
	return s
}

type DescribeDomainStatusCodeListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainStatusCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainStatusCodeListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainStatusCodeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListResponse) SetHeaders(v map[string]*string) *DescribeDomainStatusCodeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) SetStatusCode(v int32) *DescribeDomainStatusCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainStatusCodeListResponse) SetBody(v *DescribeDomainStatusCodeListResponseBody) *DescribeDomainStatusCodeListResponse {
	s.Body = v
	return s
}

type DescribeDomainsRequest struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	QueryDomainPattern *string `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) SetDomain(v string) *DescribeDomainsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainsRequest) SetInstanceIds(v []*string) *DescribeDomainsRequest {
	s.InstanceIds = v
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

func (s *DescribeDomainsRequest) SetQueryDomainPattern(v string) *DescribeDomainsRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetSourceIp(v string) *DescribeDomainsRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainsResponseBody struct {
	Domains []*DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CcEnabled *bool `json:"CcEnabled,omitempty" xml:"CcEnabled,omitempty"`
	// example:
	//
	// true
	CcRuleEnabled *bool `json:"CcRuleEnabled,omitempty" xml:"CcRuleEnabled,omitempty"`
	// example:
	//
	// normal
	CcTemplate *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	// example:
	//
	// testCertName
	CertName   *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// example:
	//
	// xxxxxxx.aliyunddos1006.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	Http2Enable   *bool                                              `json:"Http2Enable,omitempty" xml:"Http2Enable,omitempty"`
	ProxyTypeList []*DescribeDomainsResponseBodyDomainsProxyTypeList `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty" type:"Repeated"`
	RealServers   []*DescribeDomainsResponseBodyDomainsRealServers   `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	SslCiphers *string `json:"SslCiphers,omitempty" xml:"SslCiphers,omitempty"`
	// example:
	//
	// xx
	SslProtocols *string   `json:"SslProtocols,omitempty" xml:"SslProtocols,omitempty"`
	WhiteList    []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) SetBlackList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.BlackList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcRuleEnabled(v bool) *DescribeDomainsResponseBodyDomains {
	s.CcRuleEnabled = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCcTemplate(v string) *DescribeDomainsResponseBodyDomains {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCertName(v string) *DescribeDomainsResponseBodyDomains {
	s.CertName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetCertRegion(v string) *DescribeDomainsResponseBodyDomains {
	s.CertRegion = &v
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

func (s *DescribeDomainsResponseBodyDomains) SetHttp2Enable(v bool) *DescribeDomainsResponseBodyDomains {
	s.Http2Enable = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetProxyTypeList(v []*DescribeDomainsResponseBodyDomainsProxyTypeList) *DescribeDomainsResponseBodyDomains {
	s.ProxyTypeList = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetRealServers(v []*DescribeDomainsResponseBodyDomainsRealServers) *DescribeDomainsResponseBodyDomains {
	s.RealServers = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslCiphers(v string) *DescribeDomainsResponseBodyDomains {
	s.SslCiphers = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetSslProtocols(v string) *DescribeDomainsResponseBodyDomains {
	s.SslProtocols = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) SetWhiteList(v []*string) *DescribeDomainsResponseBodyDomains {
	s.WhiteList = v
	return s
}

type DescribeDomainsResponseBodyDomainsProxyTypeList struct {
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// example:
	//
	// http
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
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
	// example:
	//
	// 1.1.1.1
	RealServer *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
	// example:
	//
	// 0
	RsType *int32 `json:"RsType,omitempty" xml:"RsType,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsRealServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsRealServers) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRealServer(v string) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RealServer = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsRealServers) SetRsType(v int32) *DescribeDomainsResponseBodyDomainsRealServers {
	s.RsType = &v
	return s
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeElasticBandwidthSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeElasticBandwidthSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecRequest) SetInstanceId(v string) *DescribeElasticBandwidthSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeElasticBandwidthSpecRequest) SetSourceIp(v string) *DescribeElasticBandwidthSpecRequest {
	s.SourceIp = &v
	return s
}

type DescribeElasticBandwidthSpecResponseBody struct {
	ElasticBandwidthSpec []*string `json:"ElasticBandwidthSpec,omitempty" xml:"ElasticBandwidthSpec,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticBandwidthSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetElasticBandwidthSpec(v []*string) *DescribeElasticBandwidthSpecResponseBody {
	s.ElasticBandwidthSpec = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponseBody) SetRequestId(v string) *DescribeElasticBandwidthSpecResponseBody {
	s.RequestId = &v
	return s
}

type DescribeElasticBandwidthSpecResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticBandwidthSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeElasticBandwidthSpecResponse) SetStatusCode(v int32) *DescribeElasticBandwidthSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetBody(v *DescribeElasticBandwidthSpecResponseBody) *DescribeElasticBandwidthSpecResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","Protocol":"tcp","FrontendPort":80}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeHealthCheckListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListRequest) SetListeners(v string) *DescribeHealthCheckListRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeHealthCheckListRequest) SetSourceIp(v string) *DescribeHealthCheckListRequest {
	s.SourceIp = &v
	return s
}

type DescribeHealthCheckListResponseBody struct {
	Listeners []*DescribeHealthCheckListResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 233
	FrontendPort *int32                                                   `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	HealthCheck  *DescribeHealthCheckListResponseBodyListenersHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
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

func (s *DescribeHealthCheckListResponseBodyListeners) SetHealthCheck(v *DescribeHealthCheckListResponseBodyListenersHealthCheck) *DescribeHealthCheckListResponseBodyListeners {
	s.HealthCheck = v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetInstanceId(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListeners) SetProtocol(v string) *DescribeHealthCheckListResponseBodyListeners {
	s.Protocol = &v
	return s
}

type DescribeHealthCheckListResponseBodyListenersHealthCheck struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 500
	Down *int32 `json:"Down,omitempty" xml:"Down,omitempty"`
	// example:
	//
	// 15
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 233
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// tcp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1000
	Up *int32 `json:"Up,omitempty" xml:"Up,omitempty"`
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckListResponseBodyListenersHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDomain(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetDown(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetInterval(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetPort(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetTimeout(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetType(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUp(v int32) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckListResponseBodyListenersHealthCheck) SetUri(v string) *DescribeHealthCheckListResponseBodyListenersHealthCheck {
	s.Uri = &v
	return s
}

type DescribeHealthCheckListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeHealthCheckListResponse) SetStatusCode(v int32) *DescribeHealthCheckListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckListResponse) SetBody(v *DescribeHealthCheckListResponseBody) *DescribeHealthCheckListResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckStatusListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","Protocol":"tcp","FrontendPort":80}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeHealthCheckStatusListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListRequest) SetListeners(v string) *DescribeHealthCheckStatusListRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeHealthCheckStatusListRequest) SetSourceIp(v string) *DescribeHealthCheckStatusListRequest {
	s.SourceIp = &v
	return s
}

type DescribeHealthCheckStatusListResponseBody struct {
	HealthCheckStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList `json:"HealthCheckStatusList,omitempty" xml:"HealthCheckStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol             *string                                                                               `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RealServerStatusList []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList `json:"RealServerStatusList,omitempty" xml:"RealServerStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetFrontendPort(v int32) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetInstanceId(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetProtocol(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetRealServerStatusList(v []*DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.RealServerStatusList = v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusList {
	s.Status = &v
	return s
}

type DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList struct {
	// example:
	//
	// 1.1.1.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetAddress(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Address = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList) SetStatus(v string) *DescribeHealthCheckStatusListResponseBodyHealthCheckStatusListRealServerStatusList {
	s.Status = &v
	return s
}

type DescribeHealthCheckStatusListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHealthCheckStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeHealthCheckStatusListResponse) SetStatusCode(v int32) *DescribeHealthCheckStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHealthCheckStatusListResponse) SetBody(v *DescribeHealthCheckStatusListResponseBody) *DescribeHealthCheckStatusListResponse {
	s.Body = v
	return s
}

type DescribeInstanceDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["ddoscoo-cn-XXXX1", "ddoscoo-cn-XXXX2"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceDetailsRequest) SetSourceIp(v string) *DescribeInstanceDetailsRequest {
	s.SourceIp = &v
	return s
}

type DescribeInstanceDetailsResponseBody struct {
	InstanceDetails []*DescribeInstanceDetailsResponseBodyInstanceDetails `json:"InstanceDetails,omitempty" xml:"InstanceDetails,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// coop-line-001
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
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

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetInstanceId(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetails) SetLine(v string) *DescribeInstanceDetailsResponseBodyInstanceDetails {
	s.Line = &v
	return s
}

type DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList struct {
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetEip(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Eip = &v
	return s
}

func (s *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList) SetStatus(v string) *DescribeInstanceDetailsResponseBodyInstanceDetailsEipInfoList {
	s.Status = &v
	return s
}

type DescribeInstanceDetailsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceDetailsResponse) SetStatusCode(v int32) *DescribeInstanceDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDetailsResponse) SetBody(v *DescribeInstanceDetailsResponseBody) *DescribeInstanceDetailsResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ["ddoscoo-cn-XXXXX"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	// example:
	//
	// 20000
	BandwidthMbps *int32 `json:"BandwidthMbps,omitempty" xml:"BandwidthMbps,omitempty"`
	// example:
	//
	// 20
	BaseBandwidth *int32 `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	// example:
	//
	// 10
	DefenseCount *int32 `json:"DefenseCount,omitempty" xml:"DefenseCount,omitempty"`
	// example:
	//
	// 50
	DomainLimit *int32 `json:"DomainLimit,omitempty" xml:"DomainLimit,omitempty"`
	// example:
	//
	// 20
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 50
	PortLimit *int32 `json:"PortLimit,omitempty" xml:"PortLimit,omitempty"`
	// example:
	//
	// 1000
	QpsLimit *int32 `json:"QpsLimit,omitempty" xml:"QpsLimit,omitempty"`
	// example:
	//
	// 10
	SiteLimit *int32 `json:"SiteLimit,omitempty" xml:"SiteLimit,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBandwidthMbps(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BandwidthMbps = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetBaseBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDomainLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DomainLimit = &v
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

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPortLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PortLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetQpsLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.QpsLimit = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetSiteLimit(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.SiteLimit = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceSpecsResponse) SetStatusCode(v int32) *DescribeInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","InstanceId":"ddoscoo-cn-YYYYY"}]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetInstanceIds(v string) *DescribeInstanceStatisticsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeInstanceStatisticsResponseBody struct {
	InstanceStatistics []*DescribeInstanceStatisticsResponseBodyInstanceStatistics `json:"InstanceStatistics,omitempty" xml:"InstanceStatistics,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	DefenseCountUsage *int32 `json:"DefenseCountUsage,omitempty" xml:"DefenseCountUsage,omitempty"`
	// example:
	//
	// 10
	DomainUsage *int32 `json:"DomainUsage,omitempty" xml:"DomainUsage,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	PortUsage *int32 `json:"PortUsage,omitempty" xml:"PortUsage,omitempty"`
	// example:
	//
	// 1
	SiteUsage *int32 `json:"SiteUsage,omitempty" xml:"SiteUsage,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBodyInstanceStatistics) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDefenseCountUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DefenseCountUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetDomainUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.DomainUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetInstanceId(v string) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetPortUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.PortUsage = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyInstanceStatistics) SetSiteUsage(v int32) *DescribeInstanceStatisticsResponseBodyInstanceStatistics {
	s.SiteUsage = &v
	return s
}

type DescribeInstanceStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstanceStatisticsResponse) SetStatusCode(v int32) *DescribeInstanceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1578931200000
	ExpireEndTime *int64 `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	// example:
	//
	// 1578931200000
	ExpireStartTime *int64 `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	// example:
	//
	// ["ddoscoo-cn-XXXXX"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// testRemark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1
	Status []*int32                       `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tag    []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetEdition(v int32) *DescribeInstancesRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnabled(v int32) *DescribeInstancesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireEndTime(v int64) *DescribeInstancesRequest {
	s.ExpireEndTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireStartTime(v int64) *DescribeInstancesRequest {
	s.ExpireStartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
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

func (s *DescribeInstancesRequest) SetRemark(v string) *DescribeInstancesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v []*int32) *DescribeInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
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
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// 0
	DebtStatus *int32 `json:"DebtStatus,omitempty" xml:"DebtStatus,omitempty"`
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 2308402384
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2308402384
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// testRemark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetDebtStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.DebtStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEdition(v int32) *DescribeInstancesResponseBodyInstances {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnabled(v int32) *DescribeInstancesResponseBodyInstances {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
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

func (s *DescribeInstancesResponseBodyInstances) SetRemark(v string) *DescribeInstancesResponseBodyInstances {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStatus(v int32) *DescribeInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeIpTrafficRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1536734120
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 233
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// http
	QueryProtocol *string `json:"QueryProtocol,omitempty" xml:"QueryProtocol,omitempty"`
	// example:
	//
	// xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1536734112
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeIpTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficRequest) SetEip(v string) *DescribeIpTrafficRequest {
	s.Eip = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetEndTime(v int64) *DescribeIpTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetInterval(v int32) *DescribeIpTrafficRequest {
	s.Interval = &v
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

func (s *DescribeIpTrafficRequest) SetResourceGroupId(v string) *DescribeIpTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetSourceIp(v string) *DescribeIpTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpTrafficRequest) SetStartTime(v int64) *DescribeIpTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeIpTrafficResponseBody struct {
	// example:
	//
	// 10000
	AvgInBps *int64 `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	// example:
	//
	// 10000
	AvgOutBps       *int64                                          `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	IpTrafficPoints []*DescribeIpTrafficResponseBodyIpTrafficPoints `json:"IpTrafficPoints,omitempty" xml:"IpTrafficPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// example:
	//
	// 10000
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpTrafficResponseBody) SetAvgInBps(v int64) *DescribeIpTrafficResponseBody {
	s.AvgInBps = &v
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

func (s *DescribeIpTrafficResponseBody) SetMaxInBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetMaxOutBps(v int64) *DescribeIpTrafficResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeIpTrafficResponseBody) SetRequestId(v string) *DescribeIpTrafficResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIpTrafficResponseBodyIpTrafficPoints struct {
	// example:
	//
	// 100
	ActConns *int32 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 100
	InactConns *int32 `json:"InactConns,omitempty" xml:"InactConns,omitempty"`
	// example:
	//
	// 10000
	MaxInbps *int64 `json:"MaxInbps,omitempty" xml:"MaxInbps,omitempty"`
	// example:
	//
	// 10000
	MaxOutbps *int64 `json:"MaxOutbps,omitempty" xml:"MaxOutbps,omitempty"`
	// example:
	//
	// 1536734112
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
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

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetCps(v int32) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Cps = &v
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

func (s *DescribeIpTrafficResponseBodyIpTrafficPoints) SetTime(v int64) *DescribeIpTrafficResponseBodyIpTrafficPoints {
	s.Time = &v
	return s
}

type DescribeIpTrafficResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeIpTrafficResponse) SetStatusCode(v int32) *DescribeIpTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpTrafficResponse) SetBody(v *DescribeIpTrafficResponseBody) *DescribeIpTrafficResponse {
	s.Body = v
	return s
}

type DescribeLayer4RuleAttributesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-XXXXX","Protocol":"tcp","FrontendPort":80}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer4RuleAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesRequest) SetListeners(v string) *DescribeLayer4RuleAttributesRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeLayer4RuleAttributesRequest) SetSourceIp(v string) *DescribeLayer4RuleAttributesRequest {
	s.SourceIp = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBody struct {
	Listeners []*DescribeLayer4RuleAttributesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Config *DescribeLayer4RuleAttributesResponseBodyListenersConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetConfig(v *DescribeLayer4RuleAttributesResponseBodyListenersConfig) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Config = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Protocol = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfig struct {
	Cc *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc `json:"Cc,omitempty" xml:"Cc,omitempty" type:"Struct"`
	// example:
	//
	// on
	NodataConn *string                                                            `json:"NodataConn,omitempty" xml:"NodataConn,omitempty"`
	PayloadLen *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" type:"Struct"`
	// example:
	//
	// 0
	PersistenceTimeout *int32                                                         `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Sla                *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla    `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	Slimit             *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
	// example:
	//
	// on
	Synproxy *string `json:"Synproxy,omitempty" xml:"Synproxy,omitempty"`
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

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetNodataConn(v string) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.NodataConn = &v
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
	// example:
	//
	// 5
	Cnt *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// example:
	//
	// 60
	During *int32 `json:"During,omitempty" xml:"During,omitempty"`
	// example:
	//
	// 1800
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetCnt(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Cnt = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetDuring(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetExpires(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetType(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Type = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen struct {
	// example:
	//
	// 2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
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
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 0
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// example:
	//
	// 1000
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.MaxconnEnable = &v
	return s
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit struct {
	// example:
	//
	// 0
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 0
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// example:
	//
	// 2
	CpsMode *int32 `json:"CpsMode,omitempty" xml:"CpsMode,omitempty"`
	// example:
	//
	// 1000
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	// example:
	//
	// 0
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetBps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsMode(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsMode = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetPps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Pps = &v
	return s
}

type DescribeLayer4RuleAttributesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer4RuleAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeLayer4RuleAttributesResponse) SetStatusCode(v int32) *DescribeLayer4RuleAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponse) SetBody(v *DescribeLayer4RuleAttributesResponseBody) *DescribeLayer4RuleAttributesResponse {
	s.Body = v
	return s
}

type DescribeLayer4RulesRequest struct {
	// The type of forwarding protocol. Values:
	//
	// - **tcp**: Indicates TCP protocol.
	//
	// - **udp**: Indicates UDP protocol.
	//
	// example:
	//
	// tcp
	ForwardProtocol *string `json:"ForwardProtocol,omitempty" xml:"ForwardProtocol,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the DDoS protection instance to be queried.
	//
	// > You can call [DescribeInstances](https://help.aliyun.com/document_detail/91478.html) to query all DDoS protection instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-zvp2ay9b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// In paginated queries, specifies which page of data to return. The minimum value is **1**, indicating the first page of data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// In paginated queries, specifies the number of results per page. The maximum value is **50**, indicating that each page can contain up to 50 results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request. You do not need to fill this in; it is automatically obtained by the system.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer4RulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesRequest) SetForwardProtocol(v string) *DescribeLayer4RulesRequest {
	s.ForwardProtocol = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetFrontendPort(v int32) *DescribeLayer4RulesRequest {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesRequest) SetInstanceId(v string) *DescribeLayer4RulesRequest {
	s.InstanceId = &v
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

func (s *DescribeLayer4RulesRequest) SetSourceIp(v string) *DescribeLayer4RulesRequest {
	s.SourceIp = &v
	return s
}

type DescribeLayer4RulesResponseBody struct {
	// Detailed configuration of port forwarding rules, including the forwarding port, forwarding protocol, and origin server addresses, etc.
	Listeners []*DescribeLayer4RulesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The ID of the current request.
	//
	// example:
	//
	// 949919A2-6636-1444-9213-AB27DD88AAA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned results.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The origin server port.
	//
	// example:
	//
	// 233
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The origin mode. Values:
	//
	// - **0**: Indicates the default origin mode.
	//
	// - **1**: Indicates the primary/backup origin mode.
	//
	// example:
	//
	// 0
	BakMode *int32 `json:"BakMode,omitempty" xml:"BakMode,omitempty"`
	// The currently effective origin server type. Values:
	//
	// - **1**: Indicates that the primary origin server settings are in effect (DDoS protection forwards business traffic to the primary origin server IP address).
	//
	// - **2**: Indicates that the backup origin server settings are in effect (DDoS protection forwards business traffic to the backup origin server IP address).
	//
	// example:
	//
	// 1
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// The IP address of the DDoS protection instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the DDoS protection instance.
	//
	// example:
	//
	// ddoscoo-cn-zvp2ay9b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the rule was automatically created. Values:
	//
	// - **true**: Indicates that the rule was automatically created by DDoS protection.
	//
	// - **false**: Indicates that the rule was manually created by you.
	//
	// example:
	//
	// false
	IsAutoCreate *bool `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	// Payload rule module switch. Values:
	//
	// - 1: Enabled
	//
	// - 0: Disabled
	//
	// example:
	//
	// 0
	PayloadRuleEnable *int64 `json:"PayloadRuleEnable,omitempty" xml:"PayloadRuleEnable,omitempty"`
	// The type of forwarding protocol.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Traffic diversion switch. Values:
	//
	// - **0*	- Off.
	//
	// - **1*	- On.
	//
	// example:
	//
	// 0
	ProxyEnable *int32 `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	// Traffic diversion status. Values:
	//
	// - on: Diversion is effective
	//
	// - off: Diversion is ineffective
	//
	// example:
	//
	// on
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The list of origin server IP addresses.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The remarks for the port forwarding rule.
	//
	// example:
	//
	// test-remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeLayer4RulesResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetBackendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.BackendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetBakMode(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.BakMode = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetCurrentIndex(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.CurrentIndex = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetEip(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Eip = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetIsAutoCreate(v bool) *DescribeLayer4RulesResponseBodyListeners {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetPayloadRuleEnable(v int64) *DescribeLayer4RulesResponseBodyListeners {
	s.PayloadRuleEnable = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProxyEnable(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.ProxyEnable = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProxyStatus(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetRealServers(v []*string) *DescribeLayer4RulesResponseBodyListeners {
	s.RealServers = v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetRemark(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Remark = &v
	return s
}

type DescribeLayer4RulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeLayer4RulesResponse) SetStatusCode(v int32) *DescribeLayer4RulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer4RulesResponse) SetBody(v *DescribeLayer4RulesResponseBody) *DescribeLayer4RulesResponse {
	s.Body = v
	return s
}

type DescribeLayer7CCRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLayer7CCRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesRequest) GoString() string {
	return s.String()
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

func (s *DescribeLayer7CCRulesRequest) SetResourceGroupId(v string) *DescribeLayer7CCRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLayer7CCRulesRequest) SetSourceIp(v string) *DescribeLayer7CCRulesRequest {
	s.SourceIp = &v
	return s
}

type DescribeLayer7CCRulesResponseBody struct {
	Layer7CCRules []*DescribeLayer7CCRulesResponseBodyLayer7CCRules `json:"Layer7CCRules,omitempty" xml:"Layer7CCRules,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// match
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// testCcRule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1000
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// /a/b/c
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeLayer7CCRulesResponseBodyLayer7CCRules) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetAct(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Act = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetCount(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Count = &v
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

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetTtl(v int32) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Ttl = &v
	return s
}

func (s *DescribeLayer7CCRulesResponseBodyLayer7CCRules) SetUri(v string) *DescribeLayer7CCRulesResponseBodyLayer7CCRules {
	s.Uri = &v
	return s
}

type DescribeLayer7CCRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer7CCRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeLayer7CCRulesResponse) SetStatusCode(v int32) *DescribeLayer7CCRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetBody(v *DescribeLayer7CCRulesResponseBody) *DescribeLayer7CCRulesResponse {
	s.Body = v
	return s
}

type DescribeLogStoreExistStatusRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeLogStoreExistStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusRequest) SetLang(v string) *DescribeLogStoreExistStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetResourceGroupId(v string) *DescribeLogStoreExistStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLogStoreExistStatusRequest) SetSourceIp(v string) *DescribeLogStoreExistStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeLogStoreExistStatusResponseBody struct {
	// example:
	//
	// true
	ExistStatus *bool `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogStoreExistStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogStoreExistStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreExistStatusResponseBody) SetExistStatus(v bool) *DescribeLogStoreExistStatusResponseBody {
	s.ExistStatus = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponseBody) SetRequestId(v string) *DescribeLogStoreExistStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLogStoreExistStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogStoreExistStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeLogStoreExistStatusResponse) SetStatusCode(v int32) *DescribeLogStoreExistStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogStoreExistStatusResponse) SetBody(v *DescribeLogStoreExistStatusResponseBody) *DescribeLogStoreExistStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1536715558000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xx
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	OpAction   *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1534123558000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityObject(v string) *DescribeOpEntitiesRequest {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEntityType(v int32) *DescribeOpEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOpAction(v int32) *DescribeOpEntitiesRequest {
	s.OpAction = &v
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

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotal(v int64) *DescribeOpEntitiesResponseBody {
	s.Total = &v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	// example:
	//
	// 2.2.2.2
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 1536715558000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 123
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// example:
	//
	// 1
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// example:
	//
	// {"newEntity":{"elasticBandwidth":30},"oldEntity":{"elasticBandwidth":200}}
	OpDesc *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribeSimpleDomainsRequest struct {
	// example:
	//
	// ddoscoo-cn-XXXXXX
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSimpleDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsRequest) SetInstanceIds(v []*string) *DescribeSimpleDomainsRequest {
	s.InstanceIds = v
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

func (s *DescribeSimpleDomainsRequest) SetSourceIp(v string) *DescribeSimpleDomainsRequest {
	s.SourceIp = &v
	return s
}

type DescribeSimpleDomainsResponseBody struct {
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSimpleDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimpleDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponseBody) SetDomainList(v []*string) *DescribeSimpleDomainsResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeSimpleDomainsResponseBody) SetRequestId(v string) *DescribeSimpleDomainsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSimpleDomainsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimpleDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSimpleDomainsResponse) SetStatusCode(v int32) *DescribeSimpleDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimpleDomainsResponse) SetBody(v *DescribeSimpleDomainsResponseBody) *DescribeSimpleDomainsResponse {
	s.Body = v
	return s
}

type DescribeSlsAuthStatusRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) SetLang(v string) *DescribeSlsAuthStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetSourceIp(v string) *DescribeSlsAuthStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeSlsAuthStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SlsAuthStatus *bool `json:"SlsAuthStatus,omitempty" xml:"SlsAuthStatus,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSlsAuthStatusResponse) SetStatusCode(v int32) *DescribeSlsAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse {
	s.Body = v
	return s
}

type DescribeSlsEmptyCountRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSlsEmptyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountRequest) SetLang(v string) *DescribeSlsEmptyCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetResourceGroupId(v string) *DescribeSlsEmptyCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsEmptyCountRequest) SetSourceIp(v string) *DescribeSlsEmptyCountRequest {
	s.SourceIp = &v
	return s
}

type DescribeSlsEmptyCountResponseBody struct {
	// example:
	//
	// 0
	AvailableCount *int32 `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsEmptyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsEmptyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponseBody) SetAvailableCount(v int32) *DescribeSlsEmptyCountResponseBody {
	s.AvailableCount = &v
	return s
}

func (s *DescribeSlsEmptyCountResponseBody) SetRequestId(v string) *DescribeSlsEmptyCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSlsEmptyCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsEmptyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSlsEmptyCountResponse) SetStatusCode(v int32) *DescribeSlsEmptyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsEmptyCountResponse) SetBody(v *DescribeSlsEmptyCountResponseBody) *DescribeSlsEmptyCountResponse {
	s.Body = v
	return s
}

type DescribeSlsLogstoreInfoRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSlsLogstoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoRequest) SetLang(v string) *DescribeSlsLogstoreInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetResourceGroupId(v string) *DescribeSlsLogstoreInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoRequest) SetSourceIp(v string) *DescribeSlsLogstoreInfoRequest {
	s.SourceIp = &v
	return s
}

type DescribeSlsLogstoreInfoResponseBody struct {
	// example:
	//
	// ddoscoo-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// example:
	//
	// ddoscoo-project-xxxx-cn-hangzhou
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// example:
	//
	// 5497558138880
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ttl       *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeSlsLogstoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsLogstoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetLogStore(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetProject(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetQuota(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetRequestId(v string) *DescribeSlsLogstoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetTtl(v int32) *DescribeSlsLogstoreInfoResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponseBody) SetUsed(v int64) *DescribeSlsLogstoreInfoResponseBody {
	s.Used = &v
	return s
}

type DescribeSlsLogstoreInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsLogstoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSlsLogstoreInfoResponse) SetStatusCode(v int32) *DescribeSlsLogstoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsLogstoreInfoResponse) SetBody(v *DescribeSlsLogstoreInfoResponseBody) *DescribeSlsLogstoreInfoResponse {
	s.Body = v
	return s
}

type DescribeSlsOpenStatusRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusRequest) SetLang(v string) *DescribeSlsOpenStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetResourceGroupId(v string) *DescribeSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSlsOpenStatusRequest) SetSourceIp(v string) *DescribeSlsOpenStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeSlsOpenStatusResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SlsOpenStatus *bool `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeSlsOpenStatusResponse) SetStatusCode(v int32) *DescribeSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetBody(v *DescribeSlsOpenStatusResponseBody) *DescribeSlsOpenStatusResponse {
	s.Body = v
	return s
}

type DescribleCertListRequest struct {
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribleCertListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListRequest) GoString() string {
	return s.String()
}

func (s *DescribleCertListRequest) SetDomain(v string) *DescribleCertListRequest {
	s.Domain = &v
	return s
}

func (s *DescribleCertListRequest) SetResourceGroupId(v string) *DescribleCertListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleCertListRequest) SetSourceIp(v string) *DescribleCertListRequest {
	s.SourceIp = &v
	return s
}

type DescribleCertListResponseBody struct {
	CertList []*DescribleCertListResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// www.aliyun.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// example:
	//
	// false
	DomainRelated *bool `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty"`
	// example:
	//
	// 2020-09-23
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// DigiCert Inc
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// testCertName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2019-09-24
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribleCertListResponseBodyCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribleCertListResponseBodyCertList) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseBodyCertList) SetCertIdentifier(v string) *DescribleCertListResponseBodyCertList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetCommon(v string) *DescribleCertListResponseBodyCertList {
	s.Common = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetDomainRelated(v bool) *DescribleCertListResponseBodyCertList {
	s.DomainRelated = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetEndDate(v string) *DescribleCertListResponseBodyCertList {
	s.EndDate = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetId(v int32) *DescribleCertListResponseBodyCertList {
	s.Id = &v
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

func (s *DescribleCertListResponseBodyCertList) SetStartDate(v string) *DescribleCertListResponseBodyCertList {
	s.StartDate = &v
	return s
}

type DescribleCertListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribleCertListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribleCertListResponse) SetStatusCode(v int32) *DescribleCertListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribleCertListResponse) SetBody(v *DescribleCertListResponseBody) *DescribleCertListResponse {
	s.Body = v
	return s
}

type DescribleLayer7InstanceRelationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	DomainList []*string `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribleLayer7InstanceRelationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsRequest) SetDomainList(v []*string) *DescribleLayer7InstanceRelationsRequest {
	s.DomainList = v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetResourceGroupId(v string) *DescribleLayer7InstanceRelationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsRequest) SetSourceIp(v string) *DescribleLayer7InstanceRelationsRequest {
	s.SourceIp = &v
	return s
}

type DescribleLayer7InstanceRelationsResponseBody struct {
	Layer7InstanceRelations []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations `json:"Layer7InstanceRelations,omitempty" xml:"Layer7InstanceRelations,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribleLayer7InstanceRelationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribleLayer7InstanceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetLayer7InstanceRelations(v []*DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations) *DescribleLayer7InstanceRelationsResponseBody {
	s.Layer7InstanceRelations = v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBody) SetRequestId(v string) *DescribleLayer7InstanceRelationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelations struct {
	// example:
	//
	// www.aliyun.com
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
	EipList []*string `json:"EipList,omitempty" xml:"EipList,omitempty" type:"Repeated"`
	// example:
	//
	// default
	FunctionVersion *string `json:"FunctionVersion,omitempty" xml:"FunctionVersion,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpMode     *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	IpVersion  *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
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

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetIpMode(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.IpMode = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails) SetIpVersion(v string) *DescribleLayer7InstanceRelationsResponseBodyLayer7InstanceRelationsInstanceDetails {
	s.IpVersion = &v
	return s
}

type DescribleLayer7InstanceRelationsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribleLayer7InstanceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribleLayer7InstanceRelationsResponse) SetStatusCode(v int32) *DescribleLayer7InstanceRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribleLayer7InstanceRelationsResponse) SetBody(v *DescribleLayer7InstanceRelationsResponseBody) *DescribleLayer7InstanceRelationsResponse {
	s.Body = v
	return s
}

type DisableLayer7CCRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DisableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRequest) SetDomain(v string) *DisableLayer7CCRequest {
	s.Domain = &v
	return s
}

func (s *DisableLayer7CCRequest) SetResourceGroupId(v string) *DisableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRequest) SetSourceIp(v string) *DisableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

type DisableLayer7CCResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableLayer7CCResponse) SetStatusCode(v int32) *DisableLayer7CCResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLayer7CCResponse) SetBody(v *DisableLayer7CCResponseBody) *DisableLayer7CCResponse {
	s.Body = v
	return s
}

type DisableLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DisableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleRequest) SetDomain(v string) *DisableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetResourceGroupId(v string) *DisableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableLayer7CCRuleRequest) SetSourceIp(v string) *DisableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

type DisableLayer7CCRuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableLayer7CCRuleResponse) SetStatusCode(v int32) *DisableLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLayer7CCRuleResponse) SetBody(v *DisableLayer7CCRuleResponseBody) *DisableLayer7CCRuleResponse {
	s.Body = v
	return s
}

type EmptySlsLogstoreRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s EmptySlsLogstoreRequest) String() string {
	return tea.Prettify(s)
}

func (s EmptySlsLogstoreRequest) GoString() string {
	return s.String()
}

func (s *EmptySlsLogstoreRequest) SetLang(v string) *EmptySlsLogstoreRequest {
	s.Lang = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetResourceGroupId(v string) *EmptySlsLogstoreRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EmptySlsLogstoreRequest) SetSourceIp(v string) *EmptySlsLogstoreRequest {
	s.SourceIp = &v
	return s
}

type EmptySlsLogstoreResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EmptySlsLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EmptySlsLogstoreResponse) SetStatusCode(v int32) *EmptySlsLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *EmptySlsLogstoreResponse) SetBody(v *EmptySlsLogstoreResponseBody) *EmptySlsLogstoreResponse {
	s.Body = v
	return s
}

type EnableLayer7CCRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s EnableLayer7CCRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRequest) SetDomain(v string) *EnableLayer7CCRequest {
	s.Domain = &v
	return s
}

func (s *EnableLayer7CCRequest) SetResourceGroupId(v string) *EnableLayer7CCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRequest) SetSourceIp(v string) *EnableLayer7CCRequest {
	s.SourceIp = &v
	return s
}

type EnableLayer7CCResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableLayer7CCResponse) SetStatusCode(v int32) *EnableLayer7CCResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLayer7CCResponse) SetBody(v *EnableLayer7CCResponseBody) *EnableLayer7CCResponse {
	s.Body = v
	return s
}

type EnableLayer7CCRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s EnableLayer7CCRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableLayer7CCRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableLayer7CCRuleRequest) SetDomain(v string) *EnableLayer7CCRuleRequest {
	s.Domain = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetResourceGroupId(v string) *EnableLayer7CCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *EnableLayer7CCRuleRequest) SetSourceIp(v string) *EnableLayer7CCRuleRequest {
	s.SourceIp = &v
	return s
}

type EnableLayer7CCRuleResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableLayer7CCRuleResponse) SetStatusCode(v int32) *EnableLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableLayer7CCRuleResponse) SetBody(v *EnableLayer7CCRuleResponseBody) *EnableLayer7CCRuleResponse {
	s.Body = v
	return s
}

type ListAsyncTaskRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskRequest) SetLang(v string) *ListAsyncTaskRequest {
	s.Lang = &v
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

func (s *ListAsyncTaskRequest) SetResourceGroupId(v string) *ListAsyncTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAsyncTaskRequest) SetSourceIp(v string) *ListAsyncTaskRequest {
	s.SourceIp = &v
	return s
}

type ListAsyncTaskResponseBody struct {
	AsyncTasks []*ListAsyncTaskResponseBodyAsyncTasks `json:"AsyncTasks,omitempty" xml:"AsyncTasks,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTaskResponseBody) SetAsyncTasks(v []*ListAsyncTaskResponseBodyAsyncTasks) *ListAsyncTaskResponseBody {
	s.AsyncTasks = v
	return s
}

func (s *ListAsyncTaskResponseBody) SetRequestId(v string) *ListAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTaskResponseBody) SetTotal(v int32) *ListAsyncTaskResponseBody {
	s.Total = &v
	return s
}

type ListAsyncTaskResponseBodyAsyncTasks struct {
	// example:
	//
	// 1533866201000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1533866201000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {"instanceId": "ddoscoo-1234-qrq2134"}
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// {"instanceId": "ddoscoo-1234-qrq2134", "url": "https://oss.xxx.xxx"}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetStartTime(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.StartTime = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskId(v int64) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskParams(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskParams = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskResult(v string) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskStatus(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTaskResponseBodyAsyncTasks) SetTaskType(v int32) *ListAsyncTaskResponseBodyAsyncTasks {
	s.TaskType = &v
	return s
}

type ListAsyncTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListAsyncTaskResponse) SetStatusCode(v int32) *ListAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTaskResponse) SetBody(v *ListAsyncTaskResponseBody) *ListAsyncTaskResponse {
	s.Body = v
	return s
}

type ListLayer7CustomPortsRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListLayer7CustomPortsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsRequest) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsRequest) SetLang(v string) *ListLayer7CustomPortsRequest {
	s.Lang = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetResourceGroupId(v string) *ListLayer7CustomPortsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLayer7CustomPortsRequest) SetSourceIp(v string) *ListLayer7CustomPortsRequest {
	s.SourceIp = &v
	return s
}

type ListLayer7CustomPortsResponseBody struct {
	Layer7CustomPorts []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts `json:"Layer7CustomPorts,omitempty" xml:"Layer7CustomPorts,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLayer7CustomPortsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBody) SetLayer7CustomPorts(v []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts) *ListLayer7CustomPortsResponseBody {
	s.Layer7CustomPorts = v
	return s
}

func (s *ListLayer7CustomPortsResponseBody) SetRequestId(v string) *ListLayer7CustomPortsResponseBody {
	s.RequestId = &v
	return s
}

type ListLayer7CustomPortsResponseBodyLayer7CustomPorts struct {
	Flag       *string   `json:"Flag,omitempty" xml:"Flag,omitempty"`
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// example:
	//
	// https
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) String() string {
	return tea.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetFlag(v string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.Flag = &v
	return s
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLayer7CustomPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListLayer7CustomPortsResponse) SetStatusCode(v int32) *ListLayer7CustomPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLayer7CustomPortsResponse) SetBody(v *ListLayer7CustomPortsResponseBody) *ListLayer7CustomPortsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
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

type ListTagKeysResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 97935DF1-0289-4AA2-9DD1-72377838B16B
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
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
	// example:
	//
	// 1
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// example:
	//
	// a
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCffomr
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ddoscoo-cn-o4017n9q9004
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
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
	// example:
	//
	// testKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testValue1
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
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCffomr
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// C3F7E6AE-43B2-4730-B6A3-FD17552B8F65
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
	// example:
	//
	// ddoscoo-cn-o4017n9q9004
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// testKey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testValue1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListValueAddedRequest struct {
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ListValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ListValueAddedRequest) SetResourceGroupId(v string) *ListValueAddedRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListValueAddedRequest) SetSourceIp(v string) *ListValueAddedRequest {
	s.SourceIp = &v
	return s
}

type ListValueAddedResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	// example:
	//
	// 1580918400000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1575527305000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// ddos_fl_pre-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5497558138880
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 1
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StoreRegion *string `json:"StoreRegion,omitempty" xml:"StoreRegion,omitempty"`
}

func (s ListValueAddedResponseBodyValueAddedList) String() string {
	return tea.Prettify(s)
}

func (s ListValueAddedResponseBodyValueAddedList) GoString() string {
	return s.String()
}

func (s *ListValueAddedResponseBodyValueAddedList) SetExpireTime(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.ExpireTime = &v
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

func (s *ListValueAddedResponseBodyValueAddedList) SetLogSize(v int64) *ListValueAddedResponseBodyValueAddedList {
	s.LogSize = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetStatus(v int32) *ListValueAddedResponseBodyValueAddedList {
	s.Status = &v
	return s
}

func (s *ListValueAddedResponseBodyValueAddedList) SetStoreRegion(v string) *ListValueAddedResponseBodyValueAddedList {
	s.StoreRegion = &v
	return s
}

type ListValueAddedResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListValueAddedResponse) SetStatusCode(v int32) *ListValueAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListValueAddedResponse) SetBody(v *ListValueAddedResponseBody) *ListValueAddedResponse {
	s.Body = v
	return s
}

type ModifyElasticBandWidthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30
	ElasticBandwidth *int32 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyElasticBandWidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandWidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandWidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetInstanceId(v string) *ModifyElasticBandWidthRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticBandWidthRequest) SetSourceIp(v string) *ModifyElasticBandWidthRequest {
	s.SourceIp = &v
	return s
}

type ModifyElasticBandWidthResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticBandWidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyElasticBandWidthResponse) SetStatusCode(v int32) *ModifyElasticBandWidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticBandWidthResponse) SetBody(v *ModifyElasticBandWidthResponseBody) *ModifyElasticBandWidthResponse {
	s.Body = v
	return s
}

type ModifyFullLogTtlRequest struct {
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) SetLang(v string) *ModifyFullLogTtlRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetSourceIp(v string) *ModifyFullLogTtlRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int32) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

type ModifyFullLogTtlResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFullLogTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyFullLogTtlResponse) SetStatusCode(v int32) *ModifyFullLogTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFullLogTtlResponse) SetBody(v *ModifyFullLogTtlResponseBody) *ModifyFullLogTtlResponse {
	s.Body = v
	return s
}

type ModifyInstanceRemarkRequest struct {
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyInstanceRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkRequest) SetInstanceId(v string) *ModifyInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetRemark(v string) *ModifyInstanceRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyInstanceRemarkRequest) SetSourceIp(v string) *ModifyInstanceRemarkRequest {
	s.SourceIp = &v
	return s
}

type ModifyInstanceRemarkResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyInstanceRemarkResponse) SetStatusCode(v int32) *ModifyInstanceRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceRemarkResponse) SetBody(v *ModifyInstanceRemarkResponseBody) *ModifyInstanceRemarkResponse {
	s.Body = v
	return s
}

type OpenDomainSlsConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s OpenDomainSlsConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigRequest) SetDomain(v string) *OpenDomainSlsConfigRequest {
	s.Domain = &v
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

func (s *OpenDomainSlsConfigRequest) SetSourceIp(v string) *OpenDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

type OpenDomainSlsConfigResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDomainSlsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *OpenDomainSlsConfigResponse) SetStatusCode(v int32) *OpenDomainSlsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDomainSlsConfigResponse) SetBody(v *OpenDomainSlsConfigResponseBody) *OpenDomainSlsConfigResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetSourceIp(v string) *ReleaseInstanceRequest {
	s.SourceIp = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ReleaseInstanceResponse) SetStatusCode(v int32) *ReleaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type ReleaseValueAddedRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ddos_fl_pre-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ReleaseValueAddedRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseValueAddedRequest) GoString() string {
	return s.String()
}

func (s *ReleaseValueAddedRequest) SetInstanceId(v string) *ReleaseValueAddedRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseValueAddedRequest) SetSourceIp(v string) *ReleaseValueAddedRequest {
	s.SourceIp = &v
	return s
}

type ReleaseValueAddedResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseValueAddedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ReleaseValueAddedResponse) SetStatusCode(v int32) *ReleaseValueAddedResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseValueAddedResponse) SetBody(v *ReleaseValueAddedResponseBody) *ReleaseValueAddedResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-v0h1fmwbc024
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// example:
	//
	// testKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// testValue1
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
	// example:
	//
	// 7078CD1E-F609-47A4-9C39-B288CC27C686
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-v0h1fmwbc024
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// testKey1
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
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

type UntagResourcesResponseBody struct {
	// example:
	//
	// F2D86AED-BA27-4584-BADC-B43BDA7EEBCA
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// @param request - AddLayer7CCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLayer7CCRuleResponse
func (client *Client) AddLayer7CCRuleWithOptions(request *AddLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *AddLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Act)) {
		query["Act"] = request.Act
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		query["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddLayer7CCRule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddLayer7CCRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddLayer7CCRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - AddLayer7CCRuleRequest
//
// @return AddLayer7CCRuleResponse
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

// @param request - CloseDomainSlsConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseDomainSlsConfigResponse
func (client *Client) CloseDomainSlsConfigWithOptions(request *CloseDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *CloseDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseDomainSlsConfig"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CloseDomainSlsConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CloseDomainSlsConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CloseDomainSlsConfigRequest
//
// @return CloseDomainSlsConfigResponse
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

// @param request - ConfigHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigHealthCheckResponse
func (client *Client) ConfigHealthCheckWithOptions(request *ConfigHealthCheckRequest, runtime *util.RuntimeOptions) (_result *ConfigHealthCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForwardProtocol)) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.FrontendPort)) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheck)) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigHealthCheck"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigHealthCheckResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigHealthCheckResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigHealthCheckRequest
//
// @return ConfigHealthCheckResponse
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

// @param request - ConfigLayer4RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RuleResponse
func (client *Client) ConfigLayer4RuleWithOptions(request *ConfigLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyEnable)) {
		query["ProxyEnable"] = request.ProxyEnable
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer4Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer4RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer4RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer4RuleRequest
//
// @return ConfigLayer4RuleResponse
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

// @param request - ConfigLayer4RuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RuleAttributeResponse
func (client *Client) ConfigLayer4RuleAttributeWithOptions(request *ConfigLayer4RuleAttributeRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer4RuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardProtocol)) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.FrontendPort)) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer4RuleAttribute"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer4RuleAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer4RuleAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer4RuleAttributeRequest
//
// @return ConfigLayer4RuleAttributeResponse
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

// @param request - ConfigLayer7BlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer7BlackWhiteListResponse
func (client *Client) ConfigLayer7BlackWhiteListWithOptions(request *ConfigLayer7BlackWhiteListRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7BlackWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackList)) {
		query["BlackList"] = request.BlackList
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteList)) {
		query["WhiteList"] = request.WhiteList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer7BlackWhiteList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer7BlackWhiteListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer7BlackWhiteListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer7BlackWhiteListRequest
//
// @return ConfigLayer7BlackWhiteListResponse
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

// @param request - ConfigLayer7CCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer7CCRuleResponse
func (client *Client) ConfigLayer7CCRuleWithOptions(request *ConfigLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Act)) {
		query["Act"] = request.Act
	}

	if !tea.BoolValue(util.IsUnset(request.Count)) {
		query["Count"] = request.Count
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		query["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer7CCRule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer7CCRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer7CCRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer7CCRuleRequest
//
// @return ConfigLayer7CCRuleResponse
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

// @param request - ConfigLayer7CCTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer7CCTemplateResponse
func (client *Client) ConfigLayer7CCTemplateWithOptions(request *ConfigLayer7CCTemplateRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CCTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Template)) {
		query["Template"] = request.Template
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer7CCTemplate"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer7CCTemplateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer7CCTemplateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer7CCTemplateRequest
//
// @return ConfigLayer7CCTemplateResponse
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

// @param request - ConfigLayer7CertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer7CertResponse
func (client *Client) ConfigLayer7CertWithOptions(request *ConfigLayer7CertRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7CertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cert)) {
		query["Cert"] = request.Cert
	}

	if !tea.BoolValue(util.IsUnset(request.CertId)) {
		query["CertId"] = request.CertId
	}

	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.CertRegion)) {
		query["CertRegion"] = request.CertRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer7Cert"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer7CertResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer7CertResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer7CertRequest
//
// @return ConfigLayer7CertResponse
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

// @param request - ConfigLayer7RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer7RuleResponse
func (client *Client) ConfigLayer7RuleWithOptions(request *ConfigLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *ConfigLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyTypeList)) {
		query["ProxyTypeList"] = request.ProxyTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyTypes)) {
		query["ProxyTypes"] = request.ProxyTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RealServers)) {
		query["RealServers"] = request.RealServers
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RsType)) {
		query["RsType"] = request.RsType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigLayer7Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConfigLayer7RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConfigLayer7RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ConfigLayer7RuleRequest
//
// @return ConfigLayer7RuleResponse
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

// @param request - CreateAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTaskResponse
func (client *Client) CreateAsyncTaskWithOptions(request *CreateAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskParams)) {
		query["TaskParams"] = request.TaskParams
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAsyncTask"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAsyncTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAsyncTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateAsyncTaskRequest
//
// @return CreateAsyncTaskResponse
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

// @param request - CreateLayer4RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayer4RuleResponse
func (client *Client) CreateLayer4RuleWithOptions(request *CreateLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyEnable)) {
		query["ProxyEnable"] = request.ProxyEnable
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayer4Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLayer4RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLayer4RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateLayer4RuleRequest
//
// @return CreateLayer4RuleResponse
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

// @param request - CreateLayer7RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLayer7RuleResponse
func (client *Client) CreateLayer7RuleWithOptions(request *CreateLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *CreateLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RsType)) {
		query["RsType"] = request.RsType
	}

	if !tea.BoolValue(util.IsUnset(request.Rules)) {
		query["Rules"] = request.Rules
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLayer7Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLayer7RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLayer7RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateLayer7RuleRequest
//
// @return CreateLayer7RuleResponse
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

// @param request - DeleteAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsyncTaskResponse
func (client *Client) DeleteAsyncTaskWithOptions(request *DeleteAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAsyncTask"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAsyncTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAsyncTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAsyncTaskRequest
//
// @return DeleteAsyncTaskResponse
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

// @param request - DeleteLayer4RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayer4RuleResponse
func (client *Client) DeleteLayer4RuleWithOptions(request *DeleteLayer4RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer4RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayer4Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLayer4RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLayer4RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteLayer4RuleRequest
//
// @return DeleteLayer4RuleResponse
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

// @param request - DeleteLayer7CCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayer7CCRuleResponse
func (client *Client) DeleteLayer7CCRuleWithOptions(request *DeleteLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayer7CCRule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLayer7CCRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLayer7CCRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteLayer7CCRuleRequest
//
// @return DeleteLayer7CCRuleResponse
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

// @param request - DeleteLayer7RuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLayer7RuleResponse
func (client *Client) DeleteLayer7RuleWithOptions(request *DeleteLayer7RuleRequest, runtime *util.RuntimeOptions) (_result *DeleteLayer7RuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLayer7Rule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLayer7RuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLayer7RuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteLayer7RuleRequest
//
// @return DeleteLayer7RuleResponse
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

// @param request - DescribeBackSourceCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackSourceCidrResponse
func (client *Client) DescribeBackSourceCidrWithOptions(request *DescribeBackSourceCidrRequest, runtime *util.RuntimeOptions) (_result *DescribeBackSourceCidrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["Line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackSourceCidr"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBackSourceCidrResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBackSourceCidrResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeBackSourceCidrRequest
//
// @return DescribeBackSourceCidrResponse
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

// Summary:
//
// rosetta迁移
//
// @param request - DescribeBatchSlsDispatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchSlsDispatchStatusResponse
func (client *Client) DescribeBatchSlsDispatchStatusWithOptions(request *DescribeBatchSlsDispatchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBatchSlsDispatchStatus"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeBatchSlsDispatchStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeBatchSlsDispatchStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// rosetta迁移
//
// @param request - DescribeBatchSlsDispatchStatusRequest
//
// @return DescribeBatchSlsDispatchStatusResponse
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

// @param request - DescribeDDoSEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSEventsResponse
func (client *Client) DescribeDDoSEventsWithOptions(request *DescribeDDoSEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Eip)) {
		query["Eip"] = request.Eip
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDDoSEvents"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDDoSEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDDoSEventsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDDoSEventsRequest
//
// @return DescribeDDoSEventsResponse
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

// @param request - DescribeDDoSTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSTrafficResponse
func (client *Client) DescribeDDoSTrafficWithOptions(request *DescribeDDoSTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeDDoSTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Eip)) {
		query["Eip"] = request.Eip
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDDoSTraffic"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDDoSTrafficResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDDoSTrafficResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDDoSTrafficRequest
//
// @return DescribeDDoSTrafficResponse
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

// @param request - DescribeDefenseCountStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseCountStatisticsResponse
func (client *Client) DescribeDefenseCountStatisticsWithOptions(request *DescribeDefenseCountStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefenseCountStatistics"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDefenseCountStatisticsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDefenseCountStatisticsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDefenseCountStatisticsRequest
//
// @return DescribeDefenseCountStatisticsResponse
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

// @param request - DescribeDomainAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAccessModeResponse
func (client *Client) DescribeDomainAccessModeWithOptions(request *DescribeDomainAccessModeRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAccessModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainList)) {
		query["DomainList"] = request.DomainList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainAccessMode"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainAccessModeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainAccessModeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainAccessModeRequest
//
// @return DescribeDomainAccessModeResponse
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

// @param request - DescribeDomainAttackEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAttackEventListResponse
func (client *Client) DescribeDomainAttackEventListWithOptions(request *DescribeDomainAttackEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainAttackEventList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainAttackEventListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainAttackEventListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainAttackEventListRequest
//
// @return DescribeDomainAttackEventListResponse
func (client *Client) DescribeDomainAttackEventList(request *DescribeDomainAttackEventListRequest) (_result *DescribeDomainAttackEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackEventListResponse{}
	_body, _err := client.DescribeDomainAttackEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainAttackMaxQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAttackMaxQpsResponse
func (client *Client) DescribeDomainAttackMaxQpsWithOptions(request *DescribeDomainAttackMaxQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainAttackMaxQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainAttackMaxQps"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainAttackMaxQpsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainAttackMaxQpsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainAttackMaxQpsRequest
//
// @return DescribeDomainAttackMaxQpsResponse
func (client *Client) DescribeDomainAttackMaxQps(request *DescribeDomainAttackMaxQpsRequest) (_result *DescribeDomainAttackMaxQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainAttackMaxQpsResponse{}
	_body, _err := client.DescribeDomainAttackMaxQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainOverviewResponse
func (client *Client) DescribeDomainOverviewWithOptions(request *DescribeDomainOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainOverview"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainOverviewResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainOverviewResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainOverviewRequest
//
// @return DescribeDomainOverviewResponse
func (client *Client) DescribeDomainOverview(request *DescribeDomainOverviewRequest) (_result *DescribeDomainOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainOverviewResponse{}
	_body, _err := client.DescribeDomainOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainQpsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainQpsListResponse
func (client *Client) DescribeDomainQpsListWithOptions(request *DescribeDomainQpsListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainQpsList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainQpsListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainQpsListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainQpsListRequest
//
// @return DescribeDomainQpsListResponse
func (client *Client) DescribeDomainQpsList(request *DescribeDomainQpsListRequest) (_result *DescribeDomainQpsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainQpsListResponse{}
	_body, _err := client.DescribeDomainQpsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainQpsWithCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainQpsWithCacheResponse
func (client *Client) DescribeDomainQpsWithCacheWithOptions(request *DescribeDomainQpsWithCacheRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainQpsWithCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainQpsWithCache"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainQpsWithCacheResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainQpsWithCacheResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainQpsWithCacheRequest
//
// @return DescribeDomainQpsWithCacheResponse
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

// @param request - DescribeDomainSlsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSlsStatusResponse
func (client *Client) DescribeDomainSlsStatusWithOptions(request *DescribeDomainSlsStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSlsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainSlsStatus"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainSlsStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainSlsStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainSlsStatusRequest
//
// @return DescribeDomainSlsStatusResponse
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

// Summary:
//
// 查询网站业务的响应状态码统计信息
//
// @param request - DescribeDomainStatusCodeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatusCodeListResponse
func (client *Client) DescribeDomainStatusCodeListWithOptions(request *DescribeDomainStatusCodeListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainStatusCodeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainStatusCodeList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainStatusCodeListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainStatusCodeListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询网站业务的响应状态码统计信息
//
// @param request - DescribeDomainStatusCodeListRequest
//
// @return DescribeDomainStatusCodeListResponse
func (client *Client) DescribeDomainStatusCodeList(request *DescribeDomainStatusCodeListRequest) (_result *DescribeDomainStatusCodeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainStatusCodeListResponse{}
	_body, _err := client.DescribeDomainStatusCodeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDomainPattern)) {
		query["QueryDomainPattern"] = request.QueryDomainPattern
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomains"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDomainsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDomainsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeDomainsRequest
//
// @return DescribeDomainsResponse
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

// @param request - DescribeElasticBandwidthSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticBandwidthSpecResponse
func (client *Client) DescribeElasticBandwidthSpecWithOptions(request *DescribeElasticBandwidthSpecRequest, runtime *util.RuntimeOptions) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeElasticBandwidthSpec"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeElasticBandwidthSpecResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeElasticBandwidthSpecResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeElasticBandwidthSpecRequest
//
// @return DescribeElasticBandwidthSpecResponse
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

// @param request - DescribeHealthCheckListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthCheckListResponse
func (client *Client) DescribeHealthCheckListWithOptions(request *DescribeHealthCheckListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHealthCheckList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeHealthCheckListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeHealthCheckListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeHealthCheckListRequest
//
// @return DescribeHealthCheckListResponse
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

// @param request - DescribeHealthCheckStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthCheckStatusListResponse
func (client *Client) DescribeHealthCheckStatusListWithOptions(request *DescribeHealthCheckStatusListRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckStatusListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHealthCheckStatusList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeHealthCheckStatusListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeHealthCheckStatusListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeHealthCheckStatusListRequest
//
// @return DescribeHealthCheckStatusListResponse
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

// @param request - DescribeInstanceDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDetailsResponse
func (client *Client) DescribeInstanceDetailsWithOptions(request *DescribeInstanceDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceDetails"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceDetailsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceDetailsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeInstanceDetailsRequest
//
// @return DescribeInstanceDetailsResponse
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

// @param request - DescribeInstanceSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecs"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceSpecsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceSpecsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeInstanceSpecsRequest
//
// @return DescribeInstanceSpecsResponse
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

// @param request - DescribeInstanceStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceStatisticsResponse
func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceStatistics"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstanceStatisticsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstanceStatisticsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeInstanceStatisticsRequest
//
// @return DescribeInstanceStatisticsResponse
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

// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["Edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireEndTime)) {
		query["ExpireEndTime"] = request.ExpireEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireStartTime)) {
		query["ExpireStartTime"] = request.ExpireStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
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

// @param request - DescribeIpTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpTrafficResponse
func (client *Client) DescribeIpTrafficWithOptions(request *DescribeIpTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeIpTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Eip)) {
		query["Eip"] = request.Eip
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.QueryProtocol)) {
		query["QueryProtocol"] = request.QueryProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpTraffic"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeIpTrafficResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeIpTrafficResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeIpTrafficRequest
//
// @return DescribeIpTrafficResponse
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

// @param request - DescribeLayer4RuleAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLayer4RuleAttributesResponse
func (client *Client) DescribeLayer4RuleAttributesWithOptions(request *DescribeLayer4RuleAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RuleAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Listeners)) {
		query["Listeners"] = request.Listeners
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLayer4RuleAttributes"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeLayer4RuleAttributesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeLayer4RuleAttributesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeLayer4RuleAttributesRequest
//
// @return DescribeLayer4RuleAttributesResponse
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

// Summary:
//
// Call DescribeLayer4Rules to query the detailed configuration of port forwarding rules for DDoS protection instances.
//
// @param request - DescribeLayer4RulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLayer4RulesResponse
func (client *Client) DescribeLayer4RulesWithOptions(request *DescribeLayer4RulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer4RulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForwardProtocol)) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.FrontendPort)) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLayer4Rules"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeLayer4RulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeLayer4RulesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Call DescribeLayer4Rules to query the detailed configuration of port forwarding rules for DDoS protection instances.
//
// @param request - DescribeLayer4RulesRequest
//
// @return DescribeLayer4RulesResponse
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

// @param request - DescribeLayer7CCRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLayer7CCRulesResponse
func (client *Client) DescribeLayer7CCRulesWithOptions(request *DescribeLayer7CCRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeLayer7CCRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLayer7CCRules"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeLayer7CCRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeLayer7CCRulesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeLayer7CCRulesRequest
//
// @return DescribeLayer7CCRulesResponse
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

// @param request - DescribeLogStoreExistStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogStoreExistStatusResponse
func (client *Client) DescribeLogStoreExistStatusWithOptions(request *DescribeLogStoreExistStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogStoreExistStatus"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeLogStoreExistStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeLogStoreExistStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeLogStoreExistStatusRequest
//
// @return DescribeLogStoreExistStatusResponse
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

// @param request - DescribeOpEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EntityObject)) {
		query["EntityObject"] = request.EntityObject
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.OpAction)) {
		query["OpAction"] = request.OpAction
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpEntities"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeOpEntitiesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeOpEntitiesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeOpEntitiesRequest
//
// @return DescribeOpEntitiesResponse
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

// @param request - DescribeSimpleDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimpleDomainsResponse
func (client *Client) DescribeSimpleDomainsWithOptions(request *DescribeSimpleDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeSimpleDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSimpleDomains"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSimpleDomainsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSimpleDomainsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSimpleDomainsRequest
//
// @return DescribeSimpleDomainsResponse
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

// @param request - DescribeSlsAuthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsAuthStatusResponse
func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsAuthStatus"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlsAuthStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlsAuthStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSlsAuthStatusRequest
//
// @return DescribeSlsAuthStatusResponse
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

// @param request - DescribeSlsEmptyCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsEmptyCountResponse
func (client *Client) DescribeSlsEmptyCountWithOptions(request *DescribeSlsEmptyCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsEmptyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsEmptyCount"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlsEmptyCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlsEmptyCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSlsEmptyCountRequest
//
// @return DescribeSlsEmptyCountResponse
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

// @param request - DescribeSlsLogstoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsLogstoreInfoResponse
func (client *Client) DescribeSlsLogstoreInfoWithOptions(request *DescribeSlsLogstoreInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsLogstoreInfo"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlsLogstoreInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlsLogstoreInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSlsLogstoreInfoRequest
//
// @return DescribeSlsLogstoreInfoResponse
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

// @param request - DescribeSlsOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsOpenStatusResponse
func (client *Client) DescribeSlsOpenStatusWithOptions(request *DescribeSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlsOpenStatus"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSlsOpenStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSlsOpenStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribeSlsOpenStatusRequest
//
// @return DescribeSlsOpenStatusResponse
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

// @param request - DescribleCertListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribleCertListResponse
func (client *Client) DescribleCertListWithOptions(request *DescribleCertListRequest, runtime *util.RuntimeOptions) (_result *DescribleCertListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribleCertList"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribleCertListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribleCertListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribleCertListRequest
//
// @return DescribleCertListResponse
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

// @param request - DescribleLayer7InstanceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribleLayer7InstanceRelationsResponse
func (client *Client) DescribleLayer7InstanceRelationsWithOptions(request *DescribleLayer7InstanceRelationsRequest, runtime *util.RuntimeOptions) (_result *DescribleLayer7InstanceRelationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainList)) {
		query["DomainList"] = request.DomainList
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribleLayer7InstanceRelations"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribleLayer7InstanceRelationsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribleLayer7InstanceRelationsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DescribleLayer7InstanceRelationsRequest
//
// @return DescribleLayer7InstanceRelationsResponse
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

// @param request - DisableLayer7CCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableLayer7CCResponse
func (client *Client) DisableLayer7CCWithOptions(request *DisableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableLayer7CC"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableLayer7CCResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableLayer7CCResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DisableLayer7CCRequest
//
// @return DisableLayer7CCResponse
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

// @param request - DisableLayer7CCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableLayer7CCRuleResponse
func (client *Client) DisableLayer7CCRuleWithOptions(request *DisableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *DisableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableLayer7CCRule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableLayer7CCRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableLayer7CCRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DisableLayer7CCRuleRequest
//
// @return DisableLayer7CCRuleResponse
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

// @param request - EmptySlsLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmptySlsLogstoreResponse
func (client *Client) EmptySlsLogstoreWithOptions(request *EmptySlsLogstoreRequest, runtime *util.RuntimeOptions) (_result *EmptySlsLogstoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EmptySlsLogstore"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EmptySlsLogstoreResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EmptySlsLogstoreResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - EmptySlsLogstoreRequest
//
// @return EmptySlsLogstoreResponse
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

// @param request - EnableLayer7CCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableLayer7CCResponse
func (client *Client) EnableLayer7CCWithOptions(request *EnableLayer7CCRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLayer7CC"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableLayer7CCResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableLayer7CCResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - EnableLayer7CCRequest
//
// @return EnableLayer7CCResponse
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

// @param request - EnableLayer7CCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableLayer7CCRuleResponse
func (client *Client) EnableLayer7CCRuleWithOptions(request *EnableLayer7CCRuleRequest, runtime *util.RuntimeOptions) (_result *EnableLayer7CCRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableLayer7CCRule"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableLayer7CCRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableLayer7CCRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - EnableLayer7CCRuleRequest
//
// @return EnableLayer7CCRuleResponse
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

// @param request - ListAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTaskResponse
func (client *Client) ListAsyncTaskWithOptions(request *ListAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *ListAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsyncTask"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAsyncTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAsyncTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAsyncTaskRequest
//
// @return ListAsyncTaskResponse
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

// @param request - ListLayer7CustomPortsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLayer7CustomPortsResponse
func (client *Client) ListLayer7CustomPortsWithOptions(request *ListLayer7CustomPortsRequest, runtime *util.RuntimeOptions) (_result *ListLayer7CustomPortsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLayer7CustomPorts"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLayer7CustomPortsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLayer7CustomPortsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListLayer7CustomPortsRequest
//
// @return ListLayer7CustomPortsResponse
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

// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagKeysResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagKeysResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
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

// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// @param request - ListValueAddedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListValueAddedResponse
func (client *Client) ListValueAddedWithOptions(request *ListValueAddedRequest, runtime *util.RuntimeOptions) (_result *ListValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListValueAdded"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListValueAddedResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListValueAddedResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListValueAddedRequest
//
// @return ListValueAddedResponse
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

// @param request - ModifyElasticBandWidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticBandWidthResponse
func (client *Client) ModifyElasticBandWidthWithOptions(request *ModifyElasticBandWidthRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticBandWidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElasticBandwidth)) {
		query["ElasticBandwidth"] = request.ElasticBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyElasticBandWidth"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyElasticBandWidthResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyElasticBandWidthResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyElasticBandWidthRequest
//
// @return ModifyElasticBandWidthResponse
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

// @param request - ModifyFullLogTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFullLogTtlResponse
func (client *Client) ModifyFullLogTtlWithOptions(request *ModifyFullLogTtlRequest, runtime *util.RuntimeOptions) (_result *ModifyFullLogTtlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFullLogTtl"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyFullLogTtlResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyFullLogTtlResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyFullLogTtlRequest
//
// @return ModifyFullLogTtlResponse
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

// @param request - ModifyInstanceRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceRemarkResponse
func (client *Client) ModifyInstanceRemarkWithOptions(request *ModifyInstanceRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceRemark"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyInstanceRemarkResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyInstanceRemarkResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyInstanceRemarkRequest
//
// @return ModifyInstanceRemarkResponse
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

// @param request - OpenDomainSlsConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDomainSlsConfigResponse
func (client *Client) OpenDomainSlsConfigWithOptions(request *OpenDomainSlsConfigRequest, runtime *util.RuntimeOptions) (_result *OpenDomainSlsConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenDomainSlsConfig"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenDomainSlsConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenDomainSlsConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - OpenDomainSlsConfigRequest
//
// @return OpenDomainSlsConfigResponse
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

// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstance"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
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

// @param request - ReleaseValueAddedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseValueAddedResponse
func (client *Client) ReleaseValueAddedWithOptions(request *ReleaseValueAddedRequest, runtime *util.RuntimeOptions) (_result *ReleaseValueAddedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseValueAdded"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseValueAddedResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseValueAddedResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ReleaseValueAddedRequest
//
// @return ReleaseValueAddedResponse
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

// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-12-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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
