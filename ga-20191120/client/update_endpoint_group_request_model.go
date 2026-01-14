// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEndpointGroupRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateEndpointGroupRequest
	GetDescription() *string
	SetEndpointConfigurations(v []*UpdateEndpointGroupRequestEndpointConfigurations) *UpdateEndpointGroupRequest
	GetEndpointConfigurations() []*UpdateEndpointGroupRequestEndpointConfigurations
	SetEndpointGroupId(v string) *UpdateEndpointGroupRequest
	GetEndpointGroupId() *string
	SetEndpointGroupRegion(v string) *UpdateEndpointGroupRequest
	GetEndpointGroupRegion() *string
	SetEndpointIpVersion(v string) *UpdateEndpointGroupRequest
	GetEndpointIpVersion() *string
	SetEndpointProtocolVersion(v string) *UpdateEndpointGroupRequest
	GetEndpointProtocolVersion() *string
	SetEndpointRequestProtocol(v string) *UpdateEndpointGroupRequest
	GetEndpointRequestProtocol() *string
	SetHealthCheckEnabled(v bool) *UpdateEndpointGroupRequest
	GetHealthCheckEnabled() *bool
	SetHealthCheckHost(v string) *UpdateEndpointGroupRequest
	GetHealthCheckHost() *string
	SetHealthCheckIntervalSeconds(v int32) *UpdateEndpointGroupRequest
	GetHealthCheckIntervalSeconds() *int32
	SetHealthCheckPath(v string) *UpdateEndpointGroupRequest
	GetHealthCheckPath() *string
	SetHealthCheckPort(v int32) *UpdateEndpointGroupRequest
	GetHealthCheckPort() *int32
	SetHealthCheckProtocol(v string) *UpdateEndpointGroupRequest
	GetHealthCheckProtocol() *string
	SetName(v string) *UpdateEndpointGroupRequest
	GetName() *string
	SetPortOverrides(v []*UpdateEndpointGroupRequestPortOverrides) *UpdateEndpointGroupRequest
	GetPortOverrides() []*UpdateEndpointGroupRequestPortOverrides
	SetRegionId(v string) *UpdateEndpointGroupRequest
	GetRegionId() *string
	SetThresholdCount(v int32) *UpdateEndpointGroupRequest
	GetThresholdCount() *int32
	SetTrafficPercentage(v int32) *UpdateEndpointGroupRequest
	GetTrafficPercentage() *int32
}

type UpdateEndpointGroupRequest struct {
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
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configurations of the endpoints in the endpoint group.
	EndpointConfigurations []*UpdateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the endpoint group is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointIpVersion   *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The protocol version that is used by the backend service. Valid values:
	//
	// 	- **HTTP1.1**
	//
	// 	- **HTTP2**
	//
	// >  This parameter takes effect only when you set EndpointRequestProtocol to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol that is used by the backend service. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// > 	- You can set this parameter only when the listener that is associated with the endpoint group uses the HTTP or HTTPS protocol.
	//
	// > 	- For an HTTP listener, the backend service protocol must be HTTP.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable the health check feature. Valid values: Valid values:
	//
	// 	- **true**: The health check feature is enabled.
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval between two consecutive health checks. Unit: seconds. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path to which health check requests are sent.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that is used for health checks. Valid values: **1*	- to **65535**.
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
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port mapping.
	PortOverrides []*UpdateEndpointGroupRequestPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of consecutive health check failures that must occur before a healthy endpoint group is considered unhealthy, or the number of consecutive health check successes that must occur before an unhealthy endpoint group is considered healthy.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 3
	ThresholdCount *int32 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The weight of the endpoint group when the listener is associated with multiple endpoint groups.
	//
	// example:
	//
	// 20
	TrafficPercentage *int32 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s UpdateEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEndpointGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEndpointGroupRequest) GetEndpointConfigurations() []*UpdateEndpointGroupRequestEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *UpdateEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateEndpointGroupRequest) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *UpdateEndpointGroupRequest) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *UpdateEndpointGroupRequest) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *UpdateEndpointGroupRequest) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckIntervalSeconds() *int32 {
	return s.HealthCheckIntervalSeconds
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *UpdateEndpointGroupRequest) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *UpdateEndpointGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEndpointGroupRequest) GetPortOverrides() []*UpdateEndpointGroupRequestPortOverrides {
	return s.PortOverrides
}

