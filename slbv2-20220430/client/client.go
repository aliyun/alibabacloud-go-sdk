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

type AddServersToServerGroupRequest struct {
	ClientToken   *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool                                    `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerGroupId *string                                  `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s AddServersToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetRegionId(v string) *AddServersToServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

type AddServersToServerGroupRequestServers struct {
	// 服务器描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 服务器端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 后端权重
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 服务器对应的zoneId
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *AddServersToServerGroupRequestServers) SetZoneId(v string) *AddServersToServerGroupRequestServers {
	s.ZoneId = &v
	return s
}

type AddServersToServerGroupResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *AddServersToServerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddServersToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) SetCode(v string) *AddServersToServerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetData(v *AddServersToServerGroupResponseBodyData) *AddServersToServerGroupResponseBody {
	s.Data = v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetDynamicCode(v string) *AddServersToServerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetDynamicMessage(v string) *AddServersToServerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetHttpStatusCode(v int32) *AddServersToServerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetMessage(v string) *AddServersToServerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetSuccess(v bool) *AddServersToServerGroupResponseBody {
	s.Success = &v
	return s
}

type AddServersToServerGroupResponseBodyData struct {
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s AddServersToServerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBodyData) SetJobId(v string) *AddServersToServerGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *AddServersToServerGroupResponseBodyData) SetServerGroupId(v string) *AddServersToServerGroupResponseBodyData {
	s.ServerGroupId = &v
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

type CreateListenerRequest struct {
	// ca 证书列表
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	CaEnabled        *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// server证书列表
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// add 必选
	LoadBalancerId  *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Tclssl监听的安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// servergroupId
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetCaCertificateIds(v []*string) *CreateListenerRequest {
	s.CaCertificateIds = v
	return s
}

func (s *CreateListenerRequest) SetCaEnabled(v bool) *CreateListenerRequest {
	s.CaEnabled = &v
	return s
}

func (s *CreateListenerRequest) SetCertificateIds(v []*string) *CreateListenerRequest {
	s.CertificateIds = v
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

func (s *CreateListenerRequest) SetIdleTimeout(v int32) *CreateListenerRequest {
	s.IdleTimeout = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetListenerPort(v int32) *CreateListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateListenerRequest) SetListenerProtocol(v string) *CreateListenerRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetRegionId(v string) *CreateListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateListenerRequest) SetResourceGroupId(v string) *CreateListenerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateListenerRequest) SetSecurityPolicyId(v string) *CreateListenerRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateListenerRequest) SetServerGroupId(v string) *CreateListenerRequest {
	s.ServerGroupId = &v
	return s
}

type CreateListenerResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateListenerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetCode(v string) *CreateListenerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateListenerResponseBody) SetData(v *CreateListenerResponseBodyData) *CreateListenerResponseBody {
	s.Data = v
	return s
}

func (s *CreateListenerResponseBody) SetDynamicCode(v string) *CreateListenerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateListenerResponseBody) SetDynamicMessage(v string) *CreateListenerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateListenerResponseBody) SetHttpStatusCode(v int32) *CreateListenerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateListenerResponseBody) SetMessage(v string) *CreateListenerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateListenerResponseBody) SetSuccess(v bool) *CreateListenerResponseBody {
	s.Success = &v
	return s
}

