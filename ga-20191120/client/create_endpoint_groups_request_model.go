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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, a 2xx HTTP status code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the endpoint groups.
	//
	// You can enter the information about up to 10 endpoint groups.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// >  If the listener protocol is **HTTP*	- or **HTTPS**, you can call the **CreateEndpointGroups*	- operation to create only one endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
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
	// The endpoints that are associated with the intelligent routing listener.
	EndpointConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The description of the endpoint group.
	//
	// The description must be up to 200 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// group1
	EndpointGroupDescription *string `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	EndpointGroupName *string `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	// The ID of the region where the endpoint group is created.
	//
	// You can enter the IDs of up to 10 regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group associated with the intelligent routing listener. Valid values:
	//
	// 	- **default*	- (default)
	//
	// 	- **virtual**: a virtual endpoint group.
	//
	// You can specify up to 10 endpoint group types.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The backend service protocol of the endpoint that is associated with the intelligent routing listener. Valid values:
	//
	// 	- **HTTP1.1*	- (default)
	//
	// 	- **HTTP2**
	//
	// >  You can specify this parameter only if the EndpointRequestProtocol parameter is set to HTTPS.
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
	// >	- The backend service protocol of an HTTP listener must be HTTP.
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// Specifies whether to enable health checks for the endpoint group. Valid values:
	//
	// 	- **true**: enables the health check feature.
	//
	// 	- **false*	- (default): disables the health check feature.
	//
	// You can enable the health check feature for up to 10 endpoint groups.
	//
	// example:
	//
	// false
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval at which health checks are performed. Unit: seconds.
	//
	// You can specify up to 10 health check intervals.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int64 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The health check path.
	//
	// You can specify up to 10 health check paths.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that you want to use for health checks. Valid values: **1*	- to **65535**.
	//
	// You can specify up to 10 ports for health checks.
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
	// You can specify up to 10 health check protocols.
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The port mappings.
	PortOverrides []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// This parameter is not in use. Ignore this parameter.
	SystemTag []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag `json:"SystemTag,omitempty" xml:"SystemTag,omitempty" type:"Repeated"`
	// The tags of the endpoint group.
	Tag []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of consecutive health check failures that must occur before a healthy endpoint group is considered unhealthy, or the number of consecutive health check successes that must occur before an unhealthy endpoint group is considered healthy. Valid values: **2*	- to **10**. Default value: **3**.
	//
	// You can specify up to 10 values (the number of consecutive health check successes or consecutive health check failures).
	//
	// example:
	//
	// 3
	ThresholdCount *int64 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The traffic distribution ratio. If an intelligent routing listener is associated with multiple endpoint groups, you can configure this parameter to specify the ratio of traffic distributed to each endpoint group.
	//
	// Valid values: **1*	- to **100**. Default value: **100**.
	//
	// You can specify the traffic distribution ratios for up to 10 endpoint groups.
	//
	// example:
	//
	// 20
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
	// The IP address or domain name of the endpoint that is associated with the intelligent routing listener.
	//
	// You can enter the IP addresses or domain names of up to 100 endpoints in an endpoint group that is associated with the intelligent routing listener.
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
	// >  This parameter is available only when you set the endpoint type to **ENI**. If you leave this parameter empty, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint that is associated with the intelligent routing listener. Valid values:
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
	// You can specify up to 100 endpoint types in the endpoint group that is associated with the intelligent routing listener.
	//
	// > 	- If you set **Type*	- to **Standard**, you can configure the endpoint group and endpoint that are associated with the intelligent routing listener. In addition, this parameter is required.
	//
	// > 	- If you set this parameter to **ECS**, **ENI**, **SLB**, **ALB**, **NLB**, or **IpTarget*	- and the AliyunServiceRoleForGaVpcEndpoint service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **ALB*	- and the AliyunServiceRoleForGaAlb service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **OSS*	- and the AliyunServiceRoleForGaOss service-linked role does not exist, the system automatically creates the role.
	//
	// > 	- If you set this parameter to **NLB*	- and the AliyunServiceRoleForGaNlb service-linked role does not exist, the system automatically creates the role.
	//
	// >>  For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/178360.html).
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
	// vpc-2zekzii824szm3hps****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// Valid values: **0*	- to **255**.
	//
	// >  If you set the weight of an endpoint to 0, GA stops distributing traffic to the endpoint. Proceed with caution.
	//
	// example:
	//
	// 20
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
	// The endpoint port that is mapped to the listener port.
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
	// This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// system-tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// public
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is not in use. Ignore this parameter.
	//
	// example:
	//
	// system-tag-value
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
	// The tag key of the endpoint group. The tag key cannot be an empty string.
	//
	// The tag key must be up to 64 characters in length, and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// You can enter up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the endpoint group. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
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
