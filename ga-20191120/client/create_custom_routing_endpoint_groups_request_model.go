// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateCustomRoutingEndpointGroupsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateCustomRoutingEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateCustomRoutingEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) *CreateCustomRoutingEndpointGroupsRequest
	GetEndpointGroupConfigurations() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations
	SetListenerId(v string) *CreateCustomRoutingEndpointGroupsRequest
	GetListenerId() *string
	SetRegionId(v string) *CreateCustomRoutingEndpointGroupsRequest
	GetRegionId() *string
}

type CreateCustomRoutingEndpointGroupsRequest struct {
	// The ID of the Alibaba Cloud Global Accelerator (GA) instance.
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
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating custom routing type endpoint groups. The system checks the required parameters, request format, and business limits. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the custom routing type endpoint groups are created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint group configurations.
	//
	// You can specify up to 5 endpoint group configurations.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the custom routing type listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **ap-southeast-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetEndpointGroupConfigurations() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	return s.EndpointGroupConfigurations
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateCustomRoutingEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetAcceleratorId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetDryRun(v bool) *CreateCustomRoutingEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) *CreateCustomRoutingEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetListenerId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequest) Validate() error {
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

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations struct {
	// The description of the endpoint group.
	//
	// The description can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// You can specify up to 5 endpoint group descriptions.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mapping configurations of the endpoint group.
	//
	// You must specify the backend service port range and protocol type for the endpoint group. The specified information forms a mapping relationship with the associated listener port range.
	//
	// You can specify up to 20 mapping port range and protocol type entries for each endpoint group.
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// The endpoint configurations.
	//
	// You can specify up to 10 endpoint configurations for each endpoint group.
	EndpointConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The region ID of the endpoint group.
	//
	// You can specify up to 5 endpoint group region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length and must start with a letter or Chinese character. The name can contain digits, underscores (_), and hyphens (-).
	//
	// You can specify up to 5 endpoint group names.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GetDestinationConfigurations() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	return s.DestinationConfigurations
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointConfigurations() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) GetName() *string {
	return s.Name
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetDescription(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.Description = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetDestinationConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) SetName(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations {
	s.Name = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations) Validate() error {
	if s.DestinationConfigurations != nil {
		for _, item := range s.DestinationConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations struct {
	// The start port of the backend service for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// You can specify up to 20 start port entries for each endpoint group.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The protocol types of the backend service for the endpoint group.
	//
	// You can specify up to 4 backend service protocol types in each mapping port range and protocol type entry for the endpoint group.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The end port of the backend service for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// You can specify up to 20 end port entries for each endpoint group.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetProtocols(v []*string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) SetToPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations {
	s.ToPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	// The name of the endpoint vSwitch instance.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The traffic destination configurations.
	//
	// You can specify up to 20 traffic destinations for each endpoint.
	PolicyConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy for the backend service. Valid values:
	//
	// - **AllowAll**: allows all traffic to access the specified backend service.
	//
	// - **DenyAll*	- (default): denies all traffic from accessing the specified backend service.
	//
	// - **AllowCustom**: allows traffic to access specified destinations.
	//
	// You must specify the IP address and port range of the destination. If the port range is left empty, all ports of the destination are supported.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The type of the backend service for the endpoint. Valid values:
	//
	//  **PrivateSubNet*	- (default): private CIDR block.
	//
	// example:
	//
	// PrivateSubNet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetPolicyConfigurations() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	if s.PolicyConfigurations != nil {
		for _, item := range s.PolicyConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations struct {
	// The IP address of the traffic destination that can receive traffic.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 traffic destination IP addresses for each endpoint.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The port range of the traffic destination that can receive traffic. The port range must fall within the backend service port range of the endpoint group.
	//
	// If this parameter is left empty, all ports of the traffic destination are supported.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for traffic destinations for each endpoint, and up to 5 port ranges for each traffic destination.
	PortRanges []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) GetPortRanges() []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations) Validate() error {
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	// The start port of the traffic destination that can receive traffic. The port value must fall within the backend service port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for traffic destinations for each endpoint, and up to 5 start ports for each traffic destination.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The end port of the traffic destination that can receive traffic. The port value must fall within the backend service port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for traffic destinations for each endpoint, and up to 5 end ports for each traffic destination.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}
