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

type AddResolverEndpointRequest struct {
	IpConfig        []*AddResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	Lang            *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name            *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityGroupId *string                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VpcId           *string                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcRegionId     *string                               `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s AddResolverEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointRequest) SetIpConfig(v []*AddResolverEndpointRequestIpConfig) *AddResolverEndpointRequest {
	s.IpConfig = v
	return s
}

func (s *AddResolverEndpointRequest) SetLang(v string) *AddResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *AddResolverEndpointRequest) SetName(v string) *AddResolverEndpointRequest {
	s.Name = &v
	return s
}

func (s *AddResolverEndpointRequest) SetSecurityGroupId(v string) *AddResolverEndpointRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AddResolverEndpointRequest) SetVpcId(v string) *AddResolverEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *AddResolverEndpointRequest) SetVpcRegionId(v string) *AddResolverEndpointRequest {
	s.VpcRegionId = &v
	return s
}

type AddResolverEndpointRequestIpConfig struct {
	AzId      *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s AddResolverEndpointRequestIpConfig) String() string {
	return tea.Prettify(s)
}

func (s AddResolverEndpointRequestIpConfig) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointRequestIpConfig) SetAzId(v string) *AddResolverEndpointRequestIpConfig {
	s.AzId = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetCidrBlock(v string) *AddResolverEndpointRequestIpConfig {
	s.CidrBlock = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetIp(v string) *AddResolverEndpointRequestIpConfig {
	s.Ip = &v
	return s
}

func (s *AddResolverEndpointRequestIpConfig) SetVSwitchId(v string) *AddResolverEndpointRequestIpConfig {
	s.VSwitchId = &v
	return s
}

type AddResolverEndpointResponseBody struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddResolverEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointResponseBody) SetEndpointId(v string) *AddResolverEndpointResponseBody {
	s.EndpointId = &v
	return s
}

func (s *AddResolverEndpointResponseBody) SetRequestId(v string) *AddResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

type AddResolverEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddResolverEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *AddResolverEndpointResponse) SetHeaders(v map[string]*string) *AddResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *AddResolverEndpointResponse) SetStatusCode(v int32) *AddResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResolverEndpointResponse) SetBody(v *AddResolverEndpointResponseBody) *AddResolverEndpointResponse {
	s.Body = v
	return s
}

type AddResolverRuleRequest struct {
	EndpointId *string                            `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	ForwardIp  []*AddResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	Lang       *string                            `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name       *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneName   *string                            `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s AddResolverRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequest) SetEndpointId(v string) *AddResolverRuleRequest {
	s.EndpointId = &v
	return s
}

func (s *AddResolverRuleRequest) SetForwardIp(v []*AddResolverRuleRequestForwardIp) *AddResolverRuleRequest {
	s.ForwardIp = v
	return s
}

func (s *AddResolverRuleRequest) SetLang(v string) *AddResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *AddResolverRuleRequest) SetName(v string) *AddResolverRuleRequest {
	s.Name = &v
	return s
}

func (s *AddResolverRuleRequest) SetType(v string) *AddResolverRuleRequest {
	s.Type = &v
	return s
}

func (s *AddResolverRuleRequest) SetZoneName(v string) *AddResolverRuleRequest {
	s.ZoneName = &v
	return s
}

type AddResolverRuleRequestForwardIp struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s AddResolverRuleRequestForwardIp) String() string {
	return tea.Prettify(s)
}

func (s AddResolverRuleRequestForwardIp) GoString() string {
	return s.String()
}

func (s *AddResolverRuleRequestForwardIp) SetIp(v string) *AddResolverRuleRequestForwardIp {
	s.Ip = &v
	return s
}

func (s *AddResolverRuleRequestForwardIp) SetPort(v int32) *AddResolverRuleRequestForwardIp {
	s.Port = &v
	return s
}

type AddResolverRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddResolverRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddResolverRuleResponseBody) SetRequestId(v string) *AddResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddResolverRuleResponseBody) SetRuleId(v string) *AddResolverRuleResponseBody {
	s.RuleId = &v
	return s
}

type AddResolverRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddResolverRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *AddResolverRuleResponse) SetHeaders(v map[string]*string) *AddResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *AddResolverRuleResponse) SetStatusCode(v int32) *AddResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResolverRuleResponse) SetBody(v *AddResolverRuleResponseBody) *AddResolverRuleResponse {
	s.Body = v
	return s
}

type AddUserVpcAuthorizationRequest struct {
	AuthChannel      *string `json:"AuthChannel,omitempty" xml:"AuthChannel,omitempty"`
	AuthCode         *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	AuthType         *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
}

func (s AddUserVpcAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserVpcAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationRequest) SetAuthChannel(v string) *AddUserVpcAuthorizationRequest {
	s.AuthChannel = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthCode(v string) *AddUserVpcAuthorizationRequest {
	s.AuthCode = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthType(v string) *AddUserVpcAuthorizationRequest {
	s.AuthType = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthorizedUserId(v int64) *AddUserVpcAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

type AddUserVpcAuthorizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserVpcAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserVpcAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationResponseBody) SetRequestId(v string) *AddUserVpcAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type AddUserVpcAuthorizationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserVpcAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserVpcAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserVpcAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationResponse) SetHeaders(v map[string]*string) *AddUserVpcAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *AddUserVpcAuthorizationResponse) SetStatusCode(v int32) *AddUserVpcAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserVpcAuthorizationResponse) SetBody(v *AddUserVpcAuthorizationResponseBody) *AddUserVpcAuthorizationResponse {
	s.Body = v
	return s
}

type AddZoneRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DnsGroup        *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProxyPattern    *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneName        *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneTag         *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	ZoneType        *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s AddZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRequest) SetClientToken(v string) *AddZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneRequest) SetDnsGroup(v string) *AddZoneRequest {
	s.DnsGroup = &v
	return s
}

func (s *AddZoneRequest) SetLang(v string) *AddZoneRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRequest) SetProxyPattern(v string) *AddZoneRequest {
	s.ProxyPattern = &v
	return s
}

func (s *AddZoneRequest) SetResourceGroupId(v string) *AddZoneRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddZoneRequest) SetZoneName(v string) *AddZoneRequest {
	s.ZoneName = &v
	return s
}

func (s *AddZoneRequest) SetZoneTag(v string) *AddZoneRequest {
	s.ZoneTag = &v
	return s
}

func (s *AddZoneRequest) SetZoneType(v string) *AddZoneRequest {
	s.ZoneType = &v
	return s
}

type AddZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// zone ID。
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s AddZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddZoneResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneResponseBody) SetRequestId(v string) *AddZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneResponseBody) SetSuccess(v bool) *AddZoneResponseBody {
	s.Success = &v
	return s
}

func (s *AddZoneResponseBody) SetZoneId(v string) *AddZoneResponseBody {
	s.ZoneId = &v
	return s
}

func (s *AddZoneResponseBody) SetZoneName(v string) *AddZoneResponseBody {
	s.ZoneName = &v
	return s
}

type AddZoneResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s AddZoneResponse) GoString() string {
	return s.String()
}

func (s *AddZoneResponse) SetHeaders(v map[string]*string) *AddZoneResponse {
	s.Headers = v
	return s
}

func (s *AddZoneResponse) SetStatusCode(v int32) *AddZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneResponse) SetBody(v *AddZoneResponseBody) *AddZoneResponse {
	s.Body = v
	return s
}

type AddZoneRecordRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Rr           *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// Zone ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRecordRequest) SetClientToken(v string) *AddZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneRecordRequest) SetLang(v string) *AddZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRecordRequest) SetLine(v string) *AddZoneRecordRequest {
	s.Line = &v
	return s
}

func (s *AddZoneRecordRequest) SetPriority(v int32) *AddZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddZoneRecordRequest) SetRemark(v string) *AddZoneRecordRequest {
	s.Remark = &v
	return s
}

func (s *AddZoneRecordRequest) SetRr(v string) *AddZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *AddZoneRecordRequest) SetTtl(v int32) *AddZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *AddZoneRecordRequest) SetType(v string) *AddZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *AddZoneRecordRequest) SetUserClientIp(v string) *AddZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddZoneRecordRequest) SetValue(v string) *AddZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *AddZoneRecordRequest) SetWeight(v int32) *AddZoneRecordRequest {
	s.Weight = &v
	return s
}

func (s *AddZoneRecordRequest) SetZoneId(v string) *AddZoneRecordRequest {
	s.ZoneId = &v
	return s
}

type AddZoneRecordResponseBody struct {
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponseBody) SetRecordId(v int64) *AddZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetRequestId(v string) *AddZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetSuccess(v bool) *AddZoneRecordResponseBody {
	s.Success = &v
	return s
}

type AddZoneRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponse) SetHeaders(v map[string]*string) *AddZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *AddZoneRecordResponse) SetStatusCode(v int32) *AddZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneRecordResponse) SetBody(v *AddZoneRecordResponseBody) *AddZoneRecordResponse {
	s.Body = v
	return s
}

type BindResolverRuleVpcRequest struct {
	Lang   *string                          `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RuleId *string                          `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Vpc    []*BindResolverRuleVpcRequestVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s BindResolverRuleVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s BindResolverRuleVpcRequest) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcRequest) SetLang(v string) *BindResolverRuleVpcRequest {
	s.Lang = &v
	return s
}

func (s *BindResolverRuleVpcRequest) SetRuleId(v string) *BindResolverRuleVpcRequest {
	s.RuleId = &v
	return s
}

func (s *BindResolverRuleVpcRequest) SetVpc(v []*BindResolverRuleVpcRequestVpc) *BindResolverRuleVpcRequest {
	s.Vpc = v
	return s
}

type BindResolverRuleVpcRequestVpc struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// vpcID
	VpcId   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s BindResolverRuleVpcRequestVpc) String() string {
	return tea.Prettify(s)
}

func (s BindResolverRuleVpcRequestVpc) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcRequestVpc) SetRegionId(v string) *BindResolverRuleVpcRequestVpc {
	s.RegionId = &v
	return s
}

func (s *BindResolverRuleVpcRequestVpc) SetVpcId(v string) *BindResolverRuleVpcRequestVpc {
	s.VpcId = &v
	return s
}

