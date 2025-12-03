// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerUDPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetAclId() *string
	SetAclStatus(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetAclStatus() *string
	SetAclType(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetAclType() *string
	SetBandwidth(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetBandwidth() *int32
	SetDescription(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetDescription() *string
	SetHealthCheckConnectPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckInterval(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckSwitch(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckSwitch() *string
	SetHealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetMasterSlaveServerGroup() *string
	SetMasterSlaveServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest
	GetOwnerId() *int64
	SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerUDPListenerAttributeRequest
	GetProxyProtocolV2Enabled() *bool
	SetRegionId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetScheduler() *string
	SetUnhealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetVServerGroup() *string
	SetVServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetVServerGroupId() *string
	SetHealthCheckExp(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckExp() *string
	SetHealthCheckReq(v string) *SetLoadBalancerUDPListenerAttributeRequest
	GetHealthCheckReq() *string
}

type SetLoadBalancerUDPListenerAttributeRequest struct {
	// The ID of the network access control list (ACL) that is associated with the listener.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// off
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
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application. Your service may be adversely affected if the whitelist is not properly configured. After a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	//     If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are denied. Blacklists apply to scenarios where you want to block access from specified IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, you can set this parameter to **-1**. This way, the bandwidth of the listener is unlimited.
	//
	// 	- **1*	- to **5120**: For a pay-by-bandwidth Internet-facing CLB instance, you can specify the maximum bandwidth of each listener. The sum of bandwidth limits that you set for all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 256 characters in length and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// udp_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period of a health check. If a backend server, such as an Elastic Compute Service (ECS) instance, does not return a health check response within the specified timeout period, the server fails the health check. Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// >  If the value of the **HealthCheckConnectTimeout*	- parameter is smaller than that of the **HealthCheckInterval*	- parameter, the timeout period specified by the **HealthCheckConnectTimeout*	- parameter is ignored and the period of time specified by the **HealthCheckInterval*	- parameter is used as the timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **on*	- (default): yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	HealthCheckSwitch *string `json:"HealthCheckSwitch,omitempty" xml:"HealthCheckSwitch,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it can be declared healthy (from **fail*	- to **success**).
	//
	// Valid values: **1*	- to **10**.
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
	// lb-bp1rtfnodmywb43ecu4sf-c****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// Specifies whether to use a primary/secondary server group. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// >  You cannot set **VserverGroup*	- and **MasterSlaveServerGroup*	- both to **on**.
	//
	// example:
	//
	// on
	MasterSlaveServerGroup *string `json:"MasterSlaveServerGroup,omitempty" xml:"MasterSlaveServerGroup,omitempty"`
	// The ID of the primary/secondary server group.
	//
	// >  You cannot specify both VServerGroupId and MasterSlaveServerGroupId.
	//
	// example:
	//
	// rsp-0bfuc****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use the Proxy protocol to pass client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *bool `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr**: Backend servers with higher weights receive more requests than those with lower weights.
	//
	//     If two backend servers have the same weight, the backend server that has fewer connections is expected to receive more requests.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: specifies consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **tch**: specifies consistent hashing that is based on the following parameters: source IP address, destination IP address, source port, and destination port. Requests that contain the same preceding information are distributed to the same backend server.
	//
	// 	- **qch**: specifies consistent hashing that is based on QUIC connection IDs. Requests that contain the same QUIC connection ID are distributed to the same backend server.
	//
	// > 	- Only high-performance CLB instances support **sch**, **tch**, and **qch**.
	//
	// > 	- You cannot switch the algorithm used by a CLB instance from **wrr*	- or **rr*	- to consistent hashing or from consistent hashing to weighted round robin or round robin.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **1*	- to **10**.
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
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The response string for UDP listener health checks. The string must be 1 to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"healthCheckExp,omitempty" xml:"healthCheckExp,omitempty"`
	// The request string for UDP listener health checks. The string must be 1 to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"healthCheckReq,omitempty" xml:"healthCheckReq,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetAclId() *string {
	return s.AclId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetAclType() *string {
	return s.AclType
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckSwitch() *string {
	return s.HealthCheckSwitch
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetMasterSlaveServerGroup() *string {
	return s.MasterSlaveServerGroup
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetVServerGroup() *string {
	return s.VServerGroup
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclStatus(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclStatus = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetAclType(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.AclType = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetBandwidth(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetDescription(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Description = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectTimeout(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckInterval(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckSwitch(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetMasterSlaveServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.MasterSlaveServerGroup = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetMasterSlaveServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetProxyProtocolV2Enabled(v bool) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetRegionId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(v int64) *SetLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetScheduler(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.Scheduler = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetUnhealthyThreshold(v int32) *SetLoadBalancerUDPListenerAttributeRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetVServerGroup(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.VServerGroup = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetVServerGroupId(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckExp(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckReq(v string) *SetLoadBalancerUDPListenerAttributeRequest {
	s.HealthCheckReq = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
