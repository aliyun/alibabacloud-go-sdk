// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerUDPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateLoadBalancerUDPListenerRequest
	GetAclId() *string
	SetAclStatus(v string) *CreateLoadBalancerUDPListenerRequest
	GetAclStatus() *string
	SetAclType(v string) *CreateLoadBalancerUDPListenerRequest
	GetAclType() *string
	SetBackendServerPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerUDPListenerRequest
	GetBandwidth() *int32
	SetDescription(v string) *CreateLoadBalancerUDPListenerRequest
	GetDescription() *string
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckSwitch(v string) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckSwitch() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *CreateLoadBalancerUDPListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerUDPListenerRequest
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest
	GetOwnerId() *int64
	SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerUDPListenerRequest
	GetProxyProtocolV2Enabled() *bool
	SetRegionId(v string) *CreateLoadBalancerUDPListenerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerUDPListenerRequest
	GetScheduler() *string
	SetTag(v []*CreateLoadBalancerUDPListenerRequestTag) *CreateLoadBalancerUDPListenerRequest
	GetTag() []*CreateLoadBalancerUDPListenerRequestTag
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest
	GetVServerGroupId() *string
	SetHealthCheckExp(v string) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckExp() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckInterval() *int32
	SetHealthCheckReq(v string) *CreateLoadBalancerUDPListenerRequest
	GetHealthCheckReq() *string
}

type CreateLoadBalancerUDPListenerRequest struct {
	// The ID of the network ACL that is associated with the listener.
	//
	// If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// 123
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off*	- (default): no
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the network ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios in which you want to allow only specific IP addresses to access an application. After a whitelist is configured, only IP addresses in the whitelist can access the CLB listener. Risks may arise if the whitelist is improperly set.
	//
	//     If a whitelist is configured but no IP address is added to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are blocked. Blacklists apply to scenarios in which you want to deny access from specific IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// If the **VServerGroupId*	- parameter is not set, this parameter is required.
	//
	// example:
	//
	// 80
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
	//
	// **-1**: For a pay-by-data-transfer Internet-facing CLB instance, you can set this parameter to **-1**. This way, the bandwidth of the listener is unlimited.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// udp_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// If this parameter is not set, the backend port specified by **BackendServerPort*	- is used for health checks.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period of a health check.
	//
	// If a backend server, such as an Elastic Compute Service (ECS) instance, does not respond to a probe packet within the specified timeout period, the server fails the health check. Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
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
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**
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
	// lb-bp1ygod3yctvg1y7****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the primary/secondary server group.
	//
	// >  You can set only one of the VServerGroupId and MasterSlaveServerGroupId parameters.
	//
	// example:
	//
	// rsp-0bfucwu****
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
	// The ID of the region where the CLB instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than those with lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: specifies consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **tch**: specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are distributed to the same backend server.
	//
	// 	- **qch**: specifies consistent hashing that is based on QUIC connection IDs. Requests that contain the same QUIC connection ID are distributed to the same backend server.
	//
	// Only high-performance CLB instances support the sch, tch, and qch consistent hashing algorithms.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerUDPListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-cige6j8****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The response string for UDP listener health checks. The string must be 1 to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"healthCheckExp,omitempty" xml:"healthCheckExp,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"healthCheckInterval,omitempty" xml:"healthCheckInterval,omitempty"`
	// The request string for UDP listener health checks. The string must be 1 to 64 characters in length and can contain only letters and digits.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"healthCheckReq,omitempty" xml:"healthCheckReq,omitempty"`
}

func (s CreateLoadBalancerUDPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerRequest) GetAclId() *string {
	return s.AclId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *CreateLoadBalancerUDPListenerRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateLoadBalancerUDPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerUDPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckSwitch() *string {
	return s.HealthCheckSwitch
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerUDPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerUDPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerUDPListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *CreateLoadBalancerUDPListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerUDPListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerUDPListenerRequest) GetTag() []*CreateLoadBalancerUDPListenerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerUDPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerUDPListenerRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerUDPListenerRequest) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetAclType(v string) *CreateLoadBalancerUDPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetDescription(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckSwitch(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerUDPListenerRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetRegionId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerUDPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerUDPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetScheduler(v string) *CreateLoadBalancerUDPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetTag(v []*CreateLoadBalancerUDPListenerRequestTag) *CreateLoadBalancerUDPListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerUDPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckExp(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckExp = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) SetHealthCheckReq(v string) *CreateLoadBalancerUDPListenerRequest {
	s.HealthCheckReq = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLoadBalancerUDPListenerRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerUDPListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerUDPListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerUDPListenerRequestTag) SetKey(v string) *CreateLoadBalancerUDPListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequestTag) SetValue(v string) *CreateLoadBalancerUDPListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerRequestTag) Validate() error {
	return dara.Validate(s)
}
