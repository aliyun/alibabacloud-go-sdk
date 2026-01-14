// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEndpointGroupAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateEndpointGroupAttributeRequest
	GetDescription() *string
	SetEndpointGroupId(v string) *UpdateEndpointGroupAttributeRequest
	GetEndpointGroupId() *string
	SetName(v string) *UpdateEndpointGroupAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateEndpointGroupAttributeRequest
	GetRegionId() *string
}

type UpdateEndpointGroupAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but make sure that the value is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the endpoint group that you want to modify. The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testEndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1ple63864a5hyj5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The name of the endpoint group that you want to modify.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEndpointGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEndpointGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEndpointGroupAttributeRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateEndpointGroupAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEndpointGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEndpointGroupAttributeRequest) SetClientToken(v string) *UpdateEndpointGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetDescription(v string) *UpdateEndpointGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetEndpointGroupId(v string) *UpdateEndpointGroupAttributeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetName(v string) *UpdateEndpointGroupAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) SetRegionId(v string) *UpdateEndpointGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEndpointGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
