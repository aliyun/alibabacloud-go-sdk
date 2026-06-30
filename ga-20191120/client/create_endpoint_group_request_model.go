// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateEndpointGroupRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateEndpointGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateEndpointGroupRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateEndpointGroupRequest
	GetDryRun() *bool
	SetEndpointConfigurations(v []*CreateEndpointGroupRequestEndpointConfigurations) *CreateEndpointGroupRequest
	GetEndpointConfigurations() []*CreateEndpointGroupRequestEndpointConfigurations
	SetEndpointGroupRegion(v string) *CreateEndpointGroupRequest
	GetEndpointGroupRegion() *string
	SetEndpointGroupType(v string) *CreateEndpointGroupRequest
	GetEndpointGroupType() *string
	SetEndpointIpVersion(v string) *CreateEndpointGroupRequest
	GetEndpointIpVersion() *string
	SetEndpointProtocolVersion(v string) *CreateEndpointGroupRequest
	GetEndpointProtocolVersion() *string
	SetEndpointRequestProtocol(v string) *CreateEndpointGroupRequest
	GetEndpointRequestProtocol() *string
	SetHealthCheckEnabled(v bool) *CreateEndpointGroupRequest
	GetHealthCheckEnabled() *bool
	SetHealthCheckHost(v string) *CreateEndpointGroupRequest
	GetHealthCheckHost() *string
	SetHealthCheckIntervalSeconds(v int32) *CreateEndpointGroupRequest
	GetHealthCheckIntervalSeconds() *int32
	SetHealthCheckPath(v string) *CreateEndpointGroupRequest
	GetHealthCheckPath() *string
	SetHealthCheckPort(v int32) *CreateEndpointGroupRequest
	GetHealthCheckPort() *int32
	SetHealthCheckProtocol(v string) *CreateEndpointGroupRequest
	GetHealthCheckProtocol() *string
	SetListenerId(v string) *CreateEndpointGroupRequest
	GetListenerId() *string
	SetName(v string) *CreateEndpointGroupRequest
	GetName() *string
	SetPortOverrides(v []*CreateEndpointGroupRequestPortOverrides) *CreateEndpointGroupRequest
	GetPortOverrides() []*CreateEndpointGroupRequestPortOverrides
	SetRegionId(v string) *CreateEndpointGroupRequest
	GetRegionId() *string
	SetTag(v []*CreateEndpointGroupRequestTag) *CreateEndpointGroupRequest
	GetTag() []*CreateEndpointGroupRequestTag
	SetThresholdCount(v int32) *CreateEndpointGroupRequest
	GetThresholdCount() *int32
	SetTrafficPercentage(v int32) *CreateEndpointGroupRequest
	GetTrafficPercentage() *int32
}

type CreateEndpointGroupRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token used to ensure request idempotence.
	//
	// You can generate this token, but you must ensure it is unique for each request. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. Each request has a unique **RequestId**.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the dry run, the system returns an error message. If the request passes the dry run, the system returns an HTTP 2xx status code.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the system returns an HTTP 2xx status code and creates the endpoint group.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint configurations.
	EndpointConfigurations []*CreateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the region where the endpoint group is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// - **default*	- (default): a default endpoint group.
	//
	// - **virtual**: a virtual endpoint group.
	//
	// > Before you create a virtual endpoint group for a Layer 4 listener, make sure that you have created a default endpoint group.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The IP version used to communicate with the backend service. Valid values:
	//
	// - **IPv4*	- (default): GA uses only IPv4 to communicate with the backend service.
	//
	// - **IPv6**: GA uses only IPv6 to communicate with the backend service.
	//
	// - **ProtocolAffinity**: GA uses the same IP version as the client request to communicate with the backend service.
	//
	// example:
	//
	// IPv4
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The version of the backend service protocol. Valid values:
	//
	// - **HTTP1.1*	- (default): HTTP/1.1.
	//
	// - **HTTP2**: HTTP/2.
	//
	// > This parameter is available only when `EndpointRequestProtocol` is set to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol used by the backend service. Valid values:
	//
	// - **HTTP*	- (default)
	//
	// - **HTTPS**
	//
	// > 	- This parameter is available only for endpoint groups of **HTTP*	- or **HTTPS*	- listeners.
	//
	// >
	//
	// > 	- For an **HTTP*	- listener, the backend service protocol must be **HTTP**.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **true**: enables health checks.
	//
	// - **false**: disables health checks.
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name used for health checks.
	//
	// example:
	//
	// www.taobao.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The health check interval, in seconds.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path used for health checks.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port used for health checks.
	//
	// example:
	//
	// 20
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol used for health checks. Valid values:
	//
	// - **tcp*	- or **TCP**: TCP
	//
	// - **http*	- or **HTTP**: HTTP
	//
	// - **https*	- or **HTTPS**: HTTPS
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The listener-to-endpoint port mappings.
	PortOverrides []*CreateEndpointGroupRequestPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags of the endpoint group.
	Tag []*CreateEndpointGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of consecutive health checks that must succeed or fail before an endpoint\\"s status changes between healthy and unhealthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 3
	ThresholdCount *int32 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The percentage of traffic distributed to the endpoint group when the listener is associated with multiple endpoint groups. Valid values: **1*	- to **100**.
	//
	// example:
	//
	// 20
	TrafficPercentage *int32 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEndpointGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEndpointGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEndpointGroupRequest) GetEndpointConfigurations() []*CreateEndpointGroupRequestEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateEndpointGroupRequest) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateEndpointGroupRequest) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *CreateEndpointGroupRequest) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *CreateEndpointGroupRequest) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *CreateEndpointGroupRequest) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *CreateEndpointGroupRequest) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateEndpointGroupRequest) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateEndpointGroupRequest) GetHealthCheckIntervalSeconds() *int32 {
	return s.HealthCheckIntervalSeconds
}

