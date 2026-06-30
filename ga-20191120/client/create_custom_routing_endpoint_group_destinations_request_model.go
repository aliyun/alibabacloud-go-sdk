// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest
	GetClientToken() *string
	SetDestinationConfigurations(v []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *CreateCustomRoutingEndpointGroupDestinationsRequest
	GetDestinationConfigurations() []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations
	SetDryRun(v bool) *CreateCustomRoutingEndpointGroupDestinationsRequest
	GetDryRun() *bool
	SetEndpointGroupId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest
	GetRegionId() *string
}

type CreateCustomRoutingEndpointGroupDestinationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The mapping configurations of the endpoint group.
	//
	// Specify the service port ranges and protocol types of the backend services for the endpoint group. The specified information is mapped to the associated listener port ranges.
	//
	// You can specify up to 20 port ranges and protocol types in a single invoke of this operation.
	//
	// This parameter is required.
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and business limitations without actually creating the mapping configurations create an endpoint group. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the mapping configurations create an endpoint group are created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **ap-southeast-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) GetDestinationConfigurations() []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	return s.DestinationConfigurations
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetDestinationConfigurations(v []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationConfigurations = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *CreateCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequest) Validate() error {
	if s.DestinationConfigurations != nil {
		for _, item := range s.DestinationConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations struct {
	// The start port of the backend service port range for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// You can specify up to 20 start ports in a single request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The protocol types of the backend services for the endpoint group. Valid values:
	//
	// - **TCP**: TCP protocol.
	//
	// - **UDP**: UDP protocol.
	//
	// - **TCP,UDP**: TCP and UDP protocols.
	//
	// The Terms of Service apply to the selected protocols.
	//
	// This parameter is required.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The end port of the backend service port range for the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be less than or equal to the value of **ToPort**.
	//
	// You can specify up to 20 end ports in a single request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetFromPort() *int32 {
	return s.FromPort
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetToPort() *int32 {
	return s.ToPort
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetFromPort(v int32) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetProtocols(v []*string) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetToPort(v int32) *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.ToPort = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) Validate() error {
	return dara.Validate(s)
}
