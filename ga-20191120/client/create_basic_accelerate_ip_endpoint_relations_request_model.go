// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpEndpointRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpEndpointRelations(v []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) *CreateBasicAccelerateIpEndpointRelationsRequest
	GetAccelerateIpEndpointRelations() []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations
	SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateBasicAccelerateIpEndpointRelationsRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest
	GetRegionId() *string
}

type CreateBasicAccelerateIpEndpointRelationsRequest struct {
	// A list of accelerated IP addresses and the endpoints with which the accelerated IP addresses are associated.
	//
	// This parameter is required.
	AccelerateIpEndpointRelations []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations `json:"AccelerateIpEndpointRelations,omitempty" xml:"AccelerateIpEndpointRelations,omitempty" type:"Repeated"`
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) GetAccelerateIpEndpointRelations() []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations {
	return s.AccelerateIpEndpointRelations
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetAccelerateIpEndpointRelations(v []*CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.AccelerateIpEndpointRelations = v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetClientToken(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) SetRegionId(v string) *CreateBasicAccelerateIpEndpointRelationsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequest) Validate() error {
	if s.AccelerateIpEndpointRelations != nil {
		for _, item := range s.AccelerateIpEndpointRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations struct {
	// The IDs of the accelerated IP addresses.
	//
	// You can call the [ListBasicAccelerateIps](https://help.aliyun.com/document_detail/2253393.html) operation to query the IDs of the accelerated IP addresses.
	//
	// You can specify up to 20 IP address IDs.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The IDs of the endpoints.
	//
	// You can call the [ListBasicEndpoints](https://help.aliyun.com/document_detail/2253406.html) to query the IDs of the endpoints.
	//
	// You can specify up to 20 endpoint IDs.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsRequestAccelerateIpEndpointRelations) Validate() error {
	return dara.Validate(s)
}
