// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAccelerateIpEndpointRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetAcceleratorId() *string
	SetEndpointAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointAddress() *string
	SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointName() *string
	SetEndpointSubAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointSubAddress() *string
	SetEndpointSubAddressType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointSubAddressType() *string
	SetEndpointType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointType() *string
	SetEndpointZoneId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointZoneId() *string
	SetIpAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetIpAddress() *string
	SetRequestId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetRequestId() *string
	SetState(v string) *GetBasicAccelerateIpEndpointRelationResponseBody
	GetState() *string
}

type GetBasicAccelerateIpEndpointRelationResponseBody struct {
	// The ID of the accelerated IP address.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp11v53zfftax68b0daws
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
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
	// The name of the endpoint.
	//
	// example:
	//
	// ep01
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
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
	// This parameter is returned if the endpoint type is **ECS**, **ENI**, or **NLB**. If the endpoint type is set to **NLB**, **primary*	- is returned.
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
	// The ID of the zone where the endpoint is created.
	//
	// This parameter is returned only when the endpoint type is **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// 116.132.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The association status between the accelerated IP address and endpoint.
	//
	// A value of **active*	- indicates that the accelerated IP address is associated with the endpoint.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetAcceleratorId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointName(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointName = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointSubAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointSubAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointSubAddressType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointSubAddressType = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointType(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetEndpointZoneId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointZoneId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetIpAddress(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.IpAddress = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *GetBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicAccelerateIpEndpointRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