func (s *BindResolverRuleVpcRequestVpc) SetVpcType(v string) *BindResolverRuleVpcRequestVpc {
	s.VpcType = &v
	return s
}

type BindResolverRuleVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindResolverRuleVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindResolverRuleVpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcResponseBody) SetRequestId(v string) *BindResolverRuleVpcResponseBody {
	s.RequestId = &v
	return s
}

type BindResolverRuleVpcResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindResolverRuleVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindResolverRuleVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s BindResolverRuleVpcResponse) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcResponse) SetHeaders(v map[string]*string) *BindResolverRuleVpcResponse {
	s.Headers = v
	return s
}

func (s *BindResolverRuleVpcResponse) SetStatusCode(v int32) *BindResolverRuleVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindResolverRuleVpcResponse) SetBody(v *BindResolverRuleVpcResponseBody) *BindResolverRuleVpcResponse {
	s.Body = v
	return s
}

type BindZoneVpcRequest struct {
	ClientToken  *string                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string                   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Vpcs         []*BindZoneVpcRequestVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
	ZoneId       *string                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s BindZoneVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcRequest) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequest) SetClientToken(v string) *BindZoneVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *BindZoneVpcRequest) SetLang(v string) *BindZoneVpcRequest {
	s.Lang = &v
	return s
}

func (s *BindZoneVpcRequest) SetUserClientIp(v string) *BindZoneVpcRequest {
	s.UserClientIp = &v
	return s
}

func (s *BindZoneVpcRequest) SetVpcs(v []*BindZoneVpcRequestVpcs) *BindZoneVpcRequest {
	s.Vpcs = v
	return s
}

func (s *BindZoneVpcRequest) SetZoneId(v string) *BindZoneVpcRequest {
	s.ZoneId = &v
	return s
}

type BindZoneVpcRequestVpcs struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType  *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s BindZoneVpcRequestVpcs) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcRequestVpcs) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequestVpcs) SetRegionId(v string) *BindZoneVpcRequestVpcs {
	s.RegionId = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) SetVpcId(v string) *BindZoneVpcRequestVpcs {
	s.VpcId = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) SetVpcType(v string) *BindZoneVpcRequestVpcs {
	s.VpcType = &v
	return s
}

type BindZoneVpcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindZoneVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcResponseBody) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponseBody) SetRequestId(v string) *BindZoneVpcResponseBody {
	s.RequestId = &v
	return s
}

type BindZoneVpcResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindZoneVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindZoneVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s BindZoneVpcResponse) GoString() string {
	return s.String()
}

func (s *BindZoneVpcResponse) SetHeaders(v map[string]*string) *BindZoneVpcResponse {
	s.Headers = v
	return s
}

func (s *BindZoneVpcResponse) SetStatusCode(v int32) *BindZoneVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindZoneVpcResponse) SetBody(v *BindZoneVpcResponseBody) *BindZoneVpcResponse {
	s.Body = v
	return s
}

type CheckZoneNameRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZoneName     *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s CheckZoneNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameRequest) GoString() string {
	return s.String()
}

func (s *CheckZoneNameRequest) SetLang(v string) *CheckZoneNameRequest {
	s.Lang = &v
	return s
}

func (s *CheckZoneNameRequest) SetUserClientIp(v string) *CheckZoneNameRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckZoneNameRequest) SetZoneName(v string) *CheckZoneNameRequest {
	s.ZoneName = &v
	return s
}

type CheckZoneNameResponseBody struct {
	Check     *bool   `json:"Check,omitempty" xml:"Check,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckZoneNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponseBody) SetCheck(v bool) *CheckZoneNameResponseBody {
	s.Check = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetRequestId(v string) *CheckZoneNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckZoneNameResponseBody) SetSuccess(v bool) *CheckZoneNameResponseBody {
	s.Success = &v
	return s
}

type CheckZoneNameResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckZoneNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckZoneNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckZoneNameResponse) GoString() string {
	return s.String()
}

func (s *CheckZoneNameResponse) SetHeaders(v map[string]*string) *CheckZoneNameResponse {
	s.Headers = v
	return s
}

func (s *CheckZoneNameResponse) SetStatusCode(v int32) *CheckZoneNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckZoneNameResponse) SetBody(v *CheckZoneNameResponseBody) *CheckZoneNameResponse {
	s.Body = v
	return s
}

type DeleteResolverEndpointRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteResolverEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointRequest) SetEndpointId(v string) *DeleteResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteResolverEndpointRequest) SetLang(v string) *DeleteResolverEndpointRequest {
	s.Lang = &v
	return s
}

type DeleteResolverEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResolverEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointResponseBody) SetRequestId(v string) *DeleteResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResolverEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResolverEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteResolverEndpointResponse) SetHeaders(v map[string]*string) *DeleteResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteResolverEndpointResponse) SetStatusCode(v int32) *DeleteResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResolverEndpointResponse) SetBody(v *DeleteResolverEndpointResponseBody) *DeleteResolverEndpointResponse {
	s.Body = v
	return s
}

type DeleteResolverRuleRequest struct {
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteResolverRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleRequest) SetLang(v string) *DeleteResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteResolverRuleRequest) SetRuleId(v string) *DeleteResolverRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteResolverRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResolverRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleResponseBody) SetRequestId(v string) *DeleteResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResolverRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResolverRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleResponse) SetHeaders(v map[string]*string) *DeleteResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteResolverRuleResponse) SetStatusCode(v int32) *DeleteResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResolverRuleResponse) SetBody(v *DeleteResolverRuleResponseBody) *DeleteResolverRuleResponse {
	s.Body = v
	return s
}

type DeleteUserVpcAuthorizationRequest struct {
	AuthType         *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
}

func (s DeleteUserVpcAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserVpcAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationRequest) SetAuthType(v string) *DeleteUserVpcAuthorizationRequest {
	s.AuthType = &v
	return s
}

func (s *DeleteUserVpcAuthorizationRequest) SetAuthorizedUserId(v int64) *DeleteUserVpcAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

type DeleteUserVpcAuthorizationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserVpcAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserVpcAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationResponseBody) SetRequestId(v string) *DeleteUserVpcAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserVpcAuthorizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserVpcAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserVpcAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserVpcAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationResponse) SetHeaders(v map[string]*string) *DeleteUserVpcAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserVpcAuthorizationResponse) SetStatusCode(v int32) *DeleteUserVpcAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserVpcAuthorizationResponse) SetBody(v *DeleteUserVpcAuthorizationResponseBody) *DeleteUserVpcAuthorizationResponse {
	s.Body = v
	return s
}

type DeleteZoneRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// zone ID
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRequest) SetClientToken(v string) *DeleteZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteZoneRequest) SetLang(v string) *DeleteZoneRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRequest) SetUserClientIp(v string) *DeleteZoneRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteZoneRequest) SetZoneId(v string) *DeleteZoneRequest {
	s.ZoneId = &v
	return s
}

type DeleteZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// zone ID
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponseBody) SetRequestId(v string) *DeleteZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZoneResponseBody) SetZoneId(v string) *DeleteZoneResponseBody {
	s.ZoneId = &v
	return s
}

type DeleteZoneResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneResponse) SetHeaders(v map[string]*string) *DeleteZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneResponse) SetStatusCode(v int32) *DeleteZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteZoneResponse) SetBody(v *DeleteZoneResponseBody) *DeleteZoneResponse {
	s.Body = v
	return s
}

type DeleteZoneRecordRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordRequest) SetClientToken(v string) *DeleteZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetLang(v string) *DeleteZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetRecordId(v int64) *DeleteZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetUserClientIp(v string) *DeleteZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

type DeleteZoneRecordResponseBody struct {
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponseBody) SetRecordId(v int64) *DeleteZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *DeleteZoneRecordResponseBody) SetRequestId(v string) *DeleteZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

type DeleteZoneRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponse) SetHeaders(v map[string]*string) *DeleteZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *DeleteZoneRecordResponse) SetStatusCode(v int32) *DeleteZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteZoneRecordResponse) SetBody(v *DeleteZoneRecordResponseBody) *DeleteZoneRecordResponse {
	s.Body = v
	return s
}

type DescribeChangeLogsRequest struct {
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	EntityType     *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Keyword        *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeChangeLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsRequest) SetEndTimestamp(v int64) *DescribeChangeLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetEntityType(v string) *DescribeChangeLogsRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetKeyword(v string) *DescribeChangeLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetLang(v string) *DescribeChangeLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageNumber(v int32) *DescribeChangeLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageSize(v int32) *DescribeChangeLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetStartTimestamp(v int64) *DescribeChangeLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetUserClientIp(v string) *DescribeChangeLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetZoneId(v string) *DescribeChangeLogsRequest {
	s.ZoneId = &v
	return s
}

type DescribeChangeLogsResponseBody struct {
	ChangeLogs *DescribeChangeLogsResponseBodyChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Struct"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems *int32                                    `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                                    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeChangeLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBody) SetChangeLogs(v *DescribeChangeLogsResponseBodyChangeLogs) *DescribeChangeLogsResponseBody {
	s.ChangeLogs = v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetPageNumber(v int32) *DescribeChangeLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetPageSize(v int32) *DescribeChangeLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetRequestId(v string) *DescribeChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalItems(v int32) *DescribeChangeLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeChangeLogsResponseBody) SetTotalPages(v int32) *DescribeChangeLogsResponseBody {
	s.TotalPages = &v
	return s
}

type DescribeChangeLogsResponseBodyChangeLogs struct {
	ChangeLog []*DescribeChangeLogsResponseBodyChangeLogsChangeLog `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty" type:"Repeated"`
}

func (s DescribeChangeLogsResponseBodyChangeLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogs) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogs) SetChangeLog(v []*DescribeChangeLogsResponseBodyChangeLogsChangeLog) *DescribeChangeLogsResponseBodyChangeLogs {
	s.ChangeLog = v
	return s
}

