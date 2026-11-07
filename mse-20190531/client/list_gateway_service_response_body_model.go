// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayServiceResponseBody
	GetCode() *int32
	SetData(v *ListGatewayServiceResponseBodyData) *ListGatewayServiceResponseBody
	GetData() *ListGatewayServiceResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayServiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayServiceResponseBody
	GetSuccess() *bool
}

type ListGatewayServiceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListGatewayServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F46CDBA4-B1EE-5C94-8A48-51C10177****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayServiceResponseBody) GetData() *ListGatewayServiceResponseBodyData {
	return s.Data
}

func (s *ListGatewayServiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayServiceResponseBody) SetCode(v int32) *ListGatewayServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayServiceResponseBody) SetData(v *ListGatewayServiceResponseBodyData) *ListGatewayServiceResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayServiceResponseBody) SetHttpStatusCode(v int32) *ListGatewayServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayServiceResponseBody) SetMessage(v string) *ListGatewayServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayServiceResponseBody) SetRequestId(v string) *ListGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayServiceResponseBody) SetSuccess(v bool) *ListGatewayServiceResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of records per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned collection.
	Result []*ListGatewayServiceResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 9
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayServiceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayServiceResponseBodyData) GetResult() []*ListGatewayServiceResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayServiceResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayServiceResponseBodyData) SetPageNumber(v int32) *ListGatewayServiceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayServiceResponseBodyData) SetPageSize(v int32) *ListGatewayServiceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayServiceResponseBodyData) SetResult(v []*ListGatewayServiceResponseBodyDataResult) *ListGatewayServiceResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayServiceResponseBodyData) SetTotalSize(v int64) *ListGatewayServiceResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayServiceResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyDataResult struct {
	// The DNS server list (for DNS domain name service only).
	DnsServerList []*string `json:"DnsServerList,omitempty" xml:"DnsServerList,omitempty" type:"Repeated"`
	// The gateway ID.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The traffic management policy.
	GatewayTrafficPolicy *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy `json:"GatewayTrafficPolicy,omitempty" xml:"GatewayTrafficPolicy,omitempty" type:"Struct"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The time when the record was created.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the record was modified.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The health status. Valid values:
	//
	// - Health
	//
	// - Unhealthy
	//
	// - Unknown
	//
	// example:
	//
	// Unhealthy
	HealehStatus *string `json:"HealehStatus,omitempty" xml:"HealehStatus,omitempty"`
	// Indicates whether health checks are performed.
	//
	// example:
	//
	// true
	HealthCheck *bool `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The health check information.
	HealthCheckInfo *ListGatewayServiceResponseBodyDataResultHealthCheckInfo `json:"HealthCheckInfo,omitempty" xml:"HealthCheckInfo,omitempty" type:"Struct"`
	// The health status. Valid values:
	//
	// - Health
	//
	// - Unhealthy
	//
	// - Unknown
	//
	// example:
	//
	// Unhealthy
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address list.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The metadata or IP address list.
	//
	// example:
	//
	// {}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	// The service name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The array of ports.
	Ports []*int32 `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The fully qualified domain name (FQDN) of the service.
	//
	// example:
	//
	// demo.default.svc.cluster.local
	ServiceFQDN *string `json:"ServiceFQDN,omitempty" xml:"ServiceFQDN,omitempty"`
	// The original service name in the registry.
	//
	// example:
	//
	// test
	ServiceNameInRegistry *string `json:"ServiceNameInRegistry,omitempty" xml:"ServiceNameInRegistry,omitempty"`
	// The service port.
	//
	// example:
	//
	// 80
	ServicePort *int64 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The source ID.
	//
	// example:
	//
	// 2
	SourceId *int64 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The array of unhealthy endpoints.
	UnhealthyEndpoints []*string `json:"UnhealthyEndpoints,omitempty" xml:"UnhealthyEndpoints,omitempty" type:"Repeated"`
	// The service versions.
	Versions []*ListGatewayServiceResponseBodyDataResultVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListGatewayServiceResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResult) GetDnsServerList() []*string {
	return s.DnsServerList
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGatewayTrafficPolicy() *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy {
	return s.GatewayTrafficPolicy
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayServiceResponseBodyDataResult) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayServiceResponseBodyDataResult) GetHealehStatus() *string {
	return s.HealehStatus
}

func (s *ListGatewayServiceResponseBodyDataResult) GetHealthCheck() *bool {
	return s.HealthCheck
}

func (s *ListGatewayServiceResponseBodyDataResult) GetHealthCheckInfo() *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	return s.HealthCheckInfo
}

func (s *ListGatewayServiceResponseBodyDataResult) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListGatewayServiceResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayServiceResponseBodyDataResult) GetIps() []*string {
	return s.Ips
}

func (s *ListGatewayServiceResponseBodyDataResult) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *ListGatewayServiceResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGatewayServiceResponseBodyDataResult) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGatewayServiceResponseBodyDataResult) GetPorts() []*int32 {
	return s.Ports
}

func (s *ListGatewayServiceResponseBodyDataResult) GetServiceFQDN() *string {
	return s.ServiceFQDN
}

func (s *ListGatewayServiceResponseBodyDataResult) GetServiceNameInRegistry() *string {
	return s.ServiceNameInRegistry
}

func (s *ListGatewayServiceResponseBodyDataResult) GetServicePort() *int64 {
	return s.ServicePort
}

func (s *ListGatewayServiceResponseBodyDataResult) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *ListGatewayServiceResponseBodyDataResult) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListGatewayServiceResponseBodyDataResult) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGatewayServiceResponseBodyDataResult) GetUnhealthyEndpoints() []*string {
	return s.UnhealthyEndpoints
}

func (s *ListGatewayServiceResponseBodyDataResult) GetVersions() []*ListGatewayServiceResponseBodyDataResultVersions {
	return s.Versions
}

func (s *ListGatewayServiceResponseBodyDataResult) SetDnsServerList(v []*string) *ListGatewayServiceResponseBodyDataResult {
	s.DnsServerList = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayServiceResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGatewayTrafficPolicy(v *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) *ListGatewayServiceResponseBodyDataResult {
	s.GatewayTrafficPolicy = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayServiceResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayServiceResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGmtModified(v string) *ListGatewayServiceResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetGroupName(v string) *ListGatewayServiceResponseBodyDataResult {
	s.GroupName = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetHealehStatus(v string) *ListGatewayServiceResponseBodyDataResult {
	s.HealehStatus = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetHealthCheck(v bool) *ListGatewayServiceResponseBodyDataResult {
	s.HealthCheck = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetHealthCheckInfo(v *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) *ListGatewayServiceResponseBodyDataResult {
	s.HealthCheckInfo = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetHealthStatus(v string) *ListGatewayServiceResponseBodyDataResult {
	s.HealthStatus = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetId(v int64) *ListGatewayServiceResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetIps(v []*string) *ListGatewayServiceResponseBodyDataResult {
	s.Ips = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetMetaInfo(v string) *ListGatewayServiceResponseBodyDataResult {
	s.MetaInfo = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetName(v string) *ListGatewayServiceResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetNamespace(v string) *ListGatewayServiceResponseBodyDataResult {
	s.Namespace = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetPorts(v []*int32) *ListGatewayServiceResponseBodyDataResult {
	s.Ports = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetServiceFQDN(v string) *ListGatewayServiceResponseBodyDataResult {
	s.ServiceFQDN = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetServiceNameInRegistry(v string) *ListGatewayServiceResponseBodyDataResult {
	s.ServiceNameInRegistry = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetServicePort(v int64) *ListGatewayServiceResponseBodyDataResult {
	s.ServicePort = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetServiceProtocol(v string) *ListGatewayServiceResponseBodyDataResult {
	s.ServiceProtocol = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetSourceId(v int64) *ListGatewayServiceResponseBodyDataResult {
	s.SourceId = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetSourceType(v string) *ListGatewayServiceResponseBodyDataResult {
	s.SourceType = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetUnhealthyEndpoints(v []*string) *ListGatewayServiceResponseBodyDataResult {
	s.UnhealthyEndpoints = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) SetVersions(v []*ListGatewayServiceResponseBodyDataResultVersions) *ListGatewayServiceResponseBodyDataResult {
	s.Versions = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResult) Validate() error {
	if s.GatewayTrafficPolicy != nil {
		if err := s.GatewayTrafficPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckInfo != nil {
		if err := s.HealthCheckInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy struct {
	// The load balancing settings.
	LoadBalancerSettings *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings `json:"LoadBalancerSettings,omitempty" xml:"LoadBalancerSettings,omitempty" type:"Struct"`
	// The TLS.
	Tls *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls `json:"Tls,omitempty" xml:"Tls,omitempty" type:"Struct"`
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) GetLoadBalancerSettings() *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings {
	return s.LoadBalancerSettings
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) GetTls() *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	return s.Tls
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) SetLoadBalancerSettings(v *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy {
	s.LoadBalancerSettings = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) SetTls(v *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy {
	s.Tls = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicy) Validate() error {
	if s.LoadBalancerSettings != nil {
		if err := s.LoadBalancerSettings.Validate(); err != nil {
			return err
		}
	}
	if s.Tls != nil {
		if err := s.Tls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings struct {
	// The consistent hashing configuration.
	ConsistentHashLBConfig *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig `json:"ConsistentHashLBConfig,omitempty" xml:"ConsistentHashLBConfig,omitempty" type:"Struct"`
	// The load balancing policy. Valid values:
	//
	// - ROUND_ROBIN
	//
	// - LEAST_CONN
	//
	// - RANDOM
	//
	// - CONSISTENT_HASH
	//
	// example:
	//
	// RANDOM
	LoadbalancerType *string `json:"LoadbalancerType,omitempty" xml:"LoadbalancerType,omitempty"`
	// The warm-up duration for the minimum number of requests.
	//
	// example:
	//
	// 10
	WarmupDuration *int32 `json:"WarmupDuration,omitempty" xml:"WarmupDuration,omitempty"`
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) GetConsistentHashLBConfig() *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	return s.ConsistentHashLBConfig
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) GetLoadbalancerType() *string {
	return s.LoadbalancerType
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) GetWarmupDuration() *int32 {
	return s.WarmupDuration
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) SetConsistentHashLBConfig(v *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings {
	s.ConsistentHashLBConfig = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) SetLoadbalancerType(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings {
	s.LoadbalancerType = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) SetWarmupDuration(v int32) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings {
	s.WarmupDuration = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettings) Validate() error {
	if s.ConsistentHashLBConfig != nil {
		if err := s.ConsistentHashLBConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig struct {
	// The consistent hashing method. Valid values:
	//
	// - HEADER
	//
	// - COOKIE
	//
	// - SOURCE_IP
	//
	// - QUERY_PARAMETER
	//
	// example:
	//
	// HEADER
	ConsistentHashLBType *string `json:"ConsistentHashLBType,omitempty" xml:"ConsistentHashLBType,omitempty"`
	// The cookie-based parameter.
	HttpCookie *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie `json:"HttpCookie,omitempty" xml:"HttpCookie,omitempty" type:"Struct"`
	// The minimum size of the hash ring.
	//
	// example:
	//
	// 10000
	MinimumRingSize *int64 `json:"MinimumRingSize,omitempty" xml:"MinimumRingSize,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// param
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetConsistentHashLBType() *string {
	return s.ConsistentHashLBType
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetHttpCookie() *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	return s.HttpCookie
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetMinimumRingSize() *int64 {
	return s.MinimumRingSize
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetConsistentHashLBType(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ConsistentHashLBType = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetHttpCookie(v *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.HttpCookie = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetMinimumRingSize(v int64) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.MinimumRingSize = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) SetParameterName(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig {
	s.ParameterName = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig) Validate() error {
	if s.HttpCookie != nil {
		if err := s.HttpCookie.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie struct {
	// The cookie name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cookie path.
	//
	// example:
	//
	// /path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The time-to-live (TTL) of the cookie.
	//
	// example:
	//
	// 360
	Ttl *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetName() *string {
	return s.Name
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetPath() *string {
	return s.Path
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) GetTtl() *string {
	return s.Ttl
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetName(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Name = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetPath(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Path = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) SetTtl(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie {
	s.Ttl = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie) Validate() error {
	return dara.Validate(s)
}

type ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls struct {
	// The public key of the CA certificate.
	//
	// example:
	//
	// content
	CaCertContent *string `json:"CaCertContent,omitempty" xml:"CaCertContent,omitempty"`
	// The CA certificate ID.
	//
	// example:
	//
	// 5******-cn-hangzhou
	CaCertId *string `json:"CaCertId,omitempty" xml:"CaCertId,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 5******-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The TLS mode. Valid values:
	//
	// - DISABLE
	//
	// - SIMPLE
	//
	// - MUTUAL
	//
	// - ISTIO_MUTUAL
	//
	// example:
	//
	// SIMPLE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The service name identity.
	//
	// example:
	//
	// name-sni
	Sni *string `json:"Sni,omitempty" xml:"Sni,omitempty"`
	// The array of subject alternative names.
	SubjectAltNames []*string `json:"SubjectAltNames,omitempty" xml:"SubjectAltNames,omitempty" type:"Repeated"`
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetCaCertContent() *string {
	return s.CaCertContent
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetCaCertId() *string {
	return s.CaCertId
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetCertId() *string {
	return s.CertId
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetMode() *string {
	return s.Mode
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetSni() *string {
	return s.Sni
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) GetSubjectAltNames() []*string {
	return s.SubjectAltNames
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetCaCertContent(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.CaCertContent = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetCaCertId(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.CaCertId = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetCertId(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.CertId = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetMode(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.Mode = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetSni(v string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.Sni = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) SetSubjectAltNames(v []*string) *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls {
	s.SubjectAltNames = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultGatewayTrafficPolicyTls) Validate() error {
	return dara.Validate(s)
}

type ListGatewayServiceResponseBodyDataResultHealthCheckInfo struct {
	// Indicates whether the check is performed.
	//
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// The expected HTTP status codes for health checks.
	ExpectedStatuses []*int32 `json:"ExpectedStatuses,omitempty" xml:"ExpectedStatuses,omitempty" type:"Repeated"`
	// The healthy threshold.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP request address for health checks.
	//
	// example:
	//
	// http://172.168.0.1
	HttpHost *string `json:"HttpHost,omitempty" xml:"HttpHost,omitempty"`
	// The HTTP request path for health checks.
	//
	// example:
	//
	// /health
	HttpPath *string `json:"HttpPath,omitempty" xml:"HttpPath,omitempty"`
	// The health check interval.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The network protocol.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 2
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The unhealthy threshold.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListGatewayServiceResponseBodyDataResultHealthCheckInfo) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetCheck() *bool {
	return s.Check
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetExpectedStatuses() []*int32 {
	return s.ExpectedStatuses
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetHttpHost() *string {
	return s.HttpHost
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetHttpPath() *string {
	return s.HttpPath
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetInterval() *int32 {
	return s.Interval
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetCheck(v bool) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.Check = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetExpectedStatuses(v []*int32) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.ExpectedStatuses = v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetHealthyThreshold(v int32) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.HealthyThreshold = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetHttpHost(v string) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.HttpHost = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetHttpPath(v string) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.HttpPath = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetInterval(v int32) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.Interval = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetProtocol(v string) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.Protocol = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetTimeout(v int32) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.Timeout = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) SetUnhealthyThreshold(v int32) *ListGatewayServiceResponseBodyDataResultHealthCheckInfo {
	s.UnhealthyThreshold = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultHealthCheckInfo) Validate() error {
	return dara.Validate(s)
}

type ListGatewayServiceResponseBodyDataResultVersions struct {
	// The version name.
	//
	// example:
	//
	// v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListGatewayServiceResponseBodyDataResultVersions) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayServiceResponseBodyDataResultVersions) GoString() string {
	return s.String()
}

func (s *ListGatewayServiceResponseBodyDataResultVersions) GetName() *string {
	return s.Name
}

func (s *ListGatewayServiceResponseBodyDataResultVersions) SetName(v string) *ListGatewayServiceResponseBodyDataResultVersions {
	s.Name = &v
	return s
}

func (s *ListGatewayServiceResponseBodyDataResultVersions) Validate() error {
	return dara.Validate(s)
}
