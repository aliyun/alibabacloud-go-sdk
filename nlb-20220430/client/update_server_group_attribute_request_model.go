// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIPVersion(v string) *UpdateServerGroupAttributeRequest
	GetAddressIPVersion() *string
	SetClientToken(v string) *UpdateServerGroupAttributeRequest
	GetClientToken() *string
	SetConnectionDrainEnabled(v bool) *UpdateServerGroupAttributeRequest
	GetConnectionDrainEnabled() *bool
	SetConnectionDrainTimeout(v int32) *UpdateServerGroupAttributeRequest
	GetConnectionDrainTimeout() *int32
	SetDryRun(v bool) *UpdateServerGroupAttributeRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest
	GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig
	SetIpVersionAffinityMode(v string) *UpdateServerGroupAttributeRequest
	GetIpVersionAffinityMode() *string
	SetPreserveClientIpEnabled(v bool) *UpdateServerGroupAttributeRequest
	GetPreserveClientIpEnabled() *bool
	SetRegionId(v string) *UpdateServerGroupAttributeRequest
	GetRegionId() *string
	SetScheduler(v string) *UpdateServerGroupAttributeRequest
	GetScheduler() *string
	SetServerGroupId(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupId() *string
	SetServerGroupName(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupName() *string
}

type UpdateServerGroupAttributeRequest struct {
	// The IP version. Valid values: **ipv4*	- and **dualStack**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
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
	// - **true**: Enables connection draining.
	//
	// - **false**: Disables connection draining.
	//
	// example:
	//
	// false
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period for connection draining. Unit: seconds. Valid values: **0*	- to **900**.
	//
	// example:
	//
	// 10
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without updating the server group. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The health check configuration.
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The traffic scheduling policy in dual-stack mode. Valid values:
	//
	// - **NonAffinity**: non-affinity mode. Requests are forwarded to healthy backend servers based on the scheduling algorithm regardless of the IP version of the source address. This is the default value.
	//
	// - **Affinity**: affinity mode. Requests are forwarded to backend servers based on the IP version of the source address. IPv4 requests are forwarded only to IPv4 backend servers, and IPv6 requests are forwarded only to IPv6 backend servers.
	//
	// > This parameter takes effect only when **AddressIPVersion*	- is set to **DualStack**.
	//
	// example:
	//
	// Affinity
	IpVersionAffinityMode *string `json:"IpVersionAffinityMode,omitempty" xml:"IpVersionAffinityMode,omitempty"`
	// Specifies whether to enable client IP preservation. Valid values:
	//
	// - **true**: Enables client IP preservation.
	//
	// - **false**: Disables client IP preservation.
	//
	// > If **PreserveClientIpEnabled*	- is set to **false*	- and the server group is associated with a **TCPSSL*	- listener, you cannot change the value of this parameter to **true**.
	//
	// example:
	//
	// false
	PreserveClientIpEnabled *bool `json:"PreserveClientIpEnabled,omitempty" xml:"PreserveClientIpEnabled,omitempty"`
	// The region ID of the network load balancing instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to obtain the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// - **Wrr**: weighted round-robin. Backend servers with higher weights receive more requests.
	//
	// - **Wlc**: weighted least connections. Requests are distributed based on the weights and actual loads (number of connections) of backend servers. If weights are the same, backend servers with fewer connections receive more requests.
	//
	// - **rr**: round-robin. Requests are distributed to backend servers in sequence.
	//
	// - **sch**: source IP hash. Requests from the same source IP address are distributed to the same backend server.
	//
	// - **tch**: 4-tuple hash. Consistent hashing based on the 4-tuple (source IP address, destination IP address, source port, and destination port). Requests from the same flow are distributed to the same backend server.
	//
	// - **qch**: QUIC ID hash. Requests with the same QUIC ID are distributed to the same backend server.
	//
	// > QUIC ID hash is supported only when the backend protocol is UDP.
	//
	// example:
	//
	// Wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The ID of the server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The new name of the server group.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// NLB_ServerGroup1
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *UpdateServerGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServerGroupAttributeRequest) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *UpdateServerGroupAttributeRequest) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *UpdateServerGroupAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServerGroupAttributeRequest) GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *UpdateServerGroupAttributeRequest) GetIpVersionAffinityMode() *string {
	return s.IpVersionAffinityMode
}

