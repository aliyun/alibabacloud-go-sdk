// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerTCPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetAclId() *string
	SetAclStatus(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetAclStatus() *string
	SetAclType(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetAclType() *string
	SetBandwidth(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetBandwidth() *int32
	SetConnectionDrain(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetConnectionDrain() *string
	SetConnectionDrainTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetConnectionDrainTimeout() *int32
	SetDescription(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetDescription() *string
	SetEstablishedTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckSwitch(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckSwitch() *string
	SetHealthCheckType(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetMasterSlaveServerGroup() *string
	SetMasterSlaveServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest
	GetOwnerId() *int64
	SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetPersistenceTimeout() *int32
	SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerTCPListenerAttributeRequest
	GetProxyProtocolV2Enabled() *bool
	SetRegionId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetScheduler() *string
	SetSynProxy(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetSynProxy() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetVServerGroup() *string
	SetVServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest
	GetVServerGroupId() *string
}

type SetLoadBalancerTCPListenerAttributeRequest struct {
	// The ID of the network access control list (ACL) that is associated with the listener.
	//
	// If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// 12333
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the network ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the allowlist is not properly configured. After a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	//     If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are rejected. Blacklists apply to scenarios where you want to block access from specified IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values: **-1*	- and **1*	- to **5120**.
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, you can set this parameter to **-1**, which specifies unlimited bandwidth.
	//
	// 	- **1*	- to **5120**: For a pay-by-bandwidth Internet-facing CLB instance, you can specify the maximum bandwidth of each listener. The sum of the maximum bandwidth values of all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// 43
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// Specifies whether to enable connection draining. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	ConnectionDrain *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	// The timeout period of connection draining. This parameter is required if **ConnectionDrain*	- is set to **on**. Unit: seconds.
	//
	// Valid values: **10*	- to **900**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// tcp_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timeout period of a connection. Unit: seconds. Valid values: **10*	- to **900**.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**.
	//
	// If you do not set this parameter, the port specified by the **BackendServerPort*	- parameter is used.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period of a health check.
	//
	// If a backend ECS instance does not return a health check response within the specified timeout period, the server fails the health check.
	//
	// Valid values: **1*	- to **300**. Unit: seconds.
	//
	// >  If the value of the **HealthCheckConnectTimeout*	- parameter is smaller than that of the **HealthCheckInterval*	- parameter, the timeout period specified by the **HCTimeout*	- parameter is ignored and the period of time specified by the **HealthCheckInterval*	- parameter is used as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. You can set this parameter when the TCP listener requires HTTP health checks. If you do not set this parameter, TCP health checks are performed.
	//
	// 	- **$_ip**: the private IP addresses of the backend servers.
	//
	//     If you do not set the HealthCheckHost parameter or set the parameter to $SERVER_IP, the CLB instance uses the private IP addresses of backend servers for health checks.
	//
	// 	- **domain**: The domain name is 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// 192.168.XX.XX
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (,).
	//
	// Valid values: **http_2xx**, **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	HealthCheckSwitch *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	// The type of the health check. Valid values: **tcp*	- and **http**.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI that is used for health checks. The URI must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
	//
	// You can set this parameter when the TCP listener requires HTTP health checks.
	//
	// If you do not set this parameter, TCP health checks are performed.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1ygod3yctvg1y****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Specifies whether to use a primary/secondary server group. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// You cannot set both **VserverGroup*	- and **MasterSlaveServerGroup*	- to **on**.
	//
	// example:
	//
	// on
	MasterSlaveServerGroup *string `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty"`
	// The ID of the primary/secondary server group.
	//
	// >  You can set only one of the VServerGroupId and MasterSlaveServerGroupId parameters.
	//
	// example:
	//
	// rsp-cige****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timeout period of session persistence. Valid values: **0*	- to **3600**. Unit: seconds.
	//
	// Default value: **0**. If the default value is used, the system disables session persistence.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *bool `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](~~DescribeRegions~~) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers that have higher weights receive more requests than backend servers that have lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: specifies consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **tch**: specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are distributed to the same backend server.
	//
	// > 	- Only high-performance CLB instances support the **sch*	- and **tch*	- algorithms.
	//
	// > 	- CLB does not support converting the **wrr*	- and **rr*	- algorithms to sch or tch. You cannot switch the hash algorithm from one to another.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Specifies whether to enable the SynProxy feature of CLB for protection. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// We recommend that you use the default value of this parameter.
	//
	// example:
	//
	// enable
	SynProxy *string `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// Specifies whether to use a vServer group. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// >  You cannot set both **VserverGroup*	- and **MasterSlaveServerGroup*	- to **on**.
	//
	// example:
	//
	// on
	VServerGroup *string `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j5****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetAclType() *string {
	return s.AclType
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetConnectionDrain() *string {
	return s.ConnectionDrain
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckSwitch() *string {
	return s.HealthCheckSwitch
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetMasterSlaveServerGroup() *string {
	return s.MasterSlaveServerGroup
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetSynProxy() *string {
	return s.SynProxy
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetVServerGroup() *string {
	return s.VServerGroup
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectionDrain(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectionDrain = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetConnectionDrainTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetEstablishedTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckDomain(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckHttpCode(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckSwitch(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckType(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckType = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckURI(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetMasterSlaveServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.MasterSlaveServerGroup = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetMasterSlaveServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetPersistenceTimeout(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetSynProxy(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.SynProxy = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerTCPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerTCPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
