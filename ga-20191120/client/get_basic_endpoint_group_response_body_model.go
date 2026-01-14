// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetBasicEndpointGroupResponseBody
	GetAcceleratorId() *string
	SetDescription(v string) *GetBasicEndpointGroupResponseBody
	GetDescription() *string
	SetEndpointAddress(v string) *GetBasicEndpointGroupResponseBody
	GetEndpointAddress() *string
	SetEndpointGroupId(v string) *GetBasicEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetEndpointGroupRegion(v string) *GetBasicEndpointGroupResponseBody
	GetEndpointGroupRegion() *string
	SetEndpointSubAddress(v string) *GetBasicEndpointGroupResponseBody
	GetEndpointSubAddress() *string
	SetEndpointType(v string) *GetBasicEndpointGroupResponseBody
	GetEndpointType() *string
	SetName(v string) *GetBasicEndpointGroupResponseBody
	GetName() *string
	SetRequestId(v string) *GetBasicEndpointGroupResponseBody
	GetRequestId() *string
	SetState(v string) *GetBasicEndpointGroupResponseBody
	GetState() *string
}

type GetBasicEndpointGroupResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The description of the endpoint group.
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
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the region where the endpoint group resides.
	//
	// example:
	//
	// cn-shanghai
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The secondary address of the endpoint.
	//
	// This parameter is returned if the endpoint type is **ECS**, **ENI**, or **NLB**.
	//
	// 	- If the endpoint type is **ECS**, **EndpointSubAddress*	- returns the primary or secondary private IP address of the primary ENI.
	//
	// 	- If the endpoint type is **ENI**, **EndpointSubAddress*	- returns the primary or secondary private IP address of the secondary ENI.
	//
	// 	- If the endpoint type is **NLB**, **EndpointSubAddress*	- returns the primary private IP address of the NLB backend server.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The type of endpoint. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// 	- **NLB**: Network Load Balancer (NLB) instance
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The name of the endpoint group.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the endpoint group. Valid values:
	//
	// 	- **init**: being initialized.
	//
	// 	- **active**: running as expected.
	//
	// 	- **updating**: being updated.
	//
	// 	- **deleting**: being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointGroupResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicEndpointGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetBasicEndpointGroupResponseBody) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *GetBasicEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *GetBasicEndpointGroupResponseBody) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *GetBasicEndpointGroupResponseBody) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *GetBasicEndpointGroupResponseBody) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetBasicEndpointGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *GetBasicEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicEndpointGroupResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicEndpointGroupResponseBody) SetAcceleratorId(v string) *GetBasicEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetDescription(v string) *GetBasicEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointAddress(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointGroupId(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointSubAddress(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetEndpointType(v string) *GetBasicEndpointGroupResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetName(v string) *GetBasicEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetRequestId(v string) *GetBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) SetState(v string) *GetBasicEndpointGroupResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
