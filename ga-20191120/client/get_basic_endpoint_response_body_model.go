// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetBasicEndpointResponseBody
	GetAcceleratorId() *string
	SetEndPointId(v string) *GetBasicEndpointResponseBody
	GetEndPointId() *string
	SetEndpointAddress(v string) *GetBasicEndpointResponseBody
	GetEndpointAddress() *string
	SetEndpointGroupId(v string) *GetBasicEndpointResponseBody
	GetEndpointGroupId() *string
	SetEndpointSubAddress(v string) *GetBasicEndpointResponseBody
	GetEndpointSubAddress() *string
	SetEndpointSubAddressType(v string) *GetBasicEndpointResponseBody
	GetEndpointSubAddressType() *string
	SetEndpointType(v string) *GetBasicEndpointResponseBody
	GetEndpointType() *string
	SetEndpointZoneId(v string) *GetBasicEndpointResponseBody
	GetEndpointZoneId() *string
	SetName(v string) *GetBasicEndpointResponseBody
	GetName() *string
	SetRequestId(v string) *GetBasicEndpointResponseBody
	GetRequestId() *string
	SetState(v string) *GetBasicEndpointResponseBody
	GetState() *string
}

type GetBasicEndpointResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the endpoint that is associated with the basic GA instance.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndPointId *string `json:"EndPointId,omitempty" xml:"EndPointId,omitempty"`
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the endpoint group to which the endpoint belongs.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
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
	// The type of the secondary address of the endpoint.
	//
	// 	- **primary**: a primary private IP address.
	//
	// 	- **secondary**: a secondary private IP address.
	//
	// This parameter is returned if the type of the endpoint is set to **ECS**, **ENI**, or **NLB**. If the endpoint type is set to **NLB**, **primary*	- is returned.
	//
	// example:
	//
	// primary
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	// The type of endpoint. Valid values:
	//
	// 	- **ENI**: elastic network interface (ENI).
	//
	// 	- **SLB**: Classic Load Balancer (CLB) instance.
	//
	// 	- **ECS**: Elastic Compute Service (ECS) instance.
	//
	// 	- **NLB**: Network Load Balancer (NLB) instance.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the zone where the endpoint is deployed.
	//
	// This parameter is returned only when the endpoint type is set to **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- **init**: The endpoint is being initialized.
	//
	// 	- **active**: The endpoint is available.
	//
	// 	- **updating**: The endpoint is being configured.
	//
	// 	- **binding**: The endpoint is being associated.
	//
	// 	- **unbinding**: The endpoint is being disassociated.
	//
	// 	- **deleting**: The endpoint is being deleted.
	//
	// 	- **bound**: The endpoint is associated.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicEndpointResponseBody) GetEndPointId() *string {
	return s.EndPointId
}

func (s *GetBasicEndpointResponseBody) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *GetBasicEndpointResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *GetBasicEndpointResponseBody) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *GetBasicEndpointResponseBody) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *GetBasicEndpointResponseBody) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetBasicEndpointResponseBody) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *GetBasicEndpointResponseBody) GetName() *string {
	return s.Name
}

func (s *GetBasicEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicEndpointResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicEndpointResponseBody) SetAcceleratorId(v string) *GetBasicEndpointResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndPointId(v string) *GetBasicEndpointResponseBody {
	s.EndPointId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointAddress(v string) *GetBasicEndpointResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointGroupId(v string) *GetBasicEndpointResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointSubAddress(v string) *GetBasicEndpointResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointSubAddressType(v string) *GetBasicEndpointResponseBody {
	s.EndpointSubAddressType = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointType(v string) *GetBasicEndpointResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetEndpointZoneId(v string) *GetBasicEndpointResponseBody {
	s.EndpointZoneId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetName(v string) *GetBasicEndpointResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetRequestId(v string) *GetBasicEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicEndpointResponseBody) SetState(v string) *GetBasicEndpointResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