func (s *UpdateEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEndpointGroupRequest) GetThresholdCount() *int32 {
	return s.ThresholdCount
}

func (s *UpdateEndpointGroupRequest) GetTrafficPercentage() *int32 {
	return s.TrafficPercentage
}

func (s *UpdateEndpointGroupRequest) SetClientToken(v string) *UpdateEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetDescription(v string) *UpdateEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointConfigurations(v []*UpdateEndpointGroupRequestEndpointConfigurations) *UpdateEndpointGroupRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupId(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointGroupRegion(v string) *UpdateEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointIpVersion(v string) *UpdateEndpointGroupRequest {
	s.EndpointIpVersion = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointProtocolVersion(v string) *UpdateEndpointGroupRequest {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetEndpointRequestProtocol(v string) *UpdateEndpointGroupRequest {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckEnabled(v bool) *UpdateEndpointGroupRequest {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckHost(v string) *UpdateEndpointGroupRequest {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckIntervalSeconds(v int32) *UpdateEndpointGroupRequest {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckPath(v string) *UpdateEndpointGroupRequest {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckPort(v int32) *UpdateEndpointGroupRequest {
	s.HealthCheckPort = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetHealthCheckProtocol(v string) *UpdateEndpointGroupRequest {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetName(v string) *UpdateEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetPortOverrides(v []*UpdateEndpointGroupRequestPortOverrides) *UpdateEndpointGroupRequest {
	s.PortOverrides = v
	return s
}

func (s *UpdateEndpointGroupRequest) SetRegionId(v string) *UpdateEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetThresholdCount(v int32) *UpdateEndpointGroupRequest {
	s.ThresholdCount = &v
	return s
}

func (s *UpdateEndpointGroupRequest) SetTrafficPercentage(v int32) *UpdateEndpointGroupRequest {
	s.TrafficPercentage = &v
	return s
}

func (s *UpdateEndpointGroupRequest) Validate() error {
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
	return nil
}

type UpdateEndpointGroupRequestEndpointConfigurations struct {
	// Specifies whether to automatically preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > 	- By default, client IP address preservation is disabled for an endpoint group of a UDP or TCP listener. You can configure this parameter based on your business requirements.
	//
	// >	- By default, client IP address preservation is enabled for an endpoint group of an HTTP or HTTP listener. You can obtain client IP addresses by using the X-Forwarded-For header. You cannot disable the feature.
	//
	// >	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >>For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
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
	// >	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >>For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// Enter the IP address, domain name, or instance ID based on the value of the Type parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.XX.XX.21
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
	// > 	- If you set this parameter to **ECS**, **ENI**, **SLB**, **ALB**, **NLB**, or **IpTarget**, and the AliyunServiceRoleForGaVpcEndpoint service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **ALB*	- and the AliyunServiceRoleForGaAlb service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **OSS*	- and the AliyunServiceRoleForGaOss service-linked role does not exist, the system automatically creates the role.
	//
	// >	- If you set this parameter to **NLB*	- and the AliyunServiceRoleForGaNlb service-linked role does not exist, the system automatically creates the role.
	//
	// >>For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
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
	// vpc-2zen6t0u7xhm0k5iz****
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

func (s UpdateEndpointGroupRequestEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEnableClientIPPreservation(v bool) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEnableProxyProtocol(v bool) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetEndpoint(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetSubAddress(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetType(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetVSwitchIds(v []*string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetVpcId(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetWeight(v int32) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *UpdateEndpointGroupRequestEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type UpdateEndpointGroupRequestPortOverrides struct {
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
	// >	- If you configure port mappings for a listener, you cannot modify the listener protocol. You can only switch between HTTP and HTTPS.
	//
	// >	- Listener port: When you modify the listener port range, make sure that the port range includes the ports configured in port mappings. For example, if you set the listener port range to 80 to 82 and map the listener ports to endpoint ports 100 to 102, you cannot change the listener port range to 80 to 81.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s UpdateEndpointGroupRequestPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupRequestPortOverrides) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupRequestPortOverrides) GetEndpointPort() *int32 {
	return s.EndpointPort
}

func (s *UpdateEndpointGroupRequestPortOverrides) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetEndpointPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *UpdateEndpointGroupRequestPortOverrides) SetListenerPort(v int32) *UpdateEndpointGroupRequestPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *UpdateEndpointGroupRequestPortOverrides) Validate() error {
	return dara.Validate(s)
}
