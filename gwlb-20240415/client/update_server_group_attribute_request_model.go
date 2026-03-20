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
	SetDryRun(v bool) *UpdateServerGroupAttributeRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *UpdateServerGroupAttributeRequestHealthCheckConfig) *UpdateServerGroupAttributeRequest
	GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig
	SetScheduler(v string) *UpdateServerGroupAttributeRequest
	GetScheduler() *string
	SetServerFailoverMode(v string) *UpdateServerGroupAttributeRequest
	GetServerFailoverMode() *string
	SetServerGroupId(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupId() *string
	SetServerGroupName(v string) *UpdateServerGroupAttributeRequest
	GetServerGroupName() *string
}

type UpdateServerGroupAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of connection draining.
	ConnectionDrainConfig *UpdateServerGroupAttributeRequestConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The health check configuration.
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **5TCH**: specifies consistent hashing that is based on the following factors: source IP address, destination IP address, source port, protocol, and destination port. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **3TCH**: indicates consistent hashing that is based on the following factors: source IP address, destination IP address, and protocol. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **2TCH**: specifies consistent hashing that is based on the following factors: source IP address and destination IP address. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// example:
	//
	// 5TCH
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Specifies how GWLB processes requests over existing connections when a backend server is not running as expected. Valid values:
	//
	// 	- **NoRebalance**: GWLB continues to forward requests over existing connections to the unavailable backend server.
	//
	// 	- **Rebalance**: GWLB forwards requests over existing connections to the remaining healthy backend servers.
	//
	// example:
	//
	// NoRebalance
	ServerFailoverMode *string `json:"ServerFailoverMode,omitempty" xml:"ServerFailoverMode,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server group name.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testServerGroupName
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
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

func (s *UpdateServerGroupAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateServerGroupAttributeRequest) GetHealthCheckConfig() *UpdateServerGroupAttributeRequestHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *UpdateServerGroupAttributeRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *UpdateServerGroupAttributeRequest) GetServerFailoverMode() *string {
	return s.ServerFailoverMode
}

func (s *UpdateServerGroupAttributeRequest) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *UpdateServerGroupAttributeRequest) GetServerGroupName() *string {
	return s.ServerGroupName
}

func (s *UpdateServerGroupAttributeRequest) SetClientToken(v string) *UpdateServerGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServerGroupAttributeRequest) SetConnectionDrainConfig(v *UpdateServerGroupAttributeRequestConnectionDrainConfig) *UpdateServerGroupAttributeRequest {
	s.ConnectionDrainConfig = v
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

func (s *UpdateServerGroupAttributeRequest) SetServerFailoverMode(v string) *UpdateServerGroupAttributeRequest {
	s.ServerFailoverMode = &v
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
	return nil
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
	// Unit: seconds
	//
	// Valid values: 1 to 3600.
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
	// The backend server port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period for a health check response.
	//
	// Unit: seconds.
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks. Valid values:
	//
	// 	- **$SERVER_IP**: the internal IP address of a backend server.
	//
	// 	- **domain**: a domain name. The domain name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), and periods (.).
	//
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckExp     *string `json:"HealthCheckExp,omitempty" xml:"HealthCheckExp,omitempty"`
	// The HTTP status codes that the system returns for health checks.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval at which health checks are performed.
	//
	// Unit: seconds.
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The URL used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: ` - / . % ? # &  `It must start with a forward slash (/).
	//
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **HTTP**: HTTP health checks simulate a process that uses a web browser to access resources by sending GET requests to an instance. These requests are used to check whether the instance is healthy.
	//
	// example:
	//
	// TCP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	HealthCheckReq      *string `json:"HealthCheckReq,omitempty" xml:"HealthCheckReq,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health check status of the backend server changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthCheckReq(v string) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthCheckReq = &v
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
