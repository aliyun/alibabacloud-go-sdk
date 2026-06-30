// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicEndpointRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateBasicEndpointRequest
	GetClientToken() *string
	SetEndpointAddress(v string) *CreateBasicEndpointRequest
	GetEndpointAddress() *string
	SetEndpointGroupId(v string) *CreateBasicEndpointRequest
	GetEndpointGroupId() *string
	SetEndpointSubAddress(v string) *CreateBasicEndpointRequest
	GetEndpointSubAddress() *string
	SetEndpointSubAddressType(v string) *CreateBasicEndpointRequest
	GetEndpointSubAddressType() *string
	SetEndpointType(v string) *CreateBasicEndpointRequest
	GetEndpointType() *string
	SetEndpointZoneId(v string) *CreateBasicEndpointRequest
	GetEndpointZoneId() *string
	SetName(v string) *CreateBasicEndpointRequest
	GetName() *string
	SetRegionId(v string) *CreateBasicEndpointRequest
	GetRegionId() *string
}

type CreateBasicEndpointRequest struct {
	// The instance ID of the basic Alibaba Cloud Global Accelerator (GA).
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The address of the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp1a05txelswuj8g****
	EndpointAddress *string `json:"EndpointAddress,omitempty" xml:"EndpointAddress,omitempty"`
	// The endpoint group ID of the basic Alibaba Cloud Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The secondary address of the endpoint.
	//
	// This parameter is required when the endpoint type is **ECS**, **ENI**, or **NLB**.
	//
	// - If the endpoint type is **ECS**, EndpointSubAddress can be set to a secondary private IP of the primary network interface controller (NIC). If you leave this parameter empty, the primary private IP of the primary network interface controller (NIC) is used.
	//
	// - If the endpoint type is **ENI**, EndpointSubAddress can be set to a secondary private IP of the secondary network interface controller (NIC). If you leave this parameter empty, the primary private IP of the secondary network interface controller (NIC) is used.
	//
	// - If the endpoint type is **NLB**, this parameter is required. Set EndpointSubAddress to the primary private IP of the NLB backend server.
	//
	// example:
	//
	// 172.16.XX.XX
	EndpointSubAddress *string `json:"EndpointSubAddress,omitempty" xml:"EndpointSubAddress,omitempty"`
	// The type of the secondary address of the endpoint. Valid values:
	//
	// - **primary**: The secondary address type is the primary private IP address.
	//
	// - **secondary**: The secondary address type is a secondary private IP address.
	//
	// This parameter is required when the endpoint type is **ECS**, **ENI**, or **NLB**. If the endpoint type is **NLB**, only **primary*	- is supported.
	//
	// example:
	//
	// primary
	EndpointSubAddressType *string `json:"EndpointSubAddressType,omitempty" xml:"EndpointSubAddressType,omitempty"`
	// The endpoint type. Valid values:
	//
	// - **ENI**: Alibaba Cloud elastic network interface (ENI).
	//
	// - **SLB**: Alibaba Cloud Classic Load Balancer (CLB) instance.
	//
	// - **ECS**: Alibaba Cloud ECS instance.
	//
	// - **NLB**: Alibaba Cloud Network Load Balancer (NLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENI
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The zone ID of the endpoint.
	//
	// Currently, this parameter is required only when the endpoint type is **NLB**.
	//
	// example:
	//
	// cn-hangzhou-g
	EndpointZoneId *string `json:"EndpointZoneId,omitempty" xml:"EndpointZoneId,omitempty"`
	// The name of the endpoint for the basic Alibaba Cloud Global Accelerator (GA) instance.
	//
	// The name must be 1 to 128 characters in length and must start with a letter or a Chinese character. The name can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// ep01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Global Accelerator instance. Set the value to **ap-southeast-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateBasicEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateBasicEndpointRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBasicEndpointRequest) GetEndpointAddress() *string {
	return s.EndpointAddress
}

func (s *CreateBasicEndpointRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateBasicEndpointRequest) GetEndpointSubAddress() *string {
	return s.EndpointSubAddress
}

func (s *CreateBasicEndpointRequest) GetEndpointSubAddressType() *string {
	return s.EndpointSubAddressType
}

func (s *CreateBasicEndpointRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateBasicEndpointRequest) GetEndpointZoneId() *string {
	return s.EndpointZoneId
}

func (s *CreateBasicEndpointRequest) GetName() *string {
	return s.Name
}

func (s *CreateBasicEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateBasicEndpointRequest) SetAcceleratorId(v string) *CreateBasicEndpointRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetClientToken(v string) *CreateBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointAddress(v string) *CreateBasicEndpointRequest {
	s.EndpointAddress = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointGroupId(v string) *CreateBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointSubAddress(v string) *CreateBasicEndpointRequest {
	s.EndpointSubAddress = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointSubAddressType(v string) *CreateBasicEndpointRequest {
	s.EndpointSubAddressType = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointType(v string) *CreateBasicEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetEndpointZoneId(v string) *CreateBasicEndpointRequest {
	s.EndpointZoneId = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetName(v string) *CreateBasicEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateBasicEndpointRequest) SetRegionId(v string) *CreateBasicEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateBasicEndpointRequest) Validate() error {
	return dara.Validate(s)
}
