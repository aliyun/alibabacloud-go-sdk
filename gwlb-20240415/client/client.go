// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddServersToServerGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers that you want to add.
	//
	// > You can add at most 200 backend servers in each call.
	//
	// This parameter is required.
	Servers []*AddServersToServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s AddServersToServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequest) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequest) SetClientToken(v string) *AddServersToServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetDryRun(v bool) *AddServersToServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServerGroupId(v string) *AddServersToServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *AddServersToServerGroupRequest) SetServers(v []*AddServersToServerGroupRequestServers) *AddServersToServerGroupRequest {
	s.Servers = v
	return s
}

type AddServersToServerGroupRequestServers struct {
	// The backend server port. Valid values:
	//
	// 	- **6081**
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The backend server ID.
	//
	// 	- If the server group is of the **Instance*	- type, set this parameter to the IDs of servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set ServerId to IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.XX.XX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: elastic network interface (ENI)
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s AddServersToServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupRequestServers) SetPort(v int32) *AddServersToServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerId(v string) *AddServersToServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerIp(v string) *AddServersToServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *AddServersToServerGroupRequestServers) SetServerType(v string) *AddServersToServerGroupRequestServers {
	s.ServerType = &v
	return s
}

type AddServersToServerGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddServersToServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddServersToServerGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddServersToServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddServersToServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddServersToServerGroupResponse) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponse) SetHeaders(v map[string]*string) *AddServersToServerGroupResponse {
	s.Headers = v
	return s
}

func (s *AddServersToServerGroupResponse) SetStatusCode(v int32) *AddServersToServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServersToServerGroupResponse) SetBody(v *AddServersToServerGroupResponseBody) *AddServersToServerGroupResponse {
	s.Body = v
	return s
}

type CreateListenerRequest struct {
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
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the listener.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener-description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-te609d6696632f7*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-ckh01px70dszof****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag []*CreateListenerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateListenerRequest) SetClientToken(v string) *CreateListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateListenerRequest) SetDryRun(v bool) *CreateListenerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateListenerRequest) SetListenerDescription(v string) *CreateListenerRequest {
	s.ListenerDescription = &v
	return s
}

func (s *CreateListenerRequest) SetLoadBalancerId(v string) *CreateListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateListenerRequest) SetServerGroupId(v string) *CreateListenerRequest {
	s.ServerGroupId = &v
	return s
}

func (s *CreateListenerRequest) SetTag(v []*CreateListenerRequestTag) *CreateListenerRequest {
	s.Tag = v
	return s
}

type CreateListenerRequestTag struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateListenerRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateListenerRequestTag) SetKey(v string) *CreateListenerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateListenerRequestTag) SetValue(v string) *CreateListenerRequestTag {
	s.Value = &v
	return s
}

type CreateListenerResponseBody struct {
	// The listener ID.
	//
	// example:
	//
	// lsn-wi3c0v30ivysrg****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A045E652-D298-5E70-A978-7247135336FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateListenerResponseBody) SetListenerId(v string) *CreateListenerResponseBody {
	s.ListenerId = &v
	return s
}

func (s *CreateListenerResponseBody) SetRequestId(v string) *CreateListenerResponseBody {
	s.RequestId = &v
	return s
}

type CreateListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateListenerResponse) SetHeaders(v map[string]*string) *CreateListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateListenerResponse) SetStatusCode(v int32) *CreateListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateListenerResponse) SetBody(v *CreateListenerResponseBody) *CreateListenerResponse {
	s.Body = v
	return s
}

type CreateLoadBalancerRequest struct {
	// The IP version. Valid values:
	//
	// 	- **Ipv4*	- (default): IPv4
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// testGwlbName
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwbufq6q3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys. You can specify at most 20 tags in each call.
	Tag []*CreateLoadBalancerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6qcgpv22ttrnnjh****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. You must specify at least one zone. You can specify at most 20 zones. If the region supports two or more zones, specify at least two zones.
	//
	// This parameter is required.
	ZoneMappings []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

type CreateLoadBalancerRequestTag struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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

func (s CreateLoadBalancerRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestTag) SetKey(v string) *CreateLoadBalancerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) SetValue(v string) *CreateLoadBalancerRequestTag {
	s.Value = &v
	return s
}

type CreateLoadBalancerRequestZoneMappings struct {
	// The ID of the vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of a GWLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2f0eb020****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the DescribeZones operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type CreateLoadBalancerResponseBody struct {
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 00B19438-66BB-58C3-8C2F-DA5B6F95CBDA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerResponse) SetStatusCode(v int32) *CreateLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerResponse) SetBody(v *CreateLoadBalancerResponseBody) *CreateLoadBalancerResponse {
	s.Body = v
	return s
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
	// The configurations of the health check feature.
	HealthCheckConfig *CreateServerGroupRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The backend protocol. Valid values:
	//
	// 	- **GENEVE**(default)
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
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	ServerFailoverMode *string `json:"ServerFailoverMode,omitempty" xml:"ServerFailoverMode,omitempty"`
	// The server group name.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testServerGroupName
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The type of server group. Valid values:
	//
	// 	- **Instance*	- (default): allows you to specify servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers of by specifying IP addresses.
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
	return tea.Prettify(s)
}

func (s CreateServerGroupRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *CreateServerGroupRequestConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *CreateServerGroupRequestConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *CreateServerGroupRequestConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

type CreateServerGroupRequestHealthCheckConfig struct {
	// The backend server port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// Default value: **80**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check response.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// Default value: **5**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that you want to use for health checks. Valid values:
	//
	// 	- **$SERVER_IP*	- (default): the private IP address of a backend server.
	//
	// 	- **domain**: a domain name. The domain name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The HTTP status codes that the system returns for health checks.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval at which health checks are performed.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **50**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), percent signs (%), question marks (?), number signs (#), and ampersands (&). The URL can also contain the following extended characters: _ ; ~ ! ( ) \\	- [ ] @ $ ^ : \\" , + =
	//
	// The URL must start with a forward slash (/).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **TCP*	- (default): GWLB performs TCP health checks by sending SYN packets to a backend server to check whether the port of the backend server is available to receive requests.
	//
	// 	- **HTTP**: GWLB performs HTTP health checks to check whether backend servers are healthy by sending HEAD or GET requests which simulate access from browsers.
	//
	// example:
	//
	// TCP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// Default value: **2**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status changes from **success*	- to **fail**.
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
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestHealthCheckConfig) GoString() string {
	return s.String()
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

func (s *CreateServerGroupRequestHealthCheckConfig) SetHealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateServerGroupRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *CreateServerGroupRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
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
	return tea.Prettify(s)
}

func (s CreateServerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServerGroupRequestTag) SetKey(v string) *CreateServerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServerGroupRequestTag) SetValue(v string) *CreateServerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateServerGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s CreateServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponseBody) SetRequestId(v string) *CreateServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerGroupResponseBody) SetServerGroupId(v string) *CreateServerGroupResponseBody {
	s.ServerGroupId = &v
	return s
}

type CreateServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateServerGroupResponse) SetHeaders(v map[string]*string) *CreateServerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateServerGroupResponse) SetStatusCode(v int32) *CreateServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServerGroupResponse) SetBody(v *CreateServerGroupResponseBody) *CreateServerGroupResponse {
	s.Body = v
	return s
}

type DeleteListenerRequest struct {
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
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-brx2y3hqdincizg***
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s *DeleteListenerRequest) SetClientToken(v string) *DeleteListenerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteListenerRequest) SetDryRun(v bool) *DeleteListenerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteListenerRequest) SetListenerId(v string) *DeleteListenerRequest {
	s.ListenerId = &v
	return s
}

type DeleteListenerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5C6E3548-086F-5FF6-A2B3-B1871B3AB488
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteListenerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponseBody) SetRequestId(v string) *DeleteListenerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteListenerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteListenerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteListenerResponse) SetHeaders(v map[string]*string) *DeleteListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteListenerResponse) SetStatusCode(v int32) *DeleteListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteListenerResponse) SetBody(v *DeleteListenerResponseBody) *DeleteListenerResponse {
	s.Body = v
	return s
}

type DeleteLoadBalancerRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) SetClientToken(v string) *DeleteLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetDryRun(v bool) *DeleteLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

type DeleteLoadBalancerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 37907828-01AB-5AC3-9DDE-25D419091D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoadBalancerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponseBody) SetRequestId(v string) *DeleteLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoadBalancerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerResponse) SetStatusCode(v int32) *DeleteLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerResponse) SetBody(v *DeleteLoadBalancerResponseBody) *DeleteLoadBalancerResponse {
	s.Body = v
	return s
}

type DeleteServerGroupRequest struct {
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
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s DeleteServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupRequest) SetClientToken(v string) *DeleteServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerGroupRequest) SetDryRun(v bool) *DeleteServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServerGroupRequest) SetServerGroupId(v string) *DeleteServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

type DeleteServerGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponseBody) SetRequestId(v string) *DeleteServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServerGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerGroupResponse) SetHeaders(v map[string]*string) *DeleteServerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerGroupResponse) SetStatusCode(v int32) *DeleteServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerGroupResponse) SetBody(v *DeleteServerGroupResponseBody) *DeleteServerGroupResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The supported language. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// 	- **ja**: Japanese
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// A list of regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 162FCE8D-CEEC-5083-90BF-B45D8C4F81FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region endpoint.
	//
	// example:
	//
	// gwlb.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// The supported language. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// 	- **ja**: Japanese
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeZonesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C605F7E-D0F6-54E2-B004-F9B132F0D8B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of zones.
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// The zone name.
	//
	// example:
	//
	// Hangzhou Zone G
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type GetListenerAttributeRequest struct {
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-brx2y3hqdinciz****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
}

func (s GetListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeRequest) SetListenerId(v string) *GetListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

type GetListenerAttributeResponseBody struct {
	// The listener description.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener_description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsn-3kbj3587mqhm3p****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The listener status. Valid values:
	//
	// 	- **Provisioning**: The listener is being created.
	//
	// 	- **Running**: The listener is running.
	//
	// 	- **Configuring**: The listener is being configured.
	//
	// 	- **Deleting**: The listener is being deleted.
	//
	// example:
	//
	// Provisioning
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-te609d6696632f76****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The region ID of the GWLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75CC3312-7757-5EE1-90D8-49CEA66608AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-sp8d2r6y7t0xtl****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags.
	Tags []*GetListenerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBody) SetListenerDescription(v string) *GetListenerAttributeResponseBody {
	s.ListenerDescription = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerId(v string) *GetListenerAttributeResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetListenerStatus(v string) *GetListenerAttributeResponseBody {
	s.ListenerStatus = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetLoadBalancerId(v string) *GetListenerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRegionId(v string) *GetListenerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetRequestId(v string) *GetListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetServerGroupId(v string) *GetListenerAttributeResponseBody {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerAttributeResponseBody) SetTags(v []*GetListenerAttributeResponseBodyTags) *GetListenerAttributeResponseBody {
	s.Tags = v
	return s
}

type GetListenerAttributeResponseBodyTags struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs: `or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetListenerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponseBodyTags) SetKey(v string) *GetListenerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetListenerAttributeResponseBodyTags) SetValue(v string) *GetListenerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type GetListenerAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetListenerAttributeResponse) SetHeaders(v map[string]*string) *GetListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetListenerAttributeResponse) SetStatusCode(v int32) *GetListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerAttributeResponse) SetBody(v *GetListenerAttributeResponseBody) *GetListenerAttributeResponse {
	s.Body = v
	return s
}

type GetListenerHealthStatusRequest struct {
	// The filter conditions. You can specify at most 20 filter conditions.
	Filter []*GetListenerHealthStatusRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-7sixpvm5fc3v0b****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 4f1d7cc9f51e18904e8a063603a6b0c3d03bc69f78734254e0b5e8707e68****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s GetListenerHealthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequest) SetFilter(v []*GetListenerHealthStatusRequestFilter) *GetListenerHealthStatusRequest {
	s.Filter = v
	return s
}

func (s *GetListenerHealthStatusRequest) SetListenerId(v string) *GetListenerHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetMaxResults(v int32) *GetListenerHealthStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetNextToken(v string) *GetListenerHealthStatusRequest {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetSkip(v int32) *GetListenerHealthStatusRequest {
	s.Skip = &v
	return s
}

type GetListenerHealthStatusRequestFilter struct {
	// The filter condition name. You can filter by one or more filter condition names. The URL must meet the following requirements:
	//
	// 	- **Status**: the health status.
	//
	// 	- **ReasonCode**: the cause of an unhealthy server.
	//
	// 	- **ServerId**: the ID of the backend server.
	//
	// 	- **ServerIp**: the IP address of the backend server.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter condition values. You can specify at most 20 condition values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusRequestFilter) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequestFilter) SetName(v string) *GetListenerHealthStatusRequestFilter {
	s.Name = &v
	return s
}

func (s *GetListenerHealthStatusRequestFilter) SetValues(v []*string) *GetListenerHealthStatusRequestFilter {
	s.Values = v
	return s
}

type GetListenerHealthStatusResponseBody struct {
	// The health check status of the server groups that are associated with the listener.
	ListenerHealthStatus []*GetListenerHealthStatusResponseBodyListenerHealthStatus `json:"ListenerHealthStatus,omitempty" xml:"ListenerHealthStatus,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// U12WEI6Ro2ol3wA54rBNSwdC5+lYy6q5SjIQEvc1wz5mjZxV+YjsHRdXV8XauY1BpOQIvwX63E0en54H3D****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED4F222-36A0-5470-8A9A-AAB4E96BAC1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetListenerHealthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBody) SetListenerHealthStatus(v []*GetListenerHealthStatusResponseBodyListenerHealthStatus) *GetListenerHealthStatusResponseBody {
	s.ListenerHealthStatus = v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetMaxResults(v int32) *GetListenerHealthStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetNextToken(v string) *GetListenerHealthStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetRequestId(v string) *GetListenerHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBody) SetTotalCount(v int32) *GetListenerHealthStatusResponseBody {
	s.TotalCount = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatus struct {
	// The listener ID.
	//
	// example:
	//
	// lsn-sg8aha6pzjavvo****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The information about the server groups.
	ServerGroupInfos []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos `json:"ServerGroupInfos,omitempty" xml:"ServerGroupInfos,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatus) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetListenerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatus) SetServerGroupInfos(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) *GetListenerHealthStatusResponseBodyListenerHealthStatus {
	s.ServerGroupInfos = v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos struct {
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-0vdsbyszro3nr6****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers.
	Servers []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetHealthCheckEnabled(v bool) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.HealthCheckEnabled = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServerGroupId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.ServerGroupId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos) SetServers(v []*GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfos {
	s.Servers = v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers struct {
	// The backend port.
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The reason why **Status*	- indicates an unhealthy status.
	Reason *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason `json:"Reason,omitempty" xml:"Reason,omitempty" type:"Struct"`
	// The backend server ID.
	//
	// example:
	//
	// i-2ze4rnh8yj9kif3z****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the server.
	//
	// example:
	//
	// 192.168.0.XXX
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The health status of the backend server. Valid values:
	//
	// 	- **Initial**: Health checks are configured for the GWLB instance, but no data is found.
	//
	// 	- **Unhealthy**: The backend server consecutively fails health checks.
	//
	// 	- **Unused**: The backend server is not in use.
	//
	// 	- **Unavailable**: Health checks are disabled.
	//
	// 	- **Healthy**: The backend server is healthy.
	//
	// example:
	//
	// Healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetPort(v int32) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Port = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetReason(v *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Reason = v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetServerId(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.ServerId = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetServerIp(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.ServerIp = &v
	return s
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers) SetStatus(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServers {
	s.Status = &v
	return s
}

type GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason struct {
	// The reason why **Status*	- indicates an unhealthy status. Valid values:
	//
	// 	- **CONNECT_TIMEOUT**: The GWLB instance failed to connect to the backend server within the specified period of time.
	//
	// 	- **CONNECT_FAILED**: The GWLB instance failed to connect to the backend server.
	//
	// 	- **RECV_RESPONSE_TIMEOUT**: The GWLB instance failed to receive a response from the backend server within the specified period of time.
	//
	// 	- **CONNECT_INTERRUPT**: The connection between the health check and the backend server was interrupted.
	//
	// 	- **HTTP_CODE_NOT_MATCH**: The HTTP status code from the backend server is not the expected one.
	//
	// 	- **HTTP_INVALID_HEADER**: The format of the response from the backend servers is invalid.
	//
	// example:
	//
	// CONNECT_TIMEOUT
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason) SetReasonCode(v string) *GetListenerHealthStatusResponseBodyListenerHealthStatusServerGroupInfosServersReason {
	s.ReasonCode = &v
	return s
}

type GetListenerHealthStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListenerHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListenerHealthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListenerHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusResponse) SetHeaders(v map[string]*string) *GetListenerHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetListenerHealthStatusResponse) SetStatusCode(v int32) *GetListenerHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListenerHealthStatusResponse) SetBody(v *GetListenerHealthStatusResponseBody) *GetListenerHealthStatusResponse {
	s.Body = v
	return s
}

type GetLoadBalancerAttributeRequest struct {
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s GetLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *GetLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

type GetLoadBalancerAttributeResponseBody struct {
	// The protocol version. Valid values:
	//
	// 	- **Ipv4**: IPv4.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-08T10:12:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The business status of the GWLB instance. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// gwlb
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The GWLB instance status. Valid values:
	//
	// 	- **Active**: The GWLB instance is running.
	//
	// 	- **Inactive**: The GWLB instance is disabled. Listeners of GWLB instances in the Inactive state do not forward traffic.
	//
	// 	- **Provisioning**: The GWLB instance is being created.
	//
	// 	- **Configuring**: The GWLB instance is being modified.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6DC5DDC-9560-59BF-80FA-ED1E5CB417DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmx7pmxcy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags        []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TrafficMode *string                                     `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-k1aajsbwbaq4todet****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. You must specify at least one zone. You can specify at most 20 zones. If the region supports two or more zones, specify at least two zones.
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTrafficMode(v string) *GetLoadBalancerAttributeResponseBody {
	s.TrafficMode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

type GetLoadBalancerAttributeResponseBodyTags struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// The GWLB instance addresses.
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of a GWLB instance.
	//
	// example:
	//
	// vsw-uf6v8l7d2f1k53xrl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// The ID of the elastic network interface (ENI) used by the GWLB instance.
	//
	// example:
	//
	// eni-bp1iahwz3rzgvltz****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IPv4 address.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEniId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIpv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIpv4Address = &v
	return s
}

type GetLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *GetLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetStatusCode(v int32) *GetLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponse) SetBody(v *GetLoadBalancerAttributeResponseBody) *GetLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type ListListenersRequest struct {
	// The listener IDs. You can specify at most 20 listener IDs.
	ListenerIds []*string `json:"ListenerIds,omitempty" xml:"ListenerIds,omitempty" type:"Repeated"`
	// The GWLB instance IDs. You can specify at most 20 instance IDs.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// d209f4e63ec942c967c50c888a13****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag []*ListListenersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListListenersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequest) GoString() string {
	return s.String()
}

func (s *ListListenersRequest) SetListenerIds(v []*string) *ListListenersRequest {
	s.ListenerIds = v
	return s
}

func (s *ListListenersRequest) SetLoadBalancerIds(v []*string) *ListListenersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListListenersRequest) SetMaxResults(v int32) *ListListenersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenersRequest) SetNextToken(v string) *ListListenersRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenersRequest) SetSkip(v int32) *ListListenersRequest {
	s.Skip = &v
	return s
}

