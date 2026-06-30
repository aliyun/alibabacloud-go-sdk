// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAccelerateIpEndpointRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteBasicAccelerateIpEndpointRelationRequest
	GetClientToken() *string
	SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest
	GetEndpointId() *string
	SetRegionId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest
	GetRegionId() *string
}

type DeleteBasicAccelerateIpEndpointRelationRequest struct {
	// The accelerated IP address instance ID of the basic GA instance.
	//
	// You can call [ListBasicAccelerateIps](https://help.aliyun.com/document_detail/2253393.html) to query the accelerated IP address instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of a request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The endpoint ID of the basic GA instance.
	//
	// You can call [ListBasicEndpoints](https://help.aliyun.com/document_detail/2253406.html) to query the endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-bp14sz7ftcwwjgrdm****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the Global Accelerator instance. Set the value to **ap-southeast-1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicAccelerateIpEndpointRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetClientToken(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) SetRegionId(v string) *DeleteBasicAccelerateIpEndpointRelationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationRequest) Validate() error {
	return dara.Validate(s)
}
