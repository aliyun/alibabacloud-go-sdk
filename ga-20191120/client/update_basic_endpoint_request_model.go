// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateBasicEndpointRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *UpdateBasicEndpointRequest
	GetEndpointGroupId() *string
	SetEndpointId(v string) *UpdateBasicEndpointRequest
	GetEndpointId() *string
	SetName(v string) *UpdateBasicEndpointRequest
	GetName() *string
	SetRegionId(v string) *UpdateBasicEndpointRequest
	GetRegionId() *string
}

type UpdateBasicEndpointRequest struct {
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
	// The ID of the endpoint group to which the endpoint that you want to modify belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint that is associated with the basic GA instance.
	//
	// The name must be 1 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBasicEndpointRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateBasicEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateBasicEndpointRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBasicEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBasicEndpointRequest) SetClientToken(v string) *UpdateBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetEndpointGroupId(v string) *UpdateBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetEndpointId(v string) *UpdateBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetName(v string) *UpdateBasicEndpointRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointRequest) SetRegionId(v string) *UpdateBasicEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBasicEndpointRequest) Validate() error {
	return dara.Validate(s)
}
