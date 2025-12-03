// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetAclId() *string
	SetAclIds(v *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetAclIds() *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds
	SetAclStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetAclStatus() *string
	SetAclType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetAclType() *string
	SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetBandwidth() *int32
	SetConnectionDrain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetConnectionDrain() *string
	SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetConnectionDrainTimeout() *int32
	SetDescription(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetDescription() *string
	SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetEstablishedTimeout() *int32
	SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheck() *string
	SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckConnectPort() *int32
	SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckConnectTimeout() *int32
	SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckHttpCode() *string
	SetHealthCheckInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckInterval() *int32
	SetHealthCheckMethod(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckMethod() *string
	SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetHealthyThreshold() *int32
	SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetMasterSlaveServerGroupId() *string
	SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetPersistenceTimeout() *int32
	SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetProxyProtocolV2Enabled() *bool
	SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetRequestId() *string
	SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetScheduler() *string
	SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetStatus() *string
	SetSynProxy(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetSynProxy() *string
	SetTags(v *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetTags() *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags
	SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody
	GetVServerGroupId() *string
}

type DescribeLoadBalancerTCPListenerAttributeResponseBody struct {
	// The ID of the network ACL that is associated with the listener.
	//
	// If **AclStatus*	- is set to **on**, this parameter is returned.
	//
	// example:
	//
	// acl-uf60jwfi******
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The IDs of the ACLs.
	AclIds *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds `json:"AclIds,omitempty" xml:"AclIds,omitempty" type:"Struct"`
	// Indicates whether access control is enabled. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The type of the ACL. Valid values:
	//
	// 	- **white**: a whitelist. Only requests from the IP addresses or CIDR blocks in the network ACL are forwarded. Whitelists apply to scenarios in which you want to allow only specific IP addresses to access an application.
	//
	//     Your service may be adversely affected if the whitelist is not properly configured. After a whitelist is configured, only requests from IP addresses that are added to the whitelist are forwarded by the listener.
	//
	//     If you enable a whitelist but do not add an IP address to the ACL, the listener forwards all requests.
	//
	// 	- **black**: a blacklist. All requests from the IP addresses or CIDR blocks in the network ACL are rejected. Blacklists apply to scenarios in which you want to block access from specified IP addresses to an application.
	//
	//     If a blacklist is configured for a listener but no IP address is added to the blacklist, the listener forwards all requests.
	//
	// >  If **AclStatus*	- is set to **on**, this parameter is returned.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The backend port used by the CLB instance.
	//
	// >  If the listener is associated with a vServer group, this parameter is not returned.
	//
	// example:
	//
	// 443
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The maximum bandwidth of the listener. Unit: Mbit/s. Valid values:
	//
	// 	- **-1**: For a pay-by-data-transfer Internet-facing CLB instance, this parameter is set to -1. This indicates that the bandwidth of the listener is unlimited.
	//
	// 	- **1*	- to **5120**: For a pay-by-bandwidth Internet-facing CLB instance, you can specify the maximum bandwidth of each listener. The sum of maximum bandwidth of all listeners cannot exceed the maximum bandwidth of the CLB instance.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// Indicates whether connection draining is enabled. If **ConnectionDrain*	- is set to **on**, the parameter is returned. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// off
	ConnectionDrain *string `json:"ConnectionDrain,omitempty" xml:"ConnectionDrain,omitempty"`
	// The timeout period of connection draining. If **ConnectionDrain*	- is set to **on**, the parameter is returned.
	//
	// Valid values: 10 to 900. Unit: seconds.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// The description of the listener.
	//
	// example:
	//
	// TCP_80
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timeout period of a connection.
	//
	// example:
	//
	// 500
	EstablishedTimeout *int32 `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// example:
	//
	// on
	HealthCheck *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**. If this parameter is not set, the port specified by BackendServerPort is used for health checks.
	//
	// example:
	//
	// 8080
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$_ip**: the private IP addresses of backend servers. If you do not set the HealthCheckDomain parameter or set the parameter to $_ip, the CLB instance uses the private IP address of each backend server for health checks.
	//
	// 	- **domain**: The domain name is 1 to 80 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// www.domain.com
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// The HTTP status code for a successful health check.
	//
	// example:
	//
	// http_2xx
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	// The interval between two consecutive health checks. Valid values: **1*	- to **50**. Unit: seconds.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The health check method.
	//
	// example:
	//
	// tcp
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The health check method that is used by the TCP listener.
	//
	// Valid values: **tcp*	- and **http**.
	//
	// example:
	//
	// tcp
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The URL that is used for health checks. The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL is not a single forward slash (/) but it starts with a forward slash (/).
	//
	// example:
	//
	// /test/index.html
	HealthCheckURI *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	// The healthy threshold. The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status is changed from **fail*	- to **success**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The frontend port used by the CLB instance.
	//
	// example:
	//
	// 110
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1ygod3yctvg1y****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The ID of the primary/secondary server group that is associated with the listener.
	//
	// example:
	//
	// rsp-0bfucw****
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	// The timeout period of session persistence.
	//
	// Valid values: **0*	- to **3600**. Unit: seconds. Default value: **0**. If the default value is used, the system disables session persistence.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// Indicates whether the Proxy protocol is used to pass client IP addresses to backend servers. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ProxyProtocolV2Enabled *bool `json:"ProxyProtocolV2Enabled,omitempty" xml:"ProxyProtocolV2Enabled,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scheduling algorithm.
	//
	// 	- **wrr*	- (default): Backend servers with higher weights receive more requests than backend servers with lower weights.
	//
	// 	- **rr**: Requests are distributed to backend servers in sequence.
	//
	// 	- **sch**: specifies consistent hashing that is based on source IP addresses. Requests from the same source IP address are distributed to the same backend server.
	//
	// 	- **tch**: specifies consistent hashing that is based on four factors: source IP address, destination IP address, source port, and destination port. Requests that contain the same information based on the four factors are distributed to the same backend server.
	//
	// > Only high-performance CLB instances support the sch and tch algorithms.
	//
	// example:
	//
	// wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **running**
	//
	// 	- **stopped**
	//
	// example:
	//
	// stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the SynProxy feature of CLB is enabled for protection.
	//
	// We recommend that you use the default value of this parameter. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// disable
	SynProxy *string `json:"SynProxy,omitempty" xml:"SynProxy,omitempty"`
	// The tags.
	Tags *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The unhealthy threshold. The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status is changed from **success*	- to **fail**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	// The ID of the associated server group.
	//
	// example:
	//
	// rsp-cige6******8
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetAclId() *string {
	return s.AclId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetAclIds() *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds {
	return s.AclIds
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetAclType() *string {
	return s.AclType
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetConnectionDrain() *string {
	return s.ConnectionDrain
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetProxyProtocolV2Enabled() *bool {
	return s.ProxyProtocolV2Enabled
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetScheduler() *string {
	return s.Scheduler
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetSynProxy() *string {
	return s.SynProxy
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetTags() *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclIds(v *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclIds = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclStatus = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetAclType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.AclType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBackendServerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectionDrain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectionDrain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetConnectionDrainTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetDescription(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetEstablishedTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.EstablishedTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheck(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheck = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckConnectTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckDomain(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckDomain = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckHttpCode(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckInterval(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckMethod(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckMethod = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckType(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthCheckURI(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthCheckURI = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetHealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetMasterSlaveServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetPersistenceTimeout(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetProxyProtocolV2Enabled(v bool) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.ProxyProtocolV2Enabled = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetScheduler(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Scheduler = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetStatus(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetSynProxy(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.SynProxy = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetTags(v *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetUnhealthyThreshold(v int32) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.UnhealthyThreshold = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) SetVServerGroupId(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBody {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBody) Validate() error {
	if s.AclIds != nil {
		if err := s.AclIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds struct {
	AclId []*string `json:"AclId,omitempty" xml:"AclId,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) GetAclId() []*string {
	return s.AclId
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) SetAclId(v []*string) *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds {
	s.AclId = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyAclIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerTCPListenerAttributeResponseBodyTags struct {
	Tag []*DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) GetTag() []*DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) SetTag(v []*DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTags) Validate() error {
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

type DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag value cannot be an empty string. The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N. Valid values of N: **1*	- to **20**. The tag value can be an empty string. The tag value can be up to 128 characters in length, and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) SetTagKey(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) SetTagValue(v string) *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
