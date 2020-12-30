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

type ConfigSwitchPriorityRequest struct {
	SourceIp *string                              `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string                              `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Config   []*ConfigSwitchPriorityRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ConfigSwitchPriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigSwitchPriorityRequest) GoString() string {
	return s.String()
}

func (s *ConfigSwitchPriorityRequest) SetSourceIp(v string) *ConfigSwitchPriorityRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigSwitchPriorityRequest) SetLang(v string) *ConfigSwitchPriorityRequest {
	s.Lang = &v
	return s
}

func (s *ConfigSwitchPriorityRequest) SetDomain(v string) *ConfigSwitchPriorityRequest {
	s.Domain = &v
	return s
}

func (s *ConfigSwitchPriorityRequest) SetConfig(v []*ConfigSwitchPriorityRequestConfig) *ConfigSwitchPriorityRequest {
	s.Config = v
	return s
}

type ConfigSwitchPriorityRequestConfig struct {
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ConfigSwitchPriorityRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s ConfigSwitchPriorityRequestConfig) GoString() string {
	return s.String()
}

func (s *ConfigSwitchPriorityRequestConfig) SetIp(v string) *ConfigSwitchPriorityRequestConfig {
	s.Ip = &v
	return s
}

func (s *ConfigSwitchPriorityRequestConfig) SetPriority(v int32) *ConfigSwitchPriorityRequestConfig {
	s.Priority = &v
	return s
}

type ConfigSwitchPriorityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSwitchPriorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigSwitchPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSwitchPriorityResponseBody) SetRequestId(v string) *ConfigSwitchPriorityResponseBody {
	s.RequestId = &v
	return s
}

type ConfigSwitchPriorityResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigSwitchPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigSwitchPriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigSwitchPriorityResponse) GoString() string {
	return s.String()
}

func (s *ConfigSwitchPriorityResponse) SetHeaders(v map[string]*string) *ConfigSwitchPriorityResponse {
	s.Headers = v
	return s
}

func (s *ConfigSwitchPriorityResponse) SetBody(v *ConfigSwitchPriorityResponseBody) *ConfigSwitchPriorityResponse {
	s.Body = v
	return s
}

type CreateCcCustomedRuleRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MatchingRule *string `json:"MatchingRule,omitempty" xml:"MatchingRule,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	VisitCount   *int32  `json:"VisitCount,omitempty" xml:"VisitCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BlockingType *string `json:"BlockingType,omitempty" xml:"BlockingType,omitempty"`
	Interval     *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	BlockingTime *int32  `json:"BlockingTime,omitempty" xml:"BlockingTime,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateCcCustomedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCcCustomedRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCcCustomedRuleRequest) SetSourceIp(v string) *CreateCcCustomedRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetLang(v string) *CreateCcCustomedRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetMatchingRule(v string) *CreateCcCustomedRuleRequest {
	s.MatchingRule = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetDomain(v string) *CreateCcCustomedRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetVisitCount(v int32) *CreateCcCustomedRuleRequest {
	s.VisitCount = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetName(v string) *CreateCcCustomedRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetBlockingType(v string) *CreateCcCustomedRuleRequest {
	s.BlockingType = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetInterval(v int32) *CreateCcCustomedRuleRequest {
	s.Interval = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetBlockingTime(v int32) *CreateCcCustomedRuleRequest {
	s.BlockingTime = &v
	return s
}

func (s *CreateCcCustomedRuleRequest) SetUri(v string) *CreateCcCustomedRuleRequest {
	s.Uri = &v
	return s
}

type CreateCcCustomedRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCcCustomedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCcCustomedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCcCustomedRuleResponseBody) SetRequestId(v string) *CreateCcCustomedRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateCcCustomedRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCcCustomedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCcCustomedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCcCustomedRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCcCustomedRuleResponse) SetHeaders(v map[string]*string) *CreateCcCustomedRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCcCustomedRuleResponse) SetBody(v *CreateCcCustomedRuleResponseBody) *CreateCcCustomedRuleResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain     *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ip         *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	CcEnable   *bool     `json:"CcEnable,omitempty" xml:"CcEnable,omitempty"`
	RealServer []*string `json:"RealServer,omitempty" xml:"RealServer,omitempty" type:"Repeated"`
	ProxyType  []*string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty" type:"Repeated"`
	Ips        []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetSourceIp(v string) *CreateDomainRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDomainRequest) SetLang(v string) *CreateDomainRequest {
	s.Lang = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetIp(v string) *CreateDomainRequest {
	s.Ip = &v
	return s
}

func (s *CreateDomainRequest) SetType(v string) *CreateDomainRequest {
	s.Type = &v
	return s
}

func (s *CreateDomainRequest) SetCcEnable(v bool) *CreateDomainRequest {
	s.CcEnable = &v
	return s
}

func (s *CreateDomainRequest) SetRealServer(v []*string) *CreateDomainRequest {
	s.RealServer = v
	return s
}

func (s *CreateDomainRequest) SetProxyType(v []*string) *CreateDomainRequest {
	s.ProxyType = v
	return s
}

func (s *CreateDomainRequest) SetIps(v []*string) *CreateDomainRequest {
	s.Ips = v
	return s
}

type CreateDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreatePortRuleRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	FrontPort      *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	BackPort       *int32  `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	ProxyType      *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	RealServerList *string `json:"RealServerList,omitempty" xml:"RealServerList,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s CreatePortRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePortRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePortRuleRequest) SetSourceIp(v string) *CreatePortRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreatePortRuleRequest) SetLang(v string) *CreatePortRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreatePortRuleRequest) SetFrontPort(v int32) *CreatePortRuleRequest {
	s.FrontPort = &v
	return s
}

func (s *CreatePortRuleRequest) SetBackPort(v int32) *CreatePortRuleRequest {
	s.BackPort = &v
	return s
}

func (s *CreatePortRuleRequest) SetProxyType(v string) *CreatePortRuleRequest {
	s.ProxyType = &v
	return s
}

func (s *CreatePortRuleRequest) SetRealServerList(v string) *CreatePortRuleRequest {
	s.RealServerList = &v
	return s
}

func (s *CreatePortRuleRequest) SetIp(v string) *CreatePortRuleRequest {
	s.Ip = &v
	return s
}

type CreatePortRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePortRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortRuleResponseBody) SetRequestId(v string) *CreatePortRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreatePortRuleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePortRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePortRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePortRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatePortRuleResponse) SetHeaders(v map[string]*string) *CreatePortRuleResponse {
	s.Headers = v
	return s
}

func (s *CreatePortRuleResponse) SetBody(v *CreatePortRuleResponseBody) *CreatePortRuleResponse {
	s.Body = v
	return s
}

type CreateTransmitLineRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type        *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain      *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ips         []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s CreateTransmitLineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransmitLineRequest) GoString() string {
	return s.String()
}

func (s *CreateTransmitLineRequest) SetSourceIp(v string) *CreateTransmitLineRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateTransmitLineRequest) SetLang(v string) *CreateTransmitLineRequest {
	s.Lang = &v
	return s
}

func (s *CreateTransmitLineRequest) SetType(v string) *CreateTransmitLineRequest {
	s.Type = &v
	return s
}

func (s *CreateTransmitLineRequest) SetDomain(v string) *CreateTransmitLineRequest {
	s.Domain = &v
	return s
}

func (s *CreateTransmitLineRequest) SetIps(v []*string) *CreateTransmitLineRequest {
	s.Ips = v
	return s
}

func (s *CreateTransmitLineRequest) SetRealServers(v []*string) *CreateTransmitLineRequest {
	s.RealServers = v
	return s
}

type CreateTransmitLineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransmitLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransmitLineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransmitLineResponseBody) SetRequestId(v string) *CreateTransmitLineResponseBody {
	s.RequestId = &v
	return s
}

type CreateTransmitLineResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransmitLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransmitLineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransmitLineResponse) GoString() string {
	return s.String()
}

func (s *CreateTransmitLineResponse) SetHeaders(v map[string]*string) *CreateTransmitLineResponse {
	s.Headers = v
	return s
}

func (s *CreateTransmitLineResponse) SetBody(v *CreateTransmitLineResponseBody) *CreateTransmitLineResponse {
	s.Body = v
	return s
}

type DeleteCcCustomedRuleRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteCcCustomedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcCustomedRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCcCustomedRuleRequest) SetSourceIp(v string) *DeleteCcCustomedRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteCcCustomedRuleRequest) SetLang(v string) *DeleteCcCustomedRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteCcCustomedRuleRequest) SetDomain(v string) *DeleteCcCustomedRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteCcCustomedRuleRequest) SetName(v string) *DeleteCcCustomedRuleRequest {
	s.Name = &v
	return s
}

type DeleteCcCustomedRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCcCustomedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcCustomedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCcCustomedRuleResponseBody) SetRequestId(v string) *DeleteCcCustomedRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCcCustomedRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCcCustomedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCcCustomedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCcCustomedRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCcCustomedRuleResponse) SetHeaders(v map[string]*string) *DeleteCcCustomedRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCcCustomedRuleResponse) SetBody(v *DeleteCcCustomedRuleResponseBody) *DeleteCcCustomedRuleResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetSourceIp(v string) *DeleteDomainRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteDomainRequest) SetLang(v string) *DeleteDomainRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainRequest) SetDomain(v string) *DeleteDomainRequest {
	s.Domain = &v
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeletePortRuleRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	FrontPort *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DeletePortRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePortRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePortRuleRequest) SetSourceIp(v string) *DeletePortRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeletePortRuleRequest) SetLang(v string) *DeletePortRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeletePortRuleRequest) SetFrontPort(v int32) *DeletePortRuleRequest {
	s.FrontPort = &v
	return s
}

func (s *DeletePortRuleRequest) SetIp(v string) *DeletePortRuleRequest {
	s.Ip = &v
	return s
}

type DeletePortRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePortRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePortRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortRuleResponseBody) SetRequestId(v string) *DeletePortRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeletePortRuleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePortRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePortRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePortRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePortRuleResponse) SetHeaders(v map[string]*string) *DeletePortRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePortRuleResponse) SetBody(v *DeletePortRuleResponseBody) *DeletePortRuleResponse {
	s.Body = v
	return s
}

type DeleteTransmitLineRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Line     *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteTransmitLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransmitLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransmitLineRequest) SetSourceIp(v string) *DeleteTransmitLineRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteTransmitLineRequest) SetLang(v string) *DeleteTransmitLineRequest {
	s.Lang = &v
	return s
}

func (s *DeleteTransmitLineRequest) SetLine(v string) *DeleteTransmitLineRequest {
	s.Line = &v
	return s
}

