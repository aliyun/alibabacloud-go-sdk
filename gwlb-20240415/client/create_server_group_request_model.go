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
	SetDryRun(v bool) *CreateServerGroupRequest
	GetDryRun() *bool
	SetHealthCheckConfig(v *CreateServerGroupRequestHealthCheckConfig) *CreateServerGroupRequest
	GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig
	SetProtocol(v string) *CreateServerGroupRequest
	GetProtocol() *string
	SetResourceGroupId(v string) *CreateServerGroupRequest
	GetResourceGroupId() *string
	SetScheduler(v string) *CreateServerGroupRequest
	GetScheduler() *string
	SetServerFailoverMode(v string) *CreateServerGroupRequest
	GetServerFailoverMode() *string
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
	ConnectionDrainConfig *CreateServerGroupRequestConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// False
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The health check configurations.
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The backend protocol. Valid value:
	//
	// 	- **GENEVE*	- (default)
	//
	// example:
	//
	// GENEVE
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **5TCH*	- (default): specifies consistent hashing that is based on the following factors: source IP address, destination IP address, source port, protocol, and destination port. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **3TCH**: specifies consistent hashing that is based on the following factors: source IP address, destination IP address, and protocol. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **2TCH**: specifies consistent hashing that is based on the following factors: source IP address and destination IP address. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// example:
	//
	// 5TCH
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// Specifies how GWLB processes requests over existing connections when a backend server is not running as expected. Valid values:
	//
	// 	- **NoRebalance*	- (default): GWLB continues to forward requests over existing connections to the unavailable backend server.
	//
	// 	- **Rebalance**: GWLB forwards requests over existing connections to the remaining healthy backend servers.
	//
	// example:
	//
	// NoRebalance
	ServerFailoverMode *string `json:"ServerFailoverMode,omitempty" xml:"ServerFailoverMode,omitempty"`
	// The server group name.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testServerGroupName
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The type of the server group. Valid values:
	//
	// 	- **Instance*	- (default): allows you to specify resources of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers by specifying their IP addresses.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The tag keys.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateServerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// > If **ServerGroupType*	- is set to **Instance**, only servers in the specified VPC can be added to the server group.
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

func (s *CreateServerGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServerGroupRequest) GetConnectionDrainConfig() *CreateServerGroupRequestConnectionDrainConfig {
	return s.ConnectionDrainConfig
}

func (s *CreateServerGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServerGroupRequest) GetHealthCheckConfig() *CreateServerGroupRequestHealthCheckConfig {
	return s.HealthCheckConfig
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

func (s *CreateServerGroupRequest) GetServerFailoverMode() *string {
	return s.ServerFailoverMode
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

func (s *CreateServerGroupRequest) SetClientToken(v string) *CreateServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServerGroupRequest) SetConnectionDrainConfig(v *CreateServerGroupRequestConnectionDrainConfig) *CreateServerGroupRequest {
	s.ConnectionDrainConfig = v
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

func (s *CreateServerGroupRequest) SetServerFailoverMode(v string) *CreateServerGroupRequest {
	s.ServerFailoverMode = &v
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

type CreateServerGroupRequestConnectionDrainConfig struct {
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
	// Valid values: **1*	- to **3600**.
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
	// The backend server port used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// Default value: **80**.
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
	// Default value: **5**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name used for health checks. Valid values:
	//
	// 	- **$SERVER_IP*	- (default): the private IP address of a backend server.
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
	// 	- **true*	- (default)
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
	// Default value: **10**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The path used for health checks.
	//
	// It must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL can also contain the following extended characters: _ ; ~ ! ( ) \\	- [ ] @ $ ^ : \\" , + =
	//
	// It must start with a forward slash (/).
	//
	// >  This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol used for health checks. Valid values:
	//
	// 	- **TCP*	- (default): GWLB performs TCP health checks by sending SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	//
	// 	- **HTTP**: GWLB performs HTTP health checks to check whether backend servers are healthy by sending GET requests which simulate access from browsers.
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
	// Default value: **2**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health check status of the backend server changes from **success*	- to **fail**.
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

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthCheckReq() *string {
	return s.HealthCheckReq
}

func (s *CreateServerGroupRequestHealthCheckConfig) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckInterval(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckInterval = &v
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthCheckReq(v string) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthCheckReq = &v
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

type CreateServerGroupRequestTag struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
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
