// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpEndpointRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody
	GetAccelerateIpId() *string
	SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody
	GetAcceleratorId() *string
	SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody
	GetEndpointId() *string
	SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody
	GetRequestId() *string
	SetState(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody
	GetState() *string
}

type CreateBasicAccelerateIpEndpointRelationResponseBody struct {
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
	// The status of the mapping between the accelerated IP address and the endpoint.
	//
	// >  This parameter is not in use.
	//
	// example:
	//
	// null
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) GetAccelerateIpId() *string {
	return s.AccelerateIpId
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) GetState() *string {
	return s.State
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetAccelerateIpId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.AccelerateIpId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetEndpointId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.EndpointId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) SetState(v string) *CreateBasicAccelerateIpEndpointRelationResponseBody {
	s.State = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