func (s *DeleteTransmitLineRequest) SetDomain(v string) *DeleteTransmitLineRequest {
	s.Domain = &v
	return s
}

type DeleteTransmitLineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransmitLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransmitLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransmitLineResponseBody) SetRequestId(v string) *DeleteTransmitLineResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTransmitLineResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTransmitLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTransmitLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTransmitLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransmitLineResponse) SetHeaders(v map[string]*string) *DeleteTransmitLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransmitLineResponse) SetBody(v *DeleteTransmitLineResponseBody) *DeleteTransmitLineResponse {
	s.Body = v
	return s
}

type DescribeBackSourceCidrRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Line     *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s *DescribeBackSourceCidrRequest) SetLang(v string) *DescribeBackSourceCidrRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetLine(v string) *DescribeBackSourceCidrRequest {
	s.Line = &v
	return s
}

func (s *DescribeBackSourceCidrRequest) SetRegion(v string) *DescribeBackSourceCidrRequest {
	s.Region = &v
	return s
}

type DescribeBackSourceCidrResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CidrList  *DescribeBackSourceCidrResponseBodyCidrList `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Struct"`
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

func (s *DescribeBackSourceCidrResponseBody) SetCidrList(v *DescribeBackSourceCidrResponseBodyCidrList) *DescribeBackSourceCidrResponseBody {
	s.CidrList = v
	return s
}

type DescribeBackSourceCidrResponseBodyCidrList struct {
	Cidr []*string `json:"Cidr,omitempty" xml:"Cidr,omitempty" type:"Repeated"`
}

func (s DescribeBackSourceCidrResponseBodyCidrList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackSourceCidrResponseBodyCidrList) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponseBodyCidrList) SetCidr(v []*string) *DescribeBackSourceCidrResponseBodyCidrList {
	s.Cidr = v
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

type DescribeBizFlowRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeBizFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizFlowRequest) SetSourceIp(v string) *DescribeBizFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBizFlowRequest) SetLang(v string) *DescribeBizFlowRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBizFlowRequest) SetStartTime(v int64) *DescribeBizFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBizFlowRequest) SetEndTime(v int64) *DescribeBizFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBizFlowRequest) SetIp(v string) *DescribeBizFlowRequest {
	s.Ip = &v
	return s
}

type DescribeBizFlowResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeBizFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeBizFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizFlowResponseBody) SetRequestId(v string) *DescribeBizFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizFlowResponseBody) SetData(v *DescribeBizFlowResponseBodyData) *DescribeBizFlowResponseBody {
	s.Data = v
	return s
}

type DescribeBizFlowResponseBodyData struct {
	InKbps    []*string                                 `json:"InKbps,omitempty" xml:"InKbps,omitempty" type:"Repeated"`
	OutKbps   []*string                                 `json:"OutKbps,omitempty" xml:"OutKbps,omitempty" type:"Repeated"`
	TimeScope *DescribeBizFlowResponseBodyDataTimeScope `json:"TimeScope,omitempty" xml:"TimeScope,omitempty" type:"Struct"`
}

func (s DescribeBizFlowResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBizFlowResponseBodyData) SetInKbps(v []*string) *DescribeBizFlowResponseBodyData {
	s.InKbps = v
	return s
}

func (s *DescribeBizFlowResponseBodyData) SetOutKbps(v []*string) *DescribeBizFlowResponseBodyData {
	s.OutKbps = v
	return s
}

func (s *DescribeBizFlowResponseBodyData) SetTimeScope(v *DescribeBizFlowResponseBodyDataTimeScope) *DescribeBizFlowResponseBodyData {
	s.TimeScope = v
	return s
}

type DescribeBizFlowResponseBodyDataTimeScope struct {
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval  *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeBizFlowResponseBodyDataTimeScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizFlowResponseBodyDataTimeScope) GoString() string {
	return s.String()
}

func (s *DescribeBizFlowResponseBodyDataTimeScope) SetStartTime(v int64) *DescribeBizFlowResponseBodyDataTimeScope {
	s.StartTime = &v
	return s
}

func (s *DescribeBizFlowResponseBodyDataTimeScope) SetInterval(v int32) *DescribeBizFlowResponseBodyDataTimeScope {
	s.Interval = &v
	return s
}

type DescribeBizFlowResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBizFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBizFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizFlowResponse) SetHeaders(v map[string]*string) *DescribeBizFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizFlowResponse) SetBody(v *DescribeBizFlowResponseBody) *DescribeBizFlowResponse {
	s.Body = v
	return s
}

type DescribeCcEventsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s DescribeCcEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCcEventsRequest) SetSourceIp(v string) *DescribeCcEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCcEventsRequest) SetLang(v string) *DescribeCcEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCcEventsRequest) SetStartTime(v int64) *DescribeCcEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCcEventsRequest) SetDomain(v string) *DescribeCcEventsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCcEventsRequest) SetEndTime(v int64) *DescribeCcEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCcEventsRequest) SetPageSize(v int32) *DescribeCcEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCcEventsRequest) SetPageNo(v int32) *DescribeCcEventsRequest {
	s.PageNo = &v
	return s
}

type DescribeCcEventsResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventList []*DescribeCcEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCcEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCcEventsResponseBody) SetRequestId(v string) *DescribeCcEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCcEventsResponseBody) SetEventList(v []*DescribeCcEventsResponseBodyEventList) *DescribeCcEventsResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeCcEventsResponseBody) SetTotal(v int32) *DescribeCcEventsResponseBody {
	s.Total = &v
	return s
}

type DescribeCcEventsResponseBodyEventList struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	AttackFinished *bool   `json:"AttackFinished,omitempty" xml:"AttackFinished,omitempty"`
	MaxQps         *int32  `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	Duration       *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	BlockedCount   *int32  `json:"BlockedCount,omitempty" xml:"BlockedCount,omitempty"`
}

func (s DescribeCcEventsResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcEventsResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeCcEventsResponseBodyEventList) SetEndTime(v string) *DescribeCcEventsResponseBodyEventList {
	s.EndTime = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetStartTime(v string) *DescribeCcEventsResponseBodyEventList {
	s.StartTime = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetDomain(v string) *DescribeCcEventsResponseBodyEventList {
	s.Domain = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetAttackFinished(v bool) *DescribeCcEventsResponseBodyEventList {
	s.AttackFinished = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetMaxQps(v int32) *DescribeCcEventsResponseBodyEventList {
	s.MaxQps = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetDuration(v int32) *DescribeCcEventsResponseBodyEventList {
	s.Duration = &v
	return s
}

func (s *DescribeCcEventsResponseBodyEventList) SetBlockedCount(v int32) *DescribeCcEventsResponseBodyEventList {
	s.BlockedCount = &v
	return s
}

type DescribeCcEventsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCcEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCcEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCcEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCcEventsResponse) SetHeaders(v map[string]*string) *DescribeCcEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCcEventsResponse) SetBody(v *DescribeCcEventsResponseBody) *DescribeCcEventsResponse {
	s.Body = v
	return s
}

type DescribeCnameAutoStatusRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeCnameAutoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCnameAutoStatusRequest) SetResourceOwnerId(v int64) *DescribeCnameAutoStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCnameAutoStatusRequest) SetDomain(v string) *DescribeCnameAutoStatusRequest {
	s.Domain = &v
	return s
}

type DescribeCnameAutoStatusResponseBody struct {
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCnameAutoStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameAutoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCnameAutoStatusResponseBody) SetStatus(v bool) *DescribeCnameAutoStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCnameAutoStatusResponseBody) SetRequestId(v string) *DescribeCnameAutoStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCnameAutoStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCnameAutoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCnameAutoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCnameAutoStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCnameAutoStatusResponse) SetHeaders(v map[string]*string) *DescribeCnameAutoStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCnameAutoStatusResponse) SetBody(v *DescribeCnameAutoStatusResponseBody) *DescribeCnameAutoStatusResponse {
	s.Body = v
	return s
}

type DescribeDdosAttackEventsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDdosAttackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventsRequest) SetSourceIp(v string) *DescribeDdosAttackEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetLang(v string) *DescribeDdosAttackEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetPageSize(v int32) *DescribeDdosAttackEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetStartTime(v int64) *DescribeDdosAttackEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetEndTime(v int64) *DescribeDdosAttackEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetIp(v string) *DescribeDdosAttackEventsRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosAttackEventsRequest) SetCurrentPage(v int32) *DescribeDdosAttackEventsRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDdosAttackEventsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDdosAttackEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDdosAttackEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventsResponseBody) SetRequestId(v string) *DescribeDdosAttackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosAttackEventsResponseBody) SetData(v *DescribeDdosAttackEventsResponseBodyData) *DescribeDdosAttackEventsResponseBody {
	s.Data = v
	return s
}

type DescribeDdosAttackEventsResponseBodyData struct {
	EventList  []*DescribeDdosAttackEventsResponseBodyDataEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDdosAttackEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventsResponseBodyData) SetEventList(v []*DescribeDdosAttackEventsResponseBodyDataEventList) *DescribeDdosAttackEventsResponseBodyData {
	s.EventList = v
	return s
}

func (s *DescribeDdosAttackEventsResponseBodyData) SetTotalCount(v int32) *DescribeDdosAttackEventsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeDdosAttackEventsResponseBodyDataEventList struct {
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	Result     *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
	Duration   *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeDdosAttackEventsResponseBodyDataEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventsResponseBodyDataEventList) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventsResponseBodyDataEventList) SetEndTime(v int64) *DescribeDdosAttackEventsResponseBodyDataEventList {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAttackEventsResponseBodyDataEventList) SetStartTime(v int64) *DescribeDdosAttackEventsResponseBodyDataEventList {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAttackEventsResponseBodyDataEventList) SetAttackType(v string) *DescribeDdosAttackEventsResponseBodyDataEventList {
	s.AttackType = &v
	return s
}

func (s *DescribeDdosAttackEventsResponseBodyDataEventList) SetResult(v int32) *DescribeDdosAttackEventsResponseBodyDataEventList {
	s.Result = &v
	return s
}

func (s *DescribeDdosAttackEventsResponseBodyDataEventList) SetDuration(v string) *DescribeDdosAttackEventsResponseBodyDataEventList {
	s.Duration = &v
	return s
}

type DescribeDdosAttackEventsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosAttackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosAttackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventsResponse) SetHeaders(v map[string]*string) *DescribeDdosAttackEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosAttackEventsResponse) SetBody(v *DescribeDdosAttackEventsResponseBody) *DescribeDdosAttackEventsResponse {
	s.Body = v
	return s
}

type DescribeDdosAttackEventSourceIpsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDdosAttackEventSourceIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventSourceIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetSourceIp(v string) *DescribeDdosAttackEventSourceIpsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetLang(v string) *DescribeDdosAttackEventSourceIpsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetPageSize(v int32) *DescribeDdosAttackEventSourceIpsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetStartTime(v int64) *DescribeDdosAttackEventSourceIpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetEndTime(v int64) *DescribeDdosAttackEventSourceIpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetIp(v string) *DescribeDdosAttackEventSourceIpsRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsRequest) SetCurrentPage(v int32) *DescribeDdosAttackEventSourceIpsRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDdosAttackEventSourceIpsResponseBody struct {
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeDdosAttackEventSourceIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeDdosAttackEventSourceIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventSourceIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventSourceIpsResponseBody) SetRequestId(v string) *DescribeDdosAttackEventSourceIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsResponseBody) SetData(v *DescribeDdosAttackEventSourceIpsResponseBodyData) *DescribeDdosAttackEventSourceIpsResponseBody {
	s.Data = v
	return s
}

type DescribeDdosAttackEventSourceIpsResponseBodyData struct {
	IpList     []*DescribeDdosAttackEventSourceIpsResponseBodyDataIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDdosAttackEventSourceIpsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventSourceIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventSourceIpsResponseBodyData) SetIpList(v []*DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) *DescribeDdosAttackEventSourceIpsResponseBodyData {
	s.IpList = v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsResponseBodyData) SetTotalCount(v int32) *DescribeDdosAttackEventSourceIpsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeDdosAttackEventSourceIpsResponseBodyDataIpList struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InBps    *int32  `json:"InBps,omitempty" xml:"InBps,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
}

