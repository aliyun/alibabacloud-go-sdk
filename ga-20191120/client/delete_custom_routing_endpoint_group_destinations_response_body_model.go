// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointGroupDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomRoutingEndpointGroupDestinationsResponseBody
	GetRequestId() *string
}

type DeleteCustomRoutingEndpointGroupDestinationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) Validate() error {
	return dara.Validate(s)
}