type CreateListenerResponseBodyData struct {
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s CreateListenerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBodyData) SetJobId(v string) *CreateListenerResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateListenerResponseBodyData) SetListenerId(v string) *CreateListenerResponseBodyData {
	s.ListenerId = &v
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

type CreateLoadBalancerRequest struct {
	AddressIpVersion         *string                                  `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType              *string                                  `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	BillingConfig            *CreateLoadBalancerRequestBillingConfig  `json:"BillingConfig,omitempty" xml:"BillingConfig,omitempty" type:"Struct"`
	ClientToken              *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommonBandwidthPackageId *string                                  `json:"CommonBandwidthPackageId,omitempty" xml:"CommonBandwidthPackageId,omitempty"`
	DryRun                   *bool                                    `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EnableCrossZone          *bool                                    `json:"EnableCrossZone,omitempty" xml:"EnableCrossZone,omitempty"`
	EnableTrafficAffinity    *bool                                    `json:"EnableTrafficAffinity,omitempty" xml:"EnableTrafficAffinity,omitempty"`
	LoadBalancerName         *string                                  `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerType         *string                                  `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	RegionId                 *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId          *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroups           []*string                                `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	VpcId                    *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneMappings             []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetBillingConfig(v *CreateLoadBalancerRequestBillingConfig) *CreateLoadBalancerRequest {
	s.BillingConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetCommonBandwidthPackageId(v string) *CreateLoadBalancerRequest {
	s.CommonBandwidthPackageId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetEnableCrossZone(v bool) *CreateLoadBalancerRequest {
	s.EnableCrossZone = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetEnableTrafficAffinity(v bool) *CreateLoadBalancerRequest {
	s.EnableTrafficAffinity = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerType(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetSecurityGroups(v []*string) *CreateLoadBalancerRequest {
	s.SecurityGroups = v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

type CreateLoadBalancerRequestBillingConfig struct {
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// PayByTraffic, PayByBandwidth, PayByLcu, PayBy95, PayByOld95, PayBy96
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// PrePay, PostPay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period  *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	// Month, Year, Day
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s CreateLoadBalancerRequestBillingConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestBillingConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestBillingConfig) SetAutoPay(v bool) *CreateLoadBalancerRequestBillingConfig {
	s.AutoPay = &v
	return s
}

func (s *CreateLoadBalancerRequestBillingConfig) SetInternetChargeType(v string) *CreateLoadBalancerRequestBillingConfig {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLoadBalancerRequestBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestBillingConfig {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequestBillingConfig) SetPeriod(v int32) *CreateLoadBalancerRequestBillingConfig {
	s.Period = &v
	return s
}

func (s *CreateLoadBalancerRequestBillingConfig) SetPricingCycle(v string) *CreateLoadBalancerRequestBillingConfig {
	s.PricingCycle = &v
	return s
}

func (s *CreateLoadBalancerRequestBillingConfig) SetSpecification(v string) *CreateLoadBalancerRequestBillingConfig {
	s.Specification = &v
	return s
}

type CreateLoadBalancerRequestZoneMappings struct {
	// 公网ipId
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// 私网ip
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) SetAllocationId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetPrivateIPv4Address(v string) *CreateLoadBalancerRequestZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateLoadBalancerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetCode(v string) *CreateLoadBalancerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetData(v *CreateLoadBalancerResponseBodyData) *CreateLoadBalancerResponseBody {
	s.Data = v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetDynamicCode(v string) *CreateLoadBalancerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetDynamicMessage(v string) *CreateLoadBalancerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetHttpStatusCode(v int32) *CreateLoadBalancerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetMessage(v string) *CreateLoadBalancerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetSuccess(v bool) *CreateLoadBalancerResponseBody {
	s.Success = &v
	return s
}

type CreateLoadBalancerResponseBodyData struct {
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	LoadbalancerId *string `json:"LoadbalancerId,omitempty" xml:"LoadbalancerId,omitempty"`
}

func (s CreateLoadBalancerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBodyData) SetJobId(v string) *CreateLoadBalancerResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateLoadBalancerResponseBodyData) SetLoadbalancerId(v string) *CreateLoadBalancerResponseBodyData {
	s.LoadbalancerId = &v
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

type CreateSecurityPolicyRequest struct {
	Ciphers            []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	ClientToken        *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun             *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId           *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId    *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityPolicyName *string   `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	TlsVersions        []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
}

func (s CreateSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyRequest) SetCiphers(v []*string) *CreateSecurityPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateSecurityPolicyRequest) SetClientToken(v string) *CreateSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetDryRun(v bool) *CreateSecurityPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSecurityPolicyRequest) SetRegionId(v string) *CreateSecurityPolicyRequest {
	s.RegionId = &v
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

func (s *CreateSecurityPolicyRequest) SetTlsVersions(v []*string) *CreateSecurityPolicyRequest {
	s.TlsVersions = v
	return s
}

type CreateSecurityPolicyResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateSecurityPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBody) SetCode(v string) *CreateSecurityPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetData(v *CreateSecurityPolicyResponseBodyData) *CreateSecurityPolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetDynamicCode(v string) *CreateSecurityPolicyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetDynamicMessage(v string) *CreateSecurityPolicyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetHttpStatusCode(v int32) *CreateSecurityPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetMessage(v string) *CreateSecurityPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetRequestId(v string) *CreateSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetSuccess(v bool) *CreateSecurityPolicyResponseBody {
	s.Success = &v
	return s
}

type CreateSecurityPolicyResponseBodyData struct {
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s CreateSecurityPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBodyData) SetJobId(v string) *CreateSecurityPolicyResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBodyData) SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBodyData {
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

type CreateServerGroupRequest struct {
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 是否开启连接优雅中断
	ConnectionDrainEnable *bool `json:"ConnectionDrainEnable,omitempty" xml:"ConnectionDrainEnable,omitempty"`
	// 连接优雅中断超时时间
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	DryRun                 *bool  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 健康检查配置
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// 是否开启会话保持
	PersistenceEnable *bool `json:"PersistenceEnable,omitempty" xml:"PersistenceEnable,omitempty"`
	// 会话保持超时时间
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// 后端服务器类型
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 调度类型
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 服务器组名称
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// 服务器组类型
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// 服务器组所在vpc的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequest) SetAddressIPVersion(v string) *CreateServerGroupRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainEnable(v bool) *CreateServerGroupRequest {
	s.ConnectionDrainEnable = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainTimeout(v int32) *CreateServerGroupRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *CreateServerGroupRequest) SetDryRun(v bool) *CreateServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServerGroupRequest) SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest {
	s.HealthCheckConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetPersistenceEnable(v bool) *CreateServerGroupRequest {
	s.PersistenceEnable = &v
	return s
}

func (s *CreateServerGroupRequest) SetPersistenceTimeout(v int32) *CreateServerGroupRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *CreateServerGroupRequest) SetProtocol(v string) *CreateServerGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateServerGroupRequest) SetRegionId(v string) *CreateServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServerGroupRequest) SetResourceGroupId(v string) *CreateServerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerGroupRequest) SetScheduler(v string) *CreateServerGroupRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupName(v string) *CreateServerGroupRequest {
	s.ServerGroupName = &v
	return s
}

func (s *CreateServerGroupRequest) SetServerGroupType(v string) *CreateServerGroupRequest {
	s.ServerGroupType = &v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// 健康检查使用的端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 健康检查响应的最大超时时间
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// 健康检查的域名
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// 是否开启健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 状态码，多个状态码用逗号分隔
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// 健康检查时间间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查协议类型
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// 健康检查的url
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckDomain(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckType(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckType = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckUrl(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckUrl = &v
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

type CreateServerGroupResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateServerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBody) SetCode(v string) *CreateServerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetData(v *CreateServerGroupResponseBodyData) *CreateServerGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateServerGroupResponseBody) SetDynamicCode(v string) *CreateServerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetDynamicMessage(v string) *CreateServerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetHttpStatusCode(v int32) *CreateServerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetMessage(v string) *CreateServerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetRequestId(v string) *CreateServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetSuccess(v bool) *CreateServerGroupResponseBody {
	s.Success = &v
	return s
}

type CreateServerGroupResponseBodyData struct {
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateServerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBodyData) SetJobId(v string) *CreateServerGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateServerGroupResponseBodyData) SetServerGroupId(v string) *CreateServerGroupResponseBodyData {
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

type DeleteListenerRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// update or delete必选, add在custom中生成
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteListenerRequest) SetRegionId(v string) *DeleteListenerRequest {
	s.RegionId = &v
	return s
}

type DeleteListenerResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DeleteListenerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetCode(v string) *DeleteListenerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteListenerResponseBody) SetData(v *DeleteListenerResponseBodyData) *DeleteListenerResponseBody {
	s.Data = v
	return s
}

func (s *DeleteListenerResponseBody) SetDynamicCode(v string) *DeleteListenerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteListenerResponseBody) SetDynamicMessage(v string) *DeleteListenerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteListenerResponseBody) SetHttpStatusCode(v int32) *DeleteListenerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteListenerResponseBody) SetMessage(v string) *DeleteListenerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteListenerResponseBody) SetSuccess(v bool) *DeleteListenerResponseBody {
	s.Success = &v
	return s
}

type DeleteListenerResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteListenerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBodyData) SetJobId(v string) *DeleteListenerResponseBodyData {
	s.JobId = &v
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

type DeleteLoadBalancerRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetRegionId(v string) *DeleteLoadBalancerRequest {
	s.RegionId = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DeleteLoadBalancerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetCode(v string) *DeleteLoadBalancerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetData(v *DeleteLoadBalancerResponseBodyData) *DeleteLoadBalancerResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetDynamicCode(v string) *DeleteLoadBalancerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetDynamicMessage(v string) *DeleteLoadBalancerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetHttpStatusCode(v int32) *DeleteLoadBalancerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetMessage(v string) *DeleteLoadBalancerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLoadBalancerResponseBody) SetSuccess(v bool) *DeleteLoadBalancerResponseBody {
	s.Success = &v
	return s
}

type DeleteLoadBalancerResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteLoadBalancerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBodyData) SetJobId(v string) *DeleteLoadBalancerResponseBodyData {
	s.JobId = &v
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

type DeleteSecurityPolicyRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *DeleteSecurityPolicyRequest) SetRegionId(v string) *DeleteSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecurityPolicyRequest) SetSecurityPolicyId(v string) *DeleteSecurityPolicyRequest {
	s.SecurityPolicyId = &v
	return s
}

type DeleteSecurityPolicyResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DeleteSecurityPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBody) SetCode(v string) *DeleteSecurityPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetData(v *DeleteSecurityPolicyResponseBodyData) *DeleteSecurityPolicyResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetDynamicCode(v string) *DeleteSecurityPolicyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetDynamicMessage(v string) *DeleteSecurityPolicyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetHttpStatusCode(v int32) *DeleteSecurityPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetMessage(v string) *DeleteSecurityPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetRequestId(v string) *DeleteSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) SetSuccess(v bool) *DeleteSecurityPolicyResponseBody {
	s.Success = &v
	return s
}

type DeleteSecurityPolicyResponseBodyData struct {
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s DeleteSecurityPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBodyData) SetJobId(v string) *DeleteSecurityPolicyResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBodyData) SetSecurityPolicyId(v string) *DeleteSecurityPolicyResponseBodyData {
	s.SecurityPolicyId = &v
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

type DeleteServerGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
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

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetRegionId(v string) *DeleteServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type DeleteServerGroupResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DeleteServerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBody) SetCode(v string) *DeleteServerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetData(v *DeleteServerGroupResponseBodyData) *DeleteServerGroupResponseBody {
	s.Data = v
	return s
}

func (s *DeleteServerGroupResponseBody) SetDynamicCode(v string) *DeleteServerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetDynamicMessage(v string) *DeleteServerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetHttpStatusCode(v int32) *DeleteServerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetMessage(v string) *DeleteServerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetRequestId(v string) *DeleteServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServerGroupResponseBody) SetSuccess(v bool) *DeleteServerGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteServerGroupResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBodyData) SetJobId(v string) *DeleteServerGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DeleteServerGroupResponseBodyData) SetServerGroupId(v string) *DeleteServerGroupResponseBodyData {
	s.ServerGroupId = &v
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

type GetJobStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// add 必选
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetClientToken(v string) *GetJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

type GetJobStatusResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetCode(v string) *GetJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobStatusResponseBody) SetData(v *GetJobStatusResponseBodyData) *GetJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetJobStatusResponseBody) SetDynamicCode(v string) *GetJobStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetJobStatusResponseBody) SetDynamicMessage(v string) *GetJobStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetJobStatusResponseBody) SetHttpStatusCode(v int32) *GetJobStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJobStatusResponseBody) SetMessage(v string) *GetJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobStatusResponseBody) SetSuccess(v bool) *GetJobStatusResponseBody {
	s.Success = &v
	return s
}

type GetJobStatusResponseBodyData struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBodyData) SetStatus(v string) *GetJobStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetJobStatusResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type GetListenerAttributeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// update or delete必选, add在custom中生成
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) SetClientToken(v string) *GetListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetListenerAttributeRequest) SetDryRun(v bool) *GetListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeRequest) SetRegionId(v string) *GetListenerAttributeRequest {
	s.RegionId = &v
	return s
}

type GetListenerAttributeResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetListenerAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) SetCode(v string) *GetListenerAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetData(v *GetListenerAttributeResponseBodyData) *GetListenerAttributeResponseBody {
	s.Data = v
	return s
}

func (s *GetListenerAttributeResponseBody) SetDynamicCode(v string) *GetListenerAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetDynamicMessage(v string) *GetListenerAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetHttpStatusCode(v int32) *GetListenerAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetMessage(v string) *GetListenerAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetSuccess(v bool) *GetListenerAttributeResponseBody {
	s.Success = &v
	return s
}

type GetListenerAttributeResponseBodyData struct {
	// 用户uid
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// ca 证书列表
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	CaEnabled        *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// server证书列表
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// 创建时间
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// 空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 监听id
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议 (TCP, UDP, TCPSSL, GENEVE)
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	ListenerStatus   *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// 列表id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 业务location
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Tclssl监听的安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// servergroupId
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s GetListenerAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyData) SetAliUid(v int64) *GetListenerAttributeResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetCaCertificateIds(v []*string) *GetListenerAttributeResponseBodyData {
	s.CaCertificateIds = v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetCaEnabled(v bool) *GetListenerAttributeResponseBodyData {
	s.CaEnabled = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetCertificateIds(v []*string) *GetListenerAttributeResponseBodyData {
	s.CertificateIds = v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetGmtCreated(v string) *GetListenerAttributeResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetIdleTimeout(v int32) *GetListenerAttributeResponseBodyData {
	s.IdleTimeout = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetListenerDescription(v string) *GetListenerAttributeResponseBodyData {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetListenerId(v string) *GetListenerAttributeResponseBodyData {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetListenerPort(v int32) *GetListenerAttributeResponseBodyData {
	s.ListenerPort = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetListenerProtocol(v string) *GetListenerAttributeResponseBodyData {
	s.ListenerProtocol = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetListenerStatus(v string) *GetListenerAttributeResponseBodyData {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetLoadBalancerId(v string) *GetListenerAttributeResponseBodyData {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetRegionId(v string) *GetListenerAttributeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetSecurityPolicyId(v string) *GetListenerAttributeResponseBodyData {
	s.SecurityPolicyId = &v
	return s
}

func (s *GetListenerAttributeResponseBodyData) SetServerGroupId(v string) *GetListenerAttributeResponseBodyData {
	s.ServerGroupId = &v
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

type GetLoadBalancerAttributeRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) SetClientToken(v string) *GetLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetDryRun(v bool) *GetLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeRequest) SetRegionId(v string) *GetLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

type GetLoadBalancerAttributeResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetLoadBalancerAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) SetCode(v string) *GetLoadBalancerAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetData(v *GetLoadBalancerAttributeResponseBodyData) *GetLoadBalancerAttributeResponseBody {
	s.Data = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDynamicCode(v string) *GetLoadBalancerAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDynamicMessage(v string) *GetLoadBalancerAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetHttpStatusCode(v int32) *GetLoadBalancerAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetMessage(v string) *GetLoadBalancerAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetSuccess(v bool) *GetLoadBalancerAttributeResponseBody {
	s.Success = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyData struct {
	AddressIpVersion           *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType                *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	Bid                        *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CapacityUnitCount          *int64  `json:"CapacityUnitCount,omitempty" xml:"CapacityUnitCount,omitempty"`
	CommonBandwidthPackageId   *string `json:"CommonBandwidthPackageId,omitempty" xml:"CommonBandwidthPackageId,omitempty"`
	CreateTime                 *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossZoneEnable            *bool   `json:"CrossZoneEnable,omitempty" xml:"CrossZoneEnable,omitempty"`
	DNSName                    *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	LoadBalancerId             *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName           *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus         *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType           *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// 实例处于锁定状态列表
	OperationLocks        []*GetLoadBalancerAttributeResponseBodyDataOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Repeated"`
	RegionId              *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds      []*string                                                 `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	TrafficAffinityEnable *bool                                                     `json:"TrafficAffinityEnable,omitempty" xml:"TrafficAffinityEnable,omitempty"`
	VpcId                 *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneMappings          []*GetLoadBalancerAttributeResponseBodyDataZoneMappings   `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetAddressType(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetBid(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetCapacityUnitCount(v int64) *GetLoadBalancerAttributeResponseBodyData {
	s.CapacityUnitCount = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetCommonBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.CommonBandwidthPackageId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetCrossZoneEnable(v bool) *GetLoadBalancerAttributeResponseBodyData {
	s.CrossZoneEnable = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetDNSName(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.DNSName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetLoadBalancerType(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.LoadBalancerType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetOperationLocks(v []*GetLoadBalancerAttributeResponseBodyDataOperationLocks) *GetLoadBalancerAttributeResponseBodyData {
	s.OperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetRegionId(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetSecurityGroupIds(v []*string) *GetLoadBalancerAttributeResponseBodyData {
	s.SecurityGroupIds = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetTrafficAffinityEnable(v bool) *GetLoadBalancerAttributeResponseBodyData {
	s.TrafficAffinityEnable = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetVpcId(v string) *GetLoadBalancerAttributeResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyData) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyDataZoneMappings) *GetLoadBalancerAttributeResponseBodyData {
	s.ZoneMappings = v
	return s
}

type GetLoadBalancerAttributeResponseBodyDataOperationLocks struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	LockType   *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDataOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDataOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDataOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyDataOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyDataOperationLocks {
	s.LockType = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyDataZoneMappings struct {
	// 公网ipId
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EniId        *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// 私网ip
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// 公网ip地址：仅Get的时候有值
	PublicIPv4Address *string `json:"PublicIPv4Address,omitempty" xml:"PublicIPv4Address,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDataZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDataZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetAllocationId(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetEniId(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.EniId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetPrivateIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetPublicIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.PublicIPv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDataZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyDataZoneMappings {
	s.ZoneId = &v
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

type ListListenersRequest struct {
	// 监听唯一标识
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// 监听协议
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// 负载均衡实例标识
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetListenerProtocol(v string) *ListListenersRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetRegionId(v string) *ListListenersRequest {
	s.RegionId = &v
	return s
}

type ListListenersResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListListenersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                        `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                        `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetCode(v string) *ListListenersResponseBody {
	s.Code = &v
	return s
}

func (s *ListListenersResponseBody) SetData(v *ListListenersResponseBodyData) *ListListenersResponseBody {
	s.Data = v
	return s
}

func (s *ListListenersResponseBody) SetDynamicCode(v string) *ListListenersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListListenersResponseBody) SetDynamicMessage(v string) *ListListenersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListListenersResponseBody) SetHttpStatusCode(v int32) *ListListenersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListListenersResponseBody) SetMessage(v string) *ListListenersResponseBody {
	s.Message = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetSuccess(v bool) *ListListenersResponseBody {
	s.Success = &v
	return s
}

type ListListenersResponseBodyData struct {
	Listeners  []*ListListenersResponseBodyDataListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	MaxResults *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyData) SetListeners(v []*ListListenersResponseBodyDataListeners) *ListListenersResponseBodyData {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBodyData) SetMaxResults(v int32) *ListListenersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBodyData) SetNextToken(v string) *ListListenersResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBodyData) SetTotalCount(v int32) *ListListenersResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyDataListeners struct {
	// 用户uid
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// ca 证书列表
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	CaEnabled        *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// server证书列表
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// 创建时间
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// 修改时间
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 空闲超时时间
	IdleTimeout *int32 `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// 自己生成后赋值
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// 监听端口
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// 监听协议 (TCP, UDP, TCPSSL, GENEVE)
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	ListenerStatus   *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	LoadBalancerId   *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// 业务location
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Tclssl监听的安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// servergroupId
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s ListListenersResponseBodyDataListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyDataListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyDataListeners) SetAliUid(v int64) *ListListenersResponseBodyDataListeners {
	s.AliUid = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetCaCertificateIds(v []*string) *ListListenersResponseBodyDataListeners {
	s.CaCertificateIds = v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetCaEnabled(v bool) *ListListenersResponseBodyDataListeners {
	s.CaEnabled = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetCertificateIds(v []*string) *ListListenersResponseBodyDataListeners {
	s.CertificateIds = v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetGmtCreated(v string) *ListListenersResponseBodyDataListeners {
	s.GmtCreated = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetGmtModified(v string) *ListListenersResponseBodyDataListeners {
	s.GmtModified = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetIdleTimeout(v int32) *ListListenersResponseBodyDataListeners {
	s.IdleTimeout = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetListenerDescription(v string) *ListListenersResponseBodyDataListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetListenerId(v string) *ListListenersResponseBodyDataListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetListenerPort(v int32) *ListListenersResponseBodyDataListeners {
	s.ListenerPort = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetListenerProtocol(v string) *ListListenersResponseBodyDataListeners {
	s.ListenerProtocol = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetListenerStatus(v string) *ListListenersResponseBodyDataListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyDataListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetRegionId(v string) *ListListenersResponseBodyDataListeners {
	s.RegionId = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetSecurityPolicyId(v string) *ListListenersResponseBodyDataListeners {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListListenersResponseBodyDataListeners) SetServerGroupId(v string) *ListListenersResponseBodyDataListeners {
	s.ServerGroupId = &v
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

type ListLoadBalancersRequest struct {
	// 负载均衡地址 todo 增加校验方法
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// 协议类型
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// 地址类型：取值 internet，intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// dns 地址
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// 实例业务状态
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// 实例列表
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// 负载均衡实例名称
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// 实例状态
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// 负载均衡类型
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 企业资源组标识
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 专有网络唯一标识
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// 负载均衡拥有的可用区
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) SetAddress(v string) *ListLoadBalancersRequest {
	s.Address = &v
	return s
}

func (s *ListLoadBalancersRequest) SetAddressIpVersion(v string) *ListLoadBalancersRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersRequest) SetAddressType(v string) *ListLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetDNSName(v string) *ListLoadBalancersRequest {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBusinessStatus = &v
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

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerType(v string) *ListLoadBalancersRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetRegionId(v string) *ListLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneId(v string) *ListLoadBalancersRequest {
	s.ZoneId = &v
	return s
}

type ListLoadBalancersResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListLoadBalancersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) SetCode(v string) *ListLoadBalancersResponseBody {
	s.Code = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetData(v *ListLoadBalancersResponseBodyData) *ListLoadBalancersResponseBody {
	s.Data = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetDynamicCode(v string) *ListLoadBalancersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetDynamicMessage(v string) *ListLoadBalancersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetHttpStatusCode(v int32) *ListLoadBalancersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMessage(v string) *ListLoadBalancersResponseBody {
	s.Message = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetSuccess(v bool) *ListLoadBalancersResponseBody {
	s.Success = &v
	return s
}

type ListLoadBalancersResponseBodyData struct {
	LoadBalancers []*ListLoadBalancersResponseBodyDataLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	MaxResults    *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount    *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyData) SetLoadBalancers(v []*ListLoadBalancersResponseBodyDataLoadBalancers) *ListLoadBalancersResponseBodyData {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBodyData) SetMaxResults(v int32) *ListLoadBalancersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBodyData) SetNextToken(v string) *ListLoadBalancersResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBodyData) SetTotalCount(v int32) *ListLoadBalancersResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListLoadBalancersResponseBodyDataLoadBalancers struct {
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// 用户uid
	AliUid                   *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid                      *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CapacityUnitCount        *int64  `json:"CapacityUnitCount,omitempty" xml:"CapacityUnitCount,omitempty"`
	CommonBandwidthPackageId *string `json:"CommonBandwidthPackageId,omitempty" xml:"CommonBandwidthPackageId,omitempty"`
	CreateTime               *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossZoneEnable          *bool   `json:"CrossZoneEnable,omitempty" xml:"CrossZoneEnable,omitempty"`
	DNSName                  *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// 创建时间
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// 修改时间
	GmtModified                *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	LoadBalancerId             *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName           *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus         *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType           *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// 实例处于锁定状态列表
	OperationLocks []*ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Repeated"`
	// 业务location
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 物理location
	RegionNo              *string   `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceGroupId       *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds      []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	ServiceManagedEnabled *bool     `json:"ServiceManagedEnabled,omitempty" xml:"ServiceManagedEnabled,omitempty"`
	// 是否为托管实例，取值Managed-1, Unmanaged-0, DependencyManaged-2
	ServiceManagedMode *string `json:"ServiceManagedMode,omitempty" xml:"ServiceManagedMode,omitempty"`
	// 托管实例服务账号UID
	ServiceUid            *int64                                                        `json:"ServiceUid,omitempty" xml:"ServiceUid,omitempty"`
	TrafficAffinityEnable *bool                                                         `json:"TrafficAffinityEnable,omitempty" xml:"TrafficAffinityEnable,omitempty"`
	VpcId                 *string                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneMappings          []*ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s ListLoadBalancersResponseBodyDataLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyDataLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetAddressIpVersion(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetAddressType(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetAliUid(v int64) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.AliUid = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetBid(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.Bid = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetCapacityUnitCount(v int64) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.CapacityUnitCount = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetCommonBandwidthPackageId(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.CommonBandwidthPackageId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetCrossZoneEnable(v bool) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.CrossZoneEnable = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetDNSName(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetGmtCreated(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.GmtCreated = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetGmtModified(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.GmtModified = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetLoadBalancerType(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.LoadBalancerType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetOperationLocks(v []*ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.OperationLocks = v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetRegionId(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.RegionId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetRegionNo(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.RegionNo = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetSecurityGroupIds(v []*string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.SecurityGroupIds = v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetServiceManagedEnabled(v bool) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.ServiceManagedEnabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetServiceManagedMode(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.ServiceManagedMode = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetServiceUid(v int64) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.ServiceUid = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetTrafficAffinityEnable(v bool) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.TrafficAffinityEnable = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.VpcId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancers) SetZoneMappings(v []*ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) *ListLoadBalancersResponseBodyDataLoadBalancers {
	s.ZoneMappings = v
	return s
}

type ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	LockType   *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks) SetLockReason(v string) *ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks {
	s.LockReason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks) SetLockType(v string) *ListLoadBalancersResponseBodyDataLoadBalancersOperationLocks {
	s.LockType = &v
	return s
}

type ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings struct {
	// 公网ipId
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EniId        *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// 私网ip
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// 公网ip地址：仅Get的时候有值
	PublicIPv4Address *string `json:"PublicIPv4Address,omitempty" xml:"PublicIPv4Address,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId            *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetAllocationId(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetEniId(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.EniId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetPrivateIPv4Address(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetPublicIPv4Address(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.PublicIPv4Address = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetVSwitchId(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings) SetZoneId(v string) *ListLoadBalancersResponseBodyDataLoadBalancersZoneMappings {
	s.ZoneId = &v
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

type ListSecurityPolicyRequest struct {
	MaxResults          *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityPolicyIds   []*string `json:"SecurityPolicyIds,omitempty" xml:"SecurityPolicyIds,omitempty" type:"Repeated"`
	SecurityPolicyNames []*string `json:"SecurityPolicyNames,omitempty" xml:"SecurityPolicyNames,omitempty" type:"Repeated"`
}

func (s ListSecurityPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyRequest) SetMaxResults(v int32) *ListSecurityPolicyRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetNextToken(v string) *ListSecurityPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetRegionId(v string) *ListSecurityPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecurityPolicyRequest) SetSecurityPolicyIds(v []*string) *ListSecurityPolicyRequest {
	s.SecurityPolicyIds = v
	return s
}

func (s *ListSecurityPolicyRequest) SetSecurityPolicyNames(v []*string) *ListSecurityPolicyRequest {
	s.SecurityPolicyNames = v
	return s
}

type ListSecurityPolicyResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListSecurityPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSecurityPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBody) SetCode(v string) *ListSecurityPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetData(v *ListSecurityPolicyResponseBodyData) *ListSecurityPolicyResponseBody {
	s.Data = v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetDynamicCode(v string) *ListSecurityPolicyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetDynamicMessage(v string) *ListSecurityPolicyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetHttpStatusCode(v int32) *ListSecurityPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetMessage(v string) *ListSecurityPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetRequestId(v string) *ListSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityPolicyResponseBody) SetSuccess(v bool) *ListSecurityPolicyResponseBody {
	s.Success = &v
	return s
}

type ListSecurityPolicyResponseBodyData struct {
	MaxResults       *int32                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SecurityPolicies []*ListSecurityPolicyResponseBodyDataSecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	TotalCount       *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBodyData) SetMaxResults(v int32) *ListSecurityPolicyResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyData) SetNextToken(v string) *ListSecurityPolicyResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyData) SetSecurityPolicies(v []*ListSecurityPolicyResponseBodyDataSecurityPolicies) *ListSecurityPolicyResponseBodyData {
	s.SecurityPolicies = v
	return s
}

func (s *ListSecurityPolicyResponseBodyData) SetTotalCount(v int32) *ListSecurityPolicyResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListSecurityPolicyResponseBodyDataSecurityPolicies struct {
	// 加密套件
	Ciphers *string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty"`
	// 业务location
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// tls策略ID
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// 名称
	SecurityPolicyName *string `json:"SecurityPolicyName,omitempty" xml:"SecurityPolicyName,omitempty"`
	// tls版本
	TlsVersion *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s ListSecurityPolicyResponseBodyDataSecurityPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponseBodyDataSecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponseBodyDataSecurityPolicies) SetCiphers(v string) *ListSecurityPolicyResponseBodyDataSecurityPolicies {
	s.Ciphers = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyDataSecurityPolicies) SetRegionId(v string) *ListSecurityPolicyResponseBodyDataSecurityPolicies {
	s.RegionId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyDataSecurityPolicies) SetSecurityPolicyId(v string) *ListSecurityPolicyResponseBodyDataSecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyDataSecurityPolicies) SetSecurityPolicyName(v string) *ListSecurityPolicyResponseBodyDataSecurityPolicies {
	s.SecurityPolicyName = &v
	return s
}

func (s *ListSecurityPolicyResponseBodyDataSecurityPolicies) SetTlsVersion(v string) *ListSecurityPolicyResponseBodyDataSecurityPolicies {
	s.TlsVersion = &v
	return s
}

type ListSecurityPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPolicyResponse) SetHeaders(v map[string]*string) *ListSecurityPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPolicyResponse) SetBody(v *ListSecurityPolicyResponseBody) *ListSecurityPolicyResponse {
	s.Body = v
	return s
}

type ListServerGroupServersRequest struct {
	MaxResults    *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerGroupId *string   `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	ServerIds     []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
}

func (s ListServerGroupServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetRegionId(v string) *ListServerGroupServersRequest {
	s.RegionId = &v
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
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListServerGroupServersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) SetCode(v string) *ListServerGroupServersResponseBody {
	s.Code = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetData(v *ListServerGroupServersResponseBodyData) *ListServerGroupServersResponseBody {
	s.Data = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetDynamicCode(v string) *ListServerGroupServersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetDynamicMessage(v string) *ListServerGroupServersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetHttpStatusCode(v int32) *ListServerGroupServersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetMessage(v string) *ListServerGroupServersResponseBody {
	s.Message = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetSuccess(v bool) *ListServerGroupServersResponseBody {
	s.Success = &v
	return s
}

type ListServerGroupServersResponseBodyData struct {
	MaxResults *int32                                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Servers    []*ListServerGroupServersResponseBodyDataServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyData) SetMaxResults(v int32) *ListServerGroupServersResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBodyData) SetNextToken(v string) *ListServerGroupServersResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBodyData) SetServers(v []*ListServerGroupServersResponseBodyDataServers) *ListServerGroupServersResponseBodyData {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBodyData) SetTotalCount(v int32) *ListServerGroupServersResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListServerGroupServersResponseBodyDataServers struct {
	// 服务器描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 服务器端口
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 服务器的状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 后端权重
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 服务器对应的zoneId
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListServerGroupServersResponseBodyDataServers) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyDataServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyDataServers) SetDescription(v string) *ListServerGroupServersResponseBodyDataServers {
	s.Description = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetPort(v int32) *ListServerGroupServersResponseBodyDataServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetServerGroupId(v string) *ListServerGroupServersResponseBodyDataServers {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetServerId(v string) *ListServerGroupServersResponseBodyDataServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetServerIp(v string) *ListServerGroupServersResponseBodyDataServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetServerType(v string) *ListServerGroupServersResponseBodyDataServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetStatus(v string) *ListServerGroupServersResponseBodyDataServers {
	s.Status = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetWeight(v int32) *ListServerGroupServersResponseBodyDataServers {
	s.Weight = &v
	return s
}

func (s *ListServerGroupServersResponseBodyDataServers) SetZoneId(v string) *ListServerGroupServersResponseBodyDataServers {
	s.ZoneId = &v
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

type ListServerGroupsRequest struct {
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId  *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerGroupIds   []*string `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	ServerGroupNames []*string `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	// 服务器组所在vpc的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetRegionId(v string) *ListServerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListServerGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                           `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) SetCode(v string) *ListServerGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetData(v *ListServerGroupsResponseBodyData) *ListServerGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListServerGroupsResponseBody) SetDynamicCode(v string) *ListServerGroupsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetDynamicMessage(v string) *ListServerGroupsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetHttpStatusCode(v int32) *ListServerGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetMessage(v string) *ListServerGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetSuccess(v bool) *ListServerGroupsResponseBody {
	s.Success = &v
	return s
}

type ListServerGroupsResponseBodyData struct {
	MaxResults   *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ServerGroups []*ListServerGroupsResponseBodyDataServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	TotalCount   *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyData) SetMaxResults(v int32) *ListServerGroupsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBodyData) SetNextToken(v string) *ListServerGroupsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBodyData) SetServerGroups(v []*ListServerGroupsResponseBodyDataServerGroups) *ListServerGroupsResponseBodyData {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBodyData) SetTotalCount(v int32) *ListServerGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListServerGroupsResponseBodyDataServerGroups struct {
	// 服务器组地址类型
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 连接优雅中断开关
	ConnectionDrain *bool `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	// 连接优雅中断超时时间
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// 会话保持开关
	ConnectionPersistence *bool `json:"ConnectionPersistence,omitempty" xml:"ConnectionPersistence,omitempty"`
	// 会话保持超时时间
	ConnectionPersistenceTimeout *int32 `json:"ConnectionPersistenceTimeout,omitempty" xml:"ConnectionPersistenceTimeout,omitempty"`
	// 健康检查配置
	HealthCheck *ListServerGroupsResponseBodyDataServerGroupsHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// 后端协议
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// 业务region
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 调度类型
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 服务器组id
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// 服务器组名称
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// 状态
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// 服务器组类型
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// 服务器组的vpcid
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyDataServerGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyDataServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetAddressType(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.AddressType = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetAliUid(v int64) *ListServerGroupsResponseBodyDataServerGroups {
	s.AliUid = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetConnectionDrain(v bool) *ListServerGroupsResponseBodyDataServerGroups {
	s.ConnectionDrain = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetConnectionDrainTimeout(v int32) *ListServerGroupsResponseBodyDataServerGroups {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetConnectionPersistence(v bool) *ListServerGroupsResponseBodyDataServerGroups {
	s.ConnectionPersistence = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetConnectionPersistenceTimeout(v int32) *ListServerGroupsResponseBodyDataServerGroups {
	s.ConnectionPersistenceTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetHealthCheck(v *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) *ListServerGroupsResponseBodyDataServerGroups {
	s.HealthCheck = v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetRegionId(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.RegionId = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetResourceId(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ResourceId = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyDataServerGroups {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBodyDataServerGroupsHealthCheck struct {
	// 健康检查使用的端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 健康检查响应的最大超时时间
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// 健康检查的域名
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// 是否开启健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 状态码，多个状态码用逗号分隔
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// 健康检查时间间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查协议类型
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// 健康检查的url
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyDataServerGroupsHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyDataServerGroupsHealthCheck) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckConnectTimeout(v int32) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckDomain(v string) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckDomain = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckHttpCode(v []*string) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckHttpCode = v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckType(v string) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckType = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthCheckUrl(v string) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthCheckUrl = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyDataServerGroupsHealthCheck) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyDataServerGroupsHealthCheck {
	s.UnhealthyThreshold = &v
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

type RemoveServersFromServerGroupRequest struct {
	ClientToken   *string                                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool                                         `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerGroupId *string                                       `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetRegionId(v string) *RemoveServersFromServerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

type RemoveServersFromServerGroupRequestServers struct {
	// 服务器端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 服务器类型
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
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *RemoveServersFromServerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) SetCode(v string) *RemoveServersFromServerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetData(v *RemoveServersFromServerGroupResponseBodyData) *RemoveServersFromServerGroupResponseBody {
	s.Data = v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetDynamicCode(v string) *RemoveServersFromServerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetDynamicMessage(v string) *RemoveServersFromServerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetHttpStatusCode(v int32) *RemoveServersFromServerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetMessage(v string) *RemoveServersFromServerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetSuccess(v bool) *RemoveServersFromServerGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveServersFromServerGroupResponseBodyData struct {
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBodyData) SetJobId(v string) *RemoveServersFromServerGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBodyData) SetServerGroupId(v string) *RemoveServersFromServerGroupResponseBodyData {
	s.ServerGroupId = &v
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

type UpdateListenerAttributeRequest struct {
	// ca 证书列表
	CaCertificateIds []*string `json:"CaCertificateIds,omitempty" xml:"CaCertificateIds,omitempty" type:"Repeated"`
	CaEnabled        *bool     `json:"CaEnabled,omitempty" xml:"CaEnabled,omitempty"`
	// server证书列表
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	ClientToken    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	IdleTimeout    *int32    `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	// 监听描述
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// update or delete必选, add在custom中生成
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// https监听的安全策略
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// 实服务组
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) SetCaCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CaCertificateIds = v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCaEnabled(v bool) *UpdateListenerAttributeRequest {
	s.CaEnabled = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetCertificateIds(v []*string) *UpdateListenerAttributeRequest {
	s.CertificateIds = v
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

func (s *UpdateListenerAttributeRequest) SetIdleTimeout(v int32) *UpdateListenerAttributeRequest {
	s.IdleTimeout = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetRegionId(v string) *UpdateListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetSecurityPolicyId(v string) *UpdateListenerAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetServerGroupId(v string) *UpdateListenerAttributeRequest {
	s.ServerGroupId = &v
	return s
}

type UpdateListenerAttributeResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateListenerAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) SetCode(v string) *UpdateListenerAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetData(v *UpdateListenerAttributeResponseBodyData) *UpdateListenerAttributeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetDynamicCode(v string) *UpdateListenerAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetDynamicMessage(v string) *UpdateListenerAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetHttpStatusCode(v int32) *UpdateListenerAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetMessage(v string) *UpdateListenerAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateListenerAttributeResponseBody) SetSuccess(v bool) *UpdateListenerAttributeResponseBody {
	s.Success = &v
	return s
}

type UpdateListenerAttributeResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateListenerAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBodyData) SetJobId(v string) *UpdateListenerAttributeResponseBodyData {
	s.JobId = &v
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

type UpdateLoadBalancerAddressTypeConfigRequest struct {
	AddressType    *string                                                   `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	ClientToken    *string                                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool                                                     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string                                                   `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneMappings   []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetAddressType(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.AddressType = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetClientToken(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetDryRun(v bool) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetRegionId(v string) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequest) SetZoneMappings(v []*UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) *UpdateLoadBalancerAddressTypeConfigRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerAddressTypeConfigRequestZoneMappings struct {
	// 公网ipId
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	VSwitchId    *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerAddressTypeConfigRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateLoadBalancerAddressTypeConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                              `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetCode(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetData(v *UpdateLoadBalancerAddressTypeConfigResponseBodyData) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetDynamicCode(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetDynamicMessage(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetHttpStatusCode(v int32) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetMessage(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetRequestId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBody) SetSuccess(v bool) *UpdateLoadBalancerAddressTypeConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponseBodyData) SetJobId(v string) *UpdateLoadBalancerAddressTypeConfigResponseBodyData {
	s.JobId = &v
	return s
}

type UpdateLoadBalancerAddressTypeConfigResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLoadBalancerAddressTypeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAddressTypeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAddressTypeConfigResponse) SetBody(v *UpdateLoadBalancerAddressTypeConfigResponseBody) *UpdateLoadBalancerAddressTypeConfigResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAttributeRequest struct {
	ClientToken           *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EnableCrossZone       *bool     `json:"EnableCrossZone,omitempty" xml:"EnableCrossZone,omitempty"`
	EnableTrafficAffinity *bool     `json:"EnableTrafficAffinity,omitempty" xml:"EnableTrafficAffinity,omitempty"`
	LoadBalancerId        *string   `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName      *string   `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	RegionId              *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroups        []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
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

func (s *UpdateLoadBalancerAttributeRequest) SetEnableCrossZone(v bool) *UpdateLoadBalancerAttributeRequest {
	s.EnableCrossZone = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetEnableTrafficAffinity(v bool) *UpdateLoadBalancerAttributeRequest {
	s.EnableTrafficAffinity = &v
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

func (s *UpdateLoadBalancerAttributeRequest) SetRegionId(v string) *UpdateLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetSecurityGroups(v []*string) *UpdateLoadBalancerAttributeRequest {
	s.SecurityGroups = v
	return s
}

type UpdateLoadBalancerAttributeResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateLoadBalancerAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                      `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                      `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetCode(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetData(v *UpdateLoadBalancerAttributeResponseBodyData) *UpdateLoadBalancerAttributeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetDynamicCode(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetDynamicMessage(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetHttpStatusCode(v int32) *UpdateLoadBalancerAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetMessage(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetSuccess(v bool) *UpdateLoadBalancerAttributeResponseBody {
	s.Success = &v
	return s
}

type UpdateLoadBalancerAttributeResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBodyData) SetJobId(v string) *UpdateLoadBalancerAttributeResponseBodyData {
	s.JobId = &v
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

type UpdateLoadBalancerZonesRequest struct {
	ClientToken    *string                                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun         *bool                                         `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	LoadBalancerId *string                                       `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	RegionId       *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneMappings   []*UpdateLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequest) SetClientToken(v string) *UpdateLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetDryRun(v bool) *UpdateLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetRegionId(v string) *UpdateLoadBalancerZonesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerZonesRequestZoneMappings struct {
	// 公网ipId
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// 私网ip
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetAllocationId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetPrivateIPv4Address(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerZonesResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateLoadBalancerZonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBody) SetCode(v string) *UpdateLoadBalancerZonesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetData(v *UpdateLoadBalancerZonesResponseBodyData) *UpdateLoadBalancerZonesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetDynamicCode(v string) *UpdateLoadBalancerZonesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetDynamicMessage(v string) *UpdateLoadBalancerZonesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetHttpStatusCode(v int32) *UpdateLoadBalancerZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetMessage(v string) *UpdateLoadBalancerZonesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponseBody) SetSuccess(v bool) *UpdateLoadBalancerZonesResponseBody {
	s.Success = &v
	return s
}

type UpdateLoadBalancerZonesResponseBodyData struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBodyData) SetJobId(v string) *UpdateLoadBalancerZonesResponseBodyData {
	s.JobId = &v
	return s
}

type UpdateLoadBalancerZonesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLoadBalancerZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLoadBalancerZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerZonesResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetBody(v *UpdateLoadBalancerZonesResponseBody) *UpdateLoadBalancerZonesResponse {
	s.Body = v
	return s
}

type UpdateSecurityPolicyAttributeRequest struct {
	Ciphers          []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestContent   *string   `json:"RequestContent,omitempty" xml:"RequestContent,omitempty"`
	SecurityPolicyId *string   `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	TlsVersions      []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
}

func (s UpdateSecurityPolicyAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeRequest) SetCiphers(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetClientToken(v string) *UpdateSecurityPolicyAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetDryRun(v bool) *UpdateSecurityPolicyAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetRegionId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetRequestContent(v string) *UpdateSecurityPolicyAttributeRequest {
	s.RequestContent = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeRequest {
	s.SecurityPolicyId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeRequest) SetTlsVersions(v []*string) *UpdateSecurityPolicyAttributeRequest {
	s.TlsVersions = v
	return s
}

type UpdateSecurityPolicyAttributeResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateSecurityPolicyAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                        `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                        `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetCode(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetData(v *UpdateSecurityPolicyAttributeResponseBodyData) *UpdateSecurityPolicyAttributeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetDynamicCode(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetDynamicMessage(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetHttpStatusCode(v int32) *UpdateSecurityPolicyAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetMessage(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetRequestId(v string) *UpdateSecurityPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBody) SetSuccess(v bool) *UpdateSecurityPolicyAttributeResponseBody {
	s.Success = &v
	return s
}

type UpdateSecurityPolicyAttributeResponseBodyData struct {
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s UpdateSecurityPolicyAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecurityPolicyAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSecurityPolicyAttributeResponseBodyData) SetJobId(v string) *UpdateSecurityPolicyAttributeResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateSecurityPolicyAttributeResponseBodyData) SetSecurityPolicyId(v string) *UpdateSecurityPolicyAttributeResponseBodyData {
	s.SecurityPolicyId = &v
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

type UpdateServerGroupAttributeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 是否开启连接优雅中断
	ConnectionDrainEnable *bool `json:"ConnectionDrainEnable,omitempty" xml:"ConnectionDrainEnable,omitempty"`
	// 连接优雅中断超时时间
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	DryRun                 *bool  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// 健康检查配置
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// 是否开启会话保持
	PersistenceEnable *bool `json:"PersistenceEnable,omitempty" xml:"PersistenceEnable,omitempty"`
	// 会话保持超时时间
	PersistenceTimeout *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 调度类型
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// 服务器组ID
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainEnable(v bool) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainEnable = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainTimeout(v int32) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainTimeout = &v
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

func (s *UpdateServerGroupAttributeRequest) SetPersistenceEnable(v bool) *UpdateServerGroupAttributeRequest {
	s.PersistenceEnable = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetPersistenceTimeout(v int32) *UpdateServerGroupAttributeRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetRegionId(v string) *UpdateServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetScheduler(v string) *UpdateServerGroupAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupId = &v
	return s
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// 健康检查使用的端口
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// 健康检查响应的最大超时时间
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// 健康检查的域名
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// 是否开启健康检查
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// 状态码，多个状态码用逗号分隔
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// 健康检查时间间隔
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// 健康检查协议类型
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// 健康检查的url
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// 健康检查连续成功多少次后，将后端服务器的健康检查状态由fail判定为success
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// 健康检查连续失败多少次后，将后端服务器的健康检查状态由success判定为fail
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckDomain(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckType(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckType = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckUrl(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckUrl = &v
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

type UpdateServerGroupAttributeResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateServerGroupAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) SetCode(v string) *UpdateServerGroupAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetData(v *UpdateServerGroupAttributeResponseBodyData) *UpdateServerGroupAttributeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetDynamicCode(v string) *UpdateServerGroupAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetDynamicMessage(v string) *UpdateServerGroupAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetHttpStatusCode(v int32) *UpdateServerGroupAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetMessage(v string) *UpdateServerGroupAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBody) SetSuccess(v bool) *UpdateServerGroupAttributeResponseBody {
	s.Success = &v
	return s
}

type UpdateServerGroupAttributeResponseBodyData struct {
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBodyData) SetJobId(v string) *UpdateServerGroupAttributeResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupAttributeResponseBodyData) SetServerGroupId(v string) *UpdateServerGroupAttributeResponseBodyData {
	s.ServerGroupId = &v
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

type UpdateServerGroupServersAttributeRequest struct {
	ClientToken   *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool                                              `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	RegionId      *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerGroupId *string                                            `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	Servers       []*UpdateServerGroupServersAttributeRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s UpdateServerGroupServersAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeRequest) SetClientToken(v string) *UpdateServerGroupServersAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetDryRun(v bool) *UpdateServerGroupServersAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetRegionId(v string) *UpdateServerGroupServersAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServerGroupId(v string) *UpdateServerGroupServersAttributeRequest {
	s.ServerGroupId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeRequest) SetServers(v []*UpdateServerGroupServersAttributeRequestServers) *UpdateServerGroupServersAttributeRequest {
	s.Servers = v
	return s
}

type UpdateServerGroupServersAttributeRequestServers struct {
	// 服务器描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 服务器端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 服务器id
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// 服务器ip
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// 服务器类型
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// 后端权重
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
	Code           *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateServerGroupServersAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode    *string                                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateServerGroupServersAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetCode(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetData(v *UpdateServerGroupServersAttributeResponseBodyData) *UpdateServerGroupServersAttributeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetDynamicCode(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetDynamicMessage(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetHttpStatusCode(v int32) *UpdateServerGroupServersAttributeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetMessage(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupServersAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBody) SetSuccess(v bool) *UpdateServerGroupServersAttributeResponseBody {
	s.Success = &v
	return s
}

type UpdateServerGroupServersAttributeResponseBodyData struct {
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateServerGroupServersAttributeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupServersAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupServersAttributeResponseBodyData) SetJobId(v string) *UpdateServerGroupServersAttributeResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateServerGroupServersAttributeResponseBodyData) SetServerGroupId(v string) *UpdateServerGroupServersAttributeResponseBodyData {
	s.ServerGroupId = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("slbv2"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddServersToServerGroupWithOptions(request *AddServersToServerGroupRequest, runtime *util.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		body["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServersToServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertificateIds)) {
		body["CaCertificateIds"] = request.CaCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		body["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateIds)) {
		body["CertificateIds"] = request.CertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerPort)) {
		body["ListenerPort"] = request.ListenerPort
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		body["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		body["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		body["AddressType"] = request.AddressType
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.BillingConfig))) {
		bodyFlat["BillingConfig"] = request.BillingConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommonBandwidthPackageId)) {
		body["CommonBandwidthPackageId"] = request.CommonBandwidthPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrossZone)) {
		body["EnableCrossZone"] = request.EnableCrossZone
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTrafficAffinity)) {
		body["EnableTrafficAffinity"] = request.EnableTrafficAffinity
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerType)) {
		body["LoadBalancerType"] = request.LoadBalancerType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroups)) {
		body["SecurityGroups"] = request.SecurityGroups
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		body["ZoneMappings"] = request.ZoneMappings
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateSecurityPolicyWithOptions(request *CreateSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		body["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyName)) {
		body["SecurityPolicyName"] = request.SecurityPolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.TlsVersions)) {
		body["TlsVersions"] = request.TlsVersions
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateServerGroupWithOptions(request *CreateServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIPVersion)) {
		body["AddressIPVersion"] = request.AddressIPVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainEnable)) {
		body["ConnectionDrainEnable"] = request.ConnectionDrainEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		body["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.HealthCheckConfig))) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceEnable)) {
		body["PersistenceEnable"] = request.PersistenceEnable
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceTimeout)) {
		body["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		body["ServerGroupName"] = request.ServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		body["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteListener"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancer"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteSecurityPolicyWithOptions(request *DeleteSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteServerGroupWithOptions(request *DeleteServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobStatus"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		query["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetLoadBalancerAttributeWithOptions(request *GetLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoadBalancerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		query["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerProtocol)) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.DNSName)) {
		query["DNSName"] = request.DNSName
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBusinessStatus)) {
		query["LoadBalancerBusinessStatus"] = request.LoadBalancerBusinessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerNames)) {
		query["LoadBalancerNames"] = request.LoadBalancerNames
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerType)) {
		query["LoadBalancerType"] = request.LoadBalancerType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcIds)) {
		query["VpcIds"] = request.VpcIds
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLoadBalancers"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListSecurityPolicyWithOptions(request *ListSecurityPolicyRequest, runtime *util.RuntimeOptions) (_result *ListSecurityPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyIds)) {
		body["SecurityPolicyIds"] = request.SecurityPolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyNames)) {
		body["SecurityPolicyNames"] = request.SecurityPolicyNames
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecurityPolicy"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecurityPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityPolicy(request *ListSecurityPolicyRequest) (_result *ListSecurityPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityPolicyResponse{}
	_body, _err := client.ListSecurityPolicyWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIds)) {
		body["ServerIds"] = request.ServerIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroupServers"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListServerGroupsWithOptions(request *ListServerGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupIds)) {
		body["ServerGroupIds"] = request.ServerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupNames)) {
		body["ServerGroupNames"] = request.ServerGroupNames
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroups"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RemoveServersFromServerGroupWithOptions(request *RemoveServersFromServerGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		body["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveServersFromServerGroup"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateListenerAttributeWithOptions(request *UpdateListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaCertificateIds)) {
		body["CaCertificateIds"] = request.CaCertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.CaEnabled)) {
		body["CaEnabled"] = request.CaEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateIds)) {
		body["CertificateIds"] = request.CertificateIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.IdleTimeout)) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListenerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateLoadBalancerAddressTypeConfigWithOptions(request *UpdateLoadBalancerAddressTypeConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		body["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		body["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAddressTypeConfig"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoadBalancerAddressTypeConfig(request *UpdateLoadBalancerAddressTypeConfigRequest) (_result *UpdateLoadBalancerAddressTypeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAddressTypeConfigResponse{}
	_body, _err := client.UpdateLoadBalancerAddressTypeConfigWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCrossZone)) {
		body["EnableCrossZone"] = request.EnableCrossZone
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTrafficAffinity)) {
		body["EnableTrafficAffinity"] = request.EnableTrafficAffinity
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroups)) {
		body["SecurityGroups"] = request.SecurityGroups
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateLoadBalancerZonesWithOptions(request *UpdateLoadBalancerZonesRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		body["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerZones"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoadBalancerZones(request *UpdateLoadBalancerZonesRequest) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.UpdateLoadBalancerZonesWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ciphers)) {
		body["Ciphers"] = request.Ciphers
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestContent)) {
		body["RequestContent"] = request.RequestContent
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyId)) {
		body["SecurityPolicyId"] = request.SecurityPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TlsVersions)) {
		body["TlsVersions"] = request.TlsVersions
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecurityPolicyAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSecurityPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateServerGroupAttributeWithOptions(request *UpdateServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainEnable)) {
		body["ConnectionDrainEnable"] = request.ConnectionDrainEnable
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainTimeout)) {
		body["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.HealthCheckConfig))) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceEnable)) {
		body["PersistenceEnable"] = request.PersistenceEnable
	}

	if !tea.BoolValue(util.IsUnset(request.PersistenceTimeout)) {
		body["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateServerGroupServersAttributeWithOptions(request *UpdateServerGroupServersAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupServersAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		body["Servers"] = request.Servers
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupServersAttribute"),
		Version:     tea.String("2022-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupServersAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