type DescribeChangeLogsResponseBodyChangeLogsChangeLog struct {
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreatorId      *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	CreatorType    *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	EntityId       *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName     *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OperAction     *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	OperIp         *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	OperObject     *string `json:"OperObject,omitempty" xml:"OperObject,omitempty"`
	OperTime       *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	OperTimestamp  *int64  `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponseBodyChangeLogsChangeLog) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetContent(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Content = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorSubType(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetCreatorType(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.CreatorType = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityId(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityId = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetEntityName(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.EntityName = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetId(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.Id = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperAction(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperAction = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperIp(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperIp = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperObject(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperObject = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTime(v string) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTime = &v
	return s
}

func (s *DescribeChangeLogsResponseBodyChangeLogsChangeLog) SetOperTimestamp(v int64) *DescribeChangeLogsResponseBodyChangeLogsChangeLog {
	s.OperTimestamp = &v
	return s
}

type DescribeChangeLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChangeLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChangeLogsResponse) SetStatusCode(v int32) *DescribeChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChangeLogsResponse) SetBody(v *DescribeChangeLogsResponseBody) *DescribeChangeLogsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage   *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Scene            *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	UserClientIp     *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	VpcType          *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
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

func (s *DescribeRegionsRequest) SetAuthorizedUserId(v int64) *DescribeRegionsRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetScene(v string) *DescribeRegionsRequest {
	s.Scene = &v
	return s
}

func (s *DescribeRegionsRequest) SetUserClientIp(v string) *DescribeRegionsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetVpcType(v string) *DescribeRegionsRequest {
	s.VpcType = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRequestGraphRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	UserClientIp   *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// zone ID
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRequestGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphRequest) SetBizId(v string) *DescribeRequestGraphRequest {
	s.BizId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetBizType(v string) *DescribeRequestGraphRequest {
	s.BizType = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetEndTimestamp(v int64) *DescribeRequestGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetLang(v string) *DescribeRequestGraphRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetStartTimestamp(v int64) *DescribeRequestGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetUserClientIp(v string) *DescribeRequestGraphRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetVpcId(v string) *DescribeRequestGraphRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetZoneId(v string) *DescribeRequestGraphRequest {
	s.ZoneId = &v
	return s
}

type DescribeRequestGraphResponseBody struct {
	RequestDetails *DescribeRequestGraphResponseBodyRequestDetails `json:"RequestDetails,omitempty" xml:"RequestDetails,omitempty" type:"Struct"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRequestGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBody) SetRequestDetails(v *DescribeRequestGraphResponseBodyRequestDetails) *DescribeRequestGraphResponseBody {
	s.RequestDetails = v
	return s
}

func (s *DescribeRequestGraphResponseBody) SetRequestId(v string) *DescribeRequestGraphResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRequestGraphResponseBodyRequestDetails struct {
	ZoneRequestTop []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeRequestGraphResponseBodyRequestDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetails) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetails) SetZoneRequestTop(v []*DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) *DescribeRequestGraphResponseBodyRequestDetails {
	s.ZoneRequestTop = v
	return s
}

type DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop struct {
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	Time         *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Timestamp    *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetRequestCount(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTime(v string) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Time = &v
	return s
}

func (s *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop) SetTimestamp(v int64) *DescribeRequestGraphResponseBodyRequestDetailsZoneRequestTop {
	s.Timestamp = &v
	return s
}

type DescribeRequestGraphResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRequestGraphResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRequestGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRequestGraphResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphResponse) SetHeaders(v map[string]*string) *DescribeRequestGraphResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestGraphResponse) SetStatusCode(v int32) *DescribeRequestGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRequestGraphResponse) SetBody(v *DescribeRequestGraphResponseBody) *DescribeRequestGraphResponse {
	s.Body = v
	return s
}

type DescribeResolverAvailableZonesRequest struct {
	AzId             *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResolverRegionId *string `json:"ResolverRegionId,omitempty" xml:"ResolverRegionId,omitempty"`
}

func (s DescribeResolverAvailableZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesRequest) SetAzId(v string) *DescribeResolverAvailableZonesRequest {
	s.AzId = &v
	return s
}

func (s *DescribeResolverAvailableZonesRequest) SetLang(v string) *DescribeResolverAvailableZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverAvailableZonesRequest) SetResolverRegionId(v string) *DescribeResolverAvailableZonesRequest {
	s.ResolverRegionId = &v
	return s
}

type DescribeResolverAvailableZonesResponseBody struct {
	AvailableZones []*DescribeResolverAvailableZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResolverAvailableZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponseBody) SetAvailableZones(v []*DescribeResolverAvailableZonesResponseBodyAvailableZones) *DescribeResolverAvailableZonesResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBody) SetRequestId(v string) *DescribeResolverAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeResolverAvailableZonesResponseBodyAvailableZones struct {
	AzId   *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResolverAvailableZonesResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) SetAzId(v string) *DescribeResolverAvailableZonesResponseBodyAvailableZones {
	s.AzId = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponseBodyAvailableZones) SetStatus(v string) *DescribeResolverAvailableZonesResponseBodyAvailableZones {
	s.Status = &v
	return s
}

type DescribeResolverAvailableZonesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverAvailableZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponse) SetHeaders(v map[string]*string) *DescribeResolverAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverAvailableZonesResponse) SetStatusCode(v int32) *DescribeResolverAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponse) SetBody(v *DescribeResolverAvailableZonesResponseBody) *DescribeResolverAvailableZonesResponse {
	s.Body = v
	return s
}

type DescribeResolverEndpointRequest struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeResolverEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointRequest) SetEndpointId(v string) *DescribeResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverEndpointRequest) SetLang(v string) *DescribeResolverEndpointRequest {
	s.Lang = &v
	return s
}

type DescribeResolverEndpointResponseBody struct {
	CreateTime      *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                           `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Id              *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	IpConfigs       []*DescribeResolverEndpointResponseBodyIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	Name            *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroupId *string                                          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *string                                          `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                           `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	VpcId           *string                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName         *string                                          `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcRegionId     *string                                          `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	VpcRegionName   *string                                          `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
}

func (s DescribeResolverEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponseBody) SetCreateTime(v string) *DescribeResolverEndpointResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetCreateTimestamp(v int64) *DescribeResolverEndpointResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetId(v string) *DescribeResolverEndpointResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetIpConfigs(v []*DescribeResolverEndpointResponseBodyIpConfigs) *DescribeResolverEndpointResponseBody {
	s.IpConfigs = v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetName(v string) *DescribeResolverEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetRequestId(v string) *DescribeResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetSecurityGroupId(v string) *DescribeResolverEndpointResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetStatus(v string) *DescribeResolverEndpointResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetUpdateTime(v string) *DescribeResolverEndpointResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetUpdateTimestamp(v int64) *DescribeResolverEndpointResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcId(v string) *DescribeResolverEndpointResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcName(v string) *DescribeResolverEndpointResponseBody {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcRegionId(v string) *DescribeResolverEndpointResponseBody {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBody) SetVpcRegionName(v string) *DescribeResolverEndpointResponseBody {
	s.VpcRegionName = &v
	return s
}

type DescribeResolverEndpointResponseBodyIpConfigs struct {
	AzId      *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeResolverEndpointResponseBodyIpConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointResponseBodyIpConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetAzId(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.AzId = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetCidrBlock(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetIp(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.Ip = &v
	return s
}

func (s *DescribeResolverEndpointResponseBodyIpConfigs) SetVSwitchId(v string) *DescribeResolverEndpointResponseBodyIpConfigs {
	s.VSwitchId = &v
	return s
}

type DescribeResolverEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponse) SetHeaders(v map[string]*string) *DescribeResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverEndpointResponse) SetStatusCode(v int32) *DescribeResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverEndpointResponse) SetBody(v *DescribeResolverEndpointResponseBody) *DescribeResolverEndpointResponse {
	s.Body = v
	return s
}

type DescribeResolverEndpointsRequest struct {
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResolverEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsRequest) SetKeyword(v string) *DescribeResolverEndpointsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetLang(v string) *DescribeResolverEndpointsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetPageNumber(v int32) *DescribeResolverEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetPageSize(v int32) *DescribeResolverEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverEndpointsRequest) SetStatus(v string) *DescribeResolverEndpointsRequest {
	s.Status = &v
	return s
}

type DescribeResolverEndpointsResponseBody struct {
	Endpoints  []*DescribeResolverEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems *int32                                            `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeResolverEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBody) SetEndpoints(v []*DescribeResolverEndpointsResponseBodyEndpoints) *DescribeResolverEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetPageNumber(v int32) *DescribeResolverEndpointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetPageSize(v int32) *DescribeResolverEndpointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetRequestId(v string) *DescribeResolverEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetTotalItems(v int32) *DescribeResolverEndpointsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBody) SetTotalPages(v int32) *DescribeResolverEndpointsResponseBody {
	s.TotalPages = &v
	return s
}

type DescribeResolverEndpointsResponseBodyEndpoints struct {
	CreateTime      *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                                     `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Id              *string                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	IpConfigs       []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	Name            *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityGroupId *string                                                    `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status          *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime      *string                                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                                     `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	VpcId           *string                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName         *string                                                    `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcRegionId     *string                                                    `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	VpcRegionName   *string                                                    `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
}

func (s DescribeResolverEndpointsResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetCreateTime(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetCreateTimestamp(v int64) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Id = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetIpConfigs(v []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.IpConfigs = v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Name = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetSecurityGroupId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetStatus(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.Status = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetUpdateTime(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetUpdateTimestamp(v int64) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcRegionId(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpoints) SetVpcRegionName(v string) *DescribeResolverEndpointsResponseBodyEndpoints {
	s.VpcRegionName = &v
	return s
}

type DescribeResolverEndpointsResponseBodyEndpointsIpConfigs struct {
	AzId      *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetAzId(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.AzId = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetCidrBlock(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetIp(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.Ip = &v
	return s
}

func (s *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs) SetVSwitchId(v string) *DescribeResolverEndpointsResponseBodyEndpointsIpConfigs {
	s.VSwitchId = &v
	return s
}

type DescribeResolverEndpointsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponse) SetHeaders(v map[string]*string) *DescribeResolverEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverEndpointsResponse) SetStatusCode(v int32) *DescribeResolverEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverEndpointsResponse) SetBody(v *DescribeResolverEndpointsResponseBody) *DescribeResolverEndpointsResponse {
	s.Body = v
	return s
}

type DescribeResolverRuleRequest struct {
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeResolverRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleRequest) SetLang(v string) *DescribeResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverRuleRequest) SetRuleId(v string) *DescribeResolverRuleRequest {
	s.RuleId = &v
	return s
}

type DescribeResolverRuleResponseBody struct {
	BindVpcs        []*DescribeResolverRuleResponseBodyBindVpcs   `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	CreateTime      *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                        `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EndpointId      *string                                       `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName    *string                                       `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	ForwardIps      []*DescribeResolverRuleResponseBodyForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
	Id              *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type            *string                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime      *string                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                        `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	ZoneName        *string                                       `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResolverRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBody) SetBindVpcs(v []*DescribeResolverRuleResponseBodyBindVpcs) *DescribeResolverRuleResponseBody {
	s.BindVpcs = v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetCreateTime(v string) *DescribeResolverRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetCreateTimestamp(v int64) *DescribeResolverRuleResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetEndpointId(v string) *DescribeResolverRuleResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetEndpointName(v string) *DescribeResolverRuleResponseBody {
	s.EndpointName = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetForwardIps(v []*DescribeResolverRuleResponseBodyForwardIps) *DescribeResolverRuleResponseBody {
	s.ForwardIps = v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetId(v string) *DescribeResolverRuleResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetName(v string) *DescribeResolverRuleResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetRequestId(v string) *DescribeResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetType(v string) *DescribeResolverRuleResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetUpdateTime(v string) *DescribeResolverRuleResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetUpdateTimestamp(v int64) *DescribeResolverRuleResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverRuleResponseBody) SetZoneName(v string) *DescribeResolverRuleResponseBody {
	s.ZoneName = &v
	return s
}

type DescribeResolverRuleResponseBodyBindVpcs struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// Vpc ID
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName   *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcType   *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	VpcUserId *string `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeResolverRuleResponseBodyBindVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRuleResponseBodyBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetRegionId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.RegionId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetRegionName(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.RegionName = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcName(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcType(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcType = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyBindVpcs) SetVpcUserId(v string) *DescribeResolverRuleResponseBodyBindVpcs {
	s.VpcUserId = &v
	return s
}

type DescribeResolverRuleResponseBodyForwardIps struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeResolverRuleResponseBodyForwardIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRuleResponseBodyForwardIps) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponseBodyForwardIps) SetIp(v string) *DescribeResolverRuleResponseBodyForwardIps {
	s.Ip = &v
	return s
}

func (s *DescribeResolverRuleResponseBodyForwardIps) SetPort(v int32) *DescribeResolverRuleResponseBodyForwardIps {
	s.Port = &v
	return s
}

type DescribeResolverRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponse) SetHeaders(v map[string]*string) *DescribeResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverRuleResponse) SetStatusCode(v int32) *DescribeResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverRuleResponse) SetBody(v *DescribeResolverRuleResponseBody) *DescribeResolverRuleResponse {
	s.Body = v
	return s
}

