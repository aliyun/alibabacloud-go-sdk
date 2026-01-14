// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateBasicEndpointGroupRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateBasicEndpointGroupRequest
	GetDescription() *string
	SetEndpointAddress(v string) *UpdateBasicEndpointGroupRequest
	GetEndpointAddress() *string
	SetEndpointGroupId(v string) *UpdateBasicEndpointGroupRequest
	GetEndpointGroupId() *string
	SetEndpointSubAddress(v string) *UpdateBasicEndpointGroupRequest
	GetEndpointSubAddress() *string
	SetEndpointType(v string) *UpdateBasicEndpointGroupRequest
	GetEndpointType() *string
	SetName(v string) *UpdateBasicEndpointGroupRequest
	GetName() *string
	SetRegionId(v string) *UpdateBasicEndpointGroupRequest
	GetRegionId() *string
}

type UpdateBasicEndpointGroupRequest struct {
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
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// BasicEndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The secondary address of the endpoint.
	//
	// This parameter is required only if the accelerated IP address is associated with the secondary private IP address of an ECS instance or an ENI.
	//
	// 	- If you set the endpoint type to **ECS**, you can set **EndpointSubAddress*	- to the secondary private IP address of the primary ENI. If you leave this parameter empty, the primary private IP address of the primary ENI is used.
	//
	// 	- If you set the endpoint type to **ENI**, you can set **EndpointSubAddress*	- to the secondary private IP address of the secondary ENI. If you leave this parameter empty, the primary private IP address of the secondary ENI is used.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 1 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the basic GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateBasicEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBasicEndpointGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateBasicEndpointGroupRequest) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *UpdateBasicEndpointGroupRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *UpdateBasicEndpointGroupRequest) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *UpdateBasicEndpointGroupRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *UpdateBasicEndpointGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBasicEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBasicEndpointGroupRequest) SetClientToken(v string) *UpdateBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetDescription(v string) *UpdateBasicEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointAddress(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointAddress = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointGroupId(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointSubAddress(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetEndpointType(v string) *UpdateBasicEndpointGroupRequest {
	s.EndpointType = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetName(v string) *UpdateBasicEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) SetRegionId(v string) *UpdateBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBasicEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
