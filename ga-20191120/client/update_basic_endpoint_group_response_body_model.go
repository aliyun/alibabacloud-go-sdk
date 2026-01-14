// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBasicEndpointGroupResponseBody
	GetRequestId() *string
}

type UpdateBasicEndpointGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicEndpointGroupResponseBody) SetRequestId(v string) *UpdateBasicEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicEndpointGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
