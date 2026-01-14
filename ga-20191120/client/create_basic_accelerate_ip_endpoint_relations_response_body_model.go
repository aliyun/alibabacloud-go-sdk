// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpEndpointRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody
	GetAcceleratorId() *string
	SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody
	GetRequestId() *string
}

type CreateBasicAccelerateIpEndpointRelationsResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicAccelerateIpEndpointRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpEndpointRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) SetAcceleratorId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) SetRequestId(v string) *CreateBasicAccelerateIpEndpointRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicAccelerateIpEndpointRelationsResponseBody) Validate() error {
	return dara.Validate(s)
}
