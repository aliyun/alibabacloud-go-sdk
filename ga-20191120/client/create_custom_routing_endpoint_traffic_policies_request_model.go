// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointTrafficPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest
	GetClientToken() *string
	SetEndpointId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest
	GetEndpointId() *string
	SetPolicyConfigurations(v []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *CreateCustomRoutingEndpointTrafficPoliciesRequest
	GetPolicyConfigurations() []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations
	SetRegionId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest
	GetRegionId() *string
}

type CreateCustomRoutingEndpointTrafficPoliciesRequest struct {
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
	// The ID of the endpoint for which you want to create traffic destinations.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-2zewuzypq5e6r3pfh****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The configurations of the traffic destinations.
	//
	// You can specify up to 500 traffic destinations for each endpoint.
	//
	// This parameter is required.
	PolicyConfigurations []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) GetPolicyConfigurations() []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	return s.PolicyConfigurations
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetClientToken(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetEndpointId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.EndpointId = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetPolicyConfigurations(v []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.PolicyConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) SetRegionId(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequest) Validate() error {
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

type CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations struct {
	// The IP address of the destination to which traffic is forwarded.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify up to 500 destination IP addresses for each endpoint.
	//
	// > This parameter is required.
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
	// You can specify port ranges for up to 500 traffic destinations in each endpoint and specify up to 10 port ranges for each traffic destination.
	PortRanges []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GetAddress() *string {
	return s.Address
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) GetPortRanges() []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	return s.PortRanges
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetAddress(v string) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.Address = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) SetPortRanges(v []*CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations {
	s.PortRanges = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurations) Validate() error {
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

type CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges struct {
	// The first port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// If the first port and the last port are not specified, traffic on all ports of the destination is allowed.
	//
	// You can specify port ranges for up to 500 destinations in each endpoint and specify up to 10 first ports for each destination.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// If the first port and the last port are not specified, traffic on all ports of the destination is allowed.
	//
	// You can specify port ranges for up to 500 destinations in each endpoint and specify up to 10 last ports for each destination.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetFromPort(v int32) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) SetToPort(v int32) *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges {
	s.ToPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesRequestPolicyConfigurationsPortRanges) Validate() error {
	return dara.Validate(s)
}
