// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupConfigurations(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurations) *UpdateEndpointGroupsRequest
	GetEndpointGroupConfigurations() []*UpdateEndpointGroupsRequestEndpointGroupConfigurations
	SetListenerId(v string) *UpdateEndpointGroupsRequest
	GetListenerId() *string
	SetRegionId(v string) *UpdateEndpointGroupsRequest
	GetRegionId() *string
}

type UpdateEndpointGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a value for this parameter on your client. Make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, a 2xx HTTP status code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configurations of the endpoint group.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*UpdateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEndpointGroupsRequest) GetEndpointGroupConfigurations() []*UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	return s.EndpointGroupConfigurations
}

func (s *UpdateEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *UpdateEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEndpointGroupsRequest) SetClientToken(v string) *UpdateEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetDryRun(v bool) *UpdateEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurations) *UpdateEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetListenerId(v string) *UpdateEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) SetRegionId(v string) *UpdateEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupsRequest) Validate() error {
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

type UpdateEndpointGroupsRequestEndpointGroupConfigurations struct {
	// Specifies whether to use the Proxy Protocol to preserve client IP addresses. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	EnableClientIPPreservationProxyProtocol *bool `json:"EnableClientIPPreservationProxyProtocol,omitempty" xml:"EnableClientIPPreservationProxyProtocol,omitempty"`
	// Specifies whether to use the TCP Option Address (TOA) module to preserve client IP addresses. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	EnableClientIPPreservationToa *bool `json:"EnableClientIPPreservationToa,omitempty" xml:"EnableClientIPPreservationToa,omitempty"`
	// The configurations of the endpoint.
	EndpointConfigurations []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// group1
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp1d2utp8qqe2a44t****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length, start with a letter or a Chinese character, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// group1
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The version of the backend service protocol for endpoints in a listener that uses smart routing. Valid values:
	//
	// - **HTTP1.1*	- (default): HTTP/1.1
	//
	// - **HTTP2**: HTTP/2
	//
	// > This parameter is available only when you set EndpointRequestProtocol to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol of the backend service. Valid values:
	//
	// - **HTTP**: HTTP
	//
	// - **HTTPS**: HTTPS
	//
	// > 	- You can set this parameter only when you create an endpoint group for an HTTP or HTTPS listener.
	//
	// >
	//
	// > 	- For an HTTP listener, the backend service protocol must be HTTP.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable the health check feature.
	//
	// - **true**: enables the health check feature.
	//
	// - **false*	- (default): disables the health check feature.
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
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path of the health check.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that is used for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 20
	HealthCheckPort *int64 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol that is used for health checks.
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
	// The port mapping.
	PortOverrides []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The number of consecutive health checks that an endpoint must pass to be considered healthy, or fail to be considered unhealthy.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic distribution ratio. If a listener is associated with multiple endpoint groups, you can specify this parameter to distribute traffic to the endpoint groups.
	//
	// Valid values: **1*	- to **100**.
	//
	// example:
	//
	// 20
	TrafficPercentage *int64 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEnableClientIPPreservationProxyProtocol() *bool {
	return s.EnableClientIPPreservationProxyProtocol
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEnableClientIPPreservationToa() *bool {
	return s.EnableClientIPPreservationToa
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointConfigurations() []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupDescription() *string {
	return s.EndpointGroupDescription
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupName() *string {
	return s.EndpointGroupName
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckIntervalSeconds() *int64 {
	return s.HealthCheckIntervalSeconds
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPort() *int64 {
	return s.HealthCheckPort
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetPortOverrides() []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	return s.PortOverrides
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetThresholdCount() *int64 {
	return s.ThresholdCount
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) GetTrafficPercentage() *int64 {
	return s.TrafficPercentage
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationProxyProtocol(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationProxyProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEnableClientIPPreservationToa(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EnableClientIPPreservationToa = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupId(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointIpVersion(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointIpVersion = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointProtocolVersion(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckHost(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckHost = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetPortOverrides(v []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurations) Validate() error {
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

type UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Specifies whether to preserve client IP addresses. Valid values:
	//
	// - **true**: preserves client IP addresses.
	//
	// - **false*	- (default): does not preserve client IP addresses.
	//
	// > 	- By default, client IP address preservation is disabled for endpoint groups of TCP and UDP listeners. You can enable it based on your business needs.
	//
	// >
	//
	// > 	- Client IP address preservation is enabled by default for endpoint groups of HTTP and HTTPS listeners. The client IP addresses are retrieved from the X-Forwarded-For header field. You cannot disable this feature.
	//
	// >
	//
	// > 	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >
	//
	// > 	- For more information, see [](t1863665.xdita#).
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Specifies whether to use the Proxy Protocol to preserve client IP addresses. Valid values:
	//
	// - **true**: uses the Proxy Protocol.
	//
	// - **false*	- (default): does not use the Proxy Protocol.
	//
	// > 	- This parameter is available only for endpoint groups of TCP listeners.
	//
	// >
	//
	// > 	- EnableClientIPPreservation and EnableProxyProtocol cannot be set to true at the same time.
	//
	// >
	//
	// > 	- For more information, see [](t1863665.xdita#).
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address, domain name, or instance ID of the endpoint, based on the value of Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.0.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the ENI.
	//
	// > - This parameter is available only when the endpoint type is ENI. You can specify this parameter. If you do not specify this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint.
	//
	// - **Domain**: a custom domain name.
	//
	// - **Ip**: a custom IP address.
	//
	// - **IpTarget**: a custom private IP address.
	//
	// - **PublicIp**: an Alibaba Cloud public IP address.
	//
	// - **ECS**: an Alibaba Cloud Elastic Compute Service (ECS) instance.
	//
	// - **SLB**: an Alibaba Cloud Server Load Balancer (SLB) instance.
	//
	// - **ALB**: an Alibaba Cloud Application Load Balancer (ALB) instance.
	//
	// - **OSS**: an Alibaba Cloud Object Storage Service (OSS) bucket.
	//
	// - **ENI**: an Alibaba Cloud Elastic Network Interface (ENI).
	//
	// - **NLB**: an Alibaba Cloud Network Load Balancer (NLB) instance.
	//
	// > 	- If you set the endpoint type to **ECS**, **ENI**, **SLB**, **NLB**, or **IpTarget**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaVpcEndpoint.
	//
	// >
	//
	// > 	- If you set the endpoint type to **ALB**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaAlb.
	//
	// >
	//
	// > 	- If you set the endpoint type to **OSS**, and the service-linked role does not exist, the system automatically creates a service-linked role named AliyunServiceRoleForGaOss.
	//
	// >
	//
	// > > For more information, see [](t1963894.xdita#).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of vSwitches in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// You can specify at most one VPC ID for an endpoint group that is associated with a listener that uses smart routing.
	//
	// > This parameter is required and takes effect only when the endpoint type is **IpTarget**.
	//
	// example:
	//
	// vpc-uf66oesmrqge1t2gs****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// > If you set the weight of an endpoint to 0, GA stops distributing traffic to the endpoint. Handle this with care.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetProvider() *string {
	return s.Provider
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetWeight() *int64 {
	return s.Weight
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetApiKeys(v []*string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.ApiKeys = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableProxyProtocol(v bool) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetProvider(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Provider = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetSubAddress(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVSwitchIds(v []*string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVpcId(v string) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides struct {
	// The endpoint port.
	//
	// Valid values: **1*	- to **65499**.
	//
	// example:
	//
	// 80
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port.
	//
	// Valid values: **1*	- to **65499**.
	//
	// > - For TCP listeners, virtual endpoint groups do not support port mapping. If a virtual endpoint group already exists under the listener, you cannot configure port mapping for the default endpoint group. If port mapping is already configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// >
	//
	// > - After you configure port mapping, the following limits apply to subsequent listener modifications: You cannot change the listener protocol, except for changing it between HTTP and HTTPS.
	//
	// >
	//
	// > - Listener port: The modified listener port range must include all listener ports that are currently mapped. For example, if the listener port range is 80-82 and the ports are mapped to endpoint ports 100-102, you cannot change the listener port range to 80-81.
	//
	// example:
	//
	// 443
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetEndpointPort() *int64 {
	return s.EndpointPort
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetListenerPort() *int64 {
	return s.ListenerPort
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) Validate() error {
	return dara.Validate(s)
}
