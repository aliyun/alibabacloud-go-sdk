// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIPVersion(v string) *CreateServerGroupRequest
	GetAddressIPVersion() *string
	SetAnyPortEnabled(v bool) *CreateServerGroupRequest
	GetAnyPortEnabled() *bool
	SetClientToken(v string) *CreateServerGroupRequest
	GetClientToken() *string
	SetConnectionDrainEnabled(v bool) *CreateServerGroupRequest
	GetConnectionDrainEnabled() *bool
	SetConnectionDrainTimeout(v int32) *CreateServerGroupRequest
	GetConnectionDrainTimeout() *int32
	SetDryRun(v bool) *CreateServerGroupRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest
	GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig
	SetIpVersionAffinityMode(v string) *CreateServerGroupRequest
	GetIpVersionAffinityMode() *string
	SetPreserveClientIpEnabled(v bool) *CreateServerGroupRequest
	GetPreserveClientIpEnabled() *bool
	SetProtocol(v string) *CreateServerGroupRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateServerGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateServerGroupRequest
	GetResourceGroupId() *string
	SetScheduler(v string) *CreateServerGroupRequest
	GetScheduler() *string
	SetServerGroupName(v string) *CreateServerGroupRequest
	GetServerGroupName() *string
	SetServerGroupType(v string) *CreateServerGroupRequest
	GetServerGroupType() *string
	SetTag(v []*CreateServerGroupRequestTag) *CreateServerGroupRequest
	GetTag() []*CreateServerGroupRequestTag
	SetVpcId(v string) *CreateServerGroupRequest
	GetVpcId() *string
}