type DescribeResolverRulesRequest struct {
	EndpointId           *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NeedDetailAttributes *bool   `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeResolverRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesRequest) SetEndpointId(v string) *DescribeResolverRulesRequest {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetKeyword(v string) *DescribeResolverRulesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetLang(v string) *DescribeResolverRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetNeedDetailAttributes(v bool) *DescribeResolverRulesRequest {
	s.NeedDetailAttributes = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetPageNumber(v int32) *DescribeResolverRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverRulesRequest) SetPageSize(v int32) *DescribeResolverRulesRequest {
	s.PageSize = &v
	return s
}

type DescribeResolverRulesResponseBody struct {
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules      []*DescribeResolverRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalItems *int32                                    `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                                    `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeResolverRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBody) SetPageNumber(v int32) *DescribeResolverRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetPageSize(v int32) *DescribeResolverRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetRequestId(v string) *DescribeResolverRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetRules(v []*DescribeResolverRulesResponseBodyRules) *DescribeResolverRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetTotalItems(v int32) *DescribeResolverRulesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeResolverRulesResponseBody) SetTotalPages(v int32) *DescribeResolverRulesResponseBody {
	s.TotalPages = &v
	return s
}

type DescribeResolverRulesResponseBodyRules struct {
	BindVpcs        []*DescribeResolverRulesResponseBodyRulesBindVpcs   `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	CreateTime      *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64                                              `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EndpointId      *string                                             `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointName    *string                                             `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	ForwardIps      []*DescribeResolverRulesResponseBodyRulesForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
	Id              *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime      *string                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64                                              `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	ZoneName        *string                                             `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRules) SetBindVpcs(v []*DescribeResolverRulesResponseBodyRulesBindVpcs) *DescribeResolverRulesResponseBodyRules {
	s.BindVpcs = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetCreateTime(v string) *DescribeResolverRulesResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetCreateTimestamp(v int64) *DescribeResolverRulesResponseBodyRules {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetEndpointId(v string) *DescribeResolverRulesResponseBodyRules {
	s.EndpointId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetEndpointName(v string) *DescribeResolverRulesResponseBodyRules {
	s.EndpointName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetForwardIps(v []*DescribeResolverRulesResponseBodyRulesForwardIps) *DescribeResolverRulesResponseBodyRules {
	s.ForwardIps = v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetId(v string) *DescribeResolverRulesResponseBodyRules {
	s.Id = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetName(v string) *DescribeResolverRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetType(v string) *DescribeResolverRulesResponseBodyRules {
	s.Type = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetUpdateTime(v string) *DescribeResolverRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetUpdateTimestamp(v int64) *DescribeResolverRulesResponseBodyRules {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRules) SetZoneName(v string) *DescribeResolverRulesResponseBodyRules {
	s.ZoneName = &v
	return s
}

type DescribeResolverRulesResponseBodyRulesBindVpcs struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// VPC ID
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName   *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcType   *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	VpcUserId *string `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesBindVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetRegionId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.RegionId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetRegionName(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.RegionName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcName(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcType(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcType = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesBindVpcs) SetVpcUserId(v string) *DescribeResolverRulesResponseBodyRulesBindVpcs {
	s.VpcUserId = &v
	return s
}

type DescribeResolverRulesResponseBodyRulesForwardIps struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeResolverRulesResponseBodyRulesForwardIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesResponseBodyRulesForwardIps) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) SetIp(v string) *DescribeResolverRulesResponseBodyRulesForwardIps {
	s.Ip = &v
	return s
}

func (s *DescribeResolverRulesResponseBodyRulesForwardIps) SetPort(v int32) *DescribeResolverRulesResponseBodyRulesForwardIps {
	s.Port = &v
	return s
}

type DescribeResolverRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResolverRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverRulesResponse) SetHeaders(v map[string]*string) *DescribeResolverRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverRulesResponse) SetStatusCode(v int32) *DescribeResolverRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverRulesResponse) SetBody(v *DescribeResolverRulesResponseBody) *DescribeResolverRulesResponse {
	s.Body = v
	return s
}

type DescribeStatisticSummaryRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeStatisticSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryRequest) SetLang(v string) *DescribeStatisticSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStatisticSummaryRequest) SetUserClientIp(v string) *DescribeStatisticSummaryRequest {
	s.UserClientIp = &v
	return s
}

type DescribeStatisticSummaryResponseBody struct {
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcRequestTops  *DescribeStatisticSummaryResponseBodyVpcRequestTops  `json:"VpcRequestTops,omitempty" xml:"VpcRequestTops,omitempty" type:"Struct"`
	ZoneRequestTops *DescribeStatisticSummaryResponseBodyZoneRequestTops `json:"ZoneRequestTops,omitempty" xml:"ZoneRequestTops,omitempty" type:"Struct"`
}

func (s DescribeStatisticSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBody) SetRequestId(v string) *DescribeStatisticSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetTotalCount(v int64) *DescribeStatisticSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetVpcRequestTops(v *DescribeStatisticSummaryResponseBodyVpcRequestTops) *DescribeStatisticSummaryResponseBody {
	s.VpcRequestTops = v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetZoneRequestTops(v *DescribeStatisticSummaryResponseBodyZoneRequestTops) *DescribeStatisticSummaryResponseBody {
	s.ZoneRequestTops = v
	return s
}

type DescribeStatisticSummaryResponseBodyVpcRequestTops struct {
	VpcRequestTop []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop `json:"VpcRequestTop,omitempty" xml:"VpcRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTops) SetVpcRequestTop(v []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) *DescribeStatisticSummaryResponseBodyVpcRequestTops {
	s.VpcRequestTop = v
	return s
}

type DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	TunnelId     *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// VPC ID
	VpcId   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionName(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionName = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetTunnelId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.TunnelId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetVpcId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.VpcId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetVpcType(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.VpcType = &v
	return s
}

type DescribeStatisticSummaryResponseBodyZoneRequestTops struct {
	ZoneRequestTop []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTops) SetZoneRequestTop(v []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) *DescribeStatisticSummaryResponseBodyZoneRequestTops {
	s.ZoneRequestTop = v
	return s
}

type DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	RequestCount *int64  `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	ZoneName     *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetBizType(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.BizType = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetZoneName(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.ZoneName = &v
	return s
}

type DescribeStatisticSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatisticSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatisticSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStatisticSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponse) SetHeaders(v map[string]*string) *DescribeStatisticSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatisticSummaryResponse) SetStatusCode(v int32) *DescribeStatisticSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatisticSummaryResponse) SetBody(v *DescribeStatisticSummaryResponseBody) *DescribeStatisticSummaryResponse {
	s.Body = v
	return s
}

type DescribeSyncEcsHostTaskRequest struct {
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSyncEcsHostTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskRequest) SetLang(v string) *DescribeSyncEcsHostTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSyncEcsHostTaskRequest) SetZoneId(v string) *DescribeSyncEcsHostTaskRequest {
	s.ZoneId = &v
	return s
}

