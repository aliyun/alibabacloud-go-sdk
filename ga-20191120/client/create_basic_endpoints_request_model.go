// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicEndpointsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateBasicEndpointsRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *CreateBasicEndpointsRequest
	GetEndpointGroupId() *string
	SetEndpoints(v []*CreateBasicEndpointsRequestEndpoints) *CreateBasicEndpointsRequest
	GetEndpoints() []*CreateBasicEndpointsRequestEndpoints
	SetRegionId(v string) *CreateBasicEndpointsRequest
	GetRegionId() *string
}

type CreateBasicEndpointsRequest struct {
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
	// The ID of the endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoints.
	//
	// This parameter is required.
	Endpoints []*CreateBasicEndpointsRequestEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointsRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicEndpointsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicEndpointsRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateBasicEndpointsRequest) GetEndpoints() []*CreateBasicEndpointsRequestEndpoints {
	return s.Endpoints
}

func (s *CreateBasicEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicEndpointsRequest) SetAcceleratorId(v string) *CreateBasicEndpointsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetClientToken(v string) *CreateBasicEndpointsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetEndpointGroupId(v string) *CreateBasicEndpointsRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointsRequest) SetEndpoints(v []*CreateBasicEndpointsRequestEndpoints) *CreateBasicEndpointsRequest {
	s.Endpoints = v
	return s
}

func (s *CreateBasicEndpointsRequest) SetRegionId(v string) *CreateBasicEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicEndpointsRequest) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBasicEndpointsRequestEndpoints struct {
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The secondary address of the endpoint.
	//
	// This parameter is required only if you set the endpoint type to **ECS**, **ENI**, or **NLB**.
	//
	// 	- If you set the endpoint type to **ECS**, you can set **EndpointSubAddress*	- to the secondary private IP address of the primary ENI. If you leave this parameter empty, the primary private IP address of the primary ENI is used.
	//
	// 	- If you set the endpoint type to **ENI**, you can set **EndpointSubAddress*	- to the secondary private IP address of the secondary ENI. If you leave this parameter empty, the primary private IP address of the secondary ENI is used.
	//
	// 	- If you set the endpoint type to **NLB**, you can set **EndpointSubAddress*	- to the primary private IP address of the NLB backend server.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The secondary address type of the endpoint. Valid values:
	//
	// 	- **primary**: a primary private IP address.
	//
	// 	- **secondary**: a secondary private IP address.
	//
	// This parameter is required only if you set the endpoint type to **ECS**, **ENI**, or **NLB**. If you set the endpoint type to **NLB**, only **primary*	- is supported.
	//
	// example:
	//
	// primary
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// 	- **NLB**: Network Load Balancer (NLB) instance.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the zone where the endpoint resides.
	//
	// This parameter is required only if you set the endpoint type to **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 1 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateBasicEndpointsRequestEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointsRequestEndpoints) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsRequestEndpoints) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *CreateBasicEndpointsRequestEndpoints) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *CreateBasicEndpointsRequestEndpoints) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *CreateBasicEndpointsRequestEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateBasicEndpointsRequestEndpoints) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *CreateBasicEndpointsRequestEndpoints) GetName() *string {
	return s.Name
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointAddress(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointSubAddress(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointSubAddressType(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointSubAddressType = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointType(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetEndpointZoneId(v string) *CreateBasicEndpointsRequestEndpoints {
	s.EndpointZoneId = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) SetName(v string) *CreateBasicEndpointsRequestEndpoints {
	s.Name = &v
	return s
}

func (s *CreateBasicEndpointsRequestEndpoints) Validate() error {
	return dara.Validate(s)
}