type CreateServerGroupRequest struct {
	// The protocol version. Valid values:
	//
	// - **ipv4*	- (default): IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// Specifies whether to enable all-port forwarding. Valid values:
	//
	// - **true**: enables all-port forwarding.
	//
	// - **false*	- (default): disables all-port forwarding.
	//
	// example:
	//
	// false
	AnyPortEnabled *bool `json:"AnyPortEnabled,omitempty" xml:"AnyPortEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable connection draining. Valid values:
	//
	// - **true**: enables connection draining.
	//
	// - **false*	- (default): disables connection draining.
	//
	// example:
	//
	// false
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period for connection draining. Unit: seconds.
	//
	// Valid values: **0*	- to **900**.
	//
	// example:
	//
	// 10
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the server group. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The health check configuration.
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The traffic scheduling policy in dual-stack mode. Valid values:
	//
	// - NonAffinity (default): non-affinity mode. Traffic is forwarded to healthy backend servers based on the scheduling algorithm regardless of the IP protocol version of the source.
	//
	// - Affinity: affinity mode. Traffic is forwarded based on the IP protocol version of the request source. IPv4 requests are forwarded only to IPv4 backend servers, and IPv6 requests are forwarded only to IPv6 backend servers.
	//
	// > This parameter takes effect only when **AddressIPVersion*	- is set to **DualStack**.
	//
	// example:
	//
	// Affinity
	IpVersionAffinityMode *string `json:"IpVersionAffinityMode,omitempty" xml:"IpVersionAffinityMode,omitempty"`
	// Specifies whether to enable client IP address preservation. Valid values:
	//
	// - **true*	- (default): enables client IP address preservation.
	//
	// - **false**: disables client IP address preservation.
	//
	//
	// >
	//
	// > If **Protocol*	- is set to **TCP*	- and this parameter is set to **true**, the server group cannot be associated with **TCPSSL*	- listeners.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	PreserveClientIpEnabled *bool `json:"PreserveClientIpEnabled,omitempty" xml:"PreserveClientIpEnabled,omitempty"`
	// The communication protocol used to forward traffic to backend servers. Valid values:
	//
	// - **TCP*	- (default)
	//
	// - **UDP**
	//
	// - **TCP_UDP**
	//
	// >
	//
	// > - If this parameter is set to **UDP**, the server group can be associated only with **UDP*	- listeners.
	//
	// > - If this parameter is set to **TCP*	- and **PreserveClientIpEnabled*	- is set to **true**, the server group can be associated only with **TCP*	- listeners.
	//
	// > - If this parameter is set to **TCP*	- and **PreserveClientIpEnabled*	- is set to **false**, the server group can be associated with **TCPSSL*	- and **TCP*	- listeners.
	//
	// > - If this parameter is set to **TCP_UDP**, the server group can be associated with **TCP*	- and **UDP*	- listeners.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the server group belongs.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **Wrr*	- (default): weighted round-robin. Backend servers with higher weights are more likely to be polled.
	//
	// - **Wlc**: weighted least connections. In addition to the weight of each backend server, the actual load (number of connections) of the backend server is also considered. If the weights are the same, backend servers with fewer current connections are more likely to be polled.
	//
	// - **rr**: round-robin. External requests are distributed to backend servers in sequence.
	//
	// - **sch**: source IP hash. Requests from the same source IP address are scheduled to the same backend server.
	//
	// - **tch**: 4-tuple hash. Consistent hashing based on the 4-tuple (source IP address, destination IP address, source port, and destination port). Requests from the same flow are scheduled to the same backend server.
	//
	// - **qch**: QUIC ID hash. Requests with the same QUIC ID are scheduled to the same backend server.
	//
	// > QUIC ID hash is supported only when the backend protocol is UDP.
	//
	// example:
	//
	// Wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The name of the server group.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// NLB_ServerGroup
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The server group type. Valid values:
	//
	// - **Instance*	- (default): server type. This type of server group supports adding **Ecs**, **Eni**, and **Eci*	- instances.
	//
	// - **Ip**: IP address type. This type of server group supports directly adding backend servers by IP address.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The list of tags.
	//
	// if can be null:
	// true
	Tag []*CreateServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC where the server group is deployed.
	//
	// > If **ServerGroupType*	- is set to **Instance**, only servers in this VPC can be added to the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateServerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateServerGroupRequest) GetAnyPortEnabled() *bool {
	return s.AnyPortEnabled
}

func (s *CreateServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServerGroupRequest) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *CreateServerGroupRequest) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *CreateServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServerGroupRequest) GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *CreateServerGroupRequest) GetIpVersionAffinityMode() *string {
	return s.IpVersionAffinityMode
}

func (s *CreateServerGroupRequest) GetPreserveClientIpEnabled() *bool {
	return s.PreserveClientIpEnabled
}

func (s *CreateServerGroupRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateServerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServerGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerGroupRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateServerGroupRequest) GetServerGroupName() *string {
	return s.ServerGroupName
}

func (s *CreateServerGroupRequest) GetServerGroupType() *string {
	return s.ServerGroupType
}

func (s *CreateServerGroupRequest) GetTag() []*CreateServerGroupRequestTag {
	return s.Tag
}

func (s *CreateServerGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateServerGroupRequest) SetAddressIPVersion(v string) *CreateServerGroupRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateServerGroupRequest) SetAnyPortEnabled(v bool) *CreateServerGroupRequest {
	s.AnyPortEnabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainEnabled(v bool) *CreateServerGroupRequest {
	s.ConnectionDrainEnabled = &v
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

func (s *CreateServerGroupRequest) SetIpVersionAffinityMode(v string) *CreateServerGroupRequest {
	s.IpVersionAffinityMode = &v
	return s
}

func (s *CreateServerGroupRequest) SetPreserveClientIpEnabled(v bool) *CreateServerGroupRequest {
	s.PreserveClientIpEnabled = &v
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

func (s *CreateServerGroupRequest) SetTag(v []*CreateServerGroupRequestTag) *CreateServerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerGroupRequest) Validate() error {
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
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

type CreateServerGroupRequestHealthCheckConfig struct {
	// The port of the backend server used for health checks.
	//
	// Valid values: **0*	- to **65535**.
	//
	// Default value: **0**. This value indicates that the port of the backend server is used for health checks.
	//
	// example:
	//
	// 0
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period for a health check response. Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// Default value: **5**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks. Valid values:
	//
	// - **$SERVER_IP**: the internal IP address of the backend server.
	//
	// - **domain**: a specified domain name. The domain name must be 1 to 80 characters in length, and can contain only letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **true*	- (default): enables health checks.
	//
	// - **false**: disables health checks.
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The query string in the request for UDP listener health checks. The string can contain only letters and digits, and can be up to 512 characters in length.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The HTTP status codes for healthy responses. Separate multiple status codes with commas (,).
	//
	// Valid values: **http_2xx*	- (default), **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The HTTP protocol version for health checks. Valid values: **HTTP1.0*	- (default) and **HTTP1.1**.
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// HTTP1.0
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Default value: **5**.
	//
	// - If **HealthCheckType*	- is set to **TCP*	- or **HTTP**, valid values are **1*	- to **50**.
	//
	// - If **HealthCheckType*	- is set to **UDP**, valid values are **1*	- to **300**. You must set the interval to be greater than or equal to the response timeout period to ensure that UDP probe responses are not misjudged as unresponsive due to a timeout.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The query string in the request for UDP listener health checks. The string can contain only letters and digits, and can be up to 512 characters in length.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The protocol used for health checks. Valid values:
	//
	// - **TCP**
	//
	// - **HTTP**
	//
	// - **UDP**
	//
	// example:
	//
	// TCP
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check path.
	//
	// The path must be 1 to 80 characters in length, and can contain only letters, digits, and the following special characters: `-/.%?#&`. It must start with a forward slash (/).
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of consecutive successful health checks that must occur before the health check status of a backend server is changed from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// Default value: **2**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP method used for health checks. Valid values: **GET*	- (default) and **HEAD**.
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// GET
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" xml:"HttpCheckMethod,omitempty"`
	// The number of consecutive failed health checks that must occur before the health check status of a backend server is changed from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// Default value: **2**.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateServerGroupRequestHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckHttpCode() []*string {
	return s.HealthCheckHttpCode
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHttpCheckMethod() *string {
	return s.HttpCheckMethod
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckExp(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckExp = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckReq(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckReq = &v
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHttpCheckMethod(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HttpCheckMethod = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type CreateServerGroupRequestTag struct {
	// The tag key. The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. It can contain letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. It can contain letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServerGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServerGroupRequestTag) SetKey(v string) *CreateServerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServerGroupRequestTag) SetValue(v string) *CreateServerGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServerGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