func (s DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) SetSourceIp(v string) *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) SetInBps(v int32) *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList {
	s.InBps = &v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList) SetCity(v string) *DescribeDdosAttackEventSourceIpsResponseBodyDataIpList {
	s.City = &v
	return s
}

type DescribeDdosAttackEventSourceIpsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosAttackEventSourceIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosAttackEventSourceIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackEventSourceIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackEventSourceIpsResponse) SetHeaders(v map[string]*string) *DescribeDdosAttackEventSourceIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosAttackEventSourceIpsResponse) SetBody(v *DescribeDdosAttackEventSourceIpsResponseBody) *DescribeDdosAttackEventSourceIpsResponse {
	s.Body = v
	return s
}

type DescribeDdosAttackTypeChartRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDdosAttackTypeChartRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackTypeChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackTypeChartRequest) SetSourceIp(v string) *DescribeDdosAttackTypeChartRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosAttackTypeChartRequest) SetLang(v string) *DescribeDdosAttackTypeChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosAttackTypeChartRequest) SetStartTime(v int64) *DescribeDdosAttackTypeChartRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAttackTypeChartRequest) SetEndTime(v int64) *DescribeDdosAttackTypeChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAttackTypeChartRequest) SetIp(v string) *DescribeDdosAttackTypeChartRequest {
	s.Ip = &v
	return s
}

type DescribeDdosAttackTypeChartResponseBody struct {
	AttckCount *int32  `json:"AttckCount,omitempty" xml:"AttckCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AttckType  *string `json:"AttckType,omitempty" xml:"AttckType,omitempty"`
	DropCount  *int32  `json:"DropCount,omitempty" xml:"DropCount,omitempty"`
	DropType   *string `json:"DropType,omitempty" xml:"DropType,omitempty"`
}

func (s DescribeDdosAttackTypeChartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackTypeChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackTypeChartResponseBody) SetAttckCount(v int32) *DescribeDdosAttackTypeChartResponseBody {
	s.AttckCount = &v
	return s
}

func (s *DescribeDdosAttackTypeChartResponseBody) SetRequestId(v string) *DescribeDdosAttackTypeChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosAttackTypeChartResponseBody) SetAttckType(v string) *DescribeDdosAttackTypeChartResponseBody {
	s.AttckType = &v
	return s
}

func (s *DescribeDdosAttackTypeChartResponseBody) SetDropCount(v int32) *DescribeDdosAttackTypeChartResponseBody {
	s.DropCount = &v
	return s
}

func (s *DescribeDdosAttackTypeChartResponseBody) SetDropType(v string) *DescribeDdosAttackTypeChartResponseBody {
	s.DropType = &v
	return s
}

type DescribeDdosAttackTypeChartResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosAttackTypeChartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosAttackTypeChartResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosAttackTypeChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosAttackTypeChartResponse) SetHeaders(v map[string]*string) *DescribeDdosAttackTypeChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosAttackTypeChartResponse) SetBody(v *DescribeDdosAttackTypeChartResponseBody) *DescribeDdosAttackTypeChartResponse {
	s.Body = v
	return s
}

type DescribeDdosFlowProportionDiagramRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDdosFlowProportionDiagramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosFlowProportionDiagramRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosFlowProportionDiagramRequest) SetSourceIp(v string) *DescribeDdosFlowProportionDiagramRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramRequest) SetLang(v string) *DescribeDdosFlowProportionDiagramRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramRequest) SetStartTime(v int64) *DescribeDdosFlowProportionDiagramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramRequest) SetEndTime(v int64) *DescribeDdosFlowProportionDiagramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramRequest) SetIp(v string) *DescribeDdosFlowProportionDiagramRequest {
	s.Ip = &v
	return s
}

type DescribeDdosFlowProportionDiagramResponseBody struct {
	TotalBps  *int32  `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DropPps   *int32  `json:"DropPps,omitempty" xml:"DropPps,omitempty"`
	DropBps   *int32  `json:"DropBps,omitempty" xml:"DropBps,omitempty"`
	TotalPps  *int32  `json:"TotalPps,omitempty" xml:"TotalPps,omitempty"`
}

func (s DescribeDdosFlowProportionDiagramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosFlowProportionDiagramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosFlowProportionDiagramResponseBody) SetTotalBps(v int32) *DescribeDdosFlowProportionDiagramResponseBody {
	s.TotalBps = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramResponseBody) SetRequestId(v string) *DescribeDdosFlowProportionDiagramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramResponseBody) SetDropPps(v int32) *DescribeDdosFlowProportionDiagramResponseBody {
	s.DropPps = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramResponseBody) SetDropBps(v int32) *DescribeDdosFlowProportionDiagramResponseBody {
	s.DropBps = &v
	return s
}

func (s *DescribeDdosFlowProportionDiagramResponseBody) SetTotalPps(v int32) *DescribeDdosFlowProportionDiagramResponseBody {
	s.TotalPps = &v
	return s
}

type DescribeDdosFlowProportionDiagramResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosFlowProportionDiagramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosFlowProportionDiagramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosFlowProportionDiagramResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosFlowProportionDiagramResponse) SetHeaders(v map[string]*string) *DescribeDdosFlowProportionDiagramResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosFlowProportionDiagramResponse) SetBody(v *DescribeDdosFlowProportionDiagramResponseBody) *DescribeDdosFlowProportionDiagramResponse {
	s.Body = v
	return s
}

type DescribeDdosIpConfigRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Index    *int32    `json:"Index,omitempty" xml:"Index,omitempty"`
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Ips      []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s DescribeDdosIpConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosIpConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosIpConfigRequest) SetSourceIp(v string) *DescribeDdosIpConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosIpConfigRequest) SetLang(v string) *DescribeDdosIpConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosIpConfigRequest) SetIndex(v int32) *DescribeDdosIpConfigRequest {
	s.Index = &v
	return s
}

func (s *DescribeDdosIpConfigRequest) SetPageSize(v int32) *DescribeDdosIpConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosIpConfigRequest) SetIps(v []*string) *DescribeDdosIpConfigRequest {
	s.Ips = v
	return s
}

type DescribeDdosIpConfigResponseBody struct {
	DataList  []*DescribeDdosIpConfigResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosIpConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosIpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosIpConfigResponseBody) SetDataList(v []*DescribeDdosIpConfigResponseBodyDataList) *DescribeDdosIpConfigResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDdosIpConfigResponseBody) SetRequestId(v string) *DescribeDdosIpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBody) SetTotal(v int32) *DescribeDdosIpConfigResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosIpConfigResponseBodyDataList struct {
	Status            *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	CleanStatus       *int32  `json:"CleanStatus,omitempty" xml:"CleanStatus,omitempty"`
	Bandwidth         *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConfigDomainCount *int32  `json:"ConfigDomainCount,omitempty" xml:"ConfigDomainCount,omitempty"`
	Line              *string `json:"Line,omitempty" xml:"Line,omitempty"`
	ElasticBandwidth  *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	LbId              *string `json:"LbId,omitempty" xml:"LbId,omitempty"`
	ConfigPortCount   *int32  `json:"ConfigPortCount,omitempty" xml:"ConfigPortCount,omitempty"`
	Ip                *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	TotalDefenseCount *int32  `json:"TotalDefenseCount,omitempty" xml:"TotalDefenseCount,omitempty"`
}

func (s DescribeDdosIpConfigResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosIpConfigResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetStatus(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetCleanStatus(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.CleanStatus = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetBandwidth(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetConfigDomainCount(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.ConfigDomainCount = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetLine(v string) *DescribeDdosIpConfigResponseBodyDataList {
	s.Line = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetElasticBandwidth(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetLbId(v string) *DescribeDdosIpConfigResponseBodyDataList {
	s.LbId = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetConfigPortCount(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.ConfigPortCount = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetIp(v string) *DescribeDdosIpConfigResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *DescribeDdosIpConfigResponseBodyDataList) SetTotalDefenseCount(v int32) *DescribeDdosIpConfigResponseBodyDataList {
	s.TotalDefenseCount = &v
	return s
}

type DescribeDdosIpConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosIpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosIpConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosIpConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosIpConfigResponse) SetHeaders(v map[string]*string) *DescribeDdosIpConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosIpConfigResponse) SetBody(v *DescribeDdosIpConfigResponseBody) *DescribeDdosIpConfigResponse {
	s.Body = v
	return s
}

type DescribeDdosPeakFlowRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDdosPeakFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosPeakFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosPeakFlowRequest) SetSourceIp(v string) *DescribeDdosPeakFlowRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosPeakFlowRequest) SetLang(v string) *DescribeDdosPeakFlowRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDdosPeakFlowRequest) SetStartTime(v int64) *DescribeDdosPeakFlowRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosPeakFlowRequest) SetEndTime(v int64) *DescribeDdosPeakFlowRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosPeakFlowRequest) SetIp(v string) *DescribeDdosPeakFlowRequest {
	s.Ip = &v
	return s
}

type DescribeDdosPeakFlowResponseBody struct {
	PeakFlow  *string `json:"PeakFlow,omitempty" xml:"PeakFlow,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDdosPeakFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosPeakFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosPeakFlowResponseBody) SetPeakFlow(v string) *DescribeDdosPeakFlowResponseBody {
	s.PeakFlow = &v
	return s
}

func (s *DescribeDdosPeakFlowResponseBody) SetRequestId(v string) *DescribeDdosPeakFlowResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDdosPeakFlowResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosPeakFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosPeakFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosPeakFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosPeakFlowResponse) SetHeaders(v map[string]*string) *DescribeDdosPeakFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosPeakFlowResponse) SetBody(v *DescribeDdosPeakFlowResponseBody) *DescribeDdosPeakFlowResponse {
	s.Body = v
	return s
}

type DescribeDomainConfigPageRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s DescribeDomainConfigPageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageRequest) SetSourceIp(v string) *DescribeDomainConfigPageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainConfigPageRequest) SetLang(v string) *DescribeDomainConfigPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainConfigPageRequest) SetDomain(v string) *DescribeDomainConfigPageRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainConfigPageRequest) SetPageSize(v int32) *DescribeDomainConfigPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainConfigPageRequest) SetPageNo(v int32) *DescribeDomainConfigPageRequest {
	s.PageNo = &v
	return s
}

type DescribeDomainConfigPageResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	ConfigList []*DescribeDomainConfigPageResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s DescribeDomainConfigPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageResponseBody) SetRequestId(v string) *DescribeDomainConfigPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBody) SetTotal(v int32) *DescribeDomainConfigPageResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBody) SetConfigList(v []*DescribeDomainConfigPageResponseBodyConfigList) *DescribeDomainConfigPageResponseBody {
	s.ConfigList = v
	return s
}

type DescribeDomainConfigPageResponseBodyConfigList struct {
	Domain    *string                                                    `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Cname     *string                                                    `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Instances []*DescribeDomainConfigPageResponseBodyConfigListInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeDomainConfigPageResponseBodyConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageResponseBodyConfigList) SetDomain(v string) *DescribeDomainConfigPageResponseBodyConfigList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigList) SetCname(v string) *DescribeDomainConfigPageResponseBodyConfigList {
	s.Cname = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigList) SetInstances(v []*DescribeDomainConfigPageResponseBodyConfigListInstances) *DescribeDomainConfigPageResponseBodyConfigList {
	s.Instances = v
	return s
}

type DescribeDomainConfigPageResponseBodyConfigListInstances struct {
	InstanceRemark *string                                                         `json:"InstanceRemark,omitempty" xml:"InstanceRemark,omitempty"`
	InstanceId     *string                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Rules          []*DescribeDomainConfigPageResponseBodyConfigListInstancesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeDomainConfigPageResponseBodyConfigListInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageResponseBodyConfigListInstances) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstances) SetInstanceRemark(v string) *DescribeDomainConfigPageResponseBodyConfigListInstances {
	s.InstanceRemark = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstances) SetInstanceId(v string) *DescribeDomainConfigPageResponseBodyConfigListInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstances) SetRules(v []*DescribeDomainConfigPageResponseBodyConfigListInstancesRules) *DescribeDomainConfigPageResponseBodyConfigListInstances {
	s.Rules = v
	return s
}

type DescribeDomainConfigPageResponseBodyConfigListInstancesRules struct {
	ProxyTypeList []*string `json:"ProxyTypeList,omitempty" xml:"ProxyTypeList,omitempty" type:"Repeated"`
	Line          *string   `json:"Line,omitempty" xml:"Line,omitempty"`
	RealServers   []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	Ip            *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeDomainConfigPageResponseBodyConfigListInstancesRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageResponseBodyConfigListInstancesRules) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstancesRules) SetProxyTypeList(v []*string) *DescribeDomainConfigPageResponseBodyConfigListInstancesRules {
	s.ProxyTypeList = v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstancesRules) SetLine(v string) *DescribeDomainConfigPageResponseBodyConfigListInstancesRules {
	s.Line = &v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstancesRules) SetRealServers(v []*string) *DescribeDomainConfigPageResponseBodyConfigListInstancesRules {
	s.RealServers = v
	return s
}

func (s *DescribeDomainConfigPageResponseBodyConfigListInstancesRules) SetIp(v string) *DescribeDomainConfigPageResponseBodyConfigListInstancesRules {
	s.Ip = &v
	return s
}

type DescribeDomainConfigPageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainConfigPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainConfigPageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainConfigPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainConfigPageResponse) SetHeaders(v map[string]*string) *DescribeDomainConfigPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainConfigPageResponse) SetBody(v *DescribeDomainConfigPageResponseBody) *DescribeDomainConfigPageResponse {
	s.Body = v
	return s
}

type DescribeDomainSecurityConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainSecurityConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSecurityConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityConfigRequest) SetSourceIp(v string) *DescribeDomainSecurityConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecurityConfigRequest) SetLang(v string) *DescribeDomainSecurityConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecurityConfigRequest) SetDomain(v string) *DescribeDomainSecurityConfigRequest {
	s.Domain = &v
	return s
}

type DescribeDomainSecurityConfigResponseBody struct {
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CcInfo      *DescribeDomainSecurityConfigResponseBodyCcInfo `json:"CcInfo,omitempty" xml:"CcInfo,omitempty" type:"Struct"`
	CnameEnable *bool                                           `json:"CnameEnable,omitempty" xml:"CnameEnable,omitempty"`
	WhiteList   *string                                         `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
	BlackList   *string                                         `json:"BlackList,omitempty" xml:"BlackList,omitempty"`
	CnameMode   *int32                                          `json:"CnameMode,omitempty" xml:"CnameMode,omitempty"`
}

func (s DescribeDomainSecurityConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSecurityConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityConfigResponseBody) SetRequestId(v string) *DescribeDomainSecurityConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBody) SetCcInfo(v *DescribeDomainSecurityConfigResponseBodyCcInfo) *DescribeDomainSecurityConfigResponseBody {
	s.CcInfo = v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBody) SetCnameEnable(v bool) *DescribeDomainSecurityConfigResponseBody {
	s.CnameEnable = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBody) SetWhiteList(v string) *DescribeDomainSecurityConfigResponseBody {
	s.WhiteList = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBody) SetBlackList(v string) *DescribeDomainSecurityConfigResponseBody {
	s.BlackList = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBody) SetCnameMode(v int32) *DescribeDomainSecurityConfigResponseBody {
	s.CnameMode = &v
	return s
}

type DescribeDomainSecurityConfigResponseBodyCcInfo struct {
	CcCustomRuleCount  *int32  `json:"CcCustomRuleCount,omitempty" xml:"CcCustomRuleCount,omitempty"`
	CcSwitch           *bool   `json:"CcSwitch,omitempty" xml:"CcSwitch,omitempty"`
	CcTemplate         *string `json:"CcTemplate,omitempty" xml:"CcTemplate,omitempty"`
	CcCustomRuleEnable *bool   `json:"CcCustomRuleEnable,omitempty" xml:"CcCustomRuleEnable,omitempty"`
}

func (s DescribeDomainSecurityConfigResponseBodyCcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSecurityConfigResponseBodyCcInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityConfigResponseBodyCcInfo) SetCcCustomRuleCount(v int32) *DescribeDomainSecurityConfigResponseBodyCcInfo {
	s.CcCustomRuleCount = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBodyCcInfo) SetCcSwitch(v bool) *DescribeDomainSecurityConfigResponseBodyCcInfo {
	s.CcSwitch = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBodyCcInfo) SetCcTemplate(v string) *DescribeDomainSecurityConfigResponseBodyCcInfo {
	s.CcTemplate = &v
	return s
}

func (s *DescribeDomainSecurityConfigResponseBodyCcInfo) SetCcCustomRuleEnable(v bool) *DescribeDomainSecurityConfigResponseBodyCcInfo {
	s.CcCustomRuleEnable = &v
	return s
}

type DescribeDomainSecurityConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainSecurityConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainSecurityConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainSecurityConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecurityConfigResponse) SetHeaders(v map[string]*string) *DescribeDomainSecurityConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecurityConfigResponse) SetBody(v *DescribeDomainSecurityConfigResponseBody) *DescribeDomainSecurityConfigResponse {
	s.Body = v
	return s
}

type DescribeHealthCheckConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeHealthCheckConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigRequest) SetSourceIp(v string) *DescribeHealthCheckConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHealthCheckConfigRequest) SetLang(v string) *DescribeHealthCheckConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHealthCheckConfigRequest) SetIp(v string) *DescribeHealthCheckConfigRequest {
	s.Ip = &v
	return s
}

type DescribeHealthCheckConfigResponseBody struct {
	Listeners []*DescribeHealthCheckConfigResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHealthCheckConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBody) SetListeners(v []*DescribeHealthCheckConfigResponseBodyListeners) *DescribeHealthCheckConfigResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeHealthCheckConfigResponseBody) SetRequestId(v string) *DescribeHealthCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeHealthCheckConfigResponseBodyListeners struct {
	FrontendPort *int32                                                `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	Check        *DescribeHealthCheckConfigResponseBodyListenersCheck  `json:"Check,omitempty" xml:"Check,omitempty" type:"Struct"`
	Protocol     *string                                               `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	BackPort     *int32                                                `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	Config       *DescribeHealthCheckConfigResponseBodyListenersConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
}

func (s DescribeHealthCheckConfigResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListeners) SetFrontendPort(v int32) *DescribeHealthCheckConfigResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListeners) SetCheck(v *DescribeHealthCheckConfigResponseBodyListenersCheck) *DescribeHealthCheckConfigResponseBodyListeners {
	s.Check = v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListeners) SetProtocol(v string) *DescribeHealthCheckConfigResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListeners) SetBackPort(v int32) *DescribeHealthCheckConfigResponseBodyListeners {
	s.BackPort = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListeners) SetConfig(v *DescribeHealthCheckConfigResponseBodyListenersConfig) *DescribeHealthCheckConfigResponseBodyListeners {
	s.Config = v
	return s
}

type DescribeHealthCheckConfigResponseBodyListenersCheck struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Timeout  *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Up       *int32  `json:"Up,omitempty" xml:"Up,omitempty"`
	Down     *int32  `json:"Down,omitempty" xml:"Down,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeHealthCheckConfigResponseBodyListenersCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListenersCheck) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetType(v string) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetTimeout(v int32) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Timeout = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetDomain(v string) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Domain = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetInterval(v int32) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Interval = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetUp(v int32) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Up = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetDown(v int32) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Down = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetPort(v int32) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Port = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersCheck) SetUri(v string) *DescribeHealthCheckConfigResponseBodyListenersCheck {
	s.Uri = &v
	return s
}

