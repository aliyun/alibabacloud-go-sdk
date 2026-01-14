// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEndpointGroupAttributeResponseBody
	GetRequestId() *string
}

type UpdateEndpointGroupAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEndpointGroupAttributeResponseBody) SetRequestId(v string) *UpdateEndpointGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEndpointGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
