// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationIds(v []*string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody
	GetDestinationIds() []*string
	SetRequestId(v string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody
	GetRequestId() *string
}

type CreateCustomRoutingEndpointGroupDestinationsResponseBody struct {
	// The IDs of the endpoint group mappings.
	DestinationIds []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	// The IDs of the endpoint group mappings.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) GetDestinationIds() []*string {
	return s.DestinationIds
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) SetDestinationIds(v []*string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.DestinationIds = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponseBody) Validate() error {
	return dara.Validate(s)
}
