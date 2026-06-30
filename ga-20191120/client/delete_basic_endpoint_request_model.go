// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteBasicEndpointRequest
	GetClientToken() *string
	SetEndpointGroupId(v string) *DeleteBasicEndpointRequest
	GetEndpointGroupId() *string
	SetEndpointId(v string) *DeleteBasicEndpointRequest
	GetEndpointId() *string
	SetRegionId(v string) *DeleteBasicEndpointRequest
	GetRegionId() *string
}

type DeleteBasicEndpointRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Set this parameter to a value that is generated from your client to make sure that the value is unique among different requests. The ClientToken value supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- value as the **ClientToken*	- value. The **RequestId*	- value is different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint group to which the endpoint that you want to delete belongs. You can call the [ListBasicEndpointGroups](~~~~) operation to query the endpoint group ID of a basic Alibaba Cloud Global Accelerator (GA) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The ID of the endpoint that you want to delete from the basic Alibaba Cloud Global Accelerator (GA) instance.
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

func (s DeleteBasicEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteBasicEndpointRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DeleteBasicEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteBasicEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBasicEndpointRequest) SetClientToken(v string) *DeleteBasicEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetEndpointGroupId(v string) *DeleteBasicEndpointRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetEndpointId(v string) *DeleteBasicEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicEndpointRequest) SetRegionId(v string) *DeleteBasicEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBasicEndpointRequest) Validate() error {
	return dara.Validate(s)
}
