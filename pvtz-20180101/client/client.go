// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddResolverEndpointRequest struct {
	// The source IP addresses of outbound traffic. You must add two to six source IP addresses to ensure high availability.
	//
	// This parameter is required.
	IpConfig []*AddResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The endpoint name. The name can be up to 20 characters in length. If the upper limit is exceeded, an error message is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kqlqlqjqqkq
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The outbound VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-129343jslslsks
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The region ID of the outbound virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
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
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The source IP address of outbound traffic. The IP address must be within the specified CIDR block.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sjqkql
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
	// The outbound endpoint ID.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32436208-E1AF-4DAB-B3B8-24F5F25B0950
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The destination IP address and port number.
	//
	// This parameter is required.
	ForwardIp []*AddResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the forwarding rule. Valid value:
	//
	// 	- OUTBOUND: forwards Domain Name System (DNS) requests to one or more external IP addresses.
	//
	// example:
	//
	// OUTBOUND
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the forward zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
	// The destination IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 725B8BED-901F-480C-BBAC-FA59A18580C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rule ID.
	//
	// example:
	//
	// hra0**
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// The authorization method. Valid values:
	//
	// 	- AUTH_CODE: An authorization code is used to associate VPCs across accounts. The system checks whether the value of AuthCode is valid.
	//
	// 	- RESOURCE_DIRECTORY: A resource directory is used to associate VPCs across accounts. The system checks whether the value of AuthorizedUserId and the current account are in the same resource directory.
	//
	// 	- If this parameter is empty, an authorization code is used to associate VPCs across accounts.
	//
	// example:
	//
	// AUTH_CODE
	AuthChannel *string `json:"AuthChannel,omitempty" xml:"AuthChannel,omitempty"`
	// The verification code.
	//
	// This parameter is required when AuthType is set to NORMAL or is left empty and AuthChannel is set to AUTH_CODE or is left empty.
	//
	// example:
	//
	// 123456
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization.
	//
	// 	- CLOUD_PRODUCT: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111222333
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical location of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- Normal zone: regular module
	//
	// 	- Fast Zone: acceleration module
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 	- Specifies whether to enable the recursive resolution proxy feature for the zone. Valid values: **ZONE**: disables the recursive resolution proxy feature for the zone.
	//
	// 	- **RECORD**: enables the recursive resolution proxy feature for the zone.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-resourcegroupid1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// This parameter is not available. You can ignore it.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// This parameter is not available. You can ignore it.
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The global ID of the zone.
	//
	// example:
	//
	// AgIDE1MQ_151
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// example.com
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line. Default value: **default**.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record. Valid values: **1 to 99**.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname.
	//
	// For example, you must set Rr to @ if you want to resolve @.example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL) of the DNS record. Default value: **60**.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. Valid values: **A**, **AAAA**, **CNAME**, **TXT**, **MX**, **PTR**, and **SRV**.
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 2.2.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The record value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the address. Valid values: **0 to 100**. Default value: 1.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// CAgICA1OA_58
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
	// The ID of the DNS record.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The VPCs.
	Vpc []*BindResolverRuleVpcRequestVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-kqk1i2o2ajsksl-vpc-test
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
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
	// The request ID.
	//
	// example:
	//
	// 12FE6E98-3885-423E-B18B-88CC17052A31
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The information about VPCs.
	Vpcs []*BindZoneVpcRequestVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE0OQ_149
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID. If you do not specify this parameter, the VPCs that are bound to the zone are unbound from the zone.
	//
	// example:
	//
	// daily-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The type of the VPC. Valid values:
	//
	// 	- **STANDARD**: standard VPC
	//
	// 	- **EDS**: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The name of the zone. This parameter is required.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
	// Indicates whether the zone name is valid. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA29B88F-A571-4123-80D5-768AC2F7F806
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 35134B4A-CEC0-43C8-AAD4-BA54AE3268B5
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
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
	// The request ID.
	//
	// example:
	//
	// 0C9959BE-3A6A-4803-8DCE-973B42ACD599
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
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization
	//
	// 	- NORMAL: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111111
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE1MA_150
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
	// The request ID.
	//
	// example:
	//
	// E246E023-F2EB-4034-83F7-B13FCF31459C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global ID of the zone.
	//
	// example:
	//
	// AgIDE1MA_150
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
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
	// The ID of the DNS record.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
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
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516779348000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The type of operation logs. Valid values:
	//
	// 	- **PV_ZONE**: the logs that record the operations on zones
	//
	// 	- **PV_RECORD**: the logs that record the operations on DNS records
	//
	// If you set this parameter to other values, all types of operation logs are queried.
	//
	// example:
	//
	// PV_ZONE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The keyword for searches in "%KeyWord%" mode. The value is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516779348000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The global ID of the zone.\\
	//
	// If you specify this parameter, the logs that record the operations on the Domain Name System (DNS) records of the specified zone are queried.\\
	//
	// If you leave this parameter empty, the logs that record the operations on all zones that belong to the current Alibaba Cloud account and the DNS records of these zones are queried.
	//
	// example:
	//
	// 6726
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The operation logs.
	ChangeLogs *DescribeChangeLogsResponseBodyChangeLogs `json:"ChangeLogs,omitempty" xml:"ChangeLogs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0FCB52A-D512-41A0-8595-40234EDCFD8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 100
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The details of the operation.
	//
	// example:
	//
	// add test-api.com
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 13270376
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The subtype of the operator. Valid values:
	//
	// 	- CUSTOMER: Alibaba Cloud account
	//
	// 	- SUB: RAM user
	//
	// 	- STS: assumed role that obtains the Security Token Service (STS) token of a RAM role
	//
	// 	- OTHER: other types
	//
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// The operator type. No value or **USER*	- is returned for this parameter.
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The ID of the object on which the operation was performed.
	//
	// example:
	//
	// CAgICA1OA_58
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The name of the object on which the operation was performed.
	//
	// example:
	//
	// test-api.com
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The log ID.
	//
	// example:
	//
	// 6726
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation type.
	//
	// example:
	//
	// add
	OperAction *string `json:"OperAction,omitempty" xml:"OperAction,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	OperIp *string `json:"OperIp,omitempty" xml:"OperIp,omitempty"`
	// The type of the object on which the operation is performed.
	//
	// example:
	//
	// PV_ZONE
	OperObject *string `json:"OperObject,omitempty" xml:"OperObject,omitempty"`
	// The time when the operation is performed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-24T07:35Z
	OperTime *string `json:"OperTime,omitempty" xml:"OperTime,omitempty"`
	// The time when the operation was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516779348000
	OperTimestamp *int64 `json:"OperTimestamp,omitempty" xml:"OperTimestamp,omitempty"`
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
	// The supported language. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// 	- ja: Japanese
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the Alibaba Cloud account to which the permissions on the resources are granted.
	//
	// example:
	//
	// 111222333
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The scenario. Valid values:
	//
	// 	- AUTH: the built-in authoritative module
	//
	// 	- FWD: the forward module
	//
	// 	- RA: the traffic analysis module
	//
	// example:
	//
	// AUTH
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.168.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The type of the virtual private cloud (VPC). Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
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
	// The regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AF7D4DCE-0776-47F2-A9B2-6FB85A87AA60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The display name of the region, which varies based on the current language.
	//
	// example:
	//
	// China (Beijing)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the service in the region.
	//
	// example:
	//
	// pvtz.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Beijing)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
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
	// The business ID. BizId is specified together with BizType.
	//
	// 	- If you set BizType to AUTH_ZONE, set BizId to a zone ID.
	//
	// 	- If you set BizType to RESOLVER_RULE, set BizId to the ID of a forwarding rule.
	//
	// example:
	//
	// b9c93a8954c4098731e863c04302f45a
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The business type. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- RESOLVER_RULE: forwarding rule
	//
	// example:
	//
	// AUTH_ZONE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1571673600000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1571587200000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-1111
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The global ID of the zone. To query the number of DNS requests for a zone, you can specify ZoneId or BizType and BizId.
	//
	// example:
	//
	// 29c752a01cd281a20ddcfaecef
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
	// The information about the DNS requests.
	RequestDetails *DescribeRequestGraphResponseBodyRequestDetails `json:"RequestDetails,omitempty" xml:"RequestDetails,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EB71815-A421-4E51-8E8D-667F44ABE633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The number of DNS requests.
	//
	// example:
	//
	// 103
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The statistical time. The value is a string. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-21T10:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The statistical timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1571652000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	// The zone ID.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
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
	// The information about the queried zones.
	AvailableZones []*DescribeResolverAvailableZonesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 23268E49-0C3E-4A2C-AB70-B4C7D092470B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The zone ID.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The state of resources in the zone. Valid values:
	//
	// 	- NORMAL: The resources are in the normal state.
	//
	// 	- SOLD_OUT: The resources are sold out.
	//
	// example:
	//
	// NORMAL
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
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// The time when the endpoint was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:45:56
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the endpoint was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608356000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hra0**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source IP address of outbound traffic.
	IpConfigs []*DescribeResolverEndpointResponseBodyIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	// The endpoint name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 45020ED9-6319-4CA7-9475-6E8D6446E84F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-8vb3sigz86xc-group-test
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- SUCCESS: The endpoint works as expected.
	//
	// 	- INIT: The endpoint is being created.
	//
	// 	- FAILED: The endpoint fails to be created.
	//
	// 	- CHANGE_INIT: The endpoint is being modified.
	//
	// 	- CHANGE_FAILED: The endpoint fails to be modified.
	//
	// 	- EXCEPTION: The endpoint encounters an exception.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the endpoint was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:48:39
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the endpoint was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608519000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The outbound VPC ID.
	//
	// example:
	//
	// vpc-8vbl8mpum-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The outbound VPC name.
	//
	// example:
	//
	// vpc-name-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The ID of the region where the outbound VPC resides.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The name of the region where the outbound virtual private cloud (VPC) resides.
	VpcRegionName *string `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
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
	// The ID of the zone where the vSwitch resides.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IPv4 address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbmks7hzrmk-vswitch-id
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
	// The keyword used to filter endpoints in %keyword% mode.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The state of the endpoint that you want to query. If you do not specify this parameter, all endpoints are returned. Valid values:
	//
	// 	- SUCCESS: The endpoint works as expected.
	//
	// 	- INIT: The endpoint is being created.
	//
	// 	- FAILED: The endpoint fails to be created.
	//
	// 	- CHANGE_INIT: The endpoint is being modified.
	//
	// 	- CHANGE_FAILED: The endpoint fails to be modified.
	//
	// 	- EXCEPTION: The endpoint encounters an exception.
	//
	// example:
	//
	// SUCCESS
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
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

