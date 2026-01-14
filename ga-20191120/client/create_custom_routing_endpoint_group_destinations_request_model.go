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
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The mapping configuration of the endpoint group.
	//
	// You need to specify the backend service ports and protocols for the endpoint group. The ports are mapped to listener ports.
	//
	// You can specify up to 20 mappings in each call.
	//
	// This parameter is required.
	DestinationConfigurations []*CreateCustomRoutingEndpointGroupDestinationsRequestDestinationConfigurations `json:"DestinationConfigurations,omitempty" xml:"DestinationConfigurations,omitempty" type:"Repeated"`
	// The endpoint group ID.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The mappings of the endpoint group.
	//
	// You need to specify the backend service ports and protocols for the endpoint group. The ports are mapped to listener ports.
	//
	// You can specify up to 20 mappings in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// The last port of the backend service port range.
	//
	// Valid values: **1*	- to **65499**. The value of **FromPort*	- must be equal to or smaller than the value of **ToPort**.
	//
	// You can specify up to 20 last ports in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	FromPort *int32 `json:"FromPort,omitempty" xml:"FromPort,omitempty"`
	// The backend service protocol of the endpoint group. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **TCP+UDP: the TCP and UDP protocols.**
	//
	// You can specify up to four backend service protocols for each endpoint group mapping.
	//
	// This parameter is required.
	Protocols []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The response parameters.
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
