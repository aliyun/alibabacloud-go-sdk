// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateServerGroupAttributeRequest
	GetClientToken() *string
	SetConnectionDrainConfig(v *UpdateServerGroupAttributeRequestConnectionDrainConfig) *UpdateServerGroupAttributeRequest
	GetConnectionDrainConfig() *UpdateServerGroupAttributeRequestConnectionDrainConfig
	SetCrossZoneEnabled(v bool) *UpdateServerGroupAttributeRequest
	GetCrossZoneEnabled() *bool
	SetDryRun(v bool) *UpdateServerGroupAttributeRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest
	GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig
	SetScheduler(v string) *UpdateServerGroupAttributeRequest
	GetScheduler() *string
	SetServerGroupId(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupId() *string
	SetServerGroupName(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupName() *string
	SetServiceName(v string) *UpdateServerGroupAttributeRequest
	GetServiceName() *string
	SetSlowStartConfig(v *UpdateServerGroupAttributeRequestSlowStartConfig) *UpdateServerGroupAttributeRequest
	GetSlowStartConfig() *UpdateServerGroupAttributeRequestSlowStartConfig
	SetStickySessionConfig(v *UpdateServerGroupAttributeRequestStickySessionConfig) *UpdateServerGroupAttributeRequest
	GetStickySessionConfig() *UpdateServerGroupAttributeRequestStickySessionConfig
	SetUchConfig(v *UpdateServerGroupAttributeRequestUchConfig) *UpdateServerGroupAttributeRequest
	GetUchConfig() *UpdateServerGroupAttributeRequestUchConfig
	SetUpstreamKeepaliveEnabled(v bool) *UpdateServerGroupAttributeRequest
	GetUpstreamKeepaliveEnabled() *bool
}

type UpdateServerGroupAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of connection draining.
	//
	// After connection draining is enabled, SLB remains data transmission for a period of time after a backend server is removed or declared unhealthy.
	//
	// > 	- Basic SLB instances do not support connection draining. Standard and WAF-enabled SLB instances support connection draining.
	//
	// > 	- Server groups of the server and IP types support connection draining. Server groups of the Function Compute type do not support connection draining.
	ConnectionDrainConfig *UpdateServerGroupAttributeRequestConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// Indicates whether cross-zone load balancing is enabled for the server group. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// > 	- Basic ALB instances do not support server groups that have cross-zone load balancing disabled. Only Standard and WAF-enabled ALB instances support server groups that have cross-zone load balancing.
	//
	// >	- Cross-zone load balancing can be disabled for server groups of the server and IP type, but not for server groups of the Function Compute type.
	//
	// >	- When cross-zone load balancing is disabled, session persistence cannot be enabled.
	//
	// example:
	//
	// true
	CrossZoneEnabled *bool `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx` HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration of health checks.
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **Wrr**: the weighted round robin algorithm. Backend servers that have higher weights receive more requests than those that have lower weights.
	//
	// 	- **Wlc**: the weighted least connections algorithm. Requests are distributed based on the weights and the number of connections to backend servers. If two backend servers have the same weight, the backend server that has fewer connections is expected to receive more requests.
	//
	// 	- **Sch**: the consistent hashing algorithm. Requests from the same source IP address are distributed to the same backend server.
	//
	// example:
	//
	// Wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtop****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server group name.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// This parameter is available only if the ALB Ingress controller is used. In this case, set this parameter to the name of the `Kubernetes Service` that is associated with the server group.
	//
	// example:
	//
	// test2
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The configurations of slow starts.
	//
	// After slow starts are enabled, ALB prefetches data to newly added backend servers. Requests distributed to the backend servers gradually increase.
	//
	// > 	- Basic ALB instances do not support slow starts. Standard and WAF-enabled ALB instances support slow starts.
	//
	// >	- Server groups of the instance and IP types support slow starts. Server groups of the Function Compute type do not support slow starts.
	//
	// >	- Slow start is supported only by the weighted round-robin scheduling algorithm.
	SlowStartConfig *UpdateServerGroupAttributeRequestSlowStartConfig `json:"SlowStartConfig,omitempty" xml:"SlowStartConfig,omitempty" type:"Struct"`
	// The configuration of session persistence.
	StickySessionConfig *UpdateServerGroupAttributeRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// The configurations of consistent hashing based on URLs.
	UchConfig *UpdateServerGroupAttributeRequestUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
	// Specifies whether to enable persistent TCP connections.
	//
	// example:
	//
	// true
	UpstreamKeepaliveEnabled *bool `json:"UpstreamKeepaliveEnabled,omitempty" xml:"UpstreamKeepaliveEnabled,omitempty"`
}