type DescribeSyncEcsHostTaskResponseBody struct {
	EcsRegions *DescribeSyncEcsHostTaskResponseBodyEcsRegions `json:"EcsRegions,omitempty" xml:"EcsRegions,omitempty" type:"Struct"`
	Regions    *DescribeSyncEcsHostTaskResponseBodyRegions    `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status     *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Success    *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	ZoneId     *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetEcsRegions(v *DescribeSyncEcsHostTaskResponseBodyEcsRegions) *DescribeSyncEcsHostTaskResponseBody {
	s.EcsRegions = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetRegions(v *DescribeSyncEcsHostTaskResponseBodyRegions) *DescribeSyncEcsHostTaskResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetRequestId(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetStatus(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetSuccess(v bool) *DescribeSyncEcsHostTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBody) SetZoneId(v string) *DescribeSyncEcsHostTaskResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegions struct {
	EcsRegion []*DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion `json:"EcsRegion,omitempty" xml:"EcsRegion,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegions) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegions) SetEcsRegion(v []*DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) *DescribeSyncEcsHostTaskResponseBodyEcsRegions {
	s.EcsRegion = v
	return s
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion struct {
	RegionIds *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	UserId    *int64                                                           `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) SetRegionIds(v *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion {
	s.RegionIds = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion) SetUserId(v int64) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegion {
	s.UserId = &v
	return s
}

type DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds) SetRegionId(v []*string) *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds {
	s.RegionId = v
	return s
}

type DescribeSyncEcsHostTaskResponseBodyRegions struct {
	RegionId []*string `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeSyncEcsHostTaskResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponseBodyRegions) SetRegionId(v []*string) *DescribeSyncEcsHostTaskResponseBodyRegions {
	s.RegionId = v
	return s
}

type DescribeSyncEcsHostTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSyncEcsHostTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSyncEcsHostTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSyncEcsHostTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSyncEcsHostTaskResponse) SetHeaders(v map[string]*string) *DescribeSyncEcsHostTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSyncEcsHostTaskResponse) SetStatusCode(v int32) *DescribeSyncEcsHostTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSyncEcsHostTaskResponse) SetBody(v *DescribeSyncEcsHostTaskResponseBody) *DescribeSyncEcsHostTaskResponse {
	s.Body = v
	return s
}

type DescribeTagsRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetLang(v string) *DescribeTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int32) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int32) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

type DescribeTagsResponseBody struct {
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags       []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) SetPageNumber(v int32) *DescribeTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsResponseBody) SetPageSize(v int32) *DescribeTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTags(v []*DescribeTagsResponseBodyTags) *DescribeTagsResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTagsResponseBody) SetTotalCount(v int32) *DescribeTagsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTagsResponseBodyTags struct {
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTags) SetKey(v string) *DescribeTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeTagsResponseBodyTags) SetValues(v []*string) *DescribeTagsResponseBodyTags {
	s.Values = v
	return s
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

type DescribeUserVpcAuthorizationsRequest struct {
	AuthType         *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthorizedUserId *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserVpcAuthorizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsRequest) SetAuthType(v string) *DescribeUserVpcAuthorizationsRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetAuthorizedUserId(v int64) *DescribeUserVpcAuthorizationsRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetPageNumber(v int32) *DescribeUserVpcAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsRequest) SetPageSize(v int32) *DescribeUserVpcAuthorizationsRequest {
	s.PageSize = &v
	return s
}

type DescribeUserVpcAuthorizationsResponseBody struct {
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems *int32                                            `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Users      []*DescribeUserVpcAuthorizationsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUserVpcAuthorizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetPageNumber(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetPageSize(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetRequestId(v string) *DescribeUserVpcAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetTotalItems(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetTotalPages(v int32) *DescribeUserVpcAuthorizationsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBody) SetUsers(v []*DescribeUserVpcAuthorizationsResponseBodyUsers) *DescribeUserVpcAuthorizationsResponseBody {
	s.Users = v
	return s
}

type DescribeUserVpcAuthorizationsResponseBodyUsers struct {
	AuthType           *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AuthorizedAliyunId *string `json:"AuthorizedAliyunId,omitempty" xml:"AuthorizedAliyunId,omitempty"`
	AuthorizedUserId   *int64  `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp    *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribeUserVpcAuthorizationsResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthType(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthType = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthorizedAliyunId(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthorizedAliyunId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetAuthorizedUserId(v int64) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetCreateTime(v string) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponseBodyUsers) SetCreateTimestamp(v int64) *DescribeUserVpcAuthorizationsResponseBodyUsers {
	s.CreateTimestamp = &v
	return s
}

type DescribeUserVpcAuthorizationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserVpcAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserVpcAuthorizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponse) SetHeaders(v map[string]*string) *DescribeUserVpcAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponse) SetStatusCode(v int32) *DescribeUserVpcAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponse) SetBody(v *DescribeUserVpcAuthorizationsResponseBody) *DescribeUserVpcAuthorizationsResponse {
	s.Body = v
	return s
}

type DescribeZoneInfoRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Zone ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoRequest) SetLang(v string) *DescribeZoneInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneInfoRequest) SetZoneId(v string) *DescribeZoneInfoRequest {
	s.ZoneId = &v
	return s
}

