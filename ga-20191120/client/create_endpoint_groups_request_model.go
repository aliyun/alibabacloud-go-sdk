// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateEndpointGroupsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurations) *CreateEndpointGroupsRequest
	GetEndpointGroupConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurations
	SetListenerId(v string) *CreateEndpointGroupsRequest
	GetListenerId() *string
	SetRegionId(v string) *CreateEndpointGroupsRequest
	GetRegionId() *string
}

type CreateEndpointGroupsRequest struct {
	// The ID of the accelerator.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token used to ensure request idempotence.
	//
	// You can generate the token on your client. Ensure that it is unique across different requests. The value of `ClientToken` can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- is unique for each API request.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run but does not create the resource. The system checks the required parameters, request format, and service limits. If the request fails the dry run, the system returns an error message. If the request passes the dry run, the system returns a 2xx HTTP status code.
	//
	// - **false*	- (default): sends a normal request and creates the resource if the request passes.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configurations of the endpoint groups.
	//
	// You can configure up to 10 endpoint groups.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// > If the listener protocol is **HTTP*	- or **HTTPS**, you can create only one endpoint group in each **CreateEndpointGroups*	- call.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the accelerator is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEndpointGroupsRequest) GetEndpointGroupConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurations {
	return s.EndpointGroupConfigurations
}