func (s *ListListenersRequest) SetTag(v []*ListListenersRequestTag) *ListListenersRequest {
	s.Tag = v
	return s
}

type ListListenersRequestTag struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListListenersRequestTag) GoString() string {
	return s.String()
}

func (s *ListListenersRequestTag) SetKey(v string) *ListListenersRequestTag {
	s.Key = &v
	return s
}

func (s *ListListenersRequestTag) SetValue(v string) *ListListenersRequestTag {
	s.Value = &v
	return s
}

type ListListenersResponseBody struct {
	// The GWLB listeners.
	Listeners []*ListListenersResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 5c281c0a0d6bfb6355ed088c2108aca8e0b5e8707e68****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7DBFC67C-A272-5952-8287-6C3EBE4E04D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBody) SetListeners(v []*ListListenersResponseBodyListeners) *ListListenersResponseBody {
	s.Listeners = v
	return s
}

func (s *ListListenersResponseBody) SetMaxResults(v int32) *ListListenersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenersResponseBody) SetNextToken(v string) *ListListenersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenersResponseBody) SetRequestId(v string) *ListListenersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenersResponseBody) SetTotalCount(v int32) *ListListenersResponseBody {
	s.TotalCount = &v
	return s
}

type ListListenersResponseBodyListeners struct {
	// The description of the listener.
	//
	// example:
	//
	// listener-description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsn-vu7folhh5ntm8u****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The status of the listener. Valid values:
	//
	// 	- **Provisioning**: The listener is being created.
	//
	// 	- **Running**: The listener is running.
	//
	// 	- **Configuring**: The listener is being configured.
	//
	// 	- **Deleting**: The listener is being deleted.
	//
	// example:
	//
	// Running
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-uf6hbeh795xlqln7g****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-5yapcb422i51ru****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The tags.
	Tags []*ListListenersResponseBodyListenersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListListenersResponseBodyListeners) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListeners) SetListenerDescription(v string) *ListListenersResponseBodyListeners {
	s.ListenerDescription = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerId(v string) *ListListenersResponseBodyListeners {
	s.ListenerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetListenerStatus(v string) *ListListenersResponseBodyListeners {
	s.ListenerStatus = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetLoadBalancerId(v string) *ListListenersResponseBodyListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetServerGroupId(v string) *ListListenersResponseBodyListeners {
	s.ServerGroupId = &v
	return s
}

func (s *ListListenersResponseBodyListeners) SetTags(v []*ListListenersResponseBodyListenersTags) *ListListenersResponseBodyListeners {
	s.Tags = v
	return s
}

type ListListenersResponseBodyListenersTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListListenersResponseBodyListenersTags) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponseBodyListenersTags) GoString() string {
	return s.String()
}

func (s *ListListenersResponseBodyListenersTags) SetKey(v string) *ListListenersResponseBodyListenersTags {
	s.Key = &v
	return s
}

func (s *ListListenersResponseBodyListenersTags) SetValue(v string) *ListListenersResponseBodyListenersTags {
	s.Value = &v
	return s
}

type ListListenersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListListenersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListListenersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListListenersResponse) GoString() string {
	return s.String()
}

func (s *ListListenersResponse) SetHeaders(v map[string]*string) *ListListenersResponse {
	s.Headers = v
	return s
}

func (s *ListListenersResponse) SetStatusCode(v int32) *ListListenersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListListenersResponse) SetBody(v *ListListenersResponseBody) *ListListenersResponse {
	s.Body = v
	return s
}

type ListLoadBalancersRequest struct {
	// The IP version. Valid values:
	//
	// 	- **Ipv4**: IPv4
	//
	// Enumeration values:
	//
	// 	- IPv4: IPv4
	//
	// 	- DualStack: DualStack
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The business status of the GWLB instance. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The GWLB instance IDs. You can query at most 20 GWLB instances in each call.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The GWLB instance names. You can specify at most 20 GWLB instance names in each call.
	LoadBalancerNames []*string `json:"LoadBalancerNames,omitempty" xml:"LoadBalancerNames,omitempty" type:"Repeated"`
	// The GWLB instance status. Valid values:
	//
	// 	- **Active**: The GWLB instance is running.
	//
	// 	- **Inactive**: The GWLB instance is disabled. Listeners of GWLB instances in the Inactive state do not forward traffic.
	//
	// 	- **Provisioning**: The GWLB instance is being created.
	//
	// 	- **Configuring**: The GWLB instance is being modified.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// WyJyb290IiwibiIsIm4iLDEsMCwxNjg1MDY1NTgyNzYwLCI2NDcwMGY2ZTc2Zjc0MWFiZGEyZjQyNzc4ZDk2MmJkOTk3ZGZmM2Nm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2htf5qsyrn****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag         []*ListLoadBalancersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TrafficMode *string                        `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The virtual private cloud (VPC) IDs. You can query at most 20 IDs in each call.
	VpcIds []*string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty" type:"Repeated"`
	// The zone IDs. You can query at most 20 zone IDs in each call.
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s ListLoadBalancersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequest) SetAddressIpVersion(v string) *ListLoadBalancersRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerIds(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerNames(v []*string) *ListLoadBalancersRequest {
	s.LoadBalancerNames = v
	return s
}

func (s *ListLoadBalancersRequest) SetLoadBalancerStatus(v string) *ListLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersRequest) SetMaxResults(v int32) *ListLoadBalancersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersRequest) SetNextToken(v string) *ListLoadBalancersRequest {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersRequest) SetResourceGroupId(v string) *ListLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersRequest) SetSkip(v int32) *ListLoadBalancersRequest {
	s.Skip = &v
	return s
}

func (s *ListLoadBalancersRequest) SetTag(v []*ListLoadBalancersRequestTag) *ListLoadBalancersRequest {
	s.Tag = v
	return s
}

func (s *ListLoadBalancersRequest) SetTrafficMode(v string) *ListLoadBalancersRequest {
	s.TrafficMode = &v
	return s
}

func (s *ListLoadBalancersRequest) SetVpcIds(v []*string) *ListLoadBalancersRequest {
	s.VpcIds = v
	return s
}

func (s *ListLoadBalancersRequest) SetZoneIds(v []*string) *ListLoadBalancersRequest {
	s.ZoneIds = v
	return s
}

