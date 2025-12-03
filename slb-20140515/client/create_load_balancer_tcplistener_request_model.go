// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerTCPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateLoadBalancerTCPListenerRequest
	GetAclId() *string
	SetAclStatus(v string) *CreateLoadBalancerTCPListenerRequest
	GetAclStatus() *string
	SetAclType(v string) *CreateLoadBalancerTCPListenerRequest
	GetAclType() *string
	SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerTCPListenerRequest
	GetBandwidth() *int32
	SetConnectionDrain(v string) *CreateLoadBalancerTCPListenerRequest
	GetConnectionDrain() *string
	SetConnectionDrainTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetConnectionDrainTimeout() *int32
	SetDescription(v string) *CreateLoadBalancerTCPListenerRequest
	GetDescription() *string
	SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetEstablishedTimeout() *int32
	SetHealthCheckConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckSwitch(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckSwitch() *string
	SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest
	GetOwnerId() *int64
	SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetPersistenceTimeout() *int32
	SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerTCPListenerRequest
	GetProxyProtocolV2Enabled() *bool
	SetRegionId(v string) *CreateLoadBalancerTCPListenerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest
	GetScheduler() *string
	SetTag(v []*CreateLoadBalancerTCPListenerRequestTag) *CreateLoadBalancerTCPListenerRequest
	GetTag() []*CreateLoadBalancerTCPListenerRequestTag
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest
	GetVServerGroupId() *string
	SetHealthCheckInterval(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckInterval() *int32
}

type CreateLoadBalancerTCPListenerRequest struct {
	// The ID of the network ACL that is associated with the listener.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// acl-uf60jwfiv6******
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
	// The type of the ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios where you want to allow only specific IP addresses to access an application.
	//
	//     Your service may be adversely affected if the whitelist is not properly configured.
	//
	//     If a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener. If a whitelist is configured but no IP address is added to the whitelist, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the ACL are rejected. Blacklists apply to scenarios where you want to block access from specified IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is required.
	//
	// example:
	//
	// black
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
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, this value can be set to -1, which specifies unlimited bandwidth.
	//
	// 	- **1*	- to **5120**: For a pay-by-bandwidth Internet-facing CLB instance, you can specify the maximum bandwidth of each listener. The sum of the maximum bandwidth values that you set for all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
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
	// The timeout period of connection draining. Unit: seconds.
	//
	// Valid values: **10*	- to **900**.
	//
	// >  This parameter is required if **ConnectionDrain*	- is set to **on**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// The name of the listener.
	//
	// The name must be 1 to 256 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// tcp_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timeout period of a connection. Unit: seconds
	//
	// Valid values: **10*	- to **900**.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
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
	// The maximum timeout period of a health check response. Unit: seconds
	//
	// Valid values: **1*	- to **300**
	//
	// Default value: **5**
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that you want to use for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP address of a backend server. If you do not set the HealthCheckDomain parameter or set the parameter to $_ip, the CLB instance uses the private IP address of each backend server for health checks.
	//
	// 	- **domain**: The domain name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// 172.XX.XX.6
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check. Separate multiple HTTP status codes with commas (,). Valid values:
	//
	// 	- **http_2xx**(default)
	//
	// 	- **http_3xx**
	//
	// 	- **http_4xx**
	//
	// 	- **http_5xx**
	//
	// example:
	//
	// http_2xx,http_3xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
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
	// The type of health checks. Valid values:
	//
	// 	- **tcp*	- (default)
	//
	// 	- **http**
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URI that is used for health checks. The URI must be 1 to 80 characters in length, and can contain only digits, letters, hyphens (-), forward slashes (/), periods (.), percent signs (%), number signs (#), and ampersands (&). The URI must start with a forward slash (/) but cannot be a single forward slash (/).
	//
	// You can set this parameter when the TCP listener requires HTTP health checks. If you do not set this parameter, TCP health checks are performed.
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
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
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the primary/secondary server group.
	//
	// >  You cannot set both VServerGroupId and MasterSlaveServerGroupId.
	//
	// example:
	//
	// rsp-0bfucw****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timeout period of session persistence. Unit: seconds
	//
	// Valid values: **0 to 3600**
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
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *bool `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// For the most recent region list, refer to [Regions and zones](https://help.aliyun.com/document_detail/40654.html) or call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation the retrieve the information.
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
	// >  Only high-performance CLB instances support the **sch*	- and **tch*	- consistent hashing algorithms.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerTCPListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// rsp-cige6j****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 3
	HealthCheckInterval *int32 `json:"healthCheckInterval,omitempty" xml:"healthCheckInterval,omitempty"`
}

func (s CreateLoadBalancerTCPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerRequest) GetAclId() *string {
	return s.AclId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *CreateLoadBalancerTCPListenerRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateLoadBalancerTCPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerTCPListenerRequest) GetConnectionDrain() *string {
	return s.ConnectionDrain
}

func (s *CreateLoadBalancerTCPListenerRequest) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLoadBalancerTCPListenerRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckSwitch() *string {
	return s.HealthCheckSwitch
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerTCPListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *CreateLoadBalancerTCPListenerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerTCPListenerRequest) GetTag() []*CreateLoadBalancerTCPListenerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerTCPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclStatus(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAclType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AclType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectionDrain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectionDrain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectionDrainTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetDescription(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Description = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckSwitch(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckSwitch = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetProxyProtocolV2Enabled(v bool) *CreateLoadBalancerTCPListenerRequest {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetRegionId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetTag(v []*CreateLoadBalancerTCPListenerRequestTag) *CreateLoadBalancerTCPListenerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckInterval(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) Validate() error {
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

type CreateLoadBalancerTCPListenerRequestTag struct {
	// The key of the tag. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be at most 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerTCPListenerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerTCPListenerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerTCPListenerRequestTag) SetKey(v string) *CreateLoadBalancerTCPListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequestTag) SetValue(v string) *CreateLoadBalancerTCPListenerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequestTag) Validate() error {
	return dara.Validate(s)
}