func (s *CreateEndpointGroupRequest) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateEndpointGroupRequest) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *CreateEndpointGroupRequest) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateEndpointGroupRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateEndpointGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateEndpointGroupRequest) GetPortOverrides() []*CreateEndpointGroupRequestPortOverrides {
	return s.PortOverrides
}

func (s *CreateEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEndpointGroupRequest) GetTag() []*CreateEndpointGroupRequestTag {
	return s.Tag
}

func (s *CreateEndpointGroupRequest) GetThresholdCount() *int32 {
	return s.ThresholdCount
}

func (s *CreateEndpointGroupRequest) GetTrafficPercentage() *int32 {
	return s.TrafficPercentage
}

func (s *CreateEndpointGroupRequest) SetAcceleratorId(v string) *CreateEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetClientToken(v string) *CreateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetDescription(v string) *CreateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetDryRun(v bool) *CreateEndpointGroupRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointConfigurations(v []*CreateEndpointGroupRequestEndpointConfigurations) *CreateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupRegion(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointGroupType(v string) *CreateEndpointGroupRequest {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointIpVersion(v string) *CreateEndpointGroupRequest {
	s.EndpointIpVersion = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointProtocolVersion(v string) *CreateEndpointGroupRequest {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *CreateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckEnabled(v bool) *CreateEndpointGroupRequest {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckHost(v string) *CreateEndpointGroupRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckIntervalSeconds(v int32) *CreateEndpointGroupRequest {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckPath(v string) *CreateEndpointGroupRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckPort(v int32) *CreateEndpointGroupRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetHealthCheckProtocol(v string) *CreateEndpointGroupRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetListenerId(v string) *CreateEndpointGroupRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetName(v string) *CreateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetPortOverrides(v []*CreateEndpointGroupRequestPortOverrides) *CreateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

func (s *CreateEndpointGroupRequest) SetRegionId(v string) *CreateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetTag(v []*CreateEndpointGroupRequestTag) *CreateEndpointGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateEndpointGroupRequest) SetThresholdCount(v int32) *CreateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupRequest) SetTrafficPercentage(v int32) *CreateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

func (s *CreateEndpointGroupRequest) Validate() error {
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortOverrides != nil {
		for _, item := range s.PortOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateEndpointGroupRequestEndpointConfigurations struct {
	// The API keys for the AI service.
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Specifies whether to preserve client source IP addresses. Valid values:
	//
	// - **true**: preserves client source IP addresses.
	//
	// - **false*	- (default): does not preserve client source IP addresses.
	//
	// > 	- By default, this feature is disabled for endpoint groups that are associated with TCP or UDP listeners. You can enable this feature based on your business requirements.
	//
	// >
	//
	// > 	- By default, this feature is enabled for endpoint groups that are associated with HTTP or HTTPS listeners. The source IP address is retrieved from the X-Forwarded-For header field. This feature cannot be disabled.
	//
	// >
	//
	// > 	- `EnableClientIPPreservation` and `EnableProxyProtocol` cannot both be set to `true`.
	//
	// >
	//
	// > 	- For more information, see [Preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the PROXY protocol to preserve client source IP addresses. Valid values:
	//
	// - **true**: uses the PROXY protocol.
	//
	// - **false*	- (default): does not use the PROXY protocol.
	//
	// > 	- This parameter can be configured only for endpoint groups that are associated with TCP listeners.
	//
	// >
	//
	// > 	- `EnableClientIPPreservation` and `EnableProxyProtocol` cannot both be set to `true`.
	//
	// >
	//
	// > 	- For more information, see [Preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address, domain name, or resource ID of the endpoint. The value of this parameter depends on the value of the `Type` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.1.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The AI service provider. Set this to `BAILIAN` to use Alibaba Cloud Model Studio.
	//
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the ENI.
	//
	// > This parameter applies only when the endpoint type is set to **ENI**. If you omit this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.X.X
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Domain**: a custom domain name.
	//
	// - **Ip**: a custom IP address.
	//
	// - **IpTarget**: a custom private IP address.
	//
	// - **PublicIp**: an Alibaba Cloud public IP address.
	//
	// - **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// - **SLB**: a Server Load Balancer (SLB) instance.
	//
	// - **ALB**: an Application Load Balancer (ALB) instance.
	//
	// - **OSS**: an Object Storage Service (OSS) bucket.
	//
	// - **ENI**: an elastic network interface (ENI).
	//
	// - **NLB**: a Network Load Balancer (NLB) instance.
	//
	// > 	- If you set the endpoint type to **ECS**, **ENI**, **SLB**, **ALB**, **NLB**, or **IpTarget**, the system automatically creates a service-linked role named AliyunServiceRoleForGaVpcEndpoint if the role does not exist.
	//
	// >
	//
	// > 	- If you set the endpoint type to **ALB**, the system automatically creates a service-linked role named AliyunServiceRoleForGaAlb if the role does not exist.
	//
	// >
	//
	// > 	- If you set the endpoint type to **OSS**, the system automatically creates a service-linked role named AliyunServiceRoleForGaOss if the role does not exist.
	//
	// >
	//
	// > 	- If you set the endpoint type to **NLB**, the system automatically creates a service-linked role named AliyunServiceRoleForGaNlb if the role does not exist.
	//
	// >
	//
	// > > For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A list of vSwitches in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud (VPC).
	//
	// You can specify at most one VPC ID for an endpoint group of an intelligent routing listener.
	//
	// > This parameter is required only when the endpoint type is set to **IpTarget**.
	//
	// example:
	//
	// vpc-bp1quce3451z5b2hv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// > If you set the weight of an endpoint to 0, GA stops distributing traffic to it. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateEndpointGroupRequestEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetProvider() *string {
	return s.Provider
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetApiKeys(v []*string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.ApiKeys = v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEnableProxyProtocol(v bool) *CreateEndpointGroupRequestEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetProvider(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Provider = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetSubAddress(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetType(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetVSwitchIds(v []*string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetVpcId(v string) *CreateEndpointGroupRequestEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *CreateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupRequestPortOverrides struct {
	// The endpoint port for the port mapping.
	//
	// example:
	//
	// 80
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port for the port mapping.
	//
	// > - For TCP listeners, you cannot configure port mappings for virtual endpoint groups. If a listener is associated with a virtual endpoint group, you cannot configure port mappings for the default endpoint group. If a default endpoint group has port mappings configured, you cannot add a virtual endpoint group.
	//
	// >
	//
	// > - After you configure port mappings, you cannot change the listener protocol, except for switching between HTTP and HTTPS.
	//
	// >
	//
	// > - When you modify the listener port range, the new range must include all listener ports used in port mappings. For example, if the listener port range is 80-82 and the listener ports are mapped to endpoint ports 100-102, you cannot change the listener port range to 80-81.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateEndpointGroupRequestPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestPortOverrides) GetEndpointPort() *int32 {
	return s.EndpointPort
}

func (s *CreateEndpointGroupRequestPortOverrides) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *CreateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *CreateEndpointGroupRequestPortOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupRequestTag struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters long and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters long and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEndpointGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEndpointGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEndpointGroupRequestTag) SetKey(v string) *CreateEndpointGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEndpointGroupRequestTag) SetValue(v string) *CreateEndpointGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEndpointGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
