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
	// A client-generated token to ensure the idempotence of the request.
	//
	// The token must be unique across requests and can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- is unique for each API request.
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
	// The configurations of the endpoints.
	EndpointConfigurations []*UpdateEndpointGroupRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the endpoint group is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// Specifies the IP protocol that GA uses to communicate with endpoints. Valid values: ● **IPv4*	- (default): Use IPv4. ● **IPv6**: Use IPv6. ● **ProtocolAffinity**: Use the same IP protocol as the client request.
	//
	// example:
	//
	// IPv4
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The version of the backend service protocol. Valid values:
	//
	// - **HTTP1.1**
	//
	// - **HTTP2**
	//
	// > You can configure this parameter only when `EndpointRequestProtocol` is set to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The backend service protocol. Valid values:
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// > 	- You can configure this parameter only for endpoint groups of HTTP or HTTPS listeners.
	//
	// >
	//
	// > 	- For an HTTP listener, the backend service protocol must be HTTP.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks. Valid values:
	//
	// - **true**: Enables health checks.
	//
	// - **false*	- (default): Disables health checks.
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name for the health check.
	//
	// example:
	//
	// www.taobao.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval between health checks, in seconds. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path for health checks.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 20
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol for health checks. Valid values:
	//
	// - **tcp*	- or **TCP**
	//
	// - **http*	- or **HTTP**
	//
	// - **https*	- or **HTTPS**
	//
	// example:
	//
	// HTTPS
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port override settings.
	PortOverrides []*UpdateEndpointGroupRequestPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of consecutive successful or failed health checks required to change an endpoint\\"s health status.
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
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Specifies whether to preserve client source IP addresses. Valid values:
	//
	// - **true**: Preserves client source IP addresses.
	//
	// - **false*	- (default): Does not preserve client source IP addresses.
	//
	// > 	- For endpoint groups of TCP or UDP listeners, this feature is disabled by default but can be enabled if needed.
	//
	// >
	//
	// > 	- For endpoint groups of HTTP or HTTPS listeners, client source IP addresses are preserved by default. The client IP addresses are retrieved from the X-Forwarded-For header. You cannot disable this feature.
	//
	// >
	//
	// > 	- You cannot set both `EnableClientIPPreservation` and `EnableProxyProtocol` to `true`.
	//
	// >
	//
	// > 	- For more information, see [preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the Proxy Protocol to preserve client source IP addresses. Valid values:
	//
	// - **true**: Preserves client source IP addresses.
	//
	// - **false*	- (default): Does not preserve client source IP addresses.
	//
	// > 	- You can configure this parameter only for endpoint groups of TCP listeners.
	//
	// >
	//
	// > 	- You cannot set both `EnableClientIPPreservation` and `EnableProxyProtocol` to `true`.
	//
	// >
	//
	// > 	- For more information, see [preserve client source IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// Enter an IP address, a domain name, or an instance ID based on the value of the `Type` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120.XX.XX.21
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the elastic network interface.
	//
	// > If the endpoint type is **ENI**, you can specify this parameter. If you omit this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of endpoint. Valid values:
	//
	// - **Domain**: a custom domain name.
	//
	// - **Ip**: a custom IP address.
	//
	// - **IpTarget**: a custom private IP address.
	//
	// - **PublicIp**: an Alibaba Cloud public IP address.
	//
	// - **ECS**: an ECS instance.
	//
	// - **SLB**: an SLB instance.
	//
	// - **ALB**: an ALB instance.
	//
	// - **OSS**: an OSS instance.
	//
	// - **ENI**: an elastic network interface.
	//
	// - **NLB**: an NLB instance.
	//
	// > 	- If the endpoint type is **ECS**, **ENI**, **SLB**, or **IpTarget**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaVpcEndpoint.
	//
	// >
	//
	// > 	- If the endpoint type is **ALB**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaAlb.
	//
	// >
	//
	// > 	- If the endpoint type is **OSS**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaOss.
	//
	// >
	//
	// > 	- If the endpoint type is **NLB**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaNlb.
	//
	// >
	//
	// > > For more information, see [service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A list of vSwitches in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// You can specify at most one VPC ID for an endpoint group that is associated with an intelligent routing listener.
	//
	// > This parameter is required only when the endpoint type is **IpTarget**.
	//
	// example:
	//
	// vpc-2zen6t0u7xhm0k5iz****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// > If you set the weight of an endpoint to 0, Global Accelerator stops distributing traffic to the endpoint. Proceed with caution.
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

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
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

func (s *UpdateEndpointGroupRequestEndpointConfigurations) GetProvider() *string {
	return s.Provider
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

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetApiKeys(v []*string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.ApiKeys = v
	return s
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

func (s *UpdateEndpointGroupRequestEndpointConfigurations) SetProvider(v string) *UpdateEndpointGroupRequestEndpointConfigurations {
	s.Provider = &v
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
	// The endpoint port in the port override settings.
	//
	// example:
	//
	// 80
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port in the port override settings.
	//
	// > - For TCP listeners, virtual endpoint groups do not support port overrides. If a listener is already associated with a virtual endpoint group, you cannot configure port overrides for the default endpoint group. If the default endpoint group has port overrides configured, you cannot add a virtual endpoint group.
	//
	// >
	//
	// > - After you configure port overrides, you can change the listener protocol only between HTTP and HTTPS.
	//
	// >
	//
	// > - The updated listener port range must include all listener ports in the configured port overrides. For example, if the listener port range is 80-82 and port overrides are configured to map the ports to endpoint ports 100-102, you cannot update the listener port range to 80-81.
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