func (s UpdateServerGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServerGroupAttributeRequest) GetConnectionDrainConfig() *UpdateServerGroupAttributeRequestConnectionDrainConfig {
	return s.ConnectionDrainConfig
}

func (s *UpdateServerGroupAttributeRequest) GetCrossZoneEnabled() *bool {
	return s.CrossZoneEnabled
}

func (s *UpdateServerGroupAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServerGroupAttributeRequest) GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig {
	return s.HealthCheckConfig
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

func (s *UpdateServerGroupAttributeRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateServerGroupAttributeRequest) GetSlowStartConfig() *UpdateServerGroupAttributeRequestSlowStartConfig {
	return s.SlowStartConfig
}

func (s *UpdateServerGroupAttributeRequest) GetStickySessionConfig() *UpdateServerGroupAttributeRequestStickySessionConfig {
	return s.StickySessionConfig
}

func (s *UpdateServerGroupAttributeRequest) GetUchConfig() *UpdateServerGroupAttributeRequestUchConfig {
	return s.UchConfig
}

func (s *UpdateServerGroupAttributeRequest) GetUpstreamKeepaliveEnabled() *bool {
	return s.UpstreamKeepaliveEnabled
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainConfig(v *UpdateServerGroupAttributeRequestConnectionDrainConfig) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetCrossZoneEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.CrossZoneEnabled = &v
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

func (s *UpdateServerGroupAttributeRequest) SetServiceName(v string) *UpdateServerGroupAttributeRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetSlowStartConfig(v *UpdateServerGroupAttributeRequestSlowStartConfig) *UpdateServerGroupAttributeRequest {
	s.SlowStartConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetStickySessionConfig(v *UpdateServerGroupAttributeRequestStickySessionConfig) *UpdateServerGroupAttributeRequest {
	s.StickySessionConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetUchConfig(v *UpdateServerGroupAttributeRequestUchConfig) *UpdateServerGroupAttributeRequest {
	s.UchConfig = v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetUpstreamKeepaliveEnabled(v bool) *UpdateServerGroupAttributeRequest {
	s.UpstreamKeepaliveEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupAttributeRequestConnectionDrainConfig struct {
	// Specifies whether to enable connection draining. Valid values:
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
	// Valid values: **0*	- to **900**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
}

func (s UpdateServerGroupAttributeRequestConnectionDrainConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *UpdateServerGroupAttributeRequestConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *UpdateServerGroupAttributeRequestConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// The HTTP status codes that indicate healthy backend servers.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The backend port that is used for health checks.
	//
	// Valid values: **0*	- to **65535**.
	//
	// If you set the value to **0**, the backend port is used for health checks.
	//
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to **true**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
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
	//     	- The domain name must be 1 to 80 characters in length.
	//
	//     	- The domain name can contain lowercase letters, digits, hyphens (-), and periods (.).
	//
	//     	- The domain name must contain at least one period (.) but cannot start or end with a period (.).
	//
	//     	- The rightmost domain label of the domain name can contain only letters, and cannot contain digits or hyphens (-).
	//
	//     	- The domain name cannot start or end with a hyphen (-).
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// example.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The HTTP version that is used for health checks. Valid values:
	//
	// 	- **HTTP1.0**
	//
	// 	- **HTTP1.1**
	//
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to true and **HealthCheckProtocol*	- to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP1.1
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to **true**.
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
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to true and **HealthCheckProtocol*	- to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: `- / . % ? # & =`. It can also contain the following extended characters: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`. The URL must start with a forward slash (`/`).
	//
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to **true*	- and **HealthCheckProtocol*	- to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that you want to use for health checks. Valid values:
	//
	// 	- **HTTP**: HTTP health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers.
	//
	// 	- **HTTPS**: HTTPS health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers. HTTPS supports encryption and provides higher security than HTTP.
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to probe the availability of backend servers.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to check whether the backend server is healthy.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The timeout period of a health check response. If a backend ECS instance does not respond within the specified timeout period, the ECS instance fails the health check. Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// >  This parameter takes effect only if you set **HealthCheckEnabled*	- to **true**.
	//
	// example:
	//
	// 3
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it can be declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it can be declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 4
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckHost(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckHost = &v
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckMethod(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckPath(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupAttributeRequestSlowStartConfig struct {
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

func (s UpdateServerGroupAttributeRequestSlowStartConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestSlowStartConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestSlowStartConfig) GetSlowStartDuration() *int32 {
	return s.SlowStartDuration
}

func (s *UpdateServerGroupAttributeRequestSlowStartConfig) GetSlowStartEnabled() *bool {
	return s.SlowStartEnabled
}

func (s *UpdateServerGroupAttributeRequestSlowStartConfig) SetSlowStartDuration(v int32) *UpdateServerGroupAttributeRequestSlowStartConfig {
	s.SlowStartDuration = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestSlowStartConfig) SetSlowStartEnabled(v bool) *UpdateServerGroupAttributeRequestSlowStartConfig {
	s.SlowStartEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestSlowStartConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupAttributeRequestStickySessionConfig struct {
	// The cookie to be configured on the server.
	//
	// The cookie must be 1 to 200 characters in length and can contain only ASCII characters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// > This parameter takes effect when the **StickySessionEnabled*	- parameter is set to **true*	- and the **StickySessionType*	- parameter is set to **Server**.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA1598
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The timeout period of a cookie. Unit: seconds.
	//
	// Valid values: **1*	- to **86400**.
	//
	// > This parameter takes effect when the **StickySessionEnabled*	- parameter is set to **true*	- and the **StickySessionType*	- parameter is set to **Insert**.
	//
	// example:
	//
	// 1000
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// The method that is used to handle a cookie. Valid values:
	//
	// 	- **Insert**: inserts a cookie.
	//
	// ALB inserts a cookie (SERVERID) into the first HTTP or HTTPS response packet that is sent to a client. The next request from the client contains this cookie and the listener forwards this request to the recorded backend server.
	//
	// 	- **Server**: rewrites a cookie.
	//
	// When ALB detects a user-defined cookie, it overwrites the original cookie with the user-defined cookie. Subsequent requests to ALB carry this user-defined cookie, and ALB determines the destination servers of the requests based on the cookies.
	//
	// > This parameter takes effect when the **StickySessionEnabled*	- parameter is set to **true*	- for the server group.
	//
	// example:
	//
	// Insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) GetCookie() *string {
	return s.Cookie
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) GetStickySessionEnabled() *bool {
	return s.StickySessionEnabled
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookie(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetCookieTimeout(v int32) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionEnabled(v bool) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) SetStickySessionType(v string) *UpdateServerGroupAttributeRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestStickySessionConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateServerGroupAttributeRequestUchConfig struct {
	// The type of the parameter. Only query strings are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// QueryString
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the parameter used for consistent hashing.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateServerGroupAttributeRequestUchConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestUchConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestUchConfig) GetType() *string {
	return s.Type
}

func (s *UpdateServerGroupAttributeRequestUchConfig) GetValue() *string {
	return s.Value
}

func (s *UpdateServerGroupAttributeRequestUchConfig) SetType(v string) *UpdateServerGroupAttributeRequestUchConfig {
	s.Type = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestUchConfig) SetValue(v string) *UpdateServerGroupAttributeRequestUchConfig {
	s.Value = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestUchConfig) Validate() error {
	return dara.Validate(s)
}