type DescribeHealthCheckConfigResponseBodyListenersConfig struct {
	SynProxy           *string                                                            `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	PersistenceTimeout *int32                                                             `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	NoDataConn         *string                                                            `json:"NoDataConn,omitempty" xml:"NoDataConn,omitempty"`
	Sla                *DescribeHealthCheckConfigResponseBodyListenersConfigSla           `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	PayloadLength      *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength `json:"PayloadLength,omitempty" xml:"PayloadLength,omitempty" type:"Struct"`
	Slimit             *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit        `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfig) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetSynProxy(v string) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.SynProxy = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetPersistenceTimeout(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetNoDataConn(v string) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.NoDataConn = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetSla(v *DescribeHealthCheckConfigResponseBodyListenersConfigSla) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.Sla = v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetPayloadLength(v *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.PayloadLength = v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfig) SetSlimit(v *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) *DescribeHealthCheckConfigResponseBodyListenersConfig {
	s.Slimit = v
	return s
}

type DescribeHealthCheckConfigResponseBodyListenersConfigSla struct {
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	MaxConnEnable *int32 `json:"MaxConnEnable,omitempty" xml:"MaxConnEnable,omitempty"`
	MaxConn       *int32 `json:"MaxConn,omitempty" xml:"MaxConn,omitempty"`
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigSla) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSla) SetCpsEnable(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSla) SetCps(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSla) SetMaxConnEnable(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSla {
	s.MaxConnEnable = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSla) SetMaxConn(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSla {
	s.MaxConn = &v
	return s
}

type DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength struct {
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength) SetMax(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength {
	s.Max = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength) SetMin(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigPayloadLength {
	s.Min = &v
	return s
}

type DescribeHealthCheckConfigResponseBodyListenersConfigSlimit struct {
	CpsEnable     *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	Cps           *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	MaxConnEnable *int32 `json:"MaxConnEnable,omitempty" xml:"MaxConnEnable,omitempty"`
	MaxConn       *int32 `json:"MaxConn,omitempty" xml:"MaxConn,omitempty"`
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) SetCpsEnable(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) SetCps(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) SetMaxConnEnable(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit {
	s.MaxConnEnable = &v
	return s
}

func (s *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit) SetMaxConn(v int32) *DescribeHealthCheckConfigResponseBodyListenersConfigSlimit {
	s.MaxConn = &v
	return s
}

type DescribeHealthCheckConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHealthCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHealthCheckConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHealthCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckConfigResponse) SetHeaders(v map[string]*string) *DescribeHealthCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeHealthCheckConfigResponse) SetBody(v *DescribeHealthCheckConfigResponseBody) *DescribeHealthCheckConfigResponse {
	s.Body = v
	return s
}

type DescribeInstancePageRequest struct {
	SourceIp       *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize       *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId     *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Line           *string   `json:"Line,omitempty" xml:"Line,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	IpList         []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
}

func (s DescribeInstancePageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePageRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePageRequest) SetSourceIp(v string) *DescribeInstancePageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancePageRequest) SetLang(v string) *DescribeInstancePageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancePageRequest) SetPageSize(v int32) *DescribeInstancePageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancePageRequest) SetCurrentPage(v int32) *DescribeInstancePageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancePageRequest) SetInstanceId(v string) *DescribeInstancePageRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePageRequest) SetLine(v string) *DescribeInstancePageRequest {
	s.Line = &v
	return s
}

func (s *DescribeInstancePageRequest) SetInstanceIdList(v []*string) *DescribeInstancePageRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstancePageRequest) SetIpList(v []*string) *DescribeInstancePageRequest {
	s.IpList = v
	return s
}

type DescribeInstancePageResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	InstanceList []*DescribeInstancePageResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s DescribeInstancePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePageResponseBody) SetRequestId(v string) *DescribeInstancePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePageResponseBody) SetTotal(v int32) *DescribeInstancePageResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstancePageResponseBody) SetInstanceList(v []*DescribeInstancePageResponseBodyInstanceList) *DescribeInstancePageResponseBody {
	s.InstanceList = v
	return s
}

type DescribeInstancePageResponseBodyInstanceList struct {
	InstanceRemark *string                                               `json:"InstanceRemark,omitempty" xml:"InstanceRemark,omitempty"`
	IpList         []*DescribeInstancePageResponseBodyInstanceListIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	InstanceId     *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstancePageResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePageResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstancePageResponseBodyInstanceList) SetInstanceRemark(v string) *DescribeInstancePageResponseBodyInstanceList {
	s.InstanceRemark = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceList) SetIpList(v []*DescribeInstancePageResponseBodyInstanceListIpList) *DescribeInstancePageResponseBodyInstanceList {
	s.IpList = v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstancePageResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

type DescribeInstancePageResponseBodyInstanceListIpList struct {
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Line             *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	BandWidth        *int32  `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	ElasticBandWidth *int32  `json:"ElasticBandWidth,omitempty" xml:"ElasticBandWidth,omitempty"`
}

func (s DescribeInstancePageResponseBodyInstanceListIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePageResponseBodyInstanceListIpList) GoString() string {
	return s.String()
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetStatus(v int32) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.Status = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetLine(v string) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.Line = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetIp(v string) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.Ip = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetInstanceId(v string) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetBandWidth(v int32) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.BandWidth = &v
	return s
}

func (s *DescribeInstancePageResponseBodyInstanceListIpList) SetElasticBandWidth(v int32) *DescribeInstancePageResponseBodyInstanceListIpList {
	s.ElasticBandWidth = &v
	return s
}

type DescribeInstancePageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancePageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancePageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePageResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePageResponse) SetHeaders(v map[string]*string) *DescribeInstancePageResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePageResponse) SetBody(v *DescribeInstancePageResponseBody) *DescribeInstancePageResponse {
	s.Body = v
	return s
}

type DescribePortRulePageRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribePortRulePageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePortRulePageRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRulePageRequest) SetSourceIp(v string) *DescribePortRulePageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePortRulePageRequest) SetLang(v string) *DescribePortRulePageRequest {
	s.Lang = &v
	return s
}

func (s *DescribePortRulePageRequest) SetPageSize(v int32) *DescribePortRulePageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePortRulePageRequest) SetIp(v string) *DescribePortRulePageRequest {
	s.Ip = &v
	return s
}

func (s *DescribePortRulePageRequest) SetCurrentPage(v int32) *DescribePortRulePageRequest {
	s.CurrentPage = &v
	return s
}

type DescribePortRulePageResponseBody struct {
	RuleList  []*DescribePortRulePageResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Count     *int32                                      `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePortRulePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePortRulePageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortRulePageResponseBody) SetRuleList(v []*DescribePortRulePageResponseBodyRuleList) *DescribePortRulePageResponseBody {
	s.RuleList = v
	return s
}

func (s *DescribePortRulePageResponseBody) SetRequestId(v string) *DescribePortRulePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortRulePageResponseBody) SetCount(v int32) *DescribePortRulePageResponseBody {
	s.Count = &v
	return s
}

type DescribePortRulePageResponseBodyRuleList struct {
	BackProtocol  *string `json:"BackProtocol,omitempty" xml:"BackProtocol,omitempty"`
	BackPort      *int32  `json:"BackPort,omitempty" xml:"BackPort,omitempty"`
	LbId          *string `json:"LbId,omitempty" xml:"LbId,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	LvsType       *string `json:"LvsType,omitempty" xml:"LvsType,omitempty"`
	RealServer    *string `json:"RealServer,omitempty" xml:"RealServer,omitempty"`
	FrontPort     *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	FrontProtocol *string `json:"FrontProtocol,omitempty" xml:"FrontProtocol,omitempty"`
}

func (s DescribePortRulePageResponseBodyRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribePortRulePageResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *DescribePortRulePageResponseBodyRuleList) SetBackProtocol(v string) *DescribePortRulePageResponseBodyRuleList {
	s.BackProtocol = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetBackPort(v int32) *DescribePortRulePageResponseBodyRuleList {
	s.BackPort = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetLbId(v string) *DescribePortRulePageResponseBodyRuleList {
	s.LbId = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetIp(v string) *DescribePortRulePageResponseBodyRuleList {
	s.Ip = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetLvsType(v string) *DescribePortRulePageResponseBodyRuleList {
	s.LvsType = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetRealServer(v string) *DescribePortRulePageResponseBodyRuleList {
	s.RealServer = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetFrontPort(v int32) *DescribePortRulePageResponseBodyRuleList {
	s.FrontPort = &v
	return s
}

func (s *DescribePortRulePageResponseBodyRuleList) SetFrontProtocol(v string) *DescribePortRulePageResponseBodyRuleList {
	s.FrontProtocol = &v
	return s
}

type DescribePortRulePageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePortRulePageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePortRulePageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePortRulePageResponse) GoString() string {
	return s.String()
}

func (s *DescribePortRulePageResponse) SetHeaders(v map[string]*string) *DescribePortRulePageResponse {
	s.Headers = v
	return s
}

func (s *DescribePortRulePageResponse) SetBody(v *DescribePortRulePageResponseBody) *DescribePortRulePageResponse {
	s.Body = v
	return s
}

type ListCcCustomedRuleRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s ListCcCustomedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCcCustomedRuleRequest) GoString() string {
	return s.String()
}

func (s *ListCcCustomedRuleRequest) SetSourceIp(v string) *ListCcCustomedRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ListCcCustomedRuleRequest) SetLang(v string) *ListCcCustomedRuleRequest {
	s.Lang = &v
	return s
}

func (s *ListCcCustomedRuleRequest) SetDomain(v string) *ListCcCustomedRuleRequest {
	s.Domain = &v
	return s
}

func (s *ListCcCustomedRuleRequest) SetPageSize(v int32) *ListCcCustomedRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListCcCustomedRuleRequest) SetCurrentPage(v int32) *ListCcCustomedRuleRequest {
	s.CurrentPage = &v
	return s
}

type ListCcCustomedRuleResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RuleList   *ListCcCustomedRuleResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCcCustomedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCcCustomedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListCcCustomedRuleResponseBody) SetTotalCount(v int32) *ListCcCustomedRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCcCustomedRuleResponseBody) SetRuleList(v *ListCcCustomedRuleResponseBodyRuleList) *ListCcCustomedRuleResponseBody {
	s.RuleList = v
	return s
}

func (s *ListCcCustomedRuleResponseBody) SetRequestId(v string) *ListCcCustomedRuleResponseBody {
	s.RequestId = &v
	return s
}

