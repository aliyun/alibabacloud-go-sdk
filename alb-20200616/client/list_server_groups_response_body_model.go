// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServerGroupsResponseBody
	GetRequestId() *string
	SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody
	GetServerGroups() []*ListServerGroupsResponseBodyServerGroups
	SetTotalCount(v int32) *ListServerGroupsResponseBody
	GetTotalCount() *int32
}

type ListServerGroupsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If **NextToken*	- is not empty, the value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server groups.
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerGroupsResponseBody) GetServerGroups() []*ListServerGroupsResponseBodyServerGroups {
	return s.ServerGroups
}

func (s *ListServerGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListServerGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroups struct {
	// Indicates whether configuration management is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConfigManagedEnabled *bool `json:"ConfigManagedEnabled,omitempty" xml:"ConfigManagedEnabled,omitempty"`
	// The configurations of connection draining.
	//
	// After connection draining is enabled, ALB maintains data transmission for a period of time after the backend server is removed or declared unhealthy.
	//
	// >
	//
	// > - Basic ALB instances do not support connection draining. Standard and WAF-enabled ALB instances support connection draining.
	//
	// > -  Server groups of the instance and IP types support connection draining. Server groups of the Function Compute type do not support connection draining.
	ConnectionDrainConfig *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether cross-zone load balancing is enabled. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CrossZoneEnabled *bool `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	// The health check configurations.
	HealthCheckConfig *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Indicates whether IPv6 is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Ipv6Enabled *bool `json:"Ipv6Enabled,omitempty" xml:"Ipv6Enabled,omitempty"`
	// The backend protocol. Valid values:
	//
	// 	- **HTTP**: allows you to associate HTTPS, HTTP, or QUIC listeners with backend servers.
	//
	// 	- **HTTPS**: allows you to associate HTTPS listeners with backend servers.
	//
	// 	- **GRPC**: allows you to associate HTTPS and QUIC listeners with backend servers.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the ALB instance associated with the server group.
	RelatedLoadBalancerIds []*string `json:"RelatedLoadBalancerIds,omitempty" xml:"RelatedLoadBalancerIds,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **Wrr**: weighted round-robin. Backend servers with higher weights receive more requests than backend servers with lower weights.
	//
	// 	- **Wlc**: weighted least connections. Requests are distributed based on the weight and load of each backend server. The load refers to the number of connections on a backend server. If multiple backend servers have the same weight, requests are forwarded to the backend server with the least number of connections.
	//
	// 	- **Sch**: consistent hashing. Requests that have the same hash factors are distributed to the same backend server. If you do not specify the UchConfig parameter, the source IP address is used as the hash factor by default. Requests that are from the same IP address are distributed to the same backend server. If you specify the UchConfig parameter, the URL string is used as the hash factor. Requests that have the same URL string are distributed to the same backend server.
	//
	// example:
	//
	// Wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of backend servers in the server group.
	//
	// example:
	//
	// 1
	ServerCount *int32 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-cige6j****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server group name.
	//
	// example:
	//
	// Group3
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The status of the server group. Valid values:
	//
	// 	- **Creating**.
	//
	// 	- **Available**
	//
	// 	- **Configuring**
	//
	// example:
	//
	// Available
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// The server group type. Valid values:
	//
	// 	- **Instance**: instances, including ECS instances, ENIs, and elastic container instances.
	//
	// 	- **Ip**: IP addresses.
	//
	// 	- **Fc**: Function Compute
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The name of the server group.
	//
	// example:
	//
	// test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The configurations of slow starts.
	//
	// After slow starts are enabled, ALB prefetches data to newly added backend servers. Requests distributed to the backend servers gradually increase.
	//
	// >
	//
	// > - Basic ALB instances do not support slow starts. Standard and WAF-enabled ALB instances support slow starts.
	//
	// > - Server groups of the instance and IP types support slow starts. Server groups of the Function Compute type do not support slow starts.
	//
	// > - Slow start is supported only by the weighted round-robin scheduling algorithm.
	SlowStartConfig *ListServerGroupsResponseBodyServerGroupsSlowStartConfig `json:"SlowStartConfig,omitempty" xml:"SlowStartConfig,omitempty" type:"Struct"`
	// The configuration of session persistence.
	StickySessionConfig *ListServerGroupsResponseBodyServerGroupsStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// The tags that are added to the server group.
	Tags []*ListServerGroupsResponseBodyServerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The configuration of consistent hashing based on URLs.
	UchConfig *ListServerGroupsResponseBodyServerGroupsUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
	// Indicates whether persistent TCP connections are enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	UpstreamKeepaliveEnabled *bool `json:"UpstreamKeepaliveEnabled,omitempty" xml:"UpstreamKeepaliveEnabled,omitempty"`
	// The ID of the VPC to which the ALB instance belongs.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) GetConfigManagedEnabled() *bool {
	return s.ConfigManagedEnabled
}

func (s *ListServerGroupsResponseBodyServerGroups) GetConnectionDrainConfig() *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	return s.ConnectionDrainConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServerGroupsResponseBodyServerGroups) GetCrossZoneEnabled() *bool {
	return s.CrossZoneEnabled
}

