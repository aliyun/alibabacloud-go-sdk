// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomRoutingEndpointsRequest
	GetClientToken() *string
	SetEndpointConfigurations(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurations) *CreateCustomRoutingEndpointsRequest
	GetEndpointConfigurations() []*CreateCustomRoutingEndpointsRequestEndpointConfigurations
	SetEndpointGroupId(v string) *CreateCustomRoutingEndpointsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *CreateCustomRoutingEndpointsRequest
	GetRegionId() *string
}

type CreateCustomRoutingEndpointsRequest struct {
	// The client token that is used to ensure the idempotence of a request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint configurations.
	//
	// You can specify up to 20 endpoint configurations.
	//
	// This parameter is required.
	EndpointConfigurations []*CreateCustomRoutingEndpointsRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group in which you want to create an endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1bpn0kn908w4nb****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomRoutingEndpointsRequest) GetEndpointConfigurations() []*CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateCustomRoutingEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateCustomRoutingEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomRoutingEndpointsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetEndpointConfigurations(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurations) *CreateCustomRoutingEndpointsRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *CreateCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequest) Validate() error {
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

type CreateCustomRoutingEndpointsRequestEndpointConfigurations struct {
	// The instance ID of the endpoint vSwitch.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The traffic destination configurations.
	//
	// You can specify up to 20 traffic destinations for each endpoint.
	PolicyConfigurations []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy for the backend service. Valid values:
	//
	// - **DenyAll*	- (default): denies all traffic to the specified backend service.
	//
	// - **AllowAll**: allows all traffic to the specified backend service.
	//
	// - **AllowCustom**: allows traffic only to specified destinations.
	//
	// You must specify the IP address and port range of the destination. If the port range is left empty, all ports of the destination are supported.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The backend service type of the endpoint. Valid values:
	//
	//  **PrivateSubNet*	- (default): private CIDR block.
	//
	// example:
	//
	// PrivateSubNet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) GetPolicyConfigurations() []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetEndpoint(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) SetType(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurations) Validate() error {
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

type CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations struct {
	// The IP address of the traffic destination that can receive traffic.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 traffic destination IP addresses for each endpoint.
	//
	// This parameter is required.
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
	// You can specify up to 20 port ranges for each endpoint, and up to 20 port ranges for each traffic destination.
	PortRanges []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GetPortRanges() []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) Validate() error {
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

type CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	// The start port of the traffic destination that can receive traffic. The port value must fall within the backend service port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for each endpoint, and up to 20 start ports for each traffic destination.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The end port of the traffic destination that can receive traffic. The port value must fall within the backend service port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 port ranges for each endpoint, and up to 20 end ports for each traffic destination.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}
