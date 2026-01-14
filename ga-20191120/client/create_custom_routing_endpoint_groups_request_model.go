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
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the endpoint groups.
	//
	// You can specify at most five endpoint groups.
	//
	// This parameter is required.
	EndpointGroupConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the custom routing listener.
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
	// The description cannot exceed 256 characters in length and cannot contain `http://` or `https://`.
	//
	// You can specify at most five endpoint group descriptions.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mapping configuration of the endpoint group.
	//
	// You need to specify the backend service ports and protocols for the endpoint group. The ports are mapped to listener ports.
	//
	// You can specify at most 20 mapping configurations for each endpoint group.
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// The information about the endpoints.
	//
	// You can specify at most 10 endpoints for each endpoint group.
	EndpointConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the region in which the endpoint group resides.
	//
	// You can specify at most five region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// You can specify at most five endpoint group names.
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
	// The first backend service port for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be smaller than or equal to the value of **ToPort**.
	//
	// You can specify at most 20 first backend service ports for each endpoint group.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The backend service protocol.
	//
	// You can specify up to four backend service protocols in each mapping configuration.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The last backend service port for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be smaller than or equal to the value of **ToPort**.
	//
	// You can specify at most 20 last backend service ports for each endpoint group.
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
	// The name of the vSwitch that is specified as an endpoint.
	//
	// example:
	//
	// vsw-test01
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The destination to which traffic is forwarded.
	//
	// You can specify at most 20 destinations for each endpoint.
	PolicyConfigurations []*CreateCustomRoutingEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurationsPolicyConfigurations `json:"PolicyConfigurations,omitempty" xml:"PolicyConfigurations,omitempty" type:"Repeated"`
	// The traffic policy that is used to process traffic to the endpoint. Valid values:
	//
	// 	- **AllowAll**: allows all traffic to the endpoint.
	//
	// 	- **DenyAll*	- (default): denies all traffic to the endpoint.
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
	// The IP address of the destination to which traffic is forwarded.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify at most 20 destination IP addresses for each endpoint.
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
	// You can specify port ranges for at most 20 destinations in each endpoint and specify at most five port ranges for each destination.
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
	// The first port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for at most 20 destinations in each endpoint and specify at most five first ports for each destination.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The last port of the destination port range. The value of this parameter must fall within the port range of the endpoint group.
	//
	// This parameter takes effect only when **TrafficToEndpointPolicy*	- is set to **AllowCustom**.
	//
	// You can specify port ranges for at most 20 destinations in each endpoint and specify at most five last ports for each destination.
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
