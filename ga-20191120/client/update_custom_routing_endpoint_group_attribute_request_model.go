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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint group description.
	//
	// The description can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// EndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoint group name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
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
