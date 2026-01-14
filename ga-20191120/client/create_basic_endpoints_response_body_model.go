// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroupId(v string) *CreateBasicEndpointsResponseBody
	GetEndpointGroupId() *string
	SetEndpoints(v []*CreateBasicEndpointsResponseBodyEndpoints) *CreateBasicEndpointsResponseBody
	GetEndpoints() []*CreateBasicEndpointsResponseBodyEndpoints
	SetRequestId(v string) *CreateBasicEndpointsResponseBody
	GetRequestId() *string
}

type CreateBasicEndpointsResponseBody struct {
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoints in the endpoint group.
	Endpoints []*CreateBasicEndpointsResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateBasicEndpointsResponseBody) GetEndpoints() []*CreateBasicEndpointsResponseBodyEndpoints {
	return s.Endpoints
}

func (s *CreateBasicEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicEndpointsResponseBody) SetEndpointGroupId(v string) *CreateBasicEndpointsResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointsResponseBody) SetEndpoints(v []*CreateBasicEndpointsResponseBodyEndpoints) *CreateBasicEndpointsResponseBody {
	s.Endpoints = v
	return s
}

func (s *CreateBasicEndpointsResponseBody) SetRequestId(v string) *CreateBasicEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicEndpointsResponseBody) Validate() error {
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

type CreateBasicEndpointsResponseBodyEndpoints struct {
	// The address of the endpoint.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
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
	// The type of the endpoint. Valid values:
	//
	// 	- **ENI:*	- ENI.
	//
	// 	- **SLB:*	- CLB instance.
	//
	// 	- **ECS:*	- ECS instance.
	//
	// 	- **NLB:*	- NLB instance.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s CreateBasicEndpointsResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointsResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointAddress(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointId(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointSubAddress(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) SetEndpointType(v string) *CreateBasicEndpointsResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointsResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}
