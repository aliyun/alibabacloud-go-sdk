// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicEndpointGroupRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateBasicEndpointGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateBasicEndpointGroupRequest
	GetDescription() *string
	SetEndpointAddress(v string) *CreateBasicEndpointGroupRequest
	GetEndpointAddress() *string
	SetEndpointGroupRegion(v string) *CreateBasicEndpointGroupRequest
	GetEndpointGroupRegion() *string
	SetEndpointSubAddress(v string) *CreateBasicEndpointGroupRequest
	GetEndpointSubAddress() *string
	SetEndpointType(v string) *CreateBasicEndpointGroupRequest
	GetEndpointType() *string
	SetName(v string) *CreateBasicEndpointGroupRequest
	GetName() *string
	SetRegionId(v string) *CreateBasicEndpointGroupRequest
	GetRegionId() *string
}

type CreateBasicEndpointGroupRequest struct {
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
	// The description of the endpoint group.
	//
	// The description can be up to 200 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// BasicEndpointGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint address.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the region to which the endpoint group belongs.
	//
	// You can call the [ListAvailableBusiRegions](https://help.aliyun.com/document_detail/2253223.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The secondary address of the endpoint.
	//
	// You must specify this parameter when the accelerated IP address is associated with the secondary private IP address of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
	//
	// 	- When the endpoint type is **ECS**, you can set **EndpointSubAddress*	- to the secondary private IP address of the primary ENI. If the parameter is left empty, the primary private IP address of the primary ENI is used.
	//
	// 	- If the endpoint type is **ENI**, you can set **EndpointSubAddress*	- to the secondary private IP address of the secondary ENI. If the parameter is left empty, the primary private IP address of the secondary ENI is used.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **ENI**
	//
	// 	- **SLB**
	//
	// 	- **ECS**
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The name of the endpoint group.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointGroupRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicEndpointGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicEndpointGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBasicEndpointGroupRequest) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *CreateBasicEndpointGroupRequest) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateBasicEndpointGroupRequest) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *CreateBasicEndpointGroupRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateBasicEndpointGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateBasicEndpointGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicEndpointGroupRequest) SetAcceleratorId(v string) *CreateBasicEndpointGroupRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetClientToken(v string) *CreateBasicEndpointGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetDescription(v string) *CreateBasicEndpointGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointAddress(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointGroupRegion(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointSubAddress(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetEndpointType(v string) *CreateBasicEndpointGroupRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetName(v string) *CreateBasicEndpointGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) SetRegionId(v string) *CreateBasicEndpointGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicEndpointGroupRequest) Validate() error {
	return dara.Validate(s)
}
