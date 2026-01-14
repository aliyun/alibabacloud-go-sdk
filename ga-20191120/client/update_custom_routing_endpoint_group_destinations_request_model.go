// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest
	GetClientToken() *string
	SetDestinationConfigurations(v []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *UpdateCustomRoutingEndpointGroupDestinationsRequest
	GetDestinationConfigurations() []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations
	SetDryRun(v bool) *UpdateCustomRoutingEndpointGroupDestinationsRequest
	GetDryRun() *bool
	SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest
	GetEndpointGroupId() *string
	SetRegionId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest
	GetRegionId() *string
}

type UpdateCustomRoutingEndpointGroupDestinationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The mapping configurations of endpoint group.
	//
	// You must specify the backend service port range and protocol of the endpoint group. The specified information is used to map the port range of the associated listener.
	//
	// You can specify at most 20 mapping configurations, which include port ranges and protocol types.
	//
	// This parameter is required.
	DestinationConfigurations []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
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

func (s UpdateCustomRoutingEndpointGroupDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) GetDestinationConfigurations() []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	return s.DestinationConfigurations
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetDestinationConfigurations(v []*UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.DestinationConfigurations = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetDryRun(v bool) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequest) Validate() error {
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

type UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations struct {
	// The ID of the mapping configuration of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dst-abc123****
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The start port of the backend service port range of the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// You can specify up to 20 start ports in each request.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The backend service protocol of the endpoint group. Valid values:
	//
	// 	- **tcp**: TCP
	//
	// 	- **udp**: UDP
	//
	// 	- **tcp,udp**: TCP and UDP
	//
	// You can specify up to four backend service protocols in each configuration.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The end port of the backend service port range of the endpoint group.
	//
	// Valid values: **1*	- to **65499**. The **FromPort*	- value must be smaller than or equal to the **ToPort*	- value.
	//
	// You can specify up to 20 end ports in each request.
	//
	// example:
	//
	// 80
	ToPort *int32 `json:"ToPort,omitempty" xml:"ToPort,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetDestinationId() *string {
	return s.DestinationId
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetFromPort() *int32 {
	return s.FromPort
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetProtocols() []*string {
	return s.Protocols
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) GetToPort() *int32 {
	return s.ToPort
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetDestinationId(v string) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.DestinationId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetFromPort(v int32) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.FromPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetProtocols(v []*string) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.Protocols = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) SetToPort(v int32) *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations {
	s.ToPort = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations) Validate() error {
	return dara.Validate(s)
}