type ListLoadBalancersRequestTag struct {
	// The tag key The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersRequestTag) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersRequestTag) SetKey(v string) *ListLoadBalancersRequestTag {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersRequestTag) SetValue(v string) *ListLoadBalancersRequestTag {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBody struct {
	// The GWLB instances.
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAIldD2UAAAAACjMDLgAAADFTNzMyZDMwMzAzMDY5NzQzNDM0NmI3NzM2NjUzNzc4NzM2YTc0NjYzOTYz****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 378A80E9-4262-5D8E-8D62-0969E52D7358
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMaxResults(v int32) *ListLoadBalancersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetNextToken(v string) *ListLoadBalancersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// The IP version. Valid values:
	//
	// 	- **IPv4**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-05 18:24:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The business status of the GWLB instance. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// testGwlbName
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The GWLB instance status. Valid values:
	//
	// 	- **Active**: The GWLB instance is running.
	//
	// 	- **Inactive**: The GWLB instance is disabled. Listeners of GWLB instances in the Inactive state do not forward traffic.
	//
	// 	- **Provisioning**: The GWLB instance is being created.
	//
	// 	- **Configuring**: The GWLB instance is being modified.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek26jasguy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListLoadBalancersResponseBodyLoadBalancersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6eg0vndlsa84n7r****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. You must specify at least one zone. You can specify at most 20 zones. If the region supports two or more zones, specify at least two zones.
	ZoneMappings []*ListLoadBalancersResponseBodyLoadBalancersZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressIpVersion(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBusinessStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTags(v []*ListLoadBalancersResponseBodyLoadBalancersTags) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Tags = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.VpcId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetZoneMappings(v []*ListLoadBalancersResponseBodyLoadBalancersZoneMappings) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ZoneMappings = v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersTags struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetKey(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetValue(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Value = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersZoneMappings struct {
	// The GWLB instance addresses.
	LoadBalancerAddresses []*ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The ID of the vSwitch in the zone. By default, each zone contains one vSwitch and one subnet.
	//
	// example:
	//
	// vsw-2zemule5dz7okwqfv****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the DescribeZones operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappings) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetLoadBalancerAddresses(v []*ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetVSwitchId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappings) SetZoneId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappings {
	s.ZoneId = &v
	return s
}

type ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses struct {
	// The ID of the elastic network interface (ENI) used by the GWLB instance.
	//
	// example:
	//
	// eni-bp17qv9zbzyqy629****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IPv4 address.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetEniId(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses) SetPrivateIpv4Address(v string) *ListLoadBalancersResponseBodyLoadBalancersZoneMappingsLoadBalancerAddresses {
	s.PrivateIpv4Address = &v
	return s
}

type ListLoadBalancersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLoadBalancersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponse) SetHeaders(v map[string]*string) *ListLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancersResponse) SetStatusCode(v int32) *ListLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancersResponse) SetBody(v *ListLoadBalancersResponseBody) *ListLoadBalancersResponse {
	s.Body = v
	return s
}

type ListServerGroupServersRequest struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server IDs.
	//
	// You can specify at most 200 servers in each call.
	ServerIds []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	// The server IP addresses.
	//
	// You can specify at most 200 servers in each call.
	ServerIps []*string `json:"ServerIps,omitempty" xml:"ServerIps,omitempty" type:"Repeated"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListServerGroupServersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersRequest) SetMaxResults(v int32) *ListServerGroupServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersRequest) SetNextToken(v string) *ListServerGroupServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerGroupId(v string) *ListServerGroupServersRequest {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIds(v []*string) *ListServerGroupServersRequest {
	s.ServerIds = v
	return s
}

func (s *ListServerGroupServersRequest) SetServerIps(v []*string) *ListServerGroupServersRequest {
	s.ServerIps = v
	return s
}

func (s *ListServerGroupServersRequest) SetSkip(v int32) *ListServerGroupServersRequest {
	s.Skip = &v
	return s
}

type ListServerGroupServersResponseBody struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The backend servers.
	Servers []*ListServerGroupServersResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBody) SetMaxResults(v int32) *ListServerGroupServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetNextToken(v string) *ListServerGroupServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetRequestId(v string) *ListServerGroupServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupServersResponseBody) SetServers(v []*ListServerGroupServersResponseBodyServers) *ListServerGroupServersResponseBody {
	s.Servers = v
	return s
}

func (s *ListServerGroupServersResponseBody) SetTotalCount(v int32) *ListServerGroupServersResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupServersResponseBodyServers struct {
	// The backend server port. Valid values:
	//
	// 	- **6081**
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend server ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.xxx.xxx
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: elastic network interface (ENI)
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// Indicates the status of the backend server. Valid values:
	//
	// 	- **Adding**: The backend server is being added.
	//
	// 	- **Available**: The backend server is available.
	//
	// 	- **Draining**: The backend server is in connection draining.
	//
	// 	- **Removing**: The backend server is being removed.
	//
	// 	- **Replacing**: The backend server is being replaced.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServerGroupServersResponseBodyServers) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponseBodyServers) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponseBodyServers) SetPort(v int32) *ListServerGroupServersResponseBodyServers {
	s.Port = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerGroupId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerId(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerIp(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerIp = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetServerType(v string) *ListServerGroupServersResponseBodyServers {
	s.ServerType = &v
	return s
}

func (s *ListServerGroupServersResponseBodyServers) SetStatus(v string) *ListServerGroupServersResponseBodyServers {
	s.Status = &v
	return s
}

type ListServerGroupServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerGroupServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerGroupServersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupServersResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupServersResponse) SetHeaders(v map[string]*string) *ListServerGroupServersResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupServersResponse) SetStatusCode(v int32) *ListServerGroupServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupServersResponse) SetBody(v *ListServerGroupServersResponseBody) *ListServerGroupServersResponse {
	s.Body = v
	return s
}

type ListServerGroupsRequest struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The server group IDs.
	//
	// You can specify at most 20 server group IDs in each call.
	ServerGroupIds []*string `json:"ServerGroupIds,omitempty" xml:"ServerGroupIds,omitempty" type:"Repeated"`
	// The server group names.
	//
	// You can specify at most 20 server group names in each call.
	ServerGroupNames []*string `json:"ServerGroupNames,omitempty" xml:"ServerGroupNames,omitempty" type:"Repeated"`
	// The server group type. Valid values:
	//
	// 	- **Instance**: allows you to specify servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers of by specifying IP addresses.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The tag keys.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListServerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequest) SetMaxResults(v int32) *ListServerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsRequest) SetNextToken(v string) *ListServerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsRequest) SetResourceGroupId(v string) *ListServerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupIds(v []*string) *ListServerGroupsRequest {
	s.ServerGroupIds = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupNames(v []*string) *ListServerGroupsRequest {
	s.ServerGroupNames = v
	return s
}

func (s *ListServerGroupsRequest) SetServerGroupType(v string) *ListServerGroupsRequest {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsRequest) SetSkip(v int32) *ListServerGroupsRequest {
	s.Skip = &v
	return s
}

func (s *ListServerGroupsRequest) SetTag(v []*ListServerGroupsRequestTag) *ListServerGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListServerGroupsRequest) SetVpcId(v string) *ListServerGroupsRequest {
	s.VpcId = &v
	return s
}

type ListServerGroupsRequestTag struct {
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

func (s ListServerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListServerGroupsRequestTag) SetKey(v string) *ListServerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListServerGroupsRequestTag) SetValue(v string) *ListServerGroupsRequestTag {
	s.Value = &v
	return s
}

type ListServerGroupsResponseBody struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the server group.
	ServerGroups []*ListServerGroupsResponseBodyServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBody) SetMaxResults(v int32) *ListServerGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetNextToken(v string) *ListServerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetRequestId(v string) *ListServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerGroupsResponseBody) SetServerGroups(v []*ListServerGroupsResponseBodyServerGroups) *ListServerGroupsResponseBody {
	s.ServerGroups = v
	return s
}

func (s *ListServerGroupsResponseBody) SetTotalCount(v int32) *ListServerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListServerGroupsResponseBodyServerGroups struct {
	// The configurations of connection draining.
	ConnectionDrainConfig *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig `json:"ConnectionDrainConfig,omitempty" xml:"ConnectionDrainConfig,omitempty" type:"Struct"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-08-05T18:24:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The configuration of health checks.
	HealthCheckConfig *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The backend protocol. Valid values:
	//
	// 	- **GENEVE**.
	//
	// example:
	//
	// GENEVE
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The IDs of the GWLB instances that are associated with the server group.
	RelatedLoadBalancerIds []*string `json:"RelatedLoadBalancerIds,omitempty" xml:"RelatedLoadBalancerIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **5TCH**: indicates consistent hashing that is based on the following factors: source IP address, destination IP address, source port, protocol, and destination port. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **3TCH**: indicates consistent hashing that is based on the following factors: source IP address, destination IP address, and protocol. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **2TCH**: indicates consistent hashing that is based on the following factors: source IP address and destination IP address. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// example:
	//
	// 5TCH
	Scheduler *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	// The number of server groups.
	//
	// example:
	//
	// 2
	ServerCount        *int32  `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	ServerFailoverMode *string `json:"ServerFailoverMode,omitempty" xml:"ServerFailoverMode,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The server group name.
	//
	// example:
	//
	// testServerGroupName
	ServerGroupName *string `json:"ServerGroupName,omitempty" xml:"ServerGroupName,omitempty"`
	// The status of the server group. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Available**
	//
	// 	- **Configuring**
	//
	// example:
	//
	// Available
	ServerGroupStatus *string `json:"ServerGroupStatus,omitempty" xml:"ServerGroupStatus,omitempty"`
	// The server group type. Valid values:
	//
	// 	- **Instance**: allows you to specify servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- **Ip**: allows you to add servers of by specifying IP addresses.
	//
	// example:
	//
	// Instance
	ServerGroupType *string `json:"ServerGroupType,omitempty" xml:"ServerGroupType,omitempty"`
	// The tags.
	Tags []*ListServerGroupsResponseBodyServerGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72zv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroups) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroups) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroups) SetConnectionDrainConfig(v *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) *ListServerGroupsResponseBodyServerGroups {
	s.ConnectionDrainConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetCreateTime(v string) *ListServerGroupsResponseBodyServerGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetHealthCheckConfig(v *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) *ListServerGroupsResponseBodyServerGroups {
	s.HealthCheckConfig = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetProtocol(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Protocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetRelatedLoadBalancerIds(v []*string) *ListServerGroupsResponseBodyServerGroups {
	s.RelatedLoadBalancerIds = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetResourceGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetScheduler(v string) *ListServerGroupsResponseBodyServerGroups {
	s.Scheduler = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerCount(v int32) *ListServerGroupsResponseBodyServerGroups {
	s.ServerCount = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerFailoverMode(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerFailoverMode = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupName(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupName = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupStatus(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupStatus = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetServerGroupType(v string) *ListServerGroupsResponseBodyServerGroups {
	s.ServerGroupType = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetTags(v []*ListServerGroupsResponseBodyServerGroupsTags) *ListServerGroupsResponseBodyServerGroups {
	s.Tags = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroups) SetVpcId(v string) *ListServerGroupsResponseBodyServerGroups {
	s.VpcId = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig struct {
	// Indicates whether connection draining is enabled. Valid values:
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

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsHealthCheckConfig struct {
	// The backend server port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP**: the internal IP address of a backend server.
	//
	// 	- **domain**: a domain name. The domain name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	// Indicates whether the health check feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The HTTP status codes that the system returns for health checks.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval at which health checks are performed.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: ` - / . % ? # &  `The URL must start with a forward slash (/).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **HTTP**: HTTP health checks simulate a process that uses a web browser to access resources by sending HEAD or GET requests to an instance. These requests are used to check whether the instance is healthy.
	//
	// example:
	//
	// TCP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectPort(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectPort = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckConnectTimeout(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckConnectTimeout = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckDomain(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckDomain = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckEnabled(v bool) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckHttpCode(v []*string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckHttpCode = v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckInterval(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckInterval = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckPath(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckPath = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthCheckProtocol(v string) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetHealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig) SetUnhealthyThreshold(v int32) *ListServerGroupsResponseBodyServerGroupsHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type ListServerGroupsResponseBodyServerGroupsTags struct {
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

func (s ListServerGroupsResponseBodyServerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponseBodyServerGroupsTags) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetKey(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Key = &v
	return s
}

func (s *ListServerGroupsResponseBodyServerGroupsTags) SetValue(v string) *ListServerGroupsResponseBodyServerGroupsTags {
	s.Value = &v
	return s
}

type ListServerGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListServerGroupsResponse) SetHeaders(v map[string]*string) *ListServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListServerGroupsResponse) SetStatusCode(v int32) *ListServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerGroupsResponse) SetBody(v *ListServerGroupsResponseBody) *ListServerGroupsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// WyI2NDQ3MWUxM2EzOWNhMmY0Y2M2YTRiNzZhOWQwNmU1Y2RlNTYzMGEzIiwibiIsIm4iLDEsLTEsMTY5ODcxMzI2NjU0MywiNjU0MDRlYjI2MmI3MDhjY2JiMjM0ZmU0ODNkNTVmMGRhZDllOTBi****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of resource. Valid values:
	//
	// 	- **loadbalancer**: Gateway Load Balancer (GWLB) instance
	//
	// 	- **listener**: listener
	//
	// 	- **servergroup**: server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// d209f4e63ec942c967c50c888a13****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A95325A2-E421-58A6-88AD-7A26CE610F45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID.
	//
	// example:
	//
	// gwlb-nrnrxwd15en27r****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **loadbalancer**: GWLB instance
	//
	// 	- **listener**: listener
	//
	// 	- **servergroup**: server group
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
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
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the resource group to which you want to move the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek253e4oit****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-nrnrxwd15en27r****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- **loadbalancer**: Gateway Load Balancer (GWLB) instance
	//
	// 	- **servergroup**: server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetClientToken(v string) *MoveResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *MoveResourceGroupRequest) SetDryRun(v bool) *MoveResourceGroupRequest {
	s.DryRun = &v
	return s
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 78828B5C-521E-50F3-84D4-7019691D1382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RemoveServersFromServerGroupRequest struct {
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
	// The server group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgp-atstuj3rtoptyui****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The backend servers that you want to remove.
	//
	// > You can remove at most 200 backend servers in each call.
	//
	// This parameter is required.
	Servers []*RemoveServersFromServerGroupRequestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
}

