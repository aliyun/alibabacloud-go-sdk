// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAccelerateIpEndpointRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody
	GetAcceleratorId() *string
	SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody
	GetRequestId() *string
	SetState(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody
	GetState() *string
}

type DeleteBasicAccelerateIpEndpointRelationResponseBody struct {
	// The ID of the accelerated IP address of the basic GA instance.
	//
	// example:
	//
	// gaip-bp1****
	AccelerateIpId *string `json:"AccelerateIpId,omitempty" xml:"AccelerateIpId,omitempty"`
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
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the mapping between the accelerated IP address and endpoint.
	//
	// >  This parameter is not in use and empty.
	//
	// example:
	//
	// null
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DeleteBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) GetState() *string {
	return s.State
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetAcceleratorId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *DeleteBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

func (s *DeleteBasicAccelerateIpEndpointRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
