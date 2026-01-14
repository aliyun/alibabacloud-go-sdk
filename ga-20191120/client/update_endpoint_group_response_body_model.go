// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEndpointGroupResponseBody
	GetRequestId() *string
}

type UpdateEndpointGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEndpointGroupResponseBody) SetRequestId(v string) *UpdateEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
