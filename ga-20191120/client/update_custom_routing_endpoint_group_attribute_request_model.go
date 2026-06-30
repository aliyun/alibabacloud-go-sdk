// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest
	GetDescription() *string
	SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest
	GetEndpointGroupId() *string
	SetName(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest
	GetRegionId() *string
}

type UpdateCustomRoutingEndpointGroupAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the endpoint group.
	//
	// The description can be up to 256 characters in length and cannot contain `http://` or `https://` characters.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the Global Accelerator instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetClientToken(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetDescription(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetEndpointGroupId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetName(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) SetRegionId(v string) *UpdateCustomRoutingEndpointGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
