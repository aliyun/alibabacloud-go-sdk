// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCustomRoutingEndpointsRequest
	GetClientToken() *string
	SetEndpointConfigurations(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations) *UpdateCustomRoutingEndpointsRequest
	GetEndpointConfigurations() []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations
	SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *UpdateCustomRoutingEndpointsRequest
	GetRegionId() *string
}

type UpdateCustomRoutingEndpointsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configurations of the endpoint.
	//
	// This parameter is required.
	EndpointConfigurations []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group to which the endpoints that you want to modify belong.
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

func (s UpdateCustomRoutingEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomRoutingEndpointsRequest) GetEndpointConfigurations() []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *UpdateCustomRoutingEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateCustomRoutingEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomRoutingEndpointsRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetEndpointConfigurations(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurations) *UpdateCustomRoutingEndpointsRequest {
	s.EndpointConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequest) Validate() error {
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

type UpdateCustomRoutingEndpointsRequestEndpointConfigurations struct {
	// The ID of the endpoint.
	//
	// You can specify up to 20 endpoint IDs.
	//
	// example:
	//
	// ep-bp1dmlohjjz4kqaun****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The configurations of the policy.
	PolicyConfigurations []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The access policy of traffic for the specified endpoint. Default value: DenyAll. Valid values:
	//
	// 	- **DenyAll**: denies all traffic to the endpoint.
	//
	// 	- **AllowAll**: allows all traffic to the endpoint.
	//
	// 	- **AllowCustom**: allows traffic only to specified destinations.
	//
	//     If you set this parameter to AllowCustom, you must specify IP addresses and port ranges of destinations to which to allow traffic. If you specify only IP addresses and do not specify port ranges, GA can forward traffic to all ports and the specified IP addresses in the destinations.
	//
	// You can specify up to 20 access policies of traffic for the specified endpoint.
	//
	// example:
	//
	// DenyAll
	TrafficToEndpointPolicy *string `json:"TrafficToEndpointPolicy,omitempty" xml:"TrafficToEndpointPolicy,omitempty"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) GetPolicyConfigurations() []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) GetTrafficToEndpointPolicy() *string {
	return s.TrafficToEndpointPolicy
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetEndpointId(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.EndpointId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetPolicyConfigurations(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.PolicyConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) SetTrafficToEndpointPolicy(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurations {
	s.TrafficToEndpointPolicy = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurations) Validate() error {
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

type UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations struct {
	// The IP address of the destination to which to allow traffic.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 20 destination IP addresses for each endpoint.
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
	PortRanges []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) GetPortRanges() []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetAddress(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) SetPortRanges(v []*UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurations) Validate() error {
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

type UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges struct {
	// The start port of the port range in the destination to which to allow traffic. The specified port must fall within the port range of the specified endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations for each endpoint and specify up to 20 start ports for each destination.
	//
	// example:
	//
	// 80
	FromPort *string `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The end port of the port range in the destination to which to allow traffic. The specified port must fall within the port range of the specified endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for up to 20 destinations for each endpoint and specify up to 20 end ports for each destination.
	//
	// example:
	//
	// 80
	ToPort *string `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GetFromPort() *string {
	return s.FromPort
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) GetToPort() *string {
	return s.ToPort
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetFromPort(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) SetToPort(v string) *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointsRequestEndpointConfigurationsPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}