func (s *UpdateServerGroupAttributeRequest) GetPreserveClientIpEnabled() *bool {
	return s.PreserveClientIpEnabled
}

func (s *UpdateServerGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServerGroupAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *UpdateServerGroupAttributeRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateServerGroupAttributeRequest) GetServerGroupName() *string {
	return s.ServerGroupName
}

func (s *UpdateServerGroupAttributeRequest) SetAddressIPVersion(v string) *UpdateServerGroupAttributeRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainEnabled = &v
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

func (s *UpdateServerGroupAttributeRequest) SetIpVersionAffinityMode(v string) *UpdateServerGroupAttributeRequest {
	s.IpVersionAffinityMode = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetPreserveClientIpEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.PreserveClientIpEnabled = &v
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

func (s *UpdateServerGroupAttributeRequest) SetServerGroupName(v string) *UpdateServerGroupAttributeRequest {
	s.ServerGroupName = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) Validate() error {
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// The port used for health checks on backend servers. Valid values: **0*	- to **65535**. A value of **0*	- indicates that the port of the backend server is used for health checks.
	//
	// example:
	//
	// 0
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period for a health check response. Unit: seconds. Valid values: **1*	- to **300**. Default value: **5**.
	//
	// example:
	//
	// 100
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks. Valid values:
	//
	// - **$SERVER_IP**: the internal IP address of the backend server.
	//
	// - **domain**: a custom domain name. The domain name must be 1 to 80 characters in length, and can contain only lowercase letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **true**: Enables health checks.
	//
	// - **false**: Disables health checks.
	//
	// example:
	//
	// false
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The response string used for UDP health checks. The string can contain only letters and digits, and can be up to 512 characters in length.
	//
	// example:
	//
	// ok
	HealthCheckExp *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The HTTP status codes that indicate a successful health check. Separate multiple HTTP status codes with commas (,). Valid values: **http_2xx*	- (default), **http_3xx**, **http_4xx**, and **http_5xx**.
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The HTTP version used for health checks. Valid values: **HTTP1.0*	- and **HTTP1.1**.
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
	// - If **HealthCheckType*	- is set to **UDP**, valid values are **1*	- to **300**. You must set the interval to a value greater than or equal to the response timeout period to ensure that UDP probe responses are not mistakenly identified as timeouts.
	//
	// example:
	//
	// 5
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The request string used for UDP health checks. The string can contain only letters and digits, and can be up to 512 characters in length.
	//
	// example:
	//
	// hello
	HealthCheckReq *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The health check protocol. Valid values:
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
	// The path must be 1 to 80 characters in length. It can contain only letters, digits, the characters `-/.%?#&=`, and extension characters `_;~!()*[]@$^:\\",+`. It must start with a forward slash (/).
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The number of consecutive successful health checks required to change the health status of a backend server from **Failed*	- to **Success**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The HTTP method used for health checks. Valid values: **GET*	- and **HEAD**.
	//
	// > This parameter takes effect only when **HealthCheckType*	- is set to **HTTP**.
	//
	// example:
	//
	// GET
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" xml:"HttpCheckMethod,omitempty"`
	// The number of consecutive failed health checks required to change the health status of a backend server from **Success*	- to **Failed**. Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckConnectTimeout() *int32 {
	return s.HealthCheckConnectTimeout
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckExp() *string {
	return s.HealthCheckExp
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckHttpCode() []*string {
	return s.HealthCheckHttpCode
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckUrl() *string {
	return s.HealthCheckUrl
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHttpCheckMethod() *string {
	return s.HttpCheckMethod
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckExp(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckExp = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHttpVersion(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHttpVersion = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckReq(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckReq = &v
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHttpCheckMethod(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HttpCheckMethod = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}
