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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the endpoints.
	//
	// You can specify information for up to 20 endpoints.
	//
	// This parameter is required.
	EndpointConfigurations []*CreateCustomRoutingEndpointsRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group in which to create endpoints.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1bpn0kn908w4nb****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
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
	// The ID of the vSwitch that is specified as an endpoint.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The destination to which traffic is forwarded.
	//
	// You can specify up to 20 destinations for each endpoint.
	PolicyConfigurations []*CreateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy that is used to process traffic to the endpoint. Valid values:
	//
	// 	- **DenyAll*	- (default): denies all traffic to the endpoint.
	//
	// 	- **AllowAll**: allows all traffic to the endpoint.
	//
	// 	- **AllowCustom**: allows traffic only to specified destinations in the endpoint.
	//
	// If you set this parameter to AllowCustom, you must specify IP addresses and port ranges as the destinations to which traffic is distributed. If you specify only IP addresses and do not specify port ranges, GA can forward traffic to the specified IP addresses over all destination ports.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
	// The type of endpoint.
	//
	// Set the value to **PrivateSubNet**, which specifies a private CIDR block. This is the default value.
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
	// The IP address of the destination to which traffic is forwarded.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 destination IP addresses for each endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The port range of the destination to which traffic is forwarded. The value of this parameter must fall within the port range of the endpoint group.
	//
	// If you leave this parameter empty, traffic is forwarded to all destination ports.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations for each endpoint and specify up to 20 port ranges for each destination.
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
	// The first port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations for each endpoint and specify up to 20 first ports for each destination.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations for each endpoint and specify up to 20 last ports for each destination.
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
