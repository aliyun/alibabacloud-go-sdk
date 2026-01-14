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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, a 2xx HTTP status code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configurations of the endpoints in the endpoint group.
	EndpointConfigurations []*CreateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the region in which to create the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// 	- **default*	- (default): a default endpoint group.
	//
	// 	- **virtual**: a virtual endpoint group.
	//
	// >  When you call this operation to create a virtual endpoint group for a Layer 4 listener, make sure that a default endpoint group is created.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The backend service protocol. Valid values:
	//
	// 	- **HTTP1.1*	- (default)
	//
	// 	- **HTTP2**
	//
	// >  This parameter is required only when you set the EndpointRequestProtocol parameter to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol that is used by the backend service. Default value: HTTP. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// > 	- You can set this parameter only when the listener that is associated with the endpoint group uses **HTTP*	- or **HTTPS**.
	//
	// >	- For an **HTTP*	- listener, the backend service protocol must be **HTTP**.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path to which to send health check requests.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that is used for health checks.
	//
	// example:
	//
	// 20
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol over which to send health check requests. Valid values:
	//
	// 	- **tcp*	- or **TCP**
	//
	// 	- **http*	- or **HTTP**
	//
	// 	- **https*	- or **HTTPS**
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
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port mappings.
	PortOverrides []*CreateEndpointGroupRequestPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Tags of GA instances.
	Tag []*CreateEndpointGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of consecutive health check failures that must occur before a healthy endpoint group is considered unhealthy, or the number of consecutive health check successes that must occur before an unhealthy endpoint group is considered healthy.
	//
	// Valid values: **2*	- to **10**. Default value: **3**.
	//
	// example:
	//
	// 3
	ThresholdCount *int32 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic ratio for the endpoint group when the specified listener is associated with multiple endpoint groups.
	//
	// Valid values: **1*	- to **100**.
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
	// Specifies whether to automatically preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- By default, client IP address preservation is disabled for an endpoint group of a UDP or TCP listener. You can configure this parameter based on your business requirements.
	//
	// > 	- By default, client IP address preservation is enabled for an endpoint group of an HTTP or HTTP listener. You can obtain client IP addresses by using the X-Forwarded-For header. You cannot disable the feature.
	//
	// > 	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// > > For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the proxy protocol to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- This parameter is available only to endpoint groups of TCP listeners.
	//
	// > 	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// > >  For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address, domain name, or instance ID based on the value of Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.1.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The private IP address of the ENI.
	//
	// >  This parameter is available only when you set the endpoint type to **ENI**. If you leave this parameter empty, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Domain**: a custom domain name.
	//
	// 	- **Ip**: a custom IP address.
	//
	// 	- **IpTarget**: a custom private IP address.
	//
	// 	- **PublicIp**: a public IP address provided by Alibaba Cloud.
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: a Server Load Balancer (SLB) instance.
	//
	// 	- **ALB**: an Application Load Balancer (ALB) instance.
	//
	// 	- **OSS**: an Object Storage Service (OSS) bucket.
	//
	// 	- **ENI**: an elastic network interface (ENI).
	//
	// 	- **NLB**: a Network Load Balancer (NLB) instance.
	//
	// > 	- If you set this parameter to **ECS**, **ENI**, **SLB**, **ALB**, **NLB**, or **IpTarget*	- and the AliyunServiceRoleForGaVpcEndpoint service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **ALB*	- and the AliyunServiceRoleForGaAlb service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **OSS*	- and the AliyunServiceRoleForGaOss service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **NLB*	- and the AliyunServiceRoleForGaNlb service-linked role does not exist, the system automatically creates the role.
	//
	// > > For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IDs of vSwitches that are deployed in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID.
	//
	// You can specify one VPC ID for an endpoint group of an intelligent routing listener.
	//
	// >  This parameter is valid and required only if Type is set to **IpTarget**.
	//
	// example:
	//
	// vpc-bp1quce3451z5b2hv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// >  If you set the weight of an endpoint to 0, GA stops distributing traffic to the endpoint. Proceed with caution.
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

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateEndpointGroupRequestEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
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
	// The endpoint port that is mapped to the listener port.
	//
	// example:
	//
	// 80
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port that is mapped to the endpoint port.
	//
	// > 	- You cannot configure port mappings for virtual endpoint groups of TCP listeners. If a virtual endpoint group already exists on the listener, you cannot configure port mappings for the default endpoint group. If port mappings are configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// > 	- If you configure port mappings for a listener, you cannot modify the listener protocol. You can only switch between HTTP and HTTPS.
	//
	// > 	- Listener port: When you modify the listener port range, make sure that the port range includes the ports configured in port mappings. For example, if you set the listener port range to 80 to 82 and map the listener ports to endpoint ports 100 to 102, you cannot change the listener port range to 80 to 81.
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
	// The tag key of the GA instance. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the GA instance. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
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