func (s *DescribeResolverEndpointsRequest) SetVpcRegionId(v string) *DescribeResolverEndpointsRequest {
	s.VpcRegionId = &v
	return s
}

type DescribeResolverEndpointsResponseBody struct {
	// The information about endpoints.
	Endpoints []*DescribeResolverEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83D1682B-B69A-4060-9FA8-2907C2A35600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The time when the endpoint was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:36:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the endpoint was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594607786000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hra0**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source IP addresses of outbound traffic.
	IpConfigs []*DescribeResolverEndpointsResponseBodyEndpointsIpConfigs `json:"IpConfigs,omitempty" xml:"IpConfigs,omitempty" type:"Repeated"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-8vb3sigz86xc-test-group
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the endpoint. Valid values:
	//
	// 	- SUCCESS: The endpoint works as expected.
	//
	// 	- INIT: The endpoint is being created.
	//
	// 	- FAILED: The endpoint fails to be created.
	//
	// 	- CHANGE_INIT: The endpoint is being modified.
	//
	// 	- CHANGE_FAILED: The endpoint fails to be modified.
	//
	// 	- EXCEPTION: The endpoint encounters an exception.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the endpoint was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:38:24
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the endpoint was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594607904000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The ID of the outbound virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-8vbl8mpum-test-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// vpc-test-name
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The region ID of the outbound VPC.
	//
	// example:
	//
	// cn-zhangjiakou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	// The name of the region where the outbound VPC resides.
	//
	// example:
	//
	// China East 1 (Hangzhou)
	VpcRegionName *string `json:"VpcRegionName,omitempty" xml:"VpcRegionName,omitempty"`
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
	// The ID of the zone where the vSwitch resides.
	//
	// example:
	//
	// cn-zhangjiakou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IPv4 address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbmks7h-test-vswitchId
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra1**
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
	// The virtual private clouds (VPCs) that are associated with the forwarding rule.
	BindVpcs []*DescribeResolverRuleResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	// The time when the forwarding rule was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the forwarding rule was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The destination IP addresses.
	ForwardIps []*DescribeResolverRuleResponseBodyForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
	// The forwarding rule ID.
	//
	// example:
	//
	// hra1**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// forward rule-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 13D5113B-7E34-407F-A9C1-D96CD2B04277
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the forwarding rule. Valid value:
	//
	// 	- OUTBOUND: forwards Domain Name System (DNS) requests to one or more external IP addresses.
	//
	// example:
	//
	// OUTBOUND
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the forwarding rule was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the forwarding rule was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The name of the forward zone.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vbl8m-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// vpc-name-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The type of the VPC. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 324542413
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
	// The IP address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// The ID of the outbound endpoint.
	//
	// example:
	//
	// hra2**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The keyword used to filter forwarding rules in %keyword% mode.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to return additional information. Default value: false.
	//
	// 	- If you set this parameter to true, additional information, such as the virtual private clouds (VPCs) that are associated with the queried forwarding rule, is returned.
	//
	// 	- If you set this parameter to false, no additional information is returned.
	//
	// example:
	//
	// true
	NeedDetailAttributes *bool `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A10E03D7-808C-422D-9144-F8586C2E2297
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The forwarding rules.
	Rules []*DescribeResolverRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The VPCs associated with the forwarding rule.
	BindVpcs []*DescribeResolverRulesResponseBodyRulesBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Repeated"`
	// The time when the forwarding was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the forwarding rule was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The destination IP addresses.
	ForwardIps []*DescribeResolverRulesResponseBodyRulesForwardIps `json:"ForwardIps,omitempty" xml:"ForwardIps,omitempty" type:"Repeated"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// hra1**
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// forward rule-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the forwarding rule. Valid value:
	//
	// 	- OUTBOUND: Domain Name System (DNS) requests are forwarded to one or more IP addresses.
	//
	// example:
	//
	// OUTBOUND
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the forwarding rule was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-07-13 10:51:44
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp when the forwarding rule was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1594608704000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The name of the forward zone.
	//
	// example:
	//
	// example.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// ap-southeast-1
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vbl8mpum-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// vpc-name-test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The type of the virtual private cloud (VPC). Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 121098702443**
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
	// The IP address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
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
	// The request ID.
	//
	// example:
	//
	// A73F3BD0-B1A8-42A9-A9B6-689BBABC4891
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2254
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The top 3 virtual private clouds (VPCs) that initiate the largest number of DNS requests.
	VpcRequestTops *DescribeStatisticSummaryResponseBodyVpcRequestTops `json:"VpcRequestTops,omitempty" xml:"VpcRequestTops,omitempty" type:"Struct"`
	// The top 3 zones with the largest number of DNS requests.
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
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China (Beijing)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The number of DNS requests.
	//
	// example:
	//
	// 2254
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// 46574
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zeisd8c0j6wk1451jr6o
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The type of the VPC. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
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
	// The business type. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- RESOLVER_RULE: forwarding rule
	//
	// example:
	//
	// AUTH_ZONE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The number of DNS requests.
	//
	// example:
	//
	// 2251
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// host.local
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pvtz-test-id-2989149d628c56f00e
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
	// The information about regions.
	EcsRegions *DescribeSyncEcsHostTaskResponseBodyEcsRegions `json:"EcsRegions,omitempty" xml:"EcsRegions,omitempty" type:"Struct"`
	// The information about the regions within the current account.
	Regions *DescribeSyncEcsHostTaskResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- ON
	//
	// 	- OFF
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the task was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// pvtz-test-id-2989149d628c56f00e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The region IDs.
	RegionIds *DescribeSyncEcsHostTaskResponseBodyEcsRegionsEcsRegionRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	// The Alibaba Cloud account to which the region belongs. This parameter is used in cross-account synchronization scenarios.
	//
	// example:
	//
	// 1234567890
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type. Valid value: ZONE.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9EA7F720-B7C0-45C1-9CF4-B6A5A1179B68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the resources.
	Tags []*DescribeTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of tags added to the resources.
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
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization.
	//
	// 	- CLOUD_PRODUCT: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 111222333
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46973D4C-E3E4-4ABA-9190-9A9DE406C7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 5
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The information about the Alibaba Cloud accounts.
	Users []*DescribeUserVpcAuthorizationsResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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
	// The authorization scope. Valid value:
	//
	// 	- NORMAL: general authorization.
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the Alibaba Cloud account.
	//
	// example:
	//
	// alidns***@test.com
	AuthorizedAliyunId *string `json:"AuthorizedAliyunId,omitempty" xml:"AuthorizedAliyunId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 111222333
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The time when the authorization was performed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-08T02:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the authorization was performed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672740294000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE1MA_149
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
	// The virtual private clouds (VPCs) bound to the zone.
	BindVpcs *DescribeZoneInfoResponseBodyBindVpcs `json:"BindVpcs,omitempty" xml:"BindVpcs,omitempty" type:"Struct"`
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-23T03:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 2312234523451342
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The type of the operator.
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The logical location of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- NORMAL_ZONE: regular module
	//
	// 	- FAST_ZONE: acceleration module
	//
	// example:
	//
	// FAST_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// 	- Indicates whether the zone is a reverse lookup zone. Valid values: true and false. The value true indicates that the zone is a reverse lookup zone.
	//
	// 	- The value false indicates that the zone is not a reverse lookup zone.
	//
	// example:
	//
	// false
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// 	- Indicates whether the recursive resolution proxy feature is enabled for the zone. Valid values: **ZONE**: The recursive resolution proxy feature is disabled for the zone.
	//
	// 	- **RECORD**: The recursive resolution proxy feature is enabled for the zone.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The total number of DNS records.
	//
	// example:
	//
	// 2
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The description of the zone.
	//
	// example:
	//
	// specialZone
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F73F41A3-B6DD-42CA-A793-FFF93277835D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the secondary Domain Name System (DNS) feature is enabled for the zone. Valid values:
	//
	// 	- **true**: The secondary DNS feature is enabled.
	//
	// 	- **false**: The secondary DNS feature is disabled.
	//
	// example:
	//
	// true
	SlaveDns *bool `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
	// The time when the zone was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-24T06:35Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the zone was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516775741000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The global ID of the zone.
	//
	// example:
	//
	// AgIDE0OQ_149<
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name.
	//
	// example:
	//
	// test.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// 	- If ZoneType is set to AUTH_ZONE, no value is returned for this parameter.
	//
	// 	- If ZoneType is set to CLOUD_PRODUCT_ZONE, the type of the cloud service is returned.
	//
	// example:
	//
	// pvtz
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The type of the zone. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- CLOUD_PRODUCT_ZONE: authoritative zone for cloud services
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// 1304
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// daily-vpc-id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// daily-vpc-name
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The type of the VPC. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
	// The ID of the user to which the VPC belongs. The value null indicates that the VPC belongs to the current user.
	//
	// example:
	//
	// vpc-bp1aevy8sofi8mh1q****
	VpcUserId *int64 `json:"VpcUserId,omitempty" xml:"VpcUserId,omitempty"`
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
	// The hostname keyword based on which the system queries the DNS records.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode. Valid values:
	//
	// 	- **LIKE**: fuzzy search
	//
	// 	- **EXACT (default)**: exact search
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The tags added to the DNS record.
	//
	// 	- This parameter is left empty by default. In this case, the DNS records of the zone are queried.
	//
	// 	- If you set Tag to ecs, the DNS records added to the hostnames of Elastic Compute Service (ECS) instances in the zone are queried.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// CAgICA1OA_58
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned DNS records.
	Records *DescribeZoneRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B07FBC3-3A53-4939-A3C6-2BDFE407BAB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 100
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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
	// The time when the DNS record was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-14T03:47Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the DNS record was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1672740294000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record.
	//
	// example:
	//
	// 60
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record.
	//
	// example:
	//
	// 5809
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// xxx
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- ENABLE: The DNS record is enabled.
	//
	// 	- DISABLE: The DNS record is disabled.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time-to-live (TTL) of the DNS record.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the DNS record was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-08T02:31Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1654777678000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The record value.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the address.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// a49f55537f3b0b1e6e43add0bf5f0033
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
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
	// The request ID.
	//
	// example:
	//
	// 7B07FBC3-3A53-4939-A3C6-2BDFE407BAB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones.
	Zones *DescribeZoneVpcTreeResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-09-18T08:20Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1568794812000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 5463564356
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The operator type.
	//
	// example:
	//
	// USER
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	// The logical location of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- NORMAL_ZONE: regular module
	//
	// 	- FAST_ZONE: acceleration module
	//
	// example:
	//
	// NORMAL_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// Indicates whether the zone is a reverse lookup zone. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// The number of Domain Name System (DNS) records.
	//
	// example:
	//
	// 1
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The description of the zone.
	//
	// example:
	//
	// demo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The time when the zone was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-09-18T08:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the zone was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1568794834000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The VPCs bound to the zones.
	Vpcs *DescribeZoneVpcTreeResponseBodyZonesZoneVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
	// The global ID of the zone.
	//
	// example:
	//
	// 6d83e3b31aa60ca4aaa7161f1b6baa95
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// localzone.demo
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The type of the cloud service.
	//
	// 	- If the value of the ZoneType parameter is AUTH_ZONE, no value is returned for this parameter.
	//
	// 	- If the value of the ZoneType parameter is CLOUD_PRODUCT_ZONE, the type of the cloud service is returned.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The type of the zone. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- CLOUD_PRODUCT_ZONE: authoritative zone for cloud services
	//
	// example:
	//
	// AUTH_ZONE
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
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China North 2
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2z21341ssdadsfzyd49ra
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// demo-vpc
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The type of the VPC. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
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
	// The keyword of the zone name. The search is performed in the %KeyWord % mode and is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	QueryRegionId *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-xxxxx
	QueryVpcId *string `json:"QueryVpcId,omitempty" xml:"QueryVpcId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag added to the resource.
	ResourceTag []*DescribeZonesRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The search mode. Valid values:
	//
	// 	- **LIKE (default)**: fuzzy search
	//
	// 	- **EXACT**: exact search
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// BLINK
	ZoneTag []*string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty" type:"Repeated"`
	// The type of zones to query. Default value: AUTH_ZONE.
	//
	// Valid values:
	//
	// 	- **AUTH_ZONE**: authoritative zone
	//
	// 	- **CLOUD_PRODUCT_ZONE**: authoritative zone for cloud services
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
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
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// daily
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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 3
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The zones.
	Zones *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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
	// The time when the zone was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the zone was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1514466483000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The creator of the zone.
	//
	// example:
	//
	// 5463564356
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The type of the user account.
	//
	// 	- **CUSTOMER**: Alibaba Cloud account
	//
	// 	- **SUB**: RAM user
	//
	// 	- **STS**: assumed role that obtains the Security Token Service (STS) token of a RAM role
	//
	// 	- **OTHER**: other types
	//
	// example:
	//
	// SUB
	CreatorSubType *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	// The logical location type of the built-in authoritative module in which the zone is added. Valid values:
	//
	// 	- **NORMAL_ZONE**: regular module
	//
	// 	- **FAST_ZONE**: acceleration module
	//
	// example:
	//
	// NORMAL_ZONE
	DnsGroup *string `json:"DnsGroup,omitempty" xml:"DnsGroup,omitempty"`
	// Indicates whether the zone is being removed to another logical location. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DnsGroupChanging *bool `json:"DnsGroupChanging,omitempty" xml:"DnsGroupChanging,omitempty"`
	// Indicates whether the zone is a reverse lookup zone. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsPtr *bool `json:"IsPtr,omitempty" xml:"IsPtr,omitempty"`
	// Indicates whether the recursive resolution proxy feature is enabled for the zone. Valid values:
	//
	// 	- **ZONE**: The recursive resolution proxy feature is disabled for the zone.
	//
	// 	- **RECORD**: The recursive resolution proxy feature is enabled for the zone.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The number of Domain Name System (DNS) records.
	//
	// example:
	//
	// 2
	RecordCount *int32 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The description of the zone.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the resources.
	ResourceTags *DescribeZonesResponseBodyZonesZoneResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Struct"`
	// The time when the zone was last modified. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-03T08:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the DNS record was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1514969843000
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// 6d83e3b31aa60ca4aaa7161f1b6b**95
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// example:
	//
	// test.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
	// The type of the cloud service. Valid values:
	//
	// 	- If ZoneType is set to AUTH_ZONE, no value is returned for this parameter.
	//
	// 	- If ZoneType is set to CLOUD_PRODUCT_ZONE, the type of the cloud service is returned.
	//
	// example:
	//
	// BLINK
	ZoneTag *string `json:"ZoneTag,omitempty" xml:"ZoneTag,omitempty"`
	// The type of zones. Valid values:
	//
	// 	- **AUTH_ZONE**: authoritative zone
	//
	// 	- **CLOUD_PRODUCT_ZONE**: authoritative zone for cloud services
	//
	// example:
	//
	// CLOUD_PRODUCT_ZONE
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
	// The key of tag N added to the zone.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the zone.
	//
	// example:
	//
	// daily
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
	// The language of the values for specific response parameters. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 234235354
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs, which are zone IDs. You can specify up to 50 zone IDs.
	//
	// example:
	//
	// 97fe9321a476d0861f624d3f738dcc38
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value: ZONE.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of entries per page. Valid values: `1 to 200`. Default value: 20.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The tags added to the resources.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// daily
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
	// The pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// 234235354
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the resources.
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
	// The resource ID, which is a zone ID.
	//
	// example:
	//
	// 97fe9321a476d0861f624d3f738dcc38
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ZONE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N added to the resource.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// daily
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the values for specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzzk7hx3glaoq
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE1MA_149
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
	// The request ID.
	//
	// example:
	//
	// D1324D48-1E23-4AEF-9EDE-466120561C6F
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to enable the recursive resolution proxy feature for the zone. Valid values:
	//
	// 	- **ZONE**: disables the recursive resolution proxy feature for the zone.
	//
	// 	- **RECORD**: enables the recursive resolution proxy feature for the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ProxyPattern *string `json:"ProxyPattern,omitempty" xml:"ProxyPattern,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The global ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE0OQ_149
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global ID of the zone.
	//
	// example:
	//
	// AgIDE0OQ_149
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5809
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- ENABLE: enables the DNS record.
	//
	// 	- DISABLE: suspends the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.XX.XX
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
	// The ID of the DNS record.
	//
	// example:
	//
	// 5809
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39CB16E5-4180-49F2-A060-23C0ECEB80D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the DNS record.
	//
	// example:
	//
	// DISABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to replace the original tags added to the resources. Valid values:
	//
	// 	- True: replaces the original tags.
	//
	// 	- False|Null: appends the specified one or more tags to the original tags. If a new tag has the same key but a different value from an original tag, the new tag replaces the original tag.
	//
	// example:
	//
	// true
	OverWrite *bool `json:"OverWrite,omitempty" xml:"OverWrite,omitempty"`
	// The resource IDs, which are zone IDs. You can specify **1 to 50*	- IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97fe9321a476d0861f624d3f738dcc38
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to add to the resources.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of tag N to add to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource.
	//
	// example:
	//
	// daily
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
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
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
	// Specifies whether to remove all tags from the specified one or more resources. This parameter is valid only if the TagKey parameter is left empty. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resource IDs, which are zone IDs. You can specify up to 50 zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97fe9321a476d0861f624d3f738dcc38
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value: ZONE.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZONE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of tags that you want to remove. You can specify up to 20 tag keys.
	//
	// example:
	//
	// env
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
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18954952
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// test record
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
	// The ID of the DNS record.
	//
	// example:
	//
	// 18954952
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
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
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The source IP addresses of outbound traffic. You must add two to six source IP addresses to ensure high availability.
	IpConfig []*UpdateResolverEndpointRequestIpConfig `json:"IpConfig,omitempty" xml:"IpConfig,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-test-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	AzId *string `json:"AzId,omitempty" xml:"AzId,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// sjqkql
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
	// The request ID.
	//
	// example:
	//
	// EC0BDA3A-A92A-4AC8-B351-322A9C79D5C5
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
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The destination IP address and port number.
	ForwardIp []*UpdateResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// forward rule-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The forwarding rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// hra0**
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateResolverRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleRequest) SetEndpointId(v string) *UpdateResolverRuleRequest {
	s.EndpointId = &v
	return s
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
	// The destination IP address.
	//
	// example:
	//
	// 172.16.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0C9959BE-3A6A-4803-8DCE-973B42ACD599
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The information about regions to be synchronized.
	//
	// This parameter is required.
	Region []*UpdateSyncEcsHostTaskRequestRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
	// The state of the task. Valid values:
	//
	// 	- ON
	//
	// 	- OFF
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test79afafec***********1d28f7889c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud account to which the region belongs. This parameter is used in cross-account synchronization scenarios.
	//
	// example:
	//
	// 1234567890
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// test-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the task was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record. Valid values: **1 to 99**.
	//
	// This parameter is required if the type of the DNS record is MX.
	//
	// example:
	//
	// 60
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5809
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The hostname.
	//
	// For example, you must set this parameter to @ if you want to resolve @.example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time-to-live (TTL) of the DNS record.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. Valid values: **A**, **AAAA**, **CNAME**, **TXT**, **MX**, **PTR**, and **SRV**.
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 2.2.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The record value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the address. Valid values: **1 to 100**.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
	// The ID of the DNS record.
	//
	// example:
	//
	// 5809
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 250E2C38-D0AD-4518-851D-1C1055805F82
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new description.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 1.1.1.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The unique ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// AgIDE1MA_149
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
	// The request ID.
	//
	// example:
	//
	// C6F1D541-E7A6-447A-A2B5-9F7A20B2A8FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// AgIDE1MA_149
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

// Summary:
//
// Creates an endpoint.
//
// @param request - AddResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverEndpointResponse
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

// Summary:
//
// Creates an endpoint.
//
// @param request - AddResolverEndpointRequest
//
// @return AddResolverEndpointResponse
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

// Summary:
//
// Creates a forwarding rule.
//
// @param request - AddResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverRuleResponse
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

// Summary:
//
// Creates a forwarding rule.
//
// @param request - AddResolverRuleRequest
//
// @return AddResolverRuleResponse
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

// Summary:
//
// Adds another account to associate one or more virtual private clouds (VPCs) of the current account with a private zone.
//
// @param request - AddUserVpcAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserVpcAuthorizationResponse
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

// Summary:
//
// Adds another account to associate one or more virtual private clouds (VPCs) of the current account with a private zone.
//
// @param request - AddUserVpcAuthorizationRequest
//
// @return AddUserVpcAuthorizationResponse
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

// Summary:
//
// Creates a zone.
//
// @param request - AddZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneResponse
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

// Summary:
//
// Creates a zone.
//
// @param request - AddZoneRequest
//
// @return AddZoneResponse
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

// Summary:
//
// Adds a Domain Name System (DNS) record for a zone.
//
// @param request - AddZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneRecordResponse
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

// Summary:
//
// Adds a Domain Name System (DNS) record for a zone.
//
// @param request - AddZoneRecordRequest
//
// @return AddZoneRecordResponse
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

// Summary:
//
// Associates a forwarding rule with virtual private clouds (VPCs).
//
// @param request - BindResolverRuleVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindResolverRuleVpcResponse
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

// Summary:
//
// Associates a forwarding rule with virtual private clouds (VPCs).
//
// @param request - BindResolverRuleVpcRequest
//
// @return BindResolverRuleVpcResponse
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

// Summary:
//
// Binds a zone to virtual private clouds (VPCs) or unbinds a zone from VPCs.
//
// @param request - BindZoneVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindZoneVpcResponse
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

// Summary:
//
// Binds a zone to virtual private clouds (VPCs) or unbinds a zone from VPCs.
//
// @param request - BindZoneVpcRequest
//
// @return BindZoneVpcResponse
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

// Summary:
//
// Checks whether the name of a zone is valid based on specific rules.
//
// @param request - CheckZoneNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckZoneNameResponse
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

// Summary:
//
// Checks whether the name of a zone is valid based on specific rules.
//
// @param request - CheckZoneNameRequest
//
// @return CheckZoneNameResponse
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

// Summary:
//
// Deletes an endpoint.
//
// @param request - DeleteResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverEndpointResponse
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

// Summary:
//
// Deletes an endpoint.
//
// @param request - DeleteResolverEndpointRequest
//
// @return DeleteResolverEndpointResponse
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

// Summary:
//
// Deletes a forwarding rule.
//
// @param request - DeleteResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverRuleResponse
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

// Summary:
//
// Deletes a forwarding rule.
//
// @param request - DeleteResolverRuleRequest
//
// @return DeleteResolverRuleResponse
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

// Summary:
//
// Deletes an account whose one or more virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DeleteUserVpcAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserVpcAuthorizationResponse
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

// Summary:
//
// Deletes an account whose one or more virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DeleteUserVpcAuthorizationRequest
//
// @return DeleteUserVpcAuthorizationResponse
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

// Summary:
//
// Deletes a zone.
//
// @param request - DeleteZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneResponse
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

// Summary:
//
// Deletes a zone.
//
// @param request - DeleteZoneRequest
//
// @return DeleteZoneResponse
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

// Summary:
//
// Deletes a Domain Name System (DNS) record of a zone.
//
// @param request - DeleteZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneRecordResponse
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

// Summary:
//
// Deletes a Domain Name System (DNS) record of a zone.
//
// @param request - DeleteZoneRecordRequest
//
// @return DeleteZoneRecordResponse
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

// Summary:
//
// Queries a list of operation logs.
//
// @param request - DescribeChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChangeLogsResponse
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

// Summary:
//
// Queries a list of operation logs.
//
// @param request - DescribeChangeLogsRequest
//
// @return DescribeChangeLogsResponse
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

// Summary:
//
// Queries a list of available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries a list of available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries the information about Domain Name System (DNS) requests.
//
// @param request - DescribeRequestGraphRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRequestGraphResponse
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

// Summary:
//
// Queries the information about Domain Name System (DNS) requests.
//
// @param request - DescribeRequestGraphRequest
//
// @return DescribeRequestGraphResponse
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

// Summary:
//
// Queries a list of available zones.
//
// @param request - DescribeResolverAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverAvailableZonesResponse
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

// Summary:
//
// Queries a list of available zones.
//
// @param request - DescribeResolverAvailableZonesRequest
//
// @return DescribeResolverAvailableZonesResponse
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

// Summary:
//
// Queries the information about an endpoint.
//
// @param request - DescribeResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointResponse
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

// Summary:
//
// Queries the information about an endpoint.
//
// @param request - DescribeResolverEndpointRequest
//
// @return DescribeResolverEndpointResponse
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

// Summary:
//
// Queries a list of endpoints.
//
// @param request - DescribeResolverEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointsResponse
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

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		query["VpcRegionId"] = request.VpcRegionId
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

// Summary:
//
// Queries a list of endpoints.
//
// @param request - DescribeResolverEndpointsRequest
//
// @return DescribeResolverEndpointsResponse
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

// Summary:
//
// Queries the information about a forwarding rule.
//
// @param request - DescribeResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRuleResponse
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

// Summary:
//
// Queries the information about a forwarding rule.
//
// @param request - DescribeResolverRuleRequest
//
// @return DescribeResolverRuleResponse
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

// Summary:
//
// Queries a list of forwarding rules.
//
// @param request - DescribeResolverRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRulesResponse
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

// Summary:
//
// Queries a list of forwarding rules.
//
// @param request - DescribeResolverRulesRequest
//
// @return DescribeResolverRulesResponse
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

// Summary:
//
// Queries the statistics on the Domain Name System (DNS) requests received on the previous day.
//
// @param request - DescribeStatisticSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatisticSummaryResponse
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

// Summary:
//
// Queries the statistics on the Domain Name System (DNS) requests received on the previous day.
//
// @param request - DescribeStatisticSummaryRequest
//
// @return DescribeStatisticSummaryResponse
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

// Summary:
//
// Queries the information about a hostname synchronization task.
//
// @param request - DescribeSyncEcsHostTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyncEcsHostTaskResponse
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

// Summary:
//
// Queries the information about a hostname synchronization task.
//
// @param request - DescribeSyncEcsHostTaskRequest
//
// @return DescribeSyncEcsHostTaskResponse
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

// Summary:
//
// Queries a list of existing tags.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
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

// Summary:
//
// Queries a list of existing tags.
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
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

// Summary:
//
// Queries a list of accounts whose virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DescribeUserVpcAuthorizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserVpcAuthorizationsResponse
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

// Summary:
//
// Queries a list of accounts whose virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DescribeUserVpcAuthorizationsRequest
//
// @return DescribeUserVpcAuthorizationsResponse
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

// Summary:
//
// Queries the information about a zone.
//
// @param request - DescribeZoneInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneInfoResponse
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

// Summary:
//
// Queries the information about a zone.
//
// @param request - DescribeZoneInfoRequest
//
// @return DescribeZoneInfoResponse
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

// Summary:
//
// Queries a list of Domain Name System (DNS) records for a zone.
//
// @param request - DescribeZoneRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneRecordsResponse
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

// Summary:
//
// Queries a list of Domain Name System (DNS) records for a zone.
//
// @param request - DescribeZoneRecordsRequest
//
// @return DescribeZoneRecordsResponse
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

// Summary:
//
// Queries a list of zones and a list of virtual private clouds (VPCs) that are bound to the zones.
//
// Description:
//
// We recommend that you do not call this API operation due to its poor performance. Instead, you can call the DescribeZones operation to query a list of zones. If you want to query the information about VPCs to which a zone is bound, you can call the DescribeZoneInfo operation based on the zone ID.
//
// @param request - DescribeZoneVpcTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneVpcTreeResponse
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

// Summary:
//
// Queries a list of zones and a list of virtual private clouds (VPCs) that are bound to the zones.
//
// Description:
//
// We recommend that you do not call this API operation due to its poor performance. Instead, you can call the DescribeZones operation to query a list of zones. If you want to query the information about VPCs to which a zone is bound, you can call the DescribeZoneInfo operation based on the zone ID.
//
// @param request - DescribeZoneVpcTreeRequest
//
// @return DescribeZoneVpcTreeResponse
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

// Summary:
//
// Queries a list of zones for a user.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
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

// Summary:
//
// Queries a list of zones for a user.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
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

// Summary:
//
// Queries a list of tags added to one or more resources.
//
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

// Summary:
//
// Queries a list of tags added to one or more resources.
//
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

// Summary:
//
// Moves a zone to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
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

// Summary:
//
// Moves a zone to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
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

// Summary:
//
// Configures the recursive resolution proxy feature.
//
// @param request - SetProxyPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetProxyPatternResponse
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

// Summary:
//
// Configures the recursive resolution proxy feature.
//
// @param request - SetProxyPatternRequest
//
// @return SetProxyPatternResponse
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

// Summary:
//
// Specifies the status of a Domain Name System (DNS) record for a zone.
//
// @param request - SetZoneRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetZoneRecordStatusResponse
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

// Summary:
//
// Specifies the status of a Domain Name System (DNS) record for a zone.
//
// @param request - SetZoneRecordStatusRequest
//
// @return SetZoneRecordStatusResponse
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

// Summary:
//
// Adds tags to resources.
//
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

// Summary:
//
// Adds tags to resources.
//
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

// Summary:
//
// Removes tags from resources.
//
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

// Summary:
//
// Removes tags from resources.
//
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

// Summary:
//
// Modifies the description of a Domain Name System (DNS) record that is added for a zone.
//
// @param request - UpdateRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordRemarkResponse
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

// Summary:
//
// Modifies the description of a Domain Name System (DNS) record that is added for a zone.
//
// @param request - UpdateRecordRemarkRequest
//
// @return UpdateRecordRemarkResponse
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

// Summary:
//
// Modifies an endpoint.
//
// @param request - UpdateResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverEndpointResponse
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

// Summary:
//
// Modifies an endpoint.
//
// @param request - UpdateResolverEndpointRequest
//
// @return UpdateResolverEndpointResponse
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

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - UpdateResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverRuleResponse
func (client *Client) UpdateResolverRuleWithOptions(request *UpdateResolverRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateResolverRuleResponse, _err error) {
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

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - UpdateResolverRuleRequest
//
// @return UpdateResolverRuleResponse
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

// Summary:
//
// Creates and updates a hostname synchronize task.
//
// @param request - UpdateSyncEcsHostTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncEcsHostTaskResponse
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

// Summary:
//
// Creates and updates a hostname synchronize task.
//
// @param request - UpdateSyncEcsHostTaskRequest
//
// @return UpdateSyncEcsHostTaskResponse
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

// Summary:
//
// Modifies a Domain Name System (DNS) record of a zone.
//
// @param request - UpdateZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRecordResponse
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

// Summary:
//
// Modifies a Domain Name System (DNS) record of a zone.
//
// @param request - UpdateZoneRecordRequest
//
// @return UpdateZoneRecordResponse
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

// Summary:
//
// Modifies the description of a zone.
//
// @param request - UpdateZoneRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRemarkResponse
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

// Summary:
//
// Modifies the description of a zone.
//
// @param request - UpdateZoneRemarkRequest
//
// @return UpdateZoneRemarkResponse
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
