// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpEndpointRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationRequest
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *GetBasicAccelerateIpEndpointRelationRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *GetBasicAccelerateIpEndpointRelationRequest
	GetClientToken() *string
	SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationRequest
	GetEndpointId() *string
	SetRegionId(v string) *GetBasicAccelerateIpEndpointRelationRequest
	GetRegionId() *string
}

type GetBasicAccelerateIpEndpointRelationRequest struct {
	// The ID of the accelerated IP address.
	//
	// >  You must specify **EndpointId*	- or **AccelerateIpId**.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
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
	// The ID of the endpoint.
	//
	// >  You must specify **EndpointId*	- or **AccelerateIpId**.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetBasicAccelerateIpEndpointRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetAcceleratorId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetClientToken(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.EndpointId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) SetRegionId(v string) *GetBasicAccelerateIpEndpointRelationRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationRequest) Validate() error {
	return dara.Validate(s)
}