func (s RemoveServersFromServerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequest) SetClientToken(v string) *RemoveServersFromServerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetDryRun(v bool) *RemoveServersFromServerGroupRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServerGroupId(v string) *RemoveServersFromServerGroupRequest {
	s.ServerGroupId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequest) SetServers(v []*RemoveServersFromServerGroupRequestServers) *RemoveServersFromServerGroupRequest {
	s.Servers = v
	return s
}

type RemoveServersFromServerGroupRequestServers struct {
	// The port that is used by the backend server. Valid values:
	//
	// 	- **6081**
	//
	// example:
	//
	// 6081
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The backend server ID.
	//
	// 	- If the server group is of the **Instance*	- type, set this parameter to the IDs of servers of the **Ecs**, **Eni**, or **Eci*	- type.
	//
	// 	- If the server group is of the **Ip*	- type, set ServerId to IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.xxx.xxx
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The type of the backend server. Valid values:
	//
	// 	- **Ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **Eni**: elastic network interface (ENI)
	//
	// 	- **Eci**: elastic container instance
	//
	// 	- **Ip**: IP address
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s RemoveServersFromServerGroupRequestServers) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupRequestServers) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupRequestServers) SetPort(v int32) *RemoveServersFromServerGroupRequestServers {
	s.Port = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerId(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerId = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerIp(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerIp = &v
	return s
}

func (s *RemoveServersFromServerGroupRequestServers) SetServerType(v string) *RemoveServersFromServerGroupRequestServers {
	s.ServerType = &v
	return s
}

type RemoveServersFromServerGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveServersFromServerGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveServersFromServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveServersFromServerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveServersFromServerGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponse) SetHeaders(v map[string]*string) *RemoveServersFromServerGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetStatusCode(v int32) *RemoveServersFromServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveServersFromServerGroupResponse) SetBody(v *RemoveServersFromServerGroupResponseBody) *RemoveServersFromServerGroupResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of resource. Valid values:
	//
	// 	- **loadbalancer**: Gateway Load Balancer (GWLB) instance
	//
	// 	- **listener**: listener
	//
	// 	- **servergroup**: server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify at most 20 tags in each call.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetDryRun(v bool) *TagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54F6E60A-1777-5C17-A6A9-BCC1A7FE945B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified resource. This parameter is ignored if the TagKey parameter is specified. Valid values:
	//
	// 	- **true**: removes all tags from the specified resource.
	//
	// 	- **false**: does not remove all tags from the specified resource. This is the default value.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The resource IDs. You can specify at most 50 resource IDs in each call.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of resource. Valid values:
	//
	// 	- **loadbalancer**: Gateway Load Balancer (GWLB) instance
	//
	// 	- **listener**: listener
	//
	// 	- **servergroup**: server group
	//
	// This parameter is required.
	//
	// example:
	//
	// loadbalancer
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The keys of the tags that you want to remove. You can remove at most 20 tags in each call.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetDryRun(v bool) *UntagResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3DC0064C-C09E-5C99-8FD4-9CDB2DA7FA21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateListenerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The listener description.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// listener_description
	ListenerDescription *string `json:"ListenerDescription,omitempty" xml:"ListenerDescription,omitempty"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-lxce8iqbof2vl0****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The server group ID.
	//
	// example:
	//
	// sgp-sp8d2r6y7t0xtl****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
}

func (s UpdateListenerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeRequest) SetClientToken(v string) *UpdateListenerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetDryRun(v bool) *UpdateListenerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerDescription(v string) *UpdateListenerAttributeRequest {
	s.ListenerDescription = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetListenerId(v string) *UpdateListenerAttributeRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateListenerAttributeRequest) SetServerGroupId(v string) *UpdateListenerAttributeRequest {
	s.ServerGroupId = &v
	return s
}

type UpdateListenerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7DBFC67C-A272-5952-8287-6C3EBE4E04D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateListenerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponseBody) SetRequestId(v string) *UpdateListenerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateListenerAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateListenerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateListenerAttributeResponse) SetHeaders(v map[string]*string) *UpdateListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateListenerAttributeResponse) SetStatusCode(v int32) *UpdateListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateListenerAttributeResponse) SetBody(v *UpdateListenerAttributeResponseBody) *UpdateListenerAttributeResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// testGwlbName
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	TrafficMode      *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
}

func (s UpdateLoadBalancerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeRequest) SetClientToken(v string) *UpdateLoadBalancerAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetDryRun(v bool) *UpdateLoadBalancerAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetLoadBalancerName(v string) *UpdateLoadBalancerAttributeRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *UpdateLoadBalancerAttributeRequest) SetTrafficMode(v string) *UpdateLoadBalancerAttributeRequest {
	s.TrafficMode = &v
	return s
}

type UpdateLoadBalancerAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B956C629-0E8C-5EFF-BAC1-B0E3A8C5CBDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetStatusCode(v int32) *UpdateLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponse) SetBody(v *UpdateLoadBalancerAttributeResponseBody) *UpdateLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

type UpdateLoadBalancerZonesRequest struct {
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
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The GWLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The mappings between zones and vSwitches. You must specify at least one zone. You can specify at most 20 zones. If the region supports two or more zones, we recommend that you select two or more zones.
	//
	// This parameter is required.
	ZoneMappings []*UpdateLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequest) SetClientToken(v string) *UpdateLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetDryRun(v bool) *UpdateLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

type UpdateLoadBalancerZonesRequestZoneMappings struct {
	// The ID of the vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of a GWLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1n75pbs77v5q6p3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the DescribeZones operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

type UpdateLoadBalancerZonesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ED8905C2-A4F6-5E43-87B7-6A5DC8757146
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponseBody) SetRequestId(v string) *UpdateLoadBalancerZonesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoadBalancerZonesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoadBalancerZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoadBalancerZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoadBalancerZonesResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerZonesResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetStatusCode(v int32) *UpdateLoadBalancerZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerZonesResponse) SetBody(v *UpdateLoadBalancerZonesResponseBody) *UpdateLoadBalancerZonesResponse {
	s.Body = v
	return s
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
	// The configurations of the health check feature.
	HealthCheckConfig *UpdateServerGroupAttributeRequestHealthCheckConfig `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty" type:"Struct"`
	// The scheduling algorithm. Valid values:
	//
	// 	- **5TCH**: specifies consistent hashing that is based on the following factors: source IP address, destination IP address, source port, protocol, and destination port. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **3TCH**: specifies consistent hashing that is based on the following factors: source IP address, destination IP address, and protocol. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// 	- **2TCH**: specifies consistent hashing that is based on the following factors: source IP address and destination IP address. Requests that contain the same information based on the preceding factors are forwarded to the same backend server.
	//
	// example:
	//
	// 5TCH
	Scheduler          *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
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
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestConnectionDrainConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) SetConnectionDrainEnabled(v bool) *UpdateServerGroupAttributeRequestConnectionDrainConfig {
	s.ConnectionDrainEnabled = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestConnectionDrainConfig) SetConnectionDrainTimeout(v int32) *UpdateServerGroupAttributeRequestConnectionDrainConfig {
	s.ConnectionDrainTimeout = &v
	return s
}