func (s *CreateEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEndpointGroupsRequest) SetAcceleratorId(v string) *CreateEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetClientToken(v string) *CreateEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetDryRun(v bool) *CreateEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurations) *CreateEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequest) SetListenerId(v string) *CreateEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetRegionId(v string) *CreateEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) Validate() error {
	if s.EndpointGroupConfigurations != nil {
		for _, item := range s.EndpointGroupConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEndpointGroupsRequestEndpointGroupConfigurations struct {
	// The configurations of the endpoints in the endpoint group.
	EndpointConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// EndpointGroup
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters long, start with a letter or a Chinese character, and contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// group1
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	// The ID of the region where the endpoint group is deployed.
	//
	// You can enter up to 10 endpoint group region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group in an intelligent routing listener. Valid values:
	//
	// - **default*	- (default): a default endpoint group.
	//
	// - **virtual**: a virtual endpoint group.
	//
	// You can enter up to 10 endpoint group types.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The IP version of the backend service. Valid values:
	//
	// - **IPv4*	- (default): Global Accelerator uses only IPv4 addresses to communicate with the backend service.
	//
	// - **IPv6**: Global Accelerator uses only IPv6 addresses to communicate with the backend service.
	//
	// - **ProtocolAffinity**: Global Accelerator communicates with the backend service using the same IP version as the client request.
	//
	// example:
	//
	// IPv4
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The protocol version of the backend service. Valid values:
	//
	// - **HTTP1.1*	- (default): HTTP 1.1.
	//
	// - **HTTP2**: HTTP 2.
	//
	// > You can set this parameter only when `EndpointRequestProtocol` is set to **HTTPS**.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol of the backend service. Valid values:
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// > 	- You can set this parameter only when you create an endpoint group for an HTTP or HTTPS listener.
	//
	// >
	//
	// > 	- For an HTTP listener, you can set this parameter only to HTTP.
	//
	// example:
	//
	// HTTPS
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks for the endpoint group. Valid values:
	//
	// - **true**: enables health checks.
	//
	// - **false*	- (default): disables health checks.
	//
	// You can enable health checks for up to 10 endpoint groups.
	//
	// example:
	//
	// false
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name to which health check requests are sent.
	//
	// example:
	//
	// www.taobao.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval between health checks, in seconds.
	//
	// You can enter up to 10 health check intervals.
	//
	// example:
	//
	// 5
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path used for health checks.
	//
	// You can enter up to 10 health check paths.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port used for health checks. Valid values: **1*	- to **65535**.
	//
	// You can enter up to 10 ports for health checks.
	//
	// example:
	//
	// 443
	HealthCheckPort *int64 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol used for health checks. Valid values:
	//
	// - **tcp*	- or **TCP**: TCP protocol.
	//
	// - **http*	- or **HTTP**: HTTP protocol.
	//
	// - **https*	- or **HTTPS**: HTTPS protocol.
	//
	// You can enter up to 10 health check protocols.
	//
	// example:
	//
	// HTTPS
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The port override settings.
	PortOverrides []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// This parameter is reserved.
	SystemTag []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag `json:"SystemTag,omitempty" xml:"SystemTag,omitempty" type:"Repeated"`
	// The tags to add to the endpoint group. You can specify up to 20 tags.
	Tag []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of consecutive health checks that must succeed for an endpoint to be considered healthy, or fail for it to be considered unhealthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// You can enter up to 10 values for the number of consecutive health checks required for a health status change.
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic distribution percentage for the endpoint group. If an intelligent routing listener is associated with multiple endpoint groups, this parameter specifies the percentage of traffic that is routed to this endpoint group.
	//
	// Valid values: **1*	- to **100**. Default value: **100**.
	//
	// You can enter traffic dial values for up to 10 endpoint groups.
	//
	// example:
	//
	// 100
	TrafficPercentage *int64 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupDescription() *string {
	return s.EndpointGroupDescription
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupName() *string {
	return s.EndpointGroupName
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckIntervalSeconds() *int64 {
	return s.HealthCheckIntervalSeconds
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPort() *int64 {
	return s.HealthCheckPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetPortOverrides() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	return s.PortOverrides
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetSystemTag() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	return s.SystemTag
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetTag() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	return s.Tag
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetThresholdCount() *int64 {
	return s.ThresholdCount
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetTrafficPercentage() *int64 {
	return s.TrafficPercentage
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointIpVersion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointIpVersion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointProtocolVersion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckHost(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetPortOverrides(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetSystemTag(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.SystemTag = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetTag(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.Tag = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) Validate() error {
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
	if s.SystemTag != nil {
		for _, item := range s.SystemTag {
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

type CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Specifies whether to preserve client IP addresses. Valid values:
	//
	// - **true**: preserves client IP addresses.
	//
	// - **false*	- (default): does not preserve client IP addresses.
	//
	// > 	- For endpoint groups of UDP and TCP listeners, the preserve client IP feature is disabled by default. You can enable this feature based on your business requirements.
	//
	// >
	//
	// > 	- For endpoint groups of HTTP and HTTPS listeners, the preserve client IP feature is enabled by default. Client IP addresses are preserved in the X-Forwarded-For header. You cannot disable this feature.
	//
	// >
	//
	// > 	- `EnableClientIPPreservation` and `EnableProxyProtocol` cannot be set to `true` at the same time.
	//
	// >
	//
	// > 	- For more information, see [preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the Proxy Protocol to preserve client IP addresses. Valid values:
	//
	// - **true**: uses the Proxy Protocol to preserve client IP addresses.
	//
	// - **false*	- (default): does not use the Proxy Protocol to preserve client IP addresses.
	//
	// > 	- This parameter is available only for endpoint groups that are associated with TCP listeners.
	//
	// >
	//
	// > 	- `EnableClientIPPreservation` and `EnableProxyProtocol` cannot be set to `true` at the same time.
	//
	// >
	//
	// > 	- For more information, see [preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address or domain name of the endpoint.
	//
	// In an endpoint group of an intelligent routing listener, you can enter a maximum of 100 endpoint IP addresses or domain names.
	//
	// example:
	//
	// 1.1.1.1
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the elastic network interface (ENI).
	//
	// > This parameter is available only when the endpoint type is **ENI**. If you do not specify this parameter, the system uses the primary private IP address of the ENI.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of endpoint in an intelligent routing listener. Valid values:
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
	// - **OSS**: an OSS bucket.
	//
	// - **ENI**: an elastic network interface.
	//
	// - **NLB**: an NLB instance.
	//
	// In an endpoint group of an intelligent routing listener, you can specify up to 100 endpoints.
	//
	// > - If the routing type of the listener is **Standard*	- (intelligent routing), you must configure the endpoint group and endpoint information for the listener. This parameter is required.
	//
	// >
	//
	// > - If you set Type to **ECS**, **ENI**, **SLB**, or **IpTarget*	- and a service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaVpcEndpoint.
	//
	// >
	//
	// > - If you set Type to **ALB*	- and a service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaAlb.
	//
	// >
	//
	// > - If you set Type to **OSS*	- and a service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaOss.
	//
	// >
	//
	// > - If you set Type to **NLB*	- and a service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaNlb.
	//
	// >
	//
	// > > For more information, see [service-linked roles](https://help.aliyun.com/document_detail/178360.html).
	//
	// example:
	//
	// Domain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A list of VSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// In an endpoint group of an intelligent routing listener, you can specify only one VPC ID.
	//
	// > This parameter is required only when you set Type to **IpTarget**.
	//
	// example:
	//
	// vpc-2zekzii824szm3hps****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// > If you set the weight of an endpoint to 0, Global Accelerator stops distributing traffic to the endpoint. Proceed with caution.
	//
	// example:
	//
	// 255
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetProvider() *string {
	return s.Provider
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetApiKeys(v []*string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.ApiKeys = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableProxyProtocol(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetProvider(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Provider = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetSubAddress(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVSwitchIds(v []*string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVpcId(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides struct {
	// The endpoint port used for the port override.
	//
	// example:
	//
	// 443
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port.
	//
	// Valid values: **1*	- to **65499**.
	//
	// > - For TCP listeners, you cannot configure port overrides for a virtual endpoint group. If a virtual endpoint group already exists for the listener, you cannot configure port overrides for the default endpoint group. If port overrides are configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// >
	//
	// > - After you configure a port override, you cannot change the listener protocol, except for switching between HTTP and HTTPS.
	//
	// >
	//
	// > - When you modify the listener port range, the new range must include all listener ports that are used in the port overrides. For example, if the listener port range is 80-82 and a port override is configured to map listener ports to endpoint ports 100-102, you cannot change the listener port range to 80-81.
	//
	// example:
	//
	// 80
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetEndpointPort() *int64 {
	return s.EndpointPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetListenerPort() *int64 {
	return s.ListenerPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag struct {
	// This parameter is reserved.
	//
	// example:
	//
	// -
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// -
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetKey() *string {
	return s.Key
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetScope() *string {
	return s.Scope
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetValue() *string {
	return s.Value
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetKey(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Key = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetScope(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Scope = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetValue(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Value = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsTag struct {
	// The key of the tag. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can enter up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can enter up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GetKey() *string {
	return s.Key
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GetValue() *string {
	return s.Value
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) SetKey(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	s.Key = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) SetValue(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	s.Value = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) Validate() error {
	return dara.Validate(s)
}