func (s *ListServerGroupsResponseBodyServerGroups) GetHealthCheckConfig() *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetIpv6Enabled() *bool {
	return s.Ipv6Enabled
}

func (s *ListServerGroupsResponseBodyServerGroups) GetProtocol() *string {
	return s.Protocol
}

func (s *ListServerGroupsResponseBodyServerGroups) GetRelatedLoadBalancerIds() []*string {
	return s.RelatedLoadBalancerIds
}

func (s *ListServerGroupsResponseBodyServerGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListServerGroupsResponseBodyServerGroups) GetScheduler() *string {
	return s.Scheduler
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerCount() *int32 {
	return s.ServerCount
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupName() *string {
	return s.ServerGroupName
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupStatus() *string {
	return s.ServerGroupStatus
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServerGroupType() *string {
	return s.ServerGroupType
}

func (s *ListServerGroupsResponseBodyServerGroups) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServerGroupsResponseBodyServerGroups) GetSlowStartConfig() *ListServerGroupsResponseBodyServerGroupsSlowStartConfig {
	return s.SlowStartConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetStickySessionConfig() *ListServerGroupsResponseBodyServerGroupsStickySessionConfig {
	return s.StickySessionConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetTags() []*ListServerGroupsResponseBodyServerGroupsTags {
	return s.Tags
}

func (s *ListServerGroupsResponseBodyServerGroups) GetUchConfig() *ListServerGroupsResponseBodyServerGroupsUchConfig {
	return s.UchConfig
}

func (s *ListServerGroupsResponseBodyServerGroups) GetUpstreamKeepaliveEnabled() *bool {
	return s.UpstreamKeepaliveEnabled
}

func (s *ListServerGroupsResponseBodyServerGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConfigManagedEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.ConfigManagedEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConnectionDrainConfig(v *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) *ListServerGroupsResponseBodyServerGroups {
	s.ConnectionDrainConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetCreateTime(v string) *ListServerGroupsResponseBodyServerGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetCrossZoneEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.CrossZoneEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheckConfig(v *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheckConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetIpv6Enabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.Ipv6Enabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetRelatedLoadBalancerIds(v []*string) *ListServerGroupsResponseBodyServerGroups {
	s.RelatedLoadBalancerIds = v
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

func (s *ListServerGroupsResponseBodyServerGroups) SetServerCount(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ServerCount = &v
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

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServiceName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServiceName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetSlowStartConfig(v *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) *ListServerGroupsResponseBodyServerGroups {
	s.SlowStartConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetStickySessionConfig(v *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) *ListServerGroupsResponseBodyServerGroups {
	s.StickySessionConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetTags(v []*ListServerGroupsResponseBodyServerGroupsTags) *ListServerGroupsResponseBodyServerGroups {
	s.Tags = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetUchConfig(v *ListServerGroupsResponseBodyServerGroupsUchConfig) *ListServerGroupsResponseBodyServerGroups {
	s.UchConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetUpstreamKeepaliveEnabled(v bool) *ListServerGroupsResponseBodyServerGroups {
	s.UpstreamKeepaliveEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig struct {
	// Indicates whether connection draining is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsHealthCheckConfig struct {
	// The HTTP status codes that indicate healthy backend servers.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The backend port that is used for health checks. Valid values: **0*	- to **65535**.
	//
	// A value of **0*	- indicates that the port of a backend server is used for health checks.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name that is used for health checks.
	//
	// 	- **Backend Server Internal IP*	- (default): Use the internal IP address of backend servers as the health check domain name.
	//
	// 	- **Custom Domain Name**: Enter a domain name.
	//
	//     	- The domain name is 1 to 80 characters in length.
	//
	//     	- The domain name contains lowercase letters, digits, hyphens (-), and periods (.).
	//
	//     	- The domain name contains at least one period (.) but does not start or end with a period (.).
	//
	//     	- The rightmost domain label of the domain name contains only letters, and does not contain digits or hyphens (-).
	//
	//     	- The domain name does not start or end with a hyphen (-).
	//
	// >  This parameter takes effect only if HealthCheckProtocol is set to HTTP, HTTPS, or gRPC.
	//
	// example:
	//
	// www.example.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks.
	//
	// Valid values: **HTTP1.0*	- and **HTTP1.1**.
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP1.1
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// 	- **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// 	- **POST**: gRPC health checks use the POST method by default.
	//
	// 	- **HEAD**: HTTP and HTTPS health checks use the HEAD method by default.
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **HTTP**: HTTP health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers.
	//
	// 	- **HTTPS**: HTTPS health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers. HTTPS supports encryption and provides higher security than HTTP.
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to check whether the backend server is healthy.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the backend server is declared unhealthy. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckCodes(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
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

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsSlowStartConfig struct {
	// The duration of a slow start.
	//
	// example:
	//
	// 30
	SlowStartDuration *int32 `json:"SlowStartDuration,omitempty" xml:"SlowStartDuration,omitempty"`
	// Indicates whether slow starts are enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SlowStartEnabled *bool `json:"SlowStartEnabled,omitempty" xml:"SlowStartEnabled,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsSlowStartConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsSlowStartConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) GetSlowStartDuration() *int32 {
	return s.SlowStartDuration
}

func (s *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) GetSlowStartEnabled() *bool {
	return s.SlowStartEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) SetSlowStartDuration(v int32) *ListServerGroupsResponseBodyServerGroupsSlowStartConfig {
	s.SlowStartDuration = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) SetSlowStartEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsSlowStartConfig {
	s.SlowStartEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsSlowStartConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsStickySessionConfig struct {
	// The cookie configured for the server.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA****
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of the cookie. Unit: seconds. Valid values: **1*	- to **86400**.
	//
	// >  This parameter takes effect only when **StickySessionEnabled*	- is set to **true*	- and **StickySessionType*	- is set to **Insert**.
	//
	// example:
	//
	// 1000
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Indicates whether session persistence is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// The method that is used to handle the cookie. Valid values:
	//
	// 	- **insert**: inserts the cookie. The first time a client accesses ALB, ALB inserts the SERVERID cookie into the HTTP or HTTPS response packet. Subsequent requests from the client that carry this cookie are forwarded to the same backend server as the first request.
	//
	// 	- **Server**: rewrites the cookie. ALB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// example:
	//
	// Insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GetCookie() *string {
	return s.Cookie
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GetStickySessionEnabled() *bool {
	return s.StickySessionEnabled
}

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) GetStickySessionType() *string {
	return s.StickySessionType
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

func (s *ListServerGroupsResponseBodyServerGroupsStickySessionConfig) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsTags struct {
	// The tag key.
	//
	// example:
	//
	// Test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsTags) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetKey(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Key = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Value = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) Validate() error {
	return dara.Validate(s)
}

type ListServerGroupsResponseBodyServerGroupsUchConfig struct {
	// The parameter type. Valid value: QueryString.
	//
	// example:
	//
	// QueryString
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The hash value.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsUchConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsUchConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) GetType() *string {
	return s.Type
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) GetValue() *string {
	return s.Value
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) SetType(v string) *ListServerGroupsResponseBodyServerGroupsUchConfig {
	s.Type = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsUchConfig {
	s.Value = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsUchConfig) Validate() error {
	return dara.Validate(s)
}