type UpdateServerGroupAttributeRequestHealthCheckConfig struct {
	// The backend server port that is used by health checks.
	//
	// Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	HealthCheckConnectPort *int32 `json:"HealthCheckConnectPort,omitempty" xml:"HealthCheckConnectPort,omitempty"`
	// The maximum timeout period of a health check response.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **300**.
	//
	// example:
	//
	// 5
	HealthCheckConnectTimeout *int32 `json:"HealthCheckConnectTimeout,omitempty" xml:"HealthCheckConnectTimeout,omitempty"`
	// The domain name that is used for health checks. Valid values:
	//
	// 	- **$SERVER_IP**: the internal IP address of a backend server.
	//
	// 	- **domain**: a domain name. The domain name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), and periods (.).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// $SERVER_IP
	HealthCheckDomain *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
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
	// The HTTP status codes that the system returns for health checks.
	HealthCheckHttpCode []*string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty" type:"Repeated"`
	// The interval at which health checks are performed.
	//
	// Unit: seconds
	//
	// Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// The URL that is used for health checks.
	//
	// The URL must be 1 to 80 characters in length, and can contain letters, digits, and the following special characters: ` - / . % ? # &  `The URL must start with a forward slash (/).
	//
	// > This parameter takes effect only if you set **HealthCheckProtocol*	- to **HTTP**.
	//
	// example:
	//
	// /test/index.html
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The protocol that is used for health checks. Valid values:
	//
	// 	- **TCP**: TCP health checks send TCP SYN packets to a backend server to check whether the port of the backend server is reachable.
	//
	// 	- **HTTP**: HTTP health checks simulate a process that uses a web browser to access resources by sending HEAD or GET requests to an instance. These requests are used to check whether the instance is healthy.
	//
	// example:
	//
	// TCP
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The number of times that an unhealthy backend server must consecutively pass health checks before it is declared healthy. In this case, the health status changes from **fail*	- to **success**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The number of times that a healthy backend server must consecutively fail health checks before it is declared unhealthy. In this case, the health status changes from **success*	- to **fail**.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeRequestHealthCheckConfig) GoString() string {
	return s.String()
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

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetHealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateServerGroupAttributeRequestHealthCheckConfig) SetUnhealthyThreshold(v int32) *UpdateServerGroupAttributeRequestHealthCheckConfig {
	s.UnhealthyThreshold = &v
	return s
}

type UpdateServerGroupAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponseBody) SetRequestId(v string) *UpdateServerGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServerGroupAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServerGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetStatusCode(v int32) *UpdateServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerGroupAttributeResponse) SetBody(v *UpdateServerGroupAttributeResponseBody) *UpdateServerGroupAttributeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gwlb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds backend servers to the server group of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *AddServersToServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the ListServerGroups operation to query the status of the server group.
//
//   - If the server group is in the **Configuring*	- state, the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the server group is running.
//
// 2.  You can call the ListServerGroupServers operation to query the status of the backend server.
//
//   - If the backend server is in the **Adding*	- state, the backend server is being added to the server group.
//
//   - If the backend server is in the **Available*	- state, the server is running.
//
// @param request - AddServersToServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServersToServerGroupResponse
func (client *Client) AddServersToServerGroupWithOptions(request *AddServersToServerGroupRequest, runtime *util.RuntimeOptions) (_result *AddServersToServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		bodyFlat["Servers"] = request.Servers
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddServersToServerGroup"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds backend servers to the server group of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *AddServersToServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the ListServerGroups operation to query the status of the server group.
//
//   - If the server group is in the **Configuring*	- state, the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the server group is running.
//
// 2.  You can call the ListServerGroupServers operation to query the status of the backend server.
//
//   - If the backend server is in the **Adding*	- state, the backend server is being added to the server group.
//
//   - If the backend server is in the **Available*	- state, the server is running.
//
// @param request - AddServersToServerGroupRequest
//
// @return AddServersToServerGroupResponse
func (client *Client) AddServersToServerGroup(request *AddServersToServerGroupRequest) (_result *AddServersToServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddServersToServerGroupResponse{}
	_body, _err := client.AddServersToServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a listener for a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *CreateListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of the task.
//
//   - If the listener is in the **Provisioning*	- state, the listener is being created.
//
//   - If the listener is in the **Running*	- state, the listener is running.
//
// @param request - CreateListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateListenerResponse
func (client *Client) CreateListenerWithOptions(request *CreateListenerRequest, runtime *util.RuntimeOptions) (_result *CreateListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateListener"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a listener for a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *CreateListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of the task.
//
//   - If the listener is in the **Provisioning*	- state, the listener is being created.
//
//   - If the listener is in the **Running*	- state, the listener is running.
//
// @param request - CreateListenerRequest
//
// @return CreateListenerResponse
func (client *Client) CreateListener(request *CreateListenerRequest) (_result *CreateListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateListenerResponse{}
	_body, _err := client.CreateListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *Make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/2806160.html) of GWLB before calling this operation**.
//
//   - When you create a GWLB instance, the service-linked role AliyunServiceRoleForGwlb is automatically created.
//
//   - CreateLoadBalancer is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of a GWLB instance.
//
//   - If the GWLB instance is in the Provisioning state, the GWLB instance is being created.
//
//   - If the GWLB instance is in the Active state, the GWLB instance is created.
//
// @param request - CreateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		body["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadBalancer"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *Make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/2806160.html) of GWLB before calling this operation**.
//
//   - When you create a GWLB instance, the service-linked role AliyunServiceRoleForGwlb is automatically created.
//
//   - CreateLoadBalancer is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of a GWLB instance.
//
//   - If the GWLB instance is in the Provisioning state, the GWLB instance is being created.
//
//   - If the GWLB instance is in the Active state, the GWLB instance is created.
//
// @param request - CreateLoadBalancerRequest
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a server group for a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *CreateServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListServerGroups operation to query the status of the task.
//
//   - If the server group is in the **Creating*	- state, it indicates that the server group is being created.
//
//   - If the server group is in the **Available*	- state, it indicates that the server group is created.
//
// @param request - CreateServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServerGroupResponse
func (client *Client) CreateServerGroupWithOptions(request *CreateServerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainConfig)) {
		bodyFlat["ConnectionDrainConfig"] = request.ConnectionDrainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerFailoverMode)) {
		body["ServerFailoverMode"] = request.ServerFailoverMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		body["ServerGroupName"] = request.ServerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		body["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServerGroup"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a server group for a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *CreateServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListServerGroups operation to query the status of the task.
//
//   - If the server group is in the **Creating*	- state, it indicates that the server group is being created.
//
//   - If the server group is in the **Available*	- state, it indicates that the server group is created.
//
// @param request - CreateServerGroupRequest
//
// @return CreateServerGroupResponse
func (client *Client) CreateServerGroup(request *CreateServerGroupRequest) (_result *CreateServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServerGroupResponse{}
	_body, _err := client.CreateServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a listener from a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *DeleteListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of the task.
//
//   - If the listener is in the **Deleting*	- state, the listener is being deleted.
//
//   - If the listener cannot be found, the listener is deleted.
//
// @param request - DeleteListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteListenerResponse
func (client *Client) DeleteListenerWithOptions(request *DeleteListenerRequest, runtime *util.RuntimeOptions) (_result *DeleteListenerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteListener"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a listener from a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *DeleteListener*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of the task.
//
//   - If the listener is in the **Deleting*	- state, the listener is being deleted.
//
//   - If the listener cannot be found, the listener is deleted.
//
// @param request - DeleteListenerRequest
//
// @return DeleteListenerResponse
func (client *Client) DeleteListener(request *DeleteListenerRequest) (_result *DeleteListenerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteListenerResponse{}
	_body, _err := client.DeleteListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Gateway Load Balancer (GWLB) instance.
//
// @param request - DeleteLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *util.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoadBalancer"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Gateway Load Balancer (GWLB) instance.
//
// @param request - DeleteLoadBalancerRequest
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a server group from a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// You can delete server groups that are not associated with listeners.
//
// @param request - DeleteServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServerGroupResponse
func (client *Client) DeleteServerGroupWithOptions(request *DeleteServerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServerGroup"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a server group from a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// You can delete server groups that are not associated with listeners.
//
// @param request - DeleteServerGroupRequest
//
// @return DeleteServerGroupResponse
func (client *Client) DeleteServerGroup(request *DeleteServerGroupRequest) (_result *DeleteServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServerGroupResponse{}
	_body, _err := client.DeleteServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the most recent region list of Gateway Load Balancer (GWLB).
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent region list of Gateway Load Balancer (GWLB).
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the most recent zone list of Gateway Load Balancer (GWLB).
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent zone list of Gateway Load Balancer (GWLB).
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Gateway Load Balancer (GWLB) listener.
//
// @param request - GetListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListenerAttributeResponse
func (client *Client) GetListenerAttributeWithOptions(request *GetListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerAttribute"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Gateway Load Balancer (GWLB) listener.
//
// @param request - GetListenerAttributeRequest
//
// @return GetListenerAttributeResponse
func (client *Client) GetListenerAttribute(request *GetListenerAttributeRequest) (_result *GetListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerAttributeResponse{}
	_body, _err := client.GetListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health check status of a Gateway Load Balancer (GWLB) listener.
//
// @param request - GetListenerHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListenerHealthStatusResponse
func (client *Client) GetListenerHealthStatusWithOptions(request *GetListenerHealthStatusRequest, runtime *util.RuntimeOptions) (_result *GetListenerHealthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		bodyFlat["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		body["Skip"] = request.Skip
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListenerHealthStatus"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health check status of a Gateway Load Balancer (GWLB) listener.
//
// @param request - GetListenerHealthStatusRequest
//
// @return GetListenerHealthStatusResponse
func (client *Client) GetListenerHealthStatus(request *GetListenerHealthStatusRequest) (_result *GetListenerHealthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetListenerHealthStatusResponse{}
	_body, _err := client.GetListenerHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Gateway Load Balancer (GWLB) instance.
//
// @param request - GetLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoadBalancerAttributeResponse
func (client *Client) GetLoadBalancerAttributeWithOptions(request *GetLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *GetLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoadBalancerAttribute"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Gateway Load Balancer (GWLB) instance.
//
// @param request - GetLoadBalancerAttributeRequest
//
// @return GetLoadBalancerAttributeResponse
func (client *Client) GetLoadBalancerAttribute(request *GetLoadBalancerAttributeRequest) (_result *GetLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoadBalancerAttributeResponse{}
	_body, _err := client.GetLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Gateway Load Balancer (GWLB) listeners.
//
// @param request - ListListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenersResponse
func (client *Client) ListListenersWithOptions(request *ListListenersRequest, runtime *util.RuntimeOptions) (_result *ListListenersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListenerIds)) {
		bodyFlat["ListenerIds"] = request.ListenerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		bodyFlat["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		body["Skip"] = request.Skip
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListListeners"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Gateway Load Balancer (GWLB) listeners.
//
// @param request - ListListenersRequest
//
// @return ListListenersResponse
func (client *Client) ListListeners(request *ListListenersRequest) (_result *ListListenersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListListenersResponse{}
	_body, _err := client.ListListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Gateway Load Balancer (GWLB) instances.
//
// @param request - ListLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLoadBalancersResponse
func (client *Client) ListLoadBalancersWithOptions(request *ListLoadBalancersRequest, runtime *util.RuntimeOptions) (_result *ListLoadBalancersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressIpVersion)) {
		body["AddressIpVersion"] = request.AddressIpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerBusinessStatus)) {
		body["LoadBalancerBusinessStatus"] = request.LoadBalancerBusinessStatus
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		bodyFlat["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerNames)) {
		bodyFlat["LoadBalancerNames"] = request.LoadBalancerNames
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerStatus)) {
		body["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		body["Skip"] = request.Skip
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficMode)) {
		body["TrafficMode"] = request.TrafficMode
	}

	if !tea.BoolValue(util.IsUnset(request.VpcIds)) {
		bodyFlat["VpcIds"] = request.VpcIds
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneIds)) {
		bodyFlat["ZoneIds"] = request.ZoneIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLoadBalancers"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Gateway Load Balancer (GWLB) instances.
//
// @param request - ListLoadBalancersRequest
//
// @return ListLoadBalancersResponse
func (client *Client) ListLoadBalancers(request *ListLoadBalancersRequest) (_result *ListLoadBalancersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLoadBalancersResponse{}
	_body, _err := client.ListLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the server groups of a Gateway Load Balancer (GWLB) instance.
//
// @param request - ListServerGroupServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerGroupServersResponse
func (client *Client) ListServerGroupServersWithOptions(request *ListServerGroupServersRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServerIds)) {
		bodyFlat["ServerIds"] = request.ServerIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIps)) {
		bodyFlat["ServerIps"] = request.ServerIps
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		body["Skip"] = request.Skip
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroupServers"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the server groups of a Gateway Load Balancer (GWLB) instance.
//
// @param request - ListServerGroupServersRequest
//
// @return ListServerGroupServersResponse
func (client *Client) ListServerGroupServers(request *ListServerGroupServersRequest) (_result *ListServerGroupServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupServersResponse{}
	_body, _err := client.ListServerGroupServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the server groups of a Gateway Load Balancer (GWLB) instance.
//
// @param request - ListServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerGroupsResponse
func (client *Client) ListServerGroupsWithOptions(request *ListServerGroupsRequest, runtime *util.RuntimeOptions) (_result *ListServerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServerGroupIds)) {
		bodyFlat["ServerGroupIds"] = request.ServerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupNames)) {
		bodyFlat["ServerGroupNames"] = request.ServerGroupNames
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupType)) {
		body["ServerGroupType"] = request.ServerGroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Skip)) {
		body["Skip"] = request.Skip
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServerGroups"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the server groups of a Gateway Load Balancer (GWLB) instance.
//
// @param request - ListServerGroupsRequest
//
// @return ListServerGroupsResponse
func (client *Client) ListServerGroups(request *ListServerGroupsRequest) (_result *ListServerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServerGroupsResponse{}
	_body, _err := client.ListServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a specified cloud resource belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a specified cloud resource belongs.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes backend servers from the server group of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *RemoveServersFromServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the ListServerGroups operation to query the status of a server group.
//
//   - If the server group is in the **Configuring*	- state, the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the server group is running.
//
// 2.  You can call the ListServerGroupServers operation to query the status of a backend server.
//
//   - If the backend server is in the **Removing*	- state, the backend server is being removed from the server group.
//
//   - If the backend server cannot be found, the backend server is no longer in the server group.
//
// >
//
//   - If connection draining id enabled (**ConnectionDrainEnabled*	- set to true) for the server group of the backend server, the backend server that you remove enters the **Removing*	- state before entering the **Draining*	- state. When the connection draining timeout period (**ConnectionDrainTimeout**) ends, the backend server is removed from the server group.
//
//   - You can add the backend server to the server group again before the connection draining timeout period ends. In this case, the status of the backend server changes from **Draining*	- to **Adding**. After the backend server is added to the server group, the backend server enters the **Available*	- state.
//
// @param request - RemoveServersFromServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveServersFromServerGroupResponse
func (client *Client) RemoveServersFromServerGroupWithOptions(request *RemoveServersFromServerGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveServersFromServerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Servers)) {
		bodyFlat["Servers"] = request.Servers
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveServersFromServerGroup"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes backend servers from the server group of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *RemoveServersFromServerGroup*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background.
//
// 1.  You can call the ListServerGroups operation to query the status of a server group.
//
//   - If the server group is in the **Configuring*	- state, the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the server group is running.
//
// 2.  You can call the ListServerGroupServers operation to query the status of a backend server.
//
//   - If the backend server is in the **Removing*	- state, the backend server is being removed from the server group.
//
//   - If the backend server cannot be found, the backend server is no longer in the server group.
//
// >
//
//   - If connection draining id enabled (**ConnectionDrainEnabled*	- set to true) for the server group of the backend server, the backend server that you remove enters the **Removing*	- state before entering the **Draining*	- state. When the connection draining timeout period (**ConnectionDrainTimeout**) ends, the backend server is removed from the server group.
//
//   - You can add the backend server to the server group again before the connection draining timeout period ends. In this case, the status of the backend server changes from **Draining*	- to **Adding**. After the backend server is added to the server group, the backend server enters the **Available*	- state.
//
// @param request - RemoveServersFromServerGroupRequest
//
// @return RemoveServersFromServerGroupResponse
func (client *Client) RemoveServersFromServerGroup(request *RemoveServersFromServerGroupRequest) (_result *RemoveServersFromServerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveServersFromServerGroupResponse{}
	_body, _err := client.RemoveServersFromServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		bodyFlat["TagKey"] = request.TagKey
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a Gateway Load Balancer (GWLB) listener.
//
// Description:
//
// *UpdateListenerAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of a listener.
//
//   - If the listener is in the **Configuring*	- state, the listener is being modified.
//
//   - If the listener is in the **Running*	- state, the listener is modified.
//
// @param request - UpdateListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateListenerAttributeResponse
func (client *Client) UpdateListenerAttributeWithOptions(request *UpdateListenerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateListenerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerDescription)) {
		body["ListenerDescription"] = request.ListenerDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ListenerId)) {
		body["ListenerId"] = request.ListenerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateListenerAttribute"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a Gateway Load Balancer (GWLB) listener.
//
// Description:
//
// *UpdateListenerAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **GetListenerAttribute*	- operation to query the status of a listener.
//
//   - If the listener is in the **Configuring*	- state, the listener is being modified.
//
//   - If the listener is in the **Running*	- state, the listener is modified.
//
// @param request - UpdateListenerAttributeRequest
//
// @return UpdateListenerAttributeResponse
func (client *Client) UpdateListenerAttribute(request *UpdateListenerAttributeRequest) (_result *UpdateListenerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateListenerAttributeResponse{}
	_body, _err := client.UpdateListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attributes of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
//	UpdateLoadBalancerAttribute is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of the GWLB instance.
//
//	  	- If the GWLB instance is in the Configuring state, the GWLB instance is being modified.
//
//	  	- If the GWLB instance is in the Active state, the GWLB instance is modified.
//
// @param request - UpdateLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerAttributeResponse
func (client *Client) UpdateLoadBalancerAttributeWithOptions(request *UpdateLoadBalancerAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerName)) {
		body["LoadBalancerName"] = request.LoadBalancerName
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficMode)) {
		body["TrafficMode"] = request.TrafficMode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerAttribute"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
//	UpdateLoadBalancerAttribute is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of the GWLB instance.
//
//	  	- If the GWLB instance is in the Configuring state, the GWLB instance is being modified.
//
//	  	- If the GWLB instance is in the Active state, the GWLB instance is modified.
//
// @param request - UpdateLoadBalancerAttributeRequest
//
// @return UpdateLoadBalancerAttributeResponse
func (client *Client) UpdateLoadBalancerAttribute(request *UpdateLoadBalancerAttributeRequest) (_result *UpdateLoadBalancerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerAttributeResponse{}
	_body, _err := client.UpdateLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the zones of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *Make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/2806160.html) of GWLB before calling this operation**.
//
// UpdateLoadBalancerZones is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of the GWLB instance.
//
//   - If the GWLB instance is in the Configuring state, the GWLB instance is being modified.
//
//   - If the GWLB instance is in the Active state, the GWLB instance is modified.
//
// >  Before you call this operation, make sure that all zone parameters, including the current zones and the zones that you want to add, are specified. If you do not specify the current zones, the current zones are deleted. You can call the GetLoadBalancerAttribute operation to query the current zones of a GWLB instance.
//
// @param request - UpdateLoadBalancerZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoadBalancerZonesResponse
func (client *Client) UpdateLoadBalancerZonesWithOptions(request *UpdateLoadBalancerZonesRequest, runtime *util.RuntimeOptions) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerId)) {
		body["LoadBalancerId"] = request.LoadBalancerId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ZoneMappings)) {
		bodyFlat["ZoneMappings"] = request.ZoneMappings
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoadBalancerZones"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the zones of a Gateway Load Balancer (GWLB) instance.
//
// Description:
//
// *Make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/2806160.html) of GWLB before calling this operation**.
//
// UpdateLoadBalancerZones is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the GetLoadBalancerAttribute operation to query the status of the GWLB instance.
//
//   - If the GWLB instance is in the Configuring state, the GWLB instance is being modified.
//
//   - If the GWLB instance is in the Active state, the GWLB instance is modified.
//
// >  Before you call this operation, make sure that all zone parameters, including the current zones and the zones that you want to add, are specified. If you do not specify the current zones, the current zones are deleted. You can call the GetLoadBalancerAttribute operation to query the current zones of a GWLB instance.
//
// @param request - UpdateLoadBalancerZonesRequest
//
// @return UpdateLoadBalancerZonesResponse
func (client *Client) UpdateLoadBalancerZones(request *UpdateLoadBalancerZonesRequest) (_result *UpdateLoadBalancerZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoadBalancerZonesResponse{}
	_body, _err := client.UpdateLoadBalancerZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attributes of a server group.
//
// Description:
//
// *UpdateServerGroupAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListServerGroups operation to query the status of the task.
//
//   - If the server group is in the **Configuring*	- state, the configuration of the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the configuration of the server group is modified.
//
// @param request - UpdateServerGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServerGroupAttributeResponse
func (client *Client) UpdateServerGroupAttributeWithOptions(request *UpdateServerGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateServerGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionDrainConfig)) {
		bodyFlat["ConnectionDrainConfig"] = request.ConnectionDrainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.HealthCheckConfig)) {
		bodyFlat["HealthCheckConfig"] = request.HealthCheckConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Scheduler)) {
		body["Scheduler"] = request.Scheduler
	}

	if !tea.BoolValue(util.IsUnset(request.ServerFailoverMode)) {
		body["ServerFailoverMode"] = request.ServerFailoverMode
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupId)) {
		body["ServerGroupId"] = request.ServerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerGroupName)) {
		body["ServerGroupName"] = request.ServerGroupName
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServerGroupAttribute"),
		Version:     tea.String("2024-04-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a server group.
//
// Description:
//
// *UpdateServerGroupAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListServerGroups operation to query the status of the task.
//
//   - If the server group is in the **Configuring*	- state, the configuration of the server group is being modified.
//
//   - If the server group is in the **Available*	- state, the configuration of the server group is modified.
//
// @param request - UpdateServerGroupAttributeRequest
//
// @return UpdateServerGroupAttributeResponse
func (client *Client) UpdateServerGroupAttribute(request *UpdateServerGroupAttributeRequest) (_result *UpdateServerGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServerGroupAttributeResponse{}
	_body, _err := client.UpdateServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