type ListCcCustomedRuleResponseBodyRuleList struct {
	Rule []*ListCcCustomedRuleResponseBodyRuleListRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s ListCcCustomedRuleResponseBodyRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListCcCustomedRuleResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *ListCcCustomedRuleResponseBodyRuleList) SetRule(v []*ListCcCustomedRuleResponseBodyRuleListRule) *ListCcCustomedRuleResponseBodyRuleList {
	s.Rule = v
	return s
}

type ListCcCustomedRuleResponseBodyRuleListRule struct {
	BlockingTime *int32  `json:"BlockingTime,omitempty" xml:"BlockingTime,omitempty"`
	BlockingType *string `json:"BlockingType,omitempty" xml:"BlockingType,omitempty"`
	Interval     *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	VisitCount   *int32  `json:"VisitCount,omitempty" xml:"VisitCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	MatchingRule *string `json:"MatchingRule,omitempty" xml:"MatchingRule,omitempty"`
}

func (s ListCcCustomedRuleResponseBodyRuleListRule) String() string {
	return tea.Prettify(s)
}

func (s ListCcCustomedRuleResponseBodyRuleListRule) GoString() string {
	return s.String()
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetBlockingTime(v int32) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.BlockingTime = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetBlockingType(v string) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.BlockingType = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetInterval(v int32) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.Interval = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetVisitCount(v int32) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.VisitCount = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetName(v string) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.Name = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetUri(v string) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.Uri = &v
	return s
}

func (s *ListCcCustomedRuleResponseBodyRuleListRule) SetMatchingRule(v string) *ListCcCustomedRuleResponseBodyRuleListRule {
	s.MatchingRule = &v
	return s
}

type ListCcCustomedRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCcCustomedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCcCustomedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCcCustomedRuleResponse) GoString() string {
	return s.String()
}

func (s *ListCcCustomedRuleResponse) SetHeaders(v map[string]*string) *ListCcCustomedRuleResponse {
	s.Headers = v
	return s
}

func (s *ListCcCustomedRuleResponse) SetBody(v *ListCcCustomedRuleResponseBody) *ListCcCustomedRuleResponse {
	s.Body = v
	return s
}

type ModifyCcCustomStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyCcCustomStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcCustomStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyCcCustomStatusRequest) SetSourceIp(v string) *ModifyCcCustomStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCcCustomStatusRequest) SetLang(v string) *ModifyCcCustomStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyCcCustomStatusRequest) SetDomain(v string) *ModifyCcCustomStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCcCustomStatusRequest) SetEnable(v bool) *ModifyCcCustomStatusRequest {
	s.Enable = &v
	return s
}

type ModifyCcCustomStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCcCustomStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcCustomStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCcCustomStatusResponseBody) SetRequestId(v string) *ModifyCcCustomStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCcCustomStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCcCustomStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCcCustomStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcCustomStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyCcCustomStatusResponse) SetHeaders(v map[string]*string) *ModifyCcCustomStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyCcCustomStatusResponse) SetBody(v *ModifyCcCustomStatusResponseBody) *ModifyCcCustomStatusResponse {
	s.Body = v
	return s
}

type ModifyCcStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyCcStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyCcStatusRequest) SetSourceIp(v string) *ModifyCcStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCcStatusRequest) SetLang(v string) *ModifyCcStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyCcStatusRequest) SetDomain(v string) *ModifyCcStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCcStatusRequest) SetEnable(v bool) *ModifyCcStatusRequest {
	s.Enable = &v
	return s
}

type ModifyCcStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCcStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCcStatusResponseBody) SetRequestId(v string) *ModifyCcStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCcStatusResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCcStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCcStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyCcStatusResponse) SetHeaders(v map[string]*string) *ModifyCcStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyCcStatusResponse) SetBody(v *ModifyCcStatusResponseBody) *ModifyCcStatusResponse {
	s.Body = v
	return s
}

type ModifyCcTemplateRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Mode     *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ModifyCcTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyCcTemplateRequest) SetSourceIp(v string) *ModifyCcTemplateRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCcTemplateRequest) SetLang(v string) *ModifyCcTemplateRequest {
	s.Lang = &v
	return s
}

func (s *ModifyCcTemplateRequest) SetDomain(v string) *ModifyCcTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ModifyCcTemplateRequest) SetMode(v int32) *ModifyCcTemplateRequest {
	s.Mode = &v
	return s
}

type ModifyCcTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCcTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCcTemplateResponseBody) SetRequestId(v string) *ModifyCcTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCcTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCcTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCcTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCcTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyCcTemplateResponse) SetHeaders(v map[string]*string) *ModifyCcTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyCcTemplateResponse) SetBody(v *ModifyCcTemplateResponseBody) *ModifyCcTemplateResponse {
	s.Body = v
	return s
}

type ModifyDDoSProtectConfigRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	FrontPort  *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	LbId       *string `json:"LbId,omitempty" xml:"LbId,omitempty"`
}

func (s ModifyDDoSProtectConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDDoSProtectConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDDoSProtectConfigRequest) SetSourceIp(v string) *ModifyDDoSProtectConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDDoSProtectConfigRequest) SetLang(v string) *ModifyDDoSProtectConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDDoSProtectConfigRequest) SetIp(v string) *ModifyDDoSProtectConfigRequest {
	s.Ip = &v
	return s
}

func (s *ModifyDDoSProtectConfigRequest) SetFrontPort(v int32) *ModifyDDoSProtectConfigRequest {
	s.FrontPort = &v
	return s
}

func (s *ModifyDDoSProtectConfigRequest) SetConfigJson(v string) *ModifyDDoSProtectConfigRequest {
	s.ConfigJson = &v
	return s
}

func (s *ModifyDDoSProtectConfigRequest) SetLbId(v string) *ModifyDDoSProtectConfigRequest {
	s.LbId = &v
	return s
}

type ModifyDDoSProtectConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDDoSProtectConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDDoSProtectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDDoSProtectConfigResponseBody) SetRequestId(v string) *ModifyDDoSProtectConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDDoSProtectConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDDoSProtectConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDDoSProtectConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDDoSProtectConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDDoSProtectConfigResponse) SetHeaders(v map[string]*string) *ModifyDDoSProtectConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDDoSProtectConfigResponse) SetBody(v *ModifyDDoSProtectConfigResponseBody) *ModifyDDoSProtectConfigResponse {
	s.Body = v
	return s
}

type ModifyDomainBlackWhiteListRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Black    []*string `json:"Black,omitempty" xml:"Black,omitempty" type:"Repeated"`
	White    []*string `json:"White,omitempty" xml:"White,omitempty" type:"Repeated"`
}

func (s ModifyDomainBlackWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainBlackWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainBlackWhiteListRequest) SetSourceIp(v string) *ModifyDomainBlackWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDomainBlackWhiteListRequest) SetLang(v string) *ModifyDomainBlackWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDomainBlackWhiteListRequest) SetDomain(v string) *ModifyDomainBlackWhiteListRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainBlackWhiteListRequest) SetBlack(v []*string) *ModifyDomainBlackWhiteListRequest {
	s.Black = v
	return s
}

func (s *ModifyDomainBlackWhiteListRequest) SetWhite(v []*string) *ModifyDomainBlackWhiteListRequest {
	s.White = v
	return s
}

type ModifyDomainBlackWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainBlackWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainBlackWhiteListResponseBody) SetRequestId(v string) *ModifyDomainBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainBlackWhiteListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainBlackWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainBlackWhiteListResponse) SetHeaders(v map[string]*string) *ModifyDomainBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainBlackWhiteListResponse) SetBody(v *ModifyDomainBlackWhiteListResponseBody) *ModifyDomainBlackWhiteListResponse {
	s.Body = v
	return s
}

type ModifyDomainProxyRequest struct {
	SourceIp  *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain    *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ProxyType []*string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty" type:"Repeated"`
}

func (s ModifyDomainProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainProxyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainProxyRequest) SetSourceIp(v string) *ModifyDomainProxyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyDomainProxyRequest) SetLang(v string) *ModifyDomainProxyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDomainProxyRequest) SetDomain(v string) *ModifyDomainProxyRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainProxyRequest) SetProxyType(v []*string) *ModifyDomainProxyRequest {
	s.ProxyType = v
	return s
}

type ModifyDomainProxyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainProxyResponseBody) SetRequestId(v string) *ModifyDomainProxyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDomainProxyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDomainProxyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDomainProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDomainProxyResponse) GoString() string {
	return s.String()
}

func (s *ModifyDomainProxyResponse) SetHeaders(v map[string]*string) *ModifyDomainProxyResponse {
	s.Headers = v
	return s
}

func (s *ModifyDomainProxyResponse) SetBody(v *ModifyDomainProxyResponseBody) *ModifyDomainProxyResponse {
	s.Body = v
	return s
}

type ModifyElasticBandwidthRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ElasticBandwidth *int32  `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ModifyElasticBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandwidthRequest) SetSourceIp(v string) *ModifyElasticBandwidthRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyElasticBandwidthRequest) SetLang(v string) *ModifyElasticBandwidthRequest {
	s.Lang = &v
	return s
}

func (s *ModifyElasticBandwidthRequest) SetElasticBandwidth(v int32) *ModifyElasticBandwidthRequest {
	s.ElasticBandwidth = &v
	return s
}

func (s *ModifyElasticBandwidthRequest) SetIp(v string) *ModifyElasticBandwidthRequest {
	s.Ip = &v
	return s
}

type ModifyElasticBandwidthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandwidthResponseBody) SetRequestId(v string) *ModifyElasticBandwidthResponseBody {
	s.RequestId = &v
	return s
}

type ModifyElasticBandwidthResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyElasticBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyElasticBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyElasticBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandwidthResponse) SetHeaders(v map[string]*string) *ModifyElasticBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticBandwidthResponse) SetBody(v *ModifyElasticBandwidthResponseBody) *ModifyElasticBandwidthResponse {
	s.Body = v
	return s
}

type ModifyHealthCheckConfigRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	FrontPort  *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	LbId       *string `json:"LbId,omitempty" xml:"LbId,omitempty"`
}

func (s ModifyHealthCheckConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigRequest) SetSourceIp(v string) *ModifyHealthCheckConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetLang(v string) *ModifyHealthCheckConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetIp(v string) *ModifyHealthCheckConfigRequest {
	s.Ip = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetFrontPort(v int32) *ModifyHealthCheckConfigRequest {
	s.FrontPort = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetConfigJson(v string) *ModifyHealthCheckConfigRequest {
	s.ConfigJson = &v
	return s
}

func (s *ModifyHealthCheckConfigRequest) SetLbId(v string) *ModifyHealthCheckConfigRequest {
	s.LbId = &v
	return s
}

type ModifyHealthCheckConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHealthCheckConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponseBody) SetRequestId(v string) *ModifyHealthCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHealthCheckConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHealthCheckConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHealthCheckConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHealthCheckConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckConfigResponse) SetHeaders(v map[string]*string) *ModifyHealthCheckConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyHealthCheckConfigResponse) SetBody(v *ModifyHealthCheckConfigResponseBody) *ModifyHealthCheckConfigResponse {
	s.Body = v
	return s
}

type ModifyIpCnameStatusRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Enable   *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ModifyIpCnameStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpCnameStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpCnameStatusRequest) SetSourceIp(v string) *ModifyIpCnameStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyIpCnameStatusRequest) SetLang(v string) *ModifyIpCnameStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyIpCnameStatusRequest) SetDomain(v string) *ModifyIpCnameStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyIpCnameStatusRequest) SetIp(v string) *ModifyIpCnameStatusRequest {
	s.Ip = &v
	return s
}

func (s *ModifyIpCnameStatusRequest) SetEnable(v bool) *ModifyIpCnameStatusRequest {
	s.Enable = &v
	return s
}

type ModifyIpCnameStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpCnameStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpCnameStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpCnameStatusResponseBody) SetRequestId(v string) *ModifyIpCnameStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpCnameStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpCnameStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpCnameStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpCnameStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpCnameStatusResponse) SetHeaders(v map[string]*string) *ModifyIpCnameStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpCnameStatusResponse) SetBody(v *ModifyIpCnameStatusResponseBody) *ModifyIpCnameStatusResponse {
	s.Body = v
	return s
}

type ModifyPersistenceTimeOutRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ip         *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	FrontPort  *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	LbId       *string `json:"LbId,omitempty" xml:"LbId,omitempty"`
}

func (s ModifyPersistenceTimeOutRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPersistenceTimeOutRequest) GoString() string {
	return s.String()
}

func (s *ModifyPersistenceTimeOutRequest) SetSourceIp(v string) *ModifyPersistenceTimeOutRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPersistenceTimeOutRequest) SetLang(v string) *ModifyPersistenceTimeOutRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPersistenceTimeOutRequest) SetIp(v string) *ModifyPersistenceTimeOutRequest {
	s.Ip = &v
	return s
}

func (s *ModifyPersistenceTimeOutRequest) SetFrontPort(v int32) *ModifyPersistenceTimeOutRequest {
	s.FrontPort = &v
	return s
}

func (s *ModifyPersistenceTimeOutRequest) SetConfigJson(v string) *ModifyPersistenceTimeOutRequest {
	s.ConfigJson = &v
	return s
}

func (s *ModifyPersistenceTimeOutRequest) SetLbId(v string) *ModifyPersistenceTimeOutRequest {
	s.LbId = &v
	return s
}

type ModifyPersistenceTimeOutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPersistenceTimeOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPersistenceTimeOutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPersistenceTimeOutResponseBody) SetRequestId(v string) *ModifyPersistenceTimeOutResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPersistenceTimeOutResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPersistenceTimeOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPersistenceTimeOutResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPersistenceTimeOutResponse) GoString() string {
	return s.String()
}

func (s *ModifyPersistenceTimeOutResponse) SetHeaders(v map[string]*string) *ModifyPersistenceTimeOutResponse {
	s.Headers = v
	return s
}

func (s *ModifyPersistenceTimeOutResponse) SetBody(v *ModifyPersistenceTimeOutResponseBody) *ModifyPersistenceTimeOutResponse {
	s.Body = v
	return s
}

type ModifyRealServersRequest struct {
	SourceIp    *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type        *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Domain      *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Line        *string   `json:"Line,omitempty" xml:"Line,omitempty"`
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
}

func (s ModifyRealServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealServersRequest) GoString() string {
	return s.String()
}

func (s *ModifyRealServersRequest) SetSourceIp(v string) *ModifyRealServersRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyRealServersRequest) SetLang(v string) *ModifyRealServersRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRealServersRequest) SetType(v string) *ModifyRealServersRequest {
	s.Type = &v
	return s
}

func (s *ModifyRealServersRequest) SetDomain(v string) *ModifyRealServersRequest {
	s.Domain = &v
	return s
}

func (s *ModifyRealServersRequest) SetLine(v string) *ModifyRealServersRequest {
	s.Line = &v
	return s
}

func (s *ModifyRealServersRequest) SetRealServers(v []*string) *ModifyRealServersRequest {
	s.RealServers = v
	return s
}

type ModifyRealServersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRealServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealServersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRealServersResponseBody) SetRequestId(v string) *ModifyRealServersResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRealServersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRealServersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRealServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRealServersResponse) GoString() string {
	return s.String()
}

func (s *ModifyRealServersResponse) SetHeaders(v map[string]*string) *ModifyRealServersResponse {
	s.Headers = v
	return s
}

func (s *ModifyRealServersResponse) SetBody(v *ModifyRealServersResponseBody) *ModifyRealServersResponse {
	s.Body = v
	return s
}

type ModifyTransmitLineRequest struct {
	SourceIp *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Domain   *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ips      []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s ModifyTransmitLineRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTransmitLineRequest) GoString() string {
	return s.String()
}

func (s *ModifyTransmitLineRequest) SetSourceIp(v string) *ModifyTransmitLineRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyTransmitLineRequest) SetLang(v string) *ModifyTransmitLineRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTransmitLineRequest) SetDomain(v string) *ModifyTransmitLineRequest {
	s.Domain = &v
	return s
}

func (s *ModifyTransmitLineRequest) SetIps(v []*string) *ModifyTransmitLineRequest {
	s.Ips = v
	return s
}

type ModifyTransmitLineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTransmitLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTransmitLineResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTransmitLineResponseBody) SetRequestId(v string) *ModifyTransmitLineResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTransmitLineResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTransmitLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTransmitLineResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTransmitLineResponse) GoString() string {
	return s.String()
}

func (s *ModifyTransmitLineResponse) SetHeaders(v map[string]*string) *ModifyTransmitLineResponse {
	s.Headers = v
	return s
}

func (s *ModifyTransmitLineResponse) SetBody(v *ModifyTransmitLineResponseBody) *ModifyTransmitLineResponse {
	s.Body = v
	return s
}

type UpdateCcCustomedRuleRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MatchingRule *string `json:"MatchingRule,omitempty" xml:"MatchingRule,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	VisitCount   *int32  `json:"VisitCount,omitempty" xml:"VisitCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BlockingType *string `json:"BlockingType,omitempty" xml:"BlockingType,omitempty"`
	Interval     *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	BlockingTime *int32  `json:"BlockingTime,omitempty" xml:"BlockingTime,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateCcCustomedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcCustomedRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCcCustomedRuleRequest) SetSourceIp(v string) *UpdateCcCustomedRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetLang(v string) *UpdateCcCustomedRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetMatchingRule(v string) *UpdateCcCustomedRuleRequest {
	s.MatchingRule = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetDomain(v string) *UpdateCcCustomedRuleRequest {
	s.Domain = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetVisitCount(v int32) *UpdateCcCustomedRuleRequest {
	s.VisitCount = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetName(v string) *UpdateCcCustomedRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetBlockingType(v string) *UpdateCcCustomedRuleRequest {
	s.BlockingType = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetInterval(v int32) *UpdateCcCustomedRuleRequest {
	s.Interval = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetBlockingTime(v int32) *UpdateCcCustomedRuleRequest {
	s.BlockingTime = &v
	return s
}

func (s *UpdateCcCustomedRuleRequest) SetUri(v string) *UpdateCcCustomedRuleRequest {
	s.Uri = &v
	return s
}

type UpdateCcCustomedRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCcCustomedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcCustomedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCcCustomedRuleResponseBody) SetRequestId(v string) *UpdateCcCustomedRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCcCustomedRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCcCustomedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCcCustomedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCcCustomedRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCcCustomedRuleResponse) SetHeaders(v map[string]*string) *UpdateCcCustomedRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCcCustomedRuleResponse) SetBody(v *UpdateCcCustomedRuleResponseBody) *UpdateCcCustomedRuleResponse {
	s.Body = v
	return s
}

type UpdatePortRuleRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	FrontPort      *int32  `json:"FrontPort,omitempty" xml:"FrontPort,omitempty"`
	RealServerList *string `json:"RealServerList,omitempty" xml:"RealServerList,omitempty"`
	Ip             *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s UpdatePortRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePortRuleRequest) SetSourceIp(v string) *UpdatePortRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdatePortRuleRequest) SetLang(v string) *UpdatePortRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdatePortRuleRequest) SetFrontPort(v int32) *UpdatePortRuleRequest {
	s.FrontPort = &v
	return s
}

func (s *UpdatePortRuleRequest) SetRealServerList(v string) *UpdatePortRuleRequest {
	s.RealServerList = &v
	return s
}

func (s *UpdatePortRuleRequest) SetIp(v string) *UpdatePortRuleRequest {
	s.Ip = &v
	return s
}

type UpdatePortRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePortRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePortRuleResponseBody) SetRequestId(v string) *UpdatePortRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePortRuleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePortRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePortRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePortRuleResponse) SetHeaders(v map[string]*string) *UpdatePortRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePortRuleResponse) SetBody(v *UpdatePortRuleResponseBody) *UpdatePortRuleResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddospro"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ConfigSwitchPriorityWithOptions(request *ConfigSwitchPriorityRequest, runtime *util.RuntimeOptions) (_result *ConfigSwitchPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigSwitchPriorityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigSwitchPriority"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigSwitchPriority(request *ConfigSwitchPriorityRequest) (_result *ConfigSwitchPriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigSwitchPriorityResponse{}
	_body, _err := client.ConfigSwitchPriorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCcCustomedRuleWithOptions(request *CreateCcCustomedRuleRequest, runtime *util.RuntimeOptions) (_result *CreateCcCustomedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCcCustomedRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCcCustomedRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCcCustomedRule(request *CreateCcCustomedRuleRequest) (_result *CreateCcCustomedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCcCustomedRuleResponse{}
	_body, _err := client.CreateCcCustomedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDomain"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreatePortRuleWithOptions(request *CreatePortRuleRequest, runtime *util.RuntimeOptions) (_result *CreatePortRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePortRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePortRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePortRule(request *CreatePortRuleRequest) (_result *CreatePortRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePortRuleResponse{}
	_body, _err := client.CreatePortRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransmitLineWithOptions(request *CreateTransmitLineRequest, runtime *util.RuntimeOptions) (_result *CreateTransmitLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTransmitLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTransmitLine"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransmitLine(request *CreateTransmitLineRequest) (_result *CreateTransmitLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransmitLineResponse{}
	_body, _err := client.CreateTransmitLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCcCustomedRuleWithOptions(request *DeleteCcCustomedRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteCcCustomedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCcCustomedRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCcCustomedRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCcCustomedRule(request *DeleteCcCustomedRuleRequest) (_result *DeleteCcCustomedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCcCustomedRuleResponse{}
	_body, _err := client.DeleteCcCustomedRuleWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDomain"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeletePortRuleWithOptions(request *DeletePortRuleRequest, runtime *util.RuntimeOptions) (_result *DeletePortRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePortRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePortRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePortRule(request *DeletePortRuleRequest) (_result *DeletePortRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePortRuleResponse{}
	_body, _err := client.DeletePortRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTransmitLineWithOptions(request *DeleteTransmitLineRequest, runtime *util.RuntimeOptions) (_result *DeleteTransmitLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTransmitLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTransmitLine"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTransmitLine(request *DeleteTransmitLineRequest) (_result *DeleteTransmitLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTransmitLineResponse{}
	_body, _err := client.DeleteTransmitLineWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackSourceCidr"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeBizFlowWithOptions(request *DescribeBizFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeBizFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBizFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBizFlow"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizFlow(request *DescribeBizFlowRequest) (_result *DescribeBizFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizFlowResponse{}
	_body, _err := client.DescribeBizFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCcEventsWithOptions(request *DescribeCcEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeCcEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCcEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCcEvents"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCcEvents(request *DescribeCcEventsRequest) (_result *DescribeCcEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCcEventsResponse{}
	_body, _err := client.DescribeCcEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCnameAutoStatusWithOptions(request *DescribeCnameAutoStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCnameAutoStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCnameAutoStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCnameAutoStatus"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCnameAutoStatus(request *DescribeCnameAutoStatusRequest) (_result *DescribeCnameAutoStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCnameAutoStatusResponse{}
	_body, _err := client.DescribeCnameAutoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosAttackEventsWithOptions(request *DescribeDdosAttackEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosAttackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosAttackEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosAttackEvents"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosAttackEvents(request *DescribeDdosAttackEventsRequest) (_result *DescribeDdosAttackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosAttackEventsResponse{}
	_body, _err := client.DescribeDdosAttackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosAttackEventSourceIpsWithOptions(request *DescribeDdosAttackEventSourceIpsRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosAttackEventSourceIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosAttackEventSourceIpsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosAttackEventSourceIps"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosAttackEventSourceIps(request *DescribeDdosAttackEventSourceIpsRequest) (_result *DescribeDdosAttackEventSourceIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosAttackEventSourceIpsResponse{}
	_body, _err := client.DescribeDdosAttackEventSourceIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosAttackTypeChartWithOptions(request *DescribeDdosAttackTypeChartRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosAttackTypeChartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosAttackTypeChartResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosAttackTypeChart"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosAttackTypeChart(request *DescribeDdosAttackTypeChartRequest) (_result *DescribeDdosAttackTypeChartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosAttackTypeChartResponse{}
	_body, _err := client.DescribeDdosAttackTypeChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosFlowProportionDiagramWithOptions(request *DescribeDdosFlowProportionDiagramRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosFlowProportionDiagramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosFlowProportionDiagramResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosFlowProportionDiagram"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosFlowProportionDiagram(request *DescribeDdosFlowProportionDiagramRequest) (_result *DescribeDdosFlowProportionDiagramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosFlowProportionDiagramResponse{}
	_body, _err := client.DescribeDdosFlowProportionDiagramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosIpConfigWithOptions(request *DescribeDdosIpConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosIpConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosIpConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosIpConfig"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosIpConfig(request *DescribeDdosIpConfigRequest) (_result *DescribeDdosIpConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosIpConfigResponse{}
	_body, _err := client.DescribeDdosIpConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosPeakFlowWithOptions(request *DescribeDdosPeakFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosPeakFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosPeakFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosPeakFlow"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosPeakFlow(request *DescribeDdosPeakFlowRequest) (_result *DescribeDdosPeakFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosPeakFlowResponse{}
	_body, _err := client.DescribeDdosPeakFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainConfigPageWithOptions(request *DescribeDomainConfigPageRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainConfigPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainConfigPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainConfigPage"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainConfigPage(request *DescribeDomainConfigPageRequest) (_result *DescribeDomainConfigPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainConfigPageResponse{}
	_body, _err := client.DescribeDomainConfigPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainSecurityConfigWithOptions(request *DescribeDomainSecurityConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainSecurityConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainSecurityConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainSecurityConfig"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainSecurityConfig(request *DescribeDomainSecurityConfigRequest) (_result *DescribeDomainSecurityConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainSecurityConfigResponse{}
	_body, _err := client.DescribeDomainSecurityConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHealthCheckConfigWithOptions(request *DescribeHealthCheckConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeHealthCheckConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHealthCheckConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHealthCheckConfig"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHealthCheckConfig(request *DescribeHealthCheckConfigRequest) (_result *DescribeHealthCheckConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHealthCheckConfigResponse{}
	_body, _err := client.DescribeHealthCheckConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancePageWithOptions(request *DescribeInstancePageRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancePageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstancePage"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstancePage(request *DescribeInstancePageRequest) (_result *DescribeInstancePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancePageResponse{}
	_body, _err := client.DescribeInstancePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePortRulePageWithOptions(request *DescribePortRulePageRequest, runtime *util.RuntimeOptions) (_result *DescribePortRulePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePortRulePageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePortRulePage"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePortRulePage(request *DescribePortRulePageRequest) (_result *DescribePortRulePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePortRulePageResponse{}
	_body, _err := client.DescribePortRulePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCcCustomedRuleWithOptions(request *ListCcCustomedRuleRequest, runtime *util.RuntimeOptions) (_result *ListCcCustomedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCcCustomedRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCcCustomedRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCcCustomedRule(request *ListCcCustomedRuleRequest) (_result *ListCcCustomedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCcCustomedRuleResponse{}
	_body, _err := client.ListCcCustomedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCcCustomStatusWithOptions(request *ModifyCcCustomStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyCcCustomStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCcCustomStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCcCustomStatus"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCcCustomStatus(request *ModifyCcCustomStatusRequest) (_result *ModifyCcCustomStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCcCustomStatusResponse{}
	_body, _err := client.ModifyCcCustomStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCcStatusWithOptions(request *ModifyCcStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyCcStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCcStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCcStatus"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCcStatus(request *ModifyCcStatusRequest) (_result *ModifyCcStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCcStatusResponse{}
	_body, _err := client.ModifyCcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCcTemplateWithOptions(request *ModifyCcTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyCcTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCcTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCcTemplate"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCcTemplate(request *ModifyCcTemplateRequest) (_result *ModifyCcTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCcTemplateResponse{}
	_body, _err := client.ModifyCcTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDDoSProtectConfigWithOptions(request *ModifyDDoSProtectConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDDoSProtectConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDDoSProtectConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDDoSProtectConfig"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDDoSProtectConfig(request *ModifyDDoSProtectConfigRequest) (_result *ModifyDDoSProtectConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDDoSProtectConfigResponse{}
	_body, _err := client.ModifyDDoSProtectConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainBlackWhiteListWithOptions(request *ModifyDomainBlackWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainBlackWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDomainBlackWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDomainBlackWhiteList"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainBlackWhiteList(request *ModifyDomainBlackWhiteListRequest) (_result *ModifyDomainBlackWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainBlackWhiteListResponse{}
	_body, _err := client.ModifyDomainBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDomainProxyWithOptions(request *ModifyDomainProxyRequest, runtime *util.RuntimeOptions) (_result *ModifyDomainProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDomainProxyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDomainProxy"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDomainProxy(request *ModifyDomainProxyRequest) (_result *ModifyDomainProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDomainProxyResponse{}
	_body, _err := client.ModifyDomainProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyElasticBandwidthWithOptions(request *ModifyElasticBandwidthRequest, runtime *util.RuntimeOptions) (_result *ModifyElasticBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyElasticBandwidthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyElasticBandwidth"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyElasticBandwidth(request *ModifyElasticBandwidthRequest) (_result *ModifyElasticBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyElasticBandwidthResponse{}
	_body, _err := client.ModifyElasticBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHealthCheckConfigWithOptions(request *ModifyHealthCheckConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyHealthCheckConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHealthCheckConfig"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHealthCheckConfig(request *ModifyHealthCheckConfigRequest) (_result *ModifyHealthCheckConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.ModifyHealthCheckConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpCnameStatusWithOptions(request *ModifyIpCnameStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyIpCnameStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIpCnameStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIpCnameStatus"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpCnameStatus(request *ModifyIpCnameStatusRequest) (_result *ModifyIpCnameStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpCnameStatusResponse{}
	_body, _err := client.ModifyIpCnameStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPersistenceTimeOutWithOptions(request *ModifyPersistenceTimeOutRequest, runtime *util.RuntimeOptions) (_result *ModifyPersistenceTimeOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPersistenceTimeOutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPersistenceTimeOut"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPersistenceTimeOut(request *ModifyPersistenceTimeOutRequest) (_result *ModifyPersistenceTimeOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPersistenceTimeOutResponse{}
	_body, _err := client.ModifyPersistenceTimeOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRealServersWithOptions(request *ModifyRealServersRequest, runtime *util.RuntimeOptions) (_result *ModifyRealServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRealServersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRealServers"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRealServers(request *ModifyRealServersRequest) (_result *ModifyRealServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRealServersResponse{}
	_body, _err := client.ModifyRealServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTransmitLineWithOptions(request *ModifyTransmitLineRequest, runtime *util.RuntimeOptions) (_result *ModifyTransmitLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTransmitLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTransmitLine"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTransmitLine(request *ModifyTransmitLineRequest) (_result *ModifyTransmitLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTransmitLineResponse{}
	_body, _err := client.ModifyTransmitLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCcCustomedRuleWithOptions(request *UpdateCcCustomedRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateCcCustomedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCcCustomedRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCcCustomedRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCcCustomedRule(request *UpdateCcCustomedRuleRequest) (_result *UpdateCcCustomedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCcCustomedRuleResponse{}
	_body, _err := client.UpdateCcCustomedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePortRuleWithOptions(request *UpdatePortRuleRequest, runtime *util.RuntimeOptions) (_result *UpdatePortRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePortRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePortRule"), tea.String("2017-07-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePortRule(request *UpdatePortRuleRequest) (_result *UpdatePortRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePortRuleResponse{}
	_body, _err := client.UpdatePortRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