type DescribeZoneInfoResponseBody struct {
	BindVpcs         *DescribeZoneInfoResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	CreateTime       *string                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp  *int64                                `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Creator          *string                               `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorType      *string                               `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	DnsGroup         *string                               `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	DnsGroupChanging *bool                                 `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	IsPtr            *bool                                 `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	ProxyPattern     *string                               `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	RecordCount      *int32                                `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Remark           *string                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId  *string                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDns         *bool                                 `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	UpdateTime       *string                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp  *int64                                `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Zone ID。
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneTag  *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZoneInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBody) SetBindVpcs(v *DescribeZoneInfoResponseBodyBindVpcs) *DescribeZoneInfoResponseBody {
	s.BindVpcs = v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTime(v string) *DescribeZoneInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreator(v string) *DescribeZoneInfoResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetCreatorType(v string) *DescribeZoneInfoResponseBody {
	s.CreatorType = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetDnsGroup(v string) *DescribeZoneInfoResponseBody {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetDnsGroupChanging(v bool) *DescribeZoneInfoResponseBody {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetIsPtr(v bool) *DescribeZoneInfoResponseBody {
	s.IsPtr = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetProxyPattern(v string) *DescribeZoneInfoResponseBody {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRecordCount(v int32) *DescribeZoneInfoResponseBody {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRemark(v string) *DescribeZoneInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetRequestId(v string) *DescribeZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetResourceGroupId(v string) *DescribeZoneInfoResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetSlaveDns(v bool) *DescribeZoneInfoResponseBody {
	s.SlaveDns = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTime(v string) *DescribeZoneInfoResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetUpdateTimestamp(v int64) *DescribeZoneInfoResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneId(v string) *DescribeZoneInfoResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneName(v string) *DescribeZoneInfoResponseBody {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneTag(v string) *DescribeZoneInfoResponseBody {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneInfoResponseBody) SetZoneType(v string) *DescribeZoneInfoResponseBody {
	s.ZoneType = &v
	return s
}

type DescribeZoneInfoResponseBodyBindVpcs struct {
	Vpc []*DescribeZoneInfoResponseBodyBindVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneInfoResponseBodyBindVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcs) SetVpc(v []*DescribeZoneInfoResponseBodyBindVpcsVpc) *DescribeZoneInfoResponseBodyBindVpcs {
	s.Vpc = v
	return s
}

type DescribeZoneInfoResponseBodyBindVpcsVpc struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// Vpc ID。
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName   *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcType   *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	VpcUserId *int64  `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponseBodyBindVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetRegionName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcId(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcName(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcType(v string) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcType = &v
	return s
}

func (s *DescribeZoneInfoResponseBodyBindVpcsVpc) SetVpcUserId(v int64) *DescribeZoneInfoResponseBodyBindVpcsVpc {
	s.VpcUserId = &v
	return s
}

type DescribeZoneInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneInfoResponse) SetHeaders(v map[string]*string) *DescribeZoneInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneInfoResponse) SetStatusCode(v int32) *DescribeZoneInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneInfoResponse) SetBody(v *DescribeZoneInfoResponseBody) *DescribeZoneInfoResponse {
	s.Body = v
	return s
}

type DescribeZoneRecordsRequest struct {
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchMode   *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// Zone ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsRequest) SetKeyword(v string) *DescribeZoneRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetLang(v string) *DescribeZoneRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageNumber(v int32) *DescribeZoneRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageSize(v int32) *DescribeZoneRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetSearchMode(v string) *DescribeZoneRecordsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetTag(v string) *DescribeZoneRecordsRequest {
	s.Tag = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetUserClientIp(v string) *DescribeZoneRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetZoneId(v string) *DescribeZoneRecordsRequest {
	s.ZoneId = &v
	return s
}

type DescribeZoneRecordsResponseBody struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records    *DescribeZoneRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems *int32                                  `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeZoneRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBody) SetPageNumber(v int32) *DescribeZoneRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetPageSize(v int32) *DescribeZoneRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRecords(v *DescribeZoneRecordsResponseBodyRecords) *DescribeZoneRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetRequestId(v string) *DescribeZoneRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalItems(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZoneRecordsResponseBody) SetTotalPages(v int32) *DescribeZoneRecordsResponseBody {
	s.TotalPages = &v
	return s
}

type DescribeZoneRecordsResponseBodyRecords struct {
	Record []*DescribeZoneRecordsResponseBodyRecordsRecord `json:"Record,omitempty" xml:"Record,omitempty" type:"Repeated"`
}

func (s DescribeZoneRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecords) SetRecord(v []*DescribeZoneRecordsResponseBodyRecordsRecord) *DescribeZoneRecordsResponseBodyRecords {
	s.Record = v
	return s
}

type DescribeZoneRecordsResponseBodyRecordsRecord struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Line            *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecordId        *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Rr              *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight          *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponseBodyRecordsRecord) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetCreateTime(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetCreateTimestamp(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetLine(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Line = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetPriority(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Priority = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRecordId(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.RecordId = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRemark(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Remark = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetRr(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Rr = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetStatus(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Status = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetTtl(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Ttl = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetType(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Type = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetUpdateTime(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetUpdateTimestamp(v int64) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetValue(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Value = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetWeight(v int32) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.Weight = &v
	return s
}

func (s *DescribeZoneRecordsResponseBodyRecordsRecord) SetZoneId(v string) *DescribeZoneRecordsResponseBodyRecordsRecord {
	s.ZoneId = &v
	return s
}

type DescribeZoneRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsResponse) SetHeaders(v map[string]*string) *DescribeZoneRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneRecordsResponse) SetStatusCode(v int32) *DescribeZoneRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneRecordsResponse) SetBody(v *DescribeZoneRecordsResponseBody) *DescribeZoneRecordsResponse {
	s.Body = v
	return s
}

type DescribeZoneVpcTreeRequest struct {
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeZoneVpcTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeRequest) SetLang(v string) *DescribeZoneVpcTreeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneVpcTreeRequest) SetUserClientIp(v string) *DescribeZoneVpcTreeRequest {
	s.UserClientIp = &v
	return s
}

type DescribeZoneVpcTreeResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZoneVpcTreeResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZoneVpcTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBody) SetRequestId(v string) *DescribeZoneVpcTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBody) SetZones(v *DescribeZoneVpcTreeResponseBodyZones) *DescribeZoneVpcTreeResponseBody {
	s.Zones = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZones struct {
	Zone []*DescribeZoneVpcTreeResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZones) SetZone(v []*DescribeZoneVpcTreeResponseBodyZonesZone) *DescribeZoneVpcTreeResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZone struct {
	CreateTime       *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp  *int64                                        `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Creator          *string                                       `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorType      *string                                       `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	DnsGroup         *string                                       `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	DnsGroupChanging *bool                                         `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	IsPtr            *bool                                         `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	RecordCount      *int32                                        `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Remark           *string                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UpdateTime       *string                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp  *int64                                        `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	Vpcs             *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
	// Zone id
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneTag  *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreator(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetCreatorType(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.CreatorType = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetDnsGroup(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetDnsGroupChanging(v bool) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetRemark(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetVpcs(v *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.Vpcs = v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneId(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneName(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneTag(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZone) SetZoneType(v string) *DescribeZoneVpcTreeResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcs struct {
	Vpc []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs) SetVpc(v []*DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs {
	s.Vpc = v
	return s
}

type DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc struct {
	// region Id
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// vpc id
	VpcId   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetRegionName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.RegionName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcId(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcName(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc) SetVpcType(v string) *DescribeZoneVpcTreeResponseBodyZonesZoneVpcsVpc {
	s.VpcType = &v
	return s
}

type DescribeZoneVpcTreeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZoneVpcTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZoneVpcTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZoneVpcTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeResponse) SetHeaders(v map[string]*string) *DescribeZoneVpcTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeZoneVpcTreeResponse) SetStatusCode(v int32) *DescribeZoneVpcTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZoneVpcTreeResponse) SetBody(v *DescribeZoneVpcTreeResponseBody) *DescribeZoneVpcTreeResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	Keyword       *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
	// VPC ID。
	QueryVpcId      *string                            `json:"QueryVpcId,omitempty" xml:"QueryVpcId,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTag     []*DescribeZonesRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	SearchMode      *string                            `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	ZoneTag         []*string                          `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty" type:"Repeated"`
	ZoneType        *string                            `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetKeyword(v string) *DescribeZonesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZonesRequest) SetLang(v string) *DescribeZonesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZonesRequest) SetPageNumber(v int32) *DescribeZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesRequest) SetPageSize(v int32) *DescribeZonesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryRegionId(v string) *DescribeZonesRequest {
	s.QueryRegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetQueryVpcId(v string) *DescribeZonesRequest {
	s.QueryVpcId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceGroupId(v string) *DescribeZonesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesRequest) SetResourceTag(v []*DescribeZonesRequestResourceTag) *DescribeZonesRequest {
	s.ResourceTag = v
	return s
}

func (s *DescribeZonesRequest) SetSearchMode(v string) *DescribeZonesRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneTag(v []*string) *DescribeZonesRequest {
	s.ZoneTag = v
	return s
}

func (s *DescribeZonesRequest) SetZoneType(v string) *DescribeZonesRequest {
	s.ZoneType = &v
	return s
}

type DescribeZonesRequestResourceTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeZonesRequestResourceTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequestResourceTag) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequestResourceTag) SetKey(v string) *DescribeZonesRequestResourceTag {
	s.Key = &v
	return s
}

func (s *DescribeZonesRequestResourceTag) SetValue(v string) *DescribeZonesRequestResourceTag {
	s.Value = &v
	return s
}

type DescribeZonesResponseBody struct {
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItems *int32                          `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	TotalPages *int32                          `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Zones      *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetPageNumber(v int32) *DescribeZonesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeZonesResponseBody) SetPageSize(v int32) *DescribeZonesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalItems(v int32) *DescribeZonesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeZonesResponseBody) SetTotalPages(v int32) *DescribeZonesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	CreateTime       *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp  *int64                                          `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Creator          *string                                         `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CreatorSubType   *string                                         `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	DnsGroup         *string                                         `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	DnsGroupChanging *bool                                           `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	IsPtr            *bool                                           `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	ProxyPattern     *string                                         `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	RecordCount      *int32                                          `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	Remark           *string                                         `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceGroupId  *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceTags     *DescribeZonesResponseBodyZonesZoneResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Struct"`
	UpdateTime       *string                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp  *int64                                          `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// zone ID。
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	ZoneTag  *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.CreateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreator(v string) *DescribeZonesResponseBodyZonesZone {
	s.Creator = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetCreatorSubType(v string) *DescribeZonesResponseBodyZonesZone {
	s.CreatorSubType = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetDnsGroup(v string) *DescribeZonesResponseBodyZonesZone {
	s.DnsGroup = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetDnsGroupChanging(v bool) *DescribeZonesResponseBodyZonesZone {
	s.DnsGroupChanging = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetIsPtr(v bool) *DescribeZonesResponseBodyZonesZone {
	s.IsPtr = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetProxyPattern(v string) *DescribeZonesResponseBodyZonesZone {
	s.ProxyPattern = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRecordCount(v int32) *DescribeZonesResponseBodyZonesZone {
	s.RecordCount = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetRemark(v string) *DescribeZonesResponseBodyZonesZone {
	s.Remark = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetResourceGroupId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetResourceTags(v *DescribeZonesResponseBodyZonesZoneResourceTags) *DescribeZonesResponseBodyZonesZone {
	s.ResourceTags = v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTime(v string) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTime = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetUpdateTimestamp(v int64) *DescribeZonesResponseBodyZonesZone {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneName(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneTag(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneTag = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneType(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneType = &v
	return s
}

type DescribeZonesResponseBodyZonesZoneResourceTags struct {
	ResourceTag []*DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZonesZoneResourceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTags) SetResourceTag(v []*DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) *DescribeZonesResponseBodyZonesZoneResourceTags {
	s.ResourceTag = v
	return s
}

type DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) SetKey(v string) *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag {
	s.Key = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag) SetValue(v string) *DescribeZonesResponseBodyZonesZoneResourceTagsResourceTag {
	s.Value = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	Lang         *string                       `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Size         *int32                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetLang(v string) *ListTagResourcesRequest {
	s.Lang = &v
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

func (s *ListTagResourcesRequest) SetSize(v int32) *ListTagResourcesRequest {
	s.Size = &v
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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

type MoveResourceGroupRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// Zone Id。
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveResourceGroupRequest) SetLang(v string) *MoveResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type SetProxyPatternRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetProxyPatternRequest) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternRequest) GoString() string {
	return s.String()
}

func (s *SetProxyPatternRequest) SetClientToken(v string) *SetProxyPatternRequest {
	s.ClientToken = &v
	return s
}

func (s *SetProxyPatternRequest) SetLang(v string) *SetProxyPatternRequest {
	s.Lang = &v
	return s
}

func (s *SetProxyPatternRequest) SetProxyPattern(v string) *SetProxyPatternRequest {
	s.ProxyPattern = &v
	return s
}

func (s *SetProxyPatternRequest) SetUserClientIp(v string) *SetProxyPatternRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetProxyPatternRequest) SetZoneId(v string) *SetProxyPatternRequest {
	s.ZoneId = &v
	return s
}

type SetProxyPatternResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SetProxyPatternResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternResponseBody) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponseBody) SetRequestId(v string) *SetProxyPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetProxyPatternResponseBody) SetZoneId(v string) *SetProxyPatternResponseBody {
	s.ZoneId = &v
	return s
}

type SetProxyPatternResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetProxyPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetProxyPatternResponse) String() string {
	return tea.Prettify(s)
}

func (s SetProxyPatternResponse) GoString() string {
	return s.String()
}

func (s *SetProxyPatternResponse) SetHeaders(v map[string]*string) *SetProxyPatternResponse {
	s.Headers = v
	return s
}

func (s *SetProxyPatternResponse) SetStatusCode(v int32) *SetProxyPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *SetProxyPatternResponse) SetBody(v *SetProxyPatternResponseBody) *SetProxyPatternResponse {
	s.Body = v
	return s
}

type SetZoneRecordStatusRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetZoneRecordStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusRequest) SetClientToken(v string) *SetZoneRecordStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetLang(v string) *SetZoneRecordStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetRecordId(v int64) *SetZoneRecordStatusRequest {
	s.RecordId = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetStatus(v string) *SetZoneRecordStatusRequest {
	s.Status = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetUserClientIp(v string) *SetZoneRecordStatusRequest {
	s.UserClientIp = &v
	return s
}

type SetZoneRecordStatusResponseBody struct {
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetZoneRecordStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponseBody) SetRecordId(v int64) *SetZoneRecordStatusResponseBody {
	s.RecordId = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetRequestId(v string) *SetZoneRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetStatus(v string) *SetZoneRecordStatusResponseBody {
	s.Status = &v
	return s
}

type SetZoneRecordStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetZoneRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetZoneRecordStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetZoneRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponse) SetHeaders(v map[string]*string) *SetZoneRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *SetZoneRecordStatusResponse) SetStatusCode(v int32) *SetZoneRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetZoneRecordStatusResponse) SetBody(v *SetZoneRecordStatusResponseBody) *SetZoneRecordStatusResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	Lang         *string                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OverWrite    *bool                     `json:"OverWrite,omitempty" xml:"OverWrite,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetLang(v string) *TagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *TagResourcesRequest) SetOverWrite(v bool) *TagResourcesRequest {
	s.OverWrite = &v
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetLang(v string) *UntagResourcesRequest {
	s.Lang = &v
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

type UpdateRecordRemarkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RecordId    *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateRecordRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkRequest) SetClientToken(v string) *UpdateRecordRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetLang(v string) *UpdateRecordRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRecordId(v int64) *UpdateRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRemarkRequest) SetRemark(v string) *UpdateRecordRemarkRequest {
	s.Remark = &v
	return s
}

type UpdateRecordRemarkResponseBody struct {
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponseBody) SetRecordId(v int64) *UpdateRecordRemarkResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRemarkResponseBody) SetRequestId(v string) *UpdateRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRecordRemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecordRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponse) SetHeaders(v map[string]*string) *UpdateRecordRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordRemarkResponse) SetStatusCode(v int32) *UpdateRecordRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordRemarkResponse) SetBody(v *UpdateRecordRemarkResponseBody) *UpdateRecordRemarkResponse {
	s.Body = v
	return s
}

type UpdateResolverEndpointRequest struct {
	EndpointId *string                                  `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	IpConfig   []*UpdateResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	Lang       *string                                  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name       *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateResolverEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointRequest) SetEndpointId(v string) *UpdateResolverEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateResolverEndpointRequest) SetIpConfig(v []*UpdateResolverEndpointRequestIpConfig) *UpdateResolverEndpointRequest {
	s.IpConfig = v
	return s
}

func (s *UpdateResolverEndpointRequest) SetLang(v string) *UpdateResolverEndpointRequest {
	s.Lang = &v
	return s
}

func (s *UpdateResolverEndpointRequest) SetName(v string) *UpdateResolverEndpointRequest {
	s.Name = &v
	return s
}

type UpdateResolverEndpointRequestIpConfig struct {
	AzId      *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateResolverEndpointRequestIpConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverEndpointRequestIpConfig) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointRequestIpConfig) SetAzId(v string) *UpdateResolverEndpointRequestIpConfig {
	s.AzId = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetCidrBlock(v string) *UpdateResolverEndpointRequestIpConfig {
	s.CidrBlock = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetIp(v string) *UpdateResolverEndpointRequestIpConfig {
	s.Ip = &v
	return s
}

func (s *UpdateResolverEndpointRequestIpConfig) SetVSwitchId(v string) *UpdateResolverEndpointRequestIpConfig {
	s.VSwitchId = &v
	return s
}

type UpdateResolverEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResolverEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointResponseBody) SetRequestId(v string) *UpdateResolverEndpointResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResolverEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResolverEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateResolverEndpointResponse) SetHeaders(v map[string]*string) *UpdateResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateResolverEndpointResponse) SetStatusCode(v int32) *UpdateResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResolverEndpointResponse) SetBody(v *UpdateResolverEndpointResponseBody) *UpdateResolverEndpointResponse {
	s.Body = v
	return s
}

type UpdateResolverRuleRequest struct {
	ForwardIp []*UpdateResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	Lang      *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name      *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId    *string                               `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateResolverRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleRequest) SetForwardIp(v []*UpdateResolverRuleRequestForwardIp) *UpdateResolverRuleRequest {
	s.ForwardIp = v
	return s
}

func (s *UpdateResolverRuleRequest) SetLang(v string) *UpdateResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateResolverRuleRequest) SetName(v string) *UpdateResolverRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateResolverRuleRequest) SetRuleId(v string) *UpdateResolverRuleRequest {
	s.RuleId = &v
	return s
}

type UpdateResolverRuleRequestForwardIp struct {
	Ip   *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateResolverRuleRequestForwardIp) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverRuleRequestForwardIp) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleRequestForwardIp) SetIp(v string) *UpdateResolverRuleRequestForwardIp {
	s.Ip = &v
	return s
}

func (s *UpdateResolverRuleRequestForwardIp) SetPort(v int32) *UpdateResolverRuleRequestForwardIp {
	s.Port = &v
	return s
}

type UpdateResolverRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResolverRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleResponseBody) SetRequestId(v string) *UpdateResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResolverRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResolverRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleResponse) SetHeaders(v map[string]*string) *UpdateResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateResolverRuleResponse) SetStatusCode(v int32) *UpdateResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResolverRuleResponse) SetBody(v *UpdateResolverRuleResponseBody) *UpdateResolverRuleResponse {
	s.Body = v
	return s
}

type UpdateSyncEcsHostTaskRequest struct {
	Lang   *string                               `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Region []*UpdateSyncEcsHostTaskRequestRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
	Status *string                               `json:"Status,omitempty" xml:"Status,omitempty"`
	ZoneId *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateSyncEcsHostTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncEcsHostTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskRequest) SetLang(v string) *UpdateSyncEcsHostTaskRequest {
	s.Lang = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetRegion(v []*UpdateSyncEcsHostTaskRequestRegion) *UpdateSyncEcsHostTaskRequest {
	s.Region = v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetStatus(v string) *UpdateSyncEcsHostTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequest) SetZoneId(v string) *UpdateSyncEcsHostTaskRequest {
	s.ZoneId = &v
	return s
}

type UpdateSyncEcsHostTaskRequestRegion struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateSyncEcsHostTaskRequestRegion) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncEcsHostTaskRequestRegion) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskRequestRegion) SetRegionId(v string) *UpdateSyncEcsHostTaskRequestRegion {
	s.RegionId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskRequestRegion) SetUserId(v int64) *UpdateSyncEcsHostTaskRequestRegion {
	s.UserId = &v
	return s
}

type UpdateSyncEcsHostTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSyncEcsHostTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncEcsHostTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskResponseBody) SetRequestId(v string) *UpdateSyncEcsHostTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSyncEcsHostTaskResponseBody) SetSuccess(v bool) *UpdateSyncEcsHostTaskResponseBody {
	s.Success = &v
	return s
}

type UpdateSyncEcsHostTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSyncEcsHostTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSyncEcsHostTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSyncEcsHostTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateSyncEcsHostTaskResponse) SetHeaders(v map[string]*string) *UpdateSyncEcsHostTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateSyncEcsHostTaskResponse) SetStatusCode(v int32) *UpdateSyncEcsHostTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSyncEcsHostTaskResponse) SetBody(v *UpdateSyncEcsHostTaskResponseBody) *UpdateSyncEcsHostTaskResponse {
	s.Body = v
	return s
}

type UpdateZoneRecordRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Line         *string `json:"Line,omitempty" xml:"Line,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecordId     *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	Rr           *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	Ttl          *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Weight       *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateZoneRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordRequest) SetClientToken(v string) *UpdateZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetLang(v string) *UpdateZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetLine(v string) *UpdateZoneRecordRequest {
	s.Line = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetPriority(v int32) *UpdateZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetRecordId(v int64) *UpdateZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetRr(v string) *UpdateZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetTtl(v int32) *UpdateZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetType(v string) *UpdateZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetUserClientIp(v string) *UpdateZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetValue(v string) *UpdateZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetWeight(v int32) *UpdateZoneRecordRequest {
	s.Weight = &v
	return s
}

type UpdateZoneRecordResponseBody struct {
	RecordId  *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateZoneRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponseBody) SetRecordId(v int64) *UpdateZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordResponseBody) SetRequestId(v string) *UpdateZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateZoneRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZoneRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZoneRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponse) SetHeaders(v map[string]*string) *UpdateZoneRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRecordResponse) SetStatusCode(v int32) *UpdateZoneRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZoneRecordResponse) SetBody(v *UpdateZoneRecordResponseBody) *UpdateZoneRecordResponse {
	s.Body = v
	return s
}

type UpdateZoneRemarkRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// Zone ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateZoneRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkRequest) SetClientToken(v string) *UpdateZoneRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetLang(v string) *UpdateZoneRemarkRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetRemark(v string) *UpdateZoneRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetUserClientIp(v string) *UpdateZoneRemarkRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateZoneRemarkRequest) SetZoneId(v string) *UpdateZoneRemarkRequest {
	s.ZoneId = &v
	return s
}

type UpdateZoneRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Zone ID。
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateZoneRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponseBody) SetRequestId(v string) *UpdateZoneRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZoneRemarkResponseBody) SetZoneId(v string) *UpdateZoneRemarkResponseBody {
	s.ZoneId = &v
	return s
}

type UpdateZoneRemarkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZoneRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZoneRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateZoneRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateZoneRemarkResponse) SetHeaders(v map[string]*string) *UpdateZoneRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateZoneRemarkResponse) SetStatusCode(v int32) *UpdateZoneRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZoneRemarkResponse) SetBody(v *UpdateZoneRemarkResponseBody) *UpdateZoneRemarkResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pvtz"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddResolverEndpointWithOptions(request *AddResolverEndpointRequest, runtime *util.RuntimeOptions) (_result *AddResolverEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpConfig)) {
		query["IpConfig"] = request.IpConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddResolverEndpoint"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResolverEndpoint(request *AddResolverEndpointRequest) (_result *AddResolverEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddResolverEndpointResponse{}
	_body, _err := client.AddResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddResolverRuleWithOptions(request *AddResolverRuleRequest, runtime *util.RuntimeOptions) (_result *AddResolverRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardIp)) {
		query["ForwardIp"] = request.ForwardIp
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneName)) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddResolverRule"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddResolverRule(request *AddResolverRuleRequest) (_result *AddResolverRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddResolverRuleResponse{}
	_body, _err := client.AddResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserVpcAuthorizationWithOptions(request *AddUserVpcAuthorizationRequest, runtime *util.RuntimeOptions) (_result *AddUserVpcAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthChannel)) {
		query["AuthChannel"] = request.AuthChannel
	}

	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserVpcAuthorization"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserVpcAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserVpcAuthorization(request *AddUserVpcAuthorizationRequest) (_result *AddUserVpcAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserVpcAuthorizationResponse{}
	_body, _err := client.AddUserVpcAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddZoneWithOptions(request *AddZoneRequest, runtime *util.RuntimeOptions) (_result *AddZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DnsGroup)) {
		query["DnsGroup"] = request.DnsGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyPattern)) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneName)) {
		query["ZoneName"] = request.ZoneName
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneTag)) {
		query["ZoneTag"] = request.ZoneTag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneType)) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddZone"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddZone(request *AddZoneRequest) (_result *AddZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddZoneResponse{}
	_body, _err := client.AddZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddZoneRecordWithOptions(request *AddZoneRecordRequest, runtime *util.RuntimeOptions) (_result *AddZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["Line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Rr)) {
		query["Rr"] = request.Rr
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		query["Weight"] = request.Weight
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddZoneRecord"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddZoneRecord(request *AddZoneRecordRequest) (_result *AddZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.AddZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindResolverRuleVpcWithOptions(request *BindResolverRuleVpcRequest, runtime *util.RuntimeOptions) (_result *BindResolverRuleVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Vpc)) {
		query["Vpc"] = request.Vpc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindResolverRuleVpc"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindResolverRuleVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindResolverRuleVpc(request *BindResolverRuleVpcRequest) (_result *BindResolverRuleVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindResolverRuleVpcResponse{}
	_body, _err := client.BindResolverRuleVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindZoneVpcWithOptions(request *BindZoneVpcRequest, runtime *util.RuntimeOptions) (_result *BindZoneVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.Vpcs)) {
		query["Vpcs"] = request.Vpcs
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindZoneVpc"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindZoneVpc(request *BindZoneVpcRequest) (_result *BindZoneVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.BindZoneVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckZoneNameWithOptions(request *CheckZoneNameRequest, runtime *util.RuntimeOptions) (_result *CheckZoneNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneName)) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckZoneName"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckZoneName(request *CheckZoneNameRequest) (_result *CheckZoneNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.CheckZoneNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResolverEndpointWithOptions(request *DeleteResolverEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteResolverEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResolverEndpoint"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResolverEndpoint(request *DeleteResolverEndpointRequest) (_result *DeleteResolverEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResolverEndpointResponse{}
	_body, _err := client.DeleteResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResolverRuleWithOptions(request *DeleteResolverRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteResolverRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResolverRule"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResolverRule(request *DeleteResolverRuleRequest) (_result *DeleteResolverRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResolverRuleResponse{}
	_body, _err := client.DeleteResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserVpcAuthorizationWithOptions(request *DeleteUserVpcAuthorizationRequest, runtime *util.RuntimeOptions) (_result *DeleteUserVpcAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserVpcAuthorization"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserVpcAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserVpcAuthorization(request *DeleteUserVpcAuthorizationRequest) (_result *DeleteUserVpcAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserVpcAuthorizationResponse{}
	_body, _err := client.DeleteUserVpcAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteZoneWithOptions(request *DeleteZoneRequest, runtime *util.RuntimeOptions) (_result *DeleteZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteZone"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteZone(request *DeleteZoneRequest) (_result *DeleteZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteZoneResponse{}
	_body, _err := client.DeleteZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteZoneRecordWithOptions(request *DeleteZoneRecordRequest, runtime *util.RuntimeOptions) (_result *DeleteZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteZoneRecord"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) (_result *DeleteZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.DeleteZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChangeLogsWithOptions(request *DescribeChangeLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeChangeLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChangeLogs"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChangeLogs(request *DescribeChangeLogsRequest) (_result *DescribeChangeLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.DescribeChangeLogsWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.VpcType)) {
		query["VpcType"] = request.VpcType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeRequestGraphWithOptions(request *DescribeRequestGraphRequest, runtime *util.RuntimeOptions) (_result *DescribeRequestGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRequestGraph"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRequestGraph(request *DescribeRequestGraphRequest) (_result *DescribeRequestGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.DescribeRequestGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResolverAvailableZonesWithOptions(request *DescribeResolverAvailableZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeResolverAvailableZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AzId)) {
		query["AzId"] = request.AzId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResolverRegionId)) {
		query["ResolverRegionId"] = request.ResolverRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResolverAvailableZones"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResolverAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResolverAvailableZones(request *DescribeResolverAvailableZonesRequest) (_result *DescribeResolverAvailableZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResolverAvailableZonesResponse{}
	_body, _err := client.DescribeResolverAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResolverEndpointWithOptions(request *DescribeResolverEndpointRequest, runtime *util.RuntimeOptions) (_result *DescribeResolverEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResolverEndpoint"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResolverEndpoint(request *DescribeResolverEndpointRequest) (_result *DescribeResolverEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResolverEndpointResponse{}
	_body, _err := client.DescribeResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResolverEndpointsWithOptions(request *DescribeResolverEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeResolverEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResolverEndpoints"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResolverEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResolverEndpoints(request *DescribeResolverEndpointsRequest) (_result *DescribeResolverEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResolverEndpointsResponse{}
	_body, _err := client.DescribeResolverEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResolverRuleWithOptions(request *DescribeResolverRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeResolverRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResolverRule"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResolverRule(request *DescribeResolverRuleRequest) (_result *DescribeResolverRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResolverRuleResponse{}
	_body, _err := client.DescribeResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResolverRulesWithOptions(request *DescribeResolverRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeResolverRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NeedDetailAttributes)) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
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
		Action:      tea.String("DescribeResolverRules"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResolverRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResolverRules(request *DescribeResolverRulesRequest) (_result *DescribeResolverRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResolverRulesResponse{}
	_body, _err := client.DescribeResolverRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStatisticSummaryWithOptions(request *DescribeStatisticSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeStatisticSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStatisticSummary"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStatisticSummary(request *DescribeStatisticSummaryRequest) (_result *DescribeStatisticSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.DescribeStatisticSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSyncEcsHostTaskWithOptions(request *DescribeSyncEcsHostTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeSyncEcsHostTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSyncEcsHostTask"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSyncEcsHostTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSyncEcsHostTask(request *DescribeSyncEcsHostTaskRequest) (_result *DescribeSyncEcsHostTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSyncEcsHostTaskResponse{}
	_body, _err := client.DescribeSyncEcsHostTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTags"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserVpcAuthorizationsWithOptions(request *DescribeUserVpcAuthorizationsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserVpcAuthorizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizedUserId)) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
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
		Action:      tea.String("DescribeUserVpcAuthorizations"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserVpcAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserVpcAuthorizations(request *DescribeUserVpcAuthorizationsRequest) (_result *DescribeUserVpcAuthorizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserVpcAuthorizationsResponse{}
	_body, _err := client.DescribeUserVpcAuthorizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneInfoWithOptions(request *DescribeZoneInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZoneInfo"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneInfo(request *DescribeZoneInfoRequest) (_result *DescribeZoneInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.DescribeZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneRecordsWithOptions(request *DescribeZoneRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchMode)) {
		query["SearchMode"] = request.SearchMode
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZoneRecords"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneRecords(request *DescribeZoneRecordsRequest) (_result *DescribeZoneRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.DescribeZoneRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZoneVpcTreeWithOptions(request *DescribeZoneVpcTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeZoneVpcTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZoneVpcTree"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZoneVpcTree(request *DescribeZoneVpcTreeRequest) (_result *DescribeZoneVpcTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.DescribeZoneVpcTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRegionId)) {
		query["QueryRegionId"] = request.QueryRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryVpcId)) {
		query["QueryVpcId"] = request.QueryVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTag)) {
		query["ResourceTag"] = request.ResourceTag
	}

	if !tea.BoolValue(util.IsUnset(request.SearchMode)) {
		query["SearchMode"] = request.SearchMode
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneTag)) {
		query["ZoneTag"] = request.ZoneTag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneType)) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SetProxyPatternWithOptions(request *SetProxyPatternRequest, runtime *util.RuntimeOptions) (_result *SetProxyPatternResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyPattern)) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetProxyPattern"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetProxyPattern(request *SetProxyPatternRequest) (_result *SetProxyPatternResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.SetProxyPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetZoneRecordStatusWithOptions(request *SetZoneRecordStatusRequest, runtime *util.RuntimeOptions) (_result *SetZoneRecordStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetZoneRecordStatus"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetZoneRecordStatus(request *SetZoneRecordStatusRequest) (_result *SetZoneRecordStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.SetZoneRecordStatusWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OverWrite)) {
		query["OverWrite"] = request.OverWrite
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
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
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
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateRecordRemarkWithOptions(request *UpdateRecordRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecordRemark"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecordRemark(request *UpdateRecordRemarkRequest) (_result *UpdateRecordRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.UpdateRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResolverEndpointWithOptions(request *UpdateResolverEndpointRequest, runtime *util.RuntimeOptions) (_result *UpdateResolverEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointId)) {
		query["EndpointId"] = request.EndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.IpConfig)) {
		query["IpConfig"] = request.IpConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResolverEndpoint"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResolverEndpoint(request *UpdateResolverEndpointRequest) (_result *UpdateResolverEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResolverEndpointResponse{}
	_body, _err := client.UpdateResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResolverRuleWithOptions(request *UpdateResolverRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateResolverRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForwardIp)) {
		query["ForwardIp"] = request.ForwardIp
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResolverRule"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResolverRule(request *UpdateResolverRuleRequest) (_result *UpdateResolverRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResolverRuleResponse{}
	_body, _err := client.UpdateResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSyncEcsHostTaskWithOptions(request *UpdateSyncEcsHostTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateSyncEcsHostTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSyncEcsHostTask"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSyncEcsHostTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSyncEcsHostTask(request *UpdateSyncEcsHostTaskRequest) (_result *UpdateSyncEcsHostTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSyncEcsHostTaskResponse{}
	_body, _err := client.UpdateSyncEcsHostTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateZoneRecordWithOptions(request *UpdateZoneRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateZoneRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["Line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RecordId)) {
		query["RecordId"] = request.RecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Rr)) {
		query["Rr"] = request.Rr
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		query["Ttl"] = request.Ttl
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		query["Weight"] = request.Weight
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateZoneRecord"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateZoneRecord(request *UpdateZoneRecordRequest) (_result *UpdateZoneRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.UpdateZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateZoneRemarkWithOptions(request *UpdateZoneRemarkRequest, runtime *util.RuntimeOptions) (_result *UpdateZoneRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateZoneRemark"),
		Version:     tea.String("2018-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateZoneRemark(request *UpdateZoneRemarkRequest) (_result *UpdateZoneRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.UpdateZoneRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
