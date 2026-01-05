// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServerGroupRequest
	GetClientToken() *string
	SetConnectionDrainConfig(v *CreateServerGroupRequestConnectionDrainConfig) *CreateServerGroupRequest
	GetConnectionDrainConfig() *CreateServerGroupRequestConnectionDrainConfig
	SetCrossZoneEnabled(v bool) *CreateServerGroupRequest
	GetCrossZoneEnabled() *bool
	SetDryRun(v bool) *CreateServerGroupRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest
	GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig
	SetIpv6Enabled(v bool) *CreateServerGroupRequest
	GetIpv6Enabled() *bool
	SetProtocol(v string) *CreateServerGroupRequest
	GetProtocol() *string
	SetResourceGroupId(v string) *CreateServerGroupRequest
	GetResourceGroupId() *string
	SetScheduler(v string) *CreateServerGroupRequest
	GetScheduler() *string
	SetServerGroupName(v string) *CreateServerGroupRequest
	GetServerGroupName() *string
	SetServerGroupType(v string) *CreateServerGroupRequest
	GetServerGroupType() *string
	SetServiceName(v string) *CreateServerGroupRequest
	GetServiceName() *string
	SetSlowStartConfig(v *CreateServerGroupRequestSlowStartConfig) *CreateServerGroupRequest
	GetSlowStartConfig() *CreateServerGroupRequestSlowStartConfig
	SetStickySessionConfig(v *CreateServerGroupRequestStickySessionConfig) *CreateServerGroupRequest
	GetStickySessionConfig() *CreateServerGroupRequestStickySessionConfig
	SetTag(v []*CreateServerGroupRequestTag) *CreateServerGroupRequest
	GetTag() []*CreateServerGroupRequestTag
	SetUchConfig(v *CreateServerGroupRequestUchConfig) *CreateServerGroupRequest
	GetUchConfig() *CreateServerGroupRequestUchConfig
	SetUpstreamKeepaliveEnabled(v bool) *CreateServerGroupRequest
	GetUpstreamKeepaliveEnabled() *bool
	SetVpcId(v string) *CreateServerGroupRequest
	GetVpcId() *string
}

type CreateServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of connection draining.
	//
	// After connection draining is enabled, ALB maintains data transmission for a period of time after the backend server is removed or declared unhealthy.
	//
	//
	// >	- Basic ALB instances do not support connection draining. Standard and WAF-enabled ALB instances support connection draining.
	//
	// >	- Server groups of the instance and IP types support connection draining. Server groups of the Function Compute type do not support connection draining.
	ConnectionDrainConfig *CreateServerGroupRequestConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// Specifies whether to enable cross-zone load balancing. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// > 	- Basic ALB instances do not support server groups that have cross-zone load balancing disabled. Only Standard and WAF-enabled ALB instances support server groups that have cross-zone load balancing.
	//
	// > 	- Cross-zone load balancing can be disabled for server groups of the server and IP type, but not for server groups of the Function Compute type.
	//
	// > 	- When cross-zone load balancing is disabled, session persistence cannot be enabled.
	//
	// example:
	//
	// true
	CrossZoneEnabled *bool `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configuration of the health check feature.
	//
	// This parameter is required.
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// Specifies whether to enable Ipv6.
	//
	// example:
	//
	// false
	Ipv6Enabled *bool `json:"Ipv6Enabled,omitempty" xml:"Ipv6Enabled,omitempty"`
	// The backend protocol. Valid values:
	//
	// 	- **HTTP**: allows you to associate an HTTPS, HTTP, or QUIC listener with the server group. This is the default value.
	//
	// 	- **HTTPS**: allows you to associate HTTPS listeners with backend servers.
	//
	// 	- **gRPC**: allows you to associate an HTTPS or QUIC listener with the server group.
	//
	// >  You do not need to specify a backend protocol if you set **ServerGroupType*	- to **Fc**.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **Wrr*	- (default): The weighted round-robin algorithm is used. Backend servers that have higher weights receive more requests than those that have lower weights.
	//
	// 	- **Wlc**: The weighted least connections algorithm is used. Requests are distributed based on the weights and the number of connections to backend servers. If two backend servers have the same weight, the backend server that has fewer connections is expected to receive more requests.
	//
	// 	- **Sch**: The consistent hashing algorithm is used. Requests from the same source IP address are distributed to the same backend server.
	//
	// > This parameter takes effect when the **ServerGroupType*	- parameter is set to **Instance*	- or **Ip**.
	//
	// example:
	//
	// Wrr
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The name of the server group. The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-atstuj3rtoptyui****
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The type of server group. Valid values:
	//
	// 	- **Instance*	- (default): allows you to add servers by specifying **Ecs**, **Eni**, or **Eci**.
	//
	// 	- **Ip**: allows you to add servers by specifying IP addresses.
	//
	// 	- **Fc**: allows you to add servers by specifying functions of Function Compute.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// This parameter is available only if the ALB Ingress controller is used. In this case, set this parameter to the name of the `Kubernetes Service` that is associated with the server group.
	//
	// example:
	//
	// test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The configurations of slow starts.
	//
	// After slow starts are enabled, SLB prefetches data to newly added backend servers. Requests distributed to the backend servers gradually increase.
	//
	// > - Basic SLB instances do not support slow starts. Standard and WAF-enabled SLB instances support slow starts.
	//
	// >	- Server groups of the server and IP types support slow starts. Server groups of the Function Compute type do not support slow starts.
	//
	// >	- Slow start is supported only by the weighted round-robin scheduling algorithm.
	SlowStartConfig *CreateServerGroupRequestSlowStartConfig `json:"SlowStartConfig,omitempty" xml:"SlowStartConfig,omitempty" type:"Struct"`
	// The configuration of session persistence.
	//
	// >  This parameter takes effect when the **ServerGroupType*	- parameter is set to **Instance*	- or **Ip**.
	StickySessionConfig *CreateServerGroupRequestStickySessionConfig `json:"StickySessionConfig,omitempty" xml:"StickySessionConfig,omitempty" type:"Struct"`
	// The tag.
	Tag []*CreateServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The configuration of consistent hashing based on URLs.
	UchConfig *CreateServerGroupRequestUchConfig `json:"UchConfig,omitempty" xml:"UchConfig,omitempty" type:"Struct"`
	// Specifies whether to enable persistent TCP connections.
	//
	// example:
	//
	// false
	UpstreamKeepaliveEnabled *bool `json:"UpstreamKeepaliveEnabled,omitempty" xml:"UpstreamKeepaliveEnabled,omitempty"`
	// The ID of the virtual private cloud (VPC). You can add only servers that are deployed in the specified VPC to the server group.
	//
	// >  This parameter takes effect when the **ServerGroupType*	- parameter is set to **Instance*	- or **Ip**.
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

func (s *CreateServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServerGroupRequest) GetConnectionDrainConfig() *CreateServerGroupRequestConnectionDrainConfig {
	return s.ConnectionDrainConfig
}

func (s *CreateServerGroupRequest) GetCrossZoneEnabled() *bool {
	return s.CrossZoneEnabled
}

func (s *CreateServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServerGroupRequest) GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *CreateServerGroupRequest) GetIpv6Enabled() *bool {
	return s.Ipv6Enabled
}

func (s *CreateServerGroupRequest) GetProtocol() *string {
	return s.Protocol
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

func (s *CreateServerGroupRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServerGroupRequest) GetSlowStartConfig() *CreateServerGroupRequestSlowStartConfig {
	return s.SlowStartConfig
}

func (s *CreateServerGroupRequest) GetStickySessionConfig() *CreateServerGroupRequestStickySessionConfig {
	return s.StickySessionConfig
}

func (s *CreateServerGroupRequest) GetTag() []*CreateServerGroupRequestTag {
	return s.Tag
}

func (s *CreateServerGroupRequest) GetUchConfig() *CreateServerGroupRequestUchConfig {
	return s.UchConfig
}

func (s *CreateServerGroupRequest) GetUpstreamKeepaliveEnabled() *bool {
	return s.UpstreamKeepaliveEnabled
}

func (s *CreateServerGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainConfig(v *CreateServerGroupRequestConnectionDrainConfig) *CreateServerGroupRequest {
	s.ConnectionDrainConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetCrossZoneEnabled(v bool) *CreateServerGroupRequest {
	s.CrossZoneEnabled = &v
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

func (s *CreateServerGroupRequest) SetIpv6Enabled(v bool) *CreateServerGroupRequest {
	s.Ipv6Enabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetProtocol(v string) *CreateServerGroupRequest {
	s.Protocol = &v
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

func (s *CreateServerGroupRequest) SetServiceName(v string) *CreateServerGroupRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServerGroupRequest) SetSlowStartConfig(v *CreateServerGroupRequestSlowStartConfig) *CreateServerGroupRequest {
	s.SlowStartConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetStickySessionConfig(v *CreateServerGroupRequestStickySessionConfig) *CreateServerGroupRequest {
	s.StickySessionConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetTag(v []*CreateServerGroupRequestTag) *CreateServerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateServerGroupRequest) SetUchConfig(v *CreateServerGroupRequestUchConfig) *CreateServerGroupRequest {
	s.UchConfig = v
	return s
}

func (s *CreateServerGroupRequest) SetUpstreamKeepaliveEnabled(v bool) *CreateServerGroupRequest {
	s.UpstreamKeepaliveEnabled = &v
	return s
}

func (s *CreateServerGroupRequest) SetVpcId(v string) *CreateServerGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateServerGroupRequest) Validate() error {
	if s.ConnectionDrainConfig != nil {
		if err := s.ConnectionDrainConfig.Validate(); err != nil {
			return err
		}
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlowStartConfig != nil {
		if err := s.SlowStartConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StickySessionConfig != nil {
		if err := s.StickySessionConfig.Validate(); err != nil {
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
	if s.UchConfig != nil {
		if err := s.UchConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServerGroupRequestConnectionDrainConfig struct {
	// Specifies whether to enable connection draining. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ConnectionDrainEnabled *bool `json:"ConnectionDrainEnabled,omitempty" xml:"ConnectionDrainEnabled,omitempty"`
	// The timeout period of connection draining.
	//
	// Valid values: **0*	- to **900**.
	//
	// Default value: **300**.
	//
	// example:
	//
	// 300
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
}

func (s CreateServerGroupRequestConnectionDrainConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestConnectionDrainConfig) GetConnectionDrainEnabled() *bool {
	return s.ConnectionDrainEnabled
}

func (s *CreateServerGroupRequestConnectionDrainConfig) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *CreateServerGroupRequestConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *CreateServerGroupRequestConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *CreateServerGroupRequestConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *CreateServerGroupRequestConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *CreateServerGroupRequestConnectionDrainConfig) Validate() error {
	return dara.Validate(s)
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// The HTTP status code that indicates healthy backend servers.
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitempty" xml:"HealthCheckCodes,omitempty" type:"Repeated"`
	// The backend port that is used for health checks.
	//
	// Valid values: **0*	- to **65535**
	//
	// The default value is **0**, which specifies that the port of a backend server is used for health checks.
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
	// >  If the **ServerGroupType*	- parameter is set to **Instance*	- or **Ip**, the health check feature is enabled by default. If the **ServerGroupType*	- parameter is set to **Fc**, the health check feature is disabled by default.
	//
	// This parameter is required.
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
	// www.example.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The version of the HTTP protocol. Valid values: **HTTP1.0*	- and **HTTP1.1**. Default value: HTTP1.1.
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP*	- or **HTTPS**.
	//
	// example:
	//
	// HTTP1.1
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitempty" xml:"HealthCheckHttpVersion,omitempty"`
	// The interval at which health checks are performed. Unit: seconds
	//
	// Valid values: **1*	- to **50**
	//
	// Default value: **2**.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The HTTP method that is used for health checks. Valid values:
	//
	// 	- **GET**: If the length of a response exceeds 8 KB, the response is truncated. However, the health check result is not affected.
	//
	// 	- **POST**: By default, gRPC health checks use the POST method.
	//
	// 	- **HEAD*	- (default): By default, HTTP and HTTPS use the HEAD method.
	//
	// >  This parameter takes effect only if **HealthCheckProtocol*	- is set to **HTTP**, **HTTPS**, or **gRPC**.
	//
	// example:
	//
	// HEAD
	HealthCheckMethod *string `json:"HealthCheckMethod,omitempty" xml:"HealthCheckMethod,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: `- / . % ? # & =`. It can also contain the following extended characters: `_ ; ~ ! ( ) 	- [ ] @ $ ^ : \\" , +`. The URL must start with a forward slash (/).
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
	// 	- **HTTPS**: HTTPS health checks simulate browser behaviors by sending HEAD or GET requests to probe the availability of backend servers. HTTPS provides higher security than HTTP because HTTPS supports data encryption.
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to probe the availability of backend servers.
	//
	// 	- **gRPC**: gRPC health checks send POST or GET requests to a backend server to check whether the backend server is healthy.
	//
	// example:
	//
	// HTTP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The timeout period of a health check response. If a backend server does not respond within the specified timeout period, the backend server is declared unhealthy. Unit: seconds
	//
	// Valid values: **1*	- to **300**
	//
	// Default value: **5**
	//
	// example:
	//
	// 5
	HealthCheckTimeout *int32 `json:"HealthCheckTimeout,omitempty" xml:"HealthCheckTimeout,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**
	//
	// Default value: **3**.
	//
	// example:
	//
	// 3
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**
	//
	// Default value: **3**
	//
	// example:
	//
	// 3
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s CreateServerGroupRequestHealthCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckCodes() []*string {
	return s.HealthCheckCodes
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckConnectPort() *int32 {
	return s.HealthCheckConnectPort
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckHttpVersion() *string {
	return s.HealthCheckHttpVersion
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckMethod() *string {
	return s.HealthCheckMethod
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckTimeout() *int32 {
	return s.HealthCheckTimeout
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckCodes(v []*string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckCodes = v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckConnectPort(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckEnabled(v bool) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckHost(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckHost = &v
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckMethod(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckMethod = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckPath(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckProtocol(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckTimeout(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckTimeout = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) Validate() error {
	return dara.Validate(s)
}

type CreateServerGroupRequestSlowStartConfig struct {
	// The duration of a slow start.
	//
	// Valid values: 30 to 900.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	SlowStartDuration *int32 `json:"SlowStartDuration,omitempty" xml:"SlowStartDuration,omitempty"`
	// Specifies whether to enable slow starts. Valid values:
	//
	// - true
	//
	// - false(default)
	//
	// example:
	//
	// false
	SlowStartEnabled *bool `json:"SlowStartEnabled,omitempty" xml:"SlowStartEnabled,omitempty"`
}

func (s CreateServerGroupRequestSlowStartConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestSlowStartConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestSlowStartConfig) GetSlowStartDuration() *int32 {
	return s.SlowStartDuration
}

func (s *CreateServerGroupRequestSlowStartConfig) GetSlowStartEnabled() *bool {
	return s.SlowStartEnabled
}

func (s *CreateServerGroupRequestSlowStartConfig) SetSlowStartDuration(v int32) *CreateServerGroupRequestSlowStartConfig {
	s.SlowStartDuration = &v
	return s
}

func (s *CreateServerGroupRequestSlowStartConfig) SetSlowStartEnabled(v bool) *CreateServerGroupRequestSlowStartConfig {
	s.SlowStartEnabled = &v
	return s
}

func (s *CreateServerGroupRequestSlowStartConfig) Validate() error {
	return dara.Validate(s)
}

type CreateServerGroupRequestStickySessionConfig struct {
	// The cookie that you want to configure for the server.
	//
	// The cookie must be 1 to 200 characters in length, and can contain only ASCII letters and digits. It cannot contain commas (,), semicolons (;), or space characters. It cannot start with a dollar sign ($).
	//
	// >  This parameter takes effect only when **StickySessionEnabled*	- is set to **true*	- and **StickySessionType*	- is set to **server**.
	//
	// example:
	//
	// B490B5EBF6F3CD402E515D22BCDA****
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// The maximum amount of time to wait before the session cookie expires. Unit: seconds
	//
	// Valid values: **1*	- to **86400**
	//
	// Default value: **1000**
	//
	// >  This parameter takes effect only when **StickySessionEnabled*	- is set to **true*	- and **StickySessionType*	- is set to **Insert**.
	//
	// example:
	//
	// 1000
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// Specifies whether to enable session persistence. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter takes effect when the **ServerGroupType*	- parameter is set to **Instance*	- or **Ip**.
	//
	// example:
	//
	// false
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitempty" xml:"StickySessionEnabled,omitempty"`
	// The method that is used to handle cookies. Valid values:
	//
	// 	- **Insert*	- (default value): inserts a cookie. The first time a client accesses ALB, ALB inserts the SERVERID cookie into the HTTP or HTTPS response packet. Subsequent requests from the client that carry this cookie are forwarded to the same backend server as the first request.
	//
	// 	- **Server**: rewrites a cookie. ALB rewrites the custom cookies in requests from a client. Subsequent requests from the client that carry the new cookie are forwarded to the same backend server as the first request.
	//
	// >  This parameter takes effect when the **StickySessionEnabled*	- parameter is set to **true**.
	//
	// example:
	//
	// Insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
}

func (s CreateServerGroupRequestStickySessionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestStickySessionConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestStickySessionConfig) GetCookie() *string {
	return s.Cookie
}

func (s *CreateServerGroupRequestStickySessionConfig) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *CreateServerGroupRequestStickySessionConfig) GetStickySessionEnabled() *bool {
	return s.StickySessionEnabled
}

func (s *CreateServerGroupRequestStickySessionConfig) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookie(v string) *CreateServerGroupRequestStickySessionConfig {
	s.Cookie = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetCookieTimeout(v int32) *CreateServerGroupRequestStickySessionConfig {
	s.CookieTimeout = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionEnabled(v bool) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionEnabled = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) SetStickySessionType(v string) *CreateServerGroupRequestStickySessionConfig {
	s.StickySessionType = &v
	return s
}

func (s *CreateServerGroupRequestStickySessionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateServerGroupRequestTag struct {
	// The tag key. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
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

type CreateServerGroupRequestUchConfig struct {
	// The type of the parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// QueryString
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value for consistent hashing.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerGroupRequestUchConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerGroupRequestUchConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestUchConfig) GetType() *string {
	return s.Type
}

func (s *CreateServerGroupRequestUchConfig) GetValue() *string {
	return s.Value
}

func (s *CreateServerGroupRequestUchConfig) SetType(v string) *CreateServerGroupRequestUchConfig {
	s.Type = &v
	return s
}

func (s *CreateServerGroupRequestUchConfig) SetValue(v string) *CreateServerGroupRequestUchConfig {
	s.Value = &v
	return s
}

func (s *CreateServerGroupRequestUchConfig) Validate() error {
	return dara.Validate(s)
}
