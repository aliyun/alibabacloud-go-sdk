// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupDestinationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomRoutingEndpointGroupDestinationsResponseBody
	GetRequestId() *string
}

type UpdateCustomRoutingEndpointGroupDestinationsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointGroupDestinationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) Validate() error {
	return dara.Validate(s)
}
