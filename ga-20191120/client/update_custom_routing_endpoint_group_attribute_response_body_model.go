// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomRoutingEndpointGroupAttributeResponseBody
	GetRequestId() *string
}

type UpdateCustomRoutingEndpointGroupAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
