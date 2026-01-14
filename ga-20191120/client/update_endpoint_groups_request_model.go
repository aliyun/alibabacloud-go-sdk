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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true:*	- performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The configurations of the endpoint groups.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*UpdateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The listener ID.
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
	// Specifies whether to use the proxy protocol to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableClientIPPreservationProxyProtocol *bool `json:"EnableClientIPPreservationProxyProtocol,omitempty" xml:"EnableClientIPPreservationProxyProtocol,omitempty"`
	// Specifies whether to use the TCP Option Address (TOA) module to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableClientIPPreservationToa *bool `json:"EnableClientIPPreservationToa,omitempty" xml:"EnableClientIPPreservationToa,omitempty"`
	// The configurations of the endpoints in the endpoint group.
	EndpointConfigurations []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group.
	//
	// The description cannot exceed 200 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// group1
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp1d2utp8qqe2a44t****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The backend service protocol of the endpoint that is associated with the intelligent routing listener. Valid values:
	//
	// 	- **HTTP1.1*	- (default)
	//
	// 	- **HTTP2**
	//
	// >  You can specify this parameter only if EndpointRequestProtocol is set to HTTPS.
	//
	// example:
	//
	// HTTP1.1
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The backend service protocol. Valid values:
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// > 	- You can specify this parameter only if the listener that is associated with the endpoint group uses HTTP or HTTPS.
	//
	// > 	- The backend service protocol of an HTTP listener must be HTTP.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable the health check feature. Valid values:
	//
	// 	- **true**: enables the health check feature.
	//
	// 	- **false*	- (default): disables the health check feature.
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval at which health checks are performed. Unit: seconds. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The health check path.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that you want to use for health checks.
	//
	// Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 20
	HealthCheckPort *int64 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol over which health check requests are sent. Valid values:
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
	// The port mappings.
	PortOverrides []*UpdateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The number of failed consecutive health checks that must occur before a healthy endpoint group is considered unhealthy or the number of successful consecutive health checks that must occur before an unhealthy endpoint group is considered healthy.
	//
	// Valid values: **2*	- to **10**.
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic ratio of the endpoint group when the specified listener is associated with multiple endpoint groups.
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
	// >>  For more information, see [Preserve client IP addresses](https://help.aliyun.com/document_detail/158080.html).
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
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address, domain name, or instance ID based on the value of Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.0.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The private IP address of the ENI.
	//
	// >  If you set the endpoint type to ENI, you can specify this parameter.
	//
	// >If you leave this parameter empty, the primary private IP address of the ENI is used.
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
	// vpc-uf66oesmrqge1t2gs****
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
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
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
	// > 	- You cannot configure port mappings for virtual endpoint groups of TCP listeners. If a virtual endpoint group already exists on the listener, you cannot configure port mappings for the default endpoint group. If port mappings are configured for the default endpoint group, you cannot add a virtual endpoint group.
	//
	// >	- If you configure port mappings for a listener, you cannot modify the listener protocol. You can only switch between HTTP and HTTPS.
	//
	// >	- Listener port: When you modify the listener port range, make sure that the port range includes the ports configured in port mappings. For example, if you set the listener port range to 80 to 82 and map the listener ports to endpoint ports 100 to 102, you cannot change the listener port range to 80 to 81.
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
